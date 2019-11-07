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

            header: ColumnLayout {
RowLayout {
            spacing: listWalletSpacing
            Layout.topMargin: 30

            Label {
                text: qsTr("Contacts")
                font.bold:true
                 font.pointSize: 15
                   horizontalAlignment: Qt.AlignHCenter
                        verticalAlignment: Qt.AlignVCenter
                Layout.leftMargin: listWalletLeftMargin
                Layout.fillWidth: true
            }
        } // RowLayout
    Rectangle {
            id: rect
            Layout.fillWidth: true
            height: 1
            color: "#DDDDDD"
        }
             }//ColumnLayout


    DialogAddContact{
    id: contactDialog
 anchors.centerIn: Overlay.overlay

        modal: true
        focus: true

        width: applicationWindow.width > 540 ? 540 - 200 : applicationWindow.width - 200
        height: applicationWindow.height > 640 ? 640 - 200 : applicationWindow.height - 200

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
onAccepted:{
abm.initAddrsBook(setpass.password)
}
}

DialogGetPassword{
id:getpass
height:180
onAccepted:{
 if(!abm.openAddrsBook(getpass.password)){
 getpass.open()
 }
}
}
       ScrollView {
                anchors.fill: parent
                clip: true
                ListView {

//                 signal pressAndHold(int index)
//                    flickDeceleration: 1497
//
//                    focus: true
//                    boundsBehavior: Flickable.StopAtBounds

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

//        Connections {
//            target: contactDelegate
//            onPressAndHold: addrsBook.pressAndHold(index)
//        }
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
      }//RoundButton


//     ListModel {
//             id: listContacts
//             ListElement { name: "My first wallet"; address:"qrxw7364w8xerusftaxkw87ues" }
//             ListElement { name: "My second wallet"; address:"8745yuetsrk8tcsku4ryj48ije"}
//             ListElement { name: "My third wallet";  address:"gfdhgs343kweru38200384uwqd"}
//             ListElement { name: "My first wallet";  address:"00qdqsdjkssvmchskjkxxdg374"}
//         }
}
