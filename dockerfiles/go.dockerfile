FROM golang:1.18-alpine
MAINTAINER mtg
WORKDIR /go/src/
COPY ./binance-collect .
RUN wget -q -O wait-for  https://raw.githubusercontent.com/eficode/wait-for/v2.2.3/wait-for && chmod +x wait-for && apk add --update --no-cache netcat-openbsd
EXPOSE 8200
EXPOSE 8201
CMD ["/bin/bash","-c","./binance-collect"]
