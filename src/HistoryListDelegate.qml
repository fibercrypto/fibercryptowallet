import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

ItemDelegate {
    id: root

    RowLayout {
        id: rowLayoutRoot

        Image {
            source: "qrc:/images/send-blue.svg"
            sourceSize: "32x32"
            fillMode: Image.PreserveAspectFit
            mirror: type === HistoryListDelegate.Type.Receive
            Layout.alignment: Qt.AlignTop | Qt.AlignLeft
        }

    } // RowLayout (root)
}
