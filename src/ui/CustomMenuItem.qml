import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

MenuItem {
    id: customMenuItem

    property alias iconSource: imageIcon.source
    property alias iconSourceSize: imageIcon.sourceSize

    leftPadding: highlighted ? imageIcon.sourceSize.width + 12 : padding
    Behavior on leftPadding { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } }

    Image {
        id: imageIcon
        anchors.left: parent.left
        anchors.leftMargin: parent.highlighted ? 6 : -width/2
        Behavior on anchors.leftMargin { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } }
        anchors.verticalCenter: parent.verticalCenter
        
        sourceSize: "24x24" // same as Qt.size(24, 24)
        fillMode: Image.PreserveAspectFit

        visible: opacity > 0
        opacity: parent.highlighted ? 1.0 : 0.0
        Behavior on opacity { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } }
    }
}