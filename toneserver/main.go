package main

import (
	"flag"
	"net/http"
	"strconv"
	"tone/topology"

	"fmt"

	"os"

	"path/filepath"

	"io/ioutil"

	"time"

	"sync"

	"github.com/gorilla/mux"
)

var confPath = flag.String("conf", "", "conf file path")
var superBlock topology.SuperBlock

var fileInfoCache = make(map[string]*topology.FilePack)

func main() {
	flag.Parse()

	superBlock = *topology.NewSuperBlock("W:\\go\\tonedfs\\storage")
	startTime := time.Now()
	wg := new(sync.WaitGroup)
	count := 0
	for count < 5000 {
		index := 0
		for index < 10 {
			wg.Add(1)
			file, err := os.Open(filepath.Join("W:\\go\\tonedfs\\storage", strconv.Itoa(index)+".jpg"))
			if err != nil {
				panic(err)
			}
			defer file.Close()
			fileByte, _ := ioutil.ReadAll(file)
			go func(fileData []byte) {

				fileblock := topology.NewFileBlock("123.png", fileData)
				needle, err := superBlock.Put("123.png", fileData)
				if err != nil {
					fmt.Println(err)
				}
				//position, length := needle.GetFilePosition()
				//fmt.Printf("fid:%v p:%v l:%v\n", needle.Fid, needle.Position, needle.TotalLength)
				//fileInfoCache[needle.Fid] = needle
				addCache(needle.Fid, needle)
				wg.Done()
			}(fileByte)
			index++
		}
		count++
	}

	wg.Wait()
	fmt.Println(len(fileInfoCache))
	fmt.Println(time.Since(startTime))
	serverStart()
}

func serverStart() {
	r := mux.NewRouter()
	r.HandleFunc("/{fid}", func(w http.ResponseWriter, r *http.Request) {
		vars := mux.Vars(r)

		if fileInfoCache[vars["fid"]] != nil {
			fileinfo := fileInfoCache[vars["fid"]]
			needle := topology.NewNeedle(fileinfo.Position, fileinfo.TotalLength, fileinfo.NameLength)
			w.Header().Add("Content-Type", "image/jpeg")
			superBlock.TakeNeedleAndWriterToHttpResponseWriter(w, &needle)
		}
		w.WriteHeader(404)

	})
	http.Handle("/", r)
	if err := http.ListenAndServe(":9230", nil); err != nil {
		panic(err)
	}
}

var cacheLock = new(sync.Mutex)

func addCache(key string, fileinfo *topology.FileInfo) {
	cacheLock.Lock()
	fileInfoCache[key] = fileinfo
	cacheLock.Unlock()
}

// func (server *Server) InitServer(confPath string) {
// 	var tsc config.ToneServerConfig
// 	tsc.InitConfig(confPath)
// }
