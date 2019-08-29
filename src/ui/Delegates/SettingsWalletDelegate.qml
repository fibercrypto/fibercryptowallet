import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Item {
    id: settingsDelegate

    property bool expanded: false
    property bool animateDisplacement: false

    implicitHeight: itemDelegateSettingsName.height + (expanded ? listViewSettings.height : 0)
    Behavior on implicitHeight { NumberAnimation { duration: animateDisplacement ? 250 : 0; easing.type: Easing.OutQuint; onRunningChanged: animateDisplacement = false } }

    clip: true

    ItemDelegate {
        id: itemDelegateSettingsName

        property color textColor: expanded ? parent.Material.accent : parent.Material.foreground
        Behavior on textColor { ColorAnimation {} }

        anchors.top: parent.top
        anchors.left: parent.left
        anchors.right: parent.right

        Material.foreground: textColor
        text: name // a role of the model
        font.bold: true

        onClicked: {
            animateDisplacement = true
            expanded = !expanded
        }
    } // ItemDelegate

    ListView {
        id: listViewSettings
        anchors.top: itemDelegateSettingsName.bottom
        anchors.left: parent.left
        anchors.right: parent.right
        height: contentItem.height

        opacity: expanded ? 1.0 : 0.0
        Behavior on opacity { NumberAnimation { duration: expanded ? 250 : 1000; easing.type: Easing.OutQuint } }

        clip: true
        interactive: false
        model: modelOptions

        delegate: OutputsListAddressDelegate {
            width: listViewSettings.width
        }
    } // ListView

    // Roles: address
    // Use listModel.append( { "address": value } )
    // Or implement the model in the backend (a more recommendable approach)
    ListModel {
        id: modelOptions
        // The first element must exist but will not be used
        ListElement {
            //
            operationName: "Wallet Source"
            walletsSource: "$HOME/.skycoin/wallets"
            nodeDirection: ""
        }
        ListElement {
            //
            operationName: "Node Direction"
            nodeDirection: ""
            walletsSource: ""
        }
    }
}
