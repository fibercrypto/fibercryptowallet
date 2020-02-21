import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Window 2.12
import Qt.labs.settings 1.0
import WalletsManager 1.0
import Config 1.0
import Utils 1.0
import DeviceInteraction 1.0

// Resource imports
// import "qrc:/ui/src/ui/Dialogs"
import "Dialogs/" // For quick UI development, switch back to resources when making a release

ApplicationWindow {
    id: applicationWindow
    
    property bool skipAccentColorAnimation: false
    property bool accentColorAnimationActive: false
    property color accentColor: settings.value("style/material/accent", Material.accent)
    Behavior on accentColor {
        SequentialAnimation {
            PropertyAction { target: applicationWindow; property: "accentColorAnimationActive"; value: true }
            ColorAnimation { duration: skipAccentColorAnimation ? 0 : 200 }
            PropertyAction { target: applicationWindow; property: "accentColorAnimationActive"; value: false }
        }
    }

    width: 680
    height: 580
    title: Qt.application.name + ' v' + Qt.application.version
    Material.theme: ~~settings.value("style/material/theme", Material.Light)
    Material.accent: accentColor

    function flash() {
        flasher.flash()
    }

    menuBar: CustomMenuBar {
        id: customMenuBar

        ConfigManager{
            id: configManager
        }

        onOutputsRequested: {
            generalStackView.openOutputsPage()
            customHeader.text = qsTr("Outputs")
            
            enableOutputs = false
            enablePendingTransactions = true
            enableBlockchain = true
            enableNetworking = true
            enableSettings = true
            enableSettingsAddressBook = false
            enableAddrsBook = true
        }
        
        onPendingTransactionsRequested: {
            generalStackView.openPendingTransactionsPage()
            customHeader.text = qsTr("Pending transactions")
            
            enableOutputs = true
            enablePendingTransactions = false
            enableBlockchain = true
            enableNetworking = true
            enableSettings = true
            enableSettingsAddressBook = false
            enableAddrsBook = true

        }

        onBlockchainRequested: {
            generalStackView.openBlockchainPage()
            customHeader.text = qsTr("Blockchain")

            enableOutputs = true
            enablePendingTransactions = true
            enableBlockchain = false
            enableNetworking = true
            enableSettings = true
            enableSettingsAddressBook = false
            enableAddrsBook = true

        }

        onNetworkingRequested: {
            generalStackView.openNetworkingPage()
            customHeader.text = qsTr("Networking")

            enableOutputs = true
            enablePendingTransactions = true
            enableBlockchain = true
            enableNetworking = false
            enableSettings = true
            enableSettingsAddressBook = false
            enableAddrsBook = true
        }

        onAddressBookRequested: {
            generalStackView.openAddressBookPage()
            customHeader.text = qsTr("Address book")

            enableOutputs = true
            enablePendingTransactions = true
            enableBlockchain = true
            enableNetworking = true
            enableSettings = true
            enableSettingsAddressBook = true
            enableAddrsBook = false
        }

        onSettingsRequested: {
            generalStackView.openSettingsPage()
            customHeader.text = qsTr("Settings")

            enableOutputs = true
            enablePendingTransactions = true
            enableBlockchain = true
            enableNetworking = true
            enableSettings = false
            enableSettingsAddressBook = false
            enableAddrsBook = true
        }

        onSettingsAddressBookRequested: {
            generalStackView.openSettingsAddressBookPage()
            customHeader.text = qsTr("Address Book Settings")

            // The back button must be used to go back to the Address Book
            enableOutputs = enablePendingTransactions = enableBlockchain = enableNetworking = enableSettings = enableSettingsAddressBook = enableAddrsBook = false
        }

        onAboutRequested: {
            dialogAbout.open()
        }

        onAboutQtRequested: {
            dialogAboutQt.open()
        }

        onLicenseRequested: {
            dialogAboutLicense.open()
        }
    } // CustomMenuBar

    Action {
        id: actionFullScreen

        property int previous: applicationWindow.visibility

        shortcut: StandardKey.FullScreen
        onTriggered: {
            if (applicationWindow.visibility !== Window.FullScreen) {
                previous = applicationWindow.visibility
            }
            if (applicationWindow.visibility === Window.FullScreen) {
                applicationWindow.showNormal() // Cannot show maximized directly due to a bug in some X11 managers
                if (previous === Window.Maximized) {
                    applicationWindow.showMaximized()
                }
            } else {
                applicationWindow.showFullScreen()
            }
        }
    }

    CustomHeader {
        id: customHeader
    }

    GeneralStackView {
        id: generalStackView
        anchors.fill: parent

        onBackRequested: {
            customMenuBar.back()
        }
       
        WalletManager {
            id: walletManager
        }
    }

    //! Settings
    Settings {
        id: settings
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
    MsgDialog {
        id: requestCancel
        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 440 ? 440 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > 280 ? 280 - 40 : applicationWindow.height - 40
        focus: true
        modal: true
        standardButtons: Dialog.Cancel
        onRejected: {
            deviceInteraction.cancelCommand()
        }
    }
    MsgDialog {
        id: requestConfirmation
        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 440 ? 440 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > 280 ? 280 - 40 : applicationWindow.height - 40
        focus: true
        modal: true
        standardButtons: Dialog.Ok | Dialog.Cancel
        onRejected: {
            deviceInteraction.cancelCommand()
        }
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

    DeviceInteraction {
        id: deviceInteraction
        onInitializeDevice: {
            dialogUnconfiguredWallet.open();
        }
        onSecureDevice: {
            dialogSkyWalletId.open();
        }
        onOpenInteractionDialog: {
            skyWalletInteractionDialog.open();
        }
    }
    QBridge {
        id: topLevelDialogLocker
    }
    QBridge {
        id: bridgeForPassword

        onGetPassword: {
            getPasswordDialog.title = message;
            getPasswordDialog.clear();
            getPasswordDialog.open();
        }
        onDeviceRequireAction: {
            msgDialog.title = title;
            msgDialog.text = message;
            msgDialog.open();
        }
        onDeviceRequireConfirmableAction: {
            requestConfirmation.title = title;
            requestConfirmation.text = message;
            requestConfirmation.open();
        }
        onDeviceRequireCancelableAction: {
            requestCancel.title = title;
            requestCancel.text = message;
            requestCancel.open();
        }
        onGetSkyHardwareWalletPin: {
            numPadDialog.clear(title);
            numPadDialog.open();
        }
        onGetBip39Word: {
            bip39WordDialog.clear(title, message);
            bip39WordDialog.open();
        }
        Component.onCompleted: {
            bridgeForPassword.onCompleted();
        }
    }

    DialogGetPassword{
        id: getPasswordDialog
        anchors.centerIn: Overlay.overlay
        property int nAddress
        width: applicationWindow.width > 540 ? 540 - 120 : applicationWindow.width - 40
        height: applicationWindow.height > 570 ? 570 - 180 : applicationWindow.height - 40

        focus: true
        modal: true
        onClosed:{
            bridgeForPassword.setResult(getPasswordDialog.password)
            bridgeForPassword.unlock()
        }
    }

    NumPadDialog {
        id: numPadDialog
        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 440 ? 440 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > 540 ? 540 - 40 : applicationWindow.height - 40

        focus: visible
        modal: true
        onClosed: {
            msgDialog.close()
        }
    }
    
    DialogGetBip39Word {
        id: bip39WordDialog
        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 440 ? 440 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > 360 ? 360 - 40 : applicationWindow.height - 40

        focus: true
        modal: true
    }

    SkyWalletInteractionDialog {
        id: skyWalletInteractionDialog
        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 440 ? 440 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > 460 ? 460 - 40 : applicationWindow.height - 40

        focus: visible
        modal: true
    }

    SecureWalletDialog {
        id: dialogSkyWalletId
        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 640 ? 640 - 40 : applicationWindow.width - 40
        height: (applicationWindow.height > 590 ? 590 - 40 : applicationWindow.height - 40) - (enableBackupWarning ^ enablePINWarning ? 100 : 0) - (!enableBackupWarning && !enablePINWarning ? 240 : 0)
        
        focus: visible
        modal: true
    }

    WalletCreatedDialog {
        id: walletCreatedDialog
        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 540 ? 540 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > 360 ? 360 - 40 : applicationWindow.height - 40
        
        focus: visible
        modal: true
    }

    // Help dialogs

    DialogAbout {
        id: dialogAbout

        parent: Overlay.overlay
        anchors.centerIn: parent
        width: applicationWindow.width > 540 ? 540 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > 640 ? 640 - 40 : applicationWindow.height - 40

        onLicenseRequested: {
            close()
            dialogAboutLicense.open()
        }
    }

    DialogAboutQt {
        id: dialogAboutQt

        parent: Overlay.overlay
        anchors.centerIn: parent
        width: applicationWindow.width > 540 ? 540 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > 640 ? 640 - 40 : applicationWindow.height - 40
    }

    DialogAboutLicense {
        id: dialogAboutLicense

        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width - 40
        height: applicationWindow.height - 40
        
        focus: visible
        modal: true
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
    } // Rectangle (flasher)
}
