import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/Controls"
import "../Controls" // For quick UI development, switch back to resources when making a release

Dialog {
    id: numPadDialog

    standardButtons: Dialog.Ok | Dialog.Cancel
    closePolicy: Popup.NoAutoClose
    title: Qt.application.name

    readonly property string numPadButtonText: "#"
    readonly property real numPadButtonImplicitSideSize: 50
    readonly property real numPadButtonPointSize: 18

    Component.onCompleted: {
        standardButton(Dialog.Ok).enabled = textInput.length >= 4
    }

    onAboutToShow: {
        textInput.forceActiveFocus()
    }

    Flickable {
        id: flickable
        anchors.fill: parent
        contentHeight: columnLayoutRoot.height
        clip: true

        ColumnLayout {
            id: columnLayoutRoot
            width: parent.width

            Label {
                id: labelInstructions
                text: qsTr("Enter a hard-to-guess PIN of between 4 and 8 digits. " +
                        "The PIN layout is displayed in the hardware wallet screen.")
                Layout.fillWidth: true
                Layout.alignment: Qt.AlignHCenter
                horizontalAlignment: Label.AlignHCenter
                wrapMode: Label.WordWrap
            }

            TextField {
                id: textInput

                placeholderText: qsTr("12345678")
                selectByMouse: true
                validator: IntValidator {
                    bottom: 11111111
                    top: 99999999
                }

                inputMethodHints: Qt.ImhDigitsOnly
                Layout.fillWidth: true
                readOnly: true
                maximumLength: 8

                onTextChanged: {
                    standardButton(Dialog.Ok).enabled = textInput.length >= 4
                }
            }

            Rectangle {
                id: numPad
                Layout.alignment: Qt.AlignHCenter
                border.width: 1
                border.color: "#20000000"
                color: Material.background
                implicitWidth: Math.max(gridNumPad.implicitWidth, backspace.implicitWidth) + 16
                implicitHeight: gridNumPad.implicitHeight + backspace.implicitHeight + 20

                ColumnLayout {
                    id: columnLayoutNumPad
                    anchors.centerIn: parent

                    GridLayout {
                        id: gridNumPad
                        Layout.alignment: Qt.AlignCenter
                        columns: 3

                        readonly property var sequences: [ 7, 8, 9, 4, 5, 6, 1, 2, 3 ]
                        readonly property var positionsStrings: [ "Top left", "Top center", "Top right", "Center left", "Center", "Center right", "Bottom left", "Bottom center", "Bottom right" ]

                        Repeater {
                            model: 9

                            delegate: ItemDelegate {
                                id: itemDelegateNumPadButton

                                text: numPadButtonText
                                font.pointSize: numPadButtonPointSize
                                implicitWidth: numPadButtonImplicitSideSize
                                implicitHeight: numPadButtonImplicitSideSize

                                Shortcut {
                                    sequence: gridNumPad.sequences[index].toString()
                                    enabled: numPadDialog.visible
                                    onActivated: {
                                        itemDelegateNumPadButton.clicked()
                                        numPadButtonAnimation.start()
                                    }
                                }
                                onClicked: {
                                    console.log(gridNumPad.positionsStrings[index], "clicked")
                                }

                                SequentialAnimation {
                                    id: numPadButtonAnimation
                                    loops: 1
                                    PropertyAction { target: itemDelegateNumPadButton; property: "downSym"; value: true }
                                    PauseAnimation { duration: 350 }
                                    PropertyAction { target: itemDelegateNumPadButton; property: "downSym"; value: false }
                                }

                                property bool downSym: down
                                property color color: (down || downSym) ? Material.color(Material.Amber) : hovered ? numPadDialog.Material.accent : numPadDialog.Material.foreground
                                Material.foreground: color

                                Behavior on color { ColorAnimation {} }
                            } // ItemDelegate (delegate)
                        } // Repeater
                    } // GridLayout

                    ItemDelegate {
                        id: backspace
                        icon.source: "qrc:/images/resources/images/icons/backspace.svg"
                        display: ItemDelegate.IconOnly
                        font.pointSize: numPadButtonPointSize
                        Layout.fillWidth: true
                        Layout.alignment: Qt.AlignCenter

                        Shortcut {
                            sequence: "Backspace"
                            onActivated: {
                                backspace.clicked()
                                backspaceAnimation.start()
                            }
                        }
                        onClicked: {
                            console.log("Backspace clicked")
                        }

                        SequentialAnimation {
                            id: backspaceAnimation
                            loops: 1
                            PropertyAction { target: backspace; property: "downSym"; value: true }
                            PauseAnimation { duration: 350 }
                            PropertyAction { target: backspace; property: "downSym"; value: false }
                        }

                        property bool downSym: down
                        property color color: (down || downSym) ? Material.color(Material.Amber) : hovered ? numPadDialog.Material.accent : numPadDialog.Material.foreground
                        Material.foreground: color

                        Behavior on color { ColorAnimation {} }
                    } // GridLayout (numpad)
                } // ColumnLayout (numpad)
            } // Rectangle (numpad)

            Label {
                id: labelInstructionsNotes
                text: qsTr("<b>NOTE:</b> You can use the numpad, but the number entered " +
                        "will be the corresponding number in the hardware wallet screen.")
                font.pointSize: 9
                font.italic: true

                Layout.fillWidth: true
                wrapMode: Label.WordWrap
                color: Material.color(Material.Red)
            }

        } // ColumnLayout (root)

        ScrollIndicator.vertical: ScrollIndicator {
            parent: msgDialog.contentItem
            anchors.top: flickable.top
            anchors.bottom: flickable.bottom
            anchors.right: parent.right
            anchors.rightMargin: -msgDialog.rightPadding + 1
        }
    } // Flickable

}
