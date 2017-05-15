// +build !minimal

#pragma once

#ifndef GO_QTDBUS_H
#define GO_QTDBUS_H

#include <stdint.h>

#ifdef __cplusplus
extern "C" {
#endif

struct QtDBus_PackedString { char* data; long long len; };
struct QtDBus_PackedList { void* data; long long len; };
void* QDBusAbstractAdaptor_NewQDBusAbstractAdaptor(void* obj);
void QDBusAbstractAdaptor_SetAutoRelaySignals(void* ptr, char enable);
void QDBusAbstractAdaptor_DestroyQDBusAbstractAdaptor(void* ptr);
char QDBusAbstractAdaptor_AutoRelaySignals(void* ptr);
void* QDBusAbstractAdaptor___dynamicPropertyNames_atList(void* ptr, int i);
void QDBusAbstractAdaptor___dynamicPropertyNames_setList(void* ptr, void* i);
void* QDBusAbstractAdaptor___dynamicPropertyNames_newList(void* ptr);
void* QDBusAbstractAdaptor___findChildren_atList2(void* ptr, int i);
void QDBusAbstractAdaptor___findChildren_setList2(void* ptr, void* i);
void* QDBusAbstractAdaptor___findChildren_newList2(void* ptr);
void* QDBusAbstractAdaptor___findChildren_atList3(void* ptr, int i);
void QDBusAbstractAdaptor___findChildren_setList3(void* ptr, void* i);
void* QDBusAbstractAdaptor___findChildren_newList3(void* ptr);
void* QDBusAbstractAdaptor___findChildren_atList(void* ptr, int i);
void QDBusAbstractAdaptor___findChildren_setList(void* ptr, void* i);
void* QDBusAbstractAdaptor___findChildren_newList(void* ptr);
void* QDBusAbstractAdaptor___children_atList(void* ptr, int i);
void QDBusAbstractAdaptor___children_setList(void* ptr, void* i);
void* QDBusAbstractAdaptor___children_newList(void* ptr);
char QDBusAbstractAdaptor_EventDefault(void* ptr, void* e);
char QDBusAbstractAdaptor_EventFilterDefault(void* ptr, void* watched, void* event);
void QDBusAbstractAdaptor_ChildEventDefault(void* ptr, void* event);
void QDBusAbstractAdaptor_ConnectNotifyDefault(void* ptr, void* sign);
void QDBusAbstractAdaptor_CustomEventDefault(void* ptr, void* event);
void QDBusAbstractAdaptor_DeleteLaterDefault(void* ptr);
void QDBusAbstractAdaptor_DisconnectNotifyDefault(void* ptr, void* sign);
void QDBusAbstractAdaptor_TimerEventDefault(void* ptr, void* event);
void* QDBusAbstractAdaptor_MetaObjectDefault(void* ptr);
void* QDBusAbstractInterface_Call2(void* ptr, long long mode, char* method, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6, void* arg7, void* arg8);
void* QDBusAbstractInterface_Call(void* ptr, char* method, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6, void* arg7, void* arg8);
void* QDBusAbstractInterface_CallWithArgumentList(void* ptr, long long mode, char* method, void* args);
void* QDBusAbstractInterface_AsyncCall(void* ptr, char* method, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6, void* arg7, void* arg8);
void* QDBusAbstractInterface_AsyncCallWithArgumentList(void* ptr, char* method, void* args);
char QDBusAbstractInterface_CallWithCallback(void* ptr, char* method, void* args, void* receiver, char* returnMethod, char* errorMethod);
char QDBusAbstractInterface_CallWithCallback2(void* ptr, char* method, void* args, void* receiver, char* slot);
void QDBusAbstractInterface_SetTimeout(void* ptr, int timeout);
void QDBusAbstractInterface_DestroyQDBusAbstractInterface(void* ptr);
void QDBusAbstractInterface_DestroyQDBusAbstractInterfaceDefault(void* ptr);
void* QDBusAbstractInterface_Connection(void* ptr);
void* QDBusAbstractInterface_LastError(void* ptr);
struct QtDBus_PackedString QDBusAbstractInterface_Interface(void* ptr);
struct QtDBus_PackedString QDBusAbstractInterface_Path(void* ptr);
struct QtDBus_PackedString QDBusAbstractInterface_Service(void* ptr);
char QDBusAbstractInterface_IsValid(void* ptr);
int QDBusAbstractInterface_Timeout(void* ptr);
void* QDBusAbstractInterface___callWithArgumentList_args_atList(void* ptr, int i);
void QDBusAbstractInterface___callWithArgumentList_args_setList(void* ptr, void* i);
void* QDBusAbstractInterface___callWithArgumentList_args_newList(void* ptr);
void* QDBusAbstractInterface___asyncCallWithArgumentList_args_atList(void* ptr, int i);
void QDBusAbstractInterface___asyncCallWithArgumentList_args_setList(void* ptr, void* i);
void* QDBusAbstractInterface___asyncCallWithArgumentList_args_newList(void* ptr);
void* QDBusAbstractInterface___callWithCallback_args_atList(void* ptr, int i);
void QDBusAbstractInterface___callWithCallback_args_setList(void* ptr, void* i);
void* QDBusAbstractInterface___callWithCallback_args_newList(void* ptr);
void* QDBusAbstractInterface___callWithCallback_args_atList2(void* ptr, int i);
void QDBusAbstractInterface___callWithCallback_args_setList2(void* ptr, void* i);
void* QDBusAbstractInterface___callWithCallback_args_newList2(void* ptr);
void* QDBusAbstractInterface___internalConstCall_args_atList(void* ptr, int i);
void QDBusAbstractInterface___internalConstCall_args_setList(void* ptr, void* i);
void* QDBusAbstractInterface___internalConstCall_args_newList(void* ptr);
void* QDBusAbstractInterface___dynamicPropertyNames_atList(void* ptr, int i);
void QDBusAbstractInterface___dynamicPropertyNames_setList(void* ptr, void* i);
void* QDBusAbstractInterface___dynamicPropertyNames_newList(void* ptr);
void* QDBusAbstractInterface___findChildren_atList2(void* ptr, int i);
void QDBusAbstractInterface___findChildren_setList2(void* ptr, void* i);
void* QDBusAbstractInterface___findChildren_newList2(void* ptr);
void* QDBusAbstractInterface___findChildren_atList3(void* ptr, int i);
void QDBusAbstractInterface___findChildren_setList3(void* ptr, void* i);
void* QDBusAbstractInterface___findChildren_newList3(void* ptr);
void* QDBusAbstractInterface___findChildren_atList(void* ptr, int i);
void QDBusAbstractInterface___findChildren_setList(void* ptr, void* i);
void* QDBusAbstractInterface___findChildren_newList(void* ptr);
void* QDBusAbstractInterface___children_atList(void* ptr, int i);
void QDBusAbstractInterface___children_setList(void* ptr, void* i);
void* QDBusAbstractInterface___children_newList(void* ptr);
char QDBusAbstractInterface_EventDefault(void* ptr, void* e);
char QDBusAbstractInterface_EventFilterDefault(void* ptr, void* watched, void* event);
void QDBusAbstractInterface_ChildEventDefault(void* ptr, void* event);
void QDBusAbstractInterface_ConnectNotifyDefault(void* ptr, void* sign);
void QDBusAbstractInterface_CustomEventDefault(void* ptr, void* event);
void QDBusAbstractInterface_DeleteLaterDefault(void* ptr);
void QDBusAbstractInterface_DisconnectNotifyDefault(void* ptr, void* sign);
void QDBusAbstractInterface_TimerEventDefault(void* ptr, void* event);
void* QDBusAbstractInterface_MetaObjectDefault(void* ptr);
void* QDBusArgument_NewQDBusArgument();
void* QDBusArgument_NewQDBusArgument3(void* other);
void* QDBusArgument_NewQDBusArgument2(void* other);
void QDBusArgument_BeginArray(void* ptr, int id);
void QDBusArgument_BeginMap(void* ptr, int kid, int vid);
void QDBusArgument_BeginMapEntry(void* ptr);
void QDBusArgument_BeginStructure(void* ptr);
void QDBusArgument_EndArray(void* ptr);
void QDBusArgument_EndMap(void* ptr);
void QDBusArgument_EndMapEntry(void* ptr);
void QDBusArgument_EndStructure(void* ptr);
void QDBusArgument_Swap(void* ptr, void* other);
void QDBusArgument_DestroyQDBusArgument(void* ptr);
long long QDBusArgument_CurrentType(void* ptr);
void* QDBusArgument_AsVariant(void* ptr);
char QDBusArgument_AtEnd(void* ptr);
void QDBusArgument_BeginArray2(void* ptr);
void QDBusArgument_BeginMap2(void* ptr);
void QDBusArgument_BeginMapEntry2(void* ptr);
void QDBusArgument_BeginStructure2(void* ptr);
void QDBusArgument_EndArray2(void* ptr);
void QDBusArgument_EndMap2(void* ptr);
void QDBusArgument_EndMapEntry2(void* ptr);
void QDBusArgument_EndStructure2(void* ptr);
void* QDBusConnection_QDBusConnection_ConnectToPeer(char* address, char* name);
void* QDBusConnection_QDBusConnection_SessionBus();
void* QDBusConnection_QDBusConnection_SystemBus();
void* QDBusConnection_QDBusConnection_LocalMachineId();
void* QDBusConnection_QDBusConnection_ConnectToBus(long long ty, char* name);
void* QDBusConnection_QDBusConnection_ConnectToBus2(char* address, char* name);
void* QDBusConnection_NewQDBusConnection3(void* other);
void* QDBusConnection_NewQDBusConnection2(void* other);
void* QDBusConnection_NewQDBusConnection(char* name);
char QDBusConnection_Connect(void* ptr, char* service, char* path, char* interfa, char* name, void* receiver, char* slot);
char QDBusConnection_Connect2(void* ptr, char* service, char* path, char* interfa, char* name, char* signature, void* receiver, char* slot);
char QDBusConnection_Connect3(void* ptr, char* service, char* path, char* interfa, char* name, char* argumentMatch, char* signature, void* receiver, char* slot);
char QDBusConnection_Disconnect(void* ptr, char* service, char* path, char* interfa, char* name, void* receiver, char* slot);
char QDBusConnection_Disconnect2(void* ptr, char* service, char* path, char* interfa, char* name, char* signature, void* receiver, char* slot);
char QDBusConnection_Disconnect3(void* ptr, char* service, char* path, char* interfa, char* name, char* argumentMatch, char* signature, void* receiver, char* slot);
char QDBusConnection_RegisterObject(void* ptr, char* path, void* object, long long options);
char QDBusConnection_RegisterObject2(void* ptr, char* path, char* interfa, void* object, long long options);
char QDBusConnection_RegisterService(void* ptr, char* serviceName);
char QDBusConnection_UnregisterService(void* ptr, char* serviceName);
void QDBusConnection_QDBusConnection_DisconnectFromBus(char* name);
void QDBusConnection_QDBusConnection_DisconnectFromPeer(char* name);
void QDBusConnection_Swap(void* ptr, void* other);
void QDBusConnection_UnregisterObject(void* ptr, char* path, long long mode);
void QDBusConnection_DestroyQDBusConnection(void* ptr);
long long QDBusConnection_ConnectionCapabilities(void* ptr);
void* QDBusConnection_Interface(void* ptr);
void* QDBusConnection_LastError(void* ptr);
void* QDBusConnection_Call(void* ptr, void* message, long long mode, int timeout);
void* QDBusConnection_AsyncCall(void* ptr, void* message, int timeout);
void* QDBusConnection_ObjectRegisteredAt(void* ptr, char* path);
struct QtDBus_PackedString QDBusConnection_BaseService(void* ptr);
struct QtDBus_PackedString QDBusConnection_Name(void* ptr);
char QDBusConnection_CallWithCallback(void* ptr, void* message, void* receiver, char* returnMethod, char* errorMethod, int timeout);
char QDBusConnection_IsConnected(void* ptr);
char QDBusConnection_Send(void* ptr, void* message);
void QDBusConnectionInterface_ConnectServiceRegistered(void* ptr);
void QDBusConnectionInterface_DisconnectServiceRegistered(void* ptr);
void QDBusConnectionInterface_ServiceRegistered(void* ptr, char* serviceName);
void QDBusConnectionInterface_ConnectCallWithCallbackFailed(void* ptr);
void QDBusConnectionInterface_DisconnectCallWithCallbackFailed(void* ptr);
void QDBusConnectionInterface_CallWithCallbackFailed(void* ptr, void* error, void* call);
void QDBusConnectionInterface_ConnectServiceUnregistered(void* ptr);
void QDBusConnectionInterface_DisconnectServiceUnregistered(void* ptr);
void QDBusConnectionInterface_ServiceUnregistered(void* ptr, char* serviceName);
void* QDBusContext_NewQDBusContext();
void QDBusContext_DestroyQDBusContext(void* ptr);
void* QDBusContext_Connection(void* ptr);
char QDBusContext_CalledFromDBus(void* ptr);
char QDBusContext_IsDelayedReply(void* ptr);
void* QDBusContext_Message(void* ptr);
void QDBusContext_SendErrorReply2(void* ptr, long long ty, char* msg);
void QDBusContext_SendErrorReply(void* ptr, char* name, char* msg);
void QDBusContext_SetDelayedReply(void* ptr, char enable);
void* QDBusError_NewQDBusError(void* other);
struct QtDBus_PackedString QDBusError_QDBusError_ErrorString(long long error);
void QDBusError_Swap(void* ptr, void* other);
long long QDBusError_Type(void* ptr);
struct QtDBus_PackedString QDBusError_Message(void* ptr);
struct QtDBus_PackedString QDBusError_Name(void* ptr);
char QDBusError_IsValid(void* ptr);
void* QDBusInterface_NewQDBusInterface(char* service, char* path, char* interfa, void* connection, void* parent);
void QDBusInterface_DestroyQDBusInterface(void* ptr);
void* QDBusMessage_QDBusMessage_CreateError3(long long ty, char* msg);
void* QDBusMessage_QDBusMessage_CreateError2(void* error);
void* QDBusMessage_QDBusMessage_CreateError(char* name, char* msg);
void* QDBusMessage_QDBusMessage_CreateMethodCall(char* service, char* path, char* interfa, char* method);
void* QDBusMessage_QDBusMessage_CreateSignal(char* path, char* interfa, char* name);
void* QDBusMessage_QDBusMessage_CreateTargetedSignal(char* service, char* path, char* interfa, char* name);
void* QDBusMessage_NewQDBusMessage();
void* QDBusMessage_NewQDBusMessage2(void* other);
void QDBusMessage_SetArguments(void* ptr, void* arguments);
void QDBusMessage_SetAutoStartService(void* ptr, char enable);
void QDBusMessage_Swap(void* ptr, void* other);
void QDBusMessage_DestroyQDBusMessage(void* ptr);
long long QDBusMessage_Type(void* ptr);
void* QDBusMessage_CreateErrorReply3(void* ptr, long long ty, char* msg);
void* QDBusMessage_CreateErrorReply2(void* ptr, void* error);
void* QDBusMessage_CreateErrorReply(void* ptr, char* name, char* msg);
void* QDBusMessage_CreateReply(void* ptr, void* arguments);
void* QDBusMessage_CreateReply2(void* ptr, void* argument);
struct QtDBus_PackedList QDBusMessage_Arguments(void* ptr);
struct QtDBus_PackedString QDBusMessage_ErrorMessage(void* ptr);
struct QtDBus_PackedString QDBusMessage_ErrorName(void* ptr);
struct QtDBus_PackedString QDBusMessage_Interface(void* ptr);
struct QtDBus_PackedString QDBusMessage_Member(void* ptr);
struct QtDBus_PackedString QDBusMessage_Path(void* ptr);
struct QtDBus_PackedString QDBusMessage_Service(void* ptr);
struct QtDBus_PackedString QDBusMessage_Signature(void* ptr);
char QDBusMessage_AutoStartService(void* ptr);
char QDBusMessage_IsDelayedReply(void* ptr);
char QDBusMessage_IsReplyRequired(void* ptr);
void QDBusMessage_SetDelayedReply(void* ptr, char enable);
void* QDBusMessage___setArguments_arguments_atList(void* ptr, int i);
void QDBusMessage___setArguments_arguments_setList(void* ptr, void* i);
void* QDBusMessage___setArguments_arguments_newList(void* ptr);
void* QDBusMessage___createReply_arguments_atList(void* ptr, int i);
void QDBusMessage___createReply_arguments_setList(void* ptr, void* i);
void* QDBusMessage___createReply_arguments_newList(void* ptr);
void* QDBusMessage___arguments_atList(void* ptr, int i);
void QDBusMessage___arguments_setList(void* ptr, void* i);
void* QDBusMessage___arguments_newList(void* ptr);
void* QDBusObjectPath_NewQDBusObjectPath();
void* QDBusObjectPath_NewQDBusObjectPath3(void* path);
void* QDBusObjectPath_NewQDBusObjectPath5(char* p);
void* QDBusObjectPath_NewQDBusObjectPath4(char* path);
void* QDBusObjectPath_NewQDBusObjectPath2(char* path);
void QDBusObjectPath_SetPath(void* ptr, char* path);
void QDBusObjectPath_Swap(void* ptr, void* other);
struct QtDBus_PackedString QDBusObjectPath_Path(void* ptr);
void* QDBusPendingCall_QDBusPendingCall_FromCompletedCall(void* msg);
void* QDBusPendingCall_QDBusPendingCall_FromError(void* error);
void* QDBusPendingCall_NewQDBusPendingCall(void* other);
void QDBusPendingCall_Swap(void* ptr, void* other);
void QDBusPendingCall_DestroyQDBusPendingCall(void* ptr);
void* QDBusPendingCallWatcher_NewQDBusPendingCallWatcher(void* call, void* parent);
void QDBusPendingCallWatcher_ConnectFinished(void* ptr);
void QDBusPendingCallWatcher_DisconnectFinished(void* ptr);
void QDBusPendingCallWatcher_Finished(void* ptr, void* self);
void QDBusPendingCallWatcher_WaitForFinished(void* ptr);
void QDBusPendingCallWatcher_DestroyQDBusPendingCallWatcher(void* ptr);
char QDBusPendingCallWatcher_IsFinished(void* ptr);
void* QDBusPendingCallWatcher___dynamicPropertyNames_atList(void* ptr, int i);
void QDBusPendingCallWatcher___dynamicPropertyNames_setList(void* ptr, void* i);
void* QDBusPendingCallWatcher___dynamicPropertyNames_newList(void* ptr);
void* QDBusPendingCallWatcher___findChildren_atList2(void* ptr, int i);
void QDBusPendingCallWatcher___findChildren_setList2(void* ptr, void* i);
void* QDBusPendingCallWatcher___findChildren_newList2(void* ptr);
void* QDBusPendingCallWatcher___findChildren_atList3(void* ptr, int i);
void QDBusPendingCallWatcher___findChildren_setList3(void* ptr, void* i);
void* QDBusPendingCallWatcher___findChildren_newList3(void* ptr);
void* QDBusPendingCallWatcher___findChildren_atList(void* ptr, int i);
void QDBusPendingCallWatcher___findChildren_setList(void* ptr, void* i);
void* QDBusPendingCallWatcher___findChildren_newList(void* ptr);
void* QDBusPendingCallWatcher___children_atList(void* ptr, int i);
void QDBusPendingCallWatcher___children_setList(void* ptr, void* i);
void* QDBusPendingCallWatcher___children_newList(void* ptr);
char QDBusPendingCallWatcher_Event(void* ptr, void* e);
char QDBusPendingCallWatcher_EventDefault(void* ptr, void* e);
char QDBusPendingCallWatcher_EventFilter(void* ptr, void* watched, void* event);
char QDBusPendingCallWatcher_EventFilterDefault(void* ptr, void* watched, void* event);
void QDBusPendingCallWatcher_ChildEvent(void* ptr, void* event);
void QDBusPendingCallWatcher_ChildEventDefault(void* ptr, void* event);
void QDBusPendingCallWatcher_ConnectNotify(void* ptr, void* sign);
void QDBusPendingCallWatcher_ConnectNotifyDefault(void* ptr, void* sign);
void QDBusPendingCallWatcher_CustomEvent(void* ptr, void* event);
void QDBusPendingCallWatcher_CustomEventDefault(void* ptr, void* event);
void QDBusPendingCallWatcher_DeleteLater(void* ptr);
void QDBusPendingCallWatcher_DeleteLaterDefault(void* ptr);
void QDBusPendingCallWatcher_DisconnectNotify(void* ptr, void* sign);
void QDBusPendingCallWatcher_DisconnectNotifyDefault(void* ptr, void* sign);
void QDBusPendingCallWatcher_TimerEvent(void* ptr, void* event);
void QDBusPendingCallWatcher_TimerEventDefault(void* ptr, void* event);
void* QDBusPendingCallWatcher_MetaObject(void* ptr);
void* QDBusPendingCallWatcher_MetaObjectDefault(void* ptr);
int QDBusPendingReplyTypes_QDBusPendingReplyTypes_MetaTypeFor2(void* vqv);
void* QDBusServer_NewQDBusServer2(void* parent);
void* QDBusServer_NewQDBusServer(char* address, void* parent);
void QDBusServer_ConnectNewConnection(void* ptr);
void QDBusServer_DisconnectNewConnection(void* ptr);
void QDBusServer_NewConnection(void* ptr, void* connection);
void QDBusServer_SetAnonymousAuthenticationAllowed(void* ptr, char value);
void QDBusServer_DestroyQDBusServer(void* ptr);
void QDBusServer_DestroyQDBusServerDefault(void* ptr);
void* QDBusServer_LastError(void* ptr);
struct QtDBus_PackedString QDBusServer_Address(void* ptr);
char QDBusServer_IsAnonymousAuthenticationAllowed(void* ptr);
char QDBusServer_IsConnected(void* ptr);
void* QDBusServer___dynamicPropertyNames_atList(void* ptr, int i);
void QDBusServer___dynamicPropertyNames_setList(void* ptr, void* i);
void* QDBusServer___dynamicPropertyNames_newList(void* ptr);
void* QDBusServer___findChildren_atList2(void* ptr, int i);
void QDBusServer___findChildren_setList2(void* ptr, void* i);
void* QDBusServer___findChildren_newList2(void* ptr);
void* QDBusServer___findChildren_atList3(void* ptr, int i);
void QDBusServer___findChildren_setList3(void* ptr, void* i);
void* QDBusServer___findChildren_newList3(void* ptr);
void* QDBusServer___findChildren_atList(void* ptr, int i);
void QDBusServer___findChildren_setList(void* ptr, void* i);
void* QDBusServer___findChildren_newList(void* ptr);
void* QDBusServer___children_atList(void* ptr, int i);
void QDBusServer___children_setList(void* ptr, void* i);
void* QDBusServer___children_newList(void* ptr);
char QDBusServer_EventDefault(void* ptr, void* e);
char QDBusServer_EventFilterDefault(void* ptr, void* watched, void* event);
void QDBusServer_ChildEventDefault(void* ptr, void* event);
void QDBusServer_ConnectNotifyDefault(void* ptr, void* sign);
void QDBusServer_CustomEventDefault(void* ptr, void* event);
void QDBusServer_DeleteLaterDefault(void* ptr);
void QDBusServer_DisconnectNotifyDefault(void* ptr, void* sign);
void QDBusServer_TimerEventDefault(void* ptr, void* event);
void* QDBusServer_MetaObjectDefault(void* ptr);
void QDBusServiceWatcher_ConnectServiceRegistered(void* ptr);
void QDBusServiceWatcher_DisconnectServiceRegistered(void* ptr);
void QDBusServiceWatcher_ServiceRegistered(void* ptr, char* serviceName);
void QDBusServiceWatcher_SetConnection(void* ptr, void* connection);
void QDBusServiceWatcher_SetWatchMode(void* ptr, long long mode);
long long QDBusServiceWatcher_WatchMode(void* ptr);
void* QDBusServiceWatcher_NewQDBusServiceWatcher(void* parent);
void* QDBusServiceWatcher_NewQDBusServiceWatcher2(char* service, void* connection, long long watchMode, void* parent);
char QDBusServiceWatcher_RemoveWatchedService(void* ptr, char* service);
void QDBusServiceWatcher_AddWatchedService(void* ptr, char* newService);
void QDBusServiceWatcher_ConnectServiceOwnerChanged(void* ptr);
void QDBusServiceWatcher_DisconnectServiceOwnerChanged(void* ptr);
void QDBusServiceWatcher_ServiceOwnerChanged(void* ptr, char* serviceName, char* oldOwner, char* newOwner);
void QDBusServiceWatcher_ConnectServiceUnregistered(void* ptr);
void QDBusServiceWatcher_DisconnectServiceUnregistered(void* ptr);
void QDBusServiceWatcher_ServiceUnregistered(void* ptr, char* serviceName);
void QDBusServiceWatcher_SetWatchedServices(void* ptr, char* services);
void QDBusServiceWatcher_DestroyQDBusServiceWatcher(void* ptr);
void* QDBusServiceWatcher_Connection(void* ptr);
struct QtDBus_PackedString QDBusServiceWatcher_WatchedServices(void* ptr);
void* QDBusServiceWatcher___dynamicPropertyNames_atList(void* ptr, int i);
void QDBusServiceWatcher___dynamicPropertyNames_setList(void* ptr, void* i);
void* QDBusServiceWatcher___dynamicPropertyNames_newList(void* ptr);
void* QDBusServiceWatcher___findChildren_atList2(void* ptr, int i);
void QDBusServiceWatcher___findChildren_setList2(void* ptr, void* i);
void* QDBusServiceWatcher___findChildren_newList2(void* ptr);
void* QDBusServiceWatcher___findChildren_atList3(void* ptr, int i);
void QDBusServiceWatcher___findChildren_setList3(void* ptr, void* i);
void* QDBusServiceWatcher___findChildren_newList3(void* ptr);
void* QDBusServiceWatcher___findChildren_atList(void* ptr, int i);
void QDBusServiceWatcher___findChildren_setList(void* ptr, void* i);
void* QDBusServiceWatcher___findChildren_newList(void* ptr);
void* QDBusServiceWatcher___children_atList(void* ptr, int i);
void QDBusServiceWatcher___children_setList(void* ptr, void* i);
void* QDBusServiceWatcher___children_newList(void* ptr);
char QDBusServiceWatcher_EventDefault(void* ptr, void* e);
char QDBusServiceWatcher_EventFilterDefault(void* ptr, void* watched, void* event);
void QDBusServiceWatcher_ChildEventDefault(void* ptr, void* event);
void QDBusServiceWatcher_ConnectNotifyDefault(void* ptr, void* sign);
void QDBusServiceWatcher_CustomEventDefault(void* ptr, void* event);
void QDBusServiceWatcher_DeleteLaterDefault(void* ptr);
void QDBusServiceWatcher_DisconnectNotifyDefault(void* ptr, void* sign);
void QDBusServiceWatcher_TimerEventDefault(void* ptr, void* event);
void* QDBusServiceWatcher_MetaObjectDefault(void* ptr);
void* QDBusSignature_NewQDBusSignature();
void* QDBusSignature_NewQDBusSignature3(void* signature);
void* QDBusSignature_NewQDBusSignature5(char* sig);
void* QDBusSignature_NewQDBusSignature4(char* signature);
void* QDBusSignature_NewQDBusSignature2(char* signature);
void QDBusSignature_SetSignature(void* ptr, char* signature);
void QDBusSignature_Swap(void* ptr, void* other);
struct QtDBus_PackedString QDBusSignature_Signature(void* ptr);
void* QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor();
void* QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor3(void* other);
void* QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor2(int fileDescriptor);
char QDBusUnixFileDescriptor_QDBusUnixFileDescriptor_IsSupported();
void QDBusUnixFileDescriptor_SetFileDescriptor(void* ptr, int fileDescriptor);
void QDBusUnixFileDescriptor_Swap(void* ptr, void* other);
void QDBusUnixFileDescriptor_DestroyQDBusUnixFileDescriptor(void* ptr);
char QDBusUnixFileDescriptor_IsValid(void* ptr);
int QDBusUnixFileDescriptor_FileDescriptor(void* ptr);
void* QDBusVariant_NewQDBusVariant();
void* QDBusVariant_NewQDBusVariant3(void* v);
void* QDBusVariant_NewQDBusVariant2(void* variant);
void QDBusVariant_SetVariant(void* ptr, void* variant);
void QDBusVariant_Swap(void* ptr, void* other);
void* QDBusVariant_Variant(void* ptr);
void* QDBusVirtualObject_NewQDBusVirtualObject(void* parent);
char QDBusVirtualObject_HandleMessage(void* ptr, void* message, void* connection);
void QDBusVirtualObject_DestroyQDBusVirtualObject(void* ptr);
void QDBusVirtualObject_DestroyQDBusVirtualObjectDefault(void* ptr);
struct QtDBus_PackedString QDBusVirtualObject_Introspect(void* ptr, char* path);
void* QDBusVirtualObject___dynamicPropertyNames_atList(void* ptr, int i);
void QDBusVirtualObject___dynamicPropertyNames_setList(void* ptr, void* i);
void* QDBusVirtualObject___dynamicPropertyNames_newList(void* ptr);
void* QDBusVirtualObject___findChildren_atList2(void* ptr, int i);
void QDBusVirtualObject___findChildren_setList2(void* ptr, void* i);
void* QDBusVirtualObject___findChildren_newList2(void* ptr);
void* QDBusVirtualObject___findChildren_atList3(void* ptr, int i);
void QDBusVirtualObject___findChildren_setList3(void* ptr, void* i);
void* QDBusVirtualObject___findChildren_newList3(void* ptr);
void* QDBusVirtualObject___findChildren_atList(void* ptr, int i);
void QDBusVirtualObject___findChildren_setList(void* ptr, void* i);
void* QDBusVirtualObject___findChildren_newList(void* ptr);
void* QDBusVirtualObject___children_atList(void* ptr, int i);
void QDBusVirtualObject___children_setList(void* ptr, void* i);
void* QDBusVirtualObject___children_newList(void* ptr);
char QDBusVirtualObject_EventDefault(void* ptr, void* e);
char QDBusVirtualObject_EventFilterDefault(void* ptr, void* watched, void* event);
void QDBusVirtualObject_ChildEventDefault(void* ptr, void* event);
void QDBusVirtualObject_ConnectNotifyDefault(void* ptr, void* sign);
void QDBusVirtualObject_CustomEventDefault(void* ptr, void* event);
void QDBusVirtualObject_DeleteLaterDefault(void* ptr);
void QDBusVirtualObject_DisconnectNotifyDefault(void* ptr, void* sign);
void QDBusVirtualObject_TimerEventDefault(void* ptr, void* event);
void* QDBusVirtualObject_MetaObjectDefault(void* ptr);

#ifdef __cplusplus
}
#endif

#endif