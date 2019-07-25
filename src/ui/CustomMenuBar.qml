import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

RowLayout {
    spacing: 0

    property alias backButtonHide: toolButtonBack.hide
    property alias toolButtonBack: toolButtonBack
    property alias menuBarReal: menuBarReal
    property alias menuBarColor: menuBarReal.color

    // Signals
    signal outputsRequested()
    signal networkingRequested()
    signal blockchainRequested()
    signal aboutRequested()
    signal aboutQtRequested()

    // Functions
    function back() {
        if (!toolButtonBack.hide) {
            toolButtonBack.clicked()
        }
    }

    ToolButton {
        id: toolButtonBack

        property bool hide: generalStackView.depth === 1 || (generalStackView.depth >= 1 && generalStackView.busy)

        focusPolicy: Qt.NoFocus

        // opacity
        opacity: hide ? 0 : 1
        Behavior on opacity { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } }

        // position
        implicitHeight: menuBarReal.height
        implicitWidth: implicitHeight
        Layout.leftMargin: hide ? -(width + padding) : 0
        Behavior on Layout.leftMargin { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } }
        z: 1

        // icon
        icon.source: "qrc:/images/resources/images/icons/back.svg"
        icon.color: "white"

        // PointingHandCursor
        MouseArea {
            anchors.fill: parent
            acceptedButtons: Qt.NoButton
            cursorShape: Qt.PointingHandCursor
        }

        onClicked: {
            generalStackView.pop()
        }
    }

    MenuBar {
        id: menuBarReal

        readonly property color iconColor: "transparent"
        property color color: Material.dialogColor
        property color menuTextColor: toolButtonBack.hide ? Material.primaryTextColor : "white"

        Layout.fillWidth: true
        leftInset:  -(toolButtonBack.width + toolButtonBack.padding)
        Material.foreground: menuTextColor
        Behavior on menuTextColor { ColorAnimation { } }

        background: Rectangle {
            id: backgroundRectangle
            implicitHeight: 40
            color: toolButtonBack.hide ? Material.dialogColor : menuBarReal.color

            Behavior on color { ColorAnimation { } }
        }

        Menu {
            id: menuTools
            title: qsTr("&Tools")
            MenuItem {
                text: qsTr("&Outputs")
                leftPadding: highlighted ? padding * 1.5 : padding
                Behavior on leftPadding { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } }

                onClicked: outputsRequested()
            }
            MenuItem {
                text: qsTr("&Blockchain")
                leftPadding: highlighted ? padding * 1.5 : padding
                Behavior on leftPadding { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } }

                onClicked: blockchainRequested()
            }
            MenuItem {
                text: qsTr("&Networking")
                leftPadding: highlighted ? padding * 1.5 : padding
                Behavior on leftPadding { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } }

                onClicked: networkingRequested()
            }
        } // menuTools
        Menu {
            id: menuHelp
            title: qsTr("&Help")
            MenuItem {
                text: qsTr("&About FiberCrypto")
                leftPadding: highlighted ? padding * 1.5 : padding
                Behavior on leftPadding { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } }

                onClicked: aboutRequested()
            }
            MenuItem {
                text: qsTr("About &Qt")
                leftPadding: highlighted ? padding * 1.5 : padding
                Behavior on leftPadding { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } }

                onClicked: aboutQtRequested()
            }
        } // menuHelp
    } // menuBarReal
} // RowLayout (menuBar)
