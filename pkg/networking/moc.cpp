

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAbstractItemModel>
#include <QAbstractListModel>
#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QHash>
#include <QMap>
#include <QMetaMethod>
#include <QMimeData>
#include <QModelIndex>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QPersistentModelIndex>
#include <QSize>
#include <QString>
#include <QTimerEvent>
#include <QVariant>
#include <QVector>
#include <QWindow>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


typedef QMap<qint32, QByteArray> type378cdd;
class QConnection580e15: public QObject
{
Q_OBJECT
Q_PROPERTY(QString peer READ peer WRITE setPeer NOTIFY peerChanged)
Q_PROPERTY(QString source READ source WRITE setSource NOTIFY sourceChanged)
Q_PROPERTY(quint64 blockHeight READ blockHeight WRITE setBlockHeight NOTIFY blockHeightChanged)
Q_PROPERTY(qint64 lastSeen READ lastSeen WRITE setLastSeen NOTIFY lastSeenChanged)
public:
	QConnection580e15(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QConnection580e15_QConnection580e15_QRegisterMetaType();QConnection580e15_QConnection580e15_QRegisterMetaTypes();callbackQConnection580e15_Constructor(this);};
	QString peer() { return ({ Moc_PackedString tempVal = callbackQConnection580e15_Peer(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setPeer(QString peer) { QByteArray t2b974b = peer.toUtf8(); Moc_PackedString peerPacked = { const_cast<char*>(t2b974b.prepend("WHITESPACE").constData()+10), t2b974b.size()-10 };callbackQConnection580e15_SetPeer(this, peerPacked); };
	void Signal_PeerChanged(QString peer) { QByteArray t2b974b = peer.toUtf8(); Moc_PackedString peerPacked = { const_cast<char*>(t2b974b.prepend("WHITESPACE").constData()+10), t2b974b.size()-10 };callbackQConnection580e15_PeerChanged(this, peerPacked); };
	QString source() { return ({ Moc_PackedString tempVal = callbackQConnection580e15_Source(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setSource(QString source) { QByteArray t828d33 = source.toUtf8(); Moc_PackedString sourcePacked = { const_cast<char*>(t828d33.prepend("WHITESPACE").constData()+10), t828d33.size()-10 };callbackQConnection580e15_SetSource(this, sourcePacked); };
	void Signal_SourceChanged(QString source) { QByteArray t828d33 = source.toUtf8(); Moc_PackedString sourcePacked = { const_cast<char*>(t828d33.prepend("WHITESPACE").constData()+10), t828d33.size()-10 };callbackQConnection580e15_SourceChanged(this, sourcePacked); };
	quint64 blockHeight() { return callbackQConnection580e15_BlockHeight(this); };
	void setBlockHeight(quint64 blockHeight) { callbackQConnection580e15_SetBlockHeight(this, blockHeight); };
	void Signal_BlockHeightChanged(quint64 blockHeight) { callbackQConnection580e15_BlockHeightChanged(this, blockHeight); };
	qint64 lastSeen() { return callbackQConnection580e15_LastSeen(this); };
	void setLastSeen(qint64 lastSeen) { callbackQConnection580e15_SetLastSeen(this, lastSeen); };
	void Signal_LastSeenChanged(qint64 lastSeen) { callbackQConnection580e15_LastSeenChanged(this, lastSeen); };
	 ~QConnection580e15() { callbackQConnection580e15_DestroyQConnection(this); };
	void childEvent(QChildEvent * event) { callbackQConnection580e15_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQConnection580e15_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQConnection580e15_CustomEvent(this, event); };
	void deleteLater() { callbackQConnection580e15_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQConnection580e15_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQConnection580e15_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQConnection580e15_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQConnection580e15_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQConnection580e15_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQConnection580e15_TimerEvent(this, event); };
	QString peerDefault() { return _peer; };
	void setPeerDefault(QString p) { if (p != _peer) { _peer = p; peerChanged(_peer); } };
	QString sourceDefault() { return _source; };
	void setSourceDefault(QString p) { if (p != _source) { _source = p; sourceChanged(_source); } };
	quint64 blockHeightDefault() { return _blockHeight; };
	void setBlockHeightDefault(quint64 p) { if (p != _blockHeight) { _blockHeight = p; blockHeightChanged(_blockHeight); } };
	qint64 lastSeenDefault() { return _lastSeen; };
	void setLastSeenDefault(qint64 p) { if (p != _lastSeen) { _lastSeen = p; lastSeenChanged(_lastSeen); } };
signals:
	void peerChanged(QString peer);
	void sourceChanged(QString source);
	void blockHeightChanged(quint64 blockHeight);
	void lastSeenChanged(qint64 lastSeen);
public slots:
private:
	QString _peer;
	QString _source;
	quint64 _blockHeight;
	qint64 _lastSeen;
};

Q_DECLARE_METATYPE(QConnection580e15*)


void QConnection580e15_QConnection580e15_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class networkingModel580e15: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<QConnection580e15*> connections READ connections WRITE setConnections NOTIFY connectionsChanged)
public:
	networkingModel580e15(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");networkingModel580e15_networkingModel580e15_QRegisterMetaType();networkingModel580e15_networkingModel580e15_QRegisterMetaTypes();callbacknetworkingModel580e15_Constructor(this);};
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbacknetworkingModel580e15_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbacknetworkingModel580e15_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbacknetworkingModel580e15_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<QConnection580e15*> connections() { return ({ QList<QConnection580e15*>* tmpP = static_cast<QList<QConnection580e15*>*>(callbacknetworkingModel580e15_Connections(this)); QList<QConnection580e15*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setConnections(QList<QConnection580e15*> connections) { callbacknetworkingModel580e15_SetConnections(this, ({ QList<QConnection580e15*>* tmpValue = new QList<QConnection580e15*>(connections); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_ConnectionsChanged(QList<QConnection580e15*> connections) { callbacknetworkingModel580e15_ConnectionsChanged(this, ({ QList<QConnection580e15*>* tmpValue = new QList<QConnection580e15*>(connections); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	 ~networkingModel580e15() { callbacknetworkingModel580e15_DestroyNetworkingModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbacknetworkingModel580e15_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbacknetworkingModel580e15_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbacknetworkingModel580e15_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbacknetworkingModel580e15_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbacknetworkingModel580e15_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbacknetworkingModel580e15_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbacknetworkingModel580e15_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbacknetworkingModel580e15_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbacknetworkingModel580e15_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbacknetworkingModel580e15_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbacknetworkingModel580e15_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbacknetworkingModel580e15_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbacknetworkingModel580e15_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbacknetworkingModel580e15_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbacknetworkingModel580e15_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbacknetworkingModel580e15_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = new QVector<int>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbacknetworkingModel580e15_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbacknetworkingModel580e15_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbacknetworkingModel580e15_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbacknetworkingModel580e15_HeaderDataChanged(this, orientation, first, last); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbacknetworkingModel580e15_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbacknetworkingModel580e15_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbacknetworkingModel580e15_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbacknetworkingModel580e15_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbacknetworkingModel580e15_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbacknetworkingModel580e15_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbacknetworkingModel580e15_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbacknetworkingModel580e15_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbacknetworkingModel580e15_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbacknetworkingModel580e15_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbacknetworkingModel580e15_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbacknetworkingModel580e15_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbacknetworkingModel580e15_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbacknetworkingModel580e15_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbacknetworkingModel580e15_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbacknetworkingModel580e15_ResetInternalData(this); };
	void revert() { callbacknetworkingModel580e15_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbacknetworkingModel580e15_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbacknetworkingModel580e15_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbacknetworkingModel580e15_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbacknetworkingModel580e15_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbacknetworkingModel580e15_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbacknetworkingModel580e15_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbacknetworkingModel580e15_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbacknetworkingModel580e15_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbacknetworkingModel580e15_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbacknetworkingModel580e15_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbacknetworkingModel580e15_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbacknetworkingModel580e15_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbacknetworkingModel580e15_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbacknetworkingModel580e15_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbacknetworkingModel580e15_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbacknetworkingModel580e15_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbacknetworkingModel580e15_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbacknetworkingModel580e15_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbacknetworkingModel580e15_CustomEvent(this, event); };
	void deleteLater() { callbacknetworkingModel580e15_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbacknetworkingModel580e15_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbacknetworkingModel580e15_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbacknetworkingModel580e15_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbacknetworkingModel580e15_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbacknetworkingModel580e15_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbacknetworkingModel580e15_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<QConnection580e15*> connectionsDefault() { return _connections; };
	void setConnectionsDefault(QList<QConnection580e15*> p) { if (p != _connections) { _connections = p; connectionsChanged(_connections); } };
signals:
	void rolesChanged(type378cdd roles);
	void connectionsChanged(QList<QConnection580e15*> connections);
public slots:
private:
	type378cdd _roles;
	QList<QConnection580e15*> _connections;
};

Q_DECLARE_METATYPE(networkingModel580e15*)


void networkingModel580e15_networkingModel580e15_QRegisterMetaTypes() {
	qRegisterMetaType<type378cdd>("type378cdd");
	qRegisterMetaType<QList<QObject*>>("QList<QConnection580e15*>");
}

struct Moc_PackedString QConnection580e15_Peer(void* ptr)
{
	return ({ QByteArray tcd9a4f = static_cast<QConnection580e15*>(ptr)->peer().toUtf8(); Moc_PackedString { const_cast<char*>(tcd9a4f.prepend("WHITESPACE").constData()+10), tcd9a4f.size()-10 }; });
}

struct Moc_PackedString QConnection580e15_PeerDefault(void* ptr)
{
	return ({ QByteArray tda11e2 = static_cast<QConnection580e15*>(ptr)->peerDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tda11e2.prepend("WHITESPACE").constData()+10), tda11e2.size()-10 }; });
}

void QConnection580e15_SetPeer(void* ptr, struct Moc_PackedString peer)
{
	static_cast<QConnection580e15*>(ptr)->setPeer(QString::fromUtf8(peer.data, peer.len));
}

void QConnection580e15_SetPeerDefault(void* ptr, struct Moc_PackedString peer)
{
	static_cast<QConnection580e15*>(ptr)->setPeerDefault(QString::fromUtf8(peer.data, peer.len));
}

void QConnection580e15_ConnectPeerChanged(void* ptr)
{
	QObject::connect(static_cast<QConnection580e15*>(ptr), static_cast<void (QConnection580e15::*)(QString)>(&QConnection580e15::peerChanged), static_cast<QConnection580e15*>(ptr), static_cast<void (QConnection580e15::*)(QString)>(&QConnection580e15::Signal_PeerChanged));
}

void QConnection580e15_DisconnectPeerChanged(void* ptr)
{
	QObject::disconnect(static_cast<QConnection580e15*>(ptr), static_cast<void (QConnection580e15::*)(QString)>(&QConnection580e15::peerChanged), static_cast<QConnection580e15*>(ptr), static_cast<void (QConnection580e15::*)(QString)>(&QConnection580e15::Signal_PeerChanged));
}

void QConnection580e15_PeerChanged(void* ptr, struct Moc_PackedString peer)
{
	static_cast<QConnection580e15*>(ptr)->peerChanged(QString::fromUtf8(peer.data, peer.len));
}

struct Moc_PackedString QConnection580e15_Source(void* ptr)
{
	return ({ QByteArray t735974 = static_cast<QConnection580e15*>(ptr)->source().toUtf8(); Moc_PackedString { const_cast<char*>(t735974.prepend("WHITESPACE").constData()+10), t735974.size()-10 }; });
}

struct Moc_PackedString QConnection580e15_SourceDefault(void* ptr)
{
	return ({ QByteArray tadcb7f = static_cast<QConnection580e15*>(ptr)->sourceDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tadcb7f.prepend("WHITESPACE").constData()+10), tadcb7f.size()-10 }; });
}

void QConnection580e15_SetSource(void* ptr, struct Moc_PackedString source)
{
	static_cast<QConnection580e15*>(ptr)->setSource(QString::fromUtf8(source.data, source.len));
}

void QConnection580e15_SetSourceDefault(void* ptr, struct Moc_PackedString source)
{
	static_cast<QConnection580e15*>(ptr)->setSourceDefault(QString::fromUtf8(source.data, source.len));
}

void QConnection580e15_ConnectSourceChanged(void* ptr)
{
	QObject::connect(static_cast<QConnection580e15*>(ptr), static_cast<void (QConnection580e15::*)(QString)>(&QConnection580e15::sourceChanged), static_cast<QConnection580e15*>(ptr), static_cast<void (QConnection580e15::*)(QString)>(&QConnection580e15::Signal_SourceChanged));
}

void QConnection580e15_DisconnectSourceChanged(void* ptr)
{
	QObject::disconnect(static_cast<QConnection580e15*>(ptr), static_cast<void (QConnection580e15::*)(QString)>(&QConnection580e15::sourceChanged), static_cast<QConnection580e15*>(ptr), static_cast<void (QConnection580e15::*)(QString)>(&QConnection580e15::Signal_SourceChanged));
}

void QConnection580e15_SourceChanged(void* ptr, struct Moc_PackedString source)
{
	static_cast<QConnection580e15*>(ptr)->sourceChanged(QString::fromUtf8(source.data, source.len));
}

unsigned long long QConnection580e15_BlockHeight(void* ptr)
{
	return static_cast<QConnection580e15*>(ptr)->blockHeight();
}

unsigned long long QConnection580e15_BlockHeightDefault(void* ptr)
{
	return static_cast<QConnection580e15*>(ptr)->blockHeightDefault();
}

void QConnection580e15_SetBlockHeight(void* ptr, unsigned long long blockHeight)
{
	static_cast<QConnection580e15*>(ptr)->setBlockHeight(blockHeight);
}

void QConnection580e15_SetBlockHeightDefault(void* ptr, unsigned long long blockHeight)
{
	static_cast<QConnection580e15*>(ptr)->setBlockHeightDefault(blockHeight);
}

void QConnection580e15_ConnectBlockHeightChanged(void* ptr)
{
	QObject::connect(static_cast<QConnection580e15*>(ptr), static_cast<void (QConnection580e15::*)(quint64)>(&QConnection580e15::blockHeightChanged), static_cast<QConnection580e15*>(ptr), static_cast<void (QConnection580e15::*)(quint64)>(&QConnection580e15::Signal_BlockHeightChanged));
}

void QConnection580e15_DisconnectBlockHeightChanged(void* ptr)
{
	QObject::disconnect(static_cast<QConnection580e15*>(ptr), static_cast<void (QConnection580e15::*)(quint64)>(&QConnection580e15::blockHeightChanged), static_cast<QConnection580e15*>(ptr), static_cast<void (QConnection580e15::*)(quint64)>(&QConnection580e15::Signal_BlockHeightChanged));
}

void QConnection580e15_BlockHeightChanged(void* ptr, unsigned long long blockHeight)
{
	static_cast<QConnection580e15*>(ptr)->blockHeightChanged(blockHeight);
}

long long QConnection580e15_LastSeen(void* ptr)
{
	return static_cast<QConnection580e15*>(ptr)->lastSeen();
}

long long QConnection580e15_LastSeenDefault(void* ptr)
{
	return static_cast<QConnection580e15*>(ptr)->lastSeenDefault();
}

void QConnection580e15_SetLastSeen(void* ptr, long long lastSeen)
{
	static_cast<QConnection580e15*>(ptr)->setLastSeen(lastSeen);
}

void QConnection580e15_SetLastSeenDefault(void* ptr, long long lastSeen)
{
	static_cast<QConnection580e15*>(ptr)->setLastSeenDefault(lastSeen);
}

void QConnection580e15_ConnectLastSeenChanged(void* ptr)
{
	QObject::connect(static_cast<QConnection580e15*>(ptr), static_cast<void (QConnection580e15::*)(qint64)>(&QConnection580e15::lastSeenChanged), static_cast<QConnection580e15*>(ptr), static_cast<void (QConnection580e15::*)(qint64)>(&QConnection580e15::Signal_LastSeenChanged));
}

void QConnection580e15_DisconnectLastSeenChanged(void* ptr)
{
	QObject::disconnect(static_cast<QConnection580e15*>(ptr), static_cast<void (QConnection580e15::*)(qint64)>(&QConnection580e15::lastSeenChanged), static_cast<QConnection580e15*>(ptr), static_cast<void (QConnection580e15::*)(qint64)>(&QConnection580e15::Signal_LastSeenChanged));
}

void QConnection580e15_LastSeenChanged(void* ptr, long long lastSeen)
{
	static_cast<QConnection580e15*>(ptr)->lastSeenChanged(lastSeen);
}

int QConnection580e15_QConnection580e15_QRegisterMetaType()
{
	return qRegisterMetaType<QConnection580e15*>();
}

int QConnection580e15_QConnection580e15_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QConnection580e15*>(const_cast<const char*>(typeName));
}

int QConnection580e15_QConnection580e15_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QConnection580e15>();
#else
	return 0;
#endif
}

int QConnection580e15_QConnection580e15_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QConnection580e15>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QConnection580e15___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QConnection580e15___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QConnection580e15___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QConnection580e15___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QConnection580e15___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QConnection580e15___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QConnection580e15___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QConnection580e15___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QConnection580e15___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QConnection580e15___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QConnection580e15___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QConnection580e15___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QConnection580e15___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QConnection580e15___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QConnection580e15___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QConnection580e15_NewQConnection(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QConnection580e15(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QConnection580e15(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QConnection580e15(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QConnection580e15(static_cast<QWindow*>(parent));
	} else {
		return new QConnection580e15(static_cast<QObject*>(parent));
	}
}

void QConnection580e15_DestroyQConnection(void* ptr)
{
	static_cast<QConnection580e15*>(ptr)->~QConnection580e15();
}

void QConnection580e15_DestroyQConnectionDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void QConnection580e15_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QConnection580e15*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QConnection580e15_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QConnection580e15*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QConnection580e15_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QConnection580e15*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QConnection580e15_DeleteLaterDefault(void* ptr)
{
	static_cast<QConnection580e15*>(ptr)->QObject::deleteLater();
}

void QConnection580e15_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QConnection580e15*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QConnection580e15_EventDefault(void* ptr, void* e)
{
	return static_cast<QConnection580e15*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QConnection580e15_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QConnection580e15*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QConnection580e15_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QConnection580e15*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

struct Moc_PackedList networkingModel580e15_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<networkingModel580e15*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList networkingModel580e15_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<networkingModel580e15*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void networkingModel580e15_SetRoles(void* ptr, void* roles)
{
	static_cast<networkingModel580e15*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void networkingModel580e15_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<networkingModel580e15*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void networkingModel580e15_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<networkingModel580e15*>(ptr), static_cast<void (networkingModel580e15::*)(QMap<qint32, QByteArray>)>(&networkingModel580e15::rolesChanged), static_cast<networkingModel580e15*>(ptr), static_cast<void (networkingModel580e15::*)(QMap<qint32, QByteArray>)>(&networkingModel580e15::Signal_RolesChanged));
}

void networkingModel580e15_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<networkingModel580e15*>(ptr), static_cast<void (networkingModel580e15::*)(QMap<qint32, QByteArray>)>(&networkingModel580e15::rolesChanged), static_cast<networkingModel580e15*>(ptr), static_cast<void (networkingModel580e15::*)(QMap<qint32, QByteArray>)>(&networkingModel580e15::Signal_RolesChanged));
}

void networkingModel580e15_RolesChanged(void* ptr, void* roles)
{
	static_cast<networkingModel580e15*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList networkingModel580e15_Connections(void* ptr)
{
	return ({ QList<QConnection580e15*>* tmpValue = new QList<QConnection580e15*>(static_cast<networkingModel580e15*>(ptr)->connections()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList networkingModel580e15_ConnectionsDefault(void* ptr)
{
	return ({ QList<QConnection580e15*>* tmpValue = new QList<QConnection580e15*>(static_cast<networkingModel580e15*>(ptr)->connectionsDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void networkingModel580e15_SetConnections(void* ptr, void* connections)
{
	static_cast<networkingModel580e15*>(ptr)->setConnections(({ QList<QConnection580e15*>* tmpP = static_cast<QList<QConnection580e15*>*>(connections); QList<QConnection580e15*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void networkingModel580e15_SetConnectionsDefault(void* ptr, void* connections)
{
	static_cast<networkingModel580e15*>(ptr)->setConnectionsDefault(({ QList<QConnection580e15*>* tmpP = static_cast<QList<QConnection580e15*>*>(connections); QList<QConnection580e15*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void networkingModel580e15_ConnectConnectionsChanged(void* ptr)
{
	QObject::connect(static_cast<networkingModel580e15*>(ptr), static_cast<void (networkingModel580e15::*)(QList<QConnection580e15*>)>(&networkingModel580e15::connectionsChanged), static_cast<networkingModel580e15*>(ptr), static_cast<void (networkingModel580e15::*)(QList<QConnection580e15*>)>(&networkingModel580e15::Signal_ConnectionsChanged));
}

void networkingModel580e15_DisconnectConnectionsChanged(void* ptr)
{
	QObject::disconnect(static_cast<networkingModel580e15*>(ptr), static_cast<void (networkingModel580e15::*)(QList<QConnection580e15*>)>(&networkingModel580e15::connectionsChanged), static_cast<networkingModel580e15*>(ptr), static_cast<void (networkingModel580e15::*)(QList<QConnection580e15*>)>(&networkingModel580e15::Signal_ConnectionsChanged));
}

void networkingModel580e15_ConnectionsChanged(void* ptr, void* connections)
{
	static_cast<networkingModel580e15*>(ptr)->connectionsChanged(({ QList<QConnection580e15*>* tmpP = static_cast<QList<QConnection580e15*>*>(connections); QList<QConnection580e15*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int networkingModel580e15_networkingModel580e15_QRegisterMetaType()
{
	return qRegisterMetaType<networkingModel580e15*>();
}

int networkingModel580e15_networkingModel580e15_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<networkingModel580e15*>(const_cast<const char*>(typeName));
}

int networkingModel580e15_networkingModel580e15_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<networkingModel580e15>();
#else
	return 0;
#endif
}

int networkingModel580e15_networkingModel580e15_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<networkingModel580e15>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int networkingModel580e15_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void networkingModel580e15_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* networkingModel580e15_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int networkingModel580e15_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void networkingModel580e15_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* networkingModel580e15_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int networkingModel580e15_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void networkingModel580e15_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* networkingModel580e15_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* networkingModel580e15___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void networkingModel580e15___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* networkingModel580e15___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* networkingModel580e15___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void networkingModel580e15___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* networkingModel580e15___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int networkingModel580e15___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void networkingModel580e15___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* networkingModel580e15___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* networkingModel580e15___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void networkingModel580e15___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* networkingModel580e15___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList networkingModel580e15___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* networkingModel580e15___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void networkingModel580e15___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* networkingModel580e15___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* networkingModel580e15___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void networkingModel580e15___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* networkingModel580e15___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* networkingModel580e15___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void networkingModel580e15___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* networkingModel580e15___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* networkingModel580e15___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void networkingModel580e15___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* networkingModel580e15___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* networkingModel580e15___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void networkingModel580e15___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* networkingModel580e15___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* networkingModel580e15___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void networkingModel580e15___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* networkingModel580e15___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList networkingModel580e15___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* networkingModel580e15___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void networkingModel580e15___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* networkingModel580e15___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList networkingModel580e15___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int networkingModel580e15_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void networkingModel580e15_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* networkingModel580e15_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int networkingModel580e15_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void networkingModel580e15_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* networkingModel580e15_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* networkingModel580e15___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void networkingModel580e15___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* networkingModel580e15___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* networkingModel580e15___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void networkingModel580e15___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* networkingModel580e15___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* networkingModel580e15___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void networkingModel580e15___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* networkingModel580e15___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* networkingModel580e15___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void networkingModel580e15___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* networkingModel580e15___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* networkingModel580e15___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void networkingModel580e15___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* networkingModel580e15___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* networkingModel580e15___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void networkingModel580e15___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* networkingModel580e15___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList networkingModel580e15___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* networkingModel580e15___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void networkingModel580e15___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* networkingModel580e15___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList networkingModel580e15___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* networkingModel580e15___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void networkingModel580e15___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* networkingModel580e15___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList networkingModel580e15___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* networkingModel580e15___connections_atList(void* ptr, int i)
{
	return ({QConnection580e15* tmp = static_cast<QList<QConnection580e15*>*>(ptr)->at(i); if (i == static_cast<QList<QConnection580e15*>*>(ptr)->size()-1) { static_cast<QList<QConnection580e15*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void networkingModel580e15___connections_setList(void* ptr, void* i)
{
	static_cast<QList<QConnection580e15*>*>(ptr)->append(static_cast<QConnection580e15*>(i));
}

void* networkingModel580e15___connections_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QConnection580e15*>();
}

void* networkingModel580e15___setConnections_connections_atList(void* ptr, int i)
{
	return ({QConnection580e15* tmp = static_cast<QList<QConnection580e15*>*>(ptr)->at(i); if (i == static_cast<QList<QConnection580e15*>*>(ptr)->size()-1) { static_cast<QList<QConnection580e15*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void networkingModel580e15___setConnections_connections_setList(void* ptr, void* i)
{
	static_cast<QList<QConnection580e15*>*>(ptr)->append(static_cast<QConnection580e15*>(i));
}

void* networkingModel580e15___setConnections_connections_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QConnection580e15*>();
}

void* networkingModel580e15___connectionsChanged_connections_atList(void* ptr, int i)
{
	return ({QConnection580e15* tmp = static_cast<QList<QConnection580e15*>*>(ptr)->at(i); if (i == static_cast<QList<QConnection580e15*>*>(ptr)->size()-1) { static_cast<QList<QConnection580e15*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void networkingModel580e15___connectionsChanged_connections_setList(void* ptr, void* i)
{
	static_cast<QList<QConnection580e15*>*>(ptr)->append(static_cast<QConnection580e15*>(i));
}

void* networkingModel580e15___connectionsChanged_connections_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QConnection580e15*>();
}

int networkingModel580e15_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void networkingModel580e15_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* networkingModel580e15_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int networkingModel580e15_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void networkingModel580e15_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* networkingModel580e15_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int networkingModel580e15_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void networkingModel580e15_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* networkingModel580e15_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* networkingModel580e15_NewNetworkingModel(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new networkingModel580e15(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new networkingModel580e15(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new networkingModel580e15(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new networkingModel580e15(static_cast<QWindow*>(parent));
	} else {
		return new networkingModel580e15(static_cast<QObject*>(parent));
	}
}

void networkingModel580e15_DestroyNetworkingModel(void* ptr)
{
	static_cast<networkingModel580e15*>(ptr)->~networkingModel580e15();
}

void networkingModel580e15_DestroyNetworkingModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char networkingModel580e15_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

long long networkingModel580e15_FlagsDefault(void* ptr, void* index)
{
	return static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

void* networkingModel580e15_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* networkingModel580e15_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* networkingModel580e15_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

char networkingModel580e15_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char networkingModel580e15_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int networkingModel580e15_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* networkingModel580e15_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void networkingModel580e15_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char networkingModel580e15_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* networkingModel580e15_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

char networkingModel580e15_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char networkingModel580e15_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct Moc_PackedList networkingModel580e15_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList networkingModel580e15_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* networkingModel580e15_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct Moc_PackedString networkingModel580e15_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray td46c3c = static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::mimeTypes().join("¡¦!").toUtf8(); Moc_PackedString { const_cast<char*>(td46c3c.prepend("WHITESPACE").constData()+10), td46c3c.size()-10 }; });
}

char networkingModel580e15_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char networkingModel580e15_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* networkingModel580e15_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

char networkingModel580e15_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char networkingModel580e15_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void networkingModel580e15_ResetInternalDataDefault(void* ptr)
{
	static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::resetInternalData();
}

void networkingModel580e15_RevertDefault(void* ptr)
{
	static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::revert();
}

struct Moc_PackedList networkingModel580e15_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int networkingModel580e15_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char networkingModel580e15_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char networkingModel580e15_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char networkingModel580e15_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void networkingModel580e15_SortDefault(void* ptr, int column, long long order)
{
	static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* networkingModel580e15_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char networkingModel580e15_SubmitDefault(void* ptr)
{
	return static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::submit();
}

long long networkingModel580e15_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long networkingModel580e15_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::supportedDropActions();
}

void networkingModel580e15_ChildEventDefault(void* ptr, void* event)
{
	static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void networkingModel580e15_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void networkingModel580e15_CustomEventDefault(void* ptr, void* event)
{
	static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void networkingModel580e15_DeleteLaterDefault(void* ptr)
{
	static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::deleteLater();
}

void networkingModel580e15_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char networkingModel580e15_EventDefault(void* ptr, void* e)
{
	return static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char networkingModel580e15_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void networkingModel580e15_TimerEventDefault(void* ptr, void* event)
{
	static_cast<networkingModel580e15*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
