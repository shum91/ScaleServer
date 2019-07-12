const http      = require('http');
const dgram     = require('dgram');
const serverUDP = dgram.createSocket('udp4');

var globalWeight = 0;

var configString   = ((process.argv.length < 3) ? "0.0.0.0::50505::2005::192.168.0.202::1111::30" : process.argv[2]).split("::");
var serverHTTPip   = configString[0];
var serverHTTPport = configString[1];
var serverUDPport  = configString[2];
var scalesUDPip    = configString[3];
var scalesUDPport  = configString[4];
var scalesPASSWORD = configString[5];

serverUDP.on('error', (err) => {
  console.log((new Date()).toTimeString(), ` UDP server error:\n${err.stack}`);
  serverUDP.close();
});

serverUDP.on('message', (msg, rinfo) => {
  console.log((new Date()).toTimeString(), ` UDP server got:`, msg, `from ${rinfo.address}:${rinfo.port}`);
  if (msg[0] == 0x2) {
	  globalWeight = (Buffer.from([msg[4], msg[5]])).readInt16LE(0);
	  console.log((new Date()).toTimeString(), ' UDP server get weight OK ' + globalWeight);
  } else console.log((new Date()).toTimeString(), ' UDP server get weight ERR');
});

serverUDP.on('listening', () => {
  const address = serverUDP.address();
  console.log((new Date()).toTimeString(), ` UDP server running ${address.address}:${address.port}`);
});

serverUDP.bind(serverUDPport);

setInterval(function () {
  getWeightMsg = Buffer.from('02053830303330', 'hex');
  serverUDP.send(getWeightMsg, scalesUDPport, scalesUDPip);
  //console.log((new Date()).toTimeString(), " UDP server send: ", getWeightMsg);
}, 1000);


const serverHTTP = http.createServer((req, res) => {
  res.statusCode = 200;
  res.setHeader('Content-Type', 'application/xml; charset = utf-8');
  res.end("<?xml version=\"1.0\" encoding=\"UTF-8\"?><Scale><ErrorText/><Weight>" + globalWeight + "</Weight></Scale>\n");
});

serverHTTP.listen(serverHTTPport, serverHTTPip, () => {
  console.log((new Date()).toTimeString(), `HTTP server running at http://${serverHTTPip}:${serverHTTPport}/`);
});