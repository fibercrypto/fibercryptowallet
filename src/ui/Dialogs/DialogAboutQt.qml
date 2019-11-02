import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12

Dialog {
    id: aboutQt

    readonly property string qtVersionString: "5.13.1"

    modal: true
    title: qsTr("About") + " Qt"
    standardButtons: Dialog.Close

    Flickable {
        id: flickable
        clip: true
        anchors.fill: parent
        contentHeight: column.height

        Column {
            id: column
            spacing: 20
            width: parent.width

            Image {
                id: logo
                width: parent.width / 2
                anchors.horizontalCenter: parent.horizontalCenter
                fillMode: Image.PreserveAspectFit
                source: "qrc:/images/resources/images/icons/qt_logo_green_rgb.svg"
            }

            Label {
                width: parent.width
                text: "<b>" + qsTr("This program uses Qt %1").arg(qtVersionString) + "</b><br>"
                    + qsTr("<p>Qt is a <i>C++ toolkit for cross-platform application " +
                           "development</i>.</p>" +
                           "<p>Qt provides single-source portability across all major desktop and mobile " +
                           "operating systems. It is also available for embedded Linux and other " +
                           "embedded operating systems.</p><br>" +
                           "<p>Qt offers both <i>commercial</i> and <i>opensource</i> licences. Please see <a href=\"http://%2/\">%2</a> " +
                           "for an overview of Qt licensing.</p><br>" +
                           "<p><i>Copyright © %1 The Qt Company Ltd</i> and other " +
                           "contributors. See <a href=\"http://%3/\">%3</a> for more information.</p>").arg("2019").arg("qt.io/licensing").arg("qt.io")
                wrapMode: Label.Wrap
                onHoveredLinkChanged: {
                    mouseAreaLinkHovered.cursorShape = hoveredLink ? Qt.PointingHandCursor : Qt.ArrowCursor
                }
                onLinkActivated: { Qt.openUrlExternally(link) }

                MouseArea {
                    id: mouseAreaLinkHovered
                    anchors.fill: parent
                    acceptedButtons: Qt.NoButton
                }
            } // Label
        } // Column

        ScrollIndicator.vertical: ScrollIndicator {
            parent: aboutQt.contentItem
            anchors.top: flickable.top
            anchors.bottom: flickable.bottom
            anchors.right: parent.right
            anchors.rightMargin: -aboutQt.rightPadding + 1
        }
    } // Flickable
}
