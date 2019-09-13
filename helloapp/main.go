package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"math"
	"net/http"
	"os"

	"github.com/prometheus/client_golang/prometheus/promhttp"
)

const (
	defaultConfigPath = "/etc/helloapp/config.json"
	defaultListenAddr = ":8080"
	ApplicationName   = "helloapp"
)

type Config struct {
	TextToShow string `json:"text_to_show"`
}

var a float64

func main() {
	configPath := flag.String("config.path", defaultConfigPath, "Path to config")
	listenAddr := flag.String("net.listen", defaultListenAddr, "Listen address")
	flag.Parse()

	config, err := getConfig(*configPath)
	if err != nil {
		log.Fatal(err)
	}
	log.Printf("Starting with config: %+v", config)

	http.Handle("/metrics", promhttp.Handler())
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		for i:=0;i<40000;i++ {
			a = math.Sin(float64(i))
			ioutil.Discard.Write([]byte(fmt.Sprint(a)))
		}
		w.Write([]byte(config.TextToShow))
		log.Printf("INFO: %d %s %s %s", http.StatusOK, r.RemoteAddr, r.Method, r.URL)
	})

	log.Printf("Starting %s on %s", ApplicationName, *listenAddr)
	if err := http.ListenAndServe(*listenAddr, nil); err != nil {
		log.Fatal(err)
	}
}

func getConfig(configPath string) (Config, error) {
	var config Config
	log.SetFlags(log.Lshortfile | log.Ldate | log.Ltime | log.Lmicroseconds)
	fp, err := os.Open(configPath)
	if err != nil {
		return config, err
	}
	defer fp.Close()
	d := json.NewDecoder(fp)
	return config, d.Decode(&config)
}
