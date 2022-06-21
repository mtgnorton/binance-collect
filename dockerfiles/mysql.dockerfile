FROM --platform=linux/amd64 mysql:5.7

MAINTAINER mtg


COPY  ./binance-deploy.sql /docker-entrypoint-initdb.d/binance-deploy.sql
