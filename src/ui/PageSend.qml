import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/Dialogs"
import "Dialogs/" // For quick UI development, switch back to resources when making a release

Page {
    id: root

    property alias advancedMode: switchAdvancedMode.checked

    footer: ToolBar {
        id: tabBarSend
        Material.primary: Material.Blue
        Material.accent: Material.Yellow

        ToolButton {
            id: buttonAddWallet
            anchors.fill: parent
            text: qsTr("Send")
            icon.source: "qrc:/images/resources/images/icons/send.svg"

            onClicked: {
                dialogSend.open()
            }
        }
    }

    GroupBox {
        id: groupBox

        readonly property real margins: 50

        anchors.fill: parent
        anchors.topMargin: 10
        anchors.leftMargin: margins
        anchors.rightMargin: margins
        anchors.bottomMargin: margins
        implicitWidth: stackView.width
        clip: true

        label: SwitchDelegate {
            id: switchAdvancedMode

            text: qsTr("Advanced mode")

            onToggled: {
                if (checked) {
                    stackView.push(componentAdvanced)
                } else {
                    stackView.pop()
                }
            }
        }

        StackView {
            id: stackView
            anchors.fill: parent
            initialItem: componentSimple
            clip: true

            Component {
                id: componentSimple
                ScrollView {
                    contentWidth: simple.implicitWidth
                    contentHeight: simple.implicitHeight
                    clip: true
                    SubPageSendSimple {
                        id: simple
                        implicitWidth: stackView.width
                    }
                }
            }

            Component {
                id: componentAdvanced
                ScrollView {
                    contentWidth: advanced.implicitWidth
                    contentHeight: advanced.implicitHeight
                    clip: true
                    SubPageSendAdvanced {
                        id: advanced
                        implicitWidth: stackView.width
                    }
                }
            }
        }
    } // GroupBox

    DialogSendTransaction {
        id: dialogSend
        anchors.centerIn: Overlay.overlay
    }
}
