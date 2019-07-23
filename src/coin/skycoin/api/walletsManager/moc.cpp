

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAbstractListModel>
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

class AddressesModelbaab1e: public QAbstractListModel{
public:
	AddressesModelbaab1e(QObject *parent) : QAbstractListModel(parent) {};

};
class WalletModelbaab1e: public QAbstractListModel{
public:
	WalletModelbaab1e(QObject *parent) : QAbstractListModel(parent) {};

};

class WalletManager3cf9d2: public QObject
{
Q_OBJECT
public:
	WalletManager3cf9d2(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");WalletManager3cf9d2_WalletManager3cf9d2_QRegisterMetaType();WalletManager3cf9d2_WalletManager3cf9d2_QRegisterMetaTypes();callbackWalletManager3cf9d2_Constructor(this);};
	 ~WalletManager3cf9d2() { callbackWalletManager3cf9d2_DestroyWalletManager(this); };
	void childEvent(QChildEvent * event) { callbackWalletManager3cf9d2_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackWalletManager3cf9d2_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackWalletManager3cf9d2_CustomEvent(this, event); };
	void deleteLater() { callbackWalletManager3cf9d2_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackWalletManager3cf9d2_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackWalletManager3cf9d2_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackWalletManager3cf9d2_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackWalletManager3cf9d2_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackWalletManager3cf9d2_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackWalletManager3cf9d2_TimerEvent(this, event); };
signals:
public slots:
	void createEncryptedWallet(WalletModelbaab1e* walletM, QString seed, QString label, QString password, qint32 scanN) { QByteArray t92713d = seed.toUtf8(); Moc_PackedString seedPacked = { const_cast<char*>(t92713d.prepend("WHITESPACE").constData()+10), t92713d.size()-10 };QByteArray t64c653 = label.toUtf8(); Moc_PackedString labelPacked = { const_cast<char*>(t64c653.prepend("WHITESPACE").constData()+10), t64c653.size()-10 };QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackWalletManager3cf9d2_CreateEncryptedWallet(this, walletM, seedPacked, labelPacked, passwordPacked, scanN); };
	void createUnencryptedWallet(WalletModelbaab1e* walletM, QString seed, QString label, qint32 scanN) { QByteArray t92713d = seed.toUtf8(); Moc_PackedString seedPacked = { const_cast<char*>(t92713d.prepend("WHITESPACE").constData()+10), t92713d.size()-10 };QByteArray t64c653 = label.toUtf8(); Moc_PackedString labelPacked = { const_cast<char*>(t64c653.prepend("WHITESPACE").constData()+10), t64c653.size()-10 };callbackWalletManager3cf9d2_CreateUnencryptedWallet(this, walletM, seedPacked, labelPacked, scanN); };
	QString getNewSeed(qint32 entropy) { return ({ Moc_PackedString tempVal = callbackWalletManager3cf9d2_GetNewSeed(this, entropy); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	qint32 verifySeed(QString seed) { QByteArray t92713d = seed.toUtf8(); Moc_PackedString seedPacked = { const_cast<char*>(t92713d.prepend("WHITESPACE").constData()+10), t92713d.size()-10 };return callbackWalletManager3cf9d2_VerifySeed(this, seedPacked); };
	void newWalletAddress(AddressesModelbaab1e* addressM, QString id, qint32 n, QString password) { QByteArray t87ea5d = id.toUtf8(); Moc_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackWalletManager3cf9d2_NewWalletAddress(this, addressM, idPacked, n, passwordPacked); };
	void encryptWallet(WalletModelbaab1e* walletM, QString id, QString password) { QByteArray t87ea5d = id.toUtf8(); Moc_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackWalletManager3cf9d2_EncryptWallet(this, walletM, idPacked, passwordPacked); };
	void decryptWallet(WalletModelbaab1e* walletM, QString id, QString password) { QByteArray t87ea5d = id.toUtf8(); Moc_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackWalletManager3cf9d2_DecryptWallet(this, walletM, idPacked, passwordPacked); };
private:
};

Q_DECLARE_METATYPE(WalletManager3cf9d2*)


void WalletManager3cf9d2_WalletManager3cf9d2_QRegisterMetaTypes() {
}

void WalletManager3cf9d2_CreateEncryptedWallet(void* ptr, void* walletM, struct Moc_PackedString seed, struct Moc_PackedString label, struct Moc_PackedString password, int scanN)
{
	QMetaObject::invokeMethod(static_cast<WalletManager3cf9d2*>(ptr), "createEncryptedWallet", Q_ARG(WalletModelbaab1e*, static_cast<WalletModelbaab1e*>(walletM)), Q_ARG(QString, QString::fromUtf8(seed.data, seed.len)), Q_ARG(QString, QString::fromUtf8(label.data, label.len)), Q_ARG(QString, QString::fromUtf8(password.data, password.len)), Q_ARG(qint32, scanN));
}

void WalletManager3cf9d2_CreateUnencryptedWallet(void* ptr, void* walletM, struct Moc_PackedString seed, struct Moc_PackedString label, int scanN)
{
	QMetaObject::invokeMethod(static_cast<WalletManager3cf9d2*>(ptr), "createUnencryptedWallet", Q_ARG(WalletModelbaab1e*, static_cast<WalletModelbaab1e*>(walletM)), Q_ARG(QString, QString::fromUtf8(seed.data, seed.len)), Q_ARG(QString, QString::fromUtf8(label.data, label.len)), Q_ARG(qint32, scanN));
}

struct Moc_PackedString WalletManager3cf9d2_GetNewSeed(void* ptr, int entropy)
{
	QString returnArg;
	QMetaObject::invokeMethod(static_cast<WalletManager3cf9d2*>(ptr), "getNewSeed", Q_RETURN_ARG(QString, returnArg), Q_ARG(qint32, entropy));
	return ({ QByteArray t8e5b69 = returnArg.toUtf8(); Moc_PackedString { const_cast<char*>(t8e5b69.prepend("WHITESPACE").constData()+10), t8e5b69.size()-10 }; });
}

int WalletManager3cf9d2_VerifySeed(void* ptr, struct Moc_PackedString seed)
{
	qint32 returnArg;
	QMetaObject::invokeMethod(static_cast<WalletManager3cf9d2*>(ptr), "verifySeed", Q_RETURN_ARG(qint32, returnArg), Q_ARG(QString, QString::fromUtf8(seed.data, seed.len)));
	return returnArg;
}

void WalletManager3cf9d2_NewWalletAddress(void* ptr, void* addressM, struct Moc_PackedString id, int n, struct Moc_PackedString password)
{
	QMetaObject::invokeMethod(static_cast<WalletManager3cf9d2*>(ptr), "newWalletAddress", Q_ARG(AddressesModelbaab1e*, static_cast<AddressesModelbaab1e*>(addressM)), Q_ARG(QString, QString::fromUtf8(id.data, id.len)), Q_ARG(qint32, n), Q_ARG(QString, QString::fromUtf8(password.data, password.len)));
}

void WalletManager3cf9d2_EncryptWallet(void* ptr, void* walletM, struct Moc_PackedString id, struct Moc_PackedString password)
{
	QMetaObject::invokeMethod(static_cast<WalletManager3cf9d2*>(ptr), "encryptWallet", Q_ARG(WalletModelbaab1e*, static_cast<WalletModelbaab1e*>(walletM)), Q_ARG(QString, QString::fromUtf8(id.data, id.len)), Q_ARG(QString, QString::fromUtf8(password.data, password.len)));
}

void WalletManager3cf9d2_DecryptWallet(void* ptr, void* walletM, struct Moc_PackedString id, struct Moc_PackedString password)
{
	QMetaObject::invokeMethod(static_cast<WalletManager3cf9d2*>(ptr), "decryptWallet", Q_ARG(WalletModelbaab1e*, static_cast<WalletModelbaab1e*>(walletM)), Q_ARG(QString, QString::fromUtf8(id.data, id.len)), Q_ARG(QString, QString::fromUtf8(password.data, password.len)));
}

int WalletManager3cf9d2_WalletManager3cf9d2_QRegisterMetaType()
{
	return qRegisterMetaType<WalletManager3cf9d2*>();
}

int WalletManager3cf9d2_WalletManager3cf9d2_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<WalletManager3cf9d2*>(const_cast<const char*>(typeName));
}

int WalletManager3cf9d2_WalletManager3cf9d2_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletManager3cf9d2>();
#else
	return 0;
#endif
}

int WalletManager3cf9d2_WalletManager3cf9d2_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletManager3cf9d2>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* WalletManager3cf9d2___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager3cf9d2___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletManager3cf9d2___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* WalletManager3cf9d2___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletManager3cf9d2___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* WalletManager3cf9d2___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* WalletManager3cf9d2___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager3cf9d2___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletManager3cf9d2___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletManager3cf9d2___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager3cf9d2___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletManager3cf9d2___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletManager3cf9d2___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager3cf9d2___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletManager3cf9d2___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletManager3cf9d2_NewWalletManager(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new WalletManager3cf9d2(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new WalletManager3cf9d2(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new WalletManager3cf9d2(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new WalletManager3cf9d2(static_cast<QWindow*>(parent));
	} else {
		return new WalletManager3cf9d2(static_cast<QObject*>(parent));
	}
}

void WalletManager3cf9d2_DestroyWalletManager(void* ptr)
{
	static_cast<WalletManager3cf9d2*>(ptr)->~WalletManager3cf9d2();
}

void WalletManager3cf9d2_DestroyWalletManagerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void WalletManager3cf9d2_ChildEventDefault(void* ptr, void* event)
{
	static_cast<WalletManager3cf9d2*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void WalletManager3cf9d2_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletManager3cf9d2*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void WalletManager3cf9d2_CustomEventDefault(void* ptr, void* event)
{
	static_cast<WalletManager3cf9d2*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void WalletManager3cf9d2_DeleteLaterDefault(void* ptr)
{
	static_cast<WalletManager3cf9d2*>(ptr)->QObject::deleteLater();
}

void WalletManager3cf9d2_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletManager3cf9d2*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char WalletManager3cf9d2_EventDefault(void* ptr, void* e)
{
	return static_cast<WalletManager3cf9d2*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char WalletManager3cf9d2_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<WalletManager3cf9d2*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void WalletManager3cf9d2_TimerEventDefault(void* ptr, void* event)
{
	static_cast<WalletManager3cf9d2*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
