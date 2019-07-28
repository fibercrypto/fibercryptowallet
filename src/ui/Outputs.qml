import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "Delegates/" // For quick UI development, switch back to resources when making a release

Page {
    id: outputs

    Frame {
        anchors.fill: parent
        anchors.margins: 20
        clip: true

        ListView {
            id: listOutputs
            anchors.fill: parent
            clip: true

            model: modelOutputs
            delegate: OutputsListDelegate {
            //     modelIp: ip
            //     modelPort: port
            //     modelSource: source
            //     modelBlock: block
            //     modelLastSeenIn: lastSeenIn
            //     modelLastSeenOut: lastSeenOut
            }
        }
    } // Frame

    // Roles: name
    // Implement the model in the backend (a more recommendable approach)
    ListModel { // EXAMPLE
        id: modelOutputs
        ListElement { name: "Wallet A" }
        ListElement { name: "Wallet B" }
        ListElement { name: "Wallet C" }
    }
}