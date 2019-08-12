import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/"
import "../" // For quick UI development, switch back to resources when making a release

// Backend imports
import HistoryModels 1.0

ItemDelegate {
    id: root

    property date modelDate: date
    property int modelType: type
    property int modelStatus: status
    property var modelStatusString: [ qsTr("Confirmed"), qsTr("Pending"), qsTr("Preview") ]
    property string modelAmount: amount
    property string modelHoursReceived: hoursTraspassed
    property string modelHoursBurned: hoursBurned
    property string modelTransactionID: transactionID
    property QAddressList modelAddresses: addresses
    property QAddressList modelInputs: inputs
    property QAddressList modelOutputs: outputs
    readonly property real delegateHeight: 30
    
    signal qrCodeRequested(var data)

    implicitWidth: parent.width
    implicitHeight: (columnLayoutMainContent.height < 78 ? 78 : columnLayoutMainContent.height) + rowLayoutRoot.anchors.topMargin + rowLayoutRoot.anchors.bottomMargin

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
            mirror: modelType === TransactionDetails.Type.Receive
            Layout.alignment: Qt.AlignTop | Qt.AlignLeft
        }

        ColumnLayout {
            id: columnLayoutMainContent
            Layout.fillWidth: true
            Layout.alignment: Qt.AlignTop | Qt.AlignLeft
            
                
                
                
            
            RowLayout {
                
                
                Layout.alignment: Qt.AlignLeft
                spacing: 20
                
                Row{
                    Layout.fillWidth:true
                    spacing: 20
                Label {
                    
                    font.bold: true
                    text: (modelType == TransactionDetails.Type.Receive ? qsTr("Received") : (modelType == TransactionDetails.Type.Send ? qsTr("Sent") : qsTr("Internal"))) + " SKY"
                }

                Label {
                    
                    Material.foreground: Material.Grey
                    text: modelDate.toLocaleString("2000-01-01 00:00") // model's role
                    font.pointSize: Qt.application.font.pointSize * 0.9
                }
                }
            }
            ColumnLayout {
                Layout.fillWidth:true
                Layout.fillHeight:true
                Layout.leftMargin:10
                
                Layout.alignment: Qt.AlignLeft
                height: delegateHeight*(modelInputs.rowCount())
                ListView{
                    Layout.alignment: Qt.AlignLeft
                    implicitHeight: delegateHeight*(modelAddresses.rowCount())
                    height: parent.height
                    id: listViewAddresses
                    model: modelAddresses
                    delegate:TransactionAddressDelegate{}
                                     
                }
                
                   
            } // ColumnLayout (addresses)
        
        } // ColumnLayout (main content)

        Label {
            text: (modelType === TransactionDetails.Type.Receive ? "" : "-") + amount + " SKY" // model's role
            font.pointSize: Qt.application.font.pointSize * 1.25
            font.bold: true
            Layout.alignment: Qt.AlignTop | Qt.AlignRight
        }

    } // RowLayout (root)
}
