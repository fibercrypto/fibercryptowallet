import QtQuick 2.12
import QtQuick.Controls 2.12

// Resource imports
// import "qrc:/ui/src/ui/"

Item {
    id: generalStackView

    property alias depth: stackView.depth
    property alias busy: stackView.busy

    function openOutputsPage() {
        if (stackView.depth > 1) {
            stackView.replace(componentOutputs)
        } else {
            stackView.push(componentOutputs)
        }
    }

    function openPendingTransactionsPage() {
        if (stackView.depth > 1) {
            stackView.replace(componentPendingTransactions)
        } else {
            stackView.push(componentPendingTransactions)
        }
    }

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
        id: componentOutputs

        Outputs {
            id: outputs
        }
    }

    Component {
        id: componentPendingTransactions

        PendingTransactions {
            id: pendingTransactions
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
