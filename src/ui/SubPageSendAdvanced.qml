import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import WalletsManager 1.0
import OutputsModels 1.0

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "Delegates/" // For quick UI development, switch back to resources when making a release

Page {
    id: subPageSendAdvanced

    ColumnLayout {
        id: columnLayoutRoot
        anchors.fill: parent
        anchors.leftMargin: 10
        anchors.rightMargin: 10
        spacing: 20

        ColumnLayout {
            id: columnLayoutSendFrom

            Layout.alignment: Qt.AlignTop

            RowLayout {

                Label { text: qsTr("Wallet") }

                ToolButton {
                    id: toolButtonWalletPopupHelp
                    icon.source: "qrc:/images/resources/images/icons/help.svg"
                    icon.color: Material.color(Material.Grey)
                }
            }

            ComboBox {
                id: comboBoxWalletsSendFrom
                Layout.fillWidth: true
                Layout.topMargin: -12
                textRole: "name"
                model: WalletModel {
                    Component.onCompleted: {
                        loadModel(walletManager.getWallets())
                    }
                } 
                onCurrentTextChanged:{
                    console.log(model.wallets[currentIndex].fileName)
                    listAddresses.loadModel(walletManager.getAddresses(model.wallets[currentIndex].fileName))
                    listAddresses.removeAddress(0)

                    
                }

                // Taken from Qt 5.13.0 source code:
                delegate: MenuItem {
                    width: parent.width
                    text: comboBoxWalletsSendFrom.textRole ? (Array.isArray(comboBoxWalletsSendFrom.model) ? modelData[comboBoxWalletsSendFrom.textRole] : model[comboBoxWalletsSendFrom.textRole]) : modelData
                    Material.foreground: comboBoxWalletsSendFrom.currentIndex === index ? parent.Material.accent : parent.Material.foreground
                    highlighted: comboBoxWalletsSendFrom.highlightedIndex === index
                    hoverEnabled: comboBoxWalletsSendFrom.hoverEnabled
                    leftPadding: highlighted ? 2*padding : padding // added
                    Behavior on leftPadding { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } } // added
                }
            } // ComboBox (wallets, send from)

            RowLayout {

                Label { text: qsTr("Address") }

                ToolButton {
                    id: toolButtonAddressPopupHelp
                    icon.source: "qrc:/images/resources/images/icons/help.svg"
                    icon.color: Material.color(Material.Grey)
                }
            }

            ComboBox {
                id: comboBoxWalletsAddressesSendFrom
                //Layout.fillWidth: true
                //Layout.topMargin: -12
                //textRole: "address"
                //model: AddressModel{
                //    id: listAddresses
                //}
//
                //// Taken from Qt 5.13.0 source code:
                //delegate: MenuItem {
                //    width: parent.width
                //    text: comboBoxWalletsAddressesSendFrom.textRole ? (Array.isArray(comboBoxWalletsAddressesSendFrom.model) ? modelData[comboBoxWalletsAddressesSendFrom.textRole] : model[comboBoxWalletsAddressesSendFrom.textRole]) : modelData
                //    Material.foreground: comboBoxWalletsAddressesSendFrom.currentIndex === index ? parent.Material.accent : parent.Material.foreground
                //    highlighted: comboBoxWalletsAddressesSendFrom.highlightedIndex === index
                //    hoverEnabled: comboBoxWalletsAddressesSendFrom.hoverEnabled
                //    leftPadding: highlighted ? 2*padding : padding // added
                //    Behavior on leftPadding { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } } // added
                //}
                // This function returns all checked index in the ComboBox's popup
                function getCheckedDelegates() {
                    var checkedItems = []
                    for (var i = 0; i < popup.contentItem.contentItem.children.length; i++) {
                        if (popup.contentItem.contentItem.children[i].checked) {
                            checkedItems.push(i)
                        }
                    }
                    return checkedItems
                }
                
                Layout.fillWidth: true
                Layout.topMargin: -12

                
                model: AddressModel{
                    id: listAddresses
                }
                textRole: "address"

                delegate: Item {
                    width: parent.width
                    height: checkDelegate.height

                    property alias checked: checkDelegate.checked

                    CheckDelegate {
                        id: checkDelegate

                        width: parent.width
                        text: comboBoxWalletsAddressesSendFrom.textRole ? (Array.isArray(comboBoxWalletsAddressesSendFrom.model) ? modelData[comboBoxWalletsAddressesSendFrom.textRole] : model[comboBoxWalletsAddressesSendFrom.textRole]) : modelData
                        font.family: "Code New Roman"

                        LayoutMirroring.enabled: true
                        contentItem: Label {
                            leftPadding: comboBoxWalletsAddressesSendFrom.indicator.width + comboBoxWalletsAddressesSendFrom.spacing
                            text: checkDelegate.text
                            verticalAlignment: Qt.AlignVCenter
                            color: checkDelegate.enabled ? checkDelegate.Material.foreground : checkDelegate.Material.hintTextColor
                        }

                        onCheckedChanged:{
                            if (checked){
                                //console.log(comboBoxWalletsSendFrom.model.wallets[comboBoxWalletsSendFrom.currentIndex])
                                listOutputs.insertOutputs(walletManager.getOutputs(comboBoxWalletsSendFrom.model.wallets[comboBoxWalletsSendFrom.currentIndex].fileName, text))
                                //console.log(text)
                            }
                            //console.log("SDFDSFS")
                        }
                    } // CheckDelegate
                } // Item (delegate)
            } // ComboBox (addresses, send from)

            RowLayout {

                Label { text: qsTr("Unspent outputs") }

                ToolButton {
                    id: toolButtonUnspentOutputsPopupHelp
                    icon.source: "qrc:/images/resources/images/icons/help.svg"
                    icon.color: Material.color(Material.Grey)
                }

                CheckBox {
                    id: checkBoxUnspentOutputsUseAllOutputs
                    text: qsTr("All outputs of the selected address")
                    checked: true
                }
            }

            ComboBox {
                id: comboBoxWalletsUnspentOutputsSendFrom

                // This function returns all checked index in the ComboBox's popup
                function getCheckedDelegates() {
                    var checkedItems = []
                    for (var i = 0; i < popup.contentItem.contentItem.children.length; i++) {
                        if (popup.contentItem.contentItem.children[i].checked) {
                            checkedItems.push(i)
                        }
                    }
                    return checkedItems
                }
                
                Layout.fillWidth: true
                Layout.topMargin: -12
                textRole: "outputID"
                enabled: !checkBoxUnspentOutputsUseAllOutputs.checked
                model: QOutputs {
                    id: listOutputs
                }

                delegate: Item {
                    width: parent.width
                    height: checkDelegate.height

                    property alias checked: checkDelegate.checked

                    CheckDelegate {
                        id: checkDelegate

                        width: parent.width
                        text: comboBoxWalletsUnspentOutputsSendFrom.textRole ? (Array.isArray(comboBoxWalletsUnspentOutputsSendFrom.model) ? modelData[comboBoxWalletsUnspentOutputsSendFrom.textRole] : model[comboBoxWalletsUnspentOutputsSendFrom.textRole]) : modelData
                        font.family: "Code New Roman"

                        LayoutMirroring.enabled: true
                        contentItem: Label {
                            leftPadding: comboBoxWalletsUnspentOutputsSendFrom.indicator.width + comboBoxWalletsUnspentOutputsSendFrom.spacing
                            text: checkDelegate.text
                            verticalAlignment: Qt.AlignVCenter
                            color: checkDelegate.enabled ? checkDelegate.Material.foreground : checkDelegate.Material.hintTextColor
                        }
                    } // CheckDelegate
                } // Item (delegate)
            } // ComboBox (outputs, send from)

        } // ColumnLayout (send from)

        Label {
            Layout.fillWidth: true
            Layout.preferredHeight: 30
            text: qsTr("With your current selection you can send up to <b>%1 SKY</b> and <b>%2 Coin Hours</b> (at least <b>%3 Coin Hours</b> must be used for the transaction fee)").arg(0).arg(0).arg(0)
            wrapMode: Text.WordWrap
            horizontalAlignment: Text.AlignHCenter
        }

        ColumnLayout {
            id: columnLayoutDestinations

            Layout.alignment: Qt.AlignTop

            RowLayout {

                Label { text: qsTr("Destinations") }

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
                clip: true

                model: listModelDestinations

                delegate: DestinationListDelegate {
                    width: listViewDestinations.width
                    implicitHeight: ListView.view.delegateHeight
                }
            } // ListView
        } // ColumnLayout (destinations)

        ColumnLayout {
            id: columnLayoutCustomChangeAddress

            Layout.alignment: Qt.AlignTop

            RowLayout {

                Label { text: qsTr("Custom change address") }

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

            CheckBox {
                id: checkBoxAutomaticCoinHoursAllocation
                text: qsTr("Automatic coin hours allocation")
                checked: true
            }

            Slider {
                id: sliderCoinHoursShareFactor

                Layout.fillWidth: true

                opacity: checkBoxAutomaticCoinHoursAllocation.checked ? 1.0 : 0.0
                Behavior on opacity { NumberAnimation {} }
                from: 0.01
                to: 1.00
                value: 0.50
                stepSize: 0.01

                ToolTip {
                    parent: sliderCoinHoursShareFactor.handle
                    visible: sliderCoinHoursShareFactor.pressed
                    text: sliderCoinHoursShareFactor.value.toFixed(2)
                    font.pointSize: Qt.application.font.pointSize * 1.2
                }
            }
        } // ColumnLayout (automatic coin hours allocation)
    } // ColumnLayout (root)

    ListModel {
        id: listModelDestinations
        ListElement { address: ""; sky: 0.0; coinHours: 0.0 }
    }
}
