import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import WalletsManager 1.0

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "Delegates/" // For quick UI development, switch back to resources when making a release

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
        }

        Rectangle {
            id: rect
            Layout.fillWidth: true
            height: 1
            color: "#DDDDDD"
        }
    }

    footer: ToolBar {
        id: tabBarCreateUpload
        Material.primary: Material.Blue
        Material.accent: Material.Yellow

        RowLayout {
            anchors.fill: parent
            ToolButton {
                id: buttonAddWallet
                text: qsTr("Add wallet")
                icon.source: "qrc:/images/resources/images/icons/add.svg"
                Layout.fillWidth: true

                
            }
            ToolButton {
                id: buttonLoadWallet
                text: qsTr("Load wallet")
                icon.source: "qrc:/images/resources/images/icons/upload.svg"
                Layout.fillWidth: true
            }
        }
    }

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
            
        }
    }

    
    WalletManager{
        id: walletManager
    }

    WalletModel{
        id: walletModel
        property Timer timer: Timer {
        
                                    id: walletModelTimer
                                    repeat: false
                                    running: true
                                    interval: 0
                                    onTriggered: {
                                        walletModel.loadModel(walletManager.getWallets())
                                        walletModelTimer.running = false
                                    }
            
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
