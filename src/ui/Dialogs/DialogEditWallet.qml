import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/Controls"
import "../Controls" // For quick UI development, switch back to resources when making a release

Dialog {
    id: dialogEditWallet

    property string originalWalletName
    property alias name: textFieldName.text

    title: qsTr("Edit wallet")
    standardButtons: Dialog.Ok | Dialog.Cancel

    function clear() {
        textFieldName.clear()
    }

    Component.onCompleted: {
    }

    onAboutToShow: {
        textFieldName.forceActiveFocus()
        standardButton(Dialog.Ok).enabled = name && name != originalWalletName
    }

    onAboutToHide: {
        clear()
    }

    TextField {
        id: textFieldName

        anchors.fill: parent
        placeholderText: qsTr("Wallet's name")
        selectByMouse: true
        focus: true

        onTextChanged: {
            standardButton(Dialog.Ok).enabled = text && text != originalWalletName
        }
    }
}