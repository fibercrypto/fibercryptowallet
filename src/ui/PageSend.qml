import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/Dialogs"
import "Dialogs/" // For quick UI development, switch back to resources when making a release

Page {
    id: root

    property alias advancedMode: switchAdvancedMode.checked
    property string walletSelected
    property string walletEncrypted
    property string destinationAddress
    property string amount

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
                if (advancedMode){
                    var outs = stackView.currentItem.advancedPage.getSelectedOutputs()
                    var addrs = stackView.currentItem.advancedPage.getSelectedAddresses()
                    var wlt = stackView.currentItem.advancedPage.getSelectedWallet()
                    if (outs.length > 0){
//                        walletManager.
                    } else if(addrs.length > 0){

                    } else{

                    }
                    console.log(stackView.currentItem.advancedPage.getSelectedAddresses())
                    console.log(stackView.currentItem.advancedPage.getSelectedWallet())
                } else{
                    walletManager.sendTo(stackView.walletSelected, stackView.destinationAddress, stackView.amount)
                }

                dialogSendTransaction.showPasswordField = true // get if the current wallet is encrypted
                dialogSendTransaction.previewDate = "2019-02-26 15:27"               
                dialogSendTransaction.previewType = TransactionDetails.Type.Receive
	            dialogSendTransaction.showPasswordField = walletEncrypted // get if the current wallet is encrypted

//                dialogSendTransaction.showPasswordField = true // get if the current wallet is encrypted
                dialogSendTransaction.previewDate = "2019-02-26 15:27"
                dialogSendTransaction.previewType = TransactionDetails.Type.Send
                dialogSendTransaction.previewAmount = 10
                dialogSendTransaction.previewHoursReceived = 16957
                dialogSendTransaction.previewHoursBurned = 33901
                dialogSendTransaction.previewtransactionID = "kq7wdkkUT736dnuyQasdhsaDJ9874jk"
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

            property string walletSelected: simple.walletSelected
            property string destinationAddress: simple.destinationAddress
            property string amount: simple.amount
            
            anchors.fill: parent
            initialItem: componentSimple
            clip: true            

            Component {
                id: componentSimple
                
                ScrollView {
                    id: scrollViewSimple
                    
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

        readonly property real maxHeight: (expanded ? 490 : 340) + (showPasswordField ? 140 : 0)
        width: applicationWindow.width > 640 ? 640 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > maxHeight ? maxHeight - 40 : applicationWindow.height - 40
        Behavior on height { NumberAnimation { duration: 1000; easing.type: Easing.OutQuint } }
        
        modal: true
        focus: true
		onAccepted: {
		    walletManager.sendTo(walletSelected, destinationAddress, amount)
		}
    }
}
