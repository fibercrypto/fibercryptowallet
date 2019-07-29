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
struct qt_meta_stringdata_BlockchainStatusModel0d0160_t {
    QByteArrayData data[17];
    char stringdata0[333];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_BlockchainStatusModel0d0160_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_BlockchainStatusModel0d0160_t qt_meta_stringdata_BlockchainStatusModel0d0160 = {
    {
QT_MOC_LITERAL(0, 0, 27), // "BlockchainStatusModel0d0160"
QT_MOC_LITERAL(1, 28, 6), // "update"
QT_MOC_LITERAL(2, 35, 0), // ""
QT_MOC_LITERAL(3, 36, 21), // "numberOfBlocksChanged"
QT_MOC_LITERAL(4, 58, 14), // "numberOfBlocks"
QT_MOC_LITERAL(5, 73, 25), // "timestampLastBlockChanged"
QT_MOC_LITERAL(6, 99, 18), // "timestampLastBlock"
QT_MOC_LITERAL(7, 118, 20), // "hashLastBlockChanged"
QT_MOC_LITERAL(8, 139, 13), // "hashLastBlock"
QT_MOC_LITERAL(9, 153, 23), // "currentSkySupplyChanged"
QT_MOC_LITERAL(10, 177, 16), // "currentSkySupply"
QT_MOC_LITERAL(11, 194, 21), // "totalSkySupplyChanged"
QT_MOC_LITERAL(12, 216, 14), // "totalSkySupply"
QT_MOC_LITERAL(13, 231, 29), // "currentCoinHoursSupplyChanged"
QT_MOC_LITERAL(14, 261, 22), // "currentCoinHoursSupply"
QT_MOC_LITERAL(15, 284, 27), // "totalCoinHoursSupplyChanged"
QT_MOC_LITERAL(16, 312, 20) // "totalCoinHoursSupply"

    },
    "BlockchainStatusModel0d0160\0update\0\0"
    "numberOfBlocksChanged\0numberOfBlocks\0"
    "timestampLastBlockChanged\0timestampLastBlock\0"
    "hashLastBlockChanged\0hashLastBlock\0"
    "currentSkySupplyChanged\0currentSkySupply\0"
    "totalSkySupplyChanged\0totalSkySupply\0"
    "currentCoinHoursSupplyChanged\0"
    "currentCoinHoursSupply\0"
    "totalCoinHoursSupplyChanged\0"
    "totalCoinHoursSupply"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_BlockchainStatusModel0d0160[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       8,   14, // methods
       7,   76, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       8,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    0,   54,    2, 0x06 /* Public */,
       3,    1,   55,    2, 0x06 /* Public */,
       5,    1,   58,    2, 0x06 /* Public */,
       7,    1,   61,    2, 0x06 /* Public */,
       9,    1,   64,    2, 0x06 /* Public */,
      11,    1,   67,    2, 0x06 /* Public */,
      13,    1,   70,    2, 0x06 /* Public */,
      15,    1,   73,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void,
    QMetaType::Void, QMetaType::QString,    4,
    QMetaType::Void, QMetaType::QDateTime,    6,
    QMetaType::Void, QMetaType::QString,    8,
    QMetaType::Void, QMetaType::QString,   10,
    QMetaType::Void, QMetaType::QString,   12,
    QMetaType::Void, QMetaType::QString,   14,
    QMetaType::Void, QMetaType::QString,   16,

 // properties: name, type, flags
       4, QMetaType::QString, 0x00495103,
       6, QMetaType::QDateTime, 0x00495103,
       8, QMetaType::QString, 0x00495103,
      10, QMetaType::QString, 0x00495103,
      12, QMetaType::QString, 0x00495103,
      14, QMetaType::QString, 0x00495103,
      16, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
       1,
       2,
       3,
       4,
       5,
       6,
       7,

       0        // eod
};

void BlockchainStatusModel0d0160::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<BlockchainStatusModel0d0160 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->update(); break;
        case 1: _t->numberOfBlocksChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 2: _t->timestampLastBlockChanged((*reinterpret_cast< QDateTime(*)>(_a[1]))); break;
        case 3: _t->hashLastBlockChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 4: _t->currentSkySupplyChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 5: _t->totalSkySupplyChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 6: _t->currentCoinHoursSupplyChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 7: _t->totalCoinHoursSupplyChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (BlockchainStatusModel0d0160::*)();
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&BlockchainStatusModel0d0160::update)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (BlockchainStatusModel0d0160::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&BlockchainStatusModel0d0160::numberOfBlocksChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (BlockchainStatusModel0d0160::*)(QDateTime );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&BlockchainStatusModel0d0160::timestampLastBlockChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (BlockchainStatusModel0d0160::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&BlockchainStatusModel0d0160::hashLastBlockChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (BlockchainStatusModel0d0160::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&BlockchainStatusModel0d0160::currentSkySupplyChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (BlockchainStatusModel0d0160::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&BlockchainStatusModel0d0160::totalSkySupplyChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (BlockchainStatusModel0d0160::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&BlockchainStatusModel0d0160::currentCoinHoursSupplyChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (BlockchainStatusModel0d0160::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&BlockchainStatusModel0d0160::totalCoinHoursSupplyChanged)) {
                *result = 7;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<BlockchainStatusModel0d0160 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QString*>(_v) = _t->numberOfBlocks(); break;
        case 1: *reinterpret_cast< QDateTime*>(_v) = _t->timestampLastBlock(); break;
        case 2: *reinterpret_cast< QString*>(_v) = _t->hashLastBlock(); break;
        case 3: *reinterpret_cast< QString*>(_v) = _t->currentSkySupply(); break;
        case 4: *reinterpret_cast< QString*>(_v) = _t->totalSkySupply(); break;
        case 5: *reinterpret_cast< QString*>(_v) = _t->currentCoinHoursSupply(); break;
        case 6: *reinterpret_cast< QString*>(_v) = _t->totalCoinHoursSupply(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<BlockchainStatusModel0d0160 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setNumberOfBlocks(*reinterpret_cast< QString*>(_v)); break;
        case 1: _t->setTimestampLastBlock(*reinterpret_cast< QDateTime*>(_v)); break;
        case 2: _t->setHashLastBlock(*reinterpret_cast< QString*>(_v)); break;
        case 3: _t->setCurrentSkySupply(*reinterpret_cast< QString*>(_v)); break;
        case 4: _t->setTotalSkySupply(*reinterpret_cast< QString*>(_v)); break;
        case 5: _t->setCurrentCoinHoursSupply(*reinterpret_cast< QString*>(_v)); break;
        case 6: _t->setTotalCoinHoursSupply(*reinterpret_cast< QString*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject BlockchainStatusModel0d0160::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_BlockchainStatusModel0d0160.data,
    qt_meta_data_BlockchainStatusModel0d0160,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *BlockchainStatusModel0d0160::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *BlockchainStatusModel0d0160::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_BlockchainStatusModel0d0160.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int BlockchainStatusModel0d0160::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 8)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 8;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 8)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 8;
    }
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 7;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 7;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 7;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 7;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 7;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 7;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void BlockchainStatusModel0d0160::update()
{
    QMetaObject::activate(this, &staticMetaObject, 0, nullptr);
}

// SIGNAL 1
void BlockchainStatusModel0d0160::numberOfBlocksChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void BlockchainStatusModel0d0160::timestampLastBlockChanged(QDateTime _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void BlockchainStatusModel0d0160::hashLastBlockChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void BlockchainStatusModel0d0160::currentSkySupplyChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void BlockchainStatusModel0d0160::totalSkySupplyChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void BlockchainStatusModel0d0160::currentCoinHoursSupplyChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void BlockchainStatusModel0d0160::totalCoinHoursSupplyChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
