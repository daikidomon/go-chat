FROM golang:1.16-buster

RUN mkdir /go/src/app
RUN apt-get update \
  && apt-get install git

RUN go get -u -t github.com/gin-gonic/gin \
  && go get -u -t gopkg.in/olahol/melody.v1 \
  && go get -u -t github.com/kamva/mgm/v3 \
  && go get -u -t go.mongodb.org/mongo-driver/bson \
  && go get  -u -t go.mongodb.org/mongo-driver/mongo \
  && go get  -u -t go.mongodb.org/mongo-driver/mongo/options \
  && go get  -u -t go.mongodb.org/mongo-driver/mongo/readpref

WORKDIR /go/src/app