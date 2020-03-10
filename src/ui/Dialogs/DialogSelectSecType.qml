import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "../Delegates" // For quick UI development, switch back to resources when making a release

Dialog {
    id: dialogCreateAddressBook

    enum SecurityType { LowSecurity, MediumSecurity, StrongSecurity }

    property alias select: listViewSecurityType.currentIndex

    standardButtons: Dialog.Ok | Dialog.Cancel

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

                text: qsTr("Select the security type for your <i>Address Book</i>")
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
                currentIndex: DialogSelectSecType.SecurityType.MediumSecurity
                model: [ qsTr("Low (plain text)"), qsTr("Medium (recommended)"), qsTr("Hard (with password)") ]
                delegate: RadioButton {
                    width: parent.width
                    text: modelData
                    checked: index === ListView.view.currentIndex

                    onCheckedChanged: {
                        if (checked) {
                            ListView.view.currentIndex = index
                        }
                    }
                } // RadioButton (delegate)
            } // ListView (log output)

            Label {
                Layout.leftMargin: 20
                Layout.rightMargin: 20
                Layout.fillWidth: true

                visible: select === DialogSelectSecType.SecurityType.StrongSecurity
                text: "<b>" + qsTr("Note:") + "</b> " + qsTr("Accessing a password-protected address book can slowdown your device")
                wrapMode: Text.Wrap
                Material.foreground: Material.Red
                font.italic: true
            }
        } // ColumnLayout
    } // ScrollView
}