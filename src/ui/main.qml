import QtQuick 2.12
import QtQuick.Controls 2.12

// Resource imports
import "qrc:/ui/src/ui/"
import "qrc:/ui/src/ui/Dialogs"

ApplicationWindow {
    visible: true
    width: 680
    height: 580
    title: Qt.application.name + ' v' + Qt.application.version

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
            MenuItem { text: qsTr("About FiberCrypto"); onClicked: dialogAbout.open() }
            MenuItem { text: qsTr("About Qt"); onClicked: dialogAboutQt.open() }
        }
    }

    GeneralStackView {
        id: general
        anchors.fill: parent
    }

    //! Dialogs

    // Help dialogs

    DialogAbout {
        id: dialogAbout

        readonly property real minimumParentSideSize: Math.min(parent.width, parent.height)

        parent: Overlay.overlay
        anchors.centerIn: parent
        width: (minimumParentSideSize / 3) * 2
        height: (parent.height / 3) * 2
    }

    DialogAboutQt {
        id: dialogAboutQt

        readonly property real minimumParentSideSize: Math.min(parent.width, parent.height)

        parent: Overlay.overlay
        anchors.centerIn: parent
        width: (minimumParentSideSize / 3) * 2
        height: (parent.height / 3) * 2
    }
}
