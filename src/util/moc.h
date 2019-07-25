

#pragma once

#ifndef GO_MOC_fdda3f_H
#define GO_MOC_fdda3f_H

#include <stdint.h>

#ifdef __cplusplus
class Walletfdda3f;
void Walletfdda3f_Walletfdda3f_QRegisterMetaTypes();
extern "C" {
#endif

struct Moc_PackedString { char* data; long long len; };
struct Moc_PackedList { void* data; long long len; };
int Walletfdda3f_Walletfdda3f_QRegisterMetaType();
int Walletfdda3f_Walletfdda3f_QRegisterMetaType2(char* typeName);
int Walletfdda3f_Walletfdda3f_QmlRegisterType();
int Walletfdda3f_Walletfdda3f_QmlRegisterType2(char* uri, int versionMajor, int versionMinor, char* qmlName);
void* Walletfdda3f___children_atList(void* ptr, int i);
void Walletfdda3f___children_setList(void* ptr, void* i);
void* Walletfdda3f___children_newList(void* ptr);
void* Walletfdda3f___dynamicPropertyNames_atList(void* ptr, int i);
void Walletfdda3f___dynamicPropertyNames_setList(void* ptr, void* i);
void* Walletfdda3f___dynamicPropertyNames_newList(void* ptr);
void* Walletfdda3f___findChildren_atList(void* ptr, int i);
void Walletfdda3f___findChildren_setList(void* ptr, void* i);
void* Walletfdda3f___findChildren_newList(void* ptr);
void* Walletfdda3f___findChildren_atList3(void* ptr, int i);
void Walletfdda3f___findChildren_setList3(void* ptr, void* i);
void* Walletfdda3f___findChildren_newList3(void* ptr);
void* Walletfdda3f___qFindChildren_atList2(void* ptr, int i);
void Walletfdda3f___qFindChildren_setList2(void* ptr, void* i);
void* Walletfdda3f___qFindChildren_newList2(void* ptr);
void* Walletfdda3f_NewWallet(void* parent);
void Walletfdda3f_DestroyWallet(void* ptr);
void Walletfdda3f_DestroyWalletDefault(void* ptr);
void Walletfdda3f_ChildEventDefault(void* ptr, void* event);
void Walletfdda3f_ConnectNotifyDefault(void* ptr, void* sign);
void Walletfdda3f_CustomEventDefault(void* ptr, void* event);
void Walletfdda3f_DeleteLaterDefault(void* ptr);
void Walletfdda3f_DisconnectNotifyDefault(void* ptr, void* sign);
char Walletfdda3f_EventDefault(void* ptr, void* e);
char Walletfdda3f_EventFilterDefault(void* ptr, void* watched, void* event);
void Walletfdda3f_TimerEventDefault(void* ptr, void* event);

#ifdef __cplusplus
}
#endif

#endif