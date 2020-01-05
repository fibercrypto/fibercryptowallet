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
    readonly property string defaultWalletPath: configManager.getDefaultValue("skycoin/walletSource/1/Source")
    readonly property bool defaultIsLocalWalletEnv: configManager.getDefaultValue("skycoin/walletSource/1/SourceType") === "local"
    readonly property string defaultNodeUrl: configManager.getDefaultValue("skycoin/node/address")
    readonly property string defaultLogLevel: configManager.getDefaultValue("skycoin/node/level")

    // These are the saved settings, must be applied when the settings are opened or when
    // the user clicks "RESET" and updated when the user clicks "APPLY"
    // TODO: This should be binded to backend properties
    property string savedWalletPath: configManager.getValue("skycoin/walletSource/1/Source")
    property bool savedIsLocalWalletEnv: configManager.getValue("skycoin/walletSource/1/SourceType") === "local"
    property url savedNodeUrl: configManager.getValue("skycoin/node/address")
    property string savedLogLevel: configManager.getValue("skycoin/node/level")

    // QtObject{
    //     id: logLevel
    //     property string modifier
    //     property string old
    // }

    // These are the properties that are actually set, so they are aliases of the respective
    // control's properties
    property alias walletPath: textFieldWalletPath.text
    property alias isLocalWalletEnv: switchLocalWalletEnv.checked
    property alias nodeUrl: textFieldNodeUrl.text
    property alias logLevel: textFieldLogLevel.text

    Component.onCompleted: {
        loadSavedSettings()
    }

    function saveCurrentSettings() {
        configManager.setValue("skycoin/walletSource/1/Source", walletPath)
        configManager.setValue("skycoin/walletSource/1/SourceType", isLocalWalletEnv ? "local" : "remote")
        configManager.setValue("skycoin/node/address", nodeUrl)
        configManager.setValue("skycoin/node/level", logLevel)
        loadSavedSettings()
    }

    function loadSavedSettings() {
        walletPath = savedWalletPath = configManager.getValue("skycoin/walletSource/1/Source")
        isLocalWalletEnv = savedIsLocalWalletEnv = configManager.getValue("skycoin/walletSource/1/SourceType") === "local"
        nodeUrl = savedNodeUrl = configManager.getValue("skycoin/node/address")
        logLevel = savedLogLevel = configManager.getValue("skycoin/node/level")

        updateFooterButtonsStatus()
    }

    function restoreDefaultSettings() {
        walletPath = defaultWalletPath
        isLocalWalletEnv = defaultIsLocalWalletEnv
        nodeUrl = defaultNodeUrl
        logLevel = defaultLogLevel

        saveCurrentSettings()
    }

    function updateFooterButtonsStatus() {
        var configChanged = (walletPath !== savedWalletPath || isLocalWalletEnv !== savedIsLocalWalletEnv || nodeUrl != savedNodeUrl || logLevel != savedLogLevel)
        var noDefaultConfig = (walletPath !== defaultWalletPath || isLocalWalletEnv !== defaultIsLocalWalletEnv || nodeUrl !== defaultNodeUrl || logLevel !== defaultLogLevel)
        footer.standardButton(Dialog.Apply).enabled = configChanged
        footer.standardButton(Dialog.Discard).enabled = configChanged
        footer.standardButton(Dialog.RestoreDefaults).enabled = noDefaultConfig
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
            title: qsTr("Wallet environment settings")
            
            RowLayout {
                anchors.fill: parent
                
                Label {
                    text: qsTr("Remote")
                    font.bold: true
                    color: Material.hintTextColor
                }
                Switch {
                    id: switchLocalWalletEnv

                    checked: savedIsLocalWalletEnv
                    font.bold: true

                    onToggled: {
                        updateFooterButtonsStatus();
                    }
                }
                Label {
                    text: qsTr("Local")
                    font.bold: true
                    color: Material.accent
                }

                Rectangle {
                    Layout.fillHeight: true
                    Layout.leftMargin: 10
                    Layout.rightMargin: 10
                    width: 1
                    color: Material.hintTextColor
                }

                TextField {
                    id: textFieldWalletPath

                    Layout.fillWidth: true
                    enabled: isLocalWalletEnv
                    selectByMouse: true
                    placeholderText: qsTr("Local wallet path")

                    onTextChanged: {
                        updateFooterButtonsStatus();
                    }
                }
            } // RowLayout
        } // GroupBox (wallet settings)

        GroupBox {
            Layout.fillWidth: true
            title: qsTr("Network settings")

            TextField {
                id: textFieldNodeUrl

                anchors.fill: parent
                selectByMouse: true
                placeholderText: qsTr("Node URL")

                onTextChanged: {
                    updateFooterButtonsStatus();
                }
            }
        } // GroupBox (network settings)

        GroupBox {
            Layout.fillWidth: true
            title: qsTr("Log output")

            RowLayout {
                anchors.fill: parent

                RadioButton { //text
                    id: radioButtonlogOutputNone
                    checked: true
                    text: qsTr("None")
                }
                RadioButton {
                    id: radioButtonlogOutputStdOut
                    text: qsTr("stdout")
                }
                RadioButton {
                    id: radioButtonlogOutputFile
                    text: qsTr("File")
                }
            }

        }

        GroupBox{
            Layout.fillWidth: true
            title: qsTr("Log level")

            RowLayout{
                anchors.fill: parent

                TextField {
                    id: textFieldLogLevel

                    selectByMouse: true
                    placeholderText: qsTr("Log level")

                    onTextChanged: {
                        updateFooterButtonsStatus();
                    }
                }
            }
        }
        // RowLayout {
        //     Layout.fillWidth: true
            
        //     Label {
        //         text: qsTr("Log level")
        //         font.bold: true
        //         color: Material.accent
        //     }

            // ComboBox{
            //     id: logLevelOption

            //     // When an item is selected, update the backend.
            //     onActivated: logLevel.modifier = qsTr(logLevelOption.currentValue)
            //     // // Set the initial currentIndex to the value stored in the backend.
            //     // Component.onCompleted: currentIndex = indexOfValue(logLevel.old)
            //     model:["debug", "info", "warn", "error", "fatal", "panic"]
            // }
        //}
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
        focus: visible

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
                loadSavedSettings()
            } else {
                restoreDefaultSettings()
            }
        }
    } // Dialog
}
