

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
class AddressDetailsfb4c44: public QObject
{
Q_OBJECT
Q_PROPERTY(QString address READ address WRITE setAddress NOTIFY addressChanged)
Q_PROPERTY(QString addressSky READ addressSky WRITE setAddressSky NOTIFY addressSkyChanged)
Q_PROPERTY(QString addressCoinHours READ addressCoinHours WRITE setAddressCoinHours NOTIFY addressCoinHoursChanged)
public:
	AddressDetailsfb4c44(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");AddressDetailsfb4c44_AddressDetailsfb4c44_QRegisterMetaType();AddressDetailsfb4c44_AddressDetailsfb4c44_QRegisterMetaTypes();callbackAddressDetailsfb4c44_Constructor(this);};
	QString address() { return ({ Moc_PackedString tempVal = callbackAddressDetailsfb4c44_Address(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setAddress(QString address) { QByteArray tc66218 = address.toUtf8(); Moc_PackedString addressPacked = { const_cast<char*>(tc66218.prepend("WHITESPACE").constData()+10), tc66218.size()-10 };callbackAddressDetailsfb4c44_SetAddress(this, addressPacked); };
	void Signal_AddressChanged(QString address) { QByteArray tc66218 = address.toUtf8(); Moc_PackedString addressPacked = { const_cast<char*>(tc66218.prepend("WHITESPACE").constData()+10), tc66218.size()-10 };callbackAddressDetailsfb4c44_AddressChanged(this, addressPacked); };
	QString addressSky() { return ({ Moc_PackedString tempVal = callbackAddressDetailsfb4c44_AddressSky(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setAddressSky(QString addressSky) { QByteArray td2b9f3 = addressSky.toUtf8(); Moc_PackedString addressSkyPacked = { const_cast<char*>(td2b9f3.prepend("WHITESPACE").constData()+10), td2b9f3.size()-10 };callbackAddressDetailsfb4c44_SetAddressSky(this, addressSkyPacked); };
	void Signal_AddressSkyChanged(QString addressSky) { QByteArray td2b9f3 = addressSky.toUtf8(); Moc_PackedString addressSkyPacked = { const_cast<char*>(td2b9f3.prepend("WHITESPACE").constData()+10), td2b9f3.size()-10 };callbackAddressDetailsfb4c44_AddressSkyChanged(this, addressSkyPacked); };
	QString addressCoinHours() { return ({ Moc_PackedString tempVal = callbackAddressDetailsfb4c44_AddressCoinHours(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setAddressCoinHours(QString addressCoinHours) { QByteArray t65df47 = addressCoinHours.toUtf8(); Moc_PackedString addressCoinHoursPacked = { const_cast<char*>(t65df47.prepend("WHITESPACE").constData()+10), t65df47.size()-10 };callbackAddressDetailsfb4c44_SetAddressCoinHours(this, addressCoinHoursPacked); };
	void Signal_AddressCoinHoursChanged(QString addressCoinHours) { QByteArray t65df47 = addressCoinHours.toUtf8(); Moc_PackedString addressCoinHoursPacked = { const_cast<char*>(t65df47.prepend("WHITESPACE").constData()+10), t65df47.size()-10 };callbackAddressDetailsfb4c44_AddressCoinHoursChanged(this, addressCoinHoursPacked); };
	 ~AddressDetailsfb4c44() { callbackAddressDetailsfb4c44_DestroyAddressDetails(this); };
	void childEvent(QChildEvent * event) { callbackAddressDetailsfb4c44_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackAddressDetailsfb4c44_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackAddressDetailsfb4c44_CustomEvent(this, event); };
	void deleteLater() { callbackAddressDetailsfb4c44_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackAddressDetailsfb4c44_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackAddressDetailsfb4c44_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackAddressDetailsfb4c44_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackAddressDetailsfb4c44_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackAddressDetailsfb4c44_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackAddressDetailsfb4c44_TimerEvent(this, event); };
	QString addressDefault() { return _address; };
	void setAddressDefault(QString p) { if (p != _address) { _address = p; addressChanged(_address); } };
	QString addressSkyDefault() { return _addressSky; };
	void setAddressSkyDefault(QString p) { if (p != _addressSky) { _addressSky = p; addressSkyChanged(_addressSky); } };
	QString addressCoinHoursDefault() { return _addressCoinHours; };
	void setAddressCoinHoursDefault(QString p) { if (p != _addressCoinHours) { _addressCoinHours = p; addressCoinHoursChanged(_addressCoinHours); } };
signals:
	void addressChanged(QString address);
	void addressSkyChanged(QString addressSky);
	void addressCoinHoursChanged(QString addressCoinHours);
public slots:
private:
	QString _address;
	QString _addressSky;
	QString _addressCoinHours;
};

Q_DECLARE_METATYPE(AddressDetailsfb4c44*)


void AddressDetailsfb4c44_AddressDetailsfb4c44_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class AddressListfb4c44: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<AddressDetailsfb4c44*> addresses READ addresses WRITE setAddresses NOTIFY addressesChanged)
public:
	AddressListfb4c44(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");AddressListfb4c44_AddressListfb4c44_QRegisterMetaType();AddressListfb4c44_AddressListfb4c44_QRegisterMetaTypes();callbackAddressListfb4c44_Constructor(this);};
	void Signal_AddAddress(AddressDetailsfb4c44* transaction) { callbackAddressListfb4c44_AddAddress(this, transaction); };
	void Signal_RemoveAddress(qint32 index) { callbackAddressListfb4c44_RemoveAddress(this, index); };
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbackAddressListfb4c44_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbackAddressListfb4c44_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbackAddressListfb4c44_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<AddressDetailsfb4c44*> addresses() { return ({ QList<AddressDetailsfb4c44*>* tmpP = static_cast<QList<AddressDetailsfb4c44*>*>(callbackAddressListfb4c44_Addresses(this)); QList<AddressDetailsfb4c44*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setAddresses(QList<AddressDetailsfb4c44*> addresses) { callbackAddressListfb4c44_SetAddresses(this, ({ QList<AddressDetailsfb4c44*>* tmpValue = new QList<AddressDetailsfb4c44*>(addresses); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_AddressesChanged(QList<AddressDetailsfb4c44*> addresses) { callbackAddressListfb4c44_AddressesChanged(this, ({ QList<AddressDetailsfb4c44*>* tmpValue = new QList<AddressDetailsfb4c44*>(addresses); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	 ~AddressListfb4c44() { callbackAddressListfb4c44_DestroyAddressList(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackAddressListfb4c44_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackAddressListfb4c44_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackAddressListfb4c44_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackAddressListfb4c44_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackAddressListfb4c44_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackAddressListfb4c44_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackAddressListfb4c44_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackAddressListfb4c44_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackAddressListfb4c44_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackAddressListfb4c44_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackAddressListfb4c44_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackAddressListfb4c44_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackAddressListfb4c44_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackAddressListfb4c44_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackAddressListfb4c44_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackAddressListfb4c44_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = new QVector<int>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackAddressListfb4c44_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackAddressListfb4c44_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackAddressListfb4c44_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackAddressListfb4c44_HeaderDataChanged(this, orientation, first, last); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackAddressListfb4c44_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackAddressListfb4c44_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackAddressListfb4c44_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackAddressListfb4c44_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackAddressListfb4c44_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackAddressListfb4c44_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackAddressListfb4c44_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackAddressListfb4c44_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackAddressListfb4c44_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackAddressListfb4c44_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackAddressListfb4c44_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackAddressListfb4c44_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackAddressListfb4c44_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackAddressListfb4c44_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackAddressListfb4c44_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackAddressListfb4c44_ResetInternalData(this); };
	void revert() { callbackAddressListfb4c44_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackAddressListfb4c44_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackAddressListfb4c44_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackAddressListfb4c44_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackAddressListfb4c44_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackAddressListfb4c44_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackAddressListfb4c44_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackAddressListfb4c44_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackAddressListfb4c44_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackAddressListfb4c44_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackAddressListfb4c44_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackAddressListfb4c44_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackAddressListfb4c44_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackAddressListfb4c44_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackAddressListfb4c44_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackAddressListfb4c44_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackAddressListfb4c44_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackAddressListfb4c44_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackAddressListfb4c44_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackAddressListfb4c44_CustomEvent(this, event); };
	void deleteLater() { callbackAddressListfb4c44_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackAddressListfb4c44_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackAddressListfb4c44_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackAddressListfb4c44_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackAddressListfb4c44_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackAddressListfb4c44_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackAddressListfb4c44_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<AddressDetailsfb4c44*> addressesDefault() { return _addresses; };
	void setAddressesDefault(QList<AddressDetailsfb4c44*> p) { if (p != _addresses) { _addresses = p; addressesChanged(_addresses); } };
signals:
	void addAddress(AddressDetailsfb4c44* transaction);
	void removeAddress(qint32 index);
	void rolesChanged(type378cdd roles);
	void addressesChanged(QList<AddressDetailsfb4c44*> addresses);
public slots:
private:
	type378cdd _roles;
	QList<AddressDetailsfb4c44*> _addresses;
};

Q_DECLARE_METATYPE(AddressListfb4c44*)


void AddressListfb4c44_AddressListfb4c44_QRegisterMetaTypes() {
	qRegisterMetaType<type378cdd>("type378cdd");
	qRegisterMetaType<QList<QObject*>>("QList<AddressDetailsfb4c44*>");
}

struct Moc_PackedString AddressDetailsfb4c44_Address(void* ptr)
{
	return ({ QByteArray t686357 = static_cast<AddressDetailsfb4c44*>(ptr)->address().toUtf8(); Moc_PackedString { const_cast<char*>(t686357.prepend("WHITESPACE").constData()+10), t686357.size()-10 }; });
}

struct Moc_PackedString AddressDetailsfb4c44_AddressDefault(void* ptr)
{
	return ({ QByteArray tbaed36 = static_cast<AddressDetailsfb4c44*>(ptr)->addressDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tbaed36.prepend("WHITESPACE").constData()+10), tbaed36.size()-10 }; });
}

void AddressDetailsfb4c44_SetAddress(void* ptr, struct Moc_PackedString address)
{
	static_cast<AddressDetailsfb4c44*>(ptr)->setAddress(QString::fromUtf8(address.data, address.len));
}

void AddressDetailsfb4c44_SetAddressDefault(void* ptr, struct Moc_PackedString address)
{
	static_cast<AddressDetailsfb4c44*>(ptr)->setAddressDefault(QString::fromUtf8(address.data, address.len));
}

void AddressDetailsfb4c44_ConnectAddressChanged(void* ptr)
{
	QObject::connect(static_cast<AddressDetailsfb4c44*>(ptr), static_cast<void (AddressDetailsfb4c44::*)(QString)>(&AddressDetailsfb4c44::addressChanged), static_cast<AddressDetailsfb4c44*>(ptr), static_cast<void (AddressDetailsfb4c44::*)(QString)>(&AddressDetailsfb4c44::Signal_AddressChanged));
}

void AddressDetailsfb4c44_DisconnectAddressChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressDetailsfb4c44*>(ptr), static_cast<void (AddressDetailsfb4c44::*)(QString)>(&AddressDetailsfb4c44::addressChanged), static_cast<AddressDetailsfb4c44*>(ptr), static_cast<void (AddressDetailsfb4c44::*)(QString)>(&AddressDetailsfb4c44::Signal_AddressChanged));
}

void AddressDetailsfb4c44_AddressChanged(void* ptr, struct Moc_PackedString address)
{
	static_cast<AddressDetailsfb4c44*>(ptr)->addressChanged(QString::fromUtf8(address.data, address.len));
}

struct Moc_PackedString AddressDetailsfb4c44_AddressSky(void* ptr)
{
	return ({ QByteArray tb04c84 = static_cast<AddressDetailsfb4c44*>(ptr)->addressSky().toUtf8(); Moc_PackedString { const_cast<char*>(tb04c84.prepend("WHITESPACE").constData()+10), tb04c84.size()-10 }; });
}

struct Moc_PackedString AddressDetailsfb4c44_AddressSkyDefault(void* ptr)
{
	return ({ QByteArray tcdca54 = static_cast<AddressDetailsfb4c44*>(ptr)->addressSkyDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tcdca54.prepend("WHITESPACE").constData()+10), tcdca54.size()-10 }; });
}

void AddressDetailsfb4c44_SetAddressSky(void* ptr, struct Moc_PackedString addressSky)
{
	static_cast<AddressDetailsfb4c44*>(ptr)->setAddressSky(QString::fromUtf8(addressSky.data, addressSky.len));
}

void AddressDetailsfb4c44_SetAddressSkyDefault(void* ptr, struct Moc_PackedString addressSky)
{
	static_cast<AddressDetailsfb4c44*>(ptr)->setAddressSkyDefault(QString::fromUtf8(addressSky.data, addressSky.len));
}

void AddressDetailsfb4c44_ConnectAddressSkyChanged(void* ptr)
{
	QObject::connect(static_cast<AddressDetailsfb4c44*>(ptr), static_cast<void (AddressDetailsfb4c44::*)(QString)>(&AddressDetailsfb4c44::addressSkyChanged), static_cast<AddressDetailsfb4c44*>(ptr), static_cast<void (AddressDetailsfb4c44::*)(QString)>(&AddressDetailsfb4c44::Signal_AddressSkyChanged));
}

void AddressDetailsfb4c44_DisconnectAddressSkyChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressDetailsfb4c44*>(ptr), static_cast<void (AddressDetailsfb4c44::*)(QString)>(&AddressDetailsfb4c44::addressSkyChanged), static_cast<AddressDetailsfb4c44*>(ptr), static_cast<void (AddressDetailsfb4c44::*)(QString)>(&AddressDetailsfb4c44::Signal_AddressSkyChanged));
}

void AddressDetailsfb4c44_AddressSkyChanged(void* ptr, struct Moc_PackedString addressSky)
{
	static_cast<AddressDetailsfb4c44*>(ptr)->addressSkyChanged(QString::fromUtf8(addressSky.data, addressSky.len));
}

struct Moc_PackedString AddressDetailsfb4c44_AddressCoinHours(void* ptr)
{
	return ({ QByteArray t505be8 = static_cast<AddressDetailsfb4c44*>(ptr)->addressCoinHours().toUtf8(); Moc_PackedString { const_cast<char*>(t505be8.prepend("WHITESPACE").constData()+10), t505be8.size()-10 }; });
}

struct Moc_PackedString AddressDetailsfb4c44_AddressCoinHoursDefault(void* ptr)
{
	return ({ QByteArray t272ab0 = static_cast<AddressDetailsfb4c44*>(ptr)->addressCoinHoursDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t272ab0.prepend("WHITESPACE").constData()+10), t272ab0.size()-10 }; });
}

void AddressDetailsfb4c44_SetAddressCoinHours(void* ptr, struct Moc_PackedString addressCoinHours)
{
	static_cast<AddressDetailsfb4c44*>(ptr)->setAddressCoinHours(QString::fromUtf8(addressCoinHours.data, addressCoinHours.len));
}

void AddressDetailsfb4c44_SetAddressCoinHoursDefault(void* ptr, struct Moc_PackedString addressCoinHours)
{
	static_cast<AddressDetailsfb4c44*>(ptr)->setAddressCoinHoursDefault(QString::fromUtf8(addressCoinHours.data, addressCoinHours.len));
}

void AddressDetailsfb4c44_ConnectAddressCoinHoursChanged(void* ptr)
{
	QObject::connect(static_cast<AddressDetailsfb4c44*>(ptr), static_cast<void (AddressDetailsfb4c44::*)(QString)>(&AddressDetailsfb4c44::addressCoinHoursChanged), static_cast<AddressDetailsfb4c44*>(ptr), static_cast<void (AddressDetailsfb4c44::*)(QString)>(&AddressDetailsfb4c44::Signal_AddressCoinHoursChanged));
}

void AddressDetailsfb4c44_DisconnectAddressCoinHoursChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressDetailsfb4c44*>(ptr), static_cast<void (AddressDetailsfb4c44::*)(QString)>(&AddressDetailsfb4c44::addressCoinHoursChanged), static_cast<AddressDetailsfb4c44*>(ptr), static_cast<void (AddressDetailsfb4c44::*)(QString)>(&AddressDetailsfb4c44::Signal_AddressCoinHoursChanged));
}

void AddressDetailsfb4c44_AddressCoinHoursChanged(void* ptr, struct Moc_PackedString addressCoinHours)
{
	static_cast<AddressDetailsfb4c44*>(ptr)->addressCoinHoursChanged(QString::fromUtf8(addressCoinHours.data, addressCoinHours.len));
}

int AddressDetailsfb4c44_AddressDetailsfb4c44_QRegisterMetaType()
{
	return qRegisterMetaType<AddressDetailsfb4c44*>();
}

int AddressDetailsfb4c44_AddressDetailsfb4c44_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<AddressDetailsfb4c44*>(const_cast<const char*>(typeName));
}

int AddressDetailsfb4c44_AddressDetailsfb4c44_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressDetailsfb4c44>();
#else
	return 0;
#endif
}

int AddressDetailsfb4c44_AddressDetailsfb4c44_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressDetailsfb4c44>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* AddressDetailsfb4c44___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressDetailsfb4c44___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressDetailsfb4c44___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* AddressDetailsfb4c44___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressDetailsfb4c44___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* AddressDetailsfb4c44___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* AddressDetailsfb4c44___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressDetailsfb4c44___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressDetailsfb4c44___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressDetailsfb4c44___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressDetailsfb4c44___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressDetailsfb4c44___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressDetailsfb4c44___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressDetailsfb4c44___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressDetailsfb4c44___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressDetailsfb4c44_NewAddressDetails(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new AddressDetailsfb4c44(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new AddressDetailsfb4c44(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new AddressDetailsfb4c44(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new AddressDetailsfb4c44(static_cast<QWindow*>(parent));
	} else {
		return new AddressDetailsfb4c44(static_cast<QObject*>(parent));
	}
}

void AddressDetailsfb4c44_DestroyAddressDetails(void* ptr)
{
	static_cast<AddressDetailsfb4c44*>(ptr)->~AddressDetailsfb4c44();
}

void AddressDetailsfb4c44_DestroyAddressDetailsDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void AddressDetailsfb4c44_ChildEventDefault(void* ptr, void* event)
{
	static_cast<AddressDetailsfb4c44*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void AddressDetailsfb4c44_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressDetailsfb4c44*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void AddressDetailsfb4c44_CustomEventDefault(void* ptr, void* event)
{
	static_cast<AddressDetailsfb4c44*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void AddressDetailsfb4c44_DeleteLaterDefault(void* ptr)
{
	static_cast<AddressDetailsfb4c44*>(ptr)->QObject::deleteLater();
}

void AddressDetailsfb4c44_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressDetailsfb4c44*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char AddressDetailsfb4c44_EventDefault(void* ptr, void* e)
{
	return static_cast<AddressDetailsfb4c44*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char AddressDetailsfb4c44_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<AddressDetailsfb4c44*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void AddressDetailsfb4c44_TimerEventDefault(void* ptr, void* event)
{
	static_cast<AddressDetailsfb4c44*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

void AddressListfb4c44_ConnectAddAddress(void* ptr)
{
	QObject::connect(static_cast<AddressListfb4c44*>(ptr), static_cast<void (AddressListfb4c44::*)(AddressDetailsfb4c44*)>(&AddressListfb4c44::addAddress), static_cast<AddressListfb4c44*>(ptr), static_cast<void (AddressListfb4c44::*)(AddressDetailsfb4c44*)>(&AddressListfb4c44::Signal_AddAddress));
}

void AddressListfb4c44_DisconnectAddAddress(void* ptr)
{
	QObject::disconnect(static_cast<AddressListfb4c44*>(ptr), static_cast<void (AddressListfb4c44::*)(AddressDetailsfb4c44*)>(&AddressListfb4c44::addAddress), static_cast<AddressListfb4c44*>(ptr), static_cast<void (AddressListfb4c44::*)(AddressDetailsfb4c44*)>(&AddressListfb4c44::Signal_AddAddress));
}

void AddressListfb4c44_AddAddress(void* ptr, void* transaction)
{
	static_cast<AddressListfb4c44*>(ptr)->addAddress(static_cast<AddressDetailsfb4c44*>(transaction));
}

void AddressListfb4c44_ConnectRemoveAddress(void* ptr)
{
	QObject::connect(static_cast<AddressListfb4c44*>(ptr), static_cast<void (AddressListfb4c44::*)(qint32)>(&AddressListfb4c44::removeAddress), static_cast<AddressListfb4c44*>(ptr), static_cast<void (AddressListfb4c44::*)(qint32)>(&AddressListfb4c44::Signal_RemoveAddress));
}

void AddressListfb4c44_DisconnectRemoveAddress(void* ptr)
{
	QObject::disconnect(static_cast<AddressListfb4c44*>(ptr), static_cast<void (AddressListfb4c44::*)(qint32)>(&AddressListfb4c44::removeAddress), static_cast<AddressListfb4c44*>(ptr), static_cast<void (AddressListfb4c44::*)(qint32)>(&AddressListfb4c44::Signal_RemoveAddress));
}

void AddressListfb4c44_RemoveAddress(void* ptr, int index)
{
	static_cast<AddressListfb4c44*>(ptr)->removeAddress(index);
}

struct Moc_PackedList AddressListfb4c44_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<AddressListfb4c44*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressListfb4c44_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<AddressListfb4c44*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void AddressListfb4c44_SetRoles(void* ptr, void* roles)
{
	static_cast<AddressListfb4c44*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void AddressListfb4c44_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<AddressListfb4c44*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void AddressListfb4c44_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<AddressListfb4c44*>(ptr), static_cast<void (AddressListfb4c44::*)(QMap<qint32, QByteArray>)>(&AddressListfb4c44::rolesChanged), static_cast<AddressListfb4c44*>(ptr), static_cast<void (AddressListfb4c44::*)(QMap<qint32, QByteArray>)>(&AddressListfb4c44::Signal_RolesChanged));
}

void AddressListfb4c44_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressListfb4c44*>(ptr), static_cast<void (AddressListfb4c44::*)(QMap<qint32, QByteArray>)>(&AddressListfb4c44::rolesChanged), static_cast<AddressListfb4c44*>(ptr), static_cast<void (AddressListfb4c44::*)(QMap<qint32, QByteArray>)>(&AddressListfb4c44::Signal_RolesChanged));
}

void AddressListfb4c44_RolesChanged(void* ptr, void* roles)
{
	static_cast<AddressListfb4c44*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList AddressListfb4c44_Addresses(void* ptr)
{
	return ({ QList<AddressDetailsfb4c44*>* tmpValue = new QList<AddressDetailsfb4c44*>(static_cast<AddressListfb4c44*>(ptr)->addresses()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressListfb4c44_AddressesDefault(void* ptr)
{
	return ({ QList<AddressDetailsfb4c44*>* tmpValue = new QList<AddressDetailsfb4c44*>(static_cast<AddressListfb4c44*>(ptr)->addressesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void AddressListfb4c44_SetAddresses(void* ptr, void* addresses)
{
	static_cast<AddressListfb4c44*>(ptr)->setAddresses(({ QList<AddressDetailsfb4c44*>* tmpP = static_cast<QList<AddressDetailsfb4c44*>*>(addresses); QList<AddressDetailsfb4c44*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void AddressListfb4c44_SetAddressesDefault(void* ptr, void* addresses)
{
	static_cast<AddressListfb4c44*>(ptr)->setAddressesDefault(({ QList<AddressDetailsfb4c44*>* tmpP = static_cast<QList<AddressDetailsfb4c44*>*>(addresses); QList<AddressDetailsfb4c44*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void AddressListfb4c44_ConnectAddressesChanged(void* ptr)
{
	QObject::connect(static_cast<AddressListfb4c44*>(ptr), static_cast<void (AddressListfb4c44::*)(QList<AddressDetailsfb4c44*>)>(&AddressListfb4c44::addressesChanged), static_cast<AddressListfb4c44*>(ptr), static_cast<void (AddressListfb4c44::*)(QList<AddressDetailsfb4c44*>)>(&AddressListfb4c44::Signal_AddressesChanged));
}

void AddressListfb4c44_DisconnectAddressesChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressListfb4c44*>(ptr), static_cast<void (AddressListfb4c44::*)(QList<AddressDetailsfb4c44*>)>(&AddressListfb4c44::addressesChanged), static_cast<AddressListfb4c44*>(ptr), static_cast<void (AddressListfb4c44::*)(QList<AddressDetailsfb4c44*>)>(&AddressListfb4c44::Signal_AddressesChanged));
}

void AddressListfb4c44_AddressesChanged(void* ptr, void* addresses)
{
	static_cast<AddressListfb4c44*>(ptr)->addressesChanged(({ QList<AddressDetailsfb4c44*>* tmpP = static_cast<QList<AddressDetailsfb4c44*>*>(addresses); QList<AddressDetailsfb4c44*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int AddressListfb4c44_AddressListfb4c44_QRegisterMetaType()
{
	return qRegisterMetaType<AddressListfb4c44*>();
}

int AddressListfb4c44_AddressListfb4c44_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<AddressListfb4c44*>(const_cast<const char*>(typeName));
}

int AddressListfb4c44_AddressListfb4c44_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressListfb4c44>();
#else
	return 0;
#endif
}

int AddressListfb4c44_AddressListfb4c44_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressListfb4c44>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int AddressListfb4c44_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressListfb4c44_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressListfb4c44_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressListfb4c44_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressListfb4c44_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressListfb4c44_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressListfb4c44_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressListfb4c44_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressListfb4c44_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* AddressListfb4c44___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressListfb4c44___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressListfb4c44___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressListfb4c44___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressListfb4c44___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressListfb4c44___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int AddressListfb4c44___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void AddressListfb4c44___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* AddressListfb4c44___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* AddressListfb4c44___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressListfb4c44___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* AddressListfb4c44___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList AddressListfb4c44___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressListfb4c44___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressListfb4c44___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* AddressListfb4c44___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* AddressListfb4c44___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressListfb4c44___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* AddressListfb4c44___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* AddressListfb4c44___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressListfb4c44___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressListfb4c44___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressListfb4c44___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressListfb4c44___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressListfb4c44___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressListfb4c44___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressListfb4c44___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressListfb4c44___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressListfb4c44___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void AddressListfb4c44___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressListfb4c44___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList AddressListfb4c44___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressListfb4c44___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressListfb4c44___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* AddressListfb4c44___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList AddressListfb4c44___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int AddressListfb4c44_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressListfb4c44_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressListfb4c44_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressListfb4c44_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressListfb4c44_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressListfb4c44_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* AddressListfb4c44___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressListfb4c44___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressListfb4c44___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* AddressListfb4c44___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressListfb4c44___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* AddressListfb4c44___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* AddressListfb4c44___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressListfb4c44___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressListfb4c44___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressListfb4c44___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressListfb4c44___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressListfb4c44___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressListfb4c44___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressListfb4c44___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressListfb4c44___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressListfb4c44___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressListfb4c44___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressListfb4c44___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressListfb4c44___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressListfb4c44___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressListfb4c44___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressListfb4c44___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressListfb4c44___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressListfb4c44___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressListfb4c44___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressListfb4c44___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressListfb4c44___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressListfb4c44___addresses_atList(void* ptr, int i)
{
	return ({AddressDetailsfb4c44* tmp = static_cast<QList<AddressDetailsfb4c44*>*>(ptr)->at(i); if (i == static_cast<QList<AddressDetailsfb4c44*>*>(ptr)->size()-1) { static_cast<QList<AddressDetailsfb4c44*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressListfb4c44___addresses_setList(void* ptr, void* i)
{
	static_cast<QList<AddressDetailsfb4c44*>*>(ptr)->append(static_cast<AddressDetailsfb4c44*>(i));
}

void* AddressListfb4c44___addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<AddressDetailsfb4c44*>();
}

void* AddressListfb4c44___setAddresses_addresses_atList(void* ptr, int i)
{
	return ({AddressDetailsfb4c44* tmp = static_cast<QList<AddressDetailsfb4c44*>*>(ptr)->at(i); if (i == static_cast<QList<AddressDetailsfb4c44*>*>(ptr)->size()-1) { static_cast<QList<AddressDetailsfb4c44*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressListfb4c44___setAddresses_addresses_setList(void* ptr, void* i)
{
	static_cast<QList<AddressDetailsfb4c44*>*>(ptr)->append(static_cast<AddressDetailsfb4c44*>(i));
}

void* AddressListfb4c44___setAddresses_addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<AddressDetailsfb4c44*>();
}

void* AddressListfb4c44___addressesChanged_addresses_atList(void* ptr, int i)
{
	return ({AddressDetailsfb4c44* tmp = static_cast<QList<AddressDetailsfb4c44*>*>(ptr)->at(i); if (i == static_cast<QList<AddressDetailsfb4c44*>*>(ptr)->size()-1) { static_cast<QList<AddressDetailsfb4c44*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressListfb4c44___addressesChanged_addresses_setList(void* ptr, void* i)
{
	static_cast<QList<AddressDetailsfb4c44*>*>(ptr)->append(static_cast<AddressDetailsfb4c44*>(i));
}

void* AddressListfb4c44___addressesChanged_addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<AddressDetailsfb4c44*>();
}

int AddressListfb4c44_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressListfb4c44_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressListfb4c44_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int AddressListfb4c44_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressListfb4c44_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressListfb4c44_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int AddressListfb4c44_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressListfb4c44_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressListfb4c44_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* AddressListfb4c44_NewAddressList(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new AddressListfb4c44(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new AddressListfb4c44(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new AddressListfb4c44(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new AddressListfb4c44(static_cast<QWindow*>(parent));
	} else {
		return new AddressListfb4c44(static_cast<QObject*>(parent));
	}
}

void AddressListfb4c44_DestroyAddressList(void* ptr)
{
	static_cast<AddressListfb4c44*>(ptr)->~AddressListfb4c44();
}

void AddressListfb4c44_DestroyAddressListDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char AddressListfb4c44_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

long long AddressListfb4c44_FlagsDefault(void* ptr, void* index)
{
	return static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

void* AddressListfb4c44_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* AddressListfb4c44_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* AddressListfb4c44_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

char AddressListfb4c44_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char AddressListfb4c44_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int AddressListfb4c44_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* AddressListfb4c44_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void AddressListfb4c44_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char AddressListfb4c44_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* AddressListfb4c44_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

char AddressListfb4c44_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char AddressListfb4c44_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct Moc_PackedList AddressListfb4c44_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressListfb4c44_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressListfb4c44_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct Moc_PackedString AddressListfb4c44_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t29e12e = static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::mimeTypes().join("¡¦!").toUtf8(); Moc_PackedString { const_cast<char*>(t29e12e.prepend("WHITESPACE").constData()+10), t29e12e.size()-10 }; });
}

char AddressListfb4c44_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char AddressListfb4c44_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* AddressListfb4c44_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

char AddressListfb4c44_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char AddressListfb4c44_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void AddressListfb4c44_ResetInternalDataDefault(void* ptr)
{
	static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::resetInternalData();
}

void AddressListfb4c44_RevertDefault(void* ptr)
{
	static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::revert();
}

struct Moc_PackedList AddressListfb4c44_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int AddressListfb4c44_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char AddressListfb4c44_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char AddressListfb4c44_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char AddressListfb4c44_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void AddressListfb4c44_SortDefault(void* ptr, int column, long long order)
{
	static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* AddressListfb4c44_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char AddressListfb4c44_SubmitDefault(void* ptr)
{
	return static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::submit();
}

long long AddressListfb4c44_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long AddressListfb4c44_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::supportedDropActions();
}

void AddressListfb4c44_ChildEventDefault(void* ptr, void* event)
{
	static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void AddressListfb4c44_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void AddressListfb4c44_CustomEventDefault(void* ptr, void* event)
{
	static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void AddressListfb4c44_DeleteLaterDefault(void* ptr)
{
	static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::deleteLater();
}

void AddressListfb4c44_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char AddressListfb4c44_EventDefault(void* ptr, void* e)
{
	return static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char AddressListfb4c44_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void AddressListfb4c44_TimerEventDefault(void* ptr, void* event)
{
	static_cast<AddressListfb4c44*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
