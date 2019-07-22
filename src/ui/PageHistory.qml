import QtQuick 2.12
import QtQuick.Controls 2.12
import QtQuick.Controls.Material 2.12
import QtQuick.Layouts 1.12
import HistoryModels 1.0

// Resource imports
import "qrc:/ui/src/ui/Dialogs"
import "qrc:/ui/src/ui/Delegates"

Page {
    id: root

    GroupBox {
        anchors.fill: parent
        anchors.margins: 20
        clip: true

        label: RowLayout {
            SwitchDelegate {
                id: switchFilters
                text: qsTr("Filters")
            }
            Button {
                id: buttonFilters
                flat: true
                enabled: switchFilters.checked
                highlighted: true
                text: qsTr("Select filters")

                onClicked: {
                    toolTipFilters.open()
                }
            }
        } // RowLayout (GroupBox label)

        ScrollView {
            anchors.fill: parent
            clip: true
            ListView {
                id: listTransactions

                model: modelTransactions
                delegate: HistoryListDelegate {
                    onClicked: {
                        dialogTransactionDetails.open()
                        listTransactions.currentIndex = index
                    }
                }
            }
        }
    } // GroupBox


    ToolTip {
        id: toolTipFilters

        anchors.centerIn: Overlay.overlay
        height: 16 + (filter.count > 5 ? filter.delegateHeight * 5 : filter.delegateHeight * filter.count)
        clip: true
        modal: true
        closePolicy: Popup.CloseOnEscape | Popup.CloseOnPressOutside

        ColumnLayout {
            anchors.fill: parent

            Label {
                font.pointSize: Qt.application.font.pointSize * 1.25
                font.bold: true
                text: qsTr("Available filters")
                Layout.alignment: Qt.AlignHCenter | Qt.AlignTop
            }

            ScrollView {
                Layout.alignment: Qt.AlignCenter
                Layout.fillWidth: true
                Layout.fillHeight: true
                clip: true
                contentHeight: filter.implicitHeight
                HistoryFilterList {
                    id: filter
                    interactive: false
                }
            }
        } // ColumnLayout
    } // ToolTip

    DialogTransactionDetails {
        id: dialogTransactionDetails
        anchors.centerIn: Overlay.overlay

        modal: true
        focus: true
        standardButtons: Dialog.Ok

        date: listTransactions.currentItem.modelDate
        status: listTransactions.currentItem.modelStatus
        type: listTransactions.currentItem.modelType
        amount: listTransactions.currentItem.modelAmount
        hoursReceived: listTransactions.currentItem.modelHoursReceived
        hoursBurned: listTransactions.currentItem.modelHoursBurned
        transactionID: listTransactions.currentItem.modelTransactionID
    }

    QHistory {
        id: modelTransactions
    }
}
