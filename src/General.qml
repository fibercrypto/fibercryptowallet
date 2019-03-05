import QtQuick 2.12
import QtQuick.Controls 2.12

Page {
    id: root
    header: TabBar {
        id: tabBar
        currentIndex: 2

        TabButton {
            text: qsTr("Wallets")
        }
        TabButton {
            text: qsTr("Send")
        }
        TabButton {
            text: qsTr("History")
        }
    }

    SwipeView {
        id: swipeView
        anchors.fill: parent
        currentIndex: tabBar.currentIndex

        PageWallets {
        }

        PageSend {
        }

        PageHistory {
        }
    }
}
