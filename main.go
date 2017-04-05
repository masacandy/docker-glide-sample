package main

// Packages
import (
	"fmt"
	"log"
	"github.com/comail/colog"
	"net"
	"net/http"
	"net/http/fcgi"

	elastic "gopkg.in/olivere/elastic.v5"
)

func handler(resp http.ResponseWriter, req *http.Request) {
	client, err := elastic.NewClient(elastic.SetURL("http://elasticsearch:9200"))
	if err != nil {
		panic(err)
	}

	esversion, err := client.ElasticsearchVersion("http://elasticsearch:9200")
	if err != nil {
		panic(err)
	}
	fmt.Fprintf(resp, "<h1>%s</h1>", esversion)

	log.Printf(esversion)
}

func main() {
   	colog.SetDefaultLevel(colog.LDebug)
	colog.SetMinLevel(colog.LTrace)
	colog.SetFormatter(&colog.StdFormatter{
		Colors: true,
		Flag:   log.Ldate | log.Ltime | log.Lshortfile,
	})
	colog.Register()

	listener, err := net.Listen("tcp", ":9000")
	if err != nil {
		panic(err)
	}

	http.HandleFunc("/", handler)
	fcgi.Serve(listener, nil)
}
