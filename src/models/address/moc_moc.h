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
struct qt_meta_stringdata_AddressDetailsfb4c44_t {
    QByteArrayData data[8];
    char stringdata0[115];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_AddressDetailsfb4c44_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_AddressDetailsfb4c44_t qt_meta_stringdata_AddressDetailsfb4c44 = {
    {
QT_MOC_LITERAL(0, 0, 20), // "AddressDetailsfb4c44"
QT_MOC_LITERAL(1, 21, 14), // "addressChanged"
QT_MOC_LITERAL(2, 36, 0), // ""
QT_MOC_LITERAL(3, 37, 7), // "address"
QT_MOC_LITERAL(4, 45, 17), // "addressSkyChanged"
QT_MOC_LITERAL(5, 63, 10), // "addressSky"
QT_MOC_LITERAL(6, 74, 23), // "addressCoinHoursChanged"
QT_MOC_LITERAL(7, 98, 16) // "addressCoinHours"

    },
    "AddressDetailsfb4c44\0addressChanged\0"
    "\0address\0addressSkyChanged\0addressSky\0"
    "addressCoinHoursChanged\0addressCoinHours"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_AddressDetailsfb4c44[] = {

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
    QMetaType::Void, QMetaType::QString,    5,
    QMetaType::Void, QMetaType::QString,    7,

 // properties: name, type, flags
       3, QMetaType::QString, 0x00495103,
       5, QMetaType::QString, 0x00495103,
       7, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
       0,
       1,
       2,

       0        // eod
};

void AddressDetailsfb4c44::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<AddressDetailsfb4c44 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->addressChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 1: _t->addressSkyChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 2: _t->addressCoinHoursChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (AddressDetailsfb4c44::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressDetailsfb4c44::addressChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (AddressDetailsfb4c44::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressDetailsfb4c44::addressSkyChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (AddressDetailsfb4c44::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressDetailsfb4c44::addressCoinHoursChanged)) {
                *result = 2;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<AddressDetailsfb4c44 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QString*>(_v) = _t->address(); break;
        case 1: *reinterpret_cast< QString*>(_v) = _t->addressSky(); break;
        case 2: *reinterpret_cast< QString*>(_v) = _t->addressCoinHours(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<AddressDetailsfb4c44 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setAddress(*reinterpret_cast< QString*>(_v)); break;
        case 1: _t->setAddressSky(*reinterpret_cast< QString*>(_v)); break;
        case 2: _t->setAddressCoinHours(*reinterpret_cast< QString*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject AddressDetailsfb4c44::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_AddressDetailsfb4c44.data,
    qt_meta_data_AddressDetailsfb4c44,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *AddressDetailsfb4c44::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *AddressDetailsfb4c44::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_AddressDetailsfb4c44.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int AddressDetailsfb4c44::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
void AddressDetailsfb4c44::addressChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void AddressDetailsfb4c44::addressSkyChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void AddressDetailsfb4c44::addressCoinHoursChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}
struct qt_meta_stringdata_AddressListfb4c44_t {
    QByteArrayData data[13];
    char stringdata0[170];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_AddressListfb4c44_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_AddressListfb4c44_t qt_meta_stringdata_AddressListfb4c44 = {
    {
QT_MOC_LITERAL(0, 0, 17), // "AddressListfb4c44"
QT_MOC_LITERAL(1, 18, 10), // "addAddress"
QT_MOC_LITERAL(2, 29, 0), // ""
QT_MOC_LITERAL(3, 30, 21), // "AddressDetailsfb4c44*"
QT_MOC_LITERAL(4, 52, 11), // "transaction"
QT_MOC_LITERAL(5, 64, 13), // "removeAddress"
QT_MOC_LITERAL(6, 78, 5), // "index"
QT_MOC_LITERAL(7, 84, 12), // "rolesChanged"
QT_MOC_LITERAL(8, 97, 10), // "type378cdd"
QT_MOC_LITERAL(9, 108, 5), // "roles"
QT_MOC_LITERAL(10, 114, 16), // "addressesChanged"
QT_MOC_LITERAL(11, 131, 28), // "QList<AddressDetailsfb4c44*>"
QT_MOC_LITERAL(12, 160, 9) // "addresses"

    },
    "AddressListfb4c44\0addAddress\0\0"
    "AddressDetailsfb4c44*\0transaction\0"
    "removeAddress\0index\0rolesChanged\0"
    "type378cdd\0roles\0addressesChanged\0"
    "QList<AddressDetailsfb4c44*>\0addresses"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_AddressListfb4c44[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       4,   14, // methods
       2,   46, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       4,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   34,    2, 0x06 /* Public */,
       5,    1,   37,    2, 0x06 /* Public */,
       7,    1,   40,    2, 0x06 /* Public */,
      10,    1,   43,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, 0x80000000 | 3,    4,
    QMetaType::Void, QMetaType::Int,    6,
    QMetaType::Void, 0x80000000 | 8,    9,
    QMetaType::Void, 0x80000000 | 11,   12,

 // properties: name, type, flags
       9, 0x80000000 | 8, 0x0049510b,
      12, 0x80000000 | 11, 0x0049510b,

 // properties: notify_signal_id
       2,
       3,

       0        // eod
};

void AddressListfb4c44::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<AddressListfb4c44 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->addAddress((*reinterpret_cast< AddressDetailsfb4c44*(*)>(_a[1]))); break;
        case 1: _t->removeAddress((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 2: _t->rolesChanged((*reinterpret_cast< type378cdd(*)>(_a[1]))); break;
        case 3: _t->addressesChanged((*reinterpret_cast< QList<AddressDetailsfb4c44*>(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 0:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< AddressDetailsfb4c44* >(); break;
            }
            break;
        case 3:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<AddressDetailsfb4c44*> >(); break;
            }
            break;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (AddressListfb4c44::*)(AddressDetailsfb4c44 * );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressListfb4c44::addAddress)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (AddressListfb4c44::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressListfb4c44::removeAddress)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (AddressListfb4c44::*)(type378cdd );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressListfb4c44::rolesChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (AddressListfb4c44::*)(QList<AddressDetailsfb4c44*> );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressListfb4c44::addressesChanged)) {
                *result = 3;
                return;
            }
        }
    } else if (_c == QMetaObject::RegisterPropertyMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<AddressDetailsfb4c44*> >(); break;
        }
    }

#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<AddressListfb4c44 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< type378cdd*>(_v) = _t->roles(); break;
        case 1: *reinterpret_cast< QList<AddressDetailsfb4c44*>*>(_v) = _t->addresses(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<AddressListfb4c44 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRoles(*reinterpret_cast< type378cdd*>(_v)); break;
        case 1: _t->setAddresses(*reinterpret_cast< QList<AddressDetailsfb4c44*>*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject AddressListfb4c44::staticMetaObject = { {
    &QAbstractListModel::staticMetaObject,
    qt_meta_stringdata_AddressListfb4c44.data,
    qt_meta_data_AddressListfb4c44,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *AddressListfb4c44::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *AddressListfb4c44::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_AddressListfb4c44.stringdata0))
        return static_cast<void*>(this);
    return QAbstractListModel::qt_metacast(_clname);
}

int AddressListfb4c44::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QAbstractListModel::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 4)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 4;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 4)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 4;
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
void AddressListfb4c44::addAddress(AddressDetailsfb4c44 * _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void AddressListfb4c44::removeAddress(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void AddressListfb4c44::rolesChanged(type378cdd _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void AddressListfb4c44::addressesChanged(QList<AddressDetailsfb4c44*> _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
