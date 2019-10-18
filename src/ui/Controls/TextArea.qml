import QtQuick 2.12
import QtQuick.Controls 2.12

TextArea {
    id: control

    Menu {
        id: contextMenu

        focus: false

        ItemDelegate { text: qsTr("Cut"); icon.source: "qrc:/images/resources/images/icons/cut-content.svg"; onClicked: { control.cut(); contextMenu.close() } enabled: control.selectedText && control.echoMode === TextInput.Normal }
        ItemDelegate { text: qsTr("Copy"); icon.source: "qrc:/images/resources/images/icons/copy-content.svg"; onClicked: { control.copy(); contextMenu.close() } enabled: control.selectedText && control.echoMode === TextInput.Normal }
        ItemDelegate { text: qsTr("Paste"); icon.source: "qrc:/images/resources/images/icons/paste-content.svg"; onClicked: { control.paste(); contextMenu.close() } enabled: control.canPaste }
        MenuSeparator {}
        ItemDelegate { text: qsTr("Select all"); icon.source: "qrc:/images/resources/images/icons/selectAll-content.svg"; onClicked: { control.selectAll(); contextMenu.close() } enabled: control.selectedText !== control.text }
        MenuSeparator {}
        ItemDelegate { text: qsTr("Undo"); icon.source: "qrc:/images/resources/images/icons/undo.svg"; onClicked: { control.undo(); contextMenu.close() } enabled: control.canUndo }
        ItemDelegate { text: qsTr("Redo"); icon.source: "qrc:/images/resources/images/icons/redo.svg"; onClicked: { control.redo(); contextMenu.close() } enabled: control.canRedo }
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