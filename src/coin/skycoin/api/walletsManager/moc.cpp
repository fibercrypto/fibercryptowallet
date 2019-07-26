

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

class QWallet445aa6: public QObject{
public:
	QWallet445aa6(QObject *parent) : QObject(parent) {};

};

class WalletManager64bdd5: public QObject
{
Q_OBJECT
public:
	WalletManager64bdd5(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");WalletManager64bdd5_WalletManager64bdd5_QRegisterMetaType();WalletManager64bdd5_WalletManager64bdd5_QRegisterMetaTypes();callbackWalletManager64bdd5_Constructor(this);};
	 ~WalletManager64bdd5() { callbackWalletManager64bdd5_DestroyWalletManager(this); };
	void childEvent(QChildEvent * event) { callbackWalletManager64bdd5_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackWalletManager64bdd5_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackWalletManager64bdd5_CustomEvent(this, event); };
	void deleteLater() { callbackWalletManager64bdd5_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackWalletManager64bdd5_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackWalletManager64bdd5_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackWalletManager64bdd5_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackWalletManager64bdd5_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackWalletManager64bdd5_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackWalletManager64bdd5_TimerEvent(this, event); };
signals:
public slots:
	QWallet445aa6* createEncryptedWallet(QString seed, QString label, QString password, qint32 scanN) { QByteArray t92713d = seed.toUtf8(); Moc_PackedString seedPacked = { const_cast<char*>(t92713d.prepend("WHITESPACE").constData()+10), t92713d.size()-10 };QByteArray t64c653 = label.toUtf8(); Moc_PackedString labelPacked = { const_cast<char*>(t64c653.prepend("WHITESPACE").constData()+10), t64c653.size()-10 };QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };return static_cast<QWallet445aa6*>(callbackWalletManager64bdd5_CreateEncryptedWallet(this, seedPacked, labelPacked, passwordPacked, scanN)); };
	QWallet445aa6* createUnencryptedWallet(QString seed, QString label, qint32 scanN) { QByteArray t92713d = seed.toUtf8(); Moc_PackedString seedPacked = { const_cast<char*>(t92713d.prepend("WHITESPACE").constData()+10), t92713d.size()-10 };QByteArray t64c653 = label.toUtf8(); Moc_PackedString labelPacked = { const_cast<char*>(t64c653.prepend("WHITESPACE").constData()+10), t64c653.size()-10 };return static_cast<QWallet445aa6*>(callbackWalletManager64bdd5_CreateUnencryptedWallet(this, seedPacked, labelPacked, scanN)); };
	QString getNewSeed(qint32 entropy) { return ({ Moc_PackedString tempVal = callbackWalletManager64bdd5_GetNewSeed(this, entropy); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	qint32 verifySeed(QString seed) { QByteArray t92713d = seed.toUtf8(); Moc_PackedString seedPacked = { const_cast<char*>(t92713d.prepend("WHITESPACE").constData()+10), t92713d.size()-10 };return callbackWalletManager64bdd5_VerifySeed(this, seedPacked); };
	void encryptWallet(QString id, QString password) { QByteArray t87ea5d = id.toUtf8(); Moc_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackWalletManager64bdd5_EncryptWallet(this, idPacked, passwordPacked); };
	void decryptWallet(QString id, QString password) { QByteArray t87ea5d = id.toUtf8(); Moc_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackWalletManager64bdd5_DecryptWallet(this, idPacked, passwordPacked); };
private:
};

Q_DECLARE_METATYPE(WalletManager64bdd5*)


void WalletManager64bdd5_WalletManager64bdd5_QRegisterMetaTypes() {
}

void* WalletManager64bdd5_CreateEncryptedWallet(void* ptr, struct Moc_PackedString seed, struct Moc_PackedString label, struct Moc_PackedString password, int scanN)
{
	QWallet445aa6* returnArg;
	QMetaObject::invokeMethod(static_cast<WalletManager64bdd5*>(ptr), "createEncryptedWallet", Q_RETURN_ARG(QWallet445aa6*, returnArg), Q_ARG(QString, QString::fromUtf8(seed.data, seed.len)), Q_ARG(QString, QString::fromUtf8(label.data, label.len)), Q_ARG(QString, QString::fromUtf8(password.data, password.len)), Q_ARG(qint32, scanN));
	return returnArg;
}

void* WalletManager64bdd5_CreateUnencryptedWallet(void* ptr, struct Moc_PackedString seed, struct Moc_PackedString label, int scanN)
{
	QWallet445aa6* returnArg;
	QMetaObject::invokeMethod(static_cast<WalletManager64bdd5*>(ptr), "createUnencryptedWallet", Q_RETURN_ARG(QWallet445aa6*, returnArg), Q_ARG(QString, QString::fromUtf8(seed.data, seed.len)), Q_ARG(QString, QString::fromUtf8(label.data, label.len)), Q_ARG(qint32, scanN));
	return returnArg;
}

struct Moc_PackedString WalletManager64bdd5_GetNewSeed(void* ptr, int entropy)
{
	QString returnArg;
	QMetaObject::invokeMethod(static_cast<WalletManager64bdd5*>(ptr), "getNewSeed", Q_RETURN_ARG(QString, returnArg), Q_ARG(qint32, entropy));
	return ({ QByteArray t8e5b69 = returnArg.toUtf8(); Moc_PackedString { const_cast<char*>(t8e5b69.prepend("WHITESPACE").constData()+10), t8e5b69.size()-10 }; });
}

int WalletManager64bdd5_VerifySeed(void* ptr, struct Moc_PackedString seed)
{
	qint32 returnArg;
	QMetaObject::invokeMethod(static_cast<WalletManager64bdd5*>(ptr), "verifySeed", Q_RETURN_ARG(qint32, returnArg), Q_ARG(QString, QString::fromUtf8(seed.data, seed.len)));
	return returnArg;
}

void WalletManager64bdd5_EncryptWallet(void* ptr, struct Moc_PackedString id, struct Moc_PackedString password)
{
	QMetaObject::invokeMethod(static_cast<WalletManager64bdd5*>(ptr), "encryptWallet", Q_ARG(QString, QString::fromUtf8(id.data, id.len)), Q_ARG(QString, QString::fromUtf8(password.data, password.len)));
}

void WalletManager64bdd5_DecryptWallet(void* ptr, struct Moc_PackedString id, struct Moc_PackedString password)
{
	QMetaObject::invokeMethod(static_cast<WalletManager64bdd5*>(ptr), "decryptWallet", Q_ARG(QString, QString::fromUtf8(id.data, id.len)), Q_ARG(QString, QString::fromUtf8(password.data, password.len)));
}

int WalletManager64bdd5_WalletManager64bdd5_QRegisterMetaType()
{
	return qRegisterMetaType<WalletManager64bdd5*>();
}

int WalletManager64bdd5_WalletManager64bdd5_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<WalletManager64bdd5*>(const_cast<const char*>(typeName));
}

int WalletManager64bdd5_WalletManager64bdd5_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletManager64bdd5>();
#else
	return 0;
#endif
}

int WalletManager64bdd5_WalletManager64bdd5_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletManager64bdd5>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* WalletManager64bdd5___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager64bdd5___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletManager64bdd5___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* WalletManager64bdd5___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletManager64bdd5___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* WalletManager64bdd5___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* WalletManager64bdd5___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager64bdd5___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletManager64bdd5___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletManager64bdd5___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager64bdd5___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletManager64bdd5___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletManager64bdd5___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager64bdd5___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletManager64bdd5___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletManager64bdd5_NewWalletManager(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new WalletManager64bdd5(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new WalletManager64bdd5(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new WalletManager64bdd5(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new WalletManager64bdd5(static_cast<QWindow*>(parent));
	} else {
		return new WalletManager64bdd5(static_cast<QObject*>(parent));
	}
}

void WalletManager64bdd5_DestroyWalletManager(void* ptr)
{
	static_cast<WalletManager64bdd5*>(ptr)->~WalletManager64bdd5();
}

void WalletManager64bdd5_DestroyWalletManagerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void WalletManager64bdd5_ChildEventDefault(void* ptr, void* event)
{
	static_cast<WalletManager64bdd5*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void WalletManager64bdd5_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletManager64bdd5*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void WalletManager64bdd5_CustomEventDefault(void* ptr, void* event)
{
	static_cast<WalletManager64bdd5*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void WalletManager64bdd5_DeleteLaterDefault(void* ptr)
{
	static_cast<WalletManager64bdd5*>(ptr)->QObject::deleteLater();
}

void WalletManager64bdd5_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletManager64bdd5*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char WalletManager64bdd5_EventDefault(void* ptr, void* e)
{
	return static_cast<WalletManager64bdd5*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char WalletManager64bdd5_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<WalletManager64bdd5*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void WalletManager64bdd5_TimerEventDefault(void* ptr, void* event)
{
	static_cast<WalletManager64bdd5*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
