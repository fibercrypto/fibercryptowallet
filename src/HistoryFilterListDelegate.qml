import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Item {
    id: root

    property alias tristate: checkDelegate.tristate
    property alias walletText: checkDelegate.text

    ColumnLayout {
        width: root.width

        CheckDelegate {
            id: checkDelegate

            Layout.fillWidth: true
            tristate: true
            text: name
            LayoutMirroring.enabled: true

            contentItem: Label {
                leftPadding: checkDelegate.indicator.width + checkDelegate.spacing
                text: checkDelegate.text
                verticalAlignment: Qt.AlignVCenter
                color: checkDelegate.enabled ? checkDelegate.Material.foreground : checkDelegate.Material.hintTextColor
            }
        }
    } // ColumnLayout
}
