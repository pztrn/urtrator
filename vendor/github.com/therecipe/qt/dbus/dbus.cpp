// +build !minimal

#define protected public
#define private public

#include "dbus.h"
#include "_cgo_export.h"

#include <QByteArray>
#include <QCamera>
#include <QCameraImageCapture>
#include <QChildEvent>
#include <QDBusAbstractAdaptor>
#include <QDBusAbstractInterface>
#include <QDBusArgument>
#include <QDBusConnection>
#include <QDBusConnectionInterface>
#include <QDBusContext>
#include <QDBusError>
#include <QDBusInterface>
#include <QDBusMessage>
#include <QDBusObjectPath>
#include <QDBusPendingCall>
#include <QDBusPendingCallWatcher>
#include <QDBusPendingReply>
#include <QDBusServer>
#include <QDBusServiceWatcher>
#include <QDBusSignature>
#include <QDBusUnixFileDescriptor>
#include <QDBusVariant>
#include <QDBusVirtualObject>
#include <QEvent>
#include <QExtensionFactory>
#include <QExtensionManager>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QLatin1String>
#include <QLayout>
#include <QList>
#include <QMediaPlaylist>
#include <QMediaRecorder>
#include <QMetaMethod>
#include <QMetaObject>
#include <QObject>
#include <QOffscreenSurface>
#include <QPaintDevice>
#include <QPaintDeviceWindow>
#include <QPdfWriter>
#include <QQuickItem>
#include <QRadioData>
#include <QSignalSpy>
#include <QString>
#include <QTime>
#include <QTimer>
#include <QTimerEvent>
#include <QVariant>
#include <QWidget>
#include <QWindow>

class MyQDBusAbstractAdaptor: public QDBusAbstractAdaptor
{
public:
	MyQDBusAbstractAdaptor(QObject *obj) : QDBusAbstractAdaptor(obj) {};
	bool event(QEvent * e) { return callbackQDBusAbstractAdaptor_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDBusAbstractAdaptor_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQDBusAbstractAdaptor_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDBusAbstractAdaptor_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDBusAbstractAdaptor_CustomEvent(this, event); };
	void deleteLater() { callbackQDBusAbstractAdaptor_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQDBusAbstractAdaptor_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDBusAbstractAdaptor_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtDBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQDBusAbstractAdaptor_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQDBusAbstractAdaptor_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDBusAbstractAdaptor_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

void* QDBusAbstractAdaptor_NewQDBusAbstractAdaptor(void* obj)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(obj))) {
		return new MyQDBusAbstractAdaptor(static_cast<QCameraImageCapture*>(obj));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(obj))) {
		return new MyQDBusAbstractAdaptor(static_cast<QDBusPendingCallWatcher*>(obj));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(obj))) {
		return new MyQDBusAbstractAdaptor(static_cast<QExtensionFactory*>(obj));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(obj))) {
		return new MyQDBusAbstractAdaptor(static_cast<QExtensionManager*>(obj));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(obj))) {
		return new MyQDBusAbstractAdaptor(static_cast<QGraphicsObject*>(obj));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(obj))) {
		return new MyQDBusAbstractAdaptor(static_cast<QGraphicsWidget*>(obj));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(obj))) {
		return new MyQDBusAbstractAdaptor(static_cast<QLayout*>(obj));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(obj))) {
		return new MyQDBusAbstractAdaptor(static_cast<QMediaPlaylist*>(obj));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(obj))) {
		return new MyQDBusAbstractAdaptor(static_cast<QMediaRecorder*>(obj));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(obj))) {
		return new MyQDBusAbstractAdaptor(static_cast<QOffscreenSurface*>(obj));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(obj))) {
		return new MyQDBusAbstractAdaptor(static_cast<QPaintDeviceWindow*>(obj));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(obj))) {
		return new MyQDBusAbstractAdaptor(static_cast<QPdfWriter*>(obj));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(obj))) {
		return new MyQDBusAbstractAdaptor(static_cast<QQuickItem*>(obj));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(obj))) {
		return new MyQDBusAbstractAdaptor(static_cast<QRadioData*>(obj));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(obj))) {
		return new MyQDBusAbstractAdaptor(static_cast<QSignalSpy*>(obj));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(obj))) {
		return new MyQDBusAbstractAdaptor(static_cast<QWidget*>(obj));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(obj))) {
		return new MyQDBusAbstractAdaptor(static_cast<QWindow*>(obj));
	} else {
		return new MyQDBusAbstractAdaptor(static_cast<QObject*>(obj));
	}
}

void QDBusAbstractAdaptor_SetAutoRelaySignals(void* ptr, char enable)
{
	static_cast<QDBusAbstractAdaptor*>(ptr)->setAutoRelaySignals(enable != 0);
}

void QDBusAbstractAdaptor_DestroyQDBusAbstractAdaptor(void* ptr)
{
	static_cast<QDBusAbstractAdaptor*>(ptr)->~QDBusAbstractAdaptor();
}

char QDBusAbstractAdaptor_AutoRelaySignals(void* ptr)
{
	return static_cast<QDBusAbstractAdaptor*>(ptr)->autoRelaySignals();
}

void* QDBusAbstractAdaptor___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QDBusAbstractAdaptor___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QDBusAbstractAdaptor___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QDBusAbstractAdaptor___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDBusAbstractAdaptor___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusAbstractAdaptor___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QDBusAbstractAdaptor___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDBusAbstractAdaptor___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusAbstractAdaptor___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QDBusAbstractAdaptor___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDBusAbstractAdaptor___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusAbstractAdaptor___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* QDBusAbstractAdaptor___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QDBusAbstractAdaptor___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusAbstractAdaptor___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

char QDBusAbstractAdaptor_EventDefault(void* ptr, void* e)
{
		return static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::event(static_cast<QEvent*>(e));
}

char QDBusAbstractAdaptor_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(watched))) {
		return static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::eventFilter(static_cast<QDBusPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QDBusAbstractAdaptor_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::childEvent(static_cast<QChildEvent*>(event));
}

void QDBusAbstractAdaptor_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusAbstractAdaptor_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::customEvent(static_cast<QEvent*>(event));
}

void QDBusAbstractAdaptor_DeleteLaterDefault(void* ptr)
{
		static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::deleteLater();
}

void QDBusAbstractAdaptor_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusAbstractAdaptor_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QDBusAbstractAdaptor_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QDBusAbstractAdaptor*>(ptr)->QDBusAbstractAdaptor::metaObject());
}

class MyQDBusAbstractInterface: public QDBusAbstractInterface
{
public:
	 ~MyQDBusAbstractInterface() { callbackQDBusAbstractInterface_DestroyQDBusAbstractInterface(this); };
	bool event(QEvent * e) { return callbackQDBusAbstractInterface_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDBusAbstractInterface_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQDBusAbstractInterface_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDBusAbstractInterface_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDBusAbstractInterface_CustomEvent(this, event); };
	void deleteLater() { callbackQDBusAbstractInterface_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQDBusAbstractInterface_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDBusAbstractInterface_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtDBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQDBusAbstractInterface_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQDBusAbstractInterface_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDBusAbstractInterface_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

void* QDBusAbstractInterface_Call2(void* ptr, long long mode, char* method, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6, void* arg7, void* arg8)
{
	return new QDBusMessage(static_cast<QDBusAbstractInterface*>(ptr)->call(static_cast<QDBus::CallMode>(mode), QString(method), *static_cast<QVariant*>(arg1), *static_cast<QVariant*>(arg2), *static_cast<QVariant*>(arg3), *static_cast<QVariant*>(arg4), *static_cast<QVariant*>(arg5), *static_cast<QVariant*>(arg6), *static_cast<QVariant*>(arg7), *static_cast<QVariant*>(arg8)));
}

void* QDBusAbstractInterface_Call(void* ptr, char* method, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6, void* arg7, void* arg8)
{
	return new QDBusMessage(static_cast<QDBusAbstractInterface*>(ptr)->call(QString(method), *static_cast<QVariant*>(arg1), *static_cast<QVariant*>(arg2), *static_cast<QVariant*>(arg3), *static_cast<QVariant*>(arg4), *static_cast<QVariant*>(arg5), *static_cast<QVariant*>(arg6), *static_cast<QVariant*>(arg7), *static_cast<QVariant*>(arg8)));
}

void* QDBusAbstractInterface_CallWithArgumentList(void* ptr, long long mode, char* method, void* args)
{
	return new QDBusMessage(static_cast<QDBusAbstractInterface*>(ptr)->callWithArgumentList(static_cast<QDBus::CallMode>(mode), QString(method), *static_cast<QList<QVariant>*>(args)));
}

void* QDBusAbstractInterface_AsyncCall(void* ptr, char* method, void* arg1, void* arg2, void* arg3, void* arg4, void* arg5, void* arg6, void* arg7, void* arg8)
{
	return new QDBusPendingCall(static_cast<QDBusAbstractInterface*>(ptr)->asyncCall(QString(method), *static_cast<QVariant*>(arg1), *static_cast<QVariant*>(arg2), *static_cast<QVariant*>(arg3), *static_cast<QVariant*>(arg4), *static_cast<QVariant*>(arg5), *static_cast<QVariant*>(arg6), *static_cast<QVariant*>(arg7), *static_cast<QVariant*>(arg8)));
}

void* QDBusAbstractInterface_AsyncCallWithArgumentList(void* ptr, char* method, void* args)
{
	return new QDBusPendingCall(static_cast<QDBusAbstractInterface*>(ptr)->asyncCallWithArgumentList(QString(method), *static_cast<QList<QVariant>*>(args)));
}

char QDBusAbstractInterface_CallWithCallback(void* ptr, char* method, void* args, void* receiver, char* returnMethod, char* errorMethod)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(receiver))) {
		return static_cast<QDBusAbstractInterface*>(ptr)->callWithCallback(QString(method), *static_cast<QList<QVariant>*>(args), static_cast<QDBusPendingCallWatcher*>(receiver), const_cast<const char*>(returnMethod), const_cast<const char*>(errorMethod));
	} else {
		return static_cast<QDBusAbstractInterface*>(ptr)->callWithCallback(QString(method), *static_cast<QList<QVariant>*>(args), static_cast<QObject*>(receiver), const_cast<const char*>(returnMethod), const_cast<const char*>(errorMethod));
	}
}

char QDBusAbstractInterface_CallWithCallback2(void* ptr, char* method, void* args, void* receiver, char* slot)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(receiver))) {
		return static_cast<QDBusAbstractInterface*>(ptr)->callWithCallback(QString(method), *static_cast<QList<QVariant>*>(args), static_cast<QDBusPendingCallWatcher*>(receiver), const_cast<const char*>(slot));
	} else {
		return static_cast<QDBusAbstractInterface*>(ptr)->callWithCallback(QString(method), *static_cast<QList<QVariant>*>(args), static_cast<QObject*>(receiver), const_cast<const char*>(slot));
	}
}

void QDBusAbstractInterface_SetTimeout(void* ptr, int timeout)
{
	static_cast<QDBusAbstractInterface*>(ptr)->setTimeout(timeout);
}

void QDBusAbstractInterface_DestroyQDBusAbstractInterface(void* ptr)
{
	static_cast<QDBusAbstractInterface*>(ptr)->~QDBusAbstractInterface();
}

void QDBusAbstractInterface_DestroyQDBusAbstractInterfaceDefault(void* ptr)
{

}

void* QDBusAbstractInterface_Connection(void* ptr)
{
	return new QDBusConnection(static_cast<QDBusAbstractInterface*>(ptr)->connection());
}

void* QDBusAbstractInterface_LastError(void* ptr)
{
	return new QDBusError(static_cast<QDBusAbstractInterface*>(ptr)->lastError());
}

struct QtDBus_PackedString QDBusAbstractInterface_Interface(void* ptr)
{
	return ({ QByteArray t8d0ae4 = static_cast<QDBusAbstractInterface*>(ptr)->interface().toUtf8(); QtDBus_PackedString { const_cast<char*>(t8d0ae4.prepend("WHITESPACE").constData()+10), t8d0ae4.size()-10 }; });
}

struct QtDBus_PackedString QDBusAbstractInterface_Path(void* ptr)
{
	return ({ QByteArray tff3773 = static_cast<QDBusAbstractInterface*>(ptr)->path().toUtf8(); QtDBus_PackedString { const_cast<char*>(tff3773.prepend("WHITESPACE").constData()+10), tff3773.size()-10 }; });
}

struct QtDBus_PackedString QDBusAbstractInterface_Service(void* ptr)
{
	return ({ QByteArray t0c0f21 = static_cast<QDBusAbstractInterface*>(ptr)->service().toUtf8(); QtDBus_PackedString { const_cast<char*>(t0c0f21.prepend("WHITESPACE").constData()+10), t0c0f21.size()-10 }; });
}

char QDBusAbstractInterface_IsValid(void* ptr)
{
	return static_cast<QDBusAbstractInterface*>(ptr)->isValid();
}

int QDBusAbstractInterface_Timeout(void* ptr)
{
	return static_cast<QDBusAbstractInterface*>(ptr)->timeout();
}

void* QDBusAbstractInterface___callWithArgumentList_args_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QList<QVariant>*>(ptr)->at(i));
}

void QDBusAbstractInterface___callWithArgumentList_args_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QDBusAbstractInterface___callWithArgumentList_args_newList(void* ptr)
{
	return new QList<QVariant>;
}

void* QDBusAbstractInterface___asyncCallWithArgumentList_args_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QList<QVariant>*>(ptr)->at(i));
}

void QDBusAbstractInterface___asyncCallWithArgumentList_args_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QDBusAbstractInterface___asyncCallWithArgumentList_args_newList(void* ptr)
{
	return new QList<QVariant>;
}

void* QDBusAbstractInterface___callWithCallback_args_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QList<QVariant>*>(ptr)->at(i));
}

void QDBusAbstractInterface___callWithCallback_args_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QDBusAbstractInterface___callWithCallback_args_newList(void* ptr)
{
	return new QList<QVariant>;
}

void* QDBusAbstractInterface___callWithCallback_args_atList2(void* ptr, int i)
{
	return new QVariant(static_cast<QList<QVariant>*>(ptr)->at(i));
}

void QDBusAbstractInterface___callWithCallback_args_setList2(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QDBusAbstractInterface___callWithCallback_args_newList2(void* ptr)
{
	return new QList<QVariant>;
}

void* QDBusAbstractInterface___internalConstCall_args_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QList<QVariant>*>(ptr)->at(i));
}

void QDBusAbstractInterface___internalConstCall_args_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QDBusAbstractInterface___internalConstCall_args_newList(void* ptr)
{
	return new QList<QVariant>;
}

void* QDBusAbstractInterface___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QDBusAbstractInterface___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QDBusAbstractInterface___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QDBusAbstractInterface___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDBusAbstractInterface___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusAbstractInterface___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QDBusAbstractInterface___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDBusAbstractInterface___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusAbstractInterface___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QDBusAbstractInterface___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDBusAbstractInterface___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusAbstractInterface___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* QDBusAbstractInterface___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QDBusAbstractInterface___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusAbstractInterface___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

char QDBusAbstractInterface_EventDefault(void* ptr, void* e)
{
	if (dynamic_cast<QDBusInterface*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDBusInterface*>(ptr)->QDBusInterface::event(static_cast<QEvent*>(e));
	} else if (dynamic_cast<QDBusConnectionInterface*>(static_cast<QObject*>(ptr))) {
		return static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::event(static_cast<QEvent*>(e));
	} else {
		return static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::event(static_cast<QEvent*>(e));
	}
}

char QDBusAbstractInterface_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QDBusInterface*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(watched))) {
			return static_cast<QDBusInterface*>(ptr)->QDBusInterface::eventFilter(static_cast<QDBusPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QDBusInterface*>(ptr)->QDBusInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else if (dynamic_cast<QDBusConnectionInterface*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(watched))) {
			return static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::eventFilter(static_cast<QDBusPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	} else {
		if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(watched))) {
			return static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::eventFilter(static_cast<QDBusPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
		} else {
			return static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
		}
	}
}

void QDBusAbstractInterface_ChildEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QDBusInterface*>(static_cast<QObject*>(ptr))) {
		static_cast<QDBusInterface*>(ptr)->QDBusInterface::childEvent(static_cast<QChildEvent*>(event));
	} else if (dynamic_cast<QDBusConnectionInterface*>(static_cast<QObject*>(ptr))) {
		static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::childEvent(static_cast<QChildEvent*>(event));
	} else {
		static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::childEvent(static_cast<QChildEvent*>(event));
	}
}

void QDBusAbstractInterface_ConnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QDBusInterface*>(static_cast<QObject*>(ptr))) {
		static_cast<QDBusInterface*>(ptr)->QDBusInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QDBusConnectionInterface*>(static_cast<QObject*>(ptr))) {
		static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::connectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QDBusAbstractInterface_CustomEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QDBusInterface*>(static_cast<QObject*>(ptr))) {
		static_cast<QDBusInterface*>(ptr)->QDBusInterface::customEvent(static_cast<QEvent*>(event));
	} else if (dynamic_cast<QDBusConnectionInterface*>(static_cast<QObject*>(ptr))) {
		static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::customEvent(static_cast<QEvent*>(event));
	} else {
		static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::customEvent(static_cast<QEvent*>(event));
	}
}

void QDBusAbstractInterface_DeleteLaterDefault(void* ptr)
{
	if (dynamic_cast<QDBusInterface*>(static_cast<QObject*>(ptr))) {
		static_cast<QDBusInterface*>(ptr)->QDBusInterface::deleteLater();
	} else if (dynamic_cast<QDBusConnectionInterface*>(static_cast<QObject*>(ptr))) {
		static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::deleteLater();
	} else {
		static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::deleteLater();
	}
}

void QDBusAbstractInterface_DisconnectNotifyDefault(void* ptr, void* sign)
{
	if (dynamic_cast<QDBusInterface*>(static_cast<QObject*>(ptr))) {
		static_cast<QDBusInterface*>(ptr)->QDBusInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else if (dynamic_cast<QDBusConnectionInterface*>(static_cast<QObject*>(ptr))) {
		static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	} else {
		static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::disconnectNotify(*static_cast<QMetaMethod*>(sign));
	}
}

void QDBusAbstractInterface_TimerEventDefault(void* ptr, void* event)
{
	if (dynamic_cast<QDBusInterface*>(static_cast<QObject*>(ptr))) {
		static_cast<QDBusInterface*>(ptr)->QDBusInterface::timerEvent(static_cast<QTimerEvent*>(event));
	} else if (dynamic_cast<QDBusConnectionInterface*>(static_cast<QObject*>(ptr))) {
		static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::timerEvent(static_cast<QTimerEvent*>(event));
	} else {
		static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::timerEvent(static_cast<QTimerEvent*>(event));
	}
}

void* QDBusAbstractInterface_MetaObjectDefault(void* ptr)
{
	if (dynamic_cast<QDBusInterface*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QDBusInterface*>(ptr)->QDBusInterface::metaObject());
	} else if (dynamic_cast<QDBusConnectionInterface*>(static_cast<QObject*>(ptr))) {
		return const_cast<QMetaObject*>(static_cast<QDBusConnectionInterface*>(ptr)->QDBusConnectionInterface::metaObject());
	} else {
		return const_cast<QMetaObject*>(static_cast<QDBusAbstractInterface*>(ptr)->QDBusAbstractInterface::metaObject());
	}
}

void* QDBusArgument_NewQDBusArgument()
{
	return new QDBusArgument();
}

void* QDBusArgument_NewQDBusArgument3(void* other)
{
	return new QDBusArgument(*static_cast<QDBusArgument*>(other));
}

void* QDBusArgument_NewQDBusArgument2(void* other)
{
	return new QDBusArgument(*static_cast<QDBusArgument*>(other));
}

void QDBusArgument_BeginArray(void* ptr, int id)
{
	static_cast<QDBusArgument*>(ptr)->beginArray(id);
}

void QDBusArgument_BeginMap(void* ptr, int kid, int vid)
{
	static_cast<QDBusArgument*>(ptr)->beginMap(kid, vid);
}

void QDBusArgument_BeginMapEntry(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->beginMapEntry();
}

void QDBusArgument_BeginStructure(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->beginStructure();
}

void QDBusArgument_EndArray(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->endArray();
}

void QDBusArgument_EndMap(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->endMap();
}

void QDBusArgument_EndMapEntry(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->endMapEntry();
}

void QDBusArgument_EndStructure(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->endStructure();
}

void QDBusArgument_Swap(void* ptr, void* other)
{
	static_cast<QDBusArgument*>(ptr)->swap(*static_cast<QDBusArgument*>(other));
}

void QDBusArgument_DestroyQDBusArgument(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->~QDBusArgument();
}

long long QDBusArgument_CurrentType(void* ptr)
{
	return static_cast<QDBusArgument*>(ptr)->currentType();
}

void* QDBusArgument_AsVariant(void* ptr)
{
	return new QVariant(static_cast<QDBusArgument*>(ptr)->asVariant());
}

char QDBusArgument_AtEnd(void* ptr)
{
	return static_cast<QDBusArgument*>(ptr)->atEnd();
}

void QDBusArgument_BeginArray2(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->beginArray();
}

void QDBusArgument_BeginMap2(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->beginMap();
}

void QDBusArgument_BeginMapEntry2(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->beginMapEntry();
}

void QDBusArgument_BeginStructure2(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->beginStructure();
}

void QDBusArgument_EndArray2(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->endArray();
}

void QDBusArgument_EndMap2(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->endMap();
}

void QDBusArgument_EndMapEntry2(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->endMapEntry();
}

void QDBusArgument_EndStructure2(void* ptr)
{
	static_cast<QDBusArgument*>(ptr)->endStructure();
}

void* QDBusConnection_QDBusConnection_ConnectToPeer(char* address, char* name)
{
	return new QDBusConnection(QDBusConnection::connectToPeer(QString(address), QString(name)));
}

void* QDBusConnection_QDBusConnection_SessionBus()
{
	return new QDBusConnection(QDBusConnection::sessionBus());
}

void* QDBusConnection_QDBusConnection_SystemBus()
{
	return new QDBusConnection(QDBusConnection::systemBus());
}

void* QDBusConnection_QDBusConnection_LocalMachineId()
{
	return new QByteArray(QDBusConnection::localMachineId());
}

void* QDBusConnection_QDBusConnection_ConnectToBus(long long ty, char* name)
{
	return new QDBusConnection(QDBusConnection::connectToBus(static_cast<QDBusConnection::BusType>(ty), QString(name)));
}

void* QDBusConnection_QDBusConnection_ConnectToBus2(char* address, char* name)
{
	return new QDBusConnection(QDBusConnection::connectToBus(QString(address), QString(name)));
}

void* QDBusConnection_NewQDBusConnection3(void* other)
{
	return new QDBusConnection(*static_cast<QDBusConnection*>(other));
}

void* QDBusConnection_NewQDBusConnection2(void* other)
{
	return new QDBusConnection(*static_cast<QDBusConnection*>(other));
}

void* QDBusConnection_NewQDBusConnection(char* name)
{
	return new QDBusConnection(QString(name));
}

char QDBusConnection_Connect(void* ptr, char* service, char* path, char* interfa, char* name, void* receiver, char* slot)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(receiver))) {
		return static_cast<QDBusConnection*>(ptr)->connect(QString(service), QString(path), QString(interfa), QString(name), static_cast<QDBusPendingCallWatcher*>(receiver), const_cast<const char*>(slot));
	} else {
		return static_cast<QDBusConnection*>(ptr)->connect(QString(service), QString(path), QString(interfa), QString(name), static_cast<QObject*>(receiver), const_cast<const char*>(slot));
	}
}

char QDBusConnection_Connect2(void* ptr, char* service, char* path, char* interfa, char* name, char* signature, void* receiver, char* slot)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(receiver))) {
		return static_cast<QDBusConnection*>(ptr)->connect(QString(service), QString(path), QString(interfa), QString(name), QString(signature), static_cast<QDBusPendingCallWatcher*>(receiver), const_cast<const char*>(slot));
	} else {
		return static_cast<QDBusConnection*>(ptr)->connect(QString(service), QString(path), QString(interfa), QString(name), QString(signature), static_cast<QObject*>(receiver), const_cast<const char*>(slot));
	}
}

char QDBusConnection_Connect3(void* ptr, char* service, char* path, char* interfa, char* name, char* argumentMatch, char* signature, void* receiver, char* slot)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(receiver))) {
		return static_cast<QDBusConnection*>(ptr)->connect(QString(service), QString(path), QString(interfa), QString(name), QString(argumentMatch).split("|", QString::SkipEmptyParts), QString(signature), static_cast<QDBusPendingCallWatcher*>(receiver), const_cast<const char*>(slot));
	} else {
		return static_cast<QDBusConnection*>(ptr)->connect(QString(service), QString(path), QString(interfa), QString(name), QString(argumentMatch).split("|", QString::SkipEmptyParts), QString(signature), static_cast<QObject*>(receiver), const_cast<const char*>(slot));
	}
}

char QDBusConnection_Disconnect(void* ptr, char* service, char* path, char* interfa, char* name, void* receiver, char* slot)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(receiver))) {
		return static_cast<QDBusConnection*>(ptr)->disconnect(QString(service), QString(path), QString(interfa), QString(name), static_cast<QDBusPendingCallWatcher*>(receiver), const_cast<const char*>(slot));
	} else {
		return static_cast<QDBusConnection*>(ptr)->disconnect(QString(service), QString(path), QString(interfa), QString(name), static_cast<QObject*>(receiver), const_cast<const char*>(slot));
	}
}

char QDBusConnection_Disconnect2(void* ptr, char* service, char* path, char* interfa, char* name, char* signature, void* receiver, char* slot)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(receiver))) {
		return static_cast<QDBusConnection*>(ptr)->disconnect(QString(service), QString(path), QString(interfa), QString(name), QString(signature), static_cast<QDBusPendingCallWatcher*>(receiver), const_cast<const char*>(slot));
	} else {
		return static_cast<QDBusConnection*>(ptr)->disconnect(QString(service), QString(path), QString(interfa), QString(name), QString(signature), static_cast<QObject*>(receiver), const_cast<const char*>(slot));
	}
}

char QDBusConnection_Disconnect3(void* ptr, char* service, char* path, char* interfa, char* name, char* argumentMatch, char* signature, void* receiver, char* slot)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(receiver))) {
		return static_cast<QDBusConnection*>(ptr)->disconnect(QString(service), QString(path), QString(interfa), QString(name), QString(argumentMatch).split("|", QString::SkipEmptyParts), QString(signature), static_cast<QDBusPendingCallWatcher*>(receiver), const_cast<const char*>(slot));
	} else {
		return static_cast<QDBusConnection*>(ptr)->disconnect(QString(service), QString(path), QString(interfa), QString(name), QString(argumentMatch).split("|", QString::SkipEmptyParts), QString(signature), static_cast<QObject*>(receiver), const_cast<const char*>(slot));
	}
}

char QDBusConnection_RegisterObject(void* ptr, char* path, void* object, long long options)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(object))) {
		return static_cast<QDBusConnection*>(ptr)->registerObject(QString(path), static_cast<QDBusPendingCallWatcher*>(object), static_cast<QDBusConnection::RegisterOption>(options));
	} else {
		return static_cast<QDBusConnection*>(ptr)->registerObject(QString(path), static_cast<QObject*>(object), static_cast<QDBusConnection::RegisterOption>(options));
	}
}

char QDBusConnection_RegisterObject2(void* ptr, char* path, char* interfa, void* object, long long options)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(object))) {
		return static_cast<QDBusConnection*>(ptr)->registerObject(QString(path), QString(interfa), static_cast<QDBusPendingCallWatcher*>(object), static_cast<QDBusConnection::RegisterOption>(options));
	} else {
		return static_cast<QDBusConnection*>(ptr)->registerObject(QString(path), QString(interfa), static_cast<QObject*>(object), static_cast<QDBusConnection::RegisterOption>(options));
	}
}

char QDBusConnection_RegisterService(void* ptr, char* serviceName)
{
	return static_cast<QDBusConnection*>(ptr)->registerService(QString(serviceName));
}

char QDBusConnection_UnregisterService(void* ptr, char* serviceName)
{
	return static_cast<QDBusConnection*>(ptr)->unregisterService(QString(serviceName));
}

void QDBusConnection_QDBusConnection_DisconnectFromBus(char* name)
{
	QDBusConnection::disconnectFromBus(QString(name));
}

void QDBusConnection_QDBusConnection_DisconnectFromPeer(char* name)
{
	QDBusConnection::disconnectFromPeer(QString(name));
}

void QDBusConnection_Swap(void* ptr, void* other)
{
	static_cast<QDBusConnection*>(ptr)->swap(*static_cast<QDBusConnection*>(other));
}

void QDBusConnection_UnregisterObject(void* ptr, char* path, long long mode)
{
	static_cast<QDBusConnection*>(ptr)->unregisterObject(QString(path), static_cast<QDBusConnection::UnregisterMode>(mode));
}

void QDBusConnection_DestroyQDBusConnection(void* ptr)
{
	static_cast<QDBusConnection*>(ptr)->~QDBusConnection();
}

long long QDBusConnection_ConnectionCapabilities(void* ptr)
{
	return static_cast<QDBusConnection*>(ptr)->connectionCapabilities();
}

void* QDBusConnection_Interface(void* ptr)
{
	return static_cast<QDBusConnection*>(ptr)->interface();
}

void* QDBusConnection_LastError(void* ptr)
{
	return new QDBusError(static_cast<QDBusConnection*>(ptr)->lastError());
}

void* QDBusConnection_Call(void* ptr, void* message, long long mode, int timeout)
{
	return new QDBusMessage(static_cast<QDBusConnection*>(ptr)->call(*static_cast<QDBusMessage*>(message), static_cast<QDBus::CallMode>(mode), timeout));
}

void* QDBusConnection_AsyncCall(void* ptr, void* message, int timeout)
{
	return new QDBusPendingCall(static_cast<QDBusConnection*>(ptr)->asyncCall(*static_cast<QDBusMessage*>(message), timeout));
}

void* QDBusConnection_ObjectRegisteredAt(void* ptr, char* path)
{
	return static_cast<QDBusConnection*>(ptr)->objectRegisteredAt(QString(path));
}

struct QtDBus_PackedString QDBusConnection_BaseService(void* ptr)
{
	return ({ QByteArray tcfa988 = static_cast<QDBusConnection*>(ptr)->baseService().toUtf8(); QtDBus_PackedString { const_cast<char*>(tcfa988.prepend("WHITESPACE").constData()+10), tcfa988.size()-10 }; });
}

struct QtDBus_PackedString QDBusConnection_Name(void* ptr)
{
	return ({ QByteArray tc0189b = static_cast<QDBusConnection*>(ptr)->name().toUtf8(); QtDBus_PackedString { const_cast<char*>(tc0189b.prepend("WHITESPACE").constData()+10), tc0189b.size()-10 }; });
}

char QDBusConnection_CallWithCallback(void* ptr, void* message, void* receiver, char* returnMethod, char* errorMethod, int timeout)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(receiver))) {
		return static_cast<QDBusConnection*>(ptr)->callWithCallback(*static_cast<QDBusMessage*>(message), static_cast<QDBusPendingCallWatcher*>(receiver), const_cast<const char*>(returnMethod), const_cast<const char*>(errorMethod), timeout);
	} else {
		return static_cast<QDBusConnection*>(ptr)->callWithCallback(*static_cast<QDBusMessage*>(message), static_cast<QObject*>(receiver), const_cast<const char*>(returnMethod), const_cast<const char*>(errorMethod), timeout);
	}
}

char QDBusConnection_IsConnected(void* ptr)
{
	return static_cast<QDBusConnection*>(ptr)->isConnected();
}

char QDBusConnection_Send(void* ptr, void* message)
{
	return static_cast<QDBusConnection*>(ptr)->send(*static_cast<QDBusMessage*>(message));
}

class MyQDBusConnectionInterface: public QDBusConnectionInterface
{
public:
	void Signal_ServiceRegistered(const QString & serviceName) { QByteArray tc151ab = serviceName.toUtf8(); QtDBus_PackedString serviceNamePacked = { const_cast<char*>(tc151ab.prepend("WHITESPACE").constData()+10), tc151ab.size()-10 };callbackQDBusConnectionInterface_ServiceRegistered(this, serviceNamePacked); };
	void Signal_CallWithCallbackFailed(const QDBusError & error, const QDBusMessage & call) { callbackQDBusConnectionInterface_CallWithCallbackFailed(this, const_cast<QDBusError*>(&error), const_cast<QDBusMessage*>(&call)); };
	void Signal_ServiceUnregistered(const QString & serviceName) { QByteArray tc151ab = serviceName.toUtf8(); QtDBus_PackedString serviceNamePacked = { const_cast<char*>(tc151ab.prepend("WHITESPACE").constData()+10), tc151ab.size()-10 };callbackQDBusConnectionInterface_ServiceUnregistered(this, serviceNamePacked); };
	bool event(QEvent * e) { return callbackQDBusAbstractInterface_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDBusAbstractInterface_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQDBusAbstractInterface_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDBusAbstractInterface_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDBusAbstractInterface_CustomEvent(this, event); };
	void deleteLater() { callbackQDBusAbstractInterface_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQDBusAbstractInterface_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDBusAbstractInterface_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtDBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQDBusAbstractInterface_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQDBusAbstractInterface_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDBusAbstractInterface_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

void QDBusConnectionInterface_ConnectServiceRegistered(void* ptr)
{
	QObject::connect(static_cast<QDBusConnectionInterface*>(ptr), static_cast<void (QDBusConnectionInterface::*)(const QString &)>(&QDBusConnectionInterface::serviceRegistered), static_cast<MyQDBusConnectionInterface*>(ptr), static_cast<void (MyQDBusConnectionInterface::*)(const QString &)>(&MyQDBusConnectionInterface::Signal_ServiceRegistered));
}

void QDBusConnectionInterface_DisconnectServiceRegistered(void* ptr)
{
	QObject::disconnect(static_cast<QDBusConnectionInterface*>(ptr), static_cast<void (QDBusConnectionInterface::*)(const QString &)>(&QDBusConnectionInterface::serviceRegistered), static_cast<MyQDBusConnectionInterface*>(ptr), static_cast<void (MyQDBusConnectionInterface::*)(const QString &)>(&MyQDBusConnectionInterface::Signal_ServiceRegistered));
}

void QDBusConnectionInterface_ServiceRegistered(void* ptr, char* serviceName)
{
	static_cast<QDBusConnectionInterface*>(ptr)->serviceRegistered(QString(serviceName));
}

void QDBusConnectionInterface_ConnectCallWithCallbackFailed(void* ptr)
{
	QObject::connect(static_cast<QDBusConnectionInterface*>(ptr), static_cast<void (QDBusConnectionInterface::*)(const QDBusError &, const QDBusMessage &)>(&QDBusConnectionInterface::callWithCallbackFailed), static_cast<MyQDBusConnectionInterface*>(ptr), static_cast<void (MyQDBusConnectionInterface::*)(const QDBusError &, const QDBusMessage &)>(&MyQDBusConnectionInterface::Signal_CallWithCallbackFailed));
}

void QDBusConnectionInterface_DisconnectCallWithCallbackFailed(void* ptr)
{
	QObject::disconnect(static_cast<QDBusConnectionInterface*>(ptr), static_cast<void (QDBusConnectionInterface::*)(const QDBusError &, const QDBusMessage &)>(&QDBusConnectionInterface::callWithCallbackFailed), static_cast<MyQDBusConnectionInterface*>(ptr), static_cast<void (MyQDBusConnectionInterface::*)(const QDBusError &, const QDBusMessage &)>(&MyQDBusConnectionInterface::Signal_CallWithCallbackFailed));
}

void QDBusConnectionInterface_CallWithCallbackFailed(void* ptr, void* error, void* call)
{
	static_cast<QDBusConnectionInterface*>(ptr)->callWithCallbackFailed(*static_cast<QDBusError*>(error), *static_cast<QDBusMessage*>(call));
}

void QDBusConnectionInterface_ConnectServiceUnregistered(void* ptr)
{
	QObject::connect(static_cast<QDBusConnectionInterface*>(ptr), static_cast<void (QDBusConnectionInterface::*)(const QString &)>(&QDBusConnectionInterface::serviceUnregistered), static_cast<MyQDBusConnectionInterface*>(ptr), static_cast<void (MyQDBusConnectionInterface::*)(const QString &)>(&MyQDBusConnectionInterface::Signal_ServiceUnregistered));
}

void QDBusConnectionInterface_DisconnectServiceUnregistered(void* ptr)
{
	QObject::disconnect(static_cast<QDBusConnectionInterface*>(ptr), static_cast<void (QDBusConnectionInterface::*)(const QString &)>(&QDBusConnectionInterface::serviceUnregistered), static_cast<MyQDBusConnectionInterface*>(ptr), static_cast<void (MyQDBusConnectionInterface::*)(const QString &)>(&MyQDBusConnectionInterface::Signal_ServiceUnregistered));
}

void QDBusConnectionInterface_ServiceUnregistered(void* ptr, char* serviceName)
{
	static_cast<QDBusConnectionInterface*>(ptr)->serviceUnregistered(QString(serviceName));
}

void* QDBusContext_NewQDBusContext()
{
	return new QDBusContext();
}

void QDBusContext_DestroyQDBusContext(void* ptr)
{
	static_cast<QDBusContext*>(ptr)->~QDBusContext();
}

void* QDBusContext_Connection(void* ptr)
{
	return new QDBusConnection(static_cast<QDBusContext*>(ptr)->connection());
}

char QDBusContext_CalledFromDBus(void* ptr)
{
	return static_cast<QDBusContext*>(ptr)->calledFromDBus();
}

char QDBusContext_IsDelayedReply(void* ptr)
{
	return static_cast<QDBusContext*>(ptr)->isDelayedReply();
}

void* QDBusContext_Message(void* ptr)
{
	return const_cast<QDBusMessage*>(&static_cast<QDBusContext*>(ptr)->message());
}

void QDBusContext_SendErrorReply2(void* ptr, long long ty, char* msg)
{
	static_cast<QDBusContext*>(ptr)->sendErrorReply(static_cast<QDBusError::ErrorType>(ty), QString(msg));
}

void QDBusContext_SendErrorReply(void* ptr, char* name, char* msg)
{
	static_cast<QDBusContext*>(ptr)->sendErrorReply(QString(name), QString(msg));
}

void QDBusContext_SetDelayedReply(void* ptr, char enable)
{
	static_cast<QDBusContext*>(ptr)->setDelayedReply(enable != 0);
}

void* QDBusError_NewQDBusError(void* other)
{
	return new QDBusError(*static_cast<QDBusError*>(other));
}

struct QtDBus_PackedString QDBusError_QDBusError_ErrorString(long long error)
{
	return ({ QByteArray t5f2f4f = QDBusError::errorString(static_cast<QDBusError::ErrorType>(error)).toUtf8(); QtDBus_PackedString { const_cast<char*>(t5f2f4f.prepend("WHITESPACE").constData()+10), t5f2f4f.size()-10 }; });
}

void QDBusError_Swap(void* ptr, void* other)
{
	static_cast<QDBusError*>(ptr)->swap(*static_cast<QDBusError*>(other));
}

long long QDBusError_Type(void* ptr)
{
	return static_cast<QDBusError*>(ptr)->type();
}

struct QtDBus_PackedString QDBusError_Message(void* ptr)
{
	return ({ QByteArray t8f05f3 = static_cast<QDBusError*>(ptr)->message().toUtf8(); QtDBus_PackedString { const_cast<char*>(t8f05f3.prepend("WHITESPACE").constData()+10), t8f05f3.size()-10 }; });
}

struct QtDBus_PackedString QDBusError_Name(void* ptr)
{
	return ({ QByteArray t93fa91 = static_cast<QDBusError*>(ptr)->name().toUtf8(); QtDBus_PackedString { const_cast<char*>(t93fa91.prepend("WHITESPACE").constData()+10), t93fa91.size()-10 }; });
}

char QDBusError_IsValid(void* ptr)
{
	return static_cast<QDBusError*>(ptr)->isValid();
}

class MyQDBusInterface: public QDBusInterface
{
public:
	MyQDBusInterface(const QString &service, const QString &path, const QString &interface = QString(), const QDBusConnection &connection = QDBusConnection::sessionBus(), QObject *parent = Q_NULLPTR) : QDBusInterface(service, path, interface, connection, parent) {};
	bool event(QEvent * e) { return callbackQDBusAbstractInterface_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDBusAbstractInterface_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQDBusAbstractInterface_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDBusAbstractInterface_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDBusAbstractInterface_CustomEvent(this, event); };
	void deleteLater() { callbackQDBusAbstractInterface_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQDBusAbstractInterface_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDBusAbstractInterface_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtDBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQDBusAbstractInterface_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQDBusAbstractInterface_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDBusAbstractInterface_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

void* QDBusInterface_NewQDBusInterface(char* service, char* path, char* interfa, void* connection, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QWindow*>(parent));
	} else {
		return new MyQDBusInterface(QString(service), QString(path), QString(interfa), *static_cast<QDBusConnection*>(connection), static_cast<QObject*>(parent));
	}
}

void QDBusInterface_DestroyQDBusInterface(void* ptr)
{
	static_cast<QDBusInterface*>(ptr)->~QDBusInterface();
}

void* QDBusMessage_QDBusMessage_CreateError3(long long ty, char* msg)
{
	return new QDBusMessage(QDBusMessage::createError(static_cast<QDBusError::ErrorType>(ty), QString(msg)));
}

void* QDBusMessage_QDBusMessage_CreateError2(void* error)
{
	return new QDBusMessage(QDBusMessage::createError(*static_cast<QDBusError*>(error)));
}

void* QDBusMessage_QDBusMessage_CreateError(char* name, char* msg)
{
	return new QDBusMessage(QDBusMessage::createError(QString(name), QString(msg)));
}

void* QDBusMessage_QDBusMessage_CreateMethodCall(char* service, char* path, char* interfa, char* method)
{
	return new QDBusMessage(QDBusMessage::createMethodCall(QString(service), QString(path), QString(interfa), QString(method)));
}

void* QDBusMessage_QDBusMessage_CreateSignal(char* path, char* interfa, char* name)
{
	return new QDBusMessage(QDBusMessage::createSignal(QString(path), QString(interfa), QString(name)));
}

void* QDBusMessage_QDBusMessage_CreateTargetedSignal(char* service, char* path, char* interfa, char* name)
{
	return new QDBusMessage(QDBusMessage::createTargetedSignal(QString(service), QString(path), QString(interfa), QString(name)));
}

void* QDBusMessage_NewQDBusMessage()
{
	return new QDBusMessage();
}

void* QDBusMessage_NewQDBusMessage2(void* other)
{
	return new QDBusMessage(*static_cast<QDBusMessage*>(other));
}

void QDBusMessage_SetArguments(void* ptr, void* arguments)
{
	static_cast<QDBusMessage*>(ptr)->setArguments(*static_cast<QList<QVariant>*>(arguments));
}

void QDBusMessage_SetAutoStartService(void* ptr, char enable)
{
	static_cast<QDBusMessage*>(ptr)->setAutoStartService(enable != 0);
}

void QDBusMessage_Swap(void* ptr, void* other)
{
	static_cast<QDBusMessage*>(ptr)->swap(*static_cast<QDBusMessage*>(other));
}

void QDBusMessage_DestroyQDBusMessage(void* ptr)
{
	static_cast<QDBusMessage*>(ptr)->~QDBusMessage();
}

long long QDBusMessage_Type(void* ptr)
{
	return static_cast<QDBusMessage*>(ptr)->type();
}

void* QDBusMessage_CreateErrorReply3(void* ptr, long long ty, char* msg)
{
	return new QDBusMessage(static_cast<QDBusMessage*>(ptr)->createErrorReply(static_cast<QDBusError::ErrorType>(ty), QString(msg)));
}

void* QDBusMessage_CreateErrorReply2(void* ptr, void* error)
{
	return new QDBusMessage(static_cast<QDBusMessage*>(ptr)->createErrorReply(*static_cast<QDBusError*>(error)));
}

void* QDBusMessage_CreateErrorReply(void* ptr, char* name, char* msg)
{
	return new QDBusMessage(static_cast<QDBusMessage*>(ptr)->createErrorReply(QString(name), QString(msg)));
}

void* QDBusMessage_CreateReply(void* ptr, void* arguments)
{
	return new QDBusMessage(static_cast<QDBusMessage*>(ptr)->createReply(*static_cast<QList<QVariant>*>(arguments)));
}

void* QDBusMessage_CreateReply2(void* ptr, void* argument)
{
	return new QDBusMessage(static_cast<QDBusMessage*>(ptr)->createReply(*static_cast<QVariant*>(argument)));
}

struct QtDBus_PackedList QDBusMessage_Arguments(void* ptr)
{
	return ({ QList<QVariant>* tmpValue = new QList<QVariant>(static_cast<QDBusMessage*>(ptr)->arguments()); QtDBus_PackedList { tmpValue, tmpValue->size() }; });
}

struct QtDBus_PackedString QDBusMessage_ErrorMessage(void* ptr)
{
	return ({ QByteArray tc477a7 = static_cast<QDBusMessage*>(ptr)->errorMessage().toUtf8(); QtDBus_PackedString { const_cast<char*>(tc477a7.prepend("WHITESPACE").constData()+10), tc477a7.size()-10 }; });
}

struct QtDBus_PackedString QDBusMessage_ErrorName(void* ptr)
{
	return ({ QByteArray td5d165 = static_cast<QDBusMessage*>(ptr)->errorName().toUtf8(); QtDBus_PackedString { const_cast<char*>(td5d165.prepend("WHITESPACE").constData()+10), td5d165.size()-10 }; });
}

struct QtDBus_PackedString QDBusMessage_Interface(void* ptr)
{
	return ({ QByteArray tf9de81 = static_cast<QDBusMessage*>(ptr)->interface().toUtf8(); QtDBus_PackedString { const_cast<char*>(tf9de81.prepend("WHITESPACE").constData()+10), tf9de81.size()-10 }; });
}

struct QtDBus_PackedString QDBusMessage_Member(void* ptr)
{
	return ({ QByteArray t2721ac = static_cast<QDBusMessage*>(ptr)->member().toUtf8(); QtDBus_PackedString { const_cast<char*>(t2721ac.prepend("WHITESPACE").constData()+10), t2721ac.size()-10 }; });
}

struct QtDBus_PackedString QDBusMessage_Path(void* ptr)
{
	return ({ QByteArray t513da4 = static_cast<QDBusMessage*>(ptr)->path().toUtf8(); QtDBus_PackedString { const_cast<char*>(t513da4.prepend("WHITESPACE").constData()+10), t513da4.size()-10 }; });
}

struct QtDBus_PackedString QDBusMessage_Service(void* ptr)
{
	return ({ QByteArray te30c43 = static_cast<QDBusMessage*>(ptr)->service().toUtf8(); QtDBus_PackedString { const_cast<char*>(te30c43.prepend("WHITESPACE").constData()+10), te30c43.size()-10 }; });
}

struct QtDBus_PackedString QDBusMessage_Signature(void* ptr)
{
	return ({ QByteArray tb0ac7d = static_cast<QDBusMessage*>(ptr)->signature().toUtf8(); QtDBus_PackedString { const_cast<char*>(tb0ac7d.prepend("WHITESPACE").constData()+10), tb0ac7d.size()-10 }; });
}

char QDBusMessage_AutoStartService(void* ptr)
{
	return static_cast<QDBusMessage*>(ptr)->autoStartService();
}

char QDBusMessage_IsDelayedReply(void* ptr)
{
	return static_cast<QDBusMessage*>(ptr)->isDelayedReply();
}

char QDBusMessage_IsReplyRequired(void* ptr)
{
	return static_cast<QDBusMessage*>(ptr)->isReplyRequired();
}

void QDBusMessage_SetDelayedReply(void* ptr, char enable)
{
	static_cast<QDBusMessage*>(ptr)->setDelayedReply(enable != 0);
}

void* QDBusMessage___setArguments_arguments_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QList<QVariant>*>(ptr)->at(i));
}

void QDBusMessage___setArguments_arguments_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QDBusMessage___setArguments_arguments_newList(void* ptr)
{
	return new QList<QVariant>;
}

void* QDBusMessage___createReply_arguments_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QList<QVariant>*>(ptr)->at(i));
}

void QDBusMessage___createReply_arguments_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QDBusMessage___createReply_arguments_newList(void* ptr)
{
	return new QList<QVariant>;
}

void* QDBusMessage___arguments_atList(void* ptr, int i)
{
	return new QVariant(static_cast<QList<QVariant>*>(ptr)->at(i));
}

void QDBusMessage___arguments_setList(void* ptr, void* i)
{
	static_cast<QList<QVariant>*>(ptr)->append(*static_cast<QVariant*>(i));
}

void* QDBusMessage___arguments_newList(void* ptr)
{
	return new QList<QVariant>;
}

void* QDBusObjectPath_NewQDBusObjectPath()
{
	return new QDBusObjectPath();
}

void* QDBusObjectPath_NewQDBusObjectPath3(void* path)
{
	return new QDBusObjectPath(*static_cast<QLatin1String*>(path));
}

void* QDBusObjectPath_NewQDBusObjectPath5(char* p)
{
	return new QDBusObjectPath(*(new QString(p)));
}

void* QDBusObjectPath_NewQDBusObjectPath4(char* path)
{
	return new QDBusObjectPath(QString(path));
}

void* QDBusObjectPath_NewQDBusObjectPath2(char* path)
{
	return new QDBusObjectPath(const_cast<const char*>(path));
}

void QDBusObjectPath_SetPath(void* ptr, char* path)
{
	static_cast<QDBusObjectPath*>(ptr)->setPath(QString(path));
}

void QDBusObjectPath_Swap(void* ptr, void* other)
{
	static_cast<QDBusObjectPath*>(ptr)->swap(*static_cast<QDBusObjectPath*>(other));
}

struct QtDBus_PackedString QDBusObjectPath_Path(void* ptr)
{
	return ({ QByteArray t6519e6 = static_cast<QDBusObjectPath*>(ptr)->path().toUtf8(); QtDBus_PackedString { const_cast<char*>(t6519e6.prepend("WHITESPACE").constData()+10), t6519e6.size()-10 }; });
}

void* QDBusPendingCall_QDBusPendingCall_FromCompletedCall(void* msg)
{
		return new QDBusPendingCall(QDBusPendingCall::fromCompletedCall(*static_cast<QDBusMessage*>(msg)));
}

void* QDBusPendingCall_QDBusPendingCall_FromError(void* error)
{
		return new QDBusPendingCall(QDBusPendingCall::fromError(*static_cast<QDBusError*>(error)));
}

void* QDBusPendingCall_NewQDBusPendingCall(void* other)
{
	return new QDBusPendingCall(*static_cast<QDBusPendingCall*>(other));
}

void QDBusPendingCall_Swap(void* ptr, void* other)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(ptr))) {
		if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(other))) {
			static_cast<QDBusPendingCallWatcher*>(ptr)->swap(*static_cast<QDBusPendingCallWatcher*>(other));
		} else {
			static_cast<QDBusPendingCallWatcher*>(ptr)->swap(*static_cast<QDBusPendingCall*>(other));
		}
	} else {
		if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(other))) {
			static_cast<QDBusPendingCall*>(ptr)->swap(*static_cast<QDBusPendingCallWatcher*>(other));
		} else {
			static_cast<QDBusPendingCall*>(ptr)->swap(*static_cast<QDBusPendingCall*>(other));
		}
	}
}

void QDBusPendingCall_DestroyQDBusPendingCall(void* ptr)
{
	static_cast<QDBusPendingCall*>(ptr)->~QDBusPendingCall();
}

class MyQDBusPendingCallWatcher: public QDBusPendingCallWatcher
{
public:
	MyQDBusPendingCallWatcher(const QDBusPendingCall &call, QObject *parent = Q_NULLPTR) : QDBusPendingCallWatcher(call, parent) {};
	void Signal_Finished(QDBusPendingCallWatcher * self) { callbackQDBusPendingCallWatcher_Finished(this, self); };
	bool event(QEvent * e) { return callbackQDBusPendingCallWatcher_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDBusPendingCallWatcher_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQDBusPendingCallWatcher_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDBusPendingCallWatcher_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDBusPendingCallWatcher_CustomEvent(this, event); };
	void deleteLater() { callbackQDBusPendingCallWatcher_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQDBusPendingCallWatcher_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDBusPendingCallWatcher_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtDBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQDBusPendingCallWatcher_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQDBusPendingCallWatcher_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDBusPendingCallWatcher_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

void* QDBusPendingCallWatcher_NewQDBusPendingCallWatcher(void* call, void* parent)
{
	return new MyQDBusPendingCallWatcher(*static_cast<QDBusPendingCall*>(call), static_cast<QObject*>(parent));
}

void QDBusPendingCallWatcher_ConnectFinished(void* ptr)
{
	QObject::connect(static_cast<QDBusPendingCallWatcher*>(ptr), static_cast<void (QDBusPendingCallWatcher::*)(QDBusPendingCallWatcher *)>(&QDBusPendingCallWatcher::finished), static_cast<MyQDBusPendingCallWatcher*>(ptr), static_cast<void (MyQDBusPendingCallWatcher::*)(QDBusPendingCallWatcher *)>(&MyQDBusPendingCallWatcher::Signal_Finished));
}

void QDBusPendingCallWatcher_DisconnectFinished(void* ptr)
{
	QObject::disconnect(static_cast<QDBusPendingCallWatcher*>(ptr), static_cast<void (QDBusPendingCallWatcher::*)(QDBusPendingCallWatcher *)>(&QDBusPendingCallWatcher::finished), static_cast<MyQDBusPendingCallWatcher*>(ptr), static_cast<void (MyQDBusPendingCallWatcher::*)(QDBusPendingCallWatcher *)>(&MyQDBusPendingCallWatcher::Signal_Finished));
}

void QDBusPendingCallWatcher_Finished(void* ptr, void* self)
{
		static_cast<QDBusPendingCallWatcher*>(ptr)->finished(static_cast<QDBusPendingCallWatcher*>(self));
}

void QDBusPendingCallWatcher_WaitForFinished(void* ptr)
{
		static_cast<QDBusPendingCallWatcher*>(ptr)->waitForFinished();
}

void QDBusPendingCallWatcher_DestroyQDBusPendingCallWatcher(void* ptr)
{
	static_cast<QDBusPendingCallWatcher*>(ptr)->~QDBusPendingCallWatcher();
}

char QDBusPendingCallWatcher_IsFinished(void* ptr)
{
		return static_cast<QDBusPendingCallWatcher*>(ptr)->isFinished();
}

void* QDBusPendingCallWatcher___dynamicPropertyNames_atList(void* ptr, int i)
{
		return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QDBusPendingCallWatcher___dynamicPropertyNames_setList(void* ptr, void* i)
{
		static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QDBusPendingCallWatcher___dynamicPropertyNames_newList(void* ptr)
{
		return new QList<QByteArray>;
}

void* QDBusPendingCallWatcher___findChildren_atList2(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDBusPendingCallWatcher___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusPendingCallWatcher___findChildren_newList2(void* ptr)
{
		return new QList<QObject*>;
}

void* QDBusPendingCallWatcher___findChildren_atList3(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDBusPendingCallWatcher___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusPendingCallWatcher___findChildren_newList3(void* ptr)
{
		return new QList<QObject*>;
}

void* QDBusPendingCallWatcher___findChildren_atList(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDBusPendingCallWatcher___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusPendingCallWatcher___findChildren_newList(void* ptr)
{
		return new QList<QObject*>;
}

void* QDBusPendingCallWatcher___children_atList(void* ptr, int i)
{
		return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QDBusPendingCallWatcher___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusPendingCallWatcher___children_newList(void* ptr)
{
		return new QList<QObject *>;
}

char QDBusPendingCallWatcher_Event(void* ptr, void* e)
{
		return static_cast<QDBusPendingCallWatcher*>(ptr)->event(static_cast<QEvent*>(e));
}

char QDBusPendingCallWatcher_EventDefault(void* ptr, void* e)
{
		return static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::event(static_cast<QEvent*>(e));
}

char QDBusPendingCallWatcher_EventFilter(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(watched))) {
		return static_cast<QDBusPendingCallWatcher*>(ptr)->eventFilter(static_cast<QDBusPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QDBusPendingCallWatcher*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

char QDBusPendingCallWatcher_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(watched))) {
		return static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::eventFilter(static_cast<QDBusPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QDBusPendingCallWatcher_ChildEvent(void* ptr, void* event)
{
		static_cast<QDBusPendingCallWatcher*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QDBusPendingCallWatcher_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::childEvent(static_cast<QChildEvent*>(event));
}

void QDBusPendingCallWatcher_ConnectNotify(void* ptr, void* sign)
{
		static_cast<QDBusPendingCallWatcher*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusPendingCallWatcher_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusPendingCallWatcher_CustomEvent(void* ptr, void* event)
{
		static_cast<QDBusPendingCallWatcher*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QDBusPendingCallWatcher_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::customEvent(static_cast<QEvent*>(event));
}

void QDBusPendingCallWatcher_DeleteLater(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QDBusPendingCallWatcher*>(ptr), "deleteLater");
}

void QDBusPendingCallWatcher_DeleteLaterDefault(void* ptr)
{
		static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::deleteLater();
}

void QDBusPendingCallWatcher_DisconnectNotify(void* ptr, void* sign)
{
		static_cast<QDBusPendingCallWatcher*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusPendingCallWatcher_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusPendingCallWatcher_TimerEvent(void* ptr, void* event)
{
		static_cast<QDBusPendingCallWatcher*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QDBusPendingCallWatcher_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QDBusPendingCallWatcher_MetaObject(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QDBusPendingCallWatcher*>(ptr)->metaObject());
}

void* QDBusPendingCallWatcher_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QDBusPendingCallWatcher*>(ptr)->QDBusPendingCallWatcher::metaObject());
}

int QDBusPendingReplyTypes_QDBusPendingReplyTypes_MetaTypeFor2(void* vqv)
{
	return QDBusPendingReplyTypes::metaTypeFor(static_cast<QVariant*>(vqv));
}

class MyQDBusServer: public QDBusServer
{
public:
	MyQDBusServer(QObject *parent = Q_NULLPTR) : QDBusServer(parent) {};
	MyQDBusServer(const QString &address, QObject *parent = Q_NULLPTR) : QDBusServer(address, parent) {};
	void Signal_NewConnection(const QDBusConnection & connection) { callbackQDBusServer_NewConnection(this, const_cast<QDBusConnection*>(&connection)); };
	 ~MyQDBusServer() { callbackQDBusServer_DestroyQDBusServer(this); };
	bool event(QEvent * e) { return callbackQDBusServer_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDBusServer_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQDBusServer_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDBusServer_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDBusServer_CustomEvent(this, event); };
	void deleteLater() { callbackQDBusServer_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQDBusServer_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDBusServer_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtDBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQDBusServer_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQDBusServer_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDBusServer_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

void* QDBusServer_NewQDBusServer2(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(static_cast<QWindow*>(parent));
	} else {
		return new MyQDBusServer(static_cast<QObject*>(parent));
	}
}

void* QDBusServer_NewQDBusServer(char* address, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(QString(address), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(QString(address), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(QString(address), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(QString(address), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(QString(address), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(QString(address), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(QString(address), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(QString(address), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(QString(address), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(QString(address), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(QString(address), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(QString(address), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(QString(address), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(QString(address), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(QString(address), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(QString(address), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServer(QString(address), static_cast<QWindow*>(parent));
	} else {
		return new MyQDBusServer(QString(address), static_cast<QObject*>(parent));
	}
}

void QDBusServer_ConnectNewConnection(void* ptr)
{
	QObject::connect(static_cast<QDBusServer*>(ptr), static_cast<void (QDBusServer::*)(const QDBusConnection &)>(&QDBusServer::newConnection), static_cast<MyQDBusServer*>(ptr), static_cast<void (MyQDBusServer::*)(const QDBusConnection &)>(&MyQDBusServer::Signal_NewConnection));
}

void QDBusServer_DisconnectNewConnection(void* ptr)
{
	QObject::disconnect(static_cast<QDBusServer*>(ptr), static_cast<void (QDBusServer::*)(const QDBusConnection &)>(&QDBusServer::newConnection), static_cast<MyQDBusServer*>(ptr), static_cast<void (MyQDBusServer::*)(const QDBusConnection &)>(&MyQDBusServer::Signal_NewConnection));
}

void QDBusServer_NewConnection(void* ptr, void* connection)
{
	static_cast<QDBusServer*>(ptr)->newConnection(*static_cast<QDBusConnection*>(connection));
}

void QDBusServer_SetAnonymousAuthenticationAllowed(void* ptr, char value)
{
	static_cast<QDBusServer*>(ptr)->setAnonymousAuthenticationAllowed(value != 0);
}

void QDBusServer_DestroyQDBusServer(void* ptr)
{
	static_cast<QDBusServer*>(ptr)->~QDBusServer();
}

void QDBusServer_DestroyQDBusServerDefault(void* ptr)
{

}

void* QDBusServer_LastError(void* ptr)
{
	return new QDBusError(static_cast<QDBusServer*>(ptr)->lastError());
}

struct QtDBus_PackedString QDBusServer_Address(void* ptr)
{
	return ({ QByteArray t8337f2 = static_cast<QDBusServer*>(ptr)->address().toUtf8(); QtDBus_PackedString { const_cast<char*>(t8337f2.prepend("WHITESPACE").constData()+10), t8337f2.size()-10 }; });
}

char QDBusServer_IsAnonymousAuthenticationAllowed(void* ptr)
{
	return static_cast<QDBusServer*>(ptr)->isAnonymousAuthenticationAllowed();
}

char QDBusServer_IsConnected(void* ptr)
{
	return static_cast<QDBusServer*>(ptr)->isConnected();
}

void* QDBusServer___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QDBusServer___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QDBusServer___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QDBusServer___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDBusServer___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusServer___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QDBusServer___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDBusServer___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusServer___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QDBusServer___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDBusServer___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusServer___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* QDBusServer___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QDBusServer___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusServer___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

char QDBusServer_EventDefault(void* ptr, void* e)
{
		return static_cast<QDBusServer*>(ptr)->QDBusServer::event(static_cast<QEvent*>(e));
}

char QDBusServer_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(watched))) {
		return static_cast<QDBusServer*>(ptr)->QDBusServer::eventFilter(static_cast<QDBusPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QDBusServer*>(ptr)->QDBusServer::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QDBusServer_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QDBusServer*>(ptr)->QDBusServer::childEvent(static_cast<QChildEvent*>(event));
}

void QDBusServer_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDBusServer*>(ptr)->QDBusServer::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusServer_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QDBusServer*>(ptr)->QDBusServer::customEvent(static_cast<QEvent*>(event));
}

void QDBusServer_DeleteLaterDefault(void* ptr)
{
		static_cast<QDBusServer*>(ptr)->QDBusServer::deleteLater();
}

void QDBusServer_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDBusServer*>(ptr)->QDBusServer::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusServer_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QDBusServer*>(ptr)->QDBusServer::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QDBusServer_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QDBusServer*>(ptr)->QDBusServer::metaObject());
}

class MyQDBusServiceWatcher: public QDBusServiceWatcher
{
public:
	MyQDBusServiceWatcher(QObject *parent = Q_NULLPTR) : QDBusServiceWatcher(parent) {};
	MyQDBusServiceWatcher(const QString &service, const QDBusConnection &connection, WatchMode watchMode = WatchForOwnerChange, QObject *parent = Q_NULLPTR) : QDBusServiceWatcher(service, connection, watchMode, parent) {};
	void Signal_ServiceRegistered(const QString & serviceName) { QByteArray tc151ab = serviceName.toUtf8(); QtDBus_PackedString serviceNamePacked = { const_cast<char*>(tc151ab.prepend("WHITESPACE").constData()+10), tc151ab.size()-10 };callbackQDBusServiceWatcher_ServiceRegistered(this, serviceNamePacked); };
	void Signal_ServiceOwnerChanged(const QString & serviceName, const QString & oldOwner, const QString & newOwner) { QByteArray tc151ab = serviceName.toUtf8(); QtDBus_PackedString serviceNamePacked = { const_cast<char*>(tc151ab.prepend("WHITESPACE").constData()+10), tc151ab.size()-10 };QByteArray t9456b1 = oldOwner.toUtf8(); QtDBus_PackedString oldOwnerPacked = { const_cast<char*>(t9456b1.prepend("WHITESPACE").constData()+10), t9456b1.size()-10 };QByteArray t157d45 = newOwner.toUtf8(); QtDBus_PackedString newOwnerPacked = { const_cast<char*>(t157d45.prepend("WHITESPACE").constData()+10), t157d45.size()-10 };callbackQDBusServiceWatcher_ServiceOwnerChanged(this, serviceNamePacked, oldOwnerPacked, newOwnerPacked); };
	void Signal_ServiceUnregistered(const QString & serviceName) { QByteArray tc151ab = serviceName.toUtf8(); QtDBus_PackedString serviceNamePacked = { const_cast<char*>(tc151ab.prepend("WHITESPACE").constData()+10), tc151ab.size()-10 };callbackQDBusServiceWatcher_ServiceUnregistered(this, serviceNamePacked); };
	bool event(QEvent * e) { return callbackQDBusServiceWatcher_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDBusServiceWatcher_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQDBusServiceWatcher_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDBusServiceWatcher_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDBusServiceWatcher_CustomEvent(this, event); };
	void deleteLater() { callbackQDBusServiceWatcher_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQDBusServiceWatcher_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDBusServiceWatcher_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtDBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQDBusServiceWatcher_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQDBusServiceWatcher_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDBusServiceWatcher_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

void QDBusServiceWatcher_ConnectServiceRegistered(void* ptr)
{
	QObject::connect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &)>(&QDBusServiceWatcher::serviceRegistered), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceRegistered));
}

void QDBusServiceWatcher_DisconnectServiceRegistered(void* ptr)
{
	QObject::disconnect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &)>(&QDBusServiceWatcher::serviceRegistered), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceRegistered));
}

void QDBusServiceWatcher_ServiceRegistered(void* ptr, char* serviceName)
{
	static_cast<QDBusServiceWatcher*>(ptr)->serviceRegistered(QString(serviceName));
}

void QDBusServiceWatcher_SetConnection(void* ptr, void* connection)
{
	static_cast<QDBusServiceWatcher*>(ptr)->setConnection(*static_cast<QDBusConnection*>(connection));
}

void QDBusServiceWatcher_SetWatchMode(void* ptr, long long mode)
{
	static_cast<QDBusServiceWatcher*>(ptr)->setWatchMode(static_cast<QDBusServiceWatcher::WatchModeFlag>(mode));
}

long long QDBusServiceWatcher_WatchMode(void* ptr)
{
	return static_cast<QDBusServiceWatcher*>(ptr)->watchMode();
}

void* QDBusServiceWatcher_NewQDBusServiceWatcher(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(static_cast<QWindow*>(parent));
	} else {
		return new MyQDBusServiceWatcher(static_cast<QObject*>(parent));
	}
}

void* QDBusServiceWatcher_NewQDBusServiceWatcher2(char* service, void* connection, long long watchMode, void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(QString(service), *static_cast<QDBusConnection*>(connection), static_cast<QDBusServiceWatcher::WatchModeFlag>(watchMode), static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(QString(service), *static_cast<QDBusConnection*>(connection), static_cast<QDBusServiceWatcher::WatchModeFlag>(watchMode), static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(QString(service), *static_cast<QDBusConnection*>(connection), static_cast<QDBusServiceWatcher::WatchModeFlag>(watchMode), static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(QString(service), *static_cast<QDBusConnection*>(connection), static_cast<QDBusServiceWatcher::WatchModeFlag>(watchMode), static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(QString(service), *static_cast<QDBusConnection*>(connection), static_cast<QDBusServiceWatcher::WatchModeFlag>(watchMode), static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(QString(service), *static_cast<QDBusConnection*>(connection), static_cast<QDBusServiceWatcher::WatchModeFlag>(watchMode), static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(QString(service), *static_cast<QDBusConnection*>(connection), static_cast<QDBusServiceWatcher::WatchModeFlag>(watchMode), static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(QString(service), *static_cast<QDBusConnection*>(connection), static_cast<QDBusServiceWatcher::WatchModeFlag>(watchMode), static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(QString(service), *static_cast<QDBusConnection*>(connection), static_cast<QDBusServiceWatcher::WatchModeFlag>(watchMode), static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(QString(service), *static_cast<QDBusConnection*>(connection), static_cast<QDBusServiceWatcher::WatchModeFlag>(watchMode), static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(QString(service), *static_cast<QDBusConnection*>(connection), static_cast<QDBusServiceWatcher::WatchModeFlag>(watchMode), static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(QString(service), *static_cast<QDBusConnection*>(connection), static_cast<QDBusServiceWatcher::WatchModeFlag>(watchMode), static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(QString(service), *static_cast<QDBusConnection*>(connection), static_cast<QDBusServiceWatcher::WatchModeFlag>(watchMode), static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(QString(service), *static_cast<QDBusConnection*>(connection), static_cast<QDBusServiceWatcher::WatchModeFlag>(watchMode), static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(QString(service), *static_cast<QDBusConnection*>(connection), static_cast<QDBusServiceWatcher::WatchModeFlag>(watchMode), static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(QString(service), *static_cast<QDBusConnection*>(connection), static_cast<QDBusServiceWatcher::WatchModeFlag>(watchMode), static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDBusServiceWatcher(QString(service), *static_cast<QDBusConnection*>(connection), static_cast<QDBusServiceWatcher::WatchModeFlag>(watchMode), static_cast<QWindow*>(parent));
	} else {
		return new MyQDBusServiceWatcher(QString(service), *static_cast<QDBusConnection*>(connection), static_cast<QDBusServiceWatcher::WatchModeFlag>(watchMode), static_cast<QObject*>(parent));
	}
}

char QDBusServiceWatcher_RemoveWatchedService(void* ptr, char* service)
{
	return static_cast<QDBusServiceWatcher*>(ptr)->removeWatchedService(QString(service));
}

void QDBusServiceWatcher_AddWatchedService(void* ptr, char* newService)
{
	static_cast<QDBusServiceWatcher*>(ptr)->addWatchedService(QString(newService));
}

void QDBusServiceWatcher_ConnectServiceOwnerChanged(void* ptr)
{
	QObject::connect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &, const QString &, const QString &)>(&QDBusServiceWatcher::serviceOwnerChanged), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &, const QString &, const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceOwnerChanged));
}

void QDBusServiceWatcher_DisconnectServiceOwnerChanged(void* ptr)
{
	QObject::disconnect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &, const QString &, const QString &)>(&QDBusServiceWatcher::serviceOwnerChanged), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &, const QString &, const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceOwnerChanged));
}

void QDBusServiceWatcher_ServiceOwnerChanged(void* ptr, char* serviceName, char* oldOwner, char* newOwner)
{
	static_cast<QDBusServiceWatcher*>(ptr)->serviceOwnerChanged(QString(serviceName), QString(oldOwner), QString(newOwner));
}

void QDBusServiceWatcher_ConnectServiceUnregistered(void* ptr)
{
	QObject::connect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &)>(&QDBusServiceWatcher::serviceUnregistered), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceUnregistered));
}

void QDBusServiceWatcher_DisconnectServiceUnregistered(void* ptr)
{
	QObject::disconnect(static_cast<QDBusServiceWatcher*>(ptr), static_cast<void (QDBusServiceWatcher::*)(const QString &)>(&QDBusServiceWatcher::serviceUnregistered), static_cast<MyQDBusServiceWatcher*>(ptr), static_cast<void (MyQDBusServiceWatcher::*)(const QString &)>(&MyQDBusServiceWatcher::Signal_ServiceUnregistered));
}

void QDBusServiceWatcher_ServiceUnregistered(void* ptr, char* serviceName)
{
	static_cast<QDBusServiceWatcher*>(ptr)->serviceUnregistered(QString(serviceName));
}

void QDBusServiceWatcher_SetWatchedServices(void* ptr, char* services)
{
	static_cast<QDBusServiceWatcher*>(ptr)->setWatchedServices(QString(services).split("|", QString::SkipEmptyParts));
}

void QDBusServiceWatcher_DestroyQDBusServiceWatcher(void* ptr)
{
	static_cast<QDBusServiceWatcher*>(ptr)->~QDBusServiceWatcher();
}

void* QDBusServiceWatcher_Connection(void* ptr)
{
	return new QDBusConnection(static_cast<QDBusServiceWatcher*>(ptr)->connection());
}

struct QtDBus_PackedString QDBusServiceWatcher_WatchedServices(void* ptr)
{
	return ({ QByteArray t4599ca = static_cast<QDBusServiceWatcher*>(ptr)->watchedServices().join("|").toUtf8(); QtDBus_PackedString { const_cast<char*>(t4599ca.prepend("WHITESPACE").constData()+10), t4599ca.size()-10 }; });
}

void* QDBusServiceWatcher___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QDBusServiceWatcher___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QDBusServiceWatcher___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QDBusServiceWatcher___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDBusServiceWatcher___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusServiceWatcher___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QDBusServiceWatcher___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDBusServiceWatcher___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusServiceWatcher___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QDBusServiceWatcher___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDBusServiceWatcher___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusServiceWatcher___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* QDBusServiceWatcher___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QDBusServiceWatcher___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusServiceWatcher___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

char QDBusServiceWatcher_EventDefault(void* ptr, void* e)
{
		return static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::event(static_cast<QEvent*>(e));
}

char QDBusServiceWatcher_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(watched))) {
		return static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::eventFilter(static_cast<QDBusPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QDBusServiceWatcher_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::childEvent(static_cast<QChildEvent*>(event));
}

void QDBusServiceWatcher_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusServiceWatcher_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::customEvent(static_cast<QEvent*>(event));
}

void QDBusServiceWatcher_DeleteLaterDefault(void* ptr)
{
		static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::deleteLater();
}

void QDBusServiceWatcher_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusServiceWatcher_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QDBusServiceWatcher_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QDBusServiceWatcher*>(ptr)->QDBusServiceWatcher::metaObject());
}

void* QDBusSignature_NewQDBusSignature()
{
	return new QDBusSignature();
}

void* QDBusSignature_NewQDBusSignature3(void* signature)
{
	return new QDBusSignature(*static_cast<QLatin1String*>(signature));
}

void* QDBusSignature_NewQDBusSignature5(char* sig)
{
	return new QDBusSignature(*(new QString(sig)));
}

void* QDBusSignature_NewQDBusSignature4(char* signature)
{
	return new QDBusSignature(QString(signature));
}

void* QDBusSignature_NewQDBusSignature2(char* signature)
{
	return new QDBusSignature(const_cast<const char*>(signature));
}

void QDBusSignature_SetSignature(void* ptr, char* signature)
{
	static_cast<QDBusSignature*>(ptr)->setSignature(QString(signature));
}

void QDBusSignature_Swap(void* ptr, void* other)
{
	static_cast<QDBusSignature*>(ptr)->swap(*static_cast<QDBusSignature*>(other));
}

struct QtDBus_PackedString QDBusSignature_Signature(void* ptr)
{
	return ({ QByteArray t4b8ef7 = static_cast<QDBusSignature*>(ptr)->signature().toUtf8(); QtDBus_PackedString { const_cast<char*>(t4b8ef7.prepend("WHITESPACE").constData()+10), t4b8ef7.size()-10 }; });
}

void* QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor()
{
	return new QDBusUnixFileDescriptor();
}

void* QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor3(void* other)
{
	return new QDBusUnixFileDescriptor(*static_cast<QDBusUnixFileDescriptor*>(other));
}

void* QDBusUnixFileDescriptor_NewQDBusUnixFileDescriptor2(int fileDescriptor)
{
	return new QDBusUnixFileDescriptor(fileDescriptor);
}

char QDBusUnixFileDescriptor_QDBusUnixFileDescriptor_IsSupported()
{
	return QDBusUnixFileDescriptor::isSupported();
}

void QDBusUnixFileDescriptor_SetFileDescriptor(void* ptr, int fileDescriptor)
{
	static_cast<QDBusUnixFileDescriptor*>(ptr)->setFileDescriptor(fileDescriptor);
}

void QDBusUnixFileDescriptor_Swap(void* ptr, void* other)
{
	static_cast<QDBusUnixFileDescriptor*>(ptr)->swap(*static_cast<QDBusUnixFileDescriptor*>(other));
}

void QDBusUnixFileDescriptor_DestroyQDBusUnixFileDescriptor(void* ptr)
{
	static_cast<QDBusUnixFileDescriptor*>(ptr)->~QDBusUnixFileDescriptor();
}

char QDBusUnixFileDescriptor_IsValid(void* ptr)
{
	return static_cast<QDBusUnixFileDescriptor*>(ptr)->isValid();
}

int QDBusUnixFileDescriptor_FileDescriptor(void* ptr)
{
	return static_cast<QDBusUnixFileDescriptor*>(ptr)->fileDescriptor();
}

void* QDBusVariant_NewQDBusVariant()
{
	return new QDBusVariant();
}

void* QDBusVariant_NewQDBusVariant3(void* v)
{
	return new QDBusVariant(*static_cast<QVariant*>(v));
}

void* QDBusVariant_NewQDBusVariant2(void* variant)
{
	return new QDBusVariant(*static_cast<QVariant*>(variant));
}

void QDBusVariant_SetVariant(void* ptr, void* variant)
{
	static_cast<QDBusVariant*>(ptr)->setVariant(*static_cast<QVariant*>(variant));
}

void QDBusVariant_Swap(void* ptr, void* other)
{
	static_cast<QDBusVariant*>(ptr)->swap(*static_cast<QDBusVariant*>(other));
}

void* QDBusVariant_Variant(void* ptr)
{
	return new QVariant(static_cast<QDBusVariant*>(ptr)->variant());
}

class MyQDBusVirtualObject: public QDBusVirtualObject
{
public:
	MyQDBusVirtualObject(QObject *parent = Q_NULLPTR) : QDBusVirtualObject(parent) {};
	bool handleMessage(const QDBusMessage & message, const QDBusConnection & connection) { return callbackQDBusVirtualObject_HandleMessage(this, const_cast<QDBusMessage*>(&message), const_cast<QDBusConnection*>(&connection)) != 0; };
	 ~MyQDBusVirtualObject() { callbackQDBusVirtualObject_DestroyQDBusVirtualObject(this); };
	QString introspect(const QString & path) const { QByteArray t3150ec = path.toUtf8(); QtDBus_PackedString pathPacked = { const_cast<char*>(t3150ec.prepend("WHITESPACE").constData()+10), t3150ec.size()-10 };return QString(callbackQDBusVirtualObject_Introspect(const_cast<void*>(static_cast<const void*>(this)), pathPacked)); };
	bool event(QEvent * e) { return callbackQDBusVirtualObject_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDBusVirtualObject_EventFilter(this, watched, event) != 0; };
	void childEvent(QChildEvent * event) { callbackQDBusVirtualObject_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDBusVirtualObject_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDBusVirtualObject_CustomEvent(this, event); };
	void deleteLater() { callbackQDBusVirtualObject_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQDBusVirtualObject_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDBusVirtualObject_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtDBus_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQDBusVirtualObject_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQDBusVirtualObject_TimerEvent(this, event); };
	const QMetaObject * metaObject() const { return static_cast<QMetaObject*>(callbackQDBusVirtualObject_MetaObject(const_cast<void*>(static_cast<const void*>(this)))); };
};

void* QDBusVirtualObject_NewQDBusVirtualObject(void* parent)
{
	if (dynamic_cast<QCameraImageCapture*>(static_cast<QObject*>(parent))) {
		return new MyQDBusVirtualObject(static_cast<QCameraImageCapture*>(parent));
	} else if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(parent))) {
		return new MyQDBusVirtualObject(static_cast<QDBusPendingCallWatcher*>(parent));
	} else if (dynamic_cast<QExtensionFactory*>(static_cast<QObject*>(parent))) {
		return new MyQDBusVirtualObject(static_cast<QExtensionFactory*>(parent));
	} else if (dynamic_cast<QExtensionManager*>(static_cast<QObject*>(parent))) {
		return new MyQDBusVirtualObject(static_cast<QExtensionManager*>(parent));
	} else if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(parent))) {
		return new MyQDBusVirtualObject(static_cast<QGraphicsObject*>(parent));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDBusVirtualObject(static_cast<QGraphicsWidget*>(parent));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(parent))) {
		return new MyQDBusVirtualObject(static_cast<QLayout*>(parent));
	} else if (dynamic_cast<QMediaPlaylist*>(static_cast<QObject*>(parent))) {
		return new MyQDBusVirtualObject(static_cast<QMediaPlaylist*>(parent));
	} else if (dynamic_cast<QMediaRecorder*>(static_cast<QObject*>(parent))) {
		return new MyQDBusVirtualObject(static_cast<QMediaRecorder*>(parent));
	} else if (dynamic_cast<QOffscreenSurface*>(static_cast<QObject*>(parent))) {
		return new MyQDBusVirtualObject(static_cast<QOffscreenSurface*>(parent));
	} else if (dynamic_cast<QPaintDeviceWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDBusVirtualObject(static_cast<QPaintDeviceWindow*>(parent));
	} else if (dynamic_cast<QPdfWriter*>(static_cast<QObject*>(parent))) {
		return new MyQDBusVirtualObject(static_cast<QPdfWriter*>(parent));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(parent))) {
		return new MyQDBusVirtualObject(static_cast<QQuickItem*>(parent));
	} else if (dynamic_cast<QRadioData*>(static_cast<QObject*>(parent))) {
		return new MyQDBusVirtualObject(static_cast<QRadioData*>(parent));
	} else if (dynamic_cast<QSignalSpy*>(static_cast<QObject*>(parent))) {
		return new MyQDBusVirtualObject(static_cast<QSignalSpy*>(parent));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(parent))) {
		return new MyQDBusVirtualObject(static_cast<QWidget*>(parent));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(parent))) {
		return new MyQDBusVirtualObject(static_cast<QWindow*>(parent));
	} else {
		return new MyQDBusVirtualObject(static_cast<QObject*>(parent));
	}
}

char QDBusVirtualObject_HandleMessage(void* ptr, void* message, void* connection)
{
	return static_cast<QDBusVirtualObject*>(ptr)->handleMessage(*static_cast<QDBusMessage*>(message), *static_cast<QDBusConnection*>(connection));
}

void QDBusVirtualObject_DestroyQDBusVirtualObject(void* ptr)
{
	static_cast<QDBusVirtualObject*>(ptr)->~QDBusVirtualObject();
}

void QDBusVirtualObject_DestroyQDBusVirtualObjectDefault(void* ptr)
{

}

struct QtDBus_PackedString QDBusVirtualObject_Introspect(void* ptr, char* path)
{
	return ({ QByteArray t027267 = static_cast<QDBusVirtualObject*>(ptr)->introspect(QString(path)).toUtf8(); QtDBus_PackedString { const_cast<char*>(t027267.prepend("WHITESPACE").constData()+10), t027267.size()-10 }; });
}

void* QDBusVirtualObject___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(static_cast<QList<QByteArray>*>(ptr)->at(i));
}

void QDBusVirtualObject___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QDBusVirtualObject___dynamicPropertyNames_newList(void* ptr)
{
	return new QList<QByteArray>;
}

void* QDBusVirtualObject___findChildren_atList2(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDBusVirtualObject___findChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusVirtualObject___findChildren_newList2(void* ptr)
{
	return new QList<QObject*>;
}

void* QDBusVirtualObject___findChildren_atList3(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDBusVirtualObject___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusVirtualObject___findChildren_newList3(void* ptr)
{
	return new QList<QObject*>;
}

void* QDBusVirtualObject___findChildren_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject*>*>(ptr)->at(i));
}

void QDBusVirtualObject___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusVirtualObject___findChildren_newList(void* ptr)
{
	return new QList<QObject*>;
}

void* QDBusVirtualObject___children_atList(void* ptr, int i)
{
	return const_cast<QObject*>(static_cast<QList<QObject *>*>(ptr)->at(i));
}

void QDBusVirtualObject___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QDBusPendingCallWatcher*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDBusVirtualObject___children_newList(void* ptr)
{
	return new QList<QObject *>;
}

char QDBusVirtualObject_EventDefault(void* ptr, void* e)
{
		return static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::event(static_cast<QEvent*>(e));
}

char QDBusVirtualObject_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QDBusPendingCallWatcher*>(static_cast<QObject*>(watched))) {
		return static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::eventFilter(static_cast<QDBusPendingCallWatcher*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QDBusVirtualObject_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::childEvent(static_cast<QChildEvent*>(event));
}

void QDBusVirtualObject_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusVirtualObject_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::customEvent(static_cast<QEvent*>(event));
}

void QDBusVirtualObject_DeleteLaterDefault(void* ptr)
{
		static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::deleteLater();
}

void QDBusVirtualObject_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDBusVirtualObject_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QDBusVirtualObject_MetaObjectDefault(void* ptr)
{
		return const_cast<QMetaObject*>(static_cast<QDBusVirtualObject*>(ptr)->QDBusVirtualObject::metaObject());
}

