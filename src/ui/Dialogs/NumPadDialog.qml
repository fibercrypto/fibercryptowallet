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
    readonly property real numPadButtonImplicitSize: 50
    readonly property real numPadButtonPointSize: 18
    property alias pin: textInput.text

    function clear(newTitle) {
        textInput.clear()
        title = newTitle
        if (newTitle === "enter new pin:") {
            labelInstructions.text =
                    qsTr("Enter a hard-to-guess PIN of between 4 and 8 digits. " +
                         "The PIN layout is displayed in the hardware wallet screen.")
        } else if (newTitle === "confirm new pin:") {
            labelInstructions.text =
                    qsTr("If you forget the new PIN you will no able to work " +
                         "with this device configuretion.")
        } else if (newTitle === "enter current pin:") {
            labelInstructions.text =
                    qsTr("This device is currently protected with a PIN, to " +
                         "be able to work with it, you need to enter the PIN.")
        }
    }

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
                text: qsTr("instructions")
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
                    id: intValidator
                    bottom: 11111111
                    top: 99999999
                }
                echoMode: TextField.Password
                passwordCharacter: '#'

                inputMethodHints: Qt.ImhDigitsOnly
                Layout.fillWidth: true
                readOnly: true
                maximumLength: 8

                onTextChanged: {
                    standardButton(Dialog.Ok).enabled = textInput.length >= 4
                }
                function admitMoreDigits() {
                    return intValidator.top.toString().length - textInput.text.length > 0;
                }
                function toggleEnabledButtonsIfAny() {
                    gridNumPad.enabled = admitMoreDigits();
                    backspace.enabled = textInput.text.length > 0;
                }
                function addChar(ch) {
                    textInput.text += ch;
                    toggleEnabledButtonsIfAny();
                }
                function removeLastChar() {
                    textInput.text = textInput.text.slice(0, -1);
                    toggleEnabledButtonsIfAny();
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

                        ItemDelegate {
                            id: topLeft
                            text: numPadButtonText
                            font.pointSize: numPadButtonPointSize
                            implicitWidth: numPadButtonImplicitSize
                            implicitHeight: numPadButtonImplicitSize

                            Shortcut {
                                sequence: "7"
                                onActivated: {
                                    topLeft.clicked()
                                    topLeftAnimation.start()
                                }
                            }
                            onClicked: {
                                textInput.addChar('7');
                            }

                            SequentialAnimation {
                                id: topLeftAnimation
                                loops: 1
                                PropertyAction { target: topLeft; property: "downSym"; value: true }
                                PauseAnimation { duration: 350 }
                                PropertyAction { target: topLeft; property: "downSym"; value: false }
                            }

                            property bool downSym: down
                            property color color: (down || downSym) ? Material.color(Material.Amber) : hovered ? numPadDialog.Material.accent : numPadDialog.Material.foreground
                            Material.foreground: color

                            Behavior on color { ColorAnimation {} }
                        }

                        ItemDelegate {
                            id: topCenter
                            text: numPadButtonText
                            font.pointSize: numPadButtonPointSize
                            implicitWidth: numPadButtonImplicitSize
                            implicitHeight: numPadButtonImplicitSize

                            Shortcut {
                                sequence: "8"
                                onActivated: {
                                    topCenter.clicked()
                                    topCenterAnimation.start()
                                }
                            }
                            onClicked: {
                                textInput.addChar('8');
                            }

                            SequentialAnimation {
                                id: topCenterAnimation
                                loops: 1
                                PropertyAction { target: topCenter; property: "downSym"; value: true }
                                PauseAnimation { duration: 350 }
                                PropertyAction { target: topCenter; property: "downSym"; value: false }
                            }

                            property bool downSym: down
                            property color color: (down || downSym) ? Material.color(Material.Amber) : hovered ? numPadDialog.Material.accent : numPadDialog.Material.foreground
                            Material.foreground: color

                            Behavior on color { ColorAnimation {} }
                        }

                        ItemDelegate {
                            id: topRight
                            text: numPadButtonText
                            font.pointSize: numPadButtonPointSize
                            implicitWidth: numPadButtonImplicitSize
                            implicitHeight: numPadButtonImplicitSize

                            Shortcut {
                                sequence: "9"
                                onActivated: {
                                    topRight.clicked()
                                    topRightAnimation.start()
                                }
                            }
                            onClicked: {
                                textInput.addChar('9');
                            }

                            SequentialAnimation {
                                id: topRightAnimation
                                loops: 1
                                PropertyAction { target: topRight; property: "downSym"; value: true }
                                PauseAnimation { duration: 350 }
                                PropertyAction { target: topRight; property: "downSym"; value: false }
                            }

                            property bool downSym: down
                            property color color: (down || downSym) ? Material.color(Material.Amber) : hovered ? numPadDialog.Material.accent : numPadDialog.Material.foreground
                            Material.foreground: color

                            Behavior on color { ColorAnimation {} }
                        }

                        ItemDelegate {
                            id: centerLeft
                            text: numPadButtonText
                            font.pointSize: numPadButtonPointSize
                            implicitWidth: numPadButtonImplicitSize
                            implicitHeight: numPadButtonImplicitSize

                            Shortcut {
                                sequence: "4"
                                onActivated: {
                                    centerLeft.clicked()
                                    centerLeftAnimation.start()
                                }
                            }
                            onClicked: {
                                textInput.addChar('4');
                            }

                            SequentialAnimation {
                                id: centerLeftAnimation
                                loops: 1
                                PropertyAction { target: centerLeft; property: "downSym"; value: true }
                                PauseAnimation { duration: 350 }
                                PropertyAction { target: centerLeft; property: "downSym"; value: false }
                            }

                            property bool downSym: down
                            property color color: (down || downSym) ? Material.color(Material.Amber) : hovered ? numPadDialog.Material.accent : numPadDialog.Material.foreground
                            Material.foreground: color

                            Behavior on color { ColorAnimation {} }
                        }

                        ItemDelegate {
                            id: centerCenter
                            text: numPadButtonText
                            font.pointSize: numPadButtonPointSize
                            implicitWidth: numPadButtonImplicitSize
                            implicitHeight: numPadButtonImplicitSize

                            Shortcut {
                                sequence: "5"
                                onActivated: {
                                    centerCenter.clicked()
                                    centerCenterAnimation.start()
                                }
                            }
                            onClicked: {
                                textInput.addChar('5');
                            }

                            SequentialAnimation {
                                id: centerCenterAnimation
                                loops: 1
                                PropertyAction { target: centerCenter; property: "downSym"; value: true }
                                PauseAnimation { duration: 350 }
                                PropertyAction { target: centerCenter; property: "downSym"; value: false }
                            }

                            property bool downSym: down
                            property color color: (down || downSym) ? Material.color(Material.Amber) : hovered ? numPadDialog.Material.accent : numPadDialog.Material.foreground
                            Material.foreground: color

                            Behavior on color { ColorAnimation {} }
                        }

                        ItemDelegate {
                            id: centerRight
                            text: numPadButtonText
                            font.pointSize: numPadButtonPointSize
                            implicitWidth: numPadButtonImplicitSize
                            implicitHeight: numPadButtonImplicitSize

                            Shortcut {
                                sequence: "6"
                                onActivated: {
                                    centerRight.clicked()
                                    centerRightAnimation.start()
                                }
                            }
                            onClicked: {
                                textInput.addChar('6');
                            }

                            SequentialAnimation {
                                id: centerRightAnimation
                                loops: 1
                                PropertyAction { target: centerRight; property: "downSym"; value: true }
                                PauseAnimation { duration: 350 }
                                PropertyAction { target: centerRight; property: "downSym"; value: false }
                            }

                            property bool downSym: down
                            property color color: (down || downSym) ? Material.color(Material.Amber) : hovered ? numPadDialog.Material.accent : numPadDialog.Material.foreground
                            Material.foreground: color

                            Behavior on color { ColorAnimation {} }
                        }

                        ItemDelegate {
                            id: bottomLeft
                            text: numPadButtonText
                            font.pointSize: numPadButtonPointSize
                            implicitWidth: numPadButtonImplicitSize
                            implicitHeight: numPadButtonImplicitSize

                            Shortcut {
                                sequence: "1"
                                onActivated: {
                                    bottomLeft.clicked()
                                    bottomLeftAnimation.start()
                                }
                            }
                            onClicked: {
                                textInput.addChar('1');
                            }

                            SequentialAnimation {
                                id: bottomLeftAnimation
                                loops: 1
                                PropertyAction { target: bottomLeft; property: "downSym"; value: true }
                                PauseAnimation { duration: 350 }
                                PropertyAction { target: bottomLeft; property: "downSym"; value: false }
                            }

                            property bool downSym: down
                            property color color: (down || downSym) ? Material.color(Material.Amber) : hovered ? numPadDialog.Material.accent : numPadDialog.Material.foreground
                            Material.foreground: color

                            Behavior on color { ColorAnimation {} }
                        }

                        ItemDelegate {
                            id: bottomCenter
                            text: numPadButtonText
                            font.pointSize: numPadButtonPointSize
                            implicitWidth: numPadButtonImplicitSize
                            implicitHeight: numPadButtonImplicitSize

                            Shortcut {
                                sequence: "2"
                                onActivated: {
                                    bottomCenter.clicked()
                                    bottomCenterAnimation.start()
                                }
                            }
                            onClicked: {
                                textInput.addChar('2');
                            }

                            SequentialAnimation {
                                id: bottomCenterAnimation
                                loops: 1
                                PropertyAction { target: bottomCenter; property: "downSym"; value: true }
                                PauseAnimation { duration: 350 }
                                PropertyAction { target: bottomCenter; property: "downSym"; value: false }
                            }

                            property bool downSym: down
                            property color color: (down || downSym) ? Material.color(Material.Amber) : hovered ? numPadDialog.Material.accent : numPadDialog.Material.foreground
                            Material.foreground: color

                            Behavior on color { ColorAnimation {} }
                        }

                        ItemDelegate {
                            id: bottomRight
                            text: numPadButtonText
                            font.pointSize: numPadButtonPointSize
                            implicitWidth: numPadButtonImplicitSize
                            implicitHeight: numPadButtonImplicitSize

                            Shortcut {
                                sequence: "3"
                                onActivated: {
                                    bottomRight.clicked()
                                    bottomRightAnimation.start()
                                }
                            }
                            onClicked: {
                                textInput.addChar('3');
                            }

                            SequentialAnimation {
                                id: bottomRightAnimation
                                loops: 1
                                PropertyAction { target: bottomRight; property: "downSym"; value: true }
                                PauseAnimation { duration: 350 }
                                PropertyAction { target: bottomRight; property: "downSym"; value: false }
                            }

                            property bool downSym: down
                            property color color: (down || downSym) ? Material.color(Material.Amber) : hovered ? numPadDialog.Material.accent : numPadDialog.Material.foreground
                            Material.foreground: color

                            Behavior on color { ColorAnimation {} }
                        }
                    }

                    ItemDelegate {
                        id: backspace
                        icon.source: "qrc:/images/resources/images/icons/backspace.svg"
                        display: ItemDelegate.IconOnly
                        font.pointSize: numPadButtonPointSize
                        Layout.fillWidth: true
                        Layout.alignment: Qt.AlignCenter
                        enabled: false

                        Shortcut {
                            sequence: "Backspace"
                            onActivated: {
                                backspace.clicked()
                                backspaceAnimation.start()
                            }
                        }
                        onClicked: {
                            textInput.removeLastChar();
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
