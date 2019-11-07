import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import Transactions 1.0

// Resource imports
// import "qrc:/ui/src/ui/Dialogs"
import "Dialogs/" // For quick UI development, switch back to resources when making a release

Page {
    id: root

    property alias advancedMode: switchAdvancedMode.checked
    property string walletSelected
    property string signerSelected
    property string walletEncrypted
    property string destinationAddress
    property string amount
    property QTransaction txn

    footer: ToolBar {
        id: tabBarSend
        Material.theme: applicationWindow.Material.theme
        Material.primary: applicationWindow.accentColor
        Material.foreground: applicationWindow.Material.background
        Material.elevation: 0

        ToolButton {
            id: buttonAddWallet
            anchors.fill: parent
            text: qsTr("Send")
            icon.source: "qrc:/images/resources/images/icons/send.svg"

            onClicked: {

                var isEncrypted
                var walletSelected
                if (advancedMode){
                    var outs = stackView.currentItem.advancedPage.getSelectedOutputs()
                    var addrs = stackView.currentItem.advancedPage.getSelectedAddresses()
                    walletSelected = stackView.currentItem.advancedPage.getSelectedWallet()[0]
                    var destinationSummary = stackView.currentItem.advancedPage.getDestinationsSummary()
                    var changeAddress = stackView.currentItem.advancedPage.getChangeAddress()
                    var automaticCoinHours = stackView.currentItem.advancedPage.getAutomaticCoinHours()
                    var burnFactor = stackView.currentItem.advancedPage.getBurnFactor()
                    if (outs.length > 0){
                        console.log(outs)
                        txn = walletManager.sendFromOutputs(walletSelected, outs, destinationSummary[0], destinationSummary[1], destinationSummary[2], changeAddress, automaticCoinHours, burnFactor)
                    } else {
                        if (addrs.length == 0){
                            addrs = stackView.currentItem.advancedPage.getAllAddresses()
                            
                        }
                        txn = walletManager.sendFromAddresses(walletSelected, addrs, destinationSummary[0], destinationSummary[1], destinationSummary[2], changeAddress, automaticCoinHours, burnFactor)
                    } 
                    
                    isEncrypted = stackView.currentItem.advancedPage.walletIsEncrypted()[0]
                } else{
                    walletSelected = stackView.currentItem.simplePage.getSelectedWallet()
                    isEncrypted = stackView.currentItem.simplePage.walletIsEncrypted()
                    txn = walletManager.sendTo(walletSelected, stackView.currentItem.simplePage.getDestinationAddress(), stackView.currentItem.simplePage.getAmount())
                }
                console.log("HT "+txn.hoursTraspassed)
                dialogSendTransaction.showPasswordField =  isEncrypted// get if the current wallet is encrypted
                //dialogSendTransaction.previewDate = "2019-02-26 15:27"               
                dialogSendTransaction.previewType = TransactionDetails.Type.Send
                dialogSendTransaction.previewAmount = txn.amount
                dialogSendTransaction.previewHoursReceived = txn.hoursTraspassed
                dialogSendTransaction.previewHoursBurned = txn.hoursBurned
                dialogSendTransaction.previewtransactionID = txn.transactionId
                dialogSendTransaction.inputs = txn.inputs
                dialogSendTransaction.outputs = txn.outputs
                dialogSendTransaction.wallet = walletSelected
                dialogSendTransaction.open()
                
                
            }
        }
    }

    GroupBox {
        id: groupBox

        readonly property real margins: 50

        anchors.fill: parent
        anchors.topMargin: 10
        anchors.leftMargin: margins
        anchors.rightMargin: margins
        anchors.bottomMargin: margins
        implicitWidth: stackView.width
        clip: true

        label: SwitchDelegate {
            id: switchAdvancedMode

            text: qsTr("Advanced mode")

            onToggled: {
                if (checked) {
                    stackView.push(componentAdvanced)
                } else {
                    stackView.pop()
                }
            }
        }

        StackView {
            id: stackView

            property string walletSelected
            property string destinationAddress
            property string amount
            
            anchors.fill: parent
            initialItem: componentSimple
            clip: true            

            Component {
                id: componentSimple
                
                ScrollView {
                    id: scrollViewSimple
                    property alias simplePage: simple
                    contentWidth: simple.width
                    contentHeight: simple.height
                    clip: true

                    SubPageSendSimple {
                        id: simple

                        width: stackView.width
                        onWalletSelectedChanged: {
                            root.walletSelected = walletSelected
                        }
                        onAmountChanged: {
                            root.amount = amount
                        }
                        onDestinationAddressChanged: {
                            root.destinationAddress = destinationAddress
                        }
						onWalletEncryptedChanged: {
							root.walletEncrypted = walletEncrypted
						}
                    }
                }
            }

            Component {
                id: componentAdvanced

                ScrollView {
                    id: scrollViewAdvanced
                    property alias advancedPage: advanced
                    contentWidth: advanced.width
                    contentHeight: advanced.height
                    clip: true

                    SubPageSendAdvanced {
                        id: advanced
                        width: stackView.width
                    }
                }
            }
        } // StackView
    } // GroupBox

    DialogSendTransaction {
        id: dialogSendTransaction
        anchors.centerIn: Overlay.overlay
        property string wallet
        readonly property real maxHeight: (expanded ? 490 : 340) + (showPasswordField ? 140 : 0)
        width: applicationWindow.width > 640 ? 640 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > maxHeight ? maxHeight - 40 : applicationWindow.height - 40
        Behavior on height { NumberAnimation { duration: 1000; easing.type: Easing.OutQuint } }
        
        modal: true
        focus: true
		onAccepted: {
            signerSelected = stackView.currentItem.simplePage.getSignerSelected()
			var signedTxn = walletManager.signTxn(wallet, signerSelected, dialogSendTransaction.passwordText, [], txn)
            // FIXME check return error from walletManager.signTxn
			var injected = walletManager.broadcastTxn(signedTxn)
		}
    }
}
