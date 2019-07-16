Author Shumilin Alexander (email a.shumilin@ucs.ru), specially for http://tracker.ucs.ru/issues/175782<br />

# SHTRIH-Com

HTTP<->COM server for scales shtrih (not network, only com protocol) version 0.1

- config string  -> serverHTTPip:serverHTTPport:scalesCOMport:scalesCOMbaud
- serverHTTPip   -> use empty (default 0.0.0.0) or ip
- serverHTTPport -> use free port 0-65535 (default 50505)
- scalesCOMport  -> use COM port scales
- scalesCOMbaud  -> use speed COM port scales
- (use default password 30)

Example: SHTRIH-Com.exe ":50505:COM1:9600"

- HTTP server listening on 0.0.0.0:50505
- server listening and sending message on port COM1 speed 9600
- URL http://127.0.0.1:50505/?cmd=getweight

# SHTRIH-Ethernet<br />
HTTP<->UDP server for scales shtrih (not com, only network protocol)<br />
serverHTTPip:serverHTTPport:serverUDPport:scalesUDPip:scalesUDPport<br />
Example: node s.js "0.0.0.0:50505:2005:192.168.0.202:1111:"<br />
> HTTP server listening on 0.0.0.0:50505<br />
> UDP server listening on 0.0.0.0:2005 and send message on 192.168.0.202:1111<br />
> URL http://127.0.0.1:50505/?cmd=getweight<br />
> scales listening on 192.168.0.202:1111 and send message on udp server<br />