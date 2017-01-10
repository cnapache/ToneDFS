package main

import (
	"flag"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"github.com/cnapache/ToneDFS/topology"
	"github.com/gogap/go-gelf/gelf"
	"github.com/gogap/logrus"
	"github.com/gogap/logrus/hooks/graylog"
	"github.com/gorilla/mux"
)

var confPath = flag.String("conf", "", "conf file path")
var superBlock topology.SuperBlock

func main() {
	//flag.Parse()
	count := 0
	for count < 10 {
		gelfWriter, err := gelf.NewWriter("192.168.31.9:12201")
		if err != nil {
			log.Fatalf("gelf.NewWriter: %s", err)
		}
		// log to both stderr and graylog2
		log.SetOutput(io.MultiWriter(os.Stderr, gelfWriter))
		log.Printf("logging to stderr & graylog2@'%s'", "graylogAddr")
		//log.Fatal("asdasd")
		count++
		time.Sleep(1 * time.Nanosecond)
	}
	return

	logrus.SetFormatter(&logrus.JSONFormatter{})
	glog, err := graylog.NewHook("192.168.31.9:12201", "yijifu", nil)
	if err != nil {
		logrus.Error(err)
		return
	}
	logrus.SetLevel(logrus.DebugLevel)
	logrus.AddHook(glog)
	//logrus.AddHook(file.NewHook("logs/ss.log"))
	logrus.Debug("asdas")
	logrus.WithField("biz", "member").Errorf("member not login,member is %s", "100")

	//serverStart()
}

func serverStart() {
	r := mux.NewRouter()
	r.HandleFunc("/{fid}", func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(404)
	})
	http.Handle("/", r)
	if err := http.ListenAndServe(":9230", nil); err != nil {
		panic(err)
	}
}
