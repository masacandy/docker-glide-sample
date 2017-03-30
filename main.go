package main

// Packages
import (
	"fmt"
	log "github.com/cihub/seelog"
	"html"
	"net/http"
	"os"
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
		log.Info("Hello from SeeLog")
		fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
	})

	log.Error(http.ListenAndServe(":8080", nil))
}
