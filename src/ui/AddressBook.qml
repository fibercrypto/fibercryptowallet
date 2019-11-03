import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import AddrsBookModels 1.0
// Resource imports
// import "qrc:/ui/src/ui/Dialogs"
import "Dialogs/" // For quick UI development, switch back to resources when making a release

Page{
    id:addressBook
    GroupBox{
     anchors.fill: parent
            anchors.margins: 20
            clip: true
    }

       ScrollView {
                anchors.fill: parent
                clip: true
                ListView {
                    id: listTransactions

                    model: AddrsBookManager{}

                }
        }
    } // GroupBox
