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

    property bool configChanged: false

    signal canceled()

    function updateFooterButtonsStatus() {
        footer.standardButton(Dialog.Apply).enabled = configChanged
    }

    footer: DialogButtonBox {
        standardButtons: Dialog.Apply | Dialog.Cancel

        onApplied: {
            saveCurrentSettings()
        }

        onRejected: {
            canceled()
        }
    }

    ScrollView {
        id: scrollView
        anchors.fill: parent
        contentWidth: width
        
        ColumnLayout {
            anchors.fill: parent

            Label { text: qsTr("Security type:"); font.bold: true }

            ListView {
                id: listViewSecurityType

                Layout.fillWidth: true
                Layout.topMargin: -20
                height: contentHeight

                spacing: -6
                interactive: false
                model: [ qsTr("Low (Plain text)"), qsTr("Medium (Recommended)"), qsTr("Hard (with password)") ]
                delegate: RadioButton {
                    width: parent.width
                    text: modelData

                    onCheckedChanged: {
                        if (checked) {
                            ListView.view.currentIndex = index
                        }
                    }

                    Button {
                        id: buttonChangePassword

                        anchors.verticalCenter: parent.verticalCenter
                        anchors.right: parent.right

                        visible: index === Settings.SecurityType.StrongSecurity
                        enabled: visible && addrsBookModel.getSecType() === Settings.SecurityType.StrongSecurity
                        text: qsTr("Change password")
                        highlighted: true

                        onClicked: {
                            dialogGetPassword.open()
                        }
                    }
                } // RadioButton (delegate)
            } // ListView (log output)
            
            /*
            RowLayout {
                id: rowLayoutSecurityType
                Layout.fillWidth: true
                
                RadioButton {
                    property int pos: 0
                    Layout.fillWidth: true
                    checked: addrsBookModel.getSecType() === 0
                    text: qsTr("Low (Plain text)")
                }
                RadioButton {
                    property int pos: 1
                    Layout.fillWidth: true
                    checked: addrsBookModel.getSecType() === 1
                    text: qsTr("Medium (Recommended)")
                }
                RadioButton {
                    property int pos: 2
                    Layout.fillWidth: true
                    checked: addrsBookModel.getSecType() === 2
                    text: qsTr("Hard (with password)")// +
                        //qsTr("(This can slowdown your device)")
                }
            } // RowLayout (radio buttons)
            */

            RowLayout {
                Layout.fillWidth: true

                Button {
                    id: buttonChangePassword
                    
                    Layout.fillWidth: true
                    enabled:addrsBookModel.getSecType() === 2
                    text: qsTr("Change Password")
                    highlighted: true

                    onClicked: {
                        dialogGetPassword.open()
                    }
                }

                Button {
                    id: buttonApplyChanges

                    Layout.fillWidth: true

                    text: qsTr("Apply Changes")
                    enabled: false
                    highlighted: true

                    onClicked: {
                        if (addrsBookModel.getSecType() === 2) {
                            dialogGetPassword.open()
                        } else {
                            if (buttonGroupSecurityType.select === 2) {
                                dialogSetPassword.open()
                            } else {
                                this.enabled = !addrsBookModel.changeSecType(buttonGroupSecurityType.select, "", "")
                                buttonChangePassword.enabled = false
                            }
                        }
                    }
                }
            } // RowLayout (buttons)

        } // ColumnLayout

        DialogGetPassword {
            id: dialogGetPassword
            anchors.centerIn: Overlay.overlay
            height: 180

            onAccepted: {
                if(!addrsBookModel.authenticate(dialogGetPassword.password)) {
                    dialogGetPassword.open()
                } else {
                    if (buttonGroupSecurityType.select === 2) {
                        dialogSetPassword.open()
                    } else {
                        buttonApplyChanges.enabled = !addrsBookModel.changeSecType(buttonGroupSecurityType.select, dialogGetPassword.password, "")
                        buttonChangePassword.enabled = false
                    }
                }
            }
        }

        DialogSetPassword {
            id: dialogSetPassword
            anchors.centerIn: Overlay.overlay

            onAccepted: {
                buttonApplyChanges.enabled = !addrsBookModel.changeSecType(2, dialogGetPassword.password, dialogSetPassword.password)
                buttonChangePassword.enabled = !buttonApplyChanges.enabled
            }
        }
    } // ScrollView

    ButtonGroup {
        id: buttonGroupSecurityType
        property int select: checkedButton.pos

        buttons: rowLayoutSecurityType.children
        onClicked: {
            buttonApplyChanges.enabled = select !== addrsBookModel.getSecType()
        }
    }
}
