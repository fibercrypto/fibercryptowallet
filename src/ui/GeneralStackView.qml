import QtQuick 2.12
import QtQuick.Controls 2.12

// Resource imports
// import "qrc:/ui/src/ui/"

Item {
    id: root

    property alias depth: stackView.depth
    property alias busy: stackView.busy

    function openBlockchainPage() {
        if (stackView.depth > 1) {
            stackView.replace(componentBlockchain)
        } else {
            stackView.push(componentBlockchain)
        }
    }

    function openNetworkingPage() {
        if (stackView.depth > 1) {
            stackView.replace(componentNetworking)
        } else {
            stackView.push(componentNetworking)
        }
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
                    stackView.replace(componentGeneralSwipeView)
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
