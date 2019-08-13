import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import HistoryModels 1.0
import WalletsManager 1.0

// Resource imports
// import "qrc:/ui/src/ui/Dialogs"
// import "qrc:/ui/src/ui/Delegates"
import "Dialogs/" // For quick UI development, switch back to resources when making a release
import "Delegates/" // For quick UI development, switch back to resources when making a release

Page {
    id: root

    GroupBox {
        anchors.fill: parent
        anchors.margins: 20
        clip: true

        label: RowLayout {
            SwitchDelegate {
                id: switchFilters
                text: qsTr("Filters")
                onClicked:{
                    if (!checked) {
                        modelTransactions.clear()
                        modelTransactions.addMultipleTransactions(historyManager.loadHistory())
                    }
                    else {
                        modelTransactions.clear()
                        modelTransactions.addMultipleTransactions(historyManager.loadHistoryWithFilters())
                    }
                }
            }
            Button {
                id: buttonFilters
                flat: true
                enabled: switchFilters.checked
                highlighted: true
                text: qsTr("Select filters")

                onClicked: {
                    toolTipFilters.open()
                }
            }
        } // RowLayout (GroupBox label)

        ScrollView {
            anchors.fill: parent
            clip: true
            ListView {
                id: listTransactions

                model: modelTransactions
                delegate: HistoryListDelegate {
                    onClicked: {
                        dialogTransactionDetails.open()
                        listTransactions.currentIndex = index
                    }
                }
            }
        }
    } // GroupBox


    Dialog {
        id: toolTipFilters

        anchors.centerIn: Overlay.overlay

        readonly property real minimumHeight: Math.min(applicationWindow.height - 100, filter.contentHeight + 150)
        width: 300
        height: minimumHeight

        modal: true
        standardButtons: Dialog.Close
        closePolicy: Popup.CloseOnEscape | Popup.CloseOnPressOutside
        title: qsTr("Available filters")
        onClosed:{
            modelTransactions.clear()
            modelTransactions.addMultipleTransactions(historyManager.loadHistoryWithFilters())
            

        }
        onOpened:{
            //filter.myModel.loadModel(walletManager.getWallets())
            //filter.walletsLoaded()
            //
        }
        HistoryFilterList {
            id: filter
            anchors.fill: parent
           
            //property WalletModel myModel: WalletModel{
            //    id: modelFilters
            //}

        }
    } // ToolTip

    DialogTransactionDetails {
        id: dialogTransactionDetails
        anchors.centerIn: Overlay.overlay

        readonly property real maxHeight: expanded ? 590 : 370
        width: applicationWindow.width > 640 ? 640 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > maxHeight ? maxHeight - 40 : applicationWindow.height - 40
        Behavior on height { NumberAnimation { duration: 1000; easing.type: Easing.OutQuint } }

        modal: true
        focus: true

        date: listTransactions.currentItem.modelDate
        status: listTransactions.currentItem.modelStatus
        type: listTransactions.currentItem.modelType
        amount: listTransactions.currentItem.modelAmount
        hoursReceived: listTransactions.currentItem.modelHoursReceived
        hoursBurned: listTransactions.currentItem.modelHoursBurned
        transactionID: listTransactions.currentItem.modelTransactionID
        modelInputs: listTransactions.currentItem.modelInputs
        modelOutputs: listTransactions.currentItem.modelOutputs
    }

    QTransactionList {
        id: modelTransactions
    }

    HistoryManager{
        id:historyManager
    }

    Timer {
        
        id: qTraModelTimer
        repeat: false
        running: true
        interval: 0
        onTriggered: {
            
            modelTransactions.addMultipleTransactions(historyManager.loadHistory())
            qTraModelTimer.running = false
            
        }
    }
}
