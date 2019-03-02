import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Page {
    id: root

    GroupBox {
        anchors.fill: parent
        anchors.margins: 20
        clip: true

        label: SwitchDelegate {
            id: switchFilters
            text: qsTr("Filters")
        }

        ScrollView {
            anchors.fill: parent
            clip: true
            ListView {
                id: listTransactions

                model: modelTransactions
                delegate: HistoryListDelegate {
                }
            }
        }
    } // GroupBox


    ToolTip {
        id: toolTipFilters

        anchors.centerIn: Overlay.overlay
        height: 16 + (filter.count > 5 ? filter.delegateHeight * 5 : filter.delegateHeight * filter.count)
        clip: true
        modal: true
        closePolicy: Popup.CloseOnEscape | Popup.CloseOnPressOutside

        ColumnLayout {
            anchors.fill: parent

            Label {
                font.pointSize: Qt.application.font.pointSize * 1.25
                font.bold: true
                text: qsTr("Available filters")
                Layout.alignment: Qt.AlignHCenter | Qt.AlignTop
            }

            ScrollView {
                Layout.alignment: Qt.AlignCenter
                Layout.fillWidth: true
                Layout.fillHeight: true
                clip: true
                HistoryFilterList {
                    id: filter
                    interactive: false
                }
            }
        } // ColumnLayout
    } // ToolTip

    // Roles: type, date, amount, sentAddress, receivedAddress
    // Use listModel.append( { "type": value, "date": value, "sentAddress": value, "receivedAddress": value } )
    // Or implement the model in the backend (a more recommendable approach)
    ListModel { // EXAMPLE
        id: modelTransactions
        ListElement { type: 0; date: "2019-03-02 10:19"; amount: 9; sentAddress: "734irwepaweygtawieta783cwyc"; receivedAddress: "uewyru63u3789jfrgpaldcxz4f6" }
        ListElement { type: 1; date: "2019-03-02 10:44"; amount: 10000; sentAddress: "nvslkjoid98wemxsoqzmap50vah"; receivedAddress: "ps73729ssksof972jwkf83owhfi" }
        ListElement { type: 0; date: "2019-03-02 15:01"; amount: 125; sentAddress: "qwe98uimhfdkfu23y9hewi8qd02"; receivedAddress: "p2837djkdjvkauyfiawbw83h48f" }
        ListElement { type: 0; date: "2019-03-03 08:13"; amount: 100; sentAddress: "89niweuq82zqur37izuwsklparu"; receivedAddress: "0j3726djf9f4w9sovgwuw9e4n9s" }
        ListElement { type: 1; date: "2019-03-03 12:23"; amount: 80; sentAddress: "t7ekwduf8045ogmcbq63nf9bm36"; receivedAddress: "aske8jdhbsmwq9204h49y25jrue" }
        ListElement { type: 1; date: "2019-03-03 13:33"; amount: 60; sentAddress: "2imdiym9w7ji8s29ydk0djwd6wj"; receivedAddress: "0gndqy725k9fj3ewdardgbkd83l" }
    }
}
