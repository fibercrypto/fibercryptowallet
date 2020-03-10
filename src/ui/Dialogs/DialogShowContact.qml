import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12

import AddrsBookManager 1.0

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "../Delegates" // For quick UI development, switch back to resources when making a release

Dialog {
     id: dialogShowContact

     property alias name: labelName.text
     property AddrsBkAddressModel addressModel

     standardButtons: Dialog.Close

     header: ColumnLayout {
          Image {
               Layout.topMargin: 20
               Layout.alignment: Qt.AlignVCenter | Qt.AlignHCenter
               source: "qrc:/images/resources/images/icons/user.svg"
               sourceSize: "128x128"
          }

          Label {
               id: labelName
               Layout.fillWidth: true
               Layout.topMargin: 20
               horizontalAlignment: Qt.AlignHCenter
               font.bold: true
          }
     }

     ScrollView {
          anchors.fill: parent
          clip: true

          ListView {
               id: addresses

               model: addressModel
               delegate: ItemDelegate {
                    id: addressdelegate
                    width: parent.width
                    text: value
               }    
               section.property: "coinType"
               section.criteria: ViewSection.FullString
               section.delegate: SectionDelegate {
                    width: addresses.width
               }
          }
     } // ScrollView
}