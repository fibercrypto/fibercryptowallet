import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import WalletsManager 1.0

// Resource imports
import "../" // For quick UI development, switch back to resources when making a release

Item {
    id: root

    readonly property bool itemVisible: index === 0 || addressSky > 0 || emptyAddressVisible
    property bool showOnlyAddresses: false

    signal addAddressesRequested()
    signal editWalletRequested()
    signal qrCodeRequested(var data)

    visible: itemVisible || opacity > 0.0
    opacity: itemVisible ? 1.0 : 0.0

    Behavior on height { NumberAnimation { duration: 250; easing.type: Easing.OutQuint } }
    Behavior on opacity { NumberAnimation { duration: 200; easing.type: Easing.OutQuint } }

    RowLayout {
        id: delegateAddressMenuRowLayout
        anchors.fill: parent
        anchors.leftMargin: listWalletLeftMargin
        anchors.rightMargin: listWalletRightMargin
        spacing: listWalletSpacing
        visible: index === 0 && !showOnlyAddresses

        ToolButton {
            id: buttonAddAddress
            text: qsTr("Add address")
            icon.source: "qrc:/images/resources/images/icons/add.svg"
            Material.foreground: Material.Teal
            Layout.fillWidth: true

            onClicked: {
                addAddressesRequested()
            }
        }
        ToolButton {
            id: buttonToggleVisibility
            text: qsTr("Show empty")
            checkable: true
            checked: emptyAddressVisible
            icon.source: "qrc:/images/resources/images/icons/visible" + (checked ? "On" : "Off") + ".svg"
            Material.accent: Material.Indigo
            Material.foreground: Material.Grey
            Layout.fillWidth: true

            onCheckedChanged: {
                emptyAddressVisible = checked
            }
        }
        ToolButton {
            id: buttonToggleEncryption
            text: qsTr("Encrypt wallet")
            checkable: true
            checked: encryptionEnabled
            icon.source: "qrc:/images/resources/images/icons/lock" + (checked ? "On" : "Off") + ".svg"
            Material.accent: Material.Amber
            Material.foreground: Material.Grey
            Layout.fillWidth: true

            onCheckedChanged: {
                encryptionEnabled = checked
                dialogRequestPassword.open()
            }
        }
        ToolButton {
            id: buttonEdit
            text: qsTr("Edit wallet")
            icon.source: "qrc:/images/resources/images/icons/edit.svg"
            Material.foreground: Material.Blue
            Layout.fillWidth: true

            onClicked: {
                editWalletRequested()
            }
        }
    } // RowLayout (menu)

    RowLayout {
        id: delegateAddressRowLayout
        anchors.fill: parent
        anchors.leftMargin: listWalletLeftMargin
        anchors.rightMargin: listWalletRightMargin
        spacing: listWalletSpacing
        visible: root.visible && index > 0

        Label {
            id: labelNumber
            visible:  !showOnlyAddresses
            text: index
        }

        ToolButtonQR {
            id: toolButtonQR

            onClicked: {
                qrCodeRequested(address)
            }
        }

        RowLayout {
            TextInput {
                id: textAddress
                text: address // a role of the model
                readOnly: true
                font.family: "Code New Roman"
            }
            ToolButton {
                id: toolButtonCopy
                visible:  !showOnlyAddresses
                icon.source: "qrc:/images/resources/images/icons/copy.svg"
                Layout.alignment: Qt.AlignLeft
                ToolTip.text: qsTr("Copy to clipboard")
                ToolTip.visible: hovered // TODO: pressed when mobile?
                ToolTip.delay: Qt.styleHints.mousePressAndHoldInterval

                Image {
                    id: imageCopied
                    anchors.centerIn: parent
                    visible:  !showOnlyAddresses
                    source: "qrc:/images/resources/images/icons/check-simple.svg"
                    fillMode: Image.PreserveAspectFit
                    sourceSize: Qt.size(toolButtonCopy.icon.width*1.5, toolButtonCopy.icon.height*1.5)
                    z: 1

                    opacity: 0.0
                }

                onClicked: {
                    textAddress.selectAll()
                    textAddress.copy()
                    textAddress.deselect()
                    if (copyAnimation.running) {
                        copyAnimation.restart()
                    } else {
                        copyAnimation.start()
                    }
                }

                SequentialAnimation {
                    id: copyAnimation
                    NumberAnimation { target: imageCopied; property: "opacity"; to: 1.0; easing.type: Easing.OutCubic }
                    PauseAnimation { duration: 1000 }
                    NumberAnimation { target: imageCopied; property: "opacity"; to: 0.0; easing.type: Easing.OutCubic }
                }
            } // ToolButton
            Rectangle {
                id: spacer
                visible:  !showOnlyAddresses
                Layout.fillWidth: true
            }
        }

        Label {
            id: labelAddressSky
            visible:  !showOnlyAddresses
            text: addressSky // a role of the model
            color: Material.accent
            horizontalAlignment: Text.AlignRight
            Layout.preferredWidth: internalLabelsWidth
        }

        Label {
            id: labelAddressCoins
            visible:  !showOnlyAddresses
            text: addressCoinHours // a role of the model
            horizontalAlignment: Text.AlignRight
            Layout.preferredWidth: internalLabelsWidth
        }
    } // RowLayout (addresses)
}
