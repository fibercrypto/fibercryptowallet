import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "../Delegates" // For quick UI development, switch back to resources when making a release

Dialog {
    id:dialogCreateAddressBook

    property int select: 0

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
                id: columnRadioButtons
                Layout.fillWidth:true
                RadioButton {
                    property int pos:0
                    checked: true
                    text: qsTr("Low (Plain text)")
                }
                RadioButton {
                    property int pos:1
                    text: qsTr("Medium (Recommended)")
                }
                RadioButton {
                    property int pos:2
                    Layout.fillWidth:true
                    text: qsTr("Hard (With password)\n"+
                          "(This can slow your device)")
                  }
              }
        } // ColumnLayout (root)
    } // Flickable

    ButtonGroup {
        id: buttonsGroup
        buttons: columnRadioButtons.children
        onClicked: {
            select = checkedButton.pos
        }
    }
}