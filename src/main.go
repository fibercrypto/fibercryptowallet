package main

import (
	"os"

	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	"github.com/therecipe/qt/widgets"
)

func main() {

	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)

	// needs to be called once before you can start using the QWidgets
	app := widgets.NewQApplication(len(os.Args), os.Args)

	// create the qml application engine
	engine := qml.NewQQmlApplicationEngine(nil)

	url := core.NewQUrl3("qrc:/qml/main.qml", 0)

	// your code here!!!

	/*
	 *  core.QObject.connect(&engine, widgets.objectCreated,)
	 *  core.QCoreApplication_Exit(-1)
	 *	core.Qt__QueuedConnection
	 */
	engine.Load(url)

	// start the main Qt event loop
	// and block until app.Exit() is called
	// or the window is closed by the user
	app.Exec()
}

func handler(url core.QUrl, object core.QObject, objUrl core.QUrl) {
	if object == nil && url == objUrl {
		core.QCoreApplication_Exit(-1)
	}

}
