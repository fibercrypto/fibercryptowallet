import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Dialog {
    id: root

    focus: visible
    modal: true
    title: Qt.application.name
    standardButtons: Dialog.YesToAll | Dialog.Cancel
    closePolicy: Dialog.NoAutoClose

    implicitWidth: 420
    implicitHeight: 300

    Component.onCompleted: {
        standardButton(Dialog.YesToAll).text = qsTr("Continue")
        standardButton(Dialog.YesToAll).enabled = wordInput.text !== ""
    }

    ColumnLayout {
        anchors.fill: parent
        spacing: 50

        RowLayout {
            spacing: 30

            Image {
                id: icon
                source: "qrc:/images/resources/images/icons/backup.svg"
                sourceSize: "64x64"
            }

            ColumnLayout {
                Label {
                    id: msgTitle
                    text: "<b>" + qsTr("Enter the word indicated in the device") + "</b>"
                    Layout.fillWidth: true
                    wrapMode: Text.WordWrap
                    Layout.alignment: Qt.AlignTop
                }

                Label {
                    id: msg
                    text: qsTr("You will be asked to enter the words " +
                               "of your backup seed in random order, " +
                               "plus a few additinal words.")
                    Layout.fillWidth: true
                    wrapMode: Text.WordWrap
                    Layout.alignment: Qt.AlignTop
                }
            }
        }

        TextField {
            id: wordInput
            placeholderText: qsTr("Requested word")
            Layout.fillWidth: true
            focus: root.focus

            onTextChanged: {
                standardButton(Dialog.YesToAll).enabled = text !== ""
            }
        }
    }
}
