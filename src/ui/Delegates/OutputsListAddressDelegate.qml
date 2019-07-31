import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Item {
    id: outputsListAddressDelegate

    property bool expanded: false
    
    implicitHeight: itemDelegateAddress.height + (expanded ? listViewAddressOutputs.height : 0)
    Behavior on implicitHeight { NumberAnimation { duration: 250; easing.type: Easing.OutQuint } }

    clip: true

    ItemDelegate {
        id: itemDelegateAddress

        property color textColor: expanded ? parent.Material.accent : parent.Material.foreground
        Behavior on textColor { ColorAnimation {} }

        anchors.top: parent.top
        anchors.left: parent.left
        anchors.leftMargin: 10
        anchors.right: parent.right

        Material.foreground: textColor
        icon.source: "qrc:/images/resources/images/icons/qr.svg"
        icon.width: 16
        icon.height: 16
        text: address // a role of the model
        font.family: "Code New Roman"
        font.bold: expanded

        onClicked: {
            expanded = !expanded
        }
    } // ItemDelegate
    
    ListView {
        id: listViewAddressOutputs
        anchors.top: itemDelegateAddress.bottom
        anchors.left: parent.left
        anchors.right: parent.right
        height: contentItem.height
        
        opacity: expanded ? 1.0 : 0.0
        Behavior on opacity { NumberAnimation { duration: expanded ? 250 : 1000; easing.type: Easing.OutQuint } }

        clip: true
        interactive: false
        model: modelOutputs

        delegate: OutputsListAddressOutputDelegate {
            width: listViewAddressOutputs.width
        }
    } // ListView

    // Roles: outputID, addressSky, addressCoinHours
    // Use listModel.append( { "outputID": value, "addressSky": value, "addressCoinHours": value } )
    // Or implement the model in the backend (a more recommendable approach)
    ListModel {
        id: modelOutputs
        // The first element must exist but will not be used
        ListElement { outputID: "qrxw7364w8xerusftaxkw87ues"; addressSky: 30; addressCoinHours: 1049 }
        ListElement { outputID: "8745yuetsrk8tcsku4ryj48ije"; addressSky: 12; addressCoinHours: 16011 }
        ListElement { outputID: "gfdhgs343kweru38200384uwqd"; addressSky: 0; addressCoinHours: 72 }
        ListElement { outputID: "00qdqsdjkssvmchskjkxxdg374"; addressSky: 521; addressCoinHours: 11 }
    }
}