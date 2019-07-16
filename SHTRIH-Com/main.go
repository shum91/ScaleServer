package main

import (
	"io"
    "os"
    "fmt"
    "log"
    "time"
	"strconv"
	"strings"
	"net/http"
	"github.com/tarm/serial"
)

var globalWeight int

func main () {
    configString := "0.0.0.0:50505:COM1:9600"
	if len(os.Args) > 1 {
        if os.Args[1] == "h" || os.Args[1] == "-h" || os.Args[1] == "help" || os.Args[1] == "-help" || os.Args[1] == "--help" {
            fmt.Println("HTTP<->COM server for scales shtrih (not network, only com protocol) version 0.1")
            fmt.Println("Author shumilin alexander (email a.shumilin@ucs.ru), specially for http://tracker.ucs.ru/issues/175782\n")
            fmt.Println("config string  -> serverHTTPip:serverHTTPport:scalesCOMport:scalesCOMbaud")
            fmt.Println("serverHTTPip   -> use empty (default 0.0.0.0) or ip")
            fmt.Println("serverHTTPport -> use free port 0-65535 (default 50505)")
            fmt.Println("scalesCOMport  -> use COM port scales")
            fmt.Println("scalesCOMbaud  -> use speed COM port scales")
            fmt.Println("(use default password 30)\n\nExample:")
            fmt.Println(os.Args[0], "\":50505:COM1:9600\"")
            fmt.Println("HTTP server listening on 0.0.0.0:50505")
            fmt.Println("     server listening and sending message on port COM1 speed 9600")
            fmt.Println("URL  http://127.0.0.1:50505/?cmd=getweight")
            os.Exit(0)
        } else {
            configString = os.Args[1]
        }
	}
    fmt.Println("NEED HELP? Enter command:", os.Args[0], "--help\n")
    
    config := strings.Split(configString, ":")
	configComBaud, _ := strconv.Atoi(config[3]);
    
    go loopGetWeight(config[2], configComBaud)
    
	log.Printf("START Server, config string %s", configString)
	
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		log.Printf("CLIENT %s %s %s", r.RemoteAddr, r.Method, r.URL)
		w.Header().Set("Content-Type", "text/xml; charset=utf-8")
		io.WriteString(w, "<?xml version=\"1.0\" encoding=\"UTF-8\"?><Scale><ErrorText/><Weight>" + strconv.Itoa(globalWeight) + "</Weight></Scale>\n")
	})
	
    log.Fatal(http.ListenAndServe(config[0] + ":" + config[1], nil))
}

func loopGetWeight(comPort string, comBaud int) {
    s, err := serial.OpenPort(&serial.Config{Name: comPort, Baud: comBaud})
    if err != nil {
        log.Println("COM ERROR function loopGetWeight dead:", err)
        return
    } else {
        log.Println("COM port OK")
    }
    for {
        if _, err := s.Write([]byte{0x2, 0x5, 0x38, 0x30, 0x30, 0x33, 0x30, 0x3E}); err != nil {
            log.Println("COM ERROR function loopGetWeight dead:", err)
            return
        }
        buf := []byte{0}
        if _, err := s.Read(buf); err != nil {
            log.Println("COM ERROR function loopGetWeight dead:", err)
            return
        } else {
            if buf[0] == 0x6 {
                log.Printf("COM RECEIVED ACK %x", buf)
                s.Read(buf);
                log.Printf("COM RECEIVED STX %x", buf)
                s.Read(buf);
                log.Printf("COM RECEIVED length %x", buf)
                tailL := int(buf[0])+1
                tail  := make([]byte, tailL)
                for i := 0; i < tailL; i++ {
                    s.Read(buf);
                    tail[i] = buf[0]
                }
                if (tail[0] != 0x38) || (tail[1] != 0x00) {
                    log.Printf("COM RECEIVED ERR %x, send NAK 0x15 and restart loopGetWeight", tail)
                    s.Write([]byte{0x15});
                } else {
                    log.Printf("COM RECEIVED tail %x", tail)
                    globalWeight = int(tail[2])+int(tail[3])<<8
                    s.Write([]byte{0x6})
                }
            } else {
                log.Printf("COM RECEIVED ERR %x (NAK = 0x15)", buf)
            }
        }
        time.Sleep(2*time.Second)
    }
    s.Close()
}