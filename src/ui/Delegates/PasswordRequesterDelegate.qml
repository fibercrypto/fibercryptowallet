import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import OutputsModels 1.0


Item {
    id: passwordRequesterDelegate

     TextField {
            id: textFieldPassword
            text: name
            placeholderText: qsTr("Password")
            selectByMouse: true
            echoMode: TextField.Password
            focus: true
            Layout.alignment: Qt.AlignTop
            Layout.fillWidth: true
        }

        Button {
            id: buttonForgot
            text: qsTr("I forgot my password")
            flat: true
            highlighted: hovered
            Layout.alignment: Qt.AlignHCenter | Qt.AlignTop
            Layout.fillWidth: true
            
            onClicked: {
                passwordForgotten()
            }
        }
}
