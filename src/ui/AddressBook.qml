import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import AddrsBookManager 1.0

// Resource imports
import "Delegates/" // For quick UI development, switch back to resources when making a release
import "Dialogs/" // For quick UI development, switch back to resources when making a release

Page{
    id:addressBook

//    property int currentContact: -1
      property int secType
      property string password

    DialogAddContact{
    id: contactDialog
 anchors.centerIn: Overlay.overlay
        focus: true
        width: applicationWindow.width > 540 ? 540  : applicationWindow.width
        height: applicationWindow.height > 640 ? 640: applicationWindow.height

    }

 Component.onCompleted: {
     if(abm.hasInit()){
     if(abm.getSecType()!=2){
        abm.loadContacts()
     }else{
     getpass.open()
     }

     }else{
        dialogCreateAddrsBk.open()
     }
     }

AddrsBookModel{
    id:abm
}

DialogSelectSecType{
id:dialogCreateAddrsBk
width:300
height:300
anchors.centerIn: Overlay.overlay
onAccepted:{
secType = select
if (secType==2){
setpass.open()
}else{
abm.initAddrsBook(secType,"")
}
}
onRejected:{
 generalStackView.pop()
}

}

DialogSetPassword{
id:setpass
anchors.centerIn: Overlay.overlay
onAccepted:{
abm.initAddrsBook(2,setpass.password)
}
onRejected:{
generalStackView.pop()
}
}

DialogGetPassword{
id:getpass
anchors.centerIn: Overlay.overlay
height:180
onAccepted:{
if(!abm.authenticate(getpass.password)){
getpass.open()
}else{
abm.loadContacts()
}
}

onRejected:{
generalStackView.pop()
console.log("asd")
}
}
       ScrollView {
                anchors.fill: parent
                clip: true
                ListView {

                    id: addrsBook
                    Layout.fillWidth: true
                    Layout.fillHeight: true
                    model: abm
                    section.property: "name"
                    section.criteria: ViewSection.FirstCharacter
                    section.delegate: SectionDelegate {
                        width: addrsBook.width
                    }

                    delegate: ContactDelegate{
                   id:contactDelegate
                     width: addrsBook.width
                    }
                }
       }// ScrollView
  RoundButton {
          text: qsTr("Add Contact")
          highlighted: true
          anchors.margins: 10
          anchors.right: parent.right
          anchors.bottom: parent.bottom
          onClicked: {
          contactDialog.isEdit=false
                    contactDialog.open()
          }
      }//RoundButton

Menu{
            id:menu
            property int index: -1
            property int cId: -1
            property string name
            property AddrsBkAddressModel address
            MenuItem{
                text: "&View"
                onTriggered: {
                console.log("Show Contact")
                    console.log(menu.name)
                    dialogShowContact.open()
                }
            }
            MenuSeparator{}

            MenuItem{
                text: "&Edit"
                onTriggered: {
                contactDialog.isEdit=true
                 contactDialog.open()
                  }
            }
             MenuSeparator{}

            MenuItem{
                text: "&Remove"
                onTriggered: {
            abm.removeContact(menu.index,menu.cId)
                }
            }
        }//Menu

DialogShowContact{
id:dialogShowContact
anchors.centerIn: Overlay.overlay
width:350
focus:true
height:400
}


//     ListModel {
//             id: listContacts
//             ListElement { name: "My first wallet"; address:"qrxw7364w8xerusftaxkw87ues" }
//             ListElement { name: "My second wallet"; address:"8745yuetsrk8tcsku4ryj48ije"}
//             ListElement { name: "My third wallet";  address:"gfdhgs343kweru38200384uwqd"}
//             ListElement { name: "My first wallet";  address:"00qdqsdjkssvmchskjkxxdg374"}
//         }
}
