#include <QtCore/QStandardPaths>
#include <QtCore/QLockFile>
#include <QtCore/QLocale>
#include <QtCore/QTranslator>
#include <QtGui/QGuiApplication>
#include <QtGui/QIcon>
#include <QtGui/QCursor>
#include <QtGui/QFontDatabase>
#include <QtGui/QFont>
#include <QtQml/QQmlApplicationEngine>

int main(int argc, char *argv[])
{
    QCoreApplication::setAttribute(Qt::AA_EnableHighDpiScaling);

    QGuiApplication app(argc, argv);

    app.setOrganizationName("Simelo.Tech");
    app.setOrganizationDomain("simelo.tech");
    app.setApplicationName("FiberCryptoWallet");
    app.setApplicationVersion("0.27.0");
    app.setWindowIcon(QIcon(":/images/icons/appIcon/appIcon.png"));
    
    // Allow a single running instance
    QString lockFilePath = QStandardPaths::writableLocation(QStandardPaths::TempLocation) + QString("/m%1.lock").arg(app.applicationName());
    QLockFile lockFile(lockFilePath);
    if (!lockFile.tryLock()) {
        // There's already a running instance. Tell user to close it (a new QQmlApplicationEngine?)
        return 1;
    }

    app.setOverrideCursor(QCursor(Qt::BusyCursor));

    // Font
#ifdef Q_OS_WINDOWS
    // Fix no-so-cool font rendering on Windows
    QFont applicationFont = app.font();
    applicationFont.setFamily("Helvetica");
    app.setFont(applicationFont);
#endif
    QFontDatabase::addApplicationFont(":/fonts/code-new-roman/code-new-roman.otf");
    QFontDatabase::addApplicationFont(":/fonts/hemi-head/hemi-head.ttf");

    // Internationalization
    QLocale::setDefault(QLocale::English);
    QTranslator qtTranslator, appTranslator;
    qtTranslator.load(QLocale::English, QStringLiteral("qt"), QStringLiteral("_"), ":/translations");
    appTranslator.load(QLocale::English, QStringLiteral("fibercryptowallet"), QStringLiteral("_"), ":/translations");
    app.installTranslator(&qtTranslator);
    app.installTranslator(&appTranslator);

    // Load the UI
    QQmlApplicationEngine engine;
    const QUrl url(QStringLiteral("qrc:/src/ui/splash.qml"));
    QObject::connect(&engine, &QQmlApplicationEngine::objectCreated,
                     &app, [url](QObject *obj, const QUrl &objUrl) {
        if (!obj && url == objUrl)
            QCoreApplication::exit(-1);
    }, Qt::QueuedConnection);
    engine.load(url);

#ifdef Q_OS_ANDROID
    // Fade the splash screen on Android
    // androidextras module was removed in Qt 6. Where is this function now?
    // QtAndroid::hideSplashScreen(50);
#endif

    app.restoreOverrideCursor(); // do this in splash.qml

    return app.exec();
}