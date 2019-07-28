import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Dialog {
    id: walletCreateDialog

    title: Qt.application.name
    standardButtons: Dialog.Ok
    closePolicy: Dialog.NoAutoClose

    Component.onCompleted: {
        standardButton(Dialog.Ok).enabled = name.text !== ""
    }

    onAboutToShow: {
        name.forceActiveFocus()
    }

    Flickable {
        id: flickable
        anchors.fill: parent
        contentHeight: columnLayoutRoot.height
        clip: true

        ColumnLayout {
            id: columnLayoutRoot
            width: parent.width
            spacing: 30

            RowLayout {
                spacing: 30

                Image {
                    id: icon
                    source: "qrc:/images/resources/images/icons/check.svg"
                    sourceSize: "64x64"
                    Layout.alignment: Qt.AlignTop
                }

                ColumnLayout {
                    Label {
                        text: qsTr("The connected hardware wallet will be added " +
                                "to the wallets list with the following name:")
                        Layout.fillWidth: true
                        wrapMode: Text.WordWrap
                        Layout.alignment: Qt.AlignTop
                    }

                    TextField {
                        id: name
                        placeholderText: qsTr("New hardware wallet")
                        Layout.fillWidth: true
                        focus: walletCreateDialog.focus

                        onTextChanged: {
                            standardButton(Dialog.Ok).enabled = text !== ""
                        }
                    }
                } // ColumnLayout (message)
            } // RowLayout (icon + message)

            Label {
                text: qsTr("Now you can check the balance and the address of the " +
                        "hardware wallet even when it is not connected. " +
                        "Additionally, you can change the name of the wallet " +
                        "or remove it from the wallets list, if you wish.")
                Layout.fillWidth: true
                wrapMode: Text.WordWrap
                Layout.alignment: Qt.AlignTop
            }
        } // ColumnLayout (root)

        ScrollIndicator.vertical: ScrollIndicator {
            parent: walletCreateDialog.contentItem
            anchors.top: flickable.top
            anchors.bottom: flickable.bottom
            anchors.right: parent.right
            anchors.rightMargin: -walletCreateDialog.rightPadding + 1
        }
    } // Flickable
}
