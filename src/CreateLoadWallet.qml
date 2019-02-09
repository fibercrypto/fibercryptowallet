import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12


Item {
    id: root

    implicitWidth: 400
    implicitHeight: 400

    Column {
        anchors.fill: parent
        spacing: 30
        ControlCustomSwitch {
            id: switchNewLoadWallet
            width: 300
            height: 70
            anchors.left: parent.left
            anchors.right: parent.right

            leftText: qsTr("New wallet")
            rightText: qsTr("Load wallet")

            backgroundColor: Material.accent
            leftColor: "white"
            rightColor: "white"

            textColor: Material.accent
        }

        TextField {
            id: walletName
            anchors.left: parent.left
            anchors.right: parent.right

            selectByMouse: true
            placeholderText: qsTr("Wallet's name")
            focus: true
        }

        ControlGenerateSeed {
            id: walletSeed
            anchors.left: parent.left
            anchors.right: parent.right
            width: 100
            height: inputControlHeight - 10

            placeholderText: qsTr("Wallet's seed")
            buttonLeftText: qsTr("12 words")
            buttonRightText: qsTr("24 words")
            buttonsVisible: switchNewLoadWallet.isInLeftSide
        }

        TextArea {
            id: walletSeedConfirm
            anchors.left: parent.left
            anchors.right: parent.right
            height: walletSeed.inputControlHeight

            selectByMouse: true
            wrapMode: TextArea.Wrap
            placeholderText: qsTr("Confirm the wallet's seed")
            opacity: switchNewLoadWallet.isInLeftSide ? 1.0 : 0.0
            visible: opacity > 0.0

            Behavior on opacity { NumberAnimation { duration: 100 } }
        }

        Button {
            id: buttonCreateWallet
            anchors.horizontalCenter: parent.horizontalCenter
            width: 120
            height: 60
            font.bold: true
            font.pointSize: 12

            text: qsTr("Create")
            highlighted: true
        }

        move: Transition {
            NumberAnimation { properties: "x,y"; duration: 250; easing.type: Easing.OutQuint }
            PropertyAction { property: "text"; value: (switchNewLoadWallet.isInLeftSide ? qsTr("Create") : qsTr("Load")) }
        }
    }
}
