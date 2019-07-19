

#define protected public
#define private public

#include "moc.h"
#include "_cgo_export.h"

#include <QAbstractItemModel>
#include <QAbstractListModel>
#include <QByteArray>
#include <QChildEvent>
#include <QEvent>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHash>
#include <QLayout>
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
#include <QWidget>
#include <QWindow>

#ifdef QT_QML_LIB
	#include <QQmlEngine>
#endif


typedef QMap<qint32, QByteArray> type378cdd;
typedef QMap<qint32, QByteArray> type378cdd;
class WalletManager268d39: public QObject
{
Q_OBJECT
public:
	WalletManager268d39(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");WalletManager268d39_WalletManager268d39_QRegisterMetaType();WalletManager268d39_WalletManager268d39_QRegisterMetaTypes();callbackWalletManager268d39_Constructor(this);};
	 ~WalletManager268d39() { callbackWalletManager268d39_DestroyWalletManager(this); };
	void childEvent(QChildEvent * event) { callbackWalletManager268d39_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackWalletManager268d39_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackWalletManager268d39_CustomEvent(this, event); };
	void deleteLater() { callbackWalletManager268d39_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackWalletManager268d39_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackWalletManager268d39_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackWalletManager268d39_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackWalletManager268d39_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackWalletManager268d39_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackWalletManager268d39_TimerEvent(this, event); };
signals:
public slots:
	void createEncryptedWallet(WalletModel268d39* walletM, QString seed, QString label, QString password, qint32 scanN) { QByteArray t92713d = seed.toUtf8(); Moc_PackedString seedPacked = { const_cast<char*>(t92713d.prepend("WHITESPACE").constData()+10), t92713d.size()-10 };QByteArray t64c653 = label.toUtf8(); Moc_PackedString labelPacked = { const_cast<char*>(t64c653.prepend("WHITESPACE").constData()+10), t64c653.size()-10 };QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackWalletManager268d39_CreateEncryptedWallet(this, walletM, seedPacked, labelPacked, passwordPacked, scanN); };
	void createUnencryptedWallet(WalletModel268d39* walletM, QString seed, QString label, qint32 scanN) { QByteArray t92713d = seed.toUtf8(); Moc_PackedString seedPacked = { const_cast<char*>(t92713d.prepend("WHITESPACE").constData()+10), t92713d.size()-10 };QByteArray t64c653 = label.toUtf8(); Moc_PackedString labelPacked = { const_cast<char*>(t64c653.prepend("WHITESPACE").constData()+10), t64c653.size()-10 };callbackWalletManager268d39_CreateUnencryptedWallet(this, walletM, seedPacked, labelPacked, scanN); };
	QString getNewSeed(qint32 entropy) { return ({ Moc_PackedString tempVal = callbackWalletManager268d39_GetNewSeed(this, entropy); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	qint32 verifySeed(QString seed) { QByteArray t92713d = seed.toUtf8(); Moc_PackedString seedPacked = { const_cast<char*>(t92713d.prepend("WHITESPACE").constData()+10), t92713d.size()-10 };return callbackWalletManager268d39_VerifySeed(this, seedPacked); };
	void newWalletAddress(AddressesModel268d39* addressM, QString id, qint32 n, QString password) { QByteArray t87ea5d = id.toUtf8(); Moc_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackWalletManager268d39_NewWalletAddress(this, addressM, idPacked, n, passwordPacked); };
	void encryptWallet(WalletModel268d39* walletM, QString id, QString password) { QByteArray t87ea5d = id.toUtf8(); Moc_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackWalletManager268d39_EncryptWallet(this, walletM, idPacked, passwordPacked); };
	void decryptWallet(WalletModel268d39* walletM, QString id, QString password) { QByteArray t87ea5d = id.toUtf8(); Moc_PackedString idPacked = { const_cast<char*>(t87ea5d.prepend("WHITESPACE").constData()+10), t87ea5d.size()-10 };QByteArray t5baa61 = password.toUtf8(); Moc_PackedString passwordPacked = { const_cast<char*>(t5baa61.prepend("WHITESPACE").constData()+10), t5baa61.size()-10 };callbackWalletManager268d39_DecryptWallet(this, walletM, idPacked, passwordPacked); };
private:
};

Q_DECLARE_METATYPE(WalletManager268d39*)


void WalletManager268d39_WalletManager268d39_QRegisterMetaTypes() {
}

class WalletModel268d39: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<QWallet268d39*> wallets READ wallets WRITE setWallets NOTIFY walletsChanged)
public:
	WalletModel268d39(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");WalletModel268d39_WalletModel268d39_QRegisterMetaType();WalletModel268d39_WalletModel268d39_QRegisterMetaTypes();callbackWalletModel268d39_Constructor(this);};
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbackWalletModel268d39_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbackWalletModel268d39_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbackWalletModel268d39_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<QWallet268d39*> wallets() { return ({ QList<QWallet268d39*>* tmpP = static_cast<QList<QWallet268d39*>*>(callbackWalletModel268d39_Wallets(this)); QList<QWallet268d39*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setWallets(QList<QWallet268d39*> wallets) { callbackWalletModel268d39_SetWallets(this, ({ QList<QWallet268d39*>* tmpValue = new QList<QWallet268d39*>(wallets); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_WalletsChanged(QList<QWallet268d39*> wallets) { callbackWalletModel268d39_WalletsChanged(this, ({ QList<QWallet268d39*>* tmpValue = new QList<QWallet268d39*>(wallets); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	 ~WalletModel268d39() { callbackWalletModel268d39_DestroyWalletModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackWalletModel268d39_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackWalletModel268d39_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackWalletModel268d39_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackWalletModel268d39_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackWalletModel268d39_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackWalletModel268d39_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackWalletModel268d39_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackWalletModel268d39_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackWalletModel268d39_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackWalletModel268d39_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModel268d39_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackWalletModel268d39_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackWalletModel268d39_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModel268d39_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackWalletModel268d39_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackWalletModel268d39_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = new QVector<int>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackWalletModel268d39_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackWalletModel268d39_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackWalletModel268d39_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackWalletModel268d39_HeaderDataChanged(this, orientation, first, last); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackWalletModel268d39_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackWalletModel268d39_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackWalletModel268d39_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackWalletModel268d39_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackWalletModel268d39_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackWalletModel268d39_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackWalletModel268d39_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackWalletModel268d39_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackWalletModel268d39_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackWalletModel268d39_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackWalletModel268d39_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackWalletModel268d39_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackWalletModel268d39_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackWalletModel268d39_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackWalletModel268d39_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackWalletModel268d39_ResetInternalData(this); };
	void revert() { callbackWalletModel268d39_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackWalletModel268d39_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackWalletModel268d39_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackWalletModel268d39_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackWalletModel268d39_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModel268d39_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackWalletModel268d39_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackWalletModel268d39_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackWalletModel268d39_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackWalletModel268d39_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackWalletModel268d39_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackWalletModel268d39_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackWalletModel268d39_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackWalletModel268d39_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackWalletModel268d39_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackWalletModel268d39_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackWalletModel268d39_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackWalletModel268d39_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackWalletModel268d39_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackWalletModel268d39_CustomEvent(this, event); };
	void deleteLater() { callbackWalletModel268d39_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackWalletModel268d39_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackWalletModel268d39_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackWalletModel268d39_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackWalletModel268d39_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackWalletModel268d39_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackWalletModel268d39_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<QWallet268d39*> walletsDefault() { return _wallets; };
	void setWalletsDefault(QList<QWallet268d39*> p) { if (p != _wallets) { _wallets = p; walletsChanged(_wallets); } };
signals:
	void rolesChanged(type378cdd roles);
	void walletsChanged(QList<QWallet268d39*> wallets);
public slots:
	void addWallet(QWallet268d39* v0) { callbackWalletModel268d39_AddWallet(this, v0); };
	void editWallet(qint32 row, QString name, bool encryptionEnabled, qint32 sky, qint32 coinHours) { QByteArray t6ae999 = name.toUtf8(); Moc_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackWalletModel268d39_EditWallet(this, row, namePacked, encryptionEnabled, sky, coinHours); };
	void removeWallet(qint32 row) { callbackWalletModel268d39_RemoveWallet(this, row); };
	void loadModel() { callbackWalletModel268d39_LoadModel(this); };
private:
	type378cdd _roles;
	QList<QWallet268d39*> _wallets;
};

Q_DECLARE_METATYPE(WalletModel268d39*)


void WalletModel268d39_WalletModel268d39_QRegisterMetaTypes() {
	qRegisterMetaType<type378cdd>("type378cdd");
	qRegisterMetaType<QList<QObject*>>("QList<QWallet268d39*>");
}

class AddressesModel268d39: public QAbstractListModel
{
Q_OBJECT
Q_PROPERTY(type378cdd roles READ roles WRITE setRoles NOTIFY rolesChanged)
Q_PROPERTY(QList<QAddress268d39*> addresses READ addresses WRITE setAddresses NOTIFY addressesChanged)
Q_PROPERTY(qint32 loaded READ loaded WRITE setLoaded NOTIFY loadedChanged)
public:
	AddressesModel268d39(QObject *parent = Q_NULLPTR) : QAbstractListModel(parent) {qRegisterMetaType<quintptr>("quintptr");AddressesModel268d39_AddressesModel268d39_QRegisterMetaType();AddressesModel268d39_AddressesModel268d39_QRegisterMetaTypes();callbackAddressesModel268d39_Constructor(this);};
	type378cdd roles() { return ({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(callbackAddressesModel268d39_Roles(this)); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void setRoles(type378cdd roles) { callbackAddressesModel268d39_SetRoles(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_RolesChanged(type378cdd roles) { callbackAddressesModel268d39_RolesChanged(this, ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	QList<QAddress268d39*> addresses() { return ({ QList<QAddress268d39*>* tmpP = static_cast<QList<QAddress268d39*>*>(callbackAddressesModel268d39_Addresses(this)); QList<QAddress268d39*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void setAddresses(QList<QAddress268d39*> addresses) { callbackAddressesModel268d39_SetAddresses(this, ({ QList<QAddress268d39*>* tmpValue = new QList<QAddress268d39*>(addresses); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void Signal_AddressesChanged(QList<QAddress268d39*> addresses) { callbackAddressesModel268d39_AddressesChanged(this, ({ QList<QAddress268d39*>* tmpValue = new QList<QAddress268d39*>(addresses); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	qint32 loaded() { return callbackAddressesModel268d39_Loaded(this); };
	void setLoaded(qint32 loaded) { callbackAddressesModel268d39_SetLoaded(this, loaded); };
	void Signal_LoadedChanged(qint32 loaded) { callbackAddressesModel268d39_LoadedChanged(this, loaded); };
	 ~AddressesModel268d39() { callbackAddressesModel268d39_DestroyAddressesModel(this); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackAddressesModel268d39_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackAddressesModel268d39_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackAddressesModel268d39_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackAddressesModel268d39_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackAddressesModel268d39_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackAddressesModel268d39_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackAddressesModel268d39_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackAddressesModel268d39_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackAddressesModel268d39_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackAddressesModel268d39_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackAddressesModel268d39_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackAddressesModel268d39_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackAddressesModel268d39_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackAddressesModel268d39_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackAddressesModel268d39_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackAddressesModel268d39_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue = new QVector<int>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackAddressesModel268d39_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackAddressesModel268d39_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackAddressesModel268d39_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackAddressesModel268d39_HeaderDataChanged(this, orientation, first, last); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackAddressesModel268d39_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackAddressesModel268d39_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackAddressesModel268d39_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackAddressesModel268d39_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackAddressesModel268d39_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValue = new QList<QPersistentModelIndex>(parents); Moc_PackedList { tmpValue, tmpValue->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackAddressesModel268d39_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackAddressesModel268d39_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(indexes); Moc_PackedList { tmpValue, tmpValue->size() }; }))); };
	QStringList mimeTypes() const { return ({ Moc_PackedString tempVal = callbackAddressesModel268d39_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackAddressesModel268d39_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackAddressesModel268d39_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackAddressesModel268d39_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackAddressesModel268d39_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackAddressesModel268d39_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackAddressesModel268d39_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackAddressesModel268d39_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackAddressesModel268d39_ResetInternalData(this); };
	void revert() { callbackAddressesModel268d39_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackAddressesModel268d39_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackAddressesModel268d39_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackAddressesModel268d39_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackAddressesModel268d39_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackAddressesModel268d39_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackAddressesModel268d39_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackAddressesModel268d39_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackAddressesModel268d39_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackAddressesModel268d39_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackAddressesModel268d39_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackAddressesModel268d39_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(roles); Moc_PackedList { tmpValue, tmpValue->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackAddressesModel268d39_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackAddressesModel268d39_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackAddressesModel268d39_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackAddressesModel268d39_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackAddressesModel268d39_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackAddressesModel268d39_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackAddressesModel268d39_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackAddressesModel268d39_CustomEvent(this, event); };
	void deleteLater() { callbackAddressesModel268d39_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackAddressesModel268d39_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackAddressesModel268d39_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackAddressesModel268d39_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackAddressesModel268d39_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackAddressesModel268d39_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackAddressesModel268d39_TimerEvent(this, event); };
	type378cdd rolesDefault() { return _roles; };
	void setRolesDefault(type378cdd p) { if (p != _roles) { _roles = p; rolesChanged(_roles); } };
	QList<QAddress268d39*> addressesDefault() { return _addresses; };
	void setAddressesDefault(QList<QAddress268d39*> p) { if (p != _addresses) { _addresses = p; addressesChanged(_addresses); } };
	qint32 loadedDefault() { return _loaded; };
	void setLoadedDefault(qint32 p) { if (p != _loaded) { _loaded = p; loadedChanged(_loaded); } };
signals:
	void rolesChanged(type378cdd roles);
	void addressesChanged(QList<QAddress268d39*> addresses);
	void loadedChanged(qint32 loaded);
public slots:
	void addAddress(QAddress268d39* v0) { callbackAddressesModel268d39_AddAddress(this, v0); };
	void removeAddress(qint32 v0) { callbackAddressesModel268d39_RemoveAddress(this, v0); };
	void editAddress(qint32 v0, QString v1, qint32 v2, qint32 v3) { QByteArray t5a6df7 = v1.toUtf8(); Moc_PackedString v1Packed = { const_cast<char*>(t5a6df7.prepend("WHITESPACE").constData()+10), t5a6df7.size()-10 };callbackAddressesModel268d39_EditAddress(this, v0, v1Packed, v2, v3); };
	void loadModel(QString v0) { QByteArray tea1dd7 = v0.toUtf8(); Moc_PackedString v0Packed = { const_cast<char*>(tea1dd7.prepend("WHITESPACE").constData()+10), tea1dd7.size()-10 };callbackAddressesModel268d39_LoadModel(this, v0Packed); };
private:
	type378cdd _roles;
	QList<QAddress268d39*> _addresses;
	qint32 _loaded;
};

Q_DECLARE_METATYPE(AddressesModel268d39*)


void AddressesModel268d39_AddressesModel268d39_QRegisterMetaTypes() {
	qRegisterMetaType<type378cdd>("type378cdd");
	qRegisterMetaType<QList<QObject*>>("QList<QAddress268d39*>");
}

class QAddress268d39: public QObject
{
Q_OBJECT
Q_PROPERTY(QString address READ address WRITE setAddress NOTIFY addressChanged)
Q_PROPERTY(qint32 addressSky READ addressSky WRITE setAddressSky NOTIFY addressSkyChanged)
Q_PROPERTY(qint32 addressCoinHours READ addressCoinHours WRITE setAddressCoinHours NOTIFY addressCoinHoursChanged)
public:
	QAddress268d39(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QAddress268d39_QAddress268d39_QRegisterMetaType();QAddress268d39_QAddress268d39_QRegisterMetaTypes();callbackQAddress268d39_Constructor(this);};
	QString address() { return ({ Moc_PackedString tempVal = callbackQAddress268d39_Address(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setAddress(QString address) { QByteArray tc66218 = address.toUtf8(); Moc_PackedString addressPacked = { const_cast<char*>(tc66218.prepend("WHITESPACE").constData()+10), tc66218.size()-10 };callbackQAddress268d39_SetAddress(this, addressPacked); };
	void Signal_AddressChanged(QString address) { QByteArray tc66218 = address.toUtf8(); Moc_PackedString addressPacked = { const_cast<char*>(tc66218.prepend("WHITESPACE").constData()+10), tc66218.size()-10 };callbackQAddress268d39_AddressChanged(this, addressPacked); };
	qint32 addressSky() { return callbackQAddress268d39_AddressSky(this); };
	void setAddressSky(qint32 addressSky) { callbackQAddress268d39_SetAddressSky(this, addressSky); };
	void Signal_AddressSkyChanged(qint32 addressSky) { callbackQAddress268d39_AddressSkyChanged(this, addressSky); };
	qint32 addressCoinHours() { return callbackQAddress268d39_AddressCoinHours(this); };
	void setAddressCoinHours(qint32 addressCoinHours) { callbackQAddress268d39_SetAddressCoinHours(this, addressCoinHours); };
	void Signal_AddressCoinHoursChanged(qint32 addressCoinHours) { callbackQAddress268d39_AddressCoinHoursChanged(this, addressCoinHours); };
	 ~QAddress268d39() { callbackQAddress268d39_DestroyQAddress(this); };
	void childEvent(QChildEvent * event) { callbackQAddress268d39_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQAddress268d39_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQAddress268d39_CustomEvent(this, event); };
	void deleteLater() { callbackQAddress268d39_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQAddress268d39_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQAddress268d39_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQAddress268d39_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQAddress268d39_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQAddress268d39_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQAddress268d39_TimerEvent(this, event); };
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

Q_DECLARE_METATYPE(QAddress268d39*)


void QAddress268d39_QAddress268d39_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

class QWallet268d39: public QObject
{
Q_OBJECT
Q_PROPERTY(QString name READ name WRITE setName NOTIFY nameChanged)
Q_PROPERTY(qint32 encryptionEnabled READ encryptionEnabled WRITE setEncryptionEnabled NOTIFY encryptionEnabledChanged)
Q_PROPERTY(qint32 sky READ sky WRITE setSky NOTIFY skyChanged)
Q_PROPERTY(qint32 coinHours READ coinHours WRITE setCoinHours NOTIFY coinHoursChanged)
Q_PROPERTY(QString fileName READ fileName WRITE setFileName NOTIFY fileNameChanged)
public:
	QWallet268d39(QObject *parent = Q_NULLPTR) : QObject(parent) {qRegisterMetaType<quintptr>("quintptr");QWallet268d39_QWallet268d39_QRegisterMetaType();QWallet268d39_QWallet268d39_QRegisterMetaTypes();callbackQWallet268d39_Constructor(this);};
	QString name() { return ({ Moc_PackedString tempVal = callbackQWallet268d39_Name(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setName(QString name) { QByteArray t6ae999 = name.toUtf8(); Moc_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQWallet268d39_SetName(this, namePacked); };
	void Signal_NameChanged(QString name) { QByteArray t6ae999 = name.toUtf8(); Moc_PackedString namePacked = { const_cast<char*>(t6ae999.prepend("WHITESPACE").constData()+10), t6ae999.size()-10 };callbackQWallet268d39_NameChanged(this, namePacked); };
	qint32 encryptionEnabled() { return callbackQWallet268d39_EncryptionEnabled(this); };
	void setEncryptionEnabled(qint32 encryptionEnabled) { callbackQWallet268d39_SetEncryptionEnabled(this, encryptionEnabled); };
	void Signal_EncryptionEnabledChanged(qint32 encryptionEnabled) { callbackQWallet268d39_EncryptionEnabledChanged(this, encryptionEnabled); };
	qint32 sky() { return callbackQWallet268d39_Sky(this); };
	void setSky(qint32 sky) { callbackQWallet268d39_SetSky(this, sky); };
	void Signal_SkyChanged(qint32 sky) { callbackQWallet268d39_SkyChanged(this, sky); };
	qint32 coinHours() { return callbackQWallet268d39_CoinHours(this); };
	void setCoinHours(qint32 coinHours) { callbackQWallet268d39_SetCoinHours(this, coinHours); };
	void Signal_CoinHoursChanged(qint32 coinHours) { callbackQWallet268d39_CoinHoursChanged(this, coinHours); };
	QString fileName() { return ({ Moc_PackedString tempVal = callbackQWallet268d39_FileName(this); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void setFileName(QString fileName) { QByteArray td83e09 = fileName.toUtf8(); Moc_PackedString fileNamePacked = { const_cast<char*>(td83e09.prepend("WHITESPACE").constData()+10), td83e09.size()-10 };callbackQWallet268d39_SetFileName(this, fileNamePacked); };
	void Signal_FileNameChanged(QString fileName) { QByteArray td83e09 = fileName.toUtf8(); Moc_PackedString fileNamePacked = { const_cast<char*>(td83e09.prepend("WHITESPACE").constData()+10), td83e09.size()-10 };callbackQWallet268d39_FileNameChanged(this, fileNamePacked); };
	 ~QWallet268d39() { callbackQWallet268d39_DestroyQWallet(this); };
	void childEvent(QChildEvent * event) { callbackQWallet268d39_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWallet268d39_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWallet268d39_CustomEvent(this, event); };
	void deleteLater() { callbackQWallet268d39_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWallet268d39_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWallet268d39_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQWallet268d39_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWallet268d39_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); Moc_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWallet268d39_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWallet268d39_TimerEvent(this, event); };
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

Q_DECLARE_METATYPE(QWallet268d39*)


void QWallet268d39_QWallet268d39_QRegisterMetaTypes() {
	qRegisterMetaType<QString>();
}

void WalletModel268d39_AddWallet(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<WalletModel268d39*>(ptr), "addWallet", Q_ARG(QWallet268d39*, static_cast<QWallet268d39*>(v0)));
}

void WalletModel268d39_EditWallet(void* ptr, int row, struct Moc_PackedString name, char encryptionEnabled, int sky, int coinHours)
{
	QMetaObject::invokeMethod(static_cast<WalletModel268d39*>(ptr), "editWallet", Q_ARG(qint32, row), Q_ARG(QString, QString::fromUtf8(name.data, name.len)), Q_ARG(bool, encryptionEnabled != 0), Q_ARG(qint32, sky), Q_ARG(qint32, coinHours));
}

void WalletModel268d39_RemoveWallet(void* ptr, int row)
{
	QMetaObject::invokeMethod(static_cast<WalletModel268d39*>(ptr), "removeWallet", Q_ARG(qint32, row));
}

void WalletModel268d39_LoadModel(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<WalletModel268d39*>(ptr), "loadModel");
}

struct Moc_PackedList WalletModel268d39_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<WalletModel268d39*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList WalletModel268d39_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<WalletModel268d39*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void WalletModel268d39_SetRoles(void* ptr, void* roles)
{
	static_cast<WalletModel268d39*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void WalletModel268d39_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<WalletModel268d39*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void WalletModel268d39_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<WalletModel268d39*>(ptr), static_cast<void (WalletModel268d39::*)(QMap<qint32, QByteArray>)>(&WalletModel268d39::rolesChanged), static_cast<WalletModel268d39*>(ptr), static_cast<void (WalletModel268d39::*)(QMap<qint32, QByteArray>)>(&WalletModel268d39::Signal_RolesChanged));
}

void WalletModel268d39_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<WalletModel268d39*>(ptr), static_cast<void (WalletModel268d39::*)(QMap<qint32, QByteArray>)>(&WalletModel268d39::rolesChanged), static_cast<WalletModel268d39*>(ptr), static_cast<void (WalletModel268d39::*)(QMap<qint32, QByteArray>)>(&WalletModel268d39::Signal_RolesChanged));
}

void WalletModel268d39_RolesChanged(void* ptr, void* roles)
{
	static_cast<WalletModel268d39*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList WalletModel268d39_Wallets(void* ptr)
{
	return ({ QList<QWallet268d39*>* tmpValue = new QList<QWallet268d39*>(static_cast<WalletModel268d39*>(ptr)->wallets()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList WalletModel268d39_WalletsDefault(void* ptr)
{
	return ({ QList<QWallet268d39*>* tmpValue = new QList<QWallet268d39*>(static_cast<WalletModel268d39*>(ptr)->walletsDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void WalletModel268d39_SetWallets(void* ptr, void* wallets)
{
	static_cast<WalletModel268d39*>(ptr)->setWallets(({ QList<QWallet268d39*>* tmpP = static_cast<QList<QWallet268d39*>*>(wallets); QList<QWallet268d39*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void WalletModel268d39_SetWalletsDefault(void* ptr, void* wallets)
{
	static_cast<WalletModel268d39*>(ptr)->setWalletsDefault(({ QList<QWallet268d39*>* tmpP = static_cast<QList<QWallet268d39*>*>(wallets); QList<QWallet268d39*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void WalletModel268d39_ConnectWalletsChanged(void* ptr)
{
	QObject::connect(static_cast<WalletModel268d39*>(ptr), static_cast<void (WalletModel268d39::*)(QList<QWallet268d39*>)>(&WalletModel268d39::walletsChanged), static_cast<WalletModel268d39*>(ptr), static_cast<void (WalletModel268d39::*)(QList<QWallet268d39*>)>(&WalletModel268d39::Signal_WalletsChanged));
}

void WalletModel268d39_DisconnectWalletsChanged(void* ptr)
{
	QObject::disconnect(static_cast<WalletModel268d39*>(ptr), static_cast<void (WalletModel268d39::*)(QList<QWallet268d39*>)>(&WalletModel268d39::walletsChanged), static_cast<WalletModel268d39*>(ptr), static_cast<void (WalletModel268d39::*)(QList<QWallet268d39*>)>(&WalletModel268d39::Signal_WalletsChanged));
}

void WalletModel268d39_WalletsChanged(void* ptr, void* wallets)
{
	static_cast<WalletModel268d39*>(ptr)->walletsChanged(({ QList<QWallet268d39*>* tmpP = static_cast<QList<QWallet268d39*>*>(wallets); QList<QWallet268d39*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int WalletModel268d39_WalletModel268d39_QRegisterMetaType()
{
	return qRegisterMetaType<WalletModel268d39*>();
}

int WalletModel268d39_WalletModel268d39_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<WalletModel268d39*>(const_cast<const char*>(typeName));
}

int WalletModel268d39_WalletModel268d39_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletModel268d39>();
#else
	return 0;
#endif
}

int WalletModel268d39_WalletModel268d39_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletModel268d39>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int WalletModel268d39_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel268d39_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel268d39_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int WalletModel268d39_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel268d39_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel268d39_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int WalletModel268d39_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel268d39_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel268d39_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* WalletModel268d39___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel268d39___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel268d39___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModel268d39___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel268d39___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel268d39___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int WalletModel268d39___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void WalletModel268d39___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* WalletModel268d39___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* WalletModel268d39___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel268d39___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* WalletModel268d39___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList WalletModel268d39___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel268d39___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel268d39___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* WalletModel268d39___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* WalletModel268d39___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel268d39___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* WalletModel268d39___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* WalletModel268d39___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel268d39___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel268d39___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModel268d39___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel268d39___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel268d39___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModel268d39___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel268d39___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* WalletModel268d39___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* WalletModel268d39___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void WalletModel268d39___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModel268d39___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList WalletModel268d39___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel268d39___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel268d39___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* WalletModel268d39___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList WalletModel268d39___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int WalletModel268d39_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel268d39_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel268d39_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int WalletModel268d39_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel268d39_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* WalletModel268d39_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* WalletModel268d39___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel268d39___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModel268d39___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* WalletModel268d39___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletModel268d39___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* WalletModel268d39___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* WalletModel268d39___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel268d39___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModel268d39___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletModel268d39___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel268d39___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModel268d39___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletModel268d39___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel268d39___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletModel268d39___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletModel268d39___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel268d39___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModel268d39___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList WalletModel268d39___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel268d39___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel268d39___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModel268d39___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList WalletModel268d39___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel268d39___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void WalletModel268d39___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* WalletModel268d39___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList WalletModel268d39___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel268d39___wallets_atList(void* ptr, int i)
{
	return ({QWallet268d39* tmp = static_cast<QList<QWallet268d39*>*>(ptr)->at(i); if (i == static_cast<QList<QWallet268d39*>*>(ptr)->size()-1) { static_cast<QList<QWallet268d39*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel268d39___wallets_setList(void* ptr, void* i)
{
	static_cast<QList<QWallet268d39*>*>(ptr)->append(static_cast<QWallet268d39*>(i));
}

void* WalletModel268d39___wallets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWallet268d39*>();
}

void* WalletModel268d39___setWallets_wallets_atList(void* ptr, int i)
{
	return ({QWallet268d39* tmp = static_cast<QList<QWallet268d39*>*>(ptr)->at(i); if (i == static_cast<QList<QWallet268d39*>*>(ptr)->size()-1) { static_cast<QList<QWallet268d39*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel268d39___setWallets_wallets_setList(void* ptr, void* i)
{
	static_cast<QList<QWallet268d39*>*>(ptr)->append(static_cast<QWallet268d39*>(i));
}

void* WalletModel268d39___setWallets_wallets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWallet268d39*>();
}

void* WalletModel268d39___walletsChanged_wallets_atList(void* ptr, int i)
{
	return ({QWallet268d39* tmp = static_cast<QList<QWallet268d39*>*>(ptr)->at(i); if (i == static_cast<QList<QWallet268d39*>*>(ptr)->size()-1) { static_cast<QList<QWallet268d39*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel268d39___walletsChanged_wallets_setList(void* ptr, void* i)
{
	static_cast<QList<QWallet268d39*>*>(ptr)->append(static_cast<QWallet268d39*>(i));
}

void* WalletModel268d39___walletsChanged_wallets_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QWallet268d39*>();
}

int WalletModel268d39_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel268d39_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* WalletModel268d39_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int WalletModel268d39_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel268d39_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* WalletModel268d39_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int WalletModel268d39_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletModel268d39_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* WalletModel268d39_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* WalletModel268d39_NewWalletModel(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new WalletModel268d39(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new WalletModel268d39(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new WalletModel268d39(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new WalletModel268d39(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new WalletModel268d39(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new WalletModel268d39(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new WalletModel268d39(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new WalletModel268d39(static_cast<QWindow*>(parent));
	} else {
		return new WalletModel268d39(static_cast<QObject*>(parent));
	}
}

void WalletModel268d39_DestroyWalletModel(void* ptr)
{
	static_cast<WalletModel268d39*>(ptr)->~WalletModel268d39();
}

void WalletModel268d39_DestroyWalletModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char WalletModel268d39_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

long long WalletModel268d39_FlagsDefault(void* ptr, void* index)
{
	return static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

void* WalletModel268d39_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* WalletModel268d39_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* WalletModel268d39_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

char WalletModel268d39_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char WalletModel268d39_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int WalletModel268d39_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* WalletModel268d39_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void WalletModel268d39_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char WalletModel268d39_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* WalletModel268d39_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

char WalletModel268d39_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char WalletModel268d39_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct Moc_PackedList WalletModel268d39_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList WalletModel268d39_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* WalletModel268d39_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct Moc_PackedString WalletModel268d39_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t2c50d2 = static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::mimeTypes().join("¡¦!").toUtf8(); Moc_PackedString { const_cast<char*>(t2c50d2.prepend("WHITESPACE").constData()+10), t2c50d2.size()-10 }; });
}

char WalletModel268d39_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char WalletModel268d39_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* WalletModel268d39_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

char WalletModel268d39_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char WalletModel268d39_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void WalletModel268d39_ResetInternalDataDefault(void* ptr)
{
	static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::resetInternalData();
}

void WalletModel268d39_RevertDefault(void* ptr)
{
	static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::revert();
}

struct Moc_PackedList WalletModel268d39_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int WalletModel268d39_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char WalletModel268d39_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char WalletModel268d39_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char WalletModel268d39_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void WalletModel268d39_SortDefault(void* ptr, int column, long long order)
{
	static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* WalletModel268d39_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char WalletModel268d39_SubmitDefault(void* ptr)
{
	return static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::submit();
}

long long WalletModel268d39_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long WalletModel268d39_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::supportedDropActions();
}

void WalletModel268d39_ChildEventDefault(void* ptr, void* event)
{
	static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void WalletModel268d39_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void WalletModel268d39_CustomEventDefault(void* ptr, void* event)
{
	static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void WalletModel268d39_DeleteLaterDefault(void* ptr)
{
	static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::deleteLater();
}

void WalletModel268d39_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char WalletModel268d39_EventDefault(void* ptr, void* e)
{
	return static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char WalletModel268d39_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void WalletModel268d39_TimerEventDefault(void* ptr, void* event)
{
	static_cast<WalletModel268d39*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

void AddressesModel268d39_AddAddress(void* ptr, void* v0)
{
	QMetaObject::invokeMethod(static_cast<AddressesModel268d39*>(ptr), "addAddress", Q_ARG(QAddress268d39*, static_cast<QAddress268d39*>(v0)));
}

void AddressesModel268d39_RemoveAddress(void* ptr, int v0)
{
	QMetaObject::invokeMethod(static_cast<AddressesModel268d39*>(ptr), "removeAddress", Q_ARG(qint32, v0));
}

void AddressesModel268d39_EditAddress(void* ptr, int v0, struct Moc_PackedString v1, int v2, int v3)
{
	QMetaObject::invokeMethod(static_cast<AddressesModel268d39*>(ptr), "editAddress", Q_ARG(qint32, v0), Q_ARG(QString, QString::fromUtf8(v1.data, v1.len)), Q_ARG(qint32, v2), Q_ARG(qint32, v3));
}

void AddressesModel268d39_LoadModel(void* ptr, struct Moc_PackedString v0)
{
	QMetaObject::invokeMethod(static_cast<AddressesModel268d39*>(ptr), "loadModel", Q_ARG(QString, QString::fromUtf8(v0.data, v0.len)));
}

struct Moc_PackedList AddressesModel268d39_Roles(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<AddressesModel268d39*>(ptr)->roles()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressesModel268d39_RolesDefault(void* ptr)
{
	return ({ QMap<qint32, QByteArray>* tmpValue = new QMap<qint32, QByteArray>(static_cast<AddressesModel268d39*>(ptr)->rolesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void AddressesModel268d39_SetRoles(void* ptr, void* roles)
{
	static_cast<AddressesModel268d39*>(ptr)->setRoles(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void AddressesModel268d39_SetRolesDefault(void* ptr, void* roles)
{
	static_cast<AddressesModel268d39*>(ptr)->setRolesDefault(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

void AddressesModel268d39_ConnectRolesChanged(void* ptr)
{
	QObject::connect(static_cast<AddressesModel268d39*>(ptr), static_cast<void (AddressesModel268d39::*)(QMap<qint32, QByteArray>)>(&AddressesModel268d39::rolesChanged), static_cast<AddressesModel268d39*>(ptr), static_cast<void (AddressesModel268d39::*)(QMap<qint32, QByteArray>)>(&AddressesModel268d39::Signal_RolesChanged));
}

void AddressesModel268d39_DisconnectRolesChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressesModel268d39*>(ptr), static_cast<void (AddressesModel268d39::*)(QMap<qint32, QByteArray>)>(&AddressesModel268d39::rolesChanged), static_cast<AddressesModel268d39*>(ptr), static_cast<void (AddressesModel268d39::*)(QMap<qint32, QByteArray>)>(&AddressesModel268d39::Signal_RolesChanged));
}

void AddressesModel268d39_RolesChanged(void* ptr, void* roles)
{
	static_cast<AddressesModel268d39*>(ptr)->rolesChanged(({ QMap<qint32, QByteArray>* tmpP = static_cast<QMap<qint32, QByteArray>*>(roles); QMap<qint32, QByteArray> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }));
}

struct Moc_PackedList AddressesModel268d39_Addresses(void* ptr)
{
	return ({ QList<QAddress268d39*>* tmpValue = new QList<QAddress268d39*>(static_cast<AddressesModel268d39*>(ptr)->addresses()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressesModel268d39_AddressesDefault(void* ptr)
{
	return ({ QList<QAddress268d39*>* tmpValue = new QList<QAddress268d39*>(static_cast<AddressesModel268d39*>(ptr)->addressesDefault()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void AddressesModel268d39_SetAddresses(void* ptr, void* addresses)
{
	static_cast<AddressesModel268d39*>(ptr)->setAddresses(({ QList<QAddress268d39*>* tmpP = static_cast<QList<QAddress268d39*>*>(addresses); QList<QAddress268d39*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void AddressesModel268d39_SetAddressesDefault(void* ptr, void* addresses)
{
	static_cast<AddressesModel268d39*>(ptr)->setAddressesDefault(({ QList<QAddress268d39*>* tmpP = static_cast<QList<QAddress268d39*>*>(addresses); QList<QAddress268d39*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void AddressesModel268d39_ConnectAddressesChanged(void* ptr)
{
	QObject::connect(static_cast<AddressesModel268d39*>(ptr), static_cast<void (AddressesModel268d39::*)(QList<QAddress268d39*>)>(&AddressesModel268d39::addressesChanged), static_cast<AddressesModel268d39*>(ptr), static_cast<void (AddressesModel268d39::*)(QList<QAddress268d39*>)>(&AddressesModel268d39::Signal_AddressesChanged));
}

void AddressesModel268d39_DisconnectAddressesChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressesModel268d39*>(ptr), static_cast<void (AddressesModel268d39::*)(QList<QAddress268d39*>)>(&AddressesModel268d39::addressesChanged), static_cast<AddressesModel268d39*>(ptr), static_cast<void (AddressesModel268d39::*)(QList<QAddress268d39*>)>(&AddressesModel268d39::Signal_AddressesChanged));
}

void AddressesModel268d39_AddressesChanged(void* ptr, void* addresses)
{
	static_cast<AddressesModel268d39*>(ptr)->addressesChanged(({ QList<QAddress268d39*>* tmpP = static_cast<QList<QAddress268d39*>*>(addresses); QList<QAddress268d39*> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

int AddressesModel268d39_Loaded(void* ptr)
{
	return static_cast<AddressesModel268d39*>(ptr)->loaded();
}

int AddressesModel268d39_LoadedDefault(void* ptr)
{
	return static_cast<AddressesModel268d39*>(ptr)->loadedDefault();
}

void AddressesModel268d39_SetLoaded(void* ptr, int loaded)
{
	static_cast<AddressesModel268d39*>(ptr)->setLoaded(loaded);
}

void AddressesModel268d39_SetLoadedDefault(void* ptr, int loaded)
{
	static_cast<AddressesModel268d39*>(ptr)->setLoadedDefault(loaded);
}

void AddressesModel268d39_ConnectLoadedChanged(void* ptr)
{
	QObject::connect(static_cast<AddressesModel268d39*>(ptr), static_cast<void (AddressesModel268d39::*)(qint32)>(&AddressesModel268d39::loadedChanged), static_cast<AddressesModel268d39*>(ptr), static_cast<void (AddressesModel268d39::*)(qint32)>(&AddressesModel268d39::Signal_LoadedChanged));
}

void AddressesModel268d39_DisconnectLoadedChanged(void* ptr)
{
	QObject::disconnect(static_cast<AddressesModel268d39*>(ptr), static_cast<void (AddressesModel268d39::*)(qint32)>(&AddressesModel268d39::loadedChanged), static_cast<AddressesModel268d39*>(ptr), static_cast<void (AddressesModel268d39::*)(qint32)>(&AddressesModel268d39::Signal_LoadedChanged));
}

void AddressesModel268d39_LoadedChanged(void* ptr, int loaded)
{
	static_cast<AddressesModel268d39*>(ptr)->loadedChanged(loaded);
}

int AddressesModel268d39_AddressesModel268d39_QRegisterMetaType()
{
	return qRegisterMetaType<AddressesModel268d39*>();
}

int AddressesModel268d39_AddressesModel268d39_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<AddressesModel268d39*>(const_cast<const char*>(typeName));
}

int AddressesModel268d39_AddressesModel268d39_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressesModel268d39>();
#else
	return 0;
#endif
}

int AddressesModel268d39_AddressesModel268d39_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<AddressesModel268d39>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

int AddressesModel268d39_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel268d39_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModel268d39_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressesModel268d39_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel268d39_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModel268d39_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressesModel268d39_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel268d39_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModel268d39_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* AddressesModel268d39___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel268d39___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModel268d39___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressesModel268d39___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel268d39___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModel268d39___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int AddressesModel268d39___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void AddressesModel268d39___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* AddressesModel268d39___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* AddressesModel268d39___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModel268d39___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* AddressesModel268d39___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList AddressesModel268d39___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel268d39___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel268d39___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* AddressesModel268d39___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* AddressesModel268d39___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel268d39___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* AddressesModel268d39___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* AddressesModel268d39___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel268d39___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModel268d39___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressesModel268d39___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel268d39___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModel268d39___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressesModel268d39___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel268d39___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* AddressesModel268d39___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* AddressesModel268d39___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void AddressesModel268d39___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressesModel268d39___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct Moc_PackedList AddressesModel268d39___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel268d39___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModel268d39___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* AddressesModel268d39___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct Moc_PackedList AddressesModel268d39___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int AddressesModel268d39_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel268d39_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModel268d39_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int AddressesModel268d39_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel268d39_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* AddressesModel268d39_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

void* AddressesModel268d39___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel268d39___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressesModel268d39___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* AddressesModel268d39___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void AddressesModel268d39___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* AddressesModel268d39___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* AddressesModel268d39___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel268d39___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressesModel268d39___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressesModel268d39___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel268d39___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressesModel268d39___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressesModel268d39___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel268d39___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* AddressesModel268d39___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* AddressesModel268d39___roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModel268d39___roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressesModel268d39___roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressesModel268d39___roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel268d39___setRoles_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModel268d39___setRoles_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressesModel268d39___setRoles_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressesModel268d39___setRoles_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel268d39___rolesChanged_roles_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QMap<qint32, QByteArray>*>(ptr)->value(v); if (i == static_cast<QMap<qint32, QByteArray>*>(ptr)->size()-1) { static_cast<QMap<qint32, QByteArray>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void AddressesModel268d39___rolesChanged_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<qint32, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* AddressesModel268d39___rolesChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<qint32, QByteArray>();
}

struct Moc_PackedList AddressesModel268d39___rolesChanged_roles_keyList(void* ptr)
{
	return ({ QList<qint32>* tmpValue = new QList<qint32>(static_cast<QMap<qint32, QByteArray>*>(ptr)->keys()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel268d39___addresses_atList(void* ptr, int i)
{
	return ({QAddress268d39* tmp = static_cast<QList<QAddress268d39*>*>(ptr)->at(i); if (i == static_cast<QList<QAddress268d39*>*>(ptr)->size()-1) { static_cast<QList<QAddress268d39*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel268d39___addresses_setList(void* ptr, void* i)
{
	static_cast<QList<QAddress268d39*>*>(ptr)->append(static_cast<QAddress268d39*>(i));
}

void* AddressesModel268d39___addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAddress268d39*>();
}

void* AddressesModel268d39___setAddresses_addresses_atList(void* ptr, int i)
{
	return ({QAddress268d39* tmp = static_cast<QList<QAddress268d39*>*>(ptr)->at(i); if (i == static_cast<QList<QAddress268d39*>*>(ptr)->size()-1) { static_cast<QList<QAddress268d39*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel268d39___setAddresses_addresses_setList(void* ptr, void* i)
{
	static_cast<QList<QAddress268d39*>*>(ptr)->append(static_cast<QAddress268d39*>(i));
}

void* AddressesModel268d39___setAddresses_addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAddress268d39*>();
}

void* AddressesModel268d39___addressesChanged_addresses_atList(void* ptr, int i)
{
	return ({QAddress268d39* tmp = static_cast<QList<QAddress268d39*>*>(ptr)->at(i); if (i == static_cast<QList<QAddress268d39*>*>(ptr)->size()-1) { static_cast<QList<QAddress268d39*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel268d39___addressesChanged_addresses_setList(void* ptr, void* i)
{
	static_cast<QList<QAddress268d39*>*>(ptr)->append(static_cast<QAddress268d39*>(i));
}

void* AddressesModel268d39___addressesChanged_addresses_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QAddress268d39*>();
}

int AddressesModel268d39_____roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel268d39_____roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressesModel268d39_____roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int AddressesModel268d39_____setRoles_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel268d39_____setRoles_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressesModel268d39_____setRoles_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

int AddressesModel268d39_____rolesChanged_roles_keyList_atList(void* ptr, int i)
{
	return ({qint32 tmp = static_cast<QList<qint32>*>(ptr)->at(i); if (i == static_cast<QList<qint32>*>(ptr)->size()-1) { static_cast<QList<qint32>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void AddressesModel268d39_____rolesChanged_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<qint32>*>(ptr)->append(i);
}

void* AddressesModel268d39_____rolesChanged_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<qint32>();
}

void* AddressesModel268d39_NewAddressesModel(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new AddressesModel268d39(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new AddressesModel268d39(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new AddressesModel268d39(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new AddressesModel268d39(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new AddressesModel268d39(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new AddressesModel268d39(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new AddressesModel268d39(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new AddressesModel268d39(static_cast<QWindow*>(parent));
	} else {
		return new AddressesModel268d39(static_cast<QObject*>(parent));
	}
}

void AddressesModel268d39_DestroyAddressesModel(void* ptr)
{
	static_cast<AddressesModel268d39*>(ptr)->~AddressesModel268d39();
}

void AddressesModel268d39_DestroyAddressesModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

char AddressesModel268d39_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

long long AddressesModel268d39_FlagsDefault(void* ptr, void* index)
{
	return static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::flags(*static_cast<QModelIndex*>(index));
}

void* AddressesModel268d39_IndexDefault(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* AddressesModel268d39_SiblingDefault(void* ptr, int row, int column, void* idx)
{
	return new QModelIndex(static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::sibling(row, column, *static_cast<QModelIndex*>(idx)));
}

void* AddressesModel268d39_BuddyDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::buddy(*static_cast<QModelIndex*>(index)));
}

char AddressesModel268d39_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char AddressesModel268d39_CanFetchMoreDefault(void* ptr, void* parent)
{
	return static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::canFetchMore(*static_cast<QModelIndex*>(parent));
}

int AddressesModel268d39_ColumnCountDefault(void* ptr, void* parent)
{
	return static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::columnCount(*static_cast<QModelIndex*>(parent));
}

void* AddressesModel268d39_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);

}

void AddressesModel268d39_FetchMoreDefault(void* ptr, void* parent)
{
	static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::fetchMore(*static_cast<QModelIndex*>(parent));
}

char AddressesModel268d39_HasChildrenDefault(void* ptr, void* parent)
{
	return static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::hasChildren(*static_cast<QModelIndex*>(parent));
}

void* AddressesModel268d39_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

char AddressesModel268d39_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char AddressesModel268d39_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

struct Moc_PackedList AddressesModel268d39_ItemDataDefault(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue = new QMap<int, QVariant>(static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::itemData(*static_cast<QModelIndex*>(index))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

struct Moc_PackedList AddressesModel268d39_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValue = new QList<QModelIndex>(static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

void* AddressesModel268d39_MimeDataDefault(void* ptr, void* indexes)
{
	return static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

struct Moc_PackedString AddressesModel268d39_MimeTypesDefault(void* ptr)
{
	return ({ QByteArray t6a4f86 = static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::mimeTypes().join("¡¦!").toUtf8(); Moc_PackedString { const_cast<char*>(t6a4f86.prepend("WHITESPACE").constData()+10), t6a4f86.size()-10 }; });
}

char AddressesModel268d39_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char AddressesModel268d39_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

void* AddressesModel268d39_ParentDefault(void* ptr, void* index)
{
	return new QModelIndex(static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::parent(*static_cast<QModelIndex*>(index)));
}

char AddressesModel268d39_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	return static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char AddressesModel268d39_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	return static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

void AddressesModel268d39_ResetInternalDataDefault(void* ptr)
{
	static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::resetInternalData();
}

void AddressesModel268d39_RevertDefault(void* ptr)
{
	static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::revert();
}

struct Moc_PackedList AddressesModel268d39_RoleNamesDefault(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue = new QHash<int, QByteArray>(static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::roleNames()); Moc_PackedList { tmpValue, tmpValue->size() }; });
}

int AddressesModel268d39_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);

}

char AddressesModel268d39_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	return static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char AddressesModel268d39_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char AddressesModel268d39_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	return static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

void AddressesModel268d39_SortDefault(void* ptr, int column, long long order)
{
	static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::sort(column, static_cast<Qt::SortOrder>(order));
}

void* AddressesModel268d39_SpanDefault(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

char AddressesModel268d39_SubmitDefault(void* ptr)
{
	return static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::submit();
}

long long AddressesModel268d39_SupportedDragActionsDefault(void* ptr)
{
	return static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::supportedDragActions();
}

long long AddressesModel268d39_SupportedDropActionsDefault(void* ptr)
{
	return static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::supportedDropActions();
}

void AddressesModel268d39_ChildEventDefault(void* ptr, void* event)
{
	static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::childEvent(static_cast<QChildEvent*>(event));
}

void AddressesModel268d39_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void AddressesModel268d39_CustomEventDefault(void* ptr, void* event)
{
	static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::customEvent(static_cast<QEvent*>(event));
}

void AddressesModel268d39_DeleteLaterDefault(void* ptr)
{
	static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::deleteLater();
}

void AddressesModel268d39_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char AddressesModel268d39_EventDefault(void* ptr, void* e)
{
	return static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::event(static_cast<QEvent*>(e));
}

char AddressesModel268d39_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void AddressesModel268d39_TimerEventDefault(void* ptr, void* event)
{
	static_cast<AddressesModel268d39*>(ptr)->QAbstractListModel::timerEvent(static_cast<QTimerEvent*>(event));
}

struct Moc_PackedString QAddress268d39_Address(void* ptr)
{
	return ({ QByteArray tbbcec3 = static_cast<QAddress268d39*>(ptr)->address().toUtf8(); Moc_PackedString { const_cast<char*>(tbbcec3.prepend("WHITESPACE").constData()+10), tbbcec3.size()-10 }; });
}

struct Moc_PackedString QAddress268d39_AddressDefault(void* ptr)
{
	return ({ QByteArray tbe1549 = static_cast<QAddress268d39*>(ptr)->addressDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tbe1549.prepend("WHITESPACE").constData()+10), tbe1549.size()-10 }; });
}

void QAddress268d39_SetAddress(void* ptr, struct Moc_PackedString address)
{
	static_cast<QAddress268d39*>(ptr)->setAddress(QString::fromUtf8(address.data, address.len));
}

void QAddress268d39_SetAddressDefault(void* ptr, struct Moc_PackedString address)
{
	static_cast<QAddress268d39*>(ptr)->setAddressDefault(QString::fromUtf8(address.data, address.len));
}

void QAddress268d39_ConnectAddressChanged(void* ptr)
{
	QObject::connect(static_cast<QAddress268d39*>(ptr), static_cast<void (QAddress268d39::*)(QString)>(&QAddress268d39::addressChanged), static_cast<QAddress268d39*>(ptr), static_cast<void (QAddress268d39::*)(QString)>(&QAddress268d39::Signal_AddressChanged));
}

void QAddress268d39_DisconnectAddressChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAddress268d39*>(ptr), static_cast<void (QAddress268d39::*)(QString)>(&QAddress268d39::addressChanged), static_cast<QAddress268d39*>(ptr), static_cast<void (QAddress268d39::*)(QString)>(&QAddress268d39::Signal_AddressChanged));
}

void QAddress268d39_AddressChanged(void* ptr, struct Moc_PackedString address)
{
	static_cast<QAddress268d39*>(ptr)->addressChanged(QString::fromUtf8(address.data, address.len));
}

int QAddress268d39_AddressSky(void* ptr)
{
	return static_cast<QAddress268d39*>(ptr)->addressSky();
}

int QAddress268d39_AddressSkyDefault(void* ptr)
{
	return static_cast<QAddress268d39*>(ptr)->addressSkyDefault();
}

void QAddress268d39_SetAddressSky(void* ptr, int addressSky)
{
	static_cast<QAddress268d39*>(ptr)->setAddressSky(addressSky);
}

void QAddress268d39_SetAddressSkyDefault(void* ptr, int addressSky)
{
	static_cast<QAddress268d39*>(ptr)->setAddressSkyDefault(addressSky);
}

void QAddress268d39_ConnectAddressSkyChanged(void* ptr)
{
	QObject::connect(static_cast<QAddress268d39*>(ptr), static_cast<void (QAddress268d39::*)(qint32)>(&QAddress268d39::addressSkyChanged), static_cast<QAddress268d39*>(ptr), static_cast<void (QAddress268d39::*)(qint32)>(&QAddress268d39::Signal_AddressSkyChanged));
}

void QAddress268d39_DisconnectAddressSkyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAddress268d39*>(ptr), static_cast<void (QAddress268d39::*)(qint32)>(&QAddress268d39::addressSkyChanged), static_cast<QAddress268d39*>(ptr), static_cast<void (QAddress268d39::*)(qint32)>(&QAddress268d39::Signal_AddressSkyChanged));
}

void QAddress268d39_AddressSkyChanged(void* ptr, int addressSky)
{
	static_cast<QAddress268d39*>(ptr)->addressSkyChanged(addressSky);
}

int QAddress268d39_AddressCoinHours(void* ptr)
{
	return static_cast<QAddress268d39*>(ptr)->addressCoinHours();
}

int QAddress268d39_AddressCoinHoursDefault(void* ptr)
{
	return static_cast<QAddress268d39*>(ptr)->addressCoinHoursDefault();
}

void QAddress268d39_SetAddressCoinHours(void* ptr, int addressCoinHours)
{
	static_cast<QAddress268d39*>(ptr)->setAddressCoinHours(addressCoinHours);
}

void QAddress268d39_SetAddressCoinHoursDefault(void* ptr, int addressCoinHours)
{
	static_cast<QAddress268d39*>(ptr)->setAddressCoinHoursDefault(addressCoinHours);
}

void QAddress268d39_ConnectAddressCoinHoursChanged(void* ptr)
{
	QObject::connect(static_cast<QAddress268d39*>(ptr), static_cast<void (QAddress268d39::*)(qint32)>(&QAddress268d39::addressCoinHoursChanged), static_cast<QAddress268d39*>(ptr), static_cast<void (QAddress268d39::*)(qint32)>(&QAddress268d39::Signal_AddressCoinHoursChanged));
}

void QAddress268d39_DisconnectAddressCoinHoursChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAddress268d39*>(ptr), static_cast<void (QAddress268d39::*)(qint32)>(&QAddress268d39::addressCoinHoursChanged), static_cast<QAddress268d39*>(ptr), static_cast<void (QAddress268d39::*)(qint32)>(&QAddress268d39::Signal_AddressCoinHoursChanged));
}

void QAddress268d39_AddressCoinHoursChanged(void* ptr, int addressCoinHours)
{
	static_cast<QAddress268d39*>(ptr)->addressCoinHoursChanged(addressCoinHours);
}

int QAddress268d39_QAddress268d39_QRegisterMetaType()
{
	return qRegisterMetaType<QAddress268d39*>();
}

int QAddress268d39_QAddress268d39_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QAddress268d39*>(const_cast<const char*>(typeName));
}

int QAddress268d39_QAddress268d39_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QAddress268d39>();
#else
	return 0;
#endif
}

int QAddress268d39_QAddress268d39_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QAddress268d39>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QAddress268d39___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAddress268d39___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAddress268d39___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QAddress268d39___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAddress268d39___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QAddress268d39___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QAddress268d39___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAddress268d39___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAddress268d39___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAddress268d39___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAddress268d39___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAddress268d39___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAddress268d39___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAddress268d39___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QAddress268d39___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QAddress268d39_NewQAddress(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new QAddress268d39(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new QAddress268d39(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new QAddress268d39(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QAddress268d39(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QAddress268d39(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QAddress268d39(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new QAddress268d39(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QAddress268d39(static_cast<QWindow*>(parent));
	} else {
		return new QAddress268d39(static_cast<QObject*>(parent));
	}
}

void QAddress268d39_DestroyQAddress(void* ptr)
{
	static_cast<QAddress268d39*>(ptr)->~QAddress268d39();
}

void QAddress268d39_DestroyQAddressDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void QAddress268d39_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QAddress268d39*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QAddress268d39_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAddress268d39*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QAddress268d39_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QAddress268d39*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QAddress268d39_DeleteLaterDefault(void* ptr)
{
	static_cast<QAddress268d39*>(ptr)->QObject::deleteLater();
}

void QAddress268d39_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QAddress268d39*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QAddress268d39_EventDefault(void* ptr, void* e)
{
	return static_cast<QAddress268d39*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QAddress268d39_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QAddress268d39*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QAddress268d39_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QAddress268d39*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

struct Moc_PackedString QWallet268d39_Name(void* ptr)
{
	return ({ QByteArray tedc734 = static_cast<QWallet268d39*>(ptr)->name().toUtf8(); Moc_PackedString { const_cast<char*>(tedc734.prepend("WHITESPACE").constData()+10), tedc734.size()-10 }; });
}

struct Moc_PackedString QWallet268d39_NameDefault(void* ptr)
{
	return ({ QByteArray tee7a91 = static_cast<QWallet268d39*>(ptr)->nameDefault().toUtf8(); Moc_PackedString { const_cast<char*>(tee7a91.prepend("WHITESPACE").constData()+10), tee7a91.size()-10 }; });
}

void QWallet268d39_SetName(void* ptr, struct Moc_PackedString name)
{
	static_cast<QWallet268d39*>(ptr)->setName(QString::fromUtf8(name.data, name.len));
}

void QWallet268d39_SetNameDefault(void* ptr, struct Moc_PackedString name)
{
	static_cast<QWallet268d39*>(ptr)->setNameDefault(QString::fromUtf8(name.data, name.len));
}

void QWallet268d39_ConnectNameChanged(void* ptr)
{
	QObject::connect(static_cast<QWallet268d39*>(ptr), static_cast<void (QWallet268d39::*)(QString)>(&QWallet268d39::nameChanged), static_cast<QWallet268d39*>(ptr), static_cast<void (QWallet268d39::*)(QString)>(&QWallet268d39::Signal_NameChanged));
}

void QWallet268d39_DisconnectNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWallet268d39*>(ptr), static_cast<void (QWallet268d39::*)(QString)>(&QWallet268d39::nameChanged), static_cast<QWallet268d39*>(ptr), static_cast<void (QWallet268d39::*)(QString)>(&QWallet268d39::Signal_NameChanged));
}

void QWallet268d39_NameChanged(void* ptr, struct Moc_PackedString name)
{
	static_cast<QWallet268d39*>(ptr)->nameChanged(QString::fromUtf8(name.data, name.len));
}

int QWallet268d39_EncryptionEnabled(void* ptr)
{
	return static_cast<QWallet268d39*>(ptr)->encryptionEnabled();
}

int QWallet268d39_EncryptionEnabledDefault(void* ptr)
{
	return static_cast<QWallet268d39*>(ptr)->encryptionEnabledDefault();
}

void QWallet268d39_SetEncryptionEnabled(void* ptr, int encryptionEnabled)
{
	static_cast<QWallet268d39*>(ptr)->setEncryptionEnabled(encryptionEnabled);
}

void QWallet268d39_SetEncryptionEnabledDefault(void* ptr, int encryptionEnabled)
{
	static_cast<QWallet268d39*>(ptr)->setEncryptionEnabledDefault(encryptionEnabled);
}

void QWallet268d39_ConnectEncryptionEnabledChanged(void* ptr)
{
	QObject::connect(static_cast<QWallet268d39*>(ptr), static_cast<void (QWallet268d39::*)(qint32)>(&QWallet268d39::encryptionEnabledChanged), static_cast<QWallet268d39*>(ptr), static_cast<void (QWallet268d39::*)(qint32)>(&QWallet268d39::Signal_EncryptionEnabledChanged));
}

void QWallet268d39_DisconnectEncryptionEnabledChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWallet268d39*>(ptr), static_cast<void (QWallet268d39::*)(qint32)>(&QWallet268d39::encryptionEnabledChanged), static_cast<QWallet268d39*>(ptr), static_cast<void (QWallet268d39::*)(qint32)>(&QWallet268d39::Signal_EncryptionEnabledChanged));
}

void QWallet268d39_EncryptionEnabledChanged(void* ptr, int encryptionEnabled)
{
	static_cast<QWallet268d39*>(ptr)->encryptionEnabledChanged(encryptionEnabled);
}

int QWallet268d39_Sky(void* ptr)
{
	return static_cast<QWallet268d39*>(ptr)->sky();
}

int QWallet268d39_SkyDefault(void* ptr)
{
	return static_cast<QWallet268d39*>(ptr)->skyDefault();
}

void QWallet268d39_SetSky(void* ptr, int sky)
{
	static_cast<QWallet268d39*>(ptr)->setSky(sky);
}

void QWallet268d39_SetSkyDefault(void* ptr, int sky)
{
	static_cast<QWallet268d39*>(ptr)->setSkyDefault(sky);
}

void QWallet268d39_ConnectSkyChanged(void* ptr)
{
	QObject::connect(static_cast<QWallet268d39*>(ptr), static_cast<void (QWallet268d39::*)(qint32)>(&QWallet268d39::skyChanged), static_cast<QWallet268d39*>(ptr), static_cast<void (QWallet268d39::*)(qint32)>(&QWallet268d39::Signal_SkyChanged));
}

void QWallet268d39_DisconnectSkyChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWallet268d39*>(ptr), static_cast<void (QWallet268d39::*)(qint32)>(&QWallet268d39::skyChanged), static_cast<QWallet268d39*>(ptr), static_cast<void (QWallet268d39::*)(qint32)>(&QWallet268d39::Signal_SkyChanged));
}

void QWallet268d39_SkyChanged(void* ptr, int sky)
{
	static_cast<QWallet268d39*>(ptr)->skyChanged(sky);
}

int QWallet268d39_CoinHours(void* ptr)
{
	return static_cast<QWallet268d39*>(ptr)->coinHours();
}

int QWallet268d39_CoinHoursDefault(void* ptr)
{
	return static_cast<QWallet268d39*>(ptr)->coinHoursDefault();
}

void QWallet268d39_SetCoinHours(void* ptr, int coinHours)
{
	static_cast<QWallet268d39*>(ptr)->setCoinHours(coinHours);
}

void QWallet268d39_SetCoinHoursDefault(void* ptr, int coinHours)
{
	static_cast<QWallet268d39*>(ptr)->setCoinHoursDefault(coinHours);
}

void QWallet268d39_ConnectCoinHoursChanged(void* ptr)
{
	QObject::connect(static_cast<QWallet268d39*>(ptr), static_cast<void (QWallet268d39::*)(qint32)>(&QWallet268d39::coinHoursChanged), static_cast<QWallet268d39*>(ptr), static_cast<void (QWallet268d39::*)(qint32)>(&QWallet268d39::Signal_CoinHoursChanged));
}

void QWallet268d39_DisconnectCoinHoursChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWallet268d39*>(ptr), static_cast<void (QWallet268d39::*)(qint32)>(&QWallet268d39::coinHoursChanged), static_cast<QWallet268d39*>(ptr), static_cast<void (QWallet268d39::*)(qint32)>(&QWallet268d39::Signal_CoinHoursChanged));
}

void QWallet268d39_CoinHoursChanged(void* ptr, int coinHours)
{
	static_cast<QWallet268d39*>(ptr)->coinHoursChanged(coinHours);
}

struct Moc_PackedString QWallet268d39_FileName(void* ptr)
{
	return ({ QByteArray teb23e1 = static_cast<QWallet268d39*>(ptr)->fileName().toUtf8(); Moc_PackedString { const_cast<char*>(teb23e1.prepend("WHITESPACE").constData()+10), teb23e1.size()-10 }; });
}

struct Moc_PackedString QWallet268d39_FileNameDefault(void* ptr)
{
	return ({ QByteArray t811a22 = static_cast<QWallet268d39*>(ptr)->fileNameDefault().toUtf8(); Moc_PackedString { const_cast<char*>(t811a22.prepend("WHITESPACE").constData()+10), t811a22.size()-10 }; });
}

void QWallet268d39_SetFileName(void* ptr, struct Moc_PackedString fileName)
{
	static_cast<QWallet268d39*>(ptr)->setFileName(QString::fromUtf8(fileName.data, fileName.len));
}

void QWallet268d39_SetFileNameDefault(void* ptr, struct Moc_PackedString fileName)
{
	static_cast<QWallet268d39*>(ptr)->setFileNameDefault(QString::fromUtf8(fileName.data, fileName.len));
}

void QWallet268d39_ConnectFileNameChanged(void* ptr)
{
	QObject::connect(static_cast<QWallet268d39*>(ptr), static_cast<void (QWallet268d39::*)(QString)>(&QWallet268d39::fileNameChanged), static_cast<QWallet268d39*>(ptr), static_cast<void (QWallet268d39::*)(QString)>(&QWallet268d39::Signal_FileNameChanged));
}

void QWallet268d39_DisconnectFileNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWallet268d39*>(ptr), static_cast<void (QWallet268d39::*)(QString)>(&QWallet268d39::fileNameChanged), static_cast<QWallet268d39*>(ptr), static_cast<void (QWallet268d39::*)(QString)>(&QWallet268d39::Signal_FileNameChanged));
}

void QWallet268d39_FileNameChanged(void* ptr, struct Moc_PackedString fileName)
{
	static_cast<QWallet268d39*>(ptr)->fileNameChanged(QString::fromUtf8(fileName.data, fileName.len));
}

int QWallet268d39_QWallet268d39_QRegisterMetaType()
{
	return qRegisterMetaType<QWallet268d39*>();
}

int QWallet268d39_QWallet268d39_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<QWallet268d39*>(const_cast<const char*>(typeName));
}

int QWallet268d39_QWallet268d39_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QWallet268d39>();
#else
	return 0;
#endif
}

int QWallet268d39_QWallet268d39_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<QWallet268d39>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* QWallet268d39___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWallet268d39___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWallet268d39___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QWallet268d39___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWallet268d39___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWallet268d39___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QWallet268d39___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWallet268d39___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWallet268d39___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWallet268d39___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWallet268d39___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWallet268d39___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWallet268d39___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWallet268d39___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QWallet268d39___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QWallet268d39_NewQWallet(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new QWallet268d39(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new QWallet268d39(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new QWallet268d39(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new QWallet268d39(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new QWallet268d39(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new QWallet268d39(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new QWallet268d39(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new QWallet268d39(static_cast<QWindow*>(parent));
	} else {
		return new QWallet268d39(static_cast<QObject*>(parent));
	}
}

void QWallet268d39_DestroyQWallet(void* ptr)
{
	static_cast<QWallet268d39*>(ptr)->~QWallet268d39();
}

void QWallet268d39_DestroyQWalletDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void QWallet268d39_ChildEventDefault(void* ptr, void* event)
{
	static_cast<QWallet268d39*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void QWallet268d39_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWallet268d39*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWallet268d39_CustomEventDefault(void* ptr, void* event)
{
	static_cast<QWallet268d39*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void QWallet268d39_DeleteLaterDefault(void* ptr)
{
	static_cast<QWallet268d39*>(ptr)->QObject::deleteLater();
}

void QWallet268d39_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<QWallet268d39*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QWallet268d39_EventDefault(void* ptr, void* e)
{
	return static_cast<QWallet268d39*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char QWallet268d39_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<QWallet268d39*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void QWallet268d39_TimerEventDefault(void* ptr, void* event)
{
	static_cast<QWallet268d39*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

void WalletManager268d39_CreateEncryptedWallet(void* ptr, void* walletM, struct Moc_PackedString seed, struct Moc_PackedString label, struct Moc_PackedString password, int scanN)
{
	QMetaObject::invokeMethod(static_cast<WalletManager268d39*>(ptr), "createEncryptedWallet", Q_ARG(WalletModel268d39*, static_cast<WalletModel268d39*>(walletM)), Q_ARG(QString, QString::fromUtf8(seed.data, seed.len)), Q_ARG(QString, QString::fromUtf8(label.data, label.len)), Q_ARG(QString, QString::fromUtf8(password.data, password.len)), Q_ARG(qint32, scanN));
}

void WalletManager268d39_CreateUnencryptedWallet(void* ptr, void* walletM, struct Moc_PackedString seed, struct Moc_PackedString label, int scanN)
{
	QMetaObject::invokeMethod(static_cast<WalletManager268d39*>(ptr), "createUnencryptedWallet", Q_ARG(WalletModel268d39*, static_cast<WalletModel268d39*>(walletM)), Q_ARG(QString, QString::fromUtf8(seed.data, seed.len)), Q_ARG(QString, QString::fromUtf8(label.data, label.len)), Q_ARG(qint32, scanN));
}

struct Moc_PackedString WalletManager268d39_GetNewSeed(void* ptr, int entropy)
{
	QString returnArg;
	QMetaObject::invokeMethod(static_cast<WalletManager268d39*>(ptr), "getNewSeed", Q_RETURN_ARG(QString, returnArg), Q_ARG(qint32, entropy));
	return ({ QByteArray t8e5b69 = returnArg.toUtf8(); Moc_PackedString { const_cast<char*>(t8e5b69.prepend("WHITESPACE").constData()+10), t8e5b69.size()-10 }; });
}

int WalletManager268d39_VerifySeed(void* ptr, struct Moc_PackedString seed)
{
	qint32 returnArg;
	QMetaObject::invokeMethod(static_cast<WalletManager268d39*>(ptr), "verifySeed", Q_RETURN_ARG(qint32, returnArg), Q_ARG(QString, QString::fromUtf8(seed.data, seed.len)));
	return returnArg;
}

void WalletManager268d39_NewWalletAddress(void* ptr, void* addressM, struct Moc_PackedString id, int n, struct Moc_PackedString password)
{
	QMetaObject::invokeMethod(static_cast<WalletManager268d39*>(ptr), "newWalletAddress", Q_ARG(AddressesModel268d39*, static_cast<AddressesModel268d39*>(addressM)), Q_ARG(QString, QString::fromUtf8(id.data, id.len)), Q_ARG(qint32, n), Q_ARG(QString, QString::fromUtf8(password.data, password.len)));
}

void WalletManager268d39_EncryptWallet(void* ptr, void* walletM, struct Moc_PackedString id, struct Moc_PackedString password)
{
	QMetaObject::invokeMethod(static_cast<WalletManager268d39*>(ptr), "encryptWallet", Q_ARG(WalletModel268d39*, static_cast<WalletModel268d39*>(walletM)), Q_ARG(QString, QString::fromUtf8(id.data, id.len)), Q_ARG(QString, QString::fromUtf8(password.data, password.len)));
}

void WalletManager268d39_DecryptWallet(void* ptr, void* walletM, struct Moc_PackedString id, struct Moc_PackedString password)
{
	QMetaObject::invokeMethod(static_cast<WalletManager268d39*>(ptr), "decryptWallet", Q_ARG(WalletModel268d39*, static_cast<WalletModel268d39*>(walletM)), Q_ARG(QString, QString::fromUtf8(id.data, id.len)), Q_ARG(QString, QString::fromUtf8(password.data, password.len)));
}

int WalletManager268d39_WalletManager268d39_QRegisterMetaType()
{
	return qRegisterMetaType<WalletManager268d39*>();
}

int WalletManager268d39_WalletManager268d39_QRegisterMetaType2(char* typeName)
{
	return qRegisterMetaType<WalletManager268d39*>(const_cast<const char*>(typeName));
}

int WalletManager268d39_WalletManager268d39_QmlRegisterType()
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletManager268d39>();
#else
	return 0;
#endif
}

int WalletManager268d39_WalletManager268d39_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName)
{
#ifdef QT_QML_LIB
	return qmlRegisterType<WalletManager268d39>(const_cast<const char*>(uri), versionMajor, versionMinor, const_cast<const char*>(qmlName));
#else
	return 0;
#endif
}

void* WalletManager268d39___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager268d39___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletManager268d39___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* WalletManager268d39___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void WalletManager268d39___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* WalletManager268d39___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* WalletManager268d39___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager268d39___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletManager268d39___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletManager268d39___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager268d39___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletManager268d39___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletManager268d39___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void WalletManager268d39___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* WalletManager268d39___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* WalletManager268d39_NewWalletManager(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new WalletManager268d39(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new WalletManager268d39(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new WalletManager268d39(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new WalletManager268d39(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new WalletManager268d39(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new WalletManager268d39(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new WalletManager268d39(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new WalletManager268d39(static_cast<QWindow*>(parent));
	} else {
		return new WalletManager268d39(static_cast<QObject*>(parent));
	}
}

void WalletManager268d39_DestroyWalletManager(void* ptr)
{
	static_cast<WalletManager268d39*>(ptr)->~WalletManager268d39();
}

void WalletManager268d39_DestroyWalletManagerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void WalletManager268d39_ChildEventDefault(void* ptr, void* event)
{
	static_cast<WalletManager268d39*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
}

void WalletManager268d39_ConnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletManager268d39*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void WalletManager268d39_CustomEventDefault(void* ptr, void* event)
{
	static_cast<WalletManager268d39*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
}

void WalletManager268d39_DeleteLaterDefault(void* ptr)
{
	static_cast<WalletManager268d39*>(ptr)->QObject::deleteLater();
}

void WalletManager268d39_DisconnectNotifyDefault(void* ptr, void* sign)
{
	static_cast<WalletManager268d39*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char WalletManager268d39_EventDefault(void* ptr, void* e)
{
	return static_cast<WalletManager268d39*>(ptr)->QObject::event(static_cast<QEvent*>(e));
}

char WalletManager268d39_EventFilterDefault(void* ptr, void* watched, void* event)
{
	return static_cast<WalletManager268d39*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

void WalletManager268d39_TimerEventDefault(void* ptr, void* event)
{
	static_cast<WalletManager268d39*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
}

#include "moc_moc.h"
