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
    onRejected: {
        bridgeForPassword.errMessage = "Action canceled";
        bridgeForPassword.unlock()
    }
    onAccepted: {
        bridgeForPassword.setResult(textInput.text)
        bridgeForPassword.unlock()
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
                                    textInput.addChar(gridNumPad.sequences[index].toString());
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
            parent: numPadDialog.contentItem
            anchors.top: flickable.top
            anchors.bottom: flickable.bottom
            anchors.right: parent.right
            anchors.rightMargin: -numPadDialog.rightPadding + 1
        }
    } // Flickable
}
