import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/Dialogs"
// import "qrc:/ui/src/ui/Controls"
import "Dialogs/" // For quick UI development, switch back to resources when making a release
import "Controls/" // For quick UI development, switch back to resources when making a release

Page {
    id: settings

    // BUG: About the wallet path: What happens on Windows?
    // TODO: Consider using `StandardPaths.standardLocations(StandardPaths.AppDataLocation)`

    // These are defaults. Will be restored when the "DEFAULT" button is clicked
    // TODO: How to get the defaults from the config manager
    readonly property string defaultWalletPath: configManager.getValue("skycoin/walletSource/1/Source")
    readonly property bool defaultIsRemoteWalletEnv: configManager.getValue("skycoin/walletSource/1/SourceType") === "remote"
    readonly property string defaultNodeUrl: configManager.getValue("skycoin/node/address")

    // These are the saved settings, must be applied when the settings are opened or when
    // the user clicks "RESET" and updated when the user clicks "APPLY"
    // TODO: This should be binded to backend properties
    property string savedWalletPath: configManager.getValue("skycoin/walletSource/1/Source")
    property bool savedIsRemoteWalletEnv: configManager.getValue("skycoin/walletSource/1/SourceType") === "remote"
    property url savedNodeUrl: configManager.getValue("skycoin/node/address")

    // These are the properties that are actually set, so they are aliases of the respective
    // control's properties
    property alias walletPath: textFieldWalletPath.text
    property alias isRemoteWalletEnv: switchRemoteWalletEnv.checked
    property alias nodeUrl: textFieldNodeUrl.text

    Component.onCompleted: {
        walletPath = savedWalletPath
        nodeUrl = savedNodeUrl
        isRemoteWalletEnv = savedIsRemoteWalletEnv
        console.log("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
        console.log("DEFAULT: " + defaultWalletPath + '\n' + defaultIsRemoteWalletEnv + '\n' + defaultNodeUrl)
        console.log("SAVED: " + savedWalletPath + '\n' + savedIsRemoteWalletEnv + '\n' + savedNodeUrl)
    }

    function saveCurrentSettings() {
        configManager.setValue("skycoin/walletSource/1/Source", walletPath)
        configManager.setValue("skycoin/walletSource/1/SourceType", isRemoteWalletEnv ? "remote" : "local")
        configManager.setValue("skycoin/node/address", nodeUrl)

        getCurrentSettings()
    }

    function getCurrentSettings() {
        walletPath = savedWalletPath = configManager.getValue("skycoin/walletSource/1/Source")
        isRemoteWalletEnv = savedIsRemoteWalletEnv = configManager.getValue("skycoin/walletSource/1/SourceType") === "remote"
        nodeUrl = configManager.getValue("skycoin/node/address")
    }

    footer: DialogButtonBox {
        standardButtons: Dialog.Apply | Dialog.Discard | Dialog.RestoreDefaults

        onApplied: {
            saveCurrentSettings()
        }

        onDiscarded: {
            dialogConfirmation.onlyDiscard = true
            dialogConfirmation.open()
        }

        onReset: {
            dialogConfirmation.onlyDiscard = false
            dialogConfirmation.open()
        }
    }

    ColumnLayout {
        anchors { top: parent.top; left: parent.left; right: parent.right; margins: 20 }

        spacing: 20

        GroupBox {
            Layout.fillWidth: true
            title: qsTr("Wallet settings")

            ColumnLayout {
                anchors.fill: parent

                TextField {
                    id: textFieldWalletPath

                    Layout.fillWidth: true
                    enabled: !isRemoteWalletEnv
                    selectByMouse: true
                    placeholderText: qsTr("Local wallet path")
                }

                RowLayout {
                    Layout.fillWidth: true

                    SwitchDelegate {
                        id: switchRemoteWalletEnv

                        property color textColor: checked ? Material.accent : Material.primaryTextColor
                        Behavior on textColor { ColorAnimation {} }

                        Layout.fillWidth: true

                        text: qsTr("Local/Remote wallet environment")
                        checked: savedIsRemoteWalletEnv
                        font.bold: true
                        Material.foreground: textColor
                    }
                }
            } // ColumnLayout
        } // GroupBox (wallet settings)

        GroupBox {
            Layout.fillWidth: true
            title: qsTr("Network settings")

            ColumnLayout {
                anchors.fill: parent

                TextField {
                    id: textFieldNodeUrl

                    selectByMouse: true
                    Layout.fillWidth: true
                    placeholderText: qsTr("Node URL")
                }
            } // ColumnLayout
        } // GroupBox (network settings)
    } // ColumnLayout

    // Confirm the discard or reset action:
    Dialog {
        id: dialogConfirmation

        property bool onlyDiscard: true

        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 300 ? 300 - 40 : applicationWindow.width - 40

        standardButtons: Dialog.Yes | Dialog.No
        title: qsTr("Confirm action")
        modal: true
        focus: true

        ColumnLayout {
            width: parent.width

            Label {
                id: labelQuestion

                Layout.fillWidth: true
                text: (dialogConfirmation.onlyDiscard ? qsTr("Discard all changes?") : qsTr("Restore defaults?"))
                font.bold: true
            }
            Label {
                id: labelDescription

                Layout.fillWidth: true
                text: qsTr("This action will set the settings to the") + " " + (dialogConfirmation.onlyDiscard ? qsTr("last saved values.") : qsTr("default values."))
                font.italic: true
                wrapMode: Text.Wrap
            }
        }

        onAccepted: {
            if (onlyDiscard) {
                getCurrentSettings()
            } else {
                getCurrentSettings() // getDefaultSettings()
            }
            saveCurrentSettings()
        }
    } // Dialog
}