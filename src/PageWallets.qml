import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Page {
    id: root

    property string statusIcon; // an empty string for no icon
    readonly property real listWalletLeftMargin: 20
    readonly property real listWalletRightMargin: 50
    readonly property real listWalletSpacing: 20
    readonly property real internalLabelsWidth: 70

    header: ColumnLayout {
        Label {
            text: qsTr("Wallets")
            font.pixelSize: Qt.application.font.pixelSize * 2
            padding: 10
        }

        RowLayout {
            spacing: listWalletSpacing
            Label {
                text: qsTr("Name")
                font.pointSize: 9
                Layout.leftMargin: listWalletLeftMargin
                Layout.fillWidth: true
            }
            Label {
                text: qsTr("Sky")
                font.pointSize: 9
                horizontalAlignment: Text.AlignRight
                Layout.preferredWidth: internalLabelsWidth
            }
            Label {
                text: qsTr("Coin hours")
                font.pointSize: 9
                horizontalAlignment: Text.AlignRight
                Layout.rightMargin: listWalletRightMargin
                Layout.preferredWidth: internalLabelsWidth
            }
        }

        Rectangle {
            Layout.fillWidth: true
            height: 1
            color: "#FFDDDDDD"
        }
    }

    footer: RowLayout {
        ItemDelegate {
            id: buttonAddWallet
            text: qsTr("Add wallet")
            Layout.fillWidth: true
            display: ItemDelegate.TextUnderIcon
        }
        ItemDelegate {
            id: buttonLoadWallet
            text: qsTr("Load wallet")
            Layout.fillWidth: true
            display: ItemDelegate.TextUnderIcon
        }
    }

    ScrollView {
        id: scrollItem
        anchors.fill: parent
        ListView {
            id: walletList
            anchors.fill: parent
            model: listModel
            delegate: delegateWalletList
        }
    }

    // Roles: name, encryptionEnabled, sky, coinHours
    // Use listModel.append( { "name": value, "encryptionEnabled": value, "sky": value, "coinHours": value } )
    // Or implement the model in the backend (a more recommendable approach)
    ListModel {
        id: listModel
    }

    Component {
        id: delegateWalletList

        ItemDelegate {
            width: walletList.width

            RowLayout {
                anchors.fill: parent
                anchors.leftMargin: 30
                anchors.rightMargin: 50
                spacing: 20

                // TODO: add an 'encryption-disabled' SVG icon
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
                    source: "qrc:/images/security.svg"
                    sourceSize: "24x24"
                }

                Label {
                    id: labelskycoins
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
        }
    }
}
