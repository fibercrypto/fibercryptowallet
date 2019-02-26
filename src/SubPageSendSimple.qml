import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Layouts 1.12

Page {
    id: root

    ColumnLayout {
        id: columnLayoutRoot
        anchors.fill: parent
        anchors.leftMargin: 10
        anchors.rightMargin: 10
        spacing: 20

        ColumnLayout {
            id: columnLayoutSendFrom

            Layout.alignment: Qt.AlignTop
            Label {
                text: qsTr("Send from")
            }
            ComboBox {
                id: comboBoxWalletsSendFrom
                Layout.fillWidth: true
            }
        }

        ColumnLayout {
            id: columnLayoutSendTo

            Layout.alignment: Qt.AlignTop
            Label {
                text: qsTr("Send to")
            }
            RowLayout {
                Layout.fillWidth: true
                spacing: 8
                Image {
                    source: "qrc:/images/qr.svg"
                    sourceSize: "24x24"
                }
                TextField {
                    id: textFieldWalletsSendTo
                    font.family: "Code New Roman"
                    placeholderText: qsTr("Destination address")
                    Layout.fillWidth: true
                    Layout.topMargin: -5
                }
            }
        }

        ColumnLayout {
            id: columnLayoutAmount

            Layout.alignment: Qt.AlignTop
            Label {
                text: qsTr("Amount")
            }
            TextField {
                id: textFieldAmount
                placeholderText: qsTr("Amount to send")
                Layout.fillWidth: true
                Layout.topMargin: -10
                validator: DoubleValidator {
                    notation: DoubleValidator.StandardNotation
                }
            }
        }
    } // ColumnLayout (root)
}
