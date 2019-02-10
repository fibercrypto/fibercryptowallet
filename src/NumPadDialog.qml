import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Dialog {
    id: root

    focus: true
    modal: true
    standardButtons: Dialog.Ok | Dialog.Cancel
    closePolicy: Popup.NoAutoClose

    implicitWidth: gridNumPad.width + 250
    implicitHeight: labelInstructions.height + textInput.height + numPad.height + 150

    readonly property string numPadButtonText: "#"
    readonly property real numPadButtonImplicitSize: 50
    readonly property real numPadButtonPointSize: 18

    Component.onCompleted: {
        standardButton(Dialog.Ok).enabled = textInput.length >= 4
    }

    ColumnLayout {
        anchors.fill: parent

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
                            console.log("Top Left clicked")
                        }

                        SequentialAnimation {
                            id: topLeftAnimation
                            loops: 1
                            PropertyAction { target: topLeft; property: "downSym"; value: true }
                            PauseAnimation { duration: 350 }
                            PropertyAction { target: topLeft; property: "downSym"; value: false }
                        }

                        property bool downSym: down
                        property color color: (down || downSym) ? Material.color(Material.Amber) : hovered ? root.Material.accent : root.Material.foreground
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
                            console.log("Top Center clicked")
                        }

                        SequentialAnimation {
                            id: topCenterAnimation
                            loops: 1
                            PropertyAction { target: topCenter; property: "downSym"; value: true }
                            PauseAnimation { duration: 350 }
                            PropertyAction { target: topCenter; property: "downSym"; value: false }
                        }

                        property bool downSym: down
                        property color color: (down || downSym) ? Material.color(Material.Amber) : hovered ? root.Material.accent : root.Material.foreground
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
                            console.log("Top Right clicked")
                        }

                        SequentialAnimation {
                            id: topRightAnimation
                            loops: 1
                            PropertyAction { target: topRight; property: "downSym"; value: true }
                            PauseAnimation { duration: 350 }
                            PropertyAction { target: topRight; property: "downSym"; value: false }
                        }

                        property bool downSym: down
                        property color color: (down || downSym) ? Material.color(Material.Amber) : hovered ? root.Material.accent : root.Material.foreground
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
                            console.log("Center Left clicked")
                        }

                        SequentialAnimation {
                            id: centerLeftAnimation
                            loops: 1
                            PropertyAction { target: centerLeft; property: "downSym"; value: true }
                            PauseAnimation { duration: 350 }
                            PropertyAction { target: centerLeft; property: "downSym"; value: false }
                        }

                        property bool downSym: down
                        property color color: (down || downSym) ? Material.color(Material.Amber) : hovered ? root.Material.accent : root.Material.foreground
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
                            console.log("Center Center clicked")
                        }

                        SequentialAnimation {
                            id: centerCenterAnimation
                            loops: 1
                            PropertyAction { target: centerCenter; property: "downSym"; value: true }
                            PauseAnimation { duration: 350 }
                            PropertyAction { target: centerCenter; property: "downSym"; value: false }
                        }

                        property bool downSym: down
                        property color color: (down || downSym) ? Material.color(Material.Amber) : hovered ? root.Material.accent : root.Material.foreground
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
                            console.log("Center Right clicked")
                        }

                        SequentialAnimation {
                            id: centerRightAnimation
                            loops: 1
                            PropertyAction { target: centerRight; property: "downSym"; value: true }
                            PauseAnimation { duration: 350 }
                            PropertyAction { target: centerRight; property: "downSym"; value: false }
                        }

                        property bool downSym: down
                        property color color: (down || downSym) ? Material.color(Material.Amber) : hovered ? root.Material.accent : root.Material.foreground
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
                            console.log("Bottom Left clicked")
                        }

                        SequentialAnimation {
                            id: bottomLeftAnimation
                            loops: 1
                            PropertyAction { target: bottomLeft; property: "downSym"; value: true }
                            PauseAnimation { duration: 350 }
                            PropertyAction { target: bottomLeft; property: "downSym"; value: false }
                        }

                        property bool downSym: down
                        property color color: (down || downSym) ? Material.color(Material.Amber) : hovered ? root.Material.accent : root.Material.foreground
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
                            console.log("Bottom Center clicked")
                        }

                        SequentialAnimation {
                            id: bottomCenterAnimation
                            loops: 1
                            PropertyAction { target: bottomCenter; property: "downSym"; value: true }
                            PauseAnimation { duration: 350 }
                            PropertyAction { target: bottomCenter; property: "downSym"; value: false }
                        }

                        property bool downSym: down
                        property color color: (down || downSym) ? Material.color(Material.Amber) : hovered ? root.Material.accent : root.Material.foreground
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
                            console.log("Bottom Right clicked")
                        }

                        SequentialAnimation {
                            id: bottomRightAnimation
                            loops: 1
                            PropertyAction { target: bottomRight; property: "downSym"; value: true }
                            PauseAnimation { duration: 350 }
                            PropertyAction { target: bottomRight; property: "downSym"; value: false }
                        }

                        property bool downSym: down
                        property color color: (down || downSym) ? Material.color(Material.Amber) : hovered ? root.Material.accent : root.Material.foreground
                        Material.foreground: color

                        Behavior on color { ColorAnimation {} }
                    }
                }

                ItemDelegate {
                    id: backspace
                    icon.source: "qrc:/images/icons/backspace.svg"
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
                    property color color: (down || downSym) ? Material.color(Material.Amber) : hovered ? root.Material.accent : root.Material.foreground
                    Material.foreground: color

                    Behavior on color { ColorAnimation {} }
                }
            }
        }

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

    }

}
