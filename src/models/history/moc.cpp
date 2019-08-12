

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

class TransactionDetailsecff1c: public QObject{
public:
	TransactionDetailsecff1c(QObject *parent) : QObject(parent) {};

};

typedef QMap<qint32, QByteArray> type378cdd;
class TransactionList554044: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<TransactionDetailsecff1c*> transactions READ transactions WRITE setTransactions NOTIFY transactionsChanged)
public:
	TransactionList554044(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");TransactionList554044_TransactionList554044_QRegisterMetaType();TransactionList554044_TransactionList554044_QRegisterMetaTypes();callbackTransactionList554044_Constructor(this);};
	void Signal_AddTransaction(TransactionDetailsecff1c* transaction) { callbackTransactionList554044_AddTransaction(this, transaction); };
	void Signal_RemoveTransaction(qint32 index) { callbackTransactionList554044_RemoveTransaction(this, index); };
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbackTransactionList554044_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbackTransactionList554044_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbackTransactionList554044_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<TransactionDetailsecff1c*> transactions() { return ({ QList<TransactionDetailsecff1c*>* tmpP = static_cast<QList<TransactionDetailsecff1c*>*>(callbackTransactionList554044_Transactions(this)); QList<TransactionDetailsecff1c*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setTransactions(QList<TransactionDetailsecff1c*> transactions) { callbackTransactionList554044_SetTransactions(this, ({ QList<TransactionDetailsecff1c*>* tmpValue = new QList<TransactionDetailsecff1c*>(transactions); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_TransactionsChanged(QList<TransactionDetailsecff1c*> transactions) { callbackTransactionList554044_TransactionsChanged(this, ({ QList<TransactionDetailsecff1c*>* tmpValue = new QList<TransactionDetailsecff1c*>(transactions); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	 ~TransactionList554044() { callbackTransactionList554044_DestroyTransactionList(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackTransactionList554044_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackTransactionList554044_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackTransactionList554044_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackTransactionList554044_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackTransactionList554044_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackTransactionList554044_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackTransactionList554044_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackTransactionList554044_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackTransactionList554044_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackTransactionList554044_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackTransactionList554044_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackTransactionList554044_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackTransactionList554044_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackTransactionList554044_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackTransactionList554044_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackTransactionList554044_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = new QVector<int>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackTransactionList554044_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackTransactionList554044_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackTransactionList554044_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackTransactionList554044_HeaderDataChanged(this, orientation, first, last); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackTransactionList554044_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackTransactionList554044_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackTransactionList554044_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackTransactionList554044_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackTransactionList554044_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackTransactionList554044_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackTransactionList554044_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackTransactionList554044_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackTransactionList554044_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackTransactionList554044_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackTransactionList554044_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackTransactionList554044_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackTransactionList554044_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackTransactionList554044_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackTransactionList554044_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackTransactionList554044_ResetInternalData(this); };
	void revert() { callbackTransactionList554044_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackTransactionList554044_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackTransactionList554044_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackTransactionList554044_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackTransactionList554044_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackTransactionList554044_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackTransactionList554044_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackTransactionList554044_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackTransactionList554044_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackTransactionList554044_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackTransactionList554044_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackTransactionList554044_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackTransactionList554044_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackTransactionList554044_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackTransactionList554044_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackTransactionList554044_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackTransactionList554044_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackTransactionList554044_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackTransactionList554044_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackTransactionList554044_CustomEvent(this, event); };
	void deleteLater() { callbackTransactionList554044_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackTransactionList554044_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackTransactionList554044_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackTransactionList554044_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackTransactionList554044_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackTransactionList554044_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackTransactionList554044_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<TransactionDetailsecff1c*> transactionsDefault() { return _transactions; };
	void setTransactionsDefault(QList<TransactionDetailsecff1c*> p) { if (p != _transactions) { _transactions = p; transactionsChanged(_transactions); } };
signals:
	void addTransaction(TransactionDetailsecff1c* transaction);
	void removeTransaction(qint32 index);
	void rolesChanged(type378cdd roles);
	void transactionsChanged(QList<TransactionDetailsecff1c*> transactions);
public slots:
	void addMultipleTransactions(QList<TransactionDetailsecff1c*> txns) { callbackTransactionList554044_AddMultipleTransactions(this, ({ QList<TransactionDetailsecff1c*>* tmpValue = new QList<TransactionDetailsecff1c*>(txns); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
private:
	type378cdd _roles;
	QList<TransactionDetailsecff1c*> _transactions;
};

Q_DECLARE_METATYPE(TransactionList554044*)


void TransactionList554044_TransactionList554044_QRegisterMetaTypes() {
	qRegisterMetaType<type378cdd>("type378cdd");
	qRegisterMetaType<QList<QObject*>>("QList<TransactionDetailsecff1c*>");
}

class HistoryManager554044: public QObject
{
Q_OBJECT
public:
	HistoryManager554044(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");HistoryManager554044_HistoryManager554044_QRegisterMetaType();HistoryManager554044_HistoryManager554044_QRegisterMetaTypes();callbackHistoryManager554044_Constructor(this);};
	 ~HistoryManager554044() { callbackHistoryManager554044_DestroyHistoryManager(this); };
	void childEvent(QChildEvent * event) { callbackHistoryManager554044_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackHistoryManager554044_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackHistoryManager554044_CustomEvent(this, event); };
	void deleteLater() { callbackHistoryManager554044_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackHistoryManager554044_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackHistoryManager554044_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackHistoryManager554044_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackHistoryManager554044_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackHistoryManager554044_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackHistoryManager554044_TimerEvent(this, event); };
signals:
public slots:
	QList<TransactionDetailsecff1c*> loadHistoryWithFilters(QStringList filterAddresses) { QByteArray t0bfc4c = filterAddresses.join("¡¦!").toUtf8(); Moc_PackedString filterAddressesPacked = { const_cast<char*>(t0bfc4c.prepend("WHITESPACE").constData()+10), t0bfc4c.size()-10 };return ({ QList<TransactionDetailsecff1c*>* tmpP = static_cast<QList<TransactionDetailsecff1c*>*>(callbackHistoryManager554044_LoadHistoryWithFilters(this, filterAddressesPacked)); QList<TransactionDetailsecff1c*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QList<TransactionDetailsecff1c*> loadHistory() { return ({ QList<TransactionDetailsecff1c*>* tmpP = static_cast<QList<TransactionDetailsecff1c*>*>(callbackHistoryManager554044_LoadHistory(this)); QList<TransactionDetailsecff1c*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
private:
};

Q_DECLARE_METATYPE(HistoryManager554044*)


void HistoryManager554044_HistoryManager554044_QRegisterMetaTypes() {
	qRegisterMetaType<QList<QObject*>>("QList<TransactionDetailsecff1c*>");
}

void TransactionList554044_ConnectAddTransaction(void* ptr)
{
	QObject::connect(static_cast<TransactionList554044*>(ptr), static_cast<void (TransactionList554044::*)(TransactionDetailsecff1c*)>(&TransactionList554044::addTransaction), static_cast<TransactionList554044*>(ptr), static_cast<void (TransactionList554044::*)(TransactionDetailsecff1c*)>(&TransactionList554044::Signal_AddTransaction));
}

void TransactionList554044_DisconnectAddTransaction(void* ptr)
{
	QObject::disconnect(static_cast<TransactionList554044*>(ptr), static_cast<void (TransactionList554044::*)(TransactionDetailsecff1c*)>(&TransactionList554044::addTransaction), static_cast<TransactionList554044*>(ptr), static_cast<void (TransactionList554044::*)(TransactionDetailsecff1c*)>(&TransactionList554044::Signal_AddTransaction));
}

void TransactionList554044_AddTransaction(void* ptr, void* transaction)
{
	static_cast<TransactionList554044*>(ptr)->addTransaction(static_cast<TransactionDetailsecff1c*>(transaction));
}

void TransactionList554044_ConnectRemoveTransaction(void* ptr)
{
	QObject::connect(static_cast<TransactionList554044*>(ptr), static_cast<void (TransactionList554044::*)(qint32)>(&TransactionList554044::removeTransaction), static_cast<TransactionList554044*>(ptr), static_cast<void (TransactionList554044::*)(qint32)>(&TransactionList554044::Signal_RemoveTransaction));
}

void TransactionList554044_DisconnectRemoveTransaction(void* ptr)
{
	QObject::disconnect(static_cast<TransactionList554044*>(ptr), static_cast<void (TransactionList554044::*)(qint32)>(&TransactionList554044::removeTransaction), static_cast<TransactionList554044*>(ptr), static_cast<void (TransactionList554044::*)(qint32)>(&TransactionList554044::Signal_RemoveTransaction));
}

void TransactionList554044_RemoveTransaction(void* ptr, int index)
{
	static_cast<TransactionList554044*>(ptr)->removeTransaction(index);
}

void TransactionList554044_AddMultipleTransactions(void* ptr, void* txns)
{
	QMetaObject::invokeMethod(static_cast<TransactionList554044*>(ptr), "addMultipleTransactions", Q_ARG(QList<TransactionDetailsecff1c*>, ({ QList<TransactionDetailsecff1c*>* tmpP = static_cast<QList<TransactionDetailsecff1c*>*>(txns); QList<TransactionDetailsecff1c*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; })));
}

struct Moc_PackedList TransactionList554044_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<TransactionList554044*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList TransactionList554044_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<TransactionList554044*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void TransactionList554044_SetRoles(void* ptr, void* roles)
{
	static_cast<TransactionList554044*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void TransactionList554044_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<TransactionList554044*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void TransactionList554044_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionList554044*>(ptr), static_cast<void (TransactionList554044::*)(QMap<qint32, QByteArray>)>(&TransactionList554044::rolesChanged), static_cast<TransactionList554044*>(ptr), static_cast<void (TransactionList554044::*)(QMap<qint32, QByteArray>)>(&TransactionList554044::Signal_RolesChanged));
}

void TransactionList554044_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionList554044*>(ptr), static_cast<void (TransactionList554044::*)(QMap<qint32, QByteArray>)>(&TransactionList554044::rolesChanged), static_cast<TransactionList554044*>(ptr), static_cast<void (TransactionList554044::*)(QMap<qint32, QByteArray>)>(&TransactionList554044::Signal_RolesChanged));
}

void TransactionList554044_RolesChanged(void* ptr, void* roles)
{
	static_cast<TransactionList554044*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList TransactionList554044_Transactions(void* ptr)
{
	return ({ QList<TransactionDetailsecff1c*>* tmpValue = new QList<TransactionDetailsecff1c*>(static_cast<TransactionList554044*>(ptr)->transactions()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList TransactionList554044_TransactionsDefault(void* ptr)
{
	return ({ QList<TransactionDetailsecff1c*>* tmpValue = new QList<TransactionDetailsecff1c*>(static_cast<TransactionList554044*>(ptr)->transactionsDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void TransactionList554044_SetTransactions(void* ptr, void* transactions)
{
	static_cast<TransactionList554044*>(ptr)->setTransactions(({ QList<TransactionDetailsecff1c*>* tmpP = static_cast<QList<TransactionDetailsecff1c*>*>(transactions); QList<TransactionDetailsecff1c*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void TransactionList554044_SetTransactionsDefault(void* ptr, void* transactions)
{
	static_cast<TransactionList554044*>(ptr)->setTransactionsDefault(({ QList<TransactionDetailsecff1c*>* tmpP = static_cast<QList<TransactionDetailsecff1c*>*>(transactions); QList<TransactionDetailsecff1c*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void TransactionList554044_ConnectTransactionsChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionList554044*>(ptr), static_cast<void (TransactionList554044::*)(QList<TransactionDetailsecff1c*>)>(&TransactionList554044::transactionsChanged), static_cast<TransactionList554044*>(ptr), static_cast<void (TransactionList554044::*)(QList<TransactionDetailsecff1c*>)>(&TransactionList554044::Signal_TransactionsChanged));
}

void TransactionList554044_DisconnectTransactionsChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionList554044*>(ptr), static_cast<void (TransactionList554044::*)(QList<TransactionDetailsecff1c*>)>(&TransactionList554044::transactionsChanged), static_cast<TransactionList554044*>(ptr), static_cast<void (TransactionList554044::*)(QList<TransactionDetailsecff1c*>)>(&TransactionList554044::Signal_TransactionsChanged));
}

void TransactionList554044_TransactionsChanged(void* ptr, void* transactions)
{
	static_cast<TransactionList554044*>(ptr)->transactionsChanged(({ QList<TransactionDetailsecff1c*>* tmpP = static_cast<QList<TransactionDetailsecff1c*>*>(transactions); QList<TransactionDetailsecff1c*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int TransactionList554044_TransactionList554044_QRegisterMetaType()
{
	return qRegisterMetaType<TransactionList554044*>();
}

int TransactionList554044_TransactionList554044_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<TransactionList554044*>(const_cast<const char*>(typeName));
}

int TransactionList554044_TransactionList554044_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TransactionList554044>();
#else
	return 0;
#endif
}

int TransactionList554044_TransactionList554044_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TransactionList554044>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int TransactionList554044_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionList554044_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* TransactionList554044_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int TransactionList554044_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionList554044_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* TransactionList554044_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int TransactionList554044_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionList554044_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* TransactionList554044_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* TransactionList554044___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TransactionList554044___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* TransactionList554044___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* TransactionList554044___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TransactionList554044___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* TransactionList554044___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int TransactionList554044___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void TransactionList554044___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* TransactionList554044___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* TransactionList554044___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void TransactionList554044___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* TransactionList554044___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList TransactionList554044___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TransactionList554044___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TransactionList554044___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* TransactionList554044___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* TransactionList554044___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TransactionList554044___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* TransactionList554044___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* TransactionList554044___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TransactionList554044___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* TransactionList554044___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* TransactionList554044___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TransactionList554044___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* TransactionList554044___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* TransactionList554044___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TransactionList554044___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* TransactionList554044___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* TransactionList554044___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void TransactionList554044___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* TransactionList554044___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList TransactionList554044___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TransactionList554044___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void TransactionList554044___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* TransactionList554044___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList TransactionList554044___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int TransactionList554044_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionList554044_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* TransactionList554044_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int TransactionList554044_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionList554044_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* TransactionList554044_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* TransactionList554044___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionList554044___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TransactionList554044___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* TransactionList554044___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TransactionList554044___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* TransactionList554044___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* TransactionList554044___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionList554044___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TransactionList554044___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TransactionList554044___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionList554044___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TransactionList554044___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TransactionList554044___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionList554044___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TransactionList554044___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TransactionList554044___addMultipleTransactions_txns_atList(void* ptr, int i)
{
	return ({TransactionDetailsecff1c* tmp = static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->at(i); if (i == static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->size()-1) { static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionList554044___addMultipleTransactions_txns_setList(void* ptr, void* i)
{
	static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->append(static_cast<TransactionDetailsecff1c*>(i));
}

void* TransactionList554044___addMultipleTransactions_txns_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<TransactionDetailsecff1c*>();
}

void* TransactionList554044___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void TransactionList554044___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* TransactionList554044___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList TransactionList554044___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TransactionList554044___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void TransactionList554044___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* TransactionList554044___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList TransactionList554044___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TransactionList554044___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void TransactionList554044___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* TransactionList554044___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList TransactionList554044___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TransactionList554044___transactions_atList(void* ptr, int i)
{
	return ({TransactionDetailsecff1c* tmp = static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->at(i); if (i == static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->size()-1) { static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionList554044___transactions_setList(void* ptr, void* i)
{
	static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->append(static_cast<TransactionDetailsecff1c*>(i));
}

void* TransactionList554044___transactions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<TransactionDetailsecff1c*>();
}

void* TransactionList554044___setTransactions_transactions_atList(void* ptr, int i)
{
	return ({TransactionDetailsecff1c* tmp = static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->at(i); if (i == static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->size()-1) { static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionList554044___setTransactions_transactions_setList(void* ptr, void* i)
{
	static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->append(static_cast<TransactionDetailsecff1c*>(i));
}

void* TransactionList554044___setTransactions_transactions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<TransactionDetailsecff1c*>();
}

void* TransactionList554044___transactionsChanged_transactions_atList(void* ptr, int i)
{
	return ({TransactionDetailsecff1c* tmp = static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->at(i); if (i == static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->size()-1) { static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionList554044___transactionsChanged_transactions_setList(void* ptr, void* i)
{
	static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->append(static_cast<TransactionDetailsecff1c*>(i));
}

void* TransactionList554044___transactionsChanged_transactions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<TransactionDetailsecff1c*>();
}

int TransactionList554044_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionList554044_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* TransactionList554044_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int TransactionList554044_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionList554044_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* TransactionList554044_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int TransactionList554044_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionList554044_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* TransactionList554044_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* TransactionList554044_NewTransactionList(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new TransactionList554044(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new TransactionList554044(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new TransactionList554044(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new TransactionList554044(static_cast<QWindow*>(parent));
	} else {
		return new TransactionList554044(static_cast<QObject*>(parent));
	}
}

void TransactionList554044_DestroyTransactionList(void* ptr)
{
	static_cast<TransactionList554044*>(ptr)->~TransactionList554044();
}

void TransactionList554044_DestroyTransactionListDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char TransactionList554044_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<TransactionList554044*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

long long TransactionList554044_FlagsDefault(void* ptr, void* index)
{
	return static_cast<TransactionList554044*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

void* TransactionList554044_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<TransactionList554044*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* TransactionList554044_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<TransactionList554044*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* TransactionList554044_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<TransactionList554044*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

char TransactionList554044_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<TransactionList554044*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char TransactionList554044_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<TransactionList554044*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int TransactionList554044_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<TransactionList554044*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* TransactionList554044_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void TransactionList554044_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<TransactionList554044*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char TransactionList554044_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<TransactionList554044*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* TransactionList554044_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<TransactionList554044*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

char TransactionList554044_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<TransactionList554044*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char TransactionList554044_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<TransactionList554044*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct Moc_PackedList TransactionList554044_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<TransactionList554044*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList TransactionList554044_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<TransactionList554044*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* TransactionList554044_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<TransactionList554044*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct Moc_PackedString TransactionList554044_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t9ca435 = static_cast<TransactionList554044*>(ptr)->QAbstractListModel::mimeTypes().join("¡¦!").toUtf8(); Moc_PackedString { const_cast<char*>(t9ca435.prepend("WHITESPACE").constData()+10), t9ca435.size()-10 }; });
}

char TransactionList554044_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<TransactionList554044*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char TransactionList554044_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<TransactionList554044*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* TransactionList554044_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<TransactionList554044*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

char TransactionList554044_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<TransactionList554044*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char TransactionList554044_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<TransactionList554044*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void TransactionList554044_ResetInternalDataDefault(void* ptr)
{
	static_cast<TransactionList554044*>(ptr)->QAbstractListModel::resetInternalData();
}

void TransactionList554044_RevertDefault(void* ptr)
{
	static_cast<TransactionList554044*>(ptr)->QAbstractListModel::revert();
}

struct Moc_PackedList TransactionList554044_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<TransactionList554044*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int TransactionList554044_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char TransactionList554044_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<TransactionList554044*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char TransactionList554044_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<TransactionList554044*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char TransactionList554044_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<TransactionList554044*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void TransactionList554044_SortDefault(void* ptr, int column, long long order)
{
	static_cast<TransactionList554044*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* TransactionList554044_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<TransactionList554044*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char TransactionList554044_SubmitDefault(void* ptr)
{
	return static_cast<TransactionList554044*>(ptr)->QAbstractListModel::submit();
}

long long TransactionList554044_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<TransactionList554044*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long TransactionList554044_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<TransactionList554044*>(ptr)->QAbstractListModel::supportedDropActions();
}

void TransactionList554044_ChildEventDefault(void* ptr, void* event)
{
	static_cast<TransactionList554044*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void TransactionList554044_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TransactionList554044*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void TransactionList554044_CustomEventDefault(void* ptr, void* event)
{
	static_cast<TransactionList554044*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void TransactionList554044_DeleteLaterDefault(void* ptr)
{
	static_cast<TransactionList554044*>(ptr)->QAbstractListModel::deleteLater();
}

void TransactionList554044_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TransactionList554044*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char TransactionList554044_EventDefault(void* ptr, void* e)
{
	return static_cast<TransactionList554044*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char TransactionList554044_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<TransactionList554044*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void TransactionList554044_TimerEventDefault(void* ptr, void* event)
{
	static_cast<TransactionList554044*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

struct Moc_PackedList HistoryManager554044_LoadHistoryWithFilters(void* ptr, struct Moc_PackedString filterAddresses)
{
	QList<TransactionDetailsecff1c*> returnArg;
	QMetaObject::invokeMethod(static_cast<HistoryManager554044*>(ptr), "loadHistoryWithFilters", Q_RETURN_ARG(QList<TransactionDetailsecff1c*>, returnArg), Q_ARG(QStringList, QString::fromUtf8(filterAddresses.data, filterAddresses.len).split("¡¦!", QString::SkipEmptyParts)));
	return ({ QList<TransactionDetailsecff1c*>* tmpValue = new QList<TransactionDetailsecff1c*>(returnArg); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList HistoryManager554044_LoadHistory(void* ptr)
{
	QList<TransactionDetailsecff1c*> returnArg;
	QMetaObject::invokeMethod(static_cast<HistoryManager554044*>(ptr), "loadHistory", Q_RETURN_ARG(QList<TransactionDetailsecff1c*>, returnArg));
	return ({ QList<TransactionDetailsecff1c*>* tmpValue = new QList<TransactionDetailsecff1c*>(returnArg); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int HistoryManager554044_HistoryManager554044_QRegisterMetaType()
{
	return qRegisterMetaType<HistoryManager554044*>();
}

int HistoryManager554044_HistoryManager554044_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<HistoryManager554044*>(const_cast<const char*>(typeName));
}

int HistoryManager554044_HistoryManager554044_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<HistoryManager554044>();
#else
	return 0;
#endif
}

int HistoryManager554044_HistoryManager554044_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<HistoryManager554044>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* HistoryManager554044___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryManager554044___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* HistoryManager554044___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* HistoryManager554044___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void HistoryManager554044___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* HistoryManager554044___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* HistoryManager554044___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryManager554044___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* HistoryManager554044___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* HistoryManager554044___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryManager554044___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* HistoryManager554044___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* HistoryManager554044___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryManager554044___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* HistoryManager554044___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* HistoryManager554044___loadHistoryWithFilters_atList(void* ptr, int i)
{
	return ({TransactionDetailsecff1c* tmp = static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->at(i); if (i == static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->size()-1) { static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryManager554044___loadHistoryWithFilters_setList(void* ptr, void* i)
{
	static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->append(static_cast<TransactionDetailsecff1c*>(i));
}

void* HistoryManager554044___loadHistoryWithFilters_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<TransactionDetailsecff1c*>();
}

void* HistoryManager554044___loadHistory_atList(void* ptr, int i)
{
	return ({TransactionDetailsecff1c* tmp = static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->at(i); if (i == static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->size()-1) { static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryManager554044___loadHistory_setList(void* ptr, void* i)
{
	static_cast<QList<TransactionDetailsecff1c*>*>(ptr)->append(static_cast<TransactionDetailsecff1c*>(i));
}

void* HistoryManager554044___loadHistory_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<TransactionDetailsecff1c*>();
}

void* HistoryManager554044_NewHistoryManager(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new HistoryManager554044(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new HistoryManager554044(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new HistoryManager554044(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new HistoryManager554044(static_cast<QWindow*>(parent));
	} else {
		return new HistoryManager554044(static_cast<QObject*>(parent));
	}
}

void HistoryManager554044_DestroyHistoryManager(void* ptr)
{
	static_cast<HistoryManager554044*>(ptr)->~HistoryManager554044();
}

void HistoryManager554044_DestroyHistoryManagerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void HistoryManager554044_ChildEventDefault(void* ptr, void* event)
{
	static_cast<HistoryManager554044*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void HistoryManager554044_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<HistoryManager554044*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void HistoryManager554044_CustomEventDefault(void* ptr, void* event)
{
	static_cast<HistoryManager554044*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void HistoryManager554044_DeleteLaterDefault(void* ptr)
{
	static_cast<HistoryManager554044*>(ptr)->QObject::deleteLater();
}

void HistoryManager554044_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<HistoryManager554044*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char HistoryManager554044_EventDefault(void* ptr, void* e)
{
	return static_cast<HistoryManager554044*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char HistoryManager554044_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<HistoryManager554044*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void HistoryManager554044_TimerEventDefault(void* ptr, void* event)
{
	static_cast<HistoryManager554044*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
