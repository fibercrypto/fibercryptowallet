import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/"
import "../" // For quick UI development, switch back to resources when making a release

Dialog {
    id: dialogTransactionsDetails

    property alias date: transactionDetails.date
    property alias status: transactionDetails.status
    property alias type: transactionDetails.type
    property alias amount: transactionDetails.amount
    property alias hoursReceived: transactionDetails.hoursReceived
    property alias hoursBurned: transactionDetails.hoursBurned
    property alias transactionID: transactionDetails.transactionID
    property alias modelInputs: transactionDetails.modelInputs
    property alias modelOutputs: transactionDetails.modelOutputs

    property alias expanded: transactionDetails.expanded

    title: qsTr("Transaction details")
    standardButtons: Dialog.Ok

    Flickable {
        id: flickable
        anchors.fill: parent
        clip: true
        contentHeight: columnLayoutRoot.height

        ColumnLayout {
            id: columnLayoutRoot
            width: parent.width
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
        } // ColumnLayout

        ScrollIndicator.vertical: ScrollIndicator {
            parent: dialogTransactionsDetails.contentItem
            anchors.top: flickable.top
            anchors.bottom: flickable.bottom
            anchors.right: parent.right
            anchors.rightMargin: -dialogTransactionsDetails.rightPadding + 1
        }
    }
}
