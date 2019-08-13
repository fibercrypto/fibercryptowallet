import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import WalletsManager 1.0
import "../"
Item {
    id: root

    readonly property real addressListHeight: listViewFilterAddress.height
    readonly property real delegateHeight: 42
    property alias tristate: checkDelegate.tristate
    property alias walletText: checkDelegate.text
    

    clip: true
    width: 300
    height: checkDelegate.height +  columnLayout.spacing + listViewFilterAddress.height
    
    
    ColumnLayout {
        id: columnLayout
        anchors.fill: parent

        CheckDelegate {
            id: checkDelegate

            Layout.fillWidth: true
            tristate: true
            text: name
            LayoutMirroring.enabled: true

            nextCheckState: function() {
                if (checkState === Qt.Checked) {
                    if (!listViewFilterAddress.allChecked) {
                        listViewFilterAddress.allChecked = true
                    }
                    listViewFilterAddress.allChecked = false
                    return Qt.Unchecked
                } else {
                    if (listViewFilterAddress.allChecked) {
                        listViewFilterAddress.allChecked = false
                    }
                    listViewFilterAddress.allChecked = true
                    return Qt.Checked
                }
            }

            contentItem: Label {
                leftPadding: checkDelegate.indicator.width + checkDelegate.spacing
                text: checkDelegate.text
                verticalAlignment: Qt.AlignVCenter
                color: checkDelegate.enabled ? checkDelegate.Material.foreground : checkDelegate.Material.hintTextColor
            }
        }

        ListView {
            id: listViewFilterAddress
            property AddressModel listAddresses
            property int checkedDelegates: 0
            property bool allChecked: false
            model: listAddresses
            onCheckedDelegatesChanged: {
                if (checkedDelegates === 0) {
                    checkDelegate.checkState = Qt.Unchecked
                } else if (checkedDelegates === count) {
                    checkDelegate.checkState = Qt.Checked
                } else {
                    checkDelegate.checkState = Qt.PartiallyChecked
                }
                
            }
            onCountChanged:{
                implicitHeight = listAddresses.count * delegateHeight
                
               
            }
            

            Component.onCompleted:{
                modelManager.setWalletManager(walletManager)
                listAddresses = modelManager.getAddressModel(fileName)
            }
            
            Layout.fillWidth: true
           

            interactive: false
            
            delegate: HistoryFilterListAddressDelegate {
                leftPadding: 20
                scale: 0.85
                checked: marked
                 
                 onCheckedChanged: {
                   
                    ListView.view.checkedDelegates += checked ? 1: -1
                    
                    if (checked == 1){
                        historyManager.addFilter(address)
                        
                    }
                    else {
                        historyManager.removeFilter(address)
                        
                    }
                    listViewFilterAddress.listAddresses.editAddress(index, address, sky, coinHours, checked)
                }
                
                Connections{
                    target: listViewFilterAddress
                    onAllCheckedChanged:{
                        if (listViewFilterAddress.allChecked){
                            listViewFilterAddress.listAddresses.editAddress(index, address, sky, coinHours, 1)
                        } else{
                            listViewFilterAddress.listAddresses.editAddress(index, address, sky, coinHours, 0)
                        }
                    }
                }
                
                
            }
        } // ListView
    } // ColumnLayout

    // This model can be the same as the wallet address list,
    // as this model need to expose all addresses for each wallet.
    // For that, it should be implemented in the backend, instead of here.
    //ListModel { // EXAMPLE
    //    id: listAddresses
//
    //    ListElement { address: "qrxw7364w8xerusftaxkw87ues" }
    //     ListElement { address: "qrxw7364w8xerusftaxkw87ues" }
    //      ListElement { address: "qrxw7364w8xerusftaxkw87ues" }
    //    ListElement { address: "8745yuetsrk8tcsku4ryj48ije" }
    //    ListElement { address: "gfdhgs343kweru38200384uwqd" }
    //}

    //AddressModel {
//
    //    id: listAddresses
    //    property Timer timer: Timer {
    //                            id: addressModelTimer2
    //                            interval: 0
    //                            repeat: false
    //                            running: true
    //                            onTriggered: {
    //                                listAddresses.loadModel(walletManager.getAddresses(fileName))
    //                                listAddresses.removeAddress(0)
    //                                addressModelTimer2.running = false 
    //                                console.log("Loaded")
    //                                
    //                                   
    //                                }
    //                        }
    //}
}
