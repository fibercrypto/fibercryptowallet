import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls.Material 2.12
import QtQuick.Controls 2.12

// Resource imports
// import "qrc:/ui/src/ui/"
import "../" // For quick UI development, switch back to resources when making a release

Dialog {
    id: deviceInteractionDialog
    title: qsTr("Device administration")
    standardButtons: Dialog.Ok

    Flickable {
        id: flickable
        anchors.fill: parent
        contentHeight: columnLayoutRoot.height
        clip: true
        ColumnLayout {
            id: columnLayoutRoot
            spacing: 30
            Image {
                id: icon
                source: "qrc:/images/resources/images/icons/backup.svg"
                sourceSize: "64x64"
                Layout.alignment: Qt.AlignTop
            }
            
            ColumnLayout {
                width: parent.width
                spacing: 30
                ColumnLayout {
                    ItemDelegate {
                        text: qsTr("Change device name")
                        Layout.fillWidth: true
                    }
                }
                ColumnLayout {
                    Label {
                        text: qsTr("--------------- Danger zone ---------------")
                        Material.foreground: Material.Pink
                    }
                    ItemDelegate {
                        text: qsTr(  "Upload new firmware")
                        Material.foreground: Material.Pink
                        Layout.fillWidth: true
                    }
                    ItemDelegate {
                        text: qsTr("Wipe device")
                        Layout.fillWidth: true
                        Material.foreground: Material.Pink
                    }
                }
            }
        } // ColumnLayout (root)

        ScrollIndicator.vertical: ScrollIndicator {
            parent: deviceInteractionDialog.contentItem
            anchors.top: flickable.top
            anchors.bottom: flickable.bottom
            anchors.right: parent.right
            anchors.rightMargin: -deviceInteractionDialog.rightPadding + 1
        }
    } // Flickable
}
