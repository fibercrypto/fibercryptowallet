# -*- coding: utf-8 -*-

# Resource object code
#
# Created: mié. mar. 27 17:23:16 2019
#      by: The Resource Compiler for PySide2 (Qt v5.12.1)
#
# WARNING! All changes made in this file will be lost!

from PySide2 import QtCore

qt_resource_data = b"\
\x00\x00\x07R\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.L\
ayouts 1.12\x0a\x0aPag\
e {\x0a    id: root\
\x0a\x0a    ColumnLayo\
ut {\x0a        id:\
 columnLayoutRoo\
t\x0a        anchor\
s.fill: parent\x0a \
       anchors.l\
eftMargin: 10\x0a  \
      anchors.ri\
ghtMargin: 10\x0a  \
      spacing: 2\
0\x0a\x0a        Colum\
nLayout {\x0a      \
      id: column\
LayoutSendFrom\x0a\x0a\
            Layo\
ut.alignment: Qt\
.AlignTop\x0a      \
      Label {\x0a  \
              te\
xt: qsTr(\x22Send f\
rom\x22)\x0a          \
  }\x0a            \
ComboBox {\x0a     \
           id: c\
omboBoxWalletsSe\
ndFrom\x0a         \
       Layout.fi\
llWidth: true\x0a  \
          }\x0a    \
    }\x0a\x0a        C\
olumnLayout {\x0a  \
          id: co\
lumnLayoutSendTo\
\x0a\x0a            La\
yout.alignment: \
Qt.AlignTop\x0a    \
        Label {\x0a\
                \
text: qsTr(\x22Send\
 to\x22)\x0a          \
  }\x0a            \
RowLayout {\x0a    \
            Layo\
ut.fillWidth: tr\
ue\x0a             \
   spacing: 8\x0a  \
              Im\
age {\x0a          \
          source\
: \x22qrc:/images/i\
cons/qr.svg\x22\x0a   \
                \
 sourceSize: \x2224\
x24\x22\x0a           \
         Layout.\
bottomMargin: 4\x0a\
                \
}\x0a              \
  TextField {\x0a  \
                \
  id: textFieldW\
alletsSendTo\x0a   \
                \
 font.family: \x22C\
ode New Roman\x22\x0a \
                \
   placeholderTe\
xt: qsTr(\x22Destin\
ation address\x22)\x0a\
                \
    Layout.fillW\
idth: true\x0a     \
               L\
ayout.topMargin:\
 -5\x0a            \
    }\x0a          \
  }\x0a        }\x0a\x0a \
       ColumnLay\
out {\x0a          \
  id: columnLayo\
utAmount\x0a\x0a      \
      Layout.ali\
gnment: Qt.Align\
Top\x0a            \
Label {\x0a        \
        text: qs\
Tr(\x22Amount\x22)\x0a   \
         }\x0a     \
       TextField\
 {\x0a             \
   id: textField\
Amount\x0a         \
       placehold\
erText: qsTr(\x22Am\
ount to send\x22)\x0a \
               L\
ayout.fillWidth:\
 true\x0a          \
      Layout.top\
Margin: -10\x0a    \
            vali\
dator: DoubleVal\
idator {\x0a       \
             not\
ation: DoubleVal\
idator.StandardN\
otation\x0a        \
        }\x0a      \
      }\x0a        \
}\x0a    } // Colum\
nLayout (root)\x0a}\
\x0a\
\x00\x00\x15O\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aPage {\x0a    id\
: root\x0a\x0a    Colu\
mnLayout {\x0a     \
   id: columnLay\
outRoot\x0a        \
anchors.fill: pa\
rent\x0a        anc\
hors.leftMargin:\
 10\x0a        anch\
ors.rightMargin:\
 10\x0a        spac\
ing: 20\x0a\x0a       \
 ColumnLayout {\x0a\
            id: \
columnLayoutAddr\
essesSendFrom\x0a\x0a \
           Layou\
t.alignment: Qt.\
AlignTop\x0a       \
     RowLayout {\
\x0a               \
 Label {\x0a       \
             tex\
t: qsTr(\x22Address\
\x22)\x0a             \
   }\x0a           \
     ToolButton \
{\x0a              \
      id: toolBu\
ttonAddressPopup\
Help\x0a           \
         icon.so\
urce: \x22qrc:/imag\
es/icons/help.sv\
g\x22\x0a             \
       icon.colo\
r: Material.colo\
r(Material.Grey)\
\x0a               \
 }\x0a            }\
\x0a            Com\
boBox {\x0a        \
        id: comb\
oBoxWalletsAddre\
ssesSendFrom\x0a   \
             Lay\
out.fillWidth: t\
rue\x0a            \
    Layout.topMa\
rgin: -12\x0a      \
      }\x0a        \
}\x0a\x0a        Colum\
nLayout {\x0a      \
      id: column\
LayoutDestinatio\
ns\x0a\x0a            \
Layout.alignment\
: Qt.AlignTop\x0a  \
          RowLay\
out {\x0a          \
      Label {\x0a  \
                \
  text: qsTr(\x22De\
stinations\x22)\x0a   \
             }\x0a \
               T\
oolButton {\x0a    \
                \
id: toolButtonDe\
stinationPopupHe\
lp\x0a             \
       icon.sour\
ce: \x22qrc:/images\
/icons/help.svg\x22\
\x0a               \
     icon.color:\
 Material.color(\
Material.Grey)\x0a \
               }\
\x0a            }\x0a \
           ListV\
iew {\x0a          \
      id: listVi\
ewDestinations\x0a\x0a\
                \
property real de\
legateHeight: 47\
\x0a\x0a              \
  Layout.fillWid\
th: true\x0a       \
         Layout.\
topMargin: -16\x0a \
               i\
mplicitHeight: c\
ount * delegateH\
eight\x0a\x0a         \
       Behavior \
on implicitHeigh\
t { NumberAnimat\
ion { duration: \
250; easing.type\
: Easing.OutQuin\
t } }\x0a\x0a         \
       interacti\
ve: false\x0a\x0a     \
           model\
: listModelDesti\
nations\x0a        \
        clip: tr\
ue\x0a\x0a            \
    delegate: De\
stinationListDel\
egate {\x0a        \
            id: \
delegateDestinat\
ion\x0a            \
        width: l\
istViewDestinati\
ons.width\x0a      \
              im\
plicitHeight: Li\
stView.view.dele\
gateHeight\x0a     \
           }\x0a   \
         } // Li\
stView\x0a        }\
 // ColumnLayout\
 (destinations)\x0a\
\x0a        ColumnL\
ayout {\x0a        \
    id: columnLa\
youtCustomChange\
Address\x0a\x0a       \
     Layout.alig\
nment: Qt.AlignT\
op\x0a            R\
owLayout {\x0a     \
           Label\
 {\x0a             \
       text: qsT\
r(\x22Custom change\
 address\x22)\x0a     \
           }\x0a   \
             Too\
lButton {\x0a      \
              id\
: toolButtonCust\
omChangeAddressP\
opupHelp\x0a       \
             ico\
n.source: \x22qrc:/\
images/icons/hel\
p.svg\x22\x0a         \
           icon.\
color: Material.\
color(Material.G\
rey)\x0a           \
     }\x0a         \
       Button {\x0a\
                \
    id: buttonSe\
lectCustomChange\
Address\x0a        \
            text\
: qsTr(\x22Select\x22)\
\x0a               \
     flat: true\x0a\
                \
    highlighted:\
 true\x0a          \
      }\x0a        \
    }\x0a          \
  TextField {\x0a  \
              id\
: textFieldCusto\
mChangeAddress\x0a \
               p\
laceholderText: \
qsTr(\x22Address to\
 receive change\x22\
)\x0a              \
  Layout.fillWid\
th: true\x0a       \
         Layout.\
topMargin: -16\x0a \
           }\x0a   \
     }\x0a\x0a        \
ColumnLayout {\x0a \
           id: c\
olumnLayoutAutom\
aticCoinHoursAll\
ocation\x0a\x0a       \
     Layout.alig\
nment: Qt.AlignT\
op\x0a            R\
owLayout {\x0a     \
           Check\
Box {\x0a          \
          id: ch\
eckBoxAutomaticC\
oinHoursAllocati\
on\x0a             \
       text: qsT\
r(\x22Automatic coi\
n hours allocati\
on\x22)\x0a           \
         checked\
: true\x0a         \
       }\x0a       \
         Button \
{\x0a              \
      id: button\
OptionsAutomatic\
CoinHoursAllocat\
ion\x0a            \
        text: qs\
Tr(\x22Options\x22)\x0a  \
                \
  flat: true\x0a   \
                \
 highlighted: tr\
ue\x0a             \
       enabled: \
checkBoxAutomati\
cCoinHoursAlloca\
tion.checked\x0a\x0a  \
                \
  onClicked: {\x0a \
                \
       toolTipAu\
tomaticCoinHours\
Allocation.show(\
\x22Coin hours shar\
e factor\x22)\x0a     \
               }\
\x0a               \
 }\x0a            }\
\x0a        }\x0a    }\
 // ColumnLayout\
 (root)\x0a\x0a    Lis\
tModel {\x0a       \
 id: listModelDe\
stinations\x0a     \
   ListElement {\
 address: \x22\x22; sk\
y: 0.0; coinHour\
s: 0.0 }\x0a    }\x0a\x0a\
\x0a    ToolTip {\x0a \
       id: toolT\
ipAutomaticCoinH\
oursAllocation\x0a\x0a\
        property\
 alias toolTipHo\
vered: hoverHand\
ler.hovered\x0a\x0a   \
     parent: but\
tonOptionsAutoma\
ticCoinHoursAllo\
cation\x0a        a\
nchors.centerIn:\
 parent\x0a        \
height: sliderCo\
inHoursShareFact\
or.height + rowL\
ayoutSlider.anch\
ors.topMargin + \
10\x0a\x0a        Hove\
rHandler {\x0a     \
       id: hover\
Handler\x0a        \
    margin: 16\x0a \
           onHov\
eredChanged: {\x0a \
               i\
f (!hovered && !\
sliderCoinHoursS\
hareFactor.press\
ed) {\x0a          \
          toolTi\
pAutomaticCoinHo\
ursAllocation.hi\
de()\x0a           \
     }\x0a         \
   }\x0a        }\x0a\x0a\
        RowLayou\
t {\x0a            \
id: rowLayoutSli\
der\x0a            \
anchors.left: pa\
rent.left\x0a      \
      anchors.ri\
ght: parent.righ\
t\x0a            an\
chors.top: paren\
t.top\x0a          \
  anchors.topMar\
gin: 16\x0a        \
    Slider {\x0a   \
             id:\
 sliderCoinHours\
ShareFactor\x0a    \
            Layo\
ut.fillWidth: tr\
ue\x0a\x0a            \
    from: 0.01\x0a \
               t\
o: 1.00\x0a        \
        value: 0\
.50\x0a            \
    stepSize: 0.\
01\x0a\x0a            \
    onPressedCha\
nged: {\x0a        \
            if (\
!hovered && !too\
lTipAutomaticCoi\
nHoursAllocation\
.toolTipHovered)\
 {\x0a             \
           toolT\
ipAutomaticCoinH\
oursAllocation.c\
lose()\x0a         \
           }\x0a   \
             }\x0a \
           } // \
Slider\x0a\x0a        \
    Label {\x0a    \
            text\
: sliderCoinHour\
sShareFactor.val\
ue.toFixed(2)\x0a  \
          }\x0a    \
    }\x0a    }\x0a}\x0a\
\x00\x00\x11}\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aDialog {\x0a    \
id: root\x0a\x0a    pr\
operty string de\
viceName: \x22<NULL\
>\x22\x0a    property \
bool enableBacku\
pWarning: true\x0a \
   property bool\
 enablePINWarnin\
g: true\x0a\x0a    foc\
us: true\x0a    mod\
al: true\x0a    tit\
le: Qt.applicati\
on.name\x0a    stan\
dardButtons: Dia\
log.Abort\x0a    cl\
osePolicy: Dialo\
g.NoAutoClose\x0a\x0a \
   implicitWidth\
: 600\x0a    implic\
itHeight: 550 - \
(enableBackupWar\
ning ^ enablePIN\
Warning ? 100 : \
0) - (!enableBac\
kupWarning && !e\
nablePINWarning \
? 240 : 0)\x0a\x0a    \
ColumnLayout {\x0a \
       anchors.f\
ill: parent\x0a    \
    spacing: 50\x0a\
        RowLayou\
t {\x0a            \
spacing: 50\x0a    \
        Image {\x0a\
                \
id: icon\x0a       \
         source:\
 \x22qrc:/images/ic\
ons/security.svg\
\x22\x0a              \
  sourceSize: \x227\
2x72\x22\x0a          \
      Layout.ali\
gnment: Qt.Align\
Top\x0a            \
}\x0a            Co\
lumnLayout {\x0a   \
             Lab\
el {\x0a           \
         id: msg\
Title\x0a          \
          text: \
qsTr(\x22Secure har\
dware wallet\x22)\x0a \
                \
   font.bold: tr\
ue\x0a             \
       Layout.fi\
llWidth: true\x0a  \
                \
  wrapMode: Text\
.WordWrap\x0a      \
              La\
yout.alignment: \
Qt.AlignTop\x0a    \
            }\x0a  \
              La\
bel {\x0a          \
          id: ms\
g\x0a              \
      text: qsTr\
(\x22Hardware walle\
t detected.<br>T\
he device is ide\
ntified in the \x22\
 +\x0a             \
                \
  \x22wallets list \
as: \x22) + \x22<b><i>\
\x22 + deviceName +\
 \x22</i></b>\x22\x0a    \
                \
Layout.fillWidth\
: true\x0a         \
           wrapM\
ode: Text.WordWr\
ap\x0a             \
       Layout.al\
ignment: Qt.Alig\
nTop\x0a           \
     }\x0a         \
   }\x0a        }\x0a\x0a\
        ColumnLa\
yout {\x0a         \
   spacing: 30\x0a\x0a\
            Colu\
mnLayout {\x0a     \
           visib\
le: enableBackup\
Warning || enabl\
ePINWarning\x0a    \
            Labe\
l {\x0a            \
        id: secu\
rityWarningsTitl\
e\x0a              \
      text: qsTr\
(\x22Security warni\
ngs\x22)\x0a          \
          font.p\
ointSize: 14\x0a   \
                \
 font.bold: true\
\x0a               \
     color: Mate\
rial.color(Mater\
ial.Pink)\x0a      \
              La\
yout.fillWidth: \
true\x0a           \
         wrapMod\
e: Text.WordWrap\
\x0a               \
     Layout.alig\
nment: Qt.AlignT\
op\x0a             \
   }\x0a           \
     Label {\x0a   \
                \
 id: securityWar\
ningBackup\x0a     \
               v\
isible: enableBa\
ckupWarning\x0a    \
                \
text: \x22<b>1)</b>\
 \x22 +\x0a           \
               q\
sTr(\x22You should \
secure the hardw\
are seed or you \
could lose \x22 +\x0a \
                \
              \x22a\
ccess to the fun\
ds in case of pr\
oblems. \x22 +\x0a    \
                \
           \x22To d\
o this, select <\
b><i>Create a ba\
ckup</i></b>.\x22)\x0a\
                \
    color: Mater\
ial.color(Materi\
al.Pink)\x0a       \
             Lay\
out.fillWidth: t\
rue\x0a            \
        wrapMode\
: Text.WordWrap\x0a\
                \
    Layout.align\
ment: Qt.AlignTo\
p\x0a              \
  }\x0a            \
    Label {\x0a    \
                \
id: securityWarn\
ingPIN\x0a         \
           visib\
le: enablePINWar\
ning\x0a           \
         text: \x22\
<b>\x22 + (enableBa\
ckupWarning ? 2 \
: 1) + \x22)</b> \x22 \
+\x0a              \
            qsTr\
(\x22The connected \
hardware wallet \
does not have a \
PIN. \x22 +\x0a       \
                \
        \x22The PIN\
 code protects t\
he hardware wall\
et in case of \x22 \
+\x0a              \
                \
 \x22loss, theft an\
d hacks. \x22 +\x0a   \
                \
            \x22To \
do this, select \
<b><i>Create PIN\
 code</i></b>.\x22)\
\x0a               \
     color: Mate\
rial.color(Mater\
ial.Pink)\x0a      \
              La\
yout.fillWidth: \
true\x0a           \
         wrapMod\
e: Text.WordWrap\
\x0a               \
     Layout.alig\
nment: Qt.AlignT\
op\x0a             \
   }\x0a           \
 }\x0a\x0a            \
ColumnLayout {\x0a \
               L\
abel {\x0a         \
           id: o\
ptions\x0a         \
           text:\
 qsTr(\x22Options:\x22\
)\x0a              \
      font.bold:\
 true\x0a          \
          Layout\
.fillWidth: true\
\x0a               \
 }\x0a             \
   ItemDelegate \
{\x0a              \
      id: button\
CreateBackup\x0a   \
                \
 visible: enable\
BackupWarning\x0a  \
                \
  text: qsTr(\x22Cr\
eate a backup\x22)\x0a\
                \
    Layout.fillW\
idth: true\x0a     \
           }\x0a   \
             Ite\
mDelegate {\x0a    \
                \
id: buttonCreate\
PIN\x0a            \
        visible:\
 enablePINWarnin\
g\x0a              \
      text: qsTr\
(\x22Create PIN cod\
e\x22)\x0a            \
        Layout.f\
illWidth: true\x0a \
               }\
\x0a               \
 ItemDelegate {\x0a\
                \
    id: buttonWi\
peDevice\x0a       \
             tex\
t: qsTr(\x22Wipe de\
vice\x22)\x0a         \
           Layou\
t.fillWidth: tru\
e\x0a              \
      Material.f\
oreground: Mater\
ial.Pink\x0a       \
         }\x0a     \
       }\x0a\x0a      \
  }\x0a    }\x0a}\x0a\
\x00\x00\x13\xd0\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aPage {\x0a    id\
: root\x0a\x0a    Grou\
pBox {\x0a        a\
nchors.fill: par\
ent\x0a        anch\
ors.margins: 20\x0a\
        clip: tr\
ue\x0a\x0a        labe\
l: RowLayout {\x0a \
           Switc\
hDelegate {\x0a    \
            id: \
switchFilters\x0a  \
              te\
xt: qsTr(\x22Filter\
s\x22)\x0a            \
}\x0a            Bu\
tton {\x0a         \
       id: butto\
nFilters\x0a       \
         flat: t\
rue\x0a            \
    enabled: swi\
tchFilters.check\
ed\x0a             \
   highlighted: \
true\x0a           \
     text: qsTr(\
\x22Select filters\x22\
)\x0a\x0a             \
   onClicked: {\x0a\
                \
    toolTipFilte\
rs.open()\x0a      \
          }\x0a    \
        }\x0a      \
  } // RowLayout\
 (GroupBox label\
)\x0a\x0a        Scrol\
lView {\x0a        \
    anchors.fill\
: parent\x0a       \
     clip: true\x0a\
            List\
View {\x0a         \
       id: listT\
ransactions\x0a\x0a   \
             mod\
el: modelTransac\
tions\x0a          \
      delegate: \
HistoryListDeleg\
ate {\x0a          \
          onClic\
ked: {\x0a         \
               d\
ialogTransaction\
Details.open()\x0a \
                \
       listTrans\
actions.currentI\
ndex = index\x0a   \
                \
 }\x0a             \
   }\x0a           \
 }\x0a        }\x0a   \
 } // GroupBox\x0a\x0a\
\x0a    ToolTip {\x0a \
       id: toolT\
ipFilters\x0a\x0a     \
   anchors.cente\
rIn: Overlay.ove\
rlay\x0a        hei\
ght: 16 + (filte\
r.count > 5 ? fi\
lter.delegateHei\
ght * 5 : filter\
.delegateHeight \
* filter.count)\x0a\
        clip: tr\
ue\x0a        modal\
: true\x0a        c\
losePolicy: Popu\
p.CloseOnEscape \
| Popup.CloseOnP\
ressOutside\x0a\x0a   \
     ColumnLayou\
t {\x0a            \
anchors.fill: pa\
rent\x0a\x0a          \
  Label {\x0a      \
          font.p\
ointSize: Qt.app\
lication.font.po\
intSize * 1.25\x0a \
               f\
ont.bold: true\x0a \
               t\
ext: qsTr(\x22Avail\
able filters\x22)\x0a \
               L\
ayout.alignment:\
 Qt.AlignHCenter\
 | Qt.AlignTop\x0a \
           }\x0a\x0a  \
          Scroll\
View {\x0a         \
       Layout.al\
ignment: Qt.Alig\
nCenter\x0a        \
        Layout.f\
illWidth: true\x0a \
               L\
ayout.fillHeight\
: true\x0a         \
       clip: tru\
e\x0a              \
  contentHeight:\
 filter.implicit\
Height\x0a         \
       HistoryFi\
lterList {\x0a     \
               i\
d: filter\x0a      \
              in\
teractive: false\
\x0a               \
 }\x0a            }\
\x0a        } // Co\
lumnLayout\x0a    }\
 // ToolTip\x0a\x0a   \
 DialogTransacti\
onDetails {\x0a    \
    id: dialogTr\
ansactionDetails\
\x0a        anchors\
.centerIn: Overl\
ay.overlay\x0a\x0a    \
    modal: true\x0a\
        focus: t\
rue\x0a        stan\
dardButtons: Dia\
log.Ok\x0a\x0a        \
date: listTransa\
ctions.currentIt\
em.modelDate\x0a   \
     status: lis\
tTransactions.cu\
rrentItem.modelS\
tatus\x0a        ty\
pe: listTransact\
ions.currentItem\
.modelType\x0a     \
   amount: listT\
ransactions.curr\
entItem.modelAmo\
unt\x0a        hour\
sReceived: listT\
ransactions.curr\
entItem.modelHou\
rsReceived\x0a     \
   hoursBurned: \
listTransactions\
.currentItem.mod\
elHoursBurned\x0a  \
      transactio\
nID: listTransac\
tions.currentIte\
m.modelTransacti\
onID\x0a    }\x0a\x0a    \
// Roles: type, \
date, amount, se\
ntAddress, recei\
vedAddress\x0a    /\
/ Use listModel.\
append( { \x22type\x22\
: value, \x22date\x22:\
 value, \x22sentAdd\
ress\x22: value, \x22r\
eceivedAddress\x22:\
 value } )\x0a    /\
/ Or implement t\
he model in the \
backend (a more \
recommendable ap\
proach)\x0a    List\
Model { // EXAMP\
LE\x0a        id: m\
odelTransactions\
\x0a        ListEle\
ment { date: \x2220\
19-03-02 10:19\x22;\
 status: Transac\
tionDetails.Stat\
us.Confirmed; ty\
pe: TransactionD\
etails.Type.Send\
;    amount: 9; \
    hoursReceive\
d: 65;   hoursBu\
rned: 6287;   tr\
ansactionID: \x22ks\
uwet837iwetr27ie\
tr\x22; sentAddress\
: \x22734irwepaweyg\
tawieta783cwyc\x22;\
 receivedAddress\
: \x22uewyru63u3789\
jfrgpaldcxz4f6\x22 \
}\x0a        ListEl\
ement { date: \x222\
019-03-02 10:44\x22\
; status: Transa\
ctionDetails.Sta\
tus.Confirmed; t\
ype: Transaction\
Details.Type.Rec\
eive; amount: 10\
000; hoursReceiv\
ed: 9421; hoursB\
urned: 218489; t\
ransactionID: \x22h\
dk84iesurys29q8i\
khf\x22; sentAddres\
s: \x22nvslkjoid98w\
emxsoqzmap50vah\x22\
; receivedAddres\
s: \x22ps73729sskso\
f972jwkf83owhfi\x22\
 }\x0a        ListE\
lement { date: \x22\
2019-03-02 15:01\
\x22; status: Trans\
actionDetails.St\
atus.Confirmed; \
type: Transactio\
nDetails.Type.Se\
nd;    amount: 1\
25;   hoursRecei\
ved: 4203; hours\
Burned: 7950;   \
transactionID: \x22\
p9438938qksaxjhs\
dkuq\x22; sentAddre\
ss: \x22qwe98uimhfd\
kfu23y9hewi8qd02\
\x22; receivedAddre\
ss: \x22p2837djkdjv\
kauyfiawbw83h48f\
\x22 }\x0a        List\
Element { date: \
\x222019-03-03 08:1\
3\x22; status: Tran\
sactionDetails.S\
tatus.Confirmed;\
 type: Transacti\
onDetails.Type.S\
end;    amount: \
100;   hoursRece\
ived: 3877; hour\
sBurned: 6911;  \
 transactionID: \
\x2273t4dunksxr8w7z\
uwe13\x22; sentAddr\
ess: \x2289niweuq82\
zqur37izuwsklpar\
u\x22; receivedAddr\
ess: \x220j3726djf9\
f4w9sovgwuw9e4n9\
s\x22 }\x0a        Lis\
tElement { date:\
 \x222019-03-03 12:\
23\x22; status: Tra\
nsactionDetails.\
Status.Confirmed\
; type: Transact\
ionDetails.Type.\
Receive; amount:\
 80;    hoursRec\
eived: 3084; hou\
rsBurned: 4603; \
  transactionID:\
 \x229782nz87e7tque\
tyn83w\x22; sentAdd\
ress: \x22t7ekwduf8\
045ogmcbq63nf9bm\
36\x22; receivedAdd\
ress: \x22aske8jdhb\
smwq9204h49y25jr\
ue\x22 }\x0a        Li\
stElement { date\
: \x222019-03-03 13\
:33\x22; status: Tr\
ansactionDetails\
.Status.Confirme\
d; type: Transac\
tionDetails.Type\
.Receive; amount\
: 60;    hoursRe\
ceived: 2796; ho\
ursBurned: 2347;\
   transactionID\
: \x229127enqiuyrt6\
tyejshd\x22; sentAd\
dress: \x222imdiym9\
w7ji8s29ydk0djwd\
6wj\x22; receivedAd\
dress: \x220gndqy72\
5k9fj3ewdardgbkd\
83l\x22 }\x0a    }\x0a}\x0a\
\x00\x00\x07\x8a\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aDialog {\x0a    \
id: root\x0a\x0a    fo\
cus: visible\x0a   \
 modal: true\x0a   \
 title: Qt.appli\
cation.name\x0a    \
standardButtons:\
 Dialog.Ok\x0a    c\
losePolicy: Dial\
og.NoAutoClose\x0a\x0a\
    implicitWidt\
h: 500\x0a    impli\
citHeight: 300\x0a\x0a\
    Component.on\
Completed: {\x0a   \
     standardBut\
ton(Dialog.Ok).e\
nabled = name.te\
xt !== \x22\x22\x0a    }\x0a\
\x0a    ColumnLayou\
t {\x0a        anch\
ors.fill: parent\
\x0a        spacing\
: 30\x0a\x0a        Ro\
wLayout {\x0a      \
      spacing: 3\
0\x0a\x0a            I\
mage {\x0a         \
       id: icon\x0a\
                \
source: \x22qrc:/im\
ages/icons/check\
.svg\x22\x0a          \
      sourceSize\
: \x2264x64\x22\x0a      \
          Layout\
.alignment: Qt.A\
lignTop\x0a        \
    }\x0a\x0a         \
   ColumnLayout \
{\x0a              \
  Label {\x0a      \
              te\
xt: qsTr(\x22The co\
nnected hardware\
 wallet will be \
added \x22 +\x0a      \
                \
         \x22to the\
 wallets list wi\
th the following\
 name:\x22)\x0a       \
             Lay\
out.fillWidth: t\
rue\x0a            \
        wrapMode\
: Text.WordWrap\x0a\
                \
    Layout.align\
ment: Qt.AlignTo\
p\x0a              \
  }\x0a\x0a           \
     TextField {\
\x0a               \
     id: name\x0a  \
                \
  placeholderTex\
t: qsTr(\x22New har\
dware wallet\x22)\x0a \
                \
   Layout.fillWi\
dth: true\x0a      \
              fo\
cus: root.focus\x0a\
\x0a               \
     onTextChang\
ed: {\x0a          \
              st\
andardButton(Dia\
log.Ok).enabled \
= text !== \x22\x22\x0a  \
                \
  }\x0a            \
    }\x0a          \
  }\x0a        }\x0a\x0a \
       Label {\x0a \
           text:\
 qsTr(\x22Now you c\
an check the bal\
ance and the add\
ress of the \x22 +\x0a\
                \
       \x22hardware\
 wallet even whe\
n it is not conn\
ected. \x22 +\x0a     \
                \
  \x22Additionally,\
 you can change \
the name of the \
wallet \x22 +\x0a     \
                \
  \x22or remove it \
from the wallets\
 list, if you wi\
sh.\x22)\x0a          \
  Layout.fillWid\
th: true\x0a       \
     wrapMode: T\
ext.WordWrap\x0a   \
         Layout.\
alignment: Qt.Al\
ignTop\x0a        }\
\x0a    }\x0a}\x0a\
\x00\x00\x04\x1f\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
\x0aItem {\x0a    id: \
root\x0a\x0a    proper\
ty bool toolPage\
Opened: false\x0a\x0a \
   function open\
BlockchainPage()\
 {\x0a        stack\
View.push(compon\
entBlockchain)\x0a \
       toolPageO\
pened = true\x0a   \
 }\x0a\x0a    StackVie\
w {\x0a        id: \
stackView\x0a      \
  initialItem: c\
omponentCreateLo\
adWallet\x0a       \
 anchors.fill: p\
arent\x0a    }\x0a\x0a   \
 Component {\x0a   \
     id: compone\
ntCreateLoadWall\
et\x0a\x0a        Rect\
angle {\x0a        \
    CreateLoadWa\
llet {\x0a         \
       id: creat\
eLoadWallet\x0a    \
            anch\
ors.centerIn: pa\
rent\x0a\x0a          \
      onWalletCr\
eationRequested:\
 {\x0a             \
       stackView\
.push(componentG\
eneralSwipeView)\
\x0a               \
 }\x0a            }\
\x0a        }\x0a    }\
\x0a\x0a    Component \
{\x0a        id: co\
mponentGeneralSw\
ipeView\x0a\x0a       \
 GeneralSwipeVie\
w {\x0a            \
id: generalSwipe\
View\x0a        }\x0a \
   }\x0a\x0a    Compon\
ent {\x0a        id\
: componentBlock\
chain\x0a\x0a        B\
lockchain {\x0a    \
        id: bloc\
kchain\x0a\x0a        \
    onBackReques\
ted: {\x0a         \
       stackView\
.pop()\x0a         \
       toolPageO\
pened = false\x0a  \
          }\x0a    \
    }\x0a    }\x0a}\x0a\
\x00\x00\x04\xa1\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aDialog {\x0a    \
id: root\x0a\x0a    Co\
lumnLayout {\x0a   \
     anchors.fil\
l: parent\x0a      \
  spacing: 20\x0a\x0a \
       Image {\x0a \
           sourc\
e: \x22qrc:/images/\
appIcon.svg\x22\x0a   \
         sourceS\
ize: Qt.size(120\
, 120)\x0a         \
   fillMode: Ima\
ge.PreserveAspec\
tFit\x0a\x0a          \
  Layout.alignme\
nt: Qt.AlignTop \
| Qt.AlignHCente\
r\x0a        }\x0a\x0a   \
     Rectangle {\
\x0a            hei\
ght: 1\x0a         \
   color: Materi\
al.color(Materia\
l.Grey)\x0a        \
    Layout.fillW\
idth: true\x0a     \
   }\x0a\x0a        Co\
lumnLayout {\x0a   \
         Label {\
\x0a               \
 text: Qt.applic\
ation.name + ' v\
' + Qt.applicati\
on.version\x0a     \
           font.\
bold: true\x0a     \
       }\x0a       \
     Label {\x0a   \
             tex\
t: \x22Copyright \xc2\xa9\
 2019 - Carlos E\
. P\xc3\xa9rez\x22\x0a      \
          font.i\
talic: true\x0a    \
        }\x0a      \
      Label {\x0a  \
              te\
xt: qsTr(\x22This p\
rogram allows sc\
aning a network \
looking for \x22 +\x0a\
                \
           \x22avai\
lable proxies to\
 connect with.\x22)\
\x0a               \
 wrapMode: Label\
.WordWrap\x0a\x0a     \
           Layou\
t.preferredWidth\
: 220\x0a          \
  }\x0a        } //\
 ColumnLayout\x0a  \
  } // ColumnLay\
out\x0a} // Dialog\x0a\
\
\x00\x00\x09\xf8\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aItem {\x0a    id\
: root\x0a    clip:\
 true\x0a\x0a    RowLa\
yout {\x0a        i\
d: rootLayout\x0a  \
      width: roo\
t.width\x0a        \
clip: true\x0a     \
   spacing: 20\x0a \
       opacity: \
0.0\x0a\x0a        // \
TODO: Use `add`,\
 `remove`, etc. \
transitions\x0a    \
    Component.on\
Completed: { opa\
city = 1.0 } // \
Not the best way\
 to do this\x0a    \
    Behavior on \
opacity { Number\
Animation { dura\
tion: 500; easin\
g.type: Easing.O\
utQuint } }\x0a\x0a   \
     RowLayout {\
\x0a            Lay\
out.fillWidth: t\
rue\x0a            \
spacing: 8\x0a     \
       Image {\x0a \
               s\
ource: \x22qrc:/ima\
ges/icons/qr.svg\
\x22\x0a              \
  sourceSize: \x222\
4x24\x22\x0a          \
  }\x0a            \
TextField {\x0a    \
            id: \
textFieldDestina\
tionAddress\x0a    \
            font\
.family: \x22Code N\
ew Roman\x22\x0a      \
          placeh\
olderText: qsTr(\
\x22Destination add\
ress\x22)\x0a         \
       text: add\
ress\x0a           \
     Layout.fill\
Width: true\x0a    \
        }\x0a      \
  }\x0a        RowL\
ayout {\x0a        \
    TextField {\x0a\
                \
id: textFieldDes\
tinationAmount\x0a \
               t\
ext: sky\x0a       \
         implici\
tWidth: 60\x0a     \
           valid\
ator: DoubleVali\
dator {\x0a        \
            nota\
tion: DoubleVali\
dator.StandardNo\
tation\x0a         \
       }\x0a       \
     }\x0a         \
   Label {\x0a     \
           text:\
 qsTr(\x22SKY\x22)\x0a   \
         }\x0a     \
   }\x0a        Row\
Layout {\x0a       \
     visible: !c\
heckBoxAutomatic\
CoinHoursAllocat\
ion.checked\x0a    \
        TextFiel\
d {\x0a            \
    id: textFiel\
dCoinHoursAmount\
\x0a               \
 text: coinHours\
\x0a               \
 implicitWidth: \
60\x0a             \
   validator: Do\
ubleValidator {\x0a\
                \
    notation: Do\
ubleValidator.St\
andardNotation\x0a \
               }\
\x0a            }\x0a \
           Label\
 {\x0a             \
   text: qsTr(\x22C\
oin hours\x22)\x0a    \
        }\x0a      \
  }\x0a        Tool\
Button {\x0a       \
     id: toolBut\
tonAddRemoveDest\
ination\x0a        \
    // The 'acce\
nt' attribute is\
 used for button\
 highlighting\x0a  \
          Materi\
al.accent: index\
 === 0 ? Materia\
l.Teal : Materia\
l.Red\x0a          \
  icon.source: \x22\
qrc:/images/icon\
s/\x22 + (index ===\
 0 ? \x22add\x22 : \x22re\
move\x22) + \x22-circl\
e.svg\x22\x0a         \
   highlighted: \
true // enable t\
he usage of the \
`Material.accent\
` attribute\x0a\x0a   \
         Layout.\
alignment: Qt.Al\
ignRight\x0a\x0a      \
      onClicked:\
 {\x0a             \
   if (index ===\
 0) {\x0a          \
          listMo\
delDestinations.\
append( { \x22addre\
ss\x22: \x22\x22, \x22sky\x22: \
0.0, \x22coinHours\x22\
: 0.0 } )\x0a      \
          } else\
 {\x0a             \
       listModel\
Destinations.rem\
ove(index)\x0a     \
           }\x0a   \
         }\x0a     \
   }\x0a    } // Ro\
wLayout (rootLay\
out)\x0a}\x0a\
\x00\x00\x03\xfc\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aItem {\x0a    id\
: root\x0a\x0a    prop\
erty alias text:\
 textFieldPasswo\
rd.text\x0a    prop\
erty alias place\
holderText: text\
FieldPassword.pl\
aceholderText\x0a\x0a \
   implicitHeigh\
t: textFieldPass\
word.implicitHei\
ght + buttonForg\
ot.implicitHeigh\
t\x0a\x0a    function \
forceTextFocus()\
 {\x0a        textF\
ieldPassword.for\
ceActiveFocus()\x0a\
    }\x0a\x0a    funct\
ion clear() {\x0a  \
      textFieldP\
assword.clear()\x0a\
    }\x0a\x0a    Colum\
nLayout {\x0a      \
  anchors.fill: \
parent\x0a\x0a        \
TextField {\x0a    \
        id: text\
FieldPassword\x0a\x0a \
           place\
holderText: qsTr\
(\x22Password\x22)\x0a   \
         echoMod\
e: TextField.Pas\
sword\x0a          \
  focus: true\x0a  \
          Layout\
.alignment: Qt.A\
lignTop\x0a        \
    Layout.fillW\
idth: true\x0a     \
   }\x0a\x0a        Bu\
tton {\x0a         \
   id: buttonFor\
got\x0a            \
text: qsTr(\x22I fo\
rgot my password\
\x22)\x0a            f\
lat: true\x0a      \
      Layout.ali\
gnment: Qt.Align\
HCenter | Qt.Ali\
gnTop\x0a          \
  Layout.fillWid\
th: true\x0a       \
 }\x0a    }\x0a}\x0a\
\x00\x00\x03\xd8\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aDialog {\x0a    \
id: root\x0a\x0a    pr\
operty alias dat\
e: transactionDe\
tails.date\x0a    p\
roperty alias st\
atus: transactio\
nDetails.status\x0a\
    property ali\
as type: transac\
tionDetails.type\
\x0a    property al\
ias amount: tran\
sactionDetails.a\
mount\x0a    proper\
ty alias hoursRe\
ceived: transact\
ionDetails.hours\
Received\x0a    pro\
perty alias hour\
sBurned: transac\
tionDetails.hour\
sBurned\x0a    prop\
erty alias trans\
actionID: transa\
ctionDetails.tra\
nsactionID\x0a\x0a    \
title: qsTr(\x22Tra\
nsaction details\
\x22)\x0a\x0a    ColumnLa\
yout {\x0a        a\
nchors.fill: par\
ent\x0a        spac\
ing: 20\x0a\x0a       \
 TransactionDeta\
ils {\x0a          \
  id: transactio\
nDetails\x0a       \
     implicitWid\
th: 500\x0a        \
    Layout.fillW\
idth: true\x0a     \
   }\x0a\x0a        Re\
ctangle {\x0a      \
      visible: t\
ransactionDetail\
s.expanded\x0a     \
       height: 1\
\x0a            col\
or: Material.col\
or(Material.Grey\
)\x0a            La\
yout.fillWidth: \
true\x0a        }\x0a \
   }\x0a}\x0a\
\x00\x00\x07\x1a\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aDialog {\x0a    \
id: root\x0a\x0a    Co\
lumnLayout {\x0a   \
     anchors.fil\
l: parent\x0a      \
  spacing: 20\x0a\x0a \
       Image {\x0a \
           sourc\
e: \x22qrc:/images/\
icons/qt_logo_gr\
een_rgb.svg\x22\x0a   \
         sourceS\
ize: Qt.size(72,\
 72)\x0a           \
 fillMode: Image\
.PreserveAspectF\
it\x0a\x0a            \
Layout.alignment\
: Qt.AlignTop | \
Qt.AlignHCenter\x0a\
        }\x0a\x0a     \
   Rectangle {\x0a \
           heigh\
t: 1\x0a           \
 color: \x22gray\x22\x0a \
           Layou\
t.fillWidth: tru\
e\x0a        }\x0a\x0a   \
     ColumnLayou\
t {\x0a            \
Label {\x0a        \
        text: qs\
Tr(\x22This program\
 uses Qt %1\x22).ar\
g(\x225.12.1\x22)\x0a    \
            font\
.bold: true\x0a    \
        }\x0a      \
      Label {\x0a  \
              te\
xt: qsTr(\x22<p>Qt \
is a <i>C++ tool\
kit for cross-pl\
atform applicati\
on \x22 +\x0a         \
                \
  \x22development</\
i>.</p>\x22 +\x0a     \
                \
      \x22<p>Qt pro\
vides single-sou\
rce portability \
across all major\
 desktop and mob\
ile \x22 +\x0a        \
                \
   \x22operating sy\
stems. It is als\
o available for \
embedded Linux a\
nd other \x22 +\x0a   \
                \
        \x22embedde\
d operating syst\
ems.</p><br>\x22 +\x0a\
                \
           \x22<p>Q\
t offers both <i\
>commercial</i> \
and <i>opensourc\
e</i> licences. \
Please see <a hr\
ef=\x5c\x22http://%2/\x5c\
\x22>%2</a> \x22 +\x0a   \
                \
        \x22for an \
overview of Qt l\
icensing.</p><br\
>\x22 +\x0a           \
                \
\x22<p><i>Copyright\
 \xc2\xa9 %1 The Qt Co\
mpany Ltd</i> an\
d other \x22 +\x0a    \
                \
       \x22contribu\
tors. See <a hre\
f=\x5c\x22http://%3/\x5c\x22\
>%3</a> for more\
 information.</p\
><br>\x22).arg(\x22201\
8\x22).arg(\x22qt.io/l\
icensing\x22).arg(\x22\
qt.io\x22)\x0a        \
        wrapMode\
: Label.WordWrap\
\x0a\x0a              \
  Layout.preferr\
edWidth: 300\x0a   \
             Lay\
out.alignment: Q\
t.AlignHCenter\x0a \
           }\x0a   \
     } // Column\
Layout\x0a    } // \
ColumnLayout\x0a} /\
/ Dialog\x0a\
\x00\x00\x025\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
\x0aPage {\x0a    id: \
root\x0a    header:\
 TabBar {\x0a      \
  id: tabBar\x0a   \
     currentInde\
x: swipeView.cur\
rentIndex\x0a\x0a     \
   TabButton {\x0a \
           text:\
 qsTr(\x22Wallets\x22)\
\x0a        }\x0a     \
   TabButton {\x0a \
           text:\
 qsTr(\x22Send\x22)\x0a  \
      }\x0a        \
TabButton {\x0a    \
        text: qs\
Tr(\x22History\x22)\x0a  \
      }\x0a    }\x0a\x0a \
   SwipeView {\x0a \
       id: swipe\
View\x0a        anc\
hors.fill: paren\
t\x0a        curren\
tIndex: tabBar.c\
urrentIndex\x0a\x0a   \
     PageWallets\
 {\x0a        }\x0a\x0a  \
      PageSend {\
\x0a        }\x0a\x0a    \
    PageHistory \
{\x0a        }\x0a    \
}\x0a}\x0a\
\x00\x00\x06\xba\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aDialog {\x0a    \
id: root\x0a\x0a    ti\
tle: qsTr(\x22Passw\
ord requested\x22)\x0a\
    modal: true\x0a\
    clip: true\x0a \
   focus: true\x0a \
   standardButto\
ns: Dialog.Ok | \
Dialog.Cancel\x0a\x0a \
   onOpened: {\x0a \
       forceActi\
veFocus()\x0a      \
  passwordReques\
ter.forceTextFoc\
us()\x0a        sta\
ndardButton(Dial\
og.Ok).enabled =\
 passwordRequest\
er.text !== \x22\x22\x0a \
   }\x0a\x0a    onClos\
ed: {\x0a        pa\
sswordRequester.\
clear()\x0a    }\x0a\x0a \
   ColumnLayout \
{\x0a        anchor\
s.fill: parent\x0a \
       spacing: \
20\x0a\x0a        Tran\
sactionDetails {\
\x0a            id:\
 transactionDeta\
il\x0a\x0a            \
// EXAMPLE\x0a     \
       date: \x2220\
19-02-26 15:27\x22\x0a\
            stat\
us: TransactionD\
etails.Status.Pr\
eview\x0a          \
  type: Transact\
ionDetails.Type.\
Receive\x0a        \
    amount: 10\x0a \
           hours\
Received: 16957\x0a\
            hour\
sBurned: 33901\x0a \
           trans\
actionID: \x22kq7wd\
kkUT736dnuyQasdh\
saDJ9874jk\x22\x0a\x0a   \
         Layout.\
fillWidth: true\x0a\
        }\x0a\x0a     \
   Rectangle {\x0a \
           visib\
le: transactionD\
etail.expanded\x0a \
           heigh\
t: 1\x0a           \
 color: Material\
.color(Material.\
Grey)\x0a          \
  Layout.fillWid\
th: true\x0a       \
 }\x0a\x0a        Colu\
mnLayout {\x0a     \
       id: colum\
nLayoutPasswordF\
ield\x0a           \
 Layout.fillWidt\
h: true\x0a\x0a       \
     Label {\x0a   \
             tex\
t: qsTr(\x22Enter t\
he password to c\
onfirm the trans\
action\x22)\x0a       \
         font.bo\
ld: true\x0a       \
     }\x0a\x0a        \
    PasswordRequ\
ester {\x0a        \
        id: pass\
wordRequester\x0a\x0a \
               L\
ayout.topMargin:\
 -10\x0a           \
     Layout.fill\
Width: true\x0a\x0a   \
             onT\
extChanged: {\x0a  \
                \
  root.standardB\
utton(Dialog.Ok)\
.enabled = text \
!== \x22\x22\x0a         \
       }\x0a       \
     }\x0a        }\
\x0a    }\x0a}\x0a\
\x00\x00D\xce\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aDialog {\x0a    \
id: root\x0a\x0a    fo\
cus: true\x0a    mo\
dal: true\x0a    st\
andardButtons: D\
ialog.Ok | Dialo\
g.Cancel\x0a    clo\
sePolicy: Popup.\
NoAutoClose\x0a    \
title: Qt.applic\
ation.name\x0a\x0a    \
implicitWidth: g\
ridNumPad.width \
+ 250\x0a    implic\
itHeight: labelI\
nstructions.heig\
ht + textInput.h\
eight + numPad.h\
eight + 200\x0a\x0a   \
 readonly proper\
ty string numPad\
ButtonText: \x22#\x22\x0a\
    readonly pro\
perty real numPa\
dButtonImplicitS\
ize: 50\x0a    read\
only property re\
al numPadButtonP\
ointSize: 18\x0a\x0a  \
  Component.onCo\
mpleted: {\x0a     \
   standardButto\
n(Dialog.Ok).ena\
bled = textInput\
.length >= 4\x0a   \
 }\x0a\x0a    ColumnLa\
yout {\x0a        a\
nchors.fill: par\
ent\x0a\x0a        Lab\
el {\x0a           \
 id: labelInstru\
ctions\x0a         \
   text: qsTr(\x22E\
nter a hard-to-g\
uess PIN of betw\
een 4 and 8 digi\
ts. \x22 +\x0a        \
               \x22\
The PIN layout i\
s displayed in t\
he hardware wall\
et screen.\x22)\x0a   \
         Layout.\
fillWidth: true\x0a\
            Layo\
ut.alignment: Qt\
.AlignHCenter\x0a  \
          horizo\
ntalAlignment: L\
abel.AlignHCente\
r\x0a            wr\
apMode: Label.Wo\
rdWrap\x0a        }\
\x0a\x0a        TextFi\
eld {\x0a          \
  id: textInput\x0a\
\x0a            pla\
ceholderText: qs\
Tr(\x2212345678\x22)\x0a \
           valid\
ator: IntValidat\
or {\x0a           \
     bottom: 111\
11111\x0a          \
      top: 99999\
999\x0a            \
}\x0a\x0a            i\
nputMethodHints:\
 Qt.ImhDigitsOnl\
y\x0a            La\
yout.fillWidth: \
true\x0a           \
 readOnly: true\x0a\
            maxi\
mumLength: 8\x0a\x0a  \
          onText\
Changed: {\x0a     \
           stand\
ardButton(Dialog\
.Ok).enabled = t\
extInput.length \
>= 4\x0a           \
 }\x0a        }\x0a\x0a  \
      Rectangle \
{\x0a            id\
: numPad\x0a       \
     Layout.alig\
nment: Qt.AlignH\
Center\x0a         \
   border.width:\
 1\x0a            b\
order.color: \x22#2\
0000000\x22\x0a       \
     color: Mate\
rial.background\x0a\
            impl\
icitWidth: Math.\
max(gridNumPad.i\
mplicitWidth, ba\
ckspace.implicit\
Width) + 16\x0a    \
        implicit\
Height: gridNumP\
ad.implicitHeigh\
t + backspace.im\
plicitHeight + 2\
0\x0a\x0a            C\
olumnLayout {\x0a  \
              an\
chors.centerIn: \
parent\x0a\x0a        \
        GridLayo\
ut {\x0a           \
         id: gri\
dNumPad\x0a        \
            Layo\
ut.alignment: Qt\
.AlignCenter\x0a   \
                \
 columns: 3\x0a\x0a   \
                \
 ItemDelegate {\x0a\
                \
        id: topL\
eft\x0a            \
            text\
: numPadButtonTe\
xt\x0a             \
           font.\
pointSize: numPa\
dButtonPointSize\
\x0a               \
         implici\
tWidth: numPadBu\
ttonImplicitSize\
\x0a               \
         implici\
tHeight: numPadB\
uttonImplicitSiz\
e\x0a\x0a             \
           Short\
cut {\x0a          \
                \
  sequence: \x227\x22\x0a\
                \
            onAc\
tivated: {\x0a     \
                \
           topLe\
ft.clicked()\x0a   \
                \
             top\
LeftAnimation.st\
art()\x0a          \
                \
  }\x0a            \
            }\x0a  \
                \
      onClicked:\
 {\x0a             \
               c\
onsole.log(\x22Top \
Left clicked\x22)\x0a \
                \
       }\x0a\x0a      \
                \
  SequentialAnim\
ation {\x0a        \
                \
    id: topLeftA\
nimation\x0a       \
                \
     loops: 1\x0a  \
                \
          Proper\
tyAction { targe\
t: topLeft; prop\
erty: \x22downSym\x22;\
 value: true }\x0a \
                \
           Pause\
Animation { dura\
tion: 350 }\x0a    \
                \
        Property\
Action { target:\
 topLeft; proper\
ty: \x22downSym\x22; v\
alue: false }\x0a  \
                \
      }\x0a\x0a       \
                \
 property bool d\
ownSym: down\x0a   \
                \
     property co\
lor color: (down\
 || downSym) ? M\
aterial.color(Ma\
terial.Amber) : \
hovered ? root.M\
aterial.accent :\
 root.Material.f\
oreground\x0a      \
                \
  Material.foreg\
round: color\x0a\x0a  \
                \
      Behavior o\
n color { ColorA\
nimation {} }\x0a  \
                \
  }\x0a\x0a           \
         ItemDel\
egate {\x0a        \
                \
id: topCenter\x0a  \
                \
      text: numP\
adButtonText\x0a   \
                \
     font.pointS\
ize: numPadButto\
nPointSize\x0a     \
                \
   implicitWidth\
: numPadButtonIm\
plicitSize\x0a     \
                \
   implicitHeigh\
t: numPadButtonI\
mplicitSize\x0a\x0a   \
                \
     Shortcut {\x0a\
                \
            sequ\
ence: \x228\x22\x0a      \
                \
      onActivate\
d: {\x0a           \
                \
     topCenter.c\
licked()\x0a       \
                \
         topCent\
erAnimation.star\
t()\x0a            \
                \
}\x0a              \
          }\x0a    \
                \
    onClicked: {\
\x0a               \
             con\
sole.log(\x22Top Ce\
nter clicked\x22)\x0a \
                \
       }\x0a\x0a      \
                \
  SequentialAnim\
ation {\x0a        \
                \
    id: topCente\
rAnimation\x0a     \
                \
       loops: 1\x0a\
                \
            Prop\
ertyAction { tar\
get: topCenter; \
property: \x22downS\
ym\x22; value: true\
 }\x0a             \
               P\
auseAnimation { \
duration: 350 }\x0a\
                \
            Prop\
ertyAction { tar\
get: topCenter; \
property: \x22downS\
ym\x22; value: fals\
e }\x0a            \
            }\x0a\x0a \
                \
       property \
bool downSym: do\
wn\x0a             \
           prope\
rty color color:\
 (down || downSy\
m) ? Material.co\
lor(Material.Amb\
er) : hovered ? \
root.Material.ac\
cent : root.Mate\
rial.foreground\x0a\
                \
        Material\
.foreground: col\
or\x0a\x0a            \
            Beha\
vior on color { \
ColorAnimation {\
} }\x0a            \
        }\x0a\x0a     \
               I\
temDelegate {\x0a  \
                \
      id: topRig\
ht\x0a             \
           text:\
 numPadButtonTex\
t\x0a              \
          font.p\
ointSize: numPad\
ButtonPointSize\x0a\
                \
        implicit\
Width: numPadBut\
tonImplicitSize\x0a\
                \
        implicit\
Height: numPadBu\
ttonImplicitSize\
\x0a\x0a              \
          Shortc\
ut {\x0a           \
                \
 sequence: \x229\x22\x0a \
                \
           onAct\
ivated: {\x0a      \
                \
          topRig\
ht.clicked()\x0a   \
                \
             top\
RightAnimation.s\
tart()\x0a         \
                \
   }\x0a           \
             }\x0a \
                \
       onClicked\
: {\x0a            \
                \
console.log(\x22Top\
 Right clicked\x22)\
\x0a               \
         }\x0a\x0a    \
                \
    SequentialAn\
imation {\x0a      \
                \
      id: topRig\
htAnimation\x0a    \
                \
        loops: 1\
\x0a               \
             Pro\
pertyAction { ta\
rget: topRight; \
property: \x22downS\
ym\x22; value: true\
 }\x0a             \
               P\
auseAnimation { \
duration: 350 }\x0a\
                \
            Prop\
ertyAction { tar\
get: topRight; p\
roperty: \x22downSy\
m\x22; value: false\
 }\x0a             \
           }\x0a\x0a  \
                \
      property b\
ool downSym: dow\
n\x0a              \
          proper\
ty color color: \
(down || downSym\
) ? Material.col\
or(Material.Ambe\
r) : hovered ? r\
oot.Material.acc\
ent : root.Mater\
ial.foreground\x0a \
                \
       Material.\
foreground: colo\
r\x0a\x0a             \
           Behav\
ior on color { C\
olorAnimation {}\
 }\x0a             \
       }\x0a\x0a      \
              It\
emDelegate {\x0a   \
                \
     id: centerL\
eft\x0a            \
            text\
: numPadButtonTe\
xt\x0a             \
           font.\
pointSize: numPa\
dButtonPointSize\
\x0a               \
         implici\
tWidth: numPadBu\
ttonImplicitSize\
\x0a               \
         implici\
tHeight: numPadB\
uttonImplicitSiz\
e\x0a\x0a             \
           Short\
cut {\x0a          \
                \
  sequence: \x224\x22\x0a\
                \
            onAc\
tivated: {\x0a     \
                \
           cente\
rLeft.clicked()\x0a\
                \
                \
centerLeftAnimat\
ion.start()\x0a    \
                \
        }\x0a      \
                \
  }\x0a            \
            onCl\
icked: {\x0a       \
                \
     console.log\
(\x22Center Left cl\
icked\x22)\x0a        \
                \
}\x0a\x0a             \
           Seque\
ntialAnimation {\
\x0a               \
             id:\
 centerLeftAnima\
tion\x0a           \
                \
 loops: 1\x0a      \
                \
      PropertyAc\
tion { target: c\
enterLeft; prope\
rty: \x22downSym\x22; \
value: true }\x0a  \
                \
          PauseA\
nimation { durat\
ion: 350 }\x0a     \
                \
       PropertyA\
ction { target: \
centerLeft; prop\
erty: \x22downSym\x22;\
 value: false }\x0a\
                \
        }\x0a\x0a     \
                \
   property bool\
 downSym: down\x0a \
                \
       property \
color color: (do\
wn || downSym) ?\
 Material.color(\
Material.Amber) \
: hovered ? root\
.Material.accent\
 : root.Material\
.foreground\x0a    \
                \
    Material.for\
eground: color\x0a\x0a\
                \
        Behavior\
 on color { Colo\
rAnimation {} }\x0a\
                \
    }\x0a\x0a         \
           ItemD\
elegate {\x0a      \
                \
  id: centerCent\
er\x0a             \
           text:\
 numPadButtonTex\
t\x0a              \
          font.p\
ointSize: numPad\
ButtonPointSize\x0a\
                \
        implicit\
Width: numPadBut\
tonImplicitSize\x0a\
                \
        implicit\
Height: numPadBu\
ttonImplicitSize\
\x0a\x0a              \
          Shortc\
ut {\x0a           \
                \
 sequence: \x225\x22\x0a \
                \
           onAct\
ivated: {\x0a      \
                \
          center\
Center.clicked()\
\x0a               \
                \
 centerCenterAni\
mation.start()\x0a \
                \
           }\x0a   \
                \
     }\x0a         \
               o\
nClicked: {\x0a    \
                \
        console.\
log(\x22Center Cent\
er clicked\x22)\x0a   \
                \
     }\x0a\x0a        \
                \
SequentialAnimat\
ion {\x0a          \
                \
  id: centerCent\
erAnimation\x0a    \
                \
        loops: 1\
\x0a               \
             Pro\
pertyAction { ta\
rget: centerCent\
er; property: \x22d\
ownSym\x22; value: \
true }\x0a         \
                \
   PauseAnimatio\
n { duration: 35\
0 }\x0a            \
                \
PropertyAction {\
 target: centerC\
enter; property:\
 \x22downSym\x22; valu\
e: false }\x0a     \
                \
   }\x0a\x0a          \
              pr\
operty bool down\
Sym: down\x0a      \
                \
  property color\
 color: (down ||\
 downSym) ? Mate\
rial.color(Mater\
ial.Amber) : hov\
ered ? root.Mate\
rial.accent : ro\
ot.Material.fore\
ground\x0a         \
               M\
aterial.foregrou\
nd: color\x0a\x0a     \
                \
   Behavior on c\
olor { ColorAnim\
ation {} }\x0a     \
               }\
\x0a\x0a              \
      ItemDelega\
te {\x0a           \
             id:\
 centerRight\x0a   \
                \
     text: numPa\
dButtonText\x0a    \
                \
    font.pointSi\
ze: numPadButton\
PointSize\x0a      \
                \
  implicitWidth:\
 numPadButtonImp\
licitSize\x0a      \
                \
  implicitHeight\
: numPadButtonIm\
plicitSize\x0a\x0a    \
                \
    Shortcut {\x0a \
                \
           seque\
nce: \x226\x22\x0a       \
                \
     onActivated\
: {\x0a            \
                \
    centerRight.\
clicked()\x0a      \
                \
          center\
RightAnimation.s\
tart()\x0a         \
                \
   }\x0a           \
             }\x0a \
                \
       onClicked\
: {\x0a            \
                \
console.log(\x22Cen\
ter Right clicke\
d\x22)\x0a            \
            }\x0a\x0a \
                \
       Sequentia\
lAnimation {\x0a   \
                \
         id: cen\
terRightAnimatio\
n\x0a              \
              lo\
ops: 1\x0a         \
                \
   PropertyActio\
n { target: cent\
erRight; propert\
y: \x22downSym\x22; va\
lue: true }\x0a    \
                \
        PauseAni\
mation { duratio\
n: 350 }\x0a       \
                \
     PropertyAct\
ion { target: ce\
nterRight; prope\
rty: \x22downSym\x22; \
value: false }\x0a \
                \
       }\x0a\x0a      \
                \
  property bool \
downSym: down\x0a  \
                \
      property c\
olor color: (dow\
n || downSym) ? \
Material.color(M\
aterial.Amber) :\
 hovered ? root.\
Material.accent \
: root.Material.\
foreground\x0a     \
                \
   Material.fore\
ground: color\x0a\x0a \
                \
       Behavior \
on color { Color\
Animation {} }\x0a \
                \
   }\x0a\x0a          \
          ItemDe\
legate {\x0a       \
                \
 id: bottomLeft\x0a\
                \
        text: nu\
mPadButtonText\x0a \
                \
       font.poin\
tSize: numPadBut\
tonPointSize\x0a   \
                \
     implicitWid\
th: numPadButton\
ImplicitSize\x0a   \
                \
     implicitHei\
ght: numPadButto\
nImplicitSize\x0a\x0a \
                \
       Shortcut \
{\x0a              \
              se\
quence: \x221\x22\x0a    \
                \
        onActiva\
ted: {\x0a         \
                \
       bottomLef\
t.clicked()\x0a    \
                \
            bott\
omLeftAnimation.\
start()\x0a        \
                \
    }\x0a          \
              }\x0a\
                \
        onClicke\
d: {\x0a           \
                \
 console.log(\x22Bo\
ttom Left clicke\
d\x22)\x0a            \
            }\x0a\x0a \
                \
       Sequentia\
lAnimation {\x0a   \
                \
         id: bot\
tomLeftAnimation\
\x0a               \
             loo\
ps: 1\x0a          \
                \
  PropertyAction\
 { target: botto\
mLeft; property:\
 \x22downSym\x22; valu\
e: true }\x0a      \
                \
      PauseAnima\
tion { duration:\
 350 }\x0a         \
                \
   PropertyActio\
n { target: bott\
omLeft; property\
: \x22downSym\x22; val\
ue: false }\x0a    \
                \
    }\x0a\x0a         \
               p\
roperty bool dow\
nSym: down\x0a     \
                \
   property colo\
r color: (down |\
| downSym) ? Mat\
erial.color(Mate\
rial.Amber) : ho\
vered ? root.Mat\
erial.accent : r\
oot.Material.for\
eground\x0a        \
                \
Material.foregro\
und: color\x0a\x0a    \
                \
    Behavior on \
color { ColorAni\
mation {} }\x0a    \
                \
}\x0a\x0a             \
       ItemDeleg\
ate {\x0a          \
              id\
: bottomCenter\x0a \
                \
       text: num\
PadButtonText\x0a  \
                \
      font.point\
Size: numPadButt\
onPointSize\x0a    \
                \
    implicitWidt\
h: numPadButtonI\
mplicitSize\x0a    \
                \
    implicitHeig\
ht: numPadButton\
ImplicitSize\x0a\x0a  \
                \
      Shortcut {\
\x0a               \
             seq\
uence: \x222\x22\x0a     \
                \
       onActivat\
ed: {\x0a          \
                \
      bottomCent\
er.clicked()\x0a   \
                \
             bot\
tomCenterAnimati\
on.start()\x0a     \
                \
       }\x0a       \
                \
 }\x0a             \
           onCli\
cked: {\x0a        \
                \
    console.log(\
\x22Bottom Center c\
licked\x22)\x0a       \
                \
 }\x0a\x0a            \
            Sequ\
entialAnimation \
{\x0a              \
              id\
: bottomCenterAn\
imation\x0a        \
                \
    loops: 1\x0a   \
                \
         Propert\
yAction { target\
: bottomCenter; \
property: \x22downS\
ym\x22; value: true\
 }\x0a             \
               P\
auseAnimation { \
duration: 350 }\x0a\
                \
            Prop\
ertyAction { tar\
get: bottomCente\
r; property: \x22do\
wnSym\x22; value: f\
alse }\x0a         \
               }\
\x0a\x0a              \
          proper\
ty bool downSym:\
 down\x0a          \
              pr\
operty color col\
or: (down || dow\
nSym) ? Material\
.color(Material.\
Amber) : hovered\
 ? root.Material\
.accent : root.M\
aterial.foregrou\
nd\x0a             \
           Mater\
ial.foreground: \
color\x0a\x0a         \
               B\
ehavior on color\
 { ColorAnimatio\
n {} }\x0a         \
           }\x0a\x0a  \
                \
  ItemDelegate {\
\x0a               \
         id: bot\
tomRight\x0a       \
                \
 text: numPadBut\
tonText\x0a        \
                \
font.pointSize: \
numPadButtonPoin\
tSize\x0a          \
              im\
plicitWidth: num\
PadButtonImplici\
tSize\x0a          \
              im\
plicitHeight: nu\
mPadButtonImplic\
itSize\x0a\x0a        \
                \
Shortcut {\x0a     \
                \
       sequence:\
 \x223\x22\x0a           \
                \
 onActivated: {\x0a\
                \
                \
bottomRight.clic\
ked()\x0a          \
                \
      bottomRigh\
tAnimation.start\
()\x0a             \
               }\
\x0a               \
         }\x0a     \
                \
   onClicked: {\x0a\
                \
            cons\
ole.log(\x22Bottom \
Right clicked\x22)\x0a\
                \
        }\x0a\x0a     \
                \
   SequentialAni\
mation {\x0a       \
                \
     id: bottomR\
ightAnimation\x0a  \
                \
          loops:\
 1\x0a             \
               P\
ropertyAction { \
target: bottomRi\
ght; property: \x22\
downSym\x22; value:\
 true }\x0a        \
                \
    PauseAnimati\
on { duration: 3\
50 }\x0a           \
                \
 PropertyAction \
{ target: bottom\
Right; property:\
 \x22downSym\x22; valu\
e: false }\x0a     \
                \
   }\x0a\x0a          \
              pr\
operty bool down\
Sym: down\x0a      \
                \
  property color\
 color: (down ||\
 downSym) ? Mate\
rial.color(Mater\
ial.Amber) : hov\
ered ? root.Mate\
rial.accent : ro\
ot.Material.fore\
ground\x0a         \
               M\
aterial.foregrou\
nd: color\x0a\x0a     \
                \
   Behavior on c\
olor { ColorAnim\
ation {} }\x0a     \
               }\
\x0a               \
 }\x0a\x0a            \
    ItemDelegate\
 {\x0a             \
       id: backs\
pace\x0a           \
         icon.so\
urce: \x22qrc:/imag\
es/icons/backspa\
ce.svg\x22\x0a        \
            disp\
lay: ItemDelegat\
e.IconOnly\x0a     \
               f\
ont.pointSize: n\
umPadButtonPoint\
Size\x0a           \
         Layout.\
fillWidth: true\x0a\
                \
    Layout.align\
ment: Qt.AlignCe\
nter\x0a\x0a          \
          Shortc\
ut {\x0a           \
             seq\
uence: \x22Backspac\
e\x22\x0a             \
           onAct\
ivated: {\x0a      \
                \
      backspace.\
clicked()\x0a      \
                \
      backspaceA\
nimation.start()\
\x0a               \
         }\x0a     \
               }\
\x0a               \
     onClicked: \
{\x0a              \
          consol\
e.log(\x22Backspace\
 clicked\x22)\x0a     \
               }\
\x0a\x0a              \
      Sequential\
Animation {\x0a    \
                \
    id: backspac\
eAnimation\x0a     \
                \
   loops: 1\x0a    \
                \
    PropertyActi\
on { target: bac\
kspace; property\
: \x22downSym\x22; val\
ue: true }\x0a     \
                \
   PauseAnimatio\
n { duration: 35\
0 }\x0a            \
            Prop\
ertyAction { tar\
get: backspace; \
property: \x22downS\
ym\x22; value: fals\
e }\x0a            \
        }\x0a\x0a     \
               p\
roperty bool dow\
nSym: down\x0a     \
               p\
roperty color co\
lor: (down || do\
wnSym) ? Materia\
l.color(Material\
.Amber) : hovere\
d ? root.Materia\
l.accent : root.\
Material.foregro\
und\x0a            \
        Material\
.foreground: col\
or\x0a\x0a            \
        Behavior\
 on color { Colo\
rAnimation {} }\x0a\
                \
}\x0a            }\x0a\
        }\x0a\x0a     \
   Label {\x0a     \
       id: label\
InstructionsNote\
s\x0a            te\
xt: qsTr(\x22<b>NOT\
E:</b> You can u\
se the numpad, b\
ut the number en\
tered \x22 +\x0a      \
                \
 \x22will be the co\
rresponding numb\
er in the hardwa\
re wallet screen\
.\x22)\x0a            \
font.pointSize: \
9\x0a            fo\
nt.italic: true\x0a\
\x0a            Lay\
out.fillWidth: t\
rue\x0a            \
wrapMode: Label.\
WordWrap\x0a       \
     color: Mate\
rial.color(Mater\
ial.Red)\x0a       \
 }\x0a\x0a    }\x0a\x0a}\x0a\
\x00\x00\x0f\xf0\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aItem {\x0a    id\
: root\x0a\x0a    read\
only property re\
al delegateHeigh\
t: 30\x0a    proper\
ty bool emptyAdd\
ressVisible: tru\
e\x0a    property b\
ool expanded: fa\
lse\x0a    // The f\
ollowing propert\
y is used to avo\
id a binding con\
flict with the `\
height` property\
.\x0a    // Also av\
oids a bug with \
the animation wh\
en collapsing a \
wallet\x0a    reado\
nly property rea\
l finalViewHeigh\
t: expanded ? de\
legateHeight*(ad\
dressList.count)\
 + 50 : 0\x0a\x0a    w\
idth: walletList\
.width\x0a    heigh\
t: itemDelegateM\
ainButton.height\
 + (expanded ? f\
inalViewHeight :\
 0)\x0a\x0a    Behavio\
r on height { Nu\
mberAnimation { \
duration: 250; e\
asing.type: Easi\
ng.OutQuint } }\x0a\
\x0a    ColumnLayou\
t {\x0a        id: \
delegateColumnLa\
yout\x0a        anc\
hors.fill: paren\
t\x0a\x0a        ItemD\
elegate {\x0a      \
      id: itemDe\
legateMainButton\
\x0a            Lay\
out.fillWidth: t\
rue\x0a            \
Layout.alignment\
: Qt.AlignTop\x0a  \
          font.b\
old: expanded\x0a\x0a \
           RowLa\
yout {\x0a         \
       id: deleg\
ateRowLayout\x0a   \
             anc\
hors.fill: paren\
t\x0a              \
  anchors.leftMa\
rgin: listWallet\
LeftMargin\x0a     \
           ancho\
rs.rightMargin: \
listWalletRightM\
argin\x0a          \
      spacing: l\
istWalletSpacing\
\x0a\x0a              \
  Image {\x0a      \
              id\
: status\x0a       \
             sou\
rce: statusIcon\x0a\
                \
    sourceSize: \
\x2224x24\x22\x0a        \
        }\x0a\x0a     \
           Label\
 {\x0a             \
       id: label\
WalletName\x0a     \
               t\
ext: name // a r\
ole of the model\
\x0a               \
     Layout.fill\
Width: true\x0a    \
            }\x0a\x0a \
               I\
mage {\x0a         \
           id: l\
ockIcon\x0a        \
            sour\
ce: \x22qrc:/images\
/icons/lock\x22 + (\
encryptionEnable\
d ? \x22On\x22 : \x22Off\x22\
) + \x22.svg\x22\x0a     \
               s\
ourceSize: \x2224x2\
4\x22\x0a             \
   }\x0a\x0a          \
      Label {\x0a  \
                \
  id: labelSky\x0a \
                \
   text: sky // \
a role of the mo\
del\x0a            \
        color: M\
aterial.accent\x0a \
                \
   horizontalAli\
gnment: Text.Ali\
gnRight\x0a        \
            Layo\
ut.preferredWidt\
h: internalLabel\
sWidth\x0a         \
       }\x0a\x0a      \
          Label \
{\x0a              \
      id: labelC\
oins\x0a           \
         text: c\
oinHours // a ro\
le of the model\x0a\
                \
    horizontalAl\
ignment: Text.Al\
ignRight\x0a       \
             Lay\
out.preferredWid\
th: internalLabe\
lsWidth\x0a        \
        }\x0a      \
      }\x0a\x0a       \
     onClicked: \
{\x0a              \
  expanded = !ex\
panded\x0a         \
   }\x0a        } /\
/ ItemDelegate\x0a\x0a\
        ListView\
 {\x0a            i\
d: addressList\x0a \
           model\
: listAddresses\x0a\
            impl\
icitHeight: expa\
nded ? delegateH\
eight*(addressLi\
st.count) + 50 :\
 0\x0a            o\
pacity: expanded\
 ? 1.0 : 0.0\x0a   \
         clip: t\
rue\x0a            \
interactive: fal\
se\x0a            L\
ayout.fillWidth:\
 true\x0a          \
  Layout.alignme\
nt: Qt.AlignTop\x0a\
\x0a            Beh\
avior on implici\
tHeight { Number\
Animation { dura\
tion: 250; easin\
g.type: Easing.O\
utQuint } }\x0a    \
        Behavior\
 on opacity { Nu\
mberAnimation { \
duration: expand\
ed ? 250 : 1000;\
 easing.type: Ea\
sing.OutQuint } \
}\x0a\x0a            d\
elegate: WalletL\
istAddressDelega\
te {\x0a           \
     width: wall\
etList.width\x0a   \
             hei\
ght: index == 0 \
? delegateHeight\
 + 20 : visible \
? delegateHeight\
 : 0\x0a           \
 }\x0a        }\x0a   \
 } // ColumnLayo\
ut\x0a\x0a    // Roles\
: address, addre\
ssSky, addressCo\
inHours\x0a    // U\
se listModel.app\
end( { \x22address\x22\
: value, \x22addres\
sSky\x22: value, \x22a\
ddressCoinHours\x22\
: value } )\x0a    \
// Or implement \
the model in the\
 backend (a more\
 recommendable a\
pproach)\x0a    Lis\
tModel {\x0a       \
 id: listAddress\
es\x0a        // Th\
e first element \
must exist but w\
ill not be used\x0a\
        ListElem\
ent { address: \x22\
----------------\
----------\x22; add\
ressSky: 0; addr\
essCoinHours: 0 \
}\x0a        ListEl\
ement { address:\
 \x22qrxw7364w8xeru\
sftaxkw87ues\x22; a\
ddressSky: 30; a\
ddressCoinHours:\
 1049 }\x0a        \
ListElement { ad\
dress: \x228745yuet\
srk8tcsku4ryj48i\
je\x22; addressSky:\
 12; addressCoin\
Hours: 16011 }\x0a \
       ListEleme\
nt { address: \x22g\
fdhgs343kweru382\
00384uwqd\x22; addr\
essSky: 0; addre\
ssCoinHours: 72 \
}\x0a        ListEl\
ement { address:\
 \x2200qdqsdjkssvmc\
hskjkxxdg374\x22; a\
ddressSky: 521; \
addressCoinHours\
: 11 }\x0a    }\x0a}\x0a\
\x00\x00\x01\xd6\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0a\x0aCheckDele\
gate {\x0a    id: r\
oot\x0a\x0a    text: a\
ddress // a role\
 of the model\x0a  \
  font.family: \x22\
Code New Roman\x22\x0a\
\x0a    LayoutMirro\
ring.enabled: tr\
ue\x0a    contentIt\
em: Label {\x0a    \
    leftPadding:\
 root.indicator.\
width + root.spa\
cing\x0a        tex\
t: root.text\x0a   \
     verticalAli\
gnment: Qt.Align\
VCenter\x0a        \
color: root.enab\
led ? root.Mater\
ial.foreground :\
 root.Material.h\
intTextColor\x0a   \
 }\x0a}\x0a\
\x00\x00\x075\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aDialog {\x0a    \
id: root\x0a\x0a    fo\
cus: visible\x0a   \
 modal: true\x0a   \
 title: Qt.appli\
cation.name\x0a    \
standardButtons:\
 Dialog.YesToAll\
 | Dialog.Cancel\
\x0a    closePolicy\
: Dialog.NoAutoC\
lose\x0a\x0a    implic\
itWidth: 420\x0a   \
 implicitHeight:\
 300\x0a\x0a    Compon\
ent.onCompleted:\
 {\x0a        stand\
ardButton(Dialog\
.YesToAll).text \
= qsTr(\x22Continue\
\x22)\x0a        stand\
ardButton(Dialog\
.YesToAll).enabl\
ed = wordInput.t\
ext !== \x22\x22\x0a    }\
\x0a\x0a    ColumnLayo\
ut {\x0a        anc\
hors.fill: paren\
t\x0a        spacin\
g: 50\x0a\x0a        R\
owLayout {\x0a     \
       spacing: \
30\x0a\x0a            \
Image {\x0a        \
        id: icon\
\x0a               \
 source: \x22qrc:/i\
mages/icons/back\
up.svg\x22\x0a        \
        sourceSi\
ze: \x2264x64\x22\x0a    \
        }\x0a\x0a     \
       ColumnLay\
out {\x0a          \
      Label {\x0a  \
                \
  id: msgTitle\x0a \
                \
   text: \x22<b>\x22 +\
 qsTr(\x22Enter the\
 word indicated \
in the device\x22) \
+ \x22</b>\x22\x0a       \
             Lay\
out.fillWidth: t\
rue\x0a            \
        wrapMode\
: Text.WordWrap\x0a\
                \
    Layout.align\
ment: Qt.AlignTo\
p\x0a              \
  }\x0a\x0a           \
     Label {\x0a   \
                \
 id: msg\x0a       \
             tex\
t: qsTr(\x22You wil\
l be asked to en\
ter the words \x22 \
+\x0a              \
                \
 \x22of your backup\
 seed in random \
order, \x22 +\x0a     \
                \
          \x22plus \
a few additinal \
words.\x22)\x0a       \
             Lay\
out.fillWidth: t\
rue\x0a            \
        wrapMode\
: Text.WordWrap\x0a\
                \
    Layout.align\
ment: Qt.AlignTo\
p\x0a              \
  }\x0a            \
}\x0a        }\x0a\x0a   \
     TextField {\
\x0a            id:\
 wordInput\x0a     \
       placehold\
erText: qsTr(\x22Re\
quested word\x22)\x0a \
           Layou\
t.fillWidth: tru\
e\x0a            fo\
cus: root.focus\x0a\
\x0a            onT\
extChanged: {\x0a  \
              st\
andardButton(Dia\
log.YesToAll).en\
abled = text !==\
 \x22\x22\x0a            \
}\x0a        }\x0a    \
}\x0a}\x0a\
\x00\x00\x04d\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
\x0aApplicationWind\
ow {\x0a    visible\
: true\x0a    width\
: 680\x0a    height\
: 580\x0a    title:\
 qsTr(\x22FiberCryp\
to Wallet\x22)\x0a\x0a   \
 menuBar: MenuBa\
r {\x0a        clip\
: true\x0a        h\
eight: general.t\
oolPageOpened ? \
0 : implicitHeig\
ht\x0a        Behav\
ior on height { \
NumberAnimation \
{ duration: 200;\
 easing.type: Ea\
sing.OutCubic } \
}\x0a        Menu {\
\x0a            id:\
 menuTools\x0a     \
       title: qs\
Tr(\x22&Tools\x22)\x0a   \
         MenuIte\
m { text: qsTr(\x22\
Outputs\x22) }\x0a    \
        MenuItem\
 { text: qsTr(\x22B\
lockchain\x22); onC\
licked: general.\
openBlockchainPa\
ge() }\x0a        }\
\x0a        Menu {\x0a\
            id: \
menuHelp\x0a       \
     title: qsTr\
(\x22&Help\x22)\x0a      \
      MenuItem {\
 text: qsTr(\x22Abo\
ut FiberCripto\x22)\
 }\x0a            M\
enuItem { text: \
qsTr(\x22About Qt\x22)\
; onClicked: dia\
logAboutQt.open(\
) }\x0a        }\x0a  \
  }\x0a\x0a    General\
StackView {\x0a    \
    id: general\x0a\
        anchors.\
fill: parent\x0a   \
 }\x0a\x0a    // About\
\x0a    DialogAbout\
 {\x0a        id: d\
ialogAbout\x0a     \
   anchors.cente\
rIn: Overlay.ove\
rlay\x0a        mod\
al: true\x0a    }\x0a \
   DialogAboutQt\
 {\x0a        id: d\
ialogAboutQt\x0a   \
     anchors.cen\
terIn: Overlay.o\
verlay\x0a        m\
odal: true\x0a    }\
\x0a}\x0a\
\x00\x00\x0c.\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aItem {\x0a    id\
: root\x0a\x0a    read\
only property re\
al addressListHe\
ight: listViewFi\
lterAddress.heig\
ht\x0a    readonly \
property real de\
legateHeight: 42\
\x0a    property al\
ias tristate: ch\
eckDelegate.tris\
tate\x0a    propert\
y alias walletTe\
xt: checkDelegat\
e.text\x0a\x0a    Colu\
mnLayout {\x0a     \
   width: root.w\
idth\x0a\x0a        Ch\
eckDelegate {\x0a  \
          id: ch\
eckDelegate\x0a\x0a   \
         Layout.\
fillWidth: true\x0a\
            tris\
tate: true\x0a     \
       text: nam\
e\x0a            La\
youtMirroring.en\
abled: true\x0a\x0a   \
         nextChe\
ckState: functio\
n() {\x0a          \
      if (checkS\
tate === Qt.Chec\
ked) {\x0a         \
           if (!\
listViewFilterAd\
dress.allChecked\
) {\x0a            \
            list\
ViewFilterAddres\
s.allChecked = t\
rue\x0a            \
        }\x0a      \
              li\
stViewFilterAddr\
ess.allChecked =\
 false\x0a         \
           retur\
n Qt.Unchecked\x0a \
               }\
 else {\x0a        \
            if (\
listViewFilterAd\
dress.allChecked\
) {\x0a            \
            list\
ViewFilterAddres\
s.allChecked = f\
alse\x0a           \
         }\x0a     \
               l\
istViewFilterAdd\
ress.allChecked \
= true\x0a         \
           retur\
n Qt.Checked\x0a   \
             }\x0a \
           }\x0a\x0a  \
          conten\
tItem: Label {\x0a \
               l\
eftPadding: chec\
kDelegate.indica\
tor.width + chec\
kDelegate.spacin\
g\x0a              \
  text: checkDel\
egate.text\x0a     \
           verti\
calAlignment: Qt\
.AlignVCenter\x0a  \
              co\
lor: checkDelega\
te.enabled ? che\
ckDelegate.Mater\
ial.foreground :\
 checkDelegate.M\
aterial.hintText\
Color\x0a          \
  }\x0a        }\x0a\x0a \
       ListView \
{\x0a            id\
: listViewFilter\
Address\x0a\x0a       \
     property in\
t checkedDelegat\
es: 0\x0a          \
  property bool \
allChecked: fals\
e\x0a\x0a            o\
nCheckedDelegate\
sChanged: {\x0a    \
            if (\
checkedDelegates\
 === 0) {\x0a      \
              ch\
eckDelegate.chec\
kState = Qt.Unch\
ecked\x0a          \
      } else if \
(checkedDelegate\
s === count) {\x0a \
                \
   checkDelegate\
.checkState = Qt\
.Checked\x0a       \
         } else \
{\x0a              \
      checkDeleg\
ate.checkState =\
 Qt.PartiallyChe\
cked\x0a           \
     }\x0a         \
   }\x0a\x0a          \
  Layout.fillWid\
th: true\x0a       \
     height: cou\
nt * delegateHei\
ght\x0a\x0a           \
 model: listAddr\
esses\x0a          \
  delegate: Hist\
oryFilterListAdd\
ressDelegate {\x0a \
               w\
idth: ListView.v\
iew.width\x0a      \
          height\
: delegateHeight\
\x0a               \
 leftPadding: 30\
\x0a               \
 scale: 0.875\x0a  \
              ch\
ecked: ListView.\
view.allChecked\x0a\
\x0a               \
 onCheckedChange\
d: {\x0a           \
         ListVie\
w.view.checkedDe\
legates += check\
ed ? 1 : -1\x0a    \
            }\x0a  \
          }\x0a    \
    } // ListVie\
w\x0a    } // Colum\
nLayout\x0a\x0a    // \
This model can b\
e the same as th\
e wallet address\
 list,\x0a    // as\
 this model need\
 to expose all a\
ddresses for eac\
h wallet.\x0a    //\
 For that, it sh\
ould be implemen\
ted in the backe\
nd, instead of h\
ere.\x0a    ListMod\
el { // EXAMPLE\x0a\
        id: list\
Addresses\x0a\x0a     \
   ListElement {\
 address: \x22qrxw7\
364w8xerusftaxkw\
87ues\x22 }\x0a       \
 ListElement { a\
ddress: \x228745yue\
tsrk8tcsku4ryj48\
ije\x22 }\x0a        L\
istElement { add\
ress: \x22gfdhgs343\
kweru38200384uwq\
d\x22 }\x0a    }\x0a}\x0a\
\x00\x00\x11\xe1\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aPage {\x0a    id\
: root\x0a\x0a    prop\
erty int numberO\
fBlocks: 0\x0a    p\
roperty date tim\
estampLastBlock:\
 \x222000-01-01 00:\
00\x22\x0a    property\
 string hashLast\
Block;\x0a\x0a    prop\
erty int current\
SkySupply: 0\x0a   \
 property int to\
talSkySupply: 0\x0a\
    property int\
 currentCoinHour\
sSupply: 0\x0a    p\
roperty int tota\
lCoinHoursSupply\
: 0\x0a\x0a    signal \
backRequested()\x0a\
\x0a    header: Row\
Layout {\x0a       \
 ToolButton {\x0a  \
          id: to\
olButtonBack\x0a   \
         icon.so\
urce: \x22qrc:/imag\
es/icons/back.sv\
g\x22\x0a            L\
ayout.alignment:\
 Qt.AlignLeft\x0a\x0a \
           onCli\
cked: {\x0a        \
        backRequ\
ested()\x0a        \
    }\x0a        }\x0a\
    }\x0a\x0a    Colum\
nLayout {\x0a      \
  id: columnLayo\
utRoot\x0a        a\
nchors.top: pare\
nt.top\x0a        a\
nchors.left: par\
ent.left\x0a       \
 anchors.right: \
parent.right\x0a   \
     anchors.mar\
gins: 20\x0a       \
 spacing: 20\x0a\x0a  \
      GroupBox {\
\x0a            id:\
 groupBoxBlockDe\
tails\x0a          \
  title: qsTr(\x22B\
lock details\x22)\x0a \
           clip:\
 true\x0a          \
  Layout.fillWid\
th: true\x0a       \
     Layout.alig\
nment: Qt.AlignT\
op | Qt.AlignHCe\
nter\x0a\x0a          \
  ColumnLayout {\
\x0a               \
 id: columnLayou\
tBlockDetails\x0a  \
              sp\
acing: 20\x0a      \
          Materi\
al.foreground: M\
aterial.Grey\x0a\x0a  \
              Co\
lumnLayout {\x0a   \
                \
 Label {\x0a       \
                \
 text: qsTr(\x22Num\
ber of blocks\x22)\x0a\
                \
        font.bol\
d: true\x0a        \
            }\x0a  \
                \
  Label {\x0a      \
                \
  id: labelNumbe\
rOfBlocks\x0a      \
                \
  text: numberOf\
Blocks\x0a         \
           }\x0a   \
             }\x0a\x0a\
                \
RowLayout {\x0a    \
                \
spacing: 20\x0a\x0a   \
                \
 ColumnLayout {\x0a\
                \
        Label {\x0a\
                \
            text\
: qsTr(\x22Timestam\
p of last block\x22\
)\x0a              \
              fo\
nt.bold: true\x0a  \
                \
      }\x0a        \
                \
Label {\x0a        \
                \
    id: labelTim\
estampLastBlock\x0a\
                \
            text\
: Qt.formatDateT\
ime(root.timesta\
mpLastBlock, Qt.\
DefaultLocaleSho\
rtDate)\x0a        \
                \
}\x0a              \
      }\x0a\x0a       \
             Col\
umnLayout {\x0a    \
                \
    Layout.fillW\
idth: true\x0a     \
                \
   Label {\x0a     \
                \
       text: qsT\
r(\x22Hash of last \
block\x22)\x0a        \
                \
    font.bold: t\
rue\x0a            \
                \
Layout.fillWidth\
: true\x0a         \
               }\
\x0a               \
         Label {\
\x0a               \
             id:\
 labelHashLastBl\
ock\x0a            \
                \
text: hashLastBl\
ock\x0a            \
                \
wrapMode: Label.\
WrapAnywhere\x0a   \
                \
     }\x0a         \
           }\x0a   \
             }\x0a \
           } // \
ColumnLayout\x0a   \
     } // GroupB\
ox (block detail\
s)\x0a\x0a        Grou\
pBox {\x0a         \
   id: groupBoxS\
kyDetails\x0a      \
      title: qsT\
r(\x22Sky details\x22)\
\x0a            cli\
p: true\x0a        \
    Layout.fillW\
idth: true\x0a     \
       Layout.al\
ignment: Qt.Alig\
nTop | Qt.AlignH\
Center\x0a\x0a        \
    GridLayout {\
\x0a               \
 rows: 2\x0a       \
         columns\
: 2\x0a            \
    Material.for\
eground: Materia\
l.Grey\x0a\x0a        \
        ColumnLa\
yout {\x0a         \
           Label\
 {\x0a             \
           text:\
 qsTr(\x22Current S\
KY supply\x22)\x0a    \
                \
    font.bold: t\
rue\x0a            \
        }\x0a      \
              La\
bel {\x0a          \
              id\
: labelCurrentSk\
ySupply\x0a        \
                \
text: currentSky\
Supply\x0a         \
           }\x0a   \
             }\x0a\x0a\
                \
ColumnLayout {\x0a \
                \
   Label {\x0a     \
                \
   text: qsTr(\x22T\
otal SKY supply\x22\
)\x0a              \
          font.b\
old: true\x0a      \
              }\x0a\
                \
    Label {\x0a    \
                \
    id: labelTot\
alSkySupply\x0a    \
                \
    text: totalS\
kySupply\x0a       \
             }\x0a \
               }\
\x0a\x0a              \
  ColumnLayout {\
\x0a               \
     Label {\x0a   \
                \
     text: qsTr(\
\x22Current Coin Ho\
urs supply\x22)\x0a   \
                \
     font.bold: \
true\x0a           \
         }\x0a     \
               L\
abel {\x0a         \
               i\
d: labelCurrentC\
oinHoursSupply\x0a \
                \
       text: cur\
rentCoinHoursSup\
ply\x0a            \
        }\x0a      \
          }\x0a\x0a   \
             Col\
umnLayout {\x0a    \
                \
Label {\x0a        \
                \
text: qsTr(\x22Tota\
l Coin Hours sup\
ply\x22)\x0a          \
              fo\
nt.bold: true\x0a  \
                \
  }\x0a            \
        Label {\x0a\
                \
        id: labe\
lTotalCoinHoursS\
upply\x0a          \
              te\
xt: totalCoinHou\
rsSupply\x0a       \
             }\x0a \
               }\
\x0a            } /\
/ GridLayout\x0a   \
     } // GroupB\
ox (sky details)\
\x0a    } // Column\
Layout (root)\x0a}\x0a\
\
\x00\x00\x05\xc8\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aItem {\x0a    id\
: root\x0a\x0a    read\
only property re\
al delegateHeigh\
t: 48\x0a    readon\
ly property alia\
s count: listVie\
wFilters.count\x0a \
   property alia\
s interactive: l\
istViewFilters.i\
nteractive\x0a\x0a    \
implicitWidth: l\
istViewFilters.w\
idth\x0a    implici\
tHeight: listVie\
wFilters.height\x0a\
    clip: true\x0a\x0a\
    ListView {\x0a \
       id: listV\
iewFilters\x0a\x0a    \
    spacing: 20\x0a\
        width: 4\
00\x0a        clip:\
 true\x0a\x0a        m\
odel: modelFilte\
rs\x0a        deleg\
ate: HistoryFilt\
erListDelegate {\
\x0a            id:\
 filterDelegate\x0a\
            widt\
h: ListView.view\
.width\x0a         \
   height: deleg\
ateHeight + addr\
essListHeight\x0a  \
          Compon\
ent.onCompleted:\
 {\x0a             \
   listViewFilte\
rs.height += hei\
ght\x0a            \
    if (index ==\
= count - 1) {\x0a \
                \
   listViewFilte\
rs.height += 24\x0a\
                \
}\x0a            }\x0a\
        }\x0a    }\x0a\
\x0a    // This mod\
el can be the sa\
me as the wallet\
 list,\x0a    // as\
 this model need\
 to expose all w\
allets and their\
 addresses.\x0a    \
// For that, it \
should be implem\
ented in the bac\
kend, instead of\
 here.\x0a    ListM\
odel { // EXAMPL\
E\x0a        id: mo\
delFilters\x0a\x0a    \
    ListElement \
{ name: \x22My firs\
t wallet\x22 }\x0a    \
    ListElement \
{ name: \x22My seco\
nd wallet\x22 }\x0a   \
     ListElement\
 { name: \x22My thi\
rd wallet\x22 }\x0a   \
     ListElement\
 { name: \x22My fou\
rth wallet\x22 }\x0a  \
      ListElemen\
t { name: \x22My fi\
veth wallet\x22 }\x0a \
       ListEleme\
nt { name: \x22My s\
ixth wallet\x22 }\x0a \
   }\x0a}\x0a\
\x00\x00\x09\xb9\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0a\x0a\x0aItem {\x0a \
   id: root\x0a\x0a   \
 implicitWidth: \
400\x0a    implicit\
Height: 400\x0a\x0a   \
 signal walletCr\
eationRequested(\
)\x0a\x0a    Column {\x0a\
        anchors.\
fill: parent\x0a   \
     spacing: 30\
\x0a        Control\
CustomSwitch {\x0a \
           id: s\
witchNewLoadWall\
et\x0a            w\
idth: 300\x0a      \
      height: 70\
\x0a            anc\
hors.left: paren\
t.left\x0a         \
   anchors.right\
: parent.right\x0a\x0a\
            left\
Text: qsTr(\x22New \
wallet\x22)\x0a       \
     rightText: \
qsTr(\x22Load walle\
t\x22)\x0a\x0a           \
 backgroundColor\
: Material.accen\
t\x0a            le\
ftColor: \x22white\x22\
\x0a            rig\
htColor: \x22white\x22\
\x0a\x0a            te\
xtColor: Materia\
l.accent\x0a       \
 }\x0a\x0a        Text\
Field {\x0a        \
    id: walletNa\
me\x0a            a\
nchors.left: par\
ent.left\x0a       \
     anchors.rig\
ht: parent.right\
\x0a\x0a            se\
lectByMouse: tru\
e\x0a            pl\
aceholderText: q\
sTr(\x22Wallet's na\
me\x22)\x0a           \
 focus: true\x0a   \
     }\x0a\x0a        \
ControlGenerateS\
eed {\x0a          \
  id: walletSeed\
\x0a            anc\
hors.left: paren\
t.left\x0a         \
   anchors.right\
: parent.right\x0a \
           width\
: 100\x0a          \
  height: inputC\
ontrolHeight - 1\
0\x0a\x0a            p\
laceholderText: \
qsTr(\x22Wallet's s\
eed\x22)\x0a          \
  buttonLeftText\
: qsTr(\x2212 words\
\x22)\x0a            b\
uttonRightText: \
qsTr(\x2224 words\x22)\
\x0a            but\
tonsVisible: swi\
tchNewLoadWallet\
.isInLeftSide\x0a  \
      }\x0a\x0a       \
 TextArea {\x0a    \
        id: wall\
etSeedConfirm\x0a  \
          anchor\
s.left: parent.l\
eft\x0a            \
anchors.right: p\
arent.right\x0a    \
        height: \
walletSeed.input\
ControlHeight\x0a\x0a \
           selec\
tByMouse: true\x0a \
           wrapM\
ode: TextArea.Wr\
ap\x0a            p\
laceholderText: \
qsTr(\x22Confirm th\
e wallet's seed\x22\
)\x0a            op\
acity: switchNew\
LoadWallet.isInL\
eftSide ? 1.0 : \
0.0\x0a            \
visible: opacity\
 > 0.0\x0a\x0a        \
    Behavior on \
opacity { Number\
Animation { dura\
tion: 100 } }\x0a  \
      }\x0a\x0a       \
 Button {\x0a      \
      id: button\
CreateWallet\x0a   \
         anchors\
.horizontalCente\
r: parent.horizo\
ntalCenter\x0a     \
       width: 12\
0\x0a            he\
ight: 60\x0a       \
     font.bold: \
true\x0a           \
 font.pointSize:\
 12\x0a\x0a           \
 text: qsTr(\x22Cre\
ate\x22)\x0a          \
  highlighted: t\
rue\x0a\x0a           \
 onClicked: {\x0a  \
              wa\
lletCreationRequ\
ested()\x0a        \
    }\x0a        }\x0a\
\x0a        move: T\
ransition {\x0a    \
        NumberAn\
imation { proper\
ties: \x22x,y\x22; dur\
ation: 250; easi\
ng.type: Easing.\
OutQuint }\x0a     \
       PropertyA\
ction { property\
: \x22text\x22; value:\
 (switchNewLoadW\
allet.isInLeftSi\
de ? qsTr(\x22Creat\
e\x22) : qsTr(\x22Load\
\x22)) }\x0a        }\x0a\
    }\x0a}\x0a\
\x00\x00\x14j\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aItem {\x0a    id\
: root\x0a\x0a    read\
only property bo\
ol itemVisible: \
index === 0 || a\
ddressSky > 0 ||\
 emptyAddressVis\
ible\x0a    propert\
y bool showOnlyA\
ddresses: false\x0a\
\x0a    visible: it\
emVisible || opa\
city > 0.0\x0a    o\
pacity: itemVisi\
ble ? 1.0 : 0.0\x0a\
\x0a    Behavior on\
 height { Number\
Animation { dura\
tion: 250; easin\
g.type: Easing.O\
utQuint } }\x0a    \
Behavior on opac\
ity { NumberAnim\
ation { duration\
: 200; easing.ty\
pe: Easing.OutQu\
int } }\x0a\x0a    Row\
Layout {\x0a       \
 id: delegateAdd\
ressMenuRowLayou\
t\x0a        anchor\
s.fill: parent\x0a \
       anchors.l\
eftMargin: listW\
alletLeftMargin\x0a\
        anchors.\
rightMargin: lis\
tWalletRightMarg\
in\x0a        spaci\
ng: listWalletSp\
acing\x0a        vi\
sible: index ===\
 0 && !showOnlyA\
ddresses\x0a\x0a      \
  ToolButton {\x0a \
           id: b\
uttonAddAddress\x0a\
            text\
: qsTr(\x22Add wall\
et\x22)\x0a           \
 icon.source: \x22q\
rc:/images/icons\
/add.svg\x22\x0a      \
      Material.f\
oreground: Mater\
ial.Teal\x0a       \
     Layout.fill\
Width: true\x0a    \
    }\x0a        To\
olButton {\x0a     \
       id: butto\
nToggleVisibilit\
y\x0a            te\
xt: qsTr(\x22Show e\
mpty\x22)\x0a         \
   checkable: tr\
ue\x0a            c\
hecked: emptyAdd\
ressVisible\x0a    \
        icon.sou\
rce: \x22qrc:/image\
s/icons/visible\x22\
 + (checked ? \x22O\
n\x22 : \x22Off\x22) + \x22.\
svg\x22\x0a           \
 Material.accent\
: Material.Indig\
o\x0a            Ma\
terial.foregroun\
d: Material.Grey\
\x0a            Lay\
out.fillWidth: t\
rue\x0a\x0a           \
 onCheckedChange\
d: {\x0a           \
     emptyAddres\
sVisible = check\
ed\x0a            }\
\x0a        }\x0a     \
   ToolButton {\x0a\
            id: \
buttonToggleEncr\
yption\x0a         \
   text: qsTr(\x22E\
ncrypt wallet\x22)\x0a\
            chec\
kable: true\x0a    \
        checked:\
 encryptionEnabl\
ed\x0a            i\
con.source: \x22qrc\
:/images/icons/l\
ock\x22 + (checked \
? \x22On\x22 : \x22Off\x22) \
+ \x22.svg\x22\x0a       \
     Material.ac\
cent: Material.A\
mber\x0a           \
 Material.foregr\
ound: Material.G\
rey\x0a            \
Layout.fillWidth\
: true\x0a\x0a        \
    onCheckedCha\
nged: {\x0a        \
        encrypti\
onEnabled = chec\
ked\x0a            \
}\x0a        }\x0a    \
    ToolButton {\
\x0a            id:\
 buttonEdit\x0a    \
        text: qs\
Tr(\x22Edit wallet\x22\
)\x0a            ic\
on.source: \x22qrc:\
/images/icons/ed\
it.svg\x22\x0a        \
    Material.for\
eground: Materia\
l.Blue\x0a         \
   Layout.fillWi\
dth: true\x0a      \
  }\x0a    } // Row\
Layout (menu)\x0a\x0a \
   RowLayout {\x0a \
       id: deleg\
ateAddressRowLay\
out\x0a        anch\
ors.fill: parent\
\x0a        anchors\
.leftMargin: lis\
tWalletLeftMargi\
n\x0a        anchor\
s.rightMargin: l\
istWalletRightMa\
rgin\x0a        spa\
cing: listWallet\
Spacing\x0a        \
visible: root.vi\
sible && index >\
 0\x0a\x0a        Labe\
l {\x0a            \
id: labelNumber\x0a\
            visi\
ble:  !showOnlyA\
ddresses\x0a       \
     text: index\
\x0a        }\x0a\x0a    \
    Image {\x0a    \
        id: qrIc\
on\x0a            v\
isible:  !showOn\
lyAddresses\x0a    \
        source: \
\x22qrc:/images/ico\
ns/qr.svg\x22\x0a     \
       sourceSiz\
e: \x2216x16\x22\x0a     \
   }\x0a\x0a        Ro\
wLayout {\x0a      \
      TextInput \
{\x0a              \
  id: textAddres\
s\x0a              \
  text: address \
// a role of the\
 model\x0a         \
       readOnly:\
 true\x0a          \
      font.famil\
y: \x22Code New Rom\
an\x22\x0a            \
}\x0a            To\
olButton {\x0a     \
           id: t\
oolButtonCopy\x0a  \
              vi\
sible:  !showOnl\
yAddresses\x0a     \
           icon.\
source: \x22qrc:/im\
ages/icons/copy.\
svg\x22\x0a           \
     Layout.alig\
nment: Qt.AlignL\
eft\x0a\x0a           \
     Image {\x0a   \
                \
 id: imageCopied\
\x0a               \
     anchors.cen\
terIn: parent\x0a  \
                \
  visible:  !sho\
wOnlyAddresses\x0a \
                \
   source: \x22qrc:\
/images/icons/ch\
eck-simple.svg\x22\x0a\
                \
    fillMode: Im\
age.PreserveAspe\
ctFit\x0a          \
          source\
Size: Qt.size(to\
olButtonCopy.ico\
n.width*1.5, too\
lButtonCopy.icon\
.height*1.5)\x0a   \
                \
 z: 1\x0a\x0a         \
           opaci\
ty: 0.0\x0a        \
        }\x0a\x0a     \
           onCli\
cked: {\x0a        \
            text\
Address.selectAl\
l()\x0a            \
        textAddr\
ess.copy()\x0a     \
               t\
extAddress.desel\
ect()\x0a          \
          if (co\
pyAnimation.runn\
ing) {\x0a         \
               c\
opyAnimation.res\
tart()\x0a         \
           } els\
e {\x0a            \
            copy\
Animation.start(\
)\x0a              \
      }\x0a        \
        }\x0a\x0a     \
           Seque\
ntialAnimation {\
\x0a               \
     id: copyAni\
mation\x0a         \
           Numbe\
rAnimation { tar\
get: imageCopied\
; property: \x22opa\
city\x22; to: 1.0; \
easing.type: Eas\
ing.OutCubic }\x0a \
                \
   PauseAnimatio\
n { duration: 10\
00 }\x0a           \
         NumberA\
nimation { targe\
t: imageCopied; \
property: \x22opaci\
ty\x22; to: 0.0; ea\
sing.type: Easin\
g.OutCubic }\x0a   \
             }\x0a\x0a\
            }\x0a  \
          Rectan\
gle {\x0a          \
      id: spacer\
\x0a               \
 visible:  !show\
OnlyAddresses\x0a  \
              La\
yout.fillWidth: \
true\x0a           \
 }\x0a        }\x0a\x0a  \
      Label {\x0a  \
          id: la\
belAddressSky\x0a  \
          visibl\
e:  !showOnlyAdd\
resses\x0a         \
   text: address\
Sky // a role of\
 the model\x0a     \
       color: Ma\
terial.accent\x0a  \
          horizo\
ntalAlignment: T\
ext.AlignRight\x0a \
           Layou\
t.preferredWidth\
: internalLabels\
Width\x0a        }\x0a\
\x0a        Label {\
\x0a            id:\
 labelAddressCoi\
ns\x0a            v\
isible:  !showOn\
lyAddresses\x0a    \
        text: ad\
dressCoinHours /\
/ a role of the \
model\x0a          \
  horizontalAlig\
nment: Text.Alig\
nRight\x0a         \
   Layout.prefer\
redWidth: intern\
alLabelsWidth\x0a  \
      }\x0a    } //\
 RowLayout (addr\
esses)\x0a}\x0a\
\x00\x00\x0c\xd2\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aItemDelegate \
{\x0a    id: root\x0a\x0a\
    property dat\
e modelDate: dat\
e\x0a    property i\
nt modelType: ty\
pe\x0a    property \
int modelStatus:\
 status\x0a    prop\
erty var modelSt\
atusString: [ qs\
Tr(\x22Confirmed\x22),\
 qsTr(\x22Pending\x22)\
, qsTr(\x22Preview\x22\
) ]\x0a    property\
 real modelAmoun\
t: amount\x0a    pr\
operty int model\
HoursReceived: h\
oursReceived\x0a   \
 property int mo\
delHoursBurned: \
hoursBurned\x0a    \
property string \
modelTransaction\
ID\x0a\x0a    implicit\
Width: parent.wi\
dth\x0a    implicit\
Height: (columnL\
ayoutMainContent\
.height < 78 ? 7\
8 : columnLayout\
MainContent.heig\
ht) + rowLayoutR\
oot.anchors.topM\
argin + rowLayou\
tRoot.anchors.bo\
ttomMargin\x0a\x0a    \
RowLayout {\x0a    \
    id: rowLayou\
tRoot\x0a        an\
chors.fill: pare\
nt\x0a        ancho\
rs.leftMargin: 2\
0\x0a        anchor\
s.rightMargin: 2\
0\x0a        anchor\
s.topMargin: 10\x0a\
        anchors.\
bottomMargin: 12\
\x0a\x0a        spacin\
g: 20\x0a\x0a        I\
mage {\x0a         \
   source: \x22qrc:\
/images/icons/se\
nd-blue.svg\x22\x0a   \
         sourceS\
ize: \x2232x32\x22\x0a   \
         fillMod\
e: Image.Preserv\
eAspectFit\x0a     \
       mirror: m\
odelType === Tra\
nsactionDetails.\
Type.Receive\x0a   \
         Layout.\
alignment: Qt.Al\
ignTop | Qt.Alig\
nLeft\x0a        }\x0a\
\x0a        ColumnL\
ayout {\x0a        \
    id: columnLa\
youtMainContent\x0a\
            Layo\
ut.fillWidth: tr\
ue\x0a            L\
ayout.alignment:\
 Qt.AlignTop\x0a\x0a  \
          RowLay\
out {\x0a          \
      spacing: 2\
0\x0a              \
  Layout.fillWid\
th: true\x0a\x0a      \
          Label \
{\x0a              \
      font.bold:\
 true\x0a          \
          text: \
(modelType === T\
ransactionDetail\
s.Type.Receive ?\
 qsTr(\x22Received\x22\
) : qsTr(\x22Sent\x22)\
) + \x22 SKY\x22\x0a     \
           }\x0a\x0a  \
              La\
bel {\x0a          \
          Materi\
al.foreground: M\
aterial.Grey\x0a   \
                \
 text: modelDate\
 // model's role\
\x0a               \
     font.pointS\
ize: Qt.applicat\
ion.font.pointSi\
ze * 0.9\x0a       \
         }\x0a     \
       }\x0a\x0a      \
      ColumnLayo\
ut {\x0a           \
     RowLayout {\
\x0a               \
     id: rowLayo\
utSent\x0a         \
           visib\
le: modelType ==\
= TransactionDet\
ails.Type.Send\x0a \
                \
   Image {\x0a     \
                \
   source: \x22qrc:\
/images/icons/qr\
.svg\x22\x0a          \
              so\
urceSize: \x2224x24\
\x22\x0a              \
      }\x0a        \
            Labe\
l {\x0a            \
            text\
: sentAddress //\
 model's role\x0a  \
                \
      font.famil\
y: \x22Code New Rom\
an\x22\x0a            \
            Layo\
ut.fillWidth: tr\
ue\x0a             \
       }\x0a       \
         }\x0a     \
           RowLa\
yout {\x0a         \
           id: r\
owLayoutReceive\x0a\
                \
    Image {\x0a    \
                \
    source: \x22qrc\
:/images/icons/q\
r.svg\x22\x0a         \
               s\
ourceSize: \x2224x2\
4\x22\x0a             \
       }\x0a       \
             Lab\
el {\x0a           \
             tex\
t: receivedAddre\
ss // model's ro\
le\x0a             \
           font.\
family: \x22Code Ne\
w Roman\x22\x0a       \
                \
 Layout.fillWidt\
h: true\x0a        \
            }\x0a  \
              }\x0a\
            } //\
 ColumnLayout (a\
ddresses)\x0a      \
  } // ColumnLay\
out (main conten\
t)\x0a\x0a        Labe\
l {\x0a            \
text: (modelType\
 === Transaction\
Details.Type.Rec\
eive ? \x22\x22 : \x22-\x22)\
 + amount + \x22 SK\
Y\x22 // model's ro\
le\x0a            f\
ont.pointSize: Q\
t.application.fo\
nt.pointSize * 1\
.25\x0a            \
font.bold: true\x0a\
            Layo\
ut.alignment: Qt\
.AlignTop | Qt.A\
lignRight\x0a      \
  }\x0a\x0a    } // Ro\
wLayout (root)\x0a}\
\x0a\
\x00\x00\x06\xc3\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aItem {\x0a    id\
: root\x0a    impli\
citHeight: 80\x0a\x0a \
   ColumnLayout \
{\x0a        id: co\
lumnLayoutRoot\x0a \
       anchors.f\
ill: parent\x0a\x0a   \
     RowLayout {\
\x0a            id:\
 rowLayoutHeader\
\x0a            spa\
cing: 10\x0a       \
     Layout.alig\
nment: Qt.AlignT\
op\x0a            L\
ayout.fillWidth:\
 true\x0a\x0a         \
   Label {\x0a     \
           id: l\
abelIndex\x0a      \
          text: \
index + 1\x0a      \
          font.p\
ointSize: Qt.app\
lication.font.po\
intSize * 0.9\x0a  \
              fo\
nt.bold: true\x0a  \
          }\x0a    \
        Label {\x0a\
                \
id: labelAddress\
\x0a               \
 text: address\x0a \
               f\
ont.family: \x22Cod\
e New Roman\x22\x0a   \
             fon\
t.pointSize: Qt.\
application.font\
.pointSize * 0.9\
\x0a               \
 font.bold: true\
\x0a               \
 Layout.fillWidt\
h: true\x0a        \
    }\x0a        }\x0a\
\x0a        GridLay\
out {\x0a          \
  columns: 2\x0a   \
         Layout.\
alignment: Qt.Al\
ignTop\x0a         \
   Layout.leftMa\
rgin: labelIndex\
.width + rowLayo\
utHeader.spacing\
\x0a            Lay\
out.fillWidth: t\
rue\x0a\x0a           \
 Label {\x0a       \
         text: q\
sTr(\x22Coins:\x22)\x0a  \
              fo\
nt.pointSize: Qt\
.application.fon\
t.pointSize * 0.\
9\x0a              \
  font.bold: tru\
e\x0a            }\x0a\
            Labe\
l {\x0a            \
    text: addres\
sSky\x0a           \
     font.pointS\
ize: Qt.applicat\
ion.font.pointSi\
ze * 0.9\x0a       \
     }\x0a\x0a        \
    Label {\x0a    \
            text\
: qsTr(\x22Hours:\x22)\
\x0a               \
 font.pointSize:\
 Qt.application.\
font.pointSize *\
 0.9\x0a           \
     font.bold: \
true\x0a           \
 }\x0a            L\
abel {\x0a         \
       text: add\
ressCoinHours\x0a  \
              fo\
nt.pointSize: Qt\
.application.fon\
t.pointSize * 0.\
9\x0a            }\x0a\
        }\x0a    } \
// ColumnLayout\x0a\
}\x0a\
\x00\x00\x09!\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aPage {\x0a    id\
: root\x0a\x0a    prop\
erty alias advan\
cedMode: switchA\
dvancedMode.chec\
ked\x0a\x0a    footer:\
 ToolBar {\x0a     \
   id: tabBarSen\
d\x0a        Materi\
al.primary: Mate\
rial.Blue\x0a      \
  Material.accen\
t: Material.Yell\
ow\x0a\x0a        Tool\
Button {\x0a       \
     id: buttonA\
ddWallet\x0a       \
     anchors.fil\
l: parent\x0a      \
      text: qsTr\
(\x22Send\x22)\x0a       \
     icon.source\
: \x22qrc:/images/i\
cons/send.svg\x22\x0a\x0a\
            onCl\
icked: {\x0a       \
         dialogS\
end.open()\x0a     \
       }\x0a       \
 }\x0a    }\x0a\x0a    Gr\
oupBox {\x0a       \
 id: groupBox\x0a\x0a \
       readonly \
property real ma\
rgins: 50\x0a\x0a     \
   anchors.fill:\
 parent\x0a        \
anchors.topMargi\
n: 10\x0a        an\
chors.leftMargin\
: margins\x0a      \
  anchors.rightM\
argin: margins\x0a \
       anchors.b\
ottomMargin: mar\
gins\x0a        imp\
licitWidth: stac\
kView.width\x0a    \
    clip: true\x0a\x0a\
        label: S\
witchDelegate {\x0a\
            id: \
switchAdvancedMo\
de\x0a\x0a            \
text: qsTr(\x22Adva\
nced mode\x22)\x0a\x0a   \
         onToggl\
ed: {\x0a          \
      if (checke\
d) {\x0a           \
         stackVi\
ew.push(componen\
tAdvanced)\x0a     \
           } els\
e {\x0a            \
        stackVie\
w.pop()\x0a        \
        }\x0a      \
      }\x0a        \
}\x0a\x0a        Stack\
View {\x0a         \
   id: stackView\
\x0a            anc\
hors.fill: paren\
t\x0a            in\
itialItem: compo\
nentSimple\x0a     \
       clip: tru\
e\x0a\x0a            C\
omponent {\x0a     \
           id: c\
omponentSimple\x0a \
               S\
crollView {\x0a    \
                \
contentWidth: si\
mple.implicitWid\
th\x0a             \
       contentHe\
ight: simple.imp\
licitHeight\x0a    \
                \
clip: true\x0a     \
               S\
ubPageSendSimple\
 {\x0a             \
           id: s\
imple\x0a          \
              im\
plicitWidth: sta\
ckView.width\x0a   \
                \
 }\x0a             \
   }\x0a           \
 }\x0a\x0a            \
Component {\x0a    \
            id: \
componentAdvance\
d\x0a              \
  ScrollView {\x0a \
                \
   contentWidth:\
 advanced.implic\
itWidth\x0a        \
            cont\
entHeight: advan\
ced.implicitHeig\
ht\x0a             \
       clip: tru\
e\x0a              \
      SubPageSen\
dAdvanced {\x0a    \
                \
    id: advanced\
\x0a               \
         implici\
tWidth: stackVie\
w.width\x0a        \
            }\x0a  \
              }\x0a\
            }\x0a  \
      }\x0a    } //\
 GroupBox\x0a\x0a    D\
ialogSendTransac\
tion {\x0a        i\
d: dialogSend\x0a  \
      anchors.ce\
nterIn: Overlay.\
overlay\x0a    }\x0a}\x0a\
\
\x00\x00\x0e\xf9\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aPage {\x0a    id\
: root\x0a\x0a    prop\
erty string stat\
usIcon; // an em\
pty string for n\
o icon\x0a    reado\
nly property rea\
l listWalletLeft\
Margin: 20\x0a    r\
eadonly property\
 real listWallet\
RightMargin: 50\x0a\
    readonly pro\
perty real listW\
alletSpacing: 20\
\x0a    readonly pr\
operty real inte\
rnalLabelsWidth:\
 70\x0a\x0a    header:\
 ColumnLayout {\x0a\
\x0a        RowLayo\
ut {\x0a           \
 spacing: listWa\
lletSpacing\x0a    \
        Layout.t\
opMargin: 30\x0a\x0a  \
          Label \
{\x0a              \
  text: qsTr(\x22Na\
me\x22)\x0a           \
     font.pointS\
ize: 9\x0a         \
       Layout.le\
ftMargin: listWa\
lletLeftMargin\x0a \
               L\
ayout.fillWidth:\
 true\x0a          \
  }\x0a            \
Label {\x0a        \
        text: qs\
Tr(\x22Sky\x22)\x0a      \
          font.p\
ointSize: 9\x0a    \
            hori\
zontalAlignment:\
 Text.AlignRight\
\x0a               \
 Layout.preferre\
dWidth: internal\
LabelsWidth\x0a    \
        }\x0a      \
      Label {\x0a  \
              te\
xt: qsTr(\x22Coin h\
ours\x22)\x0a         \
       font.poin\
tSize: 9\x0a       \
         horizon\
talAlignment: Te\
xt.AlignRight\x0a  \
              La\
yout.rightMargin\
: listWalletRigh\
tMargin\x0a        \
        Layout.p\
referredWidth: i\
nternalLabelsWid\
th\x0a            }\
\x0a        }\x0a\x0a    \
    Rectangle {\x0a\
            Layo\
ut.fillWidth: tr\
ue\x0a            h\
eight: 1\x0a       \
     color: \x22#DD\
DDDD\x22\x0a        }\x0a\
    }\x0a\x0a    foote\
r: ToolBar {\x0a   \
     id: tabBarC\
reateUpload\x0a    \
    Material.pri\
mary: Material.B\
lue\x0a        Mate\
rial.accent: Mat\
erial.Yellow\x0a\x0a  \
      RowLayout \
{\x0a            an\
chors.fill: pare\
nt\x0a            T\
oolButton {\x0a    \
            id: \
buttonAddWallet\x0a\
                \
text: qsTr(\x22Add \
wallet\x22)\x0a       \
         icon.so\
urce: \x22qrc:/imag\
es/icons/add.svg\
\x22\x0a              \
  Layout.fillWid\
th: true\x0a       \
     }\x0a         \
   ToolButton {\x0a\
                \
id: buttonLoadWa\
llet\x0a           \
     text: qsTr(\
\x22Load wallet\x22)\x0a \
               i\
con.source: \x22qrc\
:/images/icons/u\
pload.svg\x22\x0a     \
           Layou\
t.fillWidth: tru\
e\x0a            }\x0a\
        }\x0a    }\x0a\
\x0a    ScrollView \
{\x0a        id: sc\
rollItem\x0a       \
 anchors.fill: p\
arent\x0a        Sc\
rollBar.horizont\
al.policy: Scrol\
lBar.AlwaysOff\x0a \
       ListView \
{\x0a            id\
: walletList\x0a   \
         anchors\
.fill: parent\x0a  \
          clip: \
true // limit th\
e painting to it\
's bounding rect\
angle\x0a          \
  model: listWal\
lets\x0a           \
 delegate: Walle\
tListDelegate {}\
\x0a        }\x0a    }\
\x0a\x0a    // Roles: \
name, encryption\
Enabled, sky, co\
inHours\x0a    // U\
se listModel.app\
end( { \x22name\x22: v\
alue, \x22encryptio\
nEnabled\x22: value\
, \x22sky\x22: value, \
\x22coinHours\x22: val\
ue } )\x0a    // Or\
 implement the m\
odel in the back\
end (a more reco\
mmendable approa\
ch)\x0a    ListMode\
l {\x0a        id: \
listWallets\x0a    \
    ListElement \
{ name: \x22My firs\
t wallet\x22; encry\
ptionEnabled: tr\
ue; sky: 5; coin\
Hours: 10 }\x0a    \
    ListElement \
{ name: \x22My seco\
nd wallet\x22; encr\
yptionEnabled: t\
rue; sky: 300; c\
oinHours: 1049 }\
\x0a        ListEle\
ment { name: \x22My\
 third wallet\x22; \
encryptionEnable\
d: false; sky: 1\
3; coinHours: 20\
1 }\x0a\x0a        Lis\
tElement { name:\
 \x22My first walle\
t\x22; encryptionEn\
abled: false; sk\
y: 5; coinHours:\
 10 }\x0a        Li\
stElement { name\
: \x22My second wal\
let\x22; encryption\
Enabled: true; s\
ky: 300; coinHou\
rs: 1049 }\x0a     \
   ListElement {\
 name: \x22My third\
 wallet\x22; encryp\
tionEnabled: tru\
e; sky: 13; coin\
Hours: 201 }\x0a\x0a  \
      ListElemen\
t { name: \x22My fi\
rst wallet\x22; enc\
ryptionEnabled: \
true; sky: 5; co\
inHours: 10 }\x0a  \
      ListElemen\
t { name: \x22My se\
cond wallet\x22; en\
cryptionEnabled:\
 false; sky: 300\
; coinHours: 104\
9 }\x0a        List\
Element { name: \
\x22My third wallet\
\x22; encryptionEna\
bled: false; sky\
: 13; coinHours:\
 201 }\x0a\x0a        \
ListElement { na\
me: \x22My first wa\
llet\x22; encryptio\
nEnabled: true; \
sky: 5; coinHour\
s: 10 }\x0a        \
ListElement { na\
me: \x22My second w\
allet\x22; encrypti\
onEnabled: false\
; sky: 300; coin\
Hours: 1049 }\x0a  \
      ListElemen\
t { name: \x22My th\
ird wallet\x22; enc\
ryptionEnabled: \
true; sky: 13; c\
oinHours: 201 }\x0a\
    }\x0a}\x0a\
\x00\x00\x0a\x1e\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aPopup {\x0a    i\
d: root\x0a\x0a    pro\
perty alias acti\
onText: actionMs\
g.text\x0a    prope\
rty alias text: \
msg.text\x0a\x0a    fo\
cus: true\x0a    mo\
dal: true\x0a    cl\
osePolicy: Popup\
.NoAutoClose\x0a\x0a  \
  RowLayout {\x0a  \
      anchors.fi\
ll: parent\x0a     \
   spacing: 13\x0a\x0a\
        BusyIndi\
cator {\x0a        \
    id: busyIndi\
cator\x0a          \
  Layout.alignme\
nt: Qt.AlignLeft\
 | Qt.AlignVCent\
er\x0a            p\
roperty color co\
lor: \x22#2196F3\x22\x0a \
           Mater\
ial.accent: colo\
r\x0a            Se\
quentialAnimatio\
n on color {\x0a   \
             loo\
ps: Animation.In\
finite\x0a         \
       running: \
root.visible\x0a   \
             Col\
orAnimation { fr\
om: \x22#2196F3\x22; t\
o: \x22#00BCD4\x22; du\
ration: 500 }\x0a  \
              Co\
lorAnimation { f\
rom: \x22#00BCD4\x22; \
to: \x22#009688\x22; d\
uration: 500 }\x0a \
               C\
olorAnimation { \
from: \x22#009688\x22;\
 to: \x22#4CAF50\x22; \
duration: 500 }\x0a\
                \
ColorAnimation {\
 from: \x22#4CAF50\x22\
; to: \x22#8BC34A\x22;\
 duration: 500 }\
\x0a               \
 ColorAnimation \
{ from: \x22#8BC34A\
\x22; to: \x22#CDDC39\x22\
; duration: 500 \
}\x0a              \
  ColorAnimation\
 { from: \x22#CDDC3\
9\x22; to: \x22#FFC107\
\x22; duration: 500\
 }\x0a             \
   ColorAnimatio\
n { from: \x22#FFC1\
07\x22; to: \x22#FF980\
0\x22; duration: 50\
0 }\x0a            \
    ColorAnimati\
on { from: \x22#FF9\
800\x22; to: \x22#FF57\
22\x22; duration: 5\
00 }\x0a           \
     ColorAnimat\
ion { from: \x22#FF\
5722\x22; to: \x22#F44\
336\x22; duration: \
500 }\x0a          \
      ColorAnima\
tion { from: \x22#F\
44336\x22; to: \x22#E9\
1E63\x22; duration:\
 500 }\x0a         \
       ColorAnim\
ation { from: \x22#\
E91E63\x22; to: \x22#9\
C27B0\x22; duration\
: 500 }\x0a        \
        ColorAni\
mation { from: \x22\
#9C27B0\x22; to: \x22#\
673AB7\x22; duratio\
n: 500 }\x0a       \
         ColorAn\
imation { from: \
\x22#673AB7\x22; to: \x22\
#3F51B5\x22; durati\
on: 500 }\x0a      \
          ColorA\
nimation { from:\
 \x22#3F51B5\x22; to: \
\x22#2196F3\x22; durat\
ion: 500 }\x0a     \
       }\x0a       \
 }\x0a\x0a        Colu\
mnLayout {\x0a\x0a    \
        Label {\x0a\
                \
id: actionMsg\x0a  \
              fo\
nt.pointSize: 13\
\x0a               \
 font.italic: tr\
ue\x0a             \
   Layout.fillWi\
dth: true\x0a      \
          Layout\
.alignment: Qt.A\
lignLeft | Qt.Al\
ignTop\x0a\x0a        \
        // To av\
oid a binding lo\
op\x0a             \
   Component.onC\
ompleted: {\x0a    \
                \
if (text === \x22\x22)\
 {\x0a             \
           text \
= qsTr(\x22Loading.\
..\x22)\x0a           \
         }\x0a     \
           }\x0a   \
         }\x0a\x0a    \
        Label {\x0a\
                \
id: msg\x0a        \
        visible:\
 text !== \x22\x22\x0a   \
             fon\
t.italic: true\x0a \
               L\
ayout.fillWidth:\
 true\x0a          \
      Layout.ali\
gnment: Qt.Align\
Left | Qt.AlignT\
op\x0a             \
   wrapMode: Tex\
t.WordWrap\x0a     \
       }\x0a\x0a      \
  }\x0a    }\x0a\x0a}\x0a\
\x00\x00 +\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aItem {\x0a    id\
: root\x0a\x0a    prop\
erty date date: \
\x222000-01-01 00:0\
0\x22\x0a    property \
int type: Transa\
ctionDetails.Typ\
e.Send\x0a    prope\
rty int status: \
TransactionDetai\
ls.Status.Previe\
w\x0a    property v\
ar statusString:\
 [ qsTr(\x22Confirm\
ed\x22), qsTr(\x22Pend\
ing\x22), qsTr(\x22Pre\
view\x22) ]\x0a    pro\
perty real amoun\
t: 0\x0a    propert\
y int hoursRecei\
ved: 0\x0a    prope\
rty int hoursBur\
ned: 0\x0a    prope\
rty string trans\
actionID\x0a\x0a    re\
adonly property \
bool expanded: b\
uttonMoreDetails\
.checked\x0a\x0a    en\
um Status {\x0a    \
    Confirmed,\x0a \
       Pending,\x0a\
        Preview\x0a\
    }\x0a    enum T\
ype {\x0a        Se\
nd,\x0a        Rece\
ive\x0a    }\x0a\x0a    i\
mplicitHeight: 8\
0 + rowLayoutBas\
icDetails.height\
 + (expanded ? r\
owLayoutMoreDeta\
ils.height : 0)\x0a\
\x0a    // TODO: Fi\
x performance pr\
oblem with the a\
nimation\x0a    //B\
ehavior on impli\
citHeight { Numb\
erAnimation { du\
ration: 1000; ea\
sing.type: Easin\
g.OutQuint } }\x0a\x0a\
    implicitWidt\
h: 650\x0a    clip:\
 true\x0a\x0a    Colum\
nLayout {\x0a      \
  id: columnLayo\
utRoot\x0a        a\
nchors.fill: par\
ent\x0a        spac\
ing: 20\x0a\x0a       \
 RowLayout {\x0a   \
         id: row\
LayoutBasicDetai\
ls\x0a            L\
ayout.alignment:\
 Qt.AlignTop\x0a   \
         Layout.\
fillWidth: true\x0a\
\x0a            Col\
umnLayout {\x0a    \
            Layo\
ut.alignment: Qt\
.AlignTop\x0a      \
          Layout\
.fillWidth: true\
\x0a\x0a              \
  Label {\x0a      \
              te\
xt: qsTr(\x22Transa\
ction\x22)\x0a        \
            font\
.bold: true\x0a    \
                \
Layout.alignment\
: Qt.AlignTop\x0a  \
                \
  Layout.fillWid\
th: true\x0a       \
         }\x0a\x0a    \
            Grid\
Layout {\x0a       \
             id:\
 gridLayoutBasic\
Info\x0a           \
         Materia\
l.foreground: Ma\
terial.Grey\x0a    \
                \
columns: 2\x0a     \
               c\
olumnSpacing: 10\
\x0a\x0a              \
      Layout.ali\
gnment: Qt.Align\
Top\x0a            \
        Layout.f\
illWidth: true\x0a\x0a\
                \
    Label {\x0a    \
                \
    text: qsTr(\x22\
Date:\x22)\x0a        \
                \
font.pointSize: \
Qt.application.f\
ont.pointSize * \
0.9\x0a            \
            font\
.bold: true\x0a    \
                \
}\x0a              \
      Label {\x0a  \
                \
      text: Qt.f\
ormatDateTime(ro\
ot.date, Qt.Defa\
ultLocaleShortDa\
te)\x0a            \
            font\
.pointSize: Qt.a\
pplication.font.\
pointSize * 0.9\x0a\
                \
    }\x0a\x0a         \
           Label\
 {\x0a             \
           text:\
 qsTr(\x22Status:\x22)\
\x0a               \
         font.po\
intSize: Qt.appl\
ication.font.poi\
ntSize * 0.9\x0a   \
                \
     font.bold: \
true\x0a           \
         }\x0a     \
               L\
abel {\x0a         \
               t\
ext: statusStrin\
g[root.status]\x0a \
                \
       font.poin\
tSize: Qt.applic\
ation.font.point\
Size * 0.9\x0a     \
               }\
\x0a\x0a              \
      Label {\x0a  \
                \
      text: qsTr\
(\x22Hours:\x22)\x0a     \
                \
   font.pointSiz\
e: Qt.applicatio\
n.font.pointSize\
 * 0.9\x0a         \
               f\
ont.bold: true\x0a \
                \
   }\x0a           \
         Label {\
\x0a               \
         text: r\
oot.hoursReceive\
d + ' ' + qsTr(\x22\
received\x22) + ' |\
 ' + hoursBurned\
 + ' ' + qsTr(\x22b\
urned\x22)\x0a        \
                \
font.pointSize: \
Qt.application.f\
ont.pointSize * \
0.9\x0a            \
        }\x0a\x0a     \
               L\
abel {\x0a         \
               t\
ext: qsTr(\x22Tx ID\
:\x22)\x0a            \
            font\
.pointSize: Qt.a\
pplication.font.\
pointSize * 0.9\x0a\
                \
        font.bol\
d: true\x0a        \
            }\x0a  \
                \
  Label {\x0a      \
                \
  text: root.tra\
nsactionID\x0a     \
                \
   font.pointSiz\
e: Qt.applicatio\
n.font.pointSize\
 * 0.9\x0a         \
               L\
ayout.fillWidth:\
 true\x0a          \
          }\x0a    \
            } //\
 GridLayout\x0a    \
        }\x0a\x0a     \
       ColumnLay\
out {\x0a          \
      Layout.ali\
gnment: Qt.Align\
Top\x0a            \
    Layout.right\
Margin: 20\x0a     \
           Image\
 {\x0a             \
       source: \x22\
qrc:/images/icon\
s/send-\x22 + (type\
 === Transaction\
Details.Type.Rec\
eive ? \x22blue\x22 : \
\x22amber\x22) + \x22.svg\
\x22\x0a              \
      sourceSize\
: \x2272x72\x22\x0a      \
              fi\
llMode: Image.Pr\
eserveAspectFit\x0a\
                \
    mirror: type\
 === Transaction\
Details.Type.Rec\
eive\x0a           \
         Layout.\
fillWidth: true\x0a\
                \
}\x0a              \
  Label {\x0a      \
              te\
xt: (type === Tr\
ansactionDetails\
.Type.Receive ? \
\x22Receive\x22 : \x22Sen\
d\x22) + ' ' + amou\
nt + ' ' + qsTr(\
\x22SKY\x22)\x0a         \
           font.\
bold: true\x0a     \
               f\
ont.pointSize: Q\
t.application.fo\
nt.pointSize * 1\
.25\x0a            \
        horizont\
alAlignment: Lab\
el.AlignHCenter\x0a\
                \
    Layout.fillW\
idth: true\x0a     \
           }\x0a   \
         }\x0a     \
   } // RowLayou\
t\x0a\x0a        RowLa\
yout {\x0a         \
   id: rowLayout\
DetailsSeparator\
\x0a\x0a            La\
yout.alignment: \
Qt.AlignTop\x0a    \
        Layout.f\
illWidth: true\x0a\x0a\
            Butt\
on {\x0a           \
     id: buttonM\
oreDetails\x0a     \
           impli\
citWidth: 200\x0a  \
              fl\
at: true\x0a       \
         checkab\
le: true\x0a       \
         text: (\
checked ? qsTr(\x22\
Less\x22) : qsTr(\x22M\
ore\x22)) + ' ' + q\
sTr(\x22details\x22)\x0a \
           }\x0a\x0a  \
          Rectan\
gle {\x0a          \
      height: 1\x0a\
                \
color: Material.\
color(Material.G\
rey)\x0a           \
     Layout.alig\
nment: Qt.AlignV\
Center\x0a         \
       Layout.fi\
llWidth: true\x0a  \
          }\x0a    \
    }\x0a\x0a        R\
owLayout {\x0a     \
       id: rowLa\
youtMoreDetails\x0a\
\x0a            Lay\
out.alignment: Q\
t.AlignTop\x0a     \
       Layout.fi\
llWidth: true\x0a\x0a \
           Colum\
nLayout {\x0a      \
          Layout\
.alignment: Qt.A\
lignTop\x0a        \
        Layout.f\
illWidth: true\x0a\x0a\
                \
Label {\x0a        \
            text\
: qsTr(\x22Inputs\x22)\
\x0a               \
     font.bold: \
true\x0a           \
         font.it\
alic: true\x0a     \
               L\
ayout.fillWidth:\
 true\x0a          \
      }\x0a\x0a       \
         ScrollV\
iew {\x0a          \
          Layout\
.alignment: Qt.A\
lignTop\x0a        \
            Layo\
ut.fillWidth: tr\
ue\x0a\x0a            \
        contentH\
eight: 160\x0a\x0a    \
                \
ListView {\x0a     \
                \
   id: listViewI\
nputs\x0a\x0a         \
               M\
aterial.foregrou\
nd: Material.Gre\
y\x0a              \
          model:\
 listModelInputs\
\x0a               \
         clip: t\
rue\x0a            \
            Layo\
ut.alignment: Qt\
.AlignTop\x0a      \
                \
  Layout.fillWid\
th: true\x0a       \
                \
 delegate: Input\
OutputDelegate {\
\x0a               \
             wid\
th: ListView.vie\
w.width\x0a        \
                \
}\x0a              \
      }\x0a        \
        }\x0a      \
      }\x0a\x0a       \
     ColumnLayou\
t {\x0a            \
    Layout.align\
ment: Qt.AlignTo\
p\x0a              \
  Layout.fillWid\
th: true\x0a\x0a      \
          Label \
{\x0a              \
      text: qsTr\
(\x22Outputs\x22)\x0a    \
                \
font.bold: true\x0a\
                \
    font.italic:\
 true\x0a          \
          Layout\
.fillWidth: true\
\x0a               \
 }\x0a\x0a            \
    ScrollView {\
\x0a               \
     Layout.alig\
nment: Qt.AlignT\
op\x0a             \
       Layout.fi\
llWidth: true\x0a\x0a \
                \
   contentHeight\
: 160\x0a\x0a         \
           ListV\
iew {\x0a          \
              id\
: listViewOutput\
s\x0a\x0a             \
           Mater\
ial.foreground: \
Material.Grey\x0a  \
                \
      model: lis\
tModelOutputs\x0a  \
                \
      clip: true\
\x0a               \
         Layout.\
alignment: Qt.Al\
ignTop\x0a         \
               L\
ayout.fillWidth:\
 true\x0a          \
              de\
legate: InputOut\
putDelegate {\x0a  \
                \
          width:\
 ListView.view.w\
idth\x0a           \
             }\x0a \
                \
   }\x0a           \
     }\x0a         \
   }\x0a        } /\
/ RowLayout\x0a    \
} // ColumnLayou\
t (root)\x0a\x0a    //\
 Roles: address,\
 addressSky, add\
ressCoinHours\x0a  \
  // Use listMod\
el.append( { \x22ad\
dress\x22: value, \x22\
addressSky\x22: val\
ue, \x22addressCoin\
Hours\x22: value } \
)\x0a    // Or impl\
ement the model \
in the backend (\
a more recommend\
able approach)\x0a \
   ListModel {\x0a \
       id: listM\
odelInputs\x0a     \
   ListElement {\
 address: \x22qrxw7\
364w8xerusftaxkw\
87ues\x22; addressS\
ky: 30; addressC\
oinHours: 1049 }\
\x0a        ListEle\
ment { address: \
\x228745yuetsrk8tcs\
ku4ryj48ije\x22; ad\
dressSky: 12; ad\
dressCoinHours: \
16011 }\x0a    }\x0a  \
  ListModel {\x0a  \
      id: listMo\
delOutputs\x0a     \
   ListElement {\
 address: \x22734ir\
weaweygtawieta78\
3cwyc\x22; addressS\
ky: 38; addressC\
oinHours: 5048 }\
\x0a        ListEle\
ment { address: \
\x22ekq03i3qerwhjqo\
qh9823yurig\x22; ad\
dressSky: 61; ad\
dressCoinHours: \
9456 }\x0a        L\
istElement { add\
ress: \x221kjher73y\
iner7wn32nkuwe94\
v\x22; addressSky: \
1; addressCoinHo\
urs: 24 }\x0a      \
  ListElement { \
address: \x22oopfww\
klfd34iuhjwe83w3\
h28r\x22; addressSk\
y: 111; addressC\
oinHours: 13548 \
}\x0a    }\x0a}\x0a\
\x00\x00\x0bh\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.impl 2.1\
2\x0aimport QtQuick\
.Controls.Materi\
al 2.12\x0a\x0a// Colo\
rs are only impl\
emented for Mate\
rial Style\x0a// TO\
DO: Implement th\
e same colors fo\
r the other styl\
es\x0a\x0aRectangle {\x0a\
    id: root\x0a\x0a  \
  readonly prope\
rty real spacing\
: 10\x0a    propert\
y color backgrou\
ndColor: Materia\
l.primary\x0a    pr\
operty bool isIn\
LeftSide: true\x0a \
   property stri\
ng leftText: \x22Le\
ft Text\x22\x0a    pro\
perty string rig\
htText: \x22Right T\
ext\x22\x0a    propert\
y color leftColo\
r: Material.colo\
r(Material.Blue)\
\x0a    property co\
lor rightColor: \
Material.color(M\
aterial.Green)\x0a \
   property int \
animationsSpeed:\
 150\x0a    propert\
y color textColo\
r: Material.fore\
ground\x0a    prope\
rty color hovere\
dColor: \x22#0A0000\
00\x22\x0a    property\
 color pressedCo\
lor: \x22#1A000000\x22\
\x0a    property re\
al controlRadius\
: 40\x0a\x0a    color:\
 backgroundColor\
\x0a    radius: con\
trolRadius\x0a\x0a    \
Button {\x0a       \
 id: switchButto\
n\x0a        clip: \
true\x0a\x0a        an\
chors.top: paren\
t.top\x0a        an\
chors.bottom: pa\
rent.bottom\x0a    \
    anchors.left\
: parent.left\x0a  \
      anchors.le\
ftMargin: root.i\
sInLeftSide ? sp\
acing : root.wid\
th/2 + spacing*2\
\x0a        anchors\
.right: parent.r\
ight\x0a        anc\
hors.rightMargin\
: root.isInLeftS\
ide ? root.width\
/2 + spacing*2 :\
 spacing\x0a\x0a      \
  Material.foreg\
round: textColor\
\x0a\x0a        Behavi\
or on anchors.le\
ftMargin  { Numb\
erAnimation { du\
ration: animatio\
nsSpeed; easing.\
type: Easing.Out\
Quint } }\x0a      \
  Behavior on an\
chors.rightMargi\
n { NumberAnimat\
ion { duration: \
animationsSpeed;\
 easing.type: Ea\
sing.OutQuint } \
}\x0a\x0a        text:\
 \x22<b>\x22 + (root.i\
sInLeftSide ? ro\
ot.leftText : ro\
ot.rightText) + \
\x22</b>\x22\x0a\x0a        \
onClicked: {\x0a   \
         root.is\
InLeftSide = !ro\
ot.isInLeftSide\x0a\
        }\x0a\x0a     \
   contentItem: \
IconLabel {\x0a    \
        spacing:\
 switchButton.sp\
acing\x0a          \
  mirrored: swit\
chButton.mirrore\
d\x0a            di\
splay: switchBut\
ton.display\x0a\x0a   \
         icon: s\
witchButton.icon\
\x0a            tex\
t: switchButton.\
text\x0a           \
 font: switchBut\
ton.font\x0a       \
     color: !swi\
tchButton.enable\
d ? switchButton\
.Material.hintTe\
xtColor : switch\
Button.Material.\
foreground\x0a     \
   }\x0a\x0a        ba\
ckground: Rectan\
gle {\x0a          \
  implicitWidth:\
 64\x0a            \
implicitHeight: \
switchButton.Mat\
erial.buttonHeig\
ht\x0a\x0a            \
radius: root.con\
trolRadius\x0a     \
       color: !s\
witchButton.enab\
led ? switchButt\
on.Material.butt\
onDisabledColor \
: root.isInLeftS\
ide ? root.leftC\
olor : root.righ\
tColor\x0a         \
   Behavior on c\
olor { ColorAnim\
ation { duration\
: animationsSpee\
d } }\x0a\x0a         \
   layer.enabled\
: switchButton.e\
nabled && switch\
Button.Material.\
buttonColor.a > \
0\x0a\x0a            R\
ectangle {\x0a     \
           id: r\
ectangleColorEff\
ect\x0a            \
    anchors.fill\
: parent\x0a       \
         radius:\
 root.controlRad\
ius\x0a\x0a           \
     color: swit\
chButton.down ? \
pressedColor : s\
witchButton.hove\
red ? hoveredCol\
or : \x22transparen\
t\x22\x0a\x0a            \
    Behavior on \
color { ColorAni\
mation { duratio\
n: animationsSpe\
ed } }\x0a         \
   }\x0a        }\x0a \
   }\x0a}\x0a\
\x00\x00\x07m\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aDialog {\x0a    \
id: root\x0a\x0a    pr\
operty alias tex\
t: msg.text\x0a    \
property alias i\
magePath: icon.s\
ource\x0a    proper\
ty alias confirm\
ationCheckBoxVis\
ible: checkBoxCo\
nfirmation.visib\
le\x0a    property \
alias confirmati\
onCheckBoxText: \
checkBoxConfirma\
tion.text\x0a    pr\
operty alias con\
firmationCheckBo\
xChecked: checkB\
oxConfirmation.c\
hecked\x0a    prope\
rty int buttonTo\
ConnectWithConfi\
rmationCheckBox:\
 Dialog.Ok\x0a\x0a    \
focus: true\x0a    \
modal: true\x0a    \
title: Qt.applic\
ation.name\x0a    s\
tandardButtons: \
Dialog.Ok\x0a\x0a    C\
omponent.onCompl\
eted: {\x0a        \
if (standardButt\
on(buttonToConne\
ctWithConfirmati\
onCheckBox)) {\x0a \
           stand\
ardButton(button\
ToConnectWithCon\
firmationCheckBo\
x).enabled = con\
firmationCheckBo\
xChecked\x0a       \
 }\x0a    }\x0a\x0a    on\
ConfirmationChec\
kBoxCheckedChang\
ed: {\x0a        if\
 (confirmationCh\
eckBoxVisible &&\
 standardButton(\
buttonToConnectW\
ithConfirmationC\
heckBox)) {\x0a    \
        standard\
Button(buttonToC\
onnectWithConfir\
mationCheckBox).\
enabled = confir\
mationCheckBoxCh\
ecked\x0a        } \
else if (standar\
dButton(buttonTo\
ConnectWithConfi\
rmationCheckBox)\
) {\x0a            \
standardButton(b\
uttonToConnectWi\
thConfirmationCh\
eckBox).enabled \
= true\x0a        }\
\x0a    }\x0a\x0a    Colu\
mnLayout {\x0a     \
   anchors.fill:\
 parent\x0a        \
spacing: 30\x0a\x0a   \
     RowLayout {\
\x0a            spa\
cing: 30\x0a\x0a      \
      Image {\x0a  \
              id\
: icon\x0a         \
       sourceSiz\
e: \x2264x64\x22\x0a     \
       }\x0a       \
     Label {\x0a   \
             id:\
 msg\x0a           \
     text: qsTr(\
\x22Your message go\
es here.\x22)\x0a     \
           Layou\
t.fillWidth: tru\
e\x0a              \
  Layout.alignme\
nt: Qt.AlignVCen\
ter\x0a            \
    wrapMode: Te\
xt.WordWrap\x0a    \
        }\x0a      \
  }\x0a\x0a        Che\
ckBox {\x0a        \
    id: checkBox\
Confirmation\x0a   \
         visible\
: false\x0a        \
    Layout.fillW\
idth: true\x0a     \
       Layout.al\
ignment: Qt.Alig\
nHCenter\x0a       \
 }\x0a    }\x0a\x0a}\x0a\
\x00\x00\x09e\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0a\x0aItem {\x0a  \
  id: root\x0a\x0a    \
implicitWidth: 1\
00\x0a    implicitH\
eight: 60\x0a\x0a    p\
roperty alias pl\
aceholderText: t\
extArea.placehol\
derText\x0a    prop\
erty alias readO\
nly: textArea.re\
adOnly\x0a    prope\
rty alias button\
LeftText: button\
Left.text\x0a    pr\
operty alias but\
tonRightText: bu\
ttonRight.text\x0a\x0a\
    property boo\
l buttonsVisible\
: buttonRight.vi\
sible\x0a\x0a    reado\
nly property ali\
as inputControlW\
idth: textArea.w\
idth\x0a    readonl\
y property alias\
 inputControlHei\
ght: textArea.he\
ight\x0a\x0a    Row {\x0a\
        id: rowB\
uttons\x0a        a\
nchors.bottom: t\
extArea.top\x0a    \
    anchors.bott\
omMargin: -10\x0a  \
      anchors.ri\
ght: textArea.ri\
ght\x0a        z: t\
extArea.z + 1\x0a\x0a \
       ItemDeleg\
ate {\x0a          \
  id: buttonLeft\
\x0a\x0a            he\
ight: buttonRigh\
t.height\x0a\x0a      \
      text: \x22But\
tonLeft\x22\x0a       \
     font.pointS\
ize: 8\x0a         \
   visible: butt\
onRight.visible\x0a\
            opac\
ity: buttonRight\
.opacity\x0a\x0a      \
      contentIte\
m: Label {\x0a     \
           text:\
 buttonLeft.text\
\x0a               \
 font: buttonLef\
t.font\x0a         \
       color: bu\
ttonLeft.enabled\
 ? buttonLeft.Ma\
terial.foregroun\
d : buttonLeft.M\
aterial.hintText\
Color\x0a          \
  }\x0a        }\x0a\x0a \
       ItemDeleg\
ate {\x0a          \
  id: buttonRigh\
t\x0a\x0a            h\
eight: 30\x0a      \
      opacity: b\
uttonsVisible ? \
1.0 : 0.0\x0a      \
      visible: o\
pacity > 0.0\x0a\x0a  \
          Behavi\
or on opacity { \
NumberAnimation \
{ easing.type: E\
asing.OutQuint }\
 }\x0a\x0a            \
text: \x22ButtonRig\
ht\x22\x0a            \
font.pointSize: \
8\x0a\x0a            c\
ontentItem: Labe\
l {\x0a            \
    text: button\
Right.text\x0a     \
           font:\
 buttonRight.fon\
t\x0a              \
  color: buttonR\
ight.enabled ? b\
uttonRight.Mater\
ial.foreground :\
 buttonRight.Mat\
erial.hintTextCo\
lor\x0a            \
}\x0a        }\x0a    \
}\x0a\x0a    /*\x0a    //\
 Put the text ar\
ea inside this t\
o make it scroll\
able\x0a    // but \
we have a bug.\x0a \
   // Set the an\
chors of the lef\
t button relativ\
e to this Scroll\
View\x0a    ScrollV\
iew {\x0a        id\
: scrollerTextAr\
ea\x0a        ancho\
rs.left: parent.\
left\x0a        anc\
hors.right: pare\
nt.right\x0a\x0a      \
  width: root.wi\
dth\x0a        heig\
ht: root.height \
- buttonRight.he\
ight\x0a\x0a        co\
ntentWidth: widt\
h\x0a        conten\
tHeight: height\x0a\
\x0a        clip: t\
rue\x0a    }\x0a    */\
\x0a\x0a    TextArea {\
\x0a        id: tex\
tArea\x0a        an\
chors.left: pare\
nt.left\x0a        \
anchors.right: p\
arent.right\x0a    \
    height: 80\x0a \
       wrapMode:\
 TextArea.Wrap\x0a\x0a\
        selectBy\
Mouse: true\x0a    \
}\x0a}\x0a\
\x00\x00\x09\x0d\
i\
mport QtQuick 2.\
12\x0aimport QtQuic\
k.Controls 2.12\x0a\
import QtQuick.C\
ontrols.Material\
 2.12\x0aimport QtQ\
uick.Layouts 1.1\
2\x0a\x0aDialog {\x0a    \
id: root\x0a\x0a    fo\
cus: true\x0a    mo\
dal: true\x0a    ti\
tle: Qt.applicat\
ion.name\x0a    sta\
ndardButtons: Di\
alog.Abort\x0a    c\
losePolicy: Dial\
og.NoAutoClose\x0a\x0a\
    implicitWidt\
h: 650\x0a    impli\
citHeight: 380\x0a\x0a\
    ColumnLayout\
 {\x0a        ancho\
rs.fill: parent\x0a\
        spacing:\
 50\x0a\x0a        Row\
Layout {\x0a       \
     spacing: 50\
\x0a\x0a            Im\
age {\x0a          \
      id: icon\x0a \
               s\
ource: \x22qrc:/ima\
ges/icons/cpu.sv\
g\x22\x0a             \
   sourceSize: \x22\
64x64\x22\x0a         \
       Layout.al\
ignment: Qt.Alig\
nTop\x0a           \
 }\x0a\x0a            \
ColumnLayout {\x0a \
               L\
abel {\x0a         \
           id: m\
sgTitle\x0a        \
            text\
: qsTr(\x22Unconfig\
ured Hardware Wa\
llet\x22)\x0a         \
           font.\
bold: true\x0a     \
               L\
ayout.fillWidth:\
 true\x0a          \
          wrapMo\
de: Text.WordWra\
p\x0a              \
      Layout.ali\
gnment: Qt.Align\
Top\x0a            \
    }\x0a          \
      Label {\x0a  \
                \
  id: msg\x0a      \
              te\
xt: qsTr(\x22A seed\
less hardware wa\
llet has been de\
tected. \x22 +\x0a    \
                \
           \x22Sele\
ct <b><i>Configu\
re automatically\
</i></b> \x22 +\x0a   \
                \
            \x22if \
you want to conf\
igure it as a br\
and new wallet \x22\
 +\x0a             \
                \
  \x22and start usi\
ng it inmediatel\
y, \x22 +\x0a         \
                \
      \x22or select\
 <b><i>Restore b\
ackup</i></b> \x22 \
+\x0a              \
                \
 \x22if you want to\
 configure it wi\
th a previously \
created \x22 +\x0a    \
                \
           \x22seed\
 backup and thus\
 be able to acce\
ss your funds ag\
ain.\x22)\x0a         \
           Layou\
t.fillWidth: tru\
e\x0a              \
      wrapMode: \
Text.WordWrap\x0a  \
                \
  Layout.alignme\
nt: Qt.AlignTop\x0a\
                \
}\x0a            }\x0a\
        }\x0a\x0a     \
   ColumnLayout \
{\x0a            La\
bel {\x0a          \
      id: option\
s\x0a              \
  text: qsTr(\x22Op\
tions:\x22)\x0a       \
         font.bo\
ld: true\x0a       \
         Layout.\
fillWidth: true\x0a\
            }\x0a  \
          ItemDe\
legate {\x0a       \
         id: but\
tonAutoConf\x0a    \
            text\
: qsTr(\x22Configur\
e automatically\x22\
)\x0a              \
  Layout.fillWid\
th: true\x0a       \
     }\x0a         \
   ItemDelegate \
{\x0a              \
  id: buttonBack\
upConf\x0a         \
       text: qsT\
r(\x22Restore backu\
p\x22)\x0a            \
    Layout.fillW\
idth: true\x0a\x0a    \
        }\x0a      \
  }\x0a    }\x0a}\x0a\
"

qt_resource_name = b"\
\x00\x15\
\x0a*\xd3|\
\x00S\
\x00u\x00b\x00P\x00a\x00g\x00e\x00S\x00e\x00n\x00d\x00S\x00i\x00m\x00p\x00l\x00e\
\x00.\x00q\x00m\x00l\
\x00\x17\
\x04\xff\x0e\xbc\
\x00S\
\x00u\x00b\x00P\x00a\x00g\x00e\x00S\x00e\x00n\x00d\x00A\x00d\x00v\x00a\x00n\x00c\
\x00e\x00d\x00.\x00q\x00m\x00l\
\x00\x16\
\x04\xd0\x87\x9c\
\x00S\
\x00e\x00c\x00u\x00r\x00e\x00W\x00a\x00l\x00l\x00e\x00t\x00D\x00i\x00a\x00l\x00o\
\x00g\x00.\x00q\x00m\x00l\
\x00\x0f\
\x0d\xa7\x1c\xbc\
\x00P\
\x00a\x00g\x00e\x00H\x00i\x00s\x00t\x00o\x00r\x00y\x00.\x00q\x00m\x00l\
\x00\x17\
\x06\xf1A\xfc\
\x00W\
\x00a\x00l\x00l\x00e\x00t\x00C\x00r\x00e\x00a\x00t\x00e\x00d\x00D\x00i\x00a\x00l\
\x00o\x00g\x00.\x00q\x00m\x00l\
\x00\x14\
\x0b\xb7\xd3<\
\x00G\
\x00e\x00n\x00e\x00r\x00a\x00l\x00S\x00t\x00a\x00c\x00k\x00V\x00i\x00e\x00w\x00.\
\x00q\x00m\x00l\
\x00\x0f\
\x06\xad\x14\xfc\
\x00D\
\x00i\x00a\x00l\x00o\x00g\x00A\x00b\x00o\x00u\x00t\x00.\x00q\x00m\x00l\
\x00\x1b\
\x0c\xc6,\xdc\
\x00D\
\x00e\x00s\x00t\x00i\x00n\x00a\x00t\x00i\x00o\x00n\x00L\x00i\x00s\x00t\x00D\x00e\
\x00l\x00e\x00g\x00a\x00t\x00e\x00.\x00q\x00m\x00l\
\x00\x15\
\x01\xef;\x5c\
\x00P\
\x00a\x00s\x00s\x00w\x00o\x00r\x00d\x00R\x00e\x00q\x00u\x00e\x00s\x00t\x00e\x00r\
\x00.\x00q\x00m\x00l\
\x00\x1c\
\x08\x92\x5c\xdc\
\x00D\
\x00i\x00a\x00l\x00o\x00g\x00T\x00r\x00a\x00n\x00s\x00a\x00c\x00t\x00i\x00o\x00n\
\x00D\x00e\x00t\x00a\x00i\x00l\x00s\x00.\x00q\x00m\x00l\
\x00\x11\
\x03\xc3\x9d|\
\x00D\
\x00i\x00a\x00l\x00o\x00g\x00A\x00b\x00o\x00u\x00t\x00Q\x00t\x00.\x00q\x00m\x00l\
\
\x00\x14\
\x0b<\xc7<\
\x00G\
\x00e\x00n\x00e\x00r\x00a\x00l\x00S\x00w\x00i\x00p\x00e\x00V\x00i\x00e\x00w\x00.\
\x00q\x00m\x00l\
\x00\x19\
\x01@h\x5c\
\x00D\
\x00i\x00a\x00l\x00o\x00g\x00S\x00e\x00n\x00d\x00T\x00r\x00a\x00n\x00s\x00a\x00c\
\x00t\x00i\x00o\x00n\x00.\x00q\x00m\x00l\
\x00\x10\
\x0f\xbez\xbc\
\x00N\
\x00u\x00m\x00P\x00a\x00d\x00D\x00i\x00a\x00l\x00o\x00g\x00.\x00q\x00m\x00l\
\x00\x16\
\x02\xc6\x04\xbc\
\x00W\
\x00a\x00l\x00l\x00e\x00t\x00L\x00i\x00s\x00t\x00D\x00e\x00l\x00e\x00g\x00a\x00t\
\x00e\x00.\x00q\x00m\x00l\
\x00$\
\x09g3\xfc\
\x00H\
\x00i\x00s\x00t\x00o\x00r\x00y\x00F\x00i\x00l\x00t\x00e\x00r\x00L\x00i\x00s\x00t\
\x00A\x00d\x00d\x00r\x00e\x00s\x00s\x00D\x00e\x00l\x00e\x00g\x00a\x00t\x00e\x00.\
\x00q\x00m\x00l\
\x00\x1c\
\x01\x9c\x8b\x1c\
\x00R\
\x00e\x00s\x00t\x00o\x00r\x00e\x00B\x00a\x00c\x00k\x00u\x00p\x00W\x00o\x00r\x00d\
\x00s\x00D\x00i\x00a\x00l\x00o\x00g\x00.\x00q\x00m\x00l\
\x00\x08\
\x08\x01Z\x5c\
\x00m\
\x00a\x00i\x00n\x00.\x00q\x00m\x00l\
\x00\x1d\
\x02\xfe+\x9c\
\x00H\
\x00i\x00s\x00t\x00o\x00r\x00y\x00F\x00i\x00l\x00t\x00e\x00r\x00L\x00i\x00s\x00t\
\x00D\x00e\x00l\x00e\x00g\x00a\x00t\x00e\x00.\x00q\x00m\x00l\
\x00\x0e\
\x01uy\xbc\
\x00B\
\x00l\x00o\x00c\x00k\x00c\x00h\x00a\x00i\x00n\x00.\x00q\x00m\x00l\
\x00\x15\
\x0a\x0a\xd9<\
\x00H\
\x00i\x00s\x00t\x00o\x00r\x00y\x00F\x00i\x00l\x00t\x00e\x00r\x00L\x00i\x00s\x00t\
\x00.\x00q\x00m\x00l\
\x00\x14\
\x0a\xd6u\xdc\
\x00C\
\x00r\x00e\x00a\x00t\x00e\x00L\x00o\x00a\x00d\x00W\x00a\x00l\x00l\x00e\x00t\x00.\
\x00q\x00m\x00l\
\x00\x1d\
\x08\x1aW\xfc\
\x00W\
\x00a\x00l\x00l\x00e\x00t\x00L\x00i\x00s\x00t\x00A\x00d\x00d\x00r\x00e\x00s\x00s\
\x00D\x00e\x00l\x00e\x00g\x00a\x00t\x00e\x00.\x00q\x00m\x00l\
\x00\x17\
\x01\xf0\x98\x9c\
\x00H\
\x00i\x00s\x00t\x00o\x00r\x00y\x00L\x00i\x00s\x00t\x00D\x00e\x00l\x00e\x00g\x00a\
\x00t\x00e\x00.\x00q\x00m\x00l\
\x00\x17\
\x0e\xd0b\xfc\
\x00I\
\x00n\x00p\x00u\x00t\x00O\x00u\x00t\x00p\x00u\x00t\x00D\x00e\x00l\x00e\x00g\x00a\
\x00t\x00e\x00.\x00q\x00m\x00l\
\x00\x0c\
\x02\x88\xe5\x1c\
\x00P\
\x00a\x00g\x00e\x00S\x00e\x00n\x00d\x00.\x00q\x00m\x00l\
\x00\x0f\
\x07\xf1\x8b\x9c\
\x00P\
\x00a\x00g\x00e\x00W\x00a\x00l\x00l\x00e\x00t\x00s\x00.\x00q\x00m\x00l\
\x00\x0d\
\x04V\x9b\xdc\
\x00B\
\x00u\x00s\x00y\x00P\x00o\x00p\x00u\x00p\x00.\x00q\x00m\x00l\
\x00\x16\
\x0e\x14\xb4\x9c\
\x00T\
\x00r\x00a\x00n\x00s\x00a\x00c\x00t\x00i\x00o\x00n\x00D\x00e\x00t\x00a\x00i\x00l\
\x00s\x00.\x00q\x00m\x00l\
\x00\x17\
\x0b\x0f\x9f\x5c\
\x00C\
\x00o\x00n\x00t\x00r\x00o\x00l\x00C\x00u\x00s\x00t\x00o\x00m\x00S\x00w\x00i\x00t\
\x00c\x00h\x00.\x00q\x00m\x00l\
\x00\x0d\
\x0aq\x0b\xfc\
\x00M\
\x00s\x00g\x00D\x00i\x00a\x00l\x00o\x00g\x00.\x00q\x00m\x00l\
\x00\x17\
\x05vQ\xfc\
\x00C\
\x00o\x00n\x00t\x00r\x00o\x00l\x00G\x00e\x00n\x00e\x00r\x00a\x00t\x00e\x00S\x00e\
\x00e\x00d\x00.\x00q\x00m\x00l\
\x00\x1c\
\x007\xda\x5c\
\x00D\
\x00i\x00a\x00l\x00o\x00g\x00U\x00n\x00c\x00o\x00n\x00f\x00i\x00g\x00u\x00r\x00e\
\x00d\x00W\x00a\x00l\x00l\x00e\x00t\x00.\x00q\x00m\x00l\
"

qt_resource_struct = b"\
\x00\x00\x00\x00\x00\x02\x00\x00\x00!\x00\x00\x00\x01\
\x00\x00\x05\xf6\x00\x00\x00\x00\x00\x01\x00\x01\x8a\xe6\
\x00\x00\x02@\x00\x00\x00\x00\x00\x01\x00\x00m\x83\
\x00\x00\x03\xb2\x00\x00\x00\x00\x00\x01\x00\x00\xe2\xb4\
\x00\x00\x03\x1e\x00\x00\x00\x00\x00\x01\x00\x00\xca\xe1\
\x00\x00\x01|\x00\x00\x00\x00\x00\x01\x00\x00\x5cP\
\x00\x00\x04r\x00\x00\x00\x00\x00\x01\x00\x01\x18\x90\
\x00\x00\x04\xda\x00\x00\x00\x00\x00\x01\x00\x01,-\
\x00\x00\x02\x9e\x00\x00\x00\x00\x00\x01\x00\x00\xb9\x13\
\x00\x00\x03r\x00\x00\x00\x00\x00\x01\x00\x00\xd6\x82\
\x00\x00\x01\xea\x00\x00\x00\x00\x00\x01\x00\x00d,\
\x00\x00\x05\x1c\x00\x00\x00\x00\x00\x01\x00\x01DO\
\x00\x00\x00d\x00\x00\x00\x00\x00\x01\x00\x00\x1c\xa9\
\x00\x00\x000\x00\x00\x00\x00\x00\x01\x00\x00\x07V\
\x00\x00\x05\xc2\x00\x00\x00\x00\x00\x01\x00\x01\x81}\
\x00\x00\x01\x1c\x00\x00\x00\x00\x00\x01\x00\x00M\xaf\
\x00\x00\x00\xba\x00\x00\x00\x00\x00\x01\x00\x00A\xfe\
\x00\x00\x04\xf8\x00\x00\x00\x00\x00\x01\x00\x015R\
\x00\x00\x03\x5c\x00\x00\x00\x00\x00\x01\x00\x00\xd2\x1a\
\x00\x00\x042\x00\x00\x00\x00\x00\x01\x00\x01\x04\x22\
\x00\x00\x01\xac\x00\x00\x00\x00\x00\x01\x00\x00`P\
\x00\x00\x02\xd0\x00\x00\x00\x00\x00\x01\x00\x00\xc9\x07\
\x00\x00\x03\xd4\x00\x00\x00\x00\x00\x01\x00\x00\xf4\x99\
\x00\x00\x00\x00\x00\x00\x00\x00\x00\x01\x00\x00\x00\x00\
\x00\x00\x05\xa2\x00\x00\x00\x00\x00\x01\x00\x01z\x0c\
\x00\x00\x04\x04\x00\x00\x00\x00\x00\x01\x00\x00\xfae\
\x00\x00\x05n\x00\x00\x00\x00\x00\x01\x00\x01n\xa0\
\x00\x00\x02\x12\x00\x00\x00\x00\x00\x01\x00\x00kJ\
\x00\x00\x00\xee\x00\x00\x00\x00\x00\x01\x00\x00I\x8c\
\x00\x00\x01@\x00\x00\x00\x00\x00\x01\x00\x00RT\
\x00\x00\x00\x96\x00\x00\x00\x00\x00\x01\x00\x00.*\
\x00\x00\x05<\x00\x00\x00\x00\x00\x01\x00\x01Nq\
\x00\x00\x04\xa6\x00\x00\x00\x00\x00\x01\x00\x01%f\
\x00\x00\x02x\x00\x00\x00\x00\x00\x01\x00\x00tA\
"

def qInitResources():
    QtCore.qRegisterResourceData(0x01, qt_resource_struct, qt_resource_name, qt_resource_data)

def qCleanupResources():
    QtCore.qUnregisterResourceData(0x01, qt_resource_struct, qt_resource_name, qt_resource_data)

qInitResources()
