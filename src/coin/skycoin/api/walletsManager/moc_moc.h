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
struct qt_meta_stringdata_WalletManager3cf9d2_t {
    QByteArrayData data[20];
    char stringdata0[231];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_WalletManager3cf9d2_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_WalletManager3cf9d2_t qt_meta_stringdata_WalletManager3cf9d2 = {
    {
QT_MOC_LITERAL(0, 0, 19), // "WalletManager3cf9d2"
QT_MOC_LITERAL(1, 20, 21), // "createEncryptedWallet"
QT_MOC_LITERAL(2, 42, 0), // ""
QT_MOC_LITERAL(3, 43, 18), // "WalletModelbaab1e*"
QT_MOC_LITERAL(4, 62, 7), // "walletM"
QT_MOC_LITERAL(5, 70, 4), // "seed"
QT_MOC_LITERAL(6, 75, 5), // "label"
QT_MOC_LITERAL(7, 81, 8), // "password"
QT_MOC_LITERAL(8, 90, 5), // "scanN"
QT_MOC_LITERAL(9, 96, 23), // "createUnencryptedWallet"
QT_MOC_LITERAL(10, 120, 10), // "getNewSeed"
QT_MOC_LITERAL(11, 131, 7), // "entropy"
QT_MOC_LITERAL(12, 139, 10), // "verifySeed"
QT_MOC_LITERAL(13, 150, 16), // "newWalletAddress"
QT_MOC_LITERAL(14, 167, 21), // "AddressesModelbaab1e*"
QT_MOC_LITERAL(15, 189, 8), // "addressM"
QT_MOC_LITERAL(16, 198, 2), // "id"
QT_MOC_LITERAL(17, 201, 1), // "n"
QT_MOC_LITERAL(18, 203, 13), // "encryptWallet"
QT_MOC_LITERAL(19, 217, 13) // "decryptWallet"

    },
    "WalletManager3cf9d2\0createEncryptedWallet\0"
    "\0WalletModelbaab1e*\0walletM\0seed\0label\0"
    "password\0scanN\0createUnencryptedWallet\0"
    "getNewSeed\0entropy\0verifySeed\0"
    "newWalletAddress\0AddressesModelbaab1e*\0"
    "addressM\0id\0n\0encryptWallet\0decryptWallet"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_WalletManager3cf9d2[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       7,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       0,       // signalCount

 // slots: name, argc, parameters, tag, flags
       1,    5,   49,    2, 0x0a /* Public */,
       9,    4,   60,    2, 0x0a /* Public */,
      10,    1,   69,    2, 0x0a /* Public */,
      12,    1,   72,    2, 0x0a /* Public */,
      13,    4,   75,    2, 0x0a /* Public */,
      18,    3,   84,    2, 0x0a /* Public */,
      19,    3,   91,    2, 0x0a /* Public */,

 // slots: parameters
    QMetaType::Void, 0x80000000 | 3, QMetaType::QString, QMetaType::QString, QMetaType::QString, QMetaType::Int,    4,    5,    6,    7,    8,
    QMetaType::Void, 0x80000000 | 3, QMetaType::QString, QMetaType::QString, QMetaType::Int,    4,    5,    6,    8,
    QMetaType::QString, QMetaType::Int,   11,
    QMetaType::Int, QMetaType::QString,    5,
    QMetaType::Void, 0x80000000 | 14, QMetaType::QString, QMetaType::Int, QMetaType::QString,   15,   16,   17,    7,
    QMetaType::Void, 0x80000000 | 3, QMetaType::QString, QMetaType::QString,    4,   16,    7,
    QMetaType::Void, 0x80000000 | 3, QMetaType::QString, QMetaType::QString,    4,   16,    7,

       0        // eod
};

void WalletManager3cf9d2::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<WalletManager3cf9d2 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: _t->createEncryptedWallet((*reinterpret_cast< WalletModelbaab1e*(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< QString(*)>(_a[3])),(*reinterpret_cast< QString(*)>(_a[4])),(*reinterpret_cast< qint32(*)>(_a[5]))); break;
        case 1: _t->createUnencryptedWallet((*reinterpret_cast< WalletModelbaab1e*(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< QString(*)>(_a[3])),(*reinterpret_cast< qint32(*)>(_a[4]))); break;
        case 2: { QString _r = _t->getNewSeed((*reinterpret_cast< qint32(*)>(_a[1])));
            if (_a[0]) *reinterpret_cast< QString*>(_a[0]) = std::move(_r); }  break;
        case 3: { qint32 _r = _t->verifySeed((*reinterpret_cast< QString(*)>(_a[1])));
            if (_a[0]) *reinterpret_cast< qint32*>(_a[0]) = std::move(_r); }  break;
        case 4: _t->newWalletAddress((*reinterpret_cast< AddressesModelbaab1e*(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< qint32(*)>(_a[3])),(*reinterpret_cast< QString(*)>(_a[4]))); break;
        case 5: _t->encryptWallet((*reinterpret_cast< WalletModelbaab1e*(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< QString(*)>(_a[3]))); break;
        case 6: _t->decryptWallet((*reinterpret_cast< WalletModelbaab1e*(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< QString(*)>(_a[3]))); break;
        default: ;
        }
    }
}

QT_INIT_METAOBJECT const QMetaObject WalletManager3cf9d2::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_WalletManager3cf9d2.data,
    qt_meta_data_WalletManager3cf9d2,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *WalletManager3cf9d2::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *WalletManager3cf9d2::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_WalletManager3cf9d2.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int WalletManager3cf9d2::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
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
    return _id;
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
