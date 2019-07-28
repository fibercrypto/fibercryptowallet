

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
class AddressesModel445aa6: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<QAddress445aa6*> addresses READ addresses WRITE setAddresses NOTIFY addressesChanged)
public:
	AddressesModel445aa6(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");AddressesModel445aa6_AddressesModel445aa6_QRegisterMetaType();AddressesModel445aa6_AddressesModel445aa6_QRegisterMetaTypes();callbackAddressesModel445aa6_Constructor(this);};
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbackAddressesModel445aa6_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbackAddressesModel445aa6_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbackAddressesModel445aa6_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<QAddress445aa6*> addresses() { return ({ QList<QAddress445aa6*>* tmpP = static_cast<QList<QAddress445aa6*>*>(callbackAddressesModel445aa6_Addresses(this)); QList<QAddress445aa6*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setAddresses(QList<QAddress445aa6*> addresses) { callbackAddressesModel445aa6_SetAddresses(this, ({ QList<QAddress445aa6*>* tmpValue = new QList<QAddress445aa6*>(addresses); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_AddressesChanged(QList<QAddress445aa6*> addresses) { callbackAddressesModel445aa6_AddressesChanged(this, ({ QList<QAddress445aa6*>* tmpValue = new QList<QAddress445aa6*>(addresses); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	 ~AddressesModel445aa6() { callbackAddressesModel445aa6_DestroyAddressesModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackAddressesModel445aa6_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackAddressesModel445aa6_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackAddressesModel445aa6_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackAddressesModel445aa6_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackAddressesModel445aa6_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackAddressesModel445aa6_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackAddressesModel445aa6_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackAddressesModel445aa6_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackAddressesModel445aa6_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackAddressesModel445aa6_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackAddressesModel445aa6_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackAddressesModel445aa6_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackAddressesModel445aa6_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackAddressesModel445aa6_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackAddressesModel445aa6_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackAddressesModel445aa6_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = new QVector<int>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackAddressesModel445aa6_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackAddressesModel445aa6_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackAddressesModel445aa6_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackAddressesModel445aa6_HeaderDataChanged(this, orientation, first, last); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackAddressesModel445aa6_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackAddressesModel445aa6_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackAddressesModel445aa6_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackAddressesModel445aa6_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackAddressesModel445aa6_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackAddressesModel445aa6_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackAddressesModel445aa6_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackAddressesModel445aa6_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackAddressesModel445aa6_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackAddressesModel445aa6_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackAddressesModel445aa6_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackAddressesModel445aa6_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackAddressesModel445aa6_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackAddressesModel445aa6_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackAddressesModel445aa6_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackAddressesModel445aa6_ResetInternalData(this); };
	void revert() { callbackAddressesModel445aa6_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackAddressesModel445aa6_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackAddressesModel445aa6_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackAddressesModel445aa6_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackAddressesModel445aa6_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackAddressesModel445aa6_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackAddressesModel445aa6_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackAddressesModel445aa6_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackAddressesModel445aa6_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackAddressesModel445aa6_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackAddressesModel445aa6_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackAddressesModel445aa6_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackAddressesModel445aa6_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackAddressesModel445aa6_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackAddressesModel445aa6_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackAddressesModel445aa6_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackAddressesModel445aa6_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackAddressesModel445aa6_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackAddressesModel445aa6_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackAddressesModel445aa6_CustomEvent(this, event); };
	void deleteLater() { callbackAddressesModel445aa6_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackAddressesModel445aa6_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackAddressesModel445aa6_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackAddressesModel445aa6_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackAddressesModel445aa6_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackAddressesModel445aa6_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackAddressesModel445aa6_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<QAddress445aa6*> addressesDefault() { return _addresses; };
	void setAddressesDefault(QList<QAddress445aa6*> p) { if (p != _addresses) { _addresses = p; addressesChanged(_addresses); } };
signals:
	void rolesChanged(type378cdd roles);
	void addressesChanged(QList<QAddress445aa6*> addresses);
public slots:
	void addAddress(QAddress445aa6* v0) { callbackAddressesModel445aa6_AddAddress(this, v0); };
	void removeAddress(qint32 v0) { callbackAddressesModel445aa6_RemoveAddress(this, v0); };
	void editAddress(qint32 v0, QString v1, quint64 v2, quint64 v3) { QByteArray t5a6df7 = v1.toUtf8(); Moc_PackedString v1Packed = { const_cast<char*>(t5a6df7.prepend("WHITESPACE").constData()+10), t5a6df7.size()-10 };callbackAddressesModel445aa6_EditAddress(this, v0, v1Packed, v2, v3); };
	void loadModel(QList<QAddress445aa6*> v0) { callbackAddressesModel445aa6_LoadModel(this, ({ QList<QAddress445aa6*>* tmpValue = new QList<QAddress445aa6*>(v0); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
private:
	type378cdd _roles;
	QList<QAddress445aa6*> _addresses;
};

Q_DECLARE_METATYPE(AddressesModel445aa6*)


void AddressesModel445aa6_AddressesModel445aa6_QRegisterMetaTypes() {
	qRegisterMetaType<QList<QObject*>>("QList<QAddress445aa6*>");
	qRegisterMetaType<type378cdd>("type378cdd");
}

class QAddress445aa6: public QObject
{
Q_OBJECT
Q_PROPERTY(QString address READ address WRITE setAddress NOTIFY addressChanged)
Q_PROPERTY(quint64 addressSky READ addressSky WRITE setAddressSky NOTIFY addressSkyChanged)
Q_PROPERTY(quint64 addressCoinHours READ addressCoinHours WRITE setAddressCoinHours NOTIFY addressCoinHoursChanged)
public:
	QAddress445aa6(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QAddress445aa6_QAddress445aa6_QRegisterMetaType();QAddress445aa6_QAddress445aa6_QRegisterMetaTypes();callbackQAddress445aa6_Constructor(this);};
	QString address() { return ({ Moc_PackedString tempVal = callbackQAddress445aa6_Address(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setAddress(QString address) { QByteArray tc66218 = address.toUtf8(); Moc_PackedString addressPacked = { const_cast<char*>(tc66218.prepend("WHITESPACE").constData()+10), tc66218.size()-10 };callbackQAddress445aa6_SetAddress(this, addressPacked); };
	void Signal_AddressChanged(QString address) { QByteArray tc66218 = address.toUtf8(); Moc_PackedString addressPacked = { const_cast<char*>(tc66218.prepend("WHITESPACE").constData()+10), tc66218.size()-10 };callbackQAddress445aa6_AddressChanged(this, addressPacked); };
	quint64 addressSky() { return callbackQAddress445aa6_AddressSky(this); };
	void setAddressSky(quint64 addressSky) { callbackQAddress445aa6_SetAddressSky(this, addressSky); };
	void Signal_AddressSkyChanged(quint64 addressSky) { callbackQAddress445aa6_AddressSkyChanged(this, addressSky); };
	quint64 addressCoinHours() { return callbackQAddress445aa6_AddressCoinHours(this); };
	void setAddressCoinHours(quint64 addressCoinHours) { callbackQAddress445aa6_SetAddressCoinHours(this, addressCoinHours); };
	void Signal_AddressCoinHoursChanged(quint64 addressCoinHours) { callbackQAddress445aa6_AddressCoinHoursChanged(this, addressCoinHours); };
	 ~QAddress445aa6() { callbackQAddress445aa6_DestroyQAddress(this); };
	void childEvent(QChildEvent * event) { callbackQAddress445aa6_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAddress445aa6_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAddress445aa6_CustomEvent(this, event); };
	void deleteLater() { callbackQAddress445aa6_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAddress445aa6_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAddress445aa6_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQAddress445aa6_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAddress445aa6_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAddress445aa6_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAddress445aa6_TimerEvent(this, event); };
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

Q_DECLARE_METATYPE(QAddress445aa6*)


void QAddress445aa6_QAddress445aa6_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class QWallet445aa6: public QObject
{
Q_OBJECT
Q_PROPERTY(QString name READ name WRITE setName NOTIFY nameChanged)
Q_PROPERTY(qint32 encryptionEnabled READ encryptionEnabled WRITE setEncryptionEnabled NOTIFY encryptionEnabledChanged)
Q_PROPERTY(quint64 sky READ sky WRITE setSky NOTIFY skyChanged)
Q_PROPERTY(quint64 coinHours READ coinHours WRITE setCoinHours NOTIFY coinHoursChanged)
Q_PROPERTY(QString fileName READ fileName WRITE setFileName NOTIFY fileNameChanged)
public:
	QWallet445aa6(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QWallet445aa6_QWallet445aa6_QRegisterMetaType();QWallet445aa6_QWallet445aa6_QRegisterMetaTypes();callbackQWallet445aa6_Constructor(this);};
	QString name() { return ({ Moc_PackedString tempVal = callbackQWallet445aa6_Name(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setName(QString name) { QByteArray t6ae999 = name.toUtf8(); Moc_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQWallet445aa6_SetName(this, namePacked); };
	void Signal_NameChanged(QString name) { QByteArray t6ae999 = name.toUtf8(); Moc_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQWallet445aa6_NameChanged(this, namePacked); };
	qint32 encryptionEnabled() { return callbackQWallet445aa6_EncryptionEnabled(this); };
	void setEncryptionEnabled(qint32 encryptionEnabled) { callbackQWallet445aa6_SetEncryptionEnabled(this, encryptionEnabled); };
	void Signal_EncryptionEnabledChanged(qint32 encryptionEnabled) { callbackQWallet445aa6_EncryptionEnabledChanged(this, encryptionEnabled); };
	quint64 sky() { return callbackQWallet445aa6_Sky(this); };
	void setSky(quint64 sky) { callbackQWallet445aa6_SetSky(this, sky); };
	void Signal_SkyChanged(quint64 sky) { callbackQWallet445aa6_SkyChanged(this, sky); };
	quint64 coinHours() { return callbackQWallet445aa6_CoinHours(this); };
	void setCoinHours(quint64 coinHours) { callbackQWallet445aa6_SetCoinHours(this, coinHours); };
	void Signal_CoinHoursChanged(quint64 coinHours) { callbackQWallet445aa6_CoinHoursChanged(this, coinHours); };
	QString fileName() { return ({ Moc_PackedString tempVal = callbackQWallet445aa6_FileName(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setFileName(QString fileName) { QByteArray td83e09 = fileName.toUtf8(); Moc_PackedString fileNamePacked = { const_cast<char*>(td83e09.prepend("WHITESPACE").constData()+10), td83e09.size()-10 };callbackQWallet445aa6_SetFileName(this, fileNamePacked); };
	void Signal_FileNameChanged(QString fileName) { QByteArray td83e09 = fileName.toUtf8(); Moc_PackedString fileNamePacked = { const_cast<char*>(td83e09.prepend("WHITESPACE").constData()+10), td83e09.size()-10 };callbackQWallet445aa6_FileNameChanged(this, fileNamePacked); };
	 ~QWallet445aa6() { callbackQWallet445aa6_DestroyQWallet(this); };
	void childEvent(QChildEvent * event) { callbackQWallet445aa6_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWallet445aa6_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWallet445aa6_CustomEvent(this, event); };
	void deleteLater() { callbackQWallet445aa6_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWallet445aa6_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWallet445aa6_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQWallet445aa6_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWallet445aa6_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWallet445aa6_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWallet445aa6_TimerEvent(this, event); };
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

Q_DECLARE_METATYPE(QWallet445aa6*)


void QWallet445aa6_QWallet445aa6_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class WalletModel445aa6: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<QWallet445aa6*> wallets READ wallets WRITE setWallets NOTIFY walletsChanged)
Q_PROPERTY(qint32 count READ count WRITE setCount NOTIFY countChanged)
public:
	WalletModel445aa6(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");WalletModel445aa6_WalletModel445aa6_QRegisterMetaType();WalletModel445aa6_WalletModel445aa6_QRegisterMetaTypes();callbackWalletModel445aa6_Constructor(this);};
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbackWalletModel445aa6_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbackWalletModel445aa6_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbackWalletModel445aa6_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<QWallet445aa6*> wallets() { return ({ QList<QWallet445aa6*>* tmpP = static_cast<QList<QWallet445aa6*>*>(callbackWalletModel445aa6_Wallets(this)); QList<QWallet445aa6*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setWallets(QList<QWallet445aa6*> wallets) { callbackWalletModel445aa6_SetWallets(this, ({ QList<QWallet445aa6*>* tmpValue = new QList<QWallet445aa6*>(wallets); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_WalletsChanged(QList<QWallet445aa6*> wallets) { callbackWalletModel445aa6_WalletsChanged(this, ({ QList<QWallet445aa6*>* tmpValue = new QList<QWallet445aa6*>(wallets); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	qint32 count() { return callbackWalletModel445aa6_Count(this); };
	void setCount(qint32 count) { callbackWalletModel445aa6_SetCount(this, count); };
	void Signal_CountChanged(qint32 count) { callbackWalletModel445aa6_CountChanged(this, count); };
	 ~WalletModel445aa6() { callbackWalletModel445aa6_DestroyWalletModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackWalletModel445aa6_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackWalletModel445aa6_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackWalletModel445aa6_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackWalletModel445aa6_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackWalletModel445aa6_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackWalletModel445aa6_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackWalletModel445aa6_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackWalletModel445aa6_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackWalletModel445aa6_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackWalletModel445aa6_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModel445aa6_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackWalletModel445aa6_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackWalletModel445aa6_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModel445aa6_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackWalletModel445aa6_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackWalletModel445aa6_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = new QVector<int>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackWalletModel445aa6_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackWalletModel445aa6_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackWalletModel445aa6_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackWalletModel445aa6_HeaderDataChanged(this, orientation, first, last); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackWalletModel445aa6_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackWalletModel445aa6_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackWalletModel445aa6_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackWalletModel445aa6_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackWalletModel445aa6_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackWalletModel445aa6_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackWalletModel445aa6_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackWalletModel445aa6_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackWalletModel445aa6_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackWalletModel445aa6_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackWalletModel445aa6_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackWalletModel445aa6_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackWalletModel445aa6_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackWalletModel445aa6_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackWalletModel445aa6_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackWalletModel445aa6_ResetInternalData(this); };
	void revert() { callbackWalletModel445aa6_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackWalletModel445aa6_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackWalletModel445aa6_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackWalletModel445aa6_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackWalletModel445aa6_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModel445aa6_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackWalletModel445aa6_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackWalletModel445aa6_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModel445aa6_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackWalletModel445aa6_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackWalletModel445aa6_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackWalletModel445aa6_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackWalletModel445aa6_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackWalletModel445aa6_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackWalletModel445aa6_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackWalletModel445aa6_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackWalletModel445aa6_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackWalletModel445aa6_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackWalletModel445aa6_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackWalletModel445aa6_CustomEvent(this, event); };
	void deleteLater() { callbackWalletModel445aa6_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackWalletModel445aa6_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackWalletModel445aa6_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackWalletModel445aa6_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackWalletModel445aa6_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackWalletModel445aa6_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackWalletModel445aa6_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<QWallet445aa6*> walletsDefault() { return _wallets; };
	void setWalletsDefault(QList<QWallet445aa6*> p) { if (p != _wallets) { _wallets = p; walletsChanged(_wallets); } };
	qint32 countDefault() { return _count; };
	void setCountDefault(qint32 p) { if (p != _count) { _count = p; countChanged(_count); } };
signals:
	void rolesChanged(type378cdd roles);
	void walletsChanged(QList<QWallet445aa6*> wallets);
	void countChanged(qint32 count);
public slots:
	void addWallet(QWallet445aa6* v0) { callbackWalletModel445aa6_AddWallet(this, v0); };
	void editWallet(qint32 row, QString name, bool encryptionEnabled, quint64 sky, quint64 coinHours) { QByteArray t6ae999 = name.toUtf8(); Moc_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackWalletModel445aa6_EditWallet(this, row, namePacked, encryptionEnabled, sky, coinHours); };
	void removeWallet(qint32 row) { callbackWalletModel445aa6_RemoveWallet(this, row); };
	void loadModel(QList<QWallet445aa6*> v0) { callbackWalletModel445aa6_LoadModel(this, ({ QList<QWallet445aa6*>* tmpValue = new QList<QWallet445aa6*>(v0); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
private:
	type378cdd _roles;
	QList<QWallet445aa6*> _wallets;
	qint32 _count;
};

Q_DECLARE_METATYPE(WalletModel445aa6*)


void WalletModel445aa6_WalletModel445aa6_QRegisterMetaTypes() {
	qRegisterMetaType<QList<QObject*>>("QList<QWallet445aa6*>");
	qRegisterMetaType<type378cdd>("type378cdd");
}

void AddressesModel445aa6_AddAddress(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<AddressesModel445aa6*>(ptr), "addAddress", Q_ARG(QAddress445aa6*, static_cast<QAddress445aa6*>(v0)));
}

void AddressesModel445aa6_RemoveAddress(void* ptr, int v0)
{
	QMetaObject::invokeMethod(static_cast<AddressesModel445aa6*>(ptr), "removeAddress", Q_ARG(qint32, v0));
}

void AddressesModel445aa6_EditAddress(void* ptr, int v0, struct Moc_PackedString v1, unsigned long long v2, unsigned long long v3)
{
	QMetaObject::invokeMethod(static_cast<AddressesModel445aa6*>(ptr), "editAddress", Q_ARG(qint32, v0), Q_ARG(QString, QString::fromUtf8(v1.data, v1.len)), Q_ARG(quint64, v2), Q_ARG(quint64, v3));
}

void AddressesModel445aa6_LoadModel(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<AddressesModel445aa6*>(ptr), "loadModel", Q_ARG(QList<QAddress445aa6*>, ({ QList<QAddress445aa6*>* tmpP = static_cast<QList<QAddress445aa6*>*>(v0); QList<QAddress445aa6*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; })));
}

struct Moc_PackedList AddressesModel445aa6_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<AddressesModel445aa6*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressesModel445aa6_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<AddressesModel445aa6*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void AddressesModel445aa6_SetRoles(void* ptr, void* roles)
{
	static_cast<AddressesModel445aa6*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void AddressesModel445aa6_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<AddressesModel445aa6*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void AddressesModel445aa6_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<AddressesModel445aa6*>(ptr), static_cast<void (AddressesModel445aa6::*)(QMap<qint32, QByteArray>)>(&AddressesModel445aa6::rolesChanged), static_cast<AddressesModel445aa6*>(ptr), static_cast<void (AddressesModel445aa6::*)(QMap<qint32, QByteArray>)>(&AddressesModel445aa6::Signal_RolesChanged));
}

void AddressesModel445aa6_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressesModel445aa6*>(ptr), static_cast<void (AddressesModel445aa6::*)(QMap<qint32, QByteArray>)>(&AddressesModel445aa6::rolesChanged), static_cast<AddressesModel445aa6*>(ptr), static_cast<void (AddressesModel445aa6::*)(QMap<qint32, QByteArray>)>(&AddressesModel445aa6::Signal_RolesChanged));
}

void AddressesModel445aa6_RolesChanged(void* ptr, void* roles)
{
	static_cast<AddressesModel445aa6*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList AddressesModel445aa6_Addresses(void* ptr)
{
	return ({ QList<QAddress445aa6*>* tmpValue = new QList<QAddress445aa6*>(static_cast<AddressesModel445aa6*>(ptr)->addresses()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressesModel445aa6_AddressesDefault(void* ptr)
{
	return ({ QList<QAddress445aa6*>* tmpValue = new QList<QAddress445aa6*>(static_cast<AddressesModel445aa6*>(ptr)->addressesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void AddressesModel445aa6_SetAddresses(void* ptr, void* addresses)
{
	static_cast<AddressesModel445aa6*>(ptr)->setAddresses(({ QList<QAddress445aa6*>* tmpP = static_cast<QList<QAddress445aa6*>*>(addresses); QList<QAddress445aa6*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void AddressesModel445aa6_SetAddressesDefault(void* ptr, void* addresses)
{
	static_cast<AddressesModel445aa6*>(ptr)->setAddressesDefault(({ QList<QAddress445aa6*>* tmpP = static_cast<QList<QAddress445aa6*>*>(addresses); QList<QAddress445aa6*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void AddressesModel445aa6_ConnectAddressesChanged(void* ptr)
{
	QObject::connect(static_cast<AddressesModel445aa6*>(ptr), static_cast<void (AddressesModel445aa6::*)(QList<QAddress445aa6*>)>(&AddressesModel445aa6::addressesChanged), static_cast<AddressesModel445aa6*>(ptr), static_cast<void (AddressesModel445aa6::*)(QList<QAddress445aa6*>)>(&AddressesModel445aa6::Signal_AddressesChanged));
}

void AddressesModel445aa6_DisconnectAddressesChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressesModel445aa6*>(ptr), static_cast<void (AddressesModel445aa6::*)(QList<QAddress445aa6*>)>(&AddressesModel445aa6::addressesChanged), static_cast<AddressesModel445aa6*>(ptr), static_cast<void (AddressesModel445aa6::*)(QList<QAddress445aa6*>)>(&AddressesModel445aa6::Signal_AddressesChanged));
}

void AddressesModel445aa6_AddressesChanged(void* ptr, void* addresses)
{
	static_cast<AddressesModel445aa6*>(ptr)->addressesChanged(({ QList<QAddress445aa6*>* tmpP = static_cast<QList<QAddress445aa6*>*>(addresses); QList<QAddress445aa6*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int AddressesModel445aa6_AddressesModel445aa6_QRegisterMetaType()
{
	return qRegisterMetaType<AddressesModel445aa6*>();
}

int AddressesModel445aa6_AddressesModel445aa6_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<AddressesModel445aa6*>(const_cast<const char*>(typeName));
}

int AddressesModel445aa6_AddressesModel445aa6_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressesModel445aa6>();
#else
	return 0;
#endif
}

int AddressesModel445aa6_AddressesModel445aa6_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressesModel445aa6>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int AddressesModel445aa6_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel445aa6_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModel445aa6_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressesModel445aa6_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel445aa6_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModel445aa6_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressesModel445aa6_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel445aa6_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModel445aa6_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* AddressesModel445aa6___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel445aa6___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModel445aa6___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressesModel445aa6___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel445aa6___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModel445aa6___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int AddressesModel445aa6___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void AddressesModel445aa6___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* AddressesModel445aa6___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* AddressesModel445aa6___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModel445aa6___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* AddressesModel445aa6___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList AddressesModel445aa6___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel445aa6___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel445aa6___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* AddressesModel445aa6___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* AddressesModel445aa6___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel445aa6___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* AddressesModel445aa6___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* AddressesModel445aa6___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel445aa6___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModel445aa6___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressesModel445aa6___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel445aa6___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModel445aa6___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressesModel445aa6___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel445aa6___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModel445aa6___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressesModel445aa6___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void AddressesModel445aa6___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressesModel445aa6___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList AddressesModel445aa6___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel445aa6___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModel445aa6___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* AddressesModel445aa6___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList AddressesModel445aa6___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int AddressesModel445aa6_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel445aa6_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModel445aa6_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressesModel445aa6_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel445aa6_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModel445aa6_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* AddressesModel445aa6___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel445aa6___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressesModel445aa6___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* AddressesModel445aa6___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel445aa6___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* AddressesModel445aa6___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* AddressesModel445aa6___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel445aa6___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressesModel445aa6___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressesModel445aa6___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel445aa6___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressesModel445aa6___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressesModel445aa6___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel445aa6___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressesModel445aa6___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressesModel445aa6___loadModel_v0_atList(void* ptr, int i)
{
	return ({QAddress445aa6* tmp = static_cast<QList<QAddress445aa6*>*>(ptr)->at(i); if (i == static_cast<QList<QAddress445aa6*>*>(ptr)->size()-1) { static_cast<QList<QAddress445aa6*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel445aa6___loadModel_v0_setList(void* ptr, void* i)
{
	static_cast<QList<QAddress445aa6*>*>(ptr)->append(static_cast<QAddress445aa6*>(i));
}

void* AddressesModel445aa6___loadModel_v0_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAddress445aa6*>();
}

void* AddressesModel445aa6___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModel445aa6___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressesModel445aa6___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressesModel445aa6___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel445aa6___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModel445aa6___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressesModel445aa6___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressesModel445aa6___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel445aa6___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModel445aa6___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressesModel445aa6___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressesModel445aa6___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel445aa6___addresses_atList(void* ptr, int i)
{
	return ({QAddress445aa6* tmp = static_cast<QList<QAddress445aa6*>*>(ptr)->at(i); if (i == static_cast<QList<QAddress445aa6*>*>(ptr)->size()-1) { static_cast<QList<QAddress445aa6*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel445aa6___addresses_setList(void* ptr, void* i)
{
	static_cast<QList<QAddress445aa6*>*>(ptr)->append(static_cast<QAddress445aa6*>(i));
}

void* AddressesModel445aa6___addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAddress445aa6*>();
}

void* AddressesModel445aa6___setAddresses_addresses_atList(void* ptr, int i)
{
	return ({QAddress445aa6* tmp = static_cast<QList<QAddress445aa6*>*>(ptr)->at(i); if (i == static_cast<QList<QAddress445aa6*>*>(ptr)->size()-1) { static_cast<QList<QAddress445aa6*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel445aa6___setAddresses_addresses_setList(void* ptr, void* i)
{
	static_cast<QList<QAddress445aa6*>*>(ptr)->append(static_cast<QAddress445aa6*>(i));
}

void* AddressesModel445aa6___setAddresses_addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAddress445aa6*>();
}

void* AddressesModel445aa6___addressesChanged_addresses_atList(void* ptr, int i)
{
	return ({QAddress445aa6* tmp = static_cast<QList<QAddress445aa6*>*>(ptr)->at(i); if (i == static_cast<QList<QAddress445aa6*>*>(ptr)->size()-1) { static_cast<QList<QAddress445aa6*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel445aa6___addressesChanged_addresses_setList(void* ptr, void* i)
{
	static_cast<QList<QAddress445aa6*>*>(ptr)->append(static_cast<QAddress445aa6*>(i));
}

void* AddressesModel445aa6___addressesChanged_addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAddress445aa6*>();
}

int AddressesModel445aa6_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel445aa6_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressesModel445aa6_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int AddressesModel445aa6_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel445aa6_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressesModel445aa6_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int AddressesModel445aa6_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel445aa6_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressesModel445aa6_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* AddressesModel445aa6_NewAddressesModel(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new AddressesModel445aa6(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new AddressesModel445aa6(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new AddressesModel445aa6(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new AddressesModel445aa6(static_cast<QWindow*>(parent));
	} else {
		return new AddressesModel445aa6(static_cast<QObject*>(parent));
	}
}

void AddressesModel445aa6_DestroyAddressesModel(void* ptr)
{
	static_cast<AddressesModel445aa6*>(ptr)->~AddressesModel445aa6();
}

void AddressesModel445aa6_DestroyAddressesModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char AddressesModel445aa6_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

long long AddressesModel445aa6_FlagsDefault(void* ptr, void* index)
{
	return static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

void* AddressesModel445aa6_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* AddressesModel445aa6_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* AddressesModel445aa6_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

char AddressesModel445aa6_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char AddressesModel445aa6_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int AddressesModel445aa6_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* AddressesModel445aa6_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void AddressesModel445aa6_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char AddressesModel445aa6_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* AddressesModel445aa6_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

char AddressesModel445aa6_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char AddressesModel445aa6_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct Moc_PackedList AddressesModel445aa6_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressesModel445aa6_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel445aa6_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct Moc_PackedString AddressesModel445aa6_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t6a4f86 = static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::mimeTypes().join("¡¦!").toUtf8(); Moc_PackedString { const_cast<char*>(t6a4f86.prepend("WHITESPACE").constData()+10), t6a4f86.size()-10 }; });
}

char AddressesModel445aa6_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char AddressesModel445aa6_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* AddressesModel445aa6_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

char AddressesModel445aa6_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char AddressesModel445aa6_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void AddressesModel445aa6_ResetInternalDataDefault(void* ptr)
{
	static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::resetInternalData();
}

void AddressesModel445aa6_RevertDefault(void* ptr)
{
	static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::revert();
}

struct Moc_PackedList AddressesModel445aa6_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int AddressesModel445aa6_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char AddressesModel445aa6_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char AddressesModel445aa6_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char AddressesModel445aa6_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void AddressesModel445aa6_SortDefault(void* ptr, int column, long long order)
{
	static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* AddressesModel445aa6_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char AddressesModel445aa6_SubmitDefault(void* ptr)
{
	return static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::submit();
}

long long AddressesModel445aa6_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long AddressesModel445aa6_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::supportedDropActions();
}

void AddressesModel445aa6_ChildEventDefault(void* ptr, void* event)
{
	static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void AddressesModel445aa6_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void AddressesModel445aa6_CustomEventDefault(void* ptr, void* event)
{
	static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void AddressesModel445aa6_DeleteLaterDefault(void* ptr)
{
	static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::deleteLater();
}

void AddressesModel445aa6_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char AddressesModel445aa6_EventDefault(void* ptr, void* e)
{
	return static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char AddressesModel445aa6_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void AddressesModel445aa6_TimerEventDefault(void* ptr, void* event)
{
	static_cast<AddressesModel445aa6*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

struct Moc_PackedString QAddress445aa6_Address(void* ptr)
{
	return ({ QByteArray tbbcec3 = static_cast<QAddress445aa6*>(ptr)->address().toUtf8(); Moc_PackedString { const_cast<char*>(tbbcec3.prepend("WHITESPACE").constData()+10), tbbcec3.size()-10 }; });
}

struct Moc_PackedString QAddress445aa6_AddressDefault(void* ptr)
{
	return ({ QByteArray tbe1549 = static_cast<QAddress445aa6*>(ptr)->addressDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tbe1549.prepend("WHITESPACE").constData()+10), tbe1549.size()-10 }; });
}

void QAddress445aa6_SetAddress(void* ptr, struct Moc_PackedString address)
{
	static_cast<QAddress445aa6*>(ptr)->setAddress(QString::fromUtf8(address.data, address.len));
}

void QAddress445aa6_SetAddressDefault(void* ptr, struct Moc_PackedString address)
{
	static_cast<QAddress445aa6*>(ptr)->setAddressDefault(QString::fromUtf8(address.data, address.len));
}

void QAddress445aa6_ConnectAddressChanged(void* ptr)
{
	QObject::connect(static_cast<QAddress445aa6*>(ptr), static_cast<void (QAddress445aa6::*)(QString)>(&QAddress445aa6::addressChanged), static_cast<QAddress445aa6*>(ptr), static_cast<void (QAddress445aa6::*)(QString)>(&QAddress445aa6::Signal_AddressChanged));
}

void QAddress445aa6_DisconnectAddressChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAddress445aa6*>(ptr), static_cast<void (QAddress445aa6::*)(QString)>(&QAddress445aa6::addressChanged), static_cast<QAddress445aa6*>(ptr), static_cast<void (QAddress445aa6::*)(QString)>(&QAddress445aa6::Signal_AddressChanged));
}

void QAddress445aa6_AddressChanged(void* ptr, struct Moc_PackedString address)
{
	static_cast<QAddress445aa6*>(ptr)->addressChanged(QString::fromUtf8(address.data, address.len));
}

unsigned long long QAddress445aa6_AddressSky(void* ptr)
{
	return static_cast<QAddress445aa6*>(ptr)->addressSky();
}

unsigned long long QAddress445aa6_AddressSkyDefault(void* ptr)
{
	return static_cast<QAddress445aa6*>(ptr)->addressSkyDefault();
}

void QAddress445aa6_SetAddressSky(void* ptr, unsigned long long addressSky)
{
	static_cast<QAddress445aa6*>(ptr)->setAddressSky(addressSky);
}

void QAddress445aa6_SetAddressSkyDefault(void* ptr, unsigned long long addressSky)
{
	static_cast<QAddress445aa6*>(ptr)->setAddressSkyDefault(addressSky);
}

void QAddress445aa6_ConnectAddressSkyChanged(void* ptr)
{
	QObject::connect(static_cast<QAddress445aa6*>(ptr), static_cast<void (QAddress445aa6::*)(quint64)>(&QAddress445aa6::addressSkyChanged), static_cast<QAddress445aa6*>(ptr), static_cast<void (QAddress445aa6::*)(quint64)>(&QAddress445aa6::Signal_AddressSkyChanged));
}

void QAddress445aa6_DisconnectAddressSkyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAddress445aa6*>(ptr), static_cast<void (QAddress445aa6::*)(quint64)>(&QAddress445aa6::addressSkyChanged), static_cast<QAddress445aa6*>(ptr), static_cast<void (QAddress445aa6::*)(quint64)>(&QAddress445aa6::Signal_AddressSkyChanged));
}

void QAddress445aa6_AddressSkyChanged(void* ptr, unsigned long long addressSky)
{
	static_cast<QAddress445aa6*>(ptr)->addressSkyChanged(addressSky);
}

unsigned long long QAddress445aa6_AddressCoinHours(void* ptr)
{
	return static_cast<QAddress445aa6*>(ptr)->addressCoinHours();
}

unsigned long long QAddress445aa6_AddressCoinHoursDefault(void* ptr)
{
	return static_cast<QAddress445aa6*>(ptr)->addressCoinHoursDefault();
}

void QAddress445aa6_SetAddressCoinHours(void* ptr, unsigned long long addressCoinHours)
{
	static_cast<QAddress445aa6*>(ptr)->setAddressCoinHours(addressCoinHours);
}

void QAddress445aa6_SetAddressCoinHoursDefault(void* ptr, unsigned long long addressCoinHours)
{
	static_cast<QAddress445aa6*>(ptr)->setAddressCoinHoursDefault(addressCoinHours);
}

void QAddress445aa6_ConnectAddressCoinHoursChanged(void* ptr)
{
	QObject::connect(static_cast<QAddress445aa6*>(ptr), static_cast<void (QAddress445aa6::*)(quint64)>(&QAddress445aa6::addressCoinHoursChanged), static_cast<QAddress445aa6*>(ptr), static_cast<void (QAddress445aa6::*)(quint64)>(&QAddress445aa6::Signal_AddressCoinHoursChanged));
}

void QAddress445aa6_DisconnectAddressCoinHoursChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAddress445aa6*>(ptr), static_cast<void (QAddress445aa6::*)(quint64)>(&QAddress445aa6::addressCoinHoursChanged), static_cast<QAddress445aa6*>(ptr), static_cast<void (QAddress445aa6::*)(quint64)>(&QAddress445aa6::Signal_AddressCoinHoursChanged));
}

void QAddress445aa6_AddressCoinHoursChanged(void* ptr, unsigned long long addressCoinHours)
{
	static_cast<QAddress445aa6*>(ptr)->addressCoinHoursChanged(addressCoinHours);
}

int QAddress445aa6_QAddress445aa6_QRegisterMetaType()
{
	return qRegisterMetaType<QAddress445aa6*>();
}

int QAddress445aa6_QAddress445aa6_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QAddress445aa6*>(const_cast<const char*>(typeName));
}

int QAddress445aa6_QAddress445aa6_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QAddress445aa6>();
#else
	return 0;
#endif
}

int QAddress445aa6_QAddress445aa6_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QAddress445aa6>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QAddress445aa6___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAddress445aa6___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAddress445aa6___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QAddress445aa6___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAddress445aa6___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QAddress445aa6___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QAddress445aa6___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAddress445aa6___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAddress445aa6___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAddress445aa6___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAddress445aa6___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAddress445aa6___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAddress445aa6___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAddress445aa6___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAddress445aa6___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAddress445aa6_NewQAddress(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QAddress445aa6(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QAddress445aa6(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QAddress445aa6(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QAddress445aa6(static_cast<QWindow*>(parent));
	} else {
		return new QAddress445aa6(static_cast<QObject*>(parent));
	}
}

void QAddress445aa6_DestroyQAddress(void* ptr)
{
	static_cast<QAddress445aa6*>(ptr)->~QAddress445aa6();
}

void QAddress445aa6_DestroyQAddressDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void QAddress445aa6_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAddress445aa6*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QAddress445aa6_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAddress445aa6*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAddress445aa6_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAddress445aa6*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QAddress445aa6_DeleteLaterDefault(void* ptr)
{
	static_cast<QAddress445aa6*>(ptr)->QObject::deleteLater();
}

void QAddress445aa6_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAddress445aa6*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QAddress445aa6_EventDefault(void* ptr, void* e)
{
	return static_cast<QAddress445aa6*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QAddress445aa6_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAddress445aa6*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QAddress445aa6_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAddress445aa6*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

struct Moc_PackedString QWallet445aa6_Name(void* ptr)
{
	return ({ QByteArray tedc734 = static_cast<QWallet445aa6*>(ptr)->name().toUtf8(); Moc_PackedString { const_cast<char*>(tedc734.prepend("WHITESPACE").constData()+10), tedc734.size()-10 }; });
}

struct Moc_PackedString QWallet445aa6_NameDefault(void* ptr)
{
	return ({ QByteArray tee7a91 = static_cast<QWallet445aa6*>(ptr)->nameDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tee7a91.prepend("WHITESPACE").constData()+10), tee7a91.size()-10 }; });
}

void QWallet445aa6_SetName(void* ptr, struct Moc_PackedString name)
{
	static_cast<QWallet445aa6*>(ptr)->setName(QString::fromUtf8(name.data, name.len));
}

void QWallet445aa6_SetNameDefault(void* ptr, struct Moc_PackedString name)
{
	static_cast<QWallet445aa6*>(ptr)->setNameDefault(QString::fromUtf8(name.data, name.len));
}

void QWallet445aa6_ConnectNameChanged(void* ptr)
{
	QObject::connect(static_cast<QWallet445aa6*>(ptr), static_cast<void (QWallet445aa6::*)(QString)>(&QWallet445aa6::nameChanged), static_cast<QWallet445aa6*>(ptr), static_cast<void (QWallet445aa6::*)(QString)>(&QWallet445aa6::Signal_NameChanged));
}

void QWallet445aa6_DisconnectNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWallet445aa6*>(ptr), static_cast<void (QWallet445aa6::*)(QString)>(&QWallet445aa6::nameChanged), static_cast<QWallet445aa6*>(ptr), static_cast<void (QWallet445aa6::*)(QString)>(&QWallet445aa6::Signal_NameChanged));
}

void QWallet445aa6_NameChanged(void* ptr, struct Moc_PackedString name)
{
	static_cast<QWallet445aa6*>(ptr)->nameChanged(QString::fromUtf8(name.data, name.len));
}

int QWallet445aa6_EncryptionEnabled(void* ptr)
{
	return static_cast<QWallet445aa6*>(ptr)->encryptionEnabled();
}

int QWallet445aa6_EncryptionEnabledDefault(void* ptr)
{
	return static_cast<QWallet445aa6*>(ptr)->encryptionEnabledDefault();
}

void QWallet445aa6_SetEncryptionEnabled(void* ptr, int encryptionEnabled)
{
	static_cast<QWallet445aa6*>(ptr)->setEncryptionEnabled(encryptionEnabled);
}

void QWallet445aa6_SetEncryptionEnabledDefault(void* ptr, int encryptionEnabled)
{
	static_cast<QWallet445aa6*>(ptr)->setEncryptionEnabledDefault(encryptionEnabled);
}

void QWallet445aa6_ConnectEncryptionEnabledChanged(void* ptr)
{
	QObject::connect(static_cast<QWallet445aa6*>(ptr), static_cast<void (QWallet445aa6::*)(qint32)>(&QWallet445aa6::encryptionEnabledChanged), static_cast<QWallet445aa6*>(ptr), static_cast<void (QWallet445aa6::*)(qint32)>(&QWallet445aa6::Signal_EncryptionEnabledChanged));
}

void QWallet445aa6_DisconnectEncryptionEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWallet445aa6*>(ptr), static_cast<void (QWallet445aa6::*)(qint32)>(&QWallet445aa6::encryptionEnabledChanged), static_cast<QWallet445aa6*>(ptr), static_cast<void (QWallet445aa6::*)(qint32)>(&QWallet445aa6::Signal_EncryptionEnabledChanged));
}

void QWallet445aa6_EncryptionEnabledChanged(void* ptr, int encryptionEnabled)
{
	static_cast<QWallet445aa6*>(ptr)->encryptionEnabledChanged(encryptionEnabled);
}

unsigned long long QWallet445aa6_Sky(void* ptr)
{
	return static_cast<QWallet445aa6*>(ptr)->sky();
}

unsigned long long QWallet445aa6_SkyDefault(void* ptr)
{
	return static_cast<QWallet445aa6*>(ptr)->skyDefault();
}

void QWallet445aa6_SetSky(void* ptr, unsigned long long sky)
{
	static_cast<QWallet445aa6*>(ptr)->setSky(sky);
}

void QWallet445aa6_SetSkyDefault(void* ptr, unsigned long long sky)
{
	static_cast<QWallet445aa6*>(ptr)->setSkyDefault(sky);
}

void QWallet445aa6_ConnectSkyChanged(void* ptr)
{
	QObject::connect(static_cast<QWallet445aa6*>(ptr), static_cast<void (QWallet445aa6::*)(quint64)>(&QWallet445aa6::skyChanged), static_cast<QWallet445aa6*>(ptr), static_cast<void (QWallet445aa6::*)(quint64)>(&QWallet445aa6::Signal_SkyChanged));
}

void QWallet445aa6_DisconnectSkyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWallet445aa6*>(ptr), static_cast<void (QWallet445aa6::*)(quint64)>(&QWallet445aa6::skyChanged), static_cast<QWallet445aa6*>(ptr), static_cast<void (QWallet445aa6::*)(quint64)>(&QWallet445aa6::Signal_SkyChanged));
}

void QWallet445aa6_SkyChanged(void* ptr, unsigned long long sky)
{
	static_cast<QWallet445aa6*>(ptr)->skyChanged(sky);
}

unsigned long long QWallet445aa6_CoinHours(void* ptr)
{
	return static_cast<QWallet445aa6*>(ptr)->coinHours();
}

unsigned long long QWallet445aa6_CoinHoursDefault(void* ptr)
{
	return static_cast<QWallet445aa6*>(ptr)->coinHoursDefault();
}

void QWallet445aa6_SetCoinHours(void* ptr, unsigned long long coinHours)
{
	static_cast<QWallet445aa6*>(ptr)->setCoinHours(coinHours);
}

void QWallet445aa6_SetCoinHoursDefault(void* ptr, unsigned long long coinHours)
{
	static_cast<QWallet445aa6*>(ptr)->setCoinHoursDefault(coinHours);
}

void QWallet445aa6_ConnectCoinHoursChanged(void* ptr)
{
	QObject::connect(static_cast<QWallet445aa6*>(ptr), static_cast<void (QWallet445aa6::*)(quint64)>(&QWallet445aa6::coinHoursChanged), static_cast<QWallet445aa6*>(ptr), static_cast<void (QWallet445aa6::*)(quint64)>(&QWallet445aa6::Signal_CoinHoursChanged));
}

void QWallet445aa6_DisconnectCoinHoursChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWallet445aa6*>(ptr), static_cast<void (QWallet445aa6::*)(quint64)>(&QWallet445aa6::coinHoursChanged), static_cast<QWallet445aa6*>(ptr), static_cast<void (QWallet445aa6::*)(quint64)>(&QWallet445aa6::Signal_CoinHoursChanged));
}

void QWallet445aa6_CoinHoursChanged(void* ptr, unsigned long long coinHours)
{
	static_cast<QWallet445aa6*>(ptr)->coinHoursChanged(coinHours);
}

struct Moc_PackedString QWallet445aa6_FileName(void* ptr)
{
	return ({ QByteArray teb23e1 = static_cast<QWallet445aa6*>(ptr)->fileName().toUtf8(); Moc_PackedString { const_cast<char*>(teb23e1.prepend("WHITESPACE").constData()+10), teb23e1.size()-10 }; });
}

struct Moc_PackedString QWallet445aa6_FileNameDefault(void* ptr)
{
	return ({ QByteArray t811a22 = static_cast<QWallet445aa6*>(ptr)->fileNameDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t811a22.prepend("WHITESPACE").constData()+10), t811a22.size()-10 }; });
}

void QWallet445aa6_SetFileName(void* ptr, struct Moc_PackedString fileName)
{
	static_cast<QWallet445aa6*>(ptr)->setFileName(QString::fromUtf8(fileName.data, fileName.len));
}

void QWallet445aa6_SetFileNameDefault(void* ptr, struct Moc_PackedString fileName)
{
	static_cast<QWallet445aa6*>(ptr)->setFileNameDefault(QString::fromUtf8(fileName.data, fileName.len));
}

void QWallet445aa6_ConnectFileNameChanged(void* ptr)
{
	QObject::connect(static_cast<QWallet445aa6*>(ptr), static_cast<void (QWallet445aa6::*)(QString)>(&QWallet445aa6::fileNameChanged), static_cast<QWallet445aa6*>(ptr), static_cast<void (QWallet445aa6::*)(QString)>(&QWallet445aa6::Signal_FileNameChanged));
}

void QWallet445aa6_DisconnectFileNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWallet445aa6*>(ptr), static_cast<void (QWallet445aa6::*)(QString)>(&QWallet445aa6::fileNameChanged), static_cast<QWallet445aa6*>(ptr), static_cast<void (QWallet445aa6::*)(QString)>(&QWallet445aa6::Signal_FileNameChanged));
}

void QWallet445aa6_FileNameChanged(void* ptr, struct Moc_PackedString fileName)
{
	static_cast<QWallet445aa6*>(ptr)->fileNameChanged(QString::fromUtf8(fileName.data, fileName.len));
}

int QWallet445aa6_QWallet445aa6_QRegisterMetaType()
{
	return qRegisterMetaType<QWallet445aa6*>();
}

int QWallet445aa6_QWallet445aa6_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QWallet445aa6*>(const_cast<const char*>(typeName));
}

int QWallet445aa6_QWallet445aa6_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QWallet445aa6>();
#else
	return 0;
#endif
}

int QWallet445aa6_QWallet445aa6_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QWallet445aa6>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QWallet445aa6___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWallet445aa6___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWallet445aa6___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QWallet445aa6___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWallet445aa6___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWallet445aa6___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QWallet445aa6___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWallet445aa6___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWallet445aa6___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWallet445aa6___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWallet445aa6___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWallet445aa6___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWallet445aa6___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWallet445aa6___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWallet445aa6___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWallet445aa6_NewQWallet(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QWallet445aa6(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QWallet445aa6(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QWallet445aa6(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QWallet445aa6(static_cast<QWindow*>(parent));
	} else {
		return new QWallet445aa6(static_cast<QObject*>(parent));
	}
}

void QWallet445aa6_DestroyQWallet(void* ptr)
{
	static_cast<QWallet445aa6*>(ptr)->~QWallet445aa6();
}

void QWallet445aa6_DestroyQWalletDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void QWallet445aa6_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QWallet445aa6*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QWallet445aa6_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWallet445aa6*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWallet445aa6_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QWallet445aa6*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QWallet445aa6_DeleteLaterDefault(void* ptr)
{
	static_cast<QWallet445aa6*>(ptr)->QObject::deleteLater();
}

void QWallet445aa6_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWallet445aa6*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QWallet445aa6_EventDefault(void* ptr, void* e)
{
	return static_cast<QWallet445aa6*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QWallet445aa6_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QWallet445aa6*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWallet445aa6_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QWallet445aa6*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

void WalletModel445aa6_AddWallet(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<WalletModel445aa6*>(ptr), "addWallet", Q_ARG(QWallet445aa6*, static_cast<QWallet445aa6*>(v0)));
}

void WalletModel445aa6_EditWallet(void* ptr, int row, struct Moc_PackedString name, char encryptionEnabled, unsigned long long sky, unsigned long long coinHours)
{
	QMetaObject::invokeMethod(static_cast<WalletModel445aa6*>(ptr), "editWallet", Q_ARG(qint32, row), Q_ARG(QString, QString::fromUtf8(name.data, name.len)), Q_ARG(bool, encryptionEnabled != 0), Q_ARG(quint64, sky), Q_ARG(quint64, coinHours));
}

void WalletModel445aa6_RemoveWallet(void* ptr, int row)
{
	QMetaObject::invokeMethod(static_cast<WalletModel445aa6*>(ptr), "removeWallet", Q_ARG(qint32, row));
}

void WalletModel445aa6_LoadModel(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<WalletModel445aa6*>(ptr), "loadModel", Q_ARG(QList<QWallet445aa6*>, ({ QList<QWallet445aa6*>* tmpP = static_cast<QList<QWallet445aa6*>*>(v0); QList<QWallet445aa6*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; })));
}

struct Moc_PackedList WalletModel445aa6_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<WalletModel445aa6*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList WalletModel445aa6_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<WalletModel445aa6*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void WalletModel445aa6_SetRoles(void* ptr, void* roles)
{
	static_cast<WalletModel445aa6*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void WalletModel445aa6_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<WalletModel445aa6*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void WalletModel445aa6_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<WalletModel445aa6*>(ptr), static_cast<void (WalletModel445aa6::*)(QMap<qint32, QByteArray>)>(&WalletModel445aa6::rolesChanged), static_cast<WalletModel445aa6*>(ptr), static_cast<void (WalletModel445aa6::*)(QMap<qint32, QByteArray>)>(&WalletModel445aa6::Signal_RolesChanged));
}

void WalletModel445aa6_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<WalletModel445aa6*>(ptr), static_cast<void (WalletModel445aa6::*)(QMap<qint32, QByteArray>)>(&WalletModel445aa6::rolesChanged), static_cast<WalletModel445aa6*>(ptr), static_cast<void (WalletModel445aa6::*)(QMap<qint32, QByteArray>)>(&WalletModel445aa6::Signal_RolesChanged));
}

void WalletModel445aa6_RolesChanged(void* ptr, void* roles)
{
	static_cast<WalletModel445aa6*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList WalletModel445aa6_Wallets(void* ptr)
{
	return ({ QList<QWallet445aa6*>* tmpValue = new QList<QWallet445aa6*>(static_cast<WalletModel445aa6*>(ptr)->wallets()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList WalletModel445aa6_WalletsDefault(void* ptr)
{
	return ({ QList<QWallet445aa6*>* tmpValue = new QList<QWallet445aa6*>(static_cast<WalletModel445aa6*>(ptr)->walletsDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void WalletModel445aa6_SetWallets(void* ptr, void* wallets)
{
	static_cast<WalletModel445aa6*>(ptr)->setWallets(({ QList<QWallet445aa6*>* tmpP = static_cast<QList<QWallet445aa6*>*>(wallets); QList<QWallet445aa6*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void WalletModel445aa6_SetWalletsDefault(void* ptr, void* wallets)
{
	static_cast<WalletModel445aa6*>(ptr)->setWalletsDefault(({ QList<QWallet445aa6*>* tmpP = static_cast<QList<QWallet445aa6*>*>(wallets); QList<QWallet445aa6*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void WalletModel445aa6_ConnectWalletsChanged(void* ptr)
{
	QObject::connect(static_cast<WalletModel445aa6*>(ptr), static_cast<void (WalletModel445aa6::*)(QList<QWallet445aa6*>)>(&WalletModel445aa6::walletsChanged), static_cast<WalletModel445aa6*>(ptr), static_cast<void (WalletModel445aa6::*)(QList<QWallet445aa6*>)>(&WalletModel445aa6::Signal_WalletsChanged));
}

void WalletModel445aa6_DisconnectWalletsChanged(void* ptr)
{
	QObject::disconnect(static_cast<WalletModel445aa6*>(ptr), static_cast<void (WalletModel445aa6::*)(QList<QWallet445aa6*>)>(&WalletModel445aa6::walletsChanged), static_cast<WalletModel445aa6*>(ptr), static_cast<void (WalletModel445aa6::*)(QList<QWallet445aa6*>)>(&WalletModel445aa6::Signal_WalletsChanged));
}

void WalletModel445aa6_WalletsChanged(void* ptr, void* wallets)
{
	static_cast<WalletModel445aa6*>(ptr)->walletsChanged(({ QList<QWallet445aa6*>* tmpP = static_cast<QList<QWallet445aa6*>*>(wallets); QList<QWallet445aa6*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int WalletModel445aa6_Count(void* ptr)
{
	return static_cast<WalletModel445aa6*>(ptr)->count();
}

int WalletModel445aa6_CountDefault(void* ptr)
{
	return static_cast<WalletModel445aa6*>(ptr)->countDefault();
}

void WalletModel445aa6_SetCount(void* ptr, int count)
{
	static_cast<WalletModel445aa6*>(ptr)->setCount(count);
}

void WalletModel445aa6_SetCountDefault(void* ptr, int count)
{
	static_cast<WalletModel445aa6*>(ptr)->setCountDefault(count);
}

void WalletModel445aa6_ConnectCountChanged(void* ptr)
{
	QObject::connect(static_cast<WalletModel445aa6*>(ptr), static_cast<void (WalletModel445aa6::*)(qint32)>(&WalletModel445aa6::countChanged), static_cast<WalletModel445aa6*>(ptr), static_cast<void (WalletModel445aa6::*)(qint32)>(&WalletModel445aa6::Signal_CountChanged));
}

void WalletModel445aa6_DisconnectCountChanged(void* ptr)
{
	QObject::disconnect(static_cast<WalletModel445aa6*>(ptr), static_cast<void (WalletModel445aa6::*)(qint32)>(&WalletModel445aa6::countChanged), static_cast<WalletModel445aa6*>(ptr), static_cast<void (WalletModel445aa6::*)(qint32)>(&WalletModel445aa6::Signal_CountChanged));
}

void WalletModel445aa6_CountChanged(void* ptr, int count)
{
	static_cast<WalletModel445aa6*>(ptr)->countChanged(count);
}

int WalletModel445aa6_WalletModel445aa6_QRegisterMetaType()
{
	return qRegisterMetaType<WalletModel445aa6*>();
}

int WalletModel445aa6_WalletModel445aa6_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<WalletModel445aa6*>(const_cast<const char*>(typeName));
}

int WalletModel445aa6_WalletModel445aa6_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletModel445aa6>();
#else
	return 0;
#endif
}

int WalletModel445aa6_WalletModel445aa6_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletModel445aa6>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int WalletModel445aa6_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel445aa6_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel445aa6_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int WalletModel445aa6_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel445aa6_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel445aa6_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int WalletModel445aa6_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel445aa6_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel445aa6_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* WalletModel445aa6___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel445aa6___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel445aa6___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModel445aa6___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel445aa6___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel445aa6___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int WalletModel445aa6___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void WalletModel445aa6___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* WalletModel445aa6___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* WalletModel445aa6___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel445aa6___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* WalletModel445aa6___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList WalletModel445aa6___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel445aa6___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel445aa6___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* WalletModel445aa6___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* WalletModel445aa6___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel445aa6___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* WalletModel445aa6___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* WalletModel445aa6___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel445aa6___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel445aa6___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModel445aa6___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel445aa6___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel445aa6___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModel445aa6___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel445aa6___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel445aa6___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModel445aa6___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void WalletModel445aa6___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModel445aa6___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList WalletModel445aa6___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel445aa6___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel445aa6___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* WalletModel445aa6___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList WalletModel445aa6___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int WalletModel445aa6_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel445aa6_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel445aa6_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int WalletModel445aa6_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel445aa6_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel445aa6_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* WalletModel445aa6___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel445aa6___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModel445aa6___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* WalletModel445aa6___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel445aa6___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* WalletModel445aa6___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* WalletModel445aa6___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel445aa6___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModel445aa6___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletModel445aa6___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel445aa6___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModel445aa6___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletModel445aa6___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel445aa6___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModel445aa6___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletModel445aa6___loadModel_v0_atList(void* ptr, int i)
{
	return ({QWallet445aa6* tmp = static_cast<QList<QWallet445aa6*>*>(ptr)->at(i); if (i == static_cast<QList<QWallet445aa6*>*>(ptr)->size()-1) { static_cast<QList<QWallet445aa6*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel445aa6___loadModel_v0_setList(void* ptr, void* i)
{
	static_cast<QList<QWallet445aa6*>*>(ptr)->append(static_cast<QWallet445aa6*>(i));
}

void* WalletModel445aa6___loadModel_v0_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWallet445aa6*>();
}

void* WalletModel445aa6___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel445aa6___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModel445aa6___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList WalletModel445aa6___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel445aa6___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel445aa6___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModel445aa6___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList WalletModel445aa6___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel445aa6___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel445aa6___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModel445aa6___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList WalletModel445aa6___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel445aa6___wallets_atList(void* ptr, int i)
{
	return ({QWallet445aa6* tmp = static_cast<QList<QWallet445aa6*>*>(ptr)->at(i); if (i == static_cast<QList<QWallet445aa6*>*>(ptr)->size()-1) { static_cast<QList<QWallet445aa6*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel445aa6___wallets_setList(void* ptr, void* i)
{
	static_cast<QList<QWallet445aa6*>*>(ptr)->append(static_cast<QWallet445aa6*>(i));
}

void* WalletModel445aa6___wallets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWallet445aa6*>();
}

void* WalletModel445aa6___setWallets_wallets_atList(void* ptr, int i)
{
	return ({QWallet445aa6* tmp = static_cast<QList<QWallet445aa6*>*>(ptr)->at(i); if (i == static_cast<QList<QWallet445aa6*>*>(ptr)->size()-1) { static_cast<QList<QWallet445aa6*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel445aa6___setWallets_wallets_setList(void* ptr, void* i)
{
	static_cast<QList<QWallet445aa6*>*>(ptr)->append(static_cast<QWallet445aa6*>(i));
}

void* WalletModel445aa6___setWallets_wallets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWallet445aa6*>();
}

void* WalletModel445aa6___walletsChanged_wallets_atList(void* ptr, int i)
{
	return ({QWallet445aa6* tmp = static_cast<QList<QWallet445aa6*>*>(ptr)->at(i); if (i == static_cast<QList<QWallet445aa6*>*>(ptr)->size()-1) { static_cast<QList<QWallet445aa6*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel445aa6___walletsChanged_wallets_setList(void* ptr, void* i)
{
	static_cast<QList<QWallet445aa6*>*>(ptr)->append(static_cast<QWallet445aa6*>(i));
}

void* WalletModel445aa6___walletsChanged_wallets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWallet445aa6*>();
}

int WalletModel445aa6_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel445aa6_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* WalletModel445aa6_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int WalletModel445aa6_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel445aa6_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* WalletModel445aa6_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int WalletModel445aa6_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel445aa6_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* WalletModel445aa6_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* WalletModel445aa6_NewWalletModel(void* parent)
{
	if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new WalletModel445aa6(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new WalletModel445aa6(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new WalletModel445aa6(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new WalletModel445aa6(static_cast<QWindow*>(parent));
	} else {
		return new WalletModel445aa6(static_cast<QObject*>(parent));
	}
}

void WalletModel445aa6_DestroyWalletModel(void* ptr)
{
	static_cast<WalletModel445aa6*>(ptr)->~WalletModel445aa6();
}

void WalletModel445aa6_DestroyWalletModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char WalletModel445aa6_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

long long WalletModel445aa6_FlagsDefault(void* ptr, void* index)
{
	return static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

void* WalletModel445aa6_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* WalletModel445aa6_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* WalletModel445aa6_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

char WalletModel445aa6_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char WalletModel445aa6_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int WalletModel445aa6_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* WalletModel445aa6_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void WalletModel445aa6_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char WalletModel445aa6_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* WalletModel445aa6_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

char WalletModel445aa6_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char WalletModel445aa6_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct Moc_PackedList WalletModel445aa6_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList WalletModel445aa6_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel445aa6_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct Moc_PackedString WalletModel445aa6_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t2c50d2 = static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::mimeTypes().join("¡¦!").toUtf8(); Moc_PackedString { const_cast<char*>(t2c50d2.prepend("WHITESPACE").constData()+10), t2c50d2.size()-10 }; });
}

char WalletModel445aa6_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char WalletModel445aa6_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* WalletModel445aa6_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

char WalletModel445aa6_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char WalletModel445aa6_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void WalletModel445aa6_ResetInternalDataDefault(void* ptr)
{
	static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::resetInternalData();
}

void WalletModel445aa6_RevertDefault(void* ptr)
{
	static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::revert();
}

struct Moc_PackedList WalletModel445aa6_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int WalletModel445aa6_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char WalletModel445aa6_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char WalletModel445aa6_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char WalletModel445aa6_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void WalletModel445aa6_SortDefault(void* ptr, int column, long long order)
{
	static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* WalletModel445aa6_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char WalletModel445aa6_SubmitDefault(void* ptr)
{
	return static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::submit();
}

long long WalletModel445aa6_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long WalletModel445aa6_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::supportedDropActions();
}

void WalletModel445aa6_ChildEventDefault(void* ptr, void* event)
{
	static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void WalletModel445aa6_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void WalletModel445aa6_CustomEventDefault(void* ptr, void* event)
{
	static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void WalletModel445aa6_DeleteLaterDefault(void* ptr)
{
	static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::deleteLater();
}

void WalletModel445aa6_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char WalletModel445aa6_EventDefault(void* ptr, void* e)
{
	return static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char WalletModel445aa6_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void WalletModel445aa6_TimerEventDefault(void* ptr, void* event)
{
	static_cast<WalletModel445aa6*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
