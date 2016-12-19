package topology

//Server 服务器结构
type Server struct {
	dataCenterName string
	rackName       string
	serverName     string
	serverIP       string
	serverPort     string
	dataDir        string
	volume         *Volume
}

func (me *Server) Start() {

}

func (me *Server) LoadVolume() {
	// sbFiles, err := tools.NewDirectory(me.dataDir).GetFilesFilterExt(".sb")
	// if err != nil {
	// 	return
	// }
	// me.volume = &Volume{}
}
