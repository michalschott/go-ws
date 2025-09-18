package main

import (
	"encoding/json"
	"log"
	"net"
	"net/http"
	"time"

	"github.com/gorilla/websocket"
)

type LogEntry struct {
	Level   string `json:"level"`
	Time    string `json:"time"`
	Message string `json:"message"`
	Client  string `json:"client,omitempty"`
}

func logJSON(level, msg, client string) {
	entry := LogEntry{
		Level:   level,
		Time:    time.Now().Format(time.RFC3339),
		Message: msg,
		Client:  client,
	}
	data, _ := json.Marshal(entry)
	log.Println(string(data))
}

var upgrader = websocket.Upgrader{
	CheckOrigin: func(r *http.Request) bool { return true },
}

func handleWS(w http.ResponseWriter, r *http.Request) {
	conn, err := upgrader.Upgrade(w, r, nil)
	if err != nil {
		logJSON("error", "Upgrade failed: "+err.Error(), "")
		return
	}
	defer conn.Close()

	// extract client IP
	clientIP, _, _ := net.SplitHostPort(r.RemoteAddr)
	logJSON("info", "Client connected", clientIP)

	for {
		mt, msg, err := conn.ReadMessage()
		if err != nil {
			logJSON("error", "Read error: "+err.Error(), clientIP)
			break
		}
		logJSON("info", "Received: "+string(msg), clientIP)

		err = conn.WriteMessage(mt, msg)
		if err != nil {
			logJSON("error", "Write error: "+err.Error(), clientIP)
			break
		}
	}
}

func handleHome(w http.ResponseWriter, r *http.Request) {
	html := `
<!DOCTYPE html>
<html>
<head>
<meta charset="utf-8">
<title>Go WebSocket Test</title>
<style>
#log { font-family: monospace; white-space: pre; border:1px solid #ccc; padding:5px; max-height:300px; overflow-y:auto;}
</style>
</head>
<body>
  <h2>WebSocket Test</h2>
  <button onclick="connect()">Connect</button>
  <input id="msg" placeholder="Type message"/>
  <button onclick="sendMsg()">Send</button>
  <div id="log"></div>
<script>
let ws;
function log(msg){
  const logDiv=document.getElementById("log");
  const line=document.createElement("div");
  line.textContent=msg;
  logDiv.appendChild(line);
  logDiv.scrollTop=logDiv.scrollHeight;
}
function connect(){
  ws=new WebSocket("ws://"+location.host+"/ws");
  ws.onopen=()=>log("Connected");
  ws.onmessage=(e)=>log("Server: "+e.data);
  ws.onclose=()=>log("Closed");
}
function sendMsg(){
  const input=document.getElementById("msg");
  ws.send(input.value);
  log("Client: "+input.value);
  input.value="";
}
</script>
</body>
</html>`
	w.Header().Set("Content-Type", "text/html")
	w.Write([]byte(html))
}

func handleHealth(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("ok"))
}

func main() {
	http.HandleFunc("/", handleHome)
	http.HandleFunc("/ws", handleWS)
	http.HandleFunc("/healthz", handleHealth)

	logJSON("info", "Server started at :8080", "")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		logJSON("fatal", "Server stopped: "+err.Error(), "")
	}
}
