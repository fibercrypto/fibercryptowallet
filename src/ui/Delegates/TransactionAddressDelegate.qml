import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/"
import "../" // For quick UI development, switch back to resources when making a release
Item {
    id: root
    
    implicitHeight: 30
    
        RowLayout {
            id: rowLayoutSent
        
        ToolButtonQR {
            
            id: toolButtonQRSent
            Layout.alignment: Qt.AlignLeft
            iconSize: "24x24"
            onClicked: {
                qrCodeRequested(address)
            }
        }
        Label {
            
              text: address // model's role
              Layout.alignment: Qt.AlignLeft
              font.family: "Code New Roman"
              Layout.fillWidth: true
            }
        }
}



