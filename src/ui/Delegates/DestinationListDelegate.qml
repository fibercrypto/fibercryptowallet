import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import AddrsBookManager 1.0

// Resource imports
// import "qrc:/ui/src/ui/"
// import "qrc:/ui/src/ui/Controls"
// import "qrc:/ui/src/ui"
import "../" // For quick UI development, switch back to resources when making a release
import "../Controls" // For quick UI development, switch back to resources when making a release
import "../Dialogs" // For quick UI development, switch back to resources when making a release

Item {
    id: root

    signal qrCodeRequested(var data)

    onQrCodeRequested: {
        dialogQR.setVars(data)
        dialogQR.open()
    }

function getAddressList(){
contactAddrsModel.clear()
console.log(addrsBkModel.count)
for(var i=0;i<addrsBkModel.count;i++){
for(var j=0;j<addrsBkModel.contacts[i].address.address.length;j++){
contactAddrsModel.append({name:addrsBkModel.contacts[i].name,
address:addrsBkModel.contacts[i].address.address[j].value,
coinType:addrsBkModel.contacts[i].address.address[j].coinType})
}
}
}

    implicitHeight: rootLayout.height
    clip: true

    RowLayout {
        id: rootLayout
        width: root.width
        clip: true
        spacing: 20
        opacity: 0.0

        // TODO: Use `add`, `remove`, etc. transitions
        Component.onCompleted: { opacity = 1.0 } // Not the best way to do this
        Behavior on opacity { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } }
            Button {
                    id: buttonSelectCustomChangeAddress
                    text: qsTr("Select")
                    flat: true
                    highlighted: true

                    onClicked: {
                   if(addrsBkModel.getSecType()!=2){
                          addrsBkModel.loadContacts()
                            dialogSelectAddressByAddressBook.open()
                         }else{
                             getpass.open()
                        }
                }
                }
                 DialogGetPassword{
                                 id:getpass
                                 anchors.centerIn: Overlay.overlay
                                 height:180
                                 onAccepted:{
                                 if(!addrsBkModel.authenticate(getpass.password)){
                                 getpass.open()
                                 }else{
                                 addrsBkModel.loadContacts()
                                 dialogSelectAddressByAddressBook.open()
                                 }
                                 }
                                 }

        RowLayout {
            Layout.fillWidth: true
            spacing: 8

            ToolButtonQR {
                id: toolButtonQR

                iconSize: "24x24"

                onClicked: {
                    qrCodeRequested(address)
                }
            }

            TextField {
                id: textFieldDestinationAddress
                font.family: "Code New Roman"
                placeholderText: qsTr("Destination address")
                text: address
                selectByMouse: true
                Layout.fillWidth: true
                Material.accent: addrsBkModel.addressIsValid(text) ? parent.Material.accent : Material.color(Material.Red)
                onTextChanged: address = text
            }

        } // RowLayout

        RowLayout {
            TextField {
                id: textFieldDestinationAmount
                onTextChanged: sky = text
                text: sky
                selectByMouse: true
                implicitWidth: 60
                validator: DoubleValidator {
                    notation: DoubleValidator.StandardNotation
                }
            }
            Label { text: qsTr("SKY") }
        }

        RowLayout {
            visible: !checkBoxAutomaticCoinHoursAllocation.checked

            TextField {
                id: textFieldCoinHoursAmount
                onTextChanged: coinHours = text
                text: coinHours
                selectByMouse: true
                implicitWidth: 60
                validator: DoubleValidator {
                    notation: DoubleValidator.StandardNotation
                }
            }
            Label {
                text: qsTr("Coin hours")
            }
        } // RowLayout

        ToolButton {
            id: toolButtonAddRemoveDestination

            Layout.topMargin: 12
            // The 'accent' attribute is used for button highlighting
            Material.accent: index === 0 ? Material.Teal : Material.Red
            icon.source: "qrc:/images/resources/images/icons/" + (index === 0 ? "add" : "remove") + "-circle.svg"
            highlighted: true // enable the usage of the `Material.accent` attribute

            Layout.alignment: Qt.AlignRight

            onClicked: {
                if (index === 0) {
                    listModelDestinations.append( { "address": "", "sky": "0.0", "coinHours": "0.0" } )
                } else {
                    listModelDestinations.remove(index)
                }
            }

        } // ToolButton (Add/Remove)
    } // RowLayout (rootLayout)

                 DialogSelectAddressByAddressBook{
                            id: dialogSelectAddressByAddressBook

                            anchors.centerIn: Overlay.overlay
                            width: applicationWindow.width > 540 ? 540 - 40 : applicationWindow.width - 40
                            height: applicationWindow.height - 40

                            listAddrsModel: contactAddrsModel
onAboutToShow:{
getAddressList()
}
                            focus: true
                            modal: true

                            onAccepted: {
                                textFieldDestinationAddress.text = selectedAddress
                            }
                        }

ListModel{
id:contactAddrsModel
}
 AddrsBookModel{
    id:addrsBkModel
    }
}
