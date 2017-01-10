package topology

import "github.com/cnapache/ToneDFS/tools"

//FID
type FID string

func NewFid() FID {
	return FID(tools.RandomUUID())
}
