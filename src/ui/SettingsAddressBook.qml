import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import AddrsBookManager 1.0

// Resource imports
// import "qrc:/ui/src/ui/Dialogs"
// import "qrc:/ui/src/ui/Controls"
import "Dialogs/" // For quick UI development, switch back to resources when making a release
import "Controls/" // For quick UI development, switch back to resources when making a release

Page {
    id: settingsAddressBook

    enum SecurityType { LowSecurity, MediumSecurity, StrongSecurity }

    property bool enableButtonChangePassword: addrsBookModel.getSecType() === SettingsAddressBook.SecurityType.StrongSecurity

    signal canceled()

    footer: DialogButtonBox {
        id: footer
        standardButtons: Dialog.Apply | Dialog.Cancel

        onApplied: {
            if (addrsBookModel.getSecType() === SettingsAddressBook.SecurityType.StrongSecurity) {
                dialogGetPassword.open()
            } else {
                if (listViewSecurityType.currentIndex === SettingsAddressBook.SecurityType.StrongSecurity) {
                    dialogSetPassword.open()
                } else {
                    footer.standardButton(Dialog.Apply).enabled = !addrsBookModel.changeSecType(listViewSecurityType.currentIndex, "", "")
                    enableButtonChangePassword = false
                }
            }
        }

        onRejected: {
            canceled()
        }
    }

    ScrollView {
        id: scrollView
        anchors.fill: parent
        contentWidth: width
        clip: true
        
        ColumnLayout {
            anchors.fill: parent

            Label {
                Layout.topMargin: 10
                Layout.leftMargin: 20
                Layout.rightMargin: 20
                Layout.fillWidth: true

                text: qsTr("Security type")
                wrapMode: Text.Wrap
                font.bold: true
            }

            ListView {
                id: listViewSecurityType

                Layout.fillWidth: true
                Layout.leftMargin: 20
                Layout.rightMargin: 20
                height: contentHeight

                spacing: -6
                interactive: false
                currentIndex: addrsBookModel.getSecType()
                model: [ qsTr("Low (plain text)"), qsTr("Medium (recommended)"), qsTr("Hard (with password)") ]
                delegate: RadioButton {
                    width: parent.width
                    text: modelData
                    checked: index === ListView.view.currentIndex

                    onCheckedChanged: {
                        if (checked) {
                            ListView.view.currentIndex = index
                            footer.standardButton(Dialog.Apply).enabled = index !== addrsBookModel.getSecType()
                        }
                    }

                    Button {
                        id: buttonChangePassword

                        anchors.verticalCenter: parent.verticalCenter
                        anchors.right: parent.right

                        visible: index === SettingsAddressBook.SecurityType.StrongSecurity
                        enabled: visible && enableButtonChangePassword
                        text: qsTr("Change password")
                        highlighted: true

                        onClicked: {
                            dialogGetPassword.open()
                        }
                    }
                } // RadioButton (delegate)
            } // ListView (log output)

            Label {
                Layout.leftMargin: 20
                Layout.rightMargin: 20
                Layout.fillWidth: true

                visible: listViewSecurityType.currentIndex === DialogSelectSecType.SecurityType.StrongSecurity
                text: "<b>" + qsTr("Note:") + "</b> " + qsTr("Accessing a password-protected address book can slowdown your device")
                wrapMode: Text.Wrap
                Material.foreground: Material.Red
                font.italic: true
            }
        } // ColumnLayout
    } // ScrollView

    AddrsBookModel {
        id: addrsBookModel
    }

    DialogGetPassword {
        id: dialogGetPassword
        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 400 ? 400 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > 280 ? 280 - 40 : applicationWindow.height - 40

        modal: true
        focus: visible

        onAccepted: {
            if(!addrsBookModel.authenticate(dialogGetPassword.password)) {
                dialogGetPassword.open()
            } else {
                if (listViewSecurityType.currentIndex === SettingsAddressBook.SecurityType.StrongSecurity) {
                    dialogSetPassword.open()
                } else {
                    footer.standardButton(Dialog.Apply).enabled = !addrsBookModel.changeSecType(listViewSecurityType.currentIndex, dialogGetPassword.password, "")
                    enableButtonChangePassword = false
                }
            }
        }
    }

    DialogSetPassword {
        id: dialogSetPassword
        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 400 ? 400 - 40 : applicationWindow.width - 40

        modal: true
        focus: visible

        onAccepted: {
            footer.standardButton(Dialog.Apply).enabled = !addrsBookModel.changeSecType(SettingsAddressBook.SecurityType.StrongSecurity, dialogGetPassword.password, dialogSetPassword.password)
            enableButtonChangePassword = true
        }
    }
}
