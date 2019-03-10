# -*- coding: utf-8 -*-

# Resource object code
#
# Created: sáb. mar. 9 22:03:24 2019
#      by: The Resource Compiler for PySide2 (Qt v5.12.1)
#
# WARNING! All changes made in this file will be lost!

from PySide2 import QtCore

qt_resource_data = b"\
\x00\x00\x00\xed\
;\
 This file can b\
e edited to chan\
ge the style of \
the application\x0a\
; Read \x22Qt Quick\
 Controls 2 Conf\
iguration File\x22 \
for details:\x0a; h\
ttp://doc.qt.io/\
qt-5/qtquickcont\
rols2-configurat\
ion.html\x0a\x0a[Contr\
ols]\x0aStyle=Mater\
ial\x0a\x0a[Material]\x0a\
Accent=Blue\x0a\
"

qt_resource_name = b"\
\x00\x15\
\x08\x1e\x16f\
\x00q\
\x00t\x00q\x00u\x00i\x00c\x00k\x00c\x00o\x00n\x00t\x00r\x00o\x00l\x00s\x002\x00.\
\x00c\x00o\x00n\x00f\
"

qt_resource_struct = b"\
\x00\x00\x00\x00\x00\x02\x00\x00\x00\x01\x00\x00\x00\x01\
\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\
"

def qInitResources():
    QtCore.qRegisterResourceData(0x01, qt_resource_struct, qt_resource_name, qt_resource_data)

def qCleanupResources():
    QtCore.qUnregisterResourceData(0x01, qt_resource_struct, qt_resource_name, qt_resource_data)

qInitResources()
