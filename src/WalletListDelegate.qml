import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

ItemDelegate {
    id: root
    width: walletList.width

    RowLayout {
        id: delegateRowLayout
        anchors.fill: parent
        anchors.leftMargin: listWalletLeftMargin
        anchors.rightMargin: listWalletRightMargin
        spacing: listWalletSpacing

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
}
