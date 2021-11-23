# SHTRIH

Author Shumilin Alexander (email dev@spy.im), specially for http://tracker.ucs.ru/issues/175782

## SHTRIH-Com

HTTP<->COM server for scales shtrih (com protocol) version 0.1  
connection string is "serverHTTPip:serverHTTPport:scalesCOMport:scalesCOMbaud" where

- serverHTTPip > empty (default 0.0.0.0) or ip
- serverHTTPport > free TCP port 0-65535 (default 50505)
- scalesCOMport > COM port scales
- scalesCOMbaud > speed COM port scales
- (use default password 30)

Example:

>SHTRIH-Com.exe ":50505:COM1:9600"
- HTTP server listening on 0.0.0.0:50505
- server listening and sending message on port COM1 speed 9600
- URL http://127.0.0.1:50505/?cmd=getweight

## SHTRIH-Ethernet

HTTP<->UDP server for scales shtrih (network protocol) version 0.1  
connection string is "serverHTTPip:serverHTTPport:serverUDPport:scalesUDPip:scalesUDPport" where

- serverHTTPip > IP
- serverHTTPport > free TCP port 0-65535 (default 50505)
- serverUDPport > free UDP port 0-65535
- scalesUDPip > IP scales
- scalesUDPport > port scales
- (use default password 30)

Example:

> node server.js "0.0.0.0:50505:2005:192.168.0.202:1111"
- HTTP server listening on 0.0.0.0:50505
- UDP server listening on 0.0.0.0:2005 and send message on 192.168.0.202:1111
- scales listening on 192.168.0.202:1111 and send message on udp server
- URL http://127.0.0.1:50505/?cmd=getweight
