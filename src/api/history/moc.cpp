

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
class HistoryModel8ba275: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<TransactionDetails8ba275*> transactions READ transactions WRITE setTransactions NOTIFY transactionsChanged)
public:
	HistoryModel8ba275(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");HistoryModel8ba275_HistoryModel8ba275_QRegisterMetaType();HistoryModel8ba275_HistoryModel8ba275_QRegisterMetaTypes();callbackHistoryModel8ba275_Constructor(this);};
	void Signal_AddTransaction(TransactionDetails8ba275* transaction) { callbackHistoryModel8ba275_AddTransaction(this, transaction); };
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbackHistoryModel8ba275_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbackHistoryModel8ba275_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbackHistoryModel8ba275_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<TransactionDetails8ba275*> transactions() { return ({ QList<TransactionDetails8ba275*>* tmpP = static_cast<QList<TransactionDetails8ba275*>*>(callbackHistoryModel8ba275_Transactions(this)); QList<TransactionDetails8ba275*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setTransactions(QList<TransactionDetails8ba275*> transactions) { callbackHistoryModel8ba275_SetTransactions(this, ({ QList<TransactionDetails8ba275*>* tmpValue = new QList<TransactionDetails8ba275*>(transactions); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_TransactionsChanged(QList<TransactionDetails8ba275*> transactions) { callbackHistoryModel8ba275_TransactionsChanged(this, ({ QList<TransactionDetails8ba275*>* tmpValue = new QList<TransactionDetails8ba275*>(transactions); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	 ~HistoryModel8ba275() { callbackHistoryModel8ba275_DestroyHistoryModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackHistoryModel8ba275_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackHistoryModel8ba275_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackHistoryModel8ba275_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackHistoryModel8ba275_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackHistoryModel8ba275_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackHistoryModel8ba275_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackHistoryModel8ba275_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackHistoryModel8ba275_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackHistoryModel8ba275_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackHistoryModel8ba275_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackHistoryModel8ba275_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackHistoryModel8ba275_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackHistoryModel8ba275_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackHistoryModel8ba275_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackHistoryModel8ba275_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackHistoryModel8ba275_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = new QVector<int>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackHistoryModel8ba275_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackHistoryModel8ba275_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackHistoryModel8ba275_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackHistoryModel8ba275_HeaderDataChanged(this, orientation, first, last); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackHistoryModel8ba275_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackHistoryModel8ba275_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackHistoryModel8ba275_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackHistoryModel8ba275_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackHistoryModel8ba275_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackHistoryModel8ba275_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackHistoryModel8ba275_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackHistoryModel8ba275_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackHistoryModel8ba275_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackHistoryModel8ba275_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackHistoryModel8ba275_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackHistoryModel8ba275_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackHistoryModel8ba275_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackHistoryModel8ba275_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackHistoryModel8ba275_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackHistoryModel8ba275_ResetInternalData(this); };
	void revert() { callbackHistoryModel8ba275_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackHistoryModel8ba275_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackHistoryModel8ba275_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackHistoryModel8ba275_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackHistoryModel8ba275_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackHistoryModel8ba275_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackHistoryModel8ba275_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackHistoryModel8ba275_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackHistoryModel8ba275_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackHistoryModel8ba275_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackHistoryModel8ba275_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackHistoryModel8ba275_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackHistoryModel8ba275_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackHistoryModel8ba275_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackHistoryModel8ba275_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackHistoryModel8ba275_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackHistoryModel8ba275_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackHistoryModel8ba275_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackHistoryModel8ba275_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackHistoryModel8ba275_CustomEvent(this, event); };
	void deleteLater() { callbackHistoryModel8ba275_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackHistoryModel8ba275_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackHistoryModel8ba275_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackHistoryModel8ba275_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackHistoryModel8ba275_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackHistoryModel8ba275_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackHistoryModel8ba275_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<TransactionDetails8ba275*> transactionsDefault() { return _transactions; };
	void setTransactionsDefault(QList<TransactionDetails8ba275*> p) { if (p != _transactions) { _transactions = p; transactionsChanged(_transactions); } };
signals:
	void addTransaction(TransactionDetails8ba275* transaction);
	void rolesChanged(type378cdd roles);
	void transactionsChanged(QList<TransactionDetails8ba275*> transactions);
public slots:
private:
	type378cdd _roles;
	QList<TransactionDetails8ba275*> _transactions;
};

Q_DECLARE_METATYPE(HistoryModel8ba275*)


void HistoryModel8ba275_HistoryModel8ba275_QRegisterMetaTypes() {
	qRegisterMetaType<type378cdd>("type378cdd");
	qRegisterMetaType<QList<QObject*>>("QList<TransactionDetails8ba275*>");
}

class TransactionDetails8ba275: public QObject
{
Q_OBJECT
Q_PROPERTY(QString date READ date WRITE setDate NOTIFY dateChanged)
Q_PROPERTY(qint32 status READ status WRITE setStatus NOTIFY statusChanged)
Q_PROPERTY(qint32 type READ type WRITE setType NOTIFY typeChanged)
Q_PROPERTY(qint32 amount READ amount WRITE setAmount NOTIFY amountChanged)
Q_PROPERTY(qint32 hoursReceived READ hoursReceived WRITE setHoursReceived NOTIFY hoursReceivedChanged)
Q_PROPERTY(qint32 hoursBurned READ hoursBurned WRITE setHoursBurned NOTIFY hoursBurnedChanged)
Q_PROPERTY(QString transactionID READ transactionID WRITE setTransactionID NOTIFY transactionIDChanged)
Q_PROPERTY(QString sentAddress READ sentAddress WRITE setSentAddress NOTIFY sentAddressChanged)
Q_PROPERTY(QString receivedAddress READ receivedAddress WRITE setReceivedAddress NOTIFY receivedAddressChanged)
public:
	TransactionDetails8ba275(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");TransactionDetails8ba275_TransactionDetails8ba275_QRegisterMetaType();TransactionDetails8ba275_TransactionDetails8ba275_QRegisterMetaTypes();callbackTransactionDetails8ba275_Constructor(this);};
	QString date() { return ({ Moc_PackedString tempVal = callbackTransactionDetails8ba275_Date(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setDate(QString date) { QByteArray te927d0 = date.toUtf8(); Moc_PackedString datePacked = { const_cast<char*>(te927d0.prepend("WHITESPACE").constData()+10), te927d0.size()-10 };callbackTransactionDetails8ba275_SetDate(this, datePacked); };
	void Signal_DateChanged(QString date) { QByteArray te927d0 = date.toUtf8(); Moc_PackedString datePacked = { const_cast<char*>(te927d0.prepend("WHITESPACE").constData()+10), te927d0.size()-10 };callbackTransactionDetails8ba275_DateChanged(this, datePacked); };
	qint32 status() { return callbackTransactionDetails8ba275_Status(this); };
	void setStatus(qint32 status) { callbackTransactionDetails8ba275_SetStatus(this, status); };
	void Signal_StatusChanged(qint32 status) { callbackTransactionDetails8ba275_StatusChanged(this, status); };
	qint32 type() { return callbackTransactionDetails8ba275_Type(this); };
	void setType(qint32 ty) { callbackTransactionDetails8ba275_SetType(this, ty); };
	void Signal_TypeChanged(qint32 ty) { callbackTransactionDetails8ba275_TypeChanged(this, ty); };
	qint32 amount() { return callbackTransactionDetails8ba275_Amount(this); };
	void setAmount(qint32 amount) { callbackTransactionDetails8ba275_SetAmount(this, amount); };
	void Signal_AmountChanged(qint32 amount) { callbackTransactionDetails8ba275_AmountChanged(this, amount); };
	qint32 hoursReceived() { return callbackTransactionDetails8ba275_HoursReceived(this); };
	void setHoursReceived(qint32 hoursReceived) { callbackTransactionDetails8ba275_SetHoursReceived(this, hoursReceived); };
	void Signal_HoursReceivedChanged(qint32 hoursReceived) { callbackTransactionDetails8ba275_HoursReceivedChanged(this, hoursReceived); };
	qint32 hoursBurned() { return callbackTransactionDetails8ba275_HoursBurned(this); };
	void setHoursBurned(qint32 hoursBurned) { callbackTransactionDetails8ba275_SetHoursBurned(this, hoursBurned); };
	void Signal_HoursBurnedChanged(qint32 hoursBurned) { callbackTransactionDetails8ba275_HoursBurnedChanged(this, hoursBurned); };
	QString transactionID() { return ({ Moc_PackedString tempVal = callbackTransactionDetails8ba275_TransactionID(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setTransactionID(QString transactionID) { QByteArray tec2ac1 = transactionID.toUtf8(); Moc_PackedString transactionIDPacked = { const_cast<char*>(tec2ac1.prepend("WHITESPACE").constData()+10), tec2ac1.size()-10 };callbackTransactionDetails8ba275_SetTransactionID(this, transactionIDPacked); };
	void Signal_TransactionIDChanged(QString transactionID) { QByteArray tec2ac1 = transactionID.toUtf8(); Moc_PackedString transactionIDPacked = { const_cast<char*>(tec2ac1.prepend("WHITESPACE").constData()+10), tec2ac1.size()-10 };callbackTransactionDetails8ba275_TransactionIDChanged(this, transactionIDPacked); };
	QString sentAddress() { return ({ Moc_PackedString tempVal = callbackTransactionDetails8ba275_SentAddress(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setSentAddress(QString sentAddress) { QByteArray ted7223 = sentAddress.toUtf8(); Moc_PackedString sentAddressPacked = { const_cast<char*>(ted7223.prepend("WHITESPACE").constData()+10), ted7223.size()-10 };callbackTransactionDetails8ba275_SetSentAddress(this, sentAddressPacked); };
	void Signal_SentAddressChanged(QString sentAddress) { QByteArray ted7223 = sentAddress.toUtf8(); Moc_PackedString sentAddressPacked = { const_cast<char*>(ted7223.prepend("WHITESPACE").constData()+10), ted7223.size()-10 };callbackTransactionDetails8ba275_SentAddressChanged(this, sentAddressPacked); };
	QString receivedAddress() { return ({ Moc_PackedString tempVal = callbackTransactionDetails8ba275_ReceivedAddress(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setReceivedAddress(QString receivedAddress) { QByteArray t55bc8c = receivedAddress.toUtf8(); Moc_PackedString receivedAddressPacked = { const_cast<char*>(t55bc8c.prepend("WHITESPACE").constData()+10), t55bc8c.size()-10 };callbackTransactionDetails8ba275_SetReceivedAddress(this, receivedAddressPacked); };
	void Signal_ReceivedAddressChanged(QString receivedAddress) { QByteArray t55bc8c = receivedAddress.toUtf8(); Moc_PackedString receivedAddressPacked = { const_cast<char*>(t55bc8c.prepend("WHITESPACE").constData()+10), t55bc8c.size()-10 };callbackTransactionDetails8ba275_ReceivedAddressChanged(this, receivedAddressPacked); };
	 ~TransactionDetails8ba275() { callbackTransactionDetails8ba275_DestroyTransactionDetails(this); };
	void childEvent(QChildEvent * event) { callbackTransactionDetails8ba275_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackTransactionDetails8ba275_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackTransactionDetails8ba275_CustomEvent(this, event); };
	void deleteLater() { callbackTransactionDetails8ba275_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackTransactionDetails8ba275_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackTransactionDetails8ba275_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackTransactionDetails8ba275_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackTransactionDetails8ba275_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackTransactionDetails8ba275_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackTransactionDetails8ba275_TimerEvent(this, event); };
	QString dateDefault() { return _date; };
	void setDateDefault(QString p) { if (p != _date) { _date = p; dateChanged(_date); } };
	qint32 statusDefault() { return _status; };
	void setStatusDefault(qint32 p) { if (p != _status) { _status = p; statusChanged(_status); } };
	qint32 typeDefault() { return _type; };
	void setTypeDefault(qint32 p) { if (p != _type) { _type = p; typeChanged(_type); } };
	qint32 amountDefault() { return _amount; };
	void setAmountDefault(qint32 p) { if (p != _amount) { _amount = p; amountChanged(_amount); } };
	qint32 hoursReceivedDefault() { return _hoursReceived; };
	void setHoursReceivedDefault(qint32 p) { if (p != _hoursReceived) { _hoursReceived = p; hoursReceivedChanged(_hoursReceived); } };
	qint32 hoursBurnedDefault() { return _hoursBurned; };
	void setHoursBurnedDefault(qint32 p) { if (p != _hoursBurned) { _hoursBurned = p; hoursBurnedChanged(_hoursBurned); } };
	QString transactionIDDefault() { return _transactionID; };
	void setTransactionIDDefault(QString p) { if (p != _transactionID) { _transactionID = p; transactionIDChanged(_transactionID); } };
	QString sentAddressDefault() { return _sentAddress; };
	void setSentAddressDefault(QString p) { if (p != _sentAddress) { _sentAddress = p; sentAddressChanged(_sentAddress); } };
	QString receivedAddressDefault() { return _receivedAddress; };
	void setReceivedAddressDefault(QString p) { if (p != _receivedAddress) { _receivedAddress = p; receivedAddressChanged(_receivedAddress); } };
signals:
	void dateChanged(QString date);
	void statusChanged(qint32 status);
	void typeChanged(qint32 ty);
	void amountChanged(qint32 amount);
	void hoursReceivedChanged(qint32 hoursReceived);
	void hoursBurnedChanged(qint32 hoursBurned);
	void transactionIDChanged(QString transactionID);
	void sentAddressChanged(QString sentAddress);
	void receivedAddressChanged(QString receivedAddress);
public slots:
private:
	QString _date;
	qint32 _status;
	qint32 _type;
	qint32 _amount;
	qint32 _hoursReceived;
	qint32 _hoursBurned;
	QString _transactionID;
	QString _sentAddress;
	QString _receivedAddress;
};

Q_DECLARE_METATYPE(TransactionDetails8ba275*)


void TransactionDetails8ba275_TransactionDetails8ba275_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

void HistoryModel8ba275_ConnectAddTransaction(void* ptr)
{
	QObject::connect(static_cast<HistoryModel8ba275*>(ptr), static_cast<void (HistoryModel8ba275::*)(TransactionDetails8ba275*)>(&HistoryModel8ba275::addTransaction), static_cast<HistoryModel8ba275*>(ptr), static_cast<void (HistoryModel8ba275::*)(TransactionDetails8ba275*)>(&HistoryModel8ba275::Signal_AddTransaction));
}

void HistoryModel8ba275_DisconnectAddTransaction(void* ptr)
{
	QObject::disconnect(static_cast<HistoryModel8ba275*>(ptr), static_cast<void (HistoryModel8ba275::*)(TransactionDetails8ba275*)>(&HistoryModel8ba275::addTransaction), static_cast<HistoryModel8ba275*>(ptr), static_cast<void (HistoryModel8ba275::*)(TransactionDetails8ba275*)>(&HistoryModel8ba275::Signal_AddTransaction));
}

void HistoryModel8ba275_AddTransaction(void* ptr, void* transaction)
{
	static_cast<HistoryModel8ba275*>(ptr)->addTransaction(static_cast<TransactionDetails8ba275*>(transaction));
}

struct Moc_PackedList HistoryModel8ba275_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<HistoryModel8ba275*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList HistoryModel8ba275_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<HistoryModel8ba275*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void HistoryModel8ba275_SetRoles(void* ptr, void* roles)
{
	static_cast<HistoryModel8ba275*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void HistoryModel8ba275_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<HistoryModel8ba275*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void HistoryModel8ba275_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<HistoryModel8ba275*>(ptr), static_cast<void (HistoryModel8ba275::*)(QMap<qint32, QByteArray>)>(&HistoryModel8ba275::rolesChanged), static_cast<HistoryModel8ba275*>(ptr), static_cast<void (HistoryModel8ba275::*)(QMap<qint32, QByteArray>)>(&HistoryModel8ba275::Signal_RolesChanged));
}

void HistoryModel8ba275_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<HistoryModel8ba275*>(ptr), static_cast<void (HistoryModel8ba275::*)(QMap<qint32, QByteArray>)>(&HistoryModel8ba275::rolesChanged), static_cast<HistoryModel8ba275*>(ptr), static_cast<void (HistoryModel8ba275::*)(QMap<qint32, QByteArray>)>(&HistoryModel8ba275::Signal_RolesChanged));
}

void HistoryModel8ba275_RolesChanged(void* ptr, void* roles)
{
	static_cast<HistoryModel8ba275*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList HistoryModel8ba275_Transactions(void* ptr)
{
	return ({ QList<TransactionDetails8ba275*>* tmpValue = new QList<TransactionDetails8ba275*>(static_cast<HistoryModel8ba275*>(ptr)->transactions()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList HistoryModel8ba275_TransactionsDefault(void* ptr)
{
	return ({ QList<TransactionDetails8ba275*>* tmpValue = new QList<TransactionDetails8ba275*>(static_cast<HistoryModel8ba275*>(ptr)->transactionsDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void HistoryModel8ba275_SetTransactions(void* ptr, void* transactions)
{
	static_cast<HistoryModel8ba275*>(ptr)->setTransactions(({ QList<TransactionDetails8ba275*>* tmpP = static_cast<QList<TransactionDetails8ba275*>*>(transactions); QList<TransactionDetails8ba275*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void HistoryModel8ba275_SetTransactionsDefault(void* ptr, void* transactions)
{
	static_cast<HistoryModel8ba275*>(ptr)->setTransactionsDefault(({ QList<TransactionDetails8ba275*>* tmpP = static_cast<QList<TransactionDetails8ba275*>*>(transactions); QList<TransactionDetails8ba275*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void HistoryModel8ba275_ConnectTransactionsChanged(void* ptr)
{
	QObject::connect(static_cast<HistoryModel8ba275*>(ptr), static_cast<void (HistoryModel8ba275::*)(QList<TransactionDetails8ba275*>)>(&HistoryModel8ba275::transactionsChanged), static_cast<HistoryModel8ba275*>(ptr), static_cast<void (HistoryModel8ba275::*)(QList<TransactionDetails8ba275*>)>(&HistoryModel8ba275::Signal_TransactionsChanged));
}

void HistoryModel8ba275_DisconnectTransactionsChanged(void* ptr)
{
	QObject::disconnect(static_cast<HistoryModel8ba275*>(ptr), static_cast<void (HistoryModel8ba275::*)(QList<TransactionDetails8ba275*>)>(&HistoryModel8ba275::transactionsChanged), static_cast<HistoryModel8ba275*>(ptr), static_cast<void (HistoryModel8ba275::*)(QList<TransactionDetails8ba275*>)>(&HistoryModel8ba275::Signal_TransactionsChanged));
}

void HistoryModel8ba275_TransactionsChanged(void* ptr, void* transactions)
{
	static_cast<HistoryModel8ba275*>(ptr)->transactionsChanged(({ QList<TransactionDetails8ba275*>* tmpP = static_cast<QList<TransactionDetails8ba275*>*>(transactions); QList<TransactionDetails8ba275*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int HistoryModel8ba275_HistoryModel8ba275_QRegisterMetaType()
{
	return qRegisterMetaType<HistoryModel8ba275*>();
}

int HistoryModel8ba275_HistoryModel8ba275_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<HistoryModel8ba275*>(const_cast<const char*>(typeName));
}

int HistoryModel8ba275_HistoryModel8ba275_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<HistoryModel8ba275>();
#else
	return 0;
#endif
}

int HistoryModel8ba275_HistoryModel8ba275_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<HistoryModel8ba275>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int HistoryModel8ba275_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel8ba275_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* HistoryModel8ba275_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int HistoryModel8ba275_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel8ba275_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* HistoryModel8ba275_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int HistoryModel8ba275_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel8ba275_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* HistoryModel8ba275_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* HistoryModel8ba275___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void HistoryModel8ba275___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* HistoryModel8ba275___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* HistoryModel8ba275___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void HistoryModel8ba275___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* HistoryModel8ba275___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int HistoryModel8ba275___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void HistoryModel8ba275___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* HistoryModel8ba275___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* HistoryModel8ba275___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void HistoryModel8ba275___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* HistoryModel8ba275___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList HistoryModel8ba275___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* HistoryModel8ba275___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void HistoryModel8ba275___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* HistoryModel8ba275___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* HistoryModel8ba275___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void HistoryModel8ba275___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* HistoryModel8ba275___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* HistoryModel8ba275___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void HistoryModel8ba275___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* HistoryModel8ba275___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* HistoryModel8ba275___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void HistoryModel8ba275___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* HistoryModel8ba275___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* HistoryModel8ba275___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void HistoryModel8ba275___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* HistoryModel8ba275___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* HistoryModel8ba275___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void HistoryModel8ba275___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* HistoryModel8ba275___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList HistoryModel8ba275___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* HistoryModel8ba275___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void HistoryModel8ba275___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* HistoryModel8ba275___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList HistoryModel8ba275___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int HistoryModel8ba275_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel8ba275_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* HistoryModel8ba275_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int HistoryModel8ba275_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel8ba275_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* HistoryModel8ba275_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* HistoryModel8ba275___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel8ba275___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* HistoryModel8ba275___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* HistoryModel8ba275___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void HistoryModel8ba275___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* HistoryModel8ba275___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* HistoryModel8ba275___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel8ba275___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* HistoryModel8ba275___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* HistoryModel8ba275___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel8ba275___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* HistoryModel8ba275___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* HistoryModel8ba275___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel8ba275___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* HistoryModel8ba275___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* HistoryModel8ba275___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void HistoryModel8ba275___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* HistoryModel8ba275___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList HistoryModel8ba275___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* HistoryModel8ba275___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void HistoryModel8ba275___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* HistoryModel8ba275___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList HistoryModel8ba275___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* HistoryModel8ba275___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void HistoryModel8ba275___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* HistoryModel8ba275___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList HistoryModel8ba275___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* HistoryModel8ba275___transactions_atList(void* ptr, int i)
{
	return ({TransactionDetails8ba275* tmp = static_cast<QList<TransactionDetails8ba275*>*>(ptr)->at(i); if (i == static_cast<QList<TransactionDetails8ba275*>*>(ptr)->size()-1) { static_cast<QList<TransactionDetails8ba275*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel8ba275___transactions_setList(void* ptr, void* i)
{
	static_cast<QList<TransactionDetails8ba275*>*>(ptr)->append(static_cast<TransactionDetails8ba275*>(i));
}

void* HistoryModel8ba275___transactions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<TransactionDetails8ba275*>();
}

void* HistoryModel8ba275___setTransactions_transactions_atList(void* ptr, int i)
{
	return ({TransactionDetails8ba275* tmp = static_cast<QList<TransactionDetails8ba275*>*>(ptr)->at(i); if (i == static_cast<QList<TransactionDetails8ba275*>*>(ptr)->size()-1) { static_cast<QList<TransactionDetails8ba275*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel8ba275___setTransactions_transactions_setList(void* ptr, void* i)
{
	static_cast<QList<TransactionDetails8ba275*>*>(ptr)->append(static_cast<TransactionDetails8ba275*>(i));
}

void* HistoryModel8ba275___setTransactions_transactions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<TransactionDetails8ba275*>();
}

void* HistoryModel8ba275___transactionsChanged_transactions_atList(void* ptr, int i)
{
	return ({TransactionDetails8ba275* tmp = static_cast<QList<TransactionDetails8ba275*>*>(ptr)->at(i); if (i == static_cast<QList<TransactionDetails8ba275*>*>(ptr)->size()-1) { static_cast<QList<TransactionDetails8ba275*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel8ba275___transactionsChanged_transactions_setList(void* ptr, void* i)
{
	static_cast<QList<TransactionDetails8ba275*>*>(ptr)->append(static_cast<TransactionDetails8ba275*>(i));
}

void* HistoryModel8ba275___transactionsChanged_transactions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<TransactionDetails8ba275*>();
}

int HistoryModel8ba275_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel8ba275_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* HistoryModel8ba275_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int HistoryModel8ba275_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel8ba275_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* HistoryModel8ba275_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int HistoryModel8ba275_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel8ba275_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* HistoryModel8ba275_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* HistoryModel8ba275_NewHistoryModel(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new HistoryModel8ba275(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new HistoryModel8ba275(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new HistoryModel8ba275(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new HistoryModel8ba275(static_cast<QWindow*>(parent));
	} else {
		return new HistoryModel8ba275(static_cast<QObject*>(parent));
	}
}

void HistoryModel8ba275_DestroyHistoryModel(void* ptr)
{
	static_cast<HistoryModel8ba275*>(ptr)->~HistoryModel8ba275();
}

void HistoryModel8ba275_DestroyHistoryModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char HistoryModel8ba275_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

long long HistoryModel8ba275_FlagsDefault(void* ptr, void* index)
{
	return static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

void* HistoryModel8ba275_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* HistoryModel8ba275_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* HistoryModel8ba275_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

char HistoryModel8ba275_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char HistoryModel8ba275_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int HistoryModel8ba275_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* HistoryModel8ba275_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void HistoryModel8ba275_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char HistoryModel8ba275_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* HistoryModel8ba275_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

char HistoryModel8ba275_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char HistoryModel8ba275_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct Moc_PackedList HistoryModel8ba275_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList HistoryModel8ba275_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* HistoryModel8ba275_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct Moc_PackedString HistoryModel8ba275_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t2ba84e = static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::mimeTypes().join("¡¦!").toUtf8(); Moc_PackedString { const_cast<char*>(t2ba84e.prepend("WHITESPACE").constData()+10), t2ba84e.size()-10 }; });
}

char HistoryModel8ba275_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char HistoryModel8ba275_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* HistoryModel8ba275_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

char HistoryModel8ba275_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char HistoryModel8ba275_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void HistoryModel8ba275_ResetInternalDataDefault(void* ptr)
{
	static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::resetInternalData();
}

void HistoryModel8ba275_RevertDefault(void* ptr)
{
	static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::revert();
}

struct Moc_PackedList HistoryModel8ba275_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int HistoryModel8ba275_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char HistoryModel8ba275_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char HistoryModel8ba275_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char HistoryModel8ba275_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void HistoryModel8ba275_SortDefault(void* ptr, int column, long long order)
{
	static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* HistoryModel8ba275_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char HistoryModel8ba275_SubmitDefault(void* ptr)
{
	return static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::submit();
}

long long HistoryModel8ba275_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long HistoryModel8ba275_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::supportedDropActions();
}

void HistoryModel8ba275_ChildEventDefault(void* ptr, void* event)
{
	static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void HistoryModel8ba275_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void HistoryModel8ba275_CustomEventDefault(void* ptr, void* event)
{
	static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void HistoryModel8ba275_DeleteLaterDefault(void* ptr)
{
	static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::deleteLater();
}

void HistoryModel8ba275_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char HistoryModel8ba275_EventDefault(void* ptr, void* e)
{
	return static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char HistoryModel8ba275_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void HistoryModel8ba275_TimerEventDefault(void* ptr, void* event)
{
	static_cast<HistoryModel8ba275*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

struct Moc_PackedString TransactionDetails8ba275_Date(void* ptr)
{
	return ({ QByteArray t3aa0fc = static_cast<TransactionDetails8ba275*>(ptr)->date().toUtf8(); Moc_PackedString { const_cast<char*>(t3aa0fc.prepend("WHITESPACE").constData()+10), t3aa0fc.size()-10 }; });
}

struct Moc_PackedString TransactionDetails8ba275_DateDefault(void* ptr)
{
	return ({ QByteArray tff1e84 = static_cast<TransactionDetails8ba275*>(ptr)->dateDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tff1e84.prepend("WHITESPACE").constData()+10), tff1e84.size()-10 }; });
}

void TransactionDetails8ba275_SetDate(void* ptr, struct Moc_PackedString date)
{
	static_cast<TransactionDetails8ba275*>(ptr)->setDate(QString::fromUtf8(date.data, date.len));
}

void TransactionDetails8ba275_SetDateDefault(void* ptr, struct Moc_PackedString date)
{
	static_cast<TransactionDetails8ba275*>(ptr)->setDateDefault(QString::fromUtf8(date.data, date.len));
}

void TransactionDetails8ba275_ConnectDateChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(QString)>(&TransactionDetails8ba275::dateChanged), static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(QString)>(&TransactionDetails8ba275::Signal_DateChanged));
}

void TransactionDetails8ba275_DisconnectDateChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(QString)>(&TransactionDetails8ba275::dateChanged), static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(QString)>(&TransactionDetails8ba275::Signal_DateChanged));
}

void TransactionDetails8ba275_DateChanged(void* ptr, struct Moc_PackedString date)
{
	static_cast<TransactionDetails8ba275*>(ptr)->dateChanged(QString::fromUtf8(date.data, date.len));
}

int TransactionDetails8ba275_Status(void* ptr)
{
	return static_cast<TransactionDetails8ba275*>(ptr)->status();
}

int TransactionDetails8ba275_StatusDefault(void* ptr)
{
	return static_cast<TransactionDetails8ba275*>(ptr)->statusDefault();
}

void TransactionDetails8ba275_SetStatus(void* ptr, int status)
{
	static_cast<TransactionDetails8ba275*>(ptr)->setStatus(status);
}

void TransactionDetails8ba275_SetStatusDefault(void* ptr, int status)
{
	static_cast<TransactionDetails8ba275*>(ptr)->setStatusDefault(status);
}

void TransactionDetails8ba275_ConnectStatusChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(qint32)>(&TransactionDetails8ba275::statusChanged), static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(qint32)>(&TransactionDetails8ba275::Signal_StatusChanged));
}

void TransactionDetails8ba275_DisconnectStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(qint32)>(&TransactionDetails8ba275::statusChanged), static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(qint32)>(&TransactionDetails8ba275::Signal_StatusChanged));
}

void TransactionDetails8ba275_StatusChanged(void* ptr, int status)
{
	static_cast<TransactionDetails8ba275*>(ptr)->statusChanged(status);
}

int TransactionDetails8ba275_Type(void* ptr)
{
	return static_cast<TransactionDetails8ba275*>(ptr)->type();
}

int TransactionDetails8ba275_TypeDefault(void* ptr)
{
	return static_cast<TransactionDetails8ba275*>(ptr)->typeDefault();
}

void TransactionDetails8ba275_SetType(void* ptr, int ty)
{
	static_cast<TransactionDetails8ba275*>(ptr)->setType(ty);
}

void TransactionDetails8ba275_SetTypeDefault(void* ptr, int ty)
{
	static_cast<TransactionDetails8ba275*>(ptr)->setTypeDefault(ty);
}

void TransactionDetails8ba275_ConnectTypeChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(qint32)>(&TransactionDetails8ba275::typeChanged), static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(qint32)>(&TransactionDetails8ba275::Signal_TypeChanged));
}

void TransactionDetails8ba275_DisconnectTypeChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(qint32)>(&TransactionDetails8ba275::typeChanged), static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(qint32)>(&TransactionDetails8ba275::Signal_TypeChanged));
}

void TransactionDetails8ba275_TypeChanged(void* ptr, int ty)
{
	static_cast<TransactionDetails8ba275*>(ptr)->typeChanged(ty);
}

int TransactionDetails8ba275_Amount(void* ptr)
{
	return static_cast<TransactionDetails8ba275*>(ptr)->amount();
}

int TransactionDetails8ba275_AmountDefault(void* ptr)
{
	return static_cast<TransactionDetails8ba275*>(ptr)->amountDefault();
}

void TransactionDetails8ba275_SetAmount(void* ptr, int amount)
{
	static_cast<TransactionDetails8ba275*>(ptr)->setAmount(amount);
}

void TransactionDetails8ba275_SetAmountDefault(void* ptr, int amount)
{
	static_cast<TransactionDetails8ba275*>(ptr)->setAmountDefault(amount);
}

void TransactionDetails8ba275_ConnectAmountChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(qint32)>(&TransactionDetails8ba275::amountChanged), static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(qint32)>(&TransactionDetails8ba275::Signal_AmountChanged));
}

void TransactionDetails8ba275_DisconnectAmountChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(qint32)>(&TransactionDetails8ba275::amountChanged), static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(qint32)>(&TransactionDetails8ba275::Signal_AmountChanged));
}

void TransactionDetails8ba275_AmountChanged(void* ptr, int amount)
{
	static_cast<TransactionDetails8ba275*>(ptr)->amountChanged(amount);
}

int TransactionDetails8ba275_HoursReceived(void* ptr)
{
	return static_cast<TransactionDetails8ba275*>(ptr)->hoursReceived();
}

int TransactionDetails8ba275_HoursReceivedDefault(void* ptr)
{
	return static_cast<TransactionDetails8ba275*>(ptr)->hoursReceivedDefault();
}

void TransactionDetails8ba275_SetHoursReceived(void* ptr, int hoursReceived)
{
	static_cast<TransactionDetails8ba275*>(ptr)->setHoursReceived(hoursReceived);
}

void TransactionDetails8ba275_SetHoursReceivedDefault(void* ptr, int hoursReceived)
{
	static_cast<TransactionDetails8ba275*>(ptr)->setHoursReceivedDefault(hoursReceived);
}

void TransactionDetails8ba275_ConnectHoursReceivedChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(qint32)>(&TransactionDetails8ba275::hoursReceivedChanged), static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(qint32)>(&TransactionDetails8ba275::Signal_HoursReceivedChanged));
}

void TransactionDetails8ba275_DisconnectHoursReceivedChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(qint32)>(&TransactionDetails8ba275::hoursReceivedChanged), static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(qint32)>(&TransactionDetails8ba275::Signal_HoursReceivedChanged));
}

void TransactionDetails8ba275_HoursReceivedChanged(void* ptr, int hoursReceived)
{
	static_cast<TransactionDetails8ba275*>(ptr)->hoursReceivedChanged(hoursReceived);
}

int TransactionDetails8ba275_HoursBurned(void* ptr)
{
	return static_cast<TransactionDetails8ba275*>(ptr)->hoursBurned();
}

int TransactionDetails8ba275_HoursBurnedDefault(void* ptr)
{
	return static_cast<TransactionDetails8ba275*>(ptr)->hoursBurnedDefault();
}

void TransactionDetails8ba275_SetHoursBurned(void* ptr, int hoursBurned)
{
	static_cast<TransactionDetails8ba275*>(ptr)->setHoursBurned(hoursBurned);
}

void TransactionDetails8ba275_SetHoursBurnedDefault(void* ptr, int hoursBurned)
{
	static_cast<TransactionDetails8ba275*>(ptr)->setHoursBurnedDefault(hoursBurned);
}

void TransactionDetails8ba275_ConnectHoursBurnedChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(qint32)>(&TransactionDetails8ba275::hoursBurnedChanged), static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(qint32)>(&TransactionDetails8ba275::Signal_HoursBurnedChanged));
}

void TransactionDetails8ba275_DisconnectHoursBurnedChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(qint32)>(&TransactionDetails8ba275::hoursBurnedChanged), static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(qint32)>(&TransactionDetails8ba275::Signal_HoursBurnedChanged));
}

void TransactionDetails8ba275_HoursBurnedChanged(void* ptr, int hoursBurned)
{
	static_cast<TransactionDetails8ba275*>(ptr)->hoursBurnedChanged(hoursBurned);
}

struct Moc_PackedString TransactionDetails8ba275_TransactionID(void* ptr)
{
	return ({ QByteArray tc8c5df = static_cast<TransactionDetails8ba275*>(ptr)->transactionID().toUtf8(); Moc_PackedString { const_cast<char*>(tc8c5df.prepend("WHITESPACE").constData()+10), tc8c5df.size()-10 }; });
}

struct Moc_PackedString TransactionDetails8ba275_TransactionIDDefault(void* ptr)
{
	return ({ QByteArray t299bde = static_cast<TransactionDetails8ba275*>(ptr)->transactionIDDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t299bde.prepend("WHITESPACE").constData()+10), t299bde.size()-10 }; });
}

void TransactionDetails8ba275_SetTransactionID(void* ptr, struct Moc_PackedString transactionID)
{
	static_cast<TransactionDetails8ba275*>(ptr)->setTransactionID(QString::fromUtf8(transactionID.data, transactionID.len));
}

void TransactionDetails8ba275_SetTransactionIDDefault(void* ptr, struct Moc_PackedString transactionID)
{
	static_cast<TransactionDetails8ba275*>(ptr)->setTransactionIDDefault(QString::fromUtf8(transactionID.data, transactionID.len));
}

void TransactionDetails8ba275_ConnectTransactionIDChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(QString)>(&TransactionDetails8ba275::transactionIDChanged), static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(QString)>(&TransactionDetails8ba275::Signal_TransactionIDChanged));
}

void TransactionDetails8ba275_DisconnectTransactionIDChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(QString)>(&TransactionDetails8ba275::transactionIDChanged), static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(QString)>(&TransactionDetails8ba275::Signal_TransactionIDChanged));
}

void TransactionDetails8ba275_TransactionIDChanged(void* ptr, struct Moc_PackedString transactionID)
{
	static_cast<TransactionDetails8ba275*>(ptr)->transactionIDChanged(QString::fromUtf8(transactionID.data, transactionID.len));
}

struct Moc_PackedString TransactionDetails8ba275_SentAddress(void* ptr)
{
	return ({ QByteArray te78448 = static_cast<TransactionDetails8ba275*>(ptr)->sentAddress().toUtf8(); Moc_PackedString { const_cast<char*>(te78448.prepend("WHITESPACE").constData()+10), te78448.size()-10 }; });
}

struct Moc_PackedString TransactionDetails8ba275_SentAddressDefault(void* ptr)
{
	return ({ QByteArray tf856c0 = static_cast<TransactionDetails8ba275*>(ptr)->sentAddressDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tf856c0.prepend("WHITESPACE").constData()+10), tf856c0.size()-10 }; });
}

void TransactionDetails8ba275_SetSentAddress(void* ptr, struct Moc_PackedString sentAddress)
{
	static_cast<TransactionDetails8ba275*>(ptr)->setSentAddress(QString::fromUtf8(sentAddress.data, sentAddress.len));
}

void TransactionDetails8ba275_SetSentAddressDefault(void* ptr, struct Moc_PackedString sentAddress)
{
	static_cast<TransactionDetails8ba275*>(ptr)->setSentAddressDefault(QString::fromUtf8(sentAddress.data, sentAddress.len));
}

void TransactionDetails8ba275_ConnectSentAddressChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(QString)>(&TransactionDetails8ba275::sentAddressChanged), static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(QString)>(&TransactionDetails8ba275::Signal_SentAddressChanged));
}

void TransactionDetails8ba275_DisconnectSentAddressChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(QString)>(&TransactionDetails8ba275::sentAddressChanged), static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(QString)>(&TransactionDetails8ba275::Signal_SentAddressChanged));
}

void TransactionDetails8ba275_SentAddressChanged(void* ptr, struct Moc_PackedString sentAddress)
{
	static_cast<TransactionDetails8ba275*>(ptr)->sentAddressChanged(QString::fromUtf8(sentAddress.data, sentAddress.len));
}

struct Moc_PackedString TransactionDetails8ba275_ReceivedAddress(void* ptr)
{
	return ({ QByteArray t3a4c8c = static_cast<TransactionDetails8ba275*>(ptr)->receivedAddress().toUtf8(); Moc_PackedString { const_cast<char*>(t3a4c8c.prepend("WHITESPACE").constData()+10), t3a4c8c.size()-10 }; });
}

struct Moc_PackedString TransactionDetails8ba275_ReceivedAddressDefault(void* ptr)
{
	return ({ QByteArray ta1cf30 = static_cast<TransactionDetails8ba275*>(ptr)->receivedAddressDefault().toUtf8(); Moc_PackedString { const_cast<char*>(ta1cf30.prepend("WHITESPACE").constData()+10), ta1cf30.size()-10 }; });
}

void TransactionDetails8ba275_SetReceivedAddress(void* ptr, struct Moc_PackedString receivedAddress)
{
	static_cast<TransactionDetails8ba275*>(ptr)->setReceivedAddress(QString::fromUtf8(receivedAddress.data, receivedAddress.len));
}

void TransactionDetails8ba275_SetReceivedAddressDefault(void* ptr, struct Moc_PackedString receivedAddress)
{
	static_cast<TransactionDetails8ba275*>(ptr)->setReceivedAddressDefault(QString::fromUtf8(receivedAddress.data, receivedAddress.len));
}

void TransactionDetails8ba275_ConnectReceivedAddressChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(QString)>(&TransactionDetails8ba275::receivedAddressChanged), static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(QString)>(&TransactionDetails8ba275::Signal_ReceivedAddressChanged));
}

void TransactionDetails8ba275_DisconnectReceivedAddressChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(QString)>(&TransactionDetails8ba275::receivedAddressChanged), static_cast<TransactionDetails8ba275*>(ptr), static_cast<void (TransactionDetails8ba275::*)(QString)>(&TransactionDetails8ba275::Signal_ReceivedAddressChanged));
}

void TransactionDetails8ba275_ReceivedAddressChanged(void* ptr, struct Moc_PackedString receivedAddress)
{
	static_cast<TransactionDetails8ba275*>(ptr)->receivedAddressChanged(QString::fromUtf8(receivedAddress.data, receivedAddress.len));
}

int TransactionDetails8ba275_TransactionDetails8ba275_QRegisterMetaType()
{
	return qRegisterMetaType<TransactionDetails8ba275*>();
}

int TransactionDetails8ba275_TransactionDetails8ba275_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<TransactionDetails8ba275*>(const_cast<const char*>(typeName));
}

int TransactionDetails8ba275_TransactionDetails8ba275_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TransactionDetails8ba275>();
#else
	return 0;
#endif
}

int TransactionDetails8ba275_TransactionDetails8ba275_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TransactionDetails8ba275>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* TransactionDetails8ba275___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionDetails8ba275___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TransactionDetails8ba275___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* TransactionDetails8ba275___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TransactionDetails8ba275___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* TransactionDetails8ba275___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* TransactionDetails8ba275___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionDetails8ba275___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TransactionDetails8ba275___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TransactionDetails8ba275___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionDetails8ba275___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TransactionDetails8ba275___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TransactionDetails8ba275___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionDetails8ba275___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TransactionDetails8ba275___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TransactionDetails8ba275_NewTransactionDetails(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new TransactionDetails8ba275(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new TransactionDetails8ba275(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new TransactionDetails8ba275(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new TransactionDetails8ba275(static_cast<QWindow*>(parent));
	} else {
		return new TransactionDetails8ba275(static_cast<QObject*>(parent));
	}
}

void TransactionDetails8ba275_DestroyTransactionDetails(void* ptr)
{
	static_cast<TransactionDetails8ba275*>(ptr)->~TransactionDetails8ba275();
}

void TransactionDetails8ba275_DestroyTransactionDetailsDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void TransactionDetails8ba275_ChildEventDefault(void* ptr, void* event)
{
	static_cast<TransactionDetails8ba275*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void TransactionDetails8ba275_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TransactionDetails8ba275*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void TransactionDetails8ba275_CustomEventDefault(void* ptr, void* event)
{
	static_cast<TransactionDetails8ba275*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void TransactionDetails8ba275_DeleteLaterDefault(void* ptr)
{
	static_cast<TransactionDetails8ba275*>(ptr)->QObject::deleteLater();
}

void TransactionDetails8ba275_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TransactionDetails8ba275*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char TransactionDetails8ba275_EventDefault(void* ptr, void* e)
{
	return static_cast<TransactionDetails8ba275*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char TransactionDetails8ba275_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<TransactionDetails8ba275*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void TransactionDetails8ba275_TimerEventDefault(void* ptr, void* event)
{
	static_cast<TransactionDetails8ba275*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
