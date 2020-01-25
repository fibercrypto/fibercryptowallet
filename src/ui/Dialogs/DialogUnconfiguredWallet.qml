import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import DeviceInteraction 1.0

Dialog {
    id: dialogUnconfiguredWallet

    title: qsTr("Unconfigured wallet")
    standardButtons: Dialog.Abort
    closePolicy: Dialog.NoAutoClose
    DeviceInteraction {
        id: deviceInteraction2
        onOperationDone: {
            deviceInteraction2.deviceFeatures();
        }
        onIsInitializedDetermined: {
            buttonAutoConfDevice.visible = !isInitialized;
            buttonRecoveryDevice.visible = !isInitialized;
        }
    }
    onAboutToShow: deviceInteraction2.deviceFeatures();
    onRejected: topLevelDialogLocker.unlock();

    Flickable {
        id: flickable
        clip: true
        anchors.fill: parent
        contentHeight: columnLayoutRoot.height

        ColumnLayout {
            id: columnLayoutRoot
            width: parent.width
            spacing: 50

            RowLayout {
                spacing: 50

                Image {
                    id: icon
                    source: "qrc:/images/resources/images/icons/cpu.svg"
                    sourceSize: "64x64"
                    Layout.alignment: Qt.AlignTop
                }

                Label {
                    id: msg
                    text: qsTr("A seedless hardware wallet has been detected. " +
                            "Select <b><i>Configure automatically</i></b> " +
                            "if you want to configure it as a brand new wallet " +
                            "and start using it inmediately, " +
                            "or select <b><i>Restore backup</i></b> " +
                            "if you want to configure it with a previously created " +
                            "seed backup and thus be able to access your funds again.")
                    wrapMode: Text.WordWrap
                    Layout.fillWidth: true
                    Layout.alignment: Qt.AlignTop
                }
            } // RowLayout (icon + message)

            ColumnLayout {
                Label {
                    id: options
                    text: qsTr("Options:")
                    font.bold: true
                    Layout.fillWidth: true
                }
                RowLayout {
                    id: buttonAutoConfDevice
                    Layout.fillWidth: true
                    ItemDelegate {
                        text: qsTr("Configure automatically")
                        Layout.fillWidth: true
                        MouseArea {
                            anchors.fill: parent
                            onClicked: {
                                deviceInteraction2.generateMnemonic(+generateWordCount.currentText, generateUsePassphrase.checked);
                            }
                        }
                    }
                    CheckBox {
                        id: generateUsePassphrase
                        text: qsTr("Use pass phrase")
                        enabled: false
                    }
                    ComboBox {
                        id: generateWordCount
                        model: [12, 24]
                    }
                }
                RowLayout {
                    id: buttonRecoveryDevice
                    Layout.fillWidth: true
                    ItemDelegate {
                        text: qsTr("Restore backup")
                        Layout.fillWidth: true
                        onClicked: {
                            deviceInteraction2.restoreBackup(+recoveryWordCount.currentText, recoveryUsePassphrase.checked)
                        }
                    }
                    CheckBox {
                        id: recoveryUsePassphrase
                        text: qsTr("Use pass phrase")
                        enabled: false
                    }
                    ComboBox {
                        id: recoveryWordCount
                        model: [12, 24]
                    }
                }
            }
        } // ColumnLayout (root)

        ScrollIndicator.vertical: ScrollIndicator {
            parent: dialogUnconfiguredWallet.contentItem
            anchors.top: flickable.top
            anchors.bottom: flickable.bottom
            anchors.right: parent.right
            anchors.rightMargin: -dialogUnconfiguredWallet.rightPadding + 1
        }
    } // Flickable
}
