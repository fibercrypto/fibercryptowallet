package main

import (
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	//"github.com/therecipe/qt/quick"
	"github.com/therecipe/qt/widgets"
	"os"
)

func main() {
	core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)


	widgets.NewQApplication(len(os.Args), os.Args)

	engine := qml.NewQQmlApplicationEngine(nil)
	engine.Load(core.NewQUrl3("qrc:/ui/qml/main.qml", 0))

	widgets.QApplication_Exec()
	//app := widgets.NewQApplication(len(os.Args), os.Args)
	//view := quick.NewQQuickView(nil)
	//view.SetTitle("gotemplate Example")
	//view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	//view.SetSource(core.NewQUrl3("qml/main.qml", 0))
	//view.Show()

	//app.Exec()
	//
	////view := quick.NewQQuickView(nil)
	//engine := qml.NewQQmlApplicationEngine(nil)
	////engine.SetBaseUrl("bridge Example")
	////view.SetResizeMode(quick.QQuickView__SizeRootObjectToView)
	////url := core.NewQUrl3("qrc:/src/qml/main.qml", 0)
	//engine.SetBaseUrl(core.NewQUrl3("qrc:/src/qml/main.qml", 0))
	////view.SetSource(core.NewQUrl3("qrc:/qml/main.qml", 0))
	////view.Show()
	//
	//
	//app.Exec()

	//core.QCoreApplication_SetAttribute(core.Qt__AA_EnableHighDpiScaling, true)
	//
	//// needs to be called once before you can start using the QWidgets
	//app := widgets.NewQApplication(len(os.Args), os.Args)
	//
	//// create the qml application engine
	//engine := qml.NewQQmlApplicationEngine(nil)
	//
	//url := core.NewQUrl3("qrc:/src/qml/main.qml", 0)
	//
	//// your code here!!!
	//
	///*
	//core.QObject.ConnectEvent(engine, widgets.objectCreated)
		//connect(&engine, widgets.objectCreated,)
	//core.QCoreApplication_Exit(-1)
	//core.Qt__QueuedConnection
	// */
	//engine.Load(url)
	//
	//// start the main Qt event loop
	//// and block until app.Exit() is called
	//// or the window is closed by the user
	//app.Exec()
}
//
//func handler(url core.QUrl, object core.QObject, objUrl core.QUrl) {
//	if object == nil && url == objUrl {
//		core.QCoreApplication_Exit(-1)
//	}
//
//}
