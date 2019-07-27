import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Dialog {
    id: restoreBackupWordsDialog

    property bool finalEntry: false    

    title: Qt.application.name
    standardButtons: Dialog.YesToAll | Dialog.Cancel
    closePolicy: Dialog.NoAutoClose

    Component.onCompleted: {
        standardButton(Dialog.YesToAll).text = finalEntry ? qsTr("Finish") : qsTr("Next")
        standardButton(Dialog.YesToAll).enabled = wordInput.text !== ""
    }

    Flickable {
        id: flickable
        anchors.fill: parent
        contentHeight: columnLayoutRoot.height
        clip: true

        ColumnLayout {
            id: columnLayoutRoot
            width: parent.width
            spacing: 50

            RowLayout {
                id: rowLayoutIconMessage
                spacing: 30

                Image {
                    id: icon
                    source: "qrc:/images/resources/images/icons/backup.svg"
                    sourceSize: "64x64"
                    Layout.alignment: Qt.AlighLeft | Qt.AlignTop
                }

                ColumnLayout {
                    id: columnLayoutMessage
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
                } // ColumnLayout (message)
            } // RowLayout (icon + message)

            TextField {
                id: wordInput
                placeholderText: qsTr("Requested word")
                Layout.fillWidth: true
                focus: restoreBackupWordsDialog.focus

                onTextChanged: {
                    standardButton(Dialog.YesToAll).enabled = text !== ""
                }
            }
        } // ColumnLayout (root)

        ScrollIndicator.vertical: ScrollIndicator {
            parent: restoreBackupWordsDialog.contentItem
            anchors.top: flickable.top
            anchors.bottom: flickable.bottom
            anchors.right: parent.right
            anchors.rightMargin: -restoreBackupWordsDialog.rightPadding + 1
        }
    } // Flickable
}
