import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12
import "../Dialogs"
ItemDelegate {
    id: contactDelegate

onClicked:{
menu.index=index
console.log(id)
menu.cId=id
menu.name=name
menu.address=address
menu.popup()
}
    checkable: true

    contentItem: RowLayout {
        spacing: 20
        Image{
        source:"../../../resources/images/icons/user_icon-icons.com_66546.svg"
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
}
