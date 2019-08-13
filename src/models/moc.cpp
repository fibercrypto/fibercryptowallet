

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
typedef QMap<qint32, QByteArray> type378cdd;
class AddressesModel539e18: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<QAddress539e18*> addresses READ addresses WRITE setAddresses NOTIFY addressesChanged)
Q_PROPERTY(qint32 count READ count WRITE setCount NOTIFY countChanged)
public:
	AddressesModel539e18(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");AddressesModel539e18_AddressesModel539e18_QRegisterMetaType();AddressesModel539e18_AddressesModel539e18_QRegisterMetaTypes();callbackAddressesModel539e18_Constructor(this);};
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbackAddressesModel539e18_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbackAddressesModel539e18_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbackAddressesModel539e18_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<QAddress539e18*> addresses() { return ({ QList<QAddress539e18*>* tmpP = static_cast<QList<QAddress539e18*>*>(callbackAddressesModel539e18_Addresses(this)); QList<QAddress539e18*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setAddresses(QList<QAddress539e18*> addresses) { callbackAddressesModel539e18_SetAddresses(this, ({ QList<QAddress539e18*>* tmpValue = new QList<QAddress539e18*>(addresses); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_AddressesChanged(QList<QAddress539e18*> addresses) { callbackAddressesModel539e18_AddressesChanged(this, ({ QList<QAddress539e18*>* tmpValue = new QList<QAddress539e18*>(addresses); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	qint32 count() { return callbackAddressesModel539e18_Count(this); };
	void setCount(qint32 count) { callbackAddressesModel539e18_SetCount(this, count); };
	void Signal_CountChanged(qint32 count) { callbackAddressesModel539e18_CountChanged(this, count); };
	 ~AddressesModel539e18() { callbackAddressesModel539e18_DestroyAddressesModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackAddressesModel539e18_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackAddressesModel539e18_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackAddressesModel539e18_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackAddressesModel539e18_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackAddressesModel539e18_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackAddressesModel539e18_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackAddressesModel539e18_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackAddressesModel539e18_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackAddressesModel539e18_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackAddressesModel539e18_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackAddressesModel539e18_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackAddressesModel539e18_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackAddressesModel539e18_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackAddressesModel539e18_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackAddressesModel539e18_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackAddressesModel539e18_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = new QVector<int>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackAddressesModel539e18_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackAddressesModel539e18_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackAddressesModel539e18_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackAddressesModel539e18_HeaderDataChanged(this, orientation, first, last); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackAddressesModel539e18_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackAddressesModel539e18_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackAddressesModel539e18_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackAddressesModel539e18_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackAddressesModel539e18_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackAddressesModel539e18_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackAddressesModel539e18_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackAddressesModel539e18_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackAddressesModel539e18_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackAddressesModel539e18_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackAddressesModel539e18_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackAddressesModel539e18_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackAddressesModel539e18_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackAddressesModel539e18_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackAddressesModel539e18_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackAddressesModel539e18_ResetInternalData(this); };
	void revert() { callbackAddressesModel539e18_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackAddressesModel539e18_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackAddressesModel539e18_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackAddressesModel539e18_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackAddressesModel539e18_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackAddressesModel539e18_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackAddressesModel539e18_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackAddressesModel539e18_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackAddressesModel539e18_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackAddressesModel539e18_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackAddressesModel539e18_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackAddressesModel539e18_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackAddressesModel539e18_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackAddressesModel539e18_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackAddressesModel539e18_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackAddressesModel539e18_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackAddressesModel539e18_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackAddressesModel539e18_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackAddressesModel539e18_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackAddressesModel539e18_CustomEvent(this, event); };
	void deleteLater() { callbackAddressesModel539e18_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackAddressesModel539e18_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackAddressesModel539e18_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackAddressesModel539e18_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackAddressesModel539e18_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackAddressesModel539e18_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackAddressesModel539e18_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<QAddress539e18*> addressesDefault() { return _addresses; };
	void setAddressesDefault(QList<QAddress539e18*> p) { if (p != _addresses) { _addresses = p; addressesChanged(_addresses); } };
	qint32 countDefault() { return _count; };
	void setCountDefault(qint32 p) { if (p != _count) { _count = p; countChanged(_count); } };
signals:
	void rolesChanged(type378cdd roles);
	void addressesChanged(QList<QAddress539e18*> addresses);
	void countChanged(qint32 count);
public slots:
	void addAddress(QAddress539e18* v0) { callbackAddressesModel539e18_AddAddress(this, v0); };
	void removeAddress(qint32 v0) { callbackAddressesModel539e18_RemoveAddress(this, v0); };
	void editAddress(qint32 v0, QString v1, quint64 v2, quint64 v3, qint32 v4) { QByteArray t5a6df7 = v1.toUtf8(); Moc_PackedString v1Packed = { const_cast<char*>(t5a6df7.prepend("WHITESPACE").constData()+10), t5a6df7.size()-10 };callbackAddressesModel539e18_EditAddress(this, v0, v1Packed, v2, v3, v4); };
	void loadModel(QList<QAddress539e18*> v0) { callbackAddressesModel539e18_LoadModel(this, ({ QList<QAddress539e18*>* tmpValue = new QList<QAddress539e18*>(v0); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
private:
	type378cdd _roles;
	QList<QAddress539e18*> _addresses;
	qint32 _count;
};

Q_DECLARE_METATYPE(AddressesModel539e18*)


void AddressesModel539e18_AddressesModel539e18_QRegisterMetaTypes() {
	qRegisterMetaType<QList<QObject*>>("QList<QAddress539e18*>");
	qRegisterMetaType<type378cdd>("type378cdd");
}

class ModelManager539e18: public QObject
{
Q_OBJECT
public:
	ModelManager539e18(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");ModelManager539e18_ModelManager539e18_QRegisterMetaType();ModelManager539e18_ModelManager539e18_QRegisterMetaTypes();callbackModelManager539e18_Constructor(this);};
	 ~ModelManager539e18() { callbackModelManager539e18_DestroyModelManager(this); };
	void childEvent(QChildEvent * event) { callbackModelManager539e18_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackModelManager539e18_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackModelManager539e18_CustomEvent(this, event); };
	void deleteLater() { callbackModelManager539e18_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackModelManager539e18_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackModelManager539e18_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackModelManager539e18_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackModelManager539e18_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackModelManager539e18_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackModelManager539e18_TimerEvent(this, event); };
signals:
public slots:
	void setWalletManager(WalletManager539e18* v0) { callbackModelManager539e18_SetWalletManager(this, v0); };
	AddressesModel539e18* getAddressModel(QString v0) { QByteArray tea1dd7 = v0.toUtf8(); Moc_PackedString v0Packed = { const_cast<char*>(tea1dd7.prepend("WHITESPACE").constData()+10), tea1dd7.size()-10 };return static_cast<AddressesModel539e18*>(callbackModelManager539e18_GetAddressModel(this, v0Packed)); };
private:
};

Q_DECLARE_METATYPE(ModelManager539e18*)


void ModelManager539e18_ModelManager539e18_QRegisterMetaTypes() {
}

class QAddress539e18: public QObject
{
Q_OBJECT
Q_PROPERTY(QString address READ address WRITE setAddress NOTIFY addressChanged)
Q_PROPERTY(quint64 addressSky READ addressSky WRITE setAddressSky NOTIFY addressSkyChanged)
Q_PROPERTY(quint64 addressCoinHours READ addressCoinHours WRITE setAddressCoinHours NOTIFY addressCoinHoursChanged)
Q_PROPERTY(qint32 marked READ marked WRITE setMarked NOTIFY markedChanged)
public:
	QAddress539e18(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QAddress539e18_QAddress539e18_QRegisterMetaType();QAddress539e18_QAddress539e18_QRegisterMetaTypes();callbackQAddress539e18_Constructor(this);};
	QString address() { return ({ Moc_PackedString tempVal = callbackQAddress539e18_Address(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setAddress(QString address) { QByteArray tc66218 = address.toUtf8(); Moc_PackedString addressPacked = { const_cast<char*>(tc66218.prepend("WHITESPACE").constData()+10), tc66218.size()-10 };callbackQAddress539e18_SetAddress(this, addressPacked); };
	void Signal_AddressChanged(QString address) { QByteArray tc66218 = address.toUtf8(); Moc_PackedString addressPacked = { const_cast<char*>(tc66218.prepend("WHITESPACE").constData()+10), tc66218.size()-10 };callbackQAddress539e18_AddressChanged(this, addressPacked); };
	quint64 addressSky() { return callbackQAddress539e18_AddressSky(this); };
	void setAddressSky(quint64 addressSky) { callbackQAddress539e18_SetAddressSky(this, addressSky); };
	void Signal_AddressSkyChanged(quint64 addressSky) { callbackQAddress539e18_AddressSkyChanged(this, addressSky); };
	quint64 addressCoinHours() { return callbackQAddress539e18_AddressCoinHours(this); };
	void setAddressCoinHours(quint64 addressCoinHours) { callbackQAddress539e18_SetAddressCoinHours(this, addressCoinHours); };
	void Signal_AddressCoinHoursChanged(quint64 addressCoinHours) { callbackQAddress539e18_AddressCoinHoursChanged(this, addressCoinHours); };
	qint32 marked() { return callbackQAddress539e18_Marked(this); };
	void setMarked(qint32 marked) { callbackQAddress539e18_SetMarked(this, marked); };
	void Signal_MarkedChanged(qint32 marked) { callbackQAddress539e18_MarkedChanged(this, marked); };
	 ~QAddress539e18() { callbackQAddress539e18_DestroyQAddress(this); };
	void childEvent(QChildEvent * event) { callbackQAddress539e18_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAddress539e18_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAddress539e18_CustomEvent(this, event); };
	void deleteLater() { callbackQAddress539e18_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAddress539e18_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAddress539e18_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQAddress539e18_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAddress539e18_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAddress539e18_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAddress539e18_TimerEvent(this, event); };
	QString addressDefault() { return _address; };
	void setAddressDefault(QString p) { if (p != _address) { _address = p; addressChanged(_address); } };
	quint64 addressSkyDefault() { return _addressSky; };
	void setAddressSkyDefault(quint64 p) { if (p != _addressSky) { _addressSky = p; addressSkyChanged(_addressSky); } };
	quint64 addressCoinHoursDefault() { return _addressCoinHours; };
	void setAddressCoinHoursDefault(quint64 p) { if (p != _addressCoinHours) { _addressCoinHours = p; addressCoinHoursChanged(_addressCoinHours); } };
	qint32 markedDefault() { return _marked; };
	void setMarkedDefault(qint32 p) { if (p != _marked) { _marked = p; markedChanged(_marked); } };
signals:
	void addressChanged(QString address);
	void addressSkyChanged(quint64 addressSky);
	void addressCoinHoursChanged(quint64 addressCoinHours);
	void markedChanged(qint32 marked);
public slots:
private:
	QString _address;
	quint64 _addressSky;
	quint64 _addressCoinHours;
	qint32 _marked;
};

Q_DECLARE_METATYPE(QAddress539e18*)


void QAddress539e18_QAddress539e18_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class QWallet539e18: public QObject
{
Q_OBJECT
Q_PROPERTY(QString name READ name WRITE setName NOTIFY nameChanged)
Q_PROPERTY(qint32 encryptionEnabled READ encryptionEnabled WRITE setEncryptionEnabled NOTIFY encryptionEnabledChanged)
Q_PROPERTY(quint64 sky READ sky WRITE setSky NOTIFY skyChanged)
Q_PROPERTY(quint64 coinHours READ coinHours WRITE setCoinHours NOTIFY coinHoursChanged)
Q_PROPERTY(QString fileName READ fileName WRITE setFileName NOTIFY fileNameChanged)
public:
	QWallet539e18(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QWallet539e18_QWallet539e18_QRegisterMetaType();QWallet539e18_QWallet539e18_QRegisterMetaTypes();callbackQWallet539e18_Constructor(this);};
	QString name() { return ({ Moc_PackedString tempVal = callbackQWallet539e18_Name(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setName(QString name) { QByteArray t6ae999 = name.toUtf8(); Moc_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQWallet539e18_SetName(this, namePacked); };
	void Signal_NameChanged(QString name) { QByteArray t6ae999 = name.toUtf8(); Moc_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQWallet539e18_NameChanged(this, namePacked); };
	qint32 encryptionEnabled() { return callbackQWallet539e18_EncryptionEnabled(this); };
	void setEncryptionEnabled(qint32 encryptionEnabled) { callbackQWallet539e18_SetEncryptionEnabled(this, encryptionEnabled); };
	void Signal_EncryptionEnabledChanged(qint32 encryptionEnabled) { callbackQWallet539e18_EncryptionEnabledChanged(this, encryptionEnabled); };
	quint64 sky() { return callbackQWallet539e18_Sky(this); };
	void setSky(quint64 sky) { callbackQWallet539e18_SetSky(this, sky); };
	void Signal_SkyChanged(quint64 sky) { callbackQWallet539e18_SkyChanged(this, sky); };
	quint64 coinHours() { return callbackQWallet539e18_CoinHours(this); };
	void setCoinHours(quint64 coinHours) { callbackQWallet539e18_SetCoinHours(this, coinHours); };
	void Signal_CoinHoursChanged(quint64 coinHours) { callbackQWallet539e18_CoinHoursChanged(this, coinHours); };
	QString fileName() { return ({ Moc_PackedString tempVal = callbackQWallet539e18_FileName(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setFileName(QString fileName) { QByteArray td83e09 = fileName.toUtf8(); Moc_PackedString fileNamePacked = { const_cast<char*>(td83e09.prepend("WHITESPACE").constData()+10), td83e09.size()-10 };callbackQWallet539e18_SetFileName(this, fileNamePacked); };
	void Signal_FileNameChanged(QString fileName) { QByteArray td83e09 = fileName.toUtf8(); Moc_PackedString fileNamePacked = { const_cast<char*>(td83e09.prepend("WHITESPACE").constData()+10), td83e09.size()-10 };callbackQWallet539e18_FileNameChanged(this, fileNamePacked); };
	 ~QWallet539e18() { callbackQWallet539e18_DestroyQWallet(this); };
	void childEvent(QChildEvent * event) { callbackQWallet539e18_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWallet539e18_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWallet539e18_CustomEvent(this, event); };
	void deleteLater() { callbackQWallet539e18_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWallet539e18_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWallet539e18_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQWallet539e18_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWallet539e18_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWallet539e18_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWallet539e18_TimerEvent(this, event); };
	QString nameDefault() { return _name; };
	void setNameDefault(QString p) { if (p != _name) { _name = p; nameChanged(_name); } };
	qint32 encryptionEnabledDefault() { return _encryptionEnabled; };
	void setEncryptionEnabledDefault(qint32 p) { if (p != _encryptionEnabled) { _encryptionEnabled = p; encryptionEnabledChanged(_encryptionEnabled); } };
	quint64 skyDefault() { return _sky; };
	void setSkyDefault(quint64 p) { if (p != _sky) { _sky = p; skyChanged(_sky); } };
	quint64 coinHoursDefault() { return _coinHours; };
	void setCoinHoursDefault(quint64 p) { if (p != _coinHours) { _coinHours = p; coinHoursChanged(_coinHours); } };
	QString fileNameDefault() { return _fileName; };
	void setFileNameDefault(QString p) { if (p != _fileName) { _fileName = p; fileNameChanged(_fileName); } };
signals:
	void nameChanged(QString name);
	void encryptionEnabledChanged(qint32 encryptionEnabled);
	void skyChanged(quint64 sky);
	void coinHoursChanged(quint64 coinHours);
	void fileNameChanged(QString fileName);
public slots:
private:
	QString _name;
	qint32 _encryptionEnabled;
	quint64 _sky;
	quint64 _coinHours;
	QString _fileName;
};

Q_DECLARE_METATYPE(QWallet539e18*)


void QWallet539e18_QWallet539e18_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class WalletManager539e18: public QObject
{
Q_OBJECT
public:
	WalletManager539e18(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");WalletManager539e18_WalletManager539e18_QRegisterMetaType();WalletManager539e18_WalletManager539e18_QRegisterMetaTypes();callbackWalletManager539e18_Constructor(this);};
	 ~WalletManager539e18() { callbackWalletManager539e18_DestroyWalletManager(this); };
	void childEvent(QChildEvent * event) { callbackWalletManager539e18_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackWalletManager539e18_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackWalletManager539e18_CustomEvent(this, event); };
	void deleteLater() { callbackWalletManager539e18_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackWalletManager539e18_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackWalletManager539e18_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackWalletManager539e18_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackWalletManager539e18_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackWalletManager539e18_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackWalletManager539e18_TimerEvent(this, event); };
signals:
public slots:
	QWallet539e18* createEncryptedWallet(QString seed, QString label, QString password, qint32 scanN) { QByteArray t92713d = seed.toUtf8(); Moc_PackedString seedPacked = { const_cast<char*>(t92713d.prepend("WHITESPACE").constData()+10), t92713d.size()-10 };QByteArray t64c653 = label.toUtf8(); Moc_PackedString labelPacked = { const_cast<char*>(t64c653.prepend("WHITESPACE").constData()+10), t64c653.size()-10 };QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };return static_cast<QWallet539e18*>(callbackWalletManager539e18_CreateEncryptedWallet(this, seedPacked, labelPacked, passwordPacked, scanN)); };
	QWallet539e18* createUnencryptedWallet(QString seed, QString label, qint32 scanN) { QByteArray t92713d = seed.toUtf8(); Moc_PackedString seedPacked = { const_cast<char*>(t92713d.prepend("WHITESPACE").constData()+10), t92713d.size()-10 };QByteArray t64c653 = label.toUtf8(); Moc_PackedString labelPacked = { const_cast<char*>(t64c653.prepend("WHITESPACE").constData()+10), t64c653.size()-10 };return static_cast<QWallet539e18*>(callbackWalletManager539e18_CreateUnencryptedWallet(this, seedPacked, labelPacked, scanN)); };
	QString getNewSeed(qint32 entropy) { return ({ Moc_PackedString tempVal = callbackWalletManager539e18_GetNewSeed(this, entropy); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	qint32 verifySeed(QString seed) { QByteArray t92713d = seed.toUtf8(); Moc_PackedString seedPacked = { const_cast<char*>(t92713d.prepend("WHITESPACE").constData()+10), t92713d.size()-10 };return callbackWalletManager539e18_VerifySeed(this, seedPacked); };
	void newWalletAddress(QString id, qint32 n, QString password) { QByteArray t87ea5d = id.toUtf8(); Moc_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackWalletManager539e18_NewWalletAddress(this, idPacked, n, passwordPacked); };
	void encryptWallet(QString id, QString password) { QByteArray t87ea5d = id.toUtf8(); Moc_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackWalletManager539e18_EncryptWallet(this, idPacked, passwordPacked); };
	void decryptWallet(QString id, QString password) { QByteArray t87ea5d = id.toUtf8(); Moc_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackWalletManager539e18_DecryptWallet(this, idPacked, passwordPacked); };
	QList<QWallet539e18*> getWallets() { return ({ QList<QWallet539e18*>* tmpP = static_cast<QList<QWallet539e18*>*>(callbackWalletManager539e18_GetWallets(this)); QList<QWallet539e18*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QList<QAddress539e18*> getAddresses(QString id) { QByteArray t87ea5d = id.toUtf8(); Moc_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };return ({ QList<QAddress539e18*>* tmpP = static_cast<QList<QAddress539e18*>*>(callbackWalletManager539e18_GetAddresses(this, idPacked)); QList<QAddress539e18*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
private:
};

Q_DECLARE_METATYPE(WalletManager539e18*)


void WalletManager539e18_WalletManager539e18_QRegisterMetaTypes() {
	qRegisterMetaType<QList<QObject*>>("QList<QWallet539e18*>");
	qRegisterMetaType<QList<QObject*>>("QList<QAddress539e18*>");
}

class WalletModel539e18: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<QWallet539e18*> wallets READ wallets WRITE setWallets NOTIFY walletsChanged)
Q_PROPERTY(qint32 count READ count WRITE setCount NOTIFY countChanged)
public:
	WalletModel539e18(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");WalletModel539e18_WalletModel539e18_QRegisterMetaType();WalletModel539e18_WalletModel539e18_QRegisterMetaTypes();callbackWalletModel539e18_Constructor(this);};
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbackWalletModel539e18_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbackWalletModel539e18_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbackWalletModel539e18_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<QWallet539e18*> wallets() { return ({ QList<QWallet539e18*>* tmpP = static_cast<QList<QWallet539e18*>*>(callbackWalletModel539e18_Wallets(this)); QList<QWallet539e18*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setWallets(QList<QWallet539e18*> wallets) { callbackWalletModel539e18_SetWallets(this, ({ QList<QWallet539e18*>* tmpValue = new QList<QWallet539e18*>(wallets); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_WalletsChanged(QList<QWallet539e18*> wallets) { callbackWalletModel539e18_WalletsChanged(this, ({ QList<QWallet539e18*>* tmpValue = new QList<QWallet539e18*>(wallets); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	qint32 count() { return callbackWalletModel539e18_Count(this); };
	void setCount(qint32 count) { callbackWalletModel539e18_SetCount(this, count); };
	void Signal_CountChanged(qint32 count) { callbackWalletModel539e18_CountChanged(this, count); };
	 ~WalletModel539e18() { callbackWalletModel539e18_DestroyWalletModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackWalletModel539e18_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackWalletModel539e18_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackWalletModel539e18_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackWalletModel539e18_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackWalletModel539e18_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackWalletModel539e18_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackWalletModel539e18_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackWalletModel539e18_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackWalletModel539e18_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackWalletModel539e18_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModel539e18_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackWalletModel539e18_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackWalletModel539e18_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModel539e18_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackWalletModel539e18_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackWalletModel539e18_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = new QVector<int>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackWalletModel539e18_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackWalletModel539e18_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackWalletModel539e18_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackWalletModel539e18_HeaderDataChanged(this, orientation, first, last); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackWalletModel539e18_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackWalletModel539e18_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackWalletModel539e18_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackWalletModel539e18_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackWalletModel539e18_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackWalletModel539e18_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackWalletModel539e18_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackWalletModel539e18_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackWalletModel539e18_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackWalletModel539e18_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackWalletModel539e18_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackWalletModel539e18_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackWalletModel539e18_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackWalletModel539e18_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackWalletModel539e18_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackWalletModel539e18_ResetInternalData(this); };
	void revert() { callbackWalletModel539e18_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackWalletModel539e18_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackWalletModel539e18_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackWalletModel539e18_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackWalletModel539e18_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModel539e18_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackWalletModel539e18_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackWalletModel539e18_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModel539e18_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackWalletModel539e18_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackWalletModel539e18_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackWalletModel539e18_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackWalletModel539e18_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackWalletModel539e18_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackWalletModel539e18_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackWalletModel539e18_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackWalletModel539e18_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackWalletModel539e18_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackWalletModel539e18_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackWalletModel539e18_CustomEvent(this, event); };
	void deleteLater() { callbackWalletModel539e18_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackWalletModel539e18_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackWalletModel539e18_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackWalletModel539e18_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackWalletModel539e18_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackWalletModel539e18_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackWalletModel539e18_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<QWallet539e18*> walletsDefault() { return _wallets; };
	void setWalletsDefault(QList<QWallet539e18*> p) { if (p != _wallets) { _wallets = p; walletsChanged(_wallets); } };
	qint32 countDefault() { return _count; };
	void setCountDefault(qint32 p) { if (p != _count) { _count = p; countChanged(_count); } };
signals:
	void rolesChanged(type378cdd roles);
	void walletsChanged(QList<QWallet539e18*> wallets);
	void countChanged(qint32 count);
public slots:
	void addWallet(QWallet539e18* v0) { callbackWalletModel539e18_AddWallet(this, v0); };
	void editWallet(qint32 row, QString name, bool encryptionEnabled, quint64 sky, quint64 coinHours) { QByteArray t6ae999 = name.toUtf8(); Moc_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackWalletModel539e18_EditWallet(this, row, namePacked, encryptionEnabled, sky, coinHours); };
	void removeWallet(qint32 row) { callbackWalletModel539e18_RemoveWallet(this, row); };
	void loadModel(QList<QWallet539e18*> v0) { callbackWalletModel539e18_LoadModel(this, ({ QList<QWallet539e18*>* tmpValue = new QList<QWallet539e18*>(v0); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
private:
	type378cdd _roles;
	QList<QWallet539e18*> _wallets;
	qint32 _count;
};

Q_DECLARE_METATYPE(WalletModel539e18*)


void WalletModel539e18_WalletModel539e18_QRegisterMetaTypes() {
	qRegisterMetaType<QList<QObject*>>("QList<QWallet539e18*>");
	qRegisterMetaType<type378cdd>("type378cdd");
}

void AddressesModel539e18_AddAddress(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<AddressesModel539e18*>(ptr), "addAddress", Q_ARG(QAddress539e18*, static_cast<QAddress539e18*>(v0)));
}

void AddressesModel539e18_RemoveAddress(void* ptr, int v0)
{
	QMetaObject::invokeMethod(static_cast<AddressesModel539e18*>(ptr), "removeAddress", Q_ARG(qint32, v0));
}

void AddressesModel539e18_EditAddress(void* ptr, int v0, struct Moc_PackedString v1, unsigned long long v2, unsigned long long v3, int v4)
{
	QMetaObject::invokeMethod(static_cast<AddressesModel539e18*>(ptr), "editAddress", Q_ARG(qint32, v0), Q_ARG(QString, QString::fromUtf8(v1.data, v1.len)), Q_ARG(quint64, v2), Q_ARG(quint64, v3), Q_ARG(qint32, v4));
}

void AddressesModel539e18_LoadModel(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<AddressesModel539e18*>(ptr), "loadModel", Q_ARG(QList<QAddress539e18*>, ({ QList<QAddress539e18*>* tmpP = static_cast<QList<QAddress539e18*>*>(v0); QList<QAddress539e18*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; })));
}

struct Moc_PackedList AddressesModel539e18_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<AddressesModel539e18*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressesModel539e18_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<AddressesModel539e18*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void AddressesModel539e18_SetRoles(void* ptr, void* roles)
{
	static_cast<AddressesModel539e18*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void AddressesModel539e18_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<AddressesModel539e18*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void AddressesModel539e18_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<AddressesModel539e18*>(ptr), static_cast<void (AddressesModel539e18::*)(QMap<qint32, QByteArray>)>(&AddressesModel539e18::rolesChanged), static_cast<AddressesModel539e18*>(ptr), static_cast<void (AddressesModel539e18::*)(QMap<qint32, QByteArray>)>(&AddressesModel539e18::Signal_RolesChanged));
}

void AddressesModel539e18_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressesModel539e18*>(ptr), static_cast<void (AddressesModel539e18::*)(QMap<qint32, QByteArray>)>(&AddressesModel539e18::rolesChanged), static_cast<AddressesModel539e18*>(ptr), static_cast<void (AddressesModel539e18::*)(QMap<qint32, QByteArray>)>(&AddressesModel539e18::Signal_RolesChanged));
}

void AddressesModel539e18_RolesChanged(void* ptr, void* roles)
{
	static_cast<AddressesModel539e18*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList AddressesModel539e18_Addresses(void* ptr)
{
	return ({ QList<QAddress539e18*>* tmpValue = new QList<QAddress539e18*>(static_cast<AddressesModel539e18*>(ptr)->addresses()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressesModel539e18_AddressesDefault(void* ptr)
{
	return ({ QList<QAddress539e18*>* tmpValue = new QList<QAddress539e18*>(static_cast<AddressesModel539e18*>(ptr)->addressesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void AddressesModel539e18_SetAddresses(void* ptr, void* addresses)
{
	static_cast<AddressesModel539e18*>(ptr)->setAddresses(({ QList<QAddress539e18*>* tmpP = static_cast<QList<QAddress539e18*>*>(addresses); QList<QAddress539e18*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void AddressesModel539e18_SetAddressesDefault(void* ptr, void* addresses)
{
	static_cast<AddressesModel539e18*>(ptr)->setAddressesDefault(({ QList<QAddress539e18*>* tmpP = static_cast<QList<QAddress539e18*>*>(addresses); QList<QAddress539e18*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void AddressesModel539e18_ConnectAddressesChanged(void* ptr)
{
	QObject::connect(static_cast<AddressesModel539e18*>(ptr), static_cast<void (AddressesModel539e18::*)(QList<QAddress539e18*>)>(&AddressesModel539e18::addressesChanged), static_cast<AddressesModel539e18*>(ptr), static_cast<void (AddressesModel539e18::*)(QList<QAddress539e18*>)>(&AddressesModel539e18::Signal_AddressesChanged));
}

void AddressesModel539e18_DisconnectAddressesChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressesModel539e18*>(ptr), static_cast<void (AddressesModel539e18::*)(QList<QAddress539e18*>)>(&AddressesModel539e18::addressesChanged), static_cast<AddressesModel539e18*>(ptr), static_cast<void (AddressesModel539e18::*)(QList<QAddress539e18*>)>(&AddressesModel539e18::Signal_AddressesChanged));
}

void AddressesModel539e18_AddressesChanged(void* ptr, void* addresses)
{
	static_cast<AddressesModel539e18*>(ptr)->addressesChanged(({ QList<QAddress539e18*>* tmpP = static_cast<QList<QAddress539e18*>*>(addresses); QList<QAddress539e18*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int AddressesModel539e18_Count(void* ptr)
{
	return static_cast<AddressesModel539e18*>(ptr)->count();
}

int AddressesModel539e18_CountDefault(void* ptr)
{
	return static_cast<AddressesModel539e18*>(ptr)->countDefault();
}

void AddressesModel539e18_SetCount(void* ptr, int count)
{
	static_cast<AddressesModel539e18*>(ptr)->setCount(count);
}

void AddressesModel539e18_SetCountDefault(void* ptr, int count)
{
	static_cast<AddressesModel539e18*>(ptr)->setCountDefault(count);
}

void AddressesModel539e18_ConnectCountChanged(void* ptr)
{
	QObject::connect(static_cast<AddressesModel539e18*>(ptr), static_cast<void (AddressesModel539e18::*)(qint32)>(&AddressesModel539e18::countChanged), static_cast<AddressesModel539e18*>(ptr), static_cast<void (AddressesModel539e18::*)(qint32)>(&AddressesModel539e18::Signal_CountChanged));
}

void AddressesModel539e18_DisconnectCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressesModel539e18*>(ptr), static_cast<void (AddressesModel539e18::*)(qint32)>(&AddressesModel539e18::countChanged), static_cast<AddressesModel539e18*>(ptr), static_cast<void (AddressesModel539e18::*)(qint32)>(&AddressesModel539e18::Signal_CountChanged));
}

void AddressesModel539e18_CountChanged(void* ptr, int count)
{
	static_cast<AddressesModel539e18*>(ptr)->countChanged(count);
}

int AddressesModel539e18_AddressesModel539e18_QRegisterMetaType()
{
	return qRegisterMetaType<AddressesModel539e18*>();
}

int AddressesModel539e18_AddressesModel539e18_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<AddressesModel539e18*>(const_cast<const char*>(typeName));
}

int AddressesModel539e18_AddressesModel539e18_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressesModel539e18>();
#else
	return 0;
#endif
}

int AddressesModel539e18_AddressesModel539e18_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressesModel539e18>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int AddressesModel539e18_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel539e18_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModel539e18_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressesModel539e18_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel539e18_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModel539e18_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressesModel539e18_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel539e18_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModel539e18_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* AddressesModel539e18___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel539e18___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModel539e18___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressesModel539e18___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel539e18___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModel539e18___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int AddressesModel539e18___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void AddressesModel539e18___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* AddressesModel539e18___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* AddressesModel539e18___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModel539e18___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* AddressesModel539e18___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList AddressesModel539e18___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel539e18___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel539e18___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* AddressesModel539e18___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* AddressesModel539e18___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel539e18___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* AddressesModel539e18___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* AddressesModel539e18___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel539e18___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModel539e18___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressesModel539e18___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel539e18___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModel539e18___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressesModel539e18___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel539e18___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModel539e18___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressesModel539e18___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void AddressesModel539e18___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressesModel539e18___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList AddressesModel539e18___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel539e18___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModel539e18___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* AddressesModel539e18___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList AddressesModel539e18___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int AddressesModel539e18_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel539e18_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModel539e18_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressesModel539e18_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel539e18_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModel539e18_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* AddressesModel539e18___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel539e18___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressesModel539e18___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* AddressesModel539e18___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel539e18___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* AddressesModel539e18___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* AddressesModel539e18___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel539e18___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressesModel539e18___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressesModel539e18___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel539e18___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressesModel539e18___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressesModel539e18___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel539e18___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressesModel539e18___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressesModel539e18___loadModel_v0_atList(void* ptr, int i)
{
	return ({QAddress539e18* tmp = static_cast<QList<QAddress539e18*>*>(ptr)->at(i); if (i == static_cast<QList<QAddress539e18*>*>(ptr)->size()-1) { static_cast<QList<QAddress539e18*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel539e18___loadModel_v0_setList(void* ptr, void* i)
{
	static_cast<QList<QAddress539e18*>*>(ptr)->append(static_cast<QAddress539e18*>(i));
}

void* AddressesModel539e18___loadModel_v0_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAddress539e18*>();
}

void* AddressesModel539e18___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModel539e18___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressesModel539e18___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressesModel539e18___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel539e18___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModel539e18___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressesModel539e18___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressesModel539e18___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel539e18___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModel539e18___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressesModel539e18___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressesModel539e18___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel539e18___addresses_atList(void* ptr, int i)
{
	return ({QAddress539e18* tmp = static_cast<QList<QAddress539e18*>*>(ptr)->at(i); if (i == static_cast<QList<QAddress539e18*>*>(ptr)->size()-1) { static_cast<QList<QAddress539e18*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel539e18___addresses_setList(void* ptr, void* i)
{
	static_cast<QList<QAddress539e18*>*>(ptr)->append(static_cast<QAddress539e18*>(i));
}

void* AddressesModel539e18___addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAddress539e18*>();
}

void* AddressesModel539e18___setAddresses_addresses_atList(void* ptr, int i)
{
	return ({QAddress539e18* tmp = static_cast<QList<QAddress539e18*>*>(ptr)->at(i); if (i == static_cast<QList<QAddress539e18*>*>(ptr)->size()-1) { static_cast<QList<QAddress539e18*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel539e18___setAddresses_addresses_setList(void* ptr, void* i)
{
	static_cast<QList<QAddress539e18*>*>(ptr)->append(static_cast<QAddress539e18*>(i));
}

void* AddressesModel539e18___setAddresses_addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAddress539e18*>();
}

void* AddressesModel539e18___addressesChanged_addresses_atList(void* ptr, int i)
{
	return ({QAddress539e18* tmp = static_cast<QList<QAddress539e18*>*>(ptr)->at(i); if (i == static_cast<QList<QAddress539e18*>*>(ptr)->size()-1) { static_cast<QList<QAddress539e18*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel539e18___addressesChanged_addresses_setList(void* ptr, void* i)
{
	static_cast<QList<QAddress539e18*>*>(ptr)->append(static_cast<QAddress539e18*>(i));
}

void* AddressesModel539e18___addressesChanged_addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAddress539e18*>();
}

int AddressesModel539e18_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel539e18_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressesModel539e18_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int AddressesModel539e18_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel539e18_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressesModel539e18_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int AddressesModel539e18_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel539e18_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressesModel539e18_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* AddressesModel539e18_NewAddressesModel(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new AddressesModel539e18(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new AddressesModel539e18(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new AddressesModel539e18(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new AddressesModel539e18(static_cast<QWindow*>(parent));
	} else {
		return new AddressesModel539e18(static_cast<QObject*>(parent));
	}
}

void AddressesModel539e18_DestroyAddressesModel(void* ptr)
{
	static_cast<AddressesModel539e18*>(ptr)->~AddressesModel539e18();
}

void AddressesModel539e18_DestroyAddressesModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char AddressesModel539e18_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

long long AddressesModel539e18_FlagsDefault(void* ptr, void* index)
{
	return static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

void* AddressesModel539e18_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* AddressesModel539e18_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* AddressesModel539e18_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

char AddressesModel539e18_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char AddressesModel539e18_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int AddressesModel539e18_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* AddressesModel539e18_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void AddressesModel539e18_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char AddressesModel539e18_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* AddressesModel539e18_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

char AddressesModel539e18_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char AddressesModel539e18_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct Moc_PackedList AddressesModel539e18_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressesModel539e18_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel539e18_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct Moc_PackedString AddressesModel539e18_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t6a4f86 = static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::mimeTypes().join("¡¦!").toUtf8(); Moc_PackedString { const_cast<char*>(t6a4f86.prepend("WHITESPACE").constData()+10), t6a4f86.size()-10 }; });
}

char AddressesModel539e18_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char AddressesModel539e18_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* AddressesModel539e18_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

char AddressesModel539e18_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char AddressesModel539e18_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void AddressesModel539e18_ResetInternalDataDefault(void* ptr)
{
	static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::resetInternalData();
}

void AddressesModel539e18_RevertDefault(void* ptr)
{
	static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::revert();
}

struct Moc_PackedList AddressesModel539e18_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int AddressesModel539e18_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char AddressesModel539e18_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char AddressesModel539e18_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char AddressesModel539e18_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void AddressesModel539e18_SortDefault(void* ptr, int column, long long order)
{
	static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* AddressesModel539e18_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char AddressesModel539e18_SubmitDefault(void* ptr)
{
	return static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::submit();
}

long long AddressesModel539e18_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long AddressesModel539e18_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::supportedDropActions();
}

void AddressesModel539e18_ChildEventDefault(void* ptr, void* event)
{
	static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void AddressesModel539e18_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void AddressesModel539e18_CustomEventDefault(void* ptr, void* event)
{
	static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void AddressesModel539e18_DeleteLaterDefault(void* ptr)
{
	static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::deleteLater();
}

void AddressesModel539e18_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char AddressesModel539e18_EventDefault(void* ptr, void* e)
{
	return static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char AddressesModel539e18_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void AddressesModel539e18_TimerEventDefault(void* ptr, void* event)
{
	static_cast<AddressesModel539e18*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void ModelManager539e18_SetWalletManager(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<ModelManager539e18*>(ptr), "setWalletManager", Q_ARG(WalletManager539e18*, static_cast<WalletManager539e18*>(v0)));
}

void* ModelManager539e18_GetAddressModel(void* ptr, struct Moc_PackedString v0)
{
	AddressesModel539e18* returnArg;
	QMetaObject::invokeMethod(static_cast<ModelManager539e18*>(ptr), "getAddressModel", Q_RETURN_ARG(AddressesModel539e18*, returnArg), Q_ARG(QString, QString::fromUtf8(v0.data, v0.len)));
	return returnArg;
}

int ModelManager539e18_ModelManager539e18_QRegisterMetaType()
{
	return qRegisterMetaType<ModelManager539e18*>();
}

int ModelManager539e18_ModelManager539e18_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<ModelManager539e18*>(const_cast<const char*>(typeName));
}

int ModelManager539e18_ModelManager539e18_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ModelManager539e18>();
#else
	return 0;
#endif
}

int ModelManager539e18_ModelManager539e18_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<ModelManager539e18>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* ModelManager539e18___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ModelManager539e18___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ModelManager539e18___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* ModelManager539e18___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void ModelManager539e18___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* ModelManager539e18___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* ModelManager539e18___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ModelManager539e18___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ModelManager539e18___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ModelManager539e18___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ModelManager539e18___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ModelManager539e18___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ModelManager539e18___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void ModelManager539e18___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* ModelManager539e18___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* ModelManager539e18_NewModelManager(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new ModelManager539e18(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new ModelManager539e18(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new ModelManager539e18(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new ModelManager539e18(static_cast<QWindow*>(parent));
	} else {
		return new ModelManager539e18(static_cast<QObject*>(parent));
	}
}

void ModelManager539e18_DestroyModelManager(void* ptr)
{
	static_cast<ModelManager539e18*>(ptr)->~ModelManager539e18();
}

void ModelManager539e18_DestroyModelManagerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void ModelManager539e18_ChildEventDefault(void* ptr, void* event)
{
	static_cast<ModelManager539e18*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void ModelManager539e18_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ModelManager539e18*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void ModelManager539e18_CustomEventDefault(void* ptr, void* event)
{
	static_cast<ModelManager539e18*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void ModelManager539e18_DeleteLaterDefault(void* ptr)
{
	static_cast<ModelManager539e18*>(ptr)->QObject::deleteLater();
}

void ModelManager539e18_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<ModelManager539e18*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char ModelManager539e18_EventDefault(void* ptr, void* e)
{
	return static_cast<ModelManager539e18*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char ModelManager539e18_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<ModelManager539e18*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void ModelManager539e18_TimerEventDefault(void* ptr, void* event)
{
	static_cast<ModelManager539e18*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

struct Moc_PackedString QAddress539e18_Address(void* ptr)
{
	return ({ QByteArray tbbcec3 = static_cast<QAddress539e18*>(ptr)->address().toUtf8(); Moc_PackedString { const_cast<char*>(tbbcec3.prepend("WHITESPACE").constData()+10), tbbcec3.size()-10 }; });
}

struct Moc_PackedString QAddress539e18_AddressDefault(void* ptr)
{
	return ({ QByteArray tbe1549 = static_cast<QAddress539e18*>(ptr)->addressDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tbe1549.prepend("WHITESPACE").constData()+10), tbe1549.size()-10 }; });
}

void QAddress539e18_SetAddress(void* ptr, struct Moc_PackedString address)
{
	static_cast<QAddress539e18*>(ptr)->setAddress(QString::fromUtf8(address.data, address.len));
}

void QAddress539e18_SetAddressDefault(void* ptr, struct Moc_PackedString address)
{
	static_cast<QAddress539e18*>(ptr)->setAddressDefault(QString::fromUtf8(address.data, address.len));
}

void QAddress539e18_ConnectAddressChanged(void* ptr)
{
	QObject::connect(static_cast<QAddress539e18*>(ptr), static_cast<void (QAddress539e18::*)(QString)>(&QAddress539e18::addressChanged), static_cast<QAddress539e18*>(ptr), static_cast<void (QAddress539e18::*)(QString)>(&QAddress539e18::Signal_AddressChanged));
}

void QAddress539e18_DisconnectAddressChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAddress539e18*>(ptr), static_cast<void (QAddress539e18::*)(QString)>(&QAddress539e18::addressChanged), static_cast<QAddress539e18*>(ptr), static_cast<void (QAddress539e18::*)(QString)>(&QAddress539e18::Signal_AddressChanged));
}

void QAddress539e18_AddressChanged(void* ptr, struct Moc_PackedString address)
{
	static_cast<QAddress539e18*>(ptr)->addressChanged(QString::fromUtf8(address.data, address.len));
}

unsigned long long QAddress539e18_AddressSky(void* ptr)
{
	return static_cast<QAddress539e18*>(ptr)->addressSky();
}

unsigned long long QAddress539e18_AddressSkyDefault(void* ptr)
{
	return static_cast<QAddress539e18*>(ptr)->addressSkyDefault();
}

void QAddress539e18_SetAddressSky(void* ptr, unsigned long long addressSky)
{
	static_cast<QAddress539e18*>(ptr)->setAddressSky(addressSky);
}

void QAddress539e18_SetAddressSkyDefault(void* ptr, unsigned long long addressSky)
{
	static_cast<QAddress539e18*>(ptr)->setAddressSkyDefault(addressSky);
}

void QAddress539e18_ConnectAddressSkyChanged(void* ptr)
{
	QObject::connect(static_cast<QAddress539e18*>(ptr), static_cast<void (QAddress539e18::*)(quint64)>(&QAddress539e18::addressSkyChanged), static_cast<QAddress539e18*>(ptr), static_cast<void (QAddress539e18::*)(quint64)>(&QAddress539e18::Signal_AddressSkyChanged));
}

void QAddress539e18_DisconnectAddressSkyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAddress539e18*>(ptr), static_cast<void (QAddress539e18::*)(quint64)>(&QAddress539e18::addressSkyChanged), static_cast<QAddress539e18*>(ptr), static_cast<void (QAddress539e18::*)(quint64)>(&QAddress539e18::Signal_AddressSkyChanged));
}

void QAddress539e18_AddressSkyChanged(void* ptr, unsigned long long addressSky)
{
	static_cast<QAddress539e18*>(ptr)->addressSkyChanged(addressSky);
}

unsigned long long QAddress539e18_AddressCoinHours(void* ptr)
{
	return static_cast<QAddress539e18*>(ptr)->addressCoinHours();
}

unsigned long long QAddress539e18_AddressCoinHoursDefault(void* ptr)
{
	return static_cast<QAddress539e18*>(ptr)->addressCoinHoursDefault();
}

void QAddress539e18_SetAddressCoinHours(void* ptr, unsigned long long addressCoinHours)
{
	static_cast<QAddress539e18*>(ptr)->setAddressCoinHours(addressCoinHours);
}

void QAddress539e18_SetAddressCoinHoursDefault(void* ptr, unsigned long long addressCoinHours)
{
	static_cast<QAddress539e18*>(ptr)->setAddressCoinHoursDefault(addressCoinHours);
}

void QAddress539e18_ConnectAddressCoinHoursChanged(void* ptr)
{
	QObject::connect(static_cast<QAddress539e18*>(ptr), static_cast<void (QAddress539e18::*)(quint64)>(&QAddress539e18::addressCoinHoursChanged), static_cast<QAddress539e18*>(ptr), static_cast<void (QAddress539e18::*)(quint64)>(&QAddress539e18::Signal_AddressCoinHoursChanged));
}

void QAddress539e18_DisconnectAddressCoinHoursChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAddress539e18*>(ptr), static_cast<void (QAddress539e18::*)(quint64)>(&QAddress539e18::addressCoinHoursChanged), static_cast<QAddress539e18*>(ptr), static_cast<void (QAddress539e18::*)(quint64)>(&QAddress539e18::Signal_AddressCoinHoursChanged));
}

void QAddress539e18_AddressCoinHoursChanged(void* ptr, unsigned long long addressCoinHours)
{
	static_cast<QAddress539e18*>(ptr)->addressCoinHoursChanged(addressCoinHours);
}

int QAddress539e18_Marked(void* ptr)
{
	return static_cast<QAddress539e18*>(ptr)->marked();
}

int QAddress539e18_MarkedDefault(void* ptr)
{
	return static_cast<QAddress539e18*>(ptr)->markedDefault();
}

void QAddress539e18_SetMarked(void* ptr, int marked)
{
	static_cast<QAddress539e18*>(ptr)->setMarked(marked);
}

void QAddress539e18_SetMarkedDefault(void* ptr, int marked)
{
	static_cast<QAddress539e18*>(ptr)->setMarkedDefault(marked);
}

void QAddress539e18_ConnectMarkedChanged(void* ptr)
{
	QObject::connect(static_cast<QAddress539e18*>(ptr), static_cast<void (QAddress539e18::*)(qint32)>(&QAddress539e18::markedChanged), static_cast<QAddress539e18*>(ptr), static_cast<void (QAddress539e18::*)(qint32)>(&QAddress539e18::Signal_MarkedChanged));
}

void QAddress539e18_DisconnectMarkedChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAddress539e18*>(ptr), static_cast<void (QAddress539e18::*)(qint32)>(&QAddress539e18::markedChanged), static_cast<QAddress539e18*>(ptr), static_cast<void (QAddress539e18::*)(qint32)>(&QAddress539e18::Signal_MarkedChanged));
}

void QAddress539e18_MarkedChanged(void* ptr, int marked)
{
	static_cast<QAddress539e18*>(ptr)->markedChanged(marked);
}

int QAddress539e18_QAddress539e18_QRegisterMetaType()
{
	return qRegisterMetaType<QAddress539e18*>();
}

int QAddress539e18_QAddress539e18_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QAddress539e18*>(const_cast<const char*>(typeName));
}

int QAddress539e18_QAddress539e18_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QAddress539e18>();
#else
	return 0;
#endif
}

int QAddress539e18_QAddress539e18_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QAddress539e18>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QAddress539e18___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAddress539e18___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAddress539e18___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QAddress539e18___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAddress539e18___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QAddress539e18___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QAddress539e18___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAddress539e18___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAddress539e18___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAddress539e18___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAddress539e18___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAddress539e18___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAddress539e18___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAddress539e18___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAddress539e18___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAddress539e18_NewQAddress(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QAddress539e18(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QAddress539e18(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QAddress539e18(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QAddress539e18(static_cast<QWindow*>(parent));
	} else {
		return new QAddress539e18(static_cast<QObject*>(parent));
	}
}

void QAddress539e18_DestroyQAddress(void* ptr)
{
	static_cast<QAddress539e18*>(ptr)->~QAddress539e18();
}

void QAddress539e18_DestroyQAddressDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void QAddress539e18_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAddress539e18*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QAddress539e18_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAddress539e18*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAddress539e18_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAddress539e18*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QAddress539e18_DeleteLaterDefault(void* ptr)
{
	static_cast<QAddress539e18*>(ptr)->QObject::deleteLater();
}

void QAddress539e18_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAddress539e18*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QAddress539e18_EventDefault(void* ptr, void* e)
{
	return static_cast<QAddress539e18*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QAddress539e18_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAddress539e18*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QAddress539e18_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAddress539e18*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

struct Moc_PackedString QWallet539e18_Name(void* ptr)
{
	return ({ QByteArray tedc734 = static_cast<QWallet539e18*>(ptr)->name().toUtf8(); Moc_PackedString { const_cast<char*>(tedc734.prepend("WHITESPACE").constData()+10), tedc734.size()-10 }; });
}

struct Moc_PackedString QWallet539e18_NameDefault(void* ptr)
{
	return ({ QByteArray tee7a91 = static_cast<QWallet539e18*>(ptr)->nameDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tee7a91.prepend("WHITESPACE").constData()+10), tee7a91.size()-10 }; });
}

void QWallet539e18_SetName(void* ptr, struct Moc_PackedString name)
{
	static_cast<QWallet539e18*>(ptr)->setName(QString::fromUtf8(name.data, name.len));
}

void QWallet539e18_SetNameDefault(void* ptr, struct Moc_PackedString name)
{
	static_cast<QWallet539e18*>(ptr)->setNameDefault(QString::fromUtf8(name.data, name.len));
}

void QWallet539e18_ConnectNameChanged(void* ptr)
{
	QObject::connect(static_cast<QWallet539e18*>(ptr), static_cast<void (QWallet539e18::*)(QString)>(&QWallet539e18::nameChanged), static_cast<QWallet539e18*>(ptr), static_cast<void (QWallet539e18::*)(QString)>(&QWallet539e18::Signal_NameChanged));
}

void QWallet539e18_DisconnectNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWallet539e18*>(ptr), static_cast<void (QWallet539e18::*)(QString)>(&QWallet539e18::nameChanged), static_cast<QWallet539e18*>(ptr), static_cast<void (QWallet539e18::*)(QString)>(&QWallet539e18::Signal_NameChanged));
}

void QWallet539e18_NameChanged(void* ptr, struct Moc_PackedString name)
{
	static_cast<QWallet539e18*>(ptr)->nameChanged(QString::fromUtf8(name.data, name.len));
}

int QWallet539e18_EncryptionEnabled(void* ptr)
{
	return static_cast<QWallet539e18*>(ptr)->encryptionEnabled();
}

int QWallet539e18_EncryptionEnabledDefault(void* ptr)
{
	return static_cast<QWallet539e18*>(ptr)->encryptionEnabledDefault();
}

void QWallet539e18_SetEncryptionEnabled(void* ptr, int encryptionEnabled)
{
	static_cast<QWallet539e18*>(ptr)->setEncryptionEnabled(encryptionEnabled);
}

void QWallet539e18_SetEncryptionEnabledDefault(void* ptr, int encryptionEnabled)
{
	static_cast<QWallet539e18*>(ptr)->setEncryptionEnabledDefault(encryptionEnabled);
}

void QWallet539e18_ConnectEncryptionEnabledChanged(void* ptr)
{
	QObject::connect(static_cast<QWallet539e18*>(ptr), static_cast<void (QWallet539e18::*)(qint32)>(&QWallet539e18::encryptionEnabledChanged), static_cast<QWallet539e18*>(ptr), static_cast<void (QWallet539e18::*)(qint32)>(&QWallet539e18::Signal_EncryptionEnabledChanged));
}

void QWallet539e18_DisconnectEncryptionEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWallet539e18*>(ptr), static_cast<void (QWallet539e18::*)(qint32)>(&QWallet539e18::encryptionEnabledChanged), static_cast<QWallet539e18*>(ptr), static_cast<void (QWallet539e18::*)(qint32)>(&QWallet539e18::Signal_EncryptionEnabledChanged));
}

void QWallet539e18_EncryptionEnabledChanged(void* ptr, int encryptionEnabled)
{
	static_cast<QWallet539e18*>(ptr)->encryptionEnabledChanged(encryptionEnabled);
}

unsigned long long QWallet539e18_Sky(void* ptr)
{
	return static_cast<QWallet539e18*>(ptr)->sky();
}

unsigned long long QWallet539e18_SkyDefault(void* ptr)
{
	return static_cast<QWallet539e18*>(ptr)->skyDefault();
}

void QWallet539e18_SetSky(void* ptr, unsigned long long sky)
{
	static_cast<QWallet539e18*>(ptr)->setSky(sky);
}

void QWallet539e18_SetSkyDefault(void* ptr, unsigned long long sky)
{
	static_cast<QWallet539e18*>(ptr)->setSkyDefault(sky);
}

void QWallet539e18_ConnectSkyChanged(void* ptr)
{
	QObject::connect(static_cast<QWallet539e18*>(ptr), static_cast<void (QWallet539e18::*)(quint64)>(&QWallet539e18::skyChanged), static_cast<QWallet539e18*>(ptr), static_cast<void (QWallet539e18::*)(quint64)>(&QWallet539e18::Signal_SkyChanged));
}

void QWallet539e18_DisconnectSkyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWallet539e18*>(ptr), static_cast<void (QWallet539e18::*)(quint64)>(&QWallet539e18::skyChanged), static_cast<QWallet539e18*>(ptr), static_cast<void (QWallet539e18::*)(quint64)>(&QWallet539e18::Signal_SkyChanged));
}

void QWallet539e18_SkyChanged(void* ptr, unsigned long long sky)
{
	static_cast<QWallet539e18*>(ptr)->skyChanged(sky);
}

unsigned long long QWallet539e18_CoinHours(void* ptr)
{
	return static_cast<QWallet539e18*>(ptr)->coinHours();
}

unsigned long long QWallet539e18_CoinHoursDefault(void* ptr)
{
	return static_cast<QWallet539e18*>(ptr)->coinHoursDefault();
}

void QWallet539e18_SetCoinHours(void* ptr, unsigned long long coinHours)
{
	static_cast<QWallet539e18*>(ptr)->setCoinHours(coinHours);
}

void QWallet539e18_SetCoinHoursDefault(void* ptr, unsigned long long coinHours)
{
	static_cast<QWallet539e18*>(ptr)->setCoinHoursDefault(coinHours);
}

void QWallet539e18_ConnectCoinHoursChanged(void* ptr)
{
	QObject::connect(static_cast<QWallet539e18*>(ptr), static_cast<void (QWallet539e18::*)(quint64)>(&QWallet539e18::coinHoursChanged), static_cast<QWallet539e18*>(ptr), static_cast<void (QWallet539e18::*)(quint64)>(&QWallet539e18::Signal_CoinHoursChanged));
}

void QWallet539e18_DisconnectCoinHoursChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWallet539e18*>(ptr), static_cast<void (QWallet539e18::*)(quint64)>(&QWallet539e18::coinHoursChanged), static_cast<QWallet539e18*>(ptr), static_cast<void (QWallet539e18::*)(quint64)>(&QWallet539e18::Signal_CoinHoursChanged));
}

void QWallet539e18_CoinHoursChanged(void* ptr, unsigned long long coinHours)
{
	static_cast<QWallet539e18*>(ptr)->coinHoursChanged(coinHours);
}

struct Moc_PackedString QWallet539e18_FileName(void* ptr)
{
	return ({ QByteArray teb23e1 = static_cast<QWallet539e18*>(ptr)->fileName().toUtf8(); Moc_PackedString { const_cast<char*>(teb23e1.prepend("WHITESPACE").constData()+10), teb23e1.size()-10 }; });
}

struct Moc_PackedString QWallet539e18_FileNameDefault(void* ptr)
{
	return ({ QByteArray t811a22 = static_cast<QWallet539e18*>(ptr)->fileNameDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t811a22.prepend("WHITESPACE").constData()+10), t811a22.size()-10 }; });
}

void QWallet539e18_SetFileName(void* ptr, struct Moc_PackedString fileName)
{
	static_cast<QWallet539e18*>(ptr)->setFileName(QString::fromUtf8(fileName.data, fileName.len));
}

void QWallet539e18_SetFileNameDefault(void* ptr, struct Moc_PackedString fileName)
{
	static_cast<QWallet539e18*>(ptr)->setFileNameDefault(QString::fromUtf8(fileName.data, fileName.len));
}

void QWallet539e18_ConnectFileNameChanged(void* ptr)
{
	QObject::connect(static_cast<QWallet539e18*>(ptr), static_cast<void (QWallet539e18::*)(QString)>(&QWallet539e18::fileNameChanged), static_cast<QWallet539e18*>(ptr), static_cast<void (QWallet539e18::*)(QString)>(&QWallet539e18::Signal_FileNameChanged));
}

void QWallet539e18_DisconnectFileNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWallet539e18*>(ptr), static_cast<void (QWallet539e18::*)(QString)>(&QWallet539e18::fileNameChanged), static_cast<QWallet539e18*>(ptr), static_cast<void (QWallet539e18::*)(QString)>(&QWallet539e18::Signal_FileNameChanged));
}

void QWallet539e18_FileNameChanged(void* ptr, struct Moc_PackedString fileName)
{
	static_cast<QWallet539e18*>(ptr)->fileNameChanged(QString::fromUtf8(fileName.data, fileName.len));
}

int QWallet539e18_QWallet539e18_QRegisterMetaType()
{
	return qRegisterMetaType<QWallet539e18*>();
}

int QWallet539e18_QWallet539e18_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QWallet539e18*>(const_cast<const char*>(typeName));
}

int QWallet539e18_QWallet539e18_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QWallet539e18>();
#else
	return 0;
#endif
}

int QWallet539e18_QWallet539e18_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QWallet539e18>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QWallet539e18___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWallet539e18___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWallet539e18___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QWallet539e18___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWallet539e18___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWallet539e18___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QWallet539e18___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWallet539e18___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWallet539e18___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWallet539e18___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWallet539e18___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWallet539e18___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWallet539e18___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWallet539e18___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWallet539e18___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWallet539e18_NewQWallet(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QWallet539e18(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QWallet539e18(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QWallet539e18(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QWallet539e18(static_cast<QWindow*>(parent));
	} else {
		return new QWallet539e18(static_cast<QObject*>(parent));
	}
}

void QWallet539e18_DestroyQWallet(void* ptr)
{
	static_cast<QWallet539e18*>(ptr)->~QWallet539e18();
}

void QWallet539e18_DestroyQWalletDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void QWallet539e18_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QWallet539e18*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QWallet539e18_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWallet539e18*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWallet539e18_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QWallet539e18*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QWallet539e18_DeleteLaterDefault(void* ptr)
{
	static_cast<QWallet539e18*>(ptr)->QObject::deleteLater();
}

void QWallet539e18_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWallet539e18*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QWallet539e18_EventDefault(void* ptr, void* e)
{
	return static_cast<QWallet539e18*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QWallet539e18_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QWallet539e18*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWallet539e18_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QWallet539e18*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

void* WalletManager539e18_CreateEncryptedWallet(void* ptr, struct Moc_PackedString seed, struct Moc_PackedString label, struct Moc_PackedString password, int scanN)
{
	QWallet539e18* returnArg;
	QMetaObject::invokeMethod(static_cast<WalletManager539e18*>(ptr), "createEncryptedWallet", Q_RETURN_ARG(QWallet539e18*, returnArg), Q_ARG(QString, QString::fromUtf8(seed.data, seed.len)), Q_ARG(QString, QString::fromUtf8(label.data, label.len)), Q_ARG(QString, QString::fromUtf8(password.data, password.len)), Q_ARG(qint32, scanN));
	return returnArg;
}

void* WalletManager539e18_CreateUnencryptedWallet(void* ptr, struct Moc_PackedString seed, struct Moc_PackedString label, int scanN)
{
	QWallet539e18* returnArg;
	QMetaObject::invokeMethod(static_cast<WalletManager539e18*>(ptr), "createUnencryptedWallet", Q_RETURN_ARG(QWallet539e18*, returnArg), Q_ARG(QString, QString::fromUtf8(seed.data, seed.len)), Q_ARG(QString, QString::fromUtf8(label.data, label.len)), Q_ARG(qint32, scanN));
	return returnArg;
}

struct Moc_PackedString WalletManager539e18_GetNewSeed(void* ptr, int entropy)
{
	QString returnArg;
	QMetaObject::invokeMethod(static_cast<WalletManager539e18*>(ptr), "getNewSeed", Q_RETURN_ARG(QString, returnArg), Q_ARG(qint32, entropy));
	return ({ QByteArray t8e5b69 = returnArg.toUtf8(); Moc_PackedString { const_cast<char*>(t8e5b69.prepend("WHITESPACE").constData()+10), t8e5b69.size()-10 }; });
}

int WalletManager539e18_VerifySeed(void* ptr, struct Moc_PackedString seed)
{
	qint32 returnArg;
	QMetaObject::invokeMethod(static_cast<WalletManager539e18*>(ptr), "verifySeed", Q_RETURN_ARG(qint32, returnArg), Q_ARG(QString, QString::fromUtf8(seed.data, seed.len)));
	return returnArg;
}

void WalletManager539e18_NewWalletAddress(void* ptr, struct Moc_PackedString id, int n, struct Moc_PackedString password)
{
	QMetaObject::invokeMethod(static_cast<WalletManager539e18*>(ptr), "newWalletAddress", Q_ARG(QString, QString::fromUtf8(id.data, id.len)), Q_ARG(qint32, n), Q_ARG(QString, QString::fromUtf8(password.data, password.len)));
}

void WalletManager539e18_EncryptWallet(void* ptr, struct Moc_PackedString id, struct Moc_PackedString password)
{
	QMetaObject::invokeMethod(static_cast<WalletManager539e18*>(ptr), "encryptWallet", Q_ARG(QString, QString::fromUtf8(id.data, id.len)), Q_ARG(QString, QString::fromUtf8(password.data, password.len)));
}

void WalletManager539e18_DecryptWallet(void* ptr, struct Moc_PackedString id, struct Moc_PackedString password)
{
	QMetaObject::invokeMethod(static_cast<WalletManager539e18*>(ptr), "decryptWallet", Q_ARG(QString, QString::fromUtf8(id.data, id.len)), Q_ARG(QString, QString::fromUtf8(password.data, password.len)));
}

struct Moc_PackedList WalletManager539e18_GetWallets(void* ptr)
{
	QList<QWallet539e18*> returnArg;
	QMetaObject::invokeMethod(static_cast<WalletManager539e18*>(ptr), "getWallets", Q_RETURN_ARG(QList<QWallet539e18*>, returnArg));
	return ({ QList<QWallet539e18*>* tmpValue = new QList<QWallet539e18*>(returnArg); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList WalletManager539e18_GetAddresses(void* ptr, struct Moc_PackedString id)
{
	QList<QAddress539e18*> returnArg;
	QMetaObject::invokeMethod(static_cast<WalletManager539e18*>(ptr), "getAddresses", Q_RETURN_ARG(QList<QAddress539e18*>, returnArg), Q_ARG(QString, QString::fromUtf8(id.data, id.len)));
	return ({ QList<QAddress539e18*>* tmpValue = new QList<QAddress539e18*>(returnArg); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int WalletManager539e18_WalletManager539e18_QRegisterMetaType()
{
	return qRegisterMetaType<WalletManager539e18*>();
}

int WalletManager539e18_WalletManager539e18_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<WalletManager539e18*>(const_cast<const char*>(typeName));
}

int WalletManager539e18_WalletManager539e18_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletManager539e18>();
#else
	return 0;
#endif
}

int WalletManager539e18_WalletManager539e18_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletManager539e18>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* WalletManager539e18___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager539e18___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletManager539e18___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* WalletManager539e18___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletManager539e18___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* WalletManager539e18___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* WalletManager539e18___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager539e18___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletManager539e18___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletManager539e18___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager539e18___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletManager539e18___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletManager539e18___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager539e18___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletManager539e18___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletManager539e18___getWallets_atList(void* ptr, int i)
{
	return ({QWallet539e18* tmp = static_cast<QList<QWallet539e18*>*>(ptr)->at(i); if (i == static_cast<QList<QWallet539e18*>*>(ptr)->size()-1) { static_cast<QList<QWallet539e18*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager539e18___getWallets_setList(void* ptr, void* i)
{
	static_cast<QList<QWallet539e18*>*>(ptr)->append(static_cast<QWallet539e18*>(i));
}

void* WalletManager539e18___getWallets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWallet539e18*>();
}

void* WalletManager539e18___getAddresses_atList(void* ptr, int i)
{
	return ({QAddress539e18* tmp = static_cast<QList<QAddress539e18*>*>(ptr)->at(i); if (i == static_cast<QList<QAddress539e18*>*>(ptr)->size()-1) { static_cast<QList<QAddress539e18*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager539e18___getAddresses_setList(void* ptr, void* i)
{
	static_cast<QList<QAddress539e18*>*>(ptr)->append(static_cast<QAddress539e18*>(i));
}

void* WalletManager539e18___getAddresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAddress539e18*>();
}

void* WalletManager539e18_NewWalletManager(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new WalletManager539e18(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new WalletManager539e18(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new WalletManager539e18(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new WalletManager539e18(static_cast<QWindow*>(parent));
	} else {
		return new WalletManager539e18(static_cast<QObject*>(parent));
	}
}

void WalletManager539e18_DestroyWalletManager(void* ptr)
{
	static_cast<WalletManager539e18*>(ptr)->~WalletManager539e18();
}

void WalletManager539e18_DestroyWalletManagerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void WalletManager539e18_ChildEventDefault(void* ptr, void* event)
{
	static_cast<WalletManager539e18*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void WalletManager539e18_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletManager539e18*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void WalletManager539e18_CustomEventDefault(void* ptr, void* event)
{
	static_cast<WalletManager539e18*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void WalletManager539e18_DeleteLaterDefault(void* ptr)
{
	static_cast<WalletManager539e18*>(ptr)->QObject::deleteLater();
}

void WalletManager539e18_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletManager539e18*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char WalletManager539e18_EventDefault(void* ptr, void* e)
{
	return static_cast<WalletManager539e18*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char WalletManager539e18_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<WalletManager539e18*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void WalletManager539e18_TimerEventDefault(void* ptr, void* event)
{
	static_cast<WalletManager539e18*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

void WalletModel539e18_AddWallet(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<WalletModel539e18*>(ptr), "addWallet", Q_ARG(QWallet539e18*, static_cast<QWallet539e18*>(v0)));
}

void WalletModel539e18_EditWallet(void* ptr, int row, struct Moc_PackedString name, char encryptionEnabled, unsigned long long sky, unsigned long long coinHours)
{
	QMetaObject::invokeMethod(static_cast<WalletModel539e18*>(ptr), "editWallet", Q_ARG(qint32, row), Q_ARG(QString, QString::fromUtf8(name.data, name.len)), Q_ARG(bool, encryptionEnabled != 0), Q_ARG(quint64, sky), Q_ARG(quint64, coinHours));
}

void WalletModel539e18_RemoveWallet(void* ptr, int row)
{
	QMetaObject::invokeMethod(static_cast<WalletModel539e18*>(ptr), "removeWallet", Q_ARG(qint32, row));
}

void WalletModel539e18_LoadModel(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<WalletModel539e18*>(ptr), "loadModel", Q_ARG(QList<QWallet539e18*>, ({ QList<QWallet539e18*>* tmpP = static_cast<QList<QWallet539e18*>*>(v0); QList<QWallet539e18*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; })));
}

struct Moc_PackedList WalletModel539e18_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<WalletModel539e18*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList WalletModel539e18_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<WalletModel539e18*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void WalletModel539e18_SetRoles(void* ptr, void* roles)
{
	static_cast<WalletModel539e18*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void WalletModel539e18_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<WalletModel539e18*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void WalletModel539e18_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<WalletModel539e18*>(ptr), static_cast<void (WalletModel539e18::*)(QMap<qint32, QByteArray>)>(&WalletModel539e18::rolesChanged), static_cast<WalletModel539e18*>(ptr), static_cast<void (WalletModel539e18::*)(QMap<qint32, QByteArray>)>(&WalletModel539e18::Signal_RolesChanged));
}

void WalletModel539e18_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<WalletModel539e18*>(ptr), static_cast<void (WalletModel539e18::*)(QMap<qint32, QByteArray>)>(&WalletModel539e18::rolesChanged), static_cast<WalletModel539e18*>(ptr), static_cast<void (WalletModel539e18::*)(QMap<qint32, QByteArray>)>(&WalletModel539e18::Signal_RolesChanged));
}

void WalletModel539e18_RolesChanged(void* ptr, void* roles)
{
	static_cast<WalletModel539e18*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList WalletModel539e18_Wallets(void* ptr)
{
	return ({ QList<QWallet539e18*>* tmpValue = new QList<QWallet539e18*>(static_cast<WalletModel539e18*>(ptr)->wallets()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList WalletModel539e18_WalletsDefault(void* ptr)
{
	return ({ QList<QWallet539e18*>* tmpValue = new QList<QWallet539e18*>(static_cast<WalletModel539e18*>(ptr)->walletsDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void WalletModel539e18_SetWallets(void* ptr, void* wallets)
{
	static_cast<WalletModel539e18*>(ptr)->setWallets(({ QList<QWallet539e18*>* tmpP = static_cast<QList<QWallet539e18*>*>(wallets); QList<QWallet539e18*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void WalletModel539e18_SetWalletsDefault(void* ptr, void* wallets)
{
	static_cast<WalletModel539e18*>(ptr)->setWalletsDefault(({ QList<QWallet539e18*>* tmpP = static_cast<QList<QWallet539e18*>*>(wallets); QList<QWallet539e18*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void WalletModel539e18_ConnectWalletsChanged(void* ptr)
{
	QObject::connect(static_cast<WalletModel539e18*>(ptr), static_cast<void (WalletModel539e18::*)(QList<QWallet539e18*>)>(&WalletModel539e18::walletsChanged), static_cast<WalletModel539e18*>(ptr), static_cast<void (WalletModel539e18::*)(QList<QWallet539e18*>)>(&WalletModel539e18::Signal_WalletsChanged));
}

void WalletModel539e18_DisconnectWalletsChanged(void* ptr)
{
	QObject::disconnect(static_cast<WalletModel539e18*>(ptr), static_cast<void (WalletModel539e18::*)(QList<QWallet539e18*>)>(&WalletModel539e18::walletsChanged), static_cast<WalletModel539e18*>(ptr), static_cast<void (WalletModel539e18::*)(QList<QWallet539e18*>)>(&WalletModel539e18::Signal_WalletsChanged));
}

void WalletModel539e18_WalletsChanged(void* ptr, void* wallets)
{
	static_cast<WalletModel539e18*>(ptr)->walletsChanged(({ QList<QWallet539e18*>* tmpP = static_cast<QList<QWallet539e18*>*>(wallets); QList<QWallet539e18*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int WalletModel539e18_Count(void* ptr)
{
	return static_cast<WalletModel539e18*>(ptr)->count();
}

int WalletModel539e18_CountDefault(void* ptr)
{
	return static_cast<WalletModel539e18*>(ptr)->countDefault();
}

void WalletModel539e18_SetCount(void* ptr, int count)
{
	static_cast<WalletModel539e18*>(ptr)->setCount(count);
}

void WalletModel539e18_SetCountDefault(void* ptr, int count)
{
	static_cast<WalletModel539e18*>(ptr)->setCountDefault(count);
}

void WalletModel539e18_ConnectCountChanged(void* ptr)
{
	QObject::connect(static_cast<WalletModel539e18*>(ptr), static_cast<void (WalletModel539e18::*)(qint32)>(&WalletModel539e18::countChanged), static_cast<WalletModel539e18*>(ptr), static_cast<void (WalletModel539e18::*)(qint32)>(&WalletModel539e18::Signal_CountChanged));
}

void WalletModel539e18_DisconnectCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<WalletModel539e18*>(ptr), static_cast<void (WalletModel539e18::*)(qint32)>(&WalletModel539e18::countChanged), static_cast<WalletModel539e18*>(ptr), static_cast<void (WalletModel539e18::*)(qint32)>(&WalletModel539e18::Signal_CountChanged));
}

void WalletModel539e18_CountChanged(void* ptr, int count)
{
	static_cast<WalletModel539e18*>(ptr)->countChanged(count);
}

int WalletModel539e18_WalletModel539e18_QRegisterMetaType()
{
	return qRegisterMetaType<WalletModel539e18*>();
}

int WalletModel539e18_WalletModel539e18_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<WalletModel539e18*>(const_cast<const char*>(typeName));
}

int WalletModel539e18_WalletModel539e18_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletModel539e18>();
#else
	return 0;
#endif
}

int WalletModel539e18_WalletModel539e18_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletModel539e18>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int WalletModel539e18_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel539e18_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel539e18_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int WalletModel539e18_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel539e18_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel539e18_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int WalletModel539e18_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel539e18_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel539e18_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* WalletModel539e18___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel539e18___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel539e18___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModel539e18___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel539e18___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel539e18___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int WalletModel539e18___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void WalletModel539e18___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* WalletModel539e18___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* WalletModel539e18___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel539e18___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* WalletModel539e18___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList WalletModel539e18___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel539e18___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel539e18___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* WalletModel539e18___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* WalletModel539e18___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel539e18___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* WalletModel539e18___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* WalletModel539e18___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel539e18___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel539e18___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModel539e18___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel539e18___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel539e18___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModel539e18___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel539e18___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel539e18___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModel539e18___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void WalletModel539e18___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModel539e18___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList WalletModel539e18___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel539e18___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel539e18___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* WalletModel539e18___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList WalletModel539e18___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int WalletModel539e18_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel539e18_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel539e18_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int WalletModel539e18_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel539e18_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel539e18_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* WalletModel539e18___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel539e18___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModel539e18___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* WalletModel539e18___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel539e18___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* WalletModel539e18___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* WalletModel539e18___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel539e18___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModel539e18___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletModel539e18___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel539e18___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModel539e18___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletModel539e18___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel539e18___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModel539e18___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletModel539e18___loadModel_v0_atList(void* ptr, int i)
{
	return ({QWallet539e18* tmp = static_cast<QList<QWallet539e18*>*>(ptr)->at(i); if (i == static_cast<QList<QWallet539e18*>*>(ptr)->size()-1) { static_cast<QList<QWallet539e18*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel539e18___loadModel_v0_setList(void* ptr, void* i)
{
	static_cast<QList<QWallet539e18*>*>(ptr)->append(static_cast<QWallet539e18*>(i));
}

void* WalletModel539e18___loadModel_v0_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWallet539e18*>();
}

void* WalletModel539e18___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel539e18___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModel539e18___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList WalletModel539e18___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel539e18___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel539e18___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModel539e18___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList WalletModel539e18___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel539e18___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel539e18___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModel539e18___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList WalletModel539e18___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel539e18___wallets_atList(void* ptr, int i)
{
	return ({QWallet539e18* tmp = static_cast<QList<QWallet539e18*>*>(ptr)->at(i); if (i == static_cast<QList<QWallet539e18*>*>(ptr)->size()-1) { static_cast<QList<QWallet539e18*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel539e18___wallets_setList(void* ptr, void* i)
{
	static_cast<QList<QWallet539e18*>*>(ptr)->append(static_cast<QWallet539e18*>(i));
}

void* WalletModel539e18___wallets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWallet539e18*>();
}

void* WalletModel539e18___setWallets_wallets_atList(void* ptr, int i)
{
	return ({QWallet539e18* tmp = static_cast<QList<QWallet539e18*>*>(ptr)->at(i); if (i == static_cast<QList<QWallet539e18*>*>(ptr)->size()-1) { static_cast<QList<QWallet539e18*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel539e18___setWallets_wallets_setList(void* ptr, void* i)
{
	static_cast<QList<QWallet539e18*>*>(ptr)->append(static_cast<QWallet539e18*>(i));
}

void* WalletModel539e18___setWallets_wallets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWallet539e18*>();
}

void* WalletModel539e18___walletsChanged_wallets_atList(void* ptr, int i)
{
	return ({QWallet539e18* tmp = static_cast<QList<QWallet539e18*>*>(ptr)->at(i); if (i == static_cast<QList<QWallet539e18*>*>(ptr)->size()-1) { static_cast<QList<QWallet539e18*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel539e18___walletsChanged_wallets_setList(void* ptr, void* i)
{
	static_cast<QList<QWallet539e18*>*>(ptr)->append(static_cast<QWallet539e18*>(i));
}

void* WalletModel539e18___walletsChanged_wallets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWallet539e18*>();
}

int WalletModel539e18_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel539e18_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* WalletModel539e18_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int WalletModel539e18_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel539e18_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* WalletModel539e18_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int WalletModel539e18_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel539e18_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* WalletModel539e18_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* WalletModel539e18_NewWalletModel(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new WalletModel539e18(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new WalletModel539e18(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new WalletModel539e18(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new WalletModel539e18(static_cast<QWindow*>(parent));
	} else {
		return new WalletModel539e18(static_cast<QObject*>(parent));
	}
}

void WalletModel539e18_DestroyWalletModel(void* ptr)
{
	static_cast<WalletModel539e18*>(ptr)->~WalletModel539e18();
}

void WalletModel539e18_DestroyWalletModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char WalletModel539e18_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

long long WalletModel539e18_FlagsDefault(void* ptr, void* index)
{
	return static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

void* WalletModel539e18_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* WalletModel539e18_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* WalletModel539e18_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

char WalletModel539e18_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char WalletModel539e18_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int WalletModel539e18_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* WalletModel539e18_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void WalletModel539e18_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char WalletModel539e18_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* WalletModel539e18_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

char WalletModel539e18_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char WalletModel539e18_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct Moc_PackedList WalletModel539e18_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList WalletModel539e18_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel539e18_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct Moc_PackedString WalletModel539e18_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t2c50d2 = static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::mimeTypes().join("¡¦!").toUtf8(); Moc_PackedString { const_cast<char*>(t2c50d2.prepend("WHITESPACE").constData()+10), t2c50d2.size()-10 }; });
}

char WalletModel539e18_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char WalletModel539e18_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* WalletModel539e18_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

char WalletModel539e18_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char WalletModel539e18_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void WalletModel539e18_ResetInternalDataDefault(void* ptr)
{
	static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::resetInternalData();
}

void WalletModel539e18_RevertDefault(void* ptr)
{
	static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::revert();
}

struct Moc_PackedList WalletModel539e18_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int WalletModel539e18_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char WalletModel539e18_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char WalletModel539e18_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char WalletModel539e18_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void WalletModel539e18_SortDefault(void* ptr, int column, long long order)
{
	static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* WalletModel539e18_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char WalletModel539e18_SubmitDefault(void* ptr)
{
	return static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::submit();
}

long long WalletModel539e18_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long WalletModel539e18_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::supportedDropActions();
}

void WalletModel539e18_ChildEventDefault(void* ptr, void* event)
{
	static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void WalletModel539e18_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void WalletModel539e18_CustomEventDefault(void* ptr, void* event)
{
	static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void WalletModel539e18_DeleteLaterDefault(void* ptr)
{
	static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::deleteLater();
}

void WalletModel539e18_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char WalletModel539e18_EventDefault(void* ptr, void* e)
{
	return static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char WalletModel539e18_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void WalletModel539e18_TimerEventDefault(void* ptr, void* event)
{
	static_cast<WalletModel539e18*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
