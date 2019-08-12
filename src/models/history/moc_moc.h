/****************************************************************************
** Meta object code from reading C++ file 'moc.cpp'
**
** Created by: The Qt Meta Object Compiler version 67 (Qt 5.13.0)
**
** WARNING! All changes made in this file will be lost!
*****************************************************************************/

#include <memory>
#include <QtCore/qbytearray.h>
#include <QtCore/qmetatype.h>
#include <QtCore/QList>
#if !defined(Q_MOC_OUTPUT_REVISION)
#error "The header file 'moc.cpp' doesn't include <QObject>."
#elif Q_MOC_OUTPUT_REVISION != 67
#error "This file was generated using the moc from 5.13.0. It"
#error "cannot be used with the include files from this version of Qt."
#error "(The moc has changed too much.)"
#endif

QT_BEGIN_MOC_NAMESPACE
QT_WARNING_PUSH
QT_WARNING_DISABLE_DEPRECATED
struct qt_meta_stringdata_TransactionList554044_t {
    QByteArrayData data[15];
    char stringdata0[225];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_TransactionList554044_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_TransactionList554044_t qt_meta_stringdata_TransactionList554044 = {
    {
QT_MOC_LITERAL(0, 0, 21), // "TransactionList554044"
QT_MOC_LITERAL(1, 22, 14), // "addTransaction"
QT_MOC_LITERAL(2, 37, 0), // ""
QT_MOC_LITERAL(3, 38, 25), // "TransactionDetailsecff1c*"
QT_MOC_LITERAL(4, 64, 11), // "transaction"
QT_MOC_LITERAL(5, 76, 17), // "removeTransaction"
QT_MOC_LITERAL(6, 94, 5), // "index"
QT_MOC_LITERAL(7, 100, 12), // "rolesChanged"
QT_MOC_LITERAL(8, 113, 10), // "type378cdd"
QT_MOC_LITERAL(9, 124, 5), // "roles"
QT_MOC_LITERAL(10, 130, 19), // "transactionsChanged"
QT_MOC_LITERAL(11, 150, 32), // "QList<TransactionDetailsecff1c*>"
QT_MOC_LITERAL(12, 183, 12), // "transactions"
QT_MOC_LITERAL(13, 196, 23), // "addMultipleTransactions"
QT_MOC_LITERAL(14, 220, 4) // "txns"

    },
    "TransactionList554044\0addTransaction\0"
    "\0TransactionDetailsecff1c*\0transaction\0"
    "removeTransaction\0index\0rolesChanged\0"
    "type378cdd\0roles\0transactionsChanged\0"
    "QList<TransactionDetailsecff1c*>\0"
    "transactions\0addMultipleTransactions\0"
    "txns"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_TransactionList554044[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       5,   14, // methods
       2,   54, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       4,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   39,    2, 0x06 /* Public */,
       5,    1,   42,    2, 0x06 /* Public */,
       7,    1,   45,    2, 0x06 /* Public */,
      10,    1,   48,    2, 0x06 /* Public */,

 // slots: name, argc, parameters, tag, flags
      13,    1,   51,    2, 0x0a /* Public */,

 // signals: parameters
    QMetaType::Void, 0x80000000 | 3,    4,
    QMetaType::Void, QMetaType::Int,    6,
    QMetaType::Void, 0x80000000 | 8,    9,
    QMetaType::Void, 0x80000000 | 11,   12,

 // slots: parameters
    QMetaType::Void, 0x80000000 | 11,   14,

 // properties: name, type, flags
       9, 0x80000000 | 8, 0x0049510b,
      12, 0x80000000 | 11, 0x0049510b,

 // properties: notify_signal_id
       2,
       3,

       0        // eod
};

void TransactionList554044::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<TransactionList554044 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->addTransaction((*reinterpret_cast< TransactionDetailsecff1c*(*)>(_a[1]))); break;
        case 1: _t->removeTransaction((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 2: _t->rolesChanged((*reinterpret_cast< type378cdd(*)>(_a[1]))); break;
        case 3: _t->transactionsChanged((*reinterpret_cast< QList<TransactionDetailsecff1c*>(*)>(_a[1]))); break;
        case 4: _t->addMultipleTransactions((*reinterpret_cast< QList<TransactionDetailsecff1c*>(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (TransactionList554044::*)(TransactionDetailsecff1c * );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionList554044::addTransaction)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (TransactionList554044::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionList554044::removeTransaction)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (TransactionList554044::*)(type378cdd );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionList554044::rolesChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (TransactionList554044::*)(QList<TransactionDetailsecff1c*> );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionList554044::transactionsChanged)) {
                *result = 3;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<TransactionList554044 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< type378cdd*>(_v) = _t->roles(); break;
        case 1: *reinterpret_cast< QList<TransactionDetailsecff1c*>*>(_v) = _t->transactions(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<TransactionList554044 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRoles(*reinterpret_cast< type378cdd*>(_v)); break;
        case 1: _t->setTransactions(*reinterpret_cast< QList<TransactionDetailsecff1c*>*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject TransactionList554044::staticMetaObject = { {
    &QAbstractListModel::staticMetaObject,
    qt_meta_stringdata_TransactionList554044.data,
    qt_meta_data_TransactionList554044,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *TransactionList554044::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *TransactionList554044::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_TransactionList554044.stringdata0))
        return static_cast<void*>(this);
    return QAbstractListModel::qt_metacast(_clname);
}

int TransactionList554044::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QAbstractListModel::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 5)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 5;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 5)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 5;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 2;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 2;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void TransactionList554044::addTransaction(TransactionDetailsecff1c * _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void TransactionList554044::removeTransaction(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void TransactionList554044::rolesChanged(type378cdd _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void TransactionList554044::transactionsChanged(QList<TransactionDetailsecff1c*> _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}
struct qt_meta_stringdata_HistoryManager554044_t {
    QByteArrayData data[6];
    char stringdata0[106];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_HistoryManager554044_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_HistoryManager554044_t qt_meta_stringdata_HistoryManager554044 = {
    {
QT_MOC_LITERAL(0, 0, 20), // "HistoryManager554044"
QT_MOC_LITERAL(1, 21, 22), // "loadHistoryWithFilters"
QT_MOC_LITERAL(2, 44, 32), // "QList<TransactionDetailsecff1c*>"
QT_MOC_LITERAL(3, 77, 0), // ""
QT_MOC_LITERAL(4, 78, 15), // "filterAddresses"
QT_MOC_LITERAL(5, 94, 11) // "loadHistory"

    },
    "HistoryManager554044\0loadHistoryWithFilters\0"
    "QList<TransactionDetailsecff1c*>\0\0"
    "filterAddresses\0loadHistory"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_HistoryManager554044[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       2,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       0,       // signalCount

 // slots: name, argc, parameters, tag, flags
       1,    1,   24,    3, 0x0a /* Public */,
       5,    0,   27,    3, 0x0a /* Public */,

 // slots: parameters
    0x80000000 | 2, QMetaType::QStringList,    4,
    0x80000000 | 2,

       0        // eod
};

void HistoryManager554044::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<HistoryManager554044 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: { QList<TransactionDetailsecff1c*> _r = _t->loadHistoryWithFilters((*reinterpret_cast< QStringList(*)>(_a[1])));
            if (_a[0]) *reinterpret_cast< QList<TransactionDetailsecff1c*>*>(_a[0]) = std::move(_r); }  break;
        case 1: { QList<TransactionDetailsecff1c*> _r = _t->loadHistory();
            if (_a[0]) *reinterpret_cast< QList<TransactionDetailsecff1c*>*>(_a[0]) = std::move(_r); }  break;
        default: ;
        }
    }
}

QT_INIT_METAOBJECT const QMetaObject HistoryManager554044::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_HistoryManager554044.data,
    qt_meta_data_HistoryManager554044,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *HistoryManager554044::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *HistoryManager554044::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_HistoryManager554044.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int HistoryManager554044::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 2)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 2;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 2)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 2;
    }
    return _id;
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
