

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QMetaMethod>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QString>
#include <QTimerEvent>
#include <QWindow>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


class Walletfdda3f: public QObject
{
Q_OBJECT
public:
	Walletfdda3f(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");Walletfdda3f_Walletfdda3f_QRegisterMetaType();Walletfdda3f_Walletfdda3f_QRegisterMetaTypes();callbackWalletfdda3f_Constructor(this);};
	 ~Walletfdda3f() { callbackWalletfdda3f_DestroyWallet(this); };
	void childEvent(QChildEvent * event) { callbackWalletfdda3f_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackWalletfdda3f_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackWalletfdda3f_CustomEvent(this, event); };
	void deleteLater() { callbackWalletfdda3f_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackWalletfdda3f_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackWalletfdda3f_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackWalletfdda3f_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackWalletfdda3f_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackWalletfdda3f_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackWalletfdda3f_TimerEvent(this, event); };
signals:
public slots:
private:
};

Q_DECLARE_METATYPE(Walletfdda3f*)


void Walletfdda3f_Walletfdda3f_QRegisterMetaTypes() {
}

int Walletfdda3f_Walletfdda3f_QRegisterMetaType()
{
	return qRegisterMetaType<Walletfdda3f*>();
}

int Walletfdda3f_Walletfdda3f_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<Walletfdda3f*>(const_cast<const char*>(typeName));
}

int Walletfdda3f_Walletfdda3f_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Walletfdda3f>();
#else
	return 0;
#endif
}

int Walletfdda3f_Walletfdda3f_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<Walletfdda3f>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* Walletfdda3f___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Walletfdda3f___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Walletfdda3f___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* Walletfdda3f___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void Walletfdda3f___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* Walletfdda3f___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* Walletfdda3f___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Walletfdda3f___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Walletfdda3f___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Walletfdda3f___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Walletfdda3f___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Walletfdda3f___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Walletfdda3f___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void Walletfdda3f___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* Walletfdda3f___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* Walletfdda3f_NewWallet(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new Walletfdda3f(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new Walletfdda3f(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new Walletfdda3f(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new Walletfdda3f(static_cast<QWindow*>(parent));
	} else {
		return new Walletfdda3f(static_cast<QObject*>(parent));
	}
}

void Walletfdda3f_DestroyWallet(void* ptr)
{
	static_cast<Walletfdda3f*>(ptr)->~Walletfdda3f();
}

void Walletfdda3f_DestroyWalletDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void Walletfdda3f_ChildEventDefault(void* ptr, void* event)
{
	static_cast<Walletfdda3f*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void Walletfdda3f_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Walletfdda3f*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void Walletfdda3f_CustomEventDefault(void* ptr, void* event)
{
	static_cast<Walletfdda3f*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void Walletfdda3f_DeleteLaterDefault(void* ptr)
{
	static_cast<Walletfdda3f*>(ptr)->QObject::deleteLater();
}

void Walletfdda3f_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<Walletfdda3f*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char Walletfdda3f_EventDefault(void* ptr, void* e)
{
	return static_cast<Walletfdda3f*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char Walletfdda3f_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<Walletfdda3f*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void Walletfdda3f_TimerEventDefault(void* ptr, void* event)
{
	static_cast<Walletfdda3f*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
