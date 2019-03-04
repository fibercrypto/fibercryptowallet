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

    ColumnLayout {
        width: root.width

        CheckDelegate {
            id: checkDelegate

            Layout.fillWidth: true
            tristate: true
            text: name
            LayoutMirroring.enabled: true

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

            Layout.fillWidth: true
            height: count * delegateHeight

            model: listAddresses
            delegate: HistoryFilterListAddressDelegate {
                width: ListView.view.width
                height: delegateHeight
                leftPadding: 30
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
