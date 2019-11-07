import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12

ItemDelegate {
    id: contactDelegate

    checkable: true

    contentItem: ColumnLayout {
        spacing: 10

        Label {
            text: name
            font.bold: true
            elide: Text.ElideRight
            Layout.fillWidth: true
        }

        GridLayout {
            id: grid
            visible: false

            columns: 2
            rowSpacing: 10
            columnSpacing: 10

            Label {
                text: qsTr("Address:")
//                Layout.leftMargin: 60
            }

            Label {
                text: address
                font.bold: true
                elide: Text.ElideRight
                Layout.fillWidth: true
            }
        }
    }

    states: [
        State {
            name: "expanded"
            when: contactDelegate.checked

            PropertyChanges {
                target: grid
                visible: true
            }
        }
    ]
}