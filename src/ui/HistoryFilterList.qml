import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import WalletsManager 1.0

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "Delegates/" // For quick UI development, switch back to resources when making a release

Item {
    id: root

    readonly property alias count: listViewFilters.count
    property alias interactive: listViewFilters.interactive
    property alias contentHeight: listViewFilters.contentHeight
    
    clip: true

    ListView {
        id: listViewFilters
        
        anchors.fill: parent
        spacing: 10
        
        model: modelFilters
        delegate: HistoryFilterListDelegate {}
    }

    ModelManager {
        id: modelManager
        
        Component.onCompleted: {
            setWalletManager(walletManager)
        }
    }

    WalletModel {
        id: modelFilters

        Component.onCompleted: {
            loadModel(walletManager.getWallets())
        }
    }
}
