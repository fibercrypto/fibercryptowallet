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
struct qt_meta_stringdata_WalletModel445aa6_t {
    QByteArrayData data[21];
    char stringdata0[216];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_WalletModel445aa6_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_WalletModel445aa6_t qt_meta_stringdata_WalletModel445aa6 = {
    {
QT_MOC_LITERAL(0, 0, 17), // "WalletModel445aa6"
QT_MOC_LITERAL(1, 18, 12), // "rolesChanged"
QT_MOC_LITERAL(2, 31, 0), // ""
QT_MOC_LITERAL(3, 32, 10), // "type378cdd"
QT_MOC_LITERAL(4, 43, 5), // "roles"
QT_MOC_LITERAL(5, 49, 14), // "walletsChanged"
QT_MOC_LITERAL(6, 64, 21), // "QList<QWallet445aa6*>"
QT_MOC_LITERAL(7, 86, 7), // "wallets"
QT_MOC_LITERAL(8, 94, 12), // "countChanged"
QT_MOC_LITERAL(9, 107, 5), // "count"
QT_MOC_LITERAL(10, 113, 9), // "addWallet"
QT_MOC_LITERAL(11, 123, 14), // "QWallet445aa6*"
QT_MOC_LITERAL(12, 138, 2), // "v0"
QT_MOC_LITERAL(13, 141, 10), // "editWallet"
QT_MOC_LITERAL(14, 152, 3), // "row"
QT_MOC_LITERAL(15, 156, 4), // "name"
QT_MOC_LITERAL(16, 161, 17), // "encryptionEnabled"
QT_MOC_LITERAL(17, 179, 3), // "sky"
QT_MOC_LITERAL(18, 183, 9), // "coinHours"
QT_MOC_LITERAL(19, 193, 12), // "removeWallet"
QT_MOC_LITERAL(20, 206, 9) // "loadModel"

    },
    "WalletModel445aa6\0rolesChanged\0\0"
    "type378cdd\0roles\0walletsChanged\0"
    "QList<QWallet445aa6*>\0wallets\0"
    "countChanged\0count\0addWallet\0"
    "QWallet445aa6*\0v0\0editWallet\0row\0name\0"
    "encryptionEnabled\0sky\0coinHours\0"
    "removeWallet\0loadModel"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_WalletModel445aa6[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       7,   14, // methods
       3,   78, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       3,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   49,    2, 0x06 /* Public */,
       5,    1,   52,    2, 0x06 /* Public */,
       8,    1,   55,    2, 0x06 /* Public */,

 // slots: name, argc, parameters, tag, flags
      10,    1,   58,    2, 0x0a /* Public */,
      13,    5,   61,    2, 0x0a /* Public */,
      19,    1,   72,    2, 0x0a /* Public */,
      20,    1,   75,    2, 0x0a /* Public */,

 // signals: parameters
    QMetaType::Void, 0x80000000 | 3,    4,
    QMetaType::Void, 0x80000000 | 6,    7,
    QMetaType::Void, QMetaType::Int,    9,

 // slots: parameters
    QMetaType::Void, 0x80000000 | 11,   12,
    QMetaType::Void, QMetaType::Int, QMetaType::QString, QMetaType::Bool, QMetaType::ULongLong, QMetaType::ULongLong,   14,   15,   16,   17,   18,
    QMetaType::Void, QMetaType::Int,   14,
    QMetaType::Void, 0x80000000 | 6,   12,

 // properties: name, type, flags
       4, 0x80000000 | 3, 0x0049510b,
       7, 0x80000000 | 6, 0x0049510b,
       9, QMetaType::Int, 0x00495103,

 // properties: notify_signal_id
       0,
       1,
       2,

       0        // eod
};

void WalletModel445aa6::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<WalletModel445aa6 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->rolesChanged((*reinterpret_cast< type378cdd(*)>(_a[1]))); break;
        case 1: _t->walletsChanged((*reinterpret_cast< QList<QWallet445aa6*>(*)>(_a[1]))); break;
        case 2: _t->countChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 3: _t->addWallet((*reinterpret_cast< QWallet445aa6*(*)>(_a[1]))); break;
        case 4: _t->editWallet((*reinterpret_cast< qint32(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< bool(*)>(_a[3])),(*reinterpret_cast< quint64(*)>(_a[4])),(*reinterpret_cast< quint64(*)>(_a[5]))); break;
        case 5: _t->removeWallet((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 6: _t->loadModel((*reinterpret_cast< QList<QWallet445aa6*>(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<QWallet445aa6*> >(); break;
            }
            break;
        case 3:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QWallet445aa6* >(); break;
            }
            break;
        case 6:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<QWallet445aa6*> >(); break;
            }
            break;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (WalletModel445aa6::*)(type378cdd );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&WalletModel445aa6::rolesChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (WalletModel445aa6::*)(QList<QWallet445aa6*> );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&WalletModel445aa6::walletsChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (WalletModel445aa6::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&WalletModel445aa6::countChanged)) {
                *result = 2;
                return;
            }
        }
    } else if (_c == QMetaObject::RegisterPropertyMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<QWallet445aa6*> >(); break;
        }
    }

#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<WalletModel445aa6 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< type378cdd*>(_v) = _t->roles(); break;
        case 1: *reinterpret_cast< QList<QWallet445aa6*>*>(_v) = _t->wallets(); break;
        case 2: *reinterpret_cast< qint32*>(_v) = _t->count(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<WalletModel445aa6 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRoles(*reinterpret_cast< type378cdd*>(_v)); break;
        case 1: _t->setWallets(*reinterpret_cast< QList<QWallet445aa6*>*>(_v)); break;
        case 2: _t->setCount(*reinterpret_cast< qint32*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject WalletModel445aa6::staticMetaObject = { {
    &QAbstractListModel::staticMetaObject,
    qt_meta_stringdata_WalletModel445aa6.data,
    qt_meta_data_WalletModel445aa6,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *WalletModel445aa6::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *WalletModel445aa6::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_WalletModel445aa6.stringdata0))
        return static_cast<void*>(this);
    return QAbstractListModel::qt_metacast(_clname);
}

int WalletModel445aa6::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QAbstractListModel::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 7)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 7;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 7)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 7;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 3;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void WalletModel445aa6::rolesChanged(type378cdd _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void WalletModel445aa6::walletsChanged(QList<QWallet445aa6*> _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void WalletModel445aa6::countChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}
struct qt_meta_stringdata_AddressesModel445aa6_t {
    QByteArrayData data[17];
    char stringdata0[177];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_AddressesModel445aa6_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_AddressesModel445aa6_t qt_meta_stringdata_AddressesModel445aa6 = {
    {
QT_MOC_LITERAL(0, 0, 20), // "AddressesModel445aa6"
QT_MOC_LITERAL(1, 21, 12), // "rolesChanged"
QT_MOC_LITERAL(2, 34, 0), // ""
QT_MOC_LITERAL(3, 35, 10), // "type378cdd"
QT_MOC_LITERAL(4, 46, 5), // "roles"
QT_MOC_LITERAL(5, 52, 16), // "addressesChanged"
QT_MOC_LITERAL(6, 69, 22), // "QList<QAddress445aa6*>"
QT_MOC_LITERAL(7, 92, 9), // "addresses"
QT_MOC_LITERAL(8, 102, 10), // "addAddress"
QT_MOC_LITERAL(9, 113, 15), // "QAddress445aa6*"
QT_MOC_LITERAL(10, 129, 2), // "v0"
QT_MOC_LITERAL(11, 132, 13), // "removeAddress"
QT_MOC_LITERAL(12, 146, 11), // "editAddress"
QT_MOC_LITERAL(13, 158, 2), // "v1"
QT_MOC_LITERAL(14, 161, 2), // "v2"
QT_MOC_LITERAL(15, 164, 2), // "v3"
QT_MOC_LITERAL(16, 167, 9) // "loadModel"

    },
    "AddressesModel445aa6\0rolesChanged\0\0"
    "type378cdd\0roles\0addressesChanged\0"
    "QList<QAddress445aa6*>\0addresses\0"
    "addAddress\0QAddress445aa6*\0v0\0"
    "removeAddress\0editAddress\0v1\0v2\0v3\0"
    "loadModel"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_AddressesModel445aa6[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       6,   14, // methods
       2,   68, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       2,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   44,    2, 0x06 /* Public */,
       5,    1,   47,    2, 0x06 /* Public */,

 // slots: name, argc, parameters, tag, flags
       8,    1,   50,    2, 0x0a /* Public */,
      11,    1,   53,    2, 0x0a /* Public */,
      12,    4,   56,    2, 0x0a /* Public */,
      16,    1,   65,    2, 0x0a /* Public */,

 // signals: parameters
    QMetaType::Void, 0x80000000 | 3,    4,
    QMetaType::Void, 0x80000000 | 6,    7,

 // slots: parameters
    QMetaType::Void, 0x80000000 | 9,   10,
    QMetaType::Void, QMetaType::Int,   10,
    QMetaType::Void, QMetaType::Int, QMetaType::QString, QMetaType::ULongLong, QMetaType::ULongLong,   10,   13,   14,   15,
    QMetaType::Void, 0x80000000 | 6,   10,

 // properties: name, type, flags
       4, 0x80000000 | 3, 0x0049510b,
       7, 0x80000000 | 6, 0x0049510b,

 // properties: notify_signal_id
       0,
       1,

       0        // eod
};

void AddressesModel445aa6::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<AddressesModel445aa6 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->rolesChanged((*reinterpret_cast< type378cdd(*)>(_a[1]))); break;
        case 1: _t->addressesChanged((*reinterpret_cast< QList<QAddress445aa6*>(*)>(_a[1]))); break;
        case 2: _t->addAddress((*reinterpret_cast< QAddress445aa6*(*)>(_a[1]))); break;
        case 3: _t->removeAddress((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 4: _t->editAddress((*reinterpret_cast< qint32(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< quint64(*)>(_a[3])),(*reinterpret_cast< quint64(*)>(_a[4]))); break;
        case 5: _t->loadModel((*reinterpret_cast< QList<QAddress445aa6*>(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<QAddress445aa6*> >(); break;
            }
            break;
        case 2:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QAddress445aa6* >(); break;
            }
            break;
        case 5:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<QAddress445aa6*> >(); break;
            }
            break;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (AddressesModel445aa6::*)(type378cdd );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressesModel445aa6::rolesChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (AddressesModel445aa6::*)(QList<QAddress445aa6*> );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressesModel445aa6::addressesChanged)) {
                *result = 1;
                return;
            }
        }
    } else if (_c == QMetaObject::RegisterPropertyMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<QAddress445aa6*> >(); break;
        }
    }

#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<AddressesModel445aa6 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< type378cdd*>(_v) = _t->roles(); break;
        case 1: *reinterpret_cast< QList<QAddress445aa6*>*>(_v) = _t->addresses(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<AddressesModel445aa6 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRoles(*reinterpret_cast< type378cdd*>(_v)); break;
        case 1: _t->setAddresses(*reinterpret_cast< QList<QAddress445aa6*>*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject AddressesModel445aa6::staticMetaObject = { {
    &QAbstractListModel::staticMetaObject,
    qt_meta_stringdata_AddressesModel445aa6.data,
    qt_meta_data_AddressesModel445aa6,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *AddressesModel445aa6::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *AddressesModel445aa6::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_AddressesModel445aa6.stringdata0))
        return static_cast<void*>(this);
    return QAbstractListModel::qt_metacast(_clname);
}

int AddressesModel445aa6::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QAbstractListModel::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 6)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 6;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 6)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 6;
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
void AddressesModel445aa6::rolesChanged(type378cdd _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void AddressesModel445aa6::addressesChanged(QList<QAddress445aa6*> _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}
struct qt_meta_stringdata_QAddress445aa6_t {
    QByteArrayData data[8];
    char stringdata0[109];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_QAddress445aa6_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_QAddress445aa6_t qt_meta_stringdata_QAddress445aa6 = {
    {
QT_MOC_LITERAL(0, 0, 14), // "QAddress445aa6"
QT_MOC_LITERAL(1, 15, 14), // "addressChanged"
QT_MOC_LITERAL(2, 30, 0), // ""
QT_MOC_LITERAL(3, 31, 7), // "address"
QT_MOC_LITERAL(4, 39, 17), // "addressSkyChanged"
QT_MOC_LITERAL(5, 57, 10), // "addressSky"
QT_MOC_LITERAL(6, 68, 23), // "addressCoinHoursChanged"
QT_MOC_LITERAL(7, 92, 16) // "addressCoinHours"

    },
    "QAddress445aa6\0addressChanged\0\0address\0"
    "addressSkyChanged\0addressSky\0"
    "addressCoinHoursChanged\0addressCoinHours"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_QAddress445aa6[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       3,   14, // methods
       3,   38, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       3,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   29,    2, 0x06 /* Public */,
       4,    1,   32,    2, 0x06 /* Public */,
       6,    1,   35,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::ULongLong,    5,
    QMetaType::Void, QMetaType::ULongLong,    7,

 // properties: name, type, flags
       3, QMetaType::QString, 0x00495103,
       5, QMetaType::ULongLong, 0x00495103,
       7, QMetaType::ULongLong, 0x00495103,

 // properties: notify_signal_id
       0,
       1,
       2,

       0        // eod
};

void QAddress445aa6::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<QAddress445aa6 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->addressChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 1: _t->addressSkyChanged((*reinterpret_cast< quint64(*)>(_a[1]))); break;
        case 2: _t->addressCoinHoursChanged((*reinterpret_cast< quint64(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (QAddress445aa6::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QAddress445aa6::addressChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (QAddress445aa6::*)(quint64 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QAddress445aa6::addressSkyChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (QAddress445aa6::*)(quint64 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QAddress445aa6::addressCoinHoursChanged)) {
                *result = 2;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<QAddress445aa6 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QString*>(_v) = _t->address(); break;
        case 1: *reinterpret_cast< quint64*>(_v) = _t->addressSky(); break;
        case 2: *reinterpret_cast< quint64*>(_v) = _t->addressCoinHours(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<QAddress445aa6 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setAddress(*reinterpret_cast< QString*>(_v)); break;
        case 1: _t->setAddressSky(*reinterpret_cast< quint64*>(_v)); break;
        case 2: _t->setAddressCoinHours(*reinterpret_cast< quint64*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject QAddress445aa6::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_QAddress445aa6.data,
    qt_meta_data_QAddress445aa6,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *QAddress445aa6::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *QAddress445aa6::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_QAddress445aa6.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int QAddress445aa6::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 3)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 3;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 3)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 3;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 3;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 3;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void QAddress445aa6::addressChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void QAddress445aa6::addressSkyChanged(quint64 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void QAddress445aa6::addressCoinHoursChanged(quint64 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}
struct qt_meta_stringdata_QWallet445aa6_t {
    QByteArrayData data[12];
    char stringdata0[142];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_QWallet445aa6_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_QWallet445aa6_t qt_meta_stringdata_QWallet445aa6 = {
    {
QT_MOC_LITERAL(0, 0, 13), // "QWallet445aa6"
QT_MOC_LITERAL(1, 14, 11), // "nameChanged"
QT_MOC_LITERAL(2, 26, 0), // ""
QT_MOC_LITERAL(3, 27, 4), // "name"
QT_MOC_LITERAL(4, 32, 24), // "encryptionEnabledChanged"
QT_MOC_LITERAL(5, 57, 17), // "encryptionEnabled"
QT_MOC_LITERAL(6, 75, 10), // "skyChanged"
QT_MOC_LITERAL(7, 86, 3), // "sky"
QT_MOC_LITERAL(8, 90, 16), // "coinHoursChanged"
QT_MOC_LITERAL(9, 107, 9), // "coinHours"
QT_MOC_LITERAL(10, 117, 15), // "fileNameChanged"
QT_MOC_LITERAL(11, 133, 8) // "fileName"

    },
    "QWallet445aa6\0nameChanged\0\0name\0"
    "encryptionEnabledChanged\0encryptionEnabled\0"
    "skyChanged\0sky\0coinHoursChanged\0"
    "coinHours\0fileNameChanged\0fileName"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_QWallet445aa6[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       5,   14, // methods
       5,   54, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       5,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   39,    2, 0x06 /* Public */,
       4,    1,   42,    2, 0x06 /* Public */,
       6,    1,   45,    2, 0x06 /* Public */,
       8,    1,   48,    2, 0x06 /* Public */,
      10,    1,   51,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::Int,    5,
    QMetaType::Void, QMetaType::ULongLong,    7,
    QMetaType::Void, QMetaType::ULongLong,    9,
    QMetaType::Void, QMetaType::QString,   11,

 // properties: name, type, flags
       3, QMetaType::QString, 0x00495103,
       5, QMetaType::Int, 0x00495103,
       7, QMetaType::ULongLong, 0x00495103,
       9, QMetaType::ULongLong, 0x00495103,
      11, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
       0,
       1,
       2,
       3,
       4,

       0        // eod
};

void QWallet445aa6::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<QWallet445aa6 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->nameChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 1: _t->encryptionEnabledChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 2: _t->skyChanged((*reinterpret_cast< quint64(*)>(_a[1]))); break;
        case 3: _t->coinHoursChanged((*reinterpret_cast< quint64(*)>(_a[1]))); break;
        case 4: _t->fileNameChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (QWallet445aa6::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QWallet445aa6::nameChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (QWallet445aa6::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QWallet445aa6::encryptionEnabledChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (QWallet445aa6::*)(quint64 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QWallet445aa6::skyChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (QWallet445aa6::*)(quint64 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QWallet445aa6::coinHoursChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (QWallet445aa6::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QWallet445aa6::fileNameChanged)) {
                *result = 4;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<QWallet445aa6 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QString*>(_v) = _t->name(); break;
        case 1: *reinterpret_cast< qint32*>(_v) = _t->encryptionEnabled(); break;
        case 2: *reinterpret_cast< quint64*>(_v) = _t->sky(); break;
        case 3: *reinterpret_cast< quint64*>(_v) = _t->coinHours(); break;
        case 4: *reinterpret_cast< QString*>(_v) = _t->fileName(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<QWallet445aa6 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setName(*reinterpret_cast< QString*>(_v)); break;
        case 1: _t->setEncryptionEnabled(*reinterpret_cast< qint32*>(_v)); break;
        case 2: _t->setSky(*reinterpret_cast< quint64*>(_v)); break;
        case 3: _t->setCoinHours(*reinterpret_cast< quint64*>(_v)); break;
        case 4: _t->setFileName(*reinterpret_cast< QString*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject QWallet445aa6::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_QWallet445aa6.data,
    qt_meta_data_QWallet445aa6,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *QWallet445aa6::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *QWallet445aa6::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_QWallet445aa6.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int QWallet445aa6::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
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
        _id -= 5;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 5;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 5;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 5;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 5;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 5;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void QWallet445aa6::nameChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void QWallet445aa6::encryptionEnabledChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void QWallet445aa6::skyChanged(quint64 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void QWallet445aa6::coinHoursChanged(quint64 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void QWallet445aa6::fileNameChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
