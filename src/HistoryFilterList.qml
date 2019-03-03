import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Item {
    id: root

    readonly property real delegateHeight: 48
    readonly property alias count: view.count
    property alias interactive: view.interactive

    implicitWidth: view.width
    implicitHeight: view.height

    ListView {
        id: view

        height: delegateHeight * count
        width: 300
        clip: true

        model: modelFilters
        delegate: HistoryFilterListDelegate {
            id: filterDelegate
            width: ListView.view.width
            height: delegateHeight
        }
    }

    // This model can be the same as the wallet list,
    // as this model need to expose all wallets and their addresses.
    // For that, it should be implemented in the backend, instead of here.
    ListModel { // EXAMPLE
        id: modelFilters

        ListElement { name: "My first wallet" }
        ListElement { name: "My second wallet" }
        ListElement { name: "My third wallet" }
        ListElement { name: "My fourth wallet" }
        ListElement { name: "My fiveth wallet" }
        ListElement { name: "My sixth wallet" }
    }
}
