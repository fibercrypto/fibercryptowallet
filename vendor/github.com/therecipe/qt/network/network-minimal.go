// +build minimal

package network

//#include <stdint.h>
//#include <stdlib.h>
//#include <string.h>
//#include "network-minimal.h"
import "C"
import (
	"strings"
	"unsafe"
)

func cGoUnpackString(s C.struct_QtNetwork_PackedString) string {
	if int(s.len) == -1 {
		return C.GoString(s.data)
	}
	return C.GoStringN(s.data, C.int(s.len))
}
func cGoUnpackBytes(s C.struct_QtNetwork_PackedString) []byte {
	if int(s.len) == -1 {
		gs := C.GoString(s.data)
		return *(*[]byte)(unsafe.Pointer(&gs))
	}
	return C.GoBytes(unsafe.Pointer(s.data), C.int(s.len))
}
func unpackStringList(s string) []string {
	if len(s) == 0 {
		return make([]string, 0)
	}
	return strings.Split(s, "¡¦!")
}

//go:generate stringer -type=QAbstractSocket__SocketType
//QAbstractSocket::SocketType
type QAbstractSocket__SocketType int64

const (
	QAbstractSocket__TcpSocket         QAbstractSocket__SocketType = QAbstractSocket__SocketType(0)
	QAbstractSocket__UdpSocket         QAbstractSocket__SocketType = QAbstractSocket__SocketType(1)
	QAbstractSocket__SctpSocket        QAbstractSocket__SocketType = QAbstractSocket__SocketType(2)
	QAbstractSocket__UnknownSocketType QAbstractSocket__SocketType = QAbstractSocket__SocketType(-1)
)

//go:generate stringer -type=QAbstractSocket__NetworkLayerProtocol
//QAbstractSocket::NetworkLayerProtocol
type QAbstractSocket__NetworkLayerProtocol int64

const (
	QAbstractSocket__IPv4Protocol                QAbstractSocket__NetworkLayerProtocol = QAbstractSocket__NetworkLayerProtocol(0)
	QAbstractSocket__IPv6Protocol                QAbstractSocket__NetworkLayerProtocol = QAbstractSocket__NetworkLayerProtocol(1)
	QAbstractSocket__AnyIPProtocol               QAbstractSocket__NetworkLayerProtocol = QAbstractSocket__NetworkLayerProtocol(2)
	QAbstractSocket__UnknownNetworkLayerProtocol QAbstractSocket__NetworkLayerProtocol = QAbstractSocket__NetworkLayerProtocol(-1)
)

//go:generate stringer -type=QAbstractSocket__SocketError
//QAbstractSocket::SocketError
type QAbstractSocket__SocketError int64

const (
	QAbstractSocket__ConnectionRefusedError           QAbstractSocket__SocketError = QAbstractSocket__SocketError(0)
	QAbstractSocket__RemoteHostClosedError            QAbstractSocket__SocketError = QAbstractSocket__SocketError(1)
	QAbstractSocket__HostNotFoundError                QAbstractSocket__SocketError = QAbstractSocket__SocketError(2)
	QAbstractSocket__SocketAccessError                QAbstractSocket__SocketError = QAbstractSocket__SocketError(3)
	QAbstractSocket__SocketResourceError              QAbstractSocket__SocketError = QAbstractSocket__SocketError(4)
	QAbstractSocket__SocketTimeoutError               QAbstractSocket__SocketError = QAbstractSocket__SocketError(5)
	QAbstractSocket__DatagramTooLargeError            QAbstractSocket__SocketError = QAbstractSocket__SocketError(6)
	QAbstractSocket__NetworkError                     QAbstractSocket__SocketError = QAbstractSocket__SocketError(7)
	QAbstractSocket__AddressInUseError                QAbstractSocket__SocketError = QAbstractSocket__SocketError(8)
	QAbstractSocket__SocketAddressNotAvailableError   QAbstractSocket__SocketError = QAbstractSocket__SocketError(9)
	QAbstractSocket__UnsupportedSocketOperationError  QAbstractSocket__SocketError = QAbstractSocket__SocketError(10)
	QAbstractSocket__UnfinishedSocketOperationError   QAbstractSocket__SocketError = QAbstractSocket__SocketError(11)
	QAbstractSocket__ProxyAuthenticationRequiredError QAbstractSocket__SocketError = QAbstractSocket__SocketError(12)
	QAbstractSocket__SslHandshakeFailedError          QAbstractSocket__SocketError = QAbstractSocket__SocketError(13)
	QAbstractSocket__ProxyConnectionRefusedError      QAbstractSocket__SocketError = QAbstractSocket__SocketError(14)
	QAbstractSocket__ProxyConnectionClosedError       QAbstractSocket__SocketError = QAbstractSocket__SocketError(15)
	QAbstractSocket__ProxyConnectionTimeoutError      QAbstractSocket__SocketError = QAbstractSocket__SocketError(16)
	QAbstractSocket__ProxyNotFoundError               QAbstractSocket__SocketError = QAbstractSocket__SocketError(17)
	QAbstractSocket__ProxyProtocolError               QAbstractSocket__SocketError = QAbstractSocket__SocketError(18)
	QAbstractSocket__OperationError                   QAbstractSocket__SocketError = QAbstractSocket__SocketError(19)
	QAbstractSocket__SslInternalError                 QAbstractSocket__SocketError = QAbstractSocket__SocketError(20)
	QAbstractSocket__SslInvalidUserDataError          QAbstractSocket__SocketError = QAbstractSocket__SocketError(21)
	QAbstractSocket__TemporaryError                   QAbstractSocket__SocketError = QAbstractSocket__SocketError(22)
	QAbstractSocket__UnknownSocketError               QAbstractSocket__SocketError = QAbstractSocket__SocketError(-1)
)

//go:generate stringer -type=QAbstractSocket__SocketState
//QAbstractSocket::SocketState
type QAbstractSocket__SocketState int64

const (
	QAbstractSocket__UnconnectedState QAbstractSocket__SocketState = QAbstractSocket__SocketState(0)
	QAbstractSocket__HostLookupState  QAbstractSocket__SocketState = QAbstractSocket__SocketState(1)
	QAbstractSocket__ConnectingState  QAbstractSocket__SocketState = QAbstractSocket__SocketState(2)
	QAbstractSocket__ConnectedState   QAbstractSocket__SocketState = QAbstractSocket__SocketState(3)
	QAbstractSocket__BoundState       QAbstractSocket__SocketState = QAbstractSocket__SocketState(4)
	QAbstractSocket__ListeningState   QAbstractSocket__SocketState = QAbstractSocket__SocketState(5)
	QAbstractSocket__ClosingState     QAbstractSocket__SocketState = QAbstractSocket__SocketState(6)
)

//go:generate stringer -type=QAbstractSocket__SocketOption
//QAbstractSocket::SocketOption
type QAbstractSocket__SocketOption int64

const (
	QAbstractSocket__LowDelayOption                QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(0)
	QAbstractSocket__KeepAliveOption               QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(1)
	QAbstractSocket__MulticastTtlOption            QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(2)
	QAbstractSocket__MulticastLoopbackOption       QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(3)
	QAbstractSocket__TypeOfServiceOption           QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(4)
	QAbstractSocket__SendBufferSizeSocketOption    QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(5)
	QAbstractSocket__ReceiveBufferSizeSocketOption QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(6)
	QAbstractSocket__PathMtuSocketOption           QAbstractSocket__SocketOption = QAbstractSocket__SocketOption(7)
)

//go:generate stringer -type=QAbstractSocket__BindFlag
//QAbstractSocket::BindFlag
type QAbstractSocket__BindFlag int64

const (
	QAbstractSocket__DefaultForPlatform QAbstractSocket__BindFlag = QAbstractSocket__BindFlag(0x0)
	QAbstractSocket__ShareAddress       QAbstractSocket__BindFlag = QAbstractSocket__BindFlag(0x1)
	QAbstractSocket__DontShareAddress   QAbstractSocket__BindFlag = QAbstractSocket__BindFlag(0x2)
	QAbstractSocket__ReuseAddressHint   QAbstractSocket__BindFlag = QAbstractSocket__BindFlag(0x4)
)

//go:generate stringer -type=QAbstractSocket__PauseMode
//QAbstractSocket::PauseMode
type QAbstractSocket__PauseMode int64

const (
	QAbstractSocket__PauseNever       QAbstractSocket__PauseMode = QAbstractSocket__PauseMode(0x0)
	QAbstractSocket__PauseOnSslErrors QAbstractSocket__PauseMode = QAbstractSocket__PauseMode(0x1)
)

//go:generate stringer -type=QDnsLookup__Error
//QDnsLookup::Error
type QDnsLookup__Error int64

const (
	QDnsLookup__NoError                 QDnsLookup__Error = QDnsLookup__Error(0)
	QDnsLookup__ResolverError           QDnsLookup__Error = QDnsLookup__Error(1)
	QDnsLookup__OperationCancelledError QDnsLookup__Error = QDnsLookup__Error(2)
	QDnsLookup__InvalidRequestError     QDnsLookup__Error = QDnsLookup__Error(3)
	QDnsLookup__InvalidReplyError       QDnsLookup__Error = QDnsLookup__Error(4)
	QDnsLookup__ServerFailureError      QDnsLookup__Error = QDnsLookup__Error(5)
	QDnsLookup__ServerRefusedError      QDnsLookup__Error = QDnsLookup__Error(6)
	QDnsLookup__NotFoundError           QDnsLookup__Error = QDnsLookup__Error(7)
)

//go:generate stringer -type=QDnsLookup__Type
//QDnsLookup::Type
type QDnsLookup__Type int64

const (
	QDnsLookup__A     QDnsLookup__Type = QDnsLookup__Type(1)
	QDnsLookup__AAAA  QDnsLookup__Type = QDnsLookup__Type(28)
	QDnsLookup__ANY   QDnsLookup__Type = QDnsLookup__Type(255)
	QDnsLookup__CNAME QDnsLookup__Type = QDnsLookup__Type(5)
	QDnsLookup__MX    QDnsLookup__Type = QDnsLookup__Type(15)
	QDnsLookup__NS    QDnsLookup__Type = QDnsLookup__Type(2)
	QDnsLookup__PTR   QDnsLookup__Type = QDnsLookup__Type(12)
	QDnsLookup__SRV   QDnsLookup__Type = QDnsLookup__Type(33)
	QDnsLookup__TXT   QDnsLookup__Type = QDnsLookup__Type(16)
)

//go:generate stringer -type=QDtls__QDtlsError
//QDtls::QDtlsError
type QDtls__QDtlsError int64

const (
	QDtls__NoError                     QDtls__QDtlsError = QDtls__QDtlsError(0)
	QDtls__InvalidInputParameters      QDtls__QDtlsError = QDtls__QDtlsError(1)
	QDtls__InvalidOperation            QDtls__QDtlsError = QDtls__QDtlsError(2)
	QDtls__UnderlyingSocketError       QDtls__QDtlsError = QDtls__QDtlsError(3)
	QDtls__RemoteClosedConnectionError QDtls__QDtlsError = QDtls__QDtlsError(4)
	QDtls__PeerVerificationError       QDtls__QDtlsError = QDtls__QDtlsError(5)
	QDtls__TlsInitializationError      QDtls__QDtlsError = QDtls__QDtlsError(6)
	QDtls__TlsFatalError               QDtls__QDtlsError = QDtls__QDtlsError(7)
	QDtls__TlsNonFatalError            QDtls__QDtlsError = QDtls__QDtlsError(8)
)

//go:generate stringer -type=QDtls__HandshakeState
//QDtls::HandshakeState
type QDtls__HandshakeState int64

const (
	QDtls__HandshakeNotStarted    QDtls__HandshakeState = QDtls__HandshakeState(0)
	QDtls__HandshakeInProgress    QDtls__HandshakeState = QDtls__HandshakeState(1)
	QDtls__PeerVerificationFailed QDtls__HandshakeState = QDtls__HandshakeState(2)
	QDtls__HandshakeComplete      QDtls__HandshakeState = QDtls__HandshakeState(3)
)

//go:generate stringer -type=QHostAddress__SpecialAddress
//QHostAddress::SpecialAddress
type QHostAddress__SpecialAddress int64

const (
	QHostAddress__Null          QHostAddress__SpecialAddress = QHostAddress__SpecialAddress(0)
	QHostAddress__Broadcast     QHostAddress__SpecialAddress = QHostAddress__SpecialAddress(1)
	QHostAddress__LocalHost     QHostAddress__SpecialAddress = QHostAddress__SpecialAddress(2)
	QHostAddress__LocalHostIPv6 QHostAddress__SpecialAddress = QHostAddress__SpecialAddress(3)
	QHostAddress__Any           QHostAddress__SpecialAddress = QHostAddress__SpecialAddress(4)
	QHostAddress__AnyIPv6       QHostAddress__SpecialAddress = QHostAddress__SpecialAddress(5)
	QHostAddress__AnyIPv4       QHostAddress__SpecialAddress = QHostAddress__SpecialAddress(6)
)

//go:generate stringer -type=QHostAddress__ConversionModeFlag
//QHostAddress::ConversionModeFlag
type QHostAddress__ConversionModeFlag int64

const (
	QHostAddress__ConvertV4MappedToIPv4     QHostAddress__ConversionModeFlag = QHostAddress__ConversionModeFlag(1)
	QHostAddress__ConvertV4CompatToIPv4     QHostAddress__ConversionModeFlag = QHostAddress__ConversionModeFlag(2)
	QHostAddress__ConvertUnspecifiedAddress QHostAddress__ConversionModeFlag = QHostAddress__ConversionModeFlag(4)
	QHostAddress__ConvertLocalHost          QHostAddress__ConversionModeFlag = QHostAddress__ConversionModeFlag(8)
	QHostAddress__TolerantConversion        QHostAddress__ConversionModeFlag = QHostAddress__ConversionModeFlag(0xff)
	QHostAddress__StrictConversion          QHostAddress__ConversionModeFlag = QHostAddress__ConversionModeFlag(0)
)

//go:generate stringer -type=QHostInfo__HostInfoError
//QHostInfo::HostInfoError
type QHostInfo__HostInfoError int64

const (
	QHostInfo__NoError      QHostInfo__HostInfoError = QHostInfo__HostInfoError(0)
	QHostInfo__HostNotFound QHostInfo__HostInfoError = QHostInfo__HostInfoError(1)
	QHostInfo__UnknownError QHostInfo__HostInfoError = QHostInfo__HostInfoError(2)
)

//go:generate stringer -type=QHstsPolicy__PolicyFlag
//QHstsPolicy::PolicyFlag
type QHstsPolicy__PolicyFlag int64

const (
	QHstsPolicy__IncludeSubDomains QHstsPolicy__PolicyFlag = QHstsPolicy__PolicyFlag(1)
)

//go:generate stringer -type=QHttpMultiPart__ContentType
//QHttpMultiPart::ContentType
type QHttpMultiPart__ContentType int64

const (
	QHttpMultiPart__MixedType       QHttpMultiPart__ContentType = QHttpMultiPart__ContentType(0)
	QHttpMultiPart__RelatedType     QHttpMultiPart__ContentType = QHttpMultiPart__ContentType(1)
	QHttpMultiPart__FormDataType    QHttpMultiPart__ContentType = QHttpMultiPart__ContentType(2)
	QHttpMultiPart__AlternativeType QHttpMultiPart__ContentType = QHttpMultiPart__ContentType(3)
)

//go:generate stringer -type=QLocalServer__SocketOption
//QLocalServer::SocketOption
type QLocalServer__SocketOption int64

const (
	QLocalServer__NoOptions         QLocalServer__SocketOption = QLocalServer__SocketOption(0x0)
	QLocalServer__UserAccessOption  QLocalServer__SocketOption = QLocalServer__SocketOption(0x01)
	QLocalServer__GroupAccessOption QLocalServer__SocketOption = QLocalServer__SocketOption(0x2)
	QLocalServer__OtherAccessOption QLocalServer__SocketOption = QLocalServer__SocketOption(0x4)
	QLocalServer__WorldAccessOption QLocalServer__SocketOption = QLocalServer__SocketOption(0x7)
)

//go:generate stringer -type=QLocalSocket__LocalSocketError
//QLocalSocket::LocalSocketError
type QLocalSocket__LocalSocketError int64

const (
	QLocalSocket__ConnectionRefusedError          QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__ConnectionRefusedError)
	QLocalSocket__PeerClosedError                 QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__RemoteHostClosedError)
	QLocalSocket__ServerNotFoundError             QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__HostNotFoundError)
	QLocalSocket__SocketAccessError               QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__SocketAccessError)
	QLocalSocket__SocketResourceError             QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__SocketResourceError)
	QLocalSocket__SocketTimeoutError              QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__SocketTimeoutError)
	QLocalSocket__DatagramTooLargeError           QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__DatagramTooLargeError)
	QLocalSocket__ConnectionError                 QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__NetworkError)
	QLocalSocket__UnsupportedSocketOperationError QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__UnsupportedSocketOperationError)
	QLocalSocket__UnknownSocketError              QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__UnknownSocketError)
	QLocalSocket__OperationError                  QLocalSocket__LocalSocketError = QLocalSocket__LocalSocketError(QAbstractSocket__OperationError)
)

//go:generate stringer -type=QLocalSocket__LocalSocketState
//QLocalSocket::LocalSocketState
type QLocalSocket__LocalSocketState int64

const (
	QLocalSocket__UnconnectedState QLocalSocket__LocalSocketState = QLocalSocket__LocalSocketState(QAbstractSocket__UnconnectedState)
	QLocalSocket__ConnectingState  QLocalSocket__LocalSocketState = QLocalSocket__LocalSocketState(QAbstractSocket__ConnectingState)
	QLocalSocket__ConnectedState   QLocalSocket__LocalSocketState = QLocalSocket__LocalSocketState(QAbstractSocket__ConnectedState)
	QLocalSocket__ClosingState     QLocalSocket__LocalSocketState = QLocalSocket__LocalSocketState(QAbstractSocket__ClosingState)
)

//go:generate stringer -type=QNetworkAccessManager__Operation
//QNetworkAccessManager::Operation
type QNetworkAccessManager__Operation int64

const (
	QNetworkAccessManager__HeadOperation    QNetworkAccessManager__Operation = QNetworkAccessManager__Operation(1)
	QNetworkAccessManager__GetOperation     QNetworkAccessManager__Operation = QNetworkAccessManager__Operation(2)
	QNetworkAccessManager__PutOperation     QNetworkAccessManager__Operation = QNetworkAccessManager__Operation(3)
	QNetworkAccessManager__PostOperation    QNetworkAccessManager__Operation = QNetworkAccessManager__Operation(4)
	QNetworkAccessManager__DeleteOperation  QNetworkAccessManager__Operation = QNetworkAccessManager__Operation(5)
	QNetworkAccessManager__CustomOperation  QNetworkAccessManager__Operation = QNetworkAccessManager__Operation(6)
	QNetworkAccessManager__UnknownOperation QNetworkAccessManager__Operation = QNetworkAccessManager__Operation(0)
)

//go:generate stringer -type=QNetworkAccessManager__NetworkAccessibility
//QNetworkAccessManager::NetworkAccessibility
type QNetworkAccessManager__NetworkAccessibility int64

const (
	QNetworkAccessManager__UnknownAccessibility QNetworkAccessManager__NetworkAccessibility = QNetworkAccessManager__NetworkAccessibility(-1)
	QNetworkAccessManager__NotAccessible        QNetworkAccessManager__NetworkAccessibility = QNetworkAccessManager__NetworkAccessibility(0)
	QNetworkAccessManager__Accessible           QNetworkAccessManager__NetworkAccessibility = QNetworkAccessManager__NetworkAccessibility(1)
)

//go:generate stringer -type=QNetworkAddressEntry__DnsEligibilityStatus
//QNetworkAddressEntry::DnsEligibilityStatus
type QNetworkAddressEntry__DnsEligibilityStatus int64

const (
	QNetworkAddressEntry__DnsEligibilityUnknown QNetworkAddressEntry__DnsEligibilityStatus = QNetworkAddressEntry__DnsEligibilityStatus(-1)
	QNetworkAddressEntry__DnsIneligible         QNetworkAddressEntry__DnsEligibilityStatus = QNetworkAddressEntry__DnsEligibilityStatus(0)
	QNetworkAddressEntry__DnsEligible           QNetworkAddressEntry__DnsEligibilityStatus = QNetworkAddressEntry__DnsEligibilityStatus(1)
)

//go:generate stringer -type=QNetworkConfiguration__Type
//QNetworkConfiguration::Type
type QNetworkConfiguration__Type int64

const (
	QNetworkConfiguration__InternetAccessPoint QNetworkConfiguration__Type = QNetworkConfiguration__Type(0)
	QNetworkConfiguration__ServiceNetwork      QNetworkConfiguration__Type = QNetworkConfiguration__Type(1)
	QNetworkConfiguration__UserChoice          QNetworkConfiguration__Type = QNetworkConfiguration__Type(2)
	QNetworkConfiguration__Invalid             QNetworkConfiguration__Type = QNetworkConfiguration__Type(3)
)

//go:generate stringer -type=QNetworkConfiguration__Purpose
//QNetworkConfiguration::Purpose
type QNetworkConfiguration__Purpose int64

const (
	QNetworkConfiguration__UnknownPurpose         QNetworkConfiguration__Purpose = QNetworkConfiguration__Purpose(0)
	QNetworkConfiguration__PublicPurpose          QNetworkConfiguration__Purpose = QNetworkConfiguration__Purpose(1)
	QNetworkConfiguration__PrivatePurpose         QNetworkConfiguration__Purpose = QNetworkConfiguration__Purpose(2)
	QNetworkConfiguration__ServiceSpecificPurpose QNetworkConfiguration__Purpose = QNetworkConfiguration__Purpose(3)
)

//go:generate stringer -type=QNetworkConfiguration__StateFlag
//QNetworkConfiguration::StateFlag
type QNetworkConfiguration__StateFlag int64

const (
	QNetworkConfiguration__Undefined  QNetworkConfiguration__StateFlag = QNetworkConfiguration__StateFlag(0x0000001)
	QNetworkConfiguration__Defined    QNetworkConfiguration__StateFlag = QNetworkConfiguration__StateFlag(0x0000002)
	QNetworkConfiguration__Discovered QNetworkConfiguration__StateFlag = QNetworkConfiguration__StateFlag(0x0000006)
	QNetworkConfiguration__Active     QNetworkConfiguration__StateFlag = QNetworkConfiguration__StateFlag(0x000000e)
)

//go:generate stringer -type=QNetworkConfiguration__BearerType
//QNetworkConfiguration::BearerType
type QNetworkConfiguration__BearerType int64

const (
	QNetworkConfiguration__BearerUnknown   QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(0)
	QNetworkConfiguration__BearerEthernet  QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(1)
	QNetworkConfiguration__BearerWLAN      QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(2)
	QNetworkConfiguration__Bearer2G        QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(3)
	QNetworkConfiguration__BearerCDMA2000  QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(4)
	QNetworkConfiguration__BearerWCDMA     QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(5)
	QNetworkConfiguration__BearerHSPA      QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(6)
	QNetworkConfiguration__BearerBluetooth QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(7)
	QNetworkConfiguration__BearerWiMAX     QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(8)
	QNetworkConfiguration__BearerEVDO      QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(9)
	QNetworkConfiguration__BearerLTE       QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(10)
	QNetworkConfiguration__Bearer3G        QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(11)
	QNetworkConfiguration__Bearer4G        QNetworkConfiguration__BearerType = QNetworkConfiguration__BearerType(12)
)

//go:generate stringer -type=QNetworkConfigurationManager__Capability
//QNetworkConfigurationManager::Capability
type QNetworkConfigurationManager__Capability int64

const (
	QNetworkConfigurationManager__CanStartAndStopInterfaces QNetworkConfigurationManager__Capability = QNetworkConfigurationManager__Capability(0x00000001)
	QNetworkConfigurationManager__DirectConnectionRouting   QNetworkConfigurationManager__Capability = QNetworkConfigurationManager__Capability(0x00000002)
	QNetworkConfigurationManager__SystemSessionSupport      QNetworkConfigurationManager__Capability = QNetworkConfigurationManager__Capability(0x00000004)
	QNetworkConfigurationManager__ApplicationLevelRoaming   QNetworkConfigurationManager__Capability = QNetworkConfigurationManager__Capability(0x00000008)
	QNetworkConfigurationManager__ForcedRoaming             QNetworkConfigurationManager__Capability = QNetworkConfigurationManager__Capability(0x00000010)
	QNetworkConfigurationManager__DataStatistics            QNetworkConfigurationManager__Capability = QNetworkConfigurationManager__Capability(0x00000020)
	QNetworkConfigurationManager__NetworkSessionRequired    QNetworkConfigurationManager__Capability = QNetworkConfigurationManager__Capability(0x00000040)
)

//go:generate stringer -type=QNetworkCookie__RawForm
//QNetworkCookie::RawForm
type QNetworkCookie__RawForm int64

const (
	QNetworkCookie__NameAndValueOnly QNetworkCookie__RawForm = QNetworkCookie__RawForm(0)
	QNetworkCookie__Full             QNetworkCookie__RawForm = QNetworkCookie__RawForm(1)
)

//go:generate stringer -type=QNetworkInterface__InterfaceType
//QNetworkInterface::InterfaceType
type QNetworkInterface__InterfaceType int64

const (
	QNetworkInterface__Loopback   QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(1)
	QNetworkInterface__Virtual    QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(2)
	QNetworkInterface__Ethernet   QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(3)
	QNetworkInterface__Slip       QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(4)
	QNetworkInterface__CanBus     QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(5)
	QNetworkInterface__Ppp        QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(6)
	QNetworkInterface__Fddi       QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(7)
	QNetworkInterface__Wifi       QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(8)
	QNetworkInterface__Ieee80211  QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(QNetworkInterface__Wifi)
	QNetworkInterface__Phonet     QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(9)
	QNetworkInterface__Ieee802154 QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(10)
	QNetworkInterface__SixLoWPAN  QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(11)
	QNetworkInterface__Ieee80216  QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(12)
	QNetworkInterface__Ieee1394   QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(13)
	QNetworkInterface__Unknown    QNetworkInterface__InterfaceType = QNetworkInterface__InterfaceType(0)
)

//go:generate stringer -type=QNetworkInterface__InterfaceFlag
//QNetworkInterface::InterfaceFlag
type QNetworkInterface__InterfaceFlag int64

const (
	QNetworkInterface__IsUp           QNetworkInterface__InterfaceFlag = QNetworkInterface__InterfaceFlag(0x1)
	QNetworkInterface__IsRunning      QNetworkInterface__InterfaceFlag = QNetworkInterface__InterfaceFlag(0x2)
	QNetworkInterface__CanBroadcast   QNetworkInterface__InterfaceFlag = QNetworkInterface__InterfaceFlag(0x4)
	QNetworkInterface__IsLoopBack     QNetworkInterface__InterfaceFlag = QNetworkInterface__InterfaceFlag(0x8)
	QNetworkInterface__IsPointToPoint QNetworkInterface__InterfaceFlag = QNetworkInterface__InterfaceFlag(0x10)
	QNetworkInterface__CanMulticast   QNetworkInterface__InterfaceFlag = QNetworkInterface__InterfaceFlag(0x20)
)

//go:generate stringer -type=QNetworkProxy__ProxyType
//QNetworkProxy::ProxyType
type QNetworkProxy__ProxyType int64

const (
	QNetworkProxy__DefaultProxy     QNetworkProxy__ProxyType = QNetworkProxy__ProxyType(0)
	QNetworkProxy__Socks5Proxy      QNetworkProxy__ProxyType = QNetworkProxy__ProxyType(1)
	QNetworkProxy__NoProxy          QNetworkProxy__ProxyType = QNetworkProxy__ProxyType(2)
	QNetworkProxy__HttpProxy        QNetworkProxy__ProxyType = QNetworkProxy__ProxyType(3)
	QNetworkProxy__HttpCachingProxy QNetworkProxy__ProxyType = QNetworkProxy__ProxyType(4)
	QNetworkProxy__FtpCachingProxy  QNetworkProxy__ProxyType = QNetworkProxy__ProxyType(5)
)

//go:generate stringer -type=QNetworkProxy__Capability
//QNetworkProxy::Capability
type QNetworkProxy__Capability int64

const (
	QNetworkProxy__TunnelingCapability      QNetworkProxy__Capability = QNetworkProxy__Capability(0x0001)
	QNetworkProxy__ListeningCapability      QNetworkProxy__Capability = QNetworkProxy__Capability(0x0002)
	QNetworkProxy__UdpTunnelingCapability   QNetworkProxy__Capability = QNetworkProxy__Capability(0x0004)
	QNetworkProxy__CachingCapability        QNetworkProxy__Capability = QNetworkProxy__Capability(0x0008)
	QNetworkProxy__HostNameLookupCapability QNetworkProxy__Capability = QNetworkProxy__Capability(0x0010)
	QNetworkProxy__SctpTunnelingCapability  QNetworkProxy__Capability = QNetworkProxy__Capability(0x00020)
	QNetworkProxy__SctpListeningCapability  QNetworkProxy__Capability = QNetworkProxy__Capability(0x00040)
)

//go:generate stringer -type=QNetworkProxyQuery__QueryType
//QNetworkProxyQuery::QueryType
type QNetworkProxyQuery__QueryType int64

const (
	QNetworkProxyQuery__TcpSocket  QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(0)
	QNetworkProxyQuery__UdpSocket  QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(1)
	QNetworkProxyQuery__SctpSocket QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(2)
	QNetworkProxyQuery__TcpServer  QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(100)
	QNetworkProxyQuery__UrlRequest QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(101)
	QNetworkProxyQuery__SctpServer QNetworkProxyQuery__QueryType = QNetworkProxyQuery__QueryType(102)
)

//go:generate stringer -type=QNetworkReply__NetworkError
//QNetworkReply::NetworkError
type QNetworkReply__NetworkError int64

const (
	QNetworkReply__NoError                           QNetworkReply__NetworkError = QNetworkReply__NetworkError(0)
	QNetworkReply__ConnectionRefusedError            QNetworkReply__NetworkError = QNetworkReply__NetworkError(1)
	QNetworkReply__RemoteHostClosedError             QNetworkReply__NetworkError = QNetworkReply__NetworkError(2)
	QNetworkReply__HostNotFoundError                 QNetworkReply__NetworkError = QNetworkReply__NetworkError(3)
	QNetworkReply__TimeoutError                      QNetworkReply__NetworkError = QNetworkReply__NetworkError(4)
	QNetworkReply__OperationCanceledError            QNetworkReply__NetworkError = QNetworkReply__NetworkError(5)
	QNetworkReply__SslHandshakeFailedError           QNetworkReply__NetworkError = QNetworkReply__NetworkError(6)
	QNetworkReply__TemporaryNetworkFailureError      QNetworkReply__NetworkError = QNetworkReply__NetworkError(7)
	QNetworkReply__NetworkSessionFailedError         QNetworkReply__NetworkError = QNetworkReply__NetworkError(8)
	QNetworkReply__BackgroundRequestNotAllowedError  QNetworkReply__NetworkError = QNetworkReply__NetworkError(9)
	QNetworkReply__TooManyRedirectsError             QNetworkReply__NetworkError = QNetworkReply__NetworkError(10)
	QNetworkReply__InsecureRedirectError             QNetworkReply__NetworkError = QNetworkReply__NetworkError(11)
	QNetworkReply__UnknownNetworkError               QNetworkReply__NetworkError = QNetworkReply__NetworkError(99)
	QNetworkReply__ProxyConnectionRefusedError       QNetworkReply__NetworkError = QNetworkReply__NetworkError(101)
	QNetworkReply__ProxyConnectionClosedError        QNetworkReply__NetworkError = QNetworkReply__NetworkError(102)
	QNetworkReply__ProxyNotFoundError                QNetworkReply__NetworkError = QNetworkReply__NetworkError(103)
	QNetworkReply__ProxyTimeoutError                 QNetworkReply__NetworkError = QNetworkReply__NetworkError(104)
	QNetworkReply__ProxyAuthenticationRequiredError  QNetworkReply__NetworkError = QNetworkReply__NetworkError(105)
	QNetworkReply__UnknownProxyError                 QNetworkReply__NetworkError = QNetworkReply__NetworkError(199)
	QNetworkReply__ContentAccessDenied               QNetworkReply__NetworkError = QNetworkReply__NetworkError(201)
	QNetworkReply__ContentOperationNotPermittedError QNetworkReply__NetworkError = QNetworkReply__NetworkError(202)
	QNetworkReply__ContentNotFoundError              QNetworkReply__NetworkError = QNetworkReply__NetworkError(203)
	QNetworkReply__AuthenticationRequiredError       QNetworkReply__NetworkError = QNetworkReply__NetworkError(204)
	QNetworkReply__ContentReSendError                QNetworkReply__NetworkError = QNetworkReply__NetworkError(205)
	QNetworkReply__ContentConflictError              QNetworkReply__NetworkError = QNetworkReply__NetworkError(206)
	QNetworkReply__ContentGoneError                  QNetworkReply__NetworkError = QNetworkReply__NetworkError(207)
	QNetworkReply__UnknownContentError               QNetworkReply__NetworkError = QNetworkReply__NetworkError(299)
	QNetworkReply__ProtocolUnknownError              QNetworkReply__NetworkError = QNetworkReply__NetworkError(301)
	QNetworkReply__ProtocolInvalidOperationError     QNetworkReply__NetworkError = QNetworkReply__NetworkError(302)
	QNetworkReply__ProtocolFailure                   QNetworkReply__NetworkError = QNetworkReply__NetworkError(399)
	QNetworkReply__InternalServerError               QNetworkReply__NetworkError = QNetworkReply__NetworkError(401)
	QNetworkReply__OperationNotImplementedError      QNetworkReply__NetworkError = QNetworkReply__NetworkError(402)
	QNetworkReply__ServiceUnavailableError           QNetworkReply__NetworkError = QNetworkReply__NetworkError(403)
	QNetworkReply__UnknownServerError                QNetworkReply__NetworkError = QNetworkReply__NetworkError(499)
)

//go:generate stringer -type=QNetworkRequest__Attribute
//QNetworkRequest::Attribute
type QNetworkRequest__Attribute int64

const (
	QNetworkRequest__HttpStatusCodeAttribute               QNetworkRequest__Attribute = QNetworkRequest__Attribute(0)
	QNetworkRequest__HttpReasonPhraseAttribute             QNetworkRequest__Attribute = QNetworkRequest__Attribute(1)
	QNetworkRequest__RedirectionTargetAttribute            QNetworkRequest__Attribute = QNetworkRequest__Attribute(2)
	QNetworkRequest__ConnectionEncryptedAttribute          QNetworkRequest__Attribute = QNetworkRequest__Attribute(3)
	QNetworkRequest__CacheLoadControlAttribute             QNetworkRequest__Attribute = QNetworkRequest__Attribute(4)
	QNetworkRequest__CacheSaveControlAttribute             QNetworkRequest__Attribute = QNetworkRequest__Attribute(5)
	QNetworkRequest__SourceIsFromCacheAttribute            QNetworkRequest__Attribute = QNetworkRequest__Attribute(6)
	QNetworkRequest__DoNotBufferUploadDataAttribute        QNetworkRequest__Attribute = QNetworkRequest__Attribute(7)
	QNetworkRequest__HttpPipeliningAllowedAttribute        QNetworkRequest__Attribute = QNetworkRequest__Attribute(8)
	QNetworkRequest__HttpPipeliningWasUsedAttribute        QNetworkRequest__Attribute = QNetworkRequest__Attribute(9)
	QNetworkRequest__CustomVerbAttribute                   QNetworkRequest__Attribute = QNetworkRequest__Attribute(10)
	QNetworkRequest__CookieLoadControlAttribute            QNetworkRequest__Attribute = QNetworkRequest__Attribute(11)
	QNetworkRequest__AuthenticationReuseAttribute          QNetworkRequest__Attribute = QNetworkRequest__Attribute(12)
	QNetworkRequest__CookieSaveControlAttribute            QNetworkRequest__Attribute = QNetworkRequest__Attribute(13)
	QNetworkRequest__MaximumDownloadBufferSizeAttribute    QNetworkRequest__Attribute = QNetworkRequest__Attribute(14)
	QNetworkRequest__DownloadBufferAttribute               QNetworkRequest__Attribute = QNetworkRequest__Attribute(15)
	QNetworkRequest__SynchronousRequestAttribute           QNetworkRequest__Attribute = QNetworkRequest__Attribute(16)
	QNetworkRequest__BackgroundRequestAttribute            QNetworkRequest__Attribute = QNetworkRequest__Attribute(17)
	QNetworkRequest__SpdyAllowedAttribute                  QNetworkRequest__Attribute = QNetworkRequest__Attribute(18)
	QNetworkRequest__SpdyWasUsedAttribute                  QNetworkRequest__Attribute = QNetworkRequest__Attribute(19)
	QNetworkRequest__EmitAllUploadProgressSignalsAttribute QNetworkRequest__Attribute = QNetworkRequest__Attribute(20)
	QNetworkRequest__FollowRedirectsAttribute              QNetworkRequest__Attribute = QNetworkRequest__Attribute(21)
	QNetworkRequest__HTTP2AllowedAttribute                 QNetworkRequest__Attribute = QNetworkRequest__Attribute(22)
	QNetworkRequest__HTTP2WasUsedAttribute                 QNetworkRequest__Attribute = QNetworkRequest__Attribute(23)
	QNetworkRequest__OriginalContentLengthAttribute        QNetworkRequest__Attribute = QNetworkRequest__Attribute(24)
	QNetworkRequest__RedirectPolicyAttribute               QNetworkRequest__Attribute = QNetworkRequest__Attribute(25)
	QNetworkRequest__Http2DirectAttribute                  QNetworkRequest__Attribute = QNetworkRequest__Attribute(26)
	QNetworkRequest__ResourceTypeAttribute                 QNetworkRequest__Attribute = QNetworkRequest__Attribute(27)
	QNetworkRequest__User                                  QNetworkRequest__Attribute = QNetworkRequest__Attribute(1000)
	QNetworkRequest__UserMax                               QNetworkRequest__Attribute = QNetworkRequest__Attribute(32767)
)

//go:generate stringer -type=QNetworkRequest__KnownHeaders
//QNetworkRequest::KnownHeaders
type QNetworkRequest__KnownHeaders int64

const (
	QNetworkRequest__ContentTypeHeader        QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(0)
	QNetworkRequest__ContentLengthHeader      QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(1)
	QNetworkRequest__LocationHeader           QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(2)
	QNetworkRequest__LastModifiedHeader       QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(3)
	QNetworkRequest__CookieHeader             QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(4)
	QNetworkRequest__SetCookieHeader          QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(5)
	QNetworkRequest__ContentDispositionHeader QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(6)
	QNetworkRequest__UserAgentHeader          QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(7)
	QNetworkRequest__ServerHeader             QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(8)
	QNetworkRequest__IfModifiedSinceHeader    QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(9)
	QNetworkRequest__ETagHeader               QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(10)
	QNetworkRequest__IfMatchHeader            QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(11)
	QNetworkRequest__IfNoneMatchHeader        QNetworkRequest__KnownHeaders = QNetworkRequest__KnownHeaders(12)
)

//go:generate stringer -type=QNetworkRequest__CacheLoadControl
//QNetworkRequest::CacheLoadControl
type QNetworkRequest__CacheLoadControl int64

const (
	QNetworkRequest__AlwaysNetwork QNetworkRequest__CacheLoadControl = QNetworkRequest__CacheLoadControl(0)
	QNetworkRequest__PreferNetwork QNetworkRequest__CacheLoadControl = QNetworkRequest__CacheLoadControl(1)
	QNetworkRequest__PreferCache   QNetworkRequest__CacheLoadControl = QNetworkRequest__CacheLoadControl(2)
	QNetworkRequest__AlwaysCache   QNetworkRequest__CacheLoadControl = QNetworkRequest__CacheLoadControl(3)
)

//go:generate stringer -type=QNetworkRequest__LoadControl
//QNetworkRequest::LoadControl
type QNetworkRequest__LoadControl int64

const (
	QNetworkRequest__Automatic QNetworkRequest__LoadControl = QNetworkRequest__LoadControl(0)
	QNetworkRequest__Manual    QNetworkRequest__LoadControl = QNetworkRequest__LoadControl(1)
)

//go:generate stringer -type=QNetworkRequest__Priority
//QNetworkRequest::Priority
type QNetworkRequest__Priority int64

const (
	QNetworkRequest__HighPriority   QNetworkRequest__Priority = QNetworkRequest__Priority(1)
	QNetworkRequest__NormalPriority QNetworkRequest__Priority = QNetworkRequest__Priority(3)
	QNetworkRequest__LowPriority    QNetworkRequest__Priority = QNetworkRequest__Priority(5)
)

//go:generate stringer -type=QNetworkRequest__RedirectPolicy
//QNetworkRequest::RedirectPolicy
type QNetworkRequest__RedirectPolicy int64

const (
	QNetworkRequest__ManualRedirectPolicy       QNetworkRequest__RedirectPolicy = QNetworkRequest__RedirectPolicy(0)
	QNetworkRequest__NoLessSafeRedirectPolicy   QNetworkRequest__RedirectPolicy = QNetworkRequest__RedirectPolicy(1)
	QNetworkRequest__SameOriginRedirectPolicy   QNetworkRequest__RedirectPolicy = QNetworkRequest__RedirectPolicy(2)
	QNetworkRequest__UserVerifiedRedirectPolicy QNetworkRequest__RedirectPolicy = QNetworkRequest__RedirectPolicy(3)
)

//go:generate stringer -type=QNetworkSession__State
//QNetworkSession::State
type QNetworkSession__State int64

const (
	QNetworkSession__Invalid      QNetworkSession__State = QNetworkSession__State(0)
	QNetworkSession__NotAvailable QNetworkSession__State = QNetworkSession__State(1)
	QNetworkSession__Connecting   QNetworkSession__State = QNetworkSession__State(2)
	QNetworkSession__Connected    QNetworkSession__State = QNetworkSession__State(3)
	QNetworkSession__Closing      QNetworkSession__State = QNetworkSession__State(4)
	QNetworkSession__Disconnected QNetworkSession__State = QNetworkSession__State(5)
	QNetworkSession__Roaming      QNetworkSession__State = QNetworkSession__State(6)
)

//go:generate stringer -type=QNetworkSession__SessionError
//QNetworkSession::SessionError
type QNetworkSession__SessionError int64

const (
	QNetworkSession__UnknownSessionError        QNetworkSession__SessionError = QNetworkSession__SessionError(0)
	QNetworkSession__SessionAbortedError        QNetworkSession__SessionError = QNetworkSession__SessionError(1)
	QNetworkSession__RoamingError               QNetworkSession__SessionError = QNetworkSession__SessionError(2)
	QNetworkSession__OperationNotSupportedError QNetworkSession__SessionError = QNetworkSession__SessionError(3)
	QNetworkSession__InvalidConfigurationError  QNetworkSession__SessionError = QNetworkSession__SessionError(4)
)

//go:generate stringer -type=QNetworkSession__UsagePolicy
//QNetworkSession::UsagePolicy
type QNetworkSession__UsagePolicy int64

const (
	QNetworkSession__NoPolicy                  QNetworkSession__UsagePolicy = QNetworkSession__UsagePolicy(0)
	QNetworkSession__NoBackgroundTrafficPolicy QNetworkSession__UsagePolicy = QNetworkSession__UsagePolicy(1)
)

//go:generate stringer -type=QOcspResponse__QOcspCertificateStatus
//QOcspResponse::QOcspCertificateStatus
type QOcspResponse__QOcspCertificateStatus int64

const (
	QOcspResponse__Good    QOcspResponse__QOcspCertificateStatus = QOcspResponse__QOcspCertificateStatus(0)
	QOcspResponse__Revoked QOcspResponse__QOcspCertificateStatus = QOcspResponse__QOcspCertificateStatus(1)
	QOcspResponse__Unknown QOcspResponse__QOcspCertificateStatus = QOcspResponse__QOcspCertificateStatus(2)
)

//go:generate stringer -type=QOcspResponse__QOcspRevocationReason
//QOcspResponse::QOcspRevocationReason
type QOcspResponse__QOcspRevocationReason int64

const (
	QOcspResponse__None                 QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(-1)
	QOcspResponse__Unspecified          QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(0)
	QOcspResponse__KeyCompromise        QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(1)
	QOcspResponse__CACompromise         QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(2)
	QOcspResponse__AffiliationChanged   QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(3)
	QOcspResponse__Superseded           QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(4)
	QOcspResponse__CessationOfOperation QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(5)
	QOcspResponse__CertificateHold      QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(6)
	QOcspResponse__RemoveFromCRL        QOcspResponse__QOcspRevocationReason = QOcspResponse__QOcspRevocationReason(7)
)

//go:generate stringer -type=QSsl__KeyType
//QSsl::KeyType
type QSsl__KeyType int64

const (
	QSsl__PrivateKey QSsl__KeyType = QSsl__KeyType(0)
	QSsl__PublicKey  QSsl__KeyType = QSsl__KeyType(1)
)

//go:generate stringer -type=QSsl__EncodingFormat
//QSsl::EncodingFormat
type QSsl__EncodingFormat int64

const (
	QSsl__Pem QSsl__EncodingFormat = QSsl__EncodingFormat(0)
	QSsl__Der QSsl__EncodingFormat = QSsl__EncodingFormat(1)
)

//go:generate stringer -type=QSsl__KeyAlgorithm
//QSsl::KeyAlgorithm
type QSsl__KeyAlgorithm int64

const (
	QSsl__Opaque QSsl__KeyAlgorithm = QSsl__KeyAlgorithm(0)
	QSsl__Rsa    QSsl__KeyAlgorithm = QSsl__KeyAlgorithm(1)
	QSsl__Dsa    QSsl__KeyAlgorithm = QSsl__KeyAlgorithm(2)
	QSsl__Ec     QSsl__KeyAlgorithm = QSsl__KeyAlgorithm(3)
	QSsl__Dh     QSsl__KeyAlgorithm = QSsl__KeyAlgorithm(4)
)

//go:generate stringer -type=QSsl__AlternativeNameEntryType
//QSsl::AlternativeNameEntryType
type QSsl__AlternativeNameEntryType int64

const (
	QSsl__EmailEntry     QSsl__AlternativeNameEntryType = QSsl__AlternativeNameEntryType(0)
	QSsl__DnsEntry       QSsl__AlternativeNameEntryType = QSsl__AlternativeNameEntryType(1)
	QSsl__IpAddressEntry QSsl__AlternativeNameEntryType = QSsl__AlternativeNameEntryType(2)
)

//go:generate stringer -type=QSsl__SslProtocol
//QSsl::SslProtocol
type QSsl__SslProtocol int64

const (
	QSsl__SslV3           QSsl__SslProtocol = QSsl__SslProtocol(0)
	QSsl__SslV2           QSsl__SslProtocol = QSsl__SslProtocol(1)
	QSsl__TlsV1_0         QSsl__SslProtocol = QSsl__SslProtocol(2)
	QSsl__TlsV1           QSsl__SslProtocol = QSsl__SslProtocol(QSsl__TlsV1_0)
	QSsl__TlsV1_1         QSsl__SslProtocol = QSsl__SslProtocol(3)
	QSsl__TlsV1_2         QSsl__SslProtocol = QSsl__SslProtocol(4)
	QSsl__AnyProtocol     QSsl__SslProtocol = QSsl__SslProtocol(5)
	QSsl__TlsV1SslV3      QSsl__SslProtocol = QSsl__SslProtocol(6)
	QSsl__SecureProtocols QSsl__SslProtocol = QSsl__SslProtocol(7)
	QSsl__TlsV1_0OrLater  QSsl__SslProtocol = QSsl__SslProtocol(8)
	QSsl__TlsV1_1OrLater  QSsl__SslProtocol = QSsl__SslProtocol(9)
	QSsl__TlsV1_2OrLater  QSsl__SslProtocol = QSsl__SslProtocol(10)
	QSsl__DtlsV1_0        QSsl__SslProtocol = QSsl__SslProtocol(11)
	QSsl__DtlsV1_0OrLater QSsl__SslProtocol = QSsl__SslProtocol(12)
	QSsl__DtlsV1_2        QSsl__SslProtocol = QSsl__SslProtocol(13)
	QSsl__DtlsV1_2OrLater QSsl__SslProtocol = QSsl__SslProtocol(14)
	QSsl__TlsV1_3         QSsl__SslProtocol = QSsl__SslProtocol(15)
	QSsl__TlsV1_3OrLater  QSsl__SslProtocol = QSsl__SslProtocol(16)
	QSsl__UnknownProtocol QSsl__SslProtocol = QSsl__SslProtocol(-1)
)

//go:generate stringer -type=QSsl__SslOption
//QSsl::SslOption
type QSsl__SslOption int64

const (
	QSsl__SslOptionDisableEmptyFragments         QSsl__SslOption = QSsl__SslOption(0x01)
	QSsl__SslOptionDisableSessionTickets         QSsl__SslOption = QSsl__SslOption(0x02)
	QSsl__SslOptionDisableCompression            QSsl__SslOption = QSsl__SslOption(0x04)
	QSsl__SslOptionDisableServerNameIndication   QSsl__SslOption = QSsl__SslOption(0x08)
	QSsl__SslOptionDisableLegacyRenegotiation    QSsl__SslOption = QSsl__SslOption(0x10)
	QSsl__SslOptionDisableSessionSharing         QSsl__SslOption = QSsl__SslOption(0x20)
	QSsl__SslOptionDisableSessionPersistence     QSsl__SslOption = QSsl__SslOption(0x40)
	QSsl__SslOptionDisableServerCipherPreference QSsl__SslOption = QSsl__SslOption(0x80)
)

//go:generate stringer -type=QSslCertificate__SubjectInfo
//QSslCertificate::SubjectInfo
type QSslCertificate__SubjectInfo int64

const (
	QSslCertificate__Organization               QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(0)
	QSslCertificate__CommonName                 QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(1)
	QSslCertificate__LocalityName               QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(2)
	QSslCertificate__OrganizationalUnitName     QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(3)
	QSslCertificate__CountryName                QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(4)
	QSslCertificate__StateOrProvinceName        QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(5)
	QSslCertificate__DistinguishedNameQualifier QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(6)
	QSslCertificate__SerialNumber               QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(7)
	QSslCertificate__EmailAddress               QSslCertificate__SubjectInfo = QSslCertificate__SubjectInfo(8)
)

//go:generate stringer -type=QSslConfiguration__NextProtocolNegotiationStatus
//QSslConfiguration::NextProtocolNegotiationStatus
type QSslConfiguration__NextProtocolNegotiationStatus int64

const (
	QSslConfiguration__NextProtocolNegotiationNone        QSslConfiguration__NextProtocolNegotiationStatus = QSslConfiguration__NextProtocolNegotiationStatus(0)
	QSslConfiguration__NextProtocolNegotiationNegotiated  QSslConfiguration__NextProtocolNegotiationStatus = QSslConfiguration__NextProtocolNegotiationStatus(1)
	QSslConfiguration__NextProtocolNegotiationUnsupported QSslConfiguration__NextProtocolNegotiationStatus = QSslConfiguration__NextProtocolNegotiationStatus(2)
)

//go:generate stringer -type=QSslDiffieHellmanParameters__Error
//QSslDiffieHellmanParameters::Error
type QSslDiffieHellmanParameters__Error int64

const (
	QSslDiffieHellmanParameters__NoError               QSslDiffieHellmanParameters__Error = QSslDiffieHellmanParameters__Error(0)
	QSslDiffieHellmanParameters__InvalidInputDataError QSslDiffieHellmanParameters__Error = QSslDiffieHellmanParameters__Error(1)
	QSslDiffieHellmanParameters__UnsafeParametersError QSslDiffieHellmanParameters__Error = QSslDiffieHellmanParameters__Error(2)
)

//go:generate stringer -type=QSslError__SslError
//QSslError::SslError
type QSslError__SslError int64

const (
	QSslError__NoError                             QSslError__SslError = QSslError__SslError(0)
	QSslError__UnableToGetIssuerCertificate        QSslError__SslError = QSslError__SslError(1)
	QSslError__UnableToDecryptCertificateSignature QSslError__SslError = QSslError__SslError(2)
	QSslError__UnableToDecodeIssuerPublicKey       QSslError__SslError = QSslError__SslError(3)
	QSslError__CertificateSignatureFailed          QSslError__SslError = QSslError__SslError(4)
	QSslError__CertificateNotYetValid              QSslError__SslError = QSslError__SslError(5)
	QSslError__CertificateExpired                  QSslError__SslError = QSslError__SslError(6)
	QSslError__InvalidNotBeforeField               QSslError__SslError = QSslError__SslError(7)
	QSslError__InvalidNotAfterField                QSslError__SslError = QSslError__SslError(8)
	QSslError__SelfSignedCertificate               QSslError__SslError = QSslError__SslError(9)
	QSslError__SelfSignedCertificateInChain        QSslError__SslError = QSslError__SslError(10)
	QSslError__UnableToGetLocalIssuerCertificate   QSslError__SslError = QSslError__SslError(11)
	QSslError__UnableToVerifyFirstCertificate      QSslError__SslError = QSslError__SslError(12)
	QSslError__CertificateRevoked                  QSslError__SslError = QSslError__SslError(13)
	QSslError__InvalidCaCertificate                QSslError__SslError = QSslError__SslError(14)
	QSslError__PathLengthExceeded                  QSslError__SslError = QSslError__SslError(15)
	QSslError__InvalidPurpose                      QSslError__SslError = QSslError__SslError(16)
	QSslError__CertificateUntrusted                QSslError__SslError = QSslError__SslError(17)
	QSslError__CertificateRejected                 QSslError__SslError = QSslError__SslError(18)
	QSslError__SubjectIssuerMismatch               QSslError__SslError = QSslError__SslError(19)
	QSslError__AuthorityIssuerSerialNumberMismatch QSslError__SslError = QSslError__SslError(20)
	QSslError__NoPeerCertificate                   QSslError__SslError = QSslError__SslError(21)
	QSslError__HostNameMismatch                    QSslError__SslError = QSslError__SslError(22)
	QSslError__NoSslSupport                        QSslError__SslError = QSslError__SslError(23)
	QSslError__CertificateBlacklisted              QSslError__SslError = QSslError__SslError(24)
	QSslError__CertificateStatusUnknown            QSslError__SslError = QSslError__SslError(25)
	QSslError__OcspNoResponseFound                 QSslError__SslError = QSslError__SslError(26)
	QSslError__OcspMalformedRequest                QSslError__SslError = QSslError__SslError(27)
	QSslError__OcspMalformedResponse               QSslError__SslError = QSslError__SslError(28)
	QSslError__OcspInternalError                   QSslError__SslError = QSslError__SslError(29)
	QSslError__OcspTryLater                        QSslError__SslError = QSslError__SslError(30)
	QSslError__OcspSigRequred                      QSslError__SslError = QSslError__SslError(31)
	QSslError__OcspUnauthorized                    QSslError__SslError = QSslError__SslError(32)
	QSslError__OcspResponseCannotBeTrusted         QSslError__SslError = QSslError__SslError(33)
	QSslError__OcspResponseCertIdUnknown           QSslError__SslError = QSslError__SslError(34)
	QSslError__OcspResponseExpired                 QSslError__SslError = QSslError__SslError(35)
	QSslError__OcspStatusUnknown                   QSslError__SslError = QSslError__SslError(36)
	QSslError__UnspecifiedError                    QSslError__SslError = QSslError__SslError(-1)
)

//go:generate stringer -type=QSslSocket__SslMode
//QSslSocket::SslMode
type QSslSocket__SslMode int64

const (
	QSslSocket__UnencryptedMode QSslSocket__SslMode = QSslSocket__SslMode(0)
	QSslSocket__SslClientMode   QSslSocket__SslMode = QSslSocket__SslMode(1)
	QSslSocket__SslServerMode   QSslSocket__SslMode = QSslSocket__SslMode(2)
)

//go:generate stringer -type=QSslSocket__PeerVerifyMode
//QSslSocket::PeerVerifyMode
type QSslSocket__PeerVerifyMode int64

const (
	QSslSocket__VerifyNone     QSslSocket__PeerVerifyMode = QSslSocket__PeerVerifyMode(0)
	QSslSocket__QueryPeer      QSslSocket__PeerVerifyMode = QSslSocket__PeerVerifyMode(1)
	QSslSocket__VerifyPeer     QSslSocket__PeerVerifyMode = QSslSocket__PeerVerifyMode(2)
	QSslSocket__AutoVerifyPeer QSslSocket__PeerVerifyMode = QSslSocket__PeerVerifyMode(3)
)
