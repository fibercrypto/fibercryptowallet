package main

import (
	_ "github.com/fibercrypto/FiberCryptoWallet/src/models"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"

	"os"

	_ "github.com/fibercrypto/FiberCryptoWallet/src/models"
	_ "github.com/fibercrypto/FiberCryptoWallet/src/models/history"
)

func main() {

	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	app := gui.NewQGuiApplication(len(os.Args), os.Args)
	app.SetOrganizationName("Simelo")
	app.SetApplicationName("FiberCrypto Wallet")
	app.SetApplicationVersion("0.1")
	app.SetWindowIcon(gui.NewQIcon5(":/images/resources/images/icons/appIcon.png"))

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

	// A `core.Qt__QueuedConnection` will allows us to remove the condition bellow, leaving only `app.Exec()`
	if allOk == true {
		app.Exec()
	} else {
		app.Exit(-1)
	}
}
