import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

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

    enum Status {
        Confirmed,
        Pending,
        Preview
    }
    enum Type {
        Send,
        Receive
    }

    implicitHeight: 400
    implicitWidth: 650
    clip: true

    ColumnLayout {
        id: columnLayoutRoot
        anchors.fill: parent

        Label {
            text: qsTr("Transaction")
            font.bold: true
        }

        RowLayout {
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignTop

            GridLayout {
                id: gridLayoutBasicInfo
                Material.foreground: Material.Grey
                columns: 2
                columnSpacing: 10

                Layout.alignment: Qt.AlignTop
                Layout.fillWidth: true

                Label {
                    text: qsTr("Date:")
                    font.bold: true
                }
                Label {
                    text: Qt.formatDateTime(root.date, Qt.DefaultLocaleShortDate)
                }

                Label {
                    text: qsTr("Status:")
                    font.bold: true
                }
                Label {
                    text: statusString[root.status]
                }

                Label {
                    text: qsTr("Hours:")
                    font.bold: true
                }
                Label {
                    text: root.hoursReceived + ' ' + qsTr("received") + ' | ' + hoursBurned + ' ' + qsTr("burned")
                }

                Label {
                    text: qsTr("Tx ID:")
                    font.bold: true
                }
                Label {
                    text: root.transactionID
                    Layout.fillWidth: true
                }
            } // GridLayout

            ColumnLayout {
                Layout.rightMargin: 20
                Image {
                    source: "qrc:/images/send-" + (type === TransactionDetails.Type.Receive ? "blue" : "amber") + ".svg"
                    sourceSize: "96x96"
                    fillMode: Image.PreserveAspectFit
                    mirror: type === TransactionDetails.Type.Receive
                    Layout.fillWidth: true
                }
                Label {
                    text: (type === TransactionDetails.Type.Receive ? "Receive" : "Send") + ' ' + amount + ' ' + qsTr("SKY")
                    font.bold: true
                    font.pointSize: 14
                    horizontalAlignment: Label.AlignHCenter
                    Layout.fillWidth: true
                }
            }
        } // RowLayout
    }
}
