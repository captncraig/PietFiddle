from golang:1.4

ADD . /go/src/github.com/captncraig/pietfiddle
WORKDIR /go/src/github.com/captncraig/pietfiddle/editor
RUN go get -v 
RUN go build

EXPOSE 3000

ENTRYPOINT /go/src/github.com/captncraig/pietfiddle/editor/editor