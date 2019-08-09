import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import OutputsModels 1.0


// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "Delegates/" // For quick UI development, switch back to resources when making a release

Page {
    id: outputs
    
    readonly property real listOutputsLeftMargin: 20
    readonly property real listOutputsRightMargin: 50
    readonly property real listOutputsSpacing: 20
    readonly property real internalLabelsWidth: 60

    Frame {
        id: frame
        anchors.fill: parent
        anchors.margins: 20
        clip: true

        ColumnLayout {
            id: columnLayoutFrame
            anchors.fill: parent

            ColumnLayout {
                id: columnLayoutHeader

                RowLayout {
                    spacing: listOutputsSpacing
                    Layout.topMargin: 30
                    
                    Label {
                        text: qsTr("Output ID")
                        font.pointSize: 9
                        Layout.leftMargin: 38
                        Layout.fillWidth: true
                    }
                    Label {
                        text: qsTr("Sky")
                        font.pointSize: 9
                        horizontalAlignment: Text.AlignRight
                        Layout.preferredWidth: internalLabelsWidth/2
                    }
                    Label {
                        text: qsTr("Coin hours")
                        font.pointSize: 9
                        horizontalAlignment: Text.AlignRight
                        Layout.rightMargin: listOutputsRightMargin
                        Layout.preferredWidth: internalLabelsWidth
                    }
                }

                Rectangle {
                    Layout.fillWidth: true
                    height: 1
                    color: "#DDDDDD"
                }
            } // ColumnLayout (header)

            ScrollView {
                id: scrollView
                Layout.fillWidth: true
                Layout.fillHeight: true
                ScrollBar.horizontal.policy: ScrollBar.AlwaysOff
                
                ListView {
                    id: listViewWallets
                    anchors.fill: parent
                    clip: true

                    // model: modelWallets
                    delegate: OutputsListDelegate {
                        width: listViewWallets.width
                    }

                    Timer {
                        id: timer
                        interval: 10000
                        onTriggered: {
                            listViewWallets.model = modelWallets
                            console.log('Model changed!')
                        }
                    }

                    Component.onCompleted: {
                        timer.start()
                        console.log('Timer started (10s)')
                    }
                } // ListView
            } // ScrollView
        } // ColumnLayout (frame)
    } // Frame

    // Roles: name
    // Implement the model in the backend (a more recommendable approach)
    QWallets{
        id: modelWallets
    }

    // ListModel { // EXAMPLE
    //     id: modelWallets
    //     ListElement { name: "Wallet A" }
    //     ListElement { name: "Wallet B" }
    //     ListElement { name: "Wallet C" }
    // }
}