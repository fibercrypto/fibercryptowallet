import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/Controls"
import "Controls" // For quick UI development, switch back to resources when making a release

Item {
    id: createLoadWallet

    enum Mode { Create, Load }

    property Item nextTabItem
    property int mode: CreateLoadWallet.Create
    property alias name: walletName.text
    property alias seed: walletSeed.text
    property alias seedConfirm: walletSeedConfirm.text

    signal dataModified()
    signal walletCreationRequested()

    function clear() {
        walletName.clear()
        walletSeed.clear()
        walletSeedConfirm.clear()
    }

    
    Component.onCompleted:{
        dataModified()
    }
    implicitHeight: walletName.height + walletSeed.height + (mode === CreateLoadWallet.Create ? walletSeedConfirm.height : 0) + 2*column.spacing
    Behavior on implicitHeight { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } }
    clip: true

    Column {
        id: column
        anchors.fill: parent
        spacing: 30

        TextField {
            id: walletName
            width: parent.width

            selectByMouse: true
            placeholderText: qsTr("Wallet's name")
            focus: true

            onTextChanged: {
                dataModified()
            }
        }

        ControlGenerateSeed {
            id: walletSeed
            width: parent.width
            height: inputControlHeight - 10

            placeholderText: qsTr("Wallet's seed")
            buttonLeftText: qsTr("12 words")
            onDataModified: createLoadWallet.dataModified()
            buttonRightText: qsTr("24 words")
            buttonsVisible: mode === CreateLoadWallet.Create
            nextTabItem: walletSeedConfirm
        }

        TextArea {
            id: walletSeedConfirm
            width: parent.width
            height: walletSeed.inputControlHeight - 10

            clip: true
            selectByMouse: true
            wrapMode: TextArea.Wrap
            placeholderText: qsTr("Confirm the wallet's seed")
            opacity: mode === CreateLoadWallet.Create ? 1 : 0
            visible: opacity > 0

            Behavior on opacity { NumberAnimation { duration: 100 } }

            KeyNavigation.priority: KeyNavigation.BeforeItem
            KeyNavigation.tab: nextTabItem

            onTextChanged: {
                dataModified()
            }
        } // TextArea
    } // Column
}
