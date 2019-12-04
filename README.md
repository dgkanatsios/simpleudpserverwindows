# simpleudpserverwindows

A simple echo UDP server in a Windows container written in Go.

UDP Server code found in <https://github.com/bediger4000/udp-echo-server/blob/master/server.go>

```bash
docker run -d -p 30001:3601/udp docker.io/dgkanatsios/simpleudpserverwindows:0.2
```