import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

ItemDelegate {
    id: root

    property string modelIp: "0.0.0.0"
    property int modelPort: 0
    property string modelSource: qsTr("Default peer")
    property int modelBlock: 0
    property string modelLastSeenIn
    property string modelLastSeenOut
    

    width: parent.width
    height: columnLayoutLastSeen.height + rowLayoutRoot.anchors.topMargin + rowLayoutRoot.anchors.bottomMargin

    RowLayout {
        id: rowLayoutRoot
        anchors.fill: parent
        anchors.leftMargin: 20
        anchors.rightMargin: 20
        anchors.topMargin: 10
        anchors.bottomMargin: 12

        spacing: 20

        Image {
            source: "qrc:/images/resources/images/icons/send-blue.svg"
            sourceSize: "32x32"
            fillMode: Image.PreserveAspectFit
            Layout.alignment: Qt.AlignLeft | Qt.AlignVCenter
        }

        Label {
            Material.foreground: Material.Grey
            text: modelIp + ':' + modelPort // model's roles
            Layout.fillWidth: true
            Layout.minimumWidth: 160
        }

        Label {
            text: modelSource // model's role
            Layout.preferredWidth: 100
        }

        Label {
            text: modelBlock // model's role
            Layout.preferredWidth: 80
        }

        ColumnLayout {
            id: columnLayoutLastSeen
            Layout.preferredWidth: 160
            spacing: 0

            RowLayout {
                Image {
                    source: "qrc:/images/resources/images/icons/up.svg"
                    sourceSize: Qt.size(labelLastSeenOut.font.pixelSize, labelLastSeenOut.font.pixelSize)
                    fillMode: Image.PreserveAspectFit
                }
                Label {
                    id: labelLastSeenOut
                    text: modelLastSeenOut +" seconds ago" // model's role
                }
                Layout.alignment: Qt.AlignLeft | Qt.AlignVCenter
            }

            RowLayout {
                Image {
                    source: "qrc:/images/resources/images/icons/down.svg"
                    sourceSize: Qt.size(labelLastSeenIn.font.pixelSize, labelLastSeenIn.font.pixelSize)
                    fillMode: Image.PreserveAspectFit
                }
                Label {
                    id: labelLastSeenIn
                    text: modelLastSeenIn +" seconds ago"// model's role
                }
                Layout.alignment: Qt.AlignLeft | Qt.AlignVCenter
            }
        } // ColumnLayout

    } // RowLayout (root)
}
