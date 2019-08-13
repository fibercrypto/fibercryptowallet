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
    //property list<AddressModel> addressModels
    //property AddressModel myAddressModel : AddressModel{}
    signal walletsLoaded
    
    
    clip: true

    ListView {

        id: listViewFilters
        property list<CheckedList> checkedList
        anchors.fill: parent
        maximumFlickVelocity: 800
        spacing: 10
        
        

        model: modelFilters
        delegate: HistoryFilterListDelegate {
            id: filterDelegate
            
           

            
        }
        
    }

    

    // This model can be the same as the wallet list,
    // as this model need to expose all wallets and their addresses.
    // For that, it should be implemented in the backend, instead of here.
    //ListModel { // EXAMPLE
    //    id: modelFilters
//
    //    ListElement { name: "My first wallet" }
    //    ListElement { name: "My second wallet" }
    //    ListElement { name: "My third wallet" }
    //    ListElement { name: "My fourth wallet" }
    //    ListElement { name: "My fiveth wallet" }
    //    ListElement { name: "My sixth wallet" }
    //}

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
