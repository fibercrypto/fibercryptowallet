import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import QtGraphicalEffects 1.12

import "../Dialogs"
Item{
    id: root

    readonly property real delegateHeight: 80
    property alias inputsModel: inputList.model
    property alias outputsModel: outputList.model
    readonly property real finalViewHeight: delegateHeight + 400

    width: parent.width
    height: itemDelegateTransaction.height+inputList.height+ 30 + outputList.height+30

    ColumnLayout {
        id: delegateColumnLayout
        anchors.fill: parent

        ItemDelegate {
            id: itemDelegateTransaction
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignTop
            Component.onCompleted:{
                inputsModel = modelData.inputs
                outputsModel = modelData.outputs
            }

            RowLayout {
                id: delegateRowLayout
                anchors.fill: parent

                Label {
                    font.bold:true
                    text: qsTr("Transaction ID:") // a role of the model
                    Layout.fillWidth: true
                }

                Label {
                    text: modelData.transactionID
                    horizontalAlignment: Text.AlignLeft
                }
            } // RowLayout
            onClicked:{

        dialogTransactionDetails.date = modelData.date
        dialogTransactionDetails.status =  modelData.status
        dialogTransactionDetails.type =  modelData.type
        dialogTransactionDetails.hoursReceived = modelData.hoursTraspassed
        dialogTransactionDetails.hoursBurned = modelData.hoursBurned
        dialogTransactionDetails.transactionID = modelData.transactionID
        dialogTransactionDetails.modelInputs =  modelData.inputs
        dialogTransactionDetails.modelOutputs = modelData.outputs


                dialogTransactionDetails.open()
            }
        } // ItemDelegate
        RowLayout{
            Label{
                text:qsTr("Input")
                font.bold:true
            }
        }
        RowLayout{
            Rectangle {
                id: rectInput
                Layout.fillWidth: true
                height: 1
                color: "#DDDDDD"
            }
        }
        ListView {
            id: inputList
            implicitHeight: delegateHeight*(inputList.count) + 20
            clip: true
            interactive: false
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignTop

            Behavior on implicitHeight { NumberAnimation { duration: 250; easing.type: Easing.OutQuint } }
            Behavior on opacity { NumberAnimation { duration: expanded ? 250 : 1000; easing.type: Easing.OutQuint } }

            delegate: ItemDelegate {
                height: delegateHeight
                ColumnLayout{
                    Layout.fillWidth:true
                    RowLayout{
                        Label{
                            width:parent.width
                            text:address
                        }
                    }
                    RowLayout{
                        Label{
                            width:parent.width
                            font.bold:true
                            text:qsTr("Coins: ")
                        }
                        Label{
                            width:parent.width
                            Component.onCompleted:{
                                console.log("label width: "+this.width)
                            }
                            text:addressSky
                        }
                    }
                    RowLayout{
                        Label{
                            width:parent.width

                            font.bold:true
                            text:qsTr("CoinsHours: ")
                        }
                        Label{
                            width:parent.width
                            Component.onCompleted:{
                                console.log("label width: "+this.width)
                            }
                            text:addressCoinHours
                        }
                    }
                }
            }

        } // ListView
        RowLayout{
            Label{
                text:qsTr("Outputs")
                font.bold:true
            }
        }
        RowLayout{
            Rectangle {
                id: rectOutput
                Layout.fillWidth: true
                height: 1
                color: "#DDDDDD"
            }
        }
        ListView {
            id: outputList
            implicitHeight: delegateHeight*(outputList.count) + 20
            clip: true
            interactive: false
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignTop

            Behavior on implicitHeight { NumberAnimation { duration: 250; easing.type: Easing.OutQuint } }
            Behavior on opacity { NumberAnimation { duration: expanded ? 250 : 1000; easing.type: Easing.OutQuint } }

            delegate: ItemDelegate {

                height: delegateHeight
                ColumnLayout{
                    Layout.fillWidth:true
                    RowLayout{
                        Label{
                            width:parent.width
                            text:address
                        }
                    }
                    RowLayout{
                        Label{
                            width:parent.width
                            font.bold:true
                            text:qsTr("Coins: ")
                        }
                        Label{
                            width:parent.width
                            Component.onCompleted:{
                                console.log("label width: "+this.width)
                            }
                            text:addressSky
                        }
                    }
                    RowLayout{
                        Label{
                            width:parent.width

                            font.bold:true
                            text:qsTr("CoinsHours: ")
                        }
                        Label{
                            width:parent.width
                            Component.onCompleted:{
                                console.log("label width: "+this.width)
                            }
                            text:addressCoinHours
                        }
                    }
                }
            }

        } // ListView
    } // ColumnLayout

DialogTransactionDetails {
        id: dialogTransactionDetails

        readonly property real maxHeight: expanded ? 590 : 370

        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 640 ? 640 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > maxHeight ? maxHeight - 40 : applicationWindow.height - 40
        Behavior on height { NumberAnimation { duration: 1000; easing.type: Easing.OutQuint } }

        modal: true
        focus: true

        date: listTransactions.currentItem !== null ? listTransactions.currentItem.modelDate : ""
        status: listTransactions.currentItem !== null ? listTransactions.currentItem.modelStatus : 0
        type: listTransactions.currentItem !== null ? listTransactions.currentItem.modelType : 0
        amount: modelData.amount !== null ? modelData.amount : ""
        hoursReceived: listTransactions.currentItem !== null ? listTransactions.currentItem.modelHoursReceived : 1
        hoursBurned: listTransactions.currentItem !== null ?  listTransactions.currentItem.modelHoursBurned : 1
        transactionID: modelData.transactionId !== null ? modelData.transactionId : ""
        modelInputs: listTransactions.currentItem !== null ? listTransactions.currentItem.modelInputs : null
        modelOutputs: listTransactions.currentItem !== null ? listTransactions.currentItem.modelOutputs : null
    }

}
