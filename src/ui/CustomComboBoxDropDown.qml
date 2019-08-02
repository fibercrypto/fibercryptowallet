import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.impl 2.12
import QtQuick.Controls.Material 2.12 
import QtQuick.Controls.Material.impl 2.12

Popup {
    id: customComboBoxDropDown

    property ComboBox comboBox

    y: comboBox.editable ? comboBox.height - 5 : 0
    width: comboBox.width
    height: Math.min(contentItem.implicitHeight, applicationWindow.height - topMargin - bottomMargin)
    transformOrigin: Item.Top
    margins: 0
    padding: 0

    Material.theme: comboBox.Material.theme
    Material.accent: comboBox.Material.accent
    Material.primary: comboBox.Material.primary

    enter: Transition {
        // grow_fade_in
        NumberAnimation { property: "scale"; from: 0.9; to: 1.0; easing.type: Easing.OutQuint; duration: 220 }
        NumberAnimation { property: "opacity"; from: 0.0; to: 1.0; easing.type: Easing.OutCubic; duration: 150 }
    }

    exit: Transition {
        // shrink_fade_out
        NumberAnimation { property: "scale"; from: 1.0; to: 0.9; easing.type: Easing.OutQuint; duration: 220 }
        NumberAnimation { property: "opacity"; from: 1.0; to: 0.0; easing.type: Easing.OutCubic; duration: 150 }
    }

    contentItem: ListView {
        clip: true
        implicitHeight: contentHeight
        model: comboBox.delegateModel
        currentIndex: comboBox.highlightedIndex
        highlightMoveDuration: 0

        section.property: "wallet"
        section.criteria: ViewSection.FullString
        section.delegate: Label {
            text: section
            leftPadding: 12
            font.bold: true
            font.pointSize: Qt.application.font.pointSize * 1.5
        }

        ScrollIndicator.vertical: ScrollIndicator { }
    }

    background: Rectangle {
        radius: 2
        color: parent.Material.dialogColor

        layer.enabled: comboBox.enabled
        layer.effect: ElevationEffect {
            elevation: 8
        }
    }
}