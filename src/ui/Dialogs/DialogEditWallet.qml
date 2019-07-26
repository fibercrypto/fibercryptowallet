import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Layouts 1.12

Dialog {
    id: dialogEditWallet
    title: Qt.application.name
    standardButtons: Dialog.Ok | Dialog.Cancel

    onAboutToShow: {
        textFieldName.forceActiveFocus()
    }

    ColumnLayout {
        id: columnLayoutRoot
        anchors.fill: parent

        Label {
            text: qsTr("Name")
            Layout.fillWidth: true
        }
        TextField {
            id: textFieldName
            Layout.fillWidth: true
            placeholderText: qsTr("Wallet's name")
            text: name
            focus: true
        }
    } // ColumnLayout (root)
}