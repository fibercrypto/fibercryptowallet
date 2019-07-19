

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
class AddressesModel611006: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<QAddress611006*> addresses READ addresses WRITE setAddresses NOTIFY addressesChanged)
Q_PROPERTY(qint32 loaded READ loaded WRITE setLoaded NOTIFY loadedChanged)
public:
	AddressesModel611006(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");AddressesModel611006_AddressesModel611006_QRegisterMetaType();AddressesModel611006_AddressesModel611006_QRegisterMetaTypes();callbackAddressesModel611006_Constructor(this);};
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbackAddressesModel611006_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbackAddressesModel611006_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbackAddressesModel611006_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<QAddress611006*> addresses() { return ({ QList<QAddress611006*>* tmpP = static_cast<QList<QAddress611006*>*>(callbackAddressesModel611006_Addresses(this)); QList<QAddress611006*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setAddresses(QList<QAddress611006*> addresses) { callbackAddressesModel611006_SetAddresses(this, ({ QList<QAddress611006*>* tmpValue = new QList<QAddress611006*>(addresses); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_AddressesChanged(QList<QAddress611006*> addresses) { callbackAddressesModel611006_AddressesChanged(this, ({ QList<QAddress611006*>* tmpValue = new QList<QAddress611006*>(addresses); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	qint32 loaded() { return callbackAddressesModel611006_Loaded(this); };
	void setLoaded(qint32 loaded) { callbackAddressesModel611006_SetLoaded(this, loaded); };
	void Signal_LoadedChanged(qint32 loaded) { callbackAddressesModel611006_LoadedChanged(this, loaded); };
	 ~AddressesModel611006() { callbackAddressesModel611006_DestroyAddressesModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackAddressesModel611006_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackAddressesModel611006_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackAddressesModel611006_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackAddressesModel611006_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackAddressesModel611006_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackAddressesModel611006_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackAddressesModel611006_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackAddressesModel611006_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackAddressesModel611006_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackAddressesModel611006_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackAddressesModel611006_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackAddressesModel611006_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackAddressesModel611006_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackAddressesModel611006_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackAddressesModel611006_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackAddressesModel611006_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = new QVector<int>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackAddressesModel611006_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackAddressesModel611006_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackAddressesModel611006_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackAddressesModel611006_HeaderDataChanged(this, orientation, first, last); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackAddressesModel611006_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackAddressesModel611006_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackAddressesModel611006_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackAddressesModel611006_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackAddressesModel611006_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackAddressesModel611006_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackAddressesModel611006_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackAddressesModel611006_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackAddressesModel611006_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackAddressesModel611006_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackAddressesModel611006_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackAddressesModel611006_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackAddressesModel611006_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackAddressesModel611006_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackAddressesModel611006_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackAddressesModel611006_ResetInternalData(this); };
	void revert() { callbackAddressesModel611006_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackAddressesModel611006_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackAddressesModel611006_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackAddressesModel611006_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackAddressesModel611006_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackAddressesModel611006_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackAddressesModel611006_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackAddressesModel611006_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackAddressesModel611006_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackAddressesModel611006_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackAddressesModel611006_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackAddressesModel611006_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackAddressesModel611006_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackAddressesModel611006_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackAddressesModel611006_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackAddressesModel611006_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackAddressesModel611006_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackAddressesModel611006_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackAddressesModel611006_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackAddressesModel611006_CustomEvent(this, event); };
	void deleteLater() { callbackAddressesModel611006_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackAddressesModel611006_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackAddressesModel611006_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackAddressesModel611006_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackAddressesModel611006_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackAddressesModel611006_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackAddressesModel611006_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<QAddress611006*> addressesDefault() { return _addresses; };
	void setAddressesDefault(QList<QAddress611006*> p) { if (p != _addresses) { _addresses = p; addressesChanged(_addresses); } };
	qint32 loadedDefault() { return _loaded; };
	void setLoadedDefault(qint32 p) { if (p != _loaded) { _loaded = p; loadedChanged(_loaded); } };
signals:
	void rolesChanged(type378cdd roles);
	void addressesChanged(QList<QAddress611006*> addresses);
	void loadedChanged(qint32 loaded);
public slots:
	void addAddress(QAddress611006* v0) { callbackAddressesModel611006_AddAddress(this, v0); };
	void removeAddress(qint32 v0) { callbackAddressesModel611006_RemoveAddress(this, v0); };
	void editAddress(qint32 v0, QString v1, qint32 v2, qint32 v3) { QByteArray t5a6df7 = v1.toUtf8(); Moc_PackedString v1Packed = { const_cast<char*>(t5a6df7.prepend("WHITESPACE").constData()+10), t5a6df7.size()-10 };callbackAddressesModel611006_EditAddress(this, v0, v1Packed, v2, v3); };
	void loadModel(QString v0) { QByteArray tea1dd7 = v0.toUtf8(); Moc_PackedString v0Packed = { const_cast<char*>(tea1dd7.prepend("WHITESPACE").constData()+10), tea1dd7.size()-10 };callbackAddressesModel611006_LoadModel(this, v0Packed); };
private:
	type378cdd _roles;
	QList<QAddress611006*> _addresses;
	qint32 _loaded;
};

Q_DECLARE_METATYPE(AddressesModel611006*)


void AddressesModel611006_AddressesModel611006_QRegisterMetaTypes() {
	qRegisterMetaType<type378cdd>("type378cdd");
	qRegisterMetaType<QList<QObject*>>("QList<QAddress611006*>");
}

class QAddress611006: public QObject
{
Q_OBJECT
Q_PROPERTY(QString address READ address WRITE setAddress NOTIFY addressChanged)
Q_PROPERTY(qint32 addressSky READ addressSky WRITE setAddressSky NOTIFY addressSkyChanged)
Q_PROPERTY(qint32 addressCoinHours READ addressCoinHours WRITE setAddressCoinHours NOTIFY addressCoinHoursChanged)
public:
	QAddress611006(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QAddress611006_QAddress611006_QRegisterMetaType();QAddress611006_QAddress611006_QRegisterMetaTypes();callbackQAddress611006_Constructor(this);};
	QString address() { return ({ Moc_PackedString tempVal = callbackQAddress611006_Address(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setAddress(QString address) { QByteArray tc66218 = address.toUtf8(); Moc_PackedString addressPacked = { const_cast<char*>(tc66218.prepend("WHITESPACE").constData()+10), tc66218.size()-10 };callbackQAddress611006_SetAddress(this, addressPacked); };
	void Signal_AddressChanged(QString address) { QByteArray tc66218 = address.toUtf8(); Moc_PackedString addressPacked = { const_cast<char*>(tc66218.prepend("WHITESPACE").constData()+10), tc66218.size()-10 };callbackQAddress611006_AddressChanged(this, addressPacked); };
	qint32 addressSky() { return callbackQAddress611006_AddressSky(this); };
	void setAddressSky(qint32 addressSky) { callbackQAddress611006_SetAddressSky(this, addressSky); };
	void Signal_AddressSkyChanged(qint32 addressSky) { callbackQAddress611006_AddressSkyChanged(this, addressSky); };
	qint32 addressCoinHours() { return callbackQAddress611006_AddressCoinHours(this); };
	void setAddressCoinHours(qint32 addressCoinHours) { callbackQAddress611006_SetAddressCoinHours(this, addressCoinHours); };
	void Signal_AddressCoinHoursChanged(qint32 addressCoinHours) { callbackQAddress611006_AddressCoinHoursChanged(this, addressCoinHours); };
	 ~QAddress611006() { callbackQAddress611006_DestroyQAddress(this); };
	void childEvent(QChildEvent * event) { callbackQAddress611006_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAddress611006_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAddress611006_CustomEvent(this, event); };
	void deleteLater() { callbackQAddress611006_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAddress611006_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAddress611006_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQAddress611006_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAddress611006_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAddress611006_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAddress611006_TimerEvent(this, event); };
	QString addressDefault() { return _address; };
	void setAddressDefault(QString p) { if (p != _address) { _address = p; addressChanged(_address); } };
	qint32 addressSkyDefault() { return _addressSky; };
	void setAddressSkyDefault(qint32 p) { if (p != _addressSky) { _addressSky = p; addressSkyChanged(_addressSky); } };
	qint32 addressCoinHoursDefault() { return _addressCoinHours; };
	void setAddressCoinHoursDefault(qint32 p) { if (p != _addressCoinHours) { _addressCoinHours = p; addressCoinHoursChanged(_addressCoinHours); } };
signals:
	void addressChanged(QString address);
	void addressSkyChanged(qint32 addressSky);
	void addressCoinHoursChanged(qint32 addressCoinHours);
public slots:
private:
	QString _address;
	qint32 _addressSky;
	qint32 _addressCoinHours;
};

Q_DECLARE_METATYPE(QAddress611006*)


void QAddress611006_QAddress611006_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class QWallet611006: public QObject
{
Q_OBJECT
Q_PROPERTY(QString name READ name WRITE setName NOTIFY nameChanged)
Q_PROPERTY(qint32 encryptionEnabled READ encryptionEnabled WRITE setEncryptionEnabled NOTIFY encryptionEnabledChanged)
Q_PROPERTY(qint32 sky READ sky WRITE setSky NOTIFY skyChanged)
Q_PROPERTY(qint32 coinHours READ coinHours WRITE setCoinHours NOTIFY coinHoursChanged)
Q_PROPERTY(QString fileName READ fileName WRITE setFileName NOTIFY fileNameChanged)
public:
	QWallet611006(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QWallet611006_QWallet611006_QRegisterMetaType();QWallet611006_QWallet611006_QRegisterMetaTypes();callbackQWallet611006_Constructor(this);};
	QString name() { return ({ Moc_PackedString tempVal = callbackQWallet611006_Name(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setName(QString name) { QByteArray t6ae999 = name.toUtf8(); Moc_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQWallet611006_SetName(this, namePacked); };
	void Signal_NameChanged(QString name) { QByteArray t6ae999 = name.toUtf8(); Moc_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQWallet611006_NameChanged(this, namePacked); };
	qint32 encryptionEnabled() { return callbackQWallet611006_EncryptionEnabled(this); };
	void setEncryptionEnabled(qint32 encryptionEnabled) { callbackQWallet611006_SetEncryptionEnabled(this, encryptionEnabled); };
	void Signal_EncryptionEnabledChanged(qint32 encryptionEnabled) { callbackQWallet611006_EncryptionEnabledChanged(this, encryptionEnabled); };
	qint32 sky() { return callbackQWallet611006_Sky(this); };
	void setSky(qint32 sky) { callbackQWallet611006_SetSky(this, sky); };
	void Signal_SkyChanged(qint32 sky) { callbackQWallet611006_SkyChanged(this, sky); };
	qint32 coinHours() { return callbackQWallet611006_CoinHours(this); };
	void setCoinHours(qint32 coinHours) { callbackQWallet611006_SetCoinHours(this, coinHours); };
	void Signal_CoinHoursChanged(qint32 coinHours) { callbackQWallet611006_CoinHoursChanged(this, coinHours); };
	QString fileName() { return ({ Moc_PackedString tempVal = callbackQWallet611006_FileName(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setFileName(QString fileName) { QByteArray td83e09 = fileName.toUtf8(); Moc_PackedString fileNamePacked = { const_cast<char*>(td83e09.prepend("WHITESPACE").constData()+10), td83e09.size()-10 };callbackQWallet611006_SetFileName(this, fileNamePacked); };
	void Signal_FileNameChanged(QString fileName) { QByteArray td83e09 = fileName.toUtf8(); Moc_PackedString fileNamePacked = { const_cast<char*>(td83e09.prepend("WHITESPACE").constData()+10), td83e09.size()-10 };callbackQWallet611006_FileNameChanged(this, fileNamePacked); };
	 ~QWallet611006() { callbackQWallet611006_DestroyQWallet(this); };
	void childEvent(QChildEvent * event) { callbackQWallet611006_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWallet611006_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWallet611006_CustomEvent(this, event); };
	void deleteLater() { callbackQWallet611006_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWallet611006_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWallet611006_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQWallet611006_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWallet611006_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWallet611006_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWallet611006_TimerEvent(this, event); };
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

Q_DECLARE_METATYPE(QWallet611006*)


void QWallet611006_QWallet611006_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class WalletManager611006: public QObject
{
Q_OBJECT
public:
	WalletManager611006(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");WalletManager611006_WalletManager611006_QRegisterMetaType();WalletManager611006_WalletManager611006_QRegisterMetaTypes();callbackWalletManager611006_Constructor(this);};
	 ~WalletManager611006() { callbackWalletManager611006_DestroyWalletManager(this); };
	void childEvent(QChildEvent * event) { callbackWalletManager611006_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackWalletManager611006_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackWalletManager611006_CustomEvent(this, event); };
	void deleteLater() { callbackWalletManager611006_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackWalletManager611006_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackWalletManager611006_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackWalletManager611006_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackWalletManager611006_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackWalletManager611006_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackWalletManager611006_TimerEvent(this, event); };
signals:
public slots:
	void createEncryptedWallet(WalletModel611006* walletM, QString seed, QString label, QString password, qint32 scanN) { QByteArray t92713d = seed.toUtf8(); Moc_PackedString seedPacked = { const_cast<char*>(t92713d.prepend("WHITESPACE").constData()+10), t92713d.size()-10 };QByteArray t64c653 = label.toUtf8(); Moc_PackedString labelPacked = { const_cast<char*>(t64c653.prepend("WHITESPACE").constData()+10), t64c653.size()-10 };QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackWalletManager611006_CreateEncryptedWallet(this, walletM, seedPacked, labelPacked, passwordPacked, scanN); };
	void createUnencryptedWallet(WalletModel611006* walletM, QString seed, QString label, qint32 scanN) { QByteArray t92713d = seed.toUtf8(); Moc_PackedString seedPacked = { const_cast<char*>(t92713d.prepend("WHITESPACE").constData()+10), t92713d.size()-10 };QByteArray t64c653 = label.toUtf8(); Moc_PackedString labelPacked = { const_cast<char*>(t64c653.prepend("WHITESPACE").constData()+10), t64c653.size()-10 };callbackWalletManager611006_CreateUnencryptedWallet(this, walletM, seedPacked, labelPacked, scanN); };
	QString getNewSeed(qint32 entropy) { return ({ Moc_PackedString tempVal = callbackWalletManager611006_GetNewSeed(this, entropy); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	qint32 verifySeed(QString seed) { QByteArray t92713d = seed.toUtf8(); Moc_PackedString seedPacked = { const_cast<char*>(t92713d.prepend("WHITESPACE").constData()+10), t92713d.size()-10 };return callbackWalletManager611006_VerifySeed(this, seedPacked); };
	void newWalletAddress(AddressesModel611006* addressM, QString id, qint32 n, QString password) { QByteArray t87ea5d = id.toUtf8(); Moc_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackWalletManager611006_NewWalletAddress(this, addressM, idPacked, n, passwordPacked); };
	void encryptWallet(WalletModel611006* walletM, QString id, QString password) { QByteArray t87ea5d = id.toUtf8(); Moc_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackWalletManager611006_EncryptWallet(this, walletM, idPacked, passwordPacked); };
	void decryptWallet(WalletModel611006* walletM, QString id, QString password) { QByteArray t87ea5d = id.toUtf8(); Moc_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackWalletManager611006_DecryptWallet(this, walletM, idPacked, passwordPacked); };
private:
};

Q_DECLARE_METATYPE(WalletManager611006*)


void WalletManager611006_WalletManager611006_QRegisterMetaTypes() {
}

class WalletModel611006: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<QWallet611006*> wallets READ wallets WRITE setWallets NOTIFY walletsChanged)
public:
	WalletModel611006(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");WalletModel611006_WalletModel611006_QRegisterMetaType();WalletModel611006_WalletModel611006_QRegisterMetaTypes();callbackWalletModel611006_Constructor(this);};
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbackWalletModel611006_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbackWalletModel611006_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbackWalletModel611006_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<QWallet611006*> wallets() { return ({ QList<QWallet611006*>* tmpP = static_cast<QList<QWallet611006*>*>(callbackWalletModel611006_Wallets(this)); QList<QWallet611006*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setWallets(QList<QWallet611006*> wallets) { callbackWalletModel611006_SetWallets(this, ({ QList<QWallet611006*>* tmpValue = new QList<QWallet611006*>(wallets); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_WalletsChanged(QList<QWallet611006*> wallets) { callbackWalletModel611006_WalletsChanged(this, ({ QList<QWallet611006*>* tmpValue = new QList<QWallet611006*>(wallets); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	 ~WalletModel611006() { callbackWalletModel611006_DestroyWalletModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackWalletModel611006_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackWalletModel611006_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackWalletModel611006_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackWalletModel611006_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackWalletModel611006_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackWalletModel611006_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackWalletModel611006_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackWalletModel611006_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackWalletModel611006_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackWalletModel611006_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModel611006_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackWalletModel611006_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackWalletModel611006_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModel611006_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackWalletModel611006_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackWalletModel611006_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = new QVector<int>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackWalletModel611006_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackWalletModel611006_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackWalletModel611006_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackWalletModel611006_HeaderDataChanged(this, orientation, first, last); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackWalletModel611006_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackWalletModel611006_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackWalletModel611006_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackWalletModel611006_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackWalletModel611006_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackWalletModel611006_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackWalletModel611006_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackWalletModel611006_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackWalletModel611006_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackWalletModel611006_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackWalletModel611006_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackWalletModel611006_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackWalletModel611006_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackWalletModel611006_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackWalletModel611006_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackWalletModel611006_ResetInternalData(this); };
	void revert() { callbackWalletModel611006_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackWalletModel611006_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackWalletModel611006_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackWalletModel611006_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackWalletModel611006_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModel611006_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackWalletModel611006_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackWalletModel611006_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModel611006_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackWalletModel611006_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackWalletModel611006_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackWalletModel611006_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackWalletModel611006_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackWalletModel611006_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackWalletModel611006_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackWalletModel611006_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackWalletModel611006_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackWalletModel611006_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackWalletModel611006_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackWalletModel611006_CustomEvent(this, event); };
	void deleteLater() { callbackWalletModel611006_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackWalletModel611006_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackWalletModel611006_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackWalletModel611006_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackWalletModel611006_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackWalletModel611006_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackWalletModel611006_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<QWallet611006*> walletsDefault() { return _wallets; };
	void setWalletsDefault(QList<QWallet611006*> p) { if (p != _wallets) { _wallets = p; walletsChanged(_wallets); } };
signals:
	void rolesChanged(type378cdd roles);
	void walletsChanged(QList<QWallet611006*> wallets);
public slots:
	void addWallet(QWallet611006* v0) { callbackWalletModel611006_AddWallet(this, v0); };
	void editWallet(qint32 row, QString name, bool encryptionEnabled, qint32 sky, qint32 coinHours) { QByteArray t6ae999 = name.toUtf8(); Moc_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackWalletModel611006_EditWallet(this, row, namePacked, encryptionEnabled, sky, coinHours); };
	void removeWallet(qint32 row) { callbackWalletModel611006_RemoveWallet(this, row); };
	void loadModel() { callbackWalletModel611006_LoadModel(this); };
private:
	type378cdd _roles;
	QList<QWallet611006*> _wallets;
};

Q_DECLARE_METATYPE(WalletModel611006*)


void WalletModel611006_WalletModel611006_QRegisterMetaTypes() {
	qRegisterMetaType<type378cdd>("type378cdd");
	qRegisterMetaType<QList<QObject*>>("QList<QWallet611006*>");
}

void AddressesModel611006_AddAddress(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<AddressesModel611006*>(ptr), "addAddress", Q_ARG(QAddress611006*, static_cast<QAddress611006*>(v0)));
}

void AddressesModel611006_RemoveAddress(void* ptr, int v0)
{
	QMetaObject::invokeMethod(static_cast<AddressesModel611006*>(ptr), "removeAddress", Q_ARG(qint32, v0));
}

void AddressesModel611006_EditAddress(void* ptr, int v0, struct Moc_PackedString v1, int v2, int v3)
{
	QMetaObject::invokeMethod(static_cast<AddressesModel611006*>(ptr), "editAddress", Q_ARG(qint32, v0), Q_ARG(QString, QString::fromUtf8(v1.data, v1.len)), Q_ARG(qint32, v2), Q_ARG(qint32, v3));
}

void AddressesModel611006_LoadModel(void* ptr, struct Moc_PackedString v0)
{
	QMetaObject::invokeMethod(static_cast<AddressesModel611006*>(ptr), "loadModel", Q_ARG(QString, QString::fromUtf8(v0.data, v0.len)));
}

struct Moc_PackedList AddressesModel611006_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<AddressesModel611006*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressesModel611006_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<AddressesModel611006*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void AddressesModel611006_SetRoles(void* ptr, void* roles)
{
	static_cast<AddressesModel611006*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void AddressesModel611006_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<AddressesModel611006*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void AddressesModel611006_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<AddressesModel611006*>(ptr), static_cast<void (AddressesModel611006::*)(QMap<qint32, QByteArray>)>(&AddressesModel611006::rolesChanged), static_cast<AddressesModel611006*>(ptr), static_cast<void (AddressesModel611006::*)(QMap<qint32, QByteArray>)>(&AddressesModel611006::Signal_RolesChanged));
}

void AddressesModel611006_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressesModel611006*>(ptr), static_cast<void (AddressesModel611006::*)(QMap<qint32, QByteArray>)>(&AddressesModel611006::rolesChanged), static_cast<AddressesModel611006*>(ptr), static_cast<void (AddressesModel611006::*)(QMap<qint32, QByteArray>)>(&AddressesModel611006::Signal_RolesChanged));
}

void AddressesModel611006_RolesChanged(void* ptr, void* roles)
{
	static_cast<AddressesModel611006*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList AddressesModel611006_Addresses(void* ptr)
{
	return ({ QList<QAddress611006*>* tmpValue = new QList<QAddress611006*>(static_cast<AddressesModel611006*>(ptr)->addresses()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressesModel611006_AddressesDefault(void* ptr)
{
	return ({ QList<QAddress611006*>* tmpValue = new QList<QAddress611006*>(static_cast<AddressesModel611006*>(ptr)->addressesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void AddressesModel611006_SetAddresses(void* ptr, void* addresses)
{
	static_cast<AddressesModel611006*>(ptr)->setAddresses(({ QList<QAddress611006*>* tmpP = static_cast<QList<QAddress611006*>*>(addresses); QList<QAddress611006*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void AddressesModel611006_SetAddressesDefault(void* ptr, void* addresses)
{
	static_cast<AddressesModel611006*>(ptr)->setAddressesDefault(({ QList<QAddress611006*>* tmpP = static_cast<QList<QAddress611006*>*>(addresses); QList<QAddress611006*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void AddressesModel611006_ConnectAddressesChanged(void* ptr)
{
	QObject::connect(static_cast<AddressesModel611006*>(ptr), static_cast<void (AddressesModel611006::*)(QList<QAddress611006*>)>(&AddressesModel611006::addressesChanged), static_cast<AddressesModel611006*>(ptr), static_cast<void (AddressesModel611006::*)(QList<QAddress611006*>)>(&AddressesModel611006::Signal_AddressesChanged));
}

void AddressesModel611006_DisconnectAddressesChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressesModel611006*>(ptr), static_cast<void (AddressesModel611006::*)(QList<QAddress611006*>)>(&AddressesModel611006::addressesChanged), static_cast<AddressesModel611006*>(ptr), static_cast<void (AddressesModel611006::*)(QList<QAddress611006*>)>(&AddressesModel611006::Signal_AddressesChanged));
}

void AddressesModel611006_AddressesChanged(void* ptr, void* addresses)
{
	static_cast<AddressesModel611006*>(ptr)->addressesChanged(({ QList<QAddress611006*>* tmpP = static_cast<QList<QAddress611006*>*>(addresses); QList<QAddress611006*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int AddressesModel611006_Loaded(void* ptr)
{
	return static_cast<AddressesModel611006*>(ptr)->loaded();
}

int AddressesModel611006_LoadedDefault(void* ptr)
{
	return static_cast<AddressesModel611006*>(ptr)->loadedDefault();
}

void AddressesModel611006_SetLoaded(void* ptr, int loaded)
{
	static_cast<AddressesModel611006*>(ptr)->setLoaded(loaded);
}

void AddressesModel611006_SetLoadedDefault(void* ptr, int loaded)
{
	static_cast<AddressesModel611006*>(ptr)->setLoadedDefault(loaded);
}

void AddressesModel611006_ConnectLoadedChanged(void* ptr)
{
	QObject::connect(static_cast<AddressesModel611006*>(ptr), static_cast<void (AddressesModel611006::*)(qint32)>(&AddressesModel611006::loadedChanged), static_cast<AddressesModel611006*>(ptr), static_cast<void (AddressesModel611006::*)(qint32)>(&AddressesModel611006::Signal_LoadedChanged));
}

void AddressesModel611006_DisconnectLoadedChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressesModel611006*>(ptr), static_cast<void (AddressesModel611006::*)(qint32)>(&AddressesModel611006::loadedChanged), static_cast<AddressesModel611006*>(ptr), static_cast<void (AddressesModel611006::*)(qint32)>(&AddressesModel611006::Signal_LoadedChanged));
}

void AddressesModel611006_LoadedChanged(void* ptr, int loaded)
{
	static_cast<AddressesModel611006*>(ptr)->loadedChanged(loaded);
}

int AddressesModel611006_AddressesModel611006_QRegisterMetaType()
{
	return qRegisterMetaType<AddressesModel611006*>();
}

int AddressesModel611006_AddressesModel611006_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<AddressesModel611006*>(const_cast<const char*>(typeName));
}

int AddressesModel611006_AddressesModel611006_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressesModel611006>();
#else
	return 0;
#endif
}

int AddressesModel611006_AddressesModel611006_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressesModel611006>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int AddressesModel611006_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel611006_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModel611006_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressesModel611006_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel611006_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModel611006_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressesModel611006_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel611006_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModel611006_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* AddressesModel611006___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel611006___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModel611006___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressesModel611006___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel611006___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModel611006___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int AddressesModel611006___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void AddressesModel611006___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* AddressesModel611006___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* AddressesModel611006___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModel611006___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* AddressesModel611006___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList AddressesModel611006___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel611006___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel611006___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* AddressesModel611006___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* AddressesModel611006___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel611006___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* AddressesModel611006___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* AddressesModel611006___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel611006___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModel611006___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressesModel611006___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel611006___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModel611006___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressesModel611006___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel611006___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModel611006___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressesModel611006___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void AddressesModel611006___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressesModel611006___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList AddressesModel611006___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel611006___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModel611006___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* AddressesModel611006___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList AddressesModel611006___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int AddressesModel611006_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel611006_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModel611006_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressesModel611006_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel611006_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModel611006_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* AddressesModel611006___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel611006___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressesModel611006___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* AddressesModel611006___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel611006___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* AddressesModel611006___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* AddressesModel611006___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel611006___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressesModel611006___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressesModel611006___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel611006___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressesModel611006___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressesModel611006___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel611006___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressesModel611006___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressesModel611006___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModel611006___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressesModel611006___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressesModel611006___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel611006___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModel611006___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressesModel611006___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressesModel611006___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel611006___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModel611006___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressesModel611006___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressesModel611006___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel611006___addresses_atList(void* ptr, int i)
{
	return ({QAddress611006* tmp = static_cast<QList<QAddress611006*>*>(ptr)->at(i); if (i == static_cast<QList<QAddress611006*>*>(ptr)->size()-1) { static_cast<QList<QAddress611006*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel611006___addresses_setList(void* ptr, void* i)
{
	static_cast<QList<QAddress611006*>*>(ptr)->append(static_cast<QAddress611006*>(i));
}

void* AddressesModel611006___addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAddress611006*>();
}

void* AddressesModel611006___setAddresses_addresses_atList(void* ptr, int i)
{
	return ({QAddress611006* tmp = static_cast<QList<QAddress611006*>*>(ptr)->at(i); if (i == static_cast<QList<QAddress611006*>*>(ptr)->size()-1) { static_cast<QList<QAddress611006*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel611006___setAddresses_addresses_setList(void* ptr, void* i)
{
	static_cast<QList<QAddress611006*>*>(ptr)->append(static_cast<QAddress611006*>(i));
}

void* AddressesModel611006___setAddresses_addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAddress611006*>();
}

void* AddressesModel611006___addressesChanged_addresses_atList(void* ptr, int i)
{
	return ({QAddress611006* tmp = static_cast<QList<QAddress611006*>*>(ptr)->at(i); if (i == static_cast<QList<QAddress611006*>*>(ptr)->size()-1) { static_cast<QList<QAddress611006*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel611006___addressesChanged_addresses_setList(void* ptr, void* i)
{
	static_cast<QList<QAddress611006*>*>(ptr)->append(static_cast<QAddress611006*>(i));
}

void* AddressesModel611006___addressesChanged_addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAddress611006*>();
}

int AddressesModel611006_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel611006_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressesModel611006_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int AddressesModel611006_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel611006_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressesModel611006_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int AddressesModel611006_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel611006_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressesModel611006_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* AddressesModel611006_NewAddressesModel(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new AddressesModel611006(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new AddressesModel611006(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new AddressesModel611006(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new AddressesModel611006(static_cast<QWindow*>(parent));
	} else {
		return new AddressesModel611006(static_cast<QObject*>(parent));
	}
}

void AddressesModel611006_DestroyAddressesModel(void* ptr)
{
	static_cast<AddressesModel611006*>(ptr)->~AddressesModel611006();
}

void AddressesModel611006_DestroyAddressesModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char AddressesModel611006_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

long long AddressesModel611006_FlagsDefault(void* ptr, void* index)
{
	return static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

void* AddressesModel611006_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* AddressesModel611006_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* AddressesModel611006_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

char AddressesModel611006_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char AddressesModel611006_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int AddressesModel611006_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* AddressesModel611006_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void AddressesModel611006_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char AddressesModel611006_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* AddressesModel611006_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

char AddressesModel611006_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char AddressesModel611006_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct Moc_PackedList AddressesModel611006_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressesModel611006_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel611006_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct Moc_PackedString AddressesModel611006_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t6a4f86 = static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::mimeTypes().join("¡¦!").toUtf8(); Moc_PackedString { const_cast<char*>(t6a4f86.prepend("WHITESPACE").constData()+10), t6a4f86.size()-10 }; });
}

char AddressesModel611006_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char AddressesModel611006_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* AddressesModel611006_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

char AddressesModel611006_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char AddressesModel611006_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void AddressesModel611006_ResetInternalDataDefault(void* ptr)
{
	static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::resetInternalData();
}

void AddressesModel611006_RevertDefault(void* ptr)
{
	static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::revert();
}

struct Moc_PackedList AddressesModel611006_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int AddressesModel611006_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char AddressesModel611006_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char AddressesModel611006_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char AddressesModel611006_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void AddressesModel611006_SortDefault(void* ptr, int column, long long order)
{
	static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* AddressesModel611006_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char AddressesModel611006_SubmitDefault(void* ptr)
{
	return static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::submit();
}

long long AddressesModel611006_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long AddressesModel611006_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::supportedDropActions();
}

void AddressesModel611006_ChildEventDefault(void* ptr, void* event)
{
	static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void AddressesModel611006_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void AddressesModel611006_CustomEventDefault(void* ptr, void* event)
{
	static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void AddressesModel611006_DeleteLaterDefault(void* ptr)
{
	static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::deleteLater();
}

void AddressesModel611006_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char AddressesModel611006_EventDefault(void* ptr, void* e)
{
	return static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char AddressesModel611006_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void AddressesModel611006_TimerEventDefault(void* ptr, void* event)
{
	static_cast<AddressesModel611006*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

struct Moc_PackedString QAddress611006_Address(void* ptr)
{
	return ({ QByteArray tbbcec3 = static_cast<QAddress611006*>(ptr)->address().toUtf8(); Moc_PackedString { const_cast<char*>(tbbcec3.prepend("WHITESPACE").constData()+10), tbbcec3.size()-10 }; });
}

struct Moc_PackedString QAddress611006_AddressDefault(void* ptr)
{
	return ({ QByteArray tbe1549 = static_cast<QAddress611006*>(ptr)->addressDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tbe1549.prepend("WHITESPACE").constData()+10), tbe1549.size()-10 }; });
}

void QAddress611006_SetAddress(void* ptr, struct Moc_PackedString address)
{
	static_cast<QAddress611006*>(ptr)->setAddress(QString::fromUtf8(address.data, address.len));
}

void QAddress611006_SetAddressDefault(void* ptr, struct Moc_PackedString address)
{
	static_cast<QAddress611006*>(ptr)->setAddressDefault(QString::fromUtf8(address.data, address.len));
}

void QAddress611006_ConnectAddressChanged(void* ptr)
{
	QObject::connect(static_cast<QAddress611006*>(ptr), static_cast<void (QAddress611006::*)(QString)>(&QAddress611006::addressChanged), static_cast<QAddress611006*>(ptr), static_cast<void (QAddress611006::*)(QString)>(&QAddress611006::Signal_AddressChanged));
}

void QAddress611006_DisconnectAddressChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAddress611006*>(ptr), static_cast<void (QAddress611006::*)(QString)>(&QAddress611006::addressChanged), static_cast<QAddress611006*>(ptr), static_cast<void (QAddress611006::*)(QString)>(&QAddress611006::Signal_AddressChanged));
}

void QAddress611006_AddressChanged(void* ptr, struct Moc_PackedString address)
{
	static_cast<QAddress611006*>(ptr)->addressChanged(QString::fromUtf8(address.data, address.len));
}

int QAddress611006_AddressSky(void* ptr)
{
	return static_cast<QAddress611006*>(ptr)->addressSky();
}

int QAddress611006_AddressSkyDefault(void* ptr)
{
	return static_cast<QAddress611006*>(ptr)->addressSkyDefault();
}

void QAddress611006_SetAddressSky(void* ptr, int addressSky)
{
	static_cast<QAddress611006*>(ptr)->setAddressSky(addressSky);
}

void QAddress611006_SetAddressSkyDefault(void* ptr, int addressSky)
{
	static_cast<QAddress611006*>(ptr)->setAddressSkyDefault(addressSky);
}

void QAddress611006_ConnectAddressSkyChanged(void* ptr)
{
	QObject::connect(static_cast<QAddress611006*>(ptr), static_cast<void (QAddress611006::*)(qint32)>(&QAddress611006::addressSkyChanged), static_cast<QAddress611006*>(ptr), static_cast<void (QAddress611006::*)(qint32)>(&QAddress611006::Signal_AddressSkyChanged));
}

void QAddress611006_DisconnectAddressSkyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAddress611006*>(ptr), static_cast<void (QAddress611006::*)(qint32)>(&QAddress611006::addressSkyChanged), static_cast<QAddress611006*>(ptr), static_cast<void (QAddress611006::*)(qint32)>(&QAddress611006::Signal_AddressSkyChanged));
}

void QAddress611006_AddressSkyChanged(void* ptr, int addressSky)
{
	static_cast<QAddress611006*>(ptr)->addressSkyChanged(addressSky);
}

int QAddress611006_AddressCoinHours(void* ptr)
{
	return static_cast<QAddress611006*>(ptr)->addressCoinHours();
}

int QAddress611006_AddressCoinHoursDefault(void* ptr)
{
	return static_cast<QAddress611006*>(ptr)->addressCoinHoursDefault();
}

void QAddress611006_SetAddressCoinHours(void* ptr, int addressCoinHours)
{
	static_cast<QAddress611006*>(ptr)->setAddressCoinHours(addressCoinHours);
}

void QAddress611006_SetAddressCoinHoursDefault(void* ptr, int addressCoinHours)
{
	static_cast<QAddress611006*>(ptr)->setAddressCoinHoursDefault(addressCoinHours);
}

void QAddress611006_ConnectAddressCoinHoursChanged(void* ptr)
{
	QObject::connect(static_cast<QAddress611006*>(ptr), static_cast<void (QAddress611006::*)(qint32)>(&QAddress611006::addressCoinHoursChanged), static_cast<QAddress611006*>(ptr), static_cast<void (QAddress611006::*)(qint32)>(&QAddress611006::Signal_AddressCoinHoursChanged));
}

void QAddress611006_DisconnectAddressCoinHoursChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAddress611006*>(ptr), static_cast<void (QAddress611006::*)(qint32)>(&QAddress611006::addressCoinHoursChanged), static_cast<QAddress611006*>(ptr), static_cast<void (QAddress611006::*)(qint32)>(&QAddress611006::Signal_AddressCoinHoursChanged));
}

void QAddress611006_AddressCoinHoursChanged(void* ptr, int addressCoinHours)
{
	static_cast<QAddress611006*>(ptr)->addressCoinHoursChanged(addressCoinHours);
}

int QAddress611006_QAddress611006_QRegisterMetaType()
{
	return qRegisterMetaType<QAddress611006*>();
}

int QAddress611006_QAddress611006_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QAddress611006*>(const_cast<const char*>(typeName));
}

int QAddress611006_QAddress611006_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QAddress611006>();
#else
	return 0;
#endif
}

int QAddress611006_QAddress611006_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QAddress611006>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QAddress611006___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAddress611006___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAddress611006___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QAddress611006___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAddress611006___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QAddress611006___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QAddress611006___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAddress611006___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAddress611006___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAddress611006___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAddress611006___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAddress611006___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAddress611006___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAddress611006___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAddress611006___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAddress611006_NewQAddress(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QAddress611006(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QAddress611006(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QAddress611006(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QAddress611006(static_cast<QWindow*>(parent));
	} else {
		return new QAddress611006(static_cast<QObject*>(parent));
	}
}

void QAddress611006_DestroyQAddress(void* ptr)
{
	static_cast<QAddress611006*>(ptr)->~QAddress611006();
}

void QAddress611006_DestroyQAddressDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void QAddress611006_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAddress611006*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QAddress611006_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAddress611006*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAddress611006_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAddress611006*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QAddress611006_DeleteLaterDefault(void* ptr)
{
	static_cast<QAddress611006*>(ptr)->QObject::deleteLater();
}

void QAddress611006_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAddress611006*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QAddress611006_EventDefault(void* ptr, void* e)
{
	return static_cast<QAddress611006*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QAddress611006_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAddress611006*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QAddress611006_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAddress611006*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

struct Moc_PackedString QWallet611006_Name(void* ptr)
{
	return ({ QByteArray tedc734 = static_cast<QWallet611006*>(ptr)->name().toUtf8(); Moc_PackedString { const_cast<char*>(tedc734.prepend("WHITESPACE").constData()+10), tedc734.size()-10 }; });
}

struct Moc_PackedString QWallet611006_NameDefault(void* ptr)
{
	return ({ QByteArray tee7a91 = static_cast<QWallet611006*>(ptr)->nameDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tee7a91.prepend("WHITESPACE").constData()+10), tee7a91.size()-10 }; });
}

void QWallet611006_SetName(void* ptr, struct Moc_PackedString name)
{
	static_cast<QWallet611006*>(ptr)->setName(QString::fromUtf8(name.data, name.len));
}

void QWallet611006_SetNameDefault(void* ptr, struct Moc_PackedString name)
{
	static_cast<QWallet611006*>(ptr)->setNameDefault(QString::fromUtf8(name.data, name.len));
}

void QWallet611006_ConnectNameChanged(void* ptr)
{
	QObject::connect(static_cast<QWallet611006*>(ptr), static_cast<void (QWallet611006::*)(QString)>(&QWallet611006::nameChanged), static_cast<QWallet611006*>(ptr), static_cast<void (QWallet611006::*)(QString)>(&QWallet611006::Signal_NameChanged));
}

void QWallet611006_DisconnectNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWallet611006*>(ptr), static_cast<void (QWallet611006::*)(QString)>(&QWallet611006::nameChanged), static_cast<QWallet611006*>(ptr), static_cast<void (QWallet611006::*)(QString)>(&QWallet611006::Signal_NameChanged));
}

void QWallet611006_NameChanged(void* ptr, struct Moc_PackedString name)
{
	static_cast<QWallet611006*>(ptr)->nameChanged(QString::fromUtf8(name.data, name.len));
}

int QWallet611006_EncryptionEnabled(void* ptr)
{
	return static_cast<QWallet611006*>(ptr)->encryptionEnabled();
}

int QWallet611006_EncryptionEnabledDefault(void* ptr)
{
	return static_cast<QWallet611006*>(ptr)->encryptionEnabledDefault();
}

void QWallet611006_SetEncryptionEnabled(void* ptr, int encryptionEnabled)
{
	static_cast<QWallet611006*>(ptr)->setEncryptionEnabled(encryptionEnabled);
}

void QWallet611006_SetEncryptionEnabledDefault(void* ptr, int encryptionEnabled)
{
	static_cast<QWallet611006*>(ptr)->setEncryptionEnabledDefault(encryptionEnabled);
}

void QWallet611006_ConnectEncryptionEnabledChanged(void* ptr)
{
	QObject::connect(static_cast<QWallet611006*>(ptr), static_cast<void (QWallet611006::*)(qint32)>(&QWallet611006::encryptionEnabledChanged), static_cast<QWallet611006*>(ptr), static_cast<void (QWallet611006::*)(qint32)>(&QWallet611006::Signal_EncryptionEnabledChanged));
}

void QWallet611006_DisconnectEncryptionEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWallet611006*>(ptr), static_cast<void (QWallet611006::*)(qint32)>(&QWallet611006::encryptionEnabledChanged), static_cast<QWallet611006*>(ptr), static_cast<void (QWallet611006::*)(qint32)>(&QWallet611006::Signal_EncryptionEnabledChanged));
}

void QWallet611006_EncryptionEnabledChanged(void* ptr, int encryptionEnabled)
{
	static_cast<QWallet611006*>(ptr)->encryptionEnabledChanged(encryptionEnabled);
}

int QWallet611006_Sky(void* ptr)
{
	return static_cast<QWallet611006*>(ptr)->sky();
}

int QWallet611006_SkyDefault(void* ptr)
{
	return static_cast<QWallet611006*>(ptr)->skyDefault();
}

void QWallet611006_SetSky(void* ptr, int sky)
{
	static_cast<QWallet611006*>(ptr)->setSky(sky);
}

void QWallet611006_SetSkyDefault(void* ptr, int sky)
{
	static_cast<QWallet611006*>(ptr)->setSkyDefault(sky);
}

void QWallet611006_ConnectSkyChanged(void* ptr)
{
	QObject::connect(static_cast<QWallet611006*>(ptr), static_cast<void (QWallet611006::*)(qint32)>(&QWallet611006::skyChanged), static_cast<QWallet611006*>(ptr), static_cast<void (QWallet611006::*)(qint32)>(&QWallet611006::Signal_SkyChanged));
}

void QWallet611006_DisconnectSkyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWallet611006*>(ptr), static_cast<void (QWallet611006::*)(qint32)>(&QWallet611006::skyChanged), static_cast<QWallet611006*>(ptr), static_cast<void (QWallet611006::*)(qint32)>(&QWallet611006::Signal_SkyChanged));
}

void QWallet611006_SkyChanged(void* ptr, int sky)
{
	static_cast<QWallet611006*>(ptr)->skyChanged(sky);
}

int QWallet611006_CoinHours(void* ptr)
{
	return static_cast<QWallet611006*>(ptr)->coinHours();
}

int QWallet611006_CoinHoursDefault(void* ptr)
{
	return static_cast<QWallet611006*>(ptr)->coinHoursDefault();
}

void QWallet611006_SetCoinHours(void* ptr, int coinHours)
{
	static_cast<QWallet611006*>(ptr)->setCoinHours(coinHours);
}

void QWallet611006_SetCoinHoursDefault(void* ptr, int coinHours)
{
	static_cast<QWallet611006*>(ptr)->setCoinHoursDefault(coinHours);
}

void QWallet611006_ConnectCoinHoursChanged(void* ptr)
{
	QObject::connect(static_cast<QWallet611006*>(ptr), static_cast<void (QWallet611006::*)(qint32)>(&QWallet611006::coinHoursChanged), static_cast<QWallet611006*>(ptr), static_cast<void (QWallet611006::*)(qint32)>(&QWallet611006::Signal_CoinHoursChanged));
}

void QWallet611006_DisconnectCoinHoursChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWallet611006*>(ptr), static_cast<void (QWallet611006::*)(qint32)>(&QWallet611006::coinHoursChanged), static_cast<QWallet611006*>(ptr), static_cast<void (QWallet611006::*)(qint32)>(&QWallet611006::Signal_CoinHoursChanged));
}

void QWallet611006_CoinHoursChanged(void* ptr, int coinHours)
{
	static_cast<QWallet611006*>(ptr)->coinHoursChanged(coinHours);
}

struct Moc_PackedString QWallet611006_FileName(void* ptr)
{
	return ({ QByteArray teb23e1 = static_cast<QWallet611006*>(ptr)->fileName().toUtf8(); Moc_PackedString { const_cast<char*>(teb23e1.prepend("WHITESPACE").constData()+10), teb23e1.size()-10 }; });
}

struct Moc_PackedString QWallet611006_FileNameDefault(void* ptr)
{
	return ({ QByteArray t811a22 = static_cast<QWallet611006*>(ptr)->fileNameDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t811a22.prepend("WHITESPACE").constData()+10), t811a22.size()-10 }; });
}

void QWallet611006_SetFileName(void* ptr, struct Moc_PackedString fileName)
{
	static_cast<QWallet611006*>(ptr)->setFileName(QString::fromUtf8(fileName.data, fileName.len));
}

void QWallet611006_SetFileNameDefault(void* ptr, struct Moc_PackedString fileName)
{
	static_cast<QWallet611006*>(ptr)->setFileNameDefault(QString::fromUtf8(fileName.data, fileName.len));
}

void QWallet611006_ConnectFileNameChanged(void* ptr)
{
	QObject::connect(static_cast<QWallet611006*>(ptr), static_cast<void (QWallet611006::*)(QString)>(&QWallet611006::fileNameChanged), static_cast<QWallet611006*>(ptr), static_cast<void (QWallet611006::*)(QString)>(&QWallet611006::Signal_FileNameChanged));
}

void QWallet611006_DisconnectFileNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWallet611006*>(ptr), static_cast<void (QWallet611006::*)(QString)>(&QWallet611006::fileNameChanged), static_cast<QWallet611006*>(ptr), static_cast<void (QWallet611006::*)(QString)>(&QWallet611006::Signal_FileNameChanged));
}

void QWallet611006_FileNameChanged(void* ptr, struct Moc_PackedString fileName)
{
	static_cast<QWallet611006*>(ptr)->fileNameChanged(QString::fromUtf8(fileName.data, fileName.len));
}

int QWallet611006_QWallet611006_QRegisterMetaType()
{
	return qRegisterMetaType<QWallet611006*>();
}

int QWallet611006_QWallet611006_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QWallet611006*>(const_cast<const char*>(typeName));
}

int QWallet611006_QWallet611006_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QWallet611006>();
#else
	return 0;
#endif
}

int QWallet611006_QWallet611006_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QWallet611006>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QWallet611006___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWallet611006___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWallet611006___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QWallet611006___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWallet611006___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWallet611006___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QWallet611006___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWallet611006___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWallet611006___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWallet611006___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWallet611006___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWallet611006___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWallet611006___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWallet611006___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWallet611006___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWallet611006_NewQWallet(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QWallet611006(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QWallet611006(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QWallet611006(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QWallet611006(static_cast<QWindow*>(parent));
	} else {
		return new QWallet611006(static_cast<QObject*>(parent));
	}
}

void QWallet611006_DestroyQWallet(void* ptr)
{
	static_cast<QWallet611006*>(ptr)->~QWallet611006();
}

void QWallet611006_DestroyQWalletDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void QWallet611006_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QWallet611006*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QWallet611006_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWallet611006*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWallet611006_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QWallet611006*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QWallet611006_DeleteLaterDefault(void* ptr)
{
	static_cast<QWallet611006*>(ptr)->QObject::deleteLater();
}

void QWallet611006_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWallet611006*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QWallet611006_EventDefault(void* ptr, void* e)
{
	return static_cast<QWallet611006*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QWallet611006_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QWallet611006*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWallet611006_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QWallet611006*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

void WalletManager611006_CreateEncryptedWallet(void* ptr, void* walletM, struct Moc_PackedString seed, struct Moc_PackedString label, struct Moc_PackedString password, int scanN)
{
	QMetaObject::invokeMethod(static_cast<WalletManager611006*>(ptr), "createEncryptedWallet", Q_ARG(WalletModel611006*, static_cast<WalletModel611006*>(walletM)), Q_ARG(QString, QString::fromUtf8(seed.data, seed.len)), Q_ARG(QString, QString::fromUtf8(label.data, label.len)), Q_ARG(QString, QString::fromUtf8(password.data, password.len)), Q_ARG(qint32, scanN));
}

void WalletManager611006_CreateUnencryptedWallet(void* ptr, void* walletM, struct Moc_PackedString seed, struct Moc_PackedString label, int scanN)
{
	QMetaObject::invokeMethod(static_cast<WalletManager611006*>(ptr), "createUnencryptedWallet", Q_ARG(WalletModel611006*, static_cast<WalletModel611006*>(walletM)), Q_ARG(QString, QString::fromUtf8(seed.data, seed.len)), Q_ARG(QString, QString::fromUtf8(label.data, label.len)), Q_ARG(qint32, scanN));
}

struct Moc_PackedString WalletManager611006_GetNewSeed(void* ptr, int entropy)
{
	QString returnArg;
	QMetaObject::invokeMethod(static_cast<WalletManager611006*>(ptr), "getNewSeed", Q_RETURN_ARG(QString, returnArg), Q_ARG(qint32, entropy));
	return ({ QByteArray t8e5b69 = returnArg.toUtf8(); Moc_PackedString { const_cast<char*>(t8e5b69.prepend("WHITESPACE").constData()+10), t8e5b69.size()-10 }; });
}

int WalletManager611006_VerifySeed(void* ptr, struct Moc_PackedString seed)
{
	qint32 returnArg;
	QMetaObject::invokeMethod(static_cast<WalletManager611006*>(ptr), "verifySeed", Q_RETURN_ARG(qint32, returnArg), Q_ARG(QString, QString::fromUtf8(seed.data, seed.len)));
	return returnArg;
}

void WalletManager611006_NewWalletAddress(void* ptr, void* addressM, struct Moc_PackedString id, int n, struct Moc_PackedString password)
{
	QMetaObject::invokeMethod(static_cast<WalletManager611006*>(ptr), "newWalletAddress", Q_ARG(AddressesModel611006*, static_cast<AddressesModel611006*>(addressM)), Q_ARG(QString, QString::fromUtf8(id.data, id.len)), Q_ARG(qint32, n), Q_ARG(QString, QString::fromUtf8(password.data, password.len)));
}

void WalletManager611006_EncryptWallet(void* ptr, void* walletM, struct Moc_PackedString id, struct Moc_PackedString password)
{
	QMetaObject::invokeMethod(static_cast<WalletManager611006*>(ptr), "encryptWallet", Q_ARG(WalletModel611006*, static_cast<WalletModel611006*>(walletM)), Q_ARG(QString, QString::fromUtf8(id.data, id.len)), Q_ARG(QString, QString::fromUtf8(password.data, password.len)));
}

void WalletManager611006_DecryptWallet(void* ptr, void* walletM, struct Moc_PackedString id, struct Moc_PackedString password)
{
	QMetaObject::invokeMethod(static_cast<WalletManager611006*>(ptr), "decryptWallet", Q_ARG(WalletModel611006*, static_cast<WalletModel611006*>(walletM)), Q_ARG(QString, QString::fromUtf8(id.data, id.len)), Q_ARG(QString, QString::fromUtf8(password.data, password.len)));
}

int WalletManager611006_WalletManager611006_QRegisterMetaType()
{
	return qRegisterMetaType<WalletManager611006*>();
}

int WalletManager611006_WalletManager611006_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<WalletManager611006*>(const_cast<const char*>(typeName));
}

int WalletManager611006_WalletManager611006_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletManager611006>();
#else
	return 0;
#endif
}

int WalletManager611006_WalletManager611006_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletManager611006>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* WalletManager611006___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager611006___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletManager611006___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* WalletManager611006___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletManager611006___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* WalletManager611006___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* WalletManager611006___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager611006___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletManager611006___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletManager611006___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager611006___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletManager611006___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletManager611006___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager611006___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletManager611006___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletManager611006_NewWalletManager(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new WalletManager611006(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new WalletManager611006(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new WalletManager611006(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new WalletManager611006(static_cast<QWindow*>(parent));
	} else {
		return new WalletManager611006(static_cast<QObject*>(parent));
	}
}

void WalletManager611006_DestroyWalletManager(void* ptr)
{
	static_cast<WalletManager611006*>(ptr)->~WalletManager611006();
}

void WalletManager611006_DestroyWalletManagerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void WalletManager611006_ChildEventDefault(void* ptr, void* event)
{
	static_cast<WalletManager611006*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void WalletManager611006_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletManager611006*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void WalletManager611006_CustomEventDefault(void* ptr, void* event)
{
	static_cast<WalletManager611006*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void WalletManager611006_DeleteLaterDefault(void* ptr)
{
	static_cast<WalletManager611006*>(ptr)->QObject::deleteLater();
}

void WalletManager611006_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletManager611006*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char WalletManager611006_EventDefault(void* ptr, void* e)
{
	return static_cast<WalletManager611006*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char WalletManager611006_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<WalletManager611006*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void WalletManager611006_TimerEventDefault(void* ptr, void* event)
{
	static_cast<WalletManager611006*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

void WalletModel611006_AddWallet(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<WalletModel611006*>(ptr), "addWallet", Q_ARG(QWallet611006*, static_cast<QWallet611006*>(v0)));
}

void WalletModel611006_EditWallet(void* ptr, int row, struct Moc_PackedString name, char encryptionEnabled, int sky, int coinHours)
{
	QMetaObject::invokeMethod(static_cast<WalletModel611006*>(ptr), "editWallet", Q_ARG(qint32, row), Q_ARG(QString, QString::fromUtf8(name.data, name.len)), Q_ARG(bool, encryptionEnabled != 0), Q_ARG(qint32, sky), Q_ARG(qint32, coinHours));
}

void WalletModel611006_RemoveWallet(void* ptr, int row)
{
	QMetaObject::invokeMethod(static_cast<WalletModel611006*>(ptr), "removeWallet", Q_ARG(qint32, row));
}

void WalletModel611006_LoadModel(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<WalletModel611006*>(ptr), "loadModel");
}

struct Moc_PackedList WalletModel611006_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<WalletModel611006*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList WalletModel611006_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<WalletModel611006*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void WalletModel611006_SetRoles(void* ptr, void* roles)
{
	static_cast<WalletModel611006*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void WalletModel611006_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<WalletModel611006*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void WalletModel611006_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<WalletModel611006*>(ptr), static_cast<void (WalletModel611006::*)(QMap<qint32, QByteArray>)>(&WalletModel611006::rolesChanged), static_cast<WalletModel611006*>(ptr), static_cast<void (WalletModel611006::*)(QMap<qint32, QByteArray>)>(&WalletModel611006::Signal_RolesChanged));
}

void WalletModel611006_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<WalletModel611006*>(ptr), static_cast<void (WalletModel611006::*)(QMap<qint32, QByteArray>)>(&WalletModel611006::rolesChanged), static_cast<WalletModel611006*>(ptr), static_cast<void (WalletModel611006::*)(QMap<qint32, QByteArray>)>(&WalletModel611006::Signal_RolesChanged));
}

void WalletModel611006_RolesChanged(void* ptr, void* roles)
{
	static_cast<WalletModel611006*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList WalletModel611006_Wallets(void* ptr)
{
	return ({ QList<QWallet611006*>* tmpValue = new QList<QWallet611006*>(static_cast<WalletModel611006*>(ptr)->wallets()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList WalletModel611006_WalletsDefault(void* ptr)
{
	return ({ QList<QWallet611006*>* tmpValue = new QList<QWallet611006*>(static_cast<WalletModel611006*>(ptr)->walletsDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void WalletModel611006_SetWallets(void* ptr, void* wallets)
{
	static_cast<WalletModel611006*>(ptr)->setWallets(({ QList<QWallet611006*>* tmpP = static_cast<QList<QWallet611006*>*>(wallets); QList<QWallet611006*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void WalletModel611006_SetWalletsDefault(void* ptr, void* wallets)
{
	static_cast<WalletModel611006*>(ptr)->setWalletsDefault(({ QList<QWallet611006*>* tmpP = static_cast<QList<QWallet611006*>*>(wallets); QList<QWallet611006*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void WalletModel611006_ConnectWalletsChanged(void* ptr)
{
	QObject::connect(static_cast<WalletModel611006*>(ptr), static_cast<void (WalletModel611006::*)(QList<QWallet611006*>)>(&WalletModel611006::walletsChanged), static_cast<WalletModel611006*>(ptr), static_cast<void (WalletModel611006::*)(QList<QWallet611006*>)>(&WalletModel611006::Signal_WalletsChanged));
}

void WalletModel611006_DisconnectWalletsChanged(void* ptr)
{
	QObject::disconnect(static_cast<WalletModel611006*>(ptr), static_cast<void (WalletModel611006::*)(QList<QWallet611006*>)>(&WalletModel611006::walletsChanged), static_cast<WalletModel611006*>(ptr), static_cast<void (WalletModel611006::*)(QList<QWallet611006*>)>(&WalletModel611006::Signal_WalletsChanged));
}

void WalletModel611006_WalletsChanged(void* ptr, void* wallets)
{
	static_cast<WalletModel611006*>(ptr)->walletsChanged(({ QList<QWallet611006*>* tmpP = static_cast<QList<QWallet611006*>*>(wallets); QList<QWallet611006*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int WalletModel611006_WalletModel611006_QRegisterMetaType()
{
	return qRegisterMetaType<WalletModel611006*>();
}

int WalletModel611006_WalletModel611006_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<WalletModel611006*>(const_cast<const char*>(typeName));
}

int WalletModel611006_WalletModel611006_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletModel611006>();
#else
	return 0;
#endif
}

int WalletModel611006_WalletModel611006_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletModel611006>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int WalletModel611006_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel611006_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel611006_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int WalletModel611006_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel611006_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel611006_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int WalletModel611006_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel611006_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel611006_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* WalletModel611006___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel611006___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel611006___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModel611006___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel611006___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel611006___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int WalletModel611006___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void WalletModel611006___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* WalletModel611006___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* WalletModel611006___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel611006___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* WalletModel611006___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList WalletModel611006___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel611006___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel611006___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* WalletModel611006___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* WalletModel611006___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel611006___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* WalletModel611006___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* WalletModel611006___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel611006___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel611006___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModel611006___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel611006___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel611006___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModel611006___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel611006___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel611006___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModel611006___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void WalletModel611006___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModel611006___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList WalletModel611006___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel611006___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel611006___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* WalletModel611006___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList WalletModel611006___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int WalletModel611006_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel611006_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel611006_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int WalletModel611006_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel611006_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel611006_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* WalletModel611006___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel611006___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModel611006___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* WalletModel611006___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel611006___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* WalletModel611006___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* WalletModel611006___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel611006___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModel611006___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletModel611006___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel611006___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModel611006___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletModel611006___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel611006___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModel611006___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletModel611006___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel611006___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModel611006___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList WalletModel611006___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel611006___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel611006___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModel611006___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList WalletModel611006___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel611006___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel611006___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModel611006___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList WalletModel611006___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel611006___wallets_atList(void* ptr, int i)
{
	return ({QWallet611006* tmp = static_cast<QList<QWallet611006*>*>(ptr)->at(i); if (i == static_cast<QList<QWallet611006*>*>(ptr)->size()-1) { static_cast<QList<QWallet611006*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel611006___wallets_setList(void* ptr, void* i)
{
	static_cast<QList<QWallet611006*>*>(ptr)->append(static_cast<QWallet611006*>(i));
}

void* WalletModel611006___wallets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWallet611006*>();
}

void* WalletModel611006___setWallets_wallets_atList(void* ptr, int i)
{
	return ({QWallet611006* tmp = static_cast<QList<QWallet611006*>*>(ptr)->at(i); if (i == static_cast<QList<QWallet611006*>*>(ptr)->size()-1) { static_cast<QList<QWallet611006*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel611006___setWallets_wallets_setList(void* ptr, void* i)
{
	static_cast<QList<QWallet611006*>*>(ptr)->append(static_cast<QWallet611006*>(i));
}

void* WalletModel611006___setWallets_wallets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWallet611006*>();
}

void* WalletModel611006___walletsChanged_wallets_atList(void* ptr, int i)
{
	return ({QWallet611006* tmp = static_cast<QList<QWallet611006*>*>(ptr)->at(i); if (i == static_cast<QList<QWallet611006*>*>(ptr)->size()-1) { static_cast<QList<QWallet611006*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel611006___walletsChanged_wallets_setList(void* ptr, void* i)
{
	static_cast<QList<QWallet611006*>*>(ptr)->append(static_cast<QWallet611006*>(i));
}

void* WalletModel611006___walletsChanged_wallets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWallet611006*>();
}

int WalletModel611006_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel611006_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* WalletModel611006_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int WalletModel611006_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel611006_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* WalletModel611006_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int WalletModel611006_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel611006_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* WalletModel611006_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* WalletModel611006_NewWalletModel(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new WalletModel611006(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new WalletModel611006(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new WalletModel611006(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new WalletModel611006(static_cast<QWindow*>(parent));
	} else {
		return new WalletModel611006(static_cast<QObject*>(parent));
	}
}

void WalletModel611006_DestroyWalletModel(void* ptr)
{
	static_cast<WalletModel611006*>(ptr)->~WalletModel611006();
}

void WalletModel611006_DestroyWalletModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char WalletModel611006_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<WalletModel611006*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

long long WalletModel611006_FlagsDefault(void* ptr, void* index)
{
	return static_cast<WalletModel611006*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

void* WalletModel611006_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<WalletModel611006*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* WalletModel611006_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<WalletModel611006*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* WalletModel611006_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<WalletModel611006*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

char WalletModel611006_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<WalletModel611006*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char WalletModel611006_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<WalletModel611006*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int WalletModel611006_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<WalletModel611006*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* WalletModel611006_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void WalletModel611006_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<WalletModel611006*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char WalletModel611006_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<WalletModel611006*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* WalletModel611006_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<WalletModel611006*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

char WalletModel611006_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<WalletModel611006*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char WalletModel611006_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<WalletModel611006*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct Moc_PackedList WalletModel611006_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<WalletModel611006*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList WalletModel611006_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<WalletModel611006*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel611006_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<WalletModel611006*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct Moc_PackedString WalletModel611006_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t2c50d2 = static_cast<WalletModel611006*>(ptr)->QAbstractListModel::mimeTypes().join("¡¦!").toUtf8(); Moc_PackedString { const_cast<char*>(t2c50d2.prepend("WHITESPACE").constData()+10), t2c50d2.size()-10 }; });
}

char WalletModel611006_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<WalletModel611006*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char WalletModel611006_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<WalletModel611006*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* WalletModel611006_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<WalletModel611006*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

char WalletModel611006_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<WalletModel611006*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char WalletModel611006_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<WalletModel611006*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void WalletModel611006_ResetInternalDataDefault(void* ptr)
{
	static_cast<WalletModel611006*>(ptr)->QAbstractListModel::resetInternalData();
}

void WalletModel611006_RevertDefault(void* ptr)
{
	static_cast<WalletModel611006*>(ptr)->QAbstractListModel::revert();
}

struct Moc_PackedList WalletModel611006_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<WalletModel611006*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int WalletModel611006_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char WalletModel611006_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<WalletModel611006*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char WalletModel611006_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<WalletModel611006*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char WalletModel611006_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<WalletModel611006*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void WalletModel611006_SortDefault(void* ptr, int column, long long order)
{
	static_cast<WalletModel611006*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* WalletModel611006_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<WalletModel611006*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char WalletModel611006_SubmitDefault(void* ptr)
{
	return static_cast<WalletModel611006*>(ptr)->QAbstractListModel::submit();
}

long long WalletModel611006_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<WalletModel611006*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long WalletModel611006_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<WalletModel611006*>(ptr)->QAbstractListModel::supportedDropActions();
}

void WalletModel611006_ChildEventDefault(void* ptr, void* event)
{
	static_cast<WalletModel611006*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void WalletModel611006_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletModel611006*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void WalletModel611006_CustomEventDefault(void* ptr, void* event)
{
	static_cast<WalletModel611006*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void WalletModel611006_DeleteLaterDefault(void* ptr)
{
	static_cast<WalletModel611006*>(ptr)->QAbstractListModel::deleteLater();
}

void WalletModel611006_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletModel611006*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char WalletModel611006_EventDefault(void* ptr, void* e)
{
	return static_cast<WalletModel611006*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char WalletModel611006_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<WalletModel611006*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void WalletModel611006_TimerEventDefault(void* ptr, void* event)
{
	static_cast<WalletModel611006*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
