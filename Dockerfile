FROM golang:1.17

WORKDIR /go/src/app
COPY . .

RUN apt-get update
RUN apt-get -y install libportmidi-dev
RUN go get -d -v ./...
RUN go install -v ./...
RUN go build
RUN ls
#COPY /go/src/app/playyoshimi .
CMD ["playyoshimi"]