// +build minimal

#define protected public
#define private public

#include "core-minimal.h"
#include "_cgo_export.h"

#ifndef QT_CORE_LIB
	#error ------------------------------------------------------------------
	#error please run: '$(go env GOPATH)/bin/qtsetup'
	#error more info here: https://github.com/therecipe/qt/wiki/Installation
	#error ------------------------------------------------------------------
#endif
#include <QAbstractItemModel>
#include <QAbstractProxyModel>
#include <QAbstractTableModel>
#include <QBitArray>
#include <QByteArray>
#include <QCborValue>
#include <QChar>
#include <QChildEvent>
#include <QCoreApplication>
#include <QDataStream>
#include <QDate>
#include <QDateTime>
#include <QDir>
#include <QEasingCurve>
#include <QEvent>
#include <QFile>
#include <QFileDevice>
#include <QFileInfo>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QHash>
#include <QIODevice>
#include <QJsonArray>
#include <QJsonDocument>
#include <QJsonObject>
#include <QJsonValue>
#include <QLatin1Char>
#include <QLatin1String>
#include <QLayout>
#include <QLine>
#include <QLineF>
#include <QLocale>
#include <QMap>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMetaProperty>
#include <QMimeData>
#include <QModelIndex>
#include <QObject>
#include <QPersistentModelIndex>
#include <QPoint>
#include <QPointF>
#include <QQuickItem>
#include <QRect>
#include <QRectF>
#include <QRegExp>
#include <QRegularExpression>
#include <QRegularExpressionMatch>
#include <QSize>
#include <QSizeF>
#include <QSortFilterProxyModel>
#include <QString>
#include <QStringRef>
#include <QSysInfo>
#include <QTime>
#include <QTimeZone>
#include <QTimer>
#include <QTimerEvent>
#include <QUrl>
#include <QUuid>
#include <QVariant>
#include <QVector>
#include <QWidget>
#include <QWindow>
#include <QTextDocument>
#include <QObject>
#include <QStringList>

class MyQAbstractItemModel: public QAbstractItemModel
{
public:
	MyQAbstractItemModel(QObject *parent = Q_NULLPTR) : QAbstractItemModel(parent) {QAbstractItemModel_QAbstractItemModel_QRegisterMetaType();};
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQAbstractItemModel_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackQAbstractItemModel_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackQAbstractItemModel_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackQAbstractItemModel_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackQAbstractItemModel_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackQAbstractItemModel_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackQAbstractItemModel_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackQAbstractItemModel_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue037c88 = new QVector<int>(roles); QtCore_PackedList { tmpValue037c88, tmpValue037c88->size() }; })); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackQAbstractItemModel_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	void fetchMore(const QModelIndex & parent) { callbackQAbstractItemModel_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackQAbstractItemModel_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool hasChildren(const QModelIndex & parent) const { return callbackQAbstractItemModel_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackQAbstractItemModel_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackQAbstractItemModel_HeaderDataChanged(this, orientation, first, last); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackQAbstractItemModel_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackQAbstractItemModel_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackQAbstractItemModel_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackQAbstractItemModel_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQAbstractItemModel_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValuea664f1 = new QList<QPersistentModelIndex>(parents); QtCore_PackedList { tmpValuea664f1, tmpValuea664f1->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQAbstractItemModel_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValuea664f1 = new QList<QPersistentModelIndex>(parents); QtCore_PackedList { tmpValuea664f1, tmpValuea664f1->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackQAbstractItemModel_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackQAbstractItemModel_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValuee0adf2 = new QList<QModelIndex>(indexes); QtCore_PackedList { tmpValuee0adf2, tmpValuee0adf2->size() }; }))); };
	QStringList mimeTypes() const { return ({ QtCore_PackedString tempVal = callbackQAbstractItemModel_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackQAbstractItemModel_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackQAbstractItemModel_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQAbstractItemModel_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQAbstractItemModel_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQAbstractItemModel_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackQAbstractItemModel_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackQAbstractItemModel_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackQAbstractItemModel_ResetInternalData(this); };
	void revert() { callbackQAbstractItemModel_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackQAbstractItemModel_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackQAbstractItemModel_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackQAbstractItemModel_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackQAbstractItemModel_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackQAbstractItemModel_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackQAbstractItemModel_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackQAbstractItemModel_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackQAbstractItemModel_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue037c88 = new QMap<int, QVariant>(roles); QtCore_PackedList { tmpValue037c88, tmpValue037c88->size() }; })) != 0; };
	QModelIndex sibling(int row, int column, const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQAbstractItemModel_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&index))); };
	void sort(int column, Qt::SortOrder order) { callbackQAbstractItemModel_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackQAbstractItemModel_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackQAbstractItemModel_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackQAbstractItemModel_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackQAbstractItemModel_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	 ~MyQAbstractItemModel() { callbackQAbstractItemModel_DestroyQAbstractItemModel(this); };
	void childEvent(QChildEvent * event) { callbackQObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQObject_CustomEvent(this, event); };
	void deleteLater() { callbackQObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQObject_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQAbstractItemModel*)

int QAbstractItemModel_QAbstractItemModel_QRegisterMetaType(){qRegisterMetaType<QAbstractItemModel*>(); return qRegisterMetaType<MyQAbstractItemModel*>();}

void* QAbstractItemModel_NewQAbstractItemModel(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractItemModel(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractItemModel(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractItemModel(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractItemModel(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractItemModel(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractItemModel(static_cast<QWindow*>(parent));
	} else {
		return new MyQAbstractItemModel(static_cast<QObject*>(parent));
	}
}

void QAbstractItemModel_BeginResetModel(void* ptr)
{
	static_cast<QAbstractItemModel*>(ptr)->beginResetModel();
}

void* QAbstractItemModel_Buddy(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QAbstractItemModel*>(ptr)->buddy(*static_cast<QModelIndex*>(index)));
}

void* QAbstractItemModel_BuddyDefault(void* ptr, void* index)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return new QModelIndex(static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::buddy(*static_cast<QModelIndex*>(index)));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return new QModelIndex(static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::buddy(*static_cast<QModelIndex*>(index)));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return new QModelIndex(static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::buddy(*static_cast<QModelIndex*>(index)));
	} else {
		return new QModelIndex(static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::buddy(*static_cast<QModelIndex*>(index)));
	}
}

char QAbstractItemModel_CanDropMimeData(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_CanDropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
	} else {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::canDropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
	}
}

char QAbstractItemModel_CanFetchMore(void* ptr, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->canFetchMore(*static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_CanFetchMoreDefault(void* ptr, void* parent)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::canFetchMore(*static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::canFetchMore(*static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::canFetchMore(*static_cast<QModelIndex*>(parent));
	} else {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::canFetchMore(*static_cast<QModelIndex*>(parent));
	}
}

int QAbstractItemModel_ColumnCount(void* ptr, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

void QAbstractItemModel_ConnectColumnsAboutToBeInserted(void* ptr, long long t)
{
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsAboutToBeInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsAboutToBeInserted), static_cast<Qt::ConnectionType>(t));
}

void QAbstractItemModel_DisconnectColumnsAboutToBeInserted(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsAboutToBeInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsAboutToBeInserted));
}

void QAbstractItemModel_ConnectColumnsAboutToBeMoved(void* ptr, long long t)
{
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsAboutToBeMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_ColumnsAboutToBeMoved), static_cast<Qt::ConnectionType>(t));
}

void QAbstractItemModel_DisconnectColumnsAboutToBeMoved(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsAboutToBeMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_ColumnsAboutToBeMoved));
}

void QAbstractItemModel_ConnectColumnsAboutToBeRemoved(void* ptr, long long t)
{
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsAboutToBeRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsAboutToBeRemoved), static_cast<Qt::ConnectionType>(t));
}

void QAbstractItemModel_DisconnectColumnsAboutToBeRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsAboutToBeRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsAboutToBeRemoved));
}

void QAbstractItemModel_ConnectColumnsInserted(void* ptr, long long t)
{
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsInserted), static_cast<Qt::ConnectionType>(t));
}

void QAbstractItemModel_DisconnectColumnsInserted(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsInserted));
}

void QAbstractItemModel_ConnectColumnsMoved(void* ptr, long long t)
{
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_ColumnsMoved), static_cast<Qt::ConnectionType>(t));
}

void QAbstractItemModel_DisconnectColumnsMoved(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_ColumnsMoved));
}

void QAbstractItemModel_ConnectColumnsRemoved(void* ptr, long long t)
{
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsRemoved), static_cast<Qt::ConnectionType>(t));
}

void QAbstractItemModel_DisconnectColumnsRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::columnsRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_ColumnsRemoved));
}

void* QAbstractItemModel_Data(void* ptr, void* index, int role)
{
	return new QVariant(static_cast<QAbstractItemModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

void QAbstractItemModel_ConnectDataChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), static_cast<void (QAbstractItemModel::*)(const QModelIndex &, const QModelIndex &, const QVector<int> &)>(&QAbstractItemModel::dataChanged), static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, const QModelIndex &, const QVector<int> &)>(&MyQAbstractItemModel::Signal_DataChanged), static_cast<Qt::ConnectionType>(t));
}

void QAbstractItemModel_DisconnectDataChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), static_cast<void (QAbstractItemModel::*)(const QModelIndex &, const QModelIndex &, const QVector<int> &)>(&QAbstractItemModel::dataChanged), static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, const QModelIndex &, const QVector<int> &)>(&MyQAbstractItemModel::Signal_DataChanged));
}

void QAbstractItemModel_DataChanged(void* ptr, void* topLeft, void* bottomRight, void* roles)
{
	static_cast<QAbstractItemModel*>(ptr)->dataChanged(*static_cast<QModelIndex*>(topLeft), *static_cast<QModelIndex*>(bottomRight), *static_cast<QVector<int>*>(roles));
}

char QAbstractItemModel_DropMimeData(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_DropMimeDataDefault(void* ptr, void* data, long long action, int row, int column, void* parent)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
	} else {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::dropMimeData(static_cast<QMimeData*>(data), static_cast<Qt::DropAction>(action), row, column, *static_cast<QModelIndex*>(parent));
	}
}

void QAbstractItemModel_EndResetModel(void* ptr)
{
	static_cast<QAbstractItemModel*>(ptr)->endResetModel();
}

void QAbstractItemModel_FetchMore(void* ptr, void* parent)
{
	static_cast<QAbstractItemModel*>(ptr)->fetchMore(*static_cast<QModelIndex*>(parent));
}

void QAbstractItemModel_FetchMoreDefault(void* ptr, void* parent)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::fetchMore(*static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::fetchMore(*static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::fetchMore(*static_cast<QModelIndex*>(parent));
	} else {
		static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::fetchMore(*static_cast<QModelIndex*>(parent));
	}
}

long long QAbstractItemModel_Flags(void* ptr, void* index)
{
	return static_cast<QAbstractItemModel*>(ptr)->flags(*static_cast<QModelIndex*>(index));
}

long long QAbstractItemModel_FlagsDefault(void* ptr, void* index)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::flags(*static_cast<QModelIndex*>(index));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::flags(*static_cast<QModelIndex*>(index));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::flags(*static_cast<QModelIndex*>(index));
	} else {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::flags(*static_cast<QModelIndex*>(index));
	}
}

char QAbstractItemModel_HasChildren(void* ptr, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->hasChildren(*static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_HasChildrenDefault(void* ptr, void* parent)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::hasChildren(*static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::hasChildren(*static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::hasChildren(*static_cast<QModelIndex*>(parent));
	} else {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::hasChildren(*static_cast<QModelIndex*>(parent));
	}
}

void* QAbstractItemModel_HeaderData(void* ptr, int section, long long orientation, int role)
{
	return new QVariant(static_cast<QAbstractItemModel*>(ptr)->headerData(section, static_cast<Qt::Orientation>(orientation), role));
}

void* QAbstractItemModel_HeaderDataDefault(void* ptr, int section, long long orientation, int role)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
	} else {
		return new QVariant(static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::headerData(section, static_cast<Qt::Orientation>(orientation), role));
	}
}

void QAbstractItemModel_ConnectHeaderDataChanged(void* ptr, long long t)
{
	qRegisterMetaType<Qt::Orientation>();
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), static_cast<void (QAbstractItemModel::*)(Qt::Orientation, int, int)>(&QAbstractItemModel::headerDataChanged), static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(Qt::Orientation, int, int)>(&MyQAbstractItemModel::Signal_HeaderDataChanged), static_cast<Qt::ConnectionType>(t));
}

void QAbstractItemModel_DisconnectHeaderDataChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), static_cast<void (QAbstractItemModel::*)(Qt::Orientation, int, int)>(&QAbstractItemModel::headerDataChanged), static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(Qt::Orientation, int, int)>(&MyQAbstractItemModel::Signal_HeaderDataChanged));
}

void QAbstractItemModel_HeaderDataChanged(void* ptr, long long orientation, int first, int last)
{
	static_cast<QAbstractItemModel*>(ptr)->headerDataChanged(static_cast<Qt::Orientation>(orientation), first, last);
}

void* QAbstractItemModel_Index(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<QAbstractItemModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)));
}

char QAbstractItemModel_InsertColumn(void* ptr, int column, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->insertColumn(column, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_InsertColumns(void* ptr, int column, int count, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->insertColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_InsertColumnsDefault(void* ptr, int column, int count, void* parent)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
	} else {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::insertColumns(column, count, *static_cast<QModelIndex*>(parent));
	}
}

char QAbstractItemModel_InsertRow(void* ptr, int row, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->insertRow(row, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_InsertRows(void* ptr, int row, int count, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->insertRows(row, count, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_InsertRowsDefault(void* ptr, int row, int count, void* parent)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
	} else {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::insertRows(row, count, *static_cast<QModelIndex*>(parent));
	}
}

struct QtCore_PackedList QAbstractItemModel_ItemData(void* ptr, void* index)
{
	return ({ QMap<int, QVariant>* tmpValue2d9f39 = new QMap<int, QVariant>(static_cast<QAbstractItemModel*>(ptr)->itemData(*static_cast<QModelIndex*>(index))); QtCore_PackedList { tmpValue2d9f39, tmpValue2d9f39->size() }; });
}

struct QtCore_PackedList QAbstractItemModel_ItemDataDefault(void* ptr, void* index)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return ({ QMap<int, QVariant>* tmpValue46a96f = new QMap<int, QVariant>(static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::itemData(*static_cast<QModelIndex*>(index))); QtCore_PackedList { tmpValue46a96f, tmpValue46a96f->size() }; });
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return ({ QMap<int, QVariant>* tmpValue46a96f = new QMap<int, QVariant>(static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::itemData(*static_cast<QModelIndex*>(index))); QtCore_PackedList { tmpValue46a96f, tmpValue46a96f->size() }; });
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return ({ QMap<int, QVariant>* tmpValue46a96f = new QMap<int, QVariant>(static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::itemData(*static_cast<QModelIndex*>(index))); QtCore_PackedList { tmpValue46a96f, tmpValue46a96f->size() }; });
	} else {
		return ({ QMap<int, QVariant>* tmpValue46a96f = new QMap<int, QVariant>(static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::itemData(*static_cast<QModelIndex*>(index))); QtCore_PackedList { tmpValue46a96f, tmpValue46a96f->size() }; });
	}
}

void QAbstractItemModel_ConnectLayoutAboutToBeChanged(void* ptr, long long t)
{
	qRegisterMetaType<QAbstractItemModel::LayoutChangeHint>();
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), static_cast<void (QAbstractItemModel::*)(const QList<QPersistentModelIndex> &, QAbstractItemModel::LayoutChangeHint)>(&QAbstractItemModel::layoutAboutToBeChanged), static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QList<QPersistentModelIndex> &, QAbstractItemModel::LayoutChangeHint)>(&MyQAbstractItemModel::Signal_LayoutAboutToBeChanged), static_cast<Qt::ConnectionType>(t));
}

void QAbstractItemModel_DisconnectLayoutAboutToBeChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), static_cast<void (QAbstractItemModel::*)(const QList<QPersistentModelIndex> &, QAbstractItemModel::LayoutChangeHint)>(&QAbstractItemModel::layoutAboutToBeChanged), static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QList<QPersistentModelIndex> &, QAbstractItemModel::LayoutChangeHint)>(&MyQAbstractItemModel::Signal_LayoutAboutToBeChanged));
}

void QAbstractItemModel_LayoutAboutToBeChanged(void* ptr, void* parents, long long hint)
{
	static_cast<QAbstractItemModel*>(ptr)->layoutAboutToBeChanged(*static_cast<QList<QPersistentModelIndex>*>(parents), static_cast<QAbstractItemModel::LayoutChangeHint>(hint));
}

void QAbstractItemModel_ConnectLayoutChanged(void* ptr, long long t)
{
	qRegisterMetaType<QAbstractItemModel::LayoutChangeHint>();
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), static_cast<void (QAbstractItemModel::*)(const QList<QPersistentModelIndex> &, QAbstractItemModel::LayoutChangeHint)>(&QAbstractItemModel::layoutChanged), static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QList<QPersistentModelIndex> &, QAbstractItemModel::LayoutChangeHint)>(&MyQAbstractItemModel::Signal_LayoutChanged), static_cast<Qt::ConnectionType>(t));
}

void QAbstractItemModel_DisconnectLayoutChanged(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), static_cast<void (QAbstractItemModel::*)(const QList<QPersistentModelIndex> &, QAbstractItemModel::LayoutChangeHint)>(&QAbstractItemModel::layoutChanged), static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QList<QPersistentModelIndex> &, QAbstractItemModel::LayoutChangeHint)>(&MyQAbstractItemModel::Signal_LayoutChanged));
}

void QAbstractItemModel_LayoutChanged(void* ptr, void* parents, long long hint)
{
	static_cast<QAbstractItemModel*>(ptr)->layoutChanged(*static_cast<QList<QPersistentModelIndex>*>(parents), static_cast<QAbstractItemModel::LayoutChangeHint>(hint));
}

struct QtCore_PackedList QAbstractItemModel_Match(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	return ({ QList<QModelIndex>* tmpValuef30906 = new QList<QModelIndex>(static_cast<QAbstractItemModel*>(ptr)->match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); QtCore_PackedList { tmpValuef30906, tmpValuef30906->size() }; });
}

struct QtCore_PackedList QAbstractItemModel_MatchDefault(void* ptr, void* start, int role, void* value, int hits, long long flags)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return ({ QList<QModelIndex>* tmpValue4d7ea4 = new QList<QModelIndex>(static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); QtCore_PackedList { tmpValue4d7ea4, tmpValue4d7ea4->size() }; });
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return ({ QList<QModelIndex>* tmpValue4d7ea4 = new QList<QModelIndex>(static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); QtCore_PackedList { tmpValue4d7ea4, tmpValue4d7ea4->size() }; });
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return ({ QList<QModelIndex>* tmpValue4d7ea4 = new QList<QModelIndex>(static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); QtCore_PackedList { tmpValue4d7ea4, tmpValue4d7ea4->size() }; });
	} else {
		return ({ QList<QModelIndex>* tmpValue4d7ea4 = new QList<QModelIndex>(static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::match(*static_cast<QModelIndex*>(start), role, *static_cast<QVariant*>(value), hits, static_cast<Qt::MatchFlag>(flags))); QtCore_PackedList { tmpValue4d7ea4, tmpValue4d7ea4->size() }; });
	}
}

void* QAbstractItemModel_MimeData(void* ptr, void* indexes)
{
	return static_cast<QAbstractItemModel*>(ptr)->mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
}

void* QAbstractItemModel_MimeDataDefault(void* ptr, void* indexes)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
	} else {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::mimeData(({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(indexes); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }));
	}
}

struct QtCore_PackedString QAbstractItemModel_MimeTypes(void* ptr)
{
	return ({ QByteArray ta50262 = static_cast<QAbstractItemModel*>(ptr)->mimeTypes().join("¡¦!").toUtf8(); QtCore_PackedString { const_cast<char*>(ta50262.prepend("WHITESPACE").constData()+10), ta50262.size()-10 }; });
}

struct QtCore_PackedString QAbstractItemModel_MimeTypesDefault(void* ptr)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return ({ QByteArray tf2dad9 = static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::mimeTypes().join("¡¦!").toUtf8(); QtCore_PackedString { const_cast<char*>(tf2dad9.prepend("WHITESPACE").constData()+10), tf2dad9.size()-10 }; });
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return ({ QByteArray tf2dad9 = static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::mimeTypes().join("¡¦!").toUtf8(); QtCore_PackedString { const_cast<char*>(tf2dad9.prepend("WHITESPACE").constData()+10), tf2dad9.size()-10 }; });
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return ({ QByteArray tf2dad9 = static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::mimeTypes().join("¡¦!").toUtf8(); QtCore_PackedString { const_cast<char*>(tf2dad9.prepend("WHITESPACE").constData()+10), tf2dad9.size()-10 }; });
	} else {
		return ({ QByteArray tf2dad9 = static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::mimeTypes().join("¡¦!").toUtf8(); QtCore_PackedString { const_cast<char*>(tf2dad9.prepend("WHITESPACE").constData()+10), tf2dad9.size()-10 }; });
	}
}

void QAbstractItemModel_ConnectModelAboutToBeReset(void* ptr, long long t)
{
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::modelAboutToBeReset, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)()>(&MyQAbstractItemModel::Signal_ModelAboutToBeReset), static_cast<Qt::ConnectionType>(t));
}

void QAbstractItemModel_DisconnectModelAboutToBeReset(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::modelAboutToBeReset, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)()>(&MyQAbstractItemModel::Signal_ModelAboutToBeReset));
}

void QAbstractItemModel_ConnectModelReset(void* ptr, long long t)
{
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::modelReset, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)()>(&MyQAbstractItemModel::Signal_ModelReset), static_cast<Qt::ConnectionType>(t));
}

void QAbstractItemModel_DisconnectModelReset(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::modelReset, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)()>(&MyQAbstractItemModel::Signal_ModelReset));
}

char QAbstractItemModel_MoveColumn(void* ptr, void* sourceParent, int sourceColumn, void* destinationParent, int destinationChild)
{
	return static_cast<QAbstractItemModel*>(ptr)->moveColumn(*static_cast<QModelIndex*>(sourceParent), sourceColumn, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char QAbstractItemModel_MoveColumns(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QAbstractItemModel*>(ptr)->moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char QAbstractItemModel_MoveColumnsDefault(void* ptr, void* sourceParent, int sourceColumn, int count, void* destinationParent, int destinationChild)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
	} else {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::moveColumns(*static_cast<QModelIndex*>(sourceParent), sourceColumn, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
	}
}

char QAbstractItemModel_MoveRow(void* ptr, void* sourceParent, int sourceRow, void* destinationParent, int destinationChild)
{
	return static_cast<QAbstractItemModel*>(ptr)->moveRow(*static_cast<QModelIndex*>(sourceParent), sourceRow, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char QAbstractItemModel_MoveRows(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	return static_cast<QAbstractItemModel*>(ptr)->moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
}

char QAbstractItemModel_MoveRowsDefault(void* ptr, void* sourceParent, int sourceRow, int count, void* destinationParent, int destinationChild)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
	} else {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::moveRows(*static_cast<QModelIndex*>(sourceParent), sourceRow, count, *static_cast<QModelIndex*>(destinationParent), destinationChild);
	}
}

void* QAbstractItemModel_Parent(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QAbstractItemModel*>(ptr)->parent(*static_cast<QModelIndex*>(index)));
}

struct QtCore_PackedList QAbstractItemModel_PersistentIndexList(void* ptr)
{
	return ({ QList<QModelIndex>* tmpValue86ddbc = new QList<QModelIndex>(static_cast<QAbstractItemModel*>(ptr)->persistentIndexList()); QtCore_PackedList { tmpValue86ddbc, tmpValue86ddbc->size() }; });
}

char QAbstractItemModel_RemoveColumn(void* ptr, int column, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->removeColumn(column, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_RemoveColumns(void* ptr, int column, int count, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->removeColumns(column, count, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_RemoveColumnsDefault(void* ptr, int column, int count, void* parent)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
	} else {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::removeColumns(column, count, *static_cast<QModelIndex*>(parent));
	}
}

char QAbstractItemModel_RemoveRow(void* ptr, int row, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->removeRow(row, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_RemoveRows(void* ptr, int row, int count, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->removeRows(row, count, *static_cast<QModelIndex*>(parent));
}

char QAbstractItemModel_RemoveRowsDefault(void* ptr, int row, int count, void* parent)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
	} else {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::removeRows(row, count, *static_cast<QModelIndex*>(parent));
	}
}

void QAbstractItemModel_ResetInternalData(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractItemModel*>(ptr), "resetInternalData");
}

void QAbstractItemModel_ResetInternalDataDefault(void* ptr)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::resetInternalData();
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::resetInternalData();
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::resetInternalData();
	} else {
		static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::resetInternalData();
	}
}

void QAbstractItemModel_Revert(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QAbstractItemModel*>(ptr), "revert");
}

void QAbstractItemModel_RevertDefault(void* ptr)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::revert();
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::revert();
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::revert();
	} else {
		static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::revert();
	}
}

struct QtCore_PackedList QAbstractItemModel_RoleNames(void* ptr)
{
	return ({ QHash<int, QByteArray>* tmpValue40dc1d = new QHash<int, QByteArray>(static_cast<QAbstractItemModel*>(ptr)->roleNames()); QtCore_PackedList { tmpValue40dc1d, tmpValue40dc1d->size() }; });
}

struct QtCore_PackedList QAbstractItemModel_RoleNamesDefault(void* ptr)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return ({ QHash<int, QByteArray>* tmpValue42442f = new QHash<int, QByteArray>(static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::roleNames()); QtCore_PackedList { tmpValue42442f, tmpValue42442f->size() }; });
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return ({ QHash<int, QByteArray>* tmpValue42442f = new QHash<int, QByteArray>(static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::roleNames()); QtCore_PackedList { tmpValue42442f, tmpValue42442f->size() }; });
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return ({ QHash<int, QByteArray>* tmpValue42442f = new QHash<int, QByteArray>(static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::roleNames()); QtCore_PackedList { tmpValue42442f, tmpValue42442f->size() }; });
	} else {
		return ({ QHash<int, QByteArray>* tmpValue42442f = new QHash<int, QByteArray>(static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::roleNames()); QtCore_PackedList { tmpValue42442f, tmpValue42442f->size() }; });
	}
}

int QAbstractItemModel_RowCount(void* ptr, void* parent)
{
	return static_cast<QAbstractItemModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

void QAbstractItemModel_ConnectRowsAboutToBeInserted(void* ptr, long long t)
{
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsAboutToBeInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsAboutToBeInserted), static_cast<Qt::ConnectionType>(t));
}

void QAbstractItemModel_DisconnectRowsAboutToBeInserted(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsAboutToBeInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsAboutToBeInserted));
}

void QAbstractItemModel_ConnectRowsAboutToBeMoved(void* ptr, long long t)
{
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsAboutToBeMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_RowsAboutToBeMoved), static_cast<Qt::ConnectionType>(t));
}

void QAbstractItemModel_DisconnectRowsAboutToBeMoved(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsAboutToBeMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_RowsAboutToBeMoved));
}

void QAbstractItemModel_ConnectRowsAboutToBeRemoved(void* ptr, long long t)
{
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsAboutToBeRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsAboutToBeRemoved), static_cast<Qt::ConnectionType>(t));
}

void QAbstractItemModel_DisconnectRowsAboutToBeRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsAboutToBeRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsAboutToBeRemoved));
}

void QAbstractItemModel_ConnectRowsInserted(void* ptr, long long t)
{
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsInserted), static_cast<Qt::ConnectionType>(t));
}

void QAbstractItemModel_DisconnectRowsInserted(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsInserted, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsInserted));
}

void QAbstractItemModel_ConnectRowsMoved(void* ptr, long long t)
{
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_RowsMoved), static_cast<Qt::ConnectionType>(t));
}

void QAbstractItemModel_DisconnectRowsMoved(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsMoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int, const QModelIndex &, int)>(&MyQAbstractItemModel::Signal_RowsMoved));
}

void QAbstractItemModel_ConnectRowsRemoved(void* ptr, long long t)
{
	QObject::connect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsRemoved), static_cast<Qt::ConnectionType>(t));
}

void QAbstractItemModel_DisconnectRowsRemoved(void* ptr)
{
	QObject::disconnect(static_cast<QAbstractItemModel*>(ptr), &QAbstractItemModel::rowsRemoved, static_cast<MyQAbstractItemModel*>(ptr), static_cast<void (MyQAbstractItemModel::*)(const QModelIndex &, int, int)>(&MyQAbstractItemModel::Signal_RowsRemoved));
}

char QAbstractItemModel_SetData(void* ptr, void* index, void* value, int role)
{
	return static_cast<QAbstractItemModel*>(ptr)->setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
}

char QAbstractItemModel_SetDataDefault(void* ptr, void* index, void* value, int role)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
	} else {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::setData(*static_cast<QModelIndex*>(index), *static_cast<QVariant*>(value), role);
	}
}

char QAbstractItemModel_SetHeaderData(void* ptr, int section, long long orientation, void* value, int role)
{
	return static_cast<QAbstractItemModel*>(ptr)->setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
}

char QAbstractItemModel_SetHeaderDataDefault(void* ptr, int section, long long orientation, void* value, int role)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
	} else {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::setHeaderData(section, static_cast<Qt::Orientation>(orientation), *static_cast<QVariant*>(value), role);
	}
}

char QAbstractItemModel_SetItemData(void* ptr, void* index, void* roles)
{
	return static_cast<QAbstractItemModel*>(ptr)->setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
}

char QAbstractItemModel_SetItemDataDefault(void* ptr, void* index, void* roles)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
	} else {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::setItemData(*static_cast<QModelIndex*>(index), *static_cast<QMap<int, QVariant>*>(roles));
	}
}

void* QAbstractItemModel_Sibling(void* ptr, int row, int column, void* index)
{
	return new QModelIndex(static_cast<QAbstractItemModel*>(ptr)->sibling(row, column, *static_cast<QModelIndex*>(index)));
}

void* QAbstractItemModel_SiblingDefault(void* ptr, int row, int column, void* index)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return new QModelIndex(static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::sibling(row, column, *static_cast<QModelIndex*>(index)));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return new QModelIndex(static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::sibling(row, column, *static_cast<QModelIndex*>(index)));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return new QModelIndex(static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::sibling(row, column, *static_cast<QModelIndex*>(index)));
	} else {
		return new QModelIndex(static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::sibling(row, column, *static_cast<QModelIndex*>(index)));
	}
}

void QAbstractItemModel_Sort(void* ptr, int column, long long order)
{
	static_cast<QAbstractItemModel*>(ptr)->sort(column, static_cast<Qt::SortOrder>(order));
}

void QAbstractItemModel_SortDefault(void* ptr, int column, long long order)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::sort(column, static_cast<Qt::SortOrder>(order));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::sort(column, static_cast<Qt::SortOrder>(order));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::sort(column, static_cast<Qt::SortOrder>(order));
	} else {
		static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::sort(column, static_cast<Qt::SortOrder>(order));
	}
}

void* QAbstractItemModel_Span(void* ptr, void* index)
{
	return ({ QSize tmpValue = static_cast<QAbstractItemModel*>(ptr)->span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QAbstractItemModel_SpanDefault(void* ptr, void* index)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSize tmpValue = static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::span(*static_cast<QModelIndex*>(index)); new QSize(tmpValue.width(), tmpValue.height()); });
	}
}

char QAbstractItemModel_Submit(void* ptr)
{
	bool returnArg;
	QMetaObject::invokeMethod(static_cast<QAbstractItemModel*>(ptr), "submit", Q_RETURN_ARG(bool, returnArg));
	return returnArg;
}

char QAbstractItemModel_SubmitDefault(void* ptr)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::submit();
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::submit();
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::submit();
	} else {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::submit();
	}
}

long long QAbstractItemModel_SupportedDragActions(void* ptr)
{
	return static_cast<QAbstractItemModel*>(ptr)->supportedDragActions();
}

long long QAbstractItemModel_SupportedDragActionsDefault(void* ptr)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::supportedDragActions();
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::supportedDragActions();
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::supportedDragActions();
	} else {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::supportedDragActions();
	}
}

long long QAbstractItemModel_SupportedDropActions(void* ptr)
{
	return static_cast<QAbstractItemModel*>(ptr)->supportedDropActions();
}

long long QAbstractItemModel_SupportedDropActionsDefault(void* ptr)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::supportedDropActions();
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::supportedDropActions();
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::supportedDropActions();
	} else {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::supportedDropActions();
	}
}

void QAbstractItemModel_DestroyQAbstractItemModel(void* ptr)
{
	static_cast<QAbstractItemModel*>(ptr)->~QAbstractItemModel();
}

void QAbstractItemModel_DestroyQAbstractItemModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QAbstractItemModel___changePersistentIndexList_from_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___changePersistentIndexList_from_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QAbstractItemModel___changePersistentIndexList_from_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QAbstractItemModel___changePersistentIndexList_to_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___changePersistentIndexList_to_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QAbstractItemModel___changePersistentIndexList_to_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

int QAbstractItemModel___dataChanged_roles_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QVector<int>*>(ptr)->at(i); if (i == static_cast<QVector<int>*>(ptr)->size()-1) { static_cast<QVector<int>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QAbstractItemModel___dataChanged_roles_setList(void* ptr, int i)
{
	static_cast<QVector<int>*>(ptr)->append(i);
}

void* QAbstractItemModel___dataChanged_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<int>();
}

void* QAbstractItemModel___doSetRoleNames_roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___doSetRoleNames_roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* QAbstractItemModel___doSetRoleNames_roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct QtCore_PackedList QAbstractItemModel___doSetRoleNames_roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue7fc3bb = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); QtCore_PackedList { tmpValue7fc3bb, tmpValue7fc3bb->size() }; });
}

void* QAbstractItemModel___encodeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___encodeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QAbstractItemModel___encodeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QAbstractItemModel___itemData_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___itemData_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QAbstractItemModel___itemData_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct QtCore_PackedList QAbstractItemModel___itemData_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue249128 = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValue249128, tmpValue249128->size() }; });
}

void* QAbstractItemModel___layoutAboutToBeChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___layoutAboutToBeChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QAbstractItemModel___layoutAboutToBeChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* QAbstractItemModel___layoutChanged_parents_atList(void* ptr, int i)
{
	return new QPersistentModelIndex(({QPersistentModelIndex tmp = static_cast<QList<QPersistentModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QPersistentModelIndex>*>(ptr)->size()-1) { static_cast<QList<QPersistentModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___layoutChanged_parents_setList(void* ptr, void* i)
{
	static_cast<QList<QPersistentModelIndex>*>(ptr)->append(*static_cast<QPersistentModelIndex*>(i));
}

void* QAbstractItemModel___layoutChanged_parents_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPersistentModelIndex>();
}

void* QAbstractItemModel___match_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___match_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QAbstractItemModel___match_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QAbstractItemModel___mimeData_indexes_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___mimeData_indexes_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QAbstractItemModel___mimeData_indexes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QAbstractItemModel___persistentIndexList_atList(void* ptr, int i)
{
	return new QModelIndex(({QModelIndex tmp = static_cast<QList<QModelIndex>*>(ptr)->at(i); if (i == static_cast<QList<QModelIndex>*>(ptr)->size()-1) { static_cast<QList<QModelIndex>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___persistentIndexList_setList(void* ptr, void* i)
{
	static_cast<QList<QModelIndex>*>(ptr)->append(*static_cast<QModelIndex*>(i));
}

void* QAbstractItemModel___persistentIndexList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QModelIndex>();
}

void* QAbstractItemModel___roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* QAbstractItemModel___roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct QtCore_PackedList QAbstractItemModel___roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue7fc3bb = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); QtCore_PackedList { tmpValue7fc3bb, tmpValue7fc3bb->size() }; });
}

void* QAbstractItemModel___setItemData_roles_atList(void* ptr, int v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<int, QVariant>*>(ptr)->value(v); if (i == static_cast<QMap<int, QVariant>*>(ptr)->size()-1) { static_cast<QMap<int, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___setItemData_roles_setList(void* ptr, int key, void* i)
{
	static_cast<QMap<int, QVariant>*>(ptr)->insert(key, *static_cast<QVariant*>(i));
}

void* QAbstractItemModel___setItemData_roles_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<int, QVariant>();
}

struct QtCore_PackedList QAbstractItemModel___setItemData_roles_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue249128 = new QList<int>(static_cast<QMap<int, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValue249128, tmpValue249128->size() }; });
}

void* QAbstractItemModel___setRoleNames_roleNames_atList(void* ptr, int v, int i)
{
	return new QByteArray(({ QByteArray tmp = static_cast<QHash<int, QByteArray>*>(ptr)->value(v); if (i == static_cast<QHash<int, QByteArray>*>(ptr)->size()-1) { static_cast<QHash<int, QByteArray>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QAbstractItemModel___setRoleNames_roleNames_setList(void* ptr, int key, void* i)
{
	static_cast<QHash<int, QByteArray>*>(ptr)->insert(key, *static_cast<QByteArray*>(i));
}

void* QAbstractItemModel___setRoleNames_roleNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<int, QByteArray>();
}

struct QtCore_PackedList QAbstractItemModel___setRoleNames_roleNames_keyList(void* ptr)
{
	return ({ QList<int>* tmpValue7fc3bb = new QList<int>(static_cast<QHash<int, QByteArray>*>(ptr)->keys()); QtCore_PackedList { tmpValue7fc3bb, tmpValue7fc3bb->size() }; });
}

int QAbstractItemModel_____doSetRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractItemModel_____doSetRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QAbstractItemModel_____doSetRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QAbstractItemModel_____itemData_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractItemModel_____itemData_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QAbstractItemModel_____itemData_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QAbstractItemModel_____roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractItemModel_____roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QAbstractItemModel_____roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QAbstractItemModel_____setItemData_roles_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractItemModel_____setItemData_roles_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QAbstractItemModel_____setItemData_roles_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

int QAbstractItemModel_____setRoleNames_roleNames_keyList_atList(void* ptr, int i)
{
	return ({int tmp = static_cast<QList<int>*>(ptr)->at(i); if (i == static_cast<QList<int>*>(ptr)->size()-1) { static_cast<QList<int>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QAbstractItemModel_____setRoleNames_roleNames_keyList_setList(void* ptr, int i)
{
	static_cast<QList<int>*>(ptr)->append(i);
}

void* QAbstractItemModel_____setRoleNames_roleNames_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<int>();
}

class MyQAbstractProxyModel: public QAbstractProxyModel
{
public:
	MyQAbstractProxyModel(QObject *parent = Q_NULLPTR) : QAbstractProxyModel(parent) {QAbstractProxyModel_QAbstractProxyModel_QRegisterMetaType();};
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQAbstractItemModel_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackQAbstractItemModel_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackQAbstractItemModel_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant data(const QModelIndex & proxyIndex, int role) const { return *static_cast<QVariant*>(callbackQAbstractProxyModel_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&proxyIndex), role)); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackQAbstractItemModel_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	void fetchMore(const QModelIndex & parent) { callbackQAbstractItemModel_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackQAbstractItemModel_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool hasChildren(const QModelIndex & parent) const { return callbackQAbstractItemModel_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackQAbstractItemModel_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	QMap<int, QVariant> itemData(const QModelIndex & proxyIndex) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackQAbstractItemModel_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&proxyIndex))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	QModelIndex mapFromSource(const QModelIndex & sourceIndex) const { return *static_cast<QModelIndex*>(callbackQAbstractProxyModel_MapFromSource(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&sourceIndex))); };
	QModelIndex mapToSource(const QModelIndex & proxyIndex) const { return *static_cast<QModelIndex*>(callbackQAbstractProxyModel_MapToSource(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&proxyIndex))); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackQAbstractItemModel_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValuee0adf2 = new QList<QModelIndex>(indexes); QtCore_PackedList { tmpValuee0adf2, tmpValuee0adf2->size() }; }))); };
	QStringList mimeTypes() const { return ({ QtCore_PackedString tempVal = callbackQAbstractItemModel_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void resetInternalData() { callbackQAbstractItemModel_ResetInternalData(this); };
	void revert() { callbackQAbstractProxyModel_Revert(this); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackQAbstractItemModel_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackQAbstractItemModel_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackQAbstractItemModel_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue037c88 = new QMap<int, QVariant>(roles); QtCore_PackedList { tmpValue037c88, tmpValue037c88->size() }; })) != 0; };
	void setSourceModel(QAbstractItemModel * sourceModel) { callbackQAbstractProxyModel_SetSourceModel(this, sourceModel); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackQAbstractItemModel_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	void sort(int column, Qt::SortOrder order) { callbackQAbstractItemModel_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackQAbstractItemModel_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackQAbstractProxyModel_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackQAbstractItemModel_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackQAbstractItemModel_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	 ~MyQAbstractProxyModel() { callbackQAbstractProxyModel_DestroyQAbstractProxyModel(this); };
	int columnCount(const QModelIndex & parent) const { return callbackQAbstractProxyModel_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackQAbstractItemModel_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackQAbstractItemModel_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackQAbstractItemModel_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue037c88 = new QVector<int>(roles); QtCore_PackedList { tmpValue037c88, tmpValue037c88->size() }; })); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackQAbstractItemModel_HeaderDataChanged(this, orientation, first, last); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackQAbstractProxyModel_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackQAbstractItemModel_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackQAbstractItemModel_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQAbstractItemModel_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValuea664f1 = new QList<QPersistentModelIndex>(parents); QtCore_PackedList { tmpValuea664f1, tmpValuea664f1->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQAbstractItemModel_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValuea664f1 = new QList<QPersistentModelIndex>(parents); QtCore_PackedList { tmpValuea664f1, tmpValuea664f1->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackQAbstractItemModel_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	void Signal_ModelAboutToBeReset() { callbackQAbstractItemModel_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackQAbstractItemModel_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQAbstractItemModel_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQAbstractItemModel_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQAbstractProxyModel_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackQAbstractItemModel_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackQAbstractItemModel_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackQAbstractItemModel_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackQAbstractProxyModel_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackQAbstractItemModel_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackQAbstractItemModel_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackQAbstractItemModel_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void childEvent(QChildEvent * event) { callbackQObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQObject_CustomEvent(this, event); };
	void deleteLater() { callbackQObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQObject_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQAbstractProxyModel*)

int QAbstractProxyModel_QAbstractProxyModel_QRegisterMetaType(){qRegisterMetaType<QAbstractProxyModel*>(); return qRegisterMetaType<MyQAbstractProxyModel*>();}

void* QAbstractProxyModel_NewQAbstractProxyModel(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractProxyModel(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractProxyModel(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractProxyModel(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractProxyModel(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractProxyModel(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractProxyModel(static_cast<QWindow*>(parent));
	} else {
		return new MyQAbstractProxyModel(static_cast<QObject*>(parent));
	}
}

void* QAbstractProxyModel_Data(void* ptr, void* proxyIndex, int role)
{
	return new QVariant(static_cast<QAbstractProxyModel*>(ptr)->data(*static_cast<QModelIndex*>(proxyIndex), role));
}

void* QAbstractProxyModel_DataDefault(void* ptr, void* proxyIndex, int role)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return new QVariant(static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::data(*static_cast<QModelIndex*>(proxyIndex), role));
	} else {
		return new QVariant(static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::data(*static_cast<QModelIndex*>(proxyIndex), role));
	}
}

void* QAbstractProxyModel_MapFromSource(void* ptr, void* sourceIndex)
{
	return new QModelIndex(static_cast<QAbstractProxyModel*>(ptr)->mapFromSource(*static_cast<QModelIndex*>(sourceIndex)));
}

void* QAbstractProxyModel_MapToSource(void* ptr, void* proxyIndex)
{
	return new QModelIndex(static_cast<QAbstractProxyModel*>(ptr)->mapToSource(*static_cast<QModelIndex*>(proxyIndex)));
}

void QAbstractProxyModel_Revert(void* ptr)
{
	static_cast<QAbstractProxyModel*>(ptr)->revert();
}

void QAbstractProxyModel_RevertDefault(void* ptr)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::revert();
	} else {
		static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::revert();
	}
}

void QAbstractProxyModel_SetSourceModel(void* ptr, void* sourceModel)
{
	static_cast<QAbstractProxyModel*>(ptr)->setSourceModel(static_cast<QAbstractItemModel*>(sourceModel));
}

void QAbstractProxyModel_SetSourceModelDefault(void* ptr, void* sourceModel)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::setSourceModel(static_cast<QAbstractItemModel*>(sourceModel));
	} else {
		static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::setSourceModel(static_cast<QAbstractItemModel*>(sourceModel));
	}
}

void* QAbstractProxyModel_SourceModel(void* ptr)
{
	return static_cast<QAbstractProxyModel*>(ptr)->sourceModel();
}

char QAbstractProxyModel_Submit(void* ptr)
{
	return static_cast<QAbstractProxyModel*>(ptr)->submit();
}

char QAbstractProxyModel_SubmitDefault(void* ptr)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::submit();
	} else {
		return static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::submit();
	}
}

void QAbstractProxyModel_DestroyQAbstractProxyModel(void* ptr)
{
	static_cast<QAbstractProxyModel*>(ptr)->~QAbstractProxyModel();
}

void QAbstractProxyModel_DestroyQAbstractProxyModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

int QAbstractProxyModel_ColumnCount(void* ptr, void* parent)
{
	return static_cast<QAbstractProxyModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

int QAbstractProxyModel_ColumnCountDefault(void* ptr, void* parent)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::columnCount(*static_cast<QModelIndex*>(parent));
	} else {
	
	}
}

void* QAbstractProxyModel_Index(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<QAbstractProxyModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* QAbstractProxyModel_IndexDefault(void* ptr, int row, int column, void* parent)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return new QModelIndex(static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::index(row, column, *static_cast<QModelIndex*>(parent)));
	} else {
	
	}
}

void* QAbstractProxyModel_Parent(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QAbstractProxyModel*>(ptr)->parent(*static_cast<QModelIndex*>(index)));
}

void* QAbstractProxyModel_ParentDefault(void* ptr, void* index)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return new QModelIndex(static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::parent(*static_cast<QModelIndex*>(index)));
	} else {
	
	}
}

int QAbstractProxyModel_RowCount(void* ptr, void* parent)
{
	return static_cast<QAbstractProxyModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QAbstractProxyModel_RowCountDefault(void* ptr, void* parent)
{
	if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::rowCount(*static_cast<QModelIndex*>(parent));
	} else {
	
	}
}

class MyQAbstractTableModel: public QAbstractTableModel
{
public:
	MyQAbstractTableModel(QObject *parent = Q_NULLPTR) : QAbstractTableModel(parent) {QAbstractTableModel_QAbstractTableModel_QRegisterMetaType();};
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackQAbstractItemModel_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackQAbstractItemModel_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackQAbstractTableModel_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackQAbstractItemModel_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	 ~MyQAbstractTableModel() { callbackQAbstractTableModel_DestroyQAbstractTableModel(this); };
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQAbstractItemModel_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackQAbstractItemModel_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	bool canFetchMore(const QModelIndex & parent) const { return callbackQAbstractItemModel_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackQAbstractTableModel_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackQAbstractItemModel_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackQAbstractItemModel_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackQAbstractTableModel_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackQAbstractItemModel_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue037c88 = new QVector<int>(roles); QtCore_PackedList { tmpValue037c88, tmpValue037c88->size() }; })); };
	void fetchMore(const QModelIndex & parent) { callbackQAbstractItemModel_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	bool hasChildren(const QModelIndex & parent) const { return callbackQAbstractItemModel_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackQAbstractItemModel_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackQAbstractItemModel_HeaderDataChanged(this, orientation, first, last); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackQAbstractItemModel_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackQAbstractItemModel_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & index) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackQAbstractItemModel_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQAbstractItemModel_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValuea664f1 = new QList<QPersistentModelIndex>(parents); QtCore_PackedList { tmpValuea664f1, tmpValuea664f1->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQAbstractItemModel_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValuea664f1 = new QList<QPersistentModelIndex>(parents); QtCore_PackedList { tmpValuea664f1, tmpValuea664f1->size() }; }), hint); };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackQAbstractItemModel_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackQAbstractItemModel_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValuee0adf2 = new QList<QModelIndex>(indexes); QtCore_PackedList { tmpValuee0adf2, tmpValuee0adf2->size() }; }))); };
	QStringList mimeTypes() const { return ({ QtCore_PackedString tempVal = callbackQAbstractItemModel_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	void Signal_ModelAboutToBeReset() { callbackQAbstractItemModel_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackQAbstractItemModel_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQAbstractItemModel_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQAbstractItemModel_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QModelIndex parent(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQAbstractTableModel_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackQAbstractItemModel_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackQAbstractItemModel_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	void resetInternalData() { callbackQAbstractItemModel_ResetInternalData(this); };
	void revert() { callbackQAbstractItemModel_Revert(this); };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackQAbstractItemModel_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	int rowCount(const QModelIndex & parent) const { return callbackQAbstractTableModel_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackQAbstractItemModel_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackQAbstractItemModel_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackQAbstractItemModel_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackQAbstractItemModel_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackQAbstractItemModel_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackQAbstractItemModel_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue037c88 = new QMap<int, QVariant>(roles); QtCore_PackedList { tmpValue037c88, tmpValue037c88->size() }; })) != 0; };
	void sort(int column, Qt::SortOrder order) { callbackQAbstractItemModel_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackQAbstractItemModel_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool submit() { return callbackQAbstractItemModel_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackQAbstractItemModel_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackQAbstractItemModel_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void childEvent(QChildEvent * event) { callbackQObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQObject_CustomEvent(this, event); };
	void deleteLater() { callbackQObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQObject_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQAbstractTableModel*)

int QAbstractTableModel_QAbstractTableModel_QRegisterMetaType(){qRegisterMetaType<QAbstractTableModel*>(); return qRegisterMetaType<MyQAbstractTableModel*>();}

void* QAbstractTableModel_NewQAbstractTableModel(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractTableModel(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractTableModel(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractTableModel(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractTableModel(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractTableModel(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQAbstractTableModel(static_cast<QWindow*>(parent));
	} else {
		return new MyQAbstractTableModel(static_cast<QObject*>(parent));
	}
}

void* QAbstractTableModel_Index(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<QAbstractTableModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* QAbstractTableModel_IndexDefault(void* ptr, int row, int column, void* parent)
{
		return new QModelIndex(static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void QAbstractTableModel_DestroyQAbstractTableModel(void* ptr)
{
	static_cast<QAbstractTableModel*>(ptr)->~QAbstractTableModel();
}

void QAbstractTableModel_DestroyQAbstractTableModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

int QAbstractTableModel_ColumnCount(void* ptr, void* parent)
{
	return static_cast<QAbstractTableModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

int QAbstractTableModel_ColumnCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);
	
}

void* QAbstractTableModel_Data(void* ptr, void* index, int role)
{
	return new QVariant(static_cast<QAbstractTableModel*>(ptr)->data(*static_cast<QModelIndex*>(index), role));
}

void* QAbstractTableModel_DataDefault(void* ptr, void* index, int role)
{
	Q_UNUSED(ptr);
	Q_UNUSED(index);
	Q_UNUSED(role);
	
}

void* QAbstractTableModel_Parent(void* ptr, void* index)
{
	return new QModelIndex(static_cast<QAbstractTableModel*>(ptr)->parent(*static_cast<QModelIndex*>(index)));
}

void* QAbstractTableModel_ParentDefault(void* ptr, void* index)
{
		return new QModelIndex(static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::parent(*static_cast<QModelIndex*>(index)));
}

int QAbstractTableModel_RowCount(void* ptr, void* parent)
{
	return static_cast<QAbstractTableModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QAbstractTableModel_RowCountDefault(void* ptr, void* parent)
{
	Q_UNUSED(ptr);
	Q_UNUSED(parent);
	
}

void* QBitArray_NewQBitArray()
{
	return new QBitArray();
}

void* QBitArray_NewQBitArray2(int size, char value)
{
	return new QBitArray(size, value != 0);
}

void* QBitArray_NewQBitArray3(void* other)
{
	return new QBitArray(*static_cast<QBitArray*>(other));
}

void* QBitArray_NewQBitArray4(void* other)
{
	return new QBitArray(*static_cast<QBitArray*>(other));
}

char QBitArray_At(void* ptr, int i)
{
	return static_cast<QBitArray*>(ptr)->at(i);
}

int QBitArray_Count(void* ptr)
{
	return static_cast<QBitArray*>(ptr)->count();
}

int QBitArray_Count2(void* ptr, char on)
{
	return static_cast<QBitArray*>(ptr)->count(on != 0);
}

char QBitArray_Fill(void* ptr, char value, int size)
{
	return static_cast<QBitArray*>(ptr)->fill(value != 0, size);
}

void QBitArray_Fill2(void* ptr, char value, int begin, int end)
{
	static_cast<QBitArray*>(ptr)->fill(value != 0, begin, end);
}

void QBitArray_Resize(void* ptr, int size)
{
	static_cast<QBitArray*>(ptr)->resize(size);
}

int QBitArray_Size(void* ptr)
{
	return static_cast<QBitArray*>(ptr)->size();
}

void* QByteArray_NewQByteArray()
{
	return new QByteArray();
}

void* QByteArray_NewQByteArray2(char* data, int size)
{
	return new QByteArray(const_cast<const char*>(data), size);
}

void* QByteArray_NewQByteArray3(int size, char* ch)
{
	return new QByteArray(size, *ch);
}

void* QByteArray_NewQByteArray4(void* other)
{
	return new QByteArray(*static_cast<QByteArray*>(other));
}

struct QtCore_PackedString QByteArray_At(void* ptr, int i)
{
	return ({ char t8d8fc5 = static_cast<QByteArray*>(ptr)->at(i); QtCore_PackedString { &t8d8fc5, -1 }; });
}

struct QtCore_PackedString QByteArray_Back(void* ptr)
{
	return ({ char te8eea8 = static_cast<QByteArray*>(ptr)->back(); QtCore_PackedString { &te8eea8, -1 }; });
}

char QByteArray_Contains(void* ptr, void* ba)
{
	return static_cast<QByteArray*>(ptr)->contains(*static_cast<QByteArray*>(ba));
}

char QByteArray_Contains2(void* ptr, char* ch)
{
	return static_cast<QByteArray*>(ptr)->contains(*ch);
}

char QByteArray_Contains3(void* ptr, char* str)
{
	return static_cast<QByteArray*>(ptr)->contains(const_cast<const char*>(str));
}

int QByteArray_Count(void* ptr, void* ba)
{
	return static_cast<QByteArray*>(ptr)->count(*static_cast<QByteArray*>(ba));
}

int QByteArray_Count2(void* ptr, char* ch)
{
	return static_cast<QByteArray*>(ptr)->count(*ch);
}

int QByteArray_Count3(void* ptr, char* str)
{
	return static_cast<QByteArray*>(ptr)->count(const_cast<const char*>(str));
}

int QByteArray_Count4(void* ptr)
{
	return static_cast<QByteArray*>(ptr)->count();
}

struct QtCore_PackedString QByteArray_Data(void* ptr)
{
	return QtCore_PackedString { static_cast<QByteArray*>(ptr)->data(), static_cast<QByteArray*>(ptr)->size() };
}

struct QtCore_PackedString QByteArray_Data2(void* ptr)
{
	return QtCore_PackedString { const_cast<char*>(static_cast<QByteArray*>(ptr)->data()), static_cast<QByteArray*>(ptr)->size() };
}

void* QByteArray_Fill(void* ptr, char* ch, int size)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->fill(*ch, size));
}

void* QByteArray_Insert(void* ptr, int i, void* ba)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->insert(i, *static_cast<QByteArray*>(ba)));
}

void* QByteArray_Insert2(void* ptr, int i, char* ch)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->insert(i, *ch));
}

void* QByteArray_Insert3(void* ptr, int i, int count, char* ch)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->insert(i, count, *ch));
}

void* QByteArray_Insert4(void* ptr, int i, char* str)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->insert(i, const_cast<const char*>(str)));
}

void* QByteArray_Insert5(void* ptr, int i, char* str, int l)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->insert(i, const_cast<const char*>(str), l));
}

void* QByteArray_Insert6(void* ptr, int i, struct QtCore_PackedString str)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->insert(i, QString::fromUtf8(str.data, str.len)));
}

void* QByteArray_Left(void* ptr, int l)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->left(l));
}

void* QByteArray_Remove(void* ptr, int pos, int l)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->remove(pos, l));
}

void QByteArray_Resize(void* ptr, int size)
{
	static_cast<QByteArray*>(ptr)->resize(size);
}

void* QByteArray_Right(void* ptr, int l)
{
	return new QByteArray(static_cast<QByteArray*>(ptr)->right(l));
}

int QByteArray_Size(void* ptr)
{
	return static_cast<QByteArray*>(ptr)->size();
}

struct QtCore_PackedList QByteArray_Split(void* ptr, char* sep)
{
	return ({ QList<QByteArray>* tmpValue17cac8 = new QList<QByteArray>(static_cast<QByteArray*>(ptr)->split(*sep)); QtCore_PackedList { tmpValue17cac8, tmpValue17cac8->size() }; });
}

int QByteArray_ToInt(void* ptr, char* ok, int base)
{
	return static_cast<QByteArray*>(ptr)->toInt(reinterpret_cast<bool*>(ok), base);
}

void QByteArray_DestroyQByteArray(void* ptr)
{
	static_cast<QByteArray*>(ptr)->~QByteArray();
}

void* QByteArray___split_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QByteArray___split_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QByteArray___split_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

int QCborValue_False_Type()
{
	return QCborValue::False;
}

int QCborValue_True_Type()
{
	return QCborValue::True;
}

int QCborValue_Null_Type()
{
	return QCborValue::Null;
}

int QCborValue_Undefined_Type()
{
	return QCborValue::Undefined;
}

void* QChar_NewQChar()
{
	return new QChar();
}

void* QChar_NewQChar2(unsigned short code)
{
	return new QChar(code);
}

void* QChar_NewQChar3(char* cell, char* row)
{
	return new QChar(*static_cast<uchar*>(static_cast<void*>(cell)), *static_cast<uchar*>(static_cast<void*>(row)));
}

void* QChar_NewQChar4(short code)
{
	return new QChar(code);
}

void* QChar_NewQChar5(unsigned int code)
{
	return new QChar(code);
}

void* QChar_NewQChar6(int code)
{
	return new QChar(code);
}

void* QChar_NewQChar7(long long ch)
{
	return new QChar(static_cast<QChar::SpecialCharacter>(ch));
}

void* QChar_NewQChar8(void* ch)
{
	return new QChar(*static_cast<QLatin1Char*>(ch));
}

void* QChar_NewQChar11(char* ch)
{
	return new QChar(*ch);
}

void* QChar_NewQChar12(char* ch)
{
	return new QChar(*static_cast<uchar*>(static_cast<void*>(ch)));
}

struct QtCore_PackedString QChar_Row(void* ptr)
{
	return ({ uchar pret52b865 = static_cast<QChar*>(ptr)->row(); char* t52b865 = static_cast<char*>(static_cast<void*>(&pret52b865)); QtCore_PackedString { t52b865, -1 }; });
}

class MyQChildEvent: public QChildEvent
{
public:
	MyQChildEvent(QEvent::Type ty, QObject *child) : QChildEvent(ty, child) {};
};

void* QChildEvent_NewQChildEvent(long long ty, void* child)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(child))) {
		return new MyQChildEvent(static_cast<QEvent::Type>(ty), static_cast<QGraphicsObject*>(child));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(child))) {
		return new MyQChildEvent(static_cast<QEvent::Type>(ty), static_cast<QGraphicsWidget*>(child));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(child))) {
		return new MyQChildEvent(static_cast<QEvent::Type>(ty), static_cast<QLayout*>(child));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(child))) {
		return new MyQChildEvent(static_cast<QEvent::Type>(ty), static_cast<QQuickItem*>(child));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(child))) {
		return new MyQChildEvent(static_cast<QEvent::Type>(ty), static_cast<QWidget*>(child));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(child))) {
		return new MyQChildEvent(static_cast<QEvent::Type>(ty), static_cast<QWindow*>(child));
	} else {
		return new MyQChildEvent(static_cast<QEvent::Type>(ty), static_cast<QObject*>(child));
	}
}

void* QChildEvent_Child(void* ptr)
{
	return static_cast<QChildEvent*>(ptr)->child();
}

char QChildEvent_Removed(void* ptr)
{
	return static_cast<QChildEvent*>(ptr)->removed();
}

class MyQCoreApplication: public QCoreApplication
{
public:
	MyQCoreApplication(int &argc, char **argv) : QCoreApplication(argc, argv) {QCoreApplication_QCoreApplication_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQObject_Event(this, e) != 0; };
	 ~MyQCoreApplication() { callbackQCoreApplication_DestroyQCoreApplication(this); };
	void childEvent(QChildEvent * event) { callbackQObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQObject_CustomEvent(this, event); };
	void deleteLater() { callbackQObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQObject_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQCoreApplication*)

int QCoreApplication_QCoreApplication_QRegisterMetaType(){qRegisterMetaType<QCoreApplication*>(); return qRegisterMetaType<MyQCoreApplication*>();}

void* QCoreApplication_NewQCoreApplication(int argc, char* argv)
{
	static int argcs = argc;
	static char** argvs = static_cast<char**>(malloc(argcs * sizeof(char*)));

	QList<QByteArray> aList = QByteArray(argv).split('|');
	for (int i = 0; i < argcs; i++)
		argvs[i] = (new QByteArray(aList.at(i)))->data();

	return new MyQCoreApplication(argcs, argvs);
}

int QCoreApplication_QCoreApplication_Exec()
{
	return QCoreApplication::exec();
}

void QCoreApplication_QCoreApplication_SetAttribute(long long attribute, char on)
{
	QCoreApplication::setAttribute(static_cast<Qt::ApplicationAttribute>(attribute), on != 0);
}

void QCoreApplication_DestroyQCoreApplication(void* ptr)
{
	static_cast<QCoreApplication*>(ptr)->~QCoreApplication();
}

void QCoreApplication_DestroyQCoreApplicationDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QDataStream_NewQDataStream()
{
	return new QDataStream();
}

void* QDataStream_NewQDataStream2(void* d)
{
	return new QDataStream(static_cast<QIODevice*>(d));
}

void* QDataStream_NewQDataStream3(void* a, long long mode)
{
	return new QDataStream(static_cast<QByteArray*>(a), static_cast<QIODevice::OpenModeFlag>(mode));
}

void* QDataStream_NewQDataStream4(void* a)
{
	return new QDataStream(*static_cast<QByteArray*>(a));
}

long long QDataStream_Status(void* ptr)
{
	return static_cast<QDataStream*>(ptr)->status();
}

void QDataStream_DestroyQDataStream(void* ptr)
{
	static_cast<QDataStream*>(ptr)->~QDataStream();
}

void* QDate_NewQDate2()
{
	return new QDate();
}

void* QDate_NewQDate3(int y, int m, int d)
{
	return new QDate(y, m, d);
}

void* QDateTime_NewQDateTime()
{
	return new QDateTime();
}

void* QDateTime_NewQDateTime2(void* date)
{
	return new QDateTime(*static_cast<QDate*>(date));
}

void* QDateTime_NewQDateTime3(void* date, void* ti, long long spec)
{
	return new QDateTime(*static_cast<QDate*>(date), *static_cast<QTime*>(ti), static_cast<Qt::TimeSpec>(spec));
}

void* QDateTime_NewQDateTime4(void* date, void* ti, long long spec, int offsetSeconds)
{
	return new QDateTime(*static_cast<QDate*>(date), *static_cast<QTime*>(ti), static_cast<Qt::TimeSpec>(spec), offsetSeconds);
}

void* QDateTime_NewQDateTime5(void* date, void* ti, void* timeZone)
{
	return new QDateTime(*static_cast<QDate*>(date), *static_cast<QTime*>(ti), *static_cast<QTimeZone*>(timeZone));
}

void* QDateTime_NewQDateTime6(void* other)
{
	return new QDateTime(*static_cast<QDateTime*>(other));
}

void* QDateTime_NewQDateTime7(void* other)
{
	return new QDateTime(*static_cast<QDateTime*>(other));
}

void* QDateTime_Date(void* ptr)
{
	return new QDate(static_cast<QDateTime*>(ptr)->date());
}

void* QDateTime_Time(void* ptr)
{
	return new QTime(static_cast<QDateTime*>(ptr)->time());
}

void QDateTime_DestroyQDateTime(void* ptr)
{
	static_cast<QDateTime*>(ptr)->~QDateTime();
}

void* QDir_NewQDir(void* dir)
{
	return new QDir(*static_cast<QDir*>(dir));
}

void* QDir_NewQDir2(struct QtCore_PackedString path)
{
	return new QDir(QString::fromUtf8(path.data, path.len));
}

void* QDir_NewQDir3(struct QtCore_PackedString path, struct QtCore_PackedString nameFilter, long long sort, long long filters)
{
	return new QDir(QString::fromUtf8(path.data, path.len), QString::fromUtf8(nameFilter.data, nameFilter.len), static_cast<QDir::SortFlag>(sort), static_cast<QDir::Filter>(filters));
}

unsigned int QDir_Count(void* ptr)
{
	return static_cast<QDir*>(ptr)->count();
}

void* QDir_QDir_Current()
{
	return new QDir(QDir::current());
}

char QDir_Exists(void* ptr, struct QtCore_PackedString name)
{
	return static_cast<QDir*>(ptr)->exists(QString::fromUtf8(name.data, name.len));
}

char QDir_Exists2(void* ptr)
{
	return static_cast<QDir*>(ptr)->exists();
}

long long QDir_Filter(void* ptr)
{
	return static_cast<QDir*>(ptr)->filter();
}

char QDir_QDir_Match(struct QtCore_PackedString filter, struct QtCore_PackedString fileName)
{
	return QDir::match(QString::fromUtf8(filter.data, filter.len), QString::fromUtf8(fileName.data, fileName.len));
}

char QDir_QDir_Match2(struct QtCore_PackedString filters, struct QtCore_PackedString fileName)
{
	return QDir::match(QString::fromUtf8(filters.data, filters.len).split("¡¦!", QString::SkipEmptyParts), QString::fromUtf8(fileName.data, fileName.len));
}

struct QtCore_PackedString QDir_Path(void* ptr)
{
	return ({ QByteArray t1e0939 = static_cast<QDir*>(ptr)->path().toUtf8(); QtCore_PackedString { const_cast<char*>(t1e0939.prepend("WHITESPACE").constData()+10), t1e0939.size()-10 }; });
}

char QDir_Remove(void* ptr, struct QtCore_PackedString fileName)
{
	return static_cast<QDir*>(ptr)->remove(QString::fromUtf8(fileName.data, fileName.len));
}

void* QDir_QDir_Root()
{
	return new QDir(QDir::root());
}

char QDir_QDir_SetCurrent(struct QtCore_PackedString path)
{
	return QDir::setCurrent(QString::fromUtf8(path.data, path.len));
}

void QDir_SetFilter(void* ptr, long long filters)
{
	static_cast<QDir*>(ptr)->setFilter(static_cast<QDir::Filter>(filters));
}

void* QDir_QDir_Temp()
{
	return new QDir(QDir::temp());
}

void QDir_DestroyQDir(void* ptr)
{
	static_cast<QDir*>(ptr)->~QDir();
}

void* QDir___drives_atList(void* ptr, int i)
{
	return new QFileInfo(({QFileInfo tmp = static_cast<QList<QFileInfo>*>(ptr)->at(i); if (i == static_cast<QList<QFileInfo>*>(ptr)->size()-1) { static_cast<QList<QFileInfo>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QDir___drives_setList(void* ptr, void* i)
{
	static_cast<QList<QFileInfo>*>(ptr)->append(*static_cast<QFileInfo*>(i));
}

void* QDir___drives_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QFileInfo>();
}

void* QDir___entryInfoList_atList(void* ptr, int i)
{
	return new QFileInfo(({QFileInfo tmp = static_cast<QList<QFileInfo>*>(ptr)->at(i); if (i == static_cast<QList<QFileInfo>*>(ptr)->size()-1) { static_cast<QList<QFileInfo>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QDir___entryInfoList_setList(void* ptr, void* i)
{
	static_cast<QList<QFileInfo>*>(ptr)->append(*static_cast<QFileInfo*>(i));
}

void* QDir___entryInfoList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QFileInfo>();
}

void* QDir___entryInfoList_atList2(void* ptr, int i)
{
	return new QFileInfo(({QFileInfo tmp = static_cast<QList<QFileInfo>*>(ptr)->at(i); if (i == static_cast<QList<QFileInfo>*>(ptr)->size()-1) { static_cast<QList<QFileInfo>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QDir___entryInfoList_setList2(void* ptr, void* i)
{
	static_cast<QList<QFileInfo>*>(ptr)->append(*static_cast<QFileInfo*>(i));
}

void* QDir___entryInfoList_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QFileInfo>();
}

void* QEasingCurve_NewQEasingCurve(long long ty)
{
	return new QEasingCurve(static_cast<QEasingCurve::Type>(ty));
}

void* QEasingCurve_NewQEasingCurve2(void* other)
{
	return new QEasingCurve(*static_cast<QEasingCurve*>(other));
}

void* QEasingCurve_NewQEasingCurve3(void* other)
{
	return new QEasingCurve(*static_cast<QEasingCurve*>(other));
}

long long QEasingCurve_Type(void* ptr)
{
	return static_cast<QEasingCurve*>(ptr)->type();
}

void QEasingCurve_DestroyQEasingCurve(void* ptr)
{
	static_cast<QEasingCurve*>(ptr)->~QEasingCurve();
}

void* QEasingCurve___cubicBezierSpline_atList(void* ptr, int i)
{
	return ({ QPointF tmpValue = ({QPointF tmp = static_cast<QList<QPointF>*>(ptr)->at(i); if (i == static_cast<QList<QPointF>*>(ptr)->size()-1) { static_cast<QList<QPointF>*>(ptr)->~QList(); free(ptr); }; tmp; }); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QEasingCurve___cubicBezierSpline_setList(void* ptr, void* i)
{
	static_cast<QList<QPointF>*>(ptr)->append(*static_cast<QPointF*>(i));
}

void* QEasingCurve___cubicBezierSpline_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QPointF>();
}

void* QEasingCurve___toCubicSpline_atList(void* ptr, int i)
{
	return ({ QPointF tmpValue = ({QPointF tmp = static_cast<QVector<QPointF>*>(ptr)->at(i); if (i == static_cast<QVector<QPointF>*>(ptr)->size()-1) { static_cast<QVector<QPointF>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QEasingCurve___toCubicSpline_setList(void* ptr, void* i)
{
	static_cast<QVector<QPointF>*>(ptr)->append(*static_cast<QPointF*>(i));
}

void* QEasingCurve___toCubicSpline_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPointF>();
}

class MyQEvent: public QEvent
{
public:
	MyQEvent(QEvent::Type ty) : QEvent(ty) {};
	 ~MyQEvent() { callbackQEvent_DestroyQEvent(this); };
};

void* QEvent_NewQEvent(long long ty)
{
	return new MyQEvent(static_cast<QEvent::Type>(ty));
}

long long QEvent_Type(void* ptr)
{
	return static_cast<QEvent*>(ptr)->type();
}

void QEvent_DestroyQEvent(void* ptr)
{
	static_cast<QEvent*>(ptr)->~QEvent();
}

void QEvent_DestroyQEventDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQFile: public QFile
{
public:
	MyQFile() : QFile() {QFile_QFile_QRegisterMetaType();};
	MyQFile(const QString &name) : QFile(name) {QFile_QFile_QRegisterMetaType();};
	MyQFile(QObject *parent) : QFile(parent) {QFile_QFile_QRegisterMetaType();};
	MyQFile(const QString &name, QObject *parent) : QFile(name, parent) {QFile_QFile_QRegisterMetaType();};
	QString fileName() const { return ({ QtCore_PackedString tempVal = callbackQFileDevice_FileName(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool open(QIODevice::OpenMode mode) { return callbackQIODevice_Open(this, mode) != 0; };
	bool resize(qint64 sz) { return callbackQFileDevice_Resize(this, sz) != 0; };
	qint64 size() const { return callbackQIODevice_Size(const_cast<void*>(static_cast<const void*>(this))); };
	 ~MyQFile() { callbackQFile_DestroyQFile(this); };
	qint64 readData(char * data, qint64 maxSize) { QtCore_PackedString dataPacked = { data, maxSize };return callbackQFileDevice_ReadData(this, dataPacked, maxSize); };
	bool reset() { return callbackQIODevice_Reset(this) != 0; };
	qint64 writeData(const char * data, qint64 maxSize) { QtCore_PackedString dataPacked = { const_cast<char*>(data), maxSize };return callbackQFileDevice_WriteData(this, dataPacked, maxSize); };
	void childEvent(QChildEvent * event) { callbackQObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQObject_CustomEvent(this, event); };
	void deleteLater() { callbackQObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQObject_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQFile*)

int QFile_QFile_QRegisterMetaType(){qRegisterMetaType<QFile*>(); return qRegisterMetaType<MyQFile*>();}

void* QFile_NewQFile()
{
	return new MyQFile();
}

void* QFile_NewQFile2(struct QtCore_PackedString name)
{
	return new MyQFile(QString::fromUtf8(name.data, name.len));
}

void* QFile_NewQFile3(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQFile(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQFile(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQFile(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQFile(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQFile(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQFile(static_cast<QWindow*>(parent));
	} else {
		return new MyQFile(static_cast<QObject*>(parent));
	}
}

void* QFile_NewQFile4(struct QtCore_PackedString name, void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQFile(QString::fromUtf8(name.data, name.len), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQFile(QString::fromUtf8(name.data, name.len), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQFile(QString::fromUtf8(name.data, name.len), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQFile(QString::fromUtf8(name.data, name.len), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQFile(QString::fromUtf8(name.data, name.len), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQFile(QString::fromUtf8(name.data, name.len), static_cast<QWindow*>(parent));
	} else {
		return new MyQFile(QString::fromUtf8(name.data, name.len), static_cast<QObject*>(parent));
	}
}

char QFile_QFile_Exists(struct QtCore_PackedString fileName)
{
	return QFile::exists(QString::fromUtf8(fileName.data, fileName.len));
}

char QFile_Exists2(void* ptr)
{
	return static_cast<QFile*>(ptr)->exists();
}

char QFile_Open3(void* ptr, int fd, long long mode, long long handleFlags)
{
	return static_cast<QFile*>(ptr)->open(fd, static_cast<QIODevice::OpenModeFlag>(mode), static_cast<QFileDevice::FileHandleFlag>(handleFlags));
}

char QFile_Remove(void* ptr)
{
	return static_cast<QFile*>(ptr)->remove();
}

char QFile_QFile_Remove2(struct QtCore_PackedString fileName)
{
	return QFile::remove(QString::fromUtf8(fileName.data, fileName.len));
}

char QFile_QFile_Resize2(struct QtCore_PackedString fileName, long long sz)
{
	return QFile::resize(QString::fromUtf8(fileName.data, fileName.len), sz);
}

void QFile_DestroyQFile(void* ptr)
{
	static_cast<QFile*>(ptr)->~QFile();
}

void QFile_DestroyQFileDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQFileDevice: public QFileDevice
{
public:
	QString fileName() const { return ({ QtCore_PackedString tempVal = callbackQFileDevice_FileName(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	bool resize(qint64 sz) { return callbackQFileDevice_Resize(this, sz) != 0; };
	qint64 size() const { return callbackQIODevice_Size(const_cast<void*>(static_cast<const void*>(this))); };
	 ~MyQFileDevice() { callbackQFileDevice_DestroyQFileDevice(this); };
	bool open(QIODevice::OpenMode mode) { return callbackQIODevice_Open(this, mode) != 0; };
	qint64 readData(char * data, qint64 maxSize) { QtCore_PackedString dataPacked = { data, maxSize };return callbackQFileDevice_ReadData(this, dataPacked, maxSize); };
	bool reset() { return callbackQIODevice_Reset(this) != 0; };
	qint64 writeData(const char * data, qint64 maxSize) { QtCore_PackedString dataPacked = { const_cast<char*>(data), maxSize };return callbackQFileDevice_WriteData(this, dataPacked, maxSize); };
	void childEvent(QChildEvent * event) { callbackQObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQObject_CustomEvent(this, event); };
	void deleteLater() { callbackQObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQObject_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQFileDevice*)

int QFileDevice_QFileDevice_QRegisterMetaType(){qRegisterMetaType<QFileDevice*>(); return qRegisterMetaType<MyQFileDevice*>();}

long long QFileDevice_Error(void* ptr)
{
	return static_cast<QFileDevice*>(ptr)->error();
}

struct QtCore_PackedString QFileDevice_FileName(void* ptr)
{
	return ({ QByteArray t4e2118 = static_cast<QFileDevice*>(ptr)->fileName().toUtf8(); QtCore_PackedString { const_cast<char*>(t4e2118.prepend("WHITESPACE").constData()+10), t4e2118.size()-10 }; });
}

struct QtCore_PackedString QFileDevice_FileNameDefault(void* ptr)
{
	if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		return ({ QByteArray tb198ab = static_cast<QFile*>(ptr)->QFile::fileName().toUtf8(); QtCore_PackedString { const_cast<char*>(tb198ab.prepend("WHITESPACE").constData()+10), tb198ab.size()-10 }; });
	} else {
		return ({ QByteArray tb198ab = static_cast<QFileDevice*>(ptr)->QFileDevice::fileName().toUtf8(); QtCore_PackedString { const_cast<char*>(tb198ab.prepend("WHITESPACE").constData()+10), tb198ab.size()-10 }; });
	}
}

char QFileDevice_Resize(void* ptr, long long sz)
{
	return static_cast<QFileDevice*>(ptr)->resize(sz);
}

char QFileDevice_ResizeDefault(void* ptr, long long sz)
{
	if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFile*>(ptr)->QFile::resize(sz);
	} else {
		return static_cast<QFileDevice*>(ptr)->QFileDevice::resize(sz);
	}
}

void QFileDevice_DestroyQFileDevice(void* ptr)
{
	static_cast<QFileDevice*>(ptr)->~QFileDevice();
}

void QFileDevice_DestroyQFileDeviceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

long long QFileDevice_ReadData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QFileDevice*>(ptr)->readData(data, maxSize);
}

long long QFileDevice_ReadDataDefault(void* ptr, char* data, long long maxSize)
{
	if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFile*>(ptr)->QFile::readData(data, maxSize);
	} else {
		return static_cast<QFileDevice*>(ptr)->QFileDevice::readData(data, maxSize);
	}
}

long long QFileDevice_WriteData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QFileDevice*>(ptr)->writeData(const_cast<const char*>(data), maxSize);
}

long long QFileDevice_WriteDataDefault(void* ptr, char* data, long long maxSize)
{
	if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFile*>(ptr)->QFile::writeData(const_cast<const char*>(data), maxSize);
	} else {
		return static_cast<QFileDevice*>(ptr)->QFileDevice::writeData(const_cast<const char*>(data), maxSize);
	}
}

void* QFileInfo_NewQFileInfo2()
{
	return new QFileInfo();
}

void* QFileInfo_NewQFileInfo3(struct QtCore_PackedString file)
{
	return new QFileInfo(QString::fromUtf8(file.data, file.len));
}

void* QFileInfo_NewQFileInfo4(void* file)
{
	return new QFileInfo(*static_cast<QFile*>(file));
}

void* QFileInfo_NewQFileInfo5(void* dir, struct QtCore_PackedString file)
{
	return new QFileInfo(*static_cast<QDir*>(dir), QString::fromUtf8(file.data, file.len));
}

void* QFileInfo_NewQFileInfo6(void* fileinfo)
{
	return new QFileInfo(*static_cast<QFileInfo*>(fileinfo));
}

void* QFileInfo_Dir(void* ptr)
{
	return new QDir(static_cast<QFileInfo*>(ptr)->dir());
}

char QFileInfo_Exists(void* ptr)
{
	return static_cast<QFileInfo*>(ptr)->exists();
}

char QFileInfo_QFileInfo_Exists2(struct QtCore_PackedString file)
{
	return QFileInfo::exists(QString::fromUtf8(file.data, file.len));
}

struct QtCore_PackedString QFileInfo_FileName(void* ptr)
{
	return ({ QByteArray t8cf8a1 = static_cast<QFileInfo*>(ptr)->fileName().toUtf8(); QtCore_PackedString { const_cast<char*>(t8cf8a1.prepend("WHITESPACE").constData()+10), t8cf8a1.size()-10 }; });
}

char QFileInfo_IsDir(void* ptr)
{
	return static_cast<QFileInfo*>(ptr)->isDir();
}

struct QtCore_PackedString QFileInfo_Path(void* ptr)
{
	return ({ QByteArray tdcd027 = static_cast<QFileInfo*>(ptr)->path().toUtf8(); QtCore_PackedString { const_cast<char*>(tdcd027.prepend("WHITESPACE").constData()+10), tdcd027.size()-10 }; });
}

void QFileInfo_SetFile(void* ptr, struct QtCore_PackedString file)
{
	static_cast<QFileInfo*>(ptr)->setFile(QString::fromUtf8(file.data, file.len));
}

void QFileInfo_SetFile2(void* ptr, void* file)
{
	static_cast<QFileInfo*>(ptr)->setFile(*static_cast<QFile*>(file));
}

void QFileInfo_SetFile3(void* ptr, void* dir, struct QtCore_PackedString file)
{
	static_cast<QFileInfo*>(ptr)->setFile(*static_cast<QDir*>(dir), QString::fromUtf8(file.data, file.len));
}

long long QFileInfo_Size(void* ptr)
{
	return static_cast<QFileInfo*>(ptr)->size();
}

struct QtCore_PackedString QFileInfo_Suffix(void* ptr)
{
	return ({ QByteArray t1b5684 = static_cast<QFileInfo*>(ptr)->suffix().toUtf8(); QtCore_PackedString { const_cast<char*>(t1b5684.prepend("WHITESPACE").constData()+10), t1b5684.size()-10 }; });
}

void QFileInfo_DestroyQFileInfo(void* ptr)
{
	static_cast<QFileInfo*>(ptr)->~QFileInfo();
}

class MyQIODevice: public QIODevice
{
public:
	MyQIODevice() : QIODevice() {QIODevice_QIODevice_QRegisterMetaType();};
	MyQIODevice(QObject *parent) : QIODevice(parent) {QIODevice_QIODevice_QRegisterMetaType();};
	bool open(QIODevice::OpenMode mode) { return callbackQIODevice_Open(this, mode) != 0; };
	qint64 readData(char * data, qint64 maxSize) { QtCore_PackedString dataPacked = { data, maxSize };return callbackQIODevice_ReadData(this, dataPacked, maxSize); };
	bool reset() { return callbackQIODevice_Reset(this) != 0; };
	qint64 size() const { return callbackQIODevice_Size(const_cast<void*>(static_cast<const void*>(this))); };
	qint64 writeData(const char * data, qint64 maxSize) { QtCore_PackedString dataPacked = { const_cast<char*>(data), maxSize };return callbackQIODevice_WriteData(this, dataPacked, maxSize); };
	 ~MyQIODevice() { callbackQIODevice_DestroyQIODevice(this); };
	void childEvent(QChildEvent * event) { callbackQObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQObject_CustomEvent(this, event); };
	void deleteLater() { callbackQObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQObject_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQIODevice*)

int QIODevice_QIODevice_QRegisterMetaType(){qRegisterMetaType<QIODevice*>(); return qRegisterMetaType<MyQIODevice*>();}

void* QIODevice_NewQIODevice()
{
	return new MyQIODevice();
}

void* QIODevice_NewQIODevice2(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQIODevice(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQIODevice(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQIODevice(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQIODevice(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQIODevice(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQIODevice(static_cast<QWindow*>(parent));
	} else {
		return new MyQIODevice(static_cast<QObject*>(parent));
	}
}

char QIODevice_Open(void* ptr, long long mode)
{
	return static_cast<QIODevice*>(ptr)->open(static_cast<QIODevice::OpenModeFlag>(mode));
}

char QIODevice_OpenDefault(void* ptr, long long mode)
{
	if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFile*>(ptr)->QFile::open(static_cast<QIODevice::OpenModeFlag>(mode));
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFileDevice*>(ptr)->QFileDevice::open(static_cast<QIODevice::OpenModeFlag>(mode));
	} else {
		return static_cast<QIODevice*>(ptr)->QIODevice::open(static_cast<QIODevice::OpenModeFlag>(mode));
	}
}

long long QIODevice_ReadData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QIODevice*>(ptr)->readData(data, maxSize);
}

char QIODevice_Reset(void* ptr)
{
	return static_cast<QIODevice*>(ptr)->reset();
}

char QIODevice_ResetDefault(void* ptr)
{
	if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFile*>(ptr)->QFile::reset();
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFileDevice*>(ptr)->QFileDevice::reset();
	} else {
		return static_cast<QIODevice*>(ptr)->QIODevice::reset();
	}
}

long long QIODevice_Size(void* ptr)
{
	return static_cast<QIODevice*>(ptr)->size();
}

long long QIODevice_SizeDefault(void* ptr)
{
	if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFile*>(ptr)->QFile::size();
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFileDevice*>(ptr)->QFileDevice::size();
	} else {
		return static_cast<QIODevice*>(ptr)->QIODevice::size();
	}
}

long long QIODevice_WriteData(void* ptr, char* data, long long maxSize)
{
	return static_cast<QIODevice*>(ptr)->writeData(const_cast<const char*>(data), maxSize);
}

void QIODevice_DestroyQIODevice(void* ptr)
{
	static_cast<QIODevice*>(ptr)->~QIODevice();
}

void QIODevice_DestroyQIODeviceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QJsonArray_NewQJsonArray()
{
	return new QJsonArray();
}

void* QJsonArray_NewQJsonArray3(void* other)
{
	return new QJsonArray(*static_cast<QJsonArray*>(other));
}

void* QJsonArray_NewQJsonArray4(void* other)
{
	return new QJsonArray(*static_cast<QJsonArray*>(other));
}

void* QJsonArray_At(void* ptr, int i)
{
	return new QJsonValue(static_cast<QJsonArray*>(ptr)->at(i));
}

char QJsonArray_Contains(void* ptr, void* value)
{
	return static_cast<QJsonArray*>(ptr)->contains(*static_cast<QJsonValue*>(value));
}

int QJsonArray_Count(void* ptr)
{
	return static_cast<QJsonArray*>(ptr)->count();
}

void QJsonArray_Insert(void* ptr, int i, void* value)
{
	static_cast<QJsonArray*>(ptr)->insert(i, *static_cast<QJsonValue*>(value));
}

int QJsonArray_Size(void* ptr)
{
	return static_cast<QJsonArray*>(ptr)->size();
}

void QJsonArray_DestroyQJsonArray(void* ptr)
{
	static_cast<QJsonArray*>(ptr)->~QJsonArray();
}

void* QJsonArray___fromVariantList_list_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QJsonArray___fromVariantList_list_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QJsonArray___fromVariantList_list_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QJsonArray___toVariantList_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QJsonArray___toVariantList_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QJsonArray___toVariantList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QJsonDocument_NewQJsonDocument()
{
	return new QJsonDocument();
}

void* QJsonDocument_NewQJsonDocument2(void* object)
{
	return new QJsonDocument(*static_cast<QJsonObject*>(object));
}

void* QJsonDocument_NewQJsonDocument3(void* array)
{
	return new QJsonDocument(*static_cast<QJsonArray*>(array));
}

void* QJsonDocument_NewQJsonDocument4(void* other)
{
	return new QJsonDocument(*static_cast<QJsonDocument*>(other));
}

void* QJsonDocument_NewQJsonDocument5(void* other)
{
	return new QJsonDocument(*static_cast<QJsonDocument*>(other));
}

void* QJsonDocument_Array(void* ptr)
{
	return new QJsonArray(static_cast<QJsonDocument*>(ptr)->array());
}

void* QJsonDocument_Object(void* ptr)
{
	return new QJsonObject(static_cast<QJsonDocument*>(ptr)->object());
}

void* QJsonDocument_ToVariant(void* ptr)
{
	return new QVariant(static_cast<QJsonDocument*>(ptr)->toVariant());
}

void QJsonDocument_DestroyQJsonDocument(void* ptr)
{
	static_cast<QJsonDocument*>(ptr)->~QJsonDocument();
}

void* QJsonObject_NewQJsonObject()
{
	return new QJsonObject();
}

void* QJsonObject_NewQJsonObject3(void* other)
{
	return new QJsonObject(*static_cast<QJsonObject*>(other));
}

void* QJsonObject_NewQJsonObject4(void* other)
{
	return new QJsonObject(*static_cast<QJsonObject*>(other));
}

char QJsonObject_Contains(void* ptr, struct QtCore_PackedString key)
{
	return static_cast<QJsonObject*>(ptr)->contains(QString::fromUtf8(key.data, key.len));
}

char QJsonObject_Contains2(void* ptr, void* key)
{
	return static_cast<QJsonObject*>(ptr)->contains(*static_cast<QLatin1String*>(key));
}

int QJsonObject_Count(void* ptr)
{
	return static_cast<QJsonObject*>(ptr)->count();
}

void QJsonObject_Remove(void* ptr, struct QtCore_PackedString key)
{
	static_cast<QJsonObject*>(ptr)->remove(QString::fromUtf8(key.data, key.len));
}

int QJsonObject_Size(void* ptr)
{
	return static_cast<QJsonObject*>(ptr)->size();
}

void* QJsonObject_Value(void* ptr, struct QtCore_PackedString key)
{
	return new QJsonValue(static_cast<QJsonObject*>(ptr)->value(QString::fromUtf8(key.data, key.len)));
}

void* QJsonObject_Value2(void* ptr, void* key)
{
	return new QJsonValue(static_cast<QJsonObject*>(ptr)->value(*static_cast<QLatin1String*>(key)));
}

void QJsonObject_DestroyQJsonObject(void* ptr)
{
	static_cast<QJsonObject*>(ptr)->~QJsonObject();
}

void* QJsonObject___fromVariantHash_hash_atList(void* ptr, struct QtCore_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QHash<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QHash<QString, QVariant>*>(ptr)->size()-1) { static_cast<QHash<QString, QVariant>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QJsonObject___fromVariantHash_hash_setList(void* ptr, struct QtCore_PackedString key, void* i)
{
	static_cast<QHash<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QJsonObject___fromVariantHash_hash_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<QString, QVariant>();
}

struct QtCore_PackedList QJsonObject___fromVariantHash_hash_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValuef43bc5 = new QList<QString>(static_cast<QHash<QString, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValuef43bc5, tmpValuef43bc5->size() }; });
}

void* QJsonObject___toVariantHash_atList(void* ptr, struct QtCore_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QHash<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QHash<QString, QVariant>*>(ptr)->size()-1) { static_cast<QHash<QString, QVariant>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QJsonObject___toVariantHash_setList(void* ptr, struct QtCore_PackedString key, void* i)
{
	static_cast<QHash<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QJsonObject___toVariantHash_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<QString, QVariant>();
}

struct QtCore_PackedList QJsonObject___toVariantHash_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValuef43bc5 = new QList<QString>(static_cast<QHash<QString, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValuef43bc5, tmpValuef43bc5->size() }; });
}

void* QJsonObject___toVariantMap_atList(void* ptr, struct QtCore_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QJsonObject___toVariantMap_setList(void* ptr, struct QtCore_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QJsonObject___toVariantMap_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtCore_PackedList QJsonObject___toVariantMap_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

struct QtCore_PackedString QJsonObject_____fromVariantHash_hash_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t94aa5e = ({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8(); QtCore_PackedString { const_cast<char*>(t94aa5e.prepend("WHITESPACE").constData()+10), t94aa5e.size()-10 }; });
}

void QJsonObject_____fromVariantHash_hash_keyList_setList(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QJsonObject_____fromVariantHash_hash_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtCore_PackedString QJsonObject_____fromVariantMap_map_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t94aa5e = ({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8(); QtCore_PackedString { const_cast<char*>(t94aa5e.prepend("WHITESPACE").constData()+10), t94aa5e.size()-10 }; });
}

void QJsonObject_____fromVariantMap_map_keyList_setList(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QJsonObject_____fromVariantMap_map_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtCore_PackedString QJsonObject_____toVariantHash_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t94aa5e = ({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8(); QtCore_PackedString { const_cast<char*>(t94aa5e.prepend("WHITESPACE").constData()+10), t94aa5e.size()-10 }; });
}

void QJsonObject_____toVariantHash_keyList_setList(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QJsonObject_____toVariantHash_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtCore_PackedString QJsonObject_____toVariantMap_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t94aa5e = ({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8(); QtCore_PackedString { const_cast<char*>(t94aa5e.prepend("WHITESPACE").constData()+10), t94aa5e.size()-10 }; });
}

void QJsonObject_____toVariantMap_keyList_setList(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QJsonObject_____toVariantMap_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

void* QJsonValue_NewQJsonValue(long long ty)
{
	return new QJsonValue(static_cast<QJsonValue::Type>(ty));
}

void* QJsonValue_NewQJsonValue2(char b)
{
	return new QJsonValue(b != 0);
}

void* QJsonValue_NewQJsonValue3(double n)
{
	return new QJsonValue(n);
}

void* QJsonValue_NewQJsonValue4(int n)
{
	return new QJsonValue(n);
}

void* QJsonValue_NewQJsonValue5(long long n)
{
	return new QJsonValue(n);
}

void* QJsonValue_NewQJsonValue6(struct QtCore_PackedString s)
{
	return new QJsonValue(QString::fromUtf8(s.data, s.len));
}

void* QJsonValue_NewQJsonValue7(void* s)
{
	return new QJsonValue(*static_cast<QLatin1String*>(s));
}

void* QJsonValue_NewQJsonValue8(char* s)
{
	return new QJsonValue(const_cast<const char*>(s));
}

void* QJsonValue_NewQJsonValue9(void* a)
{
	return new QJsonValue(*static_cast<QJsonArray*>(a));
}

void* QJsonValue_NewQJsonValue10(void* o)
{
	return new QJsonValue(*static_cast<QJsonObject*>(o));
}

void* QJsonValue_NewQJsonValue11(void* other)
{
	return new QJsonValue(*static_cast<QJsonValue*>(other));
}

void* QJsonValue_NewQJsonValue12(void* other)
{
	return new QJsonValue(*static_cast<QJsonValue*>(other));
}

int QJsonValue_ToInt(void* ptr, int defaultValue)
{
	return static_cast<QJsonValue*>(ptr)->toInt(defaultValue);
}

void* QJsonValue_ToVariant(void* ptr)
{
	return new QVariant(static_cast<QJsonValue*>(ptr)->toVariant());
}

long long QJsonValue_Type(void* ptr)
{
	return static_cast<QJsonValue*>(ptr)->type();
}

void QJsonValue_DestroyQJsonValue(void* ptr)
{
	static_cast<QJsonValue*>(ptr)->~QJsonValue();
}

void* QLatin1Char_NewQLatin1Char(char* c)
{
	return new QLatin1Char(*c);
}

void* QLatin1String_NewQLatin1String()
{
	return new QLatin1String();
}

void* QLatin1String_NewQLatin1String2(char* str)
{
	return new QLatin1String(const_cast<const char*>(str));
}

void* QLatin1String_NewQLatin1String3(char* first, char* last)
{
	return new QLatin1String(const_cast<const char*>(first), const_cast<const char*>(last));
}

void* QLatin1String_NewQLatin1String4(char* str, int size)
{
	return new QLatin1String(const_cast<const char*>(str), size);
}

void* QLatin1String_NewQLatin1String5(void* str)
{
	return new QLatin1String(*static_cast<QByteArray*>(str));
}

struct QtCore_PackedString QLatin1String_Data(void* ptr)
{
	return QtCore_PackedString { const_cast<char*>(static_cast<QLatin1String*>(ptr)->data()), static_cast<QLatin1String*>(ptr)->size() };
}

void* QLatin1String_Left(void* ptr, int length)
{
	return new QLatin1String(static_cast<QLatin1String*>(ptr)->left(length));
}

void* QLatin1String_Right(void* ptr, int length)
{
	return new QLatin1String(static_cast<QLatin1String*>(ptr)->right(length));
}

int QLatin1String_Size(void* ptr)
{
	return static_cast<QLatin1String*>(ptr)->size();
}

void* QLine_NewQLine()
{
	return new QLine();
}

void* QLine_NewQLine2(void* p1, void* p2)
{
	return new QLine(*static_cast<QPoint*>(p1), *static_cast<QPoint*>(p2));
}

void* QLine_NewQLine3(int x1, int y1, int x2, int y2)
{
	return new QLine(x1, y1, x2, y2);
}

void* QLineF_NewQLineF()
{
	return new QLineF();
}

void* QLineF_NewQLineF2(void* p1, void* p2)
{
	return new QLineF(*static_cast<QPointF*>(p1), *static_cast<QPointF*>(p2));
}

void* QLineF_NewQLineF3(double x1, double y1, double x2, double y2)
{
	return new QLineF(x1, y1, x2, y2);
}

void* QLineF_NewQLineF4(void* line)
{
	return new QLineF(*static_cast<QLine*>(line));
}

void* QLocale_NewQLocale()
{
	return new QLocale();
}

void* QLocale_NewQLocale2(struct QtCore_PackedString name)
{
	return new QLocale(QString::fromUtf8(name.data, name.len));
}

void* QLocale_NewQLocale3(long long language, long long country)
{
	return new QLocale(static_cast<QLocale::Language>(language), static_cast<QLocale::Country>(country));
}

void* QLocale_NewQLocale4(long long language, long long scri, long long country)
{
	return new QLocale(static_cast<QLocale::Language>(language), static_cast<QLocale::Script>(scri), static_cast<QLocale::Country>(country));
}

void* QLocale_NewQLocale5(void* other)
{
	return new QLocale(*static_cast<QLocale*>(other));
}

void* QLocale_QLocale_C()
{
	return new QLocale(QLocale::c());
}

struct QtCore_PackedString QLocale_Name(void* ptr)
{
	return ({ QByteArray tee867e = static_cast<QLocale*>(ptr)->name().toUtf8(); QtCore_PackedString { const_cast<char*>(tee867e.prepend("WHITESPACE").constData()+10), tee867e.size()-10 }; });
}

int QLocale_ToInt(void* ptr, struct QtCore_PackedString s, char* ok)
{
	return static_cast<QLocale*>(ptr)->toInt(QString::fromUtf8(s.data, s.len), reinterpret_cast<bool*>(ok));
}

int QLocale_ToInt2(void* ptr, void* s, char* ok)
{
	return static_cast<QLocale*>(ptr)->toInt(*static_cast<QStringRef*>(s), reinterpret_cast<bool*>(ok));
}

int QLocale_ToInt3(void* ptr, void* s, char* ok)
{
	return static_cast<QLocale*>(ptr)->toInt(*static_cast<QStringView*>(s), reinterpret_cast<bool*>(ok));
}

void QLocale_DestroyQLocale(void* ptr)
{
	static_cast<QLocale*>(ptr)->~QLocale();
}

void* QLocale___matchingLocales_atList(void* ptr, int i)
{
	return new QLocale(({QLocale tmp = static_cast<QList<QLocale>*>(ptr)->at(i); if (i == static_cast<QList<QLocale>*>(ptr)->size()-1) { static_cast<QList<QLocale>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QLocale___matchingLocales_setList(void* ptr, void* i)
{
	static_cast<QList<QLocale>*>(ptr)->append(*static_cast<QLocale*>(i));
}

void* QLocale___matchingLocales_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QLocale>();
}

long long QLocale___weekdays_atList(void* ptr, int i)
{
	return ({Qt::DayOfWeek tmp = static_cast<QList<Qt::DayOfWeek>*>(ptr)->at(i); if (i == static_cast<QList<Qt::DayOfWeek>*>(ptr)->size()-1) { static_cast<QList<Qt::DayOfWeek>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QLocale___weekdays_setList(void* ptr, long long i)
{
	static_cast<QList<Qt::DayOfWeek>*>(ptr)->append(static_cast<Qt::DayOfWeek>(i));
}

void* QLocale___weekdays_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<Qt::DayOfWeek>();
}

void* QMetaMethod_Name(void* ptr)
{
	return new QByteArray(static_cast<QMetaMethod*>(ptr)->name());
}

void* QMetaMethod___parameterNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QMetaMethod___parameterNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QMetaMethod___parameterNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QMetaMethod___parameterTypes_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QMetaMethod___parameterTypes_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QMetaMethod___parameterTypes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QMetaObject_Constructor(void* ptr, int index)
{
	return new QMetaMethod(static_cast<QMetaObject*>(ptr)->constructor(index));
}

void* QMetaObject_Method(void* ptr, int index)
{
	return new QMetaMethod(static_cast<QMetaObject*>(ptr)->method(index));
}

struct QtCore_PackedString QMetaProperty_Name(void* ptr)
{
	return QtCore_PackedString { const_cast<char*>(static_cast<QMetaProperty*>(ptr)->name()), -1 };
}

char QMetaProperty_Reset(void* ptr, void* object)
{
	return static_cast<QMetaProperty*>(ptr)->reset(static_cast<QObject*>(object));
}

long long QMetaProperty_Type(void* ptr)
{
	return static_cast<QMetaProperty*>(ptr)->type();
}

class MyQMimeData: public QMimeData
{
public:
	MyQMimeData() : QMimeData() {QMimeData_QMimeData_QRegisterMetaType();};
	 ~MyQMimeData() { callbackQMimeData_DestroyQMimeData(this); };
	void childEvent(QChildEvent * event) { callbackQObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQObject_CustomEvent(this, event); };
	void deleteLater() { callbackQObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQObject_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQMimeData*)

int QMimeData_QMimeData_QRegisterMetaType(){qRegisterMetaType<QMimeData*>(); return qRegisterMetaType<MyQMimeData*>();}

void* QMimeData_NewQMimeData()
{
	return new MyQMimeData();
}

void* QMimeData_Data(void* ptr, struct QtCore_PackedString mimeType)
{
	return new QByteArray(static_cast<QMimeData*>(ptr)->data(QString::fromUtf8(mimeType.data, mimeType.len)));
}

void QMimeData_SetData(void* ptr, struct QtCore_PackedString mimeType, void* data)
{
	static_cast<QMimeData*>(ptr)->setData(QString::fromUtf8(mimeType.data, mimeType.len), *static_cast<QByteArray*>(data));
}

void QMimeData_SetText(void* ptr, struct QtCore_PackedString text)
{
	static_cast<QMimeData*>(ptr)->setText(QString::fromUtf8(text.data, text.len));
}

struct QtCore_PackedString QMimeData_Text(void* ptr)
{
	return ({ QByteArray t2355ec = static_cast<QMimeData*>(ptr)->text().toUtf8(); QtCore_PackedString { const_cast<char*>(t2355ec.prepend("WHITESPACE").constData()+10), t2355ec.size()-10 }; });
}

void QMimeData_DestroyQMimeData(void* ptr)
{
	static_cast<QMimeData*>(ptr)->~QMimeData();
}

void QMimeData_DestroyQMimeDataDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QMimeData___setUrls_urls_atList(void* ptr, int i)
{
	return new QUrl(({QUrl tmp = static_cast<QList<QUrl>*>(ptr)->at(i); if (i == static_cast<QList<QUrl>*>(ptr)->size()-1) { static_cast<QList<QUrl>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QMimeData___setUrls_urls_setList(void* ptr, void* i)
{
	static_cast<QList<QUrl>*>(ptr)->append(*static_cast<QUrl*>(i));
}

void* QMimeData___setUrls_urls_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QUrl>();
}

void* QMimeData___urls_atList(void* ptr, int i)
{
	return new QUrl(({QUrl tmp = static_cast<QList<QUrl>*>(ptr)->at(i); if (i == static_cast<QList<QUrl>*>(ptr)->size()-1) { static_cast<QList<QUrl>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QMimeData___urls_setList(void* ptr, void* i)
{
	static_cast<QList<QUrl>*>(ptr)->append(*static_cast<QUrl*>(i));
}

void* QMimeData___urls_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QUrl>();
}

void* QModelIndex_NewQModelIndex()
{
	return new QModelIndex();
}

int QModelIndex_Column(void* ptr)
{
	return static_cast<QModelIndex*>(ptr)->column();
}

void* QModelIndex_Data(void* ptr, int role)
{
	return new QVariant(static_cast<QModelIndex*>(ptr)->data(role));
}

long long QModelIndex_Flags(void* ptr)
{
	return static_cast<QModelIndex*>(ptr)->flags();
}

void* QModelIndex_Model(void* ptr)
{
	return const_cast<QAbstractItemModel*>(static_cast<QModelIndex*>(ptr)->model());
}

void* QModelIndex_Parent(void* ptr)
{
	return new QModelIndex(static_cast<QModelIndex*>(ptr)->parent());
}

int QModelIndex_Row(void* ptr)
{
	return static_cast<QModelIndex*>(ptr)->row();
}

void* QModelIndex_Sibling(void* ptr, int row, int column)
{
	return new QModelIndex(static_cast<QModelIndex*>(ptr)->sibling(row, column));
}

class MyQObject: public QObject
{
public:
	MyQObject(QObject *parent = Q_NULLPTR) : QObject(parent) {QObject_QObject_QRegisterMetaType();};
	void childEvent(QChildEvent * event) { callbackQObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQObject_CustomEvent(this, event); };
	void deleteLater() { callbackQObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQObject_TimerEvent(this, event); };
	 ~MyQObject() { callbackQObject_DestroyQObject(this); };
};

Q_DECLARE_METATYPE(MyQObject*)

int QObject_QObject_QRegisterMetaType(){qRegisterMetaType<QObject*>(); return qRegisterMetaType<MyQObject*>();}

void* QObject_NewQObject(void* parent)
{
	return new MyQObject(static_cast<QObject*>(parent));
}

void QObject_ChildEvent(void* ptr, void* event)
{
	static_cast<QObject*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QObject_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QTimer*>(static_cast<QObject*>(ptr))) {
		static_cast<QTimer*>(ptr)->QTimer::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QMimeData*>(static_cast<QObject*>(ptr))) {
		static_cast<QMimeData*>(ptr)->QMimeData::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QCoreApplication*>(static_cast<QObject*>(ptr))) {
		static_cast<QCoreApplication*>(ptr)->QCoreApplication::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		static_cast<QFile*>(ptr)->QFile::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDevice*>(ptr)->QFileDevice::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QIODevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QIODevice*>(ptr)->QIODevice::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QAbstractItemModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QObject*>(ptr)->QObject::childEvent(static_cast<QChildEvent*>(event));
	}
}

struct QtCore_PackedList QObject_Children(void* ptr)
{
	return ({ QList<QObject *>* tmpValue53f390 = new QList<QObject *>(static_cast<QObject*>(ptr)->children()); QtCore_PackedList { tmpValue53f390, tmpValue53f390->size() }; });
}

void QObject_ConnectNotify(void* ptr, void* sign)
{
	static_cast<QObject*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QObject_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QTimer*>(static_cast<QObject*>(ptr))) {
		static_cast<QTimer*>(ptr)->QTimer::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QMimeData*>(static_cast<QObject*>(ptr))) {
		static_cast<QMimeData*>(ptr)->QMimeData::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QCoreApplication*>(static_cast<QObject*>(ptr))) {
		static_cast<QCoreApplication*>(ptr)->QCoreApplication::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		static_cast<QFile*>(ptr)->QFile::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDevice*>(ptr)->QFileDevice::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QIODevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QIODevice*>(ptr)->QIODevice::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAbstractItemModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QObject*>(ptr)->QObject::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QObject_CustomEvent(void* ptr, void* event)
{
	static_cast<QObject*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QObject_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QTimer*>(static_cast<QObject*>(ptr))) {
		static_cast<QTimer*>(ptr)->QTimer::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QMimeData*>(static_cast<QObject*>(ptr))) {
		static_cast<QMimeData*>(ptr)->QMimeData::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QCoreApplication*>(static_cast<QObject*>(ptr))) {
		static_cast<QCoreApplication*>(ptr)->QCoreApplication::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		static_cast<QFile*>(ptr)->QFile::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDevice*>(ptr)->QFileDevice::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QIODevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QIODevice*>(ptr)->QIODevice::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAbstractItemModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QObject*>(ptr)->QObject::customEvent(static_cast<QEvent*>(event));
	}
}

void QObject_DeleteLater(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QObject*>(ptr), "deleteLater");
}

void QObject_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QTimer*>(static_cast<QObject*>(ptr))) {
		static_cast<QTimer*>(ptr)->QTimer::deleteLater();
	} else if (dynamic_cast<QMimeData*>(static_cast<QObject*>(ptr))) {
		static_cast<QMimeData*>(ptr)->QMimeData::deleteLater();
	} else if (dynamic_cast<QCoreApplication*>(static_cast<QObject*>(ptr))) {
		static_cast<QCoreApplication*>(ptr)->QCoreApplication::deleteLater();
	} else if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		static_cast<QFile*>(ptr)->QFile::deleteLater();
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDevice*>(ptr)->QFileDevice::deleteLater();
	} else if (dynamic_cast<QIODevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QIODevice*>(ptr)->QIODevice::deleteLater();
	} else if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::deleteLater();
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::deleteLater();
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::deleteLater();
	} else if (dynamic_cast<QAbstractItemModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::deleteLater();
	} else {
		static_cast<QObject*>(ptr)->QObject::deleteLater();
	}
}

void QObject_ConnectDestroyed(void* ptr, long long t)
{
	QObject::connect(static_cast<QObject*>(ptr), static_cast<void (QObject::*)(QObject *)>(&QObject::destroyed), static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(QObject *)>(&MyQObject::Signal_Destroyed), static_cast<Qt::ConnectionType>(t));
}

void QObject_DisconnectDestroyed(void* ptr)
{
	QObject::disconnect(static_cast<QObject*>(ptr), static_cast<void (QObject::*)(QObject *)>(&QObject::destroyed), static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(QObject *)>(&MyQObject::Signal_Destroyed));
}

void QObject_Destroyed(void* ptr, void* obj)
{
	static_cast<QObject*>(ptr)->destroyed(static_cast<QObject*>(obj));
}

char QObject_QObject_Disconnect(void* sender, char* sign, void* receiver, char* method)
{
	return QObject::disconnect(static_cast<QObject*>(sender), const_cast<const char*>(sign), static_cast<QObject*>(receiver), const_cast<const char*>(method));
}

char QObject_QObject_Disconnect2(void* sender, void* sign, void* receiver, void* method)
{
	return QObject::disconnect(static_cast<QObject*>(sender), *static_cast<QMetaMethod*>(sign), static_cast<QObject*>(receiver), *static_cast<QMetaMethod*>(method));
}

char QObject_Disconnect3(void* ptr, char* sign, void* receiver, char* method)
{
	return static_cast<QObject*>(ptr)->disconnect(const_cast<const char*>(sign), static_cast<QObject*>(receiver), const_cast<const char*>(method));
}

char QObject_Disconnect4(void* ptr, void* receiver, char* method)
{
	return static_cast<QObject*>(ptr)->disconnect(static_cast<QObject*>(receiver), const_cast<const char*>(method));
}

void QObject_DisconnectNotify(void* ptr, void* sign)
{
	static_cast<QObject*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QObject_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QTimer*>(static_cast<QObject*>(ptr))) {
		static_cast<QTimer*>(ptr)->QTimer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QMimeData*>(static_cast<QObject*>(ptr))) {
		static_cast<QMimeData*>(ptr)->QMimeData::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QCoreApplication*>(static_cast<QObject*>(ptr))) {
		static_cast<QCoreApplication*>(ptr)->QCoreApplication::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		static_cast<QFile*>(ptr)->QFile::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDevice*>(ptr)->QFileDevice::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QIODevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QIODevice*>(ptr)->QIODevice::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QAbstractItemModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QObject*>(ptr)->QObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

char QObject_Event(void* ptr, void* e)
{
	return static_cast<QObject*>(ptr)->event(static_cast<QEvent*>(e));
}

char QObject_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QTimer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTimer*>(ptr)->QTimer::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QMimeData*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMimeData*>(ptr)->QMimeData::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QCoreApplication*>(static_cast<QObject*>(ptr))) {
		return static_cast<QCoreApplication*>(ptr)->QCoreApplication::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFile*>(ptr)->QFile::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFileDevice*>(ptr)->QFileDevice::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QIODevice*>(static_cast<QObject*>(ptr))) {
		return static_cast<QIODevice*>(ptr)->QIODevice::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QAbstractItemModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QObject*>(ptr)->QObject::event(static_cast<QEvent*>(e));
	}
}

char QObject_EventFilter(void* ptr, void* watched, void* event)
{
	return static_cast<QObject*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
}

char QObject_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QTimer*>(static_cast<QObject*>(ptr))) {
		return static_cast<QTimer*>(ptr)->QTimer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QMimeData*>(static_cast<QObject*>(ptr))) {
		return static_cast<QMimeData*>(ptr)->QMimeData::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QCoreApplication*>(static_cast<QObject*>(ptr))) {
		return static_cast<QCoreApplication*>(ptr)->QCoreApplication::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFile*>(ptr)->QFile::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		return static_cast<QFileDevice*>(ptr)->QFileDevice::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QIODevice*>(static_cast<QObject*>(ptr))) {
		return static_cast<QIODevice*>(ptr)->QIODevice::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else if (dynamic_cast<QAbstractItemModel*>(static_cast<QObject*>(ptr))) {
		return static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QObject*>(ptr)->QObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void* QObject_FindChild(void* ptr, struct QtCore_PackedString name, long long options)
{
	return static_cast<QObject*>(ptr)->findChild<QObject*>(QString::fromUtf8(name.data, name.len), static_cast<Qt::FindChildOption>(options));
}

struct QtCore_PackedList QObject_FindChildren(void* ptr, struct QtCore_PackedString name, long long options)
{
	return ({ QList<QObject*>* tmpValue03dce3 = new QList<QObject*>(static_cast<QObject*>(ptr)->findChildren<QObject*>(QString::fromUtf8(name.data, name.len), static_cast<Qt::FindChildOption>(options))); QtCore_PackedList { tmpValue03dce3, tmpValue03dce3->size() }; });
}

struct QtCore_PackedList QObject_FindChildren3(void* ptr, void* re, long long options)
{
	return ({ QList<QObject*>* tmpValuec8bf88 = new QList<QObject*>(static_cast<QObject*>(ptr)->findChildren<QObject*>(*static_cast<QRegularExpression*>(re), static_cast<Qt::FindChildOption>(options))); QtCore_PackedList { tmpValuec8bf88, tmpValuec8bf88->size() }; });
}

struct QtCore_PackedString QObject_ObjectName(void* ptr)
{
	return ({ QByteArray te77be1 = static_cast<QObject*>(ptr)->objectName().toUtf8(); QtCore_PackedString { const_cast<char*>(te77be1.prepend("WHITESPACE").constData()+10), te77be1.size()-10 }; });
}

void QObject_ConnectObjectNameChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QObject*>(ptr), &QObject::objectNameChanged, static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(const QString &)>(&MyQObject::Signal_ObjectNameChanged), static_cast<Qt::ConnectionType>(t));
}

void QObject_DisconnectObjectNameChanged(void* ptr)
{
	QObject::disconnect(static_cast<QObject*>(ptr), &QObject::objectNameChanged, static_cast<MyQObject*>(ptr), static_cast<void (MyQObject::*)(const QString &)>(&MyQObject::Signal_ObjectNameChanged));
}

void* QObject_Parent(void* ptr)
{
	return static_cast<QObject*>(ptr)->parent();
}

void* QObject_Property(void* ptr, char* name)
{
	return new QVariant(static_cast<QObject*>(ptr)->property(const_cast<const char*>(name)));
}

void QObject_TimerEvent(void* ptr, void* event)
{
	static_cast<QObject*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QObject_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QTimer*>(static_cast<QObject*>(ptr))) {
		static_cast<QTimer*>(ptr)->QTimer::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QMimeData*>(static_cast<QObject*>(ptr))) {
		static_cast<QMimeData*>(ptr)->QMimeData::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QCoreApplication*>(static_cast<QObject*>(ptr))) {
		static_cast<QCoreApplication*>(ptr)->QCoreApplication::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QFile*>(static_cast<QObject*>(ptr))) {
		static_cast<QFile*>(ptr)->QFile::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QFileDevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QFileDevice*>(ptr)->QFileDevice::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QIODevice*>(static_cast<QObject*>(ptr))) {
		static_cast<QIODevice*>(ptr)->QIODevice::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QSortFilterProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QAbstractProxyModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractProxyModel*>(ptr)->QAbstractProxyModel::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QAbstractTableModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractTableModel*>(ptr)->QAbstractTableModel::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QAbstractItemModel*>(static_cast<QObject*>(ptr))) {
		static_cast<QAbstractItemModel*>(ptr)->QAbstractItemModel::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QObject*>(ptr)->QObject::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

struct QtCore_PackedString QObject_QObject_Tr(char* sourceText, char* disambiguation, int n)
{
	return ({ QByteArray te5b32b = QObject::tr(const_cast<const char*>(sourceText), const_cast<const char*>(disambiguation), n).toUtf8(); QtCore_PackedString { const_cast<char*>(te5b32b.prepend("WHITESPACE").constData()+10), te5b32b.size()-10 }; });
}

void QObject_DestroyQObject(void* ptr)
{
	static_cast<QObject*>(ptr)->~QObject();
}

void QObject_DestroyQObjectDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QObject_ToVariant(void* ptr)
{
	return new QVariant(QVariant::fromValue(static_cast<QObject*>(ptr)));
}

void* QObject___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QObject___children_setList(void* ptr, void* i)
{
	static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QObject___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QObject___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QObject___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QObject___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QObject___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QObject___findChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QObject___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QObject___findChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QObject___findChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QObject___findChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QObject___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QObject___findChildren_setList3(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QObject___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QObject___qFindChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QObject___qFindChildren_setList(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QObject___qFindChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QObject___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QObject___qFindChildren_setList2(void* ptr, void* i)
{
	static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
}

void* QObject___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPersistentModelIndex_NewQPersistentModelIndex2(void* index)
{
	return new QPersistentModelIndex(*static_cast<QModelIndex*>(index));
}

void* QPersistentModelIndex_NewQPersistentModelIndex3(void* other)
{
	return new QPersistentModelIndex(*static_cast<QPersistentModelIndex*>(other));
}

void* QPersistentModelIndex_NewQPersistentModelIndex4(void* other)
{
	return new QPersistentModelIndex(*static_cast<QPersistentModelIndex*>(other));
}

int QPersistentModelIndex_Column(void* ptr)
{
	return static_cast<QPersistentModelIndex*>(ptr)->column();
}

void* QPersistentModelIndex_Data(void* ptr, int role)
{
	return new QVariant(static_cast<QPersistentModelIndex*>(ptr)->data(role));
}

long long QPersistentModelIndex_Flags(void* ptr)
{
	return static_cast<QPersistentModelIndex*>(ptr)->flags();
}

void* QPersistentModelIndex_Model(void* ptr)
{
	return const_cast<QAbstractItemModel*>(static_cast<QPersistentModelIndex*>(ptr)->model());
}

void* QPersistentModelIndex_Parent(void* ptr)
{
	return new QModelIndex(static_cast<QPersistentModelIndex*>(ptr)->parent());
}

int QPersistentModelIndex_Row(void* ptr)
{
	return static_cast<QPersistentModelIndex*>(ptr)->row();
}

void* QPersistentModelIndex_Sibling(void* ptr, int row, int column)
{
	return new QModelIndex(static_cast<QPersistentModelIndex*>(ptr)->sibling(row, column));
}

void* QPoint_NewQPoint()
{
	return new QPoint();
}

void* QPoint_NewQPoint2(int xpos, int ypos)
{
	return new QPoint(xpos, ypos);
}

int QPoint_X(void* ptr)
{
	return static_cast<QPoint*>(ptr)->x();
}

int QPoint_Y(void* ptr)
{
	return static_cast<QPoint*>(ptr)->y();
}

void* QPointF_NewQPointF()
{
	return new QPointF();
}

void* QPointF_NewQPointF2(void* point)
{
	return new QPointF(*static_cast<QPoint*>(point));
}

void* QPointF_NewQPointF3(double xpos, double ypos)
{
	return new QPointF(xpos, ypos);
}

double QPointF_X(void* ptr)
{
	return static_cast<QPointF*>(ptr)->x();
}

double QPointF_Y(void* ptr)
{
	return static_cast<QPointF*>(ptr)->y();
}

void* QRect_NewQRect()
{
	return new QRect();
}

void* QRect_NewQRect2(void* topLeft, void* bottomRight)
{
	return new QRect(*static_cast<QPoint*>(topLeft), *static_cast<QPoint*>(bottomRight));
}

void* QRect_NewQRect3(void* topLeft, void* size)
{
	return new QRect(*static_cast<QPoint*>(topLeft), *static_cast<QSize*>(size));
}

void* QRect_NewQRect4(int x, int y, int width, int height)
{
	return new QRect(x, y, width, height);
}

char QRect_Contains(void* ptr, void* point, char proper)
{
	return static_cast<QRect*>(ptr)->contains(*static_cast<QPoint*>(point), proper != 0);
}

char QRect_Contains2(void* ptr, void* rectangle, char proper)
{
	return static_cast<QRect*>(ptr)->contains(*static_cast<QRect*>(rectangle), proper != 0);
}

char QRect_Contains3(void* ptr, int x, int y)
{
	return static_cast<QRect*>(ptr)->contains(x, y);
}

char QRect_Contains4(void* ptr, int x, int y, char proper)
{
	return static_cast<QRect*>(ptr)->contains(x, y, proper != 0);
}

int QRect_Height(void* ptr)
{
	return static_cast<QRect*>(ptr)->height();
}

int QRect_Left(void* ptr)
{
	return static_cast<QRect*>(ptr)->left();
}

int QRect_Right(void* ptr)
{
	return static_cast<QRect*>(ptr)->right();
}

void QRect_SetHeight(void* ptr, int height)
{
	static_cast<QRect*>(ptr)->setHeight(height);
}

void* QRect_Size(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QRect*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
}

int QRect_Top(void* ptr)
{
	return static_cast<QRect*>(ptr)->top();
}

int QRect_Width(void* ptr)
{
	return static_cast<QRect*>(ptr)->width();
}

int QRect_X(void* ptr)
{
	return static_cast<QRect*>(ptr)->x();
}

int QRect_Y(void* ptr)
{
	return static_cast<QRect*>(ptr)->y();
}

void* QRectF_NewQRectF()
{
	return new QRectF();
}

void* QRectF_NewQRectF2(void* topLeft, void* size)
{
	return new QRectF(*static_cast<QPointF*>(topLeft), *static_cast<QSizeF*>(size));
}

void* QRectF_NewQRectF3(void* topLeft, void* bottomRight)
{
	return new QRectF(*static_cast<QPointF*>(topLeft), *static_cast<QPointF*>(bottomRight));
}

void* QRectF_NewQRectF4(double x, double y, double width, double height)
{
	return new QRectF(x, y, width, height);
}

void* QRectF_NewQRectF5(void* rectangle)
{
	return new QRectF(*static_cast<QRect*>(rectangle));
}

char QRectF_Contains(void* ptr, void* point)
{
	return static_cast<QRectF*>(ptr)->contains(*static_cast<QPointF*>(point));
}

char QRectF_Contains2(void* ptr, void* rectangle)
{
	return static_cast<QRectF*>(ptr)->contains(*static_cast<QRectF*>(rectangle));
}

char QRectF_Contains3(void* ptr, double x, double y)
{
	return static_cast<QRectF*>(ptr)->contains(x, y);
}

double QRectF_Height(void* ptr)
{
	return static_cast<QRectF*>(ptr)->height();
}

double QRectF_Left(void* ptr)
{
	return static_cast<QRectF*>(ptr)->left();
}

double QRectF_Right(void* ptr)
{
	return static_cast<QRectF*>(ptr)->right();
}

void QRectF_SetHeight(void* ptr, double height)
{
	static_cast<QRectF*>(ptr)->setHeight(height);
}

void* QRectF_Size(void* ptr)
{
	return ({ QSizeF tmpValue = static_cast<QRectF*>(ptr)->size(); new QSizeF(tmpValue.width(), tmpValue.height()); });
}

double QRectF_Top(void* ptr)
{
	return static_cast<QRectF*>(ptr)->top();
}

double QRectF_Width(void* ptr)
{
	return static_cast<QRectF*>(ptr)->width();
}

double QRectF_X(void* ptr)
{
	return static_cast<QRectF*>(ptr)->x();
}

double QRectF_Y(void* ptr)
{
	return static_cast<QRectF*>(ptr)->y();
}

void* QRegExp_NewQRegExp()
{
	return new QRegExp();
}

void* QRegExp_NewQRegExp2(struct QtCore_PackedString pattern, long long cs, long long syntax)
{
	return new QRegExp(QString::fromUtf8(pattern.data, pattern.len), static_cast<Qt::CaseSensitivity>(cs), static_cast<QRegExp::PatternSyntax>(syntax));
}

void* QRegExp_NewQRegExp3(void* rx)
{
	return new QRegExp(*static_cast<QRegExp*>(rx));
}

long long QRegExp_CaseSensitivity(void* ptr)
{
	return static_cast<QRegExp*>(ptr)->caseSensitivity();
}

void QRegExp_DestroyQRegExp(void* ptr)
{
	static_cast<QRegExp*>(ptr)->~QRegExp();
}

void* QRegularExpression_NewQRegularExpression()
{
	return new QRegularExpression();
}

void* QRegularExpression_NewQRegularExpression2(struct QtCore_PackedString pattern, long long options)
{
	return new QRegularExpression(QString::fromUtf8(pattern.data, pattern.len), static_cast<QRegularExpression::PatternOption>(options));
}

void* QRegularExpression_NewQRegularExpression3(void* re)
{
	return new QRegularExpression(*static_cast<QRegularExpression*>(re));
}

void* QRegularExpression_Match(void* ptr, struct QtCore_PackedString subject, int offset, long long matchType, long long matchOptions)
{
	return new QRegularExpressionMatch(static_cast<QRegularExpression*>(ptr)->match(QString::fromUtf8(subject.data, subject.len), offset, static_cast<QRegularExpression::MatchType>(matchType), static_cast<QRegularExpression::MatchOption>(matchOptions)));
}

void* QRegularExpression_Match2(void* ptr, void* subjectRef, int offset, long long matchType, long long matchOptions)
{
	return new QRegularExpressionMatch(static_cast<QRegularExpression*>(ptr)->match(*static_cast<QStringRef*>(subjectRef), offset, static_cast<QRegularExpression::MatchType>(matchType), static_cast<QRegularExpression::MatchOption>(matchOptions)));
}

void QRegularExpression_DestroyQRegularExpression(void* ptr)
{
	static_cast<QRegularExpression*>(ptr)->~QRegularExpression();
}

void* QRegularExpressionMatch_NewQRegularExpressionMatch()
{
	return new QRegularExpressionMatch();
}

void* QRegularExpressionMatch_NewQRegularExpressionMatch2(void* match)
{
	return new QRegularExpressionMatch(*static_cast<QRegularExpressionMatch*>(match));
}

void QRegularExpressionMatch_DestroyQRegularExpressionMatch(void* ptr)
{
	static_cast<QRegularExpressionMatch*>(ptr)->~QRegularExpressionMatch();
}

void* QSize_NewQSize()
{
	return new QSize();
}

void* QSize_NewQSize2(int width, int height)
{
	return new QSize(width, height);
}

int QSize_Height(void* ptr)
{
	return static_cast<QSize*>(ptr)->height();
}

void QSize_Scale(void* ptr, int width, int height, long long mode)
{
	static_cast<QSize*>(ptr)->scale(width, height, static_cast<Qt::AspectRatioMode>(mode));
}

void QSize_Scale2(void* ptr, void* size, long long mode)
{
	static_cast<QSize*>(ptr)->scale(*static_cast<QSize*>(size), static_cast<Qt::AspectRatioMode>(mode));
}

void QSize_SetHeight(void* ptr, int height)
{
	static_cast<QSize*>(ptr)->setHeight(height);
}

int QSize_Width(void* ptr)
{
	return static_cast<QSize*>(ptr)->width();
}

void* QSizeF_NewQSizeF()
{
	return new QSizeF();
}

void* QSizeF_NewQSizeF2(void* size)
{
	return new QSizeF(*static_cast<QSize*>(size));
}

void* QSizeF_NewQSizeF3(double width, double height)
{
	return new QSizeF(width, height);
}

double QSizeF_Height(void* ptr)
{
	return static_cast<QSizeF*>(ptr)->height();
}

void QSizeF_Scale(void* ptr, double width, double height, long long mode)
{
	static_cast<QSizeF*>(ptr)->scale(width, height, static_cast<Qt::AspectRatioMode>(mode));
}

void QSizeF_Scale2(void* ptr, void* size, long long mode)
{
	static_cast<QSizeF*>(ptr)->scale(*static_cast<QSizeF*>(size), static_cast<Qt::AspectRatioMode>(mode));
}

void QSizeF_SetHeight(void* ptr, double height)
{
	static_cast<QSizeF*>(ptr)->setHeight(height);
}

double QSizeF_Width(void* ptr)
{
	return static_cast<QSizeF*>(ptr)->width();
}

class MyQSortFilterProxyModel: public QSortFilterProxyModel
{
public:
	MyQSortFilterProxyModel(QObject *parent = Q_NULLPTR) : QSortFilterProxyModel(parent) {QSortFilterProxyModel_QSortFilterProxyModel_QRegisterMetaType();};
	QModelIndex buddy(const QModelIndex & index) const { return *static_cast<QModelIndex*>(callbackQAbstractItemModel_Buddy(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool canFetchMore(const QModelIndex & parent) const { return callbackQAbstractItemModel_CanFetchMore(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	int columnCount(const QModelIndex & parent) const { return callbackQSortFilterProxyModel_ColumnCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	QVariant data(const QModelIndex & index, int role) const { return *static_cast<QVariant*>(callbackQAbstractProxyModel_Data(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index), role)); };
	bool dropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) { return callbackQAbstractItemModel_DropMimeData(this, const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	void fetchMore(const QModelIndex & parent) { callbackQAbstractItemModel_FetchMore(this, const_cast<QModelIndex*>(&parent)); };
	Qt::ItemFlags flags(const QModelIndex & index) const { return static_cast<Qt::ItemFlag>(callbackQAbstractItemModel_Flags(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	bool hasChildren(const QModelIndex & parent) const { return callbackQAbstractItemModel_HasChildren(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)) != 0; };
	QVariant headerData(int section, Qt::Orientation orientation, int role) const { return *static_cast<QVariant*>(callbackQAbstractItemModel_HeaderData(const_cast<void*>(static_cast<const void*>(this)), section, orientation, role)); };
	QModelIndex index(int row, int column, const QModelIndex & parent) const { return *static_cast<QModelIndex*>(callbackQSortFilterProxyModel_Index(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&parent))); };
	bool insertColumns(int column, int count, const QModelIndex & parent) { return callbackQAbstractItemModel_InsertColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool insertRows(int row, int count, const QModelIndex & parent) { return callbackQAbstractItemModel_InsertRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	QList<QModelIndex> match(const QModelIndex & start, int role, const QVariant & value, int hits, Qt::MatchFlags flags) const { return ({ QList<QModelIndex>* tmpP = static_cast<QList<QModelIndex>*>(callbackQAbstractItemModel_Match(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&start), role, const_cast<QVariant*>(&value), hits, flags)); QList<QModelIndex> tmpV = *tmpP; tmpP->~QList(); free(tmpP); tmpV; }); };
	QMimeData * mimeData(const QModelIndexList & indexes) const { return static_cast<QMimeData*>(callbackQAbstractItemModel_MimeData(const_cast<void*>(static_cast<const void*>(this)), ({ QList<QModelIndex>* tmpValuee0adf2 = new QList<QModelIndex>(indexes); QtCore_PackedList { tmpValuee0adf2, tmpValuee0adf2->size() }; }))); };
	QStringList mimeTypes() const { return ({ QtCore_PackedString tempVal = callbackQAbstractItemModel_MimeTypes(const_cast<void*>(static_cast<const void*>(this))); QStringList ret = QString::fromUtf8(tempVal.data, tempVal.len).split("¡¦!", QString::SkipEmptyParts); free(tempVal.data); ret; }); };
	QModelIndex parent(const QModelIndex & child) const { return *static_cast<QModelIndex*>(callbackQSortFilterProxyModel_Parent(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&child))); };
	bool removeColumns(int column, int count, const QModelIndex & parent) { return callbackQAbstractItemModel_RemoveColumns(this, column, count, const_cast<QModelIndex*>(&parent)) != 0; };
	bool removeRows(int row, int count, const QModelIndex & parent) { return callbackQAbstractItemModel_RemoveRows(this, row, count, const_cast<QModelIndex*>(&parent)) != 0; };
	int rowCount(const QModelIndex & parent) const { return callbackQSortFilterProxyModel_RowCount(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&parent)); };
	bool setData(const QModelIndex & index, const QVariant & value, int role) { return callbackQAbstractItemModel_SetData(this, const_cast<QModelIndex*>(&index), const_cast<QVariant*>(&value), role) != 0; };
	void setFilterFixedString(const QString & pattern) { QByteArray t91cc2e = pattern.toUtf8(); QtCore_PackedString patternPacked = { const_cast<char*>(t91cc2e.prepend("WHITESPACE").constData()+10), t91cc2e.size()-10 };callbackQSortFilterProxyModel_SetFilterFixedString(this, patternPacked); };
	bool setHeaderData(int section, Qt::Orientation orientation, const QVariant & value, int role) { return callbackQAbstractItemModel_SetHeaderData(this, section, orientation, const_cast<QVariant*>(&value), role) != 0; };
	void setSourceModel(QAbstractItemModel * sourceModel) { callbackQAbstractProxyModel_SetSourceModel(this, sourceModel); };
	QModelIndex sibling(int row, int column, const QModelIndex & idx) const { return *static_cast<QModelIndex*>(callbackQAbstractItemModel_Sibling(const_cast<void*>(static_cast<const void*>(this)), row, column, const_cast<QModelIndex*>(&idx))); };
	void sort(int column, Qt::SortOrder order) { callbackQAbstractItemModel_Sort(this, column, order); };
	QSize span(const QModelIndex & index) const { return *static_cast<QSize*>(callbackQAbstractItemModel_Span(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&index))); };
	Qt::DropActions supportedDropActions() const { return static_cast<Qt::DropAction>(callbackQAbstractItemModel_SupportedDropActions(const_cast<void*>(static_cast<const void*>(this)))); };
	 ~MyQSortFilterProxyModel() { callbackQSortFilterProxyModel_DestroyQSortFilterProxyModel(this); };
	bool canDropMimeData(const QMimeData * data, Qt::DropAction action, int row, int column, const QModelIndex & parent) const { return callbackQAbstractItemModel_CanDropMimeData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QMimeData*>(data), action, row, column, const_cast<QModelIndex*>(&parent)) != 0; };
	QMap<int, QVariant> itemData(const QModelIndex & proxyIndex) const { return ({ QMap<int, QVariant>* tmpP = static_cast<QMap<int, QVariant>*>(callbackQAbstractItemModel_ItemData(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&proxyIndex))); QMap<int, QVariant> tmpV = *tmpP; tmpP->~QMap(); free(tmpP); tmpV; }); };
	QModelIndex mapFromSource(const QModelIndex & sourceIndex) const { return *static_cast<QModelIndex*>(callbackQSortFilterProxyModel_MapFromSource(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&sourceIndex))); };
	QModelIndex mapToSource(const QModelIndex & proxyIndex) const { return *static_cast<QModelIndex*>(callbackQSortFilterProxyModel_MapToSource(const_cast<void*>(static_cast<const void*>(this)), const_cast<QModelIndex*>(&proxyIndex))); };
	void resetInternalData() { callbackQAbstractItemModel_ResetInternalData(this); };
	void revert() { callbackQAbstractProxyModel_Revert(this); };
	bool setItemData(const QModelIndex & index, const QMap<int, QVariant> & roles) { return callbackQAbstractItemModel_SetItemData(this, const_cast<QModelIndex*>(&index), ({ QMap<int, QVariant>* tmpValue037c88 = new QMap<int, QVariant>(roles); QtCore_PackedList { tmpValue037c88, tmpValue037c88->size() }; })) != 0; };
	bool submit() { return callbackQAbstractProxyModel_Submit(this) != 0; };
	Qt::DropActions supportedDragActions() const { return static_cast<Qt::DropAction>(callbackQAbstractItemModel_SupportedDragActions(const_cast<void*>(static_cast<const void*>(this)))); };
	void Signal_ColumnsAboutToBeInserted(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_ColumnsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationColumn) { callbackQAbstractItemModel_ColumnsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationColumn); };
	void Signal_ColumnsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_ColumnsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsInserted(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_ColumnsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_ColumnsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int column) { callbackQAbstractItemModel_ColumnsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), column); };
	void Signal_ColumnsRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_ColumnsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_DataChanged(const QModelIndex & topLeft, const QModelIndex & bottomRight, const QVector<int> & roles) { callbackQAbstractItemModel_DataChanged(this, const_cast<QModelIndex*>(&topLeft), const_cast<QModelIndex*>(&bottomRight), ({ QVector<int>* tmpValue037c88 = new QVector<int>(roles); QtCore_PackedList { tmpValue037c88, tmpValue037c88->size() }; })); };
	void Signal_HeaderDataChanged(Qt::Orientation orientation, int first, int last) { callbackQAbstractItemModel_HeaderDataChanged(this, orientation, first, last); };
	void Signal_LayoutAboutToBeChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQAbstractItemModel_LayoutAboutToBeChanged(this, ({ QList<QPersistentModelIndex>* tmpValuea664f1 = new QList<QPersistentModelIndex>(parents); QtCore_PackedList { tmpValuea664f1, tmpValuea664f1->size() }; }), hint); };
	void Signal_LayoutChanged(const QList<QPersistentModelIndex> & parents, QAbstractItemModel::LayoutChangeHint hint) { callbackQAbstractItemModel_LayoutChanged(this, ({ QList<QPersistentModelIndex>* tmpValuea664f1 = new QList<QPersistentModelIndex>(parents); QtCore_PackedList { tmpValuea664f1, tmpValuea664f1->size() }; }), hint); };
	void Signal_ModelAboutToBeReset() { callbackQAbstractItemModel_ModelAboutToBeReset(this); };
	void Signal_ModelReset() { callbackQAbstractItemModel_ModelReset(this); };
	bool moveColumns(const QModelIndex & sourceParent, int sourceColumn, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQAbstractItemModel_MoveColumns(this, const_cast<QModelIndex*>(&sourceParent), sourceColumn, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	bool moveRows(const QModelIndex & sourceParent, int sourceRow, int count, const QModelIndex & destinationParent, int destinationChild) { return callbackQAbstractItemModel_MoveRows(this, const_cast<QModelIndex*>(&sourceParent), sourceRow, count, const_cast<QModelIndex*>(&destinationParent), destinationChild) != 0; };
	QHash<int, QByteArray> roleNames() const { return ({ QHash<int, QByteArray>* tmpP = static_cast<QHash<int, QByteArray>*>(callbackQAbstractItemModel_RoleNames(const_cast<void*>(static_cast<const void*>(this)))); QHash<int, QByteArray> tmpV = *tmpP; tmpP->~QHash(); free(tmpP); tmpV; }); };
	void Signal_RowsAboutToBeInserted(const QModelIndex & parent, int start, int end) { callbackQAbstractItemModel_RowsAboutToBeInserted(this, const_cast<QModelIndex*>(&parent), start, end); };
	void Signal_RowsAboutToBeMoved(const QModelIndex & sourceParent, int sourceStart, int sourceEnd, const QModelIndex & destinationParent, int destinationRow) { callbackQAbstractItemModel_RowsAboutToBeMoved(this, const_cast<QModelIndex*>(&sourceParent), sourceStart, sourceEnd, const_cast<QModelIndex*>(&destinationParent), destinationRow); };
	void Signal_RowsAboutToBeRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_RowsAboutToBeRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsInserted(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_RowsInserted(this, const_cast<QModelIndex*>(&parent), first, last); };
	void Signal_RowsMoved(const QModelIndex & parent, int start, int end, const QModelIndex & destination, int row) { callbackQAbstractItemModel_RowsMoved(this, const_cast<QModelIndex*>(&parent), start, end, const_cast<QModelIndex*>(&destination), row); };
	void Signal_RowsRemoved(const QModelIndex & parent, int first, int last) { callbackQAbstractItemModel_RowsRemoved(this, const_cast<QModelIndex*>(&parent), first, last); };
	void childEvent(QChildEvent * event) { callbackQObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQObject_CustomEvent(this, event); };
	void deleteLater() { callbackQObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQObject_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQSortFilterProxyModel*)

int QSortFilterProxyModel_QSortFilterProxyModel_QRegisterMetaType(){qRegisterMetaType<QSortFilterProxyModel*>(); return qRegisterMetaType<MyQSortFilterProxyModel*>();}

void* QSortFilterProxyModel_NewQSortFilterProxyModel(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQSortFilterProxyModel(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSortFilterProxyModel(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQSortFilterProxyModel(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQSortFilterProxyModel(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQSortFilterProxyModel(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQSortFilterProxyModel(static_cast<QWindow*>(parent));
	} else {
		return new MyQSortFilterProxyModel(static_cast<QObject*>(parent));
	}
}

int QSortFilterProxyModel_ColumnCount(void* ptr, void* parent)
{
	return static_cast<QSortFilterProxyModel*>(ptr)->columnCount(*static_cast<QModelIndex*>(parent));
}

int QSortFilterProxyModel_ColumnCountDefault(void* ptr, void* parent)
{
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::columnCount(*static_cast<QModelIndex*>(parent));
}

long long QSortFilterProxyModel_FilterCaseSensitivity(void* ptr)
{
	return static_cast<QSortFilterProxyModel*>(ptr)->filterCaseSensitivity();
}

int QSortFilterProxyModel_FilterKeyColumn(void* ptr)
{
	return static_cast<QSortFilterProxyModel*>(ptr)->filterKeyColumn();
}

void* QSortFilterProxyModel_Index(void* ptr, int row, int column, void* parent)
{
	return new QModelIndex(static_cast<QSortFilterProxyModel*>(ptr)->index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* QSortFilterProxyModel_IndexDefault(void* ptr, int row, int column, void* parent)
{
		return new QModelIndex(static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::index(row, column, *static_cast<QModelIndex*>(parent)));
}

void* QSortFilterProxyModel_Parent(void* ptr, void* child)
{
	return new QModelIndex(static_cast<QSortFilterProxyModel*>(ptr)->parent(*static_cast<QModelIndex*>(child)));
}

void* QSortFilterProxyModel_ParentDefault(void* ptr, void* child)
{
		return new QModelIndex(static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::parent(*static_cast<QModelIndex*>(child)));
}

int QSortFilterProxyModel_RowCount(void* ptr, void* parent)
{
	return static_cast<QSortFilterProxyModel*>(ptr)->rowCount(*static_cast<QModelIndex*>(parent));
}

int QSortFilterProxyModel_RowCountDefault(void* ptr, void* parent)
{
		return static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::rowCount(*static_cast<QModelIndex*>(parent));
}

void QSortFilterProxyModel_SetFilterCaseSensitivity(void* ptr, long long cs)
{
	static_cast<QSortFilterProxyModel*>(ptr)->setFilterCaseSensitivity(static_cast<Qt::CaseSensitivity>(cs));
}

void QSortFilterProxyModel_SetFilterFixedString(void* ptr, struct QtCore_PackedString pattern)
{
	QMetaObject::invokeMethod(static_cast<QSortFilterProxyModel*>(ptr), "setFilterFixedString", Q_ARG(const QString, QString::fromUtf8(pattern.data, pattern.len)));
}

void QSortFilterProxyModel_SetFilterFixedStringDefault(void* ptr, struct QtCore_PackedString pattern)
{
		static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::setFilterFixedString(QString::fromUtf8(pattern.data, pattern.len));
}

void QSortFilterProxyModel_SetFilterKeyColumn(void* ptr, int column)
{
	static_cast<QSortFilterProxyModel*>(ptr)->setFilterKeyColumn(column);
}

long long QSortFilterProxyModel_SortOrder(void* ptr)
{
	return static_cast<QSortFilterProxyModel*>(ptr)->sortOrder();
}

void QSortFilterProxyModel_DestroyQSortFilterProxyModel(void* ptr)
{
	static_cast<QSortFilterProxyModel*>(ptr)->~QSortFilterProxyModel();
}

void QSortFilterProxyModel_DestroyQSortFilterProxyModelDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QSortFilterProxyModel_MapFromSource(void* ptr, void* sourceIndex)
{
	return new QModelIndex(static_cast<QSortFilterProxyModel*>(ptr)->mapFromSource(*static_cast<QModelIndex*>(sourceIndex)));
}

void* QSortFilterProxyModel_MapFromSourceDefault(void* ptr, void* sourceIndex)
{
		return new QModelIndex(static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::mapFromSource(*static_cast<QModelIndex*>(sourceIndex)));
}

void* QSortFilterProxyModel_MapToSource(void* ptr, void* proxyIndex)
{
	return new QModelIndex(static_cast<QSortFilterProxyModel*>(ptr)->mapToSource(*static_cast<QModelIndex*>(proxyIndex)));
}

void* QSortFilterProxyModel_MapToSourceDefault(void* ptr, void* proxyIndex)
{
		return new QModelIndex(static_cast<QSortFilterProxyModel*>(ptr)->QSortFilterProxyModel::mapToSource(*static_cast<QModelIndex*>(proxyIndex)));
}

void* QStringRef_NewQStringRef()
{
	return new QStringRef();
}

void* QStringRef_NewQStringRef2(struct QtCore_PackedString stri, int position, int length)
{
	return new QStringRef(new QString(QString::fromUtf8(stri.data, stri.len)), position, length);
}

void* QStringRef_NewQStringRef3(struct QtCore_PackedString stri)
{
	return new QStringRef(new QString(QString::fromUtf8(stri.data, stri.len)));
}

void* QStringRef_NewQStringRef4(void* other)
{
	return new QStringRef(*static_cast<QStringRef*>(other));
}

void* QStringRef_At(void* ptr, int position)
{
	return new QChar(static_cast<QStringRef*>(ptr)->at(position));
}

void* QStringRef_Back(void* ptr)
{
	return new QChar(static_cast<QStringRef*>(ptr)->back());
}

char QStringRef_Contains(void* ptr, struct QtCore_PackedString str, long long cs)
{
	return static_cast<QStringRef*>(ptr)->contains(QString::fromUtf8(str.data, str.len), static_cast<Qt::CaseSensitivity>(cs));
}

char QStringRef_Contains2(void* ptr, void* ch, long long cs)
{
	return static_cast<QStringRef*>(ptr)->contains(*static_cast<QChar*>(ch), static_cast<Qt::CaseSensitivity>(cs));
}

char QStringRef_Contains3(void* ptr, void* str, long long cs)
{
	return static_cast<QStringRef*>(ptr)->contains(*static_cast<QLatin1String*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

char QStringRef_Contains4(void* ptr, void* str, long long cs)
{
	return static_cast<QStringRef*>(ptr)->contains(*static_cast<QStringRef*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Count(void* ptr)
{
	return static_cast<QStringRef*>(ptr)->count();
}

int QStringRef_Count2(void* ptr, struct QtCore_PackedString str, long long cs)
{
	return static_cast<QStringRef*>(ptr)->count(QString::fromUtf8(str.data, str.len), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Count3(void* ptr, void* ch, long long cs)
{
	return static_cast<QStringRef*>(ptr)->count(*static_cast<QChar*>(ch), static_cast<Qt::CaseSensitivity>(cs));
}

int QStringRef_Count4(void* ptr, void* str, long long cs)
{
	return static_cast<QStringRef*>(ptr)->count(*static_cast<QStringRef*>(str), static_cast<Qt::CaseSensitivity>(cs));
}

void* QStringRef_Data(void* ptr)
{
	return const_cast<QChar*>(static_cast<QStringRef*>(ptr)->data());
}

void* QStringRef_Left(void* ptr, int n)
{
	return new QStringRef(static_cast<QStringRef*>(ptr)->left(n));
}

void* QStringRef_Right(void* ptr, int n)
{
	return new QStringRef(static_cast<QStringRef*>(ptr)->right(n));
}

int QStringRef_Size(void* ptr)
{
	return static_cast<QStringRef*>(ptr)->size();
}

struct QtCore_PackedList QStringRef_Split(void* ptr, struct QtCore_PackedString sep, long long behavior, long long cs)
{
	return ({ QVector<QStringRef>* tmpValued1bf8d = new QVector<QStringRef>(static_cast<QStringRef*>(ptr)->split(QString::fromUtf8(sep.data, sep.len), static_cast<QString::SplitBehavior>(behavior), static_cast<Qt::CaseSensitivity>(cs))); QtCore_PackedList { tmpValued1bf8d, tmpValued1bf8d->size() }; });
}

struct QtCore_PackedList QStringRef_Split2(void* ptr, void* sep, long long behavior, long long cs)
{
	return ({ QVector<QStringRef>* tmpValue643751 = new QVector<QStringRef>(static_cast<QStringRef*>(ptr)->split(*static_cast<QChar*>(sep), static_cast<QString::SplitBehavior>(behavior), static_cast<Qt::CaseSensitivity>(cs))); QtCore_PackedList { tmpValue643751, tmpValue643751->size() }; });
}

struct QtCore_PackedString QStringRef_String(void* ptr)
{
	return ({ QByteArray t75a779 = static_cast<QStringRef*>(ptr)->string()->toUtf8(); QtCore_PackedString { const_cast<char*>(t75a779.prepend("WHITESPACE").constData()+10), t75a779.size()-10 }; });
}

int QStringRef_ToInt(void* ptr, char* ok, int base)
{
	return static_cast<QStringRef*>(ptr)->toInt(reinterpret_cast<bool*>(ok), base);
}

void QStringRef_DestroyQStringRef(void* ptr)
{
	static_cast<QStringRef*>(ptr)->~QStringRef();
}

void* QStringRef___split_atList(void* ptr, int i)
{
	return new QStringRef(({QStringRef tmp = static_cast<QVector<QStringRef>*>(ptr)->at(i); if (i == static_cast<QVector<QStringRef>*>(ptr)->size()-1) { static_cast<QVector<QStringRef>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QStringRef___split_setList(void* ptr, void* i)
{
	static_cast<QVector<QStringRef>*>(ptr)->append(*static_cast<QStringRef*>(i));
}

void* QStringRef___split_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QStringRef>();
}

void* QStringRef___split_atList2(void* ptr, int i)
{
	return new QStringRef(({QStringRef tmp = static_cast<QVector<QStringRef>*>(ptr)->at(i); if (i == static_cast<QVector<QStringRef>*>(ptr)->size()-1) { static_cast<QVector<QStringRef>*>(ptr)->~QVector(); free(ptr); }; tmp; }));
}

void QStringRef___split_setList2(void* ptr, void* i)
{
	static_cast<QVector<QStringRef>*>(ptr)->append(*static_cast<QStringRef*>(i));
}

void* QStringRef___split_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QStringRef>();
}

unsigned int QStringRef___toUcs4_atList(void* ptr, int i)
{
	return ({uint tmp = static_cast<QVector<uint>*>(ptr)->at(i); if (i == static_cast<QVector<uint>*>(ptr)->size()-1) { static_cast<QVector<uint>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QStringRef___toUcs4_setList(void* ptr, unsigned int i)
{
	static_cast<QVector<uint>*>(ptr)->append(i);
}

void* QStringRef___toUcs4_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<uint>();
}

void* QStringView_NewQStringView()
{
	return new QStringView();
}

void* QStringView_NewQStringView7(struct QtCore_PackedString str)
{
	return new QStringView(QString::fromUtf8(str.data, str.len));
}

void* QStringView_NewQStringView8(void* str)
{
	return new QStringView(*static_cast<QStringRef*>(str));
}

void* QStringView_Back(void* ptr)
{
	return new QChar(static_cast<QStringView*>(ptr)->back());
}

unsigned int QStringView___convertToUcs4_atList(void* ptr, int i)
{
	return ({uint tmp = static_cast<QVector<uint>*>(ptr)->at(i); if (i == static_cast<QVector<uint>*>(ptr)->size()-1) { static_cast<QVector<uint>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QStringView___convertToUcs4_setList(void* ptr, unsigned int i)
{
	static_cast<QVector<uint>*>(ptr)->append(i);
}

void* QStringView___convertToUcs4_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<uint>();
}

unsigned int QStringView___toUcs4_atList(void* ptr, int i)
{
	return ({uint tmp = static_cast<QVector<uint>*>(ptr)->at(i); if (i == static_cast<QVector<uint>*>(ptr)->size()-1) { static_cast<QVector<uint>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QStringView___toUcs4_setList(void* ptr, unsigned int i)
{
	static_cast<QVector<uint>*>(ptr)->append(i);
}

void* QStringView___toUcs4_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<uint>();
}

int QSysInfo_WordSize_Type()
{
	return QSysInfo::WordSize;
}

void* QTime_NewQTime2()
{
	return new QTime();
}

void* QTime_NewQTime3(int h, int m, int s, int ms)
{
	return new QTime(h, m, s, ms);
}

int QTime_Minute(void* ptr)
{
	return static_cast<QTime*>(ptr)->minute();
}

int QTime_Second(void* ptr)
{
	return static_cast<QTime*>(ptr)->second();
}

void QTime_Start(void* ptr)
{
	static_cast<QTime*>(ptr)->start();
}

void* QTimeZone_NewQTimeZone()
{
	return new QTimeZone();
}

void* QTimeZone_NewQTimeZone2(void* ianaId)
{
	return new QTimeZone(*static_cast<QByteArray*>(ianaId));
}

void* QTimeZone_NewQTimeZone3(int offsetSeconds)
{
	return new QTimeZone(offsetSeconds);
}

void* QTimeZone_NewQTimeZone4(void* ianaId, int offsetSeconds, struct QtCore_PackedString name, struct QtCore_PackedString abbreviation, long long country, struct QtCore_PackedString comment)
{
	return new QTimeZone(*static_cast<QByteArray*>(ianaId), offsetSeconds, QString::fromUtf8(name.data, name.len), QString::fromUtf8(abbreviation.data, abbreviation.len), static_cast<QLocale::Country>(country), QString::fromUtf8(comment.data, comment.len));
}

void* QTimeZone_NewQTimeZone5(void* other)
{
	return new QTimeZone(*static_cast<QTimeZone*>(other));
}

void QTimeZone_DestroyQTimeZone(void* ptr)
{
	static_cast<QTimeZone*>(ptr)->~QTimeZone();
}

void* QTimeZone___availableTimeZoneIds_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTimeZone___availableTimeZoneIds_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTimeZone___availableTimeZoneIds_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QTimeZone___availableTimeZoneIds_atList2(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTimeZone___availableTimeZoneIds_setList2(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTimeZone___availableTimeZoneIds_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QTimeZone___availableTimeZoneIds_atList3(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTimeZone___availableTimeZoneIds_setList3(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTimeZone___availableTimeZoneIds_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QTimeZone___windowsIdToIanaIds_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTimeZone___windowsIdToIanaIds_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTimeZone___windowsIdToIanaIds_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QTimeZone___windowsIdToIanaIds_atList2(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTimeZone___windowsIdToIanaIds_setList2(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QTimeZone___windowsIdToIanaIds_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

class MyQTimer: public QTimer
{
public:
	MyQTimer(QObject *parent = Q_NULLPTR) : QTimer(parent) {QTimer_QTimer_QRegisterMetaType();};
	void start(int msec) { callbackQTimer_Start(this, msec); };
	void start() { callbackQTimer_Start2(this); };
	void timerEvent(QTimerEvent * e) { callbackQObject_TimerEvent(this, e); };
	 ~MyQTimer() { callbackQTimer_DestroyQTimer(this); };
	void childEvent(QChildEvent * event) { callbackQObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQObject_CustomEvent(this, event); };
	void deleteLater() { callbackQObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQObject_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtCore_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQObject_ObjectNameChanged(this, objectNamePacked); };
};

Q_DECLARE_METATYPE(MyQTimer*)

int QTimer_QTimer_QRegisterMetaType(){qRegisterMetaType<QTimer*>(); return qRegisterMetaType<MyQTimer*>();}

void* QTimer_NewQTimer(void* parent)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQTimer(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTimer(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQTimer(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQTimer(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQTimer(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQTimer(static_cast<QWindow*>(parent));
	} else {
		return new MyQTimer(static_cast<QObject*>(parent));
	}
}

void QTimer_Start(void* ptr, int msec)
{
	QMetaObject::invokeMethod(static_cast<QTimer*>(ptr), "start", Q_ARG(int, msec));
}

void QTimer_StartDefault(void* ptr, int msec)
{
		static_cast<QTimer*>(ptr)->QTimer::start(msec);
}

void QTimer_Start2(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QTimer*>(ptr), "start");
}

void QTimer_Start2Default(void* ptr)
{
		static_cast<QTimer*>(ptr)->QTimer::start();
}

void QTimer_DestroyQTimer(void* ptr)
{
	static_cast<QTimer*>(ptr)->~QTimer();
}

void QTimer_DestroyQTimerDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQTimerEvent: public QTimerEvent
{
public:
	MyQTimerEvent(int timerId) : QTimerEvent(timerId) {};
};

void* QTimerEvent_NewQTimerEvent(int timerId)
{
	return new MyQTimerEvent(timerId);
}

void* QUrl_NewQUrl()
{
	return new QUrl();
}

void* QUrl_NewQUrl2(void* other)
{
	return new QUrl(*static_cast<QUrl*>(other));
}

void* QUrl_NewQUrl3(struct QtCore_PackedString url, long long parsingMode)
{
	return new QUrl(QString::fromUtf8(url.data, url.len), static_cast<QUrl::ParsingMode>(parsingMode));
}

void* QUrl_NewQUrl4(void* other)
{
	return new QUrl(*static_cast<QUrl*>(other));
}

struct QtCore_PackedString QUrl_FileName(void* ptr, long long options)
{
	return ({ QByteArray t4c468f = static_cast<QUrl*>(ptr)->fileName(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8(); QtCore_PackedString { const_cast<char*>(t4c468f.prepend("WHITESPACE").constData()+10), t4c468f.size()-10 }; });
}

void* QUrl_QUrl_FromLocalFile(struct QtCore_PackedString localFile)
{
	return new QUrl(QUrl::fromLocalFile(QString::fromUtf8(localFile.data, localFile.len)));
}

struct QtCore_PackedString QUrl_Path(void* ptr, long long options)
{
	return ({ QByteArray t70ef65 = static_cast<QUrl*>(ptr)->path(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8(); QtCore_PackedString { const_cast<char*>(t70ef65.prepend("WHITESPACE").constData()+10), t70ef65.size()-10 }; });
}

struct QtCore_PackedString QUrl_Query(void* ptr, long long options)
{
	return ({ QByteArray t911c73 = static_cast<QUrl*>(ptr)->query(static_cast<QUrl::ComponentFormattingOption>(options)).toUtf8(); QtCore_PackedString { const_cast<char*>(t911c73.prepend("WHITESPACE").constData()+10), t911c73.size()-10 }; });
}

struct QtCore_PackedString QUrl_Url(void* ptr, long long options)
{
	return ({ QByteArray t2ea726 = static_cast<QUrl*>(ptr)->url(static_cast<QUrl::UrlFormattingOption>(options)).toUtf8(); QtCore_PackedString { const_cast<char*>(t2ea726.prepend("WHITESPACE").constData()+10), t2ea726.size()-10 }; });
}

void QUrl_DestroyQUrl(void* ptr)
{
	static_cast<QUrl*>(ptr)->~QUrl();
}

void* QUrl___allEncodedQueryItemValues_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QUrl___allEncodedQueryItemValues_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QUrl___allEncodedQueryItemValues_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QUrl___fromStringList_atList(void* ptr, int i)
{
	return new QUrl(({QUrl tmp = static_cast<QList<QUrl>*>(ptr)->at(i); if (i == static_cast<QList<QUrl>*>(ptr)->size()-1) { static_cast<QList<QUrl>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QUrl___fromStringList_setList(void* ptr, void* i)
{
	static_cast<QList<QUrl>*>(ptr)->append(*static_cast<QUrl*>(i));
}

void* QUrl___fromStringList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QUrl>();
}

void* QUrl___toStringList_urls_atList(void* ptr, int i)
{
	return new QUrl(({QUrl tmp = static_cast<QList<QUrl>*>(ptr)->at(i); if (i == static_cast<QList<QUrl>*>(ptr)->size()-1) { static_cast<QList<QUrl>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QUrl___toStringList_urls_setList(void* ptr, void* i)
{
	static_cast<QList<QUrl>*>(ptr)->append(*static_cast<QUrl*>(i));
}

void* QUrl___toStringList_urls_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QUrl>();
}

void* QUuid_NewQUuid2()
{
	return new QUuid();
}

void* QUuid_NewQUuid3(unsigned int l, unsigned short w1, unsigned short w2, char* b1, char* b2, char* b3, char* b4, char* b5, char* b6, char* b7, char* b8)
{
	return new QUuid(l, w1, w2, *static_cast<uchar*>(static_cast<void*>(b1)), *static_cast<uchar*>(static_cast<void*>(b2)), *static_cast<uchar*>(static_cast<void*>(b3)), *static_cast<uchar*>(static_cast<void*>(b4)), *static_cast<uchar*>(static_cast<void*>(b5)), *static_cast<uchar*>(static_cast<void*>(b6)), *static_cast<uchar*>(static_cast<void*>(b7)), *static_cast<uchar*>(static_cast<void*>(b8)));
}

void* QUuid_NewQUuid4(struct QtCore_PackedString text)
{
	return new QUuid(QString::fromUtf8(text.data, text.len));
}

void* QUuid_NewQUuid(void* text)
{
	return new QUuid(*static_cast<QByteArray*>(text));
}

long long QUuid_Variant(void* ptr)
{
	return static_cast<QUuid*>(ptr)->variant();
}

void* QVariant_NewQVariant()
{
	return new QVariant();
}

void* QVariant_NewQVariant2(long long ty)
{
	return new QVariant(static_cast<QVariant::Type>(ty));
}

void* QVariant_NewQVariant3(int typeId, void* copy)
{
	return new QVariant(typeId, copy);
}

void* QVariant_NewQVariant4(void* s)
{
	return new QVariant(*static_cast<QDataStream*>(s));
}

void* QVariant_NewQVariant5(int val)
{
	return new QVariant(val);
}

void* QVariant_NewQVariant6(unsigned int val)
{
	return new QVariant(val);
}

void* QVariant_NewQVariant7(long long val)
{
	return new QVariant(val);
}

void* QVariant_NewQVariant8(unsigned long long val)
{
	return new QVariant(val);
}

void* QVariant_NewQVariant9(char val)
{
	return new QVariant(val != 0);
}

void* QVariant_NewQVariant10(double val)
{
	return new QVariant(val);
}

void* QVariant_NewQVariant11(float val)
{
	return new QVariant(val);
}

void* QVariant_NewQVariant12(char* val)
{
	return new QVariant(const_cast<const char*>(val));
}

void* QVariant_NewQVariant13(void* val)
{
	return new QVariant(*static_cast<QByteArray*>(val));
}

void* QVariant_NewQVariant14(void* val)
{
	return new QVariant(*static_cast<QBitArray*>(val));
}

void* QVariant_NewQVariant15(struct QtCore_PackedString val)
{
	return new QVariant(QString::fromUtf8(val.data, val.len));
}

void* QVariant_NewQVariant16(void* val)
{
	return new QVariant(*static_cast<QLatin1String*>(val));
}

void* QVariant_NewQVariant17(struct QtCore_PackedString val)
{
	return new QVariant(QString::fromUtf8(val.data, val.len).split("¡¦!", QString::SkipEmptyParts));
}

void* QVariant_NewQVariant18(void* c)
{
	return new QVariant(*static_cast<QChar*>(c));
}

void* QVariant_NewQVariant19(void* val)
{
	return new QVariant(*static_cast<QDate*>(val));
}

void* QVariant_NewQVariant20(void* val)
{
	return new QVariant(*static_cast<QTime*>(val));
}

void* QVariant_NewQVariant21(void* val)
{
	return new QVariant(*static_cast<QDateTime*>(val));
}

void* QVariant_NewQVariant22(void* val)
{
	return new QVariant(*static_cast<QList<QVariant>*>(val));
}

void* QVariant_NewQVariant23(void* val)
{
	return new QVariant(*static_cast<QMap<QString, QVariant>*>(val));
}

void* QVariant_NewQVariant24(void* val)
{
	return new QVariant(*static_cast<QHash<QString, QVariant>*>(val));
}

void* QVariant_NewQVariant25(void* val)
{
	return new QVariant(*static_cast<QSize*>(val));
}

void* QVariant_NewQVariant26(void* val)
{
	return new QVariant(*static_cast<QSizeF*>(val));
}

void* QVariant_NewQVariant27(void* val)
{
	return new QVariant(*static_cast<QPoint*>(val));
}

void* QVariant_NewQVariant28(void* val)
{
	return new QVariant(*static_cast<QPointF*>(val));
}

void* QVariant_NewQVariant29(void* val)
{
	return new QVariant(*static_cast<QLine*>(val));
}

void* QVariant_NewQVariant30(void* val)
{
	return new QVariant(*static_cast<QLineF*>(val));
}

void* QVariant_NewQVariant31(void* val)
{
	return new QVariant(*static_cast<QRect*>(val));
}

void* QVariant_NewQVariant32(void* val)
{
	return new QVariant(*static_cast<QRectF*>(val));
}

void* QVariant_NewQVariant33(void* l)
{
	return new QVariant(*static_cast<QLocale*>(l));
}

void* QVariant_NewQVariant34(void* regExp)
{
	return new QVariant(*static_cast<QRegExp*>(regExp));
}

void* QVariant_NewQVariant35(void* re)
{
	return new QVariant(*static_cast<QRegularExpression*>(re));
}

void* QVariant_NewQVariant36(void* val)
{
	return new QVariant(*static_cast<QUrl*>(val));
}

void* QVariant_NewQVariant37(void* val)
{
	return new QVariant(*static_cast<QEasingCurve*>(val));
}

void* QVariant_NewQVariant38(void* val)
{
	return new QVariant(*static_cast<QUuid*>(val));
}

void* QVariant_NewQVariant39(void* val)
{
	return new QVariant(*static_cast<QJsonValue*>(val));
}

void* QVariant_NewQVariant40(void* val)
{
	return new QVariant(*static_cast<QJsonObject*>(val));
}

void* QVariant_NewQVariant41(void* val)
{
	return new QVariant(*static_cast<QJsonArray*>(val));
}

void* QVariant_NewQVariant42(void* val)
{
	return new QVariant(*static_cast<QJsonDocument*>(val));
}

void* QVariant_NewQVariant43(void* val)
{
	return new QVariant(*static_cast<QModelIndex*>(val));
}

void* QVariant_NewQVariant44(void* val)
{
	return new QVariant(*static_cast<QPersistentModelIndex*>(val));
}

void* QVariant_NewQVariant45(void* other)
{
	return new QVariant(*static_cast<QVariant*>(other));
}

char QVariant_ToBool(void* ptr)
{
	return static_cast<QVariant*>(ptr)->toBool();
}

double QVariant_ToDouble(void* ptr, char* ok)
{
	return static_cast<QVariant*>(ptr)->toDouble(reinterpret_cast<bool*>(ok));
}

struct QtCore_PackedList QVariant_ToHash(void* ptr)
{
	return ({ QHash<QString, QVariant>* tmpValue00701e = new QHash<QString, QVariant>(static_cast<QVariant*>(ptr)->toHash()); QtCore_PackedList { tmpValue00701e, tmpValue00701e->size() }; });
}

int QVariant_ToInt(void* ptr, char* ok)
{
	return static_cast<QVariant*>(ptr)->toInt(reinterpret_cast<bool*>(ok));
}

struct QtCore_PackedList QVariant_ToList(void* ptr)
{
	return ({ QList<QVariant>* tmpValue8f6950 = new QList<QVariant>(static_cast<QVariant*>(ptr)->toList()); QtCore_PackedList { tmpValue8f6950, tmpValue8f6950->size() }; });
}

long long QVariant_ToLongLong(void* ptr, char* ok)
{
	return static_cast<QVariant*>(ptr)->toLongLong(reinterpret_cast<bool*>(ok));
}

struct QtCore_PackedList QVariant_ToMap(void* ptr)
{
	return ({ QMap<QString, QVariant>* tmpValue1e0d76 = new QMap<QString, QVariant>(static_cast<QVariant*>(ptr)->toMap()); QtCore_PackedList { tmpValue1e0d76, tmpValue1e0d76->size() }; });
}

struct QtCore_PackedString QVariant_ToString(void* ptr)
{
	return ({ QByteArray tf9e1e4 = static_cast<QVariant*>(ptr)->toString().toUtf8(); QtCore_PackedString { const_cast<char*>(tf9e1e4.prepend("WHITESPACE").constData()+10), tf9e1e4.size()-10 }; });
}

struct QtCore_PackedString QVariant_ToStringList(void* ptr)
{
	return ({ QByteArray tf99cb6 = static_cast<QVariant*>(ptr)->toStringList().join("¡¦!").toUtf8(); QtCore_PackedString { const_cast<char*>(tf99cb6.prepend("WHITESPACE").constData()+10), tf99cb6.size()-10 }; });
}

unsigned int QVariant_ToUInt(void* ptr, char* ok)
{
	return static_cast<QVariant*>(ptr)->toUInt(reinterpret_cast<bool*>(ok));
}

unsigned long long QVariant_ToULongLong(void* ptr, char* ok)
{
	return static_cast<QVariant*>(ptr)->toULongLong(reinterpret_cast<bool*>(ok));
}

long long QVariant_Type(void* ptr)
{
	return static_cast<QVariant*>(ptr)->type();
}

void QVariant_DestroyQVariant(void* ptr)
{
	static_cast<QVariant*>(ptr)->~QVariant();
}

void* QVariant___QVariant_val_atList22(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVariant___QVariant_val_setList22(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QVariant___QVariant_val_newList22(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QVariant___QVariant_val_atList23(void* ptr, struct QtCore_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QVariant___QVariant_val_setList23(void* ptr, struct QtCore_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QVariant___QVariant_val_newList23(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtCore_PackedList QVariant___QVariant_val_keyList23(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

void* QVariant___QVariant_val_atList24(void* ptr, struct QtCore_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QHash<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QHash<QString, QVariant>*>(ptr)->size()-1) { static_cast<QHash<QString, QVariant>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QVariant___QVariant_val_setList24(void* ptr, struct QtCore_PackedString key, void* i)
{
	static_cast<QHash<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QVariant___QVariant_val_newList24(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<QString, QVariant>();
}

struct QtCore_PackedList QVariant___QVariant_val_keyList24(void* ptr)
{
	return ({ QList<QString>* tmpValuef43bc5 = new QList<QString>(static_cast<QHash<QString, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValuef43bc5, tmpValuef43bc5->size() }; });
}

void* QVariant___toHash_atList(void* ptr, struct QtCore_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QHash<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QHash<QString, QVariant>*>(ptr)->size()-1) { static_cast<QHash<QString, QVariant>*>(ptr)->~QHash(); free(ptr); }; tmp; }));
}

void QVariant___toHash_setList(void* ptr, struct QtCore_PackedString key, void* i)
{
	static_cast<QHash<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QVariant___toHash_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QHash<QString, QVariant>();
}

struct QtCore_PackedList QVariant___toHash_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValuef43bc5 = new QList<QString>(static_cast<QHash<QString, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValuef43bc5, tmpValuef43bc5->size() }; });
}

void* QVariant___toList_atList(void* ptr, int i)
{
	return new QVariant(({QVariant tmp = static_cast<QList<QVariant>*>(ptr)->at(i); if (i == static_cast<QList<QVariant>*>(ptr)->size()-1) { static_cast<QList<QVariant>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QVariant___toList_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QVariant___toList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QVariant>();
}

void* QVariant___toMap_atList(void* ptr, struct QtCore_PackedString v, int i)
{
	return new QVariant(({ QVariant tmp = static_cast<QMap<QString, QVariant>*>(ptr)->value(QString::fromUtf8(v.data, v.len)); if (i == static_cast<QMap<QString, QVariant>*>(ptr)->size()-1) { static_cast<QMap<QString, QVariant>*>(ptr)->~QMap(); free(ptr); }; tmp; }));
}

void QVariant___toMap_setList(void* ptr, struct QtCore_PackedString key, void* i)
{
	static_cast<QMap<QString, QVariant>*>(ptr)->insert(QString::fromUtf8(key.data, key.len), *static_cast<QVariant*>(i));
}

void* QVariant___toMap_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QMap<QString, QVariant>();
}

struct QtCore_PackedList QVariant___toMap_keyList(void* ptr)
{
	return ({ QList<QString>* tmpValue1ab909 = new QList<QString>(static_cast<QMap<QString, QVariant>*>(ptr)->keys()); QtCore_PackedList { tmpValue1ab909, tmpValue1ab909->size() }; });
}

struct QtCore_PackedString QVariant_____QVariant_val_keyList_atList23(void* ptr, int i)
{
	return ({ QByteArray t94aa5e = ({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8(); QtCore_PackedString { const_cast<char*>(t94aa5e.prepend("WHITESPACE").constData()+10), t94aa5e.size()-10 }; });
}

void QVariant_____QVariant_val_keyList_setList23(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QVariant_____QVariant_val_keyList_newList23(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtCore_PackedString QVariant_____QVariant_val_keyList_atList24(void* ptr, int i)
{
	return ({ QByteArray t94aa5e = ({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8(); QtCore_PackedString { const_cast<char*>(t94aa5e.prepend("WHITESPACE").constData()+10), t94aa5e.size()-10 }; });
}

void QVariant_____QVariant_val_keyList_setList24(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QVariant_____QVariant_val_keyList_newList24(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtCore_PackedString QVariant_____toHash_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t94aa5e = ({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8(); QtCore_PackedString { const_cast<char*>(t94aa5e.prepend("WHITESPACE").constData()+10), t94aa5e.size()-10 }; });
}

void QVariant_____toHash_keyList_setList(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QVariant_____toHash_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

struct QtCore_PackedString QVariant_____toMap_keyList_atList(void* ptr, int i)
{
	return ({ QByteArray t94aa5e = ({QString tmp = static_cast<QList<QString>*>(ptr)->at(i); if (i == static_cast<QList<QString>*>(ptr)->size()-1) { static_cast<QList<QString>*>(ptr)->~QList(); free(ptr); }; tmp; }).toUtf8(); QtCore_PackedString { const_cast<char*>(t94aa5e.prepend("WHITESPACE").constData()+10), t94aa5e.size()-10 }; });
}

void QVariant_____toMap_keyList_setList(void* ptr, struct QtCore_PackedString i)
{
	static_cast<QList<QString>*>(ptr)->append(QString::fromUtf8(i.data, i.len));
}

void* QVariant_____toMap_keyList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QString>();
}

int Qt_LastGestureType_Type()
{
	return Qt::LastGestureType;
}

