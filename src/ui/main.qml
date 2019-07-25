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
            if (false) {
                menuBarColor = Material.color(Material.Teal)
                customHeader.text = qsTr("Outputs")
                
                enableOutputs = false
                enableBlockchain = true
                enableNetworking = true
            }
        }

        onBlockchainRequested: {
            generalStackView.openBlockchainPage()
            menuBarColor = Material.color(Material.Red)
            customHeader.text = qsTr("Blockchain")

            enableOutputs = true
            enableBlockchain = false
            enableNetworking = true
        }

        onNetworkingRequested: {
            generalStackView.openNetworkingPage()
            menuBarColor = Material.color(Material.Blue)
            customHeader.text = qsTr("Networking")

            enableOutputs = true
            enableBlockchain = true
            enableNetworking = false
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

    // Message dialogs

    MsgDialog {
        id: msgDialog
        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 440 ? 440 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > 280 ? 280 - 40 : applicationWindow.height - 40

        focus: true
        modal: true

        text: qsTr("Lorem ipsum dolor sit amet, consectetur adipiscing elit, "
                 + "sed eiusmod tempor incidunt ut labore et dolore magna aliqua. "
                 + "Ut enim ad minim veniam, quis nostrud exercitation ullamco laboris "
                 + "nisi ut aliquid ex ea commodi consequat. "
                 + "Quis aute iure reprehenderit in voluptate velit esse cillum dolore "
                 + "eu fugiat nulla pariatur. Excepteur sint obcaecat cupiditat non proident, "
                 + "sunt in culpa qui officia deserunt mollit anim id est laborum.")
    }

    // Hardware dialogs

    DialogUnconfiguredWallet {
        id: dialogUnconfiguredWallet
        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 690 ? 690 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > 420 ? 420 - 40 : applicationWindow.height - 40

        focus: true
        modal: true
    }


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
