import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12


Item {
    id: root

    implicitWidth: 400
    implicitHeight: 400

    signal walletCreationRequested()
    signal walletLoadingRequested()
    property alias name: createLoadWallet.name
    property alias seed: createLoadWallet.seed
    property alias confirmedSeed: createLoadWallet.seedConfirm
    
    Column {
        anchors.fill: parent
        spacing: 30
        
        ControlCustomSwitch {
            id: switchNewLoadWallet
            width: parent.width
            height: 70

            leftText: qsTr("New wallet")
            rightText: qsTr("Load wallet")

            backgroundColor: Material.accent

            textColor: Material.accent

            onToggled: {
                createLoadWallet.mode = isInLeftSide ? CreateLoadWallet.Create : CreateLoadWallet.Load
            }
        }

        CreateLoadWallet {
            id: createLoadWallet
            width: parent.width
        }

        Button {
            id: buttonCreateLoadWallet
            anchors.horizontalCenter: parent.horizontalCenter
            width: 120
            height: 60
            font.bold: true
            font.pointSize: Qt.application.font.pointSize * 1.2
            Connections{
                target: createLoadWallet
                onModeChanged:{
                    if( createLoadWallet.mode === CreateLoadWallet.Load){
                        buttonCreateLoadWallet.enabled = true
                    }
                    
                }
                onSeedChanged:{
                    if (createLoadWallet.mode === CreateLoadWallet.Create){
                    
                        if (createLoadWallet.seed != createLoadWallet.seedConfirm){
                            buttonCreateLoadWallet.enabled = false
                        } else{
                            buttonCreateLoadWallet.enabled = true
                        }
                    }

                }
                onSeedConfirmChanged:{
                    if (createLoadWallet.mode === CreateLoadWallet.Create){
                         if (createLoadWallet.seed != createLoadWallet.seedConfirm){
                            buttonCreateLoadWallet.enabled = false
                        } else{
                            buttonCreateLoadWallet.enabled = true
                        }
                    }
                   

                }
            }
            text: createLoadWallet.mode === CreateLoadWallet.Create ? qsTr("Create") : qsTr("Load")
            highlighted: true
            
            onClicked: {
                if (createLoadWallet.mode === CreateLoadWallet.Create) {
                    walletCreationRequested()
                } else {
                    walletLoadingRequested()
                }
            }
        }
    } // Column
}
