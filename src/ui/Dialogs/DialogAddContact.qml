import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import AddrsBookManager 1.0


import "../Controls" // For quick UI development, switch back to resources when making a release
import "../" // For quick UI development, switch back to resources when making a release

Dialog{
  id: dialogAddContact
  title: Qt.application.name
  standardButtons: Dialog.Ok | Dialog.Cancel
    Component.onCompleted: {
        standardButton(Dialog.Ok).enabled=false
    }


    onAccepted:{
updateAcceptButtonStatus()

    }

    function updateAcceptButtonStatus() {
abm.newContact(name.text,address.text)
    } // function updateAcceptButtonStatus()


    onAboutToShow: {
//        createLoadWallet.clear()
    }

Flickable{
        id:flickable
        anchors.fill: parent
        contentHeight: columnLayoutRoot.height
        clip: true
        ColumnLayout{
            id: columnLayoutRoot
            //anchors.fill: parent
            width: parent.width
            spacing: 30

                    Behavior on Layout.preferredHeight {NumberAnimation{duration: 500;easing.type:Easing.OutQuint}}
                    TextField{
                        id:name
                        placeholderText: "name"
                        Layout.fillWidth: true
                        onTextChanged:{
                        standardButton(Dialog.Ok).enabled=(name.text!="")&&(address.text!="")
                        }

                    }
                    TextField{
                        id:address
                        placeholderText: "address"
                        Layout.fillWidth: true
                        onTextChanged:{
                        standardButton(Dialog.Ok).enabled=(name.text!="")&&(address.text!="")
                        }

                    }
        }//ColumnLayoutRoot

        ScrollIndicator.vertical: ScrollIndicator{
        parent: dialogAddContact.contentItem
        anchors.top: flickable.top
        anchors.bottom: flickable.bottom
        anchors.right: parent.right
        anchors.rightMargin: -dialogAddContact.rightMargin+1
        }
    }//Flickable
}