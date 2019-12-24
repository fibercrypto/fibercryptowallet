import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12


import "../Controls" // For quick UI development, switch back to resources when making a release
import "../" // For quick UI development, switch back to resources when making a release
import "../Delegates"


Dialog{
  id: dialogAddContact
  property bool isEdit:false

  title: Qt.application.name
  standardButtons: Dialog.Ok | Dialog.Cancel
    Component.onCompleted: {
    standardButton(Dialog.Ok).enabled=false
    }
    onAboutToShow:{
if(isEdit){
    name.text=menu.name
listModelAddresses.clear()
for(var i=0;i<menu.address.rowCount();i++){
listModelAddresses.append({value:menu.address.address[i].value,
coinType:menu.address.address[i].coinType})
}
}else{
    name.text=""
listModelAddresses.append({value:"",coinType:""})
}
}
    onAccepted:{
updateAcceptButtonStatus()
    name.text=""
    listModelAddresses.clear()
    }

    onRejected:{
    name.text=""
    listModelAddresses.clear()
//    listModelAddresses.append( { value: "", coinType: "" } )
    }

function enableOkBtn(){
for(var i=0;i<listModelAddresses.count;i++){
if (!(addrsBkModel.addressIsValid(listModelAddresses.get(i).value))||
     name.text==""|| addrsBkModel.nameExist(menu.index,name.text)||
     addrsBkModel.addressExist(menu.index,listModelAddresses.get(i).value,listModelAddresses.get(i).coinType)){
standardButton(Dialog.Ok).enabled=false
return
}
}
standardButton(Dialog.Ok).enabled=true
}

    function updateAcceptButtonStatus() {
    for(var i=0;i<listModelAddresses.count;i++){
    addrsBkModel.addAddress(listModelAddresses.get(i).value,listModelAddresses.get(i).coinType)
    }

    if (isEdit){
addrsBkModel.editContact(menu.index, menu.cId, name.text)
    }else{
    addrsBkModel.newContact(name.text)
    }
    } // function updateAcceptButtonStatus()



Flickable{
        id:flickable
        anchors.fill: parent
        contentHeight: columnLayoutRoot.height
        clip: true
        ColumnLayout{
            id: columnLayoutRoot
            width: parent.width
            spacing: 30

                    Behavior on Layout.preferredHeight {NumberAnimation{duration: 500;easing.type:Easing.OutQuint}}
                    TextField{
                        id:name
                        placeholderText: qsTr("Name")
                        Layout.fillWidth: true
                        text: menu.name
                        onTextChanged:{
                        enableOkBtn()
                        }

                    }
                   ColumnLayout {
                               id: columnLayoutDestinations

                               Layout.alignment: Qt.AlignTop

                               ListView {
                                   id: listViewDestinations

                                   property real delegateHeight: 47

                                   Layout.fillWidth: true
                                   Layout.topMargin: -16
                                   implicitHeight: count * delegateHeight

                                   Behavior on implicitHeight { NumberAnimation { duration: 250; easing.type: Easing.OutQuint } }

                                   interactive: false
                                   clip: true

                                   model: listModelAddresses

                                   delegate: AddressListDelegate {
                                       width: listViewDestinations.width
                                       implicitHeight: ListView.view.delegateHeight
                                   }
                               } // ListView
                           } // ColumnLayout (destinations)
ListModel {
        id: listModelAddresses
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