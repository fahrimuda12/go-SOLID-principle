package config

import (
	"go-simple-MVC/app/controllers"
)

type Initialization struct {
	WalletCtrl controllers.WalletController
}

func NewInitialization(
	walletCtrl controllers.WalletController,
) *Initialization {
	return &Initialization{
		WalletCtrl: walletCtrl,
	}
}