import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Dialog {
    id: root

    property alias text: msg.text
    property alias imagePath: icon.source
    property alias confirmationCheckBoxVisible: checkBoxConfirmation.visible
    property alias confirmationCheckBoxText: checkBoxConfirmation.text
    property alias confirmationCheckBoxChecked: checkBoxConfirmation.checked
    property int buttonToConnectWithConfirmationCheckBox: Dialog.Ok

    focus: true
    modal: true
    title: Qt.application.name
    standardButtons: Dialog.Ok

    Component.onCompleted: {
        if (standardButton(buttonToConnectWithConfirmationCheckBox)) {
            standardButton(buttonToConnectWithConfirmationCheckBox).enabled = confirmationCheckBoxChecked
        }
    }

    onConfirmationCheckBoxCheckedChanged: {
        if (confirmationCheckBoxVisible && standardButton(buttonToConnectWithConfirmationCheckBox)) {
            standardButton(buttonToConnectWithConfirmationCheckBox).enabled = confirmationCheckBoxChecked
        } else if (standardButton(buttonToConnectWithConfirmationCheckBox)) {
            standardButton(buttonToConnectWithConfirmationCheckBox).enabled = true
        }
    }

    ColumnLayout {
        anchors.fill: parent
        spacing: 30

        RowLayout {
            spacing: 30

            Image {
                id: icon
                sourceSize: "64x64"
            }
            Label {
                id: msg
                text: qsTr("Your message goes here.")
                Layout.fillWidth: true
                Layout.alignment: Qt.AlignVCenter
                wrapMode: Text.WordWrap
            }
        }

        CheckBox {
            id: checkBoxConfirmation
            visible: false
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignHCenter
        }
    }

}
