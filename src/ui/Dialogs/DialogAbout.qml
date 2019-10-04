import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12

Dialog {
    id: about

    modal: true
    title: qsTr("About") + " " + Qt.application.name
    standardButtons: Dialog.Close

    signal licenseRequested()

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
                source: "qrc:/images/resources/images/icons/appIcon.png"
            }

            Label {
                width: parent.width
                text: "<b>" + Qt.application.name + ' v' + Qt.application.version + "</b>"
                    + qsTr("<p>This program is for [...???...].</p><br>")

                    + qsTr("<p><b>License terms and disclaimer</b><br>"
                    + "This program is free software; you can redistribute it and/or modify "
                    + "it under the terms of the <a href=\"License\">GNU General Public License</a> as published by "
                    + "the Free Software Foundation; either version 3 of the License, or "
                    + "(at your option) any later version.<br>"
                    + "This program is distributed in the hope that it will be useful, "
                    + "but WITHOUT ANY WARRANTY; without even the implied warranty of "
                    + "MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the "
                    + "<a href=\"License\">GNU General Public License</a> for more details.<br>"
                    + "You should have received a copy of the <a href=\"License\">GNU General Public License</a> along "
                    + "with this program; if not, see <a href=\"%1\">%1</a>.</p><br>").arg("http://www.gnu.org/licenses/")

                    + qsTr("<b>Contact information:</b> <a href=\"mailto:%1\">%1</a>").arg("simelo@gmail.com")
                wrapMode: Label.Wrap
                onHoveredLinkChanged: {
                    mouseAreaLinkHovered.cursorShape = hoveredLink ? Qt.PointingHandCursor : Qt.ArrowCursor
                }
                onLinkActivated: {
                    if (link === "License") {
                        licenseRequested()
                    } else {
                        Qt.openUrlExternally(link)
                    }
                }

                MouseArea {
                    id: mouseAreaLinkHovered
                    anchors.fill: parent
                    acceptedButtons: Qt.NoButton
                }
            } // Label
        } // Column

        ScrollIndicator.vertical: ScrollIndicator {
            parent: about.contentItem
            anchors.top: flickable.top
            anchors.bottom: flickable.bottom
            anchors.right: parent.right
            anchors.rightMargin: -about.rightPadding + 1
        }
    } // Flickable
}
