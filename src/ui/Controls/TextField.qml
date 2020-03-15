import QtQuick 2.12
import QtQuick.Controls 2.12

TextField {
    id: control

    Menu {
        id: contextMenu

        focus: false

        ItemDelegate { text: qsTr("Cut"); onClicked: { control.cut(); contextMenu.close() } enabled: control.selectedText && control.echoMode === TextInput.Normal }
        ItemDelegate { text: qsTr("Copy"); onClicked: { control.copy(); contextMenu.close() } enabled: control.selectedText && control.echoMode === TextInput.Normal }
        ItemDelegate { text: qsTr("Paste"); onClicked: { control.paste(); contextMenu.close() } enabled: control.canPaste }
        MenuSeparator {}
        ItemDelegate { text: qsTr("Select all"); onClicked: { control.selectAll(); contextMenu.close() } enabled: control.selectedText !== control.text }
        MenuSeparator {}
        ItemDelegate { text: qsTr("Undo"); onClicked: { control.undo(); contextMenu.close() } enabled: control.canUndo }
        ItemDelegate { text: qsTr("Redo"); onClicked: { control.redo(); contextMenu.close() } enabled: control.canRedo }
    }

    MouseArea {
        anchors.fill: parent
        acceptedButtons: Qt.RightButton
        cursorShape: Qt.IBeamCursor

        onClicked: {
            control.focus = true
            contextMenu.popup()
        }
    }
}
