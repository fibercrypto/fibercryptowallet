import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12
import "../Dialogs"

ItemDelegate {
    id: contactDelegate

    checkable: true

    contentItem: RowLayout {
        spacing: 20
        Image {
        source: "qrc:/images/resources/images/icons/user.svg"
        sourceSize.width: 24
            sourceSize.height:24
        }
        Label {
            text: name
            font.bold: true

            elide: Text.ElideRight
            Layout.fillWidth: true
        }
    }

    onClicked: {
        menu.index = index
        console.log(id)
        menu.cId = id
        menu.name = name
        menu.address = address
        menu.popup()
    }
}
