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
        initialItem: componentPageCreateLoadWallet
        anchors.fill: parent
    }

    Component {
        id: componentPageCreateLoadWallet

        Item {
            PageCreateLoadWallet {
                id: pageCreateLoadWallet
                anchors.centerIn: parent
                width: 400
                height: 400

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
