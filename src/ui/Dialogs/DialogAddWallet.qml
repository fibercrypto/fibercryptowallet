import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui"
import "../" // For quick UI development, switch back to resources when making a release

Dialog {
    id: dialogAddWallet

    title: Qt.application.name
    standardButtons: Dialog.Ok | Dialog.Cancel

    onAboutToShow: {
        createLoadWallet.clear()
    }

    ColumnLayout {
        id: columnLayoutRoot
        anchors.fill: parent

        CreateLoadWallet {
            id: createLoadWallet
            width: parent.width
        }
    } // ColumnLayout (root)
}