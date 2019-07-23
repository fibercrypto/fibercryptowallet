import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "Delegates/" // For quick UI development, switch back to resources when making a release

Page {
    id: root

    signal backRequested()

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
                modelIp: ip
                modelPort: port
                modelSource: source
                modelBlock: block
                modelLastSeenIn: lastSeenIn
                modelLastSeenOut: lastSeenOut
            }
        }
    } // Frame

    // Roles: ip, port, source, block, lastSeenIn, lastSeenOut
    // Implement the model in the backend (a more recommendable approach)
    ListModel { // EXAMPLE
        id: modelNetworking
        ListElement { ip: "192.168.137.1"; port: "8080"; source: qsTr("Default peer"); block: 17700; lastSeenIn: "a few seconds ago"; lastSeenOut: "two minutes ago" }
        ListElement { ip: "255.255.255.255"; port: "65535"; source: qsTr("Default peer"); block: 60978432; lastSeenIn: "a few seconds ago"; lastSeenOut: "a few seconds ago" }
        ListElement { ip: "192.168.137.3"; port: "5"; source: qsTr("Default peer"); block: 500; lastSeenIn: "4 days ago"; lastSeenOut: "one day ago" }
    }
}