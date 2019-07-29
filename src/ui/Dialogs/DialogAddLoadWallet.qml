import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui"
import "../" // For quick UI development, switch back to resources when making a release

Dialog {
    id: dialogAddWallet

    property alias mode: createLoadWallet.mode

    Component.onCompleted: {
        standardButton(Dialog.Ok).text = mode === CreateLoadWallet.Create ? qsTr("Create") : qsTr("Load")
    }

    onModeChanged: {
        standardButton(Dialog.Ok).text = mode === CreateLoadWallet.Create ? qsTr("Create") : qsTr("Load")
    }

    function updateAcceptButtonStatus() {
        var walletName = createLoadWallet.name
        var walletSeed = createLoadWallet.seed
        var walletSeedConfirm = createLoadWallet.seedConfirm
        
        var words = walletSeed.split(' ')

        var invalidNumberOfWords = words.length != 12 && words.length != 24
        var invalidChars = walletSeed.search("[^a-z ]|[ ]{2}") != -1
        var unconventionalSeed = invalidNumberOfWords || invalidChars
        var continueWithUnconventionalSeed = checkBoxContinueWithSeedWarning.checked
        
        var seedMatchConfirmation = walletSeed === walletSeedConfirm

        var passwordNeeded = checkBoxEncryptWallet.checked
        var passwordSet = textFieldPassword.text
        var passwordMatchConfirmation = textFieldPassword.text === textFieldConfirmPassword.text

        columnLayoutSeedWarning.warn = walletName && walletSeed && seedMatchConfirmation && !(!unconventionalSeed)

        var okButton = standardButton(Dialog.Ok)
        okButton.enabled = walletName && walletSeed && seedMatchConfirmation && ((passwordSet && passwordMatchConfirmation) || !passwordNeeded) && (!unconventionalSeed || continueWithUnconventionalSeed)
    } // function updateAcceptButtonStatus()

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
                nextTabItem: columnLayoutSeedWarning.warn ? checkBoxContinueWithSeedWarning : checkBoxEncryptWallet

                onDataModified: {
                    updateAcceptButtonStatus()
                }
            }

            ColumnLayout {
                id: columnLayoutSeedWarning

                property bool warn: false

                Layout.fillWidth: true
                Layout.preferredHeight: warn ? implicitHeight : 0
                Behavior on Layout.preferredHeight { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } }
                opacity: warn ? 1 : 0
                Behavior on opacity { NumberAnimation {} }
                clip: true
                
                Material.foreground: Material.Pink
                Material.accent: Material.Pink

                Label {
                    Layout.fillWidth: true
                    text: qsTr("Warning")
                    font.pointSize: 14
                    font.bold: true
                    wrapMode: Text.WordWrap
                }
                Label {
                    Layout.fillWidth: true
                    text: qsTr("You introduced an unconventional seed. "
                             + "If you did it for any special reason, "
                             + "you can continue (only recommended for advanced users). "
                             + "However, if your intention is to use a normal system seed, "
                             + "you must delete all the additional text and special characters.")
                    wrapMode: Text.WordWrap
                }
                CheckBox {
                    id: checkBoxContinueWithSeedWarning
                    text: qsTr("Continue with the unconventional seed")

                    onCheckedChanged: {
                        updateAcceptButtonStatus()
                    }
                }
            } // ColumnLayoutSeedWarning

            CheckBox {
                id: checkBoxEncryptWallet
                text: qsTr("Encrypt wallet")
                checked: true

                onCheckedChanged: {
                    updateAcceptButtonStatus()
                    if (checked && visible) {
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
                Layout.fillWidth: true
                spacing: 10
                enabled: checkBoxEncryptWallet.checked

                TextField {
                    id: textFieldPassword
                    Layout.fillWidth: true
                    
                    placeholderText: qsTr("Password")
                    echoMode: TextField.Password
                    selectByMouse: true

                    onTextChanged: {
                        updateAcceptButtonStatus()
                    }
                }
                TextField {
                    id: textFieldConfirmPassword
                    Layout.fillWidth: true
                    
                    placeholderText: qsTr("Confirm password")
                    echoMode: TextField.Password
                    selectByMouse: true

                    onTextChanged: {
                        updateAcceptButtonStatus()
                    }
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