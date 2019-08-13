import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Page {
    id: root

    signal qrCodeRequested(var data)

    Component.onCompleted: {
        root.qrCodeRequested.connect(genQR)
    }

    function genQR(data) {
        dialogQR.text = data
        dialogQR.open()

    }

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
                model: ["Wallet A", "Wallet B", "Wallet C", "Wallet D", "Wallet E"]

                // Taken from Qt 5.13.0 source code:
                delegate: MenuItem {
                    width: parent.width
                    text: comboBoxWalletsSendFrom.textRole ? (Array.isArray(comboBoxWalletsSendFrom.model) ? modelData[comboBoxWalletsSendFrom.textRole] : model[comboBoxWalletsAddressesSendFrom.textRole]) : modelData
                    Material.foreground: comboBoxWalletsSendFrom.currentIndex === index ? parent.Material.accent : parent.Material.foreground
                    highlighted: comboBoxWalletsSendFrom.highlightedIndex === index
                    hoverEnabled: comboBoxWalletsSendFrom.hoverEnabled
                    leftPadding: highlighted ? 2*padding : padding // added
                    Behavior on leftPadding { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } } // added
                }
            } // ComboBox
        } // ColumnLayout (send from)

        ColumnLayout {
            id: columnLayoutSendTo

            Layout.alignment: Qt.AlignTop
            Label {
                text: qsTr("Send to")
            }
            RowLayout {
                Layout.fillWidth: true
                spacing: 8

                ToolButtonQR {
                    id: toolButtonQR
                    Layout.bottomMargin: 4

                    iconSize: "24x24"

                    onClicked: {
                        qrCodeRequested(textFieldWalletsSendTo.text)
                    }
                }
                
                TextField {
                    id: textFieldWalletsSendTo
                    font.family: "Code New Roman"
                    placeholderText: qsTr("Destination address")
                    Layout.fillWidth: true
                    Layout.topMargin: -5
                }
            } // RowLayout
        } // ColumnLayout (send to)

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
        } // ColumnLayout (root)
    } // ColumnLayout (root)
}
