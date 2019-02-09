import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Dialog {
    id: root

    property string deviceName: "<NULL>"
    property bool enableBackupWarning: true
    property bool enablePINWarning: true

    focus: true
    modal: true
    title: Qt.application.name
    standardButtons: Dialog.Abort
    closePolicy: Dialog.NoAutoClose

    implicitWidth: 600
    implicitHeight: 550 - (enableBackupWarning ^ enablePINWarning ? 100 : 0) - (!enableBackupWarning && !enablePINWarning ? 240 : 0)

    ColumnLayout {
        anchors.fill: parent
        spacing: 50

        RowLayout {
            spacing: 30
            
            Image {
                id: icon
                source: "qrc:/images/icons/security.svg"
                sourceSize: "72x72"
                Layout.alignment: Qt.AlignTop
            }
            ColumnLayout {
                Label {
                    id: msgTitle
                    text: qsTr("Secure hardware wallet")
                    font.bold: true
                    Layout.fillWidth: true
                    wrapMode: Text.WordWrap
                    Layout.alignment: Qt.AlignTop
                }
                Label {
                    id: msg
                    text: qsTr("Hardware wallet detected.<br>The device is identified in the " +
                               "wallets list as: ") + "<b><i>" + deviceName + "</i></b>"
                    Layout.fillWidth: true
                    wrapMode: Text.WordWrap
                    Layout.alignment: Qt.AlignTop
                }
            }
        }

        ColumnLayout {
            spacing: 30

            ColumnLayout {
                visible: enableBackupWarning || enablePINWarning
                Label {
                    id: securityWarningsTitle
                    text: qsTr("Security warnings")
                    font.pointSize: 14
                    font.bold: true
                    color: Material.color(Material.Pink)
                    Layout.fillWidth: true
                    wrapMode: Text.WordWrap
                    Layout.alignment: Qt.AlignTop
                }
                Label {
                    id: securityWarningBackup
                    visible: enableBackupWarning
                    text: "<b>1)</b> " +
                          qsTr("You should secure the hardware seed or you could lose " +
                               "access to the funds in case of problems. " +
                               "To do this, select <b><i>Create a backup</i></b>.")
                    color: Material.color(Material.Pink)
                    Layout.fillWidth: true
                    wrapMode: Text.WordWrap
                    Layout.alignment: Qt.AlignTop
                }
                Label {
                    id: securityWarningPIN
                    visible: enablePINWarning
                    text: "<b>" + (enableBackupWarning ? 2 : 1) + ")</b> " +
                          qsTr("The connected hardware wallet does not have a PIN. " +
                               "The PIN code protects the hardware wallet in case of " +
                               "loss, theft and hacks." +
                               "To do this, select <b><i>Create PIN code</i></b>.")
                    color: Material.color(Material.Pink)
                    Layout.fillWidth: true
                    wrapMode: Text.WordWrap
                    Layout.alignment: Qt.AlignTop
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
                    id: buttonCreateBackup
                    visible: enableBackupWarning
                    text: qsTr("Create a backup")
                    Layout.fillWidth: true
                }
                ItemDelegate {
                    id: buttonCreatePIN
                    visible: enablePINWarning
                    text: qsTr("Create PIN code")
                    Layout.fillWidth: true
                }
                ItemDelegate {
                    id: buttonWipeDevice
                    text: qsTr("Wipe device")
                    Layout.fillWidth: true
                    Material.foreground: Material.Pink
                }
            }

        }
    }
}
