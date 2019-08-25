import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

ColumnLayout {
    width: scrollSettings.width
    property alias isSetWalletEnvLocal: walletEnvSwitch.checked
    property alias walletPath: walletPathTextEdit.text

    RowLayout {
        width: scrollSettings.width
        height: 40
        Label {
            font.pointSize: Qt.application.font.pointSize * 0.9
            font.bold: true
            Layout.alignment: Qt.AlignCenter
            text: "Wallet Path:"
        }
        TextField {
            id: walletPathTextEdit
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignCenter
            placeholderText: "$HOME/.skycoin/wallets"
        }
    }

    RowLayout {
        width: scrollSettings.width
        height: 40
        Label {
            font.pointSize: Qt.application.font.pointSize * 0.9
            font.bold: true
            Layout.alignment: Qt.AlignCenter
            text: "Wallet Enviroment (Local/Remote):"
        }
        Switch {
            id: walletEnvSwitch
            checked: false
        }
    }
}