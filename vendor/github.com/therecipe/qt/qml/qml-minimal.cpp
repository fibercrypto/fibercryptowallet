// +build minimal

#define protected public
#define private public

#include "qml-minimal.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QJSEngine>
#include <QJSValue>
#include <QLatin1String>
#include <QLayout>
#include <QMetaMethod>
#include <QObject>
#include <QQmlEngine>
#include <QQmlError>
#include <QQmlParserStatus>
#include <QQuickItem>
#include <QString>
#include <QTimerEvent>
#include <QUrl>
#include <QVariant>
#include <QWidget>
#include <QWindow>
#include <QStringList>

class MyQJSEngine: public QJSEngine
{
public:
	MyQJSEngine() : QJSEngine() {QJSEngine_QJSEngine_QRegisterMetaType();};
	MyQJSEngine(QObject *parent) : QJSEngine(parent) {QJSEngine_QJSEngine_QRegisterMetaType();};
	 ~MyQJSEngine() { callbackQJSEngine_DestroyQJSEngine(this); };
	void childEvent(QChildEvent * event) { callbackQJSEngine_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQJSEngine_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQJSEngine_CustomEvent(this, event); };
	void deleteLater() { callbackQJSEngine_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQJSEngine_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQJSEngine_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQJSEngine_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQJSEngine_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtQml_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQJSEngine_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQJSEngine_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQJSEngine*)

int QJSEngine_QJSEngine_QRegisterMetaType(){qRegisterMetaType<QJSEngine*>(); return qRegisterMetaType<MyQJSEngine*>();}

void* QJSEngine_NewQJSEngine()
{
	return new MyQJSEngine();
}

void* QJSEngine_NewQJSEngine2(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQJSEngine(static_cast<QWindow*>(parent));
	} else {
		return new MyQJSEngine(static_cast<QObject*>(parent));
	}
}

void* QJSEngine_NewQObject(void* ptr, void* object)
{
	return new QJSValue(static_cast<QJSEngine*>(ptr)->newQObject(static_cast<QObject*>(object)));
}

void QJSEngine_DestroyQJSEngine(void* ptr)
{
	static_cast<QJSEngine*>(ptr)->~QJSEngine();
}

void QJSEngine_DestroyQJSEngineDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QJSEngine___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QJSEngine___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QJSEngine___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QJSEngine___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QJSEngine___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QJSEngine___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QJSEngine___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QJSEngine___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QJSEngine___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QJSEngine___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QJSEngine___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QJSEngine___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QJSEngine___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QJSEngine___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QJSEngine___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QJSEngine_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQmlEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlEngine*>(ptr)->QQmlEngine::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QJSEngine*>(ptr)->QJSEngine::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QJSEngine_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QQmlEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlEngine*>(ptr)->QQmlEngine::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QJSEngine*>(ptr)->QJSEngine::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QJSEngine_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQmlEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlEngine*>(ptr)->QQmlEngine::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QJSEngine*>(ptr)->QJSEngine::customEvent(static_cast<QEvent*>(event));
	}
}

void QJSEngine_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QQmlEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlEngine*>(ptr)->QQmlEngine::deleteLater();
	} else {
		static_cast<QJSEngine*>(ptr)->QJSEngine::deleteLater();
	}
}

void QJSEngine_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QQmlEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlEngine*>(ptr)->QQmlEngine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QJSEngine*>(ptr)->QJSEngine::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

char QJSEngine_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QQmlEngine*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQmlEngine*>(ptr)->QQmlEngine::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QJSEngine*>(ptr)->QJSEngine::event(static_cast<QEvent*>(e));
	}
}

char QJSEngine_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QQmlEngine*>(static_cast<QObject*>(ptr))) {
		return static_cast<QQmlEngine*>(ptr)->QQmlEngine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QJSEngine*>(ptr)->QJSEngine::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QJSEngine_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QQmlEngine*>(static_cast<QObject*>(ptr))) {
		static_cast<QQmlEngine*>(ptr)->QQmlEngine::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QJSEngine*>(ptr)->QJSEngine::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

void* QJSValue_NewQJSValue(long long value)
{
	return new QJSValue(static_cast<QJSValue::SpecialValue>(value));
}

void* QJSValue_NewQJSValue2(void* other)
{
	return new QJSValue(*static_cast<QJSValue*>(other));
}

void* QJSValue_NewQJSValue3(void* other)
{
	return new QJSValue(*static_cast<QJSValue*>(other));
}

void* QJSValue_NewQJSValue4(char value)
{
	return new QJSValue(value != 0);
}

void* QJSValue_NewQJSValue5(int value)
{
	return new QJSValue(value);
}

void* QJSValue_NewQJSValue6(unsigned int value)
{
	return new QJSValue(value);
}

void* QJSValue_NewQJSValue7(double value)
{
	return new QJSValue(value);
}

void* QJSValue_NewQJSValue8(struct QtQml_PackedString value)
{
	return new QJSValue(QString::fromUtf8(value.data, value.len));
}

void* QJSValue_NewQJSValue9(void* value)
{
	return new QJSValue(*static_cast<QLatin1String*>(value));
}

void* QJSValue_NewQJSValue10(char* value)
{
	return new QJSValue(const_cast<const char*>(value));
}

void* QJSValue_Property(void* ptr, struct QtQml_PackedString name)
{
	return new QJSValue(static_cast<QJSValue*>(ptr)->property(QString::fromUtf8(name.data, name.len)));
}

void* QJSValue_Property2(void* ptr, unsigned int arrayIndex)
{
	return new QJSValue(static_cast<QJSValue*>(ptr)->property(arrayIndex));
}

int QJSValue_ToInt(void* ptr)
{
	return static_cast<QJSValue*>(ptr)->toInt();
}

void* QJSValue_ToVariant(void* ptr)
{
	return new QVariant(static_cast<QJSValue*>(ptr)->toVariant());
}

void QJSValue_DestroyQJSValue(void* ptr)
{
	static_cast<QJSValue*>(ptr)->~QJSValue();
}

void* QJSValue___call_args_atList(void* ptr, int i)
{
	return new QJSValue(({QJSValue tmp = static_cast<QList<QJSValue>*>(ptr)->at(i); if (i == static_cast<QList<QJSValue>*>(ptr)->size()-1) { static_cast<QList<QJSValue>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QJSValue___call_args_setList(void* ptr, void* i)
{
	static_cast<QList<QJSValue>*>(ptr)->append(*static_cast<QJSValue*>(i));
}

void* QJSValue___call_args_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QJSValue>();
}

void* QJSValue___callAsConstructor_args_atList(void* ptr, int i)
{
	return new QJSValue(({QJSValue tmp = static_cast<QList<QJSValue>*>(ptr)->at(i); if (i == static_cast<QList<QJSValue>*>(ptr)->size()-1) { static_cast<QList<QJSValue>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QJSValue___callAsConstructor_args_setList(void* ptr, void* i)
{
	static_cast<QList<QJSValue>*>(ptr)->append(*static_cast<QJSValue*>(i));
}

void* QJSValue___callAsConstructor_args_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QJSValue>();
}

void* QJSValue___callWithInstance_args_atList(void* ptr, int i)
{
	return new QJSValue(({QJSValue tmp = static_cast<QList<QJSValue>*>(ptr)->at(i); if (i == static_cast<QList<QJSValue>*>(ptr)->size()-1) { static_cast<QList<QJSValue>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QJSValue___callWithInstance_args_setList(void* ptr, void* i)
{
	static_cast<QList<QJSValue>*>(ptr)->append(*static_cast<QJSValue*>(i));
}

void* QJSValue___callWithInstance_args_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QJSValue>();
}

class MyQQmlEngine: public QQmlEngine
{
public:
	MyQQmlEngine(QObject *parent = Q_NULLPTR) : QQmlEngine(parent) {QQmlEngine_QQmlEngine_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQJSEngine_Event(this, e) != 0; };
	 ~MyQQmlEngine() { callbackQQmlEngine_DestroyQQmlEngine(this); };
	void childEvent(QChildEvent * event) { callbackQJSEngine_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQJSEngine_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQJSEngine_CustomEvent(this, event); };
	void deleteLater() { callbackQJSEngine_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQJSEngine_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQJSEngine_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQJSEngine_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtQml_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQJSEngine_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQJSEngine_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQQmlEngine*)

int QQmlEngine_QQmlEngine_QRegisterMetaType(){qRegisterMetaType<QQmlEngine*>(); return qRegisterMetaType<MyQQmlEngine*>();}

void* QQmlEngine_NewQQmlEngine(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQQmlEngine(static_cast<QWindow*>(parent));
	} else {
		return new MyQQmlEngine(static_cast<QObject*>(parent));
	}
}

void QQmlEngine_AddImportPath(void* ptr, struct QtQml_PackedString path)
{
	static_cast<QQmlEngine*>(ptr)->addImportPath(QString::fromUtf8(path.data, path.len));
}

void QQmlEngine_DestroyQQmlEngine(void* ptr)
{
	static_cast<QQmlEngine*>(ptr)->~QQmlEngine();
}

void QQmlEngine_DestroyQQmlEngineDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

int QQmlEngine_QQmlEngine_QmlRegisterType(void* url, char* uri, int versionMajor, int versionMinor, char* qmlName)
{
	return qmlRegisterType(*static_cast<QUrl*>(url), const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
}

void* QQmlEngine___importPlugin_errors_atList(void* ptr, int i)
{
	return new QQmlError(({QQmlError tmp = static_cast<QList<QQmlError>*>(ptr)->at(i); if (i == static_cast<QList<QQmlError>*>(ptr)->size()-1) { static_cast<QList<QQmlError>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QQmlEngine___importPlugin_errors_setList(void* ptr, void* i)
{
	static_cast<QList<QQmlError>*>(ptr)->append(*static_cast<QQmlError*>(i));
}

void* QQmlEngine___importPlugin_errors_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QQmlError>();
}

void* QQmlEngine___qmlDebug_errors_atList3(void* ptr, int i)
{
	return new QQmlError(({QQmlError tmp = static_cast<QList<QQmlError>*>(ptr)->at(i); if (i == static_cast<QList<QQmlError>*>(ptr)->size()-1) { static_cast<QList<QQmlError>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QQmlEngine___qmlDebug_errors_setList3(void* ptr, void* i)
{
	static_cast<QList<QQmlError>*>(ptr)->append(*static_cast<QQmlError*>(i));
}

void* QQmlEngine___qmlDebug_errors_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QQmlError>();
}

void* QQmlEngine___qmlInfo_errors_atList3(void* ptr, int i)
{
	return new QQmlError(({QQmlError tmp = static_cast<QList<QQmlError>*>(ptr)->at(i); if (i == static_cast<QList<QQmlError>*>(ptr)->size()-1) { static_cast<QList<QQmlError>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QQmlEngine___qmlInfo_errors_setList3(void* ptr, void* i)
{
	static_cast<QList<QQmlError>*>(ptr)->append(*static_cast<QQmlError*>(i));
}

void* QQmlEngine___qmlInfo_errors_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QQmlError>();
}

void* QQmlEngine___qmlWarning_errors_atList3(void* ptr, int i)
{
	return new QQmlError(({QQmlError tmp = static_cast<QList<QQmlError>*>(ptr)->at(i); if (i == static_cast<QList<QQmlError>*>(ptr)->size()-1) { static_cast<QList<QQmlError>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QQmlEngine___qmlWarning_errors_setList3(void* ptr, void* i)
{
	static_cast<QList<QQmlError>*>(ptr)->append(*static_cast<QQmlError*>(i));
}

void* QQmlEngine___qmlWarning_errors_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QQmlError>();
}

void* QQmlEngine___warnings_warnings_atList(void* ptr, int i)
{
	return new QQmlError(({QQmlError tmp = static_cast<QList<QQmlError>*>(ptr)->at(i); if (i == static_cast<QList<QQmlError>*>(ptr)->size()-1) { static_cast<QList<QQmlError>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QQmlEngine___warnings_warnings_setList(void* ptr, void* i)
{
	static_cast<QList<QQmlError>*>(ptr)->append(*static_cast<QQmlError*>(i));
}

void* QQmlEngine___warnings_warnings_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QQmlError>();
}

void* QQmlError_NewQQmlError()
{
	return new QQmlError();
}

void* QQmlError_NewQQmlError2(void* other)
{
	return new QQmlError(*static_cast<QQmlError*>(other));
}

int QQmlError_Column(void* ptr)
{
	return static_cast<QQmlError*>(ptr)->column();
}

void* QQmlError_Object(void* ptr)
{
	return static_cast<QQmlError*>(ptr)->object();
}

void* QQmlError_Url(void* ptr)
{
	return new QUrl(static_cast<QQmlError*>(ptr)->url());
}

class MyQQmlParserStatus: public QQmlParserStatus
{
public:
	void classBegin() { callbackQQmlParserStatus_ClassBegin(this); };
	void componentComplete() { callbackQQmlParserStatus_ComponentComplete(this); };
};

void QQmlParserStatus_ClassBegin(void* ptr)
{
	static_cast<QQmlParserStatus*>(ptr)->classBegin();
}

void QQmlParserStatus_ComponentComplete(void* ptr)
{
	static_cast<QQmlParserStatus*>(ptr)->componentComplete();
}

