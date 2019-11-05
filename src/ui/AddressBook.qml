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
    GroupBox{
     anchors.fill: parent
            anchors.topMargin: 30
            anchors.bottomMargin: 40
            anchors.rightMargin: 30
            anchors.leftMargin: 30
            clip: true

    DialogAddContact{
    id: contactDialog
 anchors.centerIn: Overlay.overlay

        modal: true
        focus: true

        width: applicationWindow.width > 540 ? 540 - 200 : applicationWindow.width - 200
        height: applicationWindow.height > 640 ? 640 - 200 : applicationWindow.height - 200

        onAccepted: {
            console.log("Add contact")
            }
    }

AddrsBookModel{
id:abm
}

       ScrollView {
                anchors.fill: parent
                clip: true
                ListView {
                    id: addrsBook
                    Layout.fillWidth: true
                    Layout.fillHeight: true
                    model: abm
                    delegate: ItemDelegate{
                    text: name+"\t\t\t"+address
                    width: parent.width
                    }
                }
       }// GroupBox
  RoundButton {
          text: qsTr("+")
          highlighted: true
          anchors.margins: 10
          anchors.right: parent.right
          anchors.bottom: parent.bottom
          onClicked: {
                    contactDialog.open()
          }
      }
     }

     ListModel {
             id: listContacts
             ListElement { name: "My first wallet"; address:"qrxw7364w8xerusftaxkw87ues" }
             ListElement { name: "My second wallet"; address:"8745yuetsrk8tcsku4ryj48ije"}
             ListElement { name: "My third wallet";  address:"gfdhgs343kweru38200384uwqd"}
             ListElement { name: "My first wallet";  address:"00qdqsdjkssvmchskjkxxdg374"}
         }
}
