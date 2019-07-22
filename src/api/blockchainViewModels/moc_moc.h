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
struct qt_meta_stringdata_BlockchainStatusModel97d618_t {
    QByteArrayData data[16];
    char stringdata0[326];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_BlockchainStatusModel97d618_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_BlockchainStatusModel97d618_t qt_meta_stringdata_BlockchainStatusModel97d618 = {
    {
QT_MOC_LITERAL(0, 0, 27), // "BlockchainStatusModel97d618"
QT_MOC_LITERAL(1, 28, 21), // "numberOfBlocksChanged"
QT_MOC_LITERAL(2, 50, 0), // ""
QT_MOC_LITERAL(3, 51, 14), // "numberOfBlocks"
QT_MOC_LITERAL(4, 66, 25), // "timestampLastBlockChanged"
QT_MOC_LITERAL(5, 92, 18), // "timestampLastBlock"
QT_MOC_LITERAL(6, 111, 20), // "hashLastBlockChanged"
QT_MOC_LITERAL(7, 132, 13), // "hashLastBlock"
QT_MOC_LITERAL(8, 146, 23), // "currentSkySupplyChanged"
QT_MOC_LITERAL(9, 170, 16), // "currentSkySupply"
QT_MOC_LITERAL(10, 187, 21), // "totalSkySupplyChanged"
QT_MOC_LITERAL(11, 209, 14), // "totalSkySupply"
QT_MOC_LITERAL(12, 224, 29), // "currentCoinHoursSupplyChanged"
QT_MOC_LITERAL(13, 254, 22), // "currentCoinHoursSupply"
QT_MOC_LITERAL(14, 277, 27), // "totalCoinHoursSupplyChanged"
QT_MOC_LITERAL(15, 305, 20) // "totalCoinHoursSupply"

    },
    "BlockchainStatusModel97d618\0"
    "numberOfBlocksChanged\0\0numberOfBlocks\0"
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

static const uint qt_meta_data_BlockchainStatusModel97d618[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       7,   14, // methods
       7,   70, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       7,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   49,    2, 0x06 /* Public */,
       4,    1,   52,    2, 0x06 /* Public */,
       6,    1,   55,    2, 0x06 /* Public */,
       8,    1,   58,    2, 0x06 /* Public */,
      10,    1,   61,    2, 0x06 /* Public */,
      12,    1,   64,    2, 0x06 /* Public */,
      14,    1,   67,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::Int,    3,
    QMetaType::Void, QMetaType::QDateTime,    5,
    QMetaType::Void, QMetaType::QString,    7,
    QMetaType::Void, QMetaType::QString,    9,
    QMetaType::Void, QMetaType::QString,   11,
    QMetaType::Void, QMetaType::QString,   13,
    QMetaType::Void, QMetaType::QString,   15,

 // properties: name, type, flags
       3, QMetaType::Int, 0x00495103,
       5, QMetaType::QDateTime, 0x00495103,
       7, QMetaType::QString, 0x00495103,
       9, QMetaType::QString, 0x00495103,
      11, QMetaType::QString, 0x00495103,
      13, QMetaType::QString, 0x00495103,
      15, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
       0,
       1,
       2,
       3,
       4,
       5,
       6,

       0        // eod
};

void BlockchainStatusModel97d618::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<BlockchainStatusModel97d618 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->numberOfBlocksChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 1: _t->timestampLastBlockChanged((*reinterpret_cast< QDateTime(*)>(_a[1]))); break;
        case 2: _t->hashLastBlockChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 3: _t->currentSkySupplyChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 4: _t->totalSkySupplyChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 5: _t->currentCoinHoursSupplyChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 6: _t->totalCoinHoursSupplyChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (BlockchainStatusModel97d618::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&BlockchainStatusModel97d618::numberOfBlocksChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (BlockchainStatusModel97d618::*)(QDateTime );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&BlockchainStatusModel97d618::timestampLastBlockChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (BlockchainStatusModel97d618::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&BlockchainStatusModel97d618::hashLastBlockChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (BlockchainStatusModel97d618::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&BlockchainStatusModel97d618::currentSkySupplyChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (BlockchainStatusModel97d618::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&BlockchainStatusModel97d618::totalSkySupplyChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (BlockchainStatusModel97d618::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&BlockchainStatusModel97d618::currentCoinHoursSupplyChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (BlockchainStatusModel97d618::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&BlockchainStatusModel97d618::totalCoinHoursSupplyChanged)) {
                *result = 6;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<BlockchainStatusModel97d618 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< qint32*>(_v) = _t->numberOfBlocks(); break;
        case 1: *reinterpret_cast< QDateTime*>(_v) = _t->timestampLastBlock(); break;
        case 2: *reinterpret_cast< QString*>(_v) = _t->hashLastBlock(); break;
        case 3: *reinterpret_cast< QString*>(_v) = _t->currentSkySupply(); break;
        case 4: *reinterpret_cast< QString*>(_v) = _t->totalSkySupply(); break;
        case 5: *reinterpret_cast< QString*>(_v) = _t->currentCoinHoursSupply(); break;
        case 6: *reinterpret_cast< QString*>(_v) = _t->totalCoinHoursSupply(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<BlockchainStatusModel97d618 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setNumberOfBlocks(*reinterpret_cast< qint32*>(_v)); break;
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

QT_INIT_METAOBJECT const QMetaObject BlockchainStatusModel97d618::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_BlockchainStatusModel97d618.data,
    qt_meta_data_BlockchainStatusModel97d618,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *BlockchainStatusModel97d618::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *BlockchainStatusModel97d618::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_BlockchainStatusModel97d618.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int BlockchainStatusModel97d618::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 7)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 7;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 7)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 7;
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
void BlockchainStatusModel97d618::numberOfBlocksChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void BlockchainStatusModel97d618::timestampLastBlockChanged(QDateTime _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void BlockchainStatusModel97d618::hashLastBlockChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void BlockchainStatusModel97d618::currentSkySupplyChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void BlockchainStatusModel97d618::totalSkySupplyChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void BlockchainStatusModel97d618::currentCoinHoursSupplyChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void BlockchainStatusModel97d618::totalCoinHoursSupplyChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
