import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12

Dialog {
    id: dialogRequestPassword

    modal: true
    title: qsTr("Encrypt Wallet")
    standardButtons: Dialog.Close | Dialog.Ok
    Flickable {
           id: flickable
           anchors.fill: parent
           contentHeight: columnLayoutRoot.height
           clip: true
           ColumnLayout {
               id: columnLayoutRoot
               width: parent.width - 10
               spacing: 30
               Text {
                   text: "We suggest that you encrypt each one of your wallets with a password. If you forget your password, you can reset it with your seed. Make sure you have your seed saved somewhere safe before encrypting your wallet."
                   color: "red"
                   Layout.alignment: Qt.AlignCenter
                   Layout.preferredWidth: parent.width - 10
                   wrapMode: Text.Wrap
//                   autoFit: TextAutoFit.FitToBounds
               }
               Label {
                   text: "Password"
                   Layout.preferredWidth: parent.width - 10
                   Layout.alignment: Qt.AlignCenter
               }
               TextField {
                   Layout.alignment: Qt.AlignCenter
                   Layout.preferredWidth: parent.width - 10
                   id: password
//                   anchors.fill: parent
                   font {
                       pixelSize: 20
                   }
                   passwordCharacter: "*"
                   echoMode: TextInput.Password
               }
               Label {
                   text: "Confirm Password"
                   Layout.preferredWidth: parent.width - 10
                   Layout.alignment: Qt.AlignCenter
               }
               TextField {
                   id: confirmPassword
                   Layout.preferredWidth: parent.width - 10
                   Layout.alignment: Qt.AlignCenter
//                   anchors.fill: parent
                   font {
                       pixelSize: 20
                   }
                   width: parent.width - 20
                   height: parent.height - 20
                   passwordCharacter: "*"
                   echoMode: TextInput.Password
                   onEditingFinished: {

                   }
               }

           } // ColumnLayout (root)
    }
}
