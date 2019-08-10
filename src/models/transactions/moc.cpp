

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAbstractListModel>
#include <QByteArray>
#include <QChildEvent>
#include <QDateTime>
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

class AddressListfb4c44: public QAbstractListModel{
public:
	AddressListfb4c44(QObject *parent) : QAbstractListModel(parent) {};

};

class TransactionDetailsecff1c: public QObject
{
Q_OBJECT
Q_PROPERTY(QDateTime date READ date WRITE setDate NOTIFY dateChanged)
Q_PROPERTY(qint32 status READ status WRITE setStatus NOTIFY statusChanged)
Q_PROPERTY(qint32 type READ type WRITE setType NOTIFY typeChanged)
Q_PROPERTY(QString amount READ amount WRITE setAmount NOTIFY amountChanged)
Q_PROPERTY(QString hoursTraspassed READ hoursTraspassed WRITE setHoursTraspassed NOTIFY hoursTraspassedChanged)
Q_PROPERTY(QString hoursBurned READ hoursBurned WRITE setHoursBurned NOTIFY hoursBurnedChanged)
Q_PROPERTY(QString transactionID READ transactionID WRITE setTransactionID NOTIFY transactionIDChanged)
Q_PROPERTY(AddressListfb4c44* addresses READ addresses WRITE setAddresses NOTIFY addressesChanged)
Q_PROPERTY(AddressListfb4c44* inputs READ inputs WRITE setInputs NOTIFY inputsChanged)
Q_PROPERTY(AddressListfb4c44* outputs READ outputs WRITE setOutputs NOTIFY outputsChanged)
public:
	TransactionDetailsecff1c(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");TransactionDetailsecff1c_TransactionDetailsecff1c_QRegisterMetaType();TransactionDetailsecff1c_TransactionDetailsecff1c_QRegisterMetaTypes();callbackTransactionDetailsecff1c_Constructor(this);};
	QDateTime date() { return *static_cast<QDateTime*>(callbackTransactionDetailsecff1c_Date(this)); };
	void setDate(QDateTime date) { callbackTransactionDetailsecff1c_SetDate(this, new QDateTime(date)); };
	void Signal_DateChanged(QDateTime date) { callbackTransactionDetailsecff1c_DateChanged(this, new QDateTime(date)); };
	qint32 status() { return callbackTransactionDetailsecff1c_Status(this); };
	void setStatus(qint32 status) { callbackTransactionDetailsecff1c_SetStatus(this, status); };
	void Signal_StatusChanged(qint32 status) { callbackTransactionDetailsecff1c_StatusChanged(this, status); };
	qint32 type() { return callbackTransactionDetailsecff1c_Type(this); };
	void setType(qint32 ty) { callbackTransactionDetailsecff1c_SetType(this, ty); };
	void Signal_TypeChanged(qint32 ty) { callbackTransactionDetailsecff1c_TypeChanged(this, ty); };
	QString amount() { return ({ Moc_PackedString tempVal = callbackTransactionDetailsecff1c_Amount(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setAmount(QString amount) { QByteArray t9cb6ff = amount.toUtf8(); Moc_PackedString amountPacked = { const_cast<char*>(t9cb6ff.prepend("WHITESPACE").constData()+10), t9cb6ff.size()-10 };callbackTransactionDetailsecff1c_SetAmount(this, amountPacked); };
	void Signal_AmountChanged(QString amount) { QByteArray t9cb6ff = amount.toUtf8(); Moc_PackedString amountPacked = { const_cast<char*>(t9cb6ff.prepend("WHITESPACE").constData()+10), t9cb6ff.size()-10 };callbackTransactionDetailsecff1c_AmountChanged(this, amountPacked); };
	QString hoursTraspassed() { return ({ Moc_PackedString tempVal = callbackTransactionDetailsecff1c_HoursTraspassed(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setHoursTraspassed(QString hoursTraspassed) { QByteArray teabfba = hoursTraspassed.toUtf8(); Moc_PackedString hoursTraspassedPacked = { const_cast<char*>(teabfba.prepend("WHITESPACE").constData()+10), teabfba.size()-10 };callbackTransactionDetailsecff1c_SetHoursTraspassed(this, hoursTraspassedPacked); };
	void Signal_HoursTraspassedChanged(QString hoursTraspassed) { QByteArray teabfba = hoursTraspassed.toUtf8(); Moc_PackedString hoursTraspassedPacked = { const_cast<char*>(teabfba.prepend("WHITESPACE").constData()+10), teabfba.size()-10 };callbackTransactionDetailsecff1c_HoursTraspassedChanged(this, hoursTraspassedPacked); };
	QString hoursBurned() { return ({ Moc_PackedString tempVal = callbackTransactionDetailsecff1c_HoursBurned(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setHoursBurned(QString hoursBurned) { QByteArray ta1ddd9 = hoursBurned.toUtf8(); Moc_PackedString hoursBurnedPacked = { const_cast<char*>(ta1ddd9.prepend("WHITESPACE").constData()+10), ta1ddd9.size()-10 };callbackTransactionDetailsecff1c_SetHoursBurned(this, hoursBurnedPacked); };
	void Signal_HoursBurnedChanged(QString hoursBurned) { QByteArray ta1ddd9 = hoursBurned.toUtf8(); Moc_PackedString hoursBurnedPacked = { const_cast<char*>(ta1ddd9.prepend("WHITESPACE").constData()+10), ta1ddd9.size()-10 };callbackTransactionDetailsecff1c_HoursBurnedChanged(this, hoursBurnedPacked); };
	QString transactionID() { return ({ Moc_PackedString tempVal = callbackTransactionDetailsecff1c_TransactionID(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setTransactionID(QString transactionID) { QByteArray tec2ac1 = transactionID.toUtf8(); Moc_PackedString transactionIDPacked = { const_cast<char*>(tec2ac1.prepend("WHITESPACE").constData()+10), tec2ac1.size()-10 };callbackTransactionDetailsecff1c_SetTransactionID(this, transactionIDPacked); };
	void Signal_TransactionIDChanged(QString transactionID) { QByteArray tec2ac1 = transactionID.toUtf8(); Moc_PackedString transactionIDPacked = { const_cast<char*>(tec2ac1.prepend("WHITESPACE").constData()+10), tec2ac1.size()-10 };callbackTransactionDetailsecff1c_TransactionIDChanged(this, transactionIDPacked); };
	AddressListfb4c44* addresses() { return static_cast<AddressListfb4c44*>(callbackTransactionDetailsecff1c_Addresses(this)); };
	void setAddresses(AddressListfb4c44* addresses) { callbackTransactionDetailsecff1c_SetAddresses(this, addresses); };
	void Signal_AddressesChanged(AddressListfb4c44* addresses) { callbackTransactionDetailsecff1c_AddressesChanged(this, addresses); };
	AddressListfb4c44* inputs() { return static_cast<AddressListfb4c44*>(callbackTransactionDetailsecff1c_Inputs(this)); };
	void setInputs(AddressListfb4c44* inputs) { callbackTransactionDetailsecff1c_SetInputs(this, inputs); };
	void Signal_InputsChanged(AddressListfb4c44* inputs) { callbackTransactionDetailsecff1c_InputsChanged(this, inputs); };
	AddressListfb4c44* outputs() { return static_cast<AddressListfb4c44*>(callbackTransactionDetailsecff1c_Outputs(this)); };
	void setOutputs(AddressListfb4c44* outputs) { callbackTransactionDetailsecff1c_SetOutputs(this, outputs); };
	void Signal_OutputsChanged(AddressListfb4c44* outputs) { callbackTransactionDetailsecff1c_OutputsChanged(this, outputs); };
	 ~TransactionDetailsecff1c() { callbackTransactionDetailsecff1c_DestroyTransactionDetails(this); };
	void childEvent(QChildEvent * event) { callbackTransactionDetailsecff1c_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackTransactionDetailsecff1c_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackTransactionDetailsecff1c_CustomEvent(this, event); };
	void deleteLater() { callbackTransactionDetailsecff1c_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackTransactionDetailsecff1c_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackTransactionDetailsecff1c_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackTransactionDetailsecff1c_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackTransactionDetailsecff1c_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackTransactionDetailsecff1c_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackTransactionDetailsecff1c_TimerEvent(this, event); };
	QDateTime dateDefault() { return _date; };
	void setDateDefault(QDateTime p) { if (p != _date) { _date = p; dateChanged(_date); } };
	qint32 statusDefault() { return _status; };
	void setStatusDefault(qint32 p) { if (p != _status) { _status = p; statusChanged(_status); } };
	qint32 typeDefault() { return _type; };
	void setTypeDefault(qint32 p) { if (p != _type) { _type = p; typeChanged(_type); } };
	QString amountDefault() { return _amount; };
	void setAmountDefault(QString p) { if (p != _amount) { _amount = p; amountChanged(_amount); } };
	QString hoursTraspassedDefault() { return _hoursTraspassed; };
	void setHoursTraspassedDefault(QString p) { if (p != _hoursTraspassed) { _hoursTraspassed = p; hoursTraspassedChanged(_hoursTraspassed); } };
	QString hoursBurnedDefault() { return _hoursBurned; };
	void setHoursBurnedDefault(QString p) { if (p != _hoursBurned) { _hoursBurned = p; hoursBurnedChanged(_hoursBurned); } };
	QString transactionIDDefault() { return _transactionID; };
	void setTransactionIDDefault(QString p) { if (p != _transactionID) { _transactionID = p; transactionIDChanged(_transactionID); } };
	AddressListfb4c44* addressesDefault() { return _addresses; };
	void setAddressesDefault(AddressListfb4c44* p) { if (p != _addresses) { _addresses = p; addressesChanged(_addresses); } };
	AddressListfb4c44* inputsDefault() { return _inputs; };
	void setInputsDefault(AddressListfb4c44* p) { if (p != _inputs) { _inputs = p; inputsChanged(_inputs); } };
	AddressListfb4c44* outputsDefault() { return _outputs; };
	void setOutputsDefault(AddressListfb4c44* p) { if (p != _outputs) { _outputs = p; outputsChanged(_outputs); } };
signals:
	void dateChanged(QDateTime date);
	void statusChanged(qint32 status);
	void typeChanged(qint32 ty);
	void amountChanged(QString amount);
	void hoursTraspassedChanged(QString hoursTraspassed);
	void hoursBurnedChanged(QString hoursBurned);
	void transactionIDChanged(QString transactionID);
	void addressesChanged(AddressListfb4c44* addresses);
	void inputsChanged(AddressListfb4c44* inputs);
	void outputsChanged(AddressListfb4c44* outputs);
public slots:
private:
	QDateTime _date;
	qint32 _status;
	qint32 _type;
	QString _amount;
	QString _hoursTraspassed;
	QString _hoursBurned;
	QString _transactionID;
	AddressListfb4c44* _addresses;
	AddressListfb4c44* _inputs;
	AddressListfb4c44* _outputs;
};

Q_DECLARE_METATYPE(TransactionDetailsecff1c*)


void TransactionDetailsecff1c_TransactionDetailsecff1c_QRegisterMetaTypes() {
	qRegisterMetaType<QDateTime>();
	qRegisterMetaType<QString>();
	qRegisterMetaType<AddressListfb4c44*>();
}

void* TransactionDetailsecff1c_Date(void* ptr)
{
	return new QDateTime(static_cast<TransactionDetailsecff1c*>(ptr)->date());
}

void* TransactionDetailsecff1c_DateDefault(void* ptr)
{
	return new QDateTime(static_cast<TransactionDetailsecff1c*>(ptr)->dateDefault());
}

void TransactionDetailsecff1c_SetDate(void* ptr, void* date)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->setDate(*static_cast<QDateTime*>(date));
}

void TransactionDetailsecff1c_SetDateDefault(void* ptr, void* date)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->setDateDefault(*static_cast<QDateTime*>(date));
}

void TransactionDetailsecff1c_ConnectDateChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(QDateTime)>(&TransactionDetailsecff1c::dateChanged), static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(QDateTime)>(&TransactionDetailsecff1c::Signal_DateChanged));
}

void TransactionDetailsecff1c_DisconnectDateChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(QDateTime)>(&TransactionDetailsecff1c::dateChanged), static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(QDateTime)>(&TransactionDetailsecff1c::Signal_DateChanged));
}

void TransactionDetailsecff1c_DateChanged(void* ptr, void* date)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->dateChanged(*static_cast<QDateTime*>(date));
}

int TransactionDetailsecff1c_Status(void* ptr)
{
	return static_cast<TransactionDetailsecff1c*>(ptr)->status();
}

int TransactionDetailsecff1c_StatusDefault(void* ptr)
{
	return static_cast<TransactionDetailsecff1c*>(ptr)->statusDefault();
}

void TransactionDetailsecff1c_SetStatus(void* ptr, int status)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->setStatus(status);
}

void TransactionDetailsecff1c_SetStatusDefault(void* ptr, int status)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->setStatusDefault(status);
}

void TransactionDetailsecff1c_ConnectStatusChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(qint32)>(&TransactionDetailsecff1c::statusChanged), static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(qint32)>(&TransactionDetailsecff1c::Signal_StatusChanged));
}

void TransactionDetailsecff1c_DisconnectStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(qint32)>(&TransactionDetailsecff1c::statusChanged), static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(qint32)>(&TransactionDetailsecff1c::Signal_StatusChanged));
}

void TransactionDetailsecff1c_StatusChanged(void* ptr, int status)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->statusChanged(status);
}

int TransactionDetailsecff1c_Type(void* ptr)
{
	return static_cast<TransactionDetailsecff1c*>(ptr)->type();
}

int TransactionDetailsecff1c_TypeDefault(void* ptr)
{
	return static_cast<TransactionDetailsecff1c*>(ptr)->typeDefault();
}

void TransactionDetailsecff1c_SetType(void* ptr, int ty)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->setType(ty);
}

void TransactionDetailsecff1c_SetTypeDefault(void* ptr, int ty)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->setTypeDefault(ty);
}

void TransactionDetailsecff1c_ConnectTypeChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(qint32)>(&TransactionDetailsecff1c::typeChanged), static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(qint32)>(&TransactionDetailsecff1c::Signal_TypeChanged));
}

void TransactionDetailsecff1c_DisconnectTypeChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(qint32)>(&TransactionDetailsecff1c::typeChanged), static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(qint32)>(&TransactionDetailsecff1c::Signal_TypeChanged));
}

void TransactionDetailsecff1c_TypeChanged(void* ptr, int ty)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->typeChanged(ty);
}

struct Moc_PackedString TransactionDetailsecff1c_Amount(void* ptr)
{
	return ({ QByteArray tc2b2b9 = static_cast<TransactionDetailsecff1c*>(ptr)->amount().toUtf8(); Moc_PackedString { const_cast<char*>(tc2b2b9.prepend("WHITESPACE").constData()+10), tc2b2b9.size()-10 }; });
}

struct Moc_PackedString TransactionDetailsecff1c_AmountDefault(void* ptr)
{
	return ({ QByteArray tb5913f = static_cast<TransactionDetailsecff1c*>(ptr)->amountDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tb5913f.prepend("WHITESPACE").constData()+10), tb5913f.size()-10 }; });
}

void TransactionDetailsecff1c_SetAmount(void* ptr, struct Moc_PackedString amount)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->setAmount(QString::fromUtf8(amount.data, amount.len));
}

void TransactionDetailsecff1c_SetAmountDefault(void* ptr, struct Moc_PackedString amount)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->setAmountDefault(QString::fromUtf8(amount.data, amount.len));
}

void TransactionDetailsecff1c_ConnectAmountChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(QString)>(&TransactionDetailsecff1c::amountChanged), static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(QString)>(&TransactionDetailsecff1c::Signal_AmountChanged));
}

void TransactionDetailsecff1c_DisconnectAmountChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(QString)>(&TransactionDetailsecff1c::amountChanged), static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(QString)>(&TransactionDetailsecff1c::Signal_AmountChanged));
}

void TransactionDetailsecff1c_AmountChanged(void* ptr, struct Moc_PackedString amount)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->amountChanged(QString::fromUtf8(amount.data, amount.len));
}

struct Moc_PackedString TransactionDetailsecff1c_HoursTraspassed(void* ptr)
{
	return ({ QByteArray ta82c91 = static_cast<TransactionDetailsecff1c*>(ptr)->hoursTraspassed().toUtf8(); Moc_PackedString { const_cast<char*>(ta82c91.prepend("WHITESPACE").constData()+10), ta82c91.size()-10 }; });
}

struct Moc_PackedString TransactionDetailsecff1c_HoursTraspassedDefault(void* ptr)
{
	return ({ QByteArray t98d51d = static_cast<TransactionDetailsecff1c*>(ptr)->hoursTraspassedDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t98d51d.prepend("WHITESPACE").constData()+10), t98d51d.size()-10 }; });
}

void TransactionDetailsecff1c_SetHoursTraspassed(void* ptr, struct Moc_PackedString hoursTraspassed)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->setHoursTraspassed(QString::fromUtf8(hoursTraspassed.data, hoursTraspassed.len));
}

void TransactionDetailsecff1c_SetHoursTraspassedDefault(void* ptr, struct Moc_PackedString hoursTraspassed)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->setHoursTraspassedDefault(QString::fromUtf8(hoursTraspassed.data, hoursTraspassed.len));
}

void TransactionDetailsecff1c_ConnectHoursTraspassedChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(QString)>(&TransactionDetailsecff1c::hoursTraspassedChanged), static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(QString)>(&TransactionDetailsecff1c::Signal_HoursTraspassedChanged));
}

void TransactionDetailsecff1c_DisconnectHoursTraspassedChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(QString)>(&TransactionDetailsecff1c::hoursTraspassedChanged), static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(QString)>(&TransactionDetailsecff1c::Signal_HoursTraspassedChanged));
}

void TransactionDetailsecff1c_HoursTraspassedChanged(void* ptr, struct Moc_PackedString hoursTraspassed)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->hoursTraspassedChanged(QString::fromUtf8(hoursTraspassed.data, hoursTraspassed.len));
}

struct Moc_PackedString TransactionDetailsecff1c_HoursBurned(void* ptr)
{
	return ({ QByteArray t880017 = static_cast<TransactionDetailsecff1c*>(ptr)->hoursBurned().toUtf8(); Moc_PackedString { const_cast<char*>(t880017.prepend("WHITESPACE").constData()+10), t880017.size()-10 }; });
}

struct Moc_PackedString TransactionDetailsecff1c_HoursBurnedDefault(void* ptr)
{
	return ({ QByteArray t7a73b3 = static_cast<TransactionDetailsecff1c*>(ptr)->hoursBurnedDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t7a73b3.prepend("WHITESPACE").constData()+10), t7a73b3.size()-10 }; });
}

void TransactionDetailsecff1c_SetHoursBurned(void* ptr, struct Moc_PackedString hoursBurned)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->setHoursBurned(QString::fromUtf8(hoursBurned.data, hoursBurned.len));
}

void TransactionDetailsecff1c_SetHoursBurnedDefault(void* ptr, struct Moc_PackedString hoursBurned)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->setHoursBurnedDefault(QString::fromUtf8(hoursBurned.data, hoursBurned.len));
}

void TransactionDetailsecff1c_ConnectHoursBurnedChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(QString)>(&TransactionDetailsecff1c::hoursBurnedChanged), static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(QString)>(&TransactionDetailsecff1c::Signal_HoursBurnedChanged));
}

void TransactionDetailsecff1c_DisconnectHoursBurnedChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(QString)>(&TransactionDetailsecff1c::hoursBurnedChanged), static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(QString)>(&TransactionDetailsecff1c::Signal_HoursBurnedChanged));
}

void TransactionDetailsecff1c_HoursBurnedChanged(void* ptr, struct Moc_PackedString hoursBurned)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->hoursBurnedChanged(QString::fromUtf8(hoursBurned.data, hoursBurned.len));
}

struct Moc_PackedString TransactionDetailsecff1c_TransactionID(void* ptr)
{
	return ({ QByteArray tc8c5df = static_cast<TransactionDetailsecff1c*>(ptr)->transactionID().toUtf8(); Moc_PackedString { const_cast<char*>(tc8c5df.prepend("WHITESPACE").constData()+10), tc8c5df.size()-10 }; });
}

struct Moc_PackedString TransactionDetailsecff1c_TransactionIDDefault(void* ptr)
{
	return ({ QByteArray t299bde = static_cast<TransactionDetailsecff1c*>(ptr)->transactionIDDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t299bde.prepend("WHITESPACE").constData()+10), t299bde.size()-10 }; });
}

void TransactionDetailsecff1c_SetTransactionID(void* ptr, struct Moc_PackedString transactionID)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->setTransactionID(QString::fromUtf8(transactionID.data, transactionID.len));
}

void TransactionDetailsecff1c_SetTransactionIDDefault(void* ptr, struct Moc_PackedString transactionID)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->setTransactionIDDefault(QString::fromUtf8(transactionID.data, transactionID.len));
}

void TransactionDetailsecff1c_ConnectTransactionIDChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(QString)>(&TransactionDetailsecff1c::transactionIDChanged), static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(QString)>(&TransactionDetailsecff1c::Signal_TransactionIDChanged));
}

void TransactionDetailsecff1c_DisconnectTransactionIDChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(QString)>(&TransactionDetailsecff1c::transactionIDChanged), static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(QString)>(&TransactionDetailsecff1c::Signal_TransactionIDChanged));
}

void TransactionDetailsecff1c_TransactionIDChanged(void* ptr, struct Moc_PackedString transactionID)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->transactionIDChanged(QString::fromUtf8(transactionID.data, transactionID.len));
}

void* TransactionDetailsecff1c_Addresses(void* ptr)
{
	return static_cast<TransactionDetailsecff1c*>(ptr)->addresses();
}

void* TransactionDetailsecff1c_AddressesDefault(void* ptr)
{
	return static_cast<TransactionDetailsecff1c*>(ptr)->addressesDefault();
}

void TransactionDetailsecff1c_SetAddresses(void* ptr, void* addresses)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->setAddresses(static_cast<AddressListfb4c44*>(addresses));
}

void TransactionDetailsecff1c_SetAddressesDefault(void* ptr, void* addresses)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->setAddressesDefault(static_cast<AddressListfb4c44*>(addresses));
}

void TransactionDetailsecff1c_ConnectAddressesChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(AddressListfb4c44*)>(&TransactionDetailsecff1c::addressesChanged), static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(AddressListfb4c44*)>(&TransactionDetailsecff1c::Signal_AddressesChanged));
}

void TransactionDetailsecff1c_DisconnectAddressesChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(AddressListfb4c44*)>(&TransactionDetailsecff1c::addressesChanged), static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(AddressListfb4c44*)>(&TransactionDetailsecff1c::Signal_AddressesChanged));
}

void TransactionDetailsecff1c_AddressesChanged(void* ptr, void* addresses)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->addressesChanged(static_cast<AddressListfb4c44*>(addresses));
}

void* TransactionDetailsecff1c_Inputs(void* ptr)
{
	return static_cast<TransactionDetailsecff1c*>(ptr)->inputs();
}

void* TransactionDetailsecff1c_InputsDefault(void* ptr)
{
	return static_cast<TransactionDetailsecff1c*>(ptr)->inputsDefault();
}

void TransactionDetailsecff1c_SetInputs(void* ptr, void* inputs)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->setInputs(static_cast<AddressListfb4c44*>(inputs));
}

void TransactionDetailsecff1c_SetInputsDefault(void* ptr, void* inputs)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->setInputsDefault(static_cast<AddressListfb4c44*>(inputs));
}

void TransactionDetailsecff1c_ConnectInputsChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(AddressListfb4c44*)>(&TransactionDetailsecff1c::inputsChanged), static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(AddressListfb4c44*)>(&TransactionDetailsecff1c::Signal_InputsChanged));
}

void TransactionDetailsecff1c_DisconnectInputsChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(AddressListfb4c44*)>(&TransactionDetailsecff1c::inputsChanged), static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(AddressListfb4c44*)>(&TransactionDetailsecff1c::Signal_InputsChanged));
}

void TransactionDetailsecff1c_InputsChanged(void* ptr, void* inputs)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->inputsChanged(static_cast<AddressListfb4c44*>(inputs));
}

void* TransactionDetailsecff1c_Outputs(void* ptr)
{
	return static_cast<TransactionDetailsecff1c*>(ptr)->outputs();
}

void* TransactionDetailsecff1c_OutputsDefault(void* ptr)
{
	return static_cast<TransactionDetailsecff1c*>(ptr)->outputsDefault();
}

void TransactionDetailsecff1c_SetOutputs(void* ptr, void* outputs)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->setOutputs(static_cast<AddressListfb4c44*>(outputs));
}

void TransactionDetailsecff1c_SetOutputsDefault(void* ptr, void* outputs)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->setOutputsDefault(static_cast<AddressListfb4c44*>(outputs));
}

void TransactionDetailsecff1c_ConnectOutputsChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(AddressListfb4c44*)>(&TransactionDetailsecff1c::outputsChanged), static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(AddressListfb4c44*)>(&TransactionDetailsecff1c::Signal_OutputsChanged));
}

void TransactionDetailsecff1c_DisconnectOutputsChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(AddressListfb4c44*)>(&TransactionDetailsecff1c::outputsChanged), static_cast<TransactionDetailsecff1c*>(ptr), static_cast<void (TransactionDetailsecff1c::*)(AddressListfb4c44*)>(&TransactionDetailsecff1c::Signal_OutputsChanged));
}

void TransactionDetailsecff1c_OutputsChanged(void* ptr, void* outputs)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->outputsChanged(static_cast<AddressListfb4c44*>(outputs));
}

int TransactionDetailsecff1c_TransactionDetailsecff1c_QRegisterMetaType()
{
	return qRegisterMetaType<TransactionDetailsecff1c*>();
}

int TransactionDetailsecff1c_TransactionDetailsecff1c_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<TransactionDetailsecff1c*>(const_cast<const char*>(typeName));
}

int TransactionDetailsecff1c_TransactionDetailsecff1c_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TransactionDetailsecff1c>();
#else
	return 0;
#endif
}

int TransactionDetailsecff1c_TransactionDetailsecff1c_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TransactionDetailsecff1c>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* TransactionDetailsecff1c___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionDetailsecff1c___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TransactionDetailsecff1c___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* TransactionDetailsecff1c___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TransactionDetailsecff1c___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* TransactionDetailsecff1c___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* TransactionDetailsecff1c___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionDetailsecff1c___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TransactionDetailsecff1c___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TransactionDetailsecff1c___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionDetailsecff1c___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TransactionDetailsecff1c___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TransactionDetailsecff1c___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionDetailsecff1c___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TransactionDetailsecff1c___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TransactionDetailsecff1c_NewTransactionDetails(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new TransactionDetailsecff1c(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new TransactionDetailsecff1c(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new TransactionDetailsecff1c(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new TransactionDetailsecff1c(static_cast<QWindow*>(parent));
	} else {
		return new TransactionDetailsecff1c(static_cast<QObject*>(parent));
	}
}

void TransactionDetailsecff1c_DestroyTransactionDetails(void* ptr)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->~TransactionDetailsecff1c();
}

void TransactionDetailsecff1c_DestroyTransactionDetailsDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void TransactionDetailsecff1c_ChildEventDefault(void* ptr, void* event)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void TransactionDetailsecff1c_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void TransactionDetailsecff1c_CustomEventDefault(void* ptr, void* event)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void TransactionDetailsecff1c_DeleteLaterDefault(void* ptr)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->QObject::deleteLater();
}

void TransactionDetailsecff1c_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char TransactionDetailsecff1c_EventDefault(void* ptr, void* e)
{
	return static_cast<TransactionDetailsecff1c*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char TransactionDetailsecff1c_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<TransactionDetailsecff1c*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void TransactionDetailsecff1c_TimerEventDefault(void* ptr, void* event)
{
	static_cast<TransactionDetailsecff1c*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
