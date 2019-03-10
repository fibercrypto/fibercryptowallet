# Python imports
import sys
from os.path import abspath, dirname, join

## Modify the sys.path environment variable
current_dir = dirname(__file__)
resource_dir = join(current_dir, "media")

ui_dir = join(current_dir, "ui")
images_dir = join(resource_dir, "images")
config_dir = join(resource_dir, "qt")

sys.path.append(ui_dir)
sys.path.append(images_dir)
sys.path.append(config_dir)

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

    # Once you have the path of the UI directory,
    # add the name of the QML file, to load it.
    qmlFile = join(ui_dir, 'main.qml')
    engine.load(abspath(qmlFile))

    if not engine.rootObjects():
        sys.exit(-1)

    sys.exit(app.exec_())