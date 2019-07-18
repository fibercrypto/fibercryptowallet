import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
import "qrc:/ui/src/ui/Dialogs"

Page {
    id: root

    GroupBox {
        anchors.fill: parent
        anchors.margins: 20
        clip: true

        label: RowLayout {
            SwitchDelegate {
                id: switchFilters
                text: qsTr("Filters")
            }
            Button {
                id: buttonFilters
                flat: true
                enabled: switchFilters.checked
                highlighted: true
                text: qsTr("Select filters")

                onClicked: {
                    toolTipFilters.open()
                }
            }
        } // RowLayout (GroupBox label)

        ScrollView {
            anchors.fill: parent
            clip: true
            ListView {
                id: listTransactions

                model: modelTransactions
                delegate: HistoryListDelegate {
                    onClicked: {
                        dialogTransactionDetails.open()
                        listTransactions.currentIndex = index
                    }
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
                contentHeight: filter.implicitHeight
                HistoryFilterList {
                    id: filter
                    interactive: false
                }
            }
        } // ColumnLayout
    } // ToolTip

    DialogTransactionDetails {
        id: dialogTransactionDetails
        anchors.centerIn: Overlay.overlay

        modal: true
        focus: true
        standardButtons: Dialog.Ok

        date: listTransactions.currentItem.modelDate
        status: listTransactions.currentItem.modelStatus
        type: listTransactions.currentItem.modelType
        amount: listTransactions.currentItem.modelAmount
        hoursReceived: listTransactions.currentItem.modelHoursReceived
        hoursBurned: listTransactions.currentItem.modelHoursBurned
        transactionID: listTransactions.currentItem.modelTransactionID
    }

    // Roles: type, date, amount, sentAddress, receivedAddress
    // Use listModel.append( { "type": value, "date": value, "sentAddress": value, "receivedAddress": value } )
    // Or implement the model in the backend (a more recommendable approach)
    ListModel { // EXAMPLE
        id: modelTransactions
        ListElement { date: "2019-03-02 10:19"; status: TransactionDetails.Status.Confirmed; type: TransactionDetails.Type.Send;    amount: 9;     hoursReceived: 65;   hoursBurned: 6287;   transactionID: "ksuwet837iwetr27ietr"; sentAddress: "734irwepaweygtawieta783cwyc"; receivedAddress: "uewyru63u3789jfrgpaldcxz4f6" }
        ListElement { date: "2019-03-02 10:44"; status: TransactionDetails.Status.Confirmed; type: TransactionDetails.Type.Receive; amount: 10000; hoursReceived: 9421; hoursBurned: 218489; transactionID: "hdk84iesurys29q8ikhf"; sentAddress: "nvslkjoid98wemxsoqzmap50vah"; receivedAddress: "ps73729ssksof972jwkf83owhfi" }
        ListElement { date: "2019-03-02 15:01"; status: TransactionDetails.Status.Confirmed; type: TransactionDetails.Type.Send;    amount: 125;   hoursReceived: 4203; hoursBurned: 7950;   transactionID: "p9438938qksaxjhsdkuq"; sentAddress: "qwe98uimhfdkfu23y9hewi8qd02"; receivedAddress: "p2837djkdjvkauyfiawbw83h48f" }
        ListElement { date: "2019-03-03 08:13"; status: TransactionDetails.Status.Confirmed; type: TransactionDetails.Type.Send;    amount: 100;   hoursReceived: 3877; hoursBurned: 6911;   transactionID: "73t4dunksxr8w7zuwe13"; sentAddress: "89niweuq82zqur37izuwsklparu"; receivedAddress: "0j3726djf9f4w9sovgwuw9e4n9s" }
        ListElement { date: "2019-03-03 12:23"; status: TransactionDetails.Status.Confirmed; type: TransactionDetails.Type.Receive; amount: 80;    hoursReceived: 3084; hoursBurned: 4603;   transactionID: "9782nz87e7tquetyn83w"; sentAddress: "t7ekwduf8045ogmcbq63nf9bm36"; receivedAddress: "aske8jdhbsmwq9204h49y25jrue" }
        ListElement { date: "2019-03-03 13:33"; status: TransactionDetails.Status.Confirmed; type: TransactionDetails.Type.Receive; amount: 60;    hoursReceived: 2796; hoursBurned: 2347;   transactionID: "9127enqiuyrt6tyejshd"; sentAddress: "2imdiym9w7ji8s29ydk0djwd6wj"; receivedAddress: "0gndqy725k9fj3ewdardgbkd83l" }
    }
}
