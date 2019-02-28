import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

Item {
    id: root
    implicitHeight: 80

    ColumnLayout {
        id: columnLayoutRoot
        anchors.fill: parent

        RowLayout {
            id: rowLayoutHeader
            spacing: 10
            Layout.alignment: Qt.AlignTop
            Layout.fillWidth: true

            Label {
                id: labelIndex
                text: index + 1
                font.pointSize: Qt.application.font.pointSize * 0.9
                font.bold: true
            }
            Label {
                id: labelAddress
                text: address
                font.family: "Code New Roman"
                font.pointSize: Qt.application.font.pointSize * 0.9
                font.bold: true
                Layout.fillWidth: true
            }
        }

        GridLayout {
            columns: 2
            Layout.alignment: Qt.AlignTop
            Layout.leftMargin: labelIndex.width + rowLayoutHeader.spacing
            Layout.fillWidth: true

            Label {
                text: qsTr("Coins:")
                font.pointSize: Qt.application.font.pointSize * 0.9
                font.bold: true
            }
            Label {
                text: addressSky
                font.pointSize: Qt.application.font.pointSize * 0.9
            }

            Label {
                text: qsTr("Hours:")
                font.pointSize: Qt.application.font.pointSize * 0.9
                font.bold: true
            }
            Label {
                text: addressCoinHours
                font.pointSize: Qt.application.font.pointSize * 0.9
            }
        }
    } // ColumnLayout
}
