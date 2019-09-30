import QtQuick 2.12
import QtQuick.Templates 2.12 as T
import QtQuick.Controls 2.12
import QtQuick.Controls.impl 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Controls.Material.impl 2.12

T.TextField {
    id: control

    implicitWidth: implicitBackgroundWidth + leftInset + rightInset
                   || Math.max(contentWidth, placeholder.implicitWidth) + leftPadding + rightPadding
    implicitHeight: Math.max(implicitBackgroundHeight + topInset + bottomInset,
                             contentHeight + topPadding + bottomPadding,
                             placeholder.implicitHeight + topPadding + bottomPadding)

    topPadding: 8
    bottomPadding: 16

    color: enabled ? Material.foreground : Material.hintTextColor
    selectionColor: Material.accentColor
    selectedTextColor: Material.primaryHighlightedTextColor
    placeholderTextColor: Material.hintTextColor
    verticalAlignment: TextInput.AlignVCenter

    cursorDelegate: CursorDelegate { }

    PlaceholderText {
        id: placeholder

        property bool floatPlaceholderText: !(!control.length && !control.preeditText && (!control.activeFocus || control.horizontalAlignment !== Qt.AlignHCenter))

        x: (floatPlaceholderText ? 0 : control.leftPadding) - width * (1-scale)/2
        // Behavior on x { NumberAnimation { duration: 150 } }
        y: floatPlaceholderText ? (-height*0.85 + control.topPadding*0.85) : control.topPadding
        Behavior on y { NumberAnimation { duration: 150; easing.type: Easing.OutQuint } }
        scale: floatPlaceholderText ? 0.85 : 1
        Behavior on scale { NumberAnimation { duration: 150; easing.type: Easing.OutQuint } }
        height: control.height - (control.topPadding + control.bottomPadding)
        text: control.placeholderText
        font: control.font
        color: floatPlaceholderText && control.activeFocus ? control.Material.accentColor : control.placeholderTextColor
        Behavior on color { ColorAnimation { duration: 150 } }
        verticalAlignment: control.verticalAlignment
        elide: Text.ElideRight
        renderType: control.renderType
    }

    background: Rectangle {
        y: control.height - height - control.bottomPadding + 8
        implicitWidth: 120
        height: control.activeFocus || control.hovered ? 2 : 1
        color: control.activeFocus ? control.Material.accentColor
                                   : (control.hovered ? control.Material.primaryTextColor : control.Material.hintTextColor)
    }

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
