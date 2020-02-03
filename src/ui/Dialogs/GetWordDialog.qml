import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12

Dialog {
    property string hint: ""
    property string label: ''
    property alias selectedName : editTextItem.text

    standardButtons: Dialog.Ok | Dialog.Cancel
    onVisibleChanged: {
        editTextItem.focus = true
        editTextItem.selectAll()
    }
//    onButtonClicked: {
//        Qt.inputMethod.hide();
//    }
    Flickable {
        id: flickable
        anchors.fill: parent
        contentHeight: columnLayoutRoot.height
        clip: true

        ColumnLayout {
            width: parent.width
//            implicitWidth: parent.width
            //        implicitHeight: 100
            Label {
                id: labelHeaderMessage
                text: label
                Layout.fillWidth: true
                wrapMode: Text.WordWrap
            }
            TextField {
                id: editTextItem
                text: hint
                color: Style.text
            }
        }
    }
}
