import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import HistoryModels 1.0

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "Delegates/" // For quick UI development, switch back to resources when making a release

// Backend imports
import HistoryModels 1.0

Item {
    id: root

    property date date: "2000-01-01 00:00"
    property int type: TransactionDetails.Type.Send
    property int status: TransactionDetails.Status.Preview
    property var statusString: [ qsTr("Confirmed"), qsTr("Pending"), qsTr("Preview") ]
    property real amount: 0
    property string hoursReceived
    property string hoursBurned
    property string transactionID
    property QAddressList modelInputs
    property QAddressList modelOutputs

    readonly property bool expanded: buttonMoreDetails.checked

    enum Status {
        Confirmed,
        Pending,
        Preview
    }
    enum Type {
        Send,
        Receive,
        Internal
    }

    implicitHeight: 80 + rowLayoutBasicDetails.height + (expanded ? rowLayoutMoreDetails.height : 0)
    Behavior on implicitHeight { NumberAnimation { duration: 1000; easing.type: Easing.OutQuint } }

    implicitWidth: 600
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
                Layout.topMargin: -10
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
                    font.pointSize: Qt.application.font.pointSize * 1.15
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

            opacity: expanded ? 1 : 0
            Behavior on opacity { NumberAnimation { duration: 1000; easing.type: Easing.OutQuint } }

            ColumnLayout {
                id: columnLayoutInputs
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
                        model: modelInputs
                        clip: true
                        Layout.alignment: Qt.AlignTop
                        Layout.fillWidth: true
                        delegate: InputOutputDelegate {
                            width: ListView.view.width
                        }
                    }
                } // ScrollView
            } // ColumnLayout (inputs)

            ColumnLayout {
                id: columnLayoutOutputs
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
                        model: modelOutputs
                        clip: true
                        Layout.alignment: Qt.AlignTop
                        Layout.fillWidth: true
                        delegate: InputOutputDelegate {
                            width: ListView.view.width
                        }
                    }
                } // ScrollView
            } // ColumnLayout (outputs)
        } // RowLayout (details)
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
    // QAddressList{
    //     id: listModelInputs
    // }
    // QAddressList{
    //     id: listModelOutputs
    // }
}
