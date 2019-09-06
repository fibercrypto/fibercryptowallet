import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Layouts 1.12

Dialog {
    id: dialogAddAddresses

    title: Qt.application.name
    standardButtons: Dialog.Ok | Dialog.Cancel
    closePolicy: Dialog.NoAutoClose
    property alias spinValue: spinBoxAmount.value

    onAboutToShow: {
        spinBoxAmount.forceActiveFocus()
    }

    ColumnLayout {
        id: columnLayoutRoot
        anchors.fill: parent

        Label {
            text: qsTr("Amount")
            font.bold: true
            Layout.fillWidth: true
        }
        SpinBox {
            id: spinBoxAmount
            Layout.fillWidth: true
            from: 1
            to: 100
            value: 1
            editable: true
            focus: true
        }
    } // ColumnLayout (root)
}