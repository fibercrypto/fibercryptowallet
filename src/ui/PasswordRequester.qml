import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Item {
    id: root

    property alias text: textFieldPassword.text
    property alias placeholderText: textFieldPassword.placeholderText

    implicitHeight: textFieldPassword.implicitHeight + buttonForgot.implicitHeight

    function forceTextFocus() {
        textFieldPassword.forceActiveFocus()
    }

    function clear() {
        textFieldPassword.clear()
    }

    ColumnLayout {
        anchors.fill: parent

        TextField {
            id: textFieldPassword

            placeholderText: qsTr("Password")
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
        }
    }
}
