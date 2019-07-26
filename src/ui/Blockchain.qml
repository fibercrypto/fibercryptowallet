import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import BlockchainModels 1.0

Page {
    id: root

    property string numberOfBlocks;//: model.numberOfBlocks
    property date timestampLastBlock;//: model.timestampLastBlock
    property string hashLastBlock;//: model.hashLastBlock

    property string currentSkySupply;//: model.currentSkySupply
    property string totalSkySupply;//: model.totalSkySupply
    property string currentCoinHoursSupply;//: model.currentCoinHoursSupply
    property string totalCoinHoursSupply;//: model.totalCoinHoursSupply

    signal backRequested()

    header: RowLayout {
        spacing: 20
        ToolButton {
            id: toolButtonBack
            icon.source: "qrc:/images/resources/images/icons/back.svg"
            Layout.alignment: Qt.AlignLeft

            onClicked: {
                backRequested()
            }
        }

        Label {
            text: qsTr("Blockchain")
            font.pointSize: Qt.application.font.pointSize * 2
            horizontalAlignment: Text.AlignHCenter
            leftPadding: -(toolButtonBack.width + parent.spacing)
            Layout.fillWidth: true
        }
    }

    ColumnLayout {
        id: columnLayoutRoot
        anchors.top: parent.top
        anchors.left: parent.left
        anchors.right: parent.right
        anchors.margins: 20
        spacing: 20

        GroupBox {
            id: groupBoxBlockDetails
            title: qsTr("Block details")
            clip: true
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignTop | Qt.AlignHCenter

            ColumnLayout {
                id: columnLayoutBlockDetails
                anchors.fill: parent
                spacing: 20

                ColumnLayout {
                    Label {
                        text: qsTr("Number of blocks")
                        font.bold: true
                    }
                    Label {
                        id: labelNumberOfBlocks
                        text: numberOfBlocks
                    }
                }

                RowLayout {
                    spacing: 20

                    ColumnLayout {
                        Label {
                            text: qsTr("Timestamp of last block")
                            font.bold: true
                        }
                        Label {
                            id: labelTimestampLastBlock
                            text: Qt.formatDateTime(root.timestampLastBlock, Qt.DefaultLocaleShortDate)
                        }
                    }

                    ColumnLayout {
                        Layout.fillWidth: true
                        Label {
                            text: qsTr("Hash of last block")
                            font.bold: true

                            ToolButton {
                                id: toolButtonCopy
                                anchors.left: parent.right
                                anchors.verticalCenter: parent.verticalCenter
                                icon.source: "qrc:/images/resources/images/icons/copy.svg"
                                visible: textInputHashLastBlock.text
                                ToolTip.text: qsTr("Copy to clipboard")
                                ToolTip.visible: hovered // TODO: pressed when mobile?
                                ToolTip.delay: Qt.styleHints.mousePressAndHoldInterval

                                Image {
                                    id: imageCopied
                                    anchors.centerIn: parent
                                    source: "qrc:/images/resources/images/icons/check-simple.svg"
                                    fillMode: Image.PreserveAspectFit
                                    sourceSize: Qt.size(toolButtonCopy.icon.width*1.5, toolButtonCopy.icon.height*1.5)
                                    z: 1

                                    opacity: 0.0
                                }

                                onClicked: {
                                    if (textInputHashLastBlock.text) {
                                        textInputHashLastBlock.selectAll()
                                        textInputHashLastBlock.copy()
                                        textInputHashLastBlock.deselect()
                                        if (copyAnimation.running) {
                                            copyAnimation.restart()
                                        } else {
                                            copyAnimation.start()
                                        }
                                    }
                                }

                                SequentialAnimation {
                                    id: copyAnimation
                                    NumberAnimation { target: imageCopied; property: "opacity"; to: 1.0; easing.type: Easing.OutCubic }
                                    PauseAnimation { duration: 1000 }
                                    NumberAnimation { target: imageCopied; property: "opacity"; to: 0.0; easing.type: Easing.OutCubic }
                                }
                            } // ToolButton
                        } // Label
                        TextInput {
                            id: textInputHashLastBlock
                            Layout.fillWidth: true
                            text: hashLastBlock
                            readOnly: true
                            font.family: "Code New Roman"
                            wrapMode: Label.WrapAnywhere
                        }
                    } // ColumnLayout
                } // RowLayout
            } // ColumnLayout (block details)
        } // GroupBox (block details)

        GroupBox {
            id: groupBoxSkyDetails
            title: qsTr("Sky details")
            clip: true
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignTop | Qt.AlignHCenter

            GridLayout {
                rows: 2
                columns: 2

                ColumnLayout {
                    Label {
                        text: qsTr("Current SKY supply")
                        font.bold: true
                    }
                    Label {
                        id: labelCurrentSkySupply
                        text: currentSkySupply
                    }
                }

                ColumnLayout {
                    Label {
                        text: qsTr("Total SKY supply")
                        font.bold: true
                    }
                    Label {
                        id: labelTotalSkySupply
                        text: totalSkySupply
                    }
                }

                ColumnLayout {
                    Label {
                        text: qsTr("Current Coin Hours supply")
                        font.bold: true
                    }
                    Label {
                        id: labelCurrentCoinHoursSupply
                        text: currentCoinHoursSupply
                    }
                }

                ColumnLayout {
                    Label {
                        text: qsTr("Total Coin Hours supply")
                        font.bold: true
                    }
                    Label {
                        id: labelTotalCoinHoursSupply
                        text: totalCoinHoursSupply
                    }
                }
            } // GridLayout
        } // GroupBox (sky details)
    } // ColumnLayout (root)
}
