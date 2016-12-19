package topology

import (
	"tone/tools"
)

//FID
type FID string

func NewFid() FID {
	return FID(tools.RandomUUID())
}
