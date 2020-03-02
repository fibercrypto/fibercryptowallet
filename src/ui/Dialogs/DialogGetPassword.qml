import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12

// Resource imports
// import "qrc:/ui/src/ui/"
import "../" // For quick UI development, switch back to resources when making a release

Dialog {
    id: dialogGetPassword

    property alias headerMessage: labelHeaderMessage.text
    property alias headerMessageColor: labelHeaderMessage.color
    property alias password: passwordRequester.text

    property url forgottenPasswordUrl: "http://skycoin.com/"

    function clear() {
        passwordRequester.clear()
        standardButton(Dialog.Ok).enabled = password
    }

    onAboutToShow: {
        clear()
        passwordRequester.forceTextFocus()
    }

    title: qsTr("Password requested")
    standardButtons: Dialog.Ok | Dialog.Cancel

    ColumnLayout {
        id: columnLayoutRoot
        width: parent.width
        spacing: 10
        clip: true

        Label {
            id: labelHeaderMessage

            Layout.fillWidth: true
            wrapMode: Text.Wrap
            visible: text
        }

        PasswordRequester {
            id: passwordRequester

            Layout.fillWidth: true

            onPasswordForgotten: {
                Qt.openUrlExternally(forgottenPasswordUrl)
            }
            onTextChanged: {
                dialogGetPassword.standardButton(Dialog.Ok).enabled = text !== ""
                password = text
            }
        }
    } // ColumnLayout (root)
}
