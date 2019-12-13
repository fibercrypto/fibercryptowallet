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

    // BUG: About the wallet path: What happens on Windows?
    // TODO: Consider using `StandardPaths.standardLocations(StandardPaths.AppDataLocation)`

    // These are defaults. Will be restored when the "DEFAULT" button is clicked
    // TODO: How to get the defaults from the config manager
    readonly property string defaultWalletPath: configManager.getDefaultValue("skycoin/walletSource/1/Source")
    readonly property bool defaultIsLocalWalletEnv: configManager.getDefaultValue("skycoin/walletSource/1/SourceType") === "local"
    readonly property string defaultNodeUrl: configManager.getDefaultValue("skycoin/node/address")

    // These are the saved settings, must be applied when the settings are opened or when
    // the user clicks "RESET" and updated when the user clicks "APPLY"
    // TODO: This should be binded to backend properties
    property string savedWalletPath: configManager.getValue("skycoin/walletSource/1/Source")
    property bool savedIsLocalWalletEnv: configManager.getValue("skycoin/walletSource/1/SourceType") === "local"
    property url savedNodeUrl: configManager.getValue("skycoin/node/address")

    // These are the properties that are actually set, so they are aliases of the respective
    // control's properties
    property alias walletPath: textFieldWalletPath.text
    property alias isLocalWalletEnv: switchLocalWalletEnv.checked
    property alias nodeUrl: textFieldNodeUrl.text

    Component.onCompleted: {
        loadSavedSettings()
    }

    function saveCurrentSettings() {
        configManager.setValue("skycoin/walletSource/1/Source", walletPath)
        configManager.setValue("skycoin/walletSource/1/SourceType", isLocalWalletEnv ? "local" : "remote")
        configManager.setValue("skycoin/node/address", nodeUrl)
        loadSavedSettings()
    }

    function loadSavedSettings() {
        walletPath = savedWalletPath = configManager.getValue("skycoin/walletSource/1/Source")
        isLocalWalletEnv = savedIsLocalWalletEnv = configManager.getValue("skycoin/walletSource/1/SourceType") === "local"
        nodeUrl = savedNodeUrl = configManager.getValue("skycoin/node/address")

        updateFooterButtonsStatus()
    }

    function restoreDefaultSettings() {
        walletPath = defaultWalletPath
        isLocalWalletEnv = defaultIsLocalWalletEnv
        nodeUrl = defaultNodeUrl

        saveCurrentSettings()
    }

    function updateFooterButtonsStatus() {
        var configChanged = (walletPath !== savedWalletPath || isLocalWalletEnv !== savedIsLocalWalletEnv || nodeUrl != savedNodeUrl)
        var noDefaultConfig = (walletPath !== defaultWalletPath || isLocalWalletEnv !== defaultIsLocalWalletEnv || nodeUrl !== defaultNodeUrl)
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
         enabled:abm.hasInit()
         AddrsBookModel{
             id:abm
         }
                    Layout.fillWidth: true
                    title: qsTr("Address Book Settings")
                    ColumnLayout {
                        anchors.fill: parent
                        RowLayout{
                            Layout.fillWidth: true
                            Label { text: qsTr("Security type:"); font.bold: true }
                        }
                        RowLayout {
                        id: groupRadBtn
                            Layout.fillWidth: true
                                             RadioButton {
                                             property int pos:0
                                                 checked: abm.getSecType()==0
                                                 anchors.margins:10
                                                         Layout.fillWidth:true
                                                 text: qsTr("Low (Plain text)")
                                             }
                                             RadioButton {
                                             property int pos:1
                                             checked: abm.getSecType()==1
                                             anchors.margins: 10
                                                         Layout.fillWidth:true
                                                 text: qsTr("Medium (Recommended)")
                                             }
                                             RadioButton {
                                             property int pos:2
                                             checked: abm.getSecType()==2
                                                anchors.margins: 10
                                                Layout.fillWidth:true
                                                 text: qsTr("Hard (With password)\n"+
                                                 "(This can slow your dispositive)")
                                             }

                        }//RowLayoutRadioButtons
                        RowLayout {
                             Layout.fillWidth: true
                            Button{
                            id:changePassBtn
                                    enabled:abm.getSecType()==2
                                      text: qsTr("Change Password")
                                    highlighted: true
                                              anchors.margins: 10
                                              Layout.fillWidth:true
                                              onClicked: {
                                              getpass.open()
                                              }
                            }
                            Button{
                            id:applyChangesBtn
                                      text: qsTr("Apply Changes")
                                      enabled:false
                                        highlighted: true
                                                anchors.margins: 10
                                                Layout.fillWidth:true
                                                onClicked: {
                                                if(abm.getSecType()==2){
                                                getpass.open()
                                                }else{
                                                if(buttonsGroup.select==2){
                                                setpass.open()
                                                }else{
                                                this.enabled=!abm.changeSecType(buttonsGroup.select,"","")
                                                changePassBtn.enabled=false
                                                }
                                                }
                                                }
                                  }

                                  }//RowLayoutButtons

                    } // ColumnLayout
                    DialogGetPassword{
                    id:getpass
                    anchors.centerIn: Overlay.overlay
                    height:180
                    onAccepted:{
                    if(!abm.authenticate(getpass.password)){
                    getpass.open()
                    }else{
                    if (buttonsGroup.select==2){
                    setpass.open()
                    }else{
                    applyChangesBtn.enabled=!abm.changeSecType(buttonsGroup.select,getpass.password,"")
                    changePassBtn.enabled=false
                    }
                    }
                    }
                    }

                    DialogSetPassword{
                    id:setpass
                    anchors.centerIn: Overlay.overlay
                    onAccepted:{
                    applyChangesBtn.enabled=!abm.changeSecType(2,getpass.password,setpass.password)
                    changePassBtn.enabled=!applyChangesBtn.enabled
                    }
                    }

                } // GroupBox (addressBook setting)
    } // ColumnLayout

  ButtonGroup {
  property int select:checkedButton.pos
       id:buttonsGroup
        buttons: groupRadBtn.children
     onClicked:{
     applyChangesBtn.enabled=select!=abm.getSecType()
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
