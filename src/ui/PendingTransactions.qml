import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import PendingModels 1.0

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "Delegates/" // For quick UI development, switch back to resources when making a release

Page {
    id: root

    property bool showOnlyMine: false

    GroupBox {
        anchors.fill: parent
        anchors.margins: 20
        clip: true

        label: CheckDelegate {
            text: qsTr("Show only mine")
            checked: showOnlyMine
            onCheckedChanged: {
                showOnlyMine = checked
            }
        }

        ColumnLayout {
            id: columnLayoutFrame
            anchors.fill: parent

            ColumnLayout {
                id: columnLayoutHeader

                RowLayout {
                    Layout.topMargin: 30

                    Label {
                        text: qsTr("Transaction ID")
                        font.pointSize: 9
                        Layout.leftMargin: 10 // Empirically obtained
                        Layout.fillWidth: true
                    }
                    Label {
                        text: qsTr("Sky")
                        font.pointSize: 9
                        Layout.rightMargin: 92
                    }
                    Label {
                        text: qsTr("Coin hours")
                        font.pointSize: 9
                        Layout.rightMargin: 33
                    }
                    Label {
                        text: qsTr("Timestamp")
                        font.pointSize: 9
                        Layout.rightMargin: 98
                    }
                } // RowLayout

                Rectangle {
                    Layout.fillWidth: true
                    height: 1
                    color: "#DDDDDD"
                }
            } // ColumnLayout (header)

            ScrollView {
                id: scrollItem
                Layout.fillWidth: true
                Layout.fillHeight: true
                ScrollBar.horizontal.policy: ScrollBar.AlwaysOff

                ListView {
                    id: listPendingTransactions
                    anchors.fill: parent
                    clip: true

                    model: modelPendingTransactions
                    delegate: PendingTransactionsDelegate {
                        property bool hide: modelMine != showOnlyMine

                        width: parent.width

                        modelMine: mine
                        modelTransactionID: transactionID
                        modelSky: sky
                        modelCoinHours: coinHours
                        modelTimestamp: timestamp

                        height: hide ? 0 : implicitHeight
                        Behavior on height { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } }
                        opacity: hide ? 0 : 1
                        Behavior on opacity { NumberAnimation { duration: 100 } }

                        clip: true
                    }
                } // ListView
            } // ScrollView
        } // ColumnLayout (frame)
    } // Frame

    // Roles: mine, transactionID, sky, coinHours, timestamp
    // Implement the model in the backend (a more recommendable approach)
    QPendingList{
        id: modelPendingTransactions
    }
    // ListModel { // EXAMPLE
    //     id: modelPendingTransactions
    //     ListElement { transactionID: "AAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAAA"; sky: 100000; coinHours: 3545; timestamp: "2019-03-02 10:19" }
    //     ListElement { transactionID: "BBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBBB"; sky: 9999999; coinHours: 3576852; timestamp: "2019-03-02 10:19" }
    //     ListElement { transactionID: "CCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCCC"; sky: 56; coinHours: 35435; timestamp: "2019-03-02 10:19" }
    //     ListElement { transactionID: "DDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDDD"; sky: 6875342; coinHours: 5469; timestamp: "2019-03-02 10:19" }
    //     ListElement { transactionID: "EEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEEE"; sky: 65; coinHours: 35; timestamp: "2019-03-02 10:19" }
    //     ListElement { transactionID: "FFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFFF"; sky: 10; coinHours: 357232; timestamp: "2019-03-02 10:19" }
    //     ListElement { transactionID: "GGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGGG"; sky: 1; coinHours: 2103835; timestamp: "2019-03-02 10:19" }
    //     ListElement { transactionID: "HHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHHH"; sky: 26987530; coinHours: 9753212; timestamp: "2019-03-02 10:19" }
    //     ListElement { transactionID: "IIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIIII"; sky: 357; coinHours: 3108465; timestamp: "2019-03-02 10:19" }
    //     ListElement { transactionID: "JJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJJ"; sky: 64588; coinHours: 687985; timestamp: "2019-03-02 10:19" }
    //     ListElement { transactionID: "KKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKKK"; sky: 47; coinHours: 67635; timestamp: "2019-03-02 10:19" }
    //     ListElement { transactionID: "LLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLLL"; sky: 9; coinHours: 3543542; timestamp: "2019-03-02 10:19" }
    // }
}