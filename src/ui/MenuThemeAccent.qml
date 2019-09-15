import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12

Menu {
    id: menuThemeAccent

    readonly property int currentTheme: applicationWindow.Material.theme
    readonly property color currentAccent: applicationWindow.Material.accent
    property int currentSelectedIndex: 5 // Material.Blue

    title: qsTr("Style configuration")

    readonly property var materialPredefinedColors: [
        Material.Red,
        Material.Pink,
        Material.Purple,
        Material.DeepPurple,
        Material.Indigo,
        Material.Blue,
        Material.LightBlue,
        Material.Cyan,
        Material.Teal,
        Material.Green,
        Material.LightGreen,
        Material.Lime,
        Material.Yellow,
        Material.Amber,
        Material.Orange,
        Material.DeepOrange,
        Material.Brown,
        Material.Grey,
        Material.BlueGrey
    ]

    width: gridAccents.width
        
    SwitchDelegate {
        id: switchDelegateTheme

        width: parent.width
        text: qsTr("Night theme")
        checked: currentTheme === Material.Dark

        onCheckedChanged: {
            applicationWindow.flash()
            applicationWindow.Material.theme = (currentTheme === Material.Light ? Material.Dark : Material.Light)
            applicationWindow.Material.accent = materialPredefinedColors[currentSelectedIndex]
        }
    }

    MenuSeparator {}

    Grid {
        id: gridAccents

        property int markedIndex: 5

        columns: 4
        columnSpacing: 10
        rowSpacing: 10
        leftPadding: 10
        rightPadding: leftPadding
        bottomPadding: 4

        Repeater {
            model: 19

            delegate: Rectangle {
                id: rectangleDelegate

                width: 48
                height: width
                radius: width/2
                color: Material.accent
                border { width: 3; color: Qt.darker(rectangleDelegate.color) }
                Material.accent: materialPredefinedColors[index]

                Rectangle {
                    id: rectangleCheckIndicator

                    property bool checked: applicationWindow.Material.accent === Material.accent

                    anchors.centerIn: parent
                    width: rectangleDelegate.width - 16
                    height: width
                    radius: width/2
                    color: rectangleDelegate.border.color
                    opacity: checked ? 1.0 : 0.0
                    Behavior on opacity { NumberAnimation { duration: 150; easing.type: Easing.OutQuint } }
                    scale: 0.5 + 0.5*opacity
                }

                ToolButton {
                    id: toolButtonChangeAccent
                    
                    anchors.fill: parent
                    anchors.margins: -6
                    z: 10

                    onClicked: {
                        currentSelectedIndex = index
                        applicationWindow.Material.accent = Material.accent
                    }
                } // ToolButton
            } // Rectangle
        } // Repeater
    } // Grid
}