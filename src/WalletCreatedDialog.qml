import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Dialog {
    id: root

    focus: visible
    modal: true
    title: Qt.application.name
    standardButtons: Dialog.Ok
    closePolicy: Dialog.NoAutoClose

    implicitWidth: 500
    implicitHeight: 300

    Component.onCompleted: {
        standardButton(Dialog.Ok).enabled = name.text !== ""
    }

    ColumnLayout {
        anchors.fill: parent
        spacing: 30

        RowLayout {
            spacing: 30

            Image {
                id: icon
                source: "qrc:/images/icons/check.svg"
                sourceSize: "64x64"
                Layout.alignment: Qt.AlignTop
            }

            ColumnLayout {
                Label {
                    text: qsTr("The connected hardware wallet will be added " +
                               "to the wallets list with the following name:")
                    Layout.fillWidth: true
                    wrapMode: Text.WordWrap
                    Layout.alignment: Qt.AlignTop
                }

                TextField {
                    id: name
                    placeholderText: qsTr("New hardware wallet")
                    Layout.fillWidth: true
                    focus: root.focus

                    onTextChanged: {
                        standardButton(Dialog.Ok).enabled = text !== ""
                    }
                }
            }
        }

        Label {
            text: qsTr("Now you can check the balance and the address of the " +
                       "hardware wallet even when it is not connected. " +
                       "Additionally, you can change the name of the wallet " +
                       "or remove it from the wallets list, if you wish.")
            Layout.fillWidth: true
            wrapMode: Text.WordWrap
            Layout.alignment: Qt.AlignTop
        }
    }
}
