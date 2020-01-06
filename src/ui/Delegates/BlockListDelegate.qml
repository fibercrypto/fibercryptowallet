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

    width: blocksList.width
    height: itemDelegateMainButton.height

    Behavior on height { NumberAnimation { duration: 250; easing.type: Easing.OutQuint } }

    ColumnLayout {
        id: delegateColumnLayout
        anchors.fill: parent

        ItemDelegate {
//        Component.onCompleted:{
//        console.log(modelData.transactionList)
//        }
            id: itemDelegateMainButton
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignTop

            onClicked:{
                generalStackView.openBlockPage(blockhash)
            }

            RowLayout {
                id: delegateRowLayout
                anchors.fill: parent
                anchors.leftMargin: listBlocksLeftMargin
                anchors.rightMargin: listBlocksRightMargin
                spacing: listBlocksSpacing

                Label {
                    id: labelBlockTime
                    text: Qt.formatDateTime(time, Qt.DefaultLocaleShortDate) // a role of the model
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
                    text: transactionLen // a role of the model
                    horizontalAlignment: Text.AlignRight
                    Layout.rightMargin: listBlocksRightMargin
                    Layout.preferredWidth: internalLabelsWidth
                }

                Label {
                    id: labelHash
                    text: blockhash // a role of the model
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