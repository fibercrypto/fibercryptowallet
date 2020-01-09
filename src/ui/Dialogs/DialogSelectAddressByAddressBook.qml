import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12

// Resource imports
// import "qrc:/ui/src/ui/Controls"
import "../Controls" // For quick UI development, switch back to resources when making a release
import "../Delegates"
Dialog {

    id: dialogSelectAddressByAddressBook
    property ListModel listAddrsModel
    property string selectedAddress
    padding: 0
    standardButtons: Dialog.Cancel

    closePolicy: Dialog.CloseOnPressOutside
    onAboutToShow: {
        textFieldFilterContact.forceActiveFocus()
}

function filterByText(text){
myDisplayAddrsModel.clear()
if (text===""){
for(var i=0;i<listAddrsModel.count;i++){
myDisplayAddrsModel.append(listAddrsModel.get(i))
}
return
}
for(var i=0;i<listAddrsModel.count;i++){
if(listAddrsModel.get(i).name.toLowerCase().includes(text.toLowerCase())||listAddrsModel.get(i).address.toLowerCase().includes(text.toLowerCase())){
myDisplayAddrsModel.append(listAddrsModel.get(i))
}
}

}

    ColumnLayout {
        anchors.fill: parent

        TextField {
            id: textFieldFilterContact

            Layout.fillWidth: true
            Layout.leftMargin: 10
            Layout.rightMargin: 10
            placeholderText: qsTr("Filter address")
            focus: true
            selectByMouse: true
        onTextChanged:{
filterByText(textFieldFilterContact.text)
console.log(myDisplayAddrsModel.count)
        }
        }

        ListView {
            id: listView
            model: myDisplayAddrsModel
            Layout.fillWidth: true
            Layout.fillHeight: true
            clip: true
            currentIndex: -1
              section.property: "name"
              section.criteria: ViewSection.FullString
              section.delegate: SectionDelegate {
              width: parent.width
               }
            delegate: ItemDelegate {
                width: parent.width
                Behavior on height { NumberAnimation { easing.type: Easing.OutQuint } }
               focusPolicy: Qt.NoFocus
               text:coinType+"-"+address
                font.family: "Code New Roman"
                Material.foreground: hovered ? parent.Material.accent : parent.Material.foreground
                highlighted: hovered
                leftPadding: highlighted ? 2*padding : padding // added
                Behavior on leftPadding { NumberAnimation { duration: 500; easing.type: Easing.OutQuint } } // added
            
                onHighlightedChanged: {
                    ListView.view.currentIndex = highlighted ? index : -1
                }

                onClicked: {
                    selectedAddress = address
                    accept()
                }
            }


ListModel{
id:myDisplayAddrsModel
}

            ScrollIndicator.vertical: ScrollIndicator { }
        } // ListView
    } // ColumnLayout

}