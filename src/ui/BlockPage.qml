import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import ExplorerModels 1.0

import "Delegates/"

Page {
    id: blockPage
    property string hash;
    property alias blockhash : textInputHashCurrentBlock.text
    property alias prevhash : textInputHashPrevBlock.text
    property alias time : labelTimestampLastBlock.text
    property alias blockNum : laberHeightOfBlock.text
    property alias totalAmount : labelTotalAmountBlock.text
    property alias blockSize : labelBlockSize.text
    property alias transactionList : blocksList.model
    QBlocks{
        id:explorerManager
    }

    Component.onCompleted:{
        explorerManager.loadBlockByHash(hash)
        blockhash = explorerManager.blockDetail.blockhash
        time = Qt.formatDateTime(explorerManager.blockDetail.time, Qt.DefaultLocaleShortDate)
        blockNum = explorerManager.blockDetail.blockNumber
        prevhash = explorerManager.blockDetail.prevBlockhash
        blockSize = explorerManager.blockDetail.size
        totalAmount = explorerManager.blockDetail.totalAmount
        transactionList=explorerManager.blockDetail.transactionList
    }

//    ColumnLayout {
//        id: columnLayoutRoot
//        anchors.top: parent.top
//        anchors.left: parent.left
//        anchors.right: parent.right
//        anchors.margins: 20
//        spacing: 20

        GroupBox {
            id: groupBoxBlockDetails
            title: qsTr("Block Details")
            clip: true
            width:parent.width
            Layout.alignment: Qt.AlignTop | Qt.AlignHCenter

            ColumnLayout {
                id: columnLayoutBlockDetails
                anchors.fill: parent
                spacing: 20
                RowLayout {
                    spacing: 82

                    ColumnLayout {
                    width:parent.width/2
                        Label {
                            text: qsTr("Height")
                            font.bold: true
                        }
                        Label {
                            id: laberHeightOfBlock
                        }
                    }
                    ColumnLayout {
                        Layout.fillWidth: true
                        Label {
                        text: qsTr("Parent Hash")
                        font.bold: true

                        ToolButton {
                                id: toolButtonCopyPrevHash
                                anchors.left: parent.right
                                anchors.verticalCenter: parent.verticalCenter
                                icon.source: "qrc:/images/resources/images/icons/copy.svg"
//                                visible: textInputHashLastBlock.text
                                ToolTip.text: qsTr("Copy to clipboard")
                                ToolTip.visible: hovered // TODO: pressed when mobile?
                                ToolTip.delay: Qt.styleHints.mousePressAndHoldInterval

                                Image {
                                    id: imageCopiedPrevHash
                                    anchors.centerIn: parent
                                    source: "qrc:/images/resources/images/icons/check-simple.svg"
                                    fillMode: Image.PreserveAspectFit
                                    sourceSize: Qt.size(toolButtonCopyPrevHash.icon.width*1.5, toolButtonCopyPrevHash.icon.height*1.5)
                                    z: 1

                                    opacity: 0.0
                                }

                                onClicked: {
                                    if (textInputHashLastBlock.text) {
                                        textInputHashLastBlock.selectAll()
                                        textInputHashLastBlock.copy()
                                        textInputHashLastBlock.deselect()
                                        if (copyAnimationPrevHash.running) {
                                            copyAnimationPrevHash.restart()
                                        } else {
                                            copyAnimationPrevHash.start()
                                        }
                                    }
                                }

                                SequentialAnimation {
                                    id: copyAnimationPrevHash
                                    NumberAnimation { target: imageCopiedPrevHash; property: "opacity"; to: 1.0; easing.type: Easing.OutCubic }
                                    PauseAnimation { duration: 1000 }
                                    NumberAnimation { target: imageCopiedPrevHash; property: "opacity"; to: 0.0; easing.type: Easing.OutCubic }
                                }
                            } // ToolButton
                        } // Label
                        TextInput {
                            anchors.right:parent.right
                            id: textInputHashPrevBlock
                            Layout.fillWidth: true
                            readOnly: true
                            font.family: "Code New Roman"
                            wrapMode: Label.WrapAnywhere
                        }
                    } // ColumnLayout
                }
                RowLayout {
                    spacing: 20

                    ColumnLayout {

                        Label {
                            text: qsTr("Block Timestamp")
                            font.bold: true
                        }
                        Label {
                            id: labelTimestampLastBlock
                        }
                    }

                    ColumnLayout {
                        Layout.fillWidth: true
                        Label {
                            text: qsTr("Block Hash")
                            font.bold: true

                            ToolButton {
                                id: toolButtonCopyCurrentHash
                                anchors.left: parent.right
                                anchors.verticalCenter: parent.verticalCenter
                                icon.source: "qrc:/images/resources/images/icons/copy.svg"
//                                visible: textInputHashLastBlock.text
                                ToolTip.text: qsTr("Copy to clipboard")
                                ToolTip.visible: hovered // TODO: pressed when mobile?
                                ToolTip.delay: Qt.styleHints.mousePressAndHoldInterval

                                Image {
                                    id: imageCopiedCurrentHash
                                    anchors.centerIn: parent
                                    source: "qrc:/images/resources/images/icons/check-simple.svg"
                                    fillMode: Image.PreserveAspectFit
                                    sourceSize: Qt.size(toolButtonCopyCurrentHash.icon.width*1.5, toolButtonCopyCurrentHash.icon.height*1.5)
                                    z: 1

                                    opacity: 0.0
                                }

                                onClicked: {
                                    if (textInputHashLastBlock.text) {
                                        textInputHashLastBlock.selectAll()
                                        textInputHashLastBlock.copy()
                                        textInputHashLastBlock.deselect()
                                        if (copyAnimationCurrentHash.running) {
                                            copyAnimationCurrentHash.restart()
                                        } else {
                                            copyAnimationCurrentHash.start()
                                        }
                                    }
                                }

                                SequentialAnimation {
                                    id: copyAnimationCurrentHash
                                    NumberAnimation { target: imageCopiedCurrentHash; property: "opacity"; to: 1.0; easing.type: Easing.OutCubic }
                                    PauseAnimation { duration: 1000 }
                                    NumberAnimation { target: imageCopiedCurrentHash; property: "opacity"; to: 0.0; easing.type: Easing.OutCubic }
                                }
                            } // ToolButton
                        } // Label
                        TextInput {
                            id: textInputHashCurrentBlock
                            Layout.fillWidth: true
                            readOnly: true
                            font.family: "Code New Roman"
                            wrapMode: Label.WrapAnywhere
                        }
                    } // ColumnLayout
                } // RowLayout
                RowLayout {
                    spacing: 60

                    ColumnLayout {

                        Label {
                            text: qsTr("Block Size")
                            font.bold: true
                        }
                        Label {
                            id: labelBlockSize
                        }
                    }

                    ColumnLayout {
                        Label {
                            text: qsTr("Total Amount")
                            font.bold: true
                        } // Label
                        Label {
                            id: labelTotalAmountBlock
                            Layout.fillWidth: true
                        }
                    } // ColumnLayout
                } // RowLayout
            } // ColumnLayout (block details)
        } // GroupBox (block details)


        GroupBox {
        anchors.bottom:parent.bottom
            id: groupBoxSkyDetails

            anchors.margins: 30
            title: qsTr("Transactions")
            height:250
            width: groupBoxBlockDetails.width
            clip: true
            Layout.alignment: Qt.AlignTop | Qt.AlignHCenter

            ScrollView{
                anchors.fill: parent
                ScrollBar.horizontal.policy: ScrollBar.AlwaysOff
                ListView {
                    id: blocksList
                    anchors.fill: parent
                    clip: true
                    delegate: TransactionListDelegate{}
                }
            }

        } // GroupBox (sky details)
//    } // ColumnLayout (root)
}
