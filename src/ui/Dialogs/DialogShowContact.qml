import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12

import "../Delegates"

Dialog{

id:dialogShowContact
 standardButtons: Dialog.Close
header:ColumnLayout {
RowLayout{
     anchors.horizontalCenter: parent.horizontalCenter
Image{
 Layout.topMargin: 20
      source:"../../../resources/images/icons/user_icon-icons.com_66546.svg"
 }

}
                RowLayout {
           Layout.topMargin: 20
Label{
 text:name
  font.bold: true
    horizontalAlignment: Qt.AlignHCenter
          verticalAlignment: Qt.AlignVCenter
  Layout.fillWidth: true
 }
}
}

  ScrollView {
                anchors.fill: parent
                clip: true
                ListView {
                    id: addresses
                    Layout.fillWidth: true
                    Layout.fillHeight: true
                    model: address
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
       }// ScrollView

       }