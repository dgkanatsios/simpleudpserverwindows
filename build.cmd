docker build -t docker.io/dgkanatsios/simpleudpserverwindows:0.1 .
docker push docker.io/dgkanatsios/simpleudpserverwindows:0.1
docker run -d -p 30001:3601/udp -e PORT=3601 docker.io/dgkanatsios/simpleudpserverwindows:0.1