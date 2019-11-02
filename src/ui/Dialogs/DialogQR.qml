import QtQuick 2.12
import QtQuick.Layouts 1.12
import QtQuick.Controls 2.12

// Resource imports
// import "qrc:/utils/src/ui/Utils"
import "../Utils" // For quick UI development, switch back to resources when making a release

Dialog {
    id: dialogQR

    title: Qt.application.name + " - QR"
    standardButtons: Dialog.Close

    signal setQRVars(var data)

    Component.onCompleted: {
        dialogQR.setQRVars.connect(setVars)
    }

    function setVars(data) {
        qrDisplay.value = data;
        textDisplay.text = data;
    }

    QRCode {
        id: qrDisplay
        anchors.top : parent.top
        anchors.topMargin : 30
        anchors.horizontalCenter : parent.horizontalCenter
        width : 320
        height : 320
        level : "H"
    }
    TextInput {
        id : textDisplay
        anchors.bottom : parent.bottom
        anchors.bottomMargin : 30
        anchors.horizontalCenter : parent.horizontalCenter
        onAccepted: {
            console.log(dialogQR.text)
            textDisplay.text = dialogQR.text
        }
    }
}