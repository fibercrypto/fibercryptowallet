import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/Controls"
import "Controls" // For quick UI development, switch back to resources when making a release

RowLayout {
    property alias nodeUrl: nodeUrlField.text
    width: parent.width
    height: 40
    Label {
        font.pointSize: Qt.application.font.pointSize * 0.9
        font.bold: true
        Layout.alignment: Qt.AlignCenter
        text: "Node Url:"
    }
    TextField {
        id: nodeUrlField
        Layout.fillWidth: true
        Layout.alignment: Qt.AlignCenter
        placeholderText: qsTr("http://127.0.0.1:6420")
        selectByMouse: true
    }
}
