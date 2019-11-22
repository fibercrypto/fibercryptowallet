import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import WalletsManager 1.0

// Resource imports
// import "qrc:/ui/src/ui/"
import "../" // For quick UI development, switch back to resources when making a release

Item {
    id: root

    readonly property real addressListHeight: listViewFilterAddress.height
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
                verticalAlignment: Qt.AlignVCenter
                text: checkDelegate.text
                color: checkDelegate.enabled ? checkDelegate.Material.foreground : checkDelegate.Material.hintTextColor
            }
        } // CheckDelegate

        ListView {
            id: listViewFilterAddress

            property AddressModel listAddresses
            property int checkedDelegates: 0
            property bool allChecked: false

            model: listAddresses
            
            Layout.fillWidth: true
            height: contentHeight
            interactive: false


            onCheckedDelegatesChanged: {
                if (checkedDelegates === 0) {
                    checkDelegate.checkState = Qt.Unchecked
                } else if (checkedDelegates === count) {
                    checkDelegate.checkState = Qt.Checked
                } else {
                    checkDelegate.checkState = Qt.PartiallyChecked

                }
            }

            onAllCheckedChanged: {
                if (listViewFilterAddress.allChecked) {
                    listViewFilterAddress.listAddresses.editAddress(index, address, sky, coinHours, 1)
                } else {
                    listViewFilterAddress.listAddresses.editAddress(index, address, sky, coinHours, 0)
                }
            }

            Component.onCompleted: {
                modelManager.setWalletManager(walletManager)
                listAddresses = modelManager.getAddressModel(fileName)
            }

            delegate: HistoryFilterListAddressDelegate {
                // BUG: Checking the wallet does not change the check state of addresses
                // Is `checked: marked` ok? Or it should be the opposite?
                checked: true 
                width: parent.width
                text: address 

                onCheckedChanged: {
                    ListView.view.checkedDelegates += checked ? 1: -1
                    
                    if (checked == true) {
                        historyManager.addFilter(address)
                    } else {
                        historyManager.removeFilter(address)  
                    }
                    listViewFilterAddress.listAddresses.editAddress(index, address, sky, coinHours, checked)
                }
            } // HistoryFilterListAddressDelegate (delegate)
        } // ListView
    } // ColumnLayout
}
