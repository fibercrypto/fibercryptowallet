import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Page {
    id: root

    property int numberOfBlocks: 0
    property date timestampLastBlock: "2000-01-01 00:00"
    property string hashLastBlock;

    property int currentSkySupply: 0
    property int totalSkySupply: 0
    property int currentCoinHoursSupply: 0
    property int totalCoinHoursSupply: 0

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
                spacing: 20
                Material.foreground: Material.Grey

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
                            Layout.fillWidth: true
                        }
                        Label {
                            id: labelHashLastBlock
                            text: hashLastBlock
                            wrapMode: Label.WrapAnywhere
                        }
                    }
                }
            } // ColumnLayout
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
                Material.foreground: Material.Grey

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
