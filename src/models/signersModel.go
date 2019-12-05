package models

import (
	"github.com/fibercrypto/fibercryptowallet/src/util/logging"
	"github.com/sirupsen/logrus"
	"github.com/therecipe/qt/core"
	"github.com/therecipe/qt/qml"
	wlcore "github.com/fibercrypto/fibercryptowallet/src/main"
)

const (
	SignerId              = iota + int(core.Qt__UserRole) + 1
	SignerDescription
)

var logSignersModel = logging.MustGetLogger("Signers Model")

type SignerModel struct {
	core.QAbstractListModel

	_ func()                       `constructor:"init"`
	_ map[int]*core.QByteArray     `property:"roles"`
	_ []*QSigner                   `property:"signers"`
	_ func(string)                 `slot:"loadModel"`
	_ int                          `property:"count"`
}

type QSigner struct {
	core.QObject

	_ string  `property:"id"`
	_ string  `property:"description"`
}

func (signerModel *SignerModel) init() {
	logWalletsModel.Debugln("Initialize Signer model")
	signerModel.SetRoles(map[int]*core.QByteArray{
		SignerId:              core.NewQByteArray2("id", -1),
		SignerDescription:     core.NewQByteArray2("description", -1),
	})
	qml.QQmlEngine_SetObjectOwnership(signerModel, qml.QQmlEngine__CppOwnership)
	signerModel.ConnectData(signerModel.data)
	signerModel.ConnectRowCount(signerModel.rowCount)
	signerModel.ConnectColumnCount(signerModel.columnCount)
	signerModel.ConnectRoleNames(signerModel.roleNames)
	signerModel.ConnectLoadModel(signerModel.loadModel)
}

func (signerModel *SignerModel) data(index *core.QModelIndex, role int) *core.QVariant {
	if !index.IsValid() {
		return core.NewQVariant()
	}
	if index.Row() >= len(signerModel.Signers()) {
		return core.NewQVariant()
	}
	var s = signerModel.Signers()[index.Row()]
	switch role {
	case SignerId:
		{
			return core.NewQVariant1(s.Id())
		}
	case SignerDescription:
		{
			return core.NewQVariant1(s.Description())
		}

	default:
		{
			return core.NewQVariant()
		}
	}
}

func (signerModel *SignerModel) rowCount(parent *core.QModelIndex) int {
	return len(signerModel.Signers())
}

func (signersModel *SignerModel) columnCount(parent *core.QModelIndex) int {
	return 1
}

func (signersModel *SignerModel) roleNames() map[int]*core.QByteArray {
	return signersModel.Roles()
}

func (signerModel *SignerModel) loadModel(wltId string) {
	logSignersModel.Debugln("Loading signers", wltId)
	wlt := walletManager.WalletEnv.GetWalletSet().GetWallet(wltId)
	if wlt == nil {
		return
	}
	am := wlcore.LoadAltcoinManager()
	signers := am.EnumerateSignServices()
	var qSigners []*QSigner
	qSigner := NewQSigner(nil)
	qml.QQmlEngine_SetObjectOwnership(qSigner, qml.QQmlEngine__CppOwnership)
	qSigner.SetId(wlt.GetId())
	qSigner.SetDescription(wlt.GetLabel())
	qSigners = append(qSigners, qSigner)
	for signers.Next() {
		qSigner = NewQSigner(nil)
		qml.QQmlEngine_SetObjectOwnership(qSigner, qml.QQmlEngine__CppOwnership)
		signer := signers.Value()
		sameSeedAsWlt, err := signer.ReadyForTxn(wlt, nil)
		if err != nil {
			logrus.WithError(err).Errorln("unable to check wallet and signer compatibility")
			return
		}
		if !sameSeedAsWlt {
			continue
		}
		uid, err := signer.GetSignerUID()
		if err != nil {
			logrus.WithError(err).Error("unable to get signer uid")
			return
		}
		qSigner.SetId(string(uid))
		desc, err := signer.GetSignerDescription()
		if err != nil {
			logrus.WithError(err).Error("unable to get signer description")
			return
		}
		qSigner.SetDescription(desc)
		qSigners = append(qSigners, qSigner)
	}
	signerModel.BeginResetModel()
	signerModel.SetSigners(qSigners)
	signerModel.SetCount(len(signerModel.Signers()))
	signerModel.EndResetModel()
}
