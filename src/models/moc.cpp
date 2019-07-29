

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

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


class BlockchainStatusModel0d0160: public QObject
{
Q_OBJECT
Q_PROPERTY(QString numberOfBlocks READ numberOfBlocks WRITE setNumberOfBlocks NOTIFY numberOfBlocksChanged)
Q_PROPERTY(QDateTime timestampLastBlock READ timestampLastBlock WRITE setTimestampLastBlock NOTIFY timestampLastBlockChanged)
Q_PROPERTY(QString hashLastBlock READ hashLastBlock WRITE setHashLastBlock NOTIFY hashLastBlockChanged)
Q_PROPERTY(QString currentSkySupply READ currentSkySupply WRITE setCurrentSkySupply NOTIFY currentSkySupplyChanged)
Q_PROPERTY(QString totalSkySupply READ totalSkySupply WRITE setTotalSkySupply NOTIFY totalSkySupplyChanged)
Q_PROPERTY(QString currentCoinHoursSupply READ currentCoinHoursSupply WRITE setCurrentCoinHoursSupply NOTIFY currentCoinHoursSupplyChanged)
Q_PROPERTY(QString totalCoinHoursSupply READ totalCoinHoursSupply WRITE setTotalCoinHoursSupply NOTIFY totalCoinHoursSupplyChanged)
public:
	BlockchainStatusModel0d0160(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");BlockchainStatusModel0d0160_BlockchainStatusModel0d0160_QRegisterMetaType();BlockchainStatusModel0d0160_BlockchainStatusModel0d0160_QRegisterMetaTypes();callbackBlockchainStatusModel0d0160_Constructor(this);};
	void Signal_Update() { callbackBlockchainStatusModel0d0160_Update(this); };
	QString numberOfBlocks() { return ({ Moc_PackedString tempVal = callbackBlockchainStatusModel0d0160_NumberOfBlocks(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setNumberOfBlocks(QString numberOfBlocks) { QByteArray tf3f052 = numberOfBlocks.toUtf8(); Moc_PackedString numberOfBlocksPacked = { const_cast<char*>(tf3f052.prepend("WHITESPACE").constData()+10), tf3f052.size()-10 };callbackBlockchainStatusModel0d0160_SetNumberOfBlocks(this, numberOfBlocksPacked); };
	void Signal_NumberOfBlocksChanged(QString numberOfBlocks) { QByteArray tf3f052 = numberOfBlocks.toUtf8(); Moc_PackedString numberOfBlocksPacked = { const_cast<char*>(tf3f052.prepend("WHITESPACE").constData()+10), tf3f052.size()-10 };callbackBlockchainStatusModel0d0160_NumberOfBlocksChanged(this, numberOfBlocksPacked); };
	QDateTime timestampLastBlock() { return *static_cast<QDateTime*>(callbackBlockchainStatusModel0d0160_TimestampLastBlock(this)); };
	void setTimestampLastBlock(QDateTime timestampLastBlock) { callbackBlockchainStatusModel0d0160_SetTimestampLastBlock(this, new QDateTime(timestampLastBlock)); };
	void Signal_TimestampLastBlockChanged(QDateTime timestampLastBlock) { callbackBlockchainStatusModel0d0160_TimestampLastBlockChanged(this, new QDateTime(timestampLastBlock)); };
	QString hashLastBlock() { return ({ Moc_PackedString tempVal = callbackBlockchainStatusModel0d0160_HashLastBlock(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setHashLastBlock(QString hashLastBlock) { QByteArray t33501e = hashLastBlock.toUtf8(); Moc_PackedString hashLastBlockPacked = { const_cast<char*>(t33501e.prepend("WHITESPACE").constData()+10), t33501e.size()-10 };callbackBlockchainStatusModel0d0160_SetHashLastBlock(this, hashLastBlockPacked); };
	void Signal_HashLastBlockChanged(QString hashLastBlock) { QByteArray t33501e = hashLastBlock.toUtf8(); Moc_PackedString hashLastBlockPacked = { const_cast<char*>(t33501e.prepend("WHITESPACE").constData()+10), t33501e.size()-10 };callbackBlockchainStatusModel0d0160_HashLastBlockChanged(this, hashLastBlockPacked); };
	QString currentSkySupply() { return ({ Moc_PackedString tempVal = callbackBlockchainStatusModel0d0160_CurrentSkySupply(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setCurrentSkySupply(QString currentSkySupply) { QByteArray t9cbeff = currentSkySupply.toUtf8(); Moc_PackedString currentSkySupplyPacked = { const_cast<char*>(t9cbeff.prepend("WHITESPACE").constData()+10), t9cbeff.size()-10 };callbackBlockchainStatusModel0d0160_SetCurrentSkySupply(this, currentSkySupplyPacked); };
	void Signal_CurrentSkySupplyChanged(QString currentSkySupply) { QByteArray t9cbeff = currentSkySupply.toUtf8(); Moc_PackedString currentSkySupplyPacked = { const_cast<char*>(t9cbeff.prepend("WHITESPACE").constData()+10), t9cbeff.size()-10 };callbackBlockchainStatusModel0d0160_CurrentSkySupplyChanged(this, currentSkySupplyPacked); };
	QString totalSkySupply() { return ({ Moc_PackedString tempVal = callbackBlockchainStatusModel0d0160_TotalSkySupply(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setTotalSkySupply(QString totalSkySupply) { QByteArray tf96a84 = totalSkySupply.toUtf8(); Moc_PackedString totalSkySupplyPacked = { const_cast<char*>(tf96a84.prepend("WHITESPACE").constData()+10), tf96a84.size()-10 };callbackBlockchainStatusModel0d0160_SetTotalSkySupply(this, totalSkySupplyPacked); };
	void Signal_TotalSkySupplyChanged(QString totalSkySupply) { QByteArray tf96a84 = totalSkySupply.toUtf8(); Moc_PackedString totalSkySupplyPacked = { const_cast<char*>(tf96a84.prepend("WHITESPACE").constData()+10), tf96a84.size()-10 };callbackBlockchainStatusModel0d0160_TotalSkySupplyChanged(this, totalSkySupplyPacked); };
	QString currentCoinHoursSupply() { return ({ Moc_PackedString tempVal = callbackBlockchainStatusModel0d0160_CurrentCoinHoursSupply(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setCurrentCoinHoursSupply(QString currentCoinHoursSupply) { QByteArray t88f3d1 = currentCoinHoursSupply.toUtf8(); Moc_PackedString currentCoinHoursSupplyPacked = { const_cast<char*>(t88f3d1.prepend("WHITESPACE").constData()+10), t88f3d1.size()-10 };callbackBlockchainStatusModel0d0160_SetCurrentCoinHoursSupply(this, currentCoinHoursSupplyPacked); };
	void Signal_CurrentCoinHoursSupplyChanged(QString currentCoinHoursSupply) { QByteArray t88f3d1 = currentCoinHoursSupply.toUtf8(); Moc_PackedString currentCoinHoursSupplyPacked = { const_cast<char*>(t88f3d1.prepend("WHITESPACE").constData()+10), t88f3d1.size()-10 };callbackBlockchainStatusModel0d0160_CurrentCoinHoursSupplyChanged(this, currentCoinHoursSupplyPacked); };
	QString totalCoinHoursSupply() { return ({ Moc_PackedString tempVal = callbackBlockchainStatusModel0d0160_TotalCoinHoursSupply(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setTotalCoinHoursSupply(QString totalCoinHoursSupply) { QByteArray t1cfb46 = totalCoinHoursSupply.toUtf8(); Moc_PackedString totalCoinHoursSupplyPacked = { const_cast<char*>(t1cfb46.prepend("WHITESPACE").constData()+10), t1cfb46.size()-10 };callbackBlockchainStatusModel0d0160_SetTotalCoinHoursSupply(this, totalCoinHoursSupplyPacked); };
	void Signal_TotalCoinHoursSupplyChanged(QString totalCoinHoursSupply) { QByteArray t1cfb46 = totalCoinHoursSupply.toUtf8(); Moc_PackedString totalCoinHoursSupplyPacked = { const_cast<char*>(t1cfb46.prepend("WHITESPACE").constData()+10), t1cfb46.size()-10 };callbackBlockchainStatusModel0d0160_TotalCoinHoursSupplyChanged(this, totalCoinHoursSupplyPacked); };
	 ~BlockchainStatusModel0d0160() { callbackBlockchainStatusModel0d0160_DestroyBlockchainStatusModel(this); };
	void childEvent(QChildEvent * event) { callbackBlockchainStatusModel0d0160_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackBlockchainStatusModel0d0160_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackBlockchainStatusModel0d0160_CustomEvent(this, event); };
	void deleteLater() { callbackBlockchainStatusModel0d0160_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackBlockchainStatusModel0d0160_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackBlockchainStatusModel0d0160_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackBlockchainStatusModel0d0160_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackBlockchainStatusModel0d0160_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackBlockchainStatusModel0d0160_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackBlockchainStatusModel0d0160_TimerEvent(this, event); };
	QString numberOfBlocksDefault() { return _numberOfBlocks; };
	void setNumberOfBlocksDefault(QString p) { if (p != _numberOfBlocks) { _numberOfBlocks = p; numberOfBlocksChanged(_numberOfBlocks); } };
	QDateTime timestampLastBlockDefault() { return _timestampLastBlock; };
	void setTimestampLastBlockDefault(QDateTime p) { if (p != _timestampLastBlock) { _timestampLastBlock = p; timestampLastBlockChanged(_timestampLastBlock); } };
	QString hashLastBlockDefault() { return _hashLastBlock; };
	void setHashLastBlockDefault(QString p) { if (p != _hashLastBlock) { _hashLastBlock = p; hashLastBlockChanged(_hashLastBlock); } };
	QString currentSkySupplyDefault() { return _currentSkySupply; };
	void setCurrentSkySupplyDefault(QString p) { if (p != _currentSkySupply) { _currentSkySupply = p; currentSkySupplyChanged(_currentSkySupply); } };
	QString totalSkySupplyDefault() { return _totalSkySupply; };
	void setTotalSkySupplyDefault(QString p) { if (p != _totalSkySupply) { _totalSkySupply = p; totalSkySupplyChanged(_totalSkySupply); } };
	QString currentCoinHoursSupplyDefault() { return _currentCoinHoursSupply; };
	void setCurrentCoinHoursSupplyDefault(QString p) { if (p != _currentCoinHoursSupply) { _currentCoinHoursSupply = p; currentCoinHoursSupplyChanged(_currentCoinHoursSupply); } };
	QString totalCoinHoursSupplyDefault() { return _totalCoinHoursSupply; };
	void setTotalCoinHoursSupplyDefault(QString p) { if (p != _totalCoinHoursSupply) { _totalCoinHoursSupply = p; totalCoinHoursSupplyChanged(_totalCoinHoursSupply); } };
signals:
	void update();
	void numberOfBlocksChanged(QString numberOfBlocks);
	void timestampLastBlockChanged(QDateTime timestampLastBlock);
	void hashLastBlockChanged(QString hashLastBlock);
	void currentSkySupplyChanged(QString currentSkySupply);
	void totalSkySupplyChanged(QString totalSkySupply);
	void currentCoinHoursSupplyChanged(QString currentCoinHoursSupply);
	void totalCoinHoursSupplyChanged(QString totalCoinHoursSupply);
public slots:
private:
	QString _numberOfBlocks;
	QDateTime _timestampLastBlock;
	QString _hashLastBlock;
	QString _currentSkySupply;
	QString _totalSkySupply;
	QString _currentCoinHoursSupply;
	QString _totalCoinHoursSupply;
};

Q_DECLARE_METATYPE(BlockchainStatusModel0d0160*)


void BlockchainStatusModel0d0160_BlockchainStatusModel0d0160_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
	qRegisterMetaType<QDateTime>();
}

void BlockchainStatusModel0d0160_ConnectUpdate(void* ptr)
{
	QObject::connect(static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)()>(&BlockchainStatusModel0d0160::update), static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)()>(&BlockchainStatusModel0d0160::Signal_Update));
}

void BlockchainStatusModel0d0160_DisconnectUpdate(void* ptr)
{
	QObject::disconnect(static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)()>(&BlockchainStatusModel0d0160::update), static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)()>(&BlockchainStatusModel0d0160::Signal_Update));
}

void BlockchainStatusModel0d0160_Update(void* ptr)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->update();
}

struct Moc_PackedString BlockchainStatusModel0d0160_NumberOfBlocks(void* ptr)
{
	return ({ QByteArray t90e1b1 = static_cast<BlockchainStatusModel0d0160*>(ptr)->numberOfBlocks().toUtf8(); Moc_PackedString { const_cast<char*>(t90e1b1.prepend("WHITESPACE").constData()+10), t90e1b1.size()-10 }; });
}

struct Moc_PackedString BlockchainStatusModel0d0160_NumberOfBlocksDefault(void* ptr)
{
	return ({ QByteArray td41e94 = static_cast<BlockchainStatusModel0d0160*>(ptr)->numberOfBlocksDefault().toUtf8(); Moc_PackedString { const_cast<char*>(td41e94.prepend("WHITESPACE").constData()+10), td41e94.size()-10 }; });
}

void BlockchainStatusModel0d0160_SetNumberOfBlocks(void* ptr, struct Moc_PackedString numberOfBlocks)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->setNumberOfBlocks(QString::fromUtf8(numberOfBlocks.data, numberOfBlocks.len));
}

void BlockchainStatusModel0d0160_SetNumberOfBlocksDefault(void* ptr, struct Moc_PackedString numberOfBlocks)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->setNumberOfBlocksDefault(QString::fromUtf8(numberOfBlocks.data, numberOfBlocks.len));
}

void BlockchainStatusModel0d0160_ConnectNumberOfBlocksChanged(void* ptr)
{
	QObject::connect(static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::numberOfBlocksChanged), static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::Signal_NumberOfBlocksChanged));
}

void BlockchainStatusModel0d0160_DisconnectNumberOfBlocksChanged(void* ptr)
{
	QObject::disconnect(static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::numberOfBlocksChanged), static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::Signal_NumberOfBlocksChanged));
}

void BlockchainStatusModel0d0160_NumberOfBlocksChanged(void* ptr, struct Moc_PackedString numberOfBlocks)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->numberOfBlocksChanged(QString::fromUtf8(numberOfBlocks.data, numberOfBlocks.len));
}

void* BlockchainStatusModel0d0160_TimestampLastBlock(void* ptr)
{
	return new QDateTime(static_cast<BlockchainStatusModel0d0160*>(ptr)->timestampLastBlock());
}

void* BlockchainStatusModel0d0160_TimestampLastBlockDefault(void* ptr)
{
	return new QDateTime(static_cast<BlockchainStatusModel0d0160*>(ptr)->timestampLastBlockDefault());
}

void BlockchainStatusModel0d0160_SetTimestampLastBlock(void* ptr, void* timestampLastBlock)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->setTimestampLastBlock(*static_cast<QDateTime*>(timestampLastBlock));
}

void BlockchainStatusModel0d0160_SetTimestampLastBlockDefault(void* ptr, void* timestampLastBlock)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->setTimestampLastBlockDefault(*static_cast<QDateTime*>(timestampLastBlock));
}

void BlockchainStatusModel0d0160_ConnectTimestampLastBlockChanged(void* ptr)
{
	QObject::connect(static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QDateTime)>(&BlockchainStatusModel0d0160::timestampLastBlockChanged), static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QDateTime)>(&BlockchainStatusModel0d0160::Signal_TimestampLastBlockChanged));
}

void BlockchainStatusModel0d0160_DisconnectTimestampLastBlockChanged(void* ptr)
{
	QObject::disconnect(static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QDateTime)>(&BlockchainStatusModel0d0160::timestampLastBlockChanged), static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QDateTime)>(&BlockchainStatusModel0d0160::Signal_TimestampLastBlockChanged));
}

void BlockchainStatusModel0d0160_TimestampLastBlockChanged(void* ptr, void* timestampLastBlock)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->timestampLastBlockChanged(*static_cast<QDateTime*>(timestampLastBlock));
}

struct Moc_PackedString BlockchainStatusModel0d0160_HashLastBlock(void* ptr)
{
	return ({ QByteArray tbaa67e = static_cast<BlockchainStatusModel0d0160*>(ptr)->hashLastBlock().toUtf8(); Moc_PackedString { const_cast<char*>(tbaa67e.prepend("WHITESPACE").constData()+10), tbaa67e.size()-10 }; });
}

struct Moc_PackedString BlockchainStatusModel0d0160_HashLastBlockDefault(void* ptr)
{
	return ({ QByteArray td6ba2e = static_cast<BlockchainStatusModel0d0160*>(ptr)->hashLastBlockDefault().toUtf8(); Moc_PackedString { const_cast<char*>(td6ba2e.prepend("WHITESPACE").constData()+10), td6ba2e.size()-10 }; });
}

void BlockchainStatusModel0d0160_SetHashLastBlock(void* ptr, struct Moc_PackedString hashLastBlock)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->setHashLastBlock(QString::fromUtf8(hashLastBlock.data, hashLastBlock.len));
}

void BlockchainStatusModel0d0160_SetHashLastBlockDefault(void* ptr, struct Moc_PackedString hashLastBlock)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->setHashLastBlockDefault(QString::fromUtf8(hashLastBlock.data, hashLastBlock.len));
}

void BlockchainStatusModel0d0160_ConnectHashLastBlockChanged(void* ptr)
{
	QObject::connect(static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::hashLastBlockChanged), static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::Signal_HashLastBlockChanged));
}

void BlockchainStatusModel0d0160_DisconnectHashLastBlockChanged(void* ptr)
{
	QObject::disconnect(static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::hashLastBlockChanged), static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::Signal_HashLastBlockChanged));
}

void BlockchainStatusModel0d0160_HashLastBlockChanged(void* ptr, struct Moc_PackedString hashLastBlock)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->hashLastBlockChanged(QString::fromUtf8(hashLastBlock.data, hashLastBlock.len));
}

struct Moc_PackedString BlockchainStatusModel0d0160_CurrentSkySupply(void* ptr)
{
	return ({ QByteArray ta7721d = static_cast<BlockchainStatusModel0d0160*>(ptr)->currentSkySupply().toUtf8(); Moc_PackedString { const_cast<char*>(ta7721d.prepend("WHITESPACE").constData()+10), ta7721d.size()-10 }; });
}

struct Moc_PackedString BlockchainStatusModel0d0160_CurrentSkySupplyDefault(void* ptr)
{
	return ({ QByteArray ta836bf = static_cast<BlockchainStatusModel0d0160*>(ptr)->currentSkySupplyDefault().toUtf8(); Moc_PackedString { const_cast<char*>(ta836bf.prepend("WHITESPACE").constData()+10), ta836bf.size()-10 }; });
}

void BlockchainStatusModel0d0160_SetCurrentSkySupply(void* ptr, struct Moc_PackedString currentSkySupply)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->setCurrentSkySupply(QString::fromUtf8(currentSkySupply.data, currentSkySupply.len));
}

void BlockchainStatusModel0d0160_SetCurrentSkySupplyDefault(void* ptr, struct Moc_PackedString currentSkySupply)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->setCurrentSkySupplyDefault(QString::fromUtf8(currentSkySupply.data, currentSkySupply.len));
}

void BlockchainStatusModel0d0160_ConnectCurrentSkySupplyChanged(void* ptr)
{
	QObject::connect(static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::currentSkySupplyChanged), static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::Signal_CurrentSkySupplyChanged));
}

void BlockchainStatusModel0d0160_DisconnectCurrentSkySupplyChanged(void* ptr)
{
	QObject::disconnect(static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::currentSkySupplyChanged), static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::Signal_CurrentSkySupplyChanged));
}

void BlockchainStatusModel0d0160_CurrentSkySupplyChanged(void* ptr, struct Moc_PackedString currentSkySupply)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->currentSkySupplyChanged(QString::fromUtf8(currentSkySupply.data, currentSkySupply.len));
}

struct Moc_PackedString BlockchainStatusModel0d0160_TotalSkySupply(void* ptr)
{
	return ({ QByteArray t35020f = static_cast<BlockchainStatusModel0d0160*>(ptr)->totalSkySupply().toUtf8(); Moc_PackedString { const_cast<char*>(t35020f.prepend("WHITESPACE").constData()+10), t35020f.size()-10 }; });
}

struct Moc_PackedString BlockchainStatusModel0d0160_TotalSkySupplyDefault(void* ptr)
{
	return ({ QByteArray tec951b = static_cast<BlockchainStatusModel0d0160*>(ptr)->totalSkySupplyDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tec951b.prepend("WHITESPACE").constData()+10), tec951b.size()-10 }; });
}

void BlockchainStatusModel0d0160_SetTotalSkySupply(void* ptr, struct Moc_PackedString totalSkySupply)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->setTotalSkySupply(QString::fromUtf8(totalSkySupply.data, totalSkySupply.len));
}

void BlockchainStatusModel0d0160_SetTotalSkySupplyDefault(void* ptr, struct Moc_PackedString totalSkySupply)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->setTotalSkySupplyDefault(QString::fromUtf8(totalSkySupply.data, totalSkySupply.len));
}

void BlockchainStatusModel0d0160_ConnectTotalSkySupplyChanged(void* ptr)
{
	QObject::connect(static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::totalSkySupplyChanged), static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::Signal_TotalSkySupplyChanged));
}

void BlockchainStatusModel0d0160_DisconnectTotalSkySupplyChanged(void* ptr)
{
	QObject::disconnect(static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::totalSkySupplyChanged), static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::Signal_TotalSkySupplyChanged));
}

void BlockchainStatusModel0d0160_TotalSkySupplyChanged(void* ptr, struct Moc_PackedString totalSkySupply)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->totalSkySupplyChanged(QString::fromUtf8(totalSkySupply.data, totalSkySupply.len));
}

struct Moc_PackedString BlockchainStatusModel0d0160_CurrentCoinHoursSupply(void* ptr)
{
	return ({ QByteArray t848392 = static_cast<BlockchainStatusModel0d0160*>(ptr)->currentCoinHoursSupply().toUtf8(); Moc_PackedString { const_cast<char*>(t848392.prepend("WHITESPACE").constData()+10), t848392.size()-10 }; });
}

struct Moc_PackedString BlockchainStatusModel0d0160_CurrentCoinHoursSupplyDefault(void* ptr)
{
	return ({ QByteArray t627637 = static_cast<BlockchainStatusModel0d0160*>(ptr)->currentCoinHoursSupplyDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t627637.prepend("WHITESPACE").constData()+10), t627637.size()-10 }; });
}

void BlockchainStatusModel0d0160_SetCurrentCoinHoursSupply(void* ptr, struct Moc_PackedString currentCoinHoursSupply)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->setCurrentCoinHoursSupply(QString::fromUtf8(currentCoinHoursSupply.data, currentCoinHoursSupply.len));
}

void BlockchainStatusModel0d0160_SetCurrentCoinHoursSupplyDefault(void* ptr, struct Moc_PackedString currentCoinHoursSupply)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->setCurrentCoinHoursSupplyDefault(QString::fromUtf8(currentCoinHoursSupply.data, currentCoinHoursSupply.len));
}

void BlockchainStatusModel0d0160_ConnectCurrentCoinHoursSupplyChanged(void* ptr)
{
	QObject::connect(static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::currentCoinHoursSupplyChanged), static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::Signal_CurrentCoinHoursSupplyChanged));
}

void BlockchainStatusModel0d0160_DisconnectCurrentCoinHoursSupplyChanged(void* ptr)
{
	QObject::disconnect(static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::currentCoinHoursSupplyChanged), static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::Signal_CurrentCoinHoursSupplyChanged));
}

void BlockchainStatusModel0d0160_CurrentCoinHoursSupplyChanged(void* ptr, struct Moc_PackedString currentCoinHoursSupply)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->currentCoinHoursSupplyChanged(QString::fromUtf8(currentCoinHoursSupply.data, currentCoinHoursSupply.len));
}

struct Moc_PackedString BlockchainStatusModel0d0160_TotalCoinHoursSupply(void* ptr)
{
	return ({ QByteArray t379ab0 = static_cast<BlockchainStatusModel0d0160*>(ptr)->totalCoinHoursSupply().toUtf8(); Moc_PackedString { const_cast<char*>(t379ab0.prepend("WHITESPACE").constData()+10), t379ab0.size()-10 }; });
}

struct Moc_PackedString BlockchainStatusModel0d0160_TotalCoinHoursSupplyDefault(void* ptr)
{
	return ({ QByteArray t08aab2 = static_cast<BlockchainStatusModel0d0160*>(ptr)->totalCoinHoursSupplyDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t08aab2.prepend("WHITESPACE").constData()+10), t08aab2.size()-10 }; });
}

void BlockchainStatusModel0d0160_SetTotalCoinHoursSupply(void* ptr, struct Moc_PackedString totalCoinHoursSupply)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->setTotalCoinHoursSupply(QString::fromUtf8(totalCoinHoursSupply.data, totalCoinHoursSupply.len));
}

void BlockchainStatusModel0d0160_SetTotalCoinHoursSupplyDefault(void* ptr, struct Moc_PackedString totalCoinHoursSupply)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->setTotalCoinHoursSupplyDefault(QString::fromUtf8(totalCoinHoursSupply.data, totalCoinHoursSupply.len));
}

void BlockchainStatusModel0d0160_ConnectTotalCoinHoursSupplyChanged(void* ptr)
{
	QObject::connect(static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::totalCoinHoursSupplyChanged), static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::Signal_TotalCoinHoursSupplyChanged));
}

void BlockchainStatusModel0d0160_DisconnectTotalCoinHoursSupplyChanged(void* ptr)
{
	QObject::disconnect(static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::totalCoinHoursSupplyChanged), static_cast<BlockchainStatusModel0d0160*>(ptr), static_cast<void (BlockchainStatusModel0d0160::*)(QString)>(&BlockchainStatusModel0d0160::Signal_TotalCoinHoursSupplyChanged));
}

void BlockchainStatusModel0d0160_TotalCoinHoursSupplyChanged(void* ptr, struct Moc_PackedString totalCoinHoursSupply)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->totalCoinHoursSupplyChanged(QString::fromUtf8(totalCoinHoursSupply.data, totalCoinHoursSupply.len));
}

int BlockchainStatusModel0d0160_BlockchainStatusModel0d0160_QRegisterMetaType()
{
	return qRegisterMetaType<BlockchainStatusModel0d0160*>();
}

int BlockchainStatusModel0d0160_BlockchainStatusModel0d0160_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<BlockchainStatusModel0d0160*>(const_cast<const char*>(typeName));
}

int BlockchainStatusModel0d0160_BlockchainStatusModel0d0160_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<BlockchainStatusModel0d0160>();
#else
	return 0;
#endif
}

int BlockchainStatusModel0d0160_BlockchainStatusModel0d0160_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<BlockchainStatusModel0d0160>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* BlockchainStatusModel0d0160___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void BlockchainStatusModel0d0160___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* BlockchainStatusModel0d0160___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* BlockchainStatusModel0d0160___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void BlockchainStatusModel0d0160___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* BlockchainStatusModel0d0160___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* BlockchainStatusModel0d0160___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void BlockchainStatusModel0d0160___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* BlockchainStatusModel0d0160___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* BlockchainStatusModel0d0160___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void BlockchainStatusModel0d0160___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* BlockchainStatusModel0d0160___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* BlockchainStatusModel0d0160___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void BlockchainStatusModel0d0160___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* BlockchainStatusModel0d0160___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* BlockchainStatusModel0d0160_NewBlockchainStatusModel(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new BlockchainStatusModel0d0160(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new BlockchainStatusModel0d0160(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new BlockchainStatusModel0d0160(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new BlockchainStatusModel0d0160(static_cast<QWindow*>(parent));
	} else {
		return new BlockchainStatusModel0d0160(static_cast<QObject*>(parent));
	}
}

void BlockchainStatusModel0d0160_DestroyBlockchainStatusModel(void* ptr)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->~BlockchainStatusModel0d0160();
}

void BlockchainStatusModel0d0160_DestroyBlockchainStatusModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void BlockchainStatusModel0d0160_ChildEventDefault(void* ptr, void* event)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void BlockchainStatusModel0d0160_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void BlockchainStatusModel0d0160_CustomEventDefault(void* ptr, void* event)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void BlockchainStatusModel0d0160_DeleteLaterDefault(void* ptr)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->QObject::deleteLater();
}

void BlockchainStatusModel0d0160_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char BlockchainStatusModel0d0160_EventDefault(void* ptr, void* e)
{
	return static_cast<BlockchainStatusModel0d0160*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char BlockchainStatusModel0d0160_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<BlockchainStatusModel0d0160*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void BlockchainStatusModel0d0160_TimerEventDefault(void* ptr, void* event)
{
	static_cast<BlockchainStatusModel0d0160*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
