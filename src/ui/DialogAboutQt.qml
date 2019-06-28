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
            source: "qrc:/images/icons/qt_logo_green_rgb.svg"
            sourceSize: Qt.size(72, 72)
            fillMode: Image.PreserveAspectFit

            Layout.alignment: Qt.AlignTop | Qt.AlignHCenter
        }

        Rectangle {
            height: 1
            color: "gray"
            Layout.fillWidth: true
        }

        ColumnLayout {
            Label {
                text: qsTr("This program uses Qt %1").arg("5.12.3")
                font.bold: true
            }
            Label {
                text: qsTr("<p>Qt is a <i>C++ toolkit for cross-platform application " +
                           "development</i>.</p>" +
                           "<p>Qt provides single-source portability across all major desktop and mobile " +
                           "operating systems. It is also available for embedded Linux and other " +
                           "embedded operating systems.</p><br>" +
                           "<p>Qt offers both <i>commercial</i> and <i>opensource</i> licences. Please see <a href=\"http://%2/\">%2</a> " +
                           "for an overview of Qt licensing.</p><br>" +
                           "<p><i>Copyright © %1 The Qt Company Ltd</i> and other " +
                           "contributors. See <a href=\"http://%3/\">%3</a> for more information.</p><br>").arg("2018").arg("qt.io/licensing").arg("qt.io")
                wrapMode: Label.WordWrap

                Layout.preferredWidth: 300
                Layout.alignment: Qt.AlignHCenter
            }
        } // ColumnLayout
    } // ColumnLayout
} // Dialog
