import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

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
                icon.source: "qrc:/images/add.svg"
                Layout.fillWidth: true
            }
            ToolButton {
                id: buttonLoadWallet
                text: qsTr("Load wallet")
                icon.source: "qrc:/images/upload.svg"
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
            model: listWallets
            delegate: WalletListDelegate {}
        }
    }

    // Roles: name, encryptionEnabled, sky, coinHours
    // Use listModel.append( { "name": value, "encryptionEnabled": value, "sky": value, "coinHours": value } )
    // Or implement the model in the backend (a more recommendable approach)
    ListModel {
        id: listWallets
    }
}
