import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

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
        id: textEdit
        Layout.fillWidth: true
        Layout.alignment: Qt.AlignCenter
        placeholderText: "$HOME/.skycoin/wallets"
    }
}
