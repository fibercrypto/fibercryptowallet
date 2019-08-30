import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Item {
    id: settingsListDelegate

    implicitHeight: Math.max(textOutputID.height, (toolButtonCopy.height - toolButtonCopy.topPadding*2), labelAddressSky.height, labelAddressCoins.height)

    RowLayout {
        id: rowLayoutRoot
        anchors.fill: parent
        anchors.leftMargin: listOutputsLeftMargin * 2
        anchors.rightMargin: listOutputsRightMargin
        spacing: listOutputsSpacing

        RowLayout {
            id: rowLayoutOutputID
            Layout.fillWidth: true

            TextInput {
                id: textOutputID
                Layout.fillWidth: true
                text: outputID // a role of the model
                readOnly: true
                font.family: "Code New Roman"
                wrapMode: TextInput.WrapAnywhere
            }
            ToolButton {
                id: toolButtonCopy
                icon.source: "qrc:/images/resources/images/icons/copy.svg"
                Layout.alignment: Qt.AlignLeft
                ToolTip.text: qsTr("Copy to clipboard")
                ToolTip.visible: hovered // TODO: pressed when mobile?
                ToolTip.delay: Qt.styleHints.mousePressAndHoldInterval

                Image {
                    id: imageCopied
                    anchors.centerIn: parent
                    source: "qrc:/images/resources/images/icons/check-simple.svg"
                    fillMode: Image.PreserveAspectFit
                    sourceSize: Qt.size(toolButtonCopy.icon.width*1.5, toolButtonCopy.icon.height*1.5)
                    z: 1

                    opacity: 0.0
                }

                onClicked: {
                    textOutputID.selectAll()
                    textOutputID.copy()
                    textOutputID.deselect()
                    if (copyAnimation.running) {
                        copyAnimation.restart()
                    } else {
                        copyAnimation.start()
                    }
                }

                SequentialAnimation {
                    id: copyAnimation
                    NumberAnimation { target: imageCopied; property: "opacity"; to: 1.0; easing.type: Easing.OutCubic }
                    PauseAnimation { duration: 1000 }
                    NumberAnimation { target: imageCopied; property: "opacity"; to: 0.0; easing.type: Easing.OutCubic }
                }
            } // ToolButton
        } // RowLayout (output ID)

        Label {
            id: labelAddressSky
            text: addressSky // a role of the model
            color: Material.accent
            horizontalAlignment: Text.AlignRight
            Layout.preferredWidth: internalLabelsWidth/2
        }

        Label {
            id: labelAddressCoins
            text: addressCoinHours // a role of the model
            horizontalAlignment: Text.AlignRight
            Layout.preferredWidth: internalLabelsWidth
        }
    } // RowLayout (addresses)
}
