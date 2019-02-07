import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.impl 2.12
import QtQuick.Controls.Material 2.12

// Colors are only implemented for Material Style
// TODO: Implement the same colors for the other styles

Rectangle {
    id: root

    readonly property real spacing: 10
    property color backgroundColor: Material.primary
    property bool isInLeftSide: true
    property string leftText: "Left Text"
    property string rightText: "Right Text"
    property color leftColor: Material.color(Material.Blue)
    property color rightColor: Material.color(Material.Green)
    property int animationsSpeed: 150
    property color textColor: Material.foreground

    width: 300
    height: 70

    radius: 40
    color: backgroundColor

    Button {
        id: switchButton
        clip: true

        anchors.top: parent.top
        anchors.bottom: parent.bottom
        anchors.left: parent.left
        anchors.leftMargin: root.isInLeftSide ? spacing : root.width/2 + spacing*2
        anchors.right: parent.right
        anchors.rightMargin: root.isInLeftSide ? root.width/2 + spacing*2 : spacing

        Material.foreground: textColor

        Behavior on anchors.leftMargin  { NumberAnimation { duration: animationsSpeed; easing.type: Easing.OutQuint } }
        Behavior on anchors.rightMargin { NumberAnimation { duration: animationsSpeed; easing.type: Easing.OutQuint } }

        text: "<b>" + (root.isInLeftSide ? root.leftText : root.rightText) + "</b>"

        onClicked: {
            root.isInLeftSide = !root.isInLeftSide
        }

        contentItem: IconLabel {
            spacing: switchButton.spacing
            mirrored: switchButton.mirrored
            display: switchButton.display

            icon: switchButton.icon
            text: switchButton.text
            font: switchButton.font
            color: !switchButton.enabled ? switchButton.Material.hintTextColor : switchButton.Material.foreground
        }

        background: Rectangle {
            implicitWidth: 64
            implicitHeight: switchButton.Material.buttonHeight

            radius: root.radius
            color: !switchButton.enabled ? switchButton.Material.buttonDisabledColor : root.isInLeftSide ? root.leftColor : root.rightColor
            Behavior on color { ColorAnimation { duration: animationsSpeed } }

            layer.enabled: switchButton.enabled && switchButton.Material.buttonColor.a > 0

            Rectangle {
                id: rectangleColorEffect
                anchors.fill: parent
                radius: root.radius

                color: switchButton.down ? "#33000000" : switchButton.hovered ? "#11000000" : "transparent"

                Behavior on color { ColorAnimation { duration: animationsSpeed } }
            }
        }
    }
}
