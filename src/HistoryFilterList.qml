import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Item {
    id: root

    ListView {
        id: view

        height: 48 * count
        width: 300
        clip: true
    }
}
