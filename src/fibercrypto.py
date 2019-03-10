# Python imports
import sys
from os.path import abspath, dirname, join

# PySide2 imports
from PySide2.QtGui import QGuiApplication
from PySide2.QtQml import QQmlApplicationEngine

# My imports
## Resources
import images_rc
import qt_rc

if __name__ == '__main__':
    app = QGuiApplication(sys.argv)
    engine = QQmlApplicationEngine()

    # Once you have the path of the current directory,
    # add the name of the QML file, to load it.
    qmlFile = join(current_dir, 'main.qml')
    engine.load(abspath(qmlFile))

    if not engine.rootObjects():
        sys.exit(-1)

    sys.exit(app.exec_())