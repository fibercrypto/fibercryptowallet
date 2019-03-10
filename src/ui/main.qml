import QtQuick 2.12
import QtQuick.Controls 2.12

ApplicationWindow {
    visible: true
    width: 680
    height: 580
    title: qsTr("FiberCrypto Wallet")

    GeneralSwipeView {
        id: general
        anchors.fill: parent
    }
}
