package topology

//Needle 指针结构
type Needle struct {
	beginPosition int64
	totalLength   int64
	nameLength    int64
}

func NewNeedle(beginPosition, totalLength, nameLength int64) Needle {
	return Needle{beginPosition: beginPosition,
		totalLength: totalLength,
		nameLength:  nameLength}
}

func (self *Needle) GetFilePosition() (filePosition, fileLength int64) {
	filePosition = self.beginPosition + 10 + 32 + 19 + 5 + 10 + self.nameLength
	fileLength = self.totalLength - 10 - 32 - 19 - 5 - 10 - self.nameLength
	return
}
