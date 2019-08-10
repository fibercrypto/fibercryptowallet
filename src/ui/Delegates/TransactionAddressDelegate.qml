import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12


import "../"
Item {
    id: root
    
    implicitHeight: 30
    
        RowLayout {
            id: rowLayoutSent
        
        ToolButtonQR {
            id: toolButtonQRSent
            iconSize: "24x24"
            onClicked: {
                qrCodeRequested(address)
            }
        }
        Label {
              text: address // model's role
              font.family: "Code New Roman"
              Layout.fillWidth: true
            }
        }
}



