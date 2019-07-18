import QtQuick 2.12
import QtQuick.Controls 2.12

// Resource imports
import "qrc:/qml"

ApplicationWindow {
    visible: true
    width: 680
    height: 580
    title: qsTr("FiberCrypto Wallet")

    menuBar: MenuBar {
        clip: true
        height: general.toolPageOpened ? 0 : implicitHeight
        Behavior on height { NumberAnimation { duration: 200; easing.type: Easing.OutCubic } }
        Menu {
            id: menuTools
            title: qsTr("&Tools")
            MenuItem { text: qsTr("Outputs") }
            MenuItem { text: qsTr("Blockchain"); onClicked: general.openBlockchainPage() }
        }
        Menu {
            id: menuHelp
            title: qsTr("&Help")
            MenuItem { text: qsTr("About FiberCripto") }
            MenuItem { text: qsTr("About Qt"); onClicked: dialogAboutQt.open() }
        }
    }

    GeneralStackView {
        id: general
        anchors.fill: parent
    }

    // About
    DialogAbout {
        id: dialogAbout
        anchors.centerIn: Overlay.overlay
        modal: true
    }
    DialogAboutQt {
        id: dialogAboutQt
        anchors.centerIn: Overlay.overlay
        modal: true
    }
}
