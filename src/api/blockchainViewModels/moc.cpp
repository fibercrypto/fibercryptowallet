

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QList>
#include <QMetaMethod>
#include <QObject>
#include <QString>
#include <QTimerEvent>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


class BlockchainStatusModel97d618: public QObject
{
Q_OBJECT
Q_PROPERTY(qint64 numberOfBlocks READ numberOfBlocks WRITE setNumberOfBlocks NOTIFY numberOfBlocksChanged)
Q_PROPERTY(QString timestampLastBlock READ timestampLastBlock WRITE setTimestampLastBlock NOTIFY timestampLastBlockChanged)
Q_PROPERTY(QString hashLastBlock READ hashLastBlock WRITE setHashLastBlock NOTIFY hashLastBlockChanged)
Q_PROPERTY(qint64 currentSkySupply READ currentSkySupply WRITE setCurrentSkySupply NOTIFY currentSkySupplyChanged)
Q_PROPERTY(qint64 totalSkySupply READ totalSkySupply WRITE setTotalSkySupply NOTIFY totalSkySupplyChanged)
Q_PROPERTY(qint64 currentCoinHoursSupply READ currentCoinHoursSupply WRITE setCurrentCoinHoursSupply NOTIFY currentCoinHoursSupplyChanged)
Q_PROPERTY(qint64 totalCoinHoursSupply READ totalCoinHoursSupply WRITE setTotalCoinHoursSupply NOTIFY totalCoinHoursSupplyChanged)
public:
	BlockchainStatusModel97d618(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");BlockchainStatusModel97d618_BlockchainStatusModel97d618_QRegisterMetaType();BlockchainStatusModel97d618_BlockchainStatusModel97d618_QRegisterMetaTypes();callbackBlockchainStatusModel97d618_Constructor(this);};
	qint64 numberOfBlocks() { return callbackBlockchainStatusModel97d618_NumberOfBlocks(this); };
	void setNumberOfBlocks(qint64 numberOfBlocks) { callbackBlockchainStatusModel97d618_SetNumberOfBlocks(this, numberOfBlocks); };
	void Signal_NumberOfBlocksChanged(qint64 numberOfBlocks) { callbackBlockchainStatusModel97d618_NumberOfBlocksChanged(this, numberOfBlocks); };
	QString timestampLastBlock() { return ({ Moc_PackedString tempVal = callbackBlockchainStatusModel97d618_TimestampLastBlock(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setTimestampLastBlock(QString timestampLastBlock) { QByteArray tf80d4b = timestampLastBlock.toUtf8(); Moc_PackedString timestampLastBlockPacked = { const_cast<char*>(tf80d4b.prepend("WHITESPACE").constData()+10), tf80d4b.size()-10 };callbackBlockchainStatusModel97d618_SetTimestampLastBlock(this, timestampLastBlockPacked); };
	void Signal_TimestampLastBlockChanged(QString timestampLastBlock) { QByteArray tf80d4b = timestampLastBlock.toUtf8(); Moc_PackedString timestampLastBlockPacked = { const_cast<char*>(tf80d4b.prepend("WHITESPACE").constData()+10), tf80d4b.size()-10 };callbackBlockchainStatusModel97d618_TimestampLastBlockChanged(this, timestampLastBlockPacked); };
	QString hashLastBlock() { return ({ Moc_PackedString tempVal = callbackBlockchainStatusModel97d618_HashLastBlock(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setHashLastBlock(QString hashLastBlock) { QByteArray t33501e = hashLastBlock.toUtf8(); Moc_PackedString hashLastBlockPacked = { const_cast<char*>(t33501e.prepend("WHITESPACE").constData()+10), t33501e.size()-10 };callbackBlockchainStatusModel97d618_SetHashLastBlock(this, hashLastBlockPacked); };
	void Signal_HashLastBlockChanged(QString hashLastBlock) { QByteArray t33501e = hashLastBlock.toUtf8(); Moc_PackedString hashLastBlockPacked = { const_cast<char*>(t33501e.prepend("WHITESPACE").constData()+10), t33501e.size()-10 };callbackBlockchainStatusModel97d618_HashLastBlockChanged(this, hashLastBlockPacked); };
	qint64 currentSkySupply() { return callbackBlockchainStatusModel97d618_CurrentSkySupply(this); };
	void setCurrentSkySupply(qint64 currentSkySupply) { callbackBlockchainStatusModel97d618_SetCurrentSkySupply(this, currentSkySupply); };
	void Signal_CurrentSkySupplyChanged(qint64 currentSkySupply) { callbackBlockchainStatusModel97d618_CurrentSkySupplyChanged(this, currentSkySupply); };
	qint64 totalSkySupply() { return callbackBlockchainStatusModel97d618_TotalSkySupply(this); };
	void setTotalSkySupply(qint64 totalSkySupply) { callbackBlockchainStatusModel97d618_SetTotalSkySupply(this, totalSkySupply); };
	void Signal_TotalSkySupplyChanged(qint64 totalSkySupply) { callbackBlockchainStatusModel97d618_TotalSkySupplyChanged(this, totalSkySupply); };
	qint64 currentCoinHoursSupply() { return callbackBlockchainStatusModel97d618_CurrentCoinHoursSupply(this); };
	void setCurrentCoinHoursSupply(qint64 currentCoinHoursSupply) { callbackBlockchainStatusModel97d618_SetCurrentCoinHoursSupply(this, currentCoinHoursSupply); };
	void Signal_CurrentCoinHoursSupplyChanged(qint64 currentCoinHoursSupply) { callbackBlockchainStatusModel97d618_CurrentCoinHoursSupplyChanged(this, currentCoinHoursSupply); };
	qint64 totalCoinHoursSupply() { return callbackBlockchainStatusModel97d618_TotalCoinHoursSupply(this); };
	void setTotalCoinHoursSupply(qint64 totalCoinHoursSupply) { callbackBlockchainStatusModel97d618_SetTotalCoinHoursSupply(this, totalCoinHoursSupply); };
	void Signal_TotalCoinHoursSupplyChanged(qint64 totalCoinHoursSupply) { callbackBlockchainStatusModel97d618_TotalCoinHoursSupplyChanged(this, totalCoinHoursSupply); };
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
	qint64 numberOfBlocksDefault() { return _numberOfBlocks; };
	void setNumberOfBlocksDefault(qint64 p) { if (p != _numberOfBlocks) { _numberOfBlocks = p; numberOfBlocksChanged(_numberOfBlocks); } };
	QString timestampLastBlockDefault() { return _timestampLastBlock; };
	void setTimestampLastBlockDefault(QString p) { if (p != _timestampLastBlock) { _timestampLastBlock = p; timestampLastBlockChanged(_timestampLastBlock); } };
	QString hashLastBlockDefault() { return _hashLastBlock; };
	void setHashLastBlockDefault(QString p) { if (p != _hashLastBlock) { _hashLastBlock = p; hashLastBlockChanged(_hashLastBlock); } };
	qint64 currentSkySupplyDefault() { return _currentSkySupply; };
	void setCurrentSkySupplyDefault(qint64 p) { if (p != _currentSkySupply) { _currentSkySupply = p; currentSkySupplyChanged(_currentSkySupply); } };
	qint64 totalSkySupplyDefault() { return _totalSkySupply; };
	void setTotalSkySupplyDefault(qint64 p) { if (p != _totalSkySupply) { _totalSkySupply = p; totalSkySupplyChanged(_totalSkySupply); } };
	qint64 currentCoinHoursSupplyDefault() { return _currentCoinHoursSupply; };
	void setCurrentCoinHoursSupplyDefault(qint64 p) { if (p != _currentCoinHoursSupply) { _currentCoinHoursSupply = p; currentCoinHoursSupplyChanged(_currentCoinHoursSupply); } };
	qint64 totalCoinHoursSupplyDefault() { return _totalCoinHoursSupply; };
	void setTotalCoinHoursSupplyDefault(qint64 p) { if (p != _totalCoinHoursSupply) { _totalCoinHoursSupply = p; totalCoinHoursSupplyChanged(_totalCoinHoursSupply); } };
signals:
	void numberOfBlocksChanged(qint64 numberOfBlocks);
	void timestampLastBlockChanged(QString timestampLastBlock);
	void hashLastBlockChanged(QString hashLastBlock);
	void currentSkySupplyChanged(qint64 currentSkySupply);
	void totalSkySupplyChanged(qint64 totalSkySupply);
	void currentCoinHoursSupplyChanged(qint64 currentCoinHoursSupply);
	void totalCoinHoursSupplyChanged(qint64 totalCoinHoursSupply);
public slots:
private:
	qint64 _numberOfBlocks;
	QString _timestampLastBlock;
	QString _hashLastBlock;
	qint64 _currentSkySupply;
	qint64 _totalSkySupply;
	qint64 _currentCoinHoursSupply;
	qint64 _totalCoinHoursSupply;
};

Q_DECLARE_METATYPE(BlockchainStatusModel97d618*)


void BlockchainStatusModel97d618_BlockchainStatusModel97d618_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

long long BlockchainStatusModel97d618_NumberOfBlocks(void* ptr)
{
	return static_cast<BlockchainStatusModel97d618*>(ptr)->numberOfBlocks();
}

long long BlockchainStatusModel97d618_NumberOfBlocksDefault(void* ptr)
{
	return static_cast<BlockchainStatusModel97d618*>(ptr)->numberOfBlocksDefault();
}

void BlockchainStatusModel97d618_SetNumberOfBlocks(void* ptr, long long numberOfBlocks)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setNumberOfBlocks(numberOfBlocks);
}

void BlockchainStatusModel97d618_SetNumberOfBlocksDefault(void* ptr, long long numberOfBlocks)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setNumberOfBlocksDefault(numberOfBlocks);
}

void BlockchainStatusModel97d618_ConnectNumberOfBlocksChanged(void* ptr)
{
	QObject::connect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(qint64)>(&BlockchainStatusModel97d618::numberOfBlocksChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(qint64)>(&BlockchainStatusModel97d618::Signal_NumberOfBlocksChanged));
}

void BlockchainStatusModel97d618_DisconnectNumberOfBlocksChanged(void* ptr)
{
	QObject::disconnect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(qint64)>(&BlockchainStatusModel97d618::numberOfBlocksChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(qint64)>(&BlockchainStatusModel97d618::Signal_NumberOfBlocksChanged));
}

void BlockchainStatusModel97d618_NumberOfBlocksChanged(void* ptr, long long numberOfBlocks)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->numberOfBlocksChanged(numberOfBlocks);
}

struct Moc_PackedString BlockchainStatusModel97d618_TimestampLastBlock(void* ptr)
{
	return ({ QByteArray t3a971b = static_cast<BlockchainStatusModel97d618*>(ptr)->timestampLastBlock().toUtf8(); Moc_PackedString { const_cast<char*>(t3a971b.prepend("WHITESPACE").constData()+10), t3a971b.size()-10 }; });
}

struct Moc_PackedString BlockchainStatusModel97d618_TimestampLastBlockDefault(void* ptr)
{
	return ({ QByteArray td66f60 = static_cast<BlockchainStatusModel97d618*>(ptr)->timestampLastBlockDefault().toUtf8(); Moc_PackedString { const_cast<char*>(td66f60.prepend("WHITESPACE").constData()+10), td66f60.size()-10 }; });
}

void BlockchainStatusModel97d618_SetTimestampLastBlock(void* ptr, struct Moc_PackedString timestampLastBlock)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setTimestampLastBlock(QString::fromUtf8(timestampLastBlock.data, timestampLastBlock.len));
}

void BlockchainStatusModel97d618_SetTimestampLastBlockDefault(void* ptr, struct Moc_PackedString timestampLastBlock)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setTimestampLastBlockDefault(QString::fromUtf8(timestampLastBlock.data, timestampLastBlock.len));
}

void BlockchainStatusModel97d618_ConnectTimestampLastBlockChanged(void* ptr)
{
	QObject::connect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::timestampLastBlockChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::Signal_TimestampLastBlockChanged));
}

void BlockchainStatusModel97d618_DisconnectTimestampLastBlockChanged(void* ptr)
{
	QObject::disconnect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::timestampLastBlockChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(QString)>(&BlockchainStatusModel97d618::Signal_TimestampLastBlockChanged));
}

void BlockchainStatusModel97d618_TimestampLastBlockChanged(void* ptr, struct Moc_PackedString timestampLastBlock)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->timestampLastBlockChanged(QString::fromUtf8(timestampLastBlock.data, timestampLastBlock.len));
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

long long BlockchainStatusModel97d618_CurrentSkySupply(void* ptr)
{
	return static_cast<BlockchainStatusModel97d618*>(ptr)->currentSkySupply();
}

long long BlockchainStatusModel97d618_CurrentSkySupplyDefault(void* ptr)
{
	return static_cast<BlockchainStatusModel97d618*>(ptr)->currentSkySupplyDefault();
}

void BlockchainStatusModel97d618_SetCurrentSkySupply(void* ptr, long long currentSkySupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setCurrentSkySupply(currentSkySupply);
}

void BlockchainStatusModel97d618_SetCurrentSkySupplyDefault(void* ptr, long long currentSkySupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setCurrentSkySupplyDefault(currentSkySupply);
}

void BlockchainStatusModel97d618_ConnectCurrentSkySupplyChanged(void* ptr)
{
	QObject::connect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(qint64)>(&BlockchainStatusModel97d618::currentSkySupplyChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(qint64)>(&BlockchainStatusModel97d618::Signal_CurrentSkySupplyChanged));
}

void BlockchainStatusModel97d618_DisconnectCurrentSkySupplyChanged(void* ptr)
{
	QObject::disconnect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(qint64)>(&BlockchainStatusModel97d618::currentSkySupplyChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(qint64)>(&BlockchainStatusModel97d618::Signal_CurrentSkySupplyChanged));
}

void BlockchainStatusModel97d618_CurrentSkySupplyChanged(void* ptr, long long currentSkySupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->currentSkySupplyChanged(currentSkySupply);
}

long long BlockchainStatusModel97d618_TotalSkySupply(void* ptr)
{
	return static_cast<BlockchainStatusModel97d618*>(ptr)->totalSkySupply();
}

long long BlockchainStatusModel97d618_TotalSkySupplyDefault(void* ptr)
{
	return static_cast<BlockchainStatusModel97d618*>(ptr)->totalSkySupplyDefault();
}

void BlockchainStatusModel97d618_SetTotalSkySupply(void* ptr, long long totalSkySupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setTotalSkySupply(totalSkySupply);
}

void BlockchainStatusModel97d618_SetTotalSkySupplyDefault(void* ptr, long long totalSkySupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setTotalSkySupplyDefault(totalSkySupply);
}

void BlockchainStatusModel97d618_ConnectTotalSkySupplyChanged(void* ptr)
{
	QObject::connect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(qint64)>(&BlockchainStatusModel97d618::totalSkySupplyChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(qint64)>(&BlockchainStatusModel97d618::Signal_TotalSkySupplyChanged));
}

void BlockchainStatusModel97d618_DisconnectTotalSkySupplyChanged(void* ptr)
{
	QObject::disconnect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(qint64)>(&BlockchainStatusModel97d618::totalSkySupplyChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(qint64)>(&BlockchainStatusModel97d618::Signal_TotalSkySupplyChanged));
}

void BlockchainStatusModel97d618_TotalSkySupplyChanged(void* ptr, long long totalSkySupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->totalSkySupplyChanged(totalSkySupply);
}

long long BlockchainStatusModel97d618_CurrentCoinHoursSupply(void* ptr)
{
	return static_cast<BlockchainStatusModel97d618*>(ptr)->currentCoinHoursSupply();
}

long long BlockchainStatusModel97d618_CurrentCoinHoursSupplyDefault(void* ptr)
{
	return static_cast<BlockchainStatusModel97d618*>(ptr)->currentCoinHoursSupplyDefault();
}

void BlockchainStatusModel97d618_SetCurrentCoinHoursSupply(void* ptr, long long currentCoinHoursSupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setCurrentCoinHoursSupply(currentCoinHoursSupply);
}

void BlockchainStatusModel97d618_SetCurrentCoinHoursSupplyDefault(void* ptr, long long currentCoinHoursSupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setCurrentCoinHoursSupplyDefault(currentCoinHoursSupply);
}

void BlockchainStatusModel97d618_ConnectCurrentCoinHoursSupplyChanged(void* ptr)
{
	QObject::connect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(qint64)>(&BlockchainStatusModel97d618::currentCoinHoursSupplyChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(qint64)>(&BlockchainStatusModel97d618::Signal_CurrentCoinHoursSupplyChanged));
}

void BlockchainStatusModel97d618_DisconnectCurrentCoinHoursSupplyChanged(void* ptr)
{
	QObject::disconnect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(qint64)>(&BlockchainStatusModel97d618::currentCoinHoursSupplyChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(qint64)>(&BlockchainStatusModel97d618::Signal_CurrentCoinHoursSupplyChanged));
}

void BlockchainStatusModel97d618_CurrentCoinHoursSupplyChanged(void* ptr, long long currentCoinHoursSupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->currentCoinHoursSupplyChanged(currentCoinHoursSupply);
}

long long BlockchainStatusModel97d618_TotalCoinHoursSupply(void* ptr)
{
	return static_cast<BlockchainStatusModel97d618*>(ptr)->totalCoinHoursSupply();
}

long long BlockchainStatusModel97d618_TotalCoinHoursSupplyDefault(void* ptr)
{
	return static_cast<BlockchainStatusModel97d618*>(ptr)->totalCoinHoursSupplyDefault();
}

void BlockchainStatusModel97d618_SetTotalCoinHoursSupply(void* ptr, long long totalCoinHoursSupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setTotalCoinHoursSupply(totalCoinHoursSupply);
}

void BlockchainStatusModel97d618_SetTotalCoinHoursSupplyDefault(void* ptr, long long totalCoinHoursSupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->setTotalCoinHoursSupplyDefault(totalCoinHoursSupply);
}

void BlockchainStatusModel97d618_ConnectTotalCoinHoursSupplyChanged(void* ptr)
{
	QObject::connect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(qint64)>(&BlockchainStatusModel97d618::totalCoinHoursSupplyChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(qint64)>(&BlockchainStatusModel97d618::Signal_TotalCoinHoursSupplyChanged));
}

void BlockchainStatusModel97d618_DisconnectTotalCoinHoursSupplyChanged(void* ptr)
{
	QObject::disconnect(static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(qint64)>(&BlockchainStatusModel97d618::totalCoinHoursSupplyChanged), static_cast<BlockchainStatusModel97d618*>(ptr), static_cast<void (BlockchainStatusModel97d618::*)(qint64)>(&BlockchainStatusModel97d618::Signal_TotalCoinHoursSupplyChanged));
}

void BlockchainStatusModel97d618_TotalCoinHoursSupplyChanged(void* ptr, long long totalCoinHoursSupply)
{
	static_cast<BlockchainStatusModel97d618*>(ptr)->totalCoinHoursSupplyChanged(totalCoinHoursSupply);
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
	return new BlockchainStatusModel97d618(static_cast<QObject*>(parent));
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
