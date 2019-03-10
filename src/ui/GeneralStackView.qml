import QtQuick 2.12
import QtQuick.Controls 2.12

Item {
    id: root

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
}