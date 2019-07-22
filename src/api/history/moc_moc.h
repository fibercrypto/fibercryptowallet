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
struct qt_meta_stringdata_TransactionDetails8ba275_t {
    QByteArrayData data[21];
    char stringdata0[276];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_TransactionDetails8ba275_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_TransactionDetails8ba275_t qt_meta_stringdata_TransactionDetails8ba275 = {
    {
QT_MOC_LITERAL(0, 0, 24), // "TransactionDetails8ba275"
QT_MOC_LITERAL(1, 25, 11), // "dateChanged"
QT_MOC_LITERAL(2, 37, 0), // ""
QT_MOC_LITERAL(3, 38, 4), // "date"
QT_MOC_LITERAL(4, 43, 13), // "statusChanged"
QT_MOC_LITERAL(5, 57, 6), // "status"
QT_MOC_LITERAL(6, 64, 11), // "typeChanged"
QT_MOC_LITERAL(7, 76, 2), // "ty"
QT_MOC_LITERAL(8, 79, 13), // "amountChanged"
QT_MOC_LITERAL(9, 93, 6), // "amount"
QT_MOC_LITERAL(10, 100, 20), // "hoursReceivedChanged"
QT_MOC_LITERAL(11, 121, 13), // "hoursReceived"
QT_MOC_LITERAL(12, 135, 18), // "hoursBurnedChanged"
QT_MOC_LITERAL(13, 154, 11), // "hoursBurned"
QT_MOC_LITERAL(14, 166, 20), // "transactionIDChanged"
QT_MOC_LITERAL(15, 187, 13), // "transactionID"
QT_MOC_LITERAL(16, 201, 18), // "sentAddressChanged"
QT_MOC_LITERAL(17, 220, 11), // "sentAddress"
QT_MOC_LITERAL(18, 232, 22), // "receivedAddressChanged"
QT_MOC_LITERAL(19, 255, 15), // "receivedAddress"
QT_MOC_LITERAL(20, 271, 4) // "type"

    },
    "TransactionDetails8ba275\0dateChanged\0"
    "\0date\0statusChanged\0status\0typeChanged\0"
    "ty\0amountChanged\0amount\0hoursReceivedChanged\0"
    "hoursReceived\0hoursBurnedChanged\0"
    "hoursBurned\0transactionIDChanged\0"
    "transactionID\0sentAddressChanged\0"
    "sentAddress\0receivedAddressChanged\0"
    "receivedAddress\0type"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_TransactionDetails8ba275[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       9,   14, // methods
       9,   86, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       9,       // signalCount

 // signals: name, argc, parameters, tag, flags
       1,    1,   59,    2, 0x06 /* Public */,
       4,    1,   62,    2, 0x06 /* Public */,
       6,    1,   65,    2, 0x06 /* Public */,
       8,    1,   68,    2, 0x06 /* Public */,
      10,    1,   71,    2, 0x06 /* Public */,
      12,    1,   74,    2, 0x06 /* Public */,
      14,    1,   77,    2, 0x06 /* Public */,
      16,    1,   80,    2, 0x06 /* Public */,
      18,    1,   83,    2, 0x06 /* Public */,

 // signals: parameters
    QMetaType::Void, QMetaType::QDateTime,    3,
    QMetaType::Void, QMetaType::Int,    5,
    QMetaType::Void, QMetaType::Int,    7,
    QMetaType::Void, QMetaType::Int,    9,
    QMetaType::Void, QMetaType::Int,   11,
    QMetaType::Void, QMetaType::Int,   13,
    QMetaType::Void, QMetaType::QString,   15,
    QMetaType::Void, QMetaType::QString,   17,
    QMetaType::Void, QMetaType::QString,   19,

 // properties: name, type, flags
       3, QMetaType::QDateTime, 0x00495103,
       5, QMetaType::Int, 0x00495103,
      20, QMetaType::Int, 0x00495103,
       9, QMetaType::Int, 0x00495103,
      11, QMetaType::Int, 0x00495103,
      13, QMetaType::Int, 0x00495103,
      15, QMetaType::QString, 0x00495103,
      17, QMetaType::QString, 0x00495103,
      19, QMetaType::QString, 0x00495103,

 // properties: notify_signal_id
       0,
       1,
       2,
       3,
       4,
       5,
       6,
       7,
       8,

       0        // eod
};

void TransactionDetails8ba275::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<TransactionDetails8ba275 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->dateChanged((*reinterpret_cast< QDateTime(*)>(_a[1]))); break;
        case 1: _t->statusChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 2: _t->typeChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 3: _t->amountChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 4: _t->hoursReceivedChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 5: _t->hoursBurnedChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 6: _t->transactionIDChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 7: _t->sentAddressChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        case 8: _t->receivedAddressChanged((*reinterpret_cast< QString(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (TransactionDetails8ba275::*)(QDateTime );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionDetails8ba275::dateChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (TransactionDetails8ba275::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionDetails8ba275::statusChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (TransactionDetails8ba275::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionDetails8ba275::typeChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (TransactionDetails8ba275::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionDetails8ba275::amountChanged)) {
                *result = 3;
                return;
            }
        }
        {
            using _t = void (TransactionDetails8ba275::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionDetails8ba275::hoursReceivedChanged)) {
                *result = 4;
                return;
            }
        }
        {
            using _t = void (TransactionDetails8ba275::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionDetails8ba275::hoursBurnedChanged)) {
                *result = 5;
                return;
            }
        }
        {
            using _t = void (TransactionDetails8ba275::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionDetails8ba275::transactionIDChanged)) {
                *result = 6;
                return;
            }
        }
        {
            using _t = void (TransactionDetails8ba275::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionDetails8ba275::sentAddressChanged)) {
                *result = 7;
                return;
            }
        }
        {
            using _t = void (TransactionDetails8ba275::*)(QString );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&TransactionDetails8ba275::receivedAddressChanged)) {
                *result = 8;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<TransactionDetails8ba275 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< QDateTime*>(_v) = _t->date(); break;
        case 1: *reinterpret_cast< qint32*>(_v) = _t->status(); break;
        case 2: *reinterpret_cast< qint32*>(_v) = _t->type(); break;
        case 3: *reinterpret_cast< qint32*>(_v) = _t->amount(); break;
        case 4: *reinterpret_cast< qint32*>(_v) = _t->hoursReceived(); break;
        case 5: *reinterpret_cast< qint32*>(_v) = _t->hoursBurned(); break;
        case 6: *reinterpret_cast< QString*>(_v) = _t->transactionID(); break;
        case 7: *reinterpret_cast< QString*>(_v) = _t->sentAddress(); break;
        case 8: *reinterpret_cast< QString*>(_v) = _t->receivedAddress(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<TransactionDetails8ba275 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setDate(*reinterpret_cast< QDateTime*>(_v)); break;
        case 1: _t->setStatus(*reinterpret_cast< qint32*>(_v)); break;
        case 2: _t->setType(*reinterpret_cast< qint32*>(_v)); break;
        case 3: _t->setAmount(*reinterpret_cast< qint32*>(_v)); break;
        case 4: _t->setHoursReceived(*reinterpret_cast< qint32*>(_v)); break;
        case 5: _t->setHoursBurned(*reinterpret_cast< qint32*>(_v)); break;
        case 6: _t->setTransactionID(*reinterpret_cast< QString*>(_v)); break;
        case 7: _t->setSentAddress(*reinterpret_cast< QString*>(_v)); break;
        case 8: _t->setReceivedAddress(*reinterpret_cast< QString*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject TransactionDetails8ba275::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_TransactionDetails8ba275.data,
    qt_meta_data_TransactionDetails8ba275,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *TransactionDetails8ba275::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *TransactionDetails8ba275::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_TransactionDetails8ba275.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int TransactionDetails8ba275::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
#ifndef QT_NO_PROPERTIES
   else if (_c == QMetaObject::ReadProperty || _c == QMetaObject::WriteProperty
            || _c == QMetaObject::ResetProperty || _c == QMetaObject::RegisterPropertyMetaType) {
        qt_static_metacall(this, _c, _id, _a);
        _id -= 9;
    } else if (_c == QMetaObject::QueryPropertyDesignable) {
        _id -= 9;
    } else if (_c == QMetaObject::QueryPropertyScriptable) {
        _id -= 9;
    } else if (_c == QMetaObject::QueryPropertyStored) {
        _id -= 9;
    } else if (_c == QMetaObject::QueryPropertyEditable) {
        _id -= 9;
    } else if (_c == QMetaObject::QueryPropertyUser) {
        _id -= 9;
    }
#endif // QT_NO_PROPERTIES
    return _id;
}

// SIGNAL 0
void TransactionDetails8ba275::dateChanged(QDateTime _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void TransactionDetails8ba275::statusChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void TransactionDetails8ba275::typeChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void TransactionDetails8ba275::amountChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}

// SIGNAL 4
void TransactionDetails8ba275::hoursReceivedChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 4, _a);
}

// SIGNAL 5
void TransactionDetails8ba275::hoursBurnedChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 5, _a);
}

// SIGNAL 6
void TransactionDetails8ba275::transactionIDChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 6, _a);
}

// SIGNAL 7
void TransactionDetails8ba275::sentAddressChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 7, _a);
}

// SIGNAL 8
void TransactionDetails8ba275::receivedAddressChanged(QString _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 8, _a);
}
struct qt_meta_stringdata_AddressDeatails8ba275_t {
    QByteArrayData data[8];
    char stringdata0[116];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_AddressDeatails8ba275_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_AddressDeatails8ba275_t qt_meta_stringdata_AddressDeatails8ba275 = {
    {
QT_MOC_LITERAL(0, 0, 21), // "AddressDeatails8ba275"
QT_MOC_LITERAL(1, 22, 14), // "addressChanged"
QT_MOC_LITERAL(2, 37, 0), // ""
QT_MOC_LITERAL(3, 38, 7), // "address"
QT_MOC_LITERAL(4, 46, 17), // "addressSkyChanged"
QT_MOC_LITERAL(5, 64, 10), // "addressSky"
QT_MOC_LITERAL(6, 75, 23), // "addressCoinHoursChanged"
QT_MOC_LITERAL(7, 99, 16) // "addressCoinHours"

    },
    "AddressDeatails8ba275\0addressChanged\0"
    "\0address\0addressSkyChanged\0addressSky\0"
    "addressCoinHoursChanged\0addressCoinHours"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_AddressDeatails8ba275[] = {

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
    QMetaType::Void, QMetaType::Float,    3,
    QMetaType::Void, QMetaType::Int,    5,
    QMetaType::Void, QMetaType::Int,    7,

 // properties: name, type, flags
       3, QMetaType::Float, 0x00495103,
       5, QMetaType::Int, 0x00495103,
       7, QMetaType::Int, 0x00495103,

 // properties: notify_signal_id
       0,
       1,
       2,

       0        // eod
};

void AddressDeatails8ba275::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<AddressDeatails8ba275 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->addressChanged((*reinterpret_cast< float(*)>(_a[1]))); break;
        case 1: _t->addressSkyChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 2: _t->addressCoinHoursChanged((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (AddressDeatails8ba275::*)(float );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressDeatails8ba275::addressChanged)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (AddressDeatails8ba275::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressDeatails8ba275::addressSkyChanged)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (AddressDeatails8ba275::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressDeatails8ba275::addressCoinHoursChanged)) {
                *result = 2;
                return;
            }
        }
    }
#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<AddressDeatails8ba275 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< float*>(_v) = _t->address(); break;
        case 1: *reinterpret_cast< qint32*>(_v) = _t->addressSky(); break;
        case 2: *reinterpret_cast< qint32*>(_v) = _t->addressCoinHours(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<AddressDeatails8ba275 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setAddress(*reinterpret_cast< float*>(_v)); break;
        case 1: _t->setAddressSky(*reinterpret_cast< qint32*>(_v)); break;
        case 2: _t->setAddressCoinHours(*reinterpret_cast< qint32*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject AddressDeatails8ba275::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_AddressDeatails8ba275.data,
    qt_meta_data_AddressDeatails8ba275,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *AddressDeatails8ba275::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *AddressDeatails8ba275::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_AddressDeatails8ba275.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int AddressDeatails8ba275::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
void AddressDeatails8ba275::addressChanged(float _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void AddressDeatails8ba275::addressSkyChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void AddressDeatails8ba275::addressCoinHoursChanged(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}
struct qt_meta_stringdata_AddressList8ba275_t {
    QByteArrayData data[13];
    char stringdata0[180];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_AddressList8ba275_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_AddressList8ba275_t qt_meta_stringdata_AddressList8ba275 = {
    {
QT_MOC_LITERAL(0, 0, 17), // "AddressList8ba275"
QT_MOC_LITERAL(1, 18, 14), // "addTransaction"
QT_MOC_LITERAL(2, 33, 0), // ""
QT_MOC_LITERAL(3, 34, 22), // "AddressDeatails8ba275*"
QT_MOC_LITERAL(4, 57, 11), // "transaction"
QT_MOC_LITERAL(5, 69, 17), // "removeTransaction"
QT_MOC_LITERAL(6, 87, 5), // "index"
QT_MOC_LITERAL(7, 93, 12), // "rolesChanged"
QT_MOC_LITERAL(8, 106, 10), // "type378cdd"
QT_MOC_LITERAL(9, 117, 5), // "roles"
QT_MOC_LITERAL(10, 123, 16), // "addressesChanged"
QT_MOC_LITERAL(11, 140, 29), // "QList<AddressDeatails8ba275*>"
QT_MOC_LITERAL(12, 170, 9) // "addresses"

    },
    "AddressList8ba275\0addTransaction\0\0"
    "AddressDeatails8ba275*\0transaction\0"
    "removeTransaction\0index\0rolesChanged\0"
    "type378cdd\0roles\0addressesChanged\0"
    "QList<AddressDeatails8ba275*>\0addresses"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_AddressList8ba275[] = {

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

void AddressList8ba275::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<AddressList8ba275 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->addTransaction((*reinterpret_cast< AddressDeatails8ba275*(*)>(_a[1]))); break;
        case 1: _t->removeTransaction((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 2: _t->rolesChanged((*reinterpret_cast< type378cdd(*)>(_a[1]))); break;
        case 3: _t->addressesChanged((*reinterpret_cast< QList<AddressDeatails8ba275*>(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 0:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< AddressDeatails8ba275* >(); break;
            }
            break;
        case 3:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<AddressDeatails8ba275*> >(); break;
            }
            break;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (AddressList8ba275::*)(AddressDeatails8ba275 * );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressList8ba275::addTransaction)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (AddressList8ba275::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressList8ba275::removeTransaction)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (AddressList8ba275::*)(type378cdd );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressList8ba275::rolesChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (AddressList8ba275::*)(QList<AddressDeatails8ba275*> );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&AddressList8ba275::addressesChanged)) {
                *result = 3;
                return;
            }
        }
    } else if (_c == QMetaObject::RegisterPropertyMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<AddressDeatails8ba275*> >(); break;
        }
    }

#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<AddressList8ba275 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< type378cdd*>(_v) = _t->roles(); break;
        case 1: *reinterpret_cast< QList<AddressDeatails8ba275*>*>(_v) = _t->addresses(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<AddressList8ba275 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRoles(*reinterpret_cast< type378cdd*>(_v)); break;
        case 1: _t->setAddresses(*reinterpret_cast< QList<AddressDeatails8ba275*>*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject AddressList8ba275::staticMetaObject = { {
    &QAbstractListModel::staticMetaObject,
    qt_meta_stringdata_AddressList8ba275.data,
    qt_meta_data_AddressList8ba275,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *AddressList8ba275::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *AddressList8ba275::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_AddressList8ba275.stringdata0))
        return static_cast<void*>(this);
    return QAbstractListModel::qt_metacast(_clname);
}

int AddressList8ba275::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
void AddressList8ba275::addTransaction(AddressDeatails8ba275 * _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void AddressList8ba275::removeTransaction(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void AddressList8ba275::rolesChanged(type378cdd _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void AddressList8ba275::addressesChanged(QList<AddressDeatails8ba275*> _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}
struct qt_meta_stringdata_HistoryModel8ba275_t {
    QByteArrayData data[13];
    char stringdata0[193];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_HistoryModel8ba275_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_HistoryModel8ba275_t qt_meta_stringdata_HistoryModel8ba275 = {
    {
QT_MOC_LITERAL(0, 0, 18), // "HistoryModel8ba275"
QT_MOC_LITERAL(1, 19, 14), // "addTransaction"
QT_MOC_LITERAL(2, 34, 0), // ""
QT_MOC_LITERAL(3, 35, 25), // "TransactionDetails8ba275*"
QT_MOC_LITERAL(4, 61, 11), // "transaction"
QT_MOC_LITERAL(5, 73, 17), // "removeTransaction"
QT_MOC_LITERAL(6, 91, 5), // "index"
QT_MOC_LITERAL(7, 97, 12), // "rolesChanged"
QT_MOC_LITERAL(8, 110, 10), // "type378cdd"
QT_MOC_LITERAL(9, 121, 5), // "roles"
QT_MOC_LITERAL(10, 127, 19), // "transactionsChanged"
QT_MOC_LITERAL(11, 147, 32), // "QList<TransactionDetails8ba275*>"
QT_MOC_LITERAL(12, 180, 12) // "transactions"

    },
    "HistoryModel8ba275\0addTransaction\0\0"
    "TransactionDetails8ba275*\0transaction\0"
    "removeTransaction\0index\0rolesChanged\0"
    "type378cdd\0roles\0transactionsChanged\0"
    "QList<TransactionDetails8ba275*>\0"
    "transactions"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_HistoryModel8ba275[] = {

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

void HistoryModel8ba275::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<HistoryModel8ba275 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->addTransaction((*reinterpret_cast< TransactionDetails8ba275*(*)>(_a[1]))); break;
        case 1: _t->removeTransaction((*reinterpret_cast< qint32(*)>(_a[1]))); break;
        case 2: _t->rolesChanged((*reinterpret_cast< type378cdd(*)>(_a[1]))); break;
        case 3: _t->transactionsChanged((*reinterpret_cast< QList<TransactionDetails8ba275*>(*)>(_a[1]))); break;
        default: ;
        }
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 0:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< TransactionDetails8ba275* >(); break;
            }
            break;
        case 3:
            switch (*reinterpret_cast<int*>(_a[1])) {
            default: *reinterpret_cast<int*>(_a[0]) = -1; break;
            case 0:
                *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<TransactionDetails8ba275*> >(); break;
            }
            break;
        }
    } else if (_c == QMetaObject::IndexOfMethod) {
        int *result = reinterpret_cast<int *>(_a[0]);
        {
            using _t = void (HistoryModel8ba275::*)(TransactionDetails8ba275 * );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&HistoryModel8ba275::addTransaction)) {
                *result = 0;
                return;
            }
        }
        {
            using _t = void (HistoryModel8ba275::*)(qint32 );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&HistoryModel8ba275::removeTransaction)) {
                *result = 1;
                return;
            }
        }
        {
            using _t = void (HistoryModel8ba275::*)(type378cdd );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&HistoryModel8ba275::rolesChanged)) {
                *result = 2;
                return;
            }
        }
        {
            using _t = void (HistoryModel8ba275::*)(QList<TransactionDetails8ba275*> );
            if (*reinterpret_cast<_t *>(_a[1]) == static_cast<_t>(&HistoryModel8ba275::transactionsChanged)) {
                *result = 3;
                return;
            }
        }
    } else if (_c == QMetaObject::RegisterPropertyMetaType) {
        switch (_id) {
        default: *reinterpret_cast<int*>(_a[0]) = -1; break;
        case 1:
            *reinterpret_cast<int*>(_a[0]) = qRegisterMetaType< QList<TransactionDetails8ba275*> >(); break;
        }
    }

#ifndef QT_NO_PROPERTIES
    else if (_c == QMetaObject::ReadProperty) {
        auto *_t = static_cast<HistoryModel8ba275 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: *reinterpret_cast< type378cdd*>(_v) = _t->roles(); break;
        case 1: *reinterpret_cast< QList<TransactionDetails8ba275*>*>(_v) = _t->transactions(); break;
        default: break;
        }
    } else if (_c == QMetaObject::WriteProperty) {
        auto *_t = static_cast<HistoryModel8ba275 *>(_o);
        Q_UNUSED(_t)
        void *_v = _a[0];
        switch (_id) {
        case 0: _t->setRoles(*reinterpret_cast< type378cdd*>(_v)); break;
        case 1: _t->setTransactions(*reinterpret_cast< QList<TransactionDetails8ba275*>*>(_v)); break;
        default: break;
        }
    } else if (_c == QMetaObject::ResetProperty) {
    }
#endif // QT_NO_PROPERTIES
}

QT_INIT_METAOBJECT const QMetaObject HistoryModel8ba275::staticMetaObject = { {
    &QAbstractListModel::staticMetaObject,
    qt_meta_stringdata_HistoryModel8ba275.data,
    qt_meta_data_HistoryModel8ba275,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *HistoryModel8ba275::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *HistoryModel8ba275::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_HistoryModel8ba275.stringdata0))
        return static_cast<void*>(this);
    return QAbstractListModel::qt_metacast(_clname);
}

int HistoryModel8ba275::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
void HistoryModel8ba275::addTransaction(TransactionDetails8ba275 * _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 0, _a);
}

// SIGNAL 1
void HistoryModel8ba275::removeTransaction(qint32 _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 1, _a);
}

// SIGNAL 2
void HistoryModel8ba275::rolesChanged(type378cdd _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 2, _a);
}

// SIGNAL 3
void HistoryModel8ba275::transactionsChanged(QList<TransactionDetails8ba275*> _t1)
{
    void *_a[] = { nullptr, const_cast<void*>(reinterpret_cast<const void*>(std::addressof(_t1))) };
    QMetaObject::activate(this, &staticMetaObject, 3, _a);
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
