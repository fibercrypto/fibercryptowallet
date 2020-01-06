import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import QtGraphicalEffects 1.12

Item{
    id: root

    readonly property real delegateHeight: 80
    property alias inputsModel: inputList.model
    property alias outputsModel: outputList.model
    readonly property real finalViewHeight: delegateHeight + 400

    width: parent.width
    height: itemDelegateMainButton.height+inputList.height+ 30 + outputList.height+30

    ColumnLayout {
        id: delegateColumnLayout
        anchors.fill: parent

        ItemDelegate {
            id: itemDelegateMainButton
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
                    text: modelData.transactionId
                    horizontalAlignment: Text.AlignLeft
                }
            } // RowLayout
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
                Component.onCompleted:{
                    console.log("delegate: "+this.width)
                }
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

}
