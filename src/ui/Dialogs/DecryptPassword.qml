import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12

Dialog {
    id: dialogRequestPassword

    modal: true
    title: qsTr("Encrypt Wallet")
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
                   text: "Warning: for security reasons, it is not recommended to keep the wallets unencrypted. Caution is advised."
                   color: "red"
                   Layout.alignment: Qt.AlignCenter
                   Layout.preferredWidth: parent.width - 10
                   wrapMode: Text.Wrap
               }
               Label {
                   text: "Password"
                   Layout.preferredWidth: parent.width - 10
                   Layout.alignment: Qt.AlignCenter
               }
               TextField {
                   id: password
                   Layout.alignment: Qt.AlignCenter
                   Layout.preferredWidth: parent.width - 10
                   font {
                       pixelSize: 20
                   }
                   placeholderText: qsTr("Enter Password")
                   passwordCharacter: "*"
                   echoMode: TextInput.Password
                   onTextEdited: {
                       if (confirmPassword.text.length == 0) {
                           buttonOk.enabled = false
                       } else {
                           buttonOk.enabled = true
                       }
                   }
               }
               Text {
                   textFormat: Text.RichText
                   text: "<a href=\"http://skycoin.net\">I forgot my password</a>."
                   linkColor: "blue"
                   onLinkActivated: console.log(link + " link activated")
               }
               Button {
                   id: buttonOk
                   Layout.alignment: Qt.AlignCenter
                   text: "Ok"
                   enabled: false
                   onClicked: {
                       if(enabled){
                           // TODO whatever to do with accept
                           password.clear()
                           confirmPassword.clear()
                           dialogRequestPassword.close()

                       }
                   }
               }
           } // ColumnLayout (root)
    }
}
