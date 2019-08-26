import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

import "../"

ItemDelegate {
    id: root

    property string modelNodeDirection: "http://localhost:6420"
    property var scrollSettingsWidth: 0
    property bool modelIsWalletEnvLocal: true
    property string modelWalletPath: qsTr("$HOME/.skycoin/wallets")
                        Layout.fillWidth: true
                        Layout.fillHeight: true


    width: parent.width
    height: walletPanel.height + panelNetwork.height

    ColumnLayout {
        id: columnSettings
        anchors.fill: parent
//        anchors.leftMargin: 20
//        anchors.rightMargin: 20
//        anchors.topMargin: 10
//        anchors.bottomMargin: 12
        spacing: 20
//        anchors.fill: parent
//        spacing: 5
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
