import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/"
// import "qrc:/ui/src/ui/Dialogs"
import "../Dialogs/" // For quick UI development, switch back to resources when making a release

Item {
    id: root

    readonly property real delegateHeight: 30
    property bool emptyAddressVisible: true
    property bool expanded: false
    // The following property is used to avoid a binding conflict with the `height` property.
    // Also avoids a bug with the animation when collapsing a wallet
    readonly property real finalViewHeight: expanded ? delegateHeight*(addressList.count) + 50 : 0

    width: walletList.width
    height: itemDelegateMainButton.height + (expanded ? finalViewHeight : 0)

    Behavior on height { NumberAnimation { duration: 250; easing.type: Easing.OutQuint } }

    ColumnLayout {
        id: delegateColumnLayout
        anchors.fill: parent

        ItemDelegate {
            id: itemDelegateMainButton
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignTop
            font.bold: expanded

            RowLayout {
                id: delegateRowLayout
                anchors.fill: parent
                anchors.leftMargin: listWalletLeftMargin
                anchors.rightMargin: listWalletRightMargin
                spacing: listWalletSpacing

                Image {
                    id: status
                    source: statusIcon
                    sourceSize: "24x24"
                }

                Label {
                    id: labelWalletName
                    text: name // a role of the model
                    Layout.fillWidth: true
                }

                Image {
                    id: lockIcon
                    source: "qrc:/images/resources/images/icons/lock" + (encryptionEnabled ? "On" : "Off") + ".svg"
                    sourceSize: "24x24"
                }

                Label {
                    id: labelSky
                    text: sky // a role of the model
                    color: Material.accent
                    horizontalAlignment: Text.AlignRight
                    Layout.preferredWidth: internalLabelsWidth
                }

                Label {
                    id: labelCoins
                    text: coinHours // a role of the model
                    horizontalAlignment: Text.AlignRight
                    Layout.preferredWidth: internalLabelsWidth
                }
            }

            onClicked: {
                expanded = !expanded
            }
        } // ItemDelegate

        ListView {
            id: addressList
            model: listAddresses
            implicitHeight: expanded ? delegateHeight*(addressList.count) + 50 : 0
            opacity: expanded ? 1.0 : 0.0
            clip: true
            interactive: false
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignTop

            Behavior on implicitHeight { NumberAnimation { duration: 250; easing.type: Easing.OutQuint } }
            Behavior on opacity { NumberAnimation { duration: expanded ? 250 : 1000; easing.type: Easing.OutQuint } }

            delegate: WalletListAddressDelegate {
                width: walletList.width
                height: index == 0 ? delegateHeight + 20 : visible ? delegateHeight : 0

                onAddAddressesRequested: {
                    dialogAddAddresses.open()
                }
            }
        }
    } // ColumnLayout

    DialogAddAddresses {
        id: dialogAddAddresses
        anchors.centerIn: Overlay.overlay

        modal: true
        focus: true

        onAccepted: {
            console.log("Editting accepted")
        }
        onRejected: {
            console.log("Editting rejected")
        }
    }

    // Roles: address, addressSky, addressCoinHours
    // Use listModel.append( { "address": value, "addressSky": value, "addressCoinHours": value } )
    // Or implement the model in the backend (a more recommendable approach)
    ListModel {
        id: listAddresses
        // The first element must exist but will not be used
        ListElement { address: "--------------------------"; addressSky: 0; addressCoinHours: 0 }
        ListElement { address: "qrxw7364w8xerusftaxkw87ues"; addressSky: 30; addressCoinHours: 1049 }
        ListElement { address: "8745yuetsrk8tcsku4ryj48ije"; addressSky: 12; addressCoinHours: 16011 }
        ListElement { address: "gfdhgs343kweru38200384uwqd"; addressSky: 0; addressCoinHours: 72 }
        ListElement { address: "00qdqsdjkssvmchskjkxxdg374"; addressSky: 521; addressCoinHours: 11 }
    }
}
