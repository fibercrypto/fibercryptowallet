import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12
import "../Dialogs"
ItemDelegate {
    id: contactDelegate

onClicked:{
//abm.removeContact(1,1)
menu.index=index
menu.id=id
menu.popup()
}
    checkable: true

DialogShowContact{
id:dialogShowContact
anchors.centerIn: Overlay.overlay
width:350
height:400
}
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

Menu{
            id:menu
            property int index: -1
            property int id: -1
            MenuItem{
                text: "&View"
                onTriggered: {
                console.log("Show Contact")
                    dialogShowContact.open()
                }
            }
            MenuSeparator{}

            MenuItem{
                text: "&Edit"
                onTriggered: {
                  abm.editContact(1,"mock","mock")
                }
            }
             MenuSeparator{}

            MenuItem{
                text: "&Remove"
                onTriggered: {
//                console.log("Remove Contact")
            abm.removeContact(menu.index,id)
                }
            }
        }//Menu
}
