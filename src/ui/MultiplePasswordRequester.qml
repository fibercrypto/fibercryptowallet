import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/Controls"
import "Controls" // For quick UI development, switch back to resources when making a release

Item {
    id: root

    property alias text: textFieldPassword.text
    property alias placeholderText: textFieldPassword.placeholderText
    property alias allowPasswordRecovery: buttonForgot.visible

    signal passwordForgotten()

    function forceTextFocus() {
        textFieldPassword.forceActiveFocus()
    }

    function clear() {
        textFieldPassword.clear()
    }

    implicitHeight: textFieldPassword.implicitHeight + buttonForgot.implicitHeight

    ColumnLayout {
        anchors.fill: parent

        TextField {
            id: textFieldPassword

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
}
