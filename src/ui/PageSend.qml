import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import Transactions 1.0
import Utils 1.0

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
                var walletSelecteds
                if (advancedMode){
                    var outs = stackView.currentItem.advancedPage.getSelectedOutputsWithWallets()
                    var addrs = stackView.currentItem.advancedPage.getSelectedAddressesWithWallets()
                    //walletSelecteds = stackView.currentItem.advancedPage.getSelectedWallet()
                    var destinationSummary = stackView.currentItem.advancedPage.getDestinationsSummary()
                    var changeAddress = stackView.currentItem.advancedPage.getChangeAddress()
                    var automaticCoinHours = stackView.currentItem.advancedPage.getAutomaticCoinHours()
                    var burnFactor = stackView.currentItem.advancedPage.getBurnFactor()
                    if (outs[0].length > 0){
                        txn = walletManager.sendFromOutputs(outs[1], outs[0], destinationSummary[0], destinationSummary[1], destinationSummary[2], changeAddress, automaticCoinHours, burnFactor)
                    } else {
                        if (addrs[0].length == 0){
                            addrs = stackView.currentItem.advancedPage.getAllAddressesWithWallets()                            
                        }
                        txn = walletManager.sendFromAddresses(addrs[1], addrs[0], destinationSummary[0], destinationSummary[1], destinationSummary[2], changeAddress, automaticCoinHours, burnFactor)
                    } 
                    
                    isEncrypted = stackView.currentItem.advancedPage.walletIsEncrypted()
                } else{
                    walletSelected = stackView.currentItem.simplePage.getSelectedWallet()
                    isEncrypted = stackView.currentItem.simplePage.walletIsEncrypted()
                    var addrs = []
                    addrs.push([])
                    addrs.push([])
                    addrs[0].push(stackView.currentItem.simplePage.getDestinationAddress())
                    addrs[1].push(walletSelected)
                    txn = walletManager.sendTo(walletSelected, stackView.currentItem.simplePage.getDestinationAddress(), stackView.currentItem.simplePage.getAmount())
                }
                dialogSendTransaction.showPasswordField =  false//isEncrypted// get if the current wallet is encrypted
                //dialogSendTransaction.previewDate = "2019-02-26 15:27"               
                dialogSendTransaction.previewType = TransactionDetails.Type.Send
                dialogSendTransaction.previewAmount = txn.amount
                dialogSendTransaction.previewHoursReceived = txn.hoursTraspassed
                dialogSendTransaction.previewHoursBurned = txn.hoursBurned
                dialogSendTransaction.previewtransactionID = txn.transactionId
                dialogSendTransaction.inputs = txn.inputs
                dialogSendTransaction.outputs = txn.outputs
                dialogSendTransaction.walletsAddresses = addrs
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
        property var walletsAddresses
        readonly property real maxHeight: (expanded ? 490 : 340) + (showPasswordField ? 140 : 0)
        width: applicationWindow.width > 640 - 40 ? 640 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > maxHeight - 40 ? maxHeight - 40 : applicationWindow.height - 40
        Behavior on height { NumberAnimation { duration: 1000; easing.type: Easing.OutQuint } }
        
        modal: true
        focus: true
		onAccepted: {
            signerSelected = stackView.currentItem.simplePage.getSignerSelected()
            walletManager.signAndBroadcastTxnAsync(walletsAddresses[1], walletsAddresses[0],signerSelected, bridgeForPassword, [], txn)
        }
    }

    DialogGetPassword{
        id: getPasswordDialog
        anchors.centerIn: Overlay.overlay
        property int nAddress
        width: applicationWindow.width > 540 ? 540 - 120 : applicationWindow.width - 40
        height: applicationWindow.height > 570 ? 570 - 180 : applicationWindow.height - 40

        focus: true
        modal: true
        onClosed:{
            bridgeForPassword.setResult(getPasswordDialog.password)
            bridgeForPassword.unlock()
        }
    }

    QBridge{
        id: bridgeForPassword

        onGetPassword:{
            getPasswordDialog.title = message
            getPasswordDialog.clear()
            getPasswordDialog.open()
        }
    }
}
