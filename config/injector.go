//go:build wireinject
// +build wireinject

package config

// go:build wireinject

import (
	"github.com/google/wire"
)

var db = wire.NewSet(DBInit)

// var walletCtrlSet = wire.NewSet(controllers.WalletControllerInit,
// 	wire.Bind(new(controllers.WalletController)),
//   )

func Init() *Initialization {
 wire.Build(NewInitialization, db)
 return nil
}