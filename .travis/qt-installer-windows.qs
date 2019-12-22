

var INSTALL_COMPONENTS = [
    "qt.qt5.5125.win32_mingw73",
    "qt.qt5.5125.win32_msvc2017",
    "qt.qt5.5125.win64_mingw73",
    "qt.qt5.5125.win64_msvc2017_64",
    "qt.qt5.5125.debug_info",
    "qt.qt5.5125.debug_info.win32_msvc2017",
    "qt.qt5.5125.debug_info.win64_msvc2017_64",
    "qt.qt5.5125.qtcharts",
    "qt.qt5.5125.qtcharts.win32_mingw73",
    "qt.qt5.5125.qtcharts.win32_msvc2017",
    "qt.qt5.5125.qtcharts.win64_mingw73",
    "qt.qt5.5125.qtcharts.win64_msvc2017_64",
    "qt.qt5.5125.qtdatavis3d",
    "qt.qt5.5125.qtdatavis3d.win32_mingw73",
    "qt.qt5.5125.qtdatavis3d.win32_msvc2017",
    "qt.qt5.5125.qtdatavis3d.win64_mingw73",
    "qt.qt5.5125.qtdatavis3d.win64_msvc2017_64",
    "qt.qt5.5125.qtnetworkauth",
    "qt.qt5.5125.qtnetworkauth.win32_mingw73",
    "qt.qt5.5125.qtnetworkauth.win32_msvc2017",
    "qt.qt5.5125.qtnetworkauth.win64_mingw73",
    "qt.qt5.5125.qtnetworkauth.win64_msvc2017_64",
    "qt.qt5.5125.qtpurchasing",
    "qt.qt5.5125.qtpurchasing.win32_mingw73",
    "qt.qt5.5125.qtpurchasing.win32_msvc2017",
    "qt.qt5.5125.qtpurchasing.win64_mingw73",
    "qt.qt5.5125.qtpurchasing.win64_msvc2017_64",
    "qt.qt5.5125.qtscript",
    "qt.qt5.5125.qtscript.win32_mingw73",
    "qt.qt5.5125.qtscript.win32_msvc2017",
    "qt.qt5.5125.qtscript.win64_mingw73",
    "qt.qt5.5125.qtscript.win64_msvc2017_64",
    "qt.qt5.5125.qtvirtualkeyboard",
    "qt.qt5.5125.qtvirtualkeyboard.win32_mingw73",
    "qt.qt5.5125.qtvirtualkeyboard.win32_msvc2017",
    "qt.qt5.5125.qtvirtualkeyboard.win64_mingw73",
    "qt.qt5.5125.qtvirtualkeyboard.win64_msvc2017_64",
    "qt.qt5.5125.qtwebengine",
    "qt.qt5.5125.qtwebengine.win32_msvc2017",
    "qt.qt5.5125.qtwebengine.win64_msvc2017_64",
    "qt.qt5.5125.qtwebglplugin",
    "qt.qt5.5125.qtwebglplugin.win32_mingw73",
    "qt.qt5.5125.qtwebglplugin.win32_msvc2017",
    "qt.qt5.5125.qtwebglplugin.win64_mingw73",
    "qt.qt5.5125.qtwebglplugin.win64_msvc2017_64",
];

function Controller() {
    installer.autoRejectMessageBoxes();
    installer.installationFinished.connect(function() {
        gui.clickButton(buttons.NextButton);
    })
}

Controller.prototype.WelcomePageCallback = function() {
    // click delay here because the next button is initially disabled for ~1 second
    gui.clickButton(buttons.NextButton, 3000);
}

Controller.prototype.CredentialsPageCallback = function() {
    gui.clickButton(buttons.NextButton);
}

Controller.prototype.IntroductionPageCallback = function() {
    gui.clickButton(buttons.NextButton);
}

Controller.prototype.DynamicTelemetryPluginFormCallback = function() {
    gui.currentPageWidget().TelemetryPluginForm.statisticGroupBox.disableStatisticRadioButton.setChecked(true);
    gui.clickButton(buttons.NextButton);

    //for(var key in widget.TelemetryPluginForm.statisticGroupBox){
    //    console.log(key);
    //}
}

Controller.prototype.TargetDirectoryPageCallback = function()
{
    // Keep default at "C:\Qt".
    //gui.currentPageWidget().TargetDirectoryLineEdit.setText("E:\\Qt");
    //gui.currentPageWidget().TargetDirectoryLineEdit.setText(installer.value("HomeDir") + "/Qt");
    gui.clickButton(buttons.NextButton);
}

Controller.prototype.ComponentSelectionPageCallback = function() {

    // https://doc-snapshots.qt.io/qtifw-3.1/noninteractive.html
    var page = gui.pageWidgetByObjectName("ComponentSelectionPage");

    var archiveCheckBox = gui.findChild(page, "Archive");
    var latestCheckBox = gui.findChild(page, "Latest releases");
    var fetchButton = gui.findChild(page, "FetchCategoryButton");

    archiveCheckBox.click();
    latestCheckBox.click();
    fetchButton.click();

    var widget = gui.currentPageWidget();

    widget.deselectAll();

    for (var i = 0; i < INSTALL_COMPONENTS.length; i++) {
        widget.selectComponent(INSTALL_COMPONENTS[i]);
    }

    // widget.deselectComponent("qt.tools.qtcreator");
    // widget.deselectComponent("qt.55.qt3d");

    gui.clickButton(buttons.NextButton);
}

Controller.prototype.LicenseAgreementPageCallback = function() {
    gui.currentPageWidget().AcceptLicenseRadioButton.setChecked(true);
    gui.clickButton(buttons.NextButton);
}

Controller.prototype.StartMenuDirectoryPageCallback = function() {
    gui.clickButton(buttons.NextButton);
}

Controller.prototype.ReadyForInstallationPageCallback = function()
{
    gui.clickButton(buttons.NextButton);
}

Controller.prototype.FinishedPageCallback = function() {
var checkBoxForm = gui.currentPageWidget().LaunchQtCreatorCheckBoxForm;
if (checkBoxForm && checkBoxForm.launchQtCreatorCheckBox) {
    checkBoxForm.launchQtCreatorCheckBox.checked = false;
}
    gui.clickButton(buttons.FinishButton);
}

