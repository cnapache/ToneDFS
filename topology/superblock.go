package topology

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"sync"
)

//SuperBlock 超块结构
type SuperBlock struct {
	superBlockID uint32
	writeLock    *sync.Mutex
	dataFile     *os.File
	endPosition  int64
}

//Put 将数据块写入超块
func (self *SuperBlock) Put(name string, data []byte) (*FilePack, error) {
	filePack := NewFilePack(name, data)

	formatByte, err := filePack.ToForamtByte()
	if err != nil {
		return nil, err
	}

	packLength := filePack.Length()

	self.writeLock.Lock()

	if self.endPosition+packLength > SuperBlockTotalLength {
		self.writeLock.Unlock()
		return nil, fmt.Errorf("超块(ID:%v)容量不足。", self.superBlockID)
	}
	readPosition := self.endPosition
	self.endPosition = self.endPosition + packLength

	self.writeLock.Unlock()

	_, err = self.dataFile.WriteAt(formatByte, readPosition)
	if err != nil {
		return nil, err
	}

	return filePack, nil
}

//TakeNeedleAndWriterToHttpResponseWriter 取出指针文件数据并发送到ResponseWrieter
//方法仅发送文件数据内容 不会发送Response.Headers以及文件名等相关信息
func (self *SuperBlock) TakeNeedleAndWriterToHttpResponseWriter(r http.ResponseWriter, nl *Needle) (bool, error) {
	var dataReader io.ReadSeeker = self.dataFile
	pos, len := nl.GetFilePosition()
	fmt.Println(len)
	_, err := dataReader.Read(make([]byte, len))
	if err != nil {
		return false, err
	}

	_, err = dataReader.Seek(pos, io.SeekStart)
	if err != nil {
		return false, err
	}

	sendLength, err := io.CopyN(r, dataReader, len)
	if err != nil || sendLength != len {
		return false, err
	}

	return true, nil
}
