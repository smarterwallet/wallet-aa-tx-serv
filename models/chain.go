package models

import (
	"gorm.io/gorm"
)

type Chain struct {
	gorm.Model
	NetworkId  uint   `gorm:"network_id" json:"netWorkId,omitempty"`
	Name       string `gorm:"not null;comment:链名称" json:"name"`
	RpcApi     string `gorm:"comment:链的rpc地址" json:"rpcApi,omitempty"`
	BundlerApi string `gorm:"comment:链的bundler地址" json:"bundlerApi,omitempty"`
}

func (Chain) TableName() string {
	return "chain"
}
