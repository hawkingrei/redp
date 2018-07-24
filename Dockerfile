FROM golang:1.9.7-stretch

RUN go get github.com/tools/godep && go get github.com/mattn/goveralls
RUN mkdir -p /go/src/github.com/hawkingrei && cd /go/src/github.com/hawkingrei && git clone https://github.com/hawkingrei/redp.git --depth 10 && cd redp && godep restore
WORKDIR $GOPATH/src/github.com/hawkingrei/redp
