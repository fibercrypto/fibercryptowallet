import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import WalletsManager 1.0
import OutputsModels 1.0

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
// import "qrc:/ui/src/ui/Dialogs"
// import "qrc:/ui/src/ui/Controls"
import "Delegates" // For quick UI development, switch back to resources when making a release
import "Dialogs" // For quick UI development, switch back to resources when making a release
import "Controls" // For quick UI development, switch back to resources when making a release

Page {
    id: subPageSendAdvanced

    property int upperCoinBound: 0
    property int upperAltCointBound: 0
    property int minFeeAmount: 0

    function updateInfo() {
		subPageSendAdvanced.updateOutputs()
	   	upperCoinBound = 0;
	   	upperAltCointBound = 0;
	   	minFeeAmount = 0;
		var valCH = 0;
		if (comboBoxWalletsUnspentOutputsSendFrom.enabled) {
       		for (var i = 0; i < comboBoxWalletsUnspentOutputsSendFrom.checkedElements.length; i++){
       			upperCoinBound += parseFloat(comboBoxWalletsUnspentOutputsSendFrom.model.outputs[comboBoxWalletsUnspentOutputsSendFrom.checkedElements[i]].addressSky, 10);
				var s = comboBoxWalletsUnspentOutputsSendFrom.model.outputs[comboBoxWalletsUnspentOutputsSendFrom.checkedElements[i]].addressCoinHours;
				console.log(parseInt(s.replace('\,',''), 10))
       		    valCH += parseInt(comboBoxWalletsUnspentOutputsSendFrom.model.outputs[comboBoxWalletsUnspentOutputsSendFrom.checkedElements[i]].addressCoinHours.replace('\,',''), 10);
       		}
		} else {

			if (comboBoxWalletsAddressesSendFrom.enabled) {
    	   		for (var i = 0; i < comboBoxWalletsAddressesSendFrom.checkedElements.length; i++){
    	   			upperCoinBound += parseFloat(comboBoxWalletsAddressesSendFrom.model.addresses[comboBoxWalletsAddressesSendFrom.checkedElements[i]].addressSky, 10);
    	   		    valCH += parseInt(comboBoxWalletsAddressesSendFrom.model.addresses[comboBoxWalletsAddressesSendFrom.checkedElements[i]].addressCoinHours.replace('\,',''), 10);
    	   		}
			} else {
				for(var i = 0; i < comboBoxWalletsAddressesSendFrom.model.addresses.length; i++) {
					upperCoinBound += parseFloat(comboBoxWalletsAddressesSendFrom.model.addresses[i].addressSky, 10)
					valCH += parseInt(comboBoxWalletsAddressesSendFrom.model.addresses[i].addressCoinHours.replace('\,',''), 10);
				} 
			}
		}
		upperAltCointBound = valCH*9/10
		minFeeAmount = valCH/10
    }
    function geteAddressesWithWallets(){
        var indexs =  comboBoxWalletsAddressesSendFrom.getCheckedDelegates()
        var addresses = []
        addresses.push([])
        addresses.push([])
        for (var i =0;i< indexs.length; i++){
            addresses[0].push(comboBoxWalletsAddressesSendFrom.model.addresses[indexs[i]].address)
            addresses[1].push(comboBoxWalletsAddressesSendFrom.model.addresses[indexs[i]].walletId)
            //addresses.push(comboBoxWalletsAddressesSendFrom.model.addresses[indexs[i]].address)
        }
        return addresses
	}

    function getSelectedOutputsWithWallets(){
        var indexs =  comboBoxWalletsUnspentOutputsSendFrom.getCheckedDelegates()
        var outputs = []
        outputs.push([])
        outputs.push([])
        for (var i =0;i< indexs.length; i++){
            outputs[0].push(comboBoxWalletsUnspentOutputsSendFrom.model.outputs[indexs[i]].outputID)
            outputs[1].push(comboBoxWalletsUnspentOutputsSendFrom.model.outputs[indexs[i]].walletOwner)
            //outputs.push(comboBoxWalletsUnspentOutputsSendFrom.model.outputs[indexs[i]].outputID)
        }
        return outputs
    }

    function getSelectedWallet(){

        var indexs = comboBoxWalletsSendFrom.getCheckedDelegates()
        var files = []
        for (var i=0; i < indexs.length; i++){
            files.push(comboBoxWalletsSendFrom.model.wallets[indexs[i]].fileName)
        }
        return files
    }

    function walletIsEncrypted(){
        var indexs = comboBoxWalletsSendFrom.getCheckedDelegates()
        var enc = []
        for (var i = 0; i < indexs.length; i++){
            enc.push(comboBoxWalletsSendFrom.model.wallets[indexs[i]].encryptionEnabled)
        }
        return enc
    }

    function getDestinationsSummary(){
        var addrs = []
        var skyAmounts = []
        var coinHoursAmount = []
        for (var i = 0; i < listModelDestinations.count; i++){
            addrs.push(listModelDestinations.get(i).address)
            skyAmounts.push(listModelDestinations.get(i).sky)
            coinHoursAmount.push(listModelDestinations.get(i).coinHours)
        }
        return [addrs, skyAmounts, coinHoursAmount]
    }

    function getChangeAddress(){
        return textFieldCustomChangeAddress.text
    }

    function getAutomaticCoinHours(){
        return checkBoxAutomaticCoinHoursAllocation.checked
    }
    function getBurnFactor(){
        return sliderCoinHoursShareFactor.value
    }

    function getAllAddressesWithWallets(){
        var addrs = []
        addrs.push([])
        addrs.push([])
        for (var i = 0; i < listAddresses.count; i++){
            addrs[0].push(listAddresses.addresses[i].address)
            addrs[1].push(listAddresses.addresses[i].walletId)
            //addrs.push(listAddresses.addresses[i].address)
        }
        return addrs
    }
	function updateOutputs() {
		listOutputs.cleanModel()
		if (checkBoxAllAddresses.checked) {
       		for (var i = 0; i < comboBoxWalletsSendFrom.checkedElements.length; i++){
				walletManager.updateAddresses(comboBoxWalletsSendFrom.model.wallets[comboBoxWalletsSendFrom.checkedElements[i]].fileName)
				var addresses = walletManager.getAddresses(comboBoxWalletsSendFrom.model.wallets[comboBoxWalletsSendFrom.checkedElements[i]].fileName)
				for(var j = 0; j < addresses.length; j++) {
					console.log(comboBoxWalletsSendFrom.model.wallets[comboBoxWalletsSendFrom.checkedElements[i]].fileName)
					walletManager.updateOutputs(comboBoxWalletsSendFrom.model.wallets[comboBoxWalletsSendFrom.checkedElements[i]].fileName, addresses[j].address)
					listOutputs.insertOutputs(walletManager.getOutputs(comboBoxWalletsSendFrom.model.wallets[comboBoxWalletsSendFrom.checkedElements[i]].fileName, addresses[j].address))
				}
			}
		} else {
       		for (var j = 0; j < comboBoxWalletsSendFrom.checkedElements.length; j++){
				walletManager.updateAddresses(comboBoxWalletsSendFrom.model.wallets[comboBoxWalletsSendFrom.checkedElements[j]].fileName)
			}
       		for (var i = 0; i < comboBoxWalletsAddressesSendFrom.checkedElements.length; i++){
				walletManager.updateOutputs(comboBoxWalletsAddressesSendFrom.model.addresses[comboBoxWalletsAddressesSendFrom.checkedElements[i]].walletId, comboBoxWalletsAddressesSendFrom.model.addresses[comboBoxWalletsAddressesSendFrom.checkedElements[i]].address)
				listOutputs.insertOutputs(walletManager.getOutputs(comboBoxWalletsAddressesSendFrom.model.addresses[comboBoxWalletsAddressesSendFrom.checkedElements[i]].walletId, comboBoxWalletsAddressesSendFrom.model.addresses[comboBoxWalletsAddressesSendFrom.checkedElements[i]].address))
			}
		}
	}

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
                function getCheckedDelegates() {
                    return checkedElements
                }

                property var checkedElements: []
                property var checkedElementsText: []
                property int numberOfCheckedElements: checkedElements.length
                property alias filterString: filterPopupWallets.filterText

                Layout.fillWidth: true
                Layout.topMargin: -12
                textRole: "name"
                displayText: numberOfCheckedElements > 1 ? (numberOfCheckedElements + ' ' + qsTr("address selected")) : numberOfCheckedElements === 1 ? checkedElementsText[0] : qsTr("No address selected")
                model: WalletModel {
                    Component.onCompleted: {
                        loadModel(walletManager.getWallets())
                    }
                }

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
                                listAddresses.addAddresses(walletManager.getAddresses(comboBoxWalletsSendFrom.model.wallets[index].fileName))
                                listOutputs.insertOutputs(walletManager.getOutputsFromWallet(comboBoxWalletsSendFrom.model.wallets[index].fileName))
                            } else {
                                var pos = comboBoxWalletsSendFrom.checkedElements.indexOf(index)
                                if (pos >= 0) {
                                    comboBoxWalletsSendFrom.checkedElements.splice(pos, 1)
                                    comboBoxWalletsSendFrom.checkedElementsText.splice(pos, 1)
                                }
                                // Update Outputs and Addresses Model
                                listAddresses.removeAddressesFromWallet(comboBoxWalletsSendFrom.model.wallets[index].fileName)
                                listOutputs.removeOutputsFromWallet(comboBoxWalletsSendFrom.model.wallets[index].fileName)
                            }
							subPageSendAdvanced.updateInfo();
                            comboBoxWalletsSendFrom.numberOfCheckedElements = comboBoxWalletsSendFrom.checkedElements.length
                        }

                        width: parent.width
                        text: comboBoxWalletsSendFrom.textRole ? (Array.isArray(comboBoxWalletsSendFrom.model) ? modelData[comboBoxWalletsSendFrom.textRole] + " - " + modelData["sky"] + " SKY (" + modelData["coinHours"] + " CoinHours)"  : model[comboBoxWalletsSendFrom.textRole] + " - " + model["sky"] + " SKY (" + model["coinHours"] + " CoinHours)") : " --- " + modelData
                        // Load the saved state when the delegate is recicled:
                        checked: comboBoxWalletsSendFrom.checkedElements.indexOf(index) >= 0
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
                CheckBox {
                    id: checkBoxAllAddresses
                    text: qsTr("All Addresses of the selected addresses")
                    checked: true
					onClicked: {
						subPageSendAdvanced.updateInfo()
					}
                }
            }

            ComboBox {
                id: comboBoxWalletsAddressesSendFrom
                function getCheckedDelegates() {
                    return checkedElements
                }
                property var checkedElements: []
                property var checkedElementsText: []
                property int numberOfCheckedElements: checkedElements.length
                property alias filterString: filterPopupAddresses.filterText

                popup: FilterComboBoxPopup {
                    id: filterPopupAddresses
                    comboBox: comboBoxWalletsAddressesSendFrom
                    filterPlaceholderText: qsTr("Filter Addresses")
                }

                Layout.fillWidth: true
                Layout.topMargin: -12

                onModelChanged: {
                    if (!model) {
                        checkedElements = []
                        checkedElementsText = []
                        numberOfCheckedElements = 0
                    }
                }

                model: AddressModel{
                    id: listAddresses
                }

                textRole: "address"
                displayText: !checkBoxAllAddresses.checked ? (numberOfCheckedElements > 1 ? (numberOfCheckedElements + ' ' + qsTr("addresses selected")) : numberOfCheckedElements === 1 ? checkedElementsText[0] : qsTr("No addresses selected")): "All address selected" 
                enabled: !checkBoxAllAddresses.checked
                delegate: Item {

                    property alias checked: checkDelegate.checked
                    property alias text: checkDelegate.text
                    readonly property bool matchFilter: !comboBoxWalletsAddressesSendFrom.filterString || text.toLowerCase().includes(comboBoxWalletsAddressesSendFrom.filterString.toLowerCase())

                    width: parent.width
                    height: matchFilter ? checkDelegate.height : 0
                    Behavior on height { NumberAnimation { easing.type: Easing.OutQuint } }
                    clip: true


                    CheckDelegate {
                        id: checkDelegate

                        width: parent.width
                        text: comboBoxWalletsAddressesSendFrom.textRole ? (Array.isArray(comboBoxWalletsAddressesSendFrom.model) ? modelData["addressSky"] + " --- " +  modelData[comboBoxWalletsAddressesSendFrom.textRole]  + " - " + modelData["addressSky"] + " SKY (" + modelData["addressCoinHours"] + " CoinHours)" : model[comboBoxWalletsAddressesSendFrom.textRole] + " - " + model["addressSky"] + " SKY (" + model["addressCoinHours"] + " CoinHours)") : modelData
                        font.family: "Code New Roman"

                        LayoutMirroring.enabled: true
                        contentItem: Label {
                            leftPadding: comboBoxWalletsAddressesSendFrom.indicator.width + comboBoxWalletsAddressesSendFrom.spacing
                            text: checkDelegate.text
                            verticalAlignment: Qt.AlignVCenter
                            color: checkDelegate.enabled ? checkDelegate.Material.foreground : checkDelegate.Material.hintTextColor
                        }
                        onClicked: {
                            if (checked) {
                                var addrText = comboBoxWalletsAddressesSendFrom.textRole ? (Array.isArray(comboBoxWalletsAddressesSendFrom.model) ? modelData[comboBoxWalletsAddressesSendFrom.textRole] : model[comboBoxWalletsAddressesSendFrom.textRole]) : modelData;
                                if (comboBoxWalletsAddressesSendFrom.getCheckedDelegates().length > 1){
                                    listOutputs.insertOutputs(walletManager.getOutputs(comboBoxWalletsAddressesSendFrom.model.addresses[index].walletId, addrText))
                                } else{
                                    listOutputs.loadModel(walletManager.getOutputs(comboBoxWalletsAddressesSendFrom.model.addresses[index].walletId, addrText))
                                }
                                var pos = comboBoxWalletsAddressesSendFrom.checkedElements.indexOf(index)
                                if (pos < 0) {
                                    comboBoxWalletsAddressesSendFrom.checkedElements.push(index)
                                    comboBoxWalletsAddressesSendFrom.checkedElementsText.push(text)
                                }
                            } else {
                                listOutputs.removeOutputsFromAddress(text)
                                if (comboBoxWalletsAddressesSendFrom.getCheckedDelegates().length == 0){
                                    var indexs = comboBoxWalletsSendFrom.getCheckedDelegates()
                                    for (var i = 0; i < indexs.length; i++){
                                        listOutputs.insertOutputs(walletManager.getOutputsFromWallet(comboBoxWalletsSendFrom.model.wallets[indexs[i]].fileName))
                                    }
                                }
                                var pos = comboBoxWalletsAddressesSendFrom.checkedElements.indexOf(index)
                                if (pos >= 0) {
                                    comboBoxWalletsAddressesSendFrom.checkedElements.splice(pos, 1)
                                    comboBoxWalletsAddressesSendFrom.checkedElementsText.splice(pos, 1)
                                }
                            }
							subPageSendAdvanced.updateInfo();
                            comboBoxWalletsAddressesSendFrom.numberOfCheckedElements = comboBoxWalletsAddressesSendFrom.checkedElements.length
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
                    text: qsTr("All outputs of the selected addresses")
                    checked: true
					onClicked: {
						subPageSendAdvanced.updateInfo()
					}
                }
            }

            ComboBox {
                id: comboBoxWalletsUnspentOutputsSendFrom

                function getCheckedDelegates() {
                    return checkedElements
                }
                property var checkedElements: []
                property var checkedElementsText: []
                property int numberOfCheckedElements: checkedElements.length
                property alias filterString: filterPopupOutputs.filterText

                Layout.fillWidth: true
                Layout.topMargin: -12
                textRole: "outputID"
                displayText: checkBoxUnspentOutputsUseAllOutputs.checked ? qsTr("All outputs selected") : numberOfCheckedElements > 1 ? (numberOfCheckedElements + ' ' + qsTr("outputs selected")) : numberOfCheckedElements === 1 ? checkedElementsText[0] : qsTr("No output selected")

                enabled: !checkBoxUnspentOutputsUseAllOutputs.checked
                model: QOutputs {
                    id: listOutputs
                }

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
							subPageSendAdvanced.updateInfo();
                        }

                        width: parent.width
                        text: comboBoxWalletsUnspentOutputsSendFrom.textRole ? (Array.isArray(comboBoxWalletsUnspentOutputsSendFrom.model) ? modelData[comboBoxWalletsUnspentOutputsSendFrom.textRole] + " - " + modelData["addressSky"] + " SKY (" + modelData["addressCoinHours"] + " CoinHours)" :model[comboBoxWalletsUnspentOutputsSendFrom.textRole] + " - " + model["addressSky"] + " SKY (" + model["addressCoinHours"] + " CoinHours)") : modelData
                        font.family: "Code New Roman"
                        // Load the saved state when the delegate is recicled:
                        checked: comboBoxWalletsUnspentOutputsSendFrom.checkedElements.indexOf(index) >= 0
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
            text: "With your current selection you can send up to <b>" + subPageSendAdvanced.upperCoinBound + " SKY</b> and <b>" + subPageSendAdvanced.upperAltCointBound + " Coin Hours</b> (at least <b>" + subPageSendAdvanced.minFeeAmount + " Coin Hours</b> must be used for the transaction fee)"
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

                Layout.fillWidth: true
                Layout.topMargin: -16
                implicitHeight: contentItem.height
                Behavior on implicitHeight { NumberAnimation { duration: 250; easing.type: Easing.OutQuint } }

                interactive: false
                clip: true

                model: listModelDestinations

                delegate: DestinationListDelegate {
                    width: listViewDestinations.width
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
                        modelAddressesByWallet.loadModel(walletManager.getAllAddresses())
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

            Layout.fillWidth: true

            Layout.alignment: Qt.AlignTop

            CheckBox {
                id: checkBoxAutomaticCoinHoursAllocation
                text: qsTr("Automatic coin hours allocation")
                checked: true
            }

            Slider {
                id: sliderCoinHoursShareFactor

                Layout.preferredWidth: parent.width < 500 ? 500 : parent.width

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

    AddressModel{
        id: modelAddressesByWallet
    }

    ListModel {
        id: listModelDestinations
        ListElement { address: ""; sky: "0.0"; coinHours: "0.0" }
    }
}
