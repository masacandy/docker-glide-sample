package main

// Packages
import (
	"fmt"
	log "github.com/cihub/seelog"
	//	"html"
	"net/http"
	"os"
	//	"io/ioutil"
	elastic "gopkg.in/olivere/elastic.v5"
)

// ログの設定
const logConfig = `
<seelog type="adaptive" mininterval="200000000" maxinterval="1000000000" critmsgcount="5">
<formats>
    <format id="main" format="%Date(2006-01-02T15:04:05.999999999Z07:00) [%File:%FuncShort:%Line] [%LEV] %Msg%n" />
</formats>
<outputs formatid="main">
    <filter levels="trace,debug,info,warn,error,critical">
        <console />
    </filter>
    <filter levels="info,warn,error,critical">
        <rollingfile filename="/tmp/log.log" type="size" maxsize="102400" maxrolls="1" />
    </filter>
</outputs>
</seelog>`

func initLogger() {
	logger, err := log.LoggerFromConfigAsBytes([]byte(logConfig))
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
	log.ReplaceLogger(logger)
}

func main() {
	initLogger()
	defer log.Flush()

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		client, err := elastic.NewClient(elastic.SetURL("http://elasticsearch:9200"))
		if err != nil {
			// Handle error
			panic(err)
		}

		esversion, err := client.ElasticsearchVersion("http://elasticsearch:9200")
		if err != nil {
			panic(err)
		}
		fmt.Fprintf(w, "%s", esversion)

		log.Info(esversion)
	})

	log.Error(http.ListenAndServe(":8080", nil))
}
