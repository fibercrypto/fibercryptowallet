import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12

Rectangle {
    id: root

    readonly property real spacing: 10
    property color backgroundColor: Material.primary
    property bool isInLeftSide: true
    property string leftText: "Left Text"
    property string rightText: "Right Text"

    anchors.centerIn: parent

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

        Behavior on anchors.leftMargin  { NumberAnimation { } }
        Behavior on anchors.rightMargin { NumberAnimation { } }

        text: "<b>" + (root.isInLeftSide ? root.leftText : root.rightText) + "</b>"
        highlighted: true

        onClicked: {
            root.isInLeftSide = !root.isInLeftSide
        }

        background: Rectangle {
            implicitWidth: 64
            implicitHeight: switchButton.Material.buttonHeight

            radius: root.radius
            color: !switchButton.enabled ? switchButton.Material.buttonDisabledColor :
                                        switchButton.highlighted ? switchButton.Material.highlightedButtonColor : switchButton.Material.buttonColor

            layer.enabled: switchButton.enabled && switchButton.Material.buttonColor.a > 0

            Rectangle {
                id: rectangleColorEffect
                anchors.fill: parent
                radius: root.radius

                color: switchButton.down ? "#33000000" : switchButton.hovered ? "#11000000" : "transparent"

                Behavior on color { ColorAnimation { } }
            }
        }
    }
}
