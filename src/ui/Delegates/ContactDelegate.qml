import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12
import "../Dialogs"
ItemDelegate {
    id: contactDelegate

onClicked:{
dialogShowContact.open()
}
    checkable: true

DialogShowContact{
id:dialogShowContact
width:300
height:300
}

    contentItem: ColumnLayout {
        spacing: 10


        Label {
            text: name
            font.bold: true
            elide: Text.ElideRight
            Layout.fillWidth: true
        }

//        GridLayout {
//            id: grid
//            visible: false
//
//            columns: 2
//            rowSpacing: 10
//            columnSpacing: 10
//
//            Label {
//                text: qsTr("Addresses:")
////                Layout.leftMargin: 60
//            }
//        }

}
//    states: [
//        State {
//            name: "expanded"
//            when: contactDelegate.checked
//
//            PropertyChanges {
//                target: grid
//                visible: true
//            }
//        }
//    ]
}
