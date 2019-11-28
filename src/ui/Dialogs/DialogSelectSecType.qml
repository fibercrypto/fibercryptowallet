import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12

import "../Delegates"

Dialog{

property int select:0

id:dialogCreateAddressBook
 standardButtons: Dialog.Ok | Dialog.Close


Flickable {
        id: flickable
        anchors.fill: parent
        contentHeight: columnLayoutRoot.height
        clip: true

        ColumnLayout {
            id: columnLayoutRoot
            width: parent.width
            spacing: 10

            Label {
                id: labelHeaderMessage
                text: qsTr("Select security type for your Address Book:")
                Layout.fillWidth: true
                wrapMode: Text.WordWrap
                visible: text
            }

             ColumnLayout {
                Layout.fillWidth:true
                 id: column
                  RadioButton {
                  property int pos:0
                      checked: true
                      text: qsTr("Low (Nothing)")
                  }
                  RadioButton {
                  property int pos:1
                      text: qsTr("Medium (Recommended)")
                  }
                  RadioButton {
                  property int pos:2
                  Layout.fillWidth:true
                      text: qsTr("Hard (With password)\n"+
                      "(This can slow your dispositive)")
                  }
              }
        } // ColumnLayout (root)
    } // Flickable

    ButtonGroup {
        id:buttonsGroup
        buttons: column.children
       onClicked:{
       select=checkedButton.pos
       }
       }
       }