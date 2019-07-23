

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAbstractItemModel>
#include <QAbstractListModel>
#include <QByteArray>
#include <QChildEvent>
#include <QDateTime>
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
typedef QMap<qint32, QByteArray> type378cdd;
class AddressDeatails742340: public QObject
{
Q_OBJECT
Q_PROPERTY(QString address READ address WRITE setAddress NOTIFY addressChanged)
Q_PROPERTY(float addressSky READ addressSky WRITE setAddressSky NOTIFY addressSkyChanged)
Q_PROPERTY(qint32 addressCoinHours READ addressCoinHours WRITE setAddressCoinHours NOTIFY addressCoinHoursChanged)
public:
	AddressDeatails742340(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");AddressDeatails742340_AddressDeatails742340_QRegisterMetaType();AddressDeatails742340_AddressDeatails742340_QRegisterMetaTypes();callbackAddressDeatails742340_Constructor(this);};
	QString address() { return ({ Moc_PackedString tempVal = callbackAddressDeatails742340_Address(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setAddress(QString address) { QByteArray tc66218 = address.toUtf8(); Moc_PackedString addressPacked = { const_cast<char*>(tc66218.prepend("WHITESPACE").constData()+10), tc66218.size()-10 };callbackAddressDeatails742340_SetAddress(this, addressPacked); };
	void Signal_AddressChanged(QString address) { QByteArray tc66218 = address.toUtf8(); Moc_PackedString addressPacked = { const_cast<char*>(tc66218.prepend("WHITESPACE").constData()+10), tc66218.size()-10 };callbackAddressDeatails742340_AddressChanged(this, addressPacked); };
	float addressSky() { return callbackAddressDeatails742340_AddressSky(this); };
	void setAddressSky(float addressSky) { callbackAddressDeatails742340_SetAddressSky(this, addressSky); };
	void Signal_AddressSkyChanged(float addressSky) { callbackAddressDeatails742340_AddressSkyChanged(this, addressSky); };
	qint32 addressCoinHours() { return callbackAddressDeatails742340_AddressCoinHours(this); };
	void setAddressCoinHours(qint32 addressCoinHours) { callbackAddressDeatails742340_SetAddressCoinHours(this, addressCoinHours); };
	void Signal_AddressCoinHoursChanged(qint32 addressCoinHours) { callbackAddressDeatails742340_AddressCoinHoursChanged(this, addressCoinHours); };
	 ~AddressDeatails742340() { callbackAddressDeatails742340_DestroyAddressDeatails(this); };
	void childEvent(QChildEvent * event) { callbackAddressDeatails742340_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackAddressDeatails742340_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackAddressDeatails742340_CustomEvent(this, event); };
	void deleteLater() { callbackAddressDeatails742340_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackAddressDeatails742340_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackAddressDeatails742340_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackAddressDeatails742340_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackAddressDeatails742340_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackAddressDeatails742340_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackAddressDeatails742340_TimerEvent(this, event); };
	QString addressDefault() { return _address; };
	void setAddressDefault(QString p) { if (p != _address) { _address = p; addressChanged(_address); } };
	float addressSkyDefault() { return _addressSky; };
	void setAddressSkyDefault(float p) { if (p != _addressSky) { _addressSky = p; addressSkyChanged(_addressSky); } };
	qint32 addressCoinHoursDefault() { return _addressCoinHours; };
	void setAddressCoinHoursDefault(qint32 p) { if (p != _addressCoinHours) { _addressCoinHours = p; addressCoinHoursChanged(_addressCoinHours); } };
signals:
	void addressChanged(QString address);
	void addressSkyChanged(float addressSky);
	void addressCoinHoursChanged(qint32 addressCoinHours);
public slots:
private:
	QString _address;
	float _addressSky;
	qint32 _addressCoinHours;
};

Q_DECLARE_METATYPE(AddressDeatails742340*)


void AddressDeatails742340_AddressDeatails742340_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class AddressList742340: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<AddressDeatails742340*> addresses READ addresses WRITE setAddresses NOTIFY addressesChanged)
public:
	AddressList742340(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");AddressList742340_AddressList742340_QRegisterMetaType();AddressList742340_AddressList742340_QRegisterMetaTypes();callbackAddressList742340_Constructor(this);};
	void Signal_AddAddress(AddressDeatails742340* transaction) { callbackAddressList742340_AddAddress(this, transaction); };
	void Signal_RemoveAddress(qint32 index) { callbackAddressList742340_RemoveAddress(this, index); };
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbackAddressList742340_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbackAddressList742340_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbackAddressList742340_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<AddressDeatails742340*> addresses() { return ({ QList<AddressDeatails742340*>* tmpP = static_cast<QList<AddressDeatails742340*>*>(callbackAddressList742340_Addresses(this)); QList<AddressDeatails742340*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setAddresses(QList<AddressDeatails742340*> addresses) { callbackAddressList742340_SetAddresses(this, ({ QList<AddressDeatails742340*>* tmpValue = new QList<AddressDeatails742340*>(addresses); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_AddressesChanged(QList<AddressDeatails742340*> addresses) { callbackAddressList742340_AddressesChanged(this, ({ QList<AddressDeatails742340*>* tmpValue = new QList<AddressDeatails742340*>(addresses); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	 ~AddressList742340() { callbackAddressList742340_DestroyAddressList(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackAddressList742340_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackAddressList742340_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackAddressList742340_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackAddressList742340_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackAddressList742340_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackAddressList742340_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackAddressList742340_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackAddressList742340_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackAddressList742340_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackAddressList742340_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackAddressList742340_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackAddressList742340_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackAddressList742340_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackAddressList742340_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackAddressList742340_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackAddressList742340_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = new QVector<int>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackAddressList742340_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackAddressList742340_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackAddressList742340_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackAddressList742340_HeaderDataChanged(this, orientation, first, last); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackAddressList742340_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackAddressList742340_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackAddressList742340_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackAddressList742340_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackAddressList742340_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackAddressList742340_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackAddressList742340_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackAddressList742340_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackAddressList742340_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackAddressList742340_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackAddressList742340_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackAddressList742340_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackAddressList742340_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackAddressList742340_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackAddressList742340_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackAddressList742340_ResetInternalData(this); };
	void revert() { callbackAddressList742340_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackAddressList742340_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackAddressList742340_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackAddressList742340_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackAddressList742340_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackAddressList742340_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackAddressList742340_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackAddressList742340_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackAddressList742340_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackAddressList742340_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackAddressList742340_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackAddressList742340_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackAddressList742340_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackAddressList742340_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackAddressList742340_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackAddressList742340_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackAddressList742340_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackAddressList742340_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackAddressList742340_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackAddressList742340_CustomEvent(this, event); };
	void deleteLater() { callbackAddressList742340_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackAddressList742340_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackAddressList742340_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackAddressList742340_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackAddressList742340_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackAddressList742340_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackAddressList742340_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<AddressDeatails742340*> addressesDefault() { return _addresses; };
	void setAddressesDefault(QList<AddressDeatails742340*> p) { if (p != _addresses) { _addresses = p; addressesChanged(_addresses); } };
signals:
	void addAddress(AddressDeatails742340* transaction);
	void removeAddress(qint32 index);
	void rolesChanged(type378cdd roles);
	void addressesChanged(QList<AddressDeatails742340*> addresses);
public slots:
private:
	type378cdd _roles;
	QList<AddressDeatails742340*> _addresses;
};

Q_DECLARE_METATYPE(AddressList742340*)


void AddressList742340_AddressList742340_QRegisterMetaTypes() {
	qRegisterMetaType<type378cdd>("type378cdd");
	qRegisterMetaType<QList<QObject*>>("QList<AddressDeatails742340*>");
}

class HistoryModel742340: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<TransactionDetails742340*> transactions READ transactions WRITE setTransactions NOTIFY transactionsChanged)
public:
	HistoryModel742340(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");HistoryModel742340_HistoryModel742340_QRegisterMetaType();HistoryModel742340_HistoryModel742340_QRegisterMetaTypes();callbackHistoryModel742340_Constructor(this);};
	void Signal_AddTransaction(TransactionDetails742340* transaction) { callbackHistoryModel742340_AddTransaction(this, transaction); };
	void Signal_RemoveTransaction(qint32 index) { callbackHistoryModel742340_RemoveTransaction(this, index); };
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbackHistoryModel742340_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbackHistoryModel742340_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbackHistoryModel742340_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<TransactionDetails742340*> transactions() { return ({ QList<TransactionDetails742340*>* tmpP = static_cast<QList<TransactionDetails742340*>*>(callbackHistoryModel742340_Transactions(this)); QList<TransactionDetails742340*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setTransactions(QList<TransactionDetails742340*> transactions) { callbackHistoryModel742340_SetTransactions(this, ({ QList<TransactionDetails742340*>* tmpValue = new QList<TransactionDetails742340*>(transactions); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_TransactionsChanged(QList<TransactionDetails742340*> transactions) { callbackHistoryModel742340_TransactionsChanged(this, ({ QList<TransactionDetails742340*>* tmpValue = new QList<TransactionDetails742340*>(transactions); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	 ~HistoryModel742340() { callbackHistoryModel742340_DestroyHistoryModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackHistoryModel742340_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackHistoryModel742340_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackHistoryModel742340_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackHistoryModel742340_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackHistoryModel742340_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackHistoryModel742340_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackHistoryModel742340_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackHistoryModel742340_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackHistoryModel742340_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackHistoryModel742340_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackHistoryModel742340_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackHistoryModel742340_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackHistoryModel742340_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackHistoryModel742340_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackHistoryModel742340_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackHistoryModel742340_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = new QVector<int>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackHistoryModel742340_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackHistoryModel742340_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackHistoryModel742340_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackHistoryModel742340_HeaderDataChanged(this, orientation, first, last); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackHistoryModel742340_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackHistoryModel742340_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackHistoryModel742340_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackHistoryModel742340_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackHistoryModel742340_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackHistoryModel742340_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackHistoryModel742340_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackHistoryModel742340_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackHistoryModel742340_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackHistoryModel742340_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackHistoryModel742340_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackHistoryModel742340_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackHistoryModel742340_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackHistoryModel742340_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackHistoryModel742340_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackHistoryModel742340_ResetInternalData(this); };
	void revert() { callbackHistoryModel742340_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackHistoryModel742340_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackHistoryModel742340_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackHistoryModel742340_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackHistoryModel742340_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackHistoryModel742340_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackHistoryModel742340_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackHistoryModel742340_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackHistoryModel742340_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackHistoryModel742340_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackHistoryModel742340_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackHistoryModel742340_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackHistoryModel742340_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackHistoryModel742340_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackHistoryModel742340_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackHistoryModel742340_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackHistoryModel742340_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackHistoryModel742340_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackHistoryModel742340_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackHistoryModel742340_CustomEvent(this, event); };
	void deleteLater() { callbackHistoryModel742340_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackHistoryModel742340_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackHistoryModel742340_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackHistoryModel742340_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackHistoryModel742340_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackHistoryModel742340_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackHistoryModel742340_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<TransactionDetails742340*> transactionsDefault() { return _transactions; };
	void setTransactionsDefault(QList<TransactionDetails742340*> p) { if (p != _transactions) { _transactions = p; transactionsChanged(_transactions); } };
signals:
	void addTransaction(TransactionDetails742340* transaction);
	void removeTransaction(qint32 index);
	void rolesChanged(type378cdd roles);
	void transactionsChanged(QList<TransactionDetails742340*> transactions);
public slots:
private:
	type378cdd _roles;
	QList<TransactionDetails742340*> _transactions;
};

Q_DECLARE_METATYPE(HistoryModel742340*)


void HistoryModel742340_HistoryModel742340_QRegisterMetaTypes() {
	qRegisterMetaType<type378cdd>("type378cdd");
	qRegisterMetaType<QList<QObject*>>("QList<TransactionDetails742340*>");
}

class TransactionDetails742340: public QObject
{
Q_OBJECT
Q_PROPERTY(QDateTime date READ date WRITE setDate NOTIFY dateChanged)
Q_PROPERTY(qint32 status READ status WRITE setStatus NOTIFY statusChanged)
Q_PROPERTY(qint32 type READ type WRITE setType NOTIFY typeChanged)
Q_PROPERTY(qint32 amount READ amount WRITE setAmount NOTIFY amountChanged)
Q_PROPERTY(qint32 hoursReceived READ hoursReceived WRITE setHoursReceived NOTIFY hoursReceivedChanged)
Q_PROPERTY(qint32 hoursBurned READ hoursBurned WRITE setHoursBurned NOTIFY hoursBurnedChanged)
Q_PROPERTY(QString transactionID READ transactionID WRITE setTransactionID NOTIFY transactionIDChanged)
Q_PROPERTY(QString sentAddress READ sentAddress WRITE setSentAddress NOTIFY sentAddressChanged)
Q_PROPERTY(QString receivedAddress READ receivedAddress WRITE setReceivedAddress NOTIFY receivedAddressChanged)
public:
	TransactionDetails742340(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");TransactionDetails742340_TransactionDetails742340_QRegisterMetaType();TransactionDetails742340_TransactionDetails742340_QRegisterMetaTypes();callbackTransactionDetails742340_Constructor(this);};
	QDateTime date() { return *static_cast<QDateTime*>(callbackTransactionDetails742340_Date(this)); };
	void setDate(QDateTime date) { callbackTransactionDetails742340_SetDate(this, new QDateTime(date)); };
	void Signal_DateChanged(QDateTime date) { callbackTransactionDetails742340_DateChanged(this, new QDateTime(date)); };
	qint32 status() { return callbackTransactionDetails742340_Status(this); };
	void setStatus(qint32 status) { callbackTransactionDetails742340_SetStatus(this, status); };
	void Signal_StatusChanged(qint32 status) { callbackTransactionDetails742340_StatusChanged(this, status); };
	qint32 type() { return callbackTransactionDetails742340_Type(this); };
	void setType(qint32 ty) { callbackTransactionDetails742340_SetType(this, ty); };
	void Signal_TypeChanged(qint32 ty) { callbackTransactionDetails742340_TypeChanged(this, ty); };
	qint32 amount() { return callbackTransactionDetails742340_Amount(this); };
	void setAmount(qint32 amount) { callbackTransactionDetails742340_SetAmount(this, amount); };
	void Signal_AmountChanged(qint32 amount) { callbackTransactionDetails742340_AmountChanged(this, amount); };
	qint32 hoursReceived() { return callbackTransactionDetails742340_HoursReceived(this); };
	void setHoursReceived(qint32 hoursReceived) { callbackTransactionDetails742340_SetHoursReceived(this, hoursReceived); };
	void Signal_HoursReceivedChanged(qint32 hoursReceived) { callbackTransactionDetails742340_HoursReceivedChanged(this, hoursReceived); };
	qint32 hoursBurned() { return callbackTransactionDetails742340_HoursBurned(this); };
	void setHoursBurned(qint32 hoursBurned) { callbackTransactionDetails742340_SetHoursBurned(this, hoursBurned); };
	void Signal_HoursBurnedChanged(qint32 hoursBurned) { callbackTransactionDetails742340_HoursBurnedChanged(this, hoursBurned); };
	QString transactionID() { return ({ Moc_PackedString tempVal = callbackTransactionDetails742340_TransactionID(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setTransactionID(QString transactionID) { QByteArray tec2ac1 = transactionID.toUtf8(); Moc_PackedString transactionIDPacked = { const_cast<char*>(tec2ac1.prepend("WHITESPACE").constData()+10), tec2ac1.size()-10 };callbackTransactionDetails742340_SetTransactionID(this, transactionIDPacked); };
	void Signal_TransactionIDChanged(QString transactionID) { QByteArray tec2ac1 = transactionID.toUtf8(); Moc_PackedString transactionIDPacked = { const_cast<char*>(tec2ac1.prepend("WHITESPACE").constData()+10), tec2ac1.size()-10 };callbackTransactionDetails742340_TransactionIDChanged(this, transactionIDPacked); };
	QString sentAddress() { return ({ Moc_PackedString tempVal = callbackTransactionDetails742340_SentAddress(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setSentAddress(QString sentAddress) { QByteArray ted7223 = sentAddress.toUtf8(); Moc_PackedString sentAddressPacked = { const_cast<char*>(ted7223.prepend("WHITESPACE").constData()+10), ted7223.size()-10 };callbackTransactionDetails742340_SetSentAddress(this, sentAddressPacked); };
	void Signal_SentAddressChanged(QString sentAddress) { QByteArray ted7223 = sentAddress.toUtf8(); Moc_PackedString sentAddressPacked = { const_cast<char*>(ted7223.prepend("WHITESPACE").constData()+10), ted7223.size()-10 };callbackTransactionDetails742340_SentAddressChanged(this, sentAddressPacked); };
	QString receivedAddress() { return ({ Moc_PackedString tempVal = callbackTransactionDetails742340_ReceivedAddress(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setReceivedAddress(QString receivedAddress) { QByteArray t55bc8c = receivedAddress.toUtf8(); Moc_PackedString receivedAddressPacked = { const_cast<char*>(t55bc8c.prepend("WHITESPACE").constData()+10), t55bc8c.size()-10 };callbackTransactionDetails742340_SetReceivedAddress(this, receivedAddressPacked); };
	void Signal_ReceivedAddressChanged(QString receivedAddress) { QByteArray t55bc8c = receivedAddress.toUtf8(); Moc_PackedString receivedAddressPacked = { const_cast<char*>(t55bc8c.prepend("WHITESPACE").constData()+10), t55bc8c.size()-10 };callbackTransactionDetails742340_ReceivedAddressChanged(this, receivedAddressPacked); };
	 ~TransactionDetails742340() { callbackTransactionDetails742340_DestroyTransactionDetails(this); };
	void childEvent(QChildEvent * event) { callbackTransactionDetails742340_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackTransactionDetails742340_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackTransactionDetails742340_CustomEvent(this, event); };
	void deleteLater() { callbackTransactionDetails742340_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackTransactionDetails742340_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackTransactionDetails742340_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackTransactionDetails742340_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackTransactionDetails742340_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackTransactionDetails742340_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackTransactionDetails742340_TimerEvent(this, event); };
	QDateTime dateDefault() { return _date; };
	void setDateDefault(QDateTime p) { if (p != _date) { _date = p; dateChanged(_date); } };
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
	void dateChanged(QDateTime date);
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
	QDateTime _date;
	qint32 _status;
	qint32 _type;
	qint32 _amount;
	qint32 _hoursReceived;
	qint32 _hoursBurned;
	QString _transactionID;
	QString _sentAddress;
	QString _receivedAddress;
};

Q_DECLARE_METATYPE(TransactionDetails742340*)


void TransactionDetails742340_TransactionDetails742340_QRegisterMetaTypes() {
	qRegisterMetaType<QDateTime>();
	qRegisterMetaType<QString>();
}

struct Moc_PackedString AddressDeatails742340_Address(void* ptr)
{
	return ({ QByteArray tc350c2 = static_cast<AddressDeatails742340*>(ptr)->address().toUtf8(); Moc_PackedString { const_cast<char*>(tc350c2.prepend("WHITESPACE").constData()+10), tc350c2.size()-10 }; });
}

struct Moc_PackedString AddressDeatails742340_AddressDefault(void* ptr)
{
	return ({ QByteArray t5e2b39 = static_cast<AddressDeatails742340*>(ptr)->addressDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t5e2b39.prepend("WHITESPACE").constData()+10), t5e2b39.size()-10 }; });
}

void AddressDeatails742340_SetAddress(void* ptr, struct Moc_PackedString address)
{
	static_cast<AddressDeatails742340*>(ptr)->setAddress(QString::fromUtf8(address.data, address.len));
}

void AddressDeatails742340_SetAddressDefault(void* ptr, struct Moc_PackedString address)
{
	static_cast<AddressDeatails742340*>(ptr)->setAddressDefault(QString::fromUtf8(address.data, address.len));
}

void AddressDeatails742340_ConnectAddressChanged(void* ptr)
{
	QObject::connect(static_cast<AddressDeatails742340*>(ptr), static_cast<void (AddressDeatails742340::*)(QString)>(&AddressDeatails742340::addressChanged), static_cast<AddressDeatails742340*>(ptr), static_cast<void (AddressDeatails742340::*)(QString)>(&AddressDeatails742340::Signal_AddressChanged));
}

void AddressDeatails742340_DisconnectAddressChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressDeatails742340*>(ptr), static_cast<void (AddressDeatails742340::*)(QString)>(&AddressDeatails742340::addressChanged), static_cast<AddressDeatails742340*>(ptr), static_cast<void (AddressDeatails742340::*)(QString)>(&AddressDeatails742340::Signal_AddressChanged));
}

void AddressDeatails742340_AddressChanged(void* ptr, struct Moc_PackedString address)
{
	static_cast<AddressDeatails742340*>(ptr)->addressChanged(QString::fromUtf8(address.data, address.len));
}

float AddressDeatails742340_AddressSky(void* ptr)
{
	return static_cast<AddressDeatails742340*>(ptr)->addressSky();
}

float AddressDeatails742340_AddressSkyDefault(void* ptr)
{
	return static_cast<AddressDeatails742340*>(ptr)->addressSkyDefault();
}

void AddressDeatails742340_SetAddressSky(void* ptr, float addressSky)
{
	static_cast<AddressDeatails742340*>(ptr)->setAddressSky(addressSky);
}

void AddressDeatails742340_SetAddressSkyDefault(void* ptr, float addressSky)
{
	static_cast<AddressDeatails742340*>(ptr)->setAddressSkyDefault(addressSky);
}

void AddressDeatails742340_ConnectAddressSkyChanged(void* ptr)
{
	QObject::connect(static_cast<AddressDeatails742340*>(ptr), static_cast<void (AddressDeatails742340::*)(float)>(&AddressDeatails742340::addressSkyChanged), static_cast<AddressDeatails742340*>(ptr), static_cast<void (AddressDeatails742340::*)(float)>(&AddressDeatails742340::Signal_AddressSkyChanged));
}

void AddressDeatails742340_DisconnectAddressSkyChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressDeatails742340*>(ptr), static_cast<void (AddressDeatails742340::*)(float)>(&AddressDeatails742340::addressSkyChanged), static_cast<AddressDeatails742340*>(ptr), static_cast<void (AddressDeatails742340::*)(float)>(&AddressDeatails742340::Signal_AddressSkyChanged));
}

void AddressDeatails742340_AddressSkyChanged(void* ptr, float addressSky)
{
	static_cast<AddressDeatails742340*>(ptr)->addressSkyChanged(addressSky);
}

int AddressDeatails742340_AddressCoinHours(void* ptr)
{
	return static_cast<AddressDeatails742340*>(ptr)->addressCoinHours();
}

int AddressDeatails742340_AddressCoinHoursDefault(void* ptr)
{
	return static_cast<AddressDeatails742340*>(ptr)->addressCoinHoursDefault();
}

void AddressDeatails742340_SetAddressCoinHours(void* ptr, int addressCoinHours)
{
	static_cast<AddressDeatails742340*>(ptr)->setAddressCoinHours(addressCoinHours);
}

void AddressDeatails742340_SetAddressCoinHoursDefault(void* ptr, int addressCoinHours)
{
	static_cast<AddressDeatails742340*>(ptr)->setAddressCoinHoursDefault(addressCoinHours);
}

void AddressDeatails742340_ConnectAddressCoinHoursChanged(void* ptr)
{
	QObject::connect(static_cast<AddressDeatails742340*>(ptr), static_cast<void (AddressDeatails742340::*)(qint32)>(&AddressDeatails742340::addressCoinHoursChanged), static_cast<AddressDeatails742340*>(ptr), static_cast<void (AddressDeatails742340::*)(qint32)>(&AddressDeatails742340::Signal_AddressCoinHoursChanged));
}

void AddressDeatails742340_DisconnectAddressCoinHoursChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressDeatails742340*>(ptr), static_cast<void (AddressDeatails742340::*)(qint32)>(&AddressDeatails742340::addressCoinHoursChanged), static_cast<AddressDeatails742340*>(ptr), static_cast<void (AddressDeatails742340::*)(qint32)>(&AddressDeatails742340::Signal_AddressCoinHoursChanged));
}

void AddressDeatails742340_AddressCoinHoursChanged(void* ptr, int addressCoinHours)
{
	static_cast<AddressDeatails742340*>(ptr)->addressCoinHoursChanged(addressCoinHours);
}

int AddressDeatails742340_AddressDeatails742340_QRegisterMetaType()
{
	return qRegisterMetaType<AddressDeatails742340*>();
}

int AddressDeatails742340_AddressDeatails742340_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<AddressDeatails742340*>(const_cast<const char*>(typeName));
}

int AddressDeatails742340_AddressDeatails742340_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressDeatails742340>();
#else
	return 0;
#endif
}

int AddressDeatails742340_AddressDeatails742340_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressDeatails742340>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* AddressDeatails742340___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressDeatails742340___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressDeatails742340___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* AddressDeatails742340___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressDeatails742340___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* AddressDeatails742340___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* AddressDeatails742340___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressDeatails742340___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressDeatails742340___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressDeatails742340___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressDeatails742340___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressDeatails742340___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressDeatails742340___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressDeatails742340___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressDeatails742340___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressDeatails742340_NewAddressDeatails(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new AddressDeatails742340(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new AddressDeatails742340(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new AddressDeatails742340(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new AddressDeatails742340(static_cast<QWindow*>(parent));
	} else {
		return new AddressDeatails742340(static_cast<QObject*>(parent));
	}
}

void AddressDeatails742340_DestroyAddressDeatails(void* ptr)
{
	static_cast<AddressDeatails742340*>(ptr)->~AddressDeatails742340();
}

void AddressDeatails742340_DestroyAddressDeatailsDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void AddressDeatails742340_ChildEventDefault(void* ptr, void* event)
{
	static_cast<AddressDeatails742340*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void AddressDeatails742340_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressDeatails742340*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void AddressDeatails742340_CustomEventDefault(void* ptr, void* event)
{
	static_cast<AddressDeatails742340*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void AddressDeatails742340_DeleteLaterDefault(void* ptr)
{
	static_cast<AddressDeatails742340*>(ptr)->QObject::deleteLater();
}

void AddressDeatails742340_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressDeatails742340*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char AddressDeatails742340_EventDefault(void* ptr, void* e)
{
	return static_cast<AddressDeatails742340*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char AddressDeatails742340_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<AddressDeatails742340*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void AddressDeatails742340_TimerEventDefault(void* ptr, void* event)
{
	static_cast<AddressDeatails742340*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

void AddressList742340_ConnectAddAddress(void* ptr)
{
	QObject::connect(static_cast<AddressList742340*>(ptr), static_cast<void (AddressList742340::*)(AddressDeatails742340*)>(&AddressList742340::addAddress), static_cast<AddressList742340*>(ptr), static_cast<void (AddressList742340::*)(AddressDeatails742340*)>(&AddressList742340::Signal_AddAddress));
}

void AddressList742340_DisconnectAddAddress(void* ptr)
{
	QObject::disconnect(static_cast<AddressList742340*>(ptr), static_cast<void (AddressList742340::*)(AddressDeatails742340*)>(&AddressList742340::addAddress), static_cast<AddressList742340*>(ptr), static_cast<void (AddressList742340::*)(AddressDeatails742340*)>(&AddressList742340::Signal_AddAddress));
}

void AddressList742340_AddAddress(void* ptr, void* transaction)
{
	static_cast<AddressList742340*>(ptr)->addAddress(static_cast<AddressDeatails742340*>(transaction));
}

void AddressList742340_ConnectRemoveAddress(void* ptr)
{
	QObject::connect(static_cast<AddressList742340*>(ptr), static_cast<void (AddressList742340::*)(qint32)>(&AddressList742340::removeAddress), static_cast<AddressList742340*>(ptr), static_cast<void (AddressList742340::*)(qint32)>(&AddressList742340::Signal_RemoveAddress));
}

void AddressList742340_DisconnectRemoveAddress(void* ptr)
{
	QObject::disconnect(static_cast<AddressList742340*>(ptr), static_cast<void (AddressList742340::*)(qint32)>(&AddressList742340::removeAddress), static_cast<AddressList742340*>(ptr), static_cast<void (AddressList742340::*)(qint32)>(&AddressList742340::Signal_RemoveAddress));
}

void AddressList742340_RemoveAddress(void* ptr, int index)
{
	static_cast<AddressList742340*>(ptr)->removeAddress(index);
}

struct Moc_PackedList AddressList742340_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<AddressList742340*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressList742340_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<AddressList742340*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void AddressList742340_SetRoles(void* ptr, void* roles)
{
	static_cast<AddressList742340*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void AddressList742340_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<AddressList742340*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void AddressList742340_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<AddressList742340*>(ptr), static_cast<void (AddressList742340::*)(QMap<qint32, QByteArray>)>(&AddressList742340::rolesChanged), static_cast<AddressList742340*>(ptr), static_cast<void (AddressList742340::*)(QMap<qint32, QByteArray>)>(&AddressList742340::Signal_RolesChanged));
}

void AddressList742340_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressList742340*>(ptr), static_cast<void (AddressList742340::*)(QMap<qint32, QByteArray>)>(&AddressList742340::rolesChanged), static_cast<AddressList742340*>(ptr), static_cast<void (AddressList742340::*)(QMap<qint32, QByteArray>)>(&AddressList742340::Signal_RolesChanged));
}

void AddressList742340_RolesChanged(void* ptr, void* roles)
{
	static_cast<AddressList742340*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList AddressList742340_Addresses(void* ptr)
{
	return ({ QList<AddressDeatails742340*>* tmpValue = new QList<AddressDeatails742340*>(static_cast<AddressList742340*>(ptr)->addresses()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressList742340_AddressesDefault(void* ptr)
{
	return ({ QList<AddressDeatails742340*>* tmpValue = new QList<AddressDeatails742340*>(static_cast<AddressList742340*>(ptr)->addressesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void AddressList742340_SetAddresses(void* ptr, void* addresses)
{
	static_cast<AddressList742340*>(ptr)->setAddresses(({ QList<AddressDeatails742340*>* tmpP = static_cast<QList<AddressDeatails742340*>*>(addresses); QList<AddressDeatails742340*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void AddressList742340_SetAddressesDefault(void* ptr, void* addresses)
{
	static_cast<AddressList742340*>(ptr)->setAddressesDefault(({ QList<AddressDeatails742340*>* tmpP = static_cast<QList<AddressDeatails742340*>*>(addresses); QList<AddressDeatails742340*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void AddressList742340_ConnectAddressesChanged(void* ptr)
{
	QObject::connect(static_cast<AddressList742340*>(ptr), static_cast<void (AddressList742340::*)(QList<AddressDeatails742340*>)>(&AddressList742340::addressesChanged), static_cast<AddressList742340*>(ptr), static_cast<void (AddressList742340::*)(QList<AddressDeatails742340*>)>(&AddressList742340::Signal_AddressesChanged));
}

void AddressList742340_DisconnectAddressesChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressList742340*>(ptr), static_cast<void (AddressList742340::*)(QList<AddressDeatails742340*>)>(&AddressList742340::addressesChanged), static_cast<AddressList742340*>(ptr), static_cast<void (AddressList742340::*)(QList<AddressDeatails742340*>)>(&AddressList742340::Signal_AddressesChanged));
}

void AddressList742340_AddressesChanged(void* ptr, void* addresses)
{
	static_cast<AddressList742340*>(ptr)->addressesChanged(({ QList<AddressDeatails742340*>* tmpP = static_cast<QList<AddressDeatails742340*>*>(addresses); QList<AddressDeatails742340*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int AddressList742340_AddressList742340_QRegisterMetaType()
{
	return qRegisterMetaType<AddressList742340*>();
}

int AddressList742340_AddressList742340_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<AddressList742340*>(const_cast<const char*>(typeName));
}

int AddressList742340_AddressList742340_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressList742340>();
#else
	return 0;
#endif
}

int AddressList742340_AddressList742340_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressList742340>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int AddressList742340_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList742340_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressList742340_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressList742340_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList742340_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressList742340_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressList742340_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList742340_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressList742340_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* AddressList742340___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressList742340___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressList742340___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressList742340___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressList742340___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressList742340___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int AddressList742340___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void AddressList742340___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* AddressList742340___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* AddressList742340___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressList742340___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* AddressList742340___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList AddressList742340___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressList742340___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressList742340___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* AddressList742340___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* AddressList742340___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressList742340___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* AddressList742340___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* AddressList742340___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressList742340___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressList742340___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressList742340___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressList742340___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressList742340___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressList742340___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressList742340___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressList742340___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressList742340___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void AddressList742340___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressList742340___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList AddressList742340___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressList742340___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressList742340___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* AddressList742340___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList AddressList742340___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int AddressList742340_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList742340_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressList742340_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressList742340_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList742340_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressList742340_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* AddressList742340___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList742340___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressList742340___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* AddressList742340___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressList742340___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* AddressList742340___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* AddressList742340___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList742340___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressList742340___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressList742340___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList742340___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressList742340___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressList742340___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList742340___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressList742340___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressList742340___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressList742340___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressList742340___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressList742340___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressList742340___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressList742340___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressList742340___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressList742340___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressList742340___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressList742340___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressList742340___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressList742340___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressList742340___addresses_atList(void* ptr, int i)
{
	return ({AddressDeatails742340* tmp = static_cast<QList<AddressDeatails742340*>*>(ptr)->at(i); if (i == static_cast<QList<AddressDeatails742340*>*>(ptr)->size()-1) { static_cast<QList<AddressDeatails742340*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList742340___addresses_setList(void* ptr, void* i)
{
	static_cast<QList<AddressDeatails742340*>*>(ptr)->append(static_cast<AddressDeatails742340*>(i));
}

void* AddressList742340___addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<AddressDeatails742340*>();
}

void* AddressList742340___setAddresses_addresses_atList(void* ptr, int i)
{
	return ({AddressDeatails742340* tmp = static_cast<QList<AddressDeatails742340*>*>(ptr)->at(i); if (i == static_cast<QList<AddressDeatails742340*>*>(ptr)->size()-1) { static_cast<QList<AddressDeatails742340*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList742340___setAddresses_addresses_setList(void* ptr, void* i)
{
	static_cast<QList<AddressDeatails742340*>*>(ptr)->append(static_cast<AddressDeatails742340*>(i));
}

void* AddressList742340___setAddresses_addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<AddressDeatails742340*>();
}

void* AddressList742340___addressesChanged_addresses_atList(void* ptr, int i)
{
	return ({AddressDeatails742340* tmp = static_cast<QList<AddressDeatails742340*>*>(ptr)->at(i); if (i == static_cast<QList<AddressDeatails742340*>*>(ptr)->size()-1) { static_cast<QList<AddressDeatails742340*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList742340___addressesChanged_addresses_setList(void* ptr, void* i)
{
	static_cast<QList<AddressDeatails742340*>*>(ptr)->append(static_cast<AddressDeatails742340*>(i));
}

void* AddressList742340___addressesChanged_addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<AddressDeatails742340*>();
}

int AddressList742340_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList742340_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressList742340_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int AddressList742340_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList742340_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressList742340_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int AddressList742340_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressList742340_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressList742340_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* AddressList742340_NewAddressList(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new AddressList742340(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new AddressList742340(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new AddressList742340(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new AddressList742340(static_cast<QWindow*>(parent));
	} else {
		return new AddressList742340(static_cast<QObject*>(parent));
	}
}

void AddressList742340_DestroyAddressList(void* ptr)
{
	static_cast<AddressList742340*>(ptr)->~AddressList742340();
}

void AddressList742340_DestroyAddressListDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char AddressList742340_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<AddressList742340*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

long long AddressList742340_FlagsDefault(void* ptr, void* index)
{
	return static_cast<AddressList742340*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

void* AddressList742340_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<AddressList742340*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* AddressList742340_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<AddressList742340*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* AddressList742340_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<AddressList742340*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

char AddressList742340_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<AddressList742340*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char AddressList742340_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<AddressList742340*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int AddressList742340_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<AddressList742340*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* AddressList742340_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void AddressList742340_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<AddressList742340*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char AddressList742340_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<AddressList742340*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* AddressList742340_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<AddressList742340*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

char AddressList742340_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<AddressList742340*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char AddressList742340_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<AddressList742340*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct Moc_PackedList AddressList742340_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<AddressList742340*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressList742340_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<AddressList742340*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressList742340_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<AddressList742340*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct Moc_PackedString AddressList742340_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t29e12e = static_cast<AddressList742340*>(ptr)->QAbstractListModel::mimeTypes().join("¡¦!").toUtf8(); Moc_PackedString { const_cast<char*>(t29e12e.prepend("WHITESPACE").constData()+10), t29e12e.size()-10 }; });
}

char AddressList742340_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<AddressList742340*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char AddressList742340_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<AddressList742340*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* AddressList742340_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<AddressList742340*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

char AddressList742340_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<AddressList742340*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char AddressList742340_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<AddressList742340*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void AddressList742340_ResetInternalDataDefault(void* ptr)
{
	static_cast<AddressList742340*>(ptr)->QAbstractListModel::resetInternalData();
}

void AddressList742340_RevertDefault(void* ptr)
{
	static_cast<AddressList742340*>(ptr)->QAbstractListModel::revert();
}

struct Moc_PackedList AddressList742340_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<AddressList742340*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int AddressList742340_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char AddressList742340_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<AddressList742340*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char AddressList742340_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<AddressList742340*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char AddressList742340_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<AddressList742340*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void AddressList742340_SortDefault(void* ptr, int column, long long order)
{
	static_cast<AddressList742340*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* AddressList742340_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<AddressList742340*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char AddressList742340_SubmitDefault(void* ptr)
{
	return static_cast<AddressList742340*>(ptr)->QAbstractListModel::submit();
}

long long AddressList742340_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<AddressList742340*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long AddressList742340_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<AddressList742340*>(ptr)->QAbstractListModel::supportedDropActions();
}

void AddressList742340_ChildEventDefault(void* ptr, void* event)
{
	static_cast<AddressList742340*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void AddressList742340_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressList742340*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void AddressList742340_CustomEventDefault(void* ptr, void* event)
{
	static_cast<AddressList742340*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void AddressList742340_DeleteLaterDefault(void* ptr)
{
	static_cast<AddressList742340*>(ptr)->QAbstractListModel::deleteLater();
}

void AddressList742340_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressList742340*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char AddressList742340_EventDefault(void* ptr, void* e)
{
	return static_cast<AddressList742340*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char AddressList742340_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<AddressList742340*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void AddressList742340_TimerEventDefault(void* ptr, void* event)
{
	static_cast<AddressList742340*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void HistoryModel742340_ConnectAddTransaction(void* ptr)
{
	QObject::connect(static_cast<HistoryModel742340*>(ptr), static_cast<void (HistoryModel742340::*)(TransactionDetails742340*)>(&HistoryModel742340::addTransaction), static_cast<HistoryModel742340*>(ptr), static_cast<void (HistoryModel742340::*)(TransactionDetails742340*)>(&HistoryModel742340::Signal_AddTransaction));
}

void HistoryModel742340_DisconnectAddTransaction(void* ptr)
{
	QObject::disconnect(static_cast<HistoryModel742340*>(ptr), static_cast<void (HistoryModel742340::*)(TransactionDetails742340*)>(&HistoryModel742340::addTransaction), static_cast<HistoryModel742340*>(ptr), static_cast<void (HistoryModel742340::*)(TransactionDetails742340*)>(&HistoryModel742340::Signal_AddTransaction));
}

void HistoryModel742340_AddTransaction(void* ptr, void* transaction)
{
	static_cast<HistoryModel742340*>(ptr)->addTransaction(static_cast<TransactionDetails742340*>(transaction));
}

void HistoryModel742340_ConnectRemoveTransaction(void* ptr)
{
	QObject::connect(static_cast<HistoryModel742340*>(ptr), static_cast<void (HistoryModel742340::*)(qint32)>(&HistoryModel742340::removeTransaction), static_cast<HistoryModel742340*>(ptr), static_cast<void (HistoryModel742340::*)(qint32)>(&HistoryModel742340::Signal_RemoveTransaction));
}

void HistoryModel742340_DisconnectRemoveTransaction(void* ptr)
{
	QObject::disconnect(static_cast<HistoryModel742340*>(ptr), static_cast<void (HistoryModel742340::*)(qint32)>(&HistoryModel742340::removeTransaction), static_cast<HistoryModel742340*>(ptr), static_cast<void (HistoryModel742340::*)(qint32)>(&HistoryModel742340::Signal_RemoveTransaction));
}

void HistoryModel742340_RemoveTransaction(void* ptr, int index)
{
	static_cast<HistoryModel742340*>(ptr)->removeTransaction(index);
}

struct Moc_PackedList HistoryModel742340_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<HistoryModel742340*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList HistoryModel742340_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<HistoryModel742340*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void HistoryModel742340_SetRoles(void* ptr, void* roles)
{
	static_cast<HistoryModel742340*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void HistoryModel742340_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<HistoryModel742340*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void HistoryModel742340_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<HistoryModel742340*>(ptr), static_cast<void (HistoryModel742340::*)(QMap<qint32, QByteArray>)>(&HistoryModel742340::rolesChanged), static_cast<HistoryModel742340*>(ptr), static_cast<void (HistoryModel742340::*)(QMap<qint32, QByteArray>)>(&HistoryModel742340::Signal_RolesChanged));
}

void HistoryModel742340_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<HistoryModel742340*>(ptr), static_cast<void (HistoryModel742340::*)(QMap<qint32, QByteArray>)>(&HistoryModel742340::rolesChanged), static_cast<HistoryModel742340*>(ptr), static_cast<void (HistoryModel742340::*)(QMap<qint32, QByteArray>)>(&HistoryModel742340::Signal_RolesChanged));
}

void HistoryModel742340_RolesChanged(void* ptr, void* roles)
{
	static_cast<HistoryModel742340*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList HistoryModel742340_Transactions(void* ptr)
{
	return ({ QList<TransactionDetails742340*>* tmpValue = new QList<TransactionDetails742340*>(static_cast<HistoryModel742340*>(ptr)->transactions()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList HistoryModel742340_TransactionsDefault(void* ptr)
{
	return ({ QList<TransactionDetails742340*>* tmpValue = new QList<TransactionDetails742340*>(static_cast<HistoryModel742340*>(ptr)->transactionsDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void HistoryModel742340_SetTransactions(void* ptr, void* transactions)
{
	static_cast<HistoryModel742340*>(ptr)->setTransactions(({ QList<TransactionDetails742340*>* tmpP = static_cast<QList<TransactionDetails742340*>*>(transactions); QList<TransactionDetails742340*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void HistoryModel742340_SetTransactionsDefault(void* ptr, void* transactions)
{
	static_cast<HistoryModel742340*>(ptr)->setTransactionsDefault(({ QList<TransactionDetails742340*>* tmpP = static_cast<QList<TransactionDetails742340*>*>(transactions); QList<TransactionDetails742340*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void HistoryModel742340_ConnectTransactionsChanged(void* ptr)
{
	QObject::connect(static_cast<HistoryModel742340*>(ptr), static_cast<void (HistoryModel742340::*)(QList<TransactionDetails742340*>)>(&HistoryModel742340::transactionsChanged), static_cast<HistoryModel742340*>(ptr), static_cast<void (HistoryModel742340::*)(QList<TransactionDetails742340*>)>(&HistoryModel742340::Signal_TransactionsChanged));
}

void HistoryModel742340_DisconnectTransactionsChanged(void* ptr)
{
	QObject::disconnect(static_cast<HistoryModel742340*>(ptr), static_cast<void (HistoryModel742340::*)(QList<TransactionDetails742340*>)>(&HistoryModel742340::transactionsChanged), static_cast<HistoryModel742340*>(ptr), static_cast<void (HistoryModel742340::*)(QList<TransactionDetails742340*>)>(&HistoryModel742340::Signal_TransactionsChanged));
}

void HistoryModel742340_TransactionsChanged(void* ptr, void* transactions)
{
	static_cast<HistoryModel742340*>(ptr)->transactionsChanged(({ QList<TransactionDetails742340*>* tmpP = static_cast<QList<TransactionDetails742340*>*>(transactions); QList<TransactionDetails742340*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int HistoryModel742340_HistoryModel742340_QRegisterMetaType()
{
	return qRegisterMetaType<HistoryModel742340*>();
}

int HistoryModel742340_HistoryModel742340_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<HistoryModel742340*>(const_cast<const char*>(typeName));
}

int HistoryModel742340_HistoryModel742340_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<HistoryModel742340>();
#else
	return 0;
#endif
}

int HistoryModel742340_HistoryModel742340_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<HistoryModel742340>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int HistoryModel742340_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel742340_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* HistoryModel742340_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int HistoryModel742340_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel742340_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* HistoryModel742340_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int HistoryModel742340_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel742340_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* HistoryModel742340_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* HistoryModel742340___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void HistoryModel742340___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* HistoryModel742340___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* HistoryModel742340___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void HistoryModel742340___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* HistoryModel742340___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int HistoryModel742340___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void HistoryModel742340___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* HistoryModel742340___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* HistoryModel742340___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void HistoryModel742340___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* HistoryModel742340___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList HistoryModel742340___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* HistoryModel742340___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void HistoryModel742340___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* HistoryModel742340___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* HistoryModel742340___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void HistoryModel742340___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* HistoryModel742340___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* HistoryModel742340___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void HistoryModel742340___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* HistoryModel742340___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* HistoryModel742340___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void HistoryModel742340___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* HistoryModel742340___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* HistoryModel742340___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void HistoryModel742340___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* HistoryModel742340___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* HistoryModel742340___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void HistoryModel742340___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* HistoryModel742340___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList HistoryModel742340___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* HistoryModel742340___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void HistoryModel742340___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* HistoryModel742340___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList HistoryModel742340___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int HistoryModel742340_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel742340_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* HistoryModel742340_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int HistoryModel742340_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel742340_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* HistoryModel742340_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* HistoryModel742340___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel742340___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* HistoryModel742340___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* HistoryModel742340___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void HistoryModel742340___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* HistoryModel742340___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* HistoryModel742340___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel742340___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* HistoryModel742340___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* HistoryModel742340___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel742340___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* HistoryModel742340___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* HistoryModel742340___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel742340___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* HistoryModel742340___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* HistoryModel742340___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void HistoryModel742340___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* HistoryModel742340___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList HistoryModel742340___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* HistoryModel742340___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void HistoryModel742340___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* HistoryModel742340___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList HistoryModel742340___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* HistoryModel742340___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void HistoryModel742340___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* HistoryModel742340___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList HistoryModel742340___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* HistoryModel742340___transactions_atList(void* ptr, int i)
{
	return ({TransactionDetails742340* tmp = static_cast<QList<TransactionDetails742340*>*>(ptr)->at(i); if (i == static_cast<QList<TransactionDetails742340*>*>(ptr)->size()-1) { static_cast<QList<TransactionDetails742340*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel742340___transactions_setList(void* ptr, void* i)
{
	static_cast<QList<TransactionDetails742340*>*>(ptr)->append(static_cast<TransactionDetails742340*>(i));
}

void* HistoryModel742340___transactions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<TransactionDetails742340*>();
}

void* HistoryModel742340___setTransactions_transactions_atList(void* ptr, int i)
{
	return ({TransactionDetails742340* tmp = static_cast<QList<TransactionDetails742340*>*>(ptr)->at(i); if (i == static_cast<QList<TransactionDetails742340*>*>(ptr)->size()-1) { static_cast<QList<TransactionDetails742340*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel742340___setTransactions_transactions_setList(void* ptr, void* i)
{
	static_cast<QList<TransactionDetails742340*>*>(ptr)->append(static_cast<TransactionDetails742340*>(i));
}

void* HistoryModel742340___setTransactions_transactions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<TransactionDetails742340*>();
}

void* HistoryModel742340___transactionsChanged_transactions_atList(void* ptr, int i)
{
	return ({TransactionDetails742340* tmp = static_cast<QList<TransactionDetails742340*>*>(ptr)->at(i); if (i == static_cast<QList<TransactionDetails742340*>*>(ptr)->size()-1) { static_cast<QList<TransactionDetails742340*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel742340___transactionsChanged_transactions_setList(void* ptr, void* i)
{
	static_cast<QList<TransactionDetails742340*>*>(ptr)->append(static_cast<TransactionDetails742340*>(i));
}

void* HistoryModel742340___transactionsChanged_transactions_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<TransactionDetails742340*>();
}

int HistoryModel742340_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel742340_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* HistoryModel742340_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int HistoryModel742340_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel742340_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* HistoryModel742340_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int HistoryModel742340_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void HistoryModel742340_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* HistoryModel742340_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* HistoryModel742340_NewHistoryModel(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new HistoryModel742340(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new HistoryModel742340(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new HistoryModel742340(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new HistoryModel742340(static_cast<QWindow*>(parent));
	} else {
		return new HistoryModel742340(static_cast<QObject*>(parent));
	}
}

void HistoryModel742340_DestroyHistoryModel(void* ptr)
{
	static_cast<HistoryModel742340*>(ptr)->~HistoryModel742340();
}

void HistoryModel742340_DestroyHistoryModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char HistoryModel742340_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

long long HistoryModel742340_FlagsDefault(void* ptr, void* index)
{
	return static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

void* HistoryModel742340_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* HistoryModel742340_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* HistoryModel742340_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

char HistoryModel742340_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char HistoryModel742340_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int HistoryModel742340_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* HistoryModel742340_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void HistoryModel742340_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char HistoryModel742340_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* HistoryModel742340_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

char HistoryModel742340_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char HistoryModel742340_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct Moc_PackedList HistoryModel742340_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList HistoryModel742340_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* HistoryModel742340_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct Moc_PackedString HistoryModel742340_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t2ba84e = static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::mimeTypes().join("¡¦!").toUtf8(); Moc_PackedString { const_cast<char*>(t2ba84e.prepend("WHITESPACE").constData()+10), t2ba84e.size()-10 }; });
}

char HistoryModel742340_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char HistoryModel742340_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* HistoryModel742340_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

char HistoryModel742340_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char HistoryModel742340_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void HistoryModel742340_ResetInternalDataDefault(void* ptr)
{
	static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::resetInternalData();
}

void HistoryModel742340_RevertDefault(void* ptr)
{
	static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::revert();
}

struct Moc_PackedList HistoryModel742340_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int HistoryModel742340_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char HistoryModel742340_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char HistoryModel742340_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char HistoryModel742340_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void HistoryModel742340_SortDefault(void* ptr, int column, long long order)
{
	static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* HistoryModel742340_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char HistoryModel742340_SubmitDefault(void* ptr)
{
	return static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::submit();
}

long long HistoryModel742340_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long HistoryModel742340_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::supportedDropActions();
}

void HistoryModel742340_ChildEventDefault(void* ptr, void* event)
{
	static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void HistoryModel742340_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void HistoryModel742340_CustomEventDefault(void* ptr, void* event)
{
	static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void HistoryModel742340_DeleteLaterDefault(void* ptr)
{
	static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::deleteLater();
}

void HistoryModel742340_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char HistoryModel742340_EventDefault(void* ptr, void* e)
{
	return static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char HistoryModel742340_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void HistoryModel742340_TimerEventDefault(void* ptr, void* event)
{
	static_cast<HistoryModel742340*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void* TransactionDetails742340_Date(void* ptr)
{
	return new QDateTime(static_cast<TransactionDetails742340*>(ptr)->date());
}

void* TransactionDetails742340_DateDefault(void* ptr)
{
	return new QDateTime(static_cast<TransactionDetails742340*>(ptr)->dateDefault());
}

void TransactionDetails742340_SetDate(void* ptr, void* date)
{
	static_cast<TransactionDetails742340*>(ptr)->setDate(*static_cast<QDateTime*>(date));
}

void TransactionDetails742340_SetDateDefault(void* ptr, void* date)
{
	static_cast<TransactionDetails742340*>(ptr)->setDateDefault(*static_cast<QDateTime*>(date));
}

void TransactionDetails742340_ConnectDateChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(QDateTime)>(&TransactionDetails742340::dateChanged), static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(QDateTime)>(&TransactionDetails742340::Signal_DateChanged));
}

void TransactionDetails742340_DisconnectDateChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(QDateTime)>(&TransactionDetails742340::dateChanged), static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(QDateTime)>(&TransactionDetails742340::Signal_DateChanged));
}

void TransactionDetails742340_DateChanged(void* ptr, void* date)
{
	static_cast<TransactionDetails742340*>(ptr)->dateChanged(*static_cast<QDateTime*>(date));
}

int TransactionDetails742340_Status(void* ptr)
{
	return static_cast<TransactionDetails742340*>(ptr)->status();
}

int TransactionDetails742340_StatusDefault(void* ptr)
{
	return static_cast<TransactionDetails742340*>(ptr)->statusDefault();
}

void TransactionDetails742340_SetStatus(void* ptr, int status)
{
	static_cast<TransactionDetails742340*>(ptr)->setStatus(status);
}

void TransactionDetails742340_SetStatusDefault(void* ptr, int status)
{
	static_cast<TransactionDetails742340*>(ptr)->setStatusDefault(status);
}

void TransactionDetails742340_ConnectStatusChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(qint32)>(&TransactionDetails742340::statusChanged), static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(qint32)>(&TransactionDetails742340::Signal_StatusChanged));
}

void TransactionDetails742340_DisconnectStatusChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(qint32)>(&TransactionDetails742340::statusChanged), static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(qint32)>(&TransactionDetails742340::Signal_StatusChanged));
}

void TransactionDetails742340_StatusChanged(void* ptr, int status)
{
	static_cast<TransactionDetails742340*>(ptr)->statusChanged(status);
}

int TransactionDetails742340_Type(void* ptr)
{
	return static_cast<TransactionDetails742340*>(ptr)->type();
}

int TransactionDetails742340_TypeDefault(void* ptr)
{
	return static_cast<TransactionDetails742340*>(ptr)->typeDefault();
}

void TransactionDetails742340_SetType(void* ptr, int ty)
{
	static_cast<TransactionDetails742340*>(ptr)->setType(ty);
}

void TransactionDetails742340_SetTypeDefault(void* ptr, int ty)
{
	static_cast<TransactionDetails742340*>(ptr)->setTypeDefault(ty);
}

void TransactionDetails742340_ConnectTypeChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(qint32)>(&TransactionDetails742340::typeChanged), static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(qint32)>(&TransactionDetails742340::Signal_TypeChanged));
}

void TransactionDetails742340_DisconnectTypeChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(qint32)>(&TransactionDetails742340::typeChanged), static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(qint32)>(&TransactionDetails742340::Signal_TypeChanged));
}

void TransactionDetails742340_TypeChanged(void* ptr, int ty)
{
	static_cast<TransactionDetails742340*>(ptr)->typeChanged(ty);
}

int TransactionDetails742340_Amount(void* ptr)
{
	return static_cast<TransactionDetails742340*>(ptr)->amount();
}

int TransactionDetails742340_AmountDefault(void* ptr)
{
	return static_cast<TransactionDetails742340*>(ptr)->amountDefault();
}

void TransactionDetails742340_SetAmount(void* ptr, int amount)
{
	static_cast<TransactionDetails742340*>(ptr)->setAmount(amount);
}

void TransactionDetails742340_SetAmountDefault(void* ptr, int amount)
{
	static_cast<TransactionDetails742340*>(ptr)->setAmountDefault(amount);
}

void TransactionDetails742340_ConnectAmountChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(qint32)>(&TransactionDetails742340::amountChanged), static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(qint32)>(&TransactionDetails742340::Signal_AmountChanged));
}

void TransactionDetails742340_DisconnectAmountChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(qint32)>(&TransactionDetails742340::amountChanged), static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(qint32)>(&TransactionDetails742340::Signal_AmountChanged));
}

void TransactionDetails742340_AmountChanged(void* ptr, int amount)
{
	static_cast<TransactionDetails742340*>(ptr)->amountChanged(amount);
}

int TransactionDetails742340_HoursReceived(void* ptr)
{
	return static_cast<TransactionDetails742340*>(ptr)->hoursReceived();
}

int TransactionDetails742340_HoursReceivedDefault(void* ptr)
{
	return static_cast<TransactionDetails742340*>(ptr)->hoursReceivedDefault();
}

void TransactionDetails742340_SetHoursReceived(void* ptr, int hoursReceived)
{
	static_cast<TransactionDetails742340*>(ptr)->setHoursReceived(hoursReceived);
}

void TransactionDetails742340_SetHoursReceivedDefault(void* ptr, int hoursReceived)
{
	static_cast<TransactionDetails742340*>(ptr)->setHoursReceivedDefault(hoursReceived);
}

void TransactionDetails742340_ConnectHoursReceivedChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(qint32)>(&TransactionDetails742340::hoursReceivedChanged), static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(qint32)>(&TransactionDetails742340::Signal_HoursReceivedChanged));
}

void TransactionDetails742340_DisconnectHoursReceivedChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(qint32)>(&TransactionDetails742340::hoursReceivedChanged), static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(qint32)>(&TransactionDetails742340::Signal_HoursReceivedChanged));
}

void TransactionDetails742340_HoursReceivedChanged(void* ptr, int hoursReceived)
{
	static_cast<TransactionDetails742340*>(ptr)->hoursReceivedChanged(hoursReceived);
}

int TransactionDetails742340_HoursBurned(void* ptr)
{
	return static_cast<TransactionDetails742340*>(ptr)->hoursBurned();
}

int TransactionDetails742340_HoursBurnedDefault(void* ptr)
{
	return static_cast<TransactionDetails742340*>(ptr)->hoursBurnedDefault();
}

void TransactionDetails742340_SetHoursBurned(void* ptr, int hoursBurned)
{
	static_cast<TransactionDetails742340*>(ptr)->setHoursBurned(hoursBurned);
}

void TransactionDetails742340_SetHoursBurnedDefault(void* ptr, int hoursBurned)
{
	static_cast<TransactionDetails742340*>(ptr)->setHoursBurnedDefault(hoursBurned);
}

void TransactionDetails742340_ConnectHoursBurnedChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(qint32)>(&TransactionDetails742340::hoursBurnedChanged), static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(qint32)>(&TransactionDetails742340::Signal_HoursBurnedChanged));
}

void TransactionDetails742340_DisconnectHoursBurnedChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(qint32)>(&TransactionDetails742340::hoursBurnedChanged), static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(qint32)>(&TransactionDetails742340::Signal_HoursBurnedChanged));
}

void TransactionDetails742340_HoursBurnedChanged(void* ptr, int hoursBurned)
{
	static_cast<TransactionDetails742340*>(ptr)->hoursBurnedChanged(hoursBurned);
}

struct Moc_PackedString TransactionDetails742340_TransactionID(void* ptr)
{
	return ({ QByteArray tc8c5df = static_cast<TransactionDetails742340*>(ptr)->transactionID().toUtf8(); Moc_PackedString { const_cast<char*>(tc8c5df.prepend("WHITESPACE").constData()+10), tc8c5df.size()-10 }; });
}

struct Moc_PackedString TransactionDetails742340_TransactionIDDefault(void* ptr)
{
	return ({ QByteArray t299bde = static_cast<TransactionDetails742340*>(ptr)->transactionIDDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t299bde.prepend("WHITESPACE").constData()+10), t299bde.size()-10 }; });
}

void TransactionDetails742340_SetTransactionID(void* ptr, struct Moc_PackedString transactionID)
{
	static_cast<TransactionDetails742340*>(ptr)->setTransactionID(QString::fromUtf8(transactionID.data, transactionID.len));
}

void TransactionDetails742340_SetTransactionIDDefault(void* ptr, struct Moc_PackedString transactionID)
{
	static_cast<TransactionDetails742340*>(ptr)->setTransactionIDDefault(QString::fromUtf8(transactionID.data, transactionID.len));
}

void TransactionDetails742340_ConnectTransactionIDChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(QString)>(&TransactionDetails742340::transactionIDChanged), static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(QString)>(&TransactionDetails742340::Signal_TransactionIDChanged));
}

void TransactionDetails742340_DisconnectTransactionIDChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(QString)>(&TransactionDetails742340::transactionIDChanged), static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(QString)>(&TransactionDetails742340::Signal_TransactionIDChanged));
}

void TransactionDetails742340_TransactionIDChanged(void* ptr, struct Moc_PackedString transactionID)
{
	static_cast<TransactionDetails742340*>(ptr)->transactionIDChanged(QString::fromUtf8(transactionID.data, transactionID.len));
}

struct Moc_PackedString TransactionDetails742340_SentAddress(void* ptr)
{
	return ({ QByteArray te78448 = static_cast<TransactionDetails742340*>(ptr)->sentAddress().toUtf8(); Moc_PackedString { const_cast<char*>(te78448.prepend("WHITESPACE").constData()+10), te78448.size()-10 }; });
}

struct Moc_PackedString TransactionDetails742340_SentAddressDefault(void* ptr)
{
	return ({ QByteArray tf856c0 = static_cast<TransactionDetails742340*>(ptr)->sentAddressDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tf856c0.prepend("WHITESPACE").constData()+10), tf856c0.size()-10 }; });
}

void TransactionDetails742340_SetSentAddress(void* ptr, struct Moc_PackedString sentAddress)
{
	static_cast<TransactionDetails742340*>(ptr)->setSentAddress(QString::fromUtf8(sentAddress.data, sentAddress.len));
}

void TransactionDetails742340_SetSentAddressDefault(void* ptr, struct Moc_PackedString sentAddress)
{
	static_cast<TransactionDetails742340*>(ptr)->setSentAddressDefault(QString::fromUtf8(sentAddress.data, sentAddress.len));
}

void TransactionDetails742340_ConnectSentAddressChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(QString)>(&TransactionDetails742340::sentAddressChanged), static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(QString)>(&TransactionDetails742340::Signal_SentAddressChanged));
}

void TransactionDetails742340_DisconnectSentAddressChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(QString)>(&TransactionDetails742340::sentAddressChanged), static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(QString)>(&TransactionDetails742340::Signal_SentAddressChanged));
}

void TransactionDetails742340_SentAddressChanged(void* ptr, struct Moc_PackedString sentAddress)
{
	static_cast<TransactionDetails742340*>(ptr)->sentAddressChanged(QString::fromUtf8(sentAddress.data, sentAddress.len));
}

struct Moc_PackedString TransactionDetails742340_ReceivedAddress(void* ptr)
{
	return ({ QByteArray t3a4c8c = static_cast<TransactionDetails742340*>(ptr)->receivedAddress().toUtf8(); Moc_PackedString { const_cast<char*>(t3a4c8c.prepend("WHITESPACE").constData()+10), t3a4c8c.size()-10 }; });
}

struct Moc_PackedString TransactionDetails742340_ReceivedAddressDefault(void* ptr)
{
	return ({ QByteArray ta1cf30 = static_cast<TransactionDetails742340*>(ptr)->receivedAddressDefault().toUtf8(); Moc_PackedString { const_cast<char*>(ta1cf30.prepend("WHITESPACE").constData()+10), ta1cf30.size()-10 }; });
}

void TransactionDetails742340_SetReceivedAddress(void* ptr, struct Moc_PackedString receivedAddress)
{
	static_cast<TransactionDetails742340*>(ptr)->setReceivedAddress(QString::fromUtf8(receivedAddress.data, receivedAddress.len));
}

void TransactionDetails742340_SetReceivedAddressDefault(void* ptr, struct Moc_PackedString receivedAddress)
{
	static_cast<TransactionDetails742340*>(ptr)->setReceivedAddressDefault(QString::fromUtf8(receivedAddress.data, receivedAddress.len));
}

void TransactionDetails742340_ConnectReceivedAddressChanged(void* ptr)
{
	QObject::connect(static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(QString)>(&TransactionDetails742340::receivedAddressChanged), static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(QString)>(&TransactionDetails742340::Signal_ReceivedAddressChanged));
}

void TransactionDetails742340_DisconnectReceivedAddressChanged(void* ptr)
{
	QObject::disconnect(static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(QString)>(&TransactionDetails742340::receivedAddressChanged), static_cast<TransactionDetails742340*>(ptr), static_cast<void (TransactionDetails742340::*)(QString)>(&TransactionDetails742340::Signal_ReceivedAddressChanged));
}

void TransactionDetails742340_ReceivedAddressChanged(void* ptr, struct Moc_PackedString receivedAddress)
{
	static_cast<TransactionDetails742340*>(ptr)->receivedAddressChanged(QString::fromUtf8(receivedAddress.data, receivedAddress.len));
}

int TransactionDetails742340_TransactionDetails742340_QRegisterMetaType()
{
	return qRegisterMetaType<TransactionDetails742340*>();
}

int TransactionDetails742340_TransactionDetails742340_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<TransactionDetails742340*>(const_cast<const char*>(typeName));
}

int TransactionDetails742340_TransactionDetails742340_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TransactionDetails742340>();
#else
	return 0;
#endif
}

int TransactionDetails742340_TransactionDetails742340_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<TransactionDetails742340>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* TransactionDetails742340___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionDetails742340___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TransactionDetails742340___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* TransactionDetails742340___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void TransactionDetails742340___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* TransactionDetails742340___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* TransactionDetails742340___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionDetails742340___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TransactionDetails742340___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TransactionDetails742340___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionDetails742340___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TransactionDetails742340___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TransactionDetails742340___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void TransactionDetails742340___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* TransactionDetails742340___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* TransactionDetails742340_NewTransactionDetails(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new TransactionDetails742340(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new TransactionDetails742340(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new TransactionDetails742340(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new TransactionDetails742340(static_cast<QWindow*>(parent));
	} else {
		return new TransactionDetails742340(static_cast<QObject*>(parent));
	}
}

void TransactionDetails742340_DestroyTransactionDetails(void* ptr)
{
	static_cast<TransactionDetails742340*>(ptr)->~TransactionDetails742340();
}

void TransactionDetails742340_DestroyTransactionDetailsDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void TransactionDetails742340_ChildEventDefault(void* ptr, void* event)
{
	static_cast<TransactionDetails742340*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void TransactionDetails742340_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TransactionDetails742340*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void TransactionDetails742340_CustomEventDefault(void* ptr, void* event)
{
	static_cast<TransactionDetails742340*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void TransactionDetails742340_DeleteLaterDefault(void* ptr)
{
	static_cast<TransactionDetails742340*>(ptr)->QObject::deleteLater();
}

void TransactionDetails742340_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<TransactionDetails742340*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char TransactionDetails742340_EventDefault(void* ptr, void* e)
{
	return static_cast<TransactionDetails742340*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char TransactionDetails742340_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<TransactionDetails742340*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void TransactionDetails742340_TimerEventDefault(void* ptr, void* event)
{
	static_cast<TransactionDetails742340*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
