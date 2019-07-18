import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Dialog {
    id: root

    ColumnLayout {
        anchors.fill: parent
        spacing: 20

        Image {
            source: "qrc:/images/resources/images/appIcon.svg"
            sourceSize: Qt.size(120, 120)
            fillMode: Image.PreserveAspectFit

            Layout.alignment: Qt.AlignTop | Qt.AlignHCenter
        }

        Rectangle {
            height: 1
            color: Material.color(Material.Grey)
            Layout.fillWidth: true
        }

        ColumnLayout {
            Label {
                text: Qt.application.name + ' v' + Qt.application.version
                font.bold: true
            }
            Label {
                text: "Copyright © 2019 - Carlos E. Pérez"
                font.italic: true
            }
            Label {
                text: qsTr("This program allows scaning a network looking for " +
                           "available proxies to connect with.")
                wrapMode: Label.WordWrap

                Layout.preferredWidth: 220
            }
        } // ColumnLayout
    } // ColumnLayout
} // Dialog
