import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Dialog {
    id: root

    ColumnLayout {
        anchors.fill: parent
        spacing: 20

        TransactionDetails {
            id: transactionDetails
            implicitWidth: 500
            Layout.fillWidth: true
        }

        Rectangle {
            visible: transactionDetails.expanded
            height: 1
            color: Material.color(Material.Grey)
            Layout.fillWidth: true
        }
    }
}
