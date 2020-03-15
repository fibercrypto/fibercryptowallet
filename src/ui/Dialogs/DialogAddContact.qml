import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12


import AddrsBookManager 1.0

// Resource imports
// import "qrc:/ui/src/ui/Controls"
// import "qrc:/ui/src/ui"
// import "qrc:/ui/src/ui/Delegates"
import "../Controls" // For quick UI development, switch back to resources when making a release
import "../" // For quick UI development, switch back to resources when making a release
import "../Delegates" // For quick UI development, switch back to resources when making a release

Dialog {
    id: dialogAddContact

    property alias name: textFieldName.text
    property int index: -1
    property int cId: -1
    property AddrsBkAddressModel addressModel
    property alias listModelAddresses: listModelAddresses
    property bool isEdit: false

    title: isEdit ? qsTr("Edit contact") : qsTr("Add contact")
    standardButtons: Dialog.Save | Dialog.Cancel

    Component.onCompleted: {
        standardButton(Dialog.Save).enabled = false
        standardButton(Dialog.Save).text = Qt.binding( function() { return isEdit ? qsTr("Save") : qsTr("Add") } )
    }

    onAboutToShow: {
        listModelAddresses.clear()
        if (isEdit) {
            for (var i = 0; i < addressModel.rowCount(); i++) {
                listModelAddresses.append( {
                    value: addressModel.address[i].value,
                    coinType: addressModel.address[i].coinType
                })
            }
        } else {
            name = ""
            listModelAddresses.append( { value: "", coinType: "" } )
        }

        updateAcceptButtonStatus()
    }

    function updateAcceptButtonStatus() {
        for (var i = 0; i < listModelAddresses.count; i++) {
            if ( !(abm.addressIsValid(listModelAddresses.get(i).value)) ||
                   name === "" || abm.nameExist(index, name) ||
                   abm.addressExist(index, listModelAddresses.get(i).value, listModelAddresses.get(i).coinType)) {
                standardButton(Dialog.Save).enabled = false
                return
            }
        }
        standardButton(Dialog.Save).enabled = true
    }

    ColumnLayout {
        anchors.fill: parent

        TextField {
            id: textFieldName
            Layout.fillWidth: true

            placeholderText: qsTr("Name")
            selectByMouse: true
            onTextChanged: {
                updateAcceptButtonStatus()
            }
        }

        ScrollView {
            id: scrollView
            Layout.fillWidth: true
            Layout.fillHeight: true

            ListView {
                id: listViewDestinations

                interactive: false
                clip: true

                model: listModelAddresses

                delegate: AddressListDelegate {
                    width: listViewDestinations.width

                    onAddressTextChanged: {
                        updateAcceptButtonStatus()
                    }
                    onNumberOfAddressesChanged: {
                        updateAcceptButtonStatus()
                    }
                }
            } // ListView
        } // ScrollView
    } // ColumnLayout

    ListModel {
        id: listModelAddresses
    }
}