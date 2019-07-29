import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui"
import "../" // For quick UI development, switch back to resources when making a release

Dialog {
    id: dialogAddWallet

    title: Qt.application.name
    standardButtons: Dialog.Ok | Dialog.Cancel

    onAboutToShow: {
        createLoadWallet.clear()
    }

    Flickable {
        id: flickable
        anchors.fill: parent
        contentHeight: columnLayoutRoot.height
        clip: true

        ColumnLayout {
            id: columnLayoutRoot
            width: parent.width

            CreateLoadWallet {
                id: createLoadWallet
                Layout.fillWidth: true
            }

            CheckBox {
                id: checkBoxEncryptWallet
                text: qsTr("Encrypt wallet")

                onCheckedChanged: {
                    if (checked) {
                        textFieldPassword.forceActiveFocus()
                    }
                }
            }

            Label {
                Layout.fillWidth: true
                Layout.leftMargin: checkBoxEncryptWallet.indicator.width + 2*checkBoxEncryptWallet.padding
                text: qsTr("We suggest that you encrypt each one of your wallets with a password. "
                         + "If you forget your password, you can reset it with your seed. "
                         + "Make sure you have your seed saved somewhere safe before encrypting your wallet.")
                wrapMode: Label.WordWrap
            }

            RowLayout {
                id: rowLayoutPassword
                spacing: 10
                enabled: checkBoxEncryptWallet.checked

                TextField {
                    id: textFieldPassword
                    Layout.fillWidth: true
                    
                    placeholderText: qsTr("Password")
                    echoMode: TextField.Password
                    selectByMouse: true
                }
                TextField {
                    id: textFieldConfirmPassword
                    Layout.fillWidth: true
                    
                    placeholderText: qsTr("Confirm password")
                    echoMode: TextField.Password
                    selectByMouse: true
                }
            } // RowLayout
        } // ColumnLayout (root)

        ScrollIndicator.vertical: ScrollIndicator {
            parent: dialogAddWallet.contentItem
            anchors.top: flickable.top
            anchors.bottom: flickable.bottom
            anchors.right: parent.right
            anchors.rightMargin: -dialogAddWallet.rightPadding + 1
        }
    } // Flickable
}