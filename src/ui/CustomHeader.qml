import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Label {
    id: labelHeaderText
    anchors.top: parent.top
    anchors.topMargin: (customMenuBar.toolButtonBack.hide) ? -(menuBar.height + height) : -menuBar.height
    anchors.left: parent.left
    anchors.right: parent.right
    height: menuBar.height
    font.pointSize: Qt.application.font.pointSize * 1.6
    font.bold: true
    color: "white"
    opacity: customMenuBar.toolButtonBack.opacity
    horizontalAlignment: Text.AlignHCenter
    verticalAlignment: Text.AlignVCenter
    z: 10
    Behavior on anchors.topMargin { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } }
}
