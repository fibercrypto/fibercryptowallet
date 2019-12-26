import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

import "Delegates/" // For quick UI development, switch back to resources when making a release

Page {
    id: explorer
    readonly property real listBlocksLeftMargin: 20
    readonly property real listBlocksRightMargin: 50
    readonly property real listBlocksSpacing: 20
    readonly property real internalLabelsWidth: 70
    header: ColumnLayout {

        RowLayout {
            spacing: listBlocksSpacing
            Layout.topMargin: 30

            Label {
                text: qsTr("Time")
                font.pointSize: 9
                Layout.leftMargin: listBlocksLeftMargin
                Layout.fillWidth: true
            }
            Label {
                text: qsTr("Block Number")
                font.pointSize: 9
                horizontalAlignment: Text.AlignLeft
                Layout.preferredWidth: internalLabelsWidth
            }
            Label {
                text: qsTr("Transactions")
                font.pointSize: 9
                horizontalAlignment: Text.AlignRight
                Layout.rightMargin: listBlocksRightMargin
                Layout.preferredWidth: internalLabelsWidth
            }
            Label {
                text: qsTr("                                                                             Hash")
                font.pointSize: 9
                visible:parent.width>750
                horizontalAlignment: Text.AlignRight
                Layout.rightMargin: listBlocksRightMargin
                Layout.fillWidth: true
            }

        } // RowLayout

        Rectangle {
            id: rect
            Layout.fillWidth: true
            height: 1
            color: "#DDDDDD"
        }
    } // ColumnLayout (header)


ScrollView {
        id: scrollItem

        anchors.fill: parent
        ScrollBar.horizontal.policy: ScrollBar.AlwaysOff

        ListView {
            id: blocksList
            anchors.fill: parent
            clip: true // limit the painting to it's bounding rectangle
            model: blockModel
            delegate: BlockListDelegate{}
        }
    }
ListModel{
id:blockModel
ListElement{
date:"01/02/1990"
blockNumber:10000898
txNumber:2
hash:"9516639399ab2b02f6f0bc873f03f2f2ac1c94853dd031685de1021be78c71d7"
}
ListElement{
date:"01/02/1990"
blockNumber:1
txNumber:2
hash:"9516639399ab2b02f6f0bc873f03f2f2ac1c94853dd031685de1021be78c71d7"

}
ListElement{
date:"01/02/1990"
blockNumber:2
txNumber:1
hash:"5ac7b5f4d170e606cc75e08c20b3ec216703e9ca77df178ea0130255bf20d5e5"
}
ListElement{
date:"01/02/1990"
blockNumber:3
txNumber:1
hash:"5ac7b5f4d170e606cc75e08c20b3ec216703e9ca77df178ea0130255bf20d5e5"
}
}//listModel

}
