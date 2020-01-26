package main

import (
	"github.com/fibercrypto/fibercryptowallet/src/params"
	"os"

	_ "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin"
	_ "github.com/fibercrypto/fibercryptowallet/src/models"
	_ "github.com/fibercrypto/fibercryptowallet/src/models/addressBook"
	_ "github.com/fibercrypto/fibercryptowallet/src/models/history"
	_ "github.com/fibercrypto/fibercryptowallet/src/models/pending"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
)

func main() {

	// Enable High DPI scaling
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// Create a Qt Application
	// Avoid using a `QApplication` as it requires `widgets` module
	app := gui.NewQGuiApplication(len(os.Args), os.Args)

	// Set the application information
	app.SetOrganizationName(params.OrganizationName)
	app.SetOrganizationDomain(params.OrganizationDomain)
	app.SetApplicationName(params.ApplicationName)
	app.SetApplicationVersion(params.ApplicationVersion)
	app.SetWindowIcon(gui.NewQIcon5(":/images/resources/images/icons/appIcon/appIcon.png"))

	// Add this monospaced font
	gui.QFontDatabase_AddApplicationFont(":/fonts/resources/fonts/code-new-roman/code-new-roman.otf")

	engine := qml.NewQQmlApplicationEngine(nil)
	// To speed up UI development, loading QML files from resources is disabled, but it must be re-enabled in order to make a release
	// url := core.NewQUrl3("qrc:/ui/src/ui/splash.qml", 0)
	url := core.NewQUrl3("src/ui/splash.qml", 0) // disable this to make a release

	engine.ConnectSignal(engine.ConnectObjectCreated, func(object *core.QObject, objUrl *core.QUrl) {
		if object.Pointer() == nil && url.ToString(0) == objUrl.ToString(0)[len(objUrl.ToString(0)) - len(url.ToString(0)):] {
			app.Exit(-1)
		}
	}, core.Qt__QueuedConnection)
	engine.Load(url)

	app.Exec()
}
