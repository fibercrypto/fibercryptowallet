import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12

// Resource imports
// import "qrc:/ui/src/ui/Controls"
import "Controls" // For quick UI development, switch back to resources when making a release

Item {
    id: root

    implicitWidth: 100
    implicitHeight: 60
    signal dataModified()
    property Item nextTabItem
    property alias textArea: textArea
    property alias text: textArea.text
    property alias placeholderText: textArea.placeholderText
    property alias readOnly: textArea.readOnly
    property alias buttonLeftText: buttonLeft.text
    property alias buttonRightText: buttonRight.text

    property bool buttonsVisible: buttonRight.visible

    readonly property alias inputControlWidth: textArea.width
    readonly property alias inputControlHeight: textArea.height

    function generateSeed( words){
        if (words == 12 ){
            textArea.text = walletManager.getNewSeed(128)
        } else if (words == 24){
            textArea.text = walletManager.getNewSeed(256)
        }
    }
    function clear() {
        textArea.clear()
        if (mode === CreateLoadWallet.Create){
            generateSeed(12)
        }
        
    }

    Row {
        id: rowButtons
        anchors.bottom: textArea.top
        anchors.bottomMargin: -10
        anchors.right: textArea.right
        z: textArea.z + 1

        ItemDelegate {
            id: buttonLeft
            onClicked: generateSeed(12)
            
            height: buttonRight.height

            text: "ButtonLeft"
            font.pointSize: 8
            visible: buttonRight.visible
            opacity: buttonRight.opacity

            contentItem: Label {
                text: buttonLeft.text
                font: buttonLeft.font
                color: buttonLeft.enabled ? buttonLeft.Material.foreground : buttonLeft.Material.hintTextColor
            }
        }

        ItemDelegate {
            id: buttonRight
            onClicked:  generateSeed(24)
                
                   
            height: 30
            opacity: buttonsVisible ? 1.0 : 0.0
            visible: opacity > 0.0

            Behavior on opacity { NumberAnimation { easing.type: Easing.OutQuint } }

            text: "ButtonRight"
            font.pointSize: 8

            contentItem: Label {
                text: buttonRight.text
                font: buttonRight.font
                color: buttonRight.enabled ? buttonRight.Material.foreground : buttonRight.Material.hintTextColor
            }
        } // ItemDelegate
    } // Row

    /*
    // Put the text area inside this to make it scrollable
    // but we have a bug.
    // Set the anchors of the left button relative to this ScrollView
    ScrollView {
        id: scrollerTextArea
        anchors.left: parent.left
        anchors.right: parent.right

        width: root.width
        height: root.height - buttonRight.height

        contentWidth: width
        contentHeight: height

        clip: true
    }
    */

    TextArea {
        id: textArea
        anchors.left: parent.left
        anchors.right: parent.right
        height: 80
        wrapMode: TextArea.Wrap
        clip: true

        selectByMouse: true
        KeyNavigation.priority: KeyNavigation.BeforeItem
        KeyNavigation.tab: nextTabItem
        onTextChanged: dataModified()
    }
}
