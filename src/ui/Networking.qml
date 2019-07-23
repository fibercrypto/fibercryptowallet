import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/Dialogs"
// import "qrc:/ui/src/ui/Delegates"
import "Dialogs/" // For quick UI development, switch back to resources when making a release
import "Delegates/" // For quick UI development, switch back to resources when making a release

Page {
    id: root

    header: RowLayout {
        spacing: 20
        ToolButton {
            id: toolButtonBack
            icon.source: "qrc:/images/resources/images/icons/back.svg"
            Layout.alignment: Qt.AlignLeft

            onClicked: {
                backRequested()
            }
        }

        Label {
            text: qsTr("Networking")
            font.pointSize: Qt.application.font.pointSize * 2
            horizontalAlignment: Text.AlignHCenter
            leftPadding: -(toolButtonBack.width + parent.spacing)
            Layout.fillWidth: true
        }
    }

    Frame {
        anchors.fill: parent
        anchors.margins: 20
        clip: true

        ListView {
            id: listNetworking
            anchors.fill: parent
            clip: true

            model: modelNetworking
            delegate: NetworkingListDelegate {
            }
        }
    } // Frame

    // Roles: ip, port, source, block, lastSeenIn, lastSeenOut
    // Implement the model in the backend (a more recommendable approach)
    ListModel { // EXAMPLE
        id: modelNetworking
        ListElement { ip: "0.0.0.0."; port: "0"; source: "Default peer"; block: 0; lastSeenIn: ""; lastSeenOut: "" }
    }
}