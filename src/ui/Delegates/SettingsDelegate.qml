import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

import "../"

Item {
    id: root

    property string modelNodeDirection
    property var scrollSettingsWidth: 0
    property bool modelIsWalletEnvLocal: true
    property string modelWalletPath: qsTr("$HOME/.skycoin/wallets")
    Layout.fillWidth: true
    height: walletPanel.height + panelNetwork.height
    Layout.alignment: Qt.AlignTop

    ColumnLayout {
        id: columnSettings
        anchors.fill: parent
        spacing: 20

        PanelItem {
            id: walletPanel
            Layout.fillWidth: true
            title: "Wallets"
            item: WalletSettings {
                id: walletSettings
                width: root.width
                isSetWalletEnvLocal: modelIsWalletEnvLocal
                walletPath: modelWalletPath
            }
        } // PanelItem
        PanelItem {
            id: panelNetwork
            Layout.fillWidth: true
            title: "Networks"
            item: NetworkSettings {
                width: root.width
                id: networkSettings
                nodeDirection: modelNodeDirection
            }
        } // PanelItem
        Component.onCompleted: {
            root.width = scrollSettingsWidth
            console.log("root -> " + root.width)
            console.log("networkSettings  -> " + panelNetwork.width)
        }
    } // ColumnLayout
} // ItemDelegate
