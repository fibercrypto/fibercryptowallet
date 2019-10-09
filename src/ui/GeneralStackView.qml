import QtQuick 2.12
import QtQuick.Controls 2.12
import BlockchainModels 1.0
import WalletsManager 1.0

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


    WalletManager {
        id: walletManager
    }

    WalletModel {
        id: walletModel

        Component.onCompleted: {
            walletModel.loadModel(walletManager.getWallets())
            if (walletModel.count) {
                if (stackView.depth > 1) {
                    stackView.replace(componentGeneralSwipeView)
                } else {
                    stackView.push(componentGeneralSwipeView)
                }
            }
        }
    }

    BlockchainStatusModel {
        id: blockchainModel
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

    function openSettingsPage() {
        if (stackView.depth > 1) {
            stackView.replace(componentSettings)
        } else {
            stackView.push(componentSettings)
        }
    }

    function pop() {
        stackView.pop()
    }

    
    StackView {
        id: stackView
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

            Component.onCompleted: {
                model = blockchainModel
                model.update()
            }
        }
    }

    Component {
        id: componentNetworking

        Networking {
            id: networking
        }
    }

    Component {
        id: componentSettings

        Settings {
            id: settings
        }
    }

    
}
