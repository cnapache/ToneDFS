package config

import "github.com/cnapache/ToneDFS/tools"

type ToneServerConfig struct {
	DataCenterName   string
	RackName         string
	ServerName       string
	ServerIP         string
	ServerPort       string
	DataDir          string
	DataFileMaxLimit int
}

var fc tools.Config

func (tsc *ToneServerConfig) InitConfig(confPath string) (serverConfig *ToneServerConfig) {
	fc.InitConfig(confPath)
	serverSection := "server"
	tsc.DataCenterName = fc.Read(serverSection, "datacentername")
	tsc.DataDir = fc.Read(serverSection, "datadir")
	tsc.DataFileMaxLimit = fc.ReadInt(serverSection, "datafilemaxlimit", 10)
	tsc.RackName = fc.Read(serverSection, "rackname")
	tsc.ServerIP = fc.Read(serverSection, "serverip")
	tsc.ServerName = fc.Read(serverSection, "servername")
	tsc.ServerPort = fc.Read(serverSection, "serverport")

	return tsc
}
