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

            Layout.fillWidth: true
            height: count * delegateHeight

            delegate: HistoryFilterListAddressDelegate {
                width: ListView.view.width
                height: delegateHeight
                leftPadding: 30
            }
        } // ListView
    } // ColumnLayout
}
