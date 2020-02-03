import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls.Material 2.12
import QtQuick.Controls 2.12
import QtQuick.Dialogs 1.1

import DeviceInteraction 1.0
// Resource imports
// import "qrc:/ui/src/ui/"
import "../" // For quick UI development, switch back to resources when making a release

Dialog {
    id: deviceInteractionDialog
    title: qsTr("Device administration")
    standardButtons: Dialog.Ok
    DeviceInteraction {
        id: deviceInteraction3
        onBootModeDetermined: {
            uploadNewFirmware.visible = bootloader
            wipeDevice.visible = !bootloader
            changeDeviceName.visible = !bootloader
        }
    }
    onAboutToShow: deviceInteraction3.deviceFeatures()
    FileDialog {
        id: fileDialog
        title: "Please choose a file"
        folder: shortcuts.home
        selectMultiple: false
        nameFilters: [ "Firmware image files (*.bin)", "All files (*)" ]
        onAccepted: {
            deviceInteraction3.firmwareUpload(fileDialog.fileUrl)
        }
    }
    GetWordDialog {
        id: deviceNameDialog
        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 440 ? 440 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > 240 ? 240 - 40 : applicationWindow.height - 40
        focus: true
        modal: true
        
        title: qsTr('Please choose the new device name')
        label: qsTr("Device name")
        hint: 'aaaaaa'
        onAccepted: {
            deviceInteraction3.changeDeviceName(deviceNameDialog.selectedName)
        }
    }

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
                        id: changeDeviceName
                        text: qsTr("Change device name")
                        Layout.fillWidth: true
                        onClicked: deviceNameDialog.open()
                    }
                }
                ColumnLayout {
                    Label {
                        text: qsTr("--------------- Danger zone ---------------")
                        Material.foreground: Material.Pink
                        visible: wipeDevice.visible || uploadNewFirmware.visible
                    }
                    ItemDelegate {
                        id: uploadNewFirmware
                        text: qsTr(  "Upload new firmware")
                        Material.foreground: Material.Pink
                        Layout.fillWidth: true
                        onClicked: fileDialog.open()
                    }
                    ItemDelegate {
                        id: wipeDevice
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
