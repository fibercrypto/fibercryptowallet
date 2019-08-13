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
struct qt_meta_stringdata_AddressesModel539e18_t {
    QByteArrayData data[20];
    char stringdata0[199];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_AddressesModel539e18_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_AddressesModel539e18_t qt_meta_stringdata_AddressesModel539e18 = {
    {
QT_MOC_LITERAL(0, 0, 20), // "AddressesModel539e18"
QT_MOC_LITERAL(1, 21, 12), // "rolesChanged"
QT_MOC_LITERAL(2, 34, 0), // ""
QT_MOC_LITERAL(3, 35, 10), // "type378cdd"
QT_MOC_LITERAL(4, 46, 5), // "roles"
QT_MOC_LITERAL(5, 52, 16), // "addressesChanged"
QT_MOC_LITERAL(6, 69, 22), // "QList<QAddress539e18*>"
QT_MOC_LITERAL(7, 92, 9), // "addresses"
QT_MOC_LITERAL(8, 102, 12), // "countChanged"
QT_MOC_LITERAL(9, 115, 5), // "count"
QT_MOC_LITERAL(10, 121, 10), // "addAddress"
QT_MOC_LITERAL(11, 132, 15), // "QAddress539e18*"
QT_MOC_LITERAL(12, 148, 2), // "v0"
QT_MOC_LITERAL(13, 151, 13), // "removeAddress"
QT_MOC_LITERAL(14, 165, 11), // "editAddress"
QT_MOC_LITERAL(15, 177, 2), // "v1"
QT_MOC_LITERAL(16, 180, 2), // "v2"
QT_MOC_LITERAL(17, 183, 2), // "v3"
QT_MOC_LITERAL(18, 186, 2), // "v4"
QT_MOC_LITERAL(19, 189, 9) // "loadModel"

    },
    "AddressesModel539e18\0rolesChanged\0\0"
    "type378cdd\0roles\0addressesChanged\0"
    "QList<QAddress539e18*>\0addresses\0"
    "countChanged\0count\0addAddress\0"
    "QAddress539e18*\0v0\0removeAddress\0"
    "editAddress\0v1\0v2\0v3\0v4\0loadModel"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_AddressesModel539e18[] = {

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
      13,    1,   61,    2, 0x0a /* Public */,
      14,    5,   64,    2, 0x0a /* Public */,
      19,    1,   75,    2, 0x0a /* Public */,

 // signals: parameters
    QMetaType::Void, 0x80000000 | 3,    4,
    QMetaType::Void, 0x80000000 | 6,    7,
    QMetaType::Void, QMetaType::Int,    9,

 // slots: parameters
    QMetaType::Void, 0x80000000 | 11,   12,
    QMetaType::Void, QMetaType::Int,   12,
    QMetaType::Void, QMetaType::Int, QMetaType::QString, QMetaType::ULongLong, QMetaType::ULongLong, QMetaType::Int,   12,   15,   16,   17,   18,
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

void AddressesModel539e18::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<AddressesModel539e18 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->rolesChanged((*reinterpret_cast< type378cdd(*)>(_a[1]))); break;
        case 1: _t->addressesChanged((*reinterpret_cast< QList<QAddress539e18*>(*)>(_a[1]))); break;
        case 2: _t->countChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 3: _t->addAddress((*reinterpret_cast< QAddress539e18*(*)>(_a[1]))); break;
        case 4: _t->removeAddress((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 5: _t->editAddress((*reinterpret_cast< qint32(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< quint64(*)>(_a[3])),(*reinterpret_cast< quint64(*)>(_a[4])),(*reinterpret_cast< qint32(*)>(_a[5]))); break;
        case 6: _t->loadModel((*reinterpret_cast< QList<QAddress539e18*>(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<QAddress539e18*> >(); break;
            }
            break;
        case 3:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QAddress539e18* >(); break;
            }
            break;
        case 6:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<QAddress539e18*> >(); break;
            }
            break;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (AddressesModel539e18::*)(type378cdd );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressesModel539e18::rolesChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (AddressesModel539e18::*)(QList<QAddress539e18*> );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressesModel539e18::addressesChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (AddressesModel539e18::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressesModel539e18::countChanged)) {
                *result = 2;
                return;
            }
        }
    } else if (_c == QMetaObject::RegisterPropertyMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<QAddress539e18*> >(); break;
        }
    }

#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<AddressesModel539e18 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< type378cdd*>(_v) = _t->roles(); break;
        case 1: *reinterpret_cast< QList<QAddress539e18*>*>(_v) = _t->addresses(); break;
        case 2: *reinterpret_cast< qint32*>(_v) = _t->count(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<AddressesModel539e18 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRoles(*reinterpret_cast< type378cdd*>(_v)); break;
        case 1: _t->setAddresses(*reinterpret_cast< QList<QAddress539e18*>*>(_v)); break;
        case 2: _t->setCount(*reinterpret_cast< qint32*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject AddressesModel539e18::staticMetaObject = { {
    &QAbstractListModel::staticMetaObject,
    qt_meta_stringdata_AddressesModel539e18.data,
    qt_meta_data_AddressesModel539e18,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *AddressesModel539e18::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *AddressesModel539e18::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_AddressesModel539e18.stringdata0))
        return static_cast<void*>(this);
    return QAbstractListModel::qt_metacast(_clname);
}

int AddressesModel539e18::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
void AddressesModel539e18::rolesChanged(type378cdd _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void AddressesModel539e18::addressesChanged(QList<QAddress539e18*> _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void AddressesModel539e18::countChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}
struct qt_meta_stringdata_ModelManager539e18_t {
    QByteArrayData data[7];
    char stringdata0[99];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_ModelManager539e18_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_ModelManager539e18_t qt_meta_stringdata_ModelManager539e18 = {
    {
QT_MOC_LITERAL(0, 0, 18), // "ModelManager539e18"
QT_MOC_LITERAL(1, 19, 16), // "setWalletManager"
QT_MOC_LITERAL(2, 36, 0), // ""
QT_MOC_LITERAL(3, 37, 20), // "WalletManager539e18*"
QT_MOC_LITERAL(4, 58, 2), // "v0"
QT_MOC_LITERAL(5, 61, 15), // "getAddressModel"
QT_MOC_LITERAL(6, 77, 21) // "AddressesModel539e18*"

    },
    "ModelManager539e18\0setWalletManager\0"
    "\0WalletManager539e18*\0v0\0getAddressModel\0"
    "AddressesModel539e18*"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_ModelManager539e18[] = {

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
       1,    1,   24,    2, 0x0a /* Public */,
       5,    1,   27,    2, 0x0a /* Public */,

 // slots: parameters
    QMetaType::Void, 0x80000000 | 3,    4,
    0x80000000 | 6, QMetaType::QString,    4,

       0        // eod
};

void ModelManager539e18::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<ModelManager539e18 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->setWalletManager((*reinterpret_cast< WalletManager539e18*(*)>(_a[1]))); break;
        case 1: { AddressesModel539e18* _r = _t->getAddressModel((*reinterpret_cast< QString(*)>(_a[1])));
            if (_a[0]) *reinterpret_cast< AddressesModel539e18**>(_a[0]) = std::move(_r); }  break;
        default: ;
        }
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 0:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< WalletManager539e18* >(); break;
            }
            break;
        }
    }
}

QT_INIT_METAOBJECT const QMetaObject ModelManager539e18::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_ModelManager539e18.data,
    qt_meta_data_ModelManager539e18,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *ModelManager539e18::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *ModelManager539e18::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_ModelManager539e18.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int ModelManager539e18::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
            qt_static_metacall(this, _c, _id, _a);
        _id -= 2;
    }
    return _id;
}
struct qt_meta_stringdata_QAddress539e18_t {
    QByteArrayData data[10];
    char stringdata0[130];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_QAddress539e18_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_QAddress539e18_t qt_meta_stringdata_QAddress539e18 = {
    {
QT_MOC_LITERAL(0, 0, 14), // "QAddress539e18"
QT_MOC_LITERAL(1, 15, 14), // "addressChanged"
QT_MOC_LITERAL(2, 30, 0), // ""
QT_MOC_LITERAL(3, 31, 7), // "address"
QT_MOC_LITERAL(4, 39, 17), // "addressSkyChanged"
QT_MOC_LITERAL(5, 57, 10), // "addressSky"
QT_MOC_LITERAL(6, 68, 23), // "addressCoinHoursChanged"
QT_MOC_LITERAL(7, 92, 16), // "addressCoinHours"
QT_MOC_LITERAL(8, 109, 13), // "markedChanged"
QT_MOC_LITERAL(9, 123, 6) // "marked"

    },
    "QAddress539e18\0addressChanged\0\0address\0"
    "addressSkyChanged\0addressSky\0"
    "addressCoinHoursChanged\0addressCoinHours\0"
    "markedChanged\0marked"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_QAddress539e18[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       4,   14, // methods
       4,   46, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       4,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   34,    2, 0x06 /* Public */,
       4,    1,   37,    2, 0x06 /* Public */,
       6,    1,   40,    2, 0x06 /* Public */,
       8,    1,   43,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::QString,    3,
    QMetaType::Void, QMetaType::ULongLong,    5,
    QMetaType::Void, QMetaType::ULongLong,    7,
    QMetaType::Void, QMetaType::Int,    9,

 // properties: name, type, flags
       3, QMetaType::QString, 0x00495103,
       5, QMetaType::ULongLong, 0x00495103,
       7, QMetaType::ULongLong, 0x00495103,
       9, QMetaType::Int, 0x00495103,

 // properties: notify_signal_id
       0,
       1,
       2,
       3,

       0        // eod
};

void QAddress539e18::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<QAddress539e18 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->addressChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 1: _t->addressSkyChanged((*reinterpret_cast< quint64(*)>(_a[1]))); break;
        case 2: _t->addressCoinHoursChanged((*reinterpret_cast< quint64(*)>(_a[1]))); break;
        case 3: _t->markedChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (QAddress539e18::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QAddress539e18::addressChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (QAddress539e18::*)(quint64 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QAddress539e18::addressSkyChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (QAddress539e18::*)(quint64 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QAddress539e18::addressCoinHoursChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (QAddress539e18::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QAddress539e18::markedChanged)) {
                *result = 3;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<QAddress539e18 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QString*>(_v) = _t->address(); break;
        case 1: *reinterpret_cast< quint64*>(_v) = _t->addressSky(); break;
        case 2: *reinterpret_cast< quint64*>(_v) = _t->addressCoinHours(); break;
        case 3: *reinterpret_cast< qint32*>(_v) = _t->marked(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<QAddress539e18 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setAddress(*reinterpret_cast< QString*>(_v)); break;
        case 1: _t->setAddressSky(*reinterpret_cast< quint64*>(_v)); break;
        case 2: _t->setAddressCoinHours(*reinterpret_cast< quint64*>(_v)); break;
        case 3: _t->setMarked(*reinterpret_cast< qint32*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject QAddress539e18::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_QAddress539e18.data,
    qt_meta_data_QAddress539e18,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *QAddress539e18::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *QAddress539e18::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_QAddress539e18.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int QAddress539e18::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 4)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 4;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 4)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 4;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 4;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 4;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 4;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 4;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 4;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 4;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void QAddress539e18::addressChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void QAddress539e18::addressSkyChanged(quint64 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void QAddress539e18::addressCoinHoursChanged(quint64 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void QAddress539e18::markedChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}
struct qt_meta_stringdata_QWallet539e18_t {
    QByteArrayData data[12];
    char stringdata0[142];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_QWallet539e18_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_QWallet539e18_t qt_meta_stringdata_QWallet539e18 = {
    {
QT_MOC_LITERAL(0, 0, 13), // "QWallet539e18"
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
    "QWallet539e18\0nameChanged\0\0name\0"
    "encryptionEnabledChanged\0encryptionEnabled\0"
    "skyChanged\0sky\0coinHoursChanged\0"
    "coinHours\0fileNameChanged\0fileName"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_QWallet539e18[] = {

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

void QWallet539e18::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<QWallet539e18 *>(_o);
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
            using _t = void (QWallet539e18::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QWallet539e18::nameChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (QWallet539e18::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QWallet539e18::encryptionEnabledChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (QWallet539e18::*)(quint64 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QWallet539e18::skyChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (QWallet539e18::*)(quint64 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QWallet539e18::coinHoursChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (QWallet539e18::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QWallet539e18::fileNameChanged)) {
                *result = 4;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<QWallet539e18 *>(_o);
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
        auto *_t = static_cast<QWallet539e18 *>(_o);
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

QT_INIT_METAOBJECT const QMetaObject QWallet539e18::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_QWallet539e18.data,
    qt_meta_data_QWallet539e18,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *QWallet539e18::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *QWallet539e18::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_QWallet539e18.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int QWallet539e18::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
void QWallet539e18::nameChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void QWallet539e18::encryptionEnabledChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void QWallet539e18::skyChanged(quint64 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void QWallet539e18::coinHoursChanged(quint64 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void QWallet539e18::fileNameChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}
struct qt_meta_stringdata_WalletManager539e18_t {
    QByteArrayData data[21];
    char stringdata0[257];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_WalletManager539e18_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_WalletManager539e18_t qt_meta_stringdata_WalletManager539e18 = {
    {
QT_MOC_LITERAL(0, 0, 19), // "WalletManager539e18"
QT_MOC_LITERAL(1, 20, 21), // "createEncryptedWallet"
QT_MOC_LITERAL(2, 42, 14), // "QWallet539e18*"
QT_MOC_LITERAL(3, 57, 0), // ""
QT_MOC_LITERAL(4, 58, 4), // "seed"
QT_MOC_LITERAL(5, 63, 5), // "label"
QT_MOC_LITERAL(6, 69, 8), // "password"
QT_MOC_LITERAL(7, 78, 5), // "scanN"
QT_MOC_LITERAL(8, 84, 23), // "createUnencryptedWallet"
QT_MOC_LITERAL(9, 108, 10), // "getNewSeed"
QT_MOC_LITERAL(10, 119, 7), // "entropy"
QT_MOC_LITERAL(11, 127, 10), // "verifySeed"
QT_MOC_LITERAL(12, 138, 16), // "newWalletAddress"
QT_MOC_LITERAL(13, 155, 2), // "id"
QT_MOC_LITERAL(14, 158, 1), // "n"
QT_MOC_LITERAL(15, 160, 13), // "encryptWallet"
QT_MOC_LITERAL(16, 174, 13), // "decryptWallet"
QT_MOC_LITERAL(17, 188, 10), // "getWallets"
QT_MOC_LITERAL(18, 199, 21), // "QList<QWallet539e18*>"
QT_MOC_LITERAL(19, 221, 12), // "getAddresses"
QT_MOC_LITERAL(20, 234, 22) // "QList<QAddress539e18*>"

    },
    "WalletManager539e18\0createEncryptedWallet\0"
    "QWallet539e18*\0\0seed\0label\0password\0"
    "scanN\0createUnencryptedWallet\0getNewSeed\0"
    "entropy\0verifySeed\0newWalletAddress\0"
    "id\0n\0encryptWallet\0decryptWallet\0"
    "getWallets\0QList<QWallet539e18*>\0"
    "getAddresses\0QList<QAddress539e18*>"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_WalletManager539e18[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       9,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       0,       // signalCount

 // slots: name, argc, parameters, tag, flags
       1,    4,   59,    3, 0x0a /* Public */,
       8,    3,   68,    3, 0x0a /* Public */,
       9,    1,   75,    3, 0x0a /* Public */,
      11,    1,   78,    3, 0x0a /* Public */,
      12,    3,   81,    3, 0x0a /* Public */,
      15,    2,   88,    3, 0x0a /* Public */,
      16,    2,   93,    3, 0x0a /* Public */,
      17,    0,   98,    3, 0x0a /* Public */,
      19,    1,   99,    3, 0x0a /* Public */,

 // slots: parameters
    0x80000000 | 2, QMetaType::QString, QMetaType::QString, QMetaType::QString, QMetaType::Int,    4,    5,    6,    7,
    0x80000000 | 2, QMetaType::QString, QMetaType::QString, QMetaType::Int,    4,    5,    7,
    QMetaType::QString, QMetaType::Int,   10,
    QMetaType::Int, QMetaType::QString,    4,
    QMetaType::Void, QMetaType::QString, QMetaType::Int, QMetaType::QString,   13,   14,    6,
    QMetaType::Void, QMetaType::QString, QMetaType::QString,   13,    6,
    QMetaType::Void, QMetaType::QString, QMetaType::QString,   13,    6,
    0x80000000 | 18,
    0x80000000 | 20, QMetaType::QString,   13,

       0        // eod
};

void WalletManager539e18::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<WalletManager539e18 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: { QWallet539e18* _r = _t->createEncryptedWallet((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< QString(*)>(_a[3])),(*reinterpret_cast< qint32(*)>(_a[4])));
            if (_a[0]) *reinterpret_cast< QWallet539e18**>(_a[0]) = std::move(_r); }  break;
        case 1: { QWallet539e18* _r = _t->createUnencryptedWallet((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< qint32(*)>(_a[3])));
            if (_a[0]) *reinterpret_cast< QWallet539e18**>(_a[0]) = std::move(_r); }  break;
        case 2: { QString _r = _t->getNewSeed((*reinterpret_cast< qint32(*)>(_a[1])));
            if (_a[0]) *reinterpret_cast< QString*>(_a[0]) = std::move(_r); }  break;
        case 3: { qint32 _r = _t->verifySeed((*reinterpret_cast< QString(*)>(_a[1])));
            if (_a[0]) *reinterpret_cast< qint32*>(_a[0]) = std::move(_r); }  break;
        case 4: _t->newWalletAddress((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< qint32(*)>(_a[2])),(*reinterpret_cast< QString(*)>(_a[3]))); break;
        case 5: _t->encryptWallet((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2]))); break;
        case 6: _t->decryptWallet((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2]))); break;
        case 7: { QList<QWallet539e18*> _r = _t->getWallets();
            if (_a[0]) *reinterpret_cast< QList<QWallet539e18*>*>(_a[0]) = std::move(_r); }  break;
        case 8: { QList<QAddress539e18*> _r = _t->getAddresses((*reinterpret_cast< QString(*)>(_a[1])));
            if (_a[0]) *reinterpret_cast< QList<QAddress539e18*>*>(_a[0]) = std::move(_r); }  break;
        default: ;
        }
    }
}

QT_INIT_METAOBJECT const QMetaObject WalletManager539e18::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_WalletManager539e18.data,
    qt_meta_data_WalletManager539e18,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *WalletManager539e18::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *WalletManager539e18::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_WalletManager539e18.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int WalletManager539e18::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 9)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 9;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 9)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 9;
    }
    return _id;
}
struct qt_meta_stringdata_WalletModel539e18_t {
    QByteArrayData data[21];
    char stringdata0[216];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_WalletModel539e18_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_WalletModel539e18_t qt_meta_stringdata_WalletModel539e18 = {
    {
QT_MOC_LITERAL(0, 0, 17), // "WalletModel539e18"
QT_MOC_LITERAL(1, 18, 12), // "rolesChanged"
QT_MOC_LITERAL(2, 31, 0), // ""
QT_MOC_LITERAL(3, 32, 10), // "type378cdd"
QT_MOC_LITERAL(4, 43, 5), // "roles"
QT_MOC_LITERAL(5, 49, 14), // "walletsChanged"
QT_MOC_LITERAL(6, 64, 21), // "QList<QWallet539e18*>"
QT_MOC_LITERAL(7, 86, 7), // "wallets"
QT_MOC_LITERAL(8, 94, 12), // "countChanged"
QT_MOC_LITERAL(9, 107, 5), // "count"
QT_MOC_LITERAL(10, 113, 9), // "addWallet"
QT_MOC_LITERAL(11, 123, 14), // "QWallet539e18*"
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
    "WalletModel539e18\0rolesChanged\0\0"
    "type378cdd\0roles\0walletsChanged\0"
    "QList<QWallet539e18*>\0wallets\0"
    "countChanged\0count\0addWallet\0"
    "QWallet539e18*\0v0\0editWallet\0row\0name\0"
    "encryptionEnabled\0sky\0coinHours\0"
    "removeWallet\0loadModel"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_WalletModel539e18[] = {

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

void WalletModel539e18::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<WalletModel539e18 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->rolesChanged((*reinterpret_cast< type378cdd(*)>(_a[1]))); break;
        case 1: _t->walletsChanged((*reinterpret_cast< QList<QWallet539e18*>(*)>(_a[1]))); break;
        case 2: _t->countChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 3: _t->addWallet((*reinterpret_cast< QWallet539e18*(*)>(_a[1]))); break;
        case 4: _t->editWallet((*reinterpret_cast< qint32(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< bool(*)>(_a[3])),(*reinterpret_cast< quint64(*)>(_a[4])),(*reinterpret_cast< quint64(*)>(_a[5]))); break;
        case 5: _t->removeWallet((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 6: _t->loadModel((*reinterpret_cast< QList<QWallet539e18*>(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<QWallet539e18*> >(); break;
            }
            break;
        case 3:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QWallet539e18* >(); break;
            }
            break;
        case 6:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<QWallet539e18*> >(); break;
            }
            break;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (WalletModel539e18::*)(type378cdd );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&WalletModel539e18::rolesChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (WalletModel539e18::*)(QList<QWallet539e18*> );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&WalletModel539e18::walletsChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (WalletModel539e18::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&WalletModel539e18::countChanged)) {
                *result = 2;
                return;
            }
        }
    } else if (_c == QMetaObject::RegisterPropertyMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<QWallet539e18*> >(); break;
        }
    }

#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<WalletModel539e18 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< type378cdd*>(_v) = _t->roles(); break;
        case 1: *reinterpret_cast< QList<QWallet539e18*>*>(_v) = _t->wallets(); break;
        case 2: *reinterpret_cast< qint32*>(_v) = _t->count(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<WalletModel539e18 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRoles(*reinterpret_cast< type378cdd*>(_v)); break;
        case 1: _t->setWallets(*reinterpret_cast< QList<QWallet539e18*>*>(_v)); break;
        case 2: _t->setCount(*reinterpret_cast< qint32*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject WalletModel539e18::staticMetaObject = { {
    &QAbstractListModel::staticMetaObject,
    qt_meta_stringdata_WalletModel539e18.data,
    qt_meta_data_WalletModel539e18,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *WalletModel539e18::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *WalletModel539e18::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_WalletModel539e18.stringdata0))
        return static_cast<void*>(this);
    return QAbstractListModel::qt_metacast(_clname);
}

int WalletModel539e18::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
void WalletModel539e18::rolesChanged(type378cdd _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void WalletModel539e18::walletsChanged(QList<QWallet539e18*> _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void WalletModel539e18::countChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
