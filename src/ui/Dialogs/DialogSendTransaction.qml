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
    property bool showPasswordField: false
    property string passwordText

    property alias previewDate: transactionDetails.date                    
    property alias previewType: transactionDetails.type                  
    property alias previewAmount: transactionDetails.amount              
    property alias previewHoursReceived: transactionDetails.hoursReceived
    property alias previewHoursBurned: transactionDetails.hoursBurned    
    property alias previewtransactionID: transactionDetails.transactionID

    title: qsTr("Confirm transaction")
    standardButtons: Dialog.Ok | Dialog.Cancel

    onOpened: {
        forceActiveFocus()
        passwordRequester.forceTextFocus()
        standardButton(Dialog.Ok).enabled = passwordRequester.text !== "" || showPasswordField
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
                opacity: showPasswordField ? 1.0 : 0.0
                Behavior on opacity { NumberAnimation {} }
                visible: opacity > 0

                Label {
                    text: qsTr("Enter the password to confirm the transaction")
                    font.bold: true
                }

                PasswordRequester {
                    id: passwordRequester

                    Layout.topMargin: -10
                    Layout.fillWidth: true

                    onTextChanged: {
                        dialogSendTransaction.standardButton(Dialog.Ok).enabled = text !== "" || showPasswordField
						passwordText = text
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
