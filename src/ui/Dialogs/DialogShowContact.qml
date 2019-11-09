import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12


Dialog{
id:dialogShowContact

//Component.onCompleted:{
//console.log(address)
//}

 standardButtons: Dialog.Ok
 Label{
 text:name
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

                    id: addresses
                    Layout.fillWidth: true
                    Layout.fillHeight: true
                    model: address

                    delegate: ItemDelegate{
                   id:addressdelegate
                   text:value
//                     width: addrsBook.width

//        Connections {
//            target: contactDelegate
//            onPressAndHold: addrsBook.pressAndHold(index)
//        }
                    }
                }
       }// ScrollView

       }