

#pragma once

#ifndef GO_MOC_64bdd5_H
#define GO_MOC_64bdd5_H

#include <stdint.h>

#ifdef __cplusplus
class WalletManager64bdd5;
void WalletManager64bdd5_WalletManager64bdd5_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
void* WalletManager64bdd5_CreateEncryptedWallet(void* ptr, struct Moc_PackedString seed, struct Moc_PackedString label, struct Moc_PackedString password, int scanN);
void* WalletManager64bdd5_CreateUnencryptedWallet(void* ptr, struct Moc_PackedString seed, struct Moc_PackedString label, int scanN);
struct Moc_PackedString WalletManager64bdd5_GetNewSeed(void* ptr, int entropy);
int WalletManager64bdd5_VerifySeed(void* ptr, struct Moc_PackedString seed);
void WalletManager64bdd5_NewWalletAddress(void* ptr, struct Moc_PackedString id, int n, struct Moc_PackedString password);
void WalletManager64bdd5_EncryptWallet(void* ptr, struct Moc_PackedString id, struct Moc_PackedString password);
void WalletManager64bdd5_DecryptWallet(void* ptr, struct Moc_PackedString id, struct Moc_PackedString password);
struct Moc_PackedList WalletManager64bdd5_GetWallets(void* ptr);
struct Moc_PackedList WalletManager64bdd5_GetAddresses(void* ptr, struct Moc_PackedString id);
int WalletManager64bdd5_WalletManager64bdd5_QRegisterMetaType();
int WalletManager64bdd5_WalletManager64bdd5_QRegisterMetaType2(char* typeName);
int WalletManager64bdd5_WalletManager64bdd5_QmlRegisterType();
int WalletManager64bdd5_WalletManager64bdd5_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* WalletManager64bdd5___children_atList(void* ptr, int i);
void WalletManager64bdd5___children_setList(void* ptr, void* i);
void* WalletManager64bdd5___children_newList(void* ptr);
void* WalletManager64bdd5___dynamicPropertyNames_atList(void* ptr, int i);
void WalletManager64bdd5___dynamicPropertyNames_setList(void* ptr, void* i);
void* WalletManager64bdd5___dynamicPropertyNames_newList(void* ptr);
void* WalletManager64bdd5___findChildren_atList(void* ptr, int i);
void WalletManager64bdd5___findChildren_setList(void* ptr, void* i);
void* WalletManager64bdd5___findChildren_newList(void* ptr);
void* WalletManager64bdd5___findChildren_atList3(void* ptr, int i);
void WalletManager64bdd5___findChildren_setList3(void* ptr, void* i);
void* WalletManager64bdd5___findChildren_newList3(void* ptr);
void* WalletManager64bdd5___qFindChildren_atList2(void* ptr, int i);
void WalletManager64bdd5___qFindChildren_setList2(void* ptr, void* i);
void* WalletManager64bdd5___qFindChildren_newList2(void* ptr);
void* WalletManager64bdd5___getWallets_atList(void* ptr, int i);
void WalletManager64bdd5___getWallets_setList(void* ptr, void* i);
void* WalletManager64bdd5___getWallets_newList(void* ptr);
void* WalletManager64bdd5___getAddresses_atList(void* ptr, int i);
void WalletManager64bdd5___getAddresses_setList(void* ptr, void* i);
void* WalletManager64bdd5___getAddresses_newList(void* ptr);
void* WalletManager64bdd5_NewWalletManager(void* parent);
void WalletManager64bdd5_DestroyWalletManager(void* ptr);
void WalletManager64bdd5_DestroyWalletManagerDefault(void* ptr);
void WalletManager64bdd5_ChildEventDefault(void* ptr, void* event);
void WalletManager64bdd5_ConnectNotifyDefault(void* ptr, void* sign);
void WalletManager64bdd5_CustomEventDefault(void* ptr, void* event);
void WalletManager64bdd5_DeleteLaterDefault(void* ptr);
void WalletManager64bdd5_DisconnectNotifyDefault(void* ptr, void* sign);
char WalletManager64bdd5_EventDefault(void* ptr, void* e);
char WalletManager64bdd5_EventFilterDefault(void* ptr, void* watched, void* event);
void WalletManager64bdd5_TimerEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif

#endif