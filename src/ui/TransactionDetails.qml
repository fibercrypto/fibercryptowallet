import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import HistoryModels 1.0

// Resource imports
import "qrc:/ui/src/ui/Delegates"

Item {
    id: root

    property date date: "2000-01-01 00:00"
    property int type: TransactionDetails.Type.Send
    property int status: TransactionDetails.Status.Preview
    property var statusString: [ qsTr("Confirmed"), qsTr("Pending"), qsTr("Preview") ]
    property real amount: 0
    property int hoursReceived: 0
    property int hoursBurned: 0
    property string transactionID

    readonly property bool expanded: buttonMoreDetails.checked

    enum Status {
        Confirmed,
        Pending,
        Preview
    }
    enum Type {
        Send,
        Receive
    }

    implicitHeight: 80 + rowLayoutBasicDetails.height + (expanded ? rowLayoutMoreDetails.height : 0)

    // TODO: Fix performance problem with the animation
    //Behavior on implicitHeight { NumberAnimation { duration: 1000; easing.type: Easing.OutQuint } }

    implicitWidth: 650
    clip: true

    ColumnLayout {
        id: columnLayoutRoot
        anchors.fill: parent
        spacing: 20

        RowLayout {
            id: rowLayoutBasicDetails
            Layout.alignment: Qt.AlignTop
            Layout.fillWidth: true

            ColumnLayout {
                Layout.alignment: Qt.AlignTop
                Layout.fillWidth: true

                Label {
                    text: qsTr("Transaction")
                    font.bold: true
                    Layout.alignment: Qt.AlignTop
                    Layout.fillWidth: true
                }

                GridLayout {
                    id: gridLayoutBasicInfo
                    Material.foreground: Material.Grey
                    columns: 2
                    columnSpacing: 10

                    Layout.alignment: Qt.AlignTop
                    Layout.fillWidth: true

                    Label {
                        text: qsTr("Date:")
                        font.pointSize: Qt.application.font.pointSize * 0.9
                        font.bold: true
                    }
                    Label {
                        text: Qt.formatDateTime(root.date, Qt.DefaultLocaleShortDate)
                        font.pointSize: Qt.application.font.pointSize * 0.9
                    }

                    Label {
                        text: qsTr("Status:")
                        font.pointSize: Qt.application.font.pointSize * 0.9
                        font.bold: true
                    }
                    Label {
                        text: statusString[root.status]
                        font.pointSize: Qt.application.font.pointSize * 0.9
                    }

                    Label {
                        text: qsTr("Hours:")
                        font.pointSize: Qt.application.font.pointSize * 0.9
                        font.bold: true
                    }
                    Label {
                        text: root.hoursReceived + ' ' + qsTr("received") + ' | ' + hoursBurned + ' ' + qsTr("burned")
                        font.pointSize: Qt.application.font.pointSize * 0.9
                    }

                    Label {
                        text: qsTr("Tx ID:")
                        font.pointSize: Qt.application.font.pointSize * 0.9
                        font.bold: true
                    }
                    Label {
                        text: root.transactionID
                        font.pointSize: Qt.application.font.pointSize * 0.9
                        Layout.fillWidth: true
                    }
                } // GridLayout
            }

            ColumnLayout {
                Layout.alignment: Qt.AlignTop
                Layout.rightMargin: 20
                Image {
                    source: "qrc:/images/resources/images/icons/send-" + (type === TransactionDetails.Type.Receive ? "blue" : "amber") + ".svg"
                    sourceSize: "72x72"
                    fillMode: Image.PreserveAspectFit
                    mirror: type === TransactionDetails.Type.Receive
                    Layout.fillWidth: true
                }
                Label {
                    text: (type === TransactionDetails.Type.Receive ? "Receive" : "Send") + ' ' + amount + ' ' + qsTr("SKY")
                    font.bold: true
                    font.pointSize: Qt.application.font.pointSize * 1.25
                    horizontalAlignment: Label.AlignHCenter
                    Layout.fillWidth: true
                }
            }
        } // RowLayout

        RowLayout {
            id: rowLayoutDetailsSeparator

            Layout.alignment: Qt.AlignTop
            Layout.fillWidth: true

            Button {
                id: buttonMoreDetails
                implicitWidth: 200
                flat: true
                checkable: true
                text: (checked ? qsTr("Less") : qsTr("More")) + ' ' + qsTr("details")
            }

            Rectangle {
                height: 1
                color: Material.color(Material.Grey)
                Layout.alignment: Qt.AlignVCenter
                Layout.fillWidth: true
            }
        }

        RowLayout {
            id: rowLayoutMoreDetails

            Layout.alignment: Qt.AlignTop
            Layout.fillWidth: true

            ColumnLayout {
                Layout.alignment: Qt.AlignTop
                Layout.fillWidth: true

                Label {
                    text: qsTr("Inputs")
                    font.bold: true
                    font.italic: true
                    Layout.fillWidth: true
                }

                ScrollView {
                    Layout.alignment: Qt.AlignTop
                    Layout.fillWidth: true

                    contentHeight: 160

                    ListView {
                        id: listViewInputs

                        Material.foreground: Material.Grey
                        model: listModelInputs
                        clip: true
                        Layout.alignment: Qt.AlignTop
                        Layout.fillWidth: true
                        delegate: InputOutputDelegate {
                            width: ListView.view.width
                        }
                    }
                }
            }

            ColumnLayout {
                Layout.alignment: Qt.AlignTop
                Layout.fillWidth: true

                Label {
                    text: qsTr("Outputs")
                    font.bold: true
                    font.italic: true
                    Layout.fillWidth: true
                }

                ScrollView {
                    Layout.alignment: Qt.AlignTop
                    Layout.fillWidth: true

                    contentHeight: 160

                    ListView {
                        id: listViewOutputs

                        Material.foreground: Material.Grey
                        model: listModelOutputs
                        clip: true
                        Layout.alignment: Qt.AlignTop
                        Layout.fillWidth: true
                        delegate: InputOutputDelegate {
                            width: ListView.view.width
                        }
                    }
                }
            }
        } // RowLayout
    } // ColumnLayout (root)

    // Roles: address, addressSky, addressCoinHours
    // Use listModel.append( { "address": value, "addressSky": value, "addressCoinHours": value } )
    // Or implement the model in the backend (a more recommendable approach)
    // ListModel {
    //     id: listModelInputs
    //     ListElement { address: "qrxw7364w8xerusftaxkw87ues"; addressSky: 30; addressCoinHours: 1049 }
    //     ListElement { address: "8745yuetsrk8tcsku4ryj48ije"; addressSky: 12; addressCoinHours: 16011 }
    // }
    // ListModel {
    //     id: listModelOutputs
    //     ListElement { address: "734irweaweygtawieta783cwyc"; addressSky: 38; addressCoinHours: 5048 }
    //     ListElement { address: "ekq03i3qerwhjqoqh9823yurig"; addressSky: 61; addressCoinHours: 9456 }
    //     ListElement { address: "1kjher73yiner7wn32nkuwe94v"; addressSky: 1; addressCoinHours: 24 }
    //     ListElement { address: "oopfwwklfd34iuhjwe83w3h28r"; addressSky: 111; addressCoinHours: 13548 }
    // }
    QAddressList{
        id: listModelInputs
    }
    QAddressList{
        id: listModelOutputs
    }
}
