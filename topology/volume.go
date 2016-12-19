package topology

import (
	"errors"
	"os"
	"path/filepath"
	"strconv"
	"sync"
)

const (
	SuperBlockNumberLimit        = 10          //超块数量限制
	SuperBlockTotalLength int64  = 34359738368 //超块总容量 bytes 1024 *1024 *1024 * 32 = 32GB
	SuperBlockFileExt     string = ".sb"       //超块文件扩展名
)

//Volume 卷结构
type Volume struct {
	dir        string
	lock       *sync.Mutex
	superBlock [SuperBlockNumberLimit]*SuperBlock
}

//NewSuperBlock 新建超块
func (me *Volume) NewSuperBlock() (*SuperBlock, error) {
	me.lock.Lock()
	defer me.lock.Unlock()

	if len(me.superBlock) >= SuperBlockNumberLimit {
		return nil, errors.New("卷SuperBlock数量达到上限。")
	}

	newSuperBlockID := len(me.superBlock) + 1

	superBlockFile, err := os.OpenFile(filepath.Join(me.dir, strconv.Itoa(newSuperBlockID)+SuperBlockFileExt), os.O_CREATE|os.O_RDWR, os.ModePerm)
	if err != nil {
		return nil, err
	}

	err = superBlockFile.Truncate(SuperBlockTotalLength)
	if err != nil {
		return nil, err
	}

	nsb := SuperBlock{superBlockID: uint32(newSuperBlockID),
		dataFile:    superBlockFile,
		endPosition: 0,
		writeLock:   new(sync.Mutex)}

	me.superBlock[newSuperBlockID] = &nsb

	return &nsb, nil
}

func (me *SuperBlock) LoadSuperBlock(filePath string) {

}
