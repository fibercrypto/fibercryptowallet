import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Window 2.12
import QtQuick.Layouts 1.12
import Qt.labs.settings 1.0

Window {
    id: windowSplash

    width: dialogSplash.width + 100
    height: dialogSplash.height + 140
    visible: true
    flags: Qt.SplashScreen | Qt.BypassWindowManagerHint | Qt.CustomizeWindowHint | Qt.WindowStaysOnTopHint
    color: "transparent"
    x: (Screen.width - width)/2
    y: (Screen.height - height)/2

    Material.theme: ~~settings.value("style/material/theme", Material.Light)

    Dialog {
        id: dialogSplash
        anchors.centerIn: Overlay.overlay
        visible: true
        closePolicy: Dialog.NoAutoClose
        standardButtons: Dialog.Abort

        Component.onCompleted: {
            standardButton(Dialog.Abort).Material.accent = Material.Red
            standardButton(Dialog.Abort).highlighted = true
        }

        onRejected: Qt.exit(-1)
        onClosed: windowSplash.visible = false

        width: implicitWidth + (implicitWidth % 2)
        height: implicitHeight + (implicitHeight % 2)

        ColumnLayout {
            anchors.fill: parent
            
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
                font.family: "Hemi Head"
                font.pointSize: Qt.application.font.pointSize * 2.2
            }

            Label {
                id: labelApplicationDescription
                Layout.alignment: Qt.AlignVCenter | Qt.AlignHCenter
                text: qsTr("Multi-coin cryptocurrency wallet")
                font.bold: true
            }

            RowLayout {
                clip: true
                Layout.alignment: Qt.AlignBottom | Qt.AlignHCenter

                BusyIndicator {
                    running: visible
                    implicitWidth: 32
                    implicitHeight: implicitWidth
                }
                Label {
                    text: qsTr("Loading...")
                    font.italic: true
                }
            }
        } // ColumnLayout
    } // Dialog

    Loader {
        id: loader
        source: "main.qml"
        asynchronous: true

        onLoaded: {
            dialogSplash.standardButton(Dialog.Abort).enabled = false
            item.visible = true
            timer.start()
        }
    }

    Timer {
        id: timer
        interval: 200
        repeat: false
        onTriggered: dialogSplash.close()
    }

    Settings {
        id: settings
    }
}
