SHTRIH-Com<br />
HTTP<->COM server for scales shtrih (not network, only com protocol) version 0.1<br />
Author Shumilin Alexander (email a.shumilin@ucs.ru), specially for http://tracker.ucs.ru/issues/175782<br />
config string  -> serverHTTPip:serverHTTPport:scalesCOMport:scalesCOMbaud<br />
serverHTTPip   -> use empty (default 0.0.0.0) or ip<br />
serverHTTPport -> use free port 0-65535 (default 50505)<br />
scalesCOMport  -> use COM port scales<br />
scalesCOMbaud  -> use speed COM port scales<br />
(use default password 30)<br />
<br />
Example: SHTRIH-Com.exe ":50505:COM1:9600"<br />
HTTP server listening on 0.0.0.0:50505<br />
     server listening and sending message on port COM1 speed 9600<br />
URL  http://127.0.0.1:50505/?cmd=getweight<br /><br /><br />
SHTRIH-Ethernet

