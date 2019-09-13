import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12

Dialog {
    id: dialogSetPassword

    property alias headerMessage: labelHeaderMessage.text
    property alias headerMessageColor: labelHeaderMessage.color
    property alias password: textFieldPassword.text
    readonly property bool acceptablePassword: password && textFieldPassword.text === textFieldPasswordConfirmation.text
    onAcceptablePasswordChanged: {
        standardButton(Dialog.Ok).enabled = acceptablePassword
    }

    function clear() {
        textFieldPassword.clear()
        textFieldPasswordConfirmation.clear()
        standardButton(Dialog.Ok).enabled = acceptablePassword
    }

    onAboutToShow: {
        clear()
        textFieldPassword.forceActiveFocus()
    }

    title: qsTr("Set a password")
    standardButtons: Dialog.Ok | Dialog.Cancel

    Flickable {
        id: flickable
        anchors.fill: parent
        contentHeight: columnLayoutRoot.height
        clip: true

        ColumnLayout {
            id: columnLayoutRoot
            width: parent.width
            spacing: 10

            Label {
                id: labelHeaderMessage

                Layout.fillWidth: true
                wrapMode: Text.WordWrap
                visible: text
            }

            TextField {
                id: textFieldPassword

                Layout.fillWidth: true
                placeholderText: qsTr("Password")
                echoMode: TextField.Password
                focus: true
            }

            TextField {
                id: textFieldPasswordConfirmation

                Layout.fillWidth: true
                placeholderText: qsTr("Confirm password")
                echoMode: TextField.Password
                Material.accent: text === textFieldPassword.text ? parent.Material.accent : Material.Red
            }
        } // ColumnLayout (root)
    } // Flickable
}
