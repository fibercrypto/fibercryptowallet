import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/"
// import "qrc:/ui/src/ui/Controls"
// import "qrc:/ui/src/ui"
import "../" // For quick UI development, switch back to resources when making a release
import "../Controls" // For quick UI development, switch back to resources when making a release

Item {
    id: root

    signal qrCodeRequested(var data)

    onQrCodeRequested: {
        dialogQR.setVars(data)
        dialogQR.open()
    }

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
                id: value
                font.family: "Code New Roman"
                placeholderText: qsTr("Address")
                text: address
                selectByMouse: true
                Layout.fillWidth: true
                onTextChanged: address = text
            }
        } // RowLayout

        RowLayout {
ComboBox{
id:cbCoinTypes
Layout.fillWidth: true
model:coins
onCurrentTextChanged:{
coinType=cbCoinTypes.currentText
}
}
// TextField {
//                id: textFieldDestinationAmount
//                font.family: "Code New Roman"
//                placeholderText: qsTr("Type")
//                text: coinType
//                selectByMouse: true
//                Layout.fillWidth: true
//                onTextChanged: coinType = text
//            }

        }

        ToolButton {
            id: toolButtonAddRemoveDestination
            // The 'accent' attribute is used for button highlighting
            Material.accent: index === 0 ? Material.Teal : Material.Red
            icon.source: "qrc:/images/resources/images/icons/" + (index === 0 ? "add" : "remove") + "-circle.svg"
            highlighted: true // enable the usage of the `Material.accent` attribute

            Layout.alignment: Qt.AlignRight

            onClicked: {
                if (index === 0) {
                    listModelAddresses.append( { "address": "", "coinType": "" } )
                } else {
                    listModelAddresses.remove(index)
                }
            }
        } // ToolButton
    } // RowLayout (rootLayout)

    ListModel{
    id:coins
    ListElement{
    type:"SKY"
    }
    ListElement{
    type:"BTC"
    }
    }
}
