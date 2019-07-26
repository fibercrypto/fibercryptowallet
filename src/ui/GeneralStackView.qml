import QtQuick 2.12
import QtQuick.Controls 2.12
import BlockchainModels 1.0

// Resource imports
// import "qrc:/ui/src/ui/"

Item {
    id: root

    property bool toolPageOpened: false

    BlockchainStatusModel {
        id: blockchainModel
    }

    function openBlockchainPage() {
        stackView.push(componentBlockchain)
        toolPageOpened = true
    }

    function openNetworkingPage() {
        stackView.push(componentNetworking)
        toolPageOpened = true
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

            onBackRequested: {
                stackView.pop()
                toolPageOpened = false
            }

            Component.onCompleted: {
                blockchainModel.update()
                model = blockchainModel
            }

        }
    }

    Component {
        id: componentNetworking

        Networking {
            id: networking

            onBackRequested: {
                stackView.pop()
                toolPageOpened = false
            }
        }
    }
}
