import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import QtGraphicalEffects 1.12

// Resource imports
// import "qrc:/ui/src/ui/Dialogs"
import "../Dialogs/" // For quick UI development, switch back to resources when making a release

Item {
    id: root

    readonly property real delegateHeight: 30
    property bool emptyAddressVisible: true
    property bool expanded: expand
    // The following property is used to avoid a binding conflict with the `height` property.
    // Also avoids a bug with the animation when collapsing a wallet
    readonly property real finalViewHeight: expanded ? delegateHeight*(addressList.count) + 50 : 0


    width: blocksList.width
    height: itemDelegateMainButton.height + (expanded ? finalViewHeight : 0)

    Behavior on height { NumberAnimation { duration: 250; easing.type: Easing.OutQuint } }

    ColumnLayout {
        id: delegateColumnLayout
        anchors.fill: parent

        ItemDelegate {
            id: itemDelegateMainButton
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignTop
            font.bold: expanded

            RowLayout {
                id: delegateRowLayout
                anchors.fill: parent
                anchors.leftMargin: listBlockLeftMargin
                anchors.rightMargin: listBlockRightMargin
                spacing: listBlockSpacing

                Label {
                    id: labelBlockTime
                    text: date // a role of the model
                    Layout.leftMargin: listBlocksLeftMargin
                    Layout.fillWidth: true
                }

                Label {
                    id: labelBlockNumber
                    text: blockNumber // a role of the model
                    color: Material.accent
                    horizontalAlignment: Text.AlignLeft
                    Layout.preferredWidth: internalLabelsWidth
                }

                Label {
                    id: labelTransactions
                    text: txNumber // a role of the model
                    horizontalAlignment: Text.AlignRight
                    Layout.rightMargin: listBlocksRightMargin
                    Layout.preferredWidth: internalLabelsWidth
                }

                Label {
                    id: labelHash
                    text: hash // a role of the model
                    color: Material.accent
                    horizontalAlignment: Text.AlignRight
                    Layout.rightMargin: listBlocksRightMargin
                    visible:parent.width>750
                    Layout.fillWidth: true

                }
            } // RowLayout
        } // ItemDelegate
    }
}