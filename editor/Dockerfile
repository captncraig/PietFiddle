from golang:1.4

ADD . /go/src/github.com/captncraig/pietfiddle/editor

RUN go get -v github.com/captncraig/pietfiddle/editor
RUN go install github.com/captncraig/pietfiddle/editor

EXPOSE 3000

ENTRYPOINT /go/bin/editor