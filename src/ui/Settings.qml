import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "Delegates/" // For quick UI development, switch back to resources when making a release
import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/Delegates"
import "Delegates/" // For quick UI development, switch back to resources when making a release

Page {
    id: root


    Frame {
        anchors.fill: parent
        anchors.margins: 20
        clip: true


            ScrollView {
                anchors.fill: parent
                id: scrollSettings
                ScrollBar.horizontal.policy: ScrollBar.AlwaysOff
                ColumnLayout{
                    id: columnSettings
                    anchors.fill: parent
                    Button {
                        text: "Apply"
                        Layout.alignment: Qt.AlignRight
                        anchors.margins: 20
                    }
                    SettingsDelegate {
                        id: settingsDelegate
                        width: scrollSettings.width
                        modelNodeDirection: "node"
                        modelIsWalletEnvLocal: true
                        modelWalletPath: "path"
                    }
                        Connections {
                            target: scrollSettings
                            onWidthChanged: {
                               columnSettings.width = scrollSettings.width
                               console.log("settingsDelegate.width changed => " + settingsDelegate.width)
                            }
                        }
                        Component.onCompleted: {
                               columnSettings.width = scrollSettings.width

                        }
                } // ColumnLayout
            } // ScrollView
    } // Frame
} // Page