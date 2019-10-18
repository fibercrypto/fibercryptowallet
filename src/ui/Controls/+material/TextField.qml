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

    topPadding: 16
    bottomPadding: 9

    color: enabled ? Material.foreground : Material.hintTextColor
    selectionColor: Material.accentColor
    selectedTextColor: Material.primaryHighlightedTextColor
    placeholderTextColor: Material.hintTextColor
    verticalAlignment: TextInput.AlignVCenter

    cursorDelegate: CursorDelegate { }

    PlaceholderText {
        id: placeholder

        property bool floatPlaceholderText: !(!control.length && !control.preeditText && (!control.activeFocus || control.horizontalAlignment !== Qt.AlignHCenter))
        readonly property real placeholderTextScaleFactor: 0.9

        x: (floatPlaceholderText ? 0 : control.leftPadding) - width * (1-scale)/2
        y: floatPlaceholderText ? -control.topPadding*(1 - placeholderTextScaleFactor) : control.topPadding
        Behavior on y { NumberAnimation { duration: 250; easing.type: Easing.OutQuint } }
        scale: floatPlaceholderText ? placeholderTextScaleFactor : 1
        Behavior on scale { NumberAnimation { duration: 250; easing.type: Easing.OutQuint } }
        height: control.height - (control.topPadding + control.bottomPadding)
        text: control.placeholderText
        color: floatPlaceholderText && control.activeFocus ? control.Material.accentColor : control.placeholderTextColor
        Behavior on color { ColorAnimation { duration: 250 } }
        verticalAlignment: control.verticalAlignment
        elide: Text.ElideRight
        renderType: control.renderType
    }

    background: Rectangle {
        y: control.height - height - control.bottomPadding + 8
        implicitWidth: 120
        height: 1
        color: control.hovered ? control.Material.primaryTextColor : control.Material.hintTextColor
        Behavior on color { ColorAnimation { duration: 200 } }

        Rectangle {
            id: accentRect

            readonly property bool controlHasActiveFocus: control.activeFocus

            onControlHasActiveFocusChanged: {
                if (controlHasActiveFocus) {
                    animationOnActiveFocus.start()
                } else {
                    animationOnUnactiveFocus.start()
                }
            }

            y: parent.y
            implicitWidth: 120
            // width: 
            Component.onCompleted: width = control.activeFocus ? parent.width : 0
            height: 2
            anchors.centerIn: parent
            color: control.Material.accentColor

            NumberAnimation {
                id: animationOnActiveFocus

                target: accentRect
                property: "width"
                to: accentRect.parent.width
                duration: 350
                easing.type: Easing.OutQuint
            }


            NumberAnimation {
                id: animationOnUnactiveFocus

                target: accentRect
                property: "opacity"
                to: 0.0
                duration: 350
                easing.type: Easing.OutQuint

                onFinished: {
                    target.width = 0
                    target.opacity = 1.0
                }
            }
        } // Rectangle (accent)
    } // Ractangle (decoration)

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
