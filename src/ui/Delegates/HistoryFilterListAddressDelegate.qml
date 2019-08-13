import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12

CheckDelegate {
    id: root
    //Component.onCompleted:{
    //    if(filterDelegate.myList.checkedAddresses.length < (index+1)){
    //        filterDelegate.myList.checkedAddresses.push(checked)
    //    }
    //    else{
    //        console.log(filterDelegate.myList.checkedAddresses[index])
    //        checked = filterDelegate.myList.checkedAddresses[index]
    //    }
    //}
    text: address // a role of the model
    font.family: "Code New Roman"

    LayoutMirroring.enabled: true
    contentItem: Label {
        leftPadding: root.indicator.width + root.spacing
        text: root.text
        verticalAlignment: Qt.AlignVCenter
        color: root.enabled ? root.Material.foreground : root.Material.hintTextColor
    }
}
