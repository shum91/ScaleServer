# ScaleServer
ScaleServer SHTRIH

http<->udp server for scales shtrih-m (not com, only network protocol)

author shumilin alexander (email a.shumilin@ucs.ru)

especially for http://tracker.ucs.ru/issues/175782

serverHTTPip::serverHTTPport::serverUDPport::scalesUDPip::scalesUDPport::scalesPASSWORD

example:

config string "0.0.0.0::50505::2005::192.168.0.202::1111::30"

http server listening on 0.0.0.0:50505

udp server listening on 0.0.0.0:2005 and send message on 192.168.0.202:1111

url http://127.0.0.1:50505/?cmd=getweight

scales listening on 192.168.0.202:1111 and send message on udp server
