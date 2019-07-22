package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/gui"
	"github.com/therecipe/qt/qml"

	//"github.com/therecipe/qt/quick"
	//"github.com/therecipe/qt/widgets"
	"os"

	_ "github.com/simelo/FiberCryptoWallet/src/api/history"
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
	url := core.NewQUrl3("qrc:/ui/src/ui/main.qml", 0)

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
