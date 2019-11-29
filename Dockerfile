FROM golang:1.12.13-windowsservercore-ltsc2016 AS builder
ADD . C:\\app
WORKDIR C:\\app
#RUN go mod download
RUN go build -a -o C:\\app\\simpleudpserverwindows.exe .

# final stage
FROM mcr.microsoft.com/windows/nanoserver:1909
COPY --from=builder C:\\app\\simpleudpserverwindows.exe C:\\simpleudpserverwindows.exe
ENTRYPOINT ["C:\\simpleudpserverwindows.exe"]