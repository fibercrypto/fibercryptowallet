import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "Delegates/" // For quick UI development, switch back to resources when making a release

Page {
    id: settings

    Frame {
        id: frame
        anchors.fill: parent
        anchors.margins: 20
        clip: true


        ScrollView {
            id: scrollSettings
            anchors.fill: parent
            clip: true

            ColumnLayout {
                id: columnSettings
                anchors.fill: parent
                spacing: 5
                PanelItem {
                    id: walletPanel
                    title: "Wallets"
                    item: WalletSettings {
                        id: walletSettings
                    }
                }
                PanelItem {
                    id: panelNetwork
                    title: "Networks"
                    item: NetworkSettings {
                        id: networkSettings
                    }
                }
            }
        }
    }
}