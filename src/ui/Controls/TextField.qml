import QtQuick 2.12
import QtQuick.Controls 2.12

TextField {
	id: textField

	Menu {
		id: contextMenu

		focus: false

		ItemDelegate { text: qsTr("Cut"); onClicked: { textField.cut(); contextMenu.close() } enabled: textField.selectedText && textField.echoMode === TextInput.Normal }
		ItemDelegate { text: qsTr("Copy"); onClicked: { textField.copy(); contextMenu.close() } enabled: textField.selectedText && textField.echoMode === TextInput.Normal }
		ItemDelegate { text: qsTr("Paste"); onClicked: { textField.paste(); contextMenu.close() } enabled: textField.canPaste }
		MenuSeparator {}
		ItemDelegate { text: qsTr("Select all"); onClicked: { textField.selectAll(); contextMenu.close() } enabled: textField.selectedText !== textField.text }
		MenuSeparator {}
		ItemDelegate { text: qsTr("Undo"); onClicked: { textField.undo(); contextMenu.close() } enabled: textField.canUndo }
		ItemDelegate { text: qsTr("Redo"); onClicked: { textField.redo(); contextMenu.close() } enabled: textField.canRedo }
	}

	MouseArea {
		anchors.fill: parent
		acceptedButtons: Qt.RightButton
		cursorShape: Qt.IBeamCursor

		onClicked: {
			textField.focus = true
			contextMenu.popup()
		}
	}
}