import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Window 2.12
import QtQuick.Layouts 1.12

Window {
    id: windowSplash

    width: dialogSplash.width + 100
    height: dialogSplash.height + 100
    visible: true
    flags: Qt.SplashScreen
    color: "transparent"

    Dialog {
        id: dialogSplash
        anchors.centerIn: Overlay.overlay
        visible: true
        closePolicy: Dialog.NoAutoClose

        onClosed: windowSplash.visible = false

        contentItem: ColumnLayout {
            Image {
                id: imageLogo
                Layout.alignment: Qt.AlignVCenter | Qt.AlignHCenter
                fillMode: Image.PreserveAspectFit
                source: "qrc:/images/resources/images/icons/appIcon/appIcon.png"
                sourceSize: "128x128"
            }

            Label {
                id: labelApplicationName
                Layout.alignment: Qt.AlignVCenter | Qt.AlignHCenter
                text: "FiberCrypto Wallet"
                // font.family: "???"
                font.pointSize: Qt.application.font.pointSize * 2.25
            }

            Label {
                id: labelApplicationDescription
                Layout.alignment: Qt.AlignVCenter | Qt.AlignHCenter
                text: qsTr("Multi-coin cryptocurrency wallet")
                font.bold: true
            }

            RowLayout {
                clip: true
                Layout.alignment: Qt.AlignVCenter | Qt.AlignHCenter

                BusyIndicator {
                    running: visible
                    implicitWidth: 32
                    implicitHeight: implicitWidth
                }
                Label {
                    text: qsTr("Loading")
                    font.italic: true
                }
            }
        } // ColumnLayout
    } // Dialog

    Loader {
        id: loader
        source: "main.qml"
        asynchronous: true
        visible: status == Loader.Ready

        onLoaded: {
            dialogSplash.close()
        }
    }
}
