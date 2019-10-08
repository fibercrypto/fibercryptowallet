import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Layouts 1.12

Popup {

    property ComboBox comboBox: null
    property alias filterText: textFieldFilter.text
    property alias filterPlaceholderText: textFieldFilter.placeholderText

    y: comboBox.editable ? comboBox.height - 5 : 0
    width: comboBox.width
    height: Math.min(contentItem.implicitHeight + topPadding + bottomPadding, Overlay.overlay.height - topMargin - bottomMargin)
    transformOrigin: Item.Top
    topMargin: 12
    bottomMargin: 12
    padding: 0
    topPadding: 6

    onAboutToShow: {
        textFieldFilter.forceActiveFocus()
    }

    contentItem: ColumnLayout {

        TextField {
            id: textFieldFilter

            Layout.fillWidth: true
            Layout.leftMargin: 12
            Layout.rightMargin: 12
            placeholderText: qsTr("Filter")
            focus: true
        }

        ListView {
            Layout.fillWidth: true
            Layout.fillHeight: true
            clip: true
            implicitHeight: contentHeight
            model: comboBox.delegateModel
            currentIndex: comboBox.highlightedIndex
            highlightMoveDuration: 0

            ScrollIndicator.vertical: ScrollIndicator { }
        }
    } // ColumnLayout (contentItem)
}
