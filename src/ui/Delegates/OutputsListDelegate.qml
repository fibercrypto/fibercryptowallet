import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import OutputsModels 1.0


Item {
    id: outputsListDelegate

    property bool expanded: false
    property bool animateDisplacement: false

    implicitHeight: itemDelegateWalletName.height + (expanded ? listViewAddresses.height : 0)
    Behavior on implicitHeight { NumberAnimation { duration: animateDisplacement ? 250 : 0; easing.type: Easing.OutQuint; onRunningChanged: animateDisplacement = false } }

    clip: true
    ItemDelegate {
        id: itemDelegateWalletName

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
        id: listViewAddresses
        anchors.top: itemDelegateWalletName.bottom
        anchors.left: parent.left
        anchors.right: parent.right
        height: contentItem.height

        opacity: expanded ? 1.0 : 0.0
        Behavior on opacity { NumberAnimation { duration: expanded ? 250 : 1000; easing.type: Easing.OutQuint } }
        
        clip: true
        interactive: false
        model: qaddresses
        
        delegate: OutputsListAddressDelegate {
            width: listViewAddresses.width
        }
    } // ListView

    // Roles: address
}
