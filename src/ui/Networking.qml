import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import NetworkingManager 1.0
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

    NetworkingModel {
        id: modelNetworking
    }

    NetworkManager {
        id: networkManager
        property Timer timer: Timer{
            id: networkTimer
            repeat: true
            running: true
            interval: 3000
            onTriggered: {
                modelNetworking.cleanNetworks()
                modelNetworking.addMultipleNetworks(networkManager.getNetworks())
            }

        }
    }

    BusyIndicator {
        id: busyIndicator

        anchors.centerIn: parent
        // Create a `busy` property in the backend and bind it to `running` here:
        running: modelNetworking.networks.length === 0

    }
}
