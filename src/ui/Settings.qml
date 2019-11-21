import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/Dialogs"
import "Dialogs/" // For quick UI development, switch back to resources when making a release

Page {
    id: settings

    // BUG: About the wallet path: What happens on Windows?
    // TODO: Consider using `StandardPaths.standardLocations(StandardPaths.AppDataLocation)`

    // These are defaults. Will be restored when the "DEFAULT" button is clicked
    // TODO: How to get the defaults from the config manager
    readonly property string defaultWalletPath: configManager.getValue("skycoin", ["walletSource"], "Source")
    readonly property string defaultNodeUrl: configManager.getValue("skycoin", ["node"], "node")
    readonly property bool defaultIsRemoteWalletEnv: configManager.getValue("skycoin", ["walletSource"], "SourceType")

    // These are the saved settings, must be applied when the settings are opened or when
    // the user clicks "RESET" and updated when the user clicks "APPLY"
    // TODO: This should be binded to backend properties
    property string savedWalletPath: configManager.getSourceString()
    property url savedNodeUrl: configManager.getNodeString()
    property bool savedIsRemoteWalletEnv: configManager.getTypeSource()

    // These are the properties that are actually set, so they are aliases of the respective
    // control's properties
    property alias walletPath: textFieldWalletPath.text
    property alias nodeUrl: textFieldNodeUrl.text
    property alias isRemoteWalletEnv: switchRemoteWalletEnv.checked

    Component.onCompleted: {
        walletPath = savedWalletPath
        nodeUrl = savedNodeUrl
        isRemoteWalletEnv = savedIsRemoteWalletEnv
        console.log("!!!!!!!!!!!!!!!!!!!!!!!!!!!!!!")
        console.log(defaultWalletPath + '\n' + defaultNodeUrl + '\n' + defaultIsRemoteWalletEnv)
    }

    footer: DialogButtonBox {
        standardButtons: Dialog.Apply | Dialog.Discard | Dialog.RestoreDefaults

        onApplied: {
            configManager.edit(nodeUrl, walletPath, isRemoteWalletEnv)
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

                RowLayout {
                    Layout.fillWidth: true

                    Label { text: qsTr("Wallet path:"); font.bold: true }
                    TextField {
                        id: textFieldWalletPath

                        selectByMouse: true
                        MouseArea {
                            anchors.fill: parent
                            cursorShape: Qt.IBeamCursor
                            acceptedButtons: Qt.NoButton
                        }
                        Layout.fillWidth: true
                        placeholderText: qsTr("Default to ") + savedWalletPath
                    }
                }
                RowLayout {
                    Layout.fillWidth: true

                    SwitchDelegate {
                        id: switchRemoteWalletEnv

                        property color textColor: checked ? Material.accent : Material.primaryTextColor
                        Behavior on textColor { ColorAnimation {} }

                        Layout.fillWidth: true

                        text: qsTr("Remote/Local wallet environment")
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

                RowLayout {
                    Layout.fillWidth: true

                    Label { text: qsTr("Node URL:"); font.bold: true }
                    TextField {
                        id: textFieldNodeUrl

                        selectByMouse: true
                        MouseArea {
                            anchors.fill: parent
                            cursorShape: Qt.IBeamCursor
                            acceptedButtons: Qt.NoButton
                        }
                        Layout.fillWidth: true
                        placeholderText: qsTr("Default to ") + savedNodeUrl.toString()
                    }
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
                text: qsTr("This action will set the settings to the") + " " + (dialogConfirmation.onlyDiscard ? qsTr("last saved values.") : qsTr("very default values."))
                font.italic: true
                wrapMode: Text.WordWrap
            }
        }

        onAccepted: {
            if (onlyDiscard) {
                walletPath = savedWalletPath
                nodeUrl = savedNodeUrl
                isRemoteWalletEnv = savedIsRemoteWalletEnv
            } else {
                walletPath = defaultWalletPath
                nodeUrl = defaultNodeUrl
                isRemoteWalletEnv = defaultIsRemoteWalletEnv
            }
        }
    } // Dialog
}