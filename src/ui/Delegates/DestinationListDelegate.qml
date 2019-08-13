import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
import "../" // For quick UI development, switch back to resources when making a release

Item {
    id: root

    signal qrCodeRequested(var data)
    Component.onCompleted: {
        root.qrCodeRequested.connect(genQR)
    }

    function genQR(data) {
        dialogQR.text = data
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
                id: textFieldDestinationAddress
                font.family: "Code New Roman"
                placeholderText: qsTr("Destination address")
                text: address
                Layout.fillWidth: true
            }
        }
        RowLayout {
            TextField {
                id: textFieldDestinationAmount
                text: sky
                implicitWidth: 60
                validator: DoubleValidator {
                    notation: DoubleValidator.StandardNotation
                }
            }
            Label {
                text: qsTr("SKY")
            }
        }
        RowLayout {
            visible: !checkBoxAutomaticCoinHoursAllocation.checked
            TextField {
                id: textFieldCoinHoursAmount
                text: coinHours
                implicitWidth: 60
                validator: DoubleValidator {
                    notation: DoubleValidator.StandardNotation
                }
            }
            Label {
                text: qsTr("Coin hours")
            }
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
                    listModelDestinations.append( { "address": "", "sky": 0.0, "coinHours": 0.0 } )
                } else {
                    listModelDestinations.remove(index)
                }
            }
        }
    } // RowLayout (rootLayout)
}
