import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12


import "../"
ItemDelegate {
    id: root
    implicitHeight: 80

    
                RowLayout {
                    id: rowLayoutSent
                    visible: modelType === TransactionDetails.Type.Send

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


