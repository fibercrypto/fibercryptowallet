import QtQuick 2.12
import QtQuick.Controls 2.12

ToolButton {
    id: toolButtonQR

    property size iconSize: "16x16"
    icon.source: "qrc:/images/resources/images/icons/qr.svg"
    icon.width: iconSize.width
    icon.height: iconSize.height
    onClicked: {
//        dialogQR.text = toolButtonQR.qrSource
//        console.log("Making var pass")
//        console.log(toolButtonQR.qrSource)
//        dialogQR.open()
    }
    ToolTip.text: qsTr("Show QR code")
    ToolTip.visible: hovered // TODO: pressed when mobile?
    ToolTip.delay: Qt.styleHints.mousePressAndHoldInterval
}