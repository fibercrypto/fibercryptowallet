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
    id: settings

    enum LogLevel { Debug, Information, Warning, Error, FatalError, Panic }
    enum LogOutput { Stdout, Stderr, None, File }

    // BUG: About the wallet path: What happens on Windows?
    // TODO: Consider using `StandardPaths.standardLocations(StandardPaths.AppDataLocation)`

    // These are defaults. Will be restored when the "DEFAULT" button is clicked
    // TODO: How to get the defaults from the config manager
    readonly property string defaultWalletPath: configManager.getDefaultValue("skycoin/walletSource/1/Source")
    readonly property bool defaultIsLocalWalletEnv: configManager.getDefaultValue("skycoin/walletSource/1/SourceType") === "local"
    readonly property string defaultNodeUrl: configManager.getDefaultValue("skycoin/node/address")
    readonly property int defaultLogLevel: configManager.getDefaultValue("skycoin/log/level")
    readonly property int defaultLogOutput: configManager.getDefaultValue("skycoin/log/output")
    readonly property string defaultLogOutputFile: configManager.getDefaultValue("skycoin/log/outputFile")
    readonly property var defaultCacheLifeTime: configManager.getDefaultValue("global/cache/lifeTime")

    // These are the saved settings, must be applied when the settings are opened or when
    // the user clicks "RESET" and updated when the user clicks "APPLY"
    // TODO: This should be binded to backend properties
    property string savedWalletPath: configManager.getValue("skycoin/walletSource/1/Source")
    property bool savedIsLocalWalletEnv: configManager.getValue("skycoin/walletSource/1/SourceType") === "local"
    property url savedNodeUrl: configManager.getValue("skycoin/node/address")
    property int savedLogLevel: configManager.getValue("skycoin/log/level")
    property int savedLogOutput: configManager.getValue("skycoin/log/output")
    property string savedLogOutputFile: configManager.getDefaultValue("skycoin/log/outputFile")
    property var savedLifeTime: configManager.getValue("global/cache/lifeTime")

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
    property alias logLevel: comboBoxLogLevel.currentIndex
    property alias logOutput: listViewLogOutput.currentIndex
    property alias logOutputFile: listViewLogOutput.outputFile
    property alias cacheLifeTime: textFieldCacheLifeTime.text

    Component.onCompleted: {
        loadSavedSettings()
    }

    function saveCurrentSettings() {
        configManager.setValue("skycoin/walletSource/1/Source", walletPath)
        configManager.setValue("skycoin/walletSource/1/SourceType", isLocalWalletEnv ? "local" : "remote")
        configManager.setValue("skycoin/node/address", nodeUrl)
        configManager.setValue("skycoin/log/level", logLevel)
        configManager.setValue("skycoin/log/output", logOutput)
        configManager.setValue("skycoin/log/outputFile", logOutputFile)
        configManager.setValue("global/cache/lifeTime", cacheLifeTime)
        loadSavedSettings()
    }

    function loadSavedSettings() {
        walletPath = savedWalletPath = configManager.getValue("skycoin/walletSource/1/Source")
        isLocalWalletEnv = savedIsLocalWalletEnv = configManager.getValue("skycoin/walletSource/1/SourceType") === "local"
        nodeUrl = savedNodeUrl = configManager.getValue("skycoin/node/address")
        logLevel = savedLogLevel = configManager.getValue("skycoin/log/level")
        logOutput = savedLogOutput = configManager.getValue("skycoin/log/output")
        logOutputFile = savedLogOutputFile = configManager.getValue("skycoin/log/outputFile")
        cacheLifeTime = savedLifeTime = configManager.getValue("global/cache/lifeTime")
        console.log("cacheLifeTime", cacheLifeTime)

        updateFooterButtonsStatus()
    }

    function restoreDefaultSettings() {
        walletPath = defaultWalletPath
        isLocalWalletEnv = defaultIsLocalWalletEnv
        nodeUrl = defaultNodeUrl
        logLevel = defaultLogLevel
        logOutput = defaultLogOutput
        cacheLifeTime = defaultCacheLifeTime

        saveCurrentSettings()
    }

    function updateFooterButtonsStatus() {
        var configChanged = (walletPath !== savedWalletPath || isLocalWalletEnv !== savedIsLocalWalletEnv || nodeUrl != savedNodeUrl || logLevel != savedLogLevel || logOutput != savedLogOutput || logOutputFile != savedLogOutputFile || cacheLifeTime != savedLifeTime)
        var noDefaultConfig = (walletPath !== defaultWalletPath || isLocalWalletEnv !== defaultIsLocalWalletEnv || nodeUrl !== defaultNodeUrl || logLevel !== defaultLogLevel || logOutput !== defaultLogOutput || logOutputFile !== defaultLogOutputFile || cacheLifeTime !== defaultCacheLifeTime)
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

    ScrollView {
        id: scrollView
        anchors.fill: parent
        contentWidth: width

        ColumnLayout {
            id: columnLayout
            width: parent.width
            spacing: 20

            GroupBox {
                id: groupBoxWalletEnvironment
                Layout.fillWidth: true
                Layout.topMargin: 10
                Layout.leftMargin: 20
                Layout.rightMargin: 20
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
                id: groupBoxNetworkSettings

                Layout.fillWidth: true
                Layout.leftMargin: 20
                Layout.rightMargin: 20

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
                id: groupBoxGlobalSettings

                Layout.fillWidth: true
                Layout.leftMargin: 20
                Layout.rightMargin: 20

                title: qsTr("Global settings")

                ColumnLayout {
                    id: columnLayoutGlobalSettings
                    anchors.fill: parent

                    spacing: 20

                    TextField {
                        id: textFieldCacheLifeTime
                        Layout.fillWidth: true
                        selectByMouse: true
                        placeholderText: qsTr("Cache lifetime")
                        onTextChanged: {
                            updateFooterButtonsStatus();
                        }
                    }

                    Label { text: qsTr("Log level") }

                    ComboBox {
                        id: comboBoxLogLevel
                        Layout.fillWidth: true
                        Layout.topMargin: -20

                        readonly property var logLevelString: [ "debug", "info", "warn", "error", "fatal", "panic" ]
                        readonly property var logLevelColor: [ Material.Teal, Material.Blue, Material.Amber, Material.DeepOrange, Material.Red, Material.primaryTextColor ]

                        currentIndex: savedLogLevel < 0 || savedLogLevel >= count ? defaultLogLevel : savedLogLevel
                        model: [ qsTr("Debug"), qsTr("Informations"), qsTr("Warnings"), qsTr("Errors"), qsTr("Fatal errors"), qsTr("Panics") ]
                        delegate: MenuItem {
                            width: parent.width
                            text: comboBoxLogLevel.textRole ? (Array.isArray(comboBoxLogLevel.model) ? modelData[comboBoxLogLevel.textRole] : model[comboBoxLogLevel.textRole]) : modelData
                            icon.source: "qrc:/images/resources/images/icons/log_level_" + comboBoxLogLevel.logLevelString[index] + ".svg"
                            icon.color: Material.accent
                            Material.accent: comboBoxLogLevel.logLevelColor[index]
                            Material.foreground: comboBoxLogLevel.currentIndex === index ? parent.Material.accent : parent.Material.foreground
                            highlighted: comboBoxLogLevel.highlightedIndex === index
                            hoverEnabled: comboBoxLogLevel.hoverEnabled
                            leftPadding: highlighted ? 2*padding : padding // added
                            Behavior on leftPadding { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } } // added
                        } // MenuItem (delegate)
                    } // ComboBox

                    Label { text: qsTr("Log output") }

                    ListView {
                        id: listViewLogOutput

                        property alias outputFile: textFieldLogOutputFile.text
                        readonly property var logOutputString: [ "stdout", "stderr", "none", "file" ]

                        Layout.fillWidth: true
                        Layout.topMargin: -20
                        height: contentHeight

                        spacing: -6
                        interactive: false
                        model: [ qsTr("Standard output"), qsTr("Standard error output"), qsTr("None"), qsTr("File") ]
                        delegate: RadioButton {
                            width: index === Settings.LogOutput.File && textFieldLogOutputFile.enabled ? implicitWidth : parent.width
                            rightInset: textFieldLogOutputFile.enabled ? textFieldLogOutputFile.width : 0
                            text: modelData
                            checked: savedLogOutput < 0 || savedLogOutput >= ListView.view.count ? index === defaultLogLevel : index === savedLogOutput

                            onCheckedChanged: {
                                if (checked) {
                                    ListView.view.currentIndex = index
                                    if (index === Settings.LogOutput.File) {
                                        textFieldLogOutputFile.forceActiveFocus()
                                    }
                                }
                            }
                        } // RadioButton (delegate)

                        Component.onCompleted: {
                            textFieldLogOutputFile.anchors.leftMargin = listViewLogOutput.itemAtIndex(3).implicitWidth
                        }

                        TextField {
                            id: textFieldLogOutputFile

                            anchors.bottom: parent.bottom
                            anchors.bottomMargin: 6
                            anchors.left: parent.left
                            anchors.right: parent.right

                            enabled: listViewLogOutput.currentIndex === Settings.LogOutput.File
                            placeholderText: qsTr("Output file")
                            selectByMouse: true
                        }
                    } // ListView (log output)
                } // ColumnLayout (global settings)
            } // GroupBox (global settings)

            GroupBox {
                id: groupBoxAddressBookSettings

                Layout.fillWidth: true
                Layout.leftMargin: 20
                Layout.rightMargin: 20
                Layout.bottomMargin: 10 // The Last `GroupBox` must have this set

                title: qsTr("Address Book Settings")
                visible: addrsBookModel.hasInit()

                AddrsBookModel {
                    id: addrsBookModel
                }
                
                ColumnLayout {
                    anchors.fill: parent

                    Label { text: qsTr("Security type:"); font.bold: true }
                    
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
            } // GroupBox (addressBook setting)
        } // ColumnLayout
    } // ScrollView

    ButtonGroup {
        id: buttonGroupSecurityType
        property int select: checkedButton.pos

        buttons: rowLayoutSecurityType.children
        onClicked: {
            buttonApplyChanges.enabled = select !== addrsBookModel.getSecType()
        }
     }

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
