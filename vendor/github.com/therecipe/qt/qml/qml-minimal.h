// +build minimal

#pragma once

#ifndef GO_QTQML_H
#define GO_QTQML_H

#include <stdint.h>

#ifdef __cplusplus
int QJSEngine_QJSEngine_QRegisterMetaType();
int QQmlEngine_QQmlEngine_QRegisterMetaType();
extern "C" {
#endif

struct QtQml_PackedString { char* data; long long len; };
struct QtQml_PackedList { void* data; long long len; };
void* QJSEngine_NewQJSEngine();
void* QJSEngine_NewQJSEngine2(void* parent);
void* QJSEngine_NewQObject(void* ptr, void* object);
void QJSEngine_DestroyQJSEngine(void* ptr);
void QJSEngine_DestroyQJSEngineDefault(void* ptr);
void* QJSEngine___children_atList(void* ptr, int i);
void QJSEngine___children_setList(void* ptr, void* i);
void* QJSEngine___children_newList(void* ptr);
void* QJSEngine___dynamicPropertyNames_atList(void* ptr, int i);
void QJSEngine___dynamicPropertyNames_setList(void* ptr, void* i);
void* QJSEngine___dynamicPropertyNames_newList(void* ptr);
void* QJSEngine___findChildren_atList(void* ptr, int i);
void QJSEngine___findChildren_setList(void* ptr, void* i);
void* QJSEngine___findChildren_newList(void* ptr);
void* QJSEngine___findChildren_atList3(void* ptr, int i);
void QJSEngine___findChildren_setList3(void* ptr, void* i);
void* QJSEngine___findChildren_newList3(void* ptr);
void* QJSEngine___qFindChildren_atList2(void* ptr, int i);
void QJSEngine___qFindChildren_setList2(void* ptr, void* i);
void* QJSEngine___qFindChildren_newList2(void* ptr);
void QJSEngine_ChildEventDefault(void* ptr, void* event);
void QJSEngine_ConnectNotifyDefault(void* ptr, void* sign);
void QJSEngine_CustomEventDefault(void* ptr, void* event);
void QJSEngine_DeleteLaterDefault(void* ptr);
void QJSEngine_DisconnectNotifyDefault(void* ptr, void* sign);
char QJSEngine_EventDefault(void* ptr, void* e);
char QJSEngine_EventFilterDefault(void* ptr, void* watched, void* event);
void QJSEngine_TimerEventDefault(void* ptr, void* event);
void* QJSValue_NewQJSValue(long long value);
void* QJSValue_NewQJSValue2(void* other);
void* QJSValue_NewQJSValue3(void* other);
void* QJSValue_NewQJSValue4(char value);
void* QJSValue_NewQJSValue5(int value);
void* QJSValue_NewQJSValue6(unsigned int value);
void* QJSValue_NewQJSValue7(double value);
void* QJSValue_NewQJSValue8(struct QtQml_PackedString value);
void* QJSValue_NewQJSValue9(void* value);
void* QJSValue_NewQJSValue10(char* value);
void* QJSValue_Property(void* ptr, struct QtQml_PackedString name);
void* QJSValue_Property2(void* ptr, unsigned int arrayIndex);
int QJSValue_ToInt(void* ptr);
void* QJSValue_ToVariant(void* ptr);
void QJSValue_DestroyQJSValue(void* ptr);
void* QJSValue___call_args_atList(void* ptr, int i);
void QJSValue___call_args_setList(void* ptr, void* i);
void* QJSValue___call_args_newList(void* ptr);
void* QJSValue___callAsConstructor_args_atList(void* ptr, int i);
void QJSValue___callAsConstructor_args_setList(void* ptr, void* i);
void* QJSValue___callAsConstructor_args_newList(void* ptr);
void* QJSValue___callWithInstance_args_atList(void* ptr, int i);
void QJSValue___callWithInstance_args_setList(void* ptr, void* i);
void* QJSValue___callWithInstance_args_newList(void* ptr);
void* QQmlEngine_NewQQmlEngine(void* parent);
void QQmlEngine_AddImportPath(void* ptr, struct QtQml_PackedString path);
void QQmlEngine_DestroyQQmlEngine(void* ptr);
void QQmlEngine_DestroyQQmlEngineDefault(void* ptr);
int QQmlEngine_QQmlEngine_QmlRegisterType(void* url, char* uri, int versionMajor, int versionMinor, char* qmlName);
void* QQmlEngine___importPlugin_errors_atList(void* ptr, int i);
void QQmlEngine___importPlugin_errors_setList(void* ptr, void* i);
void* QQmlEngine___importPlugin_errors_newList(void* ptr);
void* QQmlEngine___qmlDebug_errors_atList3(void* ptr, int i);
void QQmlEngine___qmlDebug_errors_setList3(void* ptr, void* i);
void* QQmlEngine___qmlDebug_errors_newList3(void* ptr);
void* QQmlEngine___qmlInfo_errors_atList3(void* ptr, int i);
void QQmlEngine___qmlInfo_errors_setList3(void* ptr, void* i);
void* QQmlEngine___qmlInfo_errors_newList3(void* ptr);
void* QQmlEngine___qmlWarning_errors_atList3(void* ptr, int i);
void QQmlEngine___qmlWarning_errors_setList3(void* ptr, void* i);
void* QQmlEngine___qmlWarning_errors_newList3(void* ptr);
void* QQmlEngine___warnings_warnings_atList(void* ptr, int i);
void QQmlEngine___warnings_warnings_setList(void* ptr, void* i);
void* QQmlEngine___warnings_warnings_newList(void* ptr);
void* QQmlError_NewQQmlError();
void* QQmlError_NewQQmlError2(void* other);
int QQmlError_Column(void* ptr);
void* QQmlError_Object(void* ptr);
void* QQmlError_Url(void* ptr);
void QQmlParserStatus_ClassBegin(void* ptr);
void QQmlParserStatus_ComponentComplete(void* ptr);

#ifdef __cplusplus
}
#endif

#endif