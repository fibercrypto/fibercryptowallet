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
        maximumFlickVelocity: 800
        spacing: 10
        
        

        model: modelFilters
        delegate: HistoryFilterListDelegate {
            id: filterDelegate
            
           

            
        }
        
    }

    

    

    ModelManager{
        id: modelManager
        property Timer timer: Timer{
            id: modelManagerTimer
            repeat: false
            running: true
            interval: 0
            onTriggered:{
                modelManager.setWalletManager(walletManager)
            }
        }
    }
    WalletModel{
        id: modelFilters
        property Timer timer: Timer {
        
                                    id: walletModelTimer
                                    repeat: false
                                    running: true
                                    interval: 0
                                    onTriggered: {
                                        modelFilters.loadModel(walletManager.getWallets())
                                        walletModelTimer.running = false

                                    }

                                }
    }
}
