import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import WalletsManager 1.0
import AddrsBookManager 1.0

// Resource imports
// import "qrc:/ui/src/ui/Controls"
import "Dialogs"
import "Controls" // For quick UI development, switch back to resources when making a release

Page {
    id: root
    property string walletSelected
    property string walletSelectedName
    property bool walletEncrypted: false 
    property string amount
    property string destinationAddress
    function getSelectedWallet(){
        return walletSelected
    }
    function getAmount(){
        return amount
    }
    function getDestinationAddress(){
        return destinationAddress
    }
    function walletIsEncrypted(){
        return [walletSelected, walletSelectedName, walletEncrypted]
    }
    signal qrCodeRequested(var data)

function getAddressList(){
addressList.clear()
for(var i=0;i<abm.count;i++){
for(var j=0;j<abm.contacts[i].address.rowCount();j++){
addressList.append({name:abm.contacts[i].name,
address:abm.contacts[i].address.address[j].value,
coinType:abm.contacts[i].address.address[j].coinType})
}
}
}


    onQrCodeRequested: {
        dialogQR.setVars(data)
        dialogQR.open()
    }

 AddrsBookModel{
    id:abm
    }

 DialogSelectAddressByAddressBook{
                            id: dialogSelectAddressByAddressBook

                            anchors.centerIn: Overlay.overlay
                            width: applicationWindow.width > 540 ? 540 - 40 : applicationWindow.width - 40
                            height: applicationWindow.height - 40

                            listAddrsModel: addressList

                            focus: true
                            modal: true

onAboutToShow:{
getAddressList()
}

         onAccepted: {
                textFieldWalletsSendTo.text = selectedAddress
                      }
                }

                 DialogGetPassword{
                 id:getpass
                 anchors.centerIn: Overlay.overlay
                 height:180
                 onAccepted:{
                 if(!abm.authenticate(getpass.password)){
                 getpass.open()
                 }else{
                 abm.loadContacts()
                 dialogSelectAddressByAddressBook.open()
                 }
                 }
                 }

    ColumnLayout {
        id: columnLayoutRoot
        anchors.fill: parent
        anchors.leftMargin: 10
        anchors.rightMargin: 10
        spacing: 20

        ColumnLayout {
            id: columnLayoutSendFrom

            Layout.alignment: Qt.AlignTop

            Label { text: qsTr("Send from") }

            ComboBox {
                id: comboBoxWalletsSendFrom

                Layout.fillWidth: true
                textRole: "name"
                displayText: comboBoxWalletsSendFrom.model.wallets[comboBoxWalletsSendFrom.currentIndex].sky ? comboBoxWalletsSendFrom.model.wallets[comboBoxWalletsSendFrom.currentIndex].name + " - " + comboBoxWalletsSendFrom.model.wallets[comboBoxWalletsSendFrom.currentIndex].sky + " SKY (" + comboBoxWalletsSendFrom.model.wallets[comboBoxWalletsSendFrom.currentIndex].coinHours + " CoinHours)": "Select a wallet"

                model: WalletModel {
                    Component.onCompleted: {
                        loadModel(walletManager.getWallets())
                    }
                } 
                
                // Taken from Qt 5.13.0 source code:
                delegate: MenuItem {
                    width: parent.width
                    text: comboBoxWalletsSendFrom.textRole ? (Array.isArray(comboBoxWalletsSendFrom.model) ? modelData[comboBoxWalletsSendFrom.textRole] : model[comboBoxWalletsSendFrom.textRole] + " - "+model["sky"] + " SKY (" + model["coinHours"] + " CoinHours)") : " --- " + modelData
                    Material.foreground: comboBoxWalletsSendFrom.currentIndex === index ? parent.Material.accent : parent.Material.foreground
                    highlighted: comboBoxWalletsSendFrom.highlightedIndex === index
                    hoverEnabled: comboBoxWalletsSendFrom.hoverEnabled
                    leftPadding: highlighted ? 2*padding : padding // added
                    Behavior on leftPadding { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } } // added
                }

                onActivated: {
                    root.walletSelected = comboBoxWalletsSendFrom.model.wallets[comboBoxWalletsSendFrom.currentIndex].fileName
                    root.walletSelectedName = comboBoxWalletsSendFrom.model.wallets[comboBoxWalletsSendFrom.currentIndex].name
                    root.walletEncrypted = comboBoxWalletsSendFrom.model.wallets[comboBoxWalletsSendFrom.currentIndex].encryptionEnabled
                }
            } // ComboBox
        } // ColumnLayout (send from)

        ColumnLayout {
            id: columnLayoutSendTo

            Layout.alignment: Qt.AlignTop

            Label { text: qsTr("Send to") }

             Button {
                                id: buttonSelectCustomChangeAddress
                                text: qsTr("Select")
                                flat: true
                                highlighted: true

                                onClicked: {
                                 if(abm.getSecType()!=2){
                                        abm.loadContacts()
                                        dialogSelectAddressByAddressBook.open()
                                     }else{
                                     getpass.open()
                                   }
                                }
                            }

            RowLayout {
                Layout.fillWidth: true
                spacing: 8

                ToolButtonQR {
                    id: toolButtonQR
                    Layout.bottomMargin: 4

                    iconSize: "24x24"

                    onClicked: {
                        qrCodeRequested(textFieldWalletsSendTo.text)
                    }
                }
                
                TextField {
                    id: textFieldWalletsSendTo
                    font.family: "Code New Roman"
                    placeholderText: qsTr("Destination address")
                    selectByMouse: true
                    Layout.fillWidth: true
                    Layout.topMargin: -5
                    Material.accent: abm.addressIsValid(text) ? parent.Material.accent : Material.color(Material.Red)
                    onTextChanged:{
                        root.destinationAddress = text

                    }
                }
            } // RowLayout
        } // ColumnLayout (send to)

        TextField {
            id: textFieldAmount
            placeholderText: qsTr("Amount to send")
            selectByMouse: true
            Layout.fillWidth: true
            Layout.topMargin: -10
            validator: DoubleValidator {
                notation: DoubleValidator.StandardNotation
            }
            onTextChanged:{
                    root.amount = text
                }
        }
    } // ColumnLayout (root)

    ListModel{
    id:addressList
    }
}
