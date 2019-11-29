# simpleudpserverwindows

A simple echo UDP server in a Windows container written in Go.

UDP Server code found in <https://github.com/bediger4000/udp-echo-server/blob/master/server.g

```bash
docker run -d -p 30001:30001/udp -e PORT=30001 docker.io/dgkanatsios/simpleudpserverwindows:0.1
```