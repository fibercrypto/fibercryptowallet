import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
ColumnLayout{
RowLayout{
    Label {
 Layout.leftMargin: 5
        id: label
        text: section
        Layout.fillWidth: true
    }
    }
    RowLayout{
    Rectangle {
                id: rect
                Layout.fillWidth: true
                height: 1
                color: "#DDDDDD"
            }
            }
    }
