// +build minimal

#define protected public
#define private public

#include "gui-minimal.h"
#include "_cgo_export.h"

#include <QBitmap>
#include <QBrush>
#include <QByteArray>
#include <QChildEvent>
#include <QColor>
#include <QDesktopServices>
#include <QDrag>
#include <QDragEnterEvent>
#include <QDragLeaveEvent>
#include <QDragMoveEvent>
#include <QDropEvent>
#include <QEvent>
#include <QFocusEvent>
#include <QFont>
#include <QGradient>
#include <QGraphicsObject>
#include <QGraphicsWidget>
#include <QGuiApplication>
#include <QHoverEvent>
#include <QIcon>
#include <QIconEngine>
#include <QImage>
#include <QInputEvent>
#include <QInputMethod>
#include <QInputMethodEvent>
#include <QKeyEvent>
#include <QKeySequence>
#include <QLatin1String>
#include <QLayout>
#include <QLine>
#include <QLineF>
#include <QMatrix>
#include <QMatrix4x4>
#include <QMetaMethod>
#include <QMetaObject>
#include <QMimeData>
#include <QMouseEvent>
#include <QMoveEvent>
#include <QObject>
#include <QPaintDevice>
#include <QPaintEngine>
#include <QPaintEngineState>
#include <QPaintEvent>
#include <QPainter>
#include <QPalette>
#include <QPixmap>
#include <QPoint>
#include <QPointF>
#include <QPolygon>
#include <QQuickItem>
#include <QRect>
#include <QRectF>
#include <QRegion>
#include <QRgba64>
#include <QScreen>
#include <QSize>
#include <QString>
#include <QSurface>
#include <QSurfaceFormat>
#include <QTimerEvent>
#include <QTouchDevice>
#include <QTouchEvent>
#include <QTransform>
#include <QUrl>
#include <QVariant>
#include <QVector>
#include <QVector2D>
#include <QVector3D>
#include <QVector4D>
#include <QWheelEvent>
#include <QWidget>
#include <QWindow>
#include <QStringList>

class MyQBitmap: public QBitmap
{
public:
	MyQBitmap() : QBitmap() {};
	MyQBitmap(const QPixmap &pixmap) : QBitmap(pixmap) {};
	MyQBitmap(int width, int height) : QBitmap(width, height) {};
	MyQBitmap(const QSize &size) : QBitmap(size) {};
	MyQBitmap(const QString &fileName, const char *format = Q_NULLPTR) : QBitmap(fileName, format) {};
	 ~MyQBitmap() { callbackQBitmap_DestroyQBitmap(this); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPixmap_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

void* QBitmap_NewQBitmap()
{
	return new MyQBitmap();
}

void* QBitmap_NewQBitmap2(void* pixmap)
{
	return new MyQBitmap(*static_cast<QPixmap*>(pixmap));
}

void* QBitmap_NewQBitmap3(int width, int height)
{
	return new MyQBitmap(width, height);
}

void* QBitmap_NewQBitmap4(void* size)
{
	return new MyQBitmap(*static_cast<QSize*>(size));
}

void* QBitmap_NewQBitmap5(struct QtGui_PackedString fileName, char* format)
{
	return new MyQBitmap(QString::fromUtf8(fileName.data, fileName.len), const_cast<const char*>(format));
}

void QBitmap_DestroyQBitmap(void* ptr)
{
	static_cast<QBitmap*>(ptr)->~QBitmap();
}

void QBitmap_DestroyQBitmapDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QBrush_NewQBrush()
{
	return new QBrush();
}

void* QBrush_NewQBrush2(long long style)
{
	return new QBrush(static_cast<Qt::BrushStyle>(style));
}

void* QBrush_NewQBrush3(void* color, long long style)
{
	return new QBrush(*static_cast<QColor*>(color), static_cast<Qt::BrushStyle>(style));
}

void* QBrush_NewQBrush4(long long color, long long style)
{
	return new QBrush(static_cast<Qt::GlobalColor>(color), static_cast<Qt::BrushStyle>(style));
}

void* QBrush_NewQBrush5(void* color, void* pixmap)
{
	return new QBrush(*static_cast<QColor*>(color), *static_cast<QPixmap*>(pixmap));
}

void* QBrush_NewQBrush6(long long color, void* pixmap)
{
	return new QBrush(static_cast<Qt::GlobalColor>(color), *static_cast<QPixmap*>(pixmap));
}

void* QBrush_NewQBrush7(void* pixmap)
{
	return new QBrush(*static_cast<QPixmap*>(pixmap));
}

void* QBrush_NewQBrush8(void* image)
{
	return new QBrush(*static_cast<QImage*>(image));
}

void* QBrush_NewQBrush9(void* other)
{
	return new QBrush(*static_cast<QBrush*>(other));
}

void* QBrush_NewQBrush10(void* gradient)
{
	return new QBrush(*static_cast<QGradient*>(gradient));
}

void* QBrush_Color(void* ptr)
{
	return const_cast<QColor*>(&static_cast<QBrush*>(ptr)->color());
}

void* QBrush_Texture(void* ptr)
{
	return new QPixmap(static_cast<QBrush*>(ptr)->texture());
}

void* QBrush_Transform(void* ptr)
{
	return new QTransform(static_cast<QBrush*>(ptr)->transform());
}

void QBrush_DestroyQBrush(void* ptr)
{
	static_cast<QBrush*>(ptr)->~QBrush();
}

void* QBrush_ToVariant(void* ptr)
{
	return new QVariant(*static_cast<QBrush*>(ptr));
}

void* QColor_NewQColor()
{
	return new QColor();
}

void* QColor_NewQColor2(long long color)
{
	return new QColor(static_cast<Qt::GlobalColor>(color));
}

void* QColor_NewQColor3(int r, int g, int b, int a)
{
	return new QColor(r, g, b, a);
}

void* QColor_NewQColor4(unsigned int color)
{
	return new QColor(color);
}

void* QColor_NewQColor5(void* rgba64)
{
	return new QColor(*static_cast<QRgba64*>(rgba64));
}

void* QColor_NewQColor6(struct QtGui_PackedString name)
{
	return new QColor(QString::fromUtf8(name.data, name.len));
}

void* QColor_NewQColor8(char* name)
{
	return new QColor(const_cast<const char*>(name));
}

void* QColor_NewQColor9(void* name)
{
	return new QColor(*static_cast<QLatin1String*>(name));
}

struct QtGui_PackedString QColor_Name(void* ptr)
{
	return ({ QByteArray t9b3be4 = static_cast<QColor*>(ptr)->name().toUtf8(); QtGui_PackedString { const_cast<char*>(t9b3be4.prepend("WHITESPACE").constData()+10), t9b3be4.size()-10 }; });
}

struct QtGui_PackedString QColor_Name2(void* ptr, long long format)
{
	return ({ QByteArray t4331f3 = static_cast<QColor*>(ptr)->name(static_cast<QColor::NameFormat>(format)).toUtf8(); QtGui_PackedString { const_cast<char*>(t4331f3.prepend("WHITESPACE").constData()+10), t4331f3.size()-10 }; });
}

int QColor_Red(void* ptr)
{
	return static_cast<QColor*>(ptr)->red();
}

unsigned int QColor_Rgb(void* ptr)
{
	return static_cast<QColor*>(ptr)->rgb();
}

int QColor_Value(void* ptr)
{
	return static_cast<QColor*>(ptr)->value();
}

void* QColor_ToVariant(void* ptr)
{
	return new QVariant(*static_cast<QColor*>(ptr));
}

char QDesktopServices_QDesktopServices_OpenUrl(void* url)
{
	return QDesktopServices::openUrl(*static_cast<QUrl*>(url));
}

class MyQDrag: public QDrag
{
public:
	MyQDrag(QObject *dragSource) : QDrag(dragSource) {QDrag_QDrag_QRegisterMetaType();};
	 ~MyQDrag() { callbackQDrag_DestroyQDrag(this); };
	void childEvent(QChildEvent * event) { callbackQDrag_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQDrag_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQDrag_CustomEvent(this, event); };
	void deleteLater() { callbackQDrag_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQDrag_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQDrag_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQDrag_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQDrag_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtGui_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQDrag_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQDrag_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQDrag*)

int QDrag_QDrag_QRegisterMetaType(){qRegisterMetaType<QDrag*>(); return qRegisterMetaType<MyQDrag*>();}

void* QDrag_NewQDrag(void* dragSource)
{
	if (dynamic_cast<QGraphicsObject*>(static_cast<QObject*>(dragSource))) {
		return new MyQDrag(static_cast<QGraphicsObject*>(dragSource));
	} else if (dynamic_cast<QGraphicsWidget*>(static_cast<QObject*>(dragSource))) {
		return new MyQDrag(static_cast<QGraphicsWidget*>(dragSource));
	} else if (dynamic_cast<QLayout*>(static_cast<QObject*>(dragSource))) {
		return new MyQDrag(static_cast<QLayout*>(dragSource));
	} else if (dynamic_cast<QQuickItem*>(static_cast<QObject*>(dragSource))) {
		return new MyQDrag(static_cast<QQuickItem*>(dragSource));
	} else if (dynamic_cast<QWidget*>(static_cast<QObject*>(dragSource))) {
		return new MyQDrag(static_cast<QWidget*>(dragSource));
	} else if (dynamic_cast<QWindow*>(static_cast<QObject*>(dragSource))) {
		return new MyQDrag(static_cast<QWindow*>(dragSource));
	} else {
		return new MyQDrag(static_cast<QObject*>(dragSource));
	}
}

long long QDrag_Exec(void* ptr, long long supportedActions)
{
	return static_cast<QDrag*>(ptr)->exec(static_cast<Qt::DropAction>(supportedActions));
}

long long QDrag_Exec2(void* ptr, long long supportedActions, long long defaultDropAction)
{
	return static_cast<QDrag*>(ptr)->exec(static_cast<Qt::DropAction>(supportedActions), static_cast<Qt::DropAction>(defaultDropAction));
}

void* QDrag_MimeData(void* ptr)
{
	return static_cast<QDrag*>(ptr)->mimeData();
}

void* QDrag_Pixmap(void* ptr)
{
	return new QPixmap(static_cast<QDrag*>(ptr)->pixmap());
}

void* QDrag_Source(void* ptr)
{
	return static_cast<QDrag*>(ptr)->source();
}

void QDrag_DestroyQDrag(void* ptr)
{
	static_cast<QDrag*>(ptr)->~QDrag();
}

void QDrag_DestroyQDragDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QDrag___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QDrag___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDrag___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QDrag___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QDrag___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QDrag___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QDrag___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QDrag___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDrag___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QDrag___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QDrag___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDrag___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QDrag___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QDrag___qFindChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QDrag___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QDrag_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QDrag*>(ptr)->QDrag::childEvent(static_cast<QChildEvent*>(event));
}

void QDrag_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDrag*>(ptr)->QDrag::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QDrag_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QDrag*>(ptr)->QDrag::customEvent(static_cast<QEvent*>(event));
}

void QDrag_DeleteLaterDefault(void* ptr)
{
		static_cast<QDrag*>(ptr)->QDrag::deleteLater();
}

void QDrag_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QDrag*>(ptr)->QDrag::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QDrag_EventDefault(void* ptr, void* e)
{
		return static_cast<QDrag*>(ptr)->QDrag::event(static_cast<QEvent*>(e));
}

char QDrag_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(watched))) {
		return static_cast<QDrag*>(ptr)->QDrag::eventFilter(static_cast<QWindow*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QDrag*>(ptr)->QDrag::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QDrag_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QDrag*>(ptr)->QDrag::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQDragEnterEvent: public QDragEnterEvent
{
public:
	MyQDragEnterEvent(const QPoint &point, Qt::DropActions actions, const QMimeData *data, Qt::MouseButtons buttons, Qt::KeyboardModifiers modifiers) : QDragEnterEvent(point, actions, data, buttons, modifiers) {};
};

void* QDragEnterEvent_NewQDragEnterEvent(void* point, long long actions, void* data, long long buttons, long long modifiers)
{
	return new MyQDragEnterEvent(*static_cast<QPoint*>(point), static_cast<Qt::DropAction>(actions), static_cast<QMimeData*>(data), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

class MyQDragLeaveEvent: public QDragLeaveEvent
{
public:
	MyQDragLeaveEvent() : QDragLeaveEvent() {};
};

void* QDragLeaveEvent_NewQDragLeaveEvent()
{
	return new MyQDragLeaveEvent();
}

class MyQDragMoveEvent: public QDragMoveEvent
{
public:
	MyQDragMoveEvent(const QPoint &pos, Qt::DropActions actions, const QMimeData *data, Qt::MouseButtons buttons, Qt::KeyboardModifiers modifiers, QEvent::Type ty = DragMove) : QDragMoveEvent(pos, actions, data, buttons, modifiers, ty) {};
	 ~MyQDragMoveEvent() { callbackQDragMoveEvent_DestroyQDragMoveEvent(this); };
};

void* QDragMoveEvent_NewQDragMoveEvent(void* pos, long long actions, void* data, long long buttons, long long modifiers, long long ty)
{
	return new MyQDragMoveEvent(*static_cast<QPoint*>(pos), static_cast<Qt::DropAction>(actions), static_cast<QMimeData*>(data), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers), static_cast<QEvent::Type>(ty));
}

void QDragMoveEvent_DestroyQDragMoveEvent(void* ptr)
{
	static_cast<QDragMoveEvent*>(ptr)->~QDragMoveEvent();
}

void QDragMoveEvent_DestroyQDragMoveEventDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQDropEvent: public QDropEvent
{
public:
	MyQDropEvent(const QPointF &pos, Qt::DropActions actions, const QMimeData *data, Qt::MouseButtons buttons, Qt::KeyboardModifiers modifiers, QEvent::Type ty = Drop) : QDropEvent(pos, actions, data, buttons, modifiers, ty) {};
};

void* QDropEvent_NewQDropEvent(void* pos, long long actions, void* data, long long buttons, long long modifiers, long long ty)
{
	return new MyQDropEvent(*static_cast<QPointF*>(pos), static_cast<Qt::DropAction>(actions), static_cast<QMimeData*>(data), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers), static_cast<QEvent::Type>(ty));
}

long long QDropEvent_DropAction(void* ptr)
{
	return static_cast<QDropEvent*>(ptr)->dropAction();
}

void* QDropEvent_MimeData(void* ptr)
{
	return const_cast<QMimeData*>(static_cast<QDropEvent*>(ptr)->mimeData());
}

void* QDropEvent_Source(void* ptr)
{
	return static_cast<QDropEvent*>(ptr)->source();
}

class MyQFocusEvent: public QFocusEvent
{
public:
	MyQFocusEvent(QEvent::Type ty, Qt::FocusReason reason = Qt::OtherFocusReason) : QFocusEvent(ty, reason) {};
};

void* QFocusEvent_NewQFocusEvent(long long ty, long long reason)
{
	return new MyQFocusEvent(static_cast<QEvent::Type>(ty), static_cast<Qt::FocusReason>(reason));
}

void* QFont_NewQFont()
{
	return new QFont();
}

void* QFont_NewQFont2(struct QtGui_PackedString family, int pointSize, int weight, char italic)
{
	return new QFont(QString::fromUtf8(family.data, family.len), pointSize, weight, italic != 0);
}

void* QFont_NewQFont4(void* font, void* pd)
{
	if (dynamic_cast<QWidget*>(static_cast<QObject*>(pd))) {
		return new QFont(*static_cast<QFont*>(font), static_cast<QWidget*>(pd));
	} else {
		return new QFont(*static_cast<QFont*>(font), static_cast<QPaintDevice*>(pd));
	}
}

void* QFont_NewQFont5(void* font)
{
	return new QFont(*static_cast<QFont*>(font));
}

struct QtGui_PackedString QFont_Key(void* ptr)
{
	return ({ QByteArray t9324a4 = static_cast<QFont*>(ptr)->key().toUtf8(); QtGui_PackedString { const_cast<char*>(t9324a4.prepend("WHITESPACE").constData()+10), t9324a4.size()-10 }; });
}

void* QFont_Resolve(void* ptr, void* other)
{
	return new QFont(static_cast<QFont*>(ptr)->resolve(*static_cast<QFont*>(other)));
}

void QFont_DestroyQFont(void* ptr)
{
	static_cast<QFont*>(ptr)->~QFont();
}

void* QFont_ToVariant(void* ptr)
{
	return new QVariant(*static_cast<QFont*>(ptr));
}

void* QGradient_NewQGradient2(long long preset)
{
	return new QGradient(static_cast<QGradient::Preset>(preset));
}

long long QGradient_Type(void* ptr)
{
	return static_cast<QGradient*>(ptr)->type();
}

class MyQGuiApplication: public QGuiApplication
{
public:
	MyQGuiApplication(int &argc, char **argv) : QGuiApplication(argc, argv) {QGuiApplication_QGuiApplication_QRegisterMetaType();};
	bool event(QEvent * e) { return callbackQGuiApplication_Event(this, e) != 0; };
	void Signal_FontChanged(const QFont & font) { callbackQGuiApplication_FontChanged(this, const_cast<QFont*>(&font)); };
	 ~MyQGuiApplication() { callbackQGuiApplication_DestroyQGuiApplication(this); };
	void childEvent(QChildEvent * event) { callbackQGuiApplication_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQGuiApplication_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQGuiApplication_CustomEvent(this, event); };
	void deleteLater() { callbackQGuiApplication_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQGuiApplication_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQGuiApplication_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQGuiApplication_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtGui_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQGuiApplication_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQGuiApplication_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQGuiApplication*)

int QGuiApplication_QGuiApplication_QRegisterMetaType(){qRegisterMetaType<QGuiApplication*>(); return qRegisterMetaType<MyQGuiApplication*>();}

void* QGuiApplication_NewQGuiApplication(int argc, char* argv)
{
	static int argcs = argc;
	static char** argvs = static_cast<char**>(malloc(argcs * sizeof(char*)));

	QList<QByteArray> aList = QByteArray(argv).split('|');
	for (int i = 0; i < argcs; i++)
		argvs[i] = (new QByteArray(aList.at(i)))->data();

	return new MyQGuiApplication(argcs, argvs);
}

char QGuiApplication_EventDefault(void* ptr, void* e)
{
		return static_cast<QGuiApplication*>(ptr)->QGuiApplication::event(static_cast<QEvent*>(e));
}

int QGuiApplication_QGuiApplication_Exec()
{
	return QGuiApplication::exec();
}

void* QGuiApplication_QGuiApplication_Font()
{
	return new QFont(QGuiApplication::font());
}

void QGuiApplication_ConnectFontChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(const QFont &)>(&QGuiApplication::fontChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(const QFont &)>(&MyQGuiApplication::Signal_FontChanged), static_cast<Qt::ConnectionType>(t));
}

void QGuiApplication_DisconnectFontChanged(void* ptr)
{
	QObject::disconnect(static_cast<QGuiApplication*>(ptr), static_cast<void (QGuiApplication::*)(const QFont &)>(&QGuiApplication::fontChanged), static_cast<MyQGuiApplication*>(ptr), static_cast<void (MyQGuiApplication::*)(const QFont &)>(&MyQGuiApplication::Signal_FontChanged));
}

void QGuiApplication_FontChanged(void* ptr, void* font)
{
	static_cast<QGuiApplication*>(ptr)->fontChanged(*static_cast<QFont*>(font));
}

void* QGuiApplication_QGuiApplication_InputMethod()
{
	return QGuiApplication::inputMethod();
}

void QGuiApplication_QGuiApplication_SetFont(void* font)
{
	QGuiApplication::setFont(*static_cast<QFont*>(font));
}

void QGuiApplication_QGuiApplication_Sync()
{
	QGuiApplication::sync();
}

void QGuiApplication_DestroyQGuiApplication(void* ptr)
{
	static_cast<QGuiApplication*>(ptr)->~QGuiApplication();
}

void QGuiApplication_DestroyQGuiApplicationDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QGuiApplication___screens_atList(void* ptr, int i)
{
	return ({QScreen * tmp = static_cast<QList<QScreen *>*>(ptr)->at(i); if (i == static_cast<QList<QScreen *>*>(ptr)->size()-1) { static_cast<QList<QScreen *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGuiApplication___screens_setList(void* ptr, void* i)
{
	static_cast<QList<QScreen *>*>(ptr)->append(static_cast<QScreen*>(i));
}

void* QGuiApplication___screens_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QScreen *>();
}

void* QGuiApplication___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGuiApplication___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QGuiApplication___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QGuiApplication___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QGuiApplication___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QGuiApplication___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QGuiApplication___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGuiApplication___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QGuiApplication___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGuiApplication___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGuiApplication___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QGuiApplication___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QGuiApplication___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QGuiApplication___qFindChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QGuiApplication___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QGuiApplication_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QGuiApplication*>(ptr)->QGuiApplication::childEvent(static_cast<QChildEvent*>(event));
}

void QGuiApplication_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGuiApplication*>(ptr)->QGuiApplication::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QGuiApplication_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QGuiApplication*>(ptr)->QGuiApplication::customEvent(static_cast<QEvent*>(event));
}

void QGuiApplication_DeleteLaterDefault(void* ptr)
{
		static_cast<QGuiApplication*>(ptr)->QGuiApplication::deleteLater();
}

void QGuiApplication_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QGuiApplication*>(ptr)->QGuiApplication::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QGuiApplication_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(watched))) {
		return static_cast<QGuiApplication*>(ptr)->QGuiApplication::eventFilter(static_cast<QWindow*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QGuiApplication*>(ptr)->QGuiApplication::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QGuiApplication_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QGuiApplication*>(ptr)->QGuiApplication::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQHoverEvent: public QHoverEvent
{
public:
	MyQHoverEvent(QEvent::Type ty, const QPointF &pos, const QPointF &oldPos, Qt::KeyboardModifiers modifiers = Qt::NoModifier) : QHoverEvent(ty, pos, oldPos, modifiers) {};
};

void* QHoverEvent_NewQHoverEvent(long long ty, void* pos, void* oldPos, long long modifiers)
{
	return new MyQHoverEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(pos), *static_cast<QPointF*>(oldPos), static_cast<Qt::KeyboardModifier>(modifiers));
}

void* QIcon_NewQIcon()
{
	return new QIcon();
}

void* QIcon_NewQIcon2(void* pixmap)
{
	return new QIcon(*static_cast<QPixmap*>(pixmap));
}

void* QIcon_NewQIcon3(void* other)
{
	return new QIcon(*static_cast<QIcon*>(other));
}

void* QIcon_NewQIcon4(void* other)
{
	return new QIcon(*static_cast<QIcon*>(other));
}

void* QIcon_NewQIcon5(struct QtGui_PackedString fileName)
{
	return new QIcon(QString::fromUtf8(fileName.data, fileName.len));
}

void* QIcon_NewQIcon6(void* engine)
{
	return new QIcon(static_cast<QIconEngine*>(engine));
}

void* QIcon_QIcon_FromTheme(struct QtGui_PackedString name)
{
	return new QIcon(QIcon::fromTheme(QString::fromUtf8(name.data, name.len)));
}

void* QIcon_QIcon_FromTheme2(struct QtGui_PackedString name, void* fallback)
{
	return new QIcon(QIcon::fromTheme(QString::fromUtf8(name.data, name.len), *static_cast<QIcon*>(fallback)));
}

struct QtGui_PackedString QIcon_Name(void* ptr)
{
	return ({ QByteArray t03700a = static_cast<QIcon*>(ptr)->name().toUtf8(); QtGui_PackedString { const_cast<char*>(t03700a.prepend("WHITESPACE").constData()+10), t03700a.size()-10 }; });
}

void* QIcon_Pixmap(void* ptr, void* size, long long mode, long long state)
{
	return new QPixmap(static_cast<QIcon*>(ptr)->pixmap(*static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state)));
}

void* QIcon_Pixmap2(void* ptr, int w, int h, long long mode, long long state)
{
	return new QPixmap(static_cast<QIcon*>(ptr)->pixmap(w, h, static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state)));
}

void* QIcon_Pixmap3(void* ptr, int extent, long long mode, long long state)
{
	return new QPixmap(static_cast<QIcon*>(ptr)->pixmap(extent, static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state)));
}

void* QIcon_Pixmap4(void* ptr, void* window, void* size, long long mode, long long state)
{
		return new QPixmap(static_cast<QIcon*>(ptr)->pixmap(static_cast<QWindow*>(window), *static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state)));
}

void QIcon_DestroyQIcon(void* ptr)
{
	static_cast<QIcon*>(ptr)->~QIcon();
}

void* QIcon_ToVariant(void* ptr)
{
	return new QVariant(*static_cast<QIcon*>(ptr));
}

void* QIcon___availableSizes_atList(void* ptr, int i)
{
	return ({ QSize tmpValue = ({QSize tmp = static_cast<QList<QSize>*>(ptr)->at(i); if (i == static_cast<QList<QSize>*>(ptr)->size()-1) { static_cast<QList<QSize>*>(ptr)->~QList(); free(ptr); }; tmp; }); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QIcon___availableSizes_setList(void* ptr, void* i)
{
	static_cast<QList<QSize>*>(ptr)->append(*static_cast<QSize*>(i));
}

void* QIcon___availableSizes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSize>();
}

class MyQIconEngine: public QIconEngine
{
public:
	MyQIconEngine() : QIconEngine() {};
	QIconEngine * clone() const { return static_cast<QIconEngine*>(callbackQIconEngine_Clone(const_cast<void*>(static_cast<const void*>(this)))); };
	QString key() const { return ({ QtGui_PackedString tempVal = callbackQIconEngine_Key(const_cast<void*>(static_cast<const void*>(this))); QString ret = QString::fromUtf8(tempVal.data, tempVal.len); free(tempVal.data); ret; }); };
	void paint(QPainter * painter, const QRect & rect, QIcon::Mode mode, QIcon::State state) { callbackQIconEngine_Paint(this, painter, const_cast<QRect*>(&rect), mode, state); };
	QPixmap pixmap(const QSize & size, QIcon::Mode mode, QIcon::State state) { return *static_cast<QPixmap*>(callbackQIconEngine_Pixmap(this, const_cast<QSize*>(&size), mode, state)); };
	 ~MyQIconEngine() { callbackQIconEngine_DestroyQIconEngine(this); };
};

void* QIconEngine_NewQIconEngine()
{
	return new MyQIconEngine();
}

void* QIconEngine_Clone(void* ptr)
{
	return static_cast<QIconEngine*>(ptr)->clone();
}

struct QtGui_PackedString QIconEngine_Key(void* ptr)
{
	return ({ QByteArray tfa2543 = static_cast<QIconEngine*>(ptr)->key().toUtf8(); QtGui_PackedString { const_cast<char*>(tfa2543.prepend("WHITESPACE").constData()+10), tfa2543.size()-10 }; });
}

struct QtGui_PackedString QIconEngine_KeyDefault(void* ptr)
{
		return ({ QByteArray t9979b6 = static_cast<QIconEngine*>(ptr)->QIconEngine::key().toUtf8(); QtGui_PackedString { const_cast<char*>(t9979b6.prepend("WHITESPACE").constData()+10), t9979b6.size()-10 }; });
}

void QIconEngine_Paint(void* ptr, void* painter, void* rect, long long mode, long long state)
{
	static_cast<QIconEngine*>(ptr)->paint(static_cast<QPainter*>(painter), *static_cast<QRect*>(rect), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state));
}

void* QIconEngine_Pixmap(void* ptr, void* size, long long mode, long long state)
{
	return new QPixmap(static_cast<QIconEngine*>(ptr)->pixmap(*static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state)));
}

void* QIconEngine_PixmapDefault(void* ptr, void* size, long long mode, long long state)
{
		return new QPixmap(static_cast<QIconEngine*>(ptr)->QIconEngine::pixmap(*static_cast<QSize*>(size), static_cast<QIcon::Mode>(mode), static_cast<QIcon::State>(state)));
}

void QIconEngine_DestroyQIconEngine(void* ptr)
{
	static_cast<QIconEngine*>(ptr)->~QIconEngine();
}

void QIconEngine_DestroyQIconEngineDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QIconEngine___availableSizes_atList(void* ptr, int i)
{
	return ({ QSize tmpValue = ({QSize tmp = static_cast<QList<QSize>*>(ptr)->at(i); if (i == static_cast<QList<QSize>*>(ptr)->size()-1) { static_cast<QList<QSize>*>(ptr)->~QList(); free(ptr); }; tmp; }); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QIconEngine___availableSizes_setList(void* ptr, void* i)
{
	static_cast<QList<QSize>*>(ptr)->append(*static_cast<QSize*>(i));
}

void* QIconEngine___availableSizes_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QSize>();
}

class MyQImage: public QImage
{
public:
	MyQImage() : QImage() {};
	MyQImage(const QSize &size, QImage::Format format) : QImage(size, format) {};
	MyQImage(int width, int height, QImage::Format format) : QImage(width, height, format) {};
	MyQImage(uchar *data, int width, int height, QImage::Format format) : QImage(data, width, height, format) {};
	MyQImage(const uchar *data, int width, int height, QImage::Format format) : QImage(data, width, height, format) {};
	MyQImage(uchar *data, int width, int height, int bytesPerLine, QImage::Format format) : QImage(data, width, height, bytesPerLine, format) {};
	MyQImage(const uchar *data, int width, int height, int bytesPerLine, QImage::Format format) : QImage(data, width, height, bytesPerLine, format) {};
	MyQImage(const QString &fileName, const char *format = Q_NULLPTR) : QImage(fileName, format) {};
	MyQImage(const QImage &image) : QImage(image) {};
	MyQImage(QImage &&other) : QImage(other) {};
	 ~MyQImage() { callbackQImage_DestroyQImage(this); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQImage_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

void* QImage_NewQImage()
{
	return new MyQImage();
}

void* QImage_NewQImage2(void* size, long long format)
{
	return new MyQImage(*static_cast<QSize*>(size), static_cast<QImage::Format>(format));
}

void* QImage_NewQImage3(int width, int height, long long format)
{
	return new MyQImage(width, height, static_cast<QImage::Format>(format));
}

void* QImage_NewQImage4(char* data, int width, int height, long long format)
{
	return new MyQImage(static_cast<uchar*>(static_cast<void*>(data)), width, height, static_cast<QImage::Format>(format));
}

void* QImage_NewQImage5(char* data, int width, int height, long long format)
{
	return new MyQImage(const_cast<const uchar*>(static_cast<uchar*>(static_cast<void*>(data))), width, height, static_cast<QImage::Format>(format));
}

void* QImage_NewQImage6(char* data, int width, int height, int bytesPerLine, long long format)
{
	return new MyQImage(static_cast<uchar*>(static_cast<void*>(data)), width, height, bytesPerLine, static_cast<QImage::Format>(format));
}

void* QImage_NewQImage7(char* data, int width, int height, int bytesPerLine, long long format)
{
	return new MyQImage(const_cast<const uchar*>(static_cast<uchar*>(static_cast<void*>(data))), width, height, bytesPerLine, static_cast<QImage::Format>(format));
}

void* QImage_NewQImage9(struct QtGui_PackedString fileName, char* format)
{
	return new MyQImage(QString::fromUtf8(fileName.data, fileName.len), const_cast<const char*>(format));
}

void* QImage_NewQImage10(void* image)
{
	return new MyQImage(*static_cast<QImage*>(image));
}

void* QImage_NewQImage11(void* other)
{
	return new MyQImage(*static_cast<QImage*>(other));
}

unsigned int QImage_Color(void* ptr, int i)
{
	return static_cast<QImage*>(ptr)->color(i);
}

void QImage_Fill(void* ptr, unsigned int pixelValue)
{
	static_cast<QImage*>(ptr)->fill(pixelValue);
}

void QImage_Fill2(void* ptr, void* color)
{
	static_cast<QImage*>(ptr)->fill(*static_cast<QColor*>(color));
}

void QImage_Fill3(void* ptr, long long color)
{
	static_cast<QImage*>(ptr)->fill(static_cast<Qt::GlobalColor>(color));
}

void* QImage_Offset(void* ptr)
{
	return ({ QPoint tmpValue = static_cast<QImage*>(ptr)->offset(); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void* QImage_Rect(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QImage*>(ptr)->rect(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QImage_SetText(void* ptr, struct QtGui_PackedString key, struct QtGui_PackedString text)
{
	static_cast<QImage*>(ptr)->setText(QString::fromUtf8(key.data, key.len), QString::fromUtf8(text.data, text.len));
}

void* QImage_Size(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QImage*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct QtGui_PackedString QImage_Text(void* ptr, struct QtGui_PackedString key)
{
	return ({ QByteArray t3cc0e5 = static_cast<QImage*>(ptr)->text(QString::fromUtf8(key.data, key.len)).toUtf8(); QtGui_PackedString { const_cast<char*>(t3cc0e5.prepend("WHITESPACE").constData()+10), t3cc0e5.size()-10 }; });
}

void QImage_DestroyQImage(void* ptr)
{
	static_cast<QImage*>(ptr)->~QImage();
}

void QImage_DestroyQImageDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QImage_ToVariant(void* ptr)
{
	return new QVariant(*static_cast<QImage*>(ptr));
}

unsigned int QImage___colorTable_atList(void* ptr, int i)
{
	return ({QRgb tmp = static_cast<QVector<QRgb>*>(ptr)->at(i); if (i == static_cast<QVector<QRgb>*>(ptr)->size()-1) { static_cast<QVector<QRgb>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QImage___colorTable_setList(void* ptr, unsigned int i)
{
	static_cast<QVector<QRgb>*>(ptr)->append(i);
}

void* QImage___colorTable_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QRgb>();
}

unsigned int QImage___convertToFormat_colorTable_atList3(void* ptr, int i)
{
	return ({QRgb tmp = static_cast<QVector<QRgb>*>(ptr)->at(i); if (i == static_cast<QVector<QRgb>*>(ptr)->size()-1) { static_cast<QVector<QRgb>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QImage___convertToFormat_colorTable_setList3(void* ptr, unsigned int i)
{
	static_cast<QVector<QRgb>*>(ptr)->append(i);
}

void* QImage___convertToFormat_colorTable_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QRgb>();
}

unsigned int QImage___setColorTable_colors_atList(void* ptr, int i)
{
	return ({QRgb tmp = static_cast<QVector<QRgb>*>(ptr)->at(i); if (i == static_cast<QVector<QRgb>*>(ptr)->size()-1) { static_cast<QVector<QRgb>*>(ptr)->~QVector(); free(ptr); }; tmp; });
}

void QImage___setColorTable_colors_setList(void* ptr, unsigned int i)
{
	static_cast<QVector<QRgb>*>(ptr)->append(i);
}

void* QImage___setColorTable_colors_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QRgb>();
}

void* QImage_PaintEngine(void* ptr)
{
	return static_cast<QImage*>(ptr)->paintEngine();
}

void* QImage_PaintEngineDefault(void* ptr)
{
		return static_cast<QImage*>(ptr)->QImage::paintEngine();
}

class MyQInputEvent: public QInputEvent
{
public:
};

unsigned long QInputEvent_Timestamp(void* ptr)
{
	return static_cast<QInputEvent*>(ptr)->timestamp();
}

class MyQInputMethod: public QInputMethod
{
public:
	void hide() { callbackQInputMethod_Hide(this); };
	void reset() { callbackQInputMethod_Reset(this); };
	void show() { callbackQInputMethod_Show(this); };
	void update(Qt::InputMethodQueries queries) { callbackQInputMethod_Update(this, queries); };
	void Signal_VisibleChanged() { callbackQInputMethod_VisibleChanged(this); };
	void childEvent(QChildEvent * event) { callbackQInputMethod_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQInputMethod_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQInputMethod_CustomEvent(this, event); };
	void deleteLater() { callbackQInputMethod_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQInputMethod_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQInputMethod_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQInputMethod_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQInputMethod_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtGui_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQInputMethod_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQInputMethod_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQInputMethod*)

int QInputMethod_QInputMethod_QRegisterMetaType(){qRegisterMetaType<QInputMethod*>(); return qRegisterMetaType<MyQInputMethod*>();}

void QInputMethod_Hide(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "hide");
}

void QInputMethod_HideDefault(void* ptr)
{
		static_cast<QInputMethod*>(ptr)->QInputMethod::hide();
}

void QInputMethod_Reset(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "reset");
}

void QInputMethod_ResetDefault(void* ptr)
{
		static_cast<QInputMethod*>(ptr)->QInputMethod::reset();
}

void QInputMethod_Show(void* ptr)
{
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "show");
}

void QInputMethod_ShowDefault(void* ptr)
{
		static_cast<QInputMethod*>(ptr)->QInputMethod::show();
}

void QInputMethod_Update(void* ptr, long long queries)
{
	qRegisterMetaType<Qt::InputMethodQueries>();
	QMetaObject::invokeMethod(static_cast<QInputMethod*>(ptr), "update", Q_ARG(Qt::InputMethodQueries, static_cast<Qt::InputMethodQuery>(queries)));
}

void QInputMethod_UpdateDefault(void* ptr, long long queries)
{
		static_cast<QInputMethod*>(ptr)->QInputMethod::update(static_cast<Qt::InputMethodQuery>(queries));
}

void QInputMethod_ConnectVisibleChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::visibleChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_VisibleChanged), static_cast<Qt::ConnectionType>(t));
}

void QInputMethod_DisconnectVisibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QInputMethod*>(ptr), static_cast<void (QInputMethod::*)()>(&QInputMethod::visibleChanged), static_cast<MyQInputMethod*>(ptr), static_cast<void (MyQInputMethod::*)()>(&MyQInputMethod::Signal_VisibleChanged));
}

void QInputMethod_VisibleChanged(void* ptr)
{
	static_cast<QInputMethod*>(ptr)->visibleChanged();
}

void* QInputMethod___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QInputMethod___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QInputMethod___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QInputMethod___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QInputMethod___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QInputMethod___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QInputMethod___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QInputMethod___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QInputMethod___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QInputMethod___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QInputMethod___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QInputMethod___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QInputMethod___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QInputMethod___qFindChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QInputMethod___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QInputMethod_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QInputMethod*>(ptr)->QInputMethod::childEvent(static_cast<QChildEvent*>(event));
}

void QInputMethod_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QInputMethod*>(ptr)->QInputMethod::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QInputMethod_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QInputMethod*>(ptr)->QInputMethod::customEvent(static_cast<QEvent*>(event));
}

void QInputMethod_DeleteLaterDefault(void* ptr)
{
		static_cast<QInputMethod*>(ptr)->QInputMethod::deleteLater();
}

void QInputMethod_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QInputMethod*>(ptr)->QInputMethod::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QInputMethod_EventDefault(void* ptr, void* e)
{
		return static_cast<QInputMethod*>(ptr)->QInputMethod::event(static_cast<QEvent*>(e));
}

char QInputMethod_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(watched))) {
		return static_cast<QInputMethod*>(ptr)->QInputMethod::eventFilter(static_cast<QWindow*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QInputMethod*>(ptr)->QInputMethod::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QInputMethod_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QInputMethod*>(ptr)->QInputMethod::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQInputMethodEvent: public QInputMethodEvent
{
public:
	MyQInputMethodEvent() : QInputMethodEvent() {};
	MyQInputMethodEvent(const QInputMethodEvent &other) : QInputMethodEvent(other) {};
};

void* QInputMethodEvent_NewQInputMethodEvent()
{
	return new MyQInputMethodEvent();
}

void* QInputMethodEvent_NewQInputMethodEvent3(void* other)
{
	return new MyQInputMethodEvent(*static_cast<QInputMethodEvent*>(other));
}

void* QInputMethodEvent___attrs_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QInputMethodEvent::Attribute>();
}

void* QInputMethodEvent___setAttrs__newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QInputMethodEvent::Attribute>();
}

class MyQKeyEvent: public QKeyEvent
{
public:
	MyQKeyEvent(QEvent::Type ty, int key, Qt::KeyboardModifiers modifiers, const QString &text = QString(), bool autorep = false, ushort count = 1) : QKeyEvent(ty, key, modifiers, text, autorep, count) {};
	MyQKeyEvent(QEvent::Type ty, int key, Qt::KeyboardModifiers modifiers, quint32 nativeScanCode, quint32 nativeVirtualKey, quint32 nativeModifiers, const QString &text = QString(), bool autorep = false, ushort count = 1) : QKeyEvent(ty, key, modifiers, nativeScanCode, nativeVirtualKey, nativeModifiers, text, autorep, count) {};
};

void* QKeyEvent_NewQKeyEvent(long long ty, int key, long long modifiers, struct QtGui_PackedString text, char autorep, unsigned short count)
{
	return new MyQKeyEvent(static_cast<QEvent::Type>(ty), key, static_cast<Qt::KeyboardModifier>(modifiers), QString::fromUtf8(text.data, text.len), autorep != 0, count);
}

void* QKeyEvent_NewQKeyEvent2(long long ty, int key, long long modifiers, unsigned int nativeScanCode, unsigned int nativeVirtualKey, unsigned int nativeModifiers, struct QtGui_PackedString text, char autorep, unsigned short count)
{
	return new MyQKeyEvent(static_cast<QEvent::Type>(ty), key, static_cast<Qt::KeyboardModifier>(modifiers), nativeScanCode, nativeVirtualKey, nativeModifiers, QString::fromUtf8(text.data, text.len), autorep != 0, count);
}

int QKeyEvent_Count(void* ptr)
{
	return static_cast<QKeyEvent*>(ptr)->count();
}

int QKeyEvent_Key(void* ptr)
{
	return static_cast<QKeyEvent*>(ptr)->key();
}

struct QtGui_PackedString QKeyEvent_Text(void* ptr)
{
	return ({ QByteArray tac962f = static_cast<QKeyEvent*>(ptr)->text().toUtf8(); QtGui_PackedString { const_cast<char*>(tac962f.prepend("WHITESPACE").constData()+10), tac962f.size()-10 }; });
}

void* QKeySequence_NewQKeySequence()
{
	return new QKeySequence();
}

void* QKeySequence_NewQKeySequence2(struct QtGui_PackedString key, long long format)
{
	return new QKeySequence(QString::fromUtf8(key.data, key.len), static_cast<QKeySequence::SequenceFormat>(format));
}

void* QKeySequence_NewQKeySequence3(int k1, int k2, int k3, int k4)
{
	return new QKeySequence(k1, k2, k3, k4);
}

void* QKeySequence_NewQKeySequence4(void* keysequence)
{
	return new QKeySequence(*static_cast<QKeySequence*>(keysequence));
}

void* QKeySequence_NewQKeySequence5(long long key)
{
	return new QKeySequence(static_cast<QKeySequence::StandardKey>(key));
}

int QKeySequence_Count(void* ptr)
{
	return static_cast<QKeySequence*>(ptr)->count();
}

void QKeySequence_DestroyQKeySequence(void* ptr)
{
	static_cast<QKeySequence*>(ptr)->~QKeySequence();
}

void* QKeySequence___keyBindings_atList(void* ptr, int i)
{
	return new QKeySequence(({QKeySequence tmp = static_cast<QList<QKeySequence>*>(ptr)->at(i); if (i == static_cast<QList<QKeySequence>*>(ptr)->size()-1) { static_cast<QList<QKeySequence>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QKeySequence___keyBindings_setList(void* ptr, void* i)
{
	static_cast<QList<QKeySequence>*>(ptr)->append(*static_cast<QKeySequence*>(i));
}

void* QKeySequence___keyBindings_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QKeySequence>();
}

void* QKeySequence___listFromString_atList(void* ptr, int i)
{
	return new QKeySequence(({QKeySequence tmp = static_cast<QList<QKeySequence>*>(ptr)->at(i); if (i == static_cast<QList<QKeySequence>*>(ptr)->size()-1) { static_cast<QList<QKeySequence>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QKeySequence___listFromString_setList(void* ptr, void* i)
{
	static_cast<QList<QKeySequence>*>(ptr)->append(*static_cast<QKeySequence*>(i));
}

void* QKeySequence___listFromString_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QKeySequence>();
}

void* QKeySequence___listToString_list_atList(void* ptr, int i)
{
	return new QKeySequence(({QKeySequence tmp = static_cast<QList<QKeySequence>*>(ptr)->at(i); if (i == static_cast<QList<QKeySequence>*>(ptr)->size()-1) { static_cast<QList<QKeySequence>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QKeySequence___listToString_list_setList(void* ptr, void* i)
{
	static_cast<QList<QKeySequence>*>(ptr)->append(*static_cast<QKeySequence*>(i));
}

void* QKeySequence___listToString_list_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QKeySequence>();
}

void* QMatrix_NewQMatrix2()
{
	return new QMatrix();
}

void* QMatrix_NewQMatrix3(double m11, double m12, double m21, double m22, double dx, double dy)
{
	return new QMatrix(m11, m12, m21, m22, dx, dy);
}

void* QMatrix_NewQMatrix5(void* matrix)
{
	return new QMatrix(*static_cast<QMatrix*>(matrix));
}

void QMatrix_Reset(void* ptr)
{
	static_cast<QMatrix*>(ptr)->reset();
}

void* QMatrix_Scale(void* ptr, double sx, double sy)
{
	return new QMatrix(static_cast<QMatrix*>(ptr)->scale(sx, sy));
}

void* QMatrix4x4_NewQMatrix4x4()
{
	return new QMatrix4x4();
}

void* QMatrix4x4_NewQMatrix4x43(float values)
{
	return new QMatrix4x4(const_cast<const float*>(&values));
}

void* QMatrix4x4_NewQMatrix4x44(float m11, float m12, float m13, float m14, float m21, float m22, float m23, float m24, float m31, float m32, float m33, float m34, float m41, float m42, float m43, float m44)
{
	return new QMatrix4x4(m11, m12, m13, m14, m21, m22, m23, m24, m31, m32, m33, m34, m41, m42, m43, m44);
}

void* QMatrix4x4_NewQMatrix4x46(void* matrix)
{
	return new QMatrix4x4(*static_cast<QMatrix*>(matrix));
}

void* QMatrix4x4_Column(void* ptr, int index)
{
	return new QVector4D(static_cast<QMatrix4x4*>(ptr)->column(index));
}

float QMatrix4x4_Data(void* ptr)
{
	return *static_cast<QMatrix4x4*>(ptr)->data();
}

float QMatrix4x4_Data2(void* ptr)
{
	return *static_cast<QMatrix4x4*>(ptr)->data();
}

void QMatrix4x4_Fill(void* ptr, float value)
{
	static_cast<QMatrix4x4*>(ptr)->fill(value);
}

void* QMatrix4x4_Row(void* ptr, int index)
{
	return new QVector4D(static_cast<QMatrix4x4*>(ptr)->row(index));
}

void QMatrix4x4_Scale(void* ptr, void* vector)
{
	static_cast<QMatrix4x4*>(ptr)->scale(*static_cast<QVector3D*>(vector));
}

void QMatrix4x4_Scale2(void* ptr, float x, float y)
{
	static_cast<QMatrix4x4*>(ptr)->scale(x, y);
}

void QMatrix4x4_Scale3(void* ptr, float x, float y, float z)
{
	static_cast<QMatrix4x4*>(ptr)->scale(x, y, z);
}

void QMatrix4x4_Scale4(void* ptr, float factor)
{
	static_cast<QMatrix4x4*>(ptr)->scale(factor);
}

class MyQMouseEvent: public QMouseEvent
{
public:
	MyQMouseEvent(QEvent::Type ty, const QPointF &localPos, Qt::MouseButton button, Qt::MouseButtons buttons, Qt::KeyboardModifiers modifiers) : QMouseEvent(ty, localPos, button, buttons, modifiers) {};
	MyQMouseEvent(QEvent::Type ty, const QPointF &localPos, const QPointF &screenPos, Qt::MouseButton button, Qt::MouseButtons buttons, Qt::KeyboardModifiers modifiers) : QMouseEvent(ty, localPos, screenPos, button, buttons, modifiers) {};
	MyQMouseEvent(QEvent::Type ty, const QPointF &localPos, const QPointF &windowPos, const QPointF &screenPos, Qt::MouseButton button, Qt::MouseButtons buttons, Qt::KeyboardModifiers modifiers) : QMouseEvent(ty, localPos, windowPos, screenPos, button, buttons, modifiers) {};
	MyQMouseEvent(QEvent::Type ty, const QPointF &localPos, const QPointF &windowPos, const QPointF &screenPos, Qt::MouseButton button, Qt::MouseButtons buttons, Qt::KeyboardModifiers modifiers, Qt::MouseEventSource source) : QMouseEvent(ty, localPos, windowPos, screenPos, button, buttons, modifiers, source) {};
};

void* QMouseEvent_NewQMouseEvent(long long ty, void* localPos, long long button, long long buttons, long long modifiers)
{
	return new MyQMouseEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(localPos), static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

void* QMouseEvent_NewQMouseEvent2(long long ty, void* localPos, void* screenPos, long long button, long long buttons, long long modifiers)
{
	return new MyQMouseEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(localPos), *static_cast<QPointF*>(screenPos), static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

void* QMouseEvent_NewQMouseEvent3(long long ty, void* localPos, void* windowPos, void* screenPos, long long button, long long buttons, long long modifiers)
{
	return new MyQMouseEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(localPos), *static_cast<QPointF*>(windowPos), *static_cast<QPointF*>(screenPos), static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

void* QMouseEvent_NewQMouseEvent4(long long ty, void* localPos, void* windowPos, void* screenPos, long long button, long long buttons, long long modifiers, long long source)
{
	return new MyQMouseEvent(static_cast<QEvent::Type>(ty), *static_cast<QPointF*>(localPos), *static_cast<QPointF*>(windowPos), *static_cast<QPointF*>(screenPos), static_cast<Qt::MouseButton>(button), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers), static_cast<Qt::MouseEventSource>(source));
}

long long QMouseEvent_Button(void* ptr)
{
	return static_cast<QMouseEvent*>(ptr)->button();
}

long long QMouseEvent_Buttons(void* ptr)
{
	return static_cast<QMouseEvent*>(ptr)->buttons();
}

long long QMouseEvent_Flags(void* ptr)
{
	return static_cast<QMouseEvent*>(ptr)->flags();
}

long long QMouseEvent_Source(void* ptr)
{
	return static_cast<QMouseEvent*>(ptr)->source();
}

int QMouseEvent_X(void* ptr)
{
	return static_cast<QMouseEvent*>(ptr)->x();
}

int QMouseEvent_Y(void* ptr)
{
	return static_cast<QMouseEvent*>(ptr)->y();
}

class MyQMoveEvent: public QMoveEvent
{
public:
	MyQMoveEvent(const QPoint &pos, const QPoint &oldPos) : QMoveEvent(pos, oldPos) {};
};

void* QMoveEvent_NewQMoveEvent(void* pos, void* oldPos)
{
	return new MyQMoveEvent(*static_cast<QPoint*>(pos), *static_cast<QPoint*>(oldPos));
}

class MyQPaintDevice: public QPaintDevice
{
public:
	MyQPaintDevice() : QPaintDevice() {};
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPaintDevice_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
	 ~MyQPaintDevice() { callbackQPaintDevice_DestroyQPaintDevice(this); };
};

void* QPaintDevice_NewQPaintDevice()
{
	return new MyQPaintDevice();
}

int QPaintDevice_Height(void* ptr)
{
	return static_cast<QPaintDevice*>(ptr)->height();
}

void* QPaintDevice_PaintEngine(void* ptr)
{
	return static_cast<QPaintDevice*>(ptr)->paintEngine();
}

int QPaintDevice_Width(void* ptr)
{
	return static_cast<QPaintDevice*>(ptr)->width();
}

void QPaintDevice_DestroyQPaintDevice(void* ptr)
{
	static_cast<QPaintDevice*>(ptr)->~QPaintDevice();
}

void QPaintDevice_DestroyQPaintDeviceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

class MyQPaintEngine: public QPaintEngine
{
public:
	MyQPaintEngine(QPaintEngine::PaintEngineFeatures caps = PaintEngineFeatures()) : QPaintEngine(caps) {};
	bool begin(QPaintDevice * pdev) { return callbackQPaintEngine_Begin(this, pdev) != 0; };
	void drawPixmap(const QRectF & r, const QPixmap & pm, const QRectF & sr) { callbackQPaintEngine_DrawPixmap(this, const_cast<QRectF*>(&r), const_cast<QPixmap*>(&pm), const_cast<QRectF*>(&sr)); };
	bool end() { return callbackQPaintEngine_End(this) != 0; };
	QPaintEngine::Type type() const { return static_cast<QPaintEngine::Type>(callbackQPaintEngine_Type(const_cast<void*>(static_cast<const void*>(this)))); };
	void updateState(const QPaintEngineState & state) { callbackQPaintEngine_UpdateState(this, const_cast<QPaintEngineState*>(&state)); };
	 ~MyQPaintEngine() { callbackQPaintEngine_DestroyQPaintEngine(this); };
};

void* QPaintEngine_NewQPaintEngine(long long caps)
{
	return new MyQPaintEngine(static_cast<QPaintEngine::PaintEngineFeature>(caps));
}

char QPaintEngine_Begin(void* ptr, void* pdev)
{
	return static_cast<QPaintEngine*>(ptr)->begin(static_cast<QPaintDevice*>(pdev));
}

void QPaintEngine_DrawPixmap(void* ptr, void* r, void* pm, void* sr)
{
	static_cast<QPaintEngine*>(ptr)->drawPixmap(*static_cast<QRectF*>(r), *static_cast<QPixmap*>(pm), *static_cast<QRectF*>(sr));
}

char QPaintEngine_End(void* ptr)
{
	return static_cast<QPaintEngine*>(ptr)->end();
}

long long QPaintEngine_Type(void* ptr)
{
	return static_cast<QPaintEngine*>(ptr)->type();
}

void QPaintEngine_UpdateState(void* ptr, void* state)
{
	static_cast<QPaintEngine*>(ptr)->updateState(*static_cast<QPaintEngineState*>(state));
}

void QPaintEngine_DestroyQPaintEngine(void* ptr)
{
	static_cast<QPaintEngine*>(ptr)->~QPaintEngine();
}

void QPaintEngine_DestroyQPaintEngineDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QPaintEngineState_Font(void* ptr)
{
	return new QFont(static_cast<QPaintEngineState*>(ptr)->font());
}

double QPaintEngineState_Opacity(void* ptr)
{
	return static_cast<QPaintEngineState*>(ptr)->opacity();
}

long long QPaintEngineState_State(void* ptr)
{
	return static_cast<QPaintEngineState*>(ptr)->state();
}

void* QPaintEngineState_Transform(void* ptr)
{
	return new QTransform(static_cast<QPaintEngineState*>(ptr)->transform());
}

class MyQPaintEvent: public QPaintEvent
{
public:
	MyQPaintEvent(const QRegion &paintRegion) : QPaintEvent(paintRegion) {};
	MyQPaintEvent(const QRect &paintRect) : QPaintEvent(paintRect) {};
};

void* QPaintEvent_NewQPaintEvent(void* paintRegion)
{
	return new MyQPaintEvent(*static_cast<QRegion*>(paintRegion));
}

void* QPaintEvent_NewQPaintEvent2(void* paintRect)
{
	return new MyQPaintEvent(*static_cast<QRect*>(paintRect));
}

void* QPaintEvent_Rect(void* ptr)
{
	return const_cast<QRect*>(&static_cast<QPaintEvent*>(ptr)->rect());
}

void* QPainter_NewQPainter()
{
	return new QPainter();
}

void* QPainter_NewQPainter2(void* device)
{
	if (dynamic_cast<QWidget*>(static_cast<QObject*>(device))) {
		return new QPainter(static_cast<QWidget*>(device));
	} else {
		return new QPainter(static_cast<QPaintDevice*>(device));
	}
}

void* QPainter_Background(void* ptr)
{
	return const_cast<QBrush*>(&static_cast<QPainter*>(ptr)->background());
}

char QPainter_Begin(void* ptr, void* device)
{
	return static_cast<QPainter*>(ptr)->begin(static_cast<QPaintDevice*>(device));
}

char QPainter_End(void* ptr)
{
	return static_cast<QPainter*>(ptr)->end();
}

void* QPainter_Font(void* ptr)
{
	return const_cast<QFont*>(&static_cast<QPainter*>(ptr)->font());
}

double QPainter_Opacity(void* ptr)
{
	return static_cast<QPainter*>(ptr)->opacity();
}

void QPainter_Scale(void* ptr, double sx, double sy)
{
	static_cast<QPainter*>(ptr)->scale(sx, sy);
}

void QPainter_SetBackground(void* ptr, void* brush)
{
	static_cast<QPainter*>(ptr)->setBackground(*static_cast<QBrush*>(brush));
}

void QPainter_SetFont(void* ptr, void* font)
{
	static_cast<QPainter*>(ptr)->setFont(*static_cast<QFont*>(font));
}

void QPainter_SetWindow(void* ptr, void* rectangle)
{
	static_cast<QPainter*>(ptr)->setWindow(*static_cast<QRect*>(rectangle));
}

void QPainter_SetWindow2(void* ptr, int x, int y, int width, int height)
{
	static_cast<QPainter*>(ptr)->setWindow(x, y, width, height);
}

void* QPainter_Transform(void* ptr)
{
	return const_cast<QTransform*>(&static_cast<QPainter*>(ptr)->transform());
}

void* QPainter_Window(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QPainter*>(ptr)->window(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QPainter_DestroyQPainter(void* ptr)
{
	static_cast<QPainter*>(ptr)->~QPainter();
}

void* QPainter___drawLines_lines_atList2(void* ptr, int i)
{
	return ({ QLineF tmpValue = ({QLineF tmp = static_cast<QVector<QLineF>*>(ptr)->at(i); if (i == static_cast<QVector<QLineF>*>(ptr)->size()-1) { static_cast<QVector<QLineF>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QLineF(tmpValue.p1(), tmpValue.p2()); });
}

void QPainter___drawLines_lines_setList2(void* ptr, void* i)
{
	static_cast<QVector<QLineF>*>(ptr)->append(*static_cast<QLineF*>(i));
}

void* QPainter___drawLines_lines_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QLineF>();
}

void* QPainter___drawLines_pointPairs_atList4(void* ptr, int i)
{
	return ({ QPointF tmpValue = ({QPointF tmp = static_cast<QVector<QPointF>*>(ptr)->at(i); if (i == static_cast<QVector<QPointF>*>(ptr)->size()-1) { static_cast<QVector<QPointF>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPointF(tmpValue.x(), tmpValue.y()); });
}

void QPainter___drawLines_pointPairs_setList4(void* ptr, void* i)
{
	static_cast<QVector<QPointF>*>(ptr)->append(*static_cast<QPointF*>(i));
}

void* QPainter___drawLines_pointPairs_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPointF>();
}

void* QPainter___drawLines_lines_atList6(void* ptr, int i)
{
	return ({ QLine tmpValue = ({QLine tmp = static_cast<QVector<QLine>*>(ptr)->at(i); if (i == static_cast<QVector<QLine>*>(ptr)->size()-1) { static_cast<QVector<QLine>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QLine(tmpValue.p1(), tmpValue.p2()); });
}

void QPainter___drawLines_lines_setList6(void* ptr, void* i)
{
	static_cast<QVector<QLine>*>(ptr)->append(*static_cast<QLine*>(i));
}

void* QPainter___drawLines_lines_newList6(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QLine>();
}

void* QPainter___drawLines_pointPairs_atList8(void* ptr, int i)
{
	return ({ QPoint tmpValue = ({QPoint tmp = static_cast<QVector<QPoint>*>(ptr)->at(i); if (i == static_cast<QVector<QPoint>*>(ptr)->size()-1) { static_cast<QVector<QPoint>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void QPainter___drawLines_pointPairs_setList8(void* ptr, void* i)
{
	static_cast<QVector<QPoint>*>(ptr)->append(*static_cast<QPoint*>(i));
}

void* QPainter___drawLines_pointPairs_newList8(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPoint>();
}

void* QPainter___drawRects_rectangles_atList2(void* ptr, int i)
{
	return ({ QRectF tmpValue = ({QRectF tmp = static_cast<QVector<QRectF>*>(ptr)->at(i); if (i == static_cast<QVector<QRectF>*>(ptr)->size()-1) { static_cast<QVector<QRectF>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QRectF(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QPainter___drawRects_rectangles_setList2(void* ptr, void* i)
{
	static_cast<QVector<QRectF>*>(ptr)->append(*static_cast<QRectF*>(i));
}

void* QPainter___drawRects_rectangles_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QRectF>();
}

void* QPainter___drawRects_rectangles_atList4(void* ptr, int i)
{
	return ({ QRect tmpValue = ({QRect tmp = static_cast<QVector<QRect>*>(ptr)->at(i); if (i == static_cast<QVector<QRect>*>(ptr)->size()-1) { static_cast<QVector<QRect>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QPainter___drawRects_rectangles_setList4(void* ptr, void* i)
{
	static_cast<QVector<QRect>*>(ptr)->append(*static_cast<QRect*>(i));
}

void* QPainter___drawRects_rectangles_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QRect>();
}

int QPalette_NColorRoles_Type()
{
	return QPalette::NColorRoles;
}

class MyQPixmap: public QPixmap
{
public:
	MyQPixmap() : QPixmap() {};
	MyQPixmap(int width, int height) : QPixmap(width, height) {};
	MyQPixmap(const QSize &size) : QPixmap(size) {};
	MyQPixmap(const QString &fileName, const char *format = Q_NULLPTR, Qt::ImageConversionFlags flags = Qt::AutoColor) : QPixmap(fileName, format, flags) {};
	MyQPixmap(const QPixmap &pixmap) : QPixmap(pixmap) {};
	 ~MyQPixmap() { callbackQPixmap_DestroyQPixmap(this); };
	QPaintEngine * paintEngine() const { return static_cast<QPaintEngine*>(callbackQPixmap_PaintEngine(const_cast<void*>(static_cast<const void*>(this)))); };
};

void* QPixmap_NewQPixmap()
{
	return new MyQPixmap();
}

void* QPixmap_NewQPixmap2(void* size)
{
	return new MyQPixmap(*static_cast<QSize*>(size));
}

void* QPixmap_NewQPixmap3(struct QtGui_PackedString fileName, char* format, long long flags)
{
	return new MyQPixmap(QString::fromUtf8(fileName.data, fileName.len), const_cast<const char*>(format), static_cast<Qt::ImageConversionFlag>(flags));
}

void* QPixmap_NewQPixmap5(void* pixmap)
{
	return new MyQPixmap(*static_cast<QPixmap*>(pixmap));
}

void QPixmap_Fill(void* ptr, void* color)
{
	static_cast<QPixmap*>(ptr)->fill(*static_cast<QColor*>(color));
}

void* QPixmap_Mask(void* ptr)
{
	return new QBitmap(static_cast<QPixmap*>(ptr)->mask());
}

void* QPixmap_Rect(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QPixmap*>(ptr)->rect(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void* QPixmap_Size(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QPixmap*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QPixmap_DestroyQPixmap(void* ptr)
{
	static_cast<QPixmap*>(ptr)->~QPixmap();
}

void QPixmap_DestroyQPixmapDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QPixmap_PaintEngine(void* ptr)
{
	return static_cast<QPixmap*>(ptr)->paintEngine();
}

void* QPixmap_PaintEngineDefault(void* ptr)
{
	if (dynamic_cast<QBitmap*>(static_cast<QPixmap*>(ptr))) {
		return static_cast<QBitmap*>(ptr)->QBitmap::paintEngine();
	} else {
		return static_cast<QPixmap*>(ptr)->QPixmap::paintEngine();
	}
}

void* QPolygon_NewQPolygon()
{
	return new QPolygon();
}

void* QPolygon_NewQPolygon2(int size)
{
	return new QPolygon(size);
}

void* QPolygon_NewQPolygon3(void* points)
{
	return new QPolygon(*static_cast<QVector<QPoint>*>(points));
}

void* QPolygon_NewQPolygon5(void* rectangle, char closed)
{
	return new QPolygon(*static_cast<QRect*>(rectangle), closed != 0);
}

void QPolygon_Point(void* ptr, int index, int x, int y)
{
	static_cast<QPolygon*>(ptr)->point(index, &x, &y);
}

void* QPolygon_Point2(void* ptr, int index)
{
	return ({ QPoint tmpValue = static_cast<QPolygon*>(ptr)->point(index); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void QPolygon_SetPoint(void* ptr, int index, int x, int y)
{
	static_cast<QPolygon*>(ptr)->setPoint(index, x, y);
}

void QPolygon_SetPoint2(void* ptr, int index, void* point)
{
	static_cast<QPolygon*>(ptr)->setPoint(index, *static_cast<QPoint*>(point));
}

void QPolygon_DestroyQPolygon(void* ptr)
{
	static_cast<QPolygon*>(ptr)->~QPolygon();
}

void* QPolygon___QPolygon_points_atList3(void* ptr, int i)
{
	return ({ QPoint tmpValue = ({QPoint tmp = static_cast<QVector<QPoint>*>(ptr)->at(i); if (i == static_cast<QVector<QPoint>*>(ptr)->size()-1) { static_cast<QVector<QPoint>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void QPolygon___QPolygon_points_setList3(void* ptr, void* i)
{
	static_cast<QVector<QPoint>*>(ptr)->append(*static_cast<QPoint*>(i));
}

void* QPolygon___QPolygon_points_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPoint>();
}

void* QPolygon___QPolygon_v_atList4(void* ptr, int i)
{
	return ({ QPoint tmpValue = ({QPoint tmp = static_cast<QVector<QPoint>*>(ptr)->at(i); if (i == static_cast<QVector<QPoint>*>(ptr)->size()-1) { static_cast<QVector<QPoint>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QPoint(tmpValue.x(), tmpValue.y()); });
}

void QPolygon___QPolygon_v_setList4(void* ptr, void* i)
{
	static_cast<QVector<QPoint>*>(ptr)->append(*static_cast<QPoint*>(i));
}

void* QPolygon___QPolygon_v_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QPoint>();
}

void* QPolygon___QVector_other_atList4(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___QVector_other_setList4(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygon___QVector_other_newList4(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___QVector_other_atList5(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___QVector_other_setList5(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygon___QVector_other_newList5(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___append_value_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___append_value_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygon___append_value_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___fill_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___fill_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygon___fill_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___fromList_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___fromList_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygon___fromList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___fromList_list_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___fromList_list_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygon___fromList_list_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___fromStdVector_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___fromStdVector_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygon___fromStdVector_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___mid_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___mid_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygon___mid_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___swap_other_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___swap_other_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygon___swap_other_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QPolygon___toList_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QPolygon___toList_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QPolygon___toList_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QRegion_NewQRegion()
{
	return new QRegion();
}

void* QRegion_NewQRegion2(int x, int y, int w, int h, long long t)
{
	return new QRegion(x, y, w, h, static_cast<QRegion::RegionType>(t));
}

void* QRegion_NewQRegion3(void* r, long long t)
{
	return new QRegion(*static_cast<QRect*>(r), static_cast<QRegion::RegionType>(t));
}

void* QRegion_NewQRegion4(void* a, long long fillRule)
{
	return new QRegion(*static_cast<QPolygon*>(a), static_cast<Qt::FillRule>(fillRule));
}

void* QRegion_NewQRegion5(void* r)
{
	return new QRegion(*static_cast<QRegion*>(r));
}

void* QRegion_NewQRegion6(void* other)
{
	return new QRegion(*static_cast<QRegion*>(other));
}

void* QRegion_NewQRegion7(void* bm)
{
	return new QRegion(*static_cast<QBitmap*>(bm));
}

char QRegion_Contains(void* ptr, void* p)
{
	return static_cast<QRegion*>(ptr)->contains(*static_cast<QPoint*>(p));
}

char QRegion_Contains2(void* ptr, void* r)
{
	return static_cast<QRegion*>(ptr)->contains(*static_cast<QRect*>(r));
}

void* QRegion___rects_atList(void* ptr, int i)
{
	return ({ QRect tmpValue = ({QRect tmp = static_cast<QVector<QRect>*>(ptr)->at(i); if (i == static_cast<QVector<QRect>*>(ptr)->size()-1) { static_cast<QVector<QRect>*>(ptr)->~QVector(); free(ptr); }; tmp; }); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QRegion___rects_setList(void* ptr, void* i)
{
	static_cast<QVector<QRect>*>(ptr)->append(*static_cast<QRect*>(i));
}

void* QRegion___rects_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QVector<QRect>();
}

unsigned short QRgba64_Red(void* ptr)
{
	return static_cast<QRgba64*>(ptr)->red();
}

class MyQScreen: public QScreen
{
public:
	void Signal_GeometryChanged(const QRect & geometry) { callbackQScreen_GeometryChanged(this, const_cast<QRect*>(&geometry)); };
	 ~MyQScreen() { callbackQScreen_DestroyQScreen(this); };
	void childEvent(QChildEvent * event) { callbackQScreen_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQScreen_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQScreen_CustomEvent(this, event); };
	void deleteLater() { callbackQScreen_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQScreen_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQScreen_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool event(QEvent * e) { return callbackQScreen_Event(this, e) != 0; };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQScreen_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtGui_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQScreen_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQScreen_TimerEvent(this, event); };
};

Q_DECLARE_METATYPE(MyQScreen*)

int QScreen_QScreen_QRegisterMetaType(){qRegisterMetaType<QScreen*>(); return qRegisterMetaType<MyQScreen*>();}

void* QScreen_Geometry(void* ptr)
{
	return ({ QRect tmpValue = static_cast<QScreen*>(ptr)->geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

void QScreen_ConnectGeometryChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(const QRect &)>(&QScreen::geometryChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(const QRect &)>(&MyQScreen::Signal_GeometryChanged), static_cast<Qt::ConnectionType>(t));
}

void QScreen_DisconnectGeometryChanged(void* ptr)
{
	QObject::disconnect(static_cast<QScreen*>(ptr), static_cast<void (QScreen::*)(const QRect &)>(&QScreen::geometryChanged), static_cast<MyQScreen*>(ptr), static_cast<void (MyQScreen::*)(const QRect &)>(&MyQScreen::Signal_GeometryChanged));
}

void QScreen_GeometryChanged(void* ptr, void* geometry)
{
	static_cast<QScreen*>(ptr)->geometryChanged(*static_cast<QRect*>(geometry));
}

struct QtGui_PackedString QScreen_Model(void* ptr)
{
#ifndef Q_OS_WIN
	return ({ QByteArray t131f94 = static_cast<QScreen*>(ptr)->model().toUtf8(); QtGui_PackedString { const_cast<char*>(t131f94.prepend("WHITESPACE").constData()+10), t131f94.size()-10 }; });
#endif
}

struct QtGui_PackedString QScreen_Name(void* ptr)
{
	return ({ QByteArray tc60f02 = static_cast<QScreen*>(ptr)->name().toUtf8(); QtGui_PackedString { const_cast<char*>(tc60f02.prepend("WHITESPACE").constData()+10), tc60f02.size()-10 }; });
}

long long QScreen_Orientation(void* ptr)
{
	return static_cast<QScreen*>(ptr)->orientation();
}

void* QScreen_Size(void* ptr)
{
	return ({ QSize tmpValue = static_cast<QScreen*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QScreen_DestroyQScreen(void* ptr)
{
	static_cast<QScreen*>(ptr)->~QScreen();
}

void QScreen_DestroyQScreenDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QScreen___virtualSiblings_atList(void* ptr, int i)
{
	return ({QScreen * tmp = static_cast<QList<QScreen *>*>(ptr)->at(i); if (i == static_cast<QList<QScreen *>*>(ptr)->size()-1) { static_cast<QList<QScreen *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QScreen___virtualSiblings_setList(void* ptr, void* i)
{
	static_cast<QList<QScreen *>*>(ptr)->append(static_cast<QScreen*>(i));
}

void* QScreen___virtualSiblings_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QScreen *>();
}

void* QScreen___children_atList(void* ptr, int i)
{
	return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QScreen___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QScreen___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject *>();
}

void* QScreen___dynamicPropertyNames_atList(void* ptr, int i)
{
	return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QScreen___dynamicPropertyNames_setList(void* ptr, void* i)
{
	static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QScreen___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QByteArray>();
}

void* QScreen___findChildren_atList(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QScreen___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QScreen___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QScreen___findChildren_atList3(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QScreen___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QScreen___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void* QScreen___qFindChildren_atList2(void* ptr, int i)
{
	return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QScreen___qFindChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QScreen___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QObject*>();
}

void QScreen_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QScreen*>(ptr)->QScreen::childEvent(static_cast<QChildEvent*>(event));
}

void QScreen_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QScreen*>(ptr)->QScreen::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QScreen_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QScreen*>(ptr)->QScreen::customEvent(static_cast<QEvent*>(event));
}

void QScreen_DeleteLaterDefault(void* ptr)
{
		static_cast<QScreen*>(ptr)->QScreen::deleteLater();
}

void QScreen_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QScreen*>(ptr)->QScreen::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QScreen_EventDefault(void* ptr, void* e)
{
		return static_cast<QScreen*>(ptr)->QScreen::event(static_cast<QEvent*>(e));
}

char QScreen_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(watched))) {
		return static_cast<QScreen*>(ptr)->QScreen::eventFilter(static_cast<QWindow*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QScreen*>(ptr)->QScreen::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QScreen_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QScreen*>(ptr)->QScreen::timerEvent(static_cast<QTimerEvent*>(event));
}

class MyQSurface: public QSurface
{
public:
	QSurfaceFormat format() const { return *static_cast<QSurfaceFormat*>(callbackQSurface_Format(const_cast<void*>(static_cast<const void*>(this)))); };
	QSize size() const { return *static_cast<QSize*>(callbackQSurface_Size(const_cast<void*>(static_cast<const void*>(this)))); };
	QSurface::SurfaceType surfaceType() const { return static_cast<QSurface::SurfaceType>(callbackQSurface_SurfaceType(const_cast<void*>(static_cast<const void*>(this)))); };
	 ~MyQSurface() { callbackQSurface_DestroyQSurface(this); };
};

void* QSurface_Format(void* ptr)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(ptr))) {
		return new QSurfaceFormat(static_cast<QWindow*>(ptr)->format());
	} else {
		return new QSurfaceFormat(static_cast<QSurface*>(ptr)->format());
	}
}

void* QSurface_Size(void* ptr)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(ptr))) {
		return ({ QSize tmpValue = static_cast<QWindow*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
	} else {
		return ({ QSize tmpValue = static_cast<QSurface*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
	}
}

long long QSurface_SurfaceType(void* ptr)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(ptr))) {
		return static_cast<QWindow*>(ptr)->surfaceType();
	} else {
		return static_cast<QSurface*>(ptr)->surfaceType();
	}
}

void QSurface_DestroyQSurface(void* ptr)
{
	static_cast<QSurface*>(ptr)->~QSurface();
}

void QSurface_DestroyQSurfaceDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QSurfaceFormat_NewQSurfaceFormat()
{
	return new QSurfaceFormat();
}

void* QSurfaceFormat_NewQSurfaceFormat2(long long options)
{
	return new QSurfaceFormat(static_cast<QSurfaceFormat::FormatOption>(options));
}

void* QSurfaceFormat_NewQSurfaceFormat3(void* other)
{
	return new QSurfaceFormat(*static_cast<QSurfaceFormat*>(other));
}

void QSurfaceFormat_SetOption(void* ptr, long long option, char on)
{
	static_cast<QSurfaceFormat*>(ptr)->setOption(static_cast<QSurfaceFormat::FormatOption>(option), on != 0);
}

void QSurfaceFormat_DestroyQSurfaceFormat(void* ptr)
{
	static_cast<QSurfaceFormat*>(ptr)->~QSurfaceFormat();
}

void* QTouchDevice_NewQTouchDevice()
{
	return new QTouchDevice();
}

struct QtGui_PackedString QTouchDevice_Name(void* ptr)
{
	return ({ QByteArray td71303 = static_cast<QTouchDevice*>(ptr)->name().toUtf8(); QtGui_PackedString { const_cast<char*>(td71303.prepend("WHITESPACE").constData()+10), td71303.size()-10 }; });
}

void QTouchDevice_SetName(void* ptr, struct QtGui_PackedString name)
{
	static_cast<QTouchDevice*>(ptr)->setName(QString::fromUtf8(name.data, name.len));
}

long long QTouchDevice_Type(void* ptr)
{
	return static_cast<QTouchDevice*>(ptr)->type();
}

void QTouchDevice_DestroyQTouchDevice(void* ptr)
{
	static_cast<QTouchDevice*>(ptr)->~QTouchDevice();
}

void* QTouchDevice___devices_atList(void* ptr, int i)
{
	return const_cast<QTouchDevice*>(({const QTouchDevice * tmp = static_cast<QList<QTouchDevice *>*>(ptr)->at(i); if (i == static_cast<QList<QTouchDevice *>*>(ptr)->size()-1) { static_cast<QList<QTouchDevice *>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QTouchDevice___devices_setList(void* ptr, void* i)
{
	static_cast<QList<QTouchDevice *>*>(ptr)->append(static_cast<QTouchDevice*>(i));
}

void* QTouchDevice___devices_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<const QTouchDevice *>();
}

class MyQTouchEvent: public QTouchEvent
{
public:
	 ~MyQTouchEvent() { callbackQTouchEvent_DestroyQTouchEvent(this); };
};

void* QTouchEvent_Window(void* ptr)
{
	return static_cast<QTouchEvent*>(ptr)->window();
}

void QTouchEvent_DestroyQTouchEvent(void* ptr)
{
	static_cast<QTouchEvent*>(ptr)->~QTouchEvent();
}

void QTouchEvent_DestroyQTouchEventDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QTouchEvent___QTouchEvent_touchPoints_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QTouchEvent::TouchPoint>();
}

void* QTouchEvent___setTouchPoints_atouchPoints_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QTouchEvent::TouchPoint>();
}

void* QTouchEvent___touchPoints_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QTouchEvent::TouchPoint>();
}

void* QTouchEvent____touchPoints_newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QTouchEvent::TouchPoint>();
}

void* QTouchEvent___set_touchPoints__newList(void* ptr)
{
	Q_UNUSED(ptr);
	return new QList<QTouchEvent::TouchPoint>();
}

void* QTransform_NewQTransform2()
{
	return new QTransform();
}

void* QTransform_NewQTransform3(double m11, double m12, double m13, double m21, double m22, double m23, double m31, double m32, double m33)
{
	return new QTransform(m11, m12, m13, m21, m22, m23, m31, m32, m33);
}

void* QTransform_NewQTransform4(double m11, double m12, double m21, double m22, double dx, double dy)
{
	return new QTransform(m11, m12, m21, m22, dx, dy);
}

void* QTransform_NewQTransform5(void* matrix)
{
	return new QTransform(*static_cast<QMatrix*>(matrix));
}

void QTransform_Reset(void* ptr)
{
	static_cast<QTransform*>(ptr)->reset();
}

void* QTransform_Scale(void* ptr, double sx, double sy)
{
	return new QTransform(static_cast<QTransform*>(ptr)->scale(sx, sy));
}

long long QTransform_Type(void* ptr)
{
	return static_cast<QTransform*>(ptr)->type();
}

void* QVector2D_NewQVector2D()
{
	return new QVector2D();
}

void* QVector2D_NewQVector2D3(float xpos, float ypos)
{
	return new QVector2D(xpos, ypos);
}

void* QVector2D_NewQVector2D4(void* point)
{
	return new QVector2D(*static_cast<QPoint*>(point));
}

void* QVector2D_NewQVector2D5(void* point)
{
	return new QVector2D(*static_cast<QPointF*>(point));
}

void* QVector2D_NewQVector2D6(void* vector)
{
	return new QVector2D(*static_cast<QVector3D*>(vector));
}

void* QVector2D_NewQVector2D7(void* vector)
{
	return new QVector2D(*static_cast<QVector4D*>(vector));
}

float QVector2D_X(void* ptr)
{
	return static_cast<QVector2D*>(ptr)->x();
}

float QVector2D_Y(void* ptr)
{
	return static_cast<QVector2D*>(ptr)->y();
}

void* QVector3D_NewQVector3D()
{
	return new QVector3D();
}

void* QVector3D_NewQVector3D3(float xpos, float ypos, float zpos)
{
	return new QVector3D(xpos, ypos, zpos);
}

void* QVector3D_NewQVector3D4(void* point)
{
	return new QVector3D(*static_cast<QPoint*>(point));
}

void* QVector3D_NewQVector3D5(void* point)
{
	return new QVector3D(*static_cast<QPointF*>(point));
}

void* QVector3D_NewQVector3D6(void* vector)
{
	return new QVector3D(*static_cast<QVector2D*>(vector));
}

void* QVector3D_NewQVector3D7(void* vector, float zpos)
{
	return new QVector3D(*static_cast<QVector2D*>(vector), zpos);
}

void* QVector3D_NewQVector3D8(void* vector)
{
	return new QVector3D(*static_cast<QVector4D*>(vector));
}

void* QVector3D_Project(void* ptr, void* modelView, void* projection, void* viewport)
{
	return new QVector3D(static_cast<QVector3D*>(ptr)->project(*static_cast<QMatrix4x4*>(modelView), *static_cast<QMatrix4x4*>(projection), *static_cast<QRect*>(viewport)));
}

float QVector3D_X(void* ptr)
{
	return static_cast<QVector3D*>(ptr)->x();
}

float QVector3D_Y(void* ptr)
{
	return static_cast<QVector3D*>(ptr)->y();
}

float QVector3D_Z(void* ptr)
{
	return static_cast<QVector3D*>(ptr)->z();
}

void* QVector4D_NewQVector4D()
{
	return new QVector4D();
}

void* QVector4D_NewQVector4D3(float xpos, float ypos, float zpos, float wpos)
{
	return new QVector4D(xpos, ypos, zpos, wpos);
}

void* QVector4D_NewQVector4D4(void* point)
{
	return new QVector4D(*static_cast<QPoint*>(point));
}

void* QVector4D_NewQVector4D5(void* point)
{
	return new QVector4D(*static_cast<QPointF*>(point));
}

void* QVector4D_NewQVector4D6(void* vector)
{
	return new QVector4D(*static_cast<QVector2D*>(vector));
}

void* QVector4D_NewQVector4D7(void* vector, float zpos, float wpos)
{
	return new QVector4D(*static_cast<QVector2D*>(vector), zpos, wpos);
}

void* QVector4D_NewQVector4D8(void* vector)
{
	return new QVector4D(*static_cast<QVector3D*>(vector));
}

void* QVector4D_NewQVector4D9(void* vector, float wpos)
{
	return new QVector4D(*static_cast<QVector3D*>(vector), wpos);
}

void QVector4D_SetW(void* ptr, float w)
{
	static_cast<QVector4D*>(ptr)->setW(w);
}

float QVector4D_W(void* ptr)
{
	return static_cast<QVector4D*>(ptr)->w();
}

float QVector4D_X(void* ptr)
{
	return static_cast<QVector4D*>(ptr)->x();
}

float QVector4D_Y(void* ptr)
{
	return static_cast<QVector4D*>(ptr)->y();
}

float QVector4D_Z(void* ptr)
{
	return static_cast<QVector4D*>(ptr)->z();
}

class MyQWheelEvent: public QWheelEvent
{
public:
	MyQWheelEvent(const QPointF &pos, const QPointF &globalPos, QPoint pixelDelta, QPoint angleDelta, int qt4Delta, Qt::Orientation qt4Orientation, Qt::MouseButtons buttons, Qt::KeyboardModifiers modifiers) : QWheelEvent(pos, globalPos, pixelDelta, angleDelta, qt4Delta, qt4Orientation, buttons, modifiers) {};
	MyQWheelEvent(const QPointF &pos, const QPointF &globalPos, QPoint pixelDelta, QPoint angleDelta, int qt4Delta, Qt::Orientation qt4Orientation, Qt::MouseButtons buttons, Qt::KeyboardModifiers modifiers, Qt::ScrollPhase phase) : QWheelEvent(pos, globalPos, pixelDelta, angleDelta, qt4Delta, qt4Orientation, buttons, modifiers, phase) {};
	MyQWheelEvent(const QPointF &pos, const QPointF &globalPos, QPoint pixelDelta, QPoint angleDelta, int qt4Delta, Qt::Orientation qt4Orientation, Qt::MouseButtons buttons, Qt::KeyboardModifiers modifiers, Qt::ScrollPhase phase, Qt::MouseEventSource source) : QWheelEvent(pos, globalPos, pixelDelta, angleDelta, qt4Delta, qt4Orientation, buttons, modifiers, phase, source) {};
	MyQWheelEvent(const QPointF &pos, const QPointF &globalPos, QPoint pixelDelta, QPoint angleDelta, int qt4Delta, Qt::Orientation qt4Orientation, Qt::MouseButtons buttons, Qt::KeyboardModifiers modifiers, Qt::ScrollPhase phase, Qt::MouseEventSource source, bool inverted) : QWheelEvent(pos, globalPos, pixelDelta, angleDelta, qt4Delta, qt4Orientation, buttons, modifiers, phase, source, inverted) {};
	MyQWheelEvent(QPointF pos, QPointF globalPos, QPoint pixelDelta, QPoint angleDelta, Qt::MouseButtons buttons, Qt::KeyboardModifiers modifiers, Qt::ScrollPhase phase, bool inverted, Qt::MouseEventSource source = Qt::MouseEventNotSynthesized) : QWheelEvent(pos, globalPos, pixelDelta, angleDelta, buttons, modifiers, phase, inverted, source) {};
};

void* QWheelEvent_NewQWheelEvent3(void* pos, void* globalPos, void* pixelDelta, void* angleDelta, int qt4Delta, long long qt4Orientation, long long buttons, long long modifiers)
{
	return new MyQWheelEvent(*static_cast<QPointF*>(pos), *static_cast<QPointF*>(globalPos), *static_cast<QPoint*>(pixelDelta), *static_cast<QPoint*>(angleDelta), qt4Delta, static_cast<Qt::Orientation>(qt4Orientation), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers));
}

void* QWheelEvent_NewQWheelEvent4(void* pos, void* globalPos, void* pixelDelta, void* angleDelta, int qt4Delta, long long qt4Orientation, long long buttons, long long modifiers, long long phase)
{
	return new MyQWheelEvent(*static_cast<QPointF*>(pos), *static_cast<QPointF*>(globalPos), *static_cast<QPoint*>(pixelDelta), *static_cast<QPoint*>(angleDelta), qt4Delta, static_cast<Qt::Orientation>(qt4Orientation), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers), static_cast<Qt::ScrollPhase>(phase));
}

void* QWheelEvent_NewQWheelEvent5(void* pos, void* globalPos, void* pixelDelta, void* angleDelta, int qt4Delta, long long qt4Orientation, long long buttons, long long modifiers, long long phase, long long source)
{
	return new MyQWheelEvent(*static_cast<QPointF*>(pos), *static_cast<QPointF*>(globalPos), *static_cast<QPoint*>(pixelDelta), *static_cast<QPoint*>(angleDelta), qt4Delta, static_cast<Qt::Orientation>(qt4Orientation), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers), static_cast<Qt::ScrollPhase>(phase), static_cast<Qt::MouseEventSource>(source));
}

void* QWheelEvent_NewQWheelEvent6(void* pos, void* globalPos, void* pixelDelta, void* angleDelta, int qt4Delta, long long qt4Orientation, long long buttons, long long modifiers, long long phase, long long source, char inverted)
{
	return new MyQWheelEvent(*static_cast<QPointF*>(pos), *static_cast<QPointF*>(globalPos), *static_cast<QPoint*>(pixelDelta), *static_cast<QPoint*>(angleDelta), qt4Delta, static_cast<Qt::Orientation>(qt4Orientation), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers), static_cast<Qt::ScrollPhase>(phase), static_cast<Qt::MouseEventSource>(source), inverted != 0);
}

void* QWheelEvent_NewQWheelEvent7(void* pos, void* globalPos, void* pixelDelta, void* angleDelta, long long buttons, long long modifiers, long long phase, char inverted, long long source)
{
	return new MyQWheelEvent(*static_cast<QPointF*>(pos), *static_cast<QPointF*>(globalPos), *static_cast<QPoint*>(pixelDelta), *static_cast<QPoint*>(angleDelta), static_cast<Qt::MouseButton>(buttons), static_cast<Qt::KeyboardModifier>(modifiers), static_cast<Qt::ScrollPhase>(phase), inverted != 0, static_cast<Qt::MouseEventSource>(source));
}

long long QWheelEvent_Buttons(void* ptr)
{
	return static_cast<QWheelEvent*>(ptr)->buttons();
}

long long QWheelEvent_Source(void* ptr)
{
	return static_cast<QWheelEvent*>(ptr)->source();
}

int QWheelEvent_X(void* ptr)
{
	return static_cast<QWheelEvent*>(ptr)->x();
}

int QWheelEvent_Y(void* ptr)
{
	return static_cast<QWheelEvent*>(ptr)->y();
}

class MyQWindow: public QWindow
{
public:
	MyQWindow(QScreen *targetScreen = Q_NULLPTR) : QWindow(targetScreen) {QWindow_QWindow_QRegisterMetaType();};
	MyQWindow(QWindow *parent) : QWindow(parent) {QWindow_QWindow_QRegisterMetaType();};
	bool event(QEvent * ev) { return callbackQWindow_Event(this, ev) != 0; };
	void focusInEvent(QFocusEvent * ev) { callbackQWindow_FocusInEvent(this, ev); };
	void focusOutEvent(QFocusEvent * ev) { callbackQWindow_FocusOutEvent(this, ev); };
	void Signal_HeightChanged(int arg) { callbackQWindow_HeightChanged(this, arg); };
	void hide() { callbackQWindow_Hide(this); };
	void keyPressEvent(QKeyEvent * ev) { callbackQWindow_KeyPressEvent(this, ev); };
	void keyReleaseEvent(QKeyEvent * ev) { callbackQWindow_KeyReleaseEvent(this, ev); };
	void mouseDoubleClickEvent(QMouseEvent * ev) { callbackQWindow_MouseDoubleClickEvent(this, ev); };
	void mouseMoveEvent(QMouseEvent * ev) { callbackQWindow_MouseMoveEvent(this, ev); };
	void mousePressEvent(QMouseEvent * ev) { callbackQWindow_MousePressEvent(this, ev); };
	void mouseReleaseEvent(QMouseEvent * ev) { callbackQWindow_MouseReleaseEvent(this, ev); };
	void moveEvent(QMoveEvent * ev) { callbackQWindow_MoveEvent(this, ev); };
	void Signal_OpacityChanged(qreal opacity) { callbackQWindow_OpacityChanged(this, opacity); };
	void setHeight(int arg) { callbackQWindow_SetHeight(this, arg); };
	void show() { callbackQWindow_Show(this); };
	QSize size() const { return *static_cast<QSize*>(callbackQWindow_Size(const_cast<void*>(static_cast<const void*>(this)))); };
	void touchEvent(QTouchEvent * ev) { callbackQWindow_TouchEvent(this, ev); };
	void Signal_VisibleChanged(bool arg) { callbackQWindow_VisibleChanged(this, arg); };
	void wheelEvent(QWheelEvent * ev) { callbackQWindow_WheelEvent(this, ev); };
	void Signal_WidthChanged(int arg) { callbackQWindow_WidthChanged(this, arg); };
	void Signal_XChanged(int arg) { callbackQWindow_XChanged(this, arg); };
	void Signal_YChanged(int arg) { callbackQWindow_YChanged(this, arg); };
	 ~MyQWindow() { callbackQWindow_DestroyQWindow(this); };
	void childEvent(QChildEvent * event) { callbackQWindow_ChildEvent(this, event); };
	void connectNotify(const QMetaMethod & sign) { callbackQWindow_ConnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	void customEvent(QEvent * event) { callbackQWindow_CustomEvent(this, event); };
	void deleteLater() { callbackQWindow_DeleteLater(this); };
	void Signal_Destroyed(QObject * obj) { callbackQWindow_Destroyed(this, obj); };
	void disconnectNotify(const QMetaMethod & sign) { callbackQWindow_DisconnectNotify(this, const_cast<QMetaMethod*>(&sign)); };
	bool eventFilter(QObject * watched, QEvent * event) { return callbackQWindow_EventFilter(this, watched, event) != 0; };
	void Signal_ObjectNameChanged(const QString & objectName) { QByteArray taa2c4f = objectName.toUtf8(); QtGui_PackedString objectNamePacked = { const_cast<char*>(taa2c4f.prepend("WHITESPACE").constData()+10), taa2c4f.size()-10 };callbackQWindow_ObjectNameChanged(this, objectNamePacked); };
	void timerEvent(QTimerEvent * event) { callbackQWindow_TimerEvent(this, event); };
	QSurfaceFormat format() const { return *static_cast<QSurfaceFormat*>(callbackQWindow_Format(const_cast<void*>(static_cast<const void*>(this)))); };
	QSurface::SurfaceType surfaceType() const { return static_cast<QSurface::SurfaceType>(callbackQWindow_SurfaceType(const_cast<void*>(static_cast<const void*>(this)))); };
};

Q_DECLARE_METATYPE(MyQWindow*)

int QWindow_QWindow_QRegisterMetaType(){qRegisterMetaType<QWindow*>(); return qRegisterMetaType<MyQWindow*>();}

void* QWindow_NewQWindow(void* targetScreen)
{
	return new MyQWindow(static_cast<QScreen*>(targetScreen));
}

void* QWindow_NewQWindow2(void* parent)
{
	return new MyQWindow(static_cast<QWindow*>(parent));
}

void QWindow_Destroy(void* ptr)
{
		static_cast<QWindow*>(ptr)->destroy();
}

char QWindow_Event(void* ptr, void* ev)
{
		return static_cast<QWindow*>(ptr)->event(static_cast<QEvent*>(ev));
}

char QWindow_EventDefault(void* ptr, void* ev)
{
		return static_cast<QWindow*>(ptr)->QWindow::event(static_cast<QEvent*>(ev));
}

long long QWindow_Flags(void* ptr)
{
		return static_cast<QWindow*>(ptr)->flags();
}

void QWindow_FocusInEvent(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->focusInEvent(static_cast<QFocusEvent*>(ev));
}

void QWindow_FocusInEventDefault(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->QWindow::focusInEvent(static_cast<QFocusEvent*>(ev));
}

void QWindow_FocusOutEvent(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void QWindow_FocusOutEventDefault(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->QWindow::focusOutEvent(static_cast<QFocusEvent*>(ev));
}

void* QWindow_Geometry(void* ptr)
{
		return ({ QRect tmpValue = static_cast<QWindow*>(ptr)->geometry(); new QRect(tmpValue.x(), tmpValue.y(), tmpValue.width(), tmpValue.height()); });
}

int QWindow_Height(void* ptr)
{
		return static_cast<QWindow*>(ptr)->height();
}

void QWindow_ConnectHeightChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::heightChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_HeightChanged), static_cast<Qt::ConnectionType>(t));
}

void QWindow_DisconnectHeightChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::heightChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_HeightChanged));
}

void QWindow_HeightChanged(void* ptr, int arg)
{
		static_cast<QWindow*>(ptr)->heightChanged(arg);
}

void QWindow_Hide(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "hide");
}

void QWindow_HideDefault(void* ptr)
{
		static_cast<QWindow*>(ptr)->QWindow::hide();
}

void* QWindow_Icon(void* ptr)
{
		return new QIcon(static_cast<QWindow*>(ptr)->icon());
}

void QWindow_KeyPressEvent(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->keyPressEvent(static_cast<QKeyEvent*>(ev));
}

void QWindow_KeyPressEventDefault(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->QWindow::keyPressEvent(static_cast<QKeyEvent*>(ev));
}

void QWindow_KeyReleaseEvent(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->keyReleaseEvent(static_cast<QKeyEvent*>(ev));
}

void QWindow_KeyReleaseEventDefault(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->QWindow::keyReleaseEvent(static_cast<QKeyEvent*>(ev));
}

void* QWindow_Mask(void* ptr)
{
		return new QRegion(static_cast<QWindow*>(ptr)->mask());
}

void* QWindow_MinimumSize(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QWindow*>(ptr)->minimumSize(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void QWindow_MouseDoubleClickEvent(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->mouseDoubleClickEvent(static_cast<QMouseEvent*>(ev));
}

void QWindow_MouseDoubleClickEventDefault(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->QWindow::mouseDoubleClickEvent(static_cast<QMouseEvent*>(ev));
}

void QWindow_MouseMoveEvent(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->mouseMoveEvent(static_cast<QMouseEvent*>(ev));
}

void QWindow_MouseMoveEventDefault(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->QWindow::mouseMoveEvent(static_cast<QMouseEvent*>(ev));
}

void QWindow_MousePressEvent(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->mousePressEvent(static_cast<QMouseEvent*>(ev));
}

void QWindow_MousePressEventDefault(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->QWindow::mousePressEvent(static_cast<QMouseEvent*>(ev));
}

void QWindow_MouseReleaseEvent(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->mouseReleaseEvent(static_cast<QMouseEvent*>(ev));
}

void QWindow_MouseReleaseEventDefault(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->QWindow::mouseReleaseEvent(static_cast<QMouseEvent*>(ev));
}

void QWindow_MoveEvent(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->moveEvent(static_cast<QMoveEvent*>(ev));
}

void QWindow_MoveEventDefault(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->QWindow::moveEvent(static_cast<QMoveEvent*>(ev));
}

double QWindow_Opacity(void* ptr)
{
		return static_cast<QWindow*>(ptr)->opacity();
}

void QWindow_ConnectOpacityChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(qreal)>(&QWindow::opacityChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(qreal)>(&MyQWindow::Signal_OpacityChanged), static_cast<Qt::ConnectionType>(t));
}

void QWindow_DisconnectOpacityChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(qreal)>(&QWindow::opacityChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(qreal)>(&MyQWindow::Signal_OpacityChanged));
}

void QWindow_OpacityChanged(void* ptr, double opacity)
{
		static_cast<QWindow*>(ptr)->opacityChanged(opacity);
}

void* QWindow_Parent(void* ptr, long long mode)
{
		return static_cast<QWindow*>(ptr)->parent(static_cast<QWindow::AncestorMode>(mode));
}

void* QWindow_Parent2(void* ptr)
{
		return static_cast<QWindow*>(ptr)->parent();
}

void QWindow_Resize(void* ptr, void* newSize)
{
		static_cast<QWindow*>(ptr)->resize(*static_cast<QSize*>(newSize));
}

void QWindow_Resize2(void* ptr, int w, int h)
{
		static_cast<QWindow*>(ptr)->resize(w, h);
}

void QWindow_SetHeight(void* ptr, int arg)
{
		QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "setHeight", Q_ARG(int, arg));
}

void QWindow_SetHeightDefault(void* ptr, int arg)
{
		static_cast<QWindow*>(ptr)->QWindow::setHeight(arg);
}

void QWindow_SetIcon(void* ptr, void* icon)
{
		static_cast<QWindow*>(ptr)->setIcon(*static_cast<QIcon*>(icon));
}

void QWindow_SetMinimumSize(void* ptr, void* size)
{
		static_cast<QWindow*>(ptr)->setMinimumSize(*static_cast<QSize*>(size));
}

void QWindow_Show(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "show");
}

void QWindow_ShowDefault(void* ptr)
{
		static_cast<QWindow*>(ptr)->QWindow::show();
}

void* QWindow_Size(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QWindow*>(ptr)->size(); new QSize(tmpValue.width(), tmpValue.height()); });
}

void* QWindow_SizeDefault(void* ptr)
{
		return ({ QSize tmpValue = static_cast<QWindow*>(ptr)->QWindow::size(); new QSize(tmpValue.width(), tmpValue.height()); });
}

struct QtGui_PackedString QWindow_Title(void* ptr)
{
		return ({ QByteArray t3f590b = static_cast<QWindow*>(ptr)->title().toUtf8(); QtGui_PackedString { const_cast<char*>(t3f590b.prepend("WHITESPACE").constData()+10), t3f590b.size()-10 }; });
}

void QWindow_TouchEvent(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->touchEvent(static_cast<QTouchEvent*>(ev));
}

void QWindow_TouchEventDefault(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->QWindow::touchEvent(static_cast<QTouchEvent*>(ev));
}

long long QWindow_Type(void* ptr)
{
		return static_cast<QWindow*>(ptr)->type();
}

void QWindow_ConnectVisibleChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(bool)>(&QWindow::visibleChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(bool)>(&MyQWindow::Signal_VisibleChanged), static_cast<Qt::ConnectionType>(t));
}

void QWindow_DisconnectVisibleChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(bool)>(&QWindow::visibleChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(bool)>(&MyQWindow::Signal_VisibleChanged));
}

void QWindow_VisibleChanged(void* ptr, char arg)
{
		static_cast<QWindow*>(ptr)->visibleChanged(arg != 0);
}

void QWindow_WheelEvent(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->wheelEvent(static_cast<QWheelEvent*>(ev));
}

void QWindow_WheelEventDefault(void* ptr, void* ev)
{
		static_cast<QWindow*>(ptr)->QWindow::wheelEvent(static_cast<QWheelEvent*>(ev));
}

int QWindow_Width(void* ptr)
{
		return static_cast<QWindow*>(ptr)->width();
}

void QWindow_ConnectWidthChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::widthChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_WidthChanged), static_cast<Qt::ConnectionType>(t));
}

void QWindow_DisconnectWidthChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::widthChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_WidthChanged));
}

void QWindow_WidthChanged(void* ptr, int arg)
{
		static_cast<QWindow*>(ptr)->widthChanged(arg);
}

int QWindow_X(void* ptr)
{
		return static_cast<QWindow*>(ptr)->x();
}

void QWindow_ConnectXChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::xChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_XChanged), static_cast<Qt::ConnectionType>(t));
}

void QWindow_DisconnectXChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::xChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_XChanged));
}

void QWindow_XChanged(void* ptr, int arg)
{
		static_cast<QWindow*>(ptr)->xChanged(arg);
}

int QWindow_Y(void* ptr)
{
		return static_cast<QWindow*>(ptr)->y();
}

void QWindow_ConnectYChanged(void* ptr, long long t)
{
	QObject::connect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::yChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_YChanged), static_cast<Qt::ConnectionType>(t));
}

void QWindow_DisconnectYChanged(void* ptr)
{
	QObject::disconnect(static_cast<QWindow*>(ptr), static_cast<void (QWindow::*)(int)>(&QWindow::yChanged), static_cast<MyQWindow*>(ptr), static_cast<void (MyQWindow::*)(int)>(&MyQWindow::Signal_YChanged));
}

void QWindow_YChanged(void* ptr, int arg)
{
		static_cast<QWindow*>(ptr)->yChanged(arg);
}

void QWindow_DestroyQWindow(void* ptr)
{
	static_cast<QWindow*>(ptr)->~QWindow();
}

void QWindow_DestroyQWindowDefault(void* ptr)
{
	Q_UNUSED(ptr);

}

void* QWindow___children_atList(void* ptr, int i)
{
		return ({QObject * tmp = static_cast<QList<QObject *>*>(ptr)->at(i); if (i == static_cast<QList<QObject *>*>(ptr)->size()-1) { static_cast<QList<QObject *>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWindow___children_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject *>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QWindow___children_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject *>();
}

void* QWindow___dynamicPropertyNames_atList(void* ptr, int i)
{
		return new QByteArray(({QByteArray tmp = static_cast<QList<QByteArray>*>(ptr)->at(i); if (i == static_cast<QList<QByteArray>*>(ptr)->size()-1) { static_cast<QList<QByteArray>*>(ptr)->~QList(); free(ptr); }; tmp; }));
}

void QWindow___dynamicPropertyNames_setList(void* ptr, void* i)
{
		static_cast<QList<QByteArray>*>(ptr)->append(*static_cast<QByteArray*>(i));
}

void* QWindow___dynamicPropertyNames_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QByteArray>();
}

void* QWindow___findChildren_atList(void* ptr, int i)
{
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWindow___findChildren_setList(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QWindow___findChildren_newList(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>();
}

void* QWindow___findChildren_atList3(void* ptr, int i)
{
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWindow___findChildren_setList3(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QWindow___findChildren_newList3(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>();
}

void* QWindow___qFindChildren_atList2(void* ptr, int i)
{
		return ({QObject* tmp = static_cast<QList<QObject*>*>(ptr)->at(i); if (i == static_cast<QList<QObject*>*>(ptr)->size()-1) { static_cast<QList<QObject*>*>(ptr)->~QList(); free(ptr); }; tmp; });
}

void QWindow___qFindChildren_setList2(void* ptr, void* i)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(i))) {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QWindow*>(i));
	} else {
		static_cast<QList<QObject*>*>(ptr)->append(static_cast<QObject*>(i));
	}
}

void* QWindow___qFindChildren_newList2(void* ptr)
{
	Q_UNUSED(ptr);
		return new QList<QObject*>();
}

void QWindow_ChildEvent(void* ptr, void* event)
{
		static_cast<QWindow*>(ptr)->childEvent(static_cast<QChildEvent*>(event));
}

void QWindow_ChildEventDefault(void* ptr, void* event)
{
		static_cast<QWindow*>(ptr)->QWindow::childEvent(static_cast<QChildEvent*>(event));
}

void QWindow_ConnectNotify(void* ptr, void* sign)
{
		static_cast<QWindow*>(ptr)->connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWindow_ConnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWindow*>(ptr)->QWindow::connectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWindow_CustomEvent(void* ptr, void* event)
{
		static_cast<QWindow*>(ptr)->customEvent(static_cast<QEvent*>(event));
}

void QWindow_CustomEventDefault(void* ptr, void* event)
{
		static_cast<QWindow*>(ptr)->QWindow::customEvent(static_cast<QEvent*>(event));
}

void QWindow_DeleteLater(void* ptr)
{
		QMetaObject::invokeMethod(static_cast<QWindow*>(ptr), "deleteLater");
}

void QWindow_DeleteLaterDefault(void* ptr)
{
		static_cast<QWindow*>(ptr)->QWindow::deleteLater();
}

void QWindow_DisconnectNotify(void* ptr, void* sign)
{
		static_cast<QWindow*>(ptr)->disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

void QWindow_DisconnectNotifyDefault(void* ptr, void* sign)
{
		static_cast<QWindow*>(ptr)->QWindow::disconnectNotify(*static_cast<QMetaMethod*>(sign));
}

char QWindow_EventFilter(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(watched))) {
		return static_cast<QWindow*>(ptr)->eventFilter(static_cast<QWindow*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QWindow*>(ptr)->eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

char QWindow_EventFilterDefault(void* ptr, void* watched, void* event)
{
	if (dynamic_cast<QWindow*>(static_cast<QObject*>(watched))) {
		return static_cast<QWindow*>(ptr)->QWindow::eventFilter(static_cast<QWindow*>(watched), static_cast<QEvent*>(event));
	} else {
		return static_cast<QWindow*>(ptr)->QWindow::eventFilter(static_cast<QObject*>(watched), static_cast<QEvent*>(event));
	}
}

void QWindow_TimerEvent(void* ptr, void* event)
{
		static_cast<QWindow*>(ptr)->timerEvent(static_cast<QTimerEvent*>(event));
}

void QWindow_TimerEventDefault(void* ptr, void* event)
{
		static_cast<QWindow*>(ptr)->QWindow::timerEvent(static_cast<QTimerEvent*>(event));
}

void* QWindow_Format(void* ptr)
{
		return new QSurfaceFormat(static_cast<QWindow*>(ptr)->format());
}

void* QWindow_FormatDefault(void* ptr)
{
		return new QSurfaceFormat(static_cast<QWindow*>(ptr)->QWindow::format());
}

long long QWindow_SurfaceType(void* ptr)
{
		return static_cast<QWindow*>(ptr)->surfaceType();
}

long long QWindow_SurfaceTypeDefault(void* ptr)
{
		return static_cast<QWindow*>(ptr)->QWindow::surfaceType();
}

