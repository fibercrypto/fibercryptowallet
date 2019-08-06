import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/"
import "../" // For quick UI development, switch back to resources when making a release

Dialog {
    id: dialogSendTransaction

    property alias expanded: transactionDetails.expanded

    title: qsTr("Password requested")
    standardButtons: Dialog.Ok | Dialog.Cancel

    onOpened: {
        forceActiveFocus()
        passwordRequester.forceTextFocus()
        standardButton(Dialog.Ok).enabled = passwordRequester.text !== ""
    }

    onClosed: {
        passwordRequester.clear()
    }

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

                // EXAMPLE
                date: "2019-02-26 15:27"
                status: TransactionDetails.Status.Preview
                type: TransactionDetails.Type.Receive
                amount: 10
                hoursReceived: 16957
                hoursBurned: 33901
                transactionID: "kq7wdkkUT736dnuyQasdhsaDJ9874jk"

                Layout.fillWidth: true
            }

            Rectangle {
                visible: expanded
                height: 1
                color: Material.color(Material.Grey)
                Layout.fillWidth: true
            }

            ColumnLayout {
                id: columnLayoutPasswordField
                Layout.fillWidth: true

                Label {
                    text: qsTr("Enter the password to confirm the transaction")
                    font.bold: true
                }

                PasswordRequester {
                    id: passwordRequester

                    Layout.topMargin: -10
                    Layout.fillWidth: true

                    onTextChanged: {
                        dialogSendTransaction.standardButton(Dialog.Ok).enabled = text !== ""
                    }
                }
            } // ColumnLayout
        } // ColumnLayout (root)

        ScrollIndicator.vertical: ScrollIndicator {
            parent: dialogSendTransaction.contentItem
            anchors.top: flickable.top
            anchors.bottom: flickable.bottom
            anchors.right: parent.right
            anchors.rightMargin: -dialogSendTransaction.rightPadding + 1
        }
    } // Flickable
}
