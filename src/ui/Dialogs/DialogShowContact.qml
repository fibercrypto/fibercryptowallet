import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "../Delegates" // For quick UI development, switch back to resources when making a release

Dialog {
     id: dialogShowContact
     standardButtons: Dialog.Close

     header: ColumnLayout {
          Image {
               Layout.topMargin: 20
               source: "qrc:/images/resources/images/icons/user.svg"
          }

          Label {
               Layout.topMargin: 20
               text: menu.name
               font.bold: true
               horizontalAlignment: Qt.AlignHCenter
               verticalAlignment: Qt.AlignVCenter
               Layout.fillWidth: true
          }
     }

     ScrollView {
          anchors.fill: parent
          clip: true
          ListView {
               id: addresses
               Layout.fillWidth: true
               Layout.fillHeight: true
               model: menu.address
               section.property: "coinType"
               section.criteria: ViewSection.FullString
               section.delegate: SectionDelegate {
                    width: addresses.width
               }
               delegate: ItemDelegate{
                    id:addressdelegate
                    Layout.fillHeight: true
                    width:parent.width
                    text:value
               }    
          }
     } // ScrollView
}