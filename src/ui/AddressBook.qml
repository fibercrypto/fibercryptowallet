import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import AddrsBookManager 1.0

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
// import "qrc:/ui/src/ui/Dialogs"
import "Delegates/" // For quick UI development, switch back to resources when making a release
import "Dialogs/" // For quick UI development, switch back to resources when making a release

Page {
    id: addressBook

    enum SecurityType { LowSecurity, MediumSecurity, StrongSecurity }

    signal canceled()

//    property int currentContact: -1
    property int secType
    property string password

    Component.onCompleted: {
        if (abm.hasInit()) {
            if (abm.getSecType() !== AddressBook.SecurityType.StrongSecurity) {
                abm.loadContacts()
            } else {
                getpass.open()
            }
        } else {
            dialogCreateAddrsBk.open()
        }
    }

    ScrollView {
        anchors.fill: parent
        clip: true

        ListView {
            id: addrsBook

            model: abm
            delegate: ContactDelegate {
                id: contactDelegate
                width: parent.width

                MouseArea {
                    anchors.fill: parent
                    acceptedButtons: Qt.RightButton

                    onClicked: {
                        menu.index = index
                        menu.cId = id
                        menu.name = name
                        menu.address = address
                        menu.popup()
                    }
                }
            }

            section.property: "name"
            section.criteria: ViewSection.FirstCharacter
            section.delegate: SectionDelegate {
                width: addrsBook.width
            }
        }
    } // ScrollView

    RoundButton {
        anchors.margins: 10
        anchors.right: parent.right
        anchors.bottom: parent.bottom
        icon.source: "qrc:/images/resources/images/icons/add.svg"
        highlighted: true
        onClicked: {
            contactDialog.isEdit = false
            contactDialog.open()
        }
    } //RoundButton

    AddrsBookModel{
        id: abm
    }

    Menu {
        id: menu

        property int index: -1
        property int cId: -1
        property string name
        property AddrsBkAddressModel address

        MenuItem {
            text: qsTr("&View")
            onTriggered: {
                console.log(menu.name)
                dialogShowContact.open()
            }
        }

        MenuSeparator {}

        MenuItem {
            text: qsTr("&Edit")
            onTriggered: {
                contactDialog.isEdit = true
                contactDialog.open()
            }
        }

        MenuSeparator {}

        MenuItem {
            text: qsTr("&Remove")
            onTriggered: {
                dialogConfirmation.index = menu.index
                dialogConfirmation.cId   = menu.cId
                dialogConfirmation.name  = menu.name
                dialogConfirmation.open()
            }
        }
    } // Menu

    DialogAddContact {
        id: contactDialog

        anchors.centerIn: Overlay.overlay
        width:  applicationWindow.width  > 540 ? 540 : applicationWindow.width
        height: applicationWindow.height > 640 ? 640 : applicationWindow.height

        focus: visible
        modal: true
    }

    DialogSelectSecType {
        id: dialogCreateAddrsBk

        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 480 ? 480 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > 400 ? 400 - 40 : applicationWindow.height - 40

        title: qsTr("Initial setup")
        focus: visible
        modal: true

        onAccepted: {
            secType = select
            if (secType === AddressBook.SecurityType.StrongSecurity) {
                setpass.open()
            } else {
                abm.initAddrsBook(secType, "")
            }
        }

        onRejected: {
            canceled()
        }
    }

    DialogSetPassword {
        id: setpass

        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 400 ? 400 - 40 : applicationWindow.width - 40

        focus: visible
        modal: true

        onAccepted: {
            abm.initAddrsBook(AddressBook.SecurityType.StrongSecurity, setpass.password)
        }

        onRejected: {
            canceled()
        }
    }

    DialogGetPassword {
        id: getpass

        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 400 ? 400 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > 280 ? 280 - 40 : applicationWindow.height - 40

        modal: true
        focus: visible

        onAccepted: {
            if (!abm.authenticate(getpass.password)) {
                getpass.open()
            } else {
                abm.loadContacts()
            }
        }

        onRejected:{
            canceled()
        }
    }

    DialogShowContact {
        id: dialogShowContact

        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 400 ? 400 - 40 : applicationWindow.width - 40
        height: applicationWindow.height > 440 ? 440 - 40 : applicationWindow.height - 40

        modal: true
        focus: visible
    }

    Dialog {
        id: dialogConfirmation

        property int index: -1
        property int cId: -1
        property string name

        anchors.centerIn: Overlay.overlay
        width: applicationWindow.width > 300 ? 300 - 40 : applicationWindow.width - 40

        standardButtons: Dialog.Ok | Dialog.Cancel
        title: qsTr("Confirm action")
        modal: true
        focus: visible

        RowLayout {
            width: parent.width
            spacing: 30

            Image {
                id: icon
                source: "qrc:/images/resources/images/icons/trash.svg"
                sourceSize: "64x64"
                Layout.alignment: Qt.AlignLeft | Qt.AlignTop
            }

            Label {
                id: labelQuestion

                Layout.fillWidth: true
                text: qsTr("Delete contact <b>%1</b>?").arg(dialogConfirmation.name)
                wrapMode: Text.Wrap
                font.italic: true
            }
        }

        Component.onCompleted: {
            standardButton(Dialog.Ok).Material.accent = Material.Red
            standardButton(Dialog.Ok).highlighted = true
            standardButton(Dialog.Ok).text = qsTr("Delete")
        }

        onAccepted: {
            abm.removeContact(dialogConfirmation.index, dialogConfirmation.cId)
        }
    } // Dialog
}
