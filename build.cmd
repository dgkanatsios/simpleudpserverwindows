docker build -t docker.io/dgkanatsios/simpleudpserverwindows:0.1 .
docker push docker.io/dgkanatsios/simpleudpserverwindows:0.1
docker run -d -p 30001:30001/udp -e PORT=30001 docker.io/dgkanatsios/simpleudpserverwindows:0.1