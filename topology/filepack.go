package topology

import (
	"bytes"
	"fmt"
	"hash/crc32"
	"math"
	"time"
)

//FilePack 文件包
type FilePack struct {
	CRC        uint32 //10byte
	Fid        FID    //32byte UUID
	Timestamp  int64  //19byte max:
	NameLength uint16 //5byte max:65535
	FileLength uint32 //10byte max:4294967295 单文件最大4G
	Name       []byte
	Data       []byte
}

//NewMetaBlock 新建数据块
func NewFilePack(name string, data []byte) *FilePack {
	nameByte := []byte(name)
	return &FilePack{
		CRC:        crc32Check(data),
		Fid:        NewFid(),
		Timestamp:  time.Now().Unix(),
		NameLength: uint16(len(nameByte)),
		FileLength: uint32(len(data)),
		Name:       nameByte,
		Data:       data}
}

//ToForamtByte 将元数据块转换为格式化的byte
//顺序 crc[10byte] fid[32byte] tstmap[19byte] nlength[5byte] dlength[10byte] name[nbyte] data[nbyte]
func (self *FilePack) ToForamtByte() ([]byte, error) {
	//NewBuffer可能是个问题
	bufferWrite := bytes.NewBuffer(make([]byte, 0))
	bufferWrite.Write(fillUint32(self.CRC))
	bufferWrite.Write([]byte(self.Fid))
	bufferWrite.Write(fillInt64(self.Timestamp))
	bufferWrite.Write(fillUint16(self.NameLength))
	bufferWrite.Write(fillUint32(self.FileLength))
	bufferWrite.Write(self.Name)
	bufferWrite.Write(self.Data)
	return bufferWrite.Bytes(), nil
}

//Length 获取数据长度
func (self *FilePack) Length() int64 {
	return int64(10 + 32 + 19 + 5 + 10 + len(self.Name) + len(self.Data))
}

var intFiller = []byte{48}

//将val byte长度填充到10byte
func fillUint32(val uint32) []byte {
	return fillStringToLength(fmt.Sprint(val), 10)
}

//将val byte长度填充到5byte
func fillUint16(val uint16) []byte {
	return fillStringToLength(fmt.Sprint(val), 5)
}

//将val byte长度填充到19byte
func fillInt64(val int64) []byte {
	return fillStringToLength(fmt.Sprint(val), 19)
}

//将val byte长度填充到指定长度
func fillStringToLength(val string, length int) []byte {
	valByte := []byte(val)
	if valByteLength := len(valByte); valByteLength < length {
		return append(bytes.Repeat(intFiller, length-valByteLength), valByte...)
	}
	return valByte
}

func crc32Check(data []byte) uint32 {
	ieee := crc32.NewIEEE()
	_, err := ieee.Write(data)
	if err != nil {
		return math.MaxUint32
	}
	return ieee.Sum32()
}
