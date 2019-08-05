import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "Delegates/" // For quick UI development, switch back to resources when making a release

Page {
    id: root

    Frame {
        anchors.fill: parent
        anchors.margins: 20
        clip: true

        ColumnLayout {
            id: columnLayoutFrame
            anchors.fill: parent

            ColumnLayout {
                id: columnLayoutHeader

                RowLayout {
                    Layout.topMargin: 30

                    Label {
                        text: qsTr("IP address and port")
                        font.pointSize: 9
                        Layout.leftMargin: 75 // Empirically obtained
                        Layout.fillWidth: true
                    }
                    Label {
                        text: qsTr("Source")
                        font.pointSize: 9
                        Layout.rightMargin: 75
                    }
                    Label {
                        text: qsTr("Block")
                        font.pointSize: 9
                        Layout.rightMargin: 65
                    }
                    Label {
                        text: qsTr("Last seen")
                        font.pointSize: 9
                        Layout.rightMargin: 90
                    }
                }

                Rectangle {
                    Layout.fillWidth: true
                    height: 1
                    color: "#DDDDDD"
                }
            } // ColumnLayout (header)

            ScrollView {
                id: scrollItem
                Layout.fillWidth: true
                Layout.fillHeight: true
                ScrollBar.horizontal.policy: ScrollBar.AlwaysOff

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
                } // ListView
            } // ScrollView
        } // ColumnLayout (frame)
    } // Frame

    NetworkingManager{
        id: networkingManager
    }


    // Roles: ip, port, source, block, lastSeenIn, lastSeenOut
    // Implement the model in the backend (a more recommendable approach)
    ListModel { // EXAMPLE
        id: modelNetworking
        ListElement { ip: "192.168.137.1"; port: "8080"; source: qsTr("Default peer"); block: 17700; lastSeenIn: "a few seconds ago"; lastSeenOut: "two minutes ago" }
        ListElement { ip: "255.255.255.255"; port: "65535"; source: qsTr("Default peer"); block: 60978432; lastSeenIn: "a few seconds ago"; lastSeenOut: "a few seconds ago" }
        ListElement { ip: "192.168.137.3"; port: "5"; source: qsTr("Default peer"); block: 500; lastSeenIn: "4 days ago"; lastSeenOut: "one day ago" }
    }
}