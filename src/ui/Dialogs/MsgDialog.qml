import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Dialog {
    id: msgDialog

    property alias text: msg.text
    property alias imagePath: icon.source
    property alias confirmationCheckBoxVisible: checkBoxConfirmation.visible
    property alias confirmationCheckBoxText: checkBoxConfirmation.text
    property alias confirmationCheckBoxChecked: checkBoxConfirmation.checked
    property int buttonToConnectWithConfirmationCheckBox: Dialog.Ok

    title: Qt.application.name
    standardButtons: Dialog.Ok

    Component.onCompleted: {
        if (confirmationCheckBoxVisible && standardButton(buttonToConnectWithConfirmationCheckBox)) {
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

    Flickable {
        id: flickable
        anchors.fill: parent
        contentHeight: columnLayoutRoot.height
        clip: true

        ColumnLayout {
            id: columnLayoutRoot
            width: parent.width
            spacing: 30

            RowLayout {
                spacing: 30
                Layout.fillWidth: true

                Image {
                    id: icon
                    source: "qrc:/images/resources/images/icons/appIcon/appIcon.png"
                    sourceSize: "64x64"
                    Layout.alignment: Qt.AlighLeft | Qt.AlignTop
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
        } // ColumnLayout (root)

        ScrollIndicator.vertical: ScrollIndicator {
            parent: msgDialog.contentItem
            anchors.top: flickable.top
            anchors.bottom: flickable.bottom
            anchors.right: parent.right
            anchors.rightMargin: -msgDialog.rightPadding + 1
        }
    } // Flickable
}
