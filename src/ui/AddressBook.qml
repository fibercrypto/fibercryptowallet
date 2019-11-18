import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import AddrsBookManager 1.0

// Resource imports
// import "qrc:/ui/src/ui/Dialogs"
import "Delegates/" // For quick UI development, switch back to resources when making a release
import "Dialogs/" // For quick UI development, switch back to resources when making a release

Page{
    id:addressBook

    property int currentContact: -1


    DialogAddContact{
    id: contactDialog
 anchors.centerIn: Overlay.overlay

        modal: true
        focus: true

        width: applicationWindow.width > 540 ? 540  : applicationWindow.width
        height: applicationWindow.height > 640 ? 640: applicationWindow.height

        onAccepted: {
//            console.log("Add contact")
            }
    }

 Component.onCompleted: {
     if(abm.exist()){
     getpass.open()
     }else{
     setpass.open()
     }
     }

AddrsBookModel{
    id:abm
}

DialogSetPassword{
id:setpass
anchors.centerIn: Overlay.overlay
onAccepted:{
abm.initAddrsBook(setpass.password)
}
onRejected:{
generalStackView.pop()
console.log("asd")
}
}

DialogGetPassword{
id:getpass
anchors.centerIn: Overlay.overlay
height:180
onAccepted:{
 if(!abm.openAddrsBook(getpass.password)){
 getpass.open()
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
                    contactDialog.open()
          }
      }//RoundButton


//     ListModel {
//             id: listContacts
//             ListElement { name: "My first wallet"; address:"qrxw7364w8xerusftaxkw87ues" }
//             ListElement { name: "My second wallet"; address:"8745yuetsrk8tcsku4ryj48ije"}
//             ListElement { name: "My third wallet";  address:"gfdhgs343kweru38200384uwqd"}
//             ListElement { name: "My first wallet";  address:"00qdqsdjkssvmchskjkxxdg374"}
//         }
}
