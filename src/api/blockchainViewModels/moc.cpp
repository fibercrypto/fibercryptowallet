

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


class BlockchainStatusModel97d618: public QObject
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
	BlockchainStatusModel97d618(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");BlockchainStatusModel97d618_BlockchainStatusModel97d618_QRegisterMetaType();BlockchainStatusModel97d618_BlockchainStatusModel97d618_QRegisterMetaTypes();callbackBlockchainStatusModel97d618_Constructor(this);};
	QString numberOfBlocks() { return ({ Moc_PackedString tempVal = callbackBlockchainStatusModel97d618_NumberOfBlocks(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setNumberOfBlocks(QString numberOfBlocks) { QByteArray tf3f052 = numberOfBlocks.toUtf8(); Moc_PackedString numberOfBlocksPacked = { const_cast<char*>(tf3f052.prepend("WHITESPACE").constData()+10), tf3f052.size()-10 };callbackBlockchainStatusModel97d618_SetNumberOfBlocks(this, numberOfBlocksPacked); };
	void Signal_NumberOfBlocksChanged(QString numberOfBlocks) { QByteArray tf3f052 = numberOfBlocks.toUtf8(); Moc_PackedString numberOfBlocksPacked = { const_cast<char*>(tf3f052.prepend("WHITESPACE").constData()+10), tf3f052.size()-10 };callbackBlockchainStatusModel97d618_NumberOfBlocksChanged(this, numberOfBlocksPacked); };
	QDateTime timestampLastBlock() { return *static_cast<QDateTime*>(callbackBlockchainStatusModel97d618_TimestampLastBlock(this)); };
	void setTimestampLastBlock(QDateTime timestampLastBlock) { callbackBlockchainStatusModel97d618_SetTimestampLastBlock(this, new QDateTime(timestampLastBlock)); };
	void Signal_TimestampLastBlockChanged(QDateTime timestampLastBlock) { callbackBlockchainStatusModel97d618_TimestampLastBlockChanged(this, new QDateTime(timestampLastBlock)); };
	QString hashLastBlock() { return ({ Moc_PackedString tempVal = callbackBlockchainStatusModel97d618_HashLastBlock(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setHashLastBlock(QString hashLastBlock) { QByteArray t33501e = hashLastBlock.toUtf8(); Moc_PackedString hashLastBlockPacked = { const_cast<char*>(t33501e.prepend("WHITESPACE").constData()+10), t33501e.size()-10 };callbackBlockchainStatusModel97d618_SetHashLastBlock(this, hashLastBlockPacked); };
	void Signal_HashLastBlockChanged(QString hashLastBlock) { QByteArray t33501e = hashLastBlock.toUtf8(); Moc_PackedString hashLastBlockPacked = { const_cast<char*>(t33501e.prepend("WHITESPACE").constData()+10), t33501e.size()-10 };callbackBlockchainStatusModel97d618_HashLastBlockChanged(this, hashLastBlockPacked); };
	QString currentSkySupply() { return ({ Moc_PackedString tempVal = callbackBlockchainStatusModel97d618_CurrentSkySupply(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setCurrentSkySupply(QString currentSkySupply) { QByteArray t9cbeff = currentSkySupply.toUtf8(); Moc_PackedString currentSkySupplyPacked = { const_cast<char*>(t9cbeff.prepend("WHITESPACE").constData()+10), t9cbeff.size()-10 };callbackBlockchainStatusModel97d618_SetCurrentSkySupply(this, currentSkySupplyPacked); };
	void Signal_CurrentSkySupplyChanged(QString currentSkySupply) { QByteArray t9cbeff = currentSkySupply.toUtf8(); Moc_PackedString currentSkySupplyPacked = { const_cast<char*>(t9cbeff.prepend("WHITESPACE").constData()+10), t9cbeff.size()-10 };callbackBlockchainStatusModel97d618_CurrentSkySupplyChanged(this, currentSkySupplyPacked); };
	QString totalSkySupply() { return ({ Moc_PackedString tempVal = callbackBlockchainStatusModel97d618_TotalSkySupply(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setTotalSkySupply(QString totalSkySupply) { QByteArray tf96a84 = totalSkySupply.toUtf8(); Moc_PackedString totalSkySupplyPacked = { const_cast<char*>(tf96a84.prepend("WHITESPACE").constData()+10), tf96a84.size()-10 };callbackBlockchainStatusModel97d618_SetTotalSkySupply(this, totalSkySupplyPacked); };
	void Signal_TotalSkySupplyChanged(QString totalSkySupply) { QByteArray tf96a84 = totalSkySupply.toUtf8(); Moc_PackedString totalSkySupplyPacked = { const_cast<char*>(tf96a84.prepend("WHITESPACE").constData()+10), tf96a84.size()-10 };callbackBlockchainStatusModel97d618_TotalSkySupplyChanged(this, totalSkySupplyPacked); };
	QString currentCoinHoursSupply() { return ({ Moc_PackedString tempVal = callbackBlockchainStatusModel97d618_CurrentCoinHoursSupply(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setCurrentCoinHoursSupply(QString currentCoinHoursSupply) { QByteArray t88f3d1 = currentCoinHoursSupply.toUtf8(); Moc_PackedString currentCoinHoursSupplyPacked = { const_cast<char*>(t88f3d1.prepend("WHITESPACE").constData()+10), t88f3d1.size()-10 };callbackBlockchainStatusModel97d618_SetCurrentCoinHoursSupply(this, currentCoinHoursSupplyPacked); };
	void Signal_CurrentCoinHoursSupplyChanged(QString currentCoinHoursSupply) { QByteArray t88f3d1 = currentCoinHoursSupply.toUtf8(); Moc_PackedString currentCoinHoursSupplyPacked = { const_cast<char*>(t88f3d1.prepend("WHITESPACE").constData()+10), t88f3d1.size()-10 };callbackBlockchainStatusModel97d618_CurrentCoinHoursSupplyChanged(this, currentCoinHoursSupplyPacked); };
	QString totalCoinHoursSupply() { return ({ Moc_PackedString tempVal = callbackBlockchainStatusModel97d618_TotalCoinHoursSupply(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setTotalCoinHoursSupply(QString totalCoinHoursSupply) { QByteArray t1cfb46 = totalCoinHoursSupply.toUtf8(); Moc_PackedString totalCoinHoursSupplyPacked = { const_cast<char*>(t1cfb46.prepend("WHITESPACE").constData()+10), t1cfb46.size()-10 };callbackBlockchainStatusModel97d618_SetTotalCoinHoursSupply(this, totalCoinHoursSupplyPacked); };
	void Signal_TotalCoinHoursSupplyChanged(QString totalCoinHoursSupply) { QByteArray t1cfb46 = totalCoinHoursSupply.toUtf8(); Moc_PackedString totalCoinHoursSupplyPacked = { const_cast<char*>(t1cfb46.prepend("WHITESPACE").constData()+10), t1cfb46.size()-10 };callbackBlockchainStatusModel97d618_TotalCoinHoursSupplyChanged(this, totalCoinHoursSupplyPacked); };
	 ~BlockchainStatusModel97d618() { callbackBlockchainStatusModel97d618_DestroyBlockchainStatusModel(this); };
	void childEvent(QChildEvent * event) { callbackBlockchainStatusModel97d618_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackBlockchainStatusModel97d618_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackBlockchainStatusModel97d618_CustomEvent(this, event); };
	void deleteLater() { callbackBlockchainStatusModel97d618_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackBlockchainStatusModel97d618_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackBlockchainStatusModel97d618_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackBlockchainStatusModel97d618_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackBlockchainStatusModel97d618_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackBlockchainStatusModel97d618_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackBlockchainStatusModel97d618_TimerEvent(this, event); };
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

Q_DECLARE_METATYPE(BlockchainStatusModel97d618*)


void BlockchainStatusModel97d618_BlockchainStatusModel97d618_QRegisterMetaTypes() {
	qRegisterMetaType<QDateTime>();
	qRegisterMetaType<QString>();
}

struct Moc_PackedString BlockchainStatusModel97d618_NumberOfBlocks(void* ptr)
{
	return ({ QByteArray t90e1b1 = static_cast<BlockchainStatusModel97d618*>(ptr)->numberOfBlocks().toUtf8(); Moc_PackedString { const_cast<char*>(t90e1b1.prepend("WHITESPACE").constData()+10), t90e1b1.size()-10 }; });
}

struct Moc_PackedString BlockchainStatusModel97d618_NumberOfBlocksDefault(void* ptr)
{
	return ({ QByteArray td41e94 = static_cast<BlockchainStatusModel97d618*>(ptr)->numberOfBlocksDefault().toUtf8(); Moc_PackedString { const_cast<char*>(td41e94.prepend("WHITESPACE").constData()+10), td41e94.size()-10 }; });
}

void BlockchainStatusModel97d618_SetNumberOfBlocks(void* ptr, struct Moc_PackedString numberOfBlocks)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setNumberOfBlocks(QString::fromUtf8(numberOfBlocks.data, numberOfBlocks.len));
}

void BlockchainStatusModel97d618_SetNumberOfBlocksDefault(void* ptr, struct Moc_PackedString numberOfBlocks)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setNumberOfBlocksDefault(QString::fromUtf8(numberOfBlocks.data, numberOfBlocks.len));
}

void BlockchainStatusModel97d618_ConnectNumberOfBlocksChanged(void* ptr)
{
	QObject::connect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::numberOfBlocksChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::Signal_NumberOfBlocksChanged));
}

void BlockchainStatusModel97d618_DisconnectNumberOfBlocksChanged(void* ptr)
{
	QObject::disconnect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::numberOfBlocksChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::Signal_NumberOfBlocksChanged));
}

void BlockchainStatusModel97d618_NumberOfBlocksChanged(void* ptr, struct Moc_PackedString numberOfBlocks)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->numberOfBlocksChanged(QString::fromUtf8(numberOfBlocks.data, numberOfBlocks.len));
}

void* BlockchainStatusModel97d618_TimestampLastBlock(void* ptr)
{
	return new QDateTime(static_cast<BlockchainStatusModel97d618*>(ptr)->timestampLastBlock());
}

void* BlockchainStatusModel97d618_TimestampLastBlockDefault(void* ptr)
{
	return new QDateTime(static_cast<BlockchainStatusModel97d618*>(ptr)->timestampLastBlockDefault());
}

void BlockchainStatusModel97d618_SetTimestampLastBlock(void* ptr, void* timestampLastBlock)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setTimestampLastBlock(*static_cast<QDateTime*>(timestampLastBlock));
}

void BlockchainStatusModel97d618_SetTimestampLastBlockDefault(void* ptr, void* timestampLastBlock)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setTimestampLastBlockDefault(*static_cast<QDateTime*>(timestampLastBlock));
}

void BlockchainStatusModel97d618_ConnectTimestampLastBlockChanged(void* ptr)
{
	QObject::connect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QDateTime)>(&BlockchainStatusModel97d618::timestampLastBlockChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QDateTime)>(&BlockchainStatusModel97d618::Signal_TimestampLastBlockChanged));
}

void BlockchainStatusModel97d618_DisconnectTimestampLastBlockChanged(void* ptr)
{
	QObject::disconnect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QDateTime)>(&BlockchainStatusModel97d618::timestampLastBlockChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QDateTime)>(&BlockchainStatusModel97d618::Signal_TimestampLastBlockChanged));
}

void BlockchainStatusModel97d618_TimestampLastBlockChanged(void* ptr, void* timestampLastBlock)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->timestampLastBlockChanged(*static_cast<QDateTime*>(timestampLastBlock));
}

struct Moc_PackedString BlockchainStatusModel97d618_HashLastBlock(void* ptr)
{
	return ({ QByteArray tbaa67e = static_cast<BlockchainStatusModel97d618*>(ptr)->hashLastBlock().toUtf8(); Moc_PackedString { const_cast<char*>(tbaa67e.prepend("WHITESPACE").constData()+10), tbaa67e.size()-10 }; });
}

struct Moc_PackedString BlockchainStatusModel97d618_HashLastBlockDefault(void* ptr)
{
	return ({ QByteArray td6ba2e = static_cast<BlockchainStatusModel97d618*>(ptr)->hashLastBlockDefault().toUtf8(); Moc_PackedString { const_cast<char*>(td6ba2e.prepend("WHITESPACE").constData()+10), td6ba2e.size()-10 }; });
}

void BlockchainStatusModel97d618_SetHashLastBlock(void* ptr, struct Moc_PackedString hashLastBlock)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setHashLastBlock(QString::fromUtf8(hashLastBlock.data, hashLastBlock.len));
}

void BlockchainStatusModel97d618_SetHashLastBlockDefault(void* ptr, struct Moc_PackedString hashLastBlock)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setHashLastBlockDefault(QString::fromUtf8(hashLastBlock.data, hashLastBlock.len));
}

void BlockchainStatusModel97d618_ConnectHashLastBlockChanged(void* ptr)
{
	QObject::connect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::hashLastBlockChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::Signal_HashLastBlockChanged));
}

void BlockchainStatusModel97d618_DisconnectHashLastBlockChanged(void* ptr)
{
	QObject::disconnect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::hashLastBlockChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::Signal_HashLastBlockChanged));
}

void BlockchainStatusModel97d618_HashLastBlockChanged(void* ptr, struct Moc_PackedString hashLastBlock)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->hashLastBlockChanged(QString::fromUtf8(hashLastBlock.data, hashLastBlock.len));
}

struct Moc_PackedString BlockchainStatusModel97d618_CurrentSkySupply(void* ptr)
{
	return ({ QByteArray ta7721d = static_cast<BlockchainStatusModel97d618*>(ptr)->currentSkySupply().toUtf8(); Moc_PackedString { const_cast<char*>(ta7721d.prepend("WHITESPACE").constData()+10), ta7721d.size()-10 }; });
}

struct Moc_PackedString BlockchainStatusModel97d618_CurrentSkySupplyDefault(void* ptr)
{
	return ({ QByteArray ta836bf = static_cast<BlockchainStatusModel97d618*>(ptr)->currentSkySupplyDefault().toUtf8(); Moc_PackedString { const_cast<char*>(ta836bf.prepend("WHITESPACE").constData()+10), ta836bf.size()-10 }; });
}

void BlockchainStatusModel97d618_SetCurrentSkySupply(void* ptr, struct Moc_PackedString currentSkySupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setCurrentSkySupply(QString::fromUtf8(currentSkySupply.data, currentSkySupply.len));
}

void BlockchainStatusModel97d618_SetCurrentSkySupplyDefault(void* ptr, struct Moc_PackedString currentSkySupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setCurrentSkySupplyDefault(QString::fromUtf8(currentSkySupply.data, currentSkySupply.len));
}

void BlockchainStatusModel97d618_ConnectCurrentSkySupplyChanged(void* ptr)
{
	QObject::connect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::currentSkySupplyChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::Signal_CurrentSkySupplyChanged));
}

void BlockchainStatusModel97d618_DisconnectCurrentSkySupplyChanged(void* ptr)
{
	QObject::disconnect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::currentSkySupplyChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::Signal_CurrentSkySupplyChanged));
}

void BlockchainStatusModel97d618_CurrentSkySupplyChanged(void* ptr, struct Moc_PackedString currentSkySupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->currentSkySupplyChanged(QString::fromUtf8(currentSkySupply.data, currentSkySupply.len));
}

struct Moc_PackedString BlockchainStatusModel97d618_TotalSkySupply(void* ptr)
{
	return ({ QByteArray t35020f = static_cast<BlockchainStatusModel97d618*>(ptr)->totalSkySupply().toUtf8(); Moc_PackedString { const_cast<char*>(t35020f.prepend("WHITESPACE").constData()+10), t35020f.size()-10 }; });
}

struct Moc_PackedString BlockchainStatusModel97d618_TotalSkySupplyDefault(void* ptr)
{
	return ({ QByteArray tec951b = static_cast<BlockchainStatusModel97d618*>(ptr)->totalSkySupplyDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tec951b.prepend("WHITESPACE").constData()+10), tec951b.size()-10 }; });
}

void BlockchainStatusModel97d618_SetTotalSkySupply(void* ptr, struct Moc_PackedString totalSkySupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setTotalSkySupply(QString::fromUtf8(totalSkySupply.data, totalSkySupply.len));
}

void BlockchainStatusModel97d618_SetTotalSkySupplyDefault(void* ptr, struct Moc_PackedString totalSkySupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setTotalSkySupplyDefault(QString::fromUtf8(totalSkySupply.data, totalSkySupply.len));
}

void BlockchainStatusModel97d618_ConnectTotalSkySupplyChanged(void* ptr)
{
	QObject::connect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::totalSkySupplyChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::Signal_TotalSkySupplyChanged));
}

void BlockchainStatusModel97d618_DisconnectTotalSkySupplyChanged(void* ptr)
{
	QObject::disconnect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::totalSkySupplyChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::Signal_TotalSkySupplyChanged));
}

void BlockchainStatusModel97d618_TotalSkySupplyChanged(void* ptr, struct Moc_PackedString totalSkySupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->totalSkySupplyChanged(QString::fromUtf8(totalSkySupply.data, totalSkySupply.len));
}

struct Moc_PackedString BlockchainStatusModel97d618_CurrentCoinHoursSupply(void* ptr)
{
	return ({ QByteArray t848392 = static_cast<BlockchainStatusModel97d618*>(ptr)->currentCoinHoursSupply().toUtf8(); Moc_PackedString { const_cast<char*>(t848392.prepend("WHITESPACE").constData()+10), t848392.size()-10 }; });
}

struct Moc_PackedString BlockchainStatusModel97d618_CurrentCoinHoursSupplyDefault(void* ptr)
{
	return ({ QByteArray t627637 = static_cast<BlockchainStatusModel97d618*>(ptr)->currentCoinHoursSupplyDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t627637.prepend("WHITESPACE").constData()+10), t627637.size()-10 }; });
}

void BlockchainStatusModel97d618_SetCurrentCoinHoursSupply(void* ptr, struct Moc_PackedString currentCoinHoursSupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setCurrentCoinHoursSupply(QString::fromUtf8(currentCoinHoursSupply.data, currentCoinHoursSupply.len));
}

void BlockchainStatusModel97d618_SetCurrentCoinHoursSupplyDefault(void* ptr, struct Moc_PackedString currentCoinHoursSupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setCurrentCoinHoursSupplyDefault(QString::fromUtf8(currentCoinHoursSupply.data, currentCoinHoursSupply.len));
}

void BlockchainStatusModel97d618_ConnectCurrentCoinHoursSupplyChanged(void* ptr)
{
	QObject::connect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::currentCoinHoursSupplyChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::Signal_CurrentCoinHoursSupplyChanged));
}

void BlockchainStatusModel97d618_DisconnectCurrentCoinHoursSupplyChanged(void* ptr)
{
	QObject::disconnect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::currentCoinHoursSupplyChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::Signal_CurrentCoinHoursSupplyChanged));
}

void BlockchainStatusModel97d618_CurrentCoinHoursSupplyChanged(void* ptr, struct Moc_PackedString currentCoinHoursSupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->currentCoinHoursSupplyChanged(QString::fromUtf8(currentCoinHoursSupply.data, currentCoinHoursSupply.len));
}

struct Moc_PackedString BlockchainStatusModel97d618_TotalCoinHoursSupply(void* ptr)
{
	return ({ QByteArray t379ab0 = static_cast<BlockchainStatusModel97d618*>(ptr)->totalCoinHoursSupply().toUtf8(); Moc_PackedString { const_cast<char*>(t379ab0.prepend("WHITESPACE").constData()+10), t379ab0.size()-10 }; });
}

struct Moc_PackedString BlockchainStatusModel97d618_TotalCoinHoursSupplyDefault(void* ptr)
{
	return ({ QByteArray t08aab2 = static_cast<BlockchainStatusModel97d618*>(ptr)->totalCoinHoursSupplyDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t08aab2.prepend("WHITESPACE").constData()+10), t08aab2.size()-10 }; });
}

void BlockchainStatusModel97d618_SetTotalCoinHoursSupply(void* ptr, struct Moc_PackedString totalCoinHoursSupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setTotalCoinHoursSupply(QString::fromUtf8(totalCoinHoursSupply.data, totalCoinHoursSupply.len));
}

void BlockchainStatusModel97d618_SetTotalCoinHoursSupplyDefault(void* ptr, struct Moc_PackedString totalCoinHoursSupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setTotalCoinHoursSupplyDefault(QString::fromUtf8(totalCoinHoursSupply.data, totalCoinHoursSupply.len));
}

void BlockchainStatusModel97d618_ConnectTotalCoinHoursSupplyChanged(void* ptr)
{
	QObject::connect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::totalCoinHoursSupplyChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::Signal_TotalCoinHoursSupplyChanged));
}

void BlockchainStatusModel97d618_DisconnectTotalCoinHoursSupplyChanged(void* ptr)
{
	QObject::disconnect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::totalCoinHoursSupplyChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::Signal_TotalCoinHoursSupplyChanged));
}

void BlockchainStatusModel97d618_TotalCoinHoursSupplyChanged(void* ptr, struct Moc_PackedString totalCoinHoursSupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->totalCoinHoursSupplyChanged(QString::fromUtf8(totalCoinHoursSupply.data, totalCoinHoursSupply.len));
}

int BlockchainStatusModel97d618_BlockchainStatusModel97d618_QRegisterMetaType()
{
	return qRegisterMetaType<BlockchainStatusModel97d618*>();
}

int BlockchainStatusModel97d618_BlockchainStatusModel97d618_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<BlockchainStatusModel97d618*>(const_cast<const char*>(typeName));
}

int BlockchainStatusModel97d618_BlockchainStatusModel97d618_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<BlockchainStatusModel97d618>();
#else
	return 0;
#endif
}

int BlockchainStatusModel97d618_BlockchainStatusModel97d618_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<BlockchainStatusModel97d618>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* BlockchainStatusModel97d618___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void BlockchainStatusModel97d618___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* BlockchainStatusModel97d618___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* BlockchainStatusModel97d618___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void BlockchainStatusModel97d618___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* BlockchainStatusModel97d618___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* BlockchainStatusModel97d618___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void BlockchainStatusModel97d618___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* BlockchainStatusModel97d618___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* BlockchainStatusModel97d618___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void BlockchainStatusModel97d618___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* BlockchainStatusModel97d618___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* BlockchainStatusModel97d618___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void BlockchainStatusModel97d618___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* BlockchainStatusModel97d618___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* BlockchainStatusModel97d618_NewBlockchainStatusModel(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new BlockchainStatusModel97d618(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new BlockchainStatusModel97d618(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new BlockchainStatusModel97d618(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new BlockchainStatusModel97d618(static_cast<QWindow*>(parent));
	} else {
		return new BlockchainStatusModel97d618(static_cast<QObject*>(parent));
	}
}

void BlockchainStatusModel97d618_DestroyBlockchainStatusModel(void* ptr)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->~BlockchainStatusModel97d618();
}

void BlockchainStatusModel97d618_DestroyBlockchainStatusModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void BlockchainStatusModel97d618_ChildEventDefault(void* ptr, void* event)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void BlockchainStatusModel97d618_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void BlockchainStatusModel97d618_CustomEventDefault(void* ptr, void* event)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void BlockchainStatusModel97d618_DeleteLaterDefault(void* ptr)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->QObject::deleteLater();
}

void BlockchainStatusModel97d618_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char BlockchainStatusModel97d618_EventDefault(void* ptr, void* e)
{
	return static_cast<BlockchainStatusModel97d618*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char BlockchainStatusModel97d618_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<BlockchainStatusModel97d618*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void BlockchainStatusModel97d618_TimerEventDefault(void* ptr, void* event)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
