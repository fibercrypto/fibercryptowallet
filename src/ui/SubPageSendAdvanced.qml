import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
// import "qrc:/ui/src/ui/Dialogs"
import "Delegates/" // For quick UI development, switch back to resources when making a release
import "Dialogs/" // For quick UI development, switch back to resources when making a release
import "Controls" // For quick UI development, switch back to resources when making a release

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

                property var checkedElements: []
                property var checkedElementsText: []
                property int numberOfCheckedElements: checkedElements.length
                property alias filterString: filterPopupWallets.filterText

                Layout.fillWidth: true
                Layout.topMargin: -12
                displayText: numberOfCheckedElements > 1 ? (numberOfCheckedElements + ' ' + qsTr("wallets selected")) : numberOfCheckedElements === 1 ? checkedElementsText[0] : qsTr("No wallet selected")

                model: ["Wallet A", "Wallet B", "Wallet C"]

                popup: FilterComboBoxPopup {
                    id: filterPopupWallets
                    comboBox: comboBoxWalletsSendFrom
                    filterPlaceholderText: qsTr("Filter wallets by name")
                }

                // Taken from Qt 5.13.0 source code:
                delegate: Item {
                    id: rootDelegate

                    property alias checked: checkDelegate.checked
                    property alias text: checkDelegate.text
                    readonly property bool matchFilter: !comboBoxWalletsSendFrom.filterString || text.toLowerCase().includes(comboBoxWalletsSendFrom.filterString.toLowerCase())

                    width: parent.width
                    height: matchFilter ? checkDelegate.height : 0
                    Behavior on height { NumberAnimation { easing.type: Easing.OutQuint } }
                    clip: true

                    CheckDelegate {
                        id: checkDelegate

                        // Update the states saved in `checkedElements`
                        onClicked: {
                            if (checked) {
                                var pos = comboBoxWalletsSendFrom.checkedElements.indexOf(index)
                                if (pos < 0) {
                                    comboBoxWalletsSendFrom.checkedElements.push(index)
                                    comboBoxWalletsSendFrom.checkedElementsText.push(text)
                                }
                            } else {
                                var pos = comboBoxWalletsSendFrom.checkedElements.indexOf(index)
                                if (pos >= 0) {
                                    comboBoxWalletsSendFrom.checkedElements.splice(pos, 1)
                                    comboBoxWalletsSendFrom.checkedElementsText.splice(pos, 1)
                                }
                            }
                            comboBoxWalletsSendFrom.numberOfCheckedElements = comboBoxWalletsSendFrom.checkedElements.length
                        }

                        width: parent.width
                        text: comboBoxWalletsSendFrom.textRole ? (Array.isArray(comboBoxWalletsSendFrom.model) ? modelData[comboBoxWalletsSendFrom.textRole] : model[comboBoxWalletsSendFrom.textRole]) : modelData
                        // Load the saved state when the delegate is recicled:
                        checked: comboBoxWalletsSendFrom.checkedElements.indexOf(index) > 0
                        hoverEnabled: comboBoxWalletsSendFrom.hoverEnabled
                        highlighted: hovered
                        Material.foreground: checked ? parent.Material.accent : parent.Material.foreground
                        leftPadding: highlighted ? 2*padding : padding // added
                        Behavior on leftPadding { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } } // added

                        LayoutMirroring.enabled: true
                        contentItem: Label {
                            leftPadding: comboBoxWalletsSendFrom.indicator.width + comboBoxWalletsSendFrom.spacing
                            text: checkDelegate.text
                            verticalAlignment: Qt.AlignVCenter
                            color: checkDelegate.enabled ? checkDelegate.Material.foreground : checkDelegate.Material.hintTextColor
                        }
                    } // CheckDelegate
                } // Item (delegate)
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

                Layout.fillWidth: true
                Layout.topMargin: -12

                model: ["sgdkaugakugxfnakusdhgf",
                        "uhrencgkhmjhsmfugwnjwh",
                        "iwyerniywetrdntwyierue",
                        "pney73snyiquemqskddqgq",
                        "inweytr82n3sr28myrxm28"]

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
                    text: qsTr("All outputs of the selected addresses")
                    checked: true
                }
            }

            ComboBox {
                id: comboBoxWalletsUnspentOutputsSendFrom

                property var checkedElements: []
                property var checkedElementsText: []
                property int numberOfCheckedElements: checkedElements.length
                property alias filterString: filterPopupOutputs.filterText
                
                Layout.fillWidth: true
                Layout.topMargin: -12
                displayText: checkBoxUnspentOutputsUseAllOutputs.checked ? qsTr("All outputs selected") : numberOfCheckedElements > 1 ? (numberOfCheckedElements + ' ' + qsTr("outputs selected")) : numberOfCheckedElements === 1 ? checkedElementsText[0] : qsTr("No output selected")

                enabled: !checkBoxUnspentOutputsUseAllOutputs.checked
                model: !enabled ? null :
                        ["sgdkaugakugxfnakusdhgf",
                        "uhrencgkhmjhsmfugwnjwh",
                        "iwyerniywetrdntwyierue",
                        "pney73snyiquemqskddqgq",
                        "inweytr82n3sr28myrxm28"]

                onModelChanged: {
                    if (!model) {
                        checkedElements = []
                        checkedElementsText = []
                        numberOfCheckedElements = 0
                    }
                }
                
                popup: FilterComboBoxPopup {
                    id: filterPopupOutputs
                    comboBox: comboBoxWalletsUnspentOutputsSendFrom
                    filterPlaceholderText: qsTr("Filter outputs")
                }

                delegate: Item {

                    property alias checked: checkDelegate.checked
                    property alias text: checkDelegate.text
                    readonly property bool matchFilter: !comboBoxWalletsUnspentOutputsSendFrom.filterString || text.toLowerCase().includes(comboBoxWalletsUnspentOutputsSendFrom.filterString.toLowerCase())
                    
                    width: parent.width
                    height: matchFilter ? checkDelegate.height : 0
                    Behavior on height { NumberAnimation { easing.type: Easing.OutQuint } }
                    clip: true

                    CheckDelegate {
                        id: checkDelegate

                        // Update the states saved in `checkedElements`
                        onClicked: {
                            if (checked) {
                                var pos = comboBoxWalletsUnspentOutputsSendFrom.checkedElements.indexOf(index)
                                if (pos < 0) {
                                    comboBoxWalletsUnspentOutputsSendFrom.checkedElements.push(index)
                                    comboBoxWalletsUnspentOutputsSendFrom.checkedElementsText.push(text)
                                }
                            } else {
                                var pos = comboBoxWalletsUnspentOutputsSendFrom.checkedElements.indexOf(index)
                                if (pos >= 0) {
                                    comboBoxWalletsUnspentOutputsSendFrom.checkedElements.splice(pos, 1)
                                    comboBoxWalletsUnspentOutputsSendFrom.checkedElementsText.splice(pos, 1)
                                }
                            }
                            comboBoxWalletsUnspentOutputsSendFrom.numberOfCheckedElements = comboBoxWalletsUnspentOutputsSendFrom.checkedElements.length
                        }

                        width: parent.width
                        text: comboBoxWalletsUnspentOutputsSendFrom.textRole ? (Array.isArray(comboBoxWalletsUnspentOutputsSendFrom.model) ? modelData[comboBoxWalletsUnspentOutputsSendFrom.textRole] : model[comboBoxWalletsUnspentOutputsSendFrom.textRole]) : modelData
                        font.family: "Code New Roman"
                        // Load the saved state when the delegate is recicled:
                        checked: comboBoxWalletsUnspentOutputsSendFrom.checkedElements.indexOf(index) > 0
                        hoverEnabled: comboBoxWalletsSendFrom.hoverEnabled
                        highlighted: hovered
                        Material.foreground: checked ? parent.Material.accent : parent.Material.foreground
                        leftPadding: highlighted ? 2*padding : padding // added
                        Behavior on leftPadding { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } } // added

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

                    onClicked: {
                        dialogSelectAddressByWallet.open()
                    }
                }
            }

            TextField {
                id: textFieldCustomChangeAddress
                
                Layout.fillWidth: true
                Layout.topMargin: -16
                placeholderText: qsTr("Address to receive change")
                selectByMouse: true
                font.family: "Code New Roman"
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

    DialogSelectAddressByWallet {
        id: dialogSelectAddressByWallet

        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 540 ? 540 - 40 : applicationWindow.width - 40
        height: applicationWindow.height - 40

        model: modelAddressesByWallet

        focus: true
        modal: true

        onAccepted: {
            textFieldCustomChangeAddress.text = selectedAddress
        }
    }

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

    ListModel {
        id: listModelDestinations
        ListElement { address: ""; sky: 0.0; coinHours: 0.0 }
    }
}
