import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Dialog {
    id: root

    title: qsTr("Unconfigured wallet")
    standardButtons: Dialog.Abort
    closePolicy: Dialog.NoAutoClose

    implicitWidth: 650
    implicitHeight: 380

    ColumnLayout {
        anchors.fill: parent
        spacing: 50

        RowLayout {
            spacing: 50

            Image {
                id: icon
                source: "qrc:/images/resources/images/icons/cpu.svg"
                sourceSize: "64x64"
                Layout.alignment: Qt.AlignTop
            }

            ColumnLayout {
                Label {
                    id: msgTitle
                    text: qsTr("Unconfigured Hardware Wallet")
                    font.bold: true
                    Layout.fillWidth: true
                    wrapMode: Text.WordWrap
                    Layout.alignment: Qt.AlignTop
                }
                Label {
                    id: msg
                    text: qsTr("A seedless hardware wallet has been detected. " +
                               "Select <b><i>Configure automatically</i></b> " +
                               "if you want to configure it as a brand new wallet " +
                               "and start using it inmediately, " +
                               "or select <b><i>Restore backup</i></b> " +
                               "if you want to configure it with a previously created " +
                               "seed backup and thus be able to access your funds again.")
                    Layout.fillWidth: true
                    wrapMode: Text.WordWrap
                    Layout.alignment: Qt.AlignTop
                }
            }
        }

        ColumnLayout {
            Label {
                id: options
                text: qsTr("Options:")
                font.bold: true
                Layout.fillWidth: true
            }
            ItemDelegate {
                id: buttonAutoConf
                text: qsTr("Configure automatically")
                Layout.fillWidth: true
            }
            ItemDelegate {
                id: buttonBackupConf
                text: qsTr("Restore backup")
                Layout.fillWidth: true

            }
        }
    }
}
