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

    // Roles: type, date, amount, sentAddress, receivedAddress
    // Use listModel.append( { "type": value, "date": value, "sentAddress": value, "receivedAddress": value } )
    // Or implement the model in the backend (a more recommendable approach)
    ListModel { // EXAMPLE
        id: modelTransactions
    }
}
