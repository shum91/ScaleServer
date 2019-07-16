:: HTTP<->UDP server for scales shtrih-m (not com, only network protocol)
:: Author Shumilin Alexander (email a.shumilin@ucs.ru)
:: Specially for http://tracker.ucs.ru/issues/175782

:: serverHTTPip::serverHTTPport::serverUDPport::scalesUDPip::scalesUDPport::scalesPASSWORD
:: example:
:: config string "0.0.0.0::50505::2005::192.168.0.202::1111::30"
:: http server listening on 0.0.0.0:50505
::  udp server listening on 0.0.0.0:2005 and send message on 192.168.0.202:1111
::  url http://127.0.0.1:50505/?cmd=getweight
:: scales listening on 192.168.0.202:1111 and send message on udp server

node.exe s.js "0.0.0.0::50505::2005::192.168.0.202::1111::30" 1>> "%Date:~6,4%%Date:~3,2%%Date:~0,2%_out.log" 2>>&1
pause