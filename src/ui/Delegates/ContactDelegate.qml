import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import "../Dialogs"

ItemDelegate {
    id: contactDelegate

    signal editRequested()
    signal deleteRequested()

    contentItem: RowLayout {
        spacing: 20
        Image {
            Layout.alignment: Qt.AlignLeft | Qt.AlignVCenter
            source: "qrc:/images/resources/images/icons/user.svg"
            sourceSize: "24x24"
        }
        Label {
            Layout.fillWidth: true
            text: name
            font.bold: true
            elide: Text.ElideRight
        }
        ToolButton {
            id: toolButtonEdit
            icon.source: "qrc:/images/resources/images/icons/edit.svg"

            onClicked: editRequested()
        }
        ToolButton {
            id: toolButtonDelete
            icon.source: "qrc:/images/resources/images/icons/delete.svg"
            icon.color: Material.accent
            Material.accent: Material.Red

            onClicked: deleteRequested()
        }
    }
}
