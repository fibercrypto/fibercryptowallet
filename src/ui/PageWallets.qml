import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import WalletsManager 1.0

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
// import "qrc:/ui/src/ui/Dialogs"
import "Delegates/" // For quick UI development, switch back to resources when making a release
import "Dialogs/" // For quick UI development, switch back to resources when making a release

Page {
    id: root

    property string statusIcon; // an empty string for no icon
    readonly property real listWalletLeftMargin: 20
    readonly property real listWalletRightMargin: 50
    readonly property real listWalletSpacing: 20
    readonly property real internalLabelsWidth: 70

    header: ColumnLayout {

        RowLayout {
            spacing: listWalletSpacing
            Layout.topMargin: 30

            Label {
                text: qsTr("Name")
                font.pointSize: 9
                Layout.leftMargin: listWalletLeftMargin
                Layout.fillWidth: true
            }
            Label {
                text: qsTr("Sky")
                font.pointSize: 9
                horizontalAlignment: Text.AlignRight
                Layout.preferredWidth: internalLabelsWidth
            }
            Label {
                text: qsTr("Coin hours")
                font.pointSize: 9
                horizontalAlignment: Text.AlignRight
                Layout.rightMargin: listWalletRightMargin
                Layout.preferredWidth: internalLabelsWidth
            }
        } // RowLayout

        Rectangle {
            id: rect
            Layout.fillWidth: true
            height: 1
            color: "#DDDDDD"
        }
    } // ColumnLayout (header)

    footer: ToolBar {
        id: tabBarCreateUpload
        Material.theme: applicationWindow.Material.theme
        Material.primary: applicationWindow.accentColor
        Material.foreground: applicationWindow.Material.background
        Material.elevation: 0

        RowLayout {
            anchors.fill: parent
            ToolButton {
                id: buttonAddWallet
                text: qsTr("Add wallet")
                icon.source: "qrc:/images/resources/images/icons/add.svg"
                Layout.fillWidth: true

                onClicked: {
                    dialogAddLoadWallet.mode = CreateLoadWallet.Create
                    dialogAddLoadWallet.open()
                }

            }
            ToolButton {
                id: buttonLoadWallet
                text: qsTr("Load wallet")
                icon.source: "qrc:/images/resources/images/icons/upload.svg"
                Layout.fillWidth: true

                onClicked: {
                    dialogAddLoadWallet.mode = CreateLoadWallet.Load
                    dialogAddLoadWallet.open()
                }
            }
        } // RowLayout
    } // ToolBar (footer)

    ScrollView {
        id: scrollItem

        anchors.fill: parent
        ScrollBar.horizontal.policy: ScrollBar.AlwaysOff

        ListView {
            id: walletList
            anchors.fill: parent
            clip: true // limit the painting to it's bounding rectangle
            model: walletModel
            delegate: WalletListDelegate {}
            function updateView(wallets, walletsExpanded) {
                var index = 0;
                for(var child in walletList.contentItem.children) {
                    walletList.contentItem.children[child].name = wallets[index].name;
                    walletList.contentItem.children[child].sky = wallets[index].sky;
                    walletList.contentItem.children[child].encryptionEnabled = wallets[index].encryptionEnabled;
                    walletList.contentItem.children[child].coinHours = wallets[index].coinHours;
                    walletList.contentItem.children[child].fileName = wallets[index].fileName;
                    walletList.contentItem.children[child].expanded = walletsExpanded[index];
                }
                if (index < wallets.length) {
                    var i;
                    for(i = index; i < wallets.length; i++) {
                        walletModel.addWallet(wallets[i]);
                        walletList.contentItem.children[i].expanded = walletsExpanded[i].expanded;
                    }
                }
            }
            function getChildWithName(contentItem, index) {
                if (index < contentItem.children.length) {
                    return contentItem.children[index];
                }
                return undefined
            }
        }
    }

    WalletModel {
        id: walletModel

        Component.onCompleted: {
            walletModel.loadModel(walletManager.getWallets())
        }
        property Timer timer: Timer {
            id: walletListTimer
            interval: 5000
            repeat: true
            running: true
            onTriggered: {
                var index;
                var walletsExpanded = []
                console.log(walletModel)
                for (index = 0; index < walletList.count; index++) {
                    walletsExpanded.push(walletList.getChildWithName(walletList.contentItem, index).expanded);
                }
                console.log("SUPERRRR");
                console.log(walletsExpanded);
////                walletManager.updateWallets()
//                walletModel.loadModelUsingExpanded(walletManager.getWallets(), walletsExpanded)
                walletList.updateView(walletManager.getWallets(), walletsExpanded)
                console.log("Finito")
//                for (index = 0; index < walletList.count; index++) {
//                    walletList.itemAtIndex(index).expanded = walletsExpanded[index]
//                }
            }
        }
    }

    DialogAddLoadWallet {
        id: dialogAddLoadWallet
        anchors.centerIn: Overlay.overlay

        modal: true
        focus: true

        width: applicationWindow.width > 540 ? 540 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > 640 ? 640 - 40 : applicationWindow.height - 40

        onAccepted: {
            console.log("Add wallet")
        }
    }


    // Roles: name, encryptionEnabled, sky, coinHours
    // Use listModel.append( { "name": value, "encryptionEnabled": value, "sky": value, "coinHours": value } )
    // Or implement the model in the backend (a more recommendable approach)
    ListModel {
        id: listWallets
        ListElement { name: "My first wallet"; encryptionEnabled: true; sky: 5; coinHours: 10 }
        ListElement { name: "My second wallet"; encryptionEnabled: true; sky: 300; coinHours: 1049 }
        ListElement { name: "My third wallet"; encryptionEnabled: false; sky: 13; coinHours: 201 }

        ListElement { name: "My first wallet"; encryptionEnabled: false; sky: 5; coinHours: 10 }
        ListElement { name: "My second wallet"; encryptionEnabled: true; sky: 300; coinHours: 1049 }
        ListElement { name: "My third wallet"; encryptionEnabled: true; sky: 13; coinHours: 201 }

        ListElement { name: "My first wallet"; encryptionEnabled: true; sky: 5; coinHours: 10 }
        ListElement { name: "My second wallet"; encryptionEnabled: false; sky: 300; coinHours: 1049 }
        ListElement { name: "My third wallet"; encryptionEnabled: false; sky: 13; coinHours: 201 }

        ListElement { name: "My first wallet"; encryptionEnabled: true; sky: 5; coinHours: 10 }
        ListElement { name: "My second wallet"; encryptionEnabled: false; sky: 300; coinHours: 1049 }
        ListElement { name: "My third wallet"; encryptionEnabled: true; sky: 13; coinHours: 201 }
    }
}
