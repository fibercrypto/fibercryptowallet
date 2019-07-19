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
struct qt_meta_stringdata_QWallet268d39_t {
    QByteArrayData data[10];
    char stringdata0[117];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_QWallet268d39_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_QWallet268d39_t qt_meta_stringdata_QWallet268d39 = {
    {
QT_MOC_LITERAL(0, 0, 13), // "QWallet268d39"
QT_MOC_LITERAL(1, 14, 11), // "nameChanged"
QT_MOC_LITERAL(2, 26, 0), // ""
QT_MOC_LITERAL(3, 27, 4), // "name"
QT_MOC_LITERAL(4, 32, 24), // "encryptionEnabledChanged"
QT_MOC_LITERAL(5, 57, 17), // "encryptionEnabled"
QT_MOC_LITERAL(6, 75, 10), // "skyChanged"
QT_MOC_LITERAL(7, 86, 3), // "sky"
QT_MOC_LITERAL(8, 90, 16), // "coinHoursChanged"
QT_MOC_LITERAL(9, 107, 9) // "coinHours"

    },
    "QWallet268d39\0nameChanged\0\0name\0"
    "encryptionEnabledChanged\0encryptionEnabled\0"
    "skyChanged\0sky\0coinHoursChanged\0"
    "coinHours"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_QWallet268d39[] = {

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
    QMetaType::Void, QMetaType::Int,    5,
    QMetaType::Void, QMetaType::Int,    7,
    QMetaType::Void, QMetaType::Int,    9,

 // properties: name, type, flags
       3, QMetaType::QString, 0x00495103,
       5, QMetaType::Int, 0x00495103,
       7, QMetaType::Int, 0x00495103,
       9, QMetaType::Int, 0x00495103,

 // properties: notify_signal_id
       0,
       1,
       2,
       3,

       0        // eod
};

void QWallet268d39::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<QWallet268d39 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->nameChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 1: _t->encryptionEnabledChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 2: _t->skyChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 3: _t->coinHoursChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (QWallet268d39::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QWallet268d39::nameChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (QWallet268d39::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QWallet268d39::encryptionEnabledChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (QWallet268d39::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QWallet268d39::skyChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (QWallet268d39::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&QWallet268d39::coinHoursChanged)) {
                *result = 3;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<QWallet268d39 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QString*>(_v) = _t->name(); break;
        case 1: *reinterpret_cast< qint32*>(_v) = _t->encryptionEnabled(); break;
        case 2: *reinterpret_cast< qint32*>(_v) = _t->sky(); break;
        case 3: *reinterpret_cast< qint32*>(_v) = _t->coinHours(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<QWallet268d39 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setName(*reinterpret_cast< QString*>(_v)); break;
        case 1: _t->setEncryptionEnabled(*reinterpret_cast< qint32*>(_v)); break;
        case 2: _t->setSky(*reinterpret_cast< qint32*>(_v)); break;
        case 3: _t->setCoinHours(*reinterpret_cast< qint32*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject QWallet268d39::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_QWallet268d39.data,
    qt_meta_data_QWallet268d39,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *QWallet268d39::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *QWallet268d39::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_QWallet268d39.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int QWallet268d39::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
void QWallet268d39::nameChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void QWallet268d39::encryptionEnabledChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void QWallet268d39::skyChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void QWallet268d39::coinHoursChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}
struct qt_meta_stringdata_WalletModel268d39_t {
    QByteArrayData data[19];
    char stringdata0[197];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_WalletModel268d39_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_WalletModel268d39_t qt_meta_stringdata_WalletModel268d39 = {
    {
QT_MOC_LITERAL(0, 0, 17), // "WalletModel268d39"
QT_MOC_LITERAL(1, 18, 12), // "rolesChanged"
QT_MOC_LITERAL(2, 31, 0), // ""
QT_MOC_LITERAL(3, 32, 10), // "type378cdd"
QT_MOC_LITERAL(4, 43, 5), // "roles"
QT_MOC_LITERAL(5, 49, 14), // "walletsChanged"
QT_MOC_LITERAL(6, 64, 21), // "QList<QWallet268d39*>"
QT_MOC_LITERAL(7, 86, 7), // "wallets"
QT_MOC_LITERAL(8, 94, 9), // "addWallet"
QT_MOC_LITERAL(9, 104, 14), // "QWallet268d39*"
QT_MOC_LITERAL(10, 119, 2), // "v0"
QT_MOC_LITERAL(11, 122, 10), // "editWallet"
QT_MOC_LITERAL(12, 133, 3), // "row"
QT_MOC_LITERAL(13, 137, 4), // "name"
QT_MOC_LITERAL(14, 142, 17), // "encryptionEnabled"
QT_MOC_LITERAL(15, 160, 3), // "sky"
QT_MOC_LITERAL(16, 164, 9), // "coinHours"
QT_MOC_LITERAL(17, 174, 12), // "removeWallet"
QT_MOC_LITERAL(18, 187, 9) // "loadModel"

    },
    "WalletModel268d39\0rolesChanged\0\0"
    "type378cdd\0roles\0walletsChanged\0"
    "QList<QWallet268d39*>\0wallets\0addWallet\0"
    "QWallet268d39*\0v0\0editWallet\0row\0name\0"
    "encryptionEnabled\0sky\0coinHours\0"
    "removeWallet\0loadModel"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_WalletModel268d39[] = {

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
      11,    5,   53,    2, 0x0a /* Public */,
      17,    1,   64,    2, 0x0a /* Public */,
      18,    0,   67,    2, 0x0a /* Public */,

 // signals: parameters
    QMetaType::Void, 0x80000000 | 3,    4,
    QMetaType::Void, 0x80000000 | 6,    7,

 // slots: parameters
    QMetaType::Void, 0x80000000 | 9,   10,
    QMetaType::Void, QMetaType::Int, QMetaType::QString, QMetaType::Bool, QMetaType::Int, QMetaType::Int,   12,   13,   14,   15,   16,
    QMetaType::Void, QMetaType::Int,   12,
    QMetaType::Void,

 // properties: name, type, flags
       4, 0x80000000 | 3, 0x0049510b,
       7, 0x80000000 | 6, 0x0049510b,

 // properties: notify_signal_id
       0,
       1,

       0        // eod
};

void WalletModel268d39::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<WalletModel268d39 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->rolesChanged((*reinterpret_cast< type378cdd(*)>(_a[1]))); break;
        case 1: _t->walletsChanged((*reinterpret_cast< QList<QWallet268d39*>(*)>(_a[1]))); break;
        case 2: _t->addWallet((*reinterpret_cast< QWallet268d39*(*)>(_a[1]))); break;
        case 3: _t->editWallet((*reinterpret_cast< qint32(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< bool(*)>(_a[3])),(*reinterpret_cast< qint32(*)>(_a[4])),(*reinterpret_cast< qint32(*)>(_a[5]))); break;
        case 4: _t->removeWallet((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 5: _t->loadModel(); break;
        default: ;
        }
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<QWallet268d39*> >(); break;
            }
            break;
        case 2:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QWallet268d39* >(); break;
            }
            break;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (WalletModel268d39::*)(type378cdd );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&WalletModel268d39::rolesChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (WalletModel268d39::*)(QList<QWallet268d39*> );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&WalletModel268d39::walletsChanged)) {
                *result = 1;
                return;
            }
        }
    } else if (_c == QMetaObject::RegisterPropertyMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<QWallet268d39*> >(); break;
        }
    }

#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<WalletModel268d39 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< type378cdd*>(_v) = _t->roles(); break;
        case 1: *reinterpret_cast< QList<QWallet268d39*>*>(_v) = _t->wallets(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<WalletModel268d39 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRoles(*reinterpret_cast< type378cdd*>(_v)); break;
        case 1: _t->setWallets(*reinterpret_cast< QList<QWallet268d39*>*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject WalletModel268d39::staticMetaObject = { {
    &QAbstractListModel::staticMetaObject,
    qt_meta_stringdata_WalletModel268d39.data,
    qt_meta_data_WalletModel268d39,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *WalletModel268d39::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *WalletModel268d39::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_WalletModel268d39.stringdata0))
        return static_cast<void*>(this);
    return QAbstractListModel::qt_metacast(_clname);
}

int WalletModel268d39::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
void WalletModel268d39::rolesChanged(type378cdd _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void WalletModel268d39::walletsChanged(QList<QWallet268d39*> _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
