import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12

// Resource imports
// import "qrc:/ui/src/ui/"
// import "qrc:/ui/src/ui/Dialogs"
import "Dialogs/" // For quick UI development, switch back to resources when making a release

ApplicationWindow {
    id: applicationWindow
    visible: true
    width: 680
    height: 580
    title: Qt.application.name + ' v' + Qt.application.version

    menuBar: CustomMenuBar {
        id: customMenuBar

        onOutputsRequested: {
            // ...
            menuBarColor = Material.color(Material.Teal)
            customHeader.text = qsTr("Outputs")
        }

        onBlockchainRequested: {
            generalStackView.openBlockchainPage()
            menuBarColor = Material.color(Material.Red)
            customHeader.text = qsTr("Blockchain")
        }

        onNetworkingRequested: {
            generalStackView.openNetworkingPage()
            menuBarColor = Material.color(Material.Blue)
            customHeader.text = qsTr("Networking")
        }

        onAboutRequested: {
            dialogAbout.open()
        }

        onAboutQtRequested: {
            dialogAboutQt.open()
        }
    } // CustomMenuBar

    CustomHeader {
        id: customHeader
    } // CustomHeader

    GeneralStackView {
        id: generalStackView
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
