import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import WalletsManager 1.0
import Config 1.0

// Resource imports
// import "qrc:/ui/src/ui/"
// import "qrc:/ui/src/ui/Dialogs"
import "Dialogs/" // For quick UI development, switch back to resources when making a release

ApplicationWindow {
    id: applicationWindow
    
    property bool skipAccentColorAnimation: false
    property bool accentColorAnimationActive: false
    onAccentColorAnimationActiveChanged: console.log(accentColorAnimationActive)
    property color accentColor: Material.accent
    Behavior on accentColor {
        SequentialAnimation {
            PropertyAction { target: applicationWindow; property: "accentColorAnimationActive"; value: true }
            ColorAnimation { duration: skipAccentColorAnimation ? 0 : 200 }
            PropertyAction { target: applicationWindow; property: "accentColorAnimationActive"; value: false }
        }
    }

    visible: true
    width: 680
    height: 580
    title: Qt.application.name + ' v' + Qt.application.version
    Material.accent: accentColor
   
    function flash() {
        flasher.flash()
    }

    menuBar: CustomMenuBar {
        id: customMenuBar

        onOutputsRequested: {
            generalStackView.openOutputsPage()
            customHeader.text = qsTr("Outputs")
            
            enableOutputs = false
            enablePendingTransactions = true
            enableBlockchain = true
            enableNetworking = true
            enableSettings = true
        }
        ConfigManager{
            id: configManager
        }

        onPendingTransactionsRequested: {
            generalStackView.openPendingTransactionsPage()
            customHeader.text = qsTr("Pending transactions")
            
            enableOutputs = true
            enablePendingTransactions = false
            enableBlockchain = true
            enableNetworking = true
            enableSettings = true

        }

        onBlockchainRequested: {
            generalStackView.openBlockchainPage()
            customHeader.text = qsTr("Blockchain")

            enableOutputs = true
            enablePendingTransactions = true
            enableBlockchain = false
            enableNetworking = true
            enableSettings = true
        }

        onNetworkingRequested: {
            generalStackView.openNetworkingPage()
            customHeader.text = qsTr("Networking")

            enableOutputs = true
            enablePendingTransactions = true
            enableBlockchain = true
            enableNetworking = false
            enableSettings = true
        }

        onSettingsRequested: {
            generalStackView.openSettingsPage()
            customHeader.text = qsTr("Settings")

            enableOutputs = true
            enablePendingTransactions = true
            enableBlockchain = true
            enableNetworking = true
            enableSettings = false
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
        //property WalletManager  walletManger: WalletManager{
        //id: walletManager
        //}
        WalletManager{
            id: walletManager
        }
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

    // QR
    DialogQR {
        id: dialogQR
        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 540 ? 540 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > 570 ? 570 - 40 : applicationWindow.height - 40

        focus: true
        modal: true
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

    NumPadDialog {
        id: numPadDialog
        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 440 ? 440 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > 540 ? 540 - 40 : applicationWindow.height - 40

        focus: true
        modal: true
    }

    RestoreBackupWordsDialog {
        id: restoreBackupWordsDialog
        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 460 ? 460 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > 340 ? 340 - 40 : applicationWindow.height - 40

        focus: true
        modal: true
    }

    SecureWalletDialog {
        id: secureWalletDialog
        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 640 ? 640 - 40 : applicationWindow.width - 40
        height: (applicationWindow.height > 590 ? 590 - 40 : applicationWindow.height - 40) - (enableBackupWarning ^ enablePINWarning ? 100 : 0) - (!enableBackupWarning && !enablePINWarning ? 240 : 0)
        
        focus: true
        modal: true
    }

   

    WalletCreatedDialog {
        id: walletCreatedDialog
        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 540 ? 540 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > 360 ? 360 - 40 : applicationWindow.height - 40
        
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

    //! This must be the last object (i.e. the one with the greater `z` value)
    Rectangle {
        id: flasher

        property int duration: 500

        function flash() {
            if (flashAnimation.running) {
                flashAnimation.restart()
            } else {
                flashAnimation.start()
            }
        }

        y: -customMenuBar.height
        width: applicationWindow.width
        height: applicationWindow.height
        color: "white"
        opacity: 0.0
        z: customMenuBar.z + 1

        NumberAnimation {
            id: flashAnimation

            target: flasher
            property: "opacity"
            from: 1.0; to: 0.0
            duration: flasher.duration
            easing.type: Easing.InQuad
        }
    }
}
