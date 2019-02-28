import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Dialog {
    id: root
    anchors.centerIn: Overlay.overlay

    title: qsTr("Password requested")
    modal: true
    clip: true
    focus: true
    standardButtons: Dialog.Ok | Dialog.Cancel

    onOpened: {
        forceActiveFocus()
        passwordRequester.forceTextFocus()
        standardButton(Dialog.Ok).enabled = passwordRequester.text !== ""
    }

    onClosed: {
        passwordRequester.clear()
    }

    ColumnLayout {
        anchors.fill: parent
        spacing: 20

        TransactionDetails {
            id: transactionDetail

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
            visible: transactionDetail.expanded
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
                    root.standardButton(Dialog.Ok).enabled = text !== ""
                }
            }
        }
    }
}
