import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "Delegates/" // For quick UI development, switch back to resources when making a release

Page {
    id: root

    ColumnLayout {
        id: columnLayoutRoot
        anchors.fill: parent
        anchors.leftMargin: 10
        anchors.rightMargin: 10
        spacing: 20

        ColumnLayout {
            id: columnLayoutAddressesSendFrom

            Layout.alignment: Qt.AlignTop
            RowLayout {
                Label {
                    text: qsTr("Address")
                }
                ToolButton {
                    id: toolButtonAddressPopupHelp
                    icon.source: "qrc:/images/resources/images/icons/help.svg"
                    icon.color: Material.color(Material.Grey)
                }
            }
            ComboBox {
                id: comboBoxWalletsAddressesSendFrom
                Layout.fillWidth: true
                Layout.topMargin: -12

                textRole: "address"
                model: modelAddressesByWallet

                // Taken from Qt 5.13.0 source code:
                delegate: MenuItem {
                    width: parent.width
                    text: comboBoxWalletsAddressesSendFrom.textRole ? (Array.isArray(comboBoxWalletsAddressesSendFrom.model) ? modelData[comboBoxWalletsAddressesSendFrom.textRole] : model[comboBoxWalletsAddressesSendFrom.textRole]) : modelData
                    Material.foreground: comboBoxWalletsAddressesSendFrom.currentIndex === index ? parent.Material.accent : parent.Material.foreground
                    highlighted: comboBoxWalletsAddressesSendFrom.highlightedIndex === index
                    hoverEnabled: comboBoxWalletsAddressesSendFrom.hoverEnabled
                    leftPadding: highlighted ? 2*padding : padding // added
                    Behavior on leftPadding { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } } // added
                }

                popup: CustomComboBoxDropDown { comboBox: comboBoxWalletsAddressesSendFrom }
            } // ComboBox
        } // ColumnLayout (send from)

        ColumnLayout {
            id: columnLayoutDestinations

            Layout.alignment: Qt.AlignTop
            RowLayout {
                Label {
                    text: qsTr("Destinations")
                }
                ToolButton {
                    id: toolButtonDestinationPopupHelp
                    icon.source: "qrc:/images/resources/images/icons/help.svg"
                    icon.color: Material.color(Material.Grey)
                }
            }
            ListView {
                id: listViewDestinations

                property real delegateHeight: 47

                Layout.fillWidth: true
                Layout.topMargin: -16
                implicitHeight: count * delegateHeight

                Behavior on implicitHeight { NumberAnimation { duration: 250; easing.type: Easing.OutQuint } }

                interactive: false

                model: listModelDestinations
                clip: true

                delegate: DestinationListDelegate {
                    id: delegateDestination
                    width: listViewDestinations.width
                    implicitHeight: ListView.view.delegateHeight
                }
            } // ListView
        } // ColumnLayout (destinations)

        ColumnLayout {
            id: columnLayoutCustomChangeAddress

            Layout.alignment: Qt.AlignTop
            RowLayout {
                Label {
                    text: qsTr("Custom change address")
                }
                ToolButton {
                    id: toolButtonCustomChangeAddressPopupHelp
                    icon.source: "qrc:/images/resources/images/icons/help.svg"
                    icon.color: Material.color(Material.Grey)
                }
                Button {
                    id: buttonSelectCustomChangeAddress
                    text: qsTr("Select")
                    flat: true
                    highlighted: true
                }
            }
            TextField {
                id: textFieldCustomChangeAddress
                placeholderText: qsTr("Address to receive change")
                Layout.fillWidth: true
                Layout.topMargin: -16
            }
        } // ColumnLayout (custom change address)

        ColumnLayout {
            id: columnLayoutAutomaticCoinHoursAllocation

            Layout.alignment: Qt.AlignTop
            RowLayout {
                CheckBox {
                    id: checkBoxAutomaticCoinHoursAllocation
                    text: qsTr("Automatic coin hours allocation")
                    checked: true
                }
                Button {
                    id: buttonOptionsAutomaticCoinHoursAllocation
                    text: qsTr("Options")
                    flat: true
                    highlighted: true
                    enabled: checkBoxAutomaticCoinHoursAllocation.checked

                    onClicked: {
                        toolTipAutomaticCoinHoursAllocation.show("Coin hours share factor")
                    }
                }
            } // RowLayout
        } // ColumnLayout (automatic coin hours allocation)
    } // ColumnLayout (root)

    ListModel {
        id: listModelDestinations
        ListElement { address: ""; sky: 0.0; coinHours: 0.0 }
    }


    ToolTip {
        id: toolTipAutomaticCoinHoursAllocation

        property alias toolTipHovered: hoverHandler.hovered

        parent: buttonOptionsAutomaticCoinHoursAllocation
        anchors.centerIn: parent
        height: sliderCoinHoursShareFactor.height + rowLayoutSlider.anchors.topMargin + 10

        HoverHandler {
            id: hoverHandler
            margin: 16
            onHoveredChanged: {
                if (!hovered && !sliderCoinHoursShareFactor.pressed) {
                    toolTipAutomaticCoinHoursAllocation.hide()
                }
            }
        }

        RowLayout {
            id: rowLayoutSlider
            anchors.left: parent.left
            anchors.right: parent.right
            anchors.top: parent.top
            anchors.topMargin: 16
            Slider {
                id: sliderCoinHoursShareFactor
                Layout.fillWidth: true

                from: 0.01
                to: 1.00
                value: 0.50
                stepSize: 0.01

                onPressedChanged: {
                    if (!hovered && !toolTipAutomaticCoinHoursAllocation.toolTipHovered) {
                        toolTipAutomaticCoinHoursAllocation.close()
                    }
                }
            } // Slider

            Label {
                text: sliderCoinHoursShareFactor.value.toFixed(2)
            }
        } // RowLayout (slider)
    } // ToolTip (automatic coin hours allocation)

    // Roles: wallet, address
    // Implement the model in the backend (a more recommendable approach)
    ListModel { // EXAMPLE
        id: modelAddressesByWallet
        ListElement { wallet: "Wallet A"; address: "qrxw7364w8xerusftaxkw87ues" }
        ListElement { wallet: "Wallet A"; address: "8745yuetsrk8tcsku4ryj48ije" }
        ListElement { wallet: "Wallet A"; address: "gfdhgs343kweru38200384uwqd" }
        ListElement { wallet: "Wallet B"; address: "00qdqsdjkssvmchskjkxxdg374" }
        ListElement { wallet: "Wallet B"; address: "hkdti34aoliwuiu3qsoiemdfhc" }
        ListElement { wallet: "Wallet C"; address: "1oiwrelkrir73o8ielukaur9qq" }
        ListElement { wallet: "Wallet C"; address: "piur948o9q8m0a8qsye8q3omxs" }
        ListElement { wallet: "Wallet C"; address: "4ntd4im93usppturm83ysniroe" }
        ListElement { wallet: "Wallet C"; address: "meje73o50ejdwumfle92rndlwm" }
    }        
}
