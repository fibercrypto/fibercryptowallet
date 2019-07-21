import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Item {
    id: root

    readonly property real addressListHeight: listViewFilterAddress.height
    readonly property real delegateHeight: 42
    property alias tristate: checkDelegate.tristate
    property alias walletText: checkDelegate.text

    clip: true
    width: 300
    height: checkDelegate.height + columnLayout.spacing + listViewFilterAddress.height

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

            property int checkedDelegates: 0
            property bool allChecked: false

            onCheckedDelegatesChanged: {
                if (checkedDelegates === 0) {
                    checkDelegate.checkState = Qt.Unchecked
                } else if (checkedDelegates === count) {
                    checkDelegate.checkState = Qt.Checked
                } else {
                    checkDelegate.checkState = Qt.PartiallyChecked
                }
            }

            Layout.fillWidth: true
            height: count * delegateHeight

            interactive: false
            model: listAddresses
            delegate: HistoryFilterListAddressDelegate {
                leftPadding: 20
                scale: 0.85
                checked: ListView.view.allChecked

                onCheckedChanged: {
                    ListView.view.checkedDelegates += checked ? 1 : -1
                }
            }
        } // ListView
    } // ColumnLayout

    // This model can be the same as the wallet address list,
    // as this model need to expose all addresses for each wallet.
    // For that, it should be implemented in the backend, instead of here.
    ListModel { // EXAMPLE
        id: listAddresses

        ListElement { address: "qrxw7364w8xerusftaxkw87ues" }
        ListElement { address: "8745yuetsrk8tcsku4ryj48ije" }
        ListElement { address: "gfdhgs343kweru38200384uwqd" }
    }
}
