

#pragma once

#ifndef GO_MOC_3cf9d2_H
#define GO_MOC_3cf9d2_H

#include <stdint.h>

#ifdef __cplusplus
class WalletManager3cf9d2;
void WalletManager3cf9d2_WalletManager3cf9d2_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
void WalletManager3cf9d2_CreateEncryptedWallet(void* ptr, void* walletM, struct Moc_PackedString seed, struct Moc_PackedString label, struct Moc_PackedString password, int scanN);
void WalletManager3cf9d2_CreateUnencryptedWallet(void* ptr, void* walletM, struct Moc_PackedString seed, struct Moc_PackedString label, int scanN);
struct Moc_PackedString WalletManager3cf9d2_GetNewSeed(void* ptr, int entropy);
int WalletManager3cf9d2_VerifySeed(void* ptr, struct Moc_PackedString seed);
void WalletManager3cf9d2_NewWalletAddress(void* ptr, void* addressM, struct Moc_PackedString id, int n, struct Moc_PackedString password);
void WalletManager3cf9d2_EncryptWallet(void* ptr, void* walletM, struct Moc_PackedString id, struct Moc_PackedString password);
void WalletManager3cf9d2_DecryptWallet(void* ptr, void* walletM, struct Moc_PackedString id, struct Moc_PackedString password);
int WalletManager3cf9d2_WalletManager3cf9d2_QRegisterMetaType();
int WalletManager3cf9d2_WalletManager3cf9d2_QRegisterMetaType2(char* typeName);
int WalletManager3cf9d2_WalletManager3cf9d2_QmlRegisterType();
int WalletManager3cf9d2_WalletManager3cf9d2_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* WalletManager3cf9d2___children_atList(void* ptr, int i);
void WalletManager3cf9d2___children_setList(void* ptr, void* i);
void* WalletManager3cf9d2___children_newList(void* ptr);
void* WalletManager3cf9d2___dynamicPropertyNames_atList(void* ptr, int i);
void WalletManager3cf9d2___dynamicPropertyNames_setList(void* ptr, void* i);
void* WalletManager3cf9d2___dynamicPropertyNames_newList(void* ptr);
void* WalletManager3cf9d2___findChildren_atList(void* ptr, int i);
void WalletManager3cf9d2___findChildren_setList(void* ptr, void* i);
void* WalletManager3cf9d2___findChildren_newList(void* ptr);
void* WalletManager3cf9d2___findChildren_atList3(void* ptr, int i);
void WalletManager3cf9d2___findChildren_setList3(void* ptr, void* i);
void* WalletManager3cf9d2___findChildren_newList3(void* ptr);
void* WalletManager3cf9d2___qFindChildren_atList2(void* ptr, int i);
void WalletManager3cf9d2___qFindChildren_setList2(void* ptr, void* i);
void* WalletManager3cf9d2___qFindChildren_newList2(void* ptr);
void* WalletManager3cf9d2_NewWalletManager(void* parent);
void WalletManager3cf9d2_DestroyWalletManager(void* ptr);
void WalletManager3cf9d2_DestroyWalletManagerDefault(void* ptr);
void WalletManager3cf9d2_ChildEventDefault(void* ptr, void* event);
void WalletManager3cf9d2_ConnectNotifyDefault(void* ptr, void* sign);
void WalletManager3cf9d2_CustomEventDefault(void* ptr, void* event);
void WalletManager3cf9d2_DeleteLaterDefault(void* ptr);
void WalletManager3cf9d2_DisconnectNotifyDefault(void* ptr, void* sign);
char WalletManager3cf9d2_EventDefault(void* ptr, void* e);
char WalletManager3cf9d2_EventFilterDefault(void* ptr, void* watched, void* event);
void WalletManager3cf9d2_TimerEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif

#endif