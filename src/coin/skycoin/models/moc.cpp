

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
class AddressesModelbaab1e: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<QAddressbaab1e*> addresses READ addresses WRITE setAddresses NOTIFY addressesChanged)
Q_PROPERTY(qint32 loaded READ loaded WRITE setLoaded NOTIFY loadedChanged)
public:
	AddressesModelbaab1e(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");AddressesModelbaab1e_AddressesModelbaab1e_QRegisterMetaType();AddressesModelbaab1e_AddressesModelbaab1e_QRegisterMetaTypes();callbackAddressesModelbaab1e_Constructor(this);};
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbackAddressesModelbaab1e_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbackAddressesModelbaab1e_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbackAddressesModelbaab1e_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<QAddressbaab1e*> addresses() { return ({ QList<QAddressbaab1e*>* tmpP = static_cast<QList<QAddressbaab1e*>*>(callbackAddressesModelbaab1e_Addresses(this)); QList<QAddressbaab1e*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setAddresses(QList<QAddressbaab1e*> addresses) { callbackAddressesModelbaab1e_SetAddresses(this, ({ QList<QAddressbaab1e*>* tmpValue = new QList<QAddressbaab1e*>(addresses); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_AddressesChanged(QList<QAddressbaab1e*> addresses) { callbackAddressesModelbaab1e_AddressesChanged(this, ({ QList<QAddressbaab1e*>* tmpValue = new QList<QAddressbaab1e*>(addresses); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	qint32 loaded() { return callbackAddressesModelbaab1e_Loaded(this); };
	void setLoaded(qint32 loaded) { callbackAddressesModelbaab1e_SetLoaded(this, loaded); };
	void Signal_LoadedChanged(qint32 loaded) { callbackAddressesModelbaab1e_LoadedChanged(this, loaded); };
	 ~AddressesModelbaab1e() { callbackAddressesModelbaab1e_DestroyAddressesModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackAddressesModelbaab1e_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackAddressesModelbaab1e_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackAddressesModelbaab1e_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackAddressesModelbaab1e_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackAddressesModelbaab1e_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackAddressesModelbaab1e_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackAddressesModelbaab1e_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackAddressesModelbaab1e_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackAddressesModelbaab1e_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackAddressesModelbaab1e_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackAddressesModelbaab1e_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackAddressesModelbaab1e_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackAddressesModelbaab1e_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackAddressesModelbaab1e_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackAddressesModelbaab1e_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackAddressesModelbaab1e_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = new QVector<int>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackAddressesModelbaab1e_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackAddressesModelbaab1e_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackAddressesModelbaab1e_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackAddressesModelbaab1e_HeaderDataChanged(this, orientation, first, last); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackAddressesModelbaab1e_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackAddressesModelbaab1e_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackAddressesModelbaab1e_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackAddressesModelbaab1e_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackAddressesModelbaab1e_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackAddressesModelbaab1e_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackAddressesModelbaab1e_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackAddressesModelbaab1e_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackAddressesModelbaab1e_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackAddressesModelbaab1e_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackAddressesModelbaab1e_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackAddressesModelbaab1e_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackAddressesModelbaab1e_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackAddressesModelbaab1e_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackAddressesModelbaab1e_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackAddressesModelbaab1e_ResetInternalData(this); };
	void revert() { callbackAddressesModelbaab1e_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackAddressesModelbaab1e_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackAddressesModelbaab1e_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackAddressesModelbaab1e_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackAddressesModelbaab1e_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackAddressesModelbaab1e_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackAddressesModelbaab1e_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackAddressesModelbaab1e_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackAddressesModelbaab1e_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackAddressesModelbaab1e_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackAddressesModelbaab1e_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackAddressesModelbaab1e_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackAddressesModelbaab1e_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackAddressesModelbaab1e_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackAddressesModelbaab1e_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackAddressesModelbaab1e_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackAddressesModelbaab1e_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackAddressesModelbaab1e_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackAddressesModelbaab1e_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackAddressesModelbaab1e_CustomEvent(this, event); };
	void deleteLater() { callbackAddressesModelbaab1e_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackAddressesModelbaab1e_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackAddressesModelbaab1e_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackAddressesModelbaab1e_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackAddressesModelbaab1e_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackAddressesModelbaab1e_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackAddressesModelbaab1e_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<QAddressbaab1e*> addressesDefault() { return _addresses; };
	void setAddressesDefault(QList<QAddressbaab1e*> p) { if (p != _addresses) { _addresses = p; addressesChanged(_addresses); } };
	qint32 loadedDefault() { return _loaded; };
	void setLoadedDefault(qint32 p) { if (p != _loaded) { _loaded = p; loadedChanged(_loaded); } };
signals:
	void rolesChanged(type378cdd roles);
	void addressesChanged(QList<QAddressbaab1e*> addresses);
	void loadedChanged(qint32 loaded);
public slots:
	void addAddress(QAddressbaab1e* v0) { callbackAddressesModelbaab1e_AddAddress(this, v0); };
	void removeAddress(qint32 v0) { callbackAddressesModelbaab1e_RemoveAddress(this, v0); };
	void editAddress(qint32 v0, QString v1, quint64 v2, quint64 v3) { QByteArray t5a6df7 = v1.toUtf8(); Moc_PackedString v1Packed = { const_cast<char*>(t5a6df7.prepend("WHITESPACE").constData()+10), t5a6df7.size()-10 };callbackAddressesModelbaab1e_EditAddress(this, v0, v1Packed, v2, v3); };
	void loadModel(QString v0) { QByteArray tea1dd7 = v0.toUtf8(); Moc_PackedString v0Packed = { const_cast<char*>(tea1dd7.prepend("WHITESPACE").constData()+10), tea1dd7.size()-10 };callbackAddressesModelbaab1e_LoadModel(this, v0Packed); };
private:
	type378cdd _roles;
	QList<QAddressbaab1e*> _addresses;
	qint32 _loaded;
};

Q_DECLARE_METATYPE(AddressesModelbaab1e*)


void AddressesModelbaab1e_AddressesModelbaab1e_QRegisterMetaTypes() {
	qRegisterMetaType<type378cdd>("type378cdd");
	qRegisterMetaType<QList<QObject*>>("QList<QAddressbaab1e*>");
}

class QAddressbaab1e: public QObject
{
Q_OBJECT
Q_PROPERTY(QString address READ address WRITE setAddress NOTIFY addressChanged)
Q_PROPERTY(quint64 addressSky READ addressSky WRITE setAddressSky NOTIFY addressSkyChanged)
Q_PROPERTY(quint64 addressCoinHours READ addressCoinHours WRITE setAddressCoinHours NOTIFY addressCoinHoursChanged)
public:
	QAddressbaab1e(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QAddressbaab1e_QAddressbaab1e_QRegisterMetaType();QAddressbaab1e_QAddressbaab1e_QRegisterMetaTypes();callbackQAddressbaab1e_Constructor(this);};
	QString address() { return ({ Moc_PackedString tempVal = callbackQAddressbaab1e_Address(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setAddress(QString address) { QByteArray tc66218 = address.toUtf8(); Moc_PackedString addressPacked = { const_cast<char*>(tc66218.prepend("WHITESPACE").constData()+10), tc66218.size()-10 };callbackQAddressbaab1e_SetAddress(this, addressPacked); };
	void Signal_AddressChanged(QString address) { QByteArray tc66218 = address.toUtf8(); Moc_PackedString addressPacked = { const_cast<char*>(tc66218.prepend("WHITESPACE").constData()+10), tc66218.size()-10 };callbackQAddressbaab1e_AddressChanged(this, addressPacked); };
	quint64 addressSky() { return callbackQAddressbaab1e_AddressSky(this); };
	void setAddressSky(quint64 addressSky) { callbackQAddressbaab1e_SetAddressSky(this, addressSky); };
	void Signal_AddressSkyChanged(quint64 addressSky) { callbackQAddressbaab1e_AddressSkyChanged(this, addressSky); };
	quint64 addressCoinHours() { return callbackQAddressbaab1e_AddressCoinHours(this); };
	void setAddressCoinHours(quint64 addressCoinHours) { callbackQAddressbaab1e_SetAddressCoinHours(this, addressCoinHours); };
	void Signal_AddressCoinHoursChanged(quint64 addressCoinHours) { callbackQAddressbaab1e_AddressCoinHoursChanged(this, addressCoinHours); };
	 ~QAddressbaab1e() { callbackQAddressbaab1e_DestroyQAddress(this); };
	void childEvent(QChildEvent * event) { callbackQAddressbaab1e_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAddressbaab1e_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAddressbaab1e_CustomEvent(this, event); };
	void deleteLater() { callbackQAddressbaab1e_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAddressbaab1e_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAddressbaab1e_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQAddressbaab1e_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAddressbaab1e_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAddressbaab1e_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAddressbaab1e_TimerEvent(this, event); };
	QString addressDefault() { return _address; };
	void setAddressDefault(QString p) { if (p != _address) { _address = p; addressChanged(_address); } };
	quint64 addressSkyDefault() { return _addressSky; };
	void setAddressSkyDefault(quint64 p) { if (p != _addressSky) { _addressSky = p; addressSkyChanged(_addressSky); } };
	quint64 addressCoinHoursDefault() { return _addressCoinHours; };
	void setAddressCoinHoursDefault(quint64 p) { if (p != _addressCoinHours) { _addressCoinHours = p; addressCoinHoursChanged(_addressCoinHours); } };
signals:
	void addressChanged(QString address);
	void addressSkyChanged(quint64 addressSky);
	void addressCoinHoursChanged(quint64 addressCoinHours);
public slots:
private:
	QString _address;
	quint64 _addressSky;
	quint64 _addressCoinHours;
};

Q_DECLARE_METATYPE(QAddressbaab1e*)


void QAddressbaab1e_QAddressbaab1e_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class QWalletbaab1e: public QObject
{
Q_OBJECT
Q_PROPERTY(QString name READ name WRITE setName NOTIFY nameChanged)
Q_PROPERTY(qint32 encryptionEnabled READ encryptionEnabled WRITE setEncryptionEnabled NOTIFY encryptionEnabledChanged)
Q_PROPERTY(qint32 sky READ sky WRITE setSky NOTIFY skyChanged)
Q_PROPERTY(qint32 coinHours READ coinHours WRITE setCoinHours NOTIFY coinHoursChanged)
Q_PROPERTY(QString fileName READ fileName WRITE setFileName NOTIFY fileNameChanged)
public:
	QWalletbaab1e(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QWalletbaab1e_QWalletbaab1e_QRegisterMetaType();QWalletbaab1e_QWalletbaab1e_QRegisterMetaTypes();callbackQWalletbaab1e_Constructor(this);};
	QString name() { return ({ Moc_PackedString tempVal = callbackQWalletbaab1e_Name(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setName(QString name) { QByteArray t6ae999 = name.toUtf8(); Moc_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQWalletbaab1e_SetName(this, namePacked); };
	void Signal_NameChanged(QString name) { QByteArray t6ae999 = name.toUtf8(); Moc_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQWalletbaab1e_NameChanged(this, namePacked); };
	qint32 encryptionEnabled() { return callbackQWalletbaab1e_EncryptionEnabled(this); };
	void setEncryptionEnabled(qint32 encryptionEnabled) { callbackQWalletbaab1e_SetEncryptionEnabled(this, encryptionEnabled); };
	void Signal_EncryptionEnabledChanged(qint32 encryptionEnabled) { callbackQWalletbaab1e_EncryptionEnabledChanged(this, encryptionEnabled); };
	qint32 sky() { return callbackQWalletbaab1e_Sky(this); };
	void setSky(qint32 sky) { callbackQWalletbaab1e_SetSky(this, sky); };
	void Signal_SkyChanged(qint32 sky) { callbackQWalletbaab1e_SkyChanged(this, sky); };
	qint32 coinHours() { return callbackQWalletbaab1e_CoinHours(this); };
	void setCoinHours(qint32 coinHours) { callbackQWalletbaab1e_SetCoinHours(this, coinHours); };
	void Signal_CoinHoursChanged(qint32 coinHours) { callbackQWalletbaab1e_CoinHoursChanged(this, coinHours); };
	QString fileName() { return ({ Moc_PackedString tempVal = callbackQWalletbaab1e_FileName(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setFileName(QString fileName) { QByteArray td83e09 = fileName.toUtf8(); Moc_PackedString fileNamePacked = { const_cast<char*>(td83e09.prepend("WHITESPACE").constData()+10), td83e09.size()-10 };callbackQWalletbaab1e_SetFileName(this, fileNamePacked); };
	void Signal_FileNameChanged(QString fileName) { QByteArray td83e09 = fileName.toUtf8(); Moc_PackedString fileNamePacked = { const_cast<char*>(td83e09.prepend("WHITESPACE").constData()+10), td83e09.size()-10 };callbackQWalletbaab1e_FileNameChanged(this, fileNamePacked); };
	 ~QWalletbaab1e() { callbackQWalletbaab1e_DestroyQWallet(this); };
	void childEvent(QChildEvent * event) { callbackQWalletbaab1e_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWalletbaab1e_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWalletbaab1e_CustomEvent(this, event); };
	void deleteLater() { callbackQWalletbaab1e_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWalletbaab1e_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWalletbaab1e_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQWalletbaab1e_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWalletbaab1e_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWalletbaab1e_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWalletbaab1e_TimerEvent(this, event); };
	QString nameDefault() { return _name; };
	void setNameDefault(QString p) { if (p != _name) { _name = p; nameChanged(_name); } };
	qint32 encryptionEnabledDefault() { return _encryptionEnabled; };
	void setEncryptionEnabledDefault(qint32 p) { if (p != _encryptionEnabled) { _encryptionEnabled = p; encryptionEnabledChanged(_encryptionEnabled); } };
	qint32 skyDefault() { return _sky; };
	void setSkyDefault(qint32 p) { if (p != _sky) { _sky = p; skyChanged(_sky); } };
	qint32 coinHoursDefault() { return _coinHours; };
	void setCoinHoursDefault(qint32 p) { if (p != _coinHours) { _coinHours = p; coinHoursChanged(_coinHours); } };
	QString fileNameDefault() { return _fileName; };
	void setFileNameDefault(QString p) { if (p != _fileName) { _fileName = p; fileNameChanged(_fileName); } };
signals:
	void nameChanged(QString name);
	void encryptionEnabledChanged(qint32 encryptionEnabled);
	void skyChanged(qint32 sky);
	void coinHoursChanged(qint32 coinHours);
	void fileNameChanged(QString fileName);
public slots:
private:
	QString _name;
	qint32 _encryptionEnabled;
	qint32 _sky;
	qint32 _coinHours;
	QString _fileName;
};

Q_DECLARE_METATYPE(QWalletbaab1e*)


void QWalletbaab1e_QWalletbaab1e_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class WalletModelbaab1e: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<QWalletbaab1e*> wallets READ wallets WRITE setWallets NOTIFY walletsChanged)
public:
	WalletModelbaab1e(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");WalletModelbaab1e_WalletModelbaab1e_QRegisterMetaType();WalletModelbaab1e_WalletModelbaab1e_QRegisterMetaTypes();callbackWalletModelbaab1e_Constructor(this);};
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbackWalletModelbaab1e_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbackWalletModelbaab1e_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbackWalletModelbaab1e_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<QWalletbaab1e*> wallets() { return ({ QList<QWalletbaab1e*>* tmpP = static_cast<QList<QWalletbaab1e*>*>(callbackWalletModelbaab1e_Wallets(this)); QList<QWalletbaab1e*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setWallets(QList<QWalletbaab1e*> wallets) { callbackWalletModelbaab1e_SetWallets(this, ({ QList<QWalletbaab1e*>* tmpValue = new QList<QWalletbaab1e*>(wallets); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_WalletsChanged(QList<QWalletbaab1e*> wallets) { callbackWalletModelbaab1e_WalletsChanged(this, ({ QList<QWalletbaab1e*>* tmpValue = new QList<QWalletbaab1e*>(wallets); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	 ~WalletModelbaab1e() { callbackWalletModelbaab1e_DestroyWalletModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackWalletModelbaab1e_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackWalletModelbaab1e_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackWalletModelbaab1e_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackWalletModelbaab1e_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackWalletModelbaab1e_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackWalletModelbaab1e_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackWalletModelbaab1e_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackWalletModelbaab1e_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackWalletModelbaab1e_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackWalletModelbaab1e_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModelbaab1e_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackWalletModelbaab1e_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackWalletModelbaab1e_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModelbaab1e_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackWalletModelbaab1e_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackWalletModelbaab1e_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = new QVector<int>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackWalletModelbaab1e_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackWalletModelbaab1e_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackWalletModelbaab1e_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackWalletModelbaab1e_HeaderDataChanged(this, orientation, first, last); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackWalletModelbaab1e_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackWalletModelbaab1e_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackWalletModelbaab1e_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackWalletModelbaab1e_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackWalletModelbaab1e_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackWalletModelbaab1e_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackWalletModelbaab1e_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackWalletModelbaab1e_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackWalletModelbaab1e_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackWalletModelbaab1e_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackWalletModelbaab1e_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackWalletModelbaab1e_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackWalletModelbaab1e_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackWalletModelbaab1e_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackWalletModelbaab1e_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackWalletModelbaab1e_ResetInternalData(this); };
	void revert() { callbackWalletModelbaab1e_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackWalletModelbaab1e_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackWalletModelbaab1e_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackWalletModelbaab1e_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackWalletModelbaab1e_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModelbaab1e_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackWalletModelbaab1e_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackWalletModelbaab1e_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModelbaab1e_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackWalletModelbaab1e_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackWalletModelbaab1e_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackWalletModelbaab1e_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackWalletModelbaab1e_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackWalletModelbaab1e_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackWalletModelbaab1e_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackWalletModelbaab1e_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackWalletModelbaab1e_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackWalletModelbaab1e_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackWalletModelbaab1e_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackWalletModelbaab1e_CustomEvent(this, event); };
	void deleteLater() { callbackWalletModelbaab1e_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackWalletModelbaab1e_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackWalletModelbaab1e_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackWalletModelbaab1e_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackWalletModelbaab1e_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackWalletModelbaab1e_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackWalletModelbaab1e_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<QWalletbaab1e*> walletsDefault() { return _wallets; };
	void setWalletsDefault(QList<QWalletbaab1e*> p) { if (p != _wallets) { _wallets = p; walletsChanged(_wallets); } };
signals:
	void rolesChanged(type378cdd roles);
	void walletsChanged(QList<QWalletbaab1e*> wallets);
public slots:
	void addWallet(QWalletbaab1e* v0) { callbackWalletModelbaab1e_AddWallet(this, v0); };
	void editWallet(qint32 row, QString name, bool encryptionEnabled, qint32 sky, qint32 coinHours) { QByteArray t6ae999 = name.toUtf8(); Moc_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackWalletModelbaab1e_EditWallet(this, row, namePacked, encryptionEnabled, sky, coinHours); };
	void removeWallet(qint32 row) { callbackWalletModelbaab1e_RemoveWallet(this, row); };
	void loadModel() { callbackWalletModelbaab1e_LoadModel(this); };
private:
	type378cdd _roles;
	QList<QWalletbaab1e*> _wallets;
};

Q_DECLARE_METATYPE(WalletModelbaab1e*)


void WalletModelbaab1e_WalletModelbaab1e_QRegisterMetaTypes() {
	qRegisterMetaType<type378cdd>("type378cdd");
	qRegisterMetaType<QList<QObject*>>("QList<QWalletbaab1e*>");
}

void AddressesModelbaab1e_AddAddress(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<AddressesModelbaab1e*>(ptr), "addAddress", Q_ARG(QAddressbaab1e*, static_cast<QAddressbaab1e*>(v0)));
}

void AddressesModelbaab1e_RemoveAddress(void* ptr, int v0)
{
	QMetaObject::invokeMethod(static_cast<AddressesModelbaab1e*>(ptr), "removeAddress", Q_ARG(qint32, v0));
}

void AddressesModelbaab1e_EditAddress(void* ptr, int v0, struct Moc_PackedString v1, unsigned long long v2, unsigned long long v3)
{
	QMetaObject::invokeMethod(static_cast<AddressesModelbaab1e*>(ptr), "editAddress", Q_ARG(qint32, v0), Q_ARG(QString, QString::fromUtf8(v1.data, v1.len)), Q_ARG(quint64, v2), Q_ARG(quint64, v3));
}

void AddressesModelbaab1e_LoadModel(void* ptr, struct Moc_PackedString v0)
{
	QMetaObject::invokeMethod(static_cast<AddressesModelbaab1e*>(ptr), "loadModel", Q_ARG(QString, QString::fromUtf8(v0.data, v0.len)));
}

struct Moc_PackedList AddressesModelbaab1e_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<AddressesModelbaab1e*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressesModelbaab1e_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<AddressesModelbaab1e*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void AddressesModelbaab1e_SetRoles(void* ptr, void* roles)
{
	static_cast<AddressesModelbaab1e*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void AddressesModelbaab1e_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<AddressesModelbaab1e*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void AddressesModelbaab1e_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<AddressesModelbaab1e*>(ptr), static_cast<void (AddressesModelbaab1e::*)(QMap<qint32, QByteArray>)>(&AddressesModelbaab1e::rolesChanged), static_cast<AddressesModelbaab1e*>(ptr), static_cast<void (AddressesModelbaab1e::*)(QMap<qint32, QByteArray>)>(&AddressesModelbaab1e::Signal_RolesChanged));
}

void AddressesModelbaab1e_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressesModelbaab1e*>(ptr), static_cast<void (AddressesModelbaab1e::*)(QMap<qint32, QByteArray>)>(&AddressesModelbaab1e::rolesChanged), static_cast<AddressesModelbaab1e*>(ptr), static_cast<void (AddressesModelbaab1e::*)(QMap<qint32, QByteArray>)>(&AddressesModelbaab1e::Signal_RolesChanged));
}

void AddressesModelbaab1e_RolesChanged(void* ptr, void* roles)
{
	static_cast<AddressesModelbaab1e*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList AddressesModelbaab1e_Addresses(void* ptr)
{
	return ({ QList<QAddressbaab1e*>* tmpValue = new QList<QAddressbaab1e*>(static_cast<AddressesModelbaab1e*>(ptr)->addresses()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressesModelbaab1e_AddressesDefault(void* ptr)
{
	return ({ QList<QAddressbaab1e*>* tmpValue = new QList<QAddressbaab1e*>(static_cast<AddressesModelbaab1e*>(ptr)->addressesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void AddressesModelbaab1e_SetAddresses(void* ptr, void* addresses)
{
	static_cast<AddressesModelbaab1e*>(ptr)->setAddresses(({ QList<QAddressbaab1e*>* tmpP = static_cast<QList<QAddressbaab1e*>*>(addresses); QList<QAddressbaab1e*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void AddressesModelbaab1e_SetAddressesDefault(void* ptr, void* addresses)
{
	static_cast<AddressesModelbaab1e*>(ptr)->setAddressesDefault(({ QList<QAddressbaab1e*>* tmpP = static_cast<QList<QAddressbaab1e*>*>(addresses); QList<QAddressbaab1e*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void AddressesModelbaab1e_ConnectAddressesChanged(void* ptr)
{
	QObject::connect(static_cast<AddressesModelbaab1e*>(ptr), static_cast<void (AddressesModelbaab1e::*)(QList<QAddressbaab1e*>)>(&AddressesModelbaab1e::addressesChanged), static_cast<AddressesModelbaab1e*>(ptr), static_cast<void (AddressesModelbaab1e::*)(QList<QAddressbaab1e*>)>(&AddressesModelbaab1e::Signal_AddressesChanged));
}

void AddressesModelbaab1e_DisconnectAddressesChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressesModelbaab1e*>(ptr), static_cast<void (AddressesModelbaab1e::*)(QList<QAddressbaab1e*>)>(&AddressesModelbaab1e::addressesChanged), static_cast<AddressesModelbaab1e*>(ptr), static_cast<void (AddressesModelbaab1e::*)(QList<QAddressbaab1e*>)>(&AddressesModelbaab1e::Signal_AddressesChanged));
}

void AddressesModelbaab1e_AddressesChanged(void* ptr, void* addresses)
{
	static_cast<AddressesModelbaab1e*>(ptr)->addressesChanged(({ QList<QAddressbaab1e*>* tmpP = static_cast<QList<QAddressbaab1e*>*>(addresses); QList<QAddressbaab1e*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int AddressesModelbaab1e_Loaded(void* ptr)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->loaded();
}

int AddressesModelbaab1e_LoadedDefault(void* ptr)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->loadedDefault();
}

void AddressesModelbaab1e_SetLoaded(void* ptr, int loaded)
{
	static_cast<AddressesModelbaab1e*>(ptr)->setLoaded(loaded);
}

void AddressesModelbaab1e_SetLoadedDefault(void* ptr, int loaded)
{
	static_cast<AddressesModelbaab1e*>(ptr)->setLoadedDefault(loaded);
}

void AddressesModelbaab1e_ConnectLoadedChanged(void* ptr)
{
	QObject::connect(static_cast<AddressesModelbaab1e*>(ptr), static_cast<void (AddressesModelbaab1e::*)(qint32)>(&AddressesModelbaab1e::loadedChanged), static_cast<AddressesModelbaab1e*>(ptr), static_cast<void (AddressesModelbaab1e::*)(qint32)>(&AddressesModelbaab1e::Signal_LoadedChanged));
}

void AddressesModelbaab1e_DisconnectLoadedChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressesModelbaab1e*>(ptr), static_cast<void (AddressesModelbaab1e::*)(qint32)>(&AddressesModelbaab1e::loadedChanged), static_cast<AddressesModelbaab1e*>(ptr), static_cast<void (AddressesModelbaab1e::*)(qint32)>(&AddressesModelbaab1e::Signal_LoadedChanged));
}

void AddressesModelbaab1e_LoadedChanged(void* ptr, int loaded)
{
	static_cast<AddressesModelbaab1e*>(ptr)->loadedChanged(loaded);
}

int AddressesModelbaab1e_AddressesModelbaab1e_QRegisterMetaType()
{
	return qRegisterMetaType<AddressesModelbaab1e*>();
}

int AddressesModelbaab1e_AddressesModelbaab1e_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<AddressesModelbaab1e*>(const_cast<const char*>(typeName));
}

int AddressesModelbaab1e_AddressesModelbaab1e_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressesModelbaab1e>();
#else
	return 0;
#endif
}

int AddressesModelbaab1e_AddressesModelbaab1e_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressesModelbaab1e>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int AddressesModelbaab1e_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModelbaab1e_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModelbaab1e_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressesModelbaab1e_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModelbaab1e_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModelbaab1e_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressesModelbaab1e_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModelbaab1e_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModelbaab1e_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* AddressesModelbaab1e___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModelbaab1e___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModelbaab1e___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressesModelbaab1e___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModelbaab1e___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModelbaab1e___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int AddressesModelbaab1e___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void AddressesModelbaab1e___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* AddressesModelbaab1e___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* AddressesModelbaab1e___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModelbaab1e___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* AddressesModelbaab1e___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList AddressesModelbaab1e___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModelbaab1e___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModelbaab1e___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* AddressesModelbaab1e___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* AddressesModelbaab1e___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModelbaab1e___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* AddressesModelbaab1e___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* AddressesModelbaab1e___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModelbaab1e___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModelbaab1e___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressesModelbaab1e___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModelbaab1e___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModelbaab1e___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressesModelbaab1e___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModelbaab1e___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModelbaab1e___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressesModelbaab1e___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void AddressesModelbaab1e___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressesModelbaab1e___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList AddressesModelbaab1e___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModelbaab1e___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModelbaab1e___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* AddressesModelbaab1e___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList AddressesModelbaab1e___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int AddressesModelbaab1e_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModelbaab1e_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModelbaab1e_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressesModelbaab1e_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModelbaab1e_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModelbaab1e_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* AddressesModelbaab1e___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModelbaab1e___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressesModelbaab1e___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* AddressesModelbaab1e___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModelbaab1e___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* AddressesModelbaab1e___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* AddressesModelbaab1e___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModelbaab1e___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressesModelbaab1e___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressesModelbaab1e___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModelbaab1e___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressesModelbaab1e___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressesModelbaab1e___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModelbaab1e___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressesModelbaab1e___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressesModelbaab1e___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModelbaab1e___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressesModelbaab1e___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressesModelbaab1e___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModelbaab1e___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModelbaab1e___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressesModelbaab1e___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressesModelbaab1e___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModelbaab1e___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModelbaab1e___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressesModelbaab1e___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressesModelbaab1e___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModelbaab1e___addresses_atList(void* ptr, int i)
{
	return ({QAddressbaab1e* tmp = static_cast<QList<QAddressbaab1e*>*>(ptr)->at(i); if (i == static_cast<QList<QAddressbaab1e*>*>(ptr)->size()-1) { static_cast<QList<QAddressbaab1e*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModelbaab1e___addresses_setList(void* ptr, void* i)
{
	static_cast<QList<QAddressbaab1e*>*>(ptr)->append(static_cast<QAddressbaab1e*>(i));
}

void* AddressesModelbaab1e___addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAddressbaab1e*>();
}

void* AddressesModelbaab1e___setAddresses_addresses_atList(void* ptr, int i)
{
	return ({QAddressbaab1e* tmp = static_cast<QList<QAddressbaab1e*>*>(ptr)->at(i); if (i == static_cast<QList<QAddressbaab1e*>*>(ptr)->size()-1) { static_cast<QList<QAddressbaab1e*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModelbaab1e___setAddresses_addresses_setList(void* ptr, void* i)
{
	static_cast<QList<QAddressbaab1e*>*>(ptr)->append(static_cast<QAddressbaab1e*>(i));
}

void* AddressesModelbaab1e___setAddresses_addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAddressbaab1e*>();
}

void* AddressesModelbaab1e___addressesChanged_addresses_atList(void* ptr, int i)
{
	return ({QAddressbaab1e* tmp = static_cast<QList<QAddressbaab1e*>*>(ptr)->at(i); if (i == static_cast<QList<QAddressbaab1e*>*>(ptr)->size()-1) { static_cast<QList<QAddressbaab1e*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModelbaab1e___addressesChanged_addresses_setList(void* ptr, void* i)
{
	static_cast<QList<QAddressbaab1e*>*>(ptr)->append(static_cast<QAddressbaab1e*>(i));
}

void* AddressesModelbaab1e___addressesChanged_addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAddressbaab1e*>();
}

int AddressesModelbaab1e_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModelbaab1e_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressesModelbaab1e_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int AddressesModelbaab1e_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModelbaab1e_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressesModelbaab1e_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int AddressesModelbaab1e_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModelbaab1e_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressesModelbaab1e_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* AddressesModelbaab1e_NewAddressesModel(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new AddressesModelbaab1e(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new AddressesModelbaab1e(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new AddressesModelbaab1e(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new AddressesModelbaab1e(static_cast<QWindow*>(parent));
	} else {
		return new AddressesModelbaab1e(static_cast<QObject*>(parent));
	}
}

void AddressesModelbaab1e_DestroyAddressesModel(void* ptr)
{
	static_cast<AddressesModelbaab1e*>(ptr)->~AddressesModelbaab1e();
}

void AddressesModelbaab1e_DestroyAddressesModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char AddressesModelbaab1e_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

long long AddressesModelbaab1e_FlagsDefault(void* ptr, void* index)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

void* AddressesModelbaab1e_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* AddressesModelbaab1e_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* AddressesModelbaab1e_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

char AddressesModelbaab1e_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char AddressesModelbaab1e_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int AddressesModelbaab1e_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* AddressesModelbaab1e_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void AddressesModelbaab1e_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char AddressesModelbaab1e_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* AddressesModelbaab1e_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

char AddressesModelbaab1e_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char AddressesModelbaab1e_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct Moc_PackedList AddressesModelbaab1e_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressesModelbaab1e_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModelbaab1e_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct Moc_PackedString AddressesModelbaab1e_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t6a4f86 = static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::mimeTypes().join("¡¦!").toUtf8(); Moc_PackedString { const_cast<char*>(t6a4f86.prepend("WHITESPACE").constData()+10), t6a4f86.size()-10 }; });
}

char AddressesModelbaab1e_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char AddressesModelbaab1e_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* AddressesModelbaab1e_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

char AddressesModelbaab1e_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char AddressesModelbaab1e_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void AddressesModelbaab1e_ResetInternalDataDefault(void* ptr)
{
	static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::resetInternalData();
}

void AddressesModelbaab1e_RevertDefault(void* ptr)
{
	static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::revert();
}

struct Moc_PackedList AddressesModelbaab1e_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int AddressesModelbaab1e_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char AddressesModelbaab1e_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char AddressesModelbaab1e_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char AddressesModelbaab1e_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void AddressesModelbaab1e_SortDefault(void* ptr, int column, long long order)
{
	static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* AddressesModelbaab1e_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char AddressesModelbaab1e_SubmitDefault(void* ptr)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::submit();
}

long long AddressesModelbaab1e_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long AddressesModelbaab1e_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::supportedDropActions();
}

void AddressesModelbaab1e_ChildEventDefault(void* ptr, void* event)
{
	static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void AddressesModelbaab1e_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void AddressesModelbaab1e_CustomEventDefault(void* ptr, void* event)
{
	static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void AddressesModelbaab1e_DeleteLaterDefault(void* ptr)
{
	static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::deleteLater();
}

void AddressesModelbaab1e_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char AddressesModelbaab1e_EventDefault(void* ptr, void* e)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char AddressesModelbaab1e_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void AddressesModelbaab1e_TimerEventDefault(void* ptr, void* event)
{
	static_cast<AddressesModelbaab1e*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

struct Moc_PackedString QAddressbaab1e_Address(void* ptr)
{
	return ({ QByteArray tbbcec3 = static_cast<QAddressbaab1e*>(ptr)->address().toUtf8(); Moc_PackedString { const_cast<char*>(tbbcec3.prepend("WHITESPACE").constData()+10), tbbcec3.size()-10 }; });
}

struct Moc_PackedString QAddressbaab1e_AddressDefault(void* ptr)
{
	return ({ QByteArray tbe1549 = static_cast<QAddressbaab1e*>(ptr)->addressDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tbe1549.prepend("WHITESPACE").constData()+10), tbe1549.size()-10 }; });
}

void QAddressbaab1e_SetAddress(void* ptr, struct Moc_PackedString address)
{
	static_cast<QAddressbaab1e*>(ptr)->setAddress(QString::fromUtf8(address.data, address.len));
}

void QAddressbaab1e_SetAddressDefault(void* ptr, struct Moc_PackedString address)
{
	static_cast<QAddressbaab1e*>(ptr)->setAddressDefault(QString::fromUtf8(address.data, address.len));
}

void QAddressbaab1e_ConnectAddressChanged(void* ptr)
{
	QObject::connect(static_cast<QAddressbaab1e*>(ptr), static_cast<void (QAddressbaab1e::*)(QString)>(&QAddressbaab1e::addressChanged), static_cast<QAddressbaab1e*>(ptr), static_cast<void (QAddressbaab1e::*)(QString)>(&QAddressbaab1e::Signal_AddressChanged));
}

void QAddressbaab1e_DisconnectAddressChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAddressbaab1e*>(ptr), static_cast<void (QAddressbaab1e::*)(QString)>(&QAddressbaab1e::addressChanged), static_cast<QAddressbaab1e*>(ptr), static_cast<void (QAddressbaab1e::*)(QString)>(&QAddressbaab1e::Signal_AddressChanged));
}

void QAddressbaab1e_AddressChanged(void* ptr, struct Moc_PackedString address)
{
	static_cast<QAddressbaab1e*>(ptr)->addressChanged(QString::fromUtf8(address.data, address.len));
}

unsigned long long QAddressbaab1e_AddressSky(void* ptr)
{
	return static_cast<QAddressbaab1e*>(ptr)->addressSky();
}

unsigned long long QAddressbaab1e_AddressSkyDefault(void* ptr)
{
	return static_cast<QAddressbaab1e*>(ptr)->addressSkyDefault();
}

void QAddressbaab1e_SetAddressSky(void* ptr, unsigned long long addressSky)
{
	static_cast<QAddressbaab1e*>(ptr)->setAddressSky(addressSky);
}

void QAddressbaab1e_SetAddressSkyDefault(void* ptr, unsigned long long addressSky)
{
	static_cast<QAddressbaab1e*>(ptr)->setAddressSkyDefault(addressSky);
}

void QAddressbaab1e_ConnectAddressSkyChanged(void* ptr)
{
	QObject::connect(static_cast<QAddressbaab1e*>(ptr), static_cast<void (QAddressbaab1e::*)(quint64)>(&QAddressbaab1e::addressSkyChanged), static_cast<QAddressbaab1e*>(ptr), static_cast<void (QAddressbaab1e::*)(quint64)>(&QAddressbaab1e::Signal_AddressSkyChanged));
}

void QAddressbaab1e_DisconnectAddressSkyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAddressbaab1e*>(ptr), static_cast<void (QAddressbaab1e::*)(quint64)>(&QAddressbaab1e::addressSkyChanged), static_cast<QAddressbaab1e*>(ptr), static_cast<void (QAddressbaab1e::*)(quint64)>(&QAddressbaab1e::Signal_AddressSkyChanged));
}

void QAddressbaab1e_AddressSkyChanged(void* ptr, unsigned long long addressSky)
{
	static_cast<QAddressbaab1e*>(ptr)->addressSkyChanged(addressSky);
}

unsigned long long QAddressbaab1e_AddressCoinHours(void* ptr)
{
	return static_cast<QAddressbaab1e*>(ptr)->addressCoinHours();
}

unsigned long long QAddressbaab1e_AddressCoinHoursDefault(void* ptr)
{
	return static_cast<QAddressbaab1e*>(ptr)->addressCoinHoursDefault();
}

void QAddressbaab1e_SetAddressCoinHours(void* ptr, unsigned long long addressCoinHours)
{
	static_cast<QAddressbaab1e*>(ptr)->setAddressCoinHours(addressCoinHours);
}

void QAddressbaab1e_SetAddressCoinHoursDefault(void* ptr, unsigned long long addressCoinHours)
{
	static_cast<QAddressbaab1e*>(ptr)->setAddressCoinHoursDefault(addressCoinHours);
}

void QAddressbaab1e_ConnectAddressCoinHoursChanged(void* ptr)
{
	QObject::connect(static_cast<QAddressbaab1e*>(ptr), static_cast<void (QAddressbaab1e::*)(quint64)>(&QAddressbaab1e::addressCoinHoursChanged), static_cast<QAddressbaab1e*>(ptr), static_cast<void (QAddressbaab1e::*)(quint64)>(&QAddressbaab1e::Signal_AddressCoinHoursChanged));
}

void QAddressbaab1e_DisconnectAddressCoinHoursChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAddressbaab1e*>(ptr), static_cast<void (QAddressbaab1e::*)(quint64)>(&QAddressbaab1e::addressCoinHoursChanged), static_cast<QAddressbaab1e*>(ptr), static_cast<void (QAddressbaab1e::*)(quint64)>(&QAddressbaab1e::Signal_AddressCoinHoursChanged));
}

void QAddressbaab1e_AddressCoinHoursChanged(void* ptr, unsigned long long addressCoinHours)
{
	static_cast<QAddressbaab1e*>(ptr)->addressCoinHoursChanged(addressCoinHours);
}

int QAddressbaab1e_QAddressbaab1e_QRegisterMetaType()
{
	return qRegisterMetaType<QAddressbaab1e*>();
}

int QAddressbaab1e_QAddressbaab1e_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QAddressbaab1e*>(const_cast<const char*>(typeName));
}

int QAddressbaab1e_QAddressbaab1e_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QAddressbaab1e>();
#else
	return 0;
#endif
}

int QAddressbaab1e_QAddressbaab1e_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QAddressbaab1e>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QAddressbaab1e___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAddressbaab1e___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAddressbaab1e___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QAddressbaab1e___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAddressbaab1e___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QAddressbaab1e___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QAddressbaab1e___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAddressbaab1e___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAddressbaab1e___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAddressbaab1e___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAddressbaab1e___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAddressbaab1e___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAddressbaab1e___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAddressbaab1e___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAddressbaab1e___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAddressbaab1e_NewQAddress(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QAddressbaab1e(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QAddressbaab1e(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QAddressbaab1e(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QAddressbaab1e(static_cast<QWindow*>(parent));
	} else {
		return new QAddressbaab1e(static_cast<QObject*>(parent));
	}
}

void QAddressbaab1e_DestroyQAddress(void* ptr)
{
	static_cast<QAddressbaab1e*>(ptr)->~QAddressbaab1e();
}

void QAddressbaab1e_DestroyQAddressDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void QAddressbaab1e_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAddressbaab1e*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QAddressbaab1e_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAddressbaab1e*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAddressbaab1e_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAddressbaab1e*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QAddressbaab1e_DeleteLaterDefault(void* ptr)
{
	static_cast<QAddressbaab1e*>(ptr)->QObject::deleteLater();
}

void QAddressbaab1e_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAddressbaab1e*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QAddressbaab1e_EventDefault(void* ptr, void* e)
{
	return static_cast<QAddressbaab1e*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QAddressbaab1e_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAddressbaab1e*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QAddressbaab1e_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAddressbaab1e*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

struct Moc_PackedString QWalletbaab1e_Name(void* ptr)
{
	return ({ QByteArray tedc734 = static_cast<QWalletbaab1e*>(ptr)->name().toUtf8(); Moc_PackedString { const_cast<char*>(tedc734.prepend("WHITESPACE").constData()+10), tedc734.size()-10 }; });
}

struct Moc_PackedString QWalletbaab1e_NameDefault(void* ptr)
{
	return ({ QByteArray tee7a91 = static_cast<QWalletbaab1e*>(ptr)->nameDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tee7a91.prepend("WHITESPACE").constData()+10), tee7a91.size()-10 }; });
}

void QWalletbaab1e_SetName(void* ptr, struct Moc_PackedString name)
{
	static_cast<QWalletbaab1e*>(ptr)->setName(QString::fromUtf8(name.data, name.len));
}

void QWalletbaab1e_SetNameDefault(void* ptr, struct Moc_PackedString name)
{
	static_cast<QWalletbaab1e*>(ptr)->setNameDefault(QString::fromUtf8(name.data, name.len));
}

void QWalletbaab1e_ConnectNameChanged(void* ptr)
{
	QObject::connect(static_cast<QWalletbaab1e*>(ptr), static_cast<void (QWalletbaab1e::*)(QString)>(&QWalletbaab1e::nameChanged), static_cast<QWalletbaab1e*>(ptr), static_cast<void (QWalletbaab1e::*)(QString)>(&QWalletbaab1e::Signal_NameChanged));
}

void QWalletbaab1e_DisconnectNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWalletbaab1e*>(ptr), static_cast<void (QWalletbaab1e::*)(QString)>(&QWalletbaab1e::nameChanged), static_cast<QWalletbaab1e*>(ptr), static_cast<void (QWalletbaab1e::*)(QString)>(&QWalletbaab1e::Signal_NameChanged));
}

void QWalletbaab1e_NameChanged(void* ptr, struct Moc_PackedString name)
{
	static_cast<QWalletbaab1e*>(ptr)->nameChanged(QString::fromUtf8(name.data, name.len));
}

int QWalletbaab1e_EncryptionEnabled(void* ptr)
{
	return static_cast<QWalletbaab1e*>(ptr)->encryptionEnabled();
}

int QWalletbaab1e_EncryptionEnabledDefault(void* ptr)
{
	return static_cast<QWalletbaab1e*>(ptr)->encryptionEnabledDefault();
}

void QWalletbaab1e_SetEncryptionEnabled(void* ptr, int encryptionEnabled)
{
	static_cast<QWalletbaab1e*>(ptr)->setEncryptionEnabled(encryptionEnabled);
}

void QWalletbaab1e_SetEncryptionEnabledDefault(void* ptr, int encryptionEnabled)
{
	static_cast<QWalletbaab1e*>(ptr)->setEncryptionEnabledDefault(encryptionEnabled);
}

void QWalletbaab1e_ConnectEncryptionEnabledChanged(void* ptr)
{
	QObject::connect(static_cast<QWalletbaab1e*>(ptr), static_cast<void (QWalletbaab1e::*)(qint32)>(&QWalletbaab1e::encryptionEnabledChanged), static_cast<QWalletbaab1e*>(ptr), static_cast<void (QWalletbaab1e::*)(qint32)>(&QWalletbaab1e::Signal_EncryptionEnabledChanged));
}

void QWalletbaab1e_DisconnectEncryptionEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWalletbaab1e*>(ptr), static_cast<void (QWalletbaab1e::*)(qint32)>(&QWalletbaab1e::encryptionEnabledChanged), static_cast<QWalletbaab1e*>(ptr), static_cast<void (QWalletbaab1e::*)(qint32)>(&QWalletbaab1e::Signal_EncryptionEnabledChanged));
}

void QWalletbaab1e_EncryptionEnabledChanged(void* ptr, int encryptionEnabled)
{
	static_cast<QWalletbaab1e*>(ptr)->encryptionEnabledChanged(encryptionEnabled);
}

int QWalletbaab1e_Sky(void* ptr)
{
	return static_cast<QWalletbaab1e*>(ptr)->sky();
}

int QWalletbaab1e_SkyDefault(void* ptr)
{
	return static_cast<QWalletbaab1e*>(ptr)->skyDefault();
}

void QWalletbaab1e_SetSky(void* ptr, int sky)
{
	static_cast<QWalletbaab1e*>(ptr)->setSky(sky);
}

void QWalletbaab1e_SetSkyDefault(void* ptr, int sky)
{
	static_cast<QWalletbaab1e*>(ptr)->setSkyDefault(sky);
}

void QWalletbaab1e_ConnectSkyChanged(void* ptr)
{
	QObject::connect(static_cast<QWalletbaab1e*>(ptr), static_cast<void (QWalletbaab1e::*)(qint32)>(&QWalletbaab1e::skyChanged), static_cast<QWalletbaab1e*>(ptr), static_cast<void (QWalletbaab1e::*)(qint32)>(&QWalletbaab1e::Signal_SkyChanged));
}

void QWalletbaab1e_DisconnectSkyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWalletbaab1e*>(ptr), static_cast<void (QWalletbaab1e::*)(qint32)>(&QWalletbaab1e::skyChanged), static_cast<QWalletbaab1e*>(ptr), static_cast<void (QWalletbaab1e::*)(qint32)>(&QWalletbaab1e::Signal_SkyChanged));
}

void QWalletbaab1e_SkyChanged(void* ptr, int sky)
{
	static_cast<QWalletbaab1e*>(ptr)->skyChanged(sky);
}

int QWalletbaab1e_CoinHours(void* ptr)
{
	return static_cast<QWalletbaab1e*>(ptr)->coinHours();
}

int QWalletbaab1e_CoinHoursDefault(void* ptr)
{
	return static_cast<QWalletbaab1e*>(ptr)->coinHoursDefault();
}

void QWalletbaab1e_SetCoinHours(void* ptr, int coinHours)
{
	static_cast<QWalletbaab1e*>(ptr)->setCoinHours(coinHours);
}

void QWalletbaab1e_SetCoinHoursDefault(void* ptr, int coinHours)
{
	static_cast<QWalletbaab1e*>(ptr)->setCoinHoursDefault(coinHours);
}

void QWalletbaab1e_ConnectCoinHoursChanged(void* ptr)
{
	QObject::connect(static_cast<QWalletbaab1e*>(ptr), static_cast<void (QWalletbaab1e::*)(qint32)>(&QWalletbaab1e::coinHoursChanged), static_cast<QWalletbaab1e*>(ptr), static_cast<void (QWalletbaab1e::*)(qint32)>(&QWalletbaab1e::Signal_CoinHoursChanged));
}

void QWalletbaab1e_DisconnectCoinHoursChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWalletbaab1e*>(ptr), static_cast<void (QWalletbaab1e::*)(qint32)>(&QWalletbaab1e::coinHoursChanged), static_cast<QWalletbaab1e*>(ptr), static_cast<void (QWalletbaab1e::*)(qint32)>(&QWalletbaab1e::Signal_CoinHoursChanged));
}

void QWalletbaab1e_CoinHoursChanged(void* ptr, int coinHours)
{
	static_cast<QWalletbaab1e*>(ptr)->coinHoursChanged(coinHours);
}

struct Moc_PackedString QWalletbaab1e_FileName(void* ptr)
{
	return ({ QByteArray teb23e1 = static_cast<QWalletbaab1e*>(ptr)->fileName().toUtf8(); Moc_PackedString { const_cast<char*>(teb23e1.prepend("WHITESPACE").constData()+10), teb23e1.size()-10 }; });
}

struct Moc_PackedString QWalletbaab1e_FileNameDefault(void* ptr)
{
	return ({ QByteArray t811a22 = static_cast<QWalletbaab1e*>(ptr)->fileNameDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t811a22.prepend("WHITESPACE").constData()+10), t811a22.size()-10 }; });
}

void QWalletbaab1e_SetFileName(void* ptr, struct Moc_PackedString fileName)
{
	static_cast<QWalletbaab1e*>(ptr)->setFileName(QString::fromUtf8(fileName.data, fileName.len));
}

void QWalletbaab1e_SetFileNameDefault(void* ptr, struct Moc_PackedString fileName)
{
	static_cast<QWalletbaab1e*>(ptr)->setFileNameDefault(QString::fromUtf8(fileName.data, fileName.len));
}

void QWalletbaab1e_ConnectFileNameChanged(void* ptr)
{
	QObject::connect(static_cast<QWalletbaab1e*>(ptr), static_cast<void (QWalletbaab1e::*)(QString)>(&QWalletbaab1e::fileNameChanged), static_cast<QWalletbaab1e*>(ptr), static_cast<void (QWalletbaab1e::*)(QString)>(&QWalletbaab1e::Signal_FileNameChanged));
}

void QWalletbaab1e_DisconnectFileNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWalletbaab1e*>(ptr), static_cast<void (QWalletbaab1e::*)(QString)>(&QWalletbaab1e::fileNameChanged), static_cast<QWalletbaab1e*>(ptr), static_cast<void (QWalletbaab1e::*)(QString)>(&QWalletbaab1e::Signal_FileNameChanged));
}

void QWalletbaab1e_FileNameChanged(void* ptr, struct Moc_PackedString fileName)
{
	static_cast<QWalletbaab1e*>(ptr)->fileNameChanged(QString::fromUtf8(fileName.data, fileName.len));
}

int QWalletbaab1e_QWalletbaab1e_QRegisterMetaType()
{
	return qRegisterMetaType<QWalletbaab1e*>();
}

int QWalletbaab1e_QWalletbaab1e_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QWalletbaab1e*>(const_cast<const char*>(typeName));
}

int QWalletbaab1e_QWalletbaab1e_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QWalletbaab1e>();
#else
	return 0;
#endif
}

int QWalletbaab1e_QWalletbaab1e_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QWalletbaab1e>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QWalletbaab1e___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWalletbaab1e___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWalletbaab1e___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QWalletbaab1e___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWalletbaab1e___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWalletbaab1e___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QWalletbaab1e___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWalletbaab1e___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWalletbaab1e___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWalletbaab1e___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWalletbaab1e___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWalletbaab1e___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWalletbaab1e___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWalletbaab1e___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWalletbaab1e___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWalletbaab1e_NewQWallet(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QWalletbaab1e(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QWalletbaab1e(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QWalletbaab1e(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QWalletbaab1e(static_cast<QWindow*>(parent));
	} else {
		return new QWalletbaab1e(static_cast<QObject*>(parent));
	}
}

void QWalletbaab1e_DestroyQWallet(void* ptr)
{
	static_cast<QWalletbaab1e*>(ptr)->~QWalletbaab1e();
}

void QWalletbaab1e_DestroyQWalletDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void QWalletbaab1e_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QWalletbaab1e*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QWalletbaab1e_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWalletbaab1e*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWalletbaab1e_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QWalletbaab1e*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QWalletbaab1e_DeleteLaterDefault(void* ptr)
{
	static_cast<QWalletbaab1e*>(ptr)->QObject::deleteLater();
}

void QWalletbaab1e_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWalletbaab1e*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QWalletbaab1e_EventDefault(void* ptr, void* e)
{
	return static_cast<QWalletbaab1e*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QWalletbaab1e_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QWalletbaab1e*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWalletbaab1e_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QWalletbaab1e*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

void WalletModelbaab1e_AddWallet(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<WalletModelbaab1e*>(ptr), "addWallet", Q_ARG(QWalletbaab1e*, static_cast<QWalletbaab1e*>(v0)));
}

void WalletModelbaab1e_EditWallet(void* ptr, int row, struct Moc_PackedString name, char encryptionEnabled, int sky, int coinHours)
{
	QMetaObject::invokeMethod(static_cast<WalletModelbaab1e*>(ptr), "editWallet", Q_ARG(qint32, row), Q_ARG(QString, QString::fromUtf8(name.data, name.len)), Q_ARG(bool, encryptionEnabled != 0), Q_ARG(qint32, sky), Q_ARG(qint32, coinHours));
}

void WalletModelbaab1e_RemoveWallet(void* ptr, int row)
{
	QMetaObject::invokeMethod(static_cast<WalletModelbaab1e*>(ptr), "removeWallet", Q_ARG(qint32, row));
}

void WalletModelbaab1e_LoadModel(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<WalletModelbaab1e*>(ptr), "loadModel");
}

struct Moc_PackedList WalletModelbaab1e_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<WalletModelbaab1e*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList WalletModelbaab1e_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<WalletModelbaab1e*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void WalletModelbaab1e_SetRoles(void* ptr, void* roles)
{
	static_cast<WalletModelbaab1e*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void WalletModelbaab1e_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<WalletModelbaab1e*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void WalletModelbaab1e_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<WalletModelbaab1e*>(ptr), static_cast<void (WalletModelbaab1e::*)(QMap<qint32, QByteArray>)>(&WalletModelbaab1e::rolesChanged), static_cast<WalletModelbaab1e*>(ptr), static_cast<void (WalletModelbaab1e::*)(QMap<qint32, QByteArray>)>(&WalletModelbaab1e::Signal_RolesChanged));
}

void WalletModelbaab1e_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<WalletModelbaab1e*>(ptr), static_cast<void (WalletModelbaab1e::*)(QMap<qint32, QByteArray>)>(&WalletModelbaab1e::rolesChanged), static_cast<WalletModelbaab1e*>(ptr), static_cast<void (WalletModelbaab1e::*)(QMap<qint32, QByteArray>)>(&WalletModelbaab1e::Signal_RolesChanged));
}

void WalletModelbaab1e_RolesChanged(void* ptr, void* roles)
{
	static_cast<WalletModelbaab1e*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList WalletModelbaab1e_Wallets(void* ptr)
{
	return ({ QList<QWalletbaab1e*>* tmpValue = new QList<QWalletbaab1e*>(static_cast<WalletModelbaab1e*>(ptr)->wallets()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList WalletModelbaab1e_WalletsDefault(void* ptr)
{
	return ({ QList<QWalletbaab1e*>* tmpValue = new QList<QWalletbaab1e*>(static_cast<WalletModelbaab1e*>(ptr)->walletsDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void WalletModelbaab1e_SetWallets(void* ptr, void* wallets)
{
	static_cast<WalletModelbaab1e*>(ptr)->setWallets(({ QList<QWalletbaab1e*>* tmpP = static_cast<QList<QWalletbaab1e*>*>(wallets); QList<QWalletbaab1e*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void WalletModelbaab1e_SetWalletsDefault(void* ptr, void* wallets)
{
	static_cast<WalletModelbaab1e*>(ptr)->setWalletsDefault(({ QList<QWalletbaab1e*>* tmpP = static_cast<QList<QWalletbaab1e*>*>(wallets); QList<QWalletbaab1e*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void WalletModelbaab1e_ConnectWalletsChanged(void* ptr)
{
	QObject::connect(static_cast<WalletModelbaab1e*>(ptr), static_cast<void (WalletModelbaab1e::*)(QList<QWalletbaab1e*>)>(&WalletModelbaab1e::walletsChanged), static_cast<WalletModelbaab1e*>(ptr), static_cast<void (WalletModelbaab1e::*)(QList<QWalletbaab1e*>)>(&WalletModelbaab1e::Signal_WalletsChanged));
}

void WalletModelbaab1e_DisconnectWalletsChanged(void* ptr)
{
	QObject::disconnect(static_cast<WalletModelbaab1e*>(ptr), static_cast<void (WalletModelbaab1e::*)(QList<QWalletbaab1e*>)>(&WalletModelbaab1e::walletsChanged), static_cast<WalletModelbaab1e*>(ptr), static_cast<void (WalletModelbaab1e::*)(QList<QWalletbaab1e*>)>(&WalletModelbaab1e::Signal_WalletsChanged));
}

void WalletModelbaab1e_WalletsChanged(void* ptr, void* wallets)
{
	static_cast<WalletModelbaab1e*>(ptr)->walletsChanged(({ QList<QWalletbaab1e*>* tmpP = static_cast<QList<QWalletbaab1e*>*>(wallets); QList<QWalletbaab1e*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int WalletModelbaab1e_WalletModelbaab1e_QRegisterMetaType()
{
	return qRegisterMetaType<WalletModelbaab1e*>();
}

int WalletModelbaab1e_WalletModelbaab1e_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<WalletModelbaab1e*>(const_cast<const char*>(typeName));
}

int WalletModelbaab1e_WalletModelbaab1e_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletModelbaab1e>();
#else
	return 0;
#endif
}

int WalletModelbaab1e_WalletModelbaab1e_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletModelbaab1e>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int WalletModelbaab1e_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModelbaab1e_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModelbaab1e_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int WalletModelbaab1e_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModelbaab1e_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModelbaab1e_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int WalletModelbaab1e_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModelbaab1e_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModelbaab1e_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* WalletModelbaab1e___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModelbaab1e___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModelbaab1e___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModelbaab1e___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModelbaab1e___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModelbaab1e___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int WalletModelbaab1e___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void WalletModelbaab1e___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* WalletModelbaab1e___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* WalletModelbaab1e___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModelbaab1e___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* WalletModelbaab1e___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList WalletModelbaab1e___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModelbaab1e___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModelbaab1e___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* WalletModelbaab1e___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* WalletModelbaab1e___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModelbaab1e___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* WalletModelbaab1e___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* WalletModelbaab1e___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModelbaab1e___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModelbaab1e___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModelbaab1e___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModelbaab1e___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModelbaab1e___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModelbaab1e___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModelbaab1e___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModelbaab1e___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModelbaab1e___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void WalletModelbaab1e___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModelbaab1e___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList WalletModelbaab1e___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModelbaab1e___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModelbaab1e___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* WalletModelbaab1e___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList WalletModelbaab1e___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int WalletModelbaab1e_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModelbaab1e_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModelbaab1e_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int WalletModelbaab1e_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModelbaab1e_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModelbaab1e_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* WalletModelbaab1e___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModelbaab1e___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModelbaab1e___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* WalletModelbaab1e___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModelbaab1e___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* WalletModelbaab1e___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* WalletModelbaab1e___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModelbaab1e___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModelbaab1e___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletModelbaab1e___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModelbaab1e___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModelbaab1e___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletModelbaab1e___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModelbaab1e___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModelbaab1e___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletModelbaab1e___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModelbaab1e___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModelbaab1e___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList WalletModelbaab1e___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModelbaab1e___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModelbaab1e___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModelbaab1e___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList WalletModelbaab1e___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModelbaab1e___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModelbaab1e___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModelbaab1e___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList WalletModelbaab1e___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModelbaab1e___wallets_atList(void* ptr, int i)
{
	return ({QWalletbaab1e* tmp = static_cast<QList<QWalletbaab1e*>*>(ptr)->at(i); if (i == static_cast<QList<QWalletbaab1e*>*>(ptr)->size()-1) { static_cast<QList<QWalletbaab1e*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModelbaab1e___wallets_setList(void* ptr, void* i)
{
	static_cast<QList<QWalletbaab1e*>*>(ptr)->append(static_cast<QWalletbaab1e*>(i));
}

void* WalletModelbaab1e___wallets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWalletbaab1e*>();
}

void* WalletModelbaab1e___setWallets_wallets_atList(void* ptr, int i)
{
	return ({QWalletbaab1e* tmp = static_cast<QList<QWalletbaab1e*>*>(ptr)->at(i); if (i == static_cast<QList<QWalletbaab1e*>*>(ptr)->size()-1) { static_cast<QList<QWalletbaab1e*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModelbaab1e___setWallets_wallets_setList(void* ptr, void* i)
{
	static_cast<QList<QWalletbaab1e*>*>(ptr)->append(static_cast<QWalletbaab1e*>(i));
}

void* WalletModelbaab1e___setWallets_wallets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWalletbaab1e*>();
}

void* WalletModelbaab1e___walletsChanged_wallets_atList(void* ptr, int i)
{
	return ({QWalletbaab1e* tmp = static_cast<QList<QWalletbaab1e*>*>(ptr)->at(i); if (i == static_cast<QList<QWalletbaab1e*>*>(ptr)->size()-1) { static_cast<QList<QWalletbaab1e*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModelbaab1e___walletsChanged_wallets_setList(void* ptr, void* i)
{
	static_cast<QList<QWalletbaab1e*>*>(ptr)->append(static_cast<QWalletbaab1e*>(i));
}

void* WalletModelbaab1e___walletsChanged_wallets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWalletbaab1e*>();
}

int WalletModelbaab1e_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModelbaab1e_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* WalletModelbaab1e_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int WalletModelbaab1e_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModelbaab1e_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* WalletModelbaab1e_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int WalletModelbaab1e_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModelbaab1e_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* WalletModelbaab1e_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* WalletModelbaab1e_NewWalletModel(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new WalletModelbaab1e(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new WalletModelbaab1e(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new WalletModelbaab1e(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new WalletModelbaab1e(static_cast<QWindow*>(parent));
	} else {
		return new WalletModelbaab1e(static_cast<QObject*>(parent));
	}
}

void WalletModelbaab1e_DestroyWalletModel(void* ptr)
{
	static_cast<WalletModelbaab1e*>(ptr)->~WalletModelbaab1e();
}

void WalletModelbaab1e_DestroyWalletModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char WalletModelbaab1e_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

long long WalletModelbaab1e_FlagsDefault(void* ptr, void* index)
{
	return static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

void* WalletModelbaab1e_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* WalletModelbaab1e_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* WalletModelbaab1e_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

char WalletModelbaab1e_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char WalletModelbaab1e_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int WalletModelbaab1e_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* WalletModelbaab1e_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void WalletModelbaab1e_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char WalletModelbaab1e_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* WalletModelbaab1e_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

char WalletModelbaab1e_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char WalletModelbaab1e_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct Moc_PackedList WalletModelbaab1e_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList WalletModelbaab1e_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModelbaab1e_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct Moc_PackedString WalletModelbaab1e_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t2c50d2 = static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::mimeTypes().join("¡¦!").toUtf8(); Moc_PackedString { const_cast<char*>(t2c50d2.prepend("WHITESPACE").constData()+10), t2c50d2.size()-10 }; });
}

char WalletModelbaab1e_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char WalletModelbaab1e_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* WalletModelbaab1e_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

char WalletModelbaab1e_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char WalletModelbaab1e_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void WalletModelbaab1e_ResetInternalDataDefault(void* ptr)
{
	static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::resetInternalData();
}

void WalletModelbaab1e_RevertDefault(void* ptr)
{
	static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::revert();
}

struct Moc_PackedList WalletModelbaab1e_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int WalletModelbaab1e_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char WalletModelbaab1e_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char WalletModelbaab1e_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char WalletModelbaab1e_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void WalletModelbaab1e_SortDefault(void* ptr, int column, long long order)
{
	static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* WalletModelbaab1e_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char WalletModelbaab1e_SubmitDefault(void* ptr)
{
	return static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::submit();
}

long long WalletModelbaab1e_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long WalletModelbaab1e_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::supportedDropActions();
}

void WalletModelbaab1e_ChildEventDefault(void* ptr, void* event)
{
	static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void WalletModelbaab1e_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void WalletModelbaab1e_CustomEventDefault(void* ptr, void* event)
{
	static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void WalletModelbaab1e_DeleteLaterDefault(void* ptr)
{
	static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::deleteLater();
}

void WalletModelbaab1e_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char WalletModelbaab1e_EventDefault(void* ptr, void* e)
{
	return static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char WalletModelbaab1e_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void WalletModelbaab1e_TimerEventDefault(void* ptr, void* event)
{
	static_cast<WalletModelbaab1e*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
