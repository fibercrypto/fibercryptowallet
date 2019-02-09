import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Dialog {
    id: root

    property alias text: msg.text
    property alias imagePath: icon.source

    focus: true
    modal: true
    title: Qt.application.name
    standardButtons: Dialog.Ok

    RowLayout {
        anchors.fill: parent
        spacing: 50
        Image {
            id: icon
            sourceSize: "64x64"
            Layout.alignment: Qt.AlignTop
        }
        Label {
            id: msg
            text: qsTr("Your message goes here.")
            wrapMode: Text.WordWrap
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignTop
        }
    }

}
