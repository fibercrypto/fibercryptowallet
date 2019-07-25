import QtQuick 2.12
import QtQuick.Controls 2.12

// Resource imports
// import "qrc:/ui/src/ui/"

Item {
    id: root

    property bool toolPageOpened: false
    property alias depth: stackView.depth
    property alias busy: stackView.busy

    function openBlockchainPage() {
        stackView.push(componentBlockchain)
        toolPageOpened = true
    }

    function openNetworkingPage() {
        stackView.push(componentNetworking)
        toolPageOpened = true
    }

    function pop() {
        stackView.pop()
    }

    StackView {
        id: stackView
        initialItem: componentCreateLoadWallet
        anchors.fill: parent
    }

    Component {
        id: componentCreateLoadWallet

        Rectangle {
            CreateLoadWallet {
                id: createLoadWallet
                anchors.centerIn: parent

                onWalletCreationRequested: {
                    stackView.push(componentGeneralSwipeView)
                }
            }
        }
    }

    Component {
        id: componentGeneralSwipeView

        GeneralSwipeView {
            id: generalSwipeView
        }
    }

    Component {
        id: componentBlockchain

        Blockchain {
            id: blockchain
        }
    }

    Component {
        id: componentNetworking

        Networking {
            id: networking
        }
    }
}
