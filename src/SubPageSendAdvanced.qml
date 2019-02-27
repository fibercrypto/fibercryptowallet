import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

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
                    icon.source: "qrc:/images/help.svg"
                    icon.color: Material.color(Material.Grey)
                }
            }
            ComboBox {
                id: comboBoxWalletsAddressesSendFrom
                Layout.fillWidth: true
                Layout.topMargin: -12
            }
        }

        ColumnLayout {
            id: columnLayoutDestinations

            Layout.alignment: Qt.AlignTop
            RowLayout {
                Label {
                    text: qsTr("Destinations")
                }
                ToolButton {
                    id: toolButtonDestinationPopupHelp
                    icon.source: "qrc:/images/help.svg"
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
                    icon.source: "qrc:/images/help.svg"
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
        }

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
            }
        }
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
        height: sliderCoinHoursShareFactor.height + sliderCoinHoursShareFactor.anchors.topMargin + 10

        HoverHandler {
            id: hoverHandler
            margin: 16
            onHoveredChanged: {
                if (!hovered && !sliderCoinHoursShareFactor.pressed) {
                    toolTipAutomaticCoinHoursAllocation.hide()
                }
            }
        }

        Slider {
            id: sliderCoinHoursShareFactor
            anchors.left: parent.left
            anchors.right: parent.right
            anchors.top: parent.top
            anchors.topMargin: 16

            from: 0.01
            to: 1.00
            value: 0.50


            onPressedChanged: {
                if (!hovered && !toolTipAutomaticCoinHoursAllocation.toolTipHovered) {
                    toolTipAutomaticCoinHoursAllocation.close()
                }
            }

            ToolTip {
                parent: sliderCoinHoursShareFactor.handle
                visible: sliderCoinHoursShareFactor.pressed
                text: sliderCoinHoursShareFactor.value.toFixed(2)
            }
        }
    }
}
