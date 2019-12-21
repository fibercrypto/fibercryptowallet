package main

import (
	"os"

	_ "github.com/fibercrypto/fibercryptowallet/src/coin/skycoin"
	_ "github.com/fibercrypto/fibercryptowallet/src/models"
	_ "github.com/fibercrypto/fibercryptowallet/src/models/addressBook"
	_ "github.com/fibercrypto/fibercryptowallet/src/models/history"
	_ "github.com/fibercrypto/fibercryptowallet/src/models/pending"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"
)

func main() {

	// Enable High DPI scaling
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// Create a Qt Application
	// On embedded devices, a `QGuiApplication` must be used instead
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// Setup the splash screen
	splashPixmap := gui.NewQPixmap3(":/images/resources/images/icons/appIcon/splash.png", "", core.Qt__AutoColor)
	splash := widgets.NewQSplashScreen(splashPixmap, core.Qt__SplashScreen)
	splash.Show()
	// Process the pending `show` event
	app.QCoreApplication.ProcessEvents(core.QEventLoop__AllEvents)

	// Set the application information
	app.SetOrganizationName("Simelo.Tech")
	app.SetOrganizationDomain("simelo.tech")
	app.SetApplicationName("FiberCryptoWallet")
	app.SetApplicationVersion("0.27.0")
	app.SetWindowIcon(gui.NewQIcon5(":/images/resources/images/icons/appIcon/appIcon.png"))

	// Add this monospaced font
	gui.QFontDatabase_AddApplicationFont(":/fonts/resources/fonts/code-new-roman/code-new-roman.otf")

	engine := qml.NewQQmlApplicationEngine(nil)
	// To speed up UI development, loading QML files from resources is disabled, but it must be re-enabled in order to make a release
	// url := core.NewQUrl3("qrc:/ui/src/ui/main.qml", 0)
	url := core.NewQUrl3("src/ui/main.qml", 0) // disable this to make a release

	// TODO: Find a way to use a `core.Qt__QueuedConnection`, so we can remove the flag `allOk`
	allOk := true
	engine.ConnectObjectCreated(func(object *core.QObject, objUrl *core.QUrl) {
		if object.Pointer() == nil && url.ToString(0) == objUrl.ToString(0) {
			allOk = false
			app.Exit(-1) // Ignored because we need a `core.Qt__QueuedConnection`
		}
	})
	engine.Load(url)
	splash.QWidget.Close()

	// A `core.Qt__QueuedConnection` will allows us to remove the condition bellow, leaving only `app.Exec()`
	if allOk == true {
		app.Exec()
	} else {
		app.Exit(-1)
	}
}
