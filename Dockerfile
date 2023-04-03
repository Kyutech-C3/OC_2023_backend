FROM golang:1.19.3-alpine
RUN mkdir -p /go/src/api
COPY . /go/src/api
WORKDIR /go/src/api

RUN go install

RUN go build -o /go/bin/air github.com/cosmtrek/air

EXPOSE 8000

CMD [ "air" ]