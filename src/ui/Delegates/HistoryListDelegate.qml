import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/"
import "../" // For quick UI development, switch back to resources when making a release

ItemDelegate {
    id: root

    property date modelDate: date
    property int modelType: type
    property int modelStatus: status
    property var modelStatusString: [ qsTr("Confirmed"), qsTr("Pending"), qsTr("Preview") ]
    property real modelAmount: amount
    property int modelHoursReceived: hoursReceived
    property int modelHoursBurned: hoursBurned
    property string modelTransactionID

    signal qrCodeRequested(var data)

    Component.onCompleted: {
        root.qrCodeRequested.connect(genQR)
    }

    function genQR(data) {
        dialogQR.text = data
        dialogQR.open()

    }

    implicitWidth: parent.width
    implicitHeight: (columnLayoutMainContent.height < 78 ? 78 : columnLayoutMainContent.height) + rowLayoutRoot.anchors.topMargin + rowLayoutRoot.anchors.bottomMargin

    RowLayout {
        id: rowLayoutRoot
        anchors.fill: parent
        anchors.leftMargin: 20
        anchors.rightMargin: 20
        anchors.topMargin: 10
        anchors.bottomMargin: 12

        spacing: 20

        Image {
            source: "qrc:/images/resources/images/icons/send-blue.svg"
            sourceSize: "32x32"
            fillMode: Image.PreserveAspectFit
            mirror: modelType === TransactionDetails.Type.Receive
            Layout.alignment: Qt.AlignTop | Qt.AlignLeft
        }

        ColumnLayout {
            id: columnLayoutMainContent
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignTop

            RowLayout {
                spacing: 20
                Layout.fillWidth: true

                Label {
                    font.bold: true
                    text: (modelType === TransactionDetails.Type.Receive ? qsTr("Received") : qsTr("Sent")) + " SKY"
                }

                Label {
                    Material.foreground: Material.Grey
                    text: modelDate // model's role
                    font.pointSize: Qt.application.font.pointSize * 0.9
                }
            }

            ColumnLayout {
                RowLayout {
                    id: rowLayoutSent
                    visible: modelType === TransactionDetails.Type.Send

                    ToolButtonQR {
                        id: toolButtonQRSent

                        iconSize: "24x24"

                        onClicked: {
                            qrCodeRequested(sentAddress)
                        }
                    }

                    Label {
                        text: sentAddress // model's role
                        font.family: "Code New Roman"
                        Layout.fillWidth: true
                    }
                }
                RowLayout {
                    id: rowLayoutReceive

                    ToolButtonQR {
                        id: toolButtonQRReceived

                        iconSize: "24x24"

                        onClicked: {
                            qrCodeRequested(receivedAddress)
                        }
                    }
                    
                    Label {
                        text: receivedAddress // model's role
                        font.family: "Code New Roman"
                        Layout.fillWidth: true
                    }
                }
            } // ColumnLayout (addresses)
        } // ColumnLayout (main content)

        Label {
            text: (modelType === TransactionDetails.Type.Receive ? "" : "-") + amount + " SKY" // model's role
            font.pointSize: Qt.application.font.pointSize * 1.25
            font.bold: true
            Layout.alignment: Qt.AlignTop | Qt.AlignRight
        }

    } // RowLayout (root)
}
