import QtQuick 2.7
import QtQuick.Layouts 1.2
import QtQuick.Controls.Material 2.12

Column {
  id: root
  property string title: "Accordion"
  default property alias item: ld.sourceComponent
  property alias itemWidth: info.width
  Rectangle {
    width: columnSettings.width
    height: 40
    color: Material.color(Material.Blue)
    MouseArea {
      anchors.fill: parent
      onClicked: info.show = !info.show
    }
    Text {
        anchors.fill: parent
        anchors.margins: 10
        horizontalAlignment: Text.AlignHCenter
        verticalAlignment: Text.AlignVCenter
        text: root.title
    }
    Text {
        anchors{
            right: parent.right
            top: parent.top
            bottom: parent.bottom
            margins: 10
        }
        horizontalAlignment: Text.AlignRight
        verticalAlignment: Text.AlignVCenter
        text: "^"
        rotation: info.show ? 0 : "180"
    }

  }
  Rectangle {
    id: info
    width: show ? ld.width : 0
    height: show ? ld.height : 0
    property bool show : false
//    color: generalStackView.color // TODO Set proper color
    clip: true
    Loader {
      id: ld
      y: info.height - height
      anchors.horizontalCenter: info.horizontalCenter
    }
    Behavior on height {
      NumberAnimation { duration: 50; easing.type: Easing.InOutQuad }
    }
  }
}