package main

import (
	"flag"
	"fmt"
	"log/slog"
	"net"
	"os"

	"github.com/patos-ufscar/http-web-server-example-go/common"
	"github.com/patos-ufscar/http-web-server-example-go/utils"
)

var configPath string

func init() {
	utils.InitSlogger()
}

func main() {
	l, err := net.Listen("tcp", "0.0.0.0:4221")
	if err != nil {
		slog.Error("Failed to bind to port 4221")
		os.Exit(1)
	}

	flag.StringVar(&configPath, "configPath", "", "the config.yml file path")
	flag.Parse()

	handlers := common.ParseConfig(configPath)

	ports := []string{}
	for _, v := range handlers {
		ports = append(ports, string(v.Port))
	}

	slog.Info(fmt.Sprintf("GoServe Listening on %s", ports))

	for {
		conn, err := l.Accept()
		if err != nil {
			slog.Error(fmt.Sprintf("Error accepting connection: %s", err.Error()))
			continue
		}

		go func (conn net.Conn) {

			defer func(conn net.Conn) {
				// we re-reply in case of error (reply missing)
				r := recover()
				if r != nil {
					slog.Error(fmt.Sprint("Recovered from: ", r))
					slog.Error(fmt.Sprintf("Error handling connection: %s", err.Error()))
					err := utils.Reply502(conn)
					if err != nil {
						slog.Error(fmt.Sprintf("Could not reply: %s", err.Error()))
					}
				}
			}(conn)
			defer conn.Close()

			// rep, err := handlers.HandleGlobal(conn, config)
			// if err != nil {
			// 	panic(err)
			// } else {
			// 	slog.Info(fmt.Sprintf("handled: %s", string(rep)))
			// }

		}(conn)
	}
}

// package main

// import (
// 	"log"
// 	"net/http"
// 	"net/http/httputil"
// 	"net/url"
// )

// func main() {
//         remote, err := url.Parse("http://google.com")
//         if err != nil {
// 			panic(err)
//         }

//         handler := func(p *httputil.ReverseProxy) func(http.ResponseWriter, *http.Request) {
// 			return func(w http.ResponseWriter, r *http.Request) {
// 				log.Println(r.URL)
// 				r.Host = remote.Host
// 				w.Header().Set("X-Ben", "Rad")
// 				p.ServeHTTP(w, r)
// 			}
//         }
        
//         proxy := httputil.NewSingleHostReverseProxy(remote)
//         http.HandleFunc("/", handler(proxy))
//         err = http.ListenAndServe(":8080", nil)
//         if err != nil {
// 			panic(err)
//         }
// }