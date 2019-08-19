import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12

Dialog {
    id: dialogRequestPassword

    modal: true
    title: qsTr("Encrypt Wallet")
    standardButtons: Dialog.Close
    Flickable {
           id: flickable
           anchors.fill: parent
           contentHeight: columnLayoutRoot.height
           clip: true
           ColumnLayout {
               id: columnLayoutRoot
               width: parent.width
               spacing: 30
               Label {
                   text: "We suggest that you encrypt each one of your wallets with a password. If you forget your password, you can reset it with your seed. Make sure you have your seed saved somewhere safe before encrypting your wallet."
               }
               Label {
                   text: "Password"
               }
               TextField {
                   id: password
//                   anchors.fill: parent
                   font {
                       pixelSize: 20
                   }
                   passwordCharacter: "*"
                   echoMode: TextInput.PasswordEchoOnEdit
               }
               Label {
                   text: "Confirm Password"
               }
               TextField {
                   id: confirmPassword
//                   anchors.fill: parent
                   font {
                       pixelSize: 20
                   }
                   passwordCharacter: "*"
                   echoMode: TextInput.PasswordEchoOnEdit
               }

           } // ColumnLayout (root)
    }
}
