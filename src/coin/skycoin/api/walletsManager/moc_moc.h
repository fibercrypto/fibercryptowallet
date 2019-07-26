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
struct qt_meta_stringdata_WalletManager64bdd5_t {
    QByteArrayData data[15];
    char stringdata0[169];
};
#define QT_MOC_LITERAL(idx, ofs, len) \
    Q_STATIC_BYTE_ARRAY_DATA_HEADER_INITIALIZER_WITH_OFFSET(len, \
    qptrdiff(offsetof(qt_meta_stringdata_WalletManager64bdd5_t, stringdata0) + ofs \
        - idx * sizeof(QByteArrayData)) \
    )
static const qt_meta_stringdata_WalletManager64bdd5_t qt_meta_stringdata_WalletManager64bdd5 = {
    {
QT_MOC_LITERAL(0, 0, 19), // "WalletManager64bdd5"
QT_MOC_LITERAL(1, 20, 21), // "createEncryptedWallet"
QT_MOC_LITERAL(2, 42, 14), // "QWallet445aa6*"
QT_MOC_LITERAL(3, 57, 0), // ""
QT_MOC_LITERAL(4, 58, 4), // "seed"
QT_MOC_LITERAL(5, 63, 5), // "label"
QT_MOC_LITERAL(6, 69, 8), // "password"
QT_MOC_LITERAL(7, 78, 5), // "scanN"
QT_MOC_LITERAL(8, 84, 23), // "createUnencryptedWallet"
QT_MOC_LITERAL(9, 108, 10), // "getNewSeed"
QT_MOC_LITERAL(10, 119, 7), // "entropy"
QT_MOC_LITERAL(11, 127, 10), // "verifySeed"
QT_MOC_LITERAL(12, 138, 13), // "encryptWallet"
QT_MOC_LITERAL(13, 152, 2), // "id"
QT_MOC_LITERAL(14, 155, 13) // "decryptWallet"

    },
    "WalletManager64bdd5\0createEncryptedWallet\0"
    "QWallet445aa6*\0\0seed\0label\0password\0"
    "scanN\0createUnencryptedWallet\0getNewSeed\0"
    "entropy\0verifySeed\0encryptWallet\0id\0"
    "decryptWallet"
};
#undef QT_MOC_LITERAL

static const uint qt_meta_data_WalletManager64bdd5[] = {

 // content:
       8,       // revision
       0,       // classname
       0,    0, // classinfo
       6,   14, // methods
       0,    0, // properties
       0,    0, // enums/sets
       0,    0, // constructors
       0,       // flags
       0,       // signalCount

 // slots: name, argc, parameters, tag, flags
       1,    4,   44,    3, 0x0a /* Public */,
       8,    3,   53,    3, 0x0a /* Public */,
       9,    1,   60,    3, 0x0a /* Public */,
      11,    1,   63,    3, 0x0a /* Public */,
      12,    2,   66,    3, 0x0a /* Public */,
      14,    2,   71,    3, 0x0a /* Public */,

 // slots: parameters
    0x80000000 | 2, QMetaType::QString, QMetaType::QString, QMetaType::QString, QMetaType::Int,    4,    5,    6,    7,
    0x80000000 | 2, QMetaType::QString, QMetaType::QString, QMetaType::Int,    4,    5,    7,
    QMetaType::QString, QMetaType::Int,   10,
    QMetaType::Int, QMetaType::QString,    4,
    QMetaType::Void, QMetaType::QString, QMetaType::QString,   13,    6,
    QMetaType::Void, QMetaType::QString, QMetaType::QString,   13,    6,

       0        // eod
};

void WalletManager64bdd5::qt_static_metacall(QObject *_o, QMetaObject::Call _c, int _id, void **_a)
{
    if (_c == QMetaObject::InvokeMetaMethod) {
        auto *_t = static_cast<WalletManager64bdd5 *>(_o);
        Q_UNUSED(_t)
        switch (_id) {
        case 0: { QWallet445aa6* _r = _t->createEncryptedWallet((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< QString(*)>(_a[3])),(*reinterpret_cast< qint32(*)>(_a[4])));
            if (_a[0]) *reinterpret_cast< QWallet445aa6**>(_a[0]) = std::move(_r); }  break;
        case 1: { QWallet445aa6* _r = _t->createUnencryptedWallet((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2])),(*reinterpret_cast< qint32(*)>(_a[3])));
            if (_a[0]) *reinterpret_cast< QWallet445aa6**>(_a[0]) = std::move(_r); }  break;
        case 2: { QString _r = _t->getNewSeed((*reinterpret_cast< qint32(*)>(_a[1])));
            if (_a[0]) *reinterpret_cast< QString*>(_a[0]) = std::move(_r); }  break;
        case 3: { qint32 _r = _t->verifySeed((*reinterpret_cast< QString(*)>(_a[1])));
            if (_a[0]) *reinterpret_cast< qint32*>(_a[0]) = std::move(_r); }  break;
        case 4: _t->encryptWallet((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2]))); break;
        case 5: _t->decryptWallet((*reinterpret_cast< QString(*)>(_a[1])),(*reinterpret_cast< QString(*)>(_a[2]))); break;
        default: ;
        }
    }
}

QT_INIT_METAOBJECT const QMetaObject WalletManager64bdd5::staticMetaObject = { {
    &QObject::staticMetaObject,
    qt_meta_stringdata_WalletManager64bdd5.data,
    qt_meta_data_WalletManager64bdd5,
    qt_static_metacall,
    nullptr,
    nullptr
} };


const QMetaObject *WalletManager64bdd5::metaObject() const
{
    return QObject::d_ptr->metaObject ? QObject::d_ptr->dynamicMetaObject() : &staticMetaObject;
}

void *WalletManager64bdd5::qt_metacast(const char *_clname)
{
    if (!_clname) return nullptr;
    if (!strcmp(_clname, qt_meta_stringdata_WalletManager64bdd5.stringdata0))
        return static_cast<void*>(this);
    return QObject::qt_metacast(_clname);
}

int WalletManager64bdd5::qt_metacall(QMetaObject::Call _c, int _id, void **_a)
{
    _id = QObject::qt_metacall(_c, _id, _a);
    if (_id < 0)
        return _id;
    if (_c == QMetaObject::InvokeMetaMethod) {
        if (_id < 6)
            qt_static_metacall(this, _c, _id, _a);
        _id -= 6;
    } else if (_c == QMetaObject::RegisterMethodArgumentMetaType) {
        if (_id < 6)
            *reinterpret_cast<int*>(_a[0]) = -1;
        _id -= 6;
    }
    return _id;
}
QT_WARNING_POP
QT_END_MOC_NAMESPACE
