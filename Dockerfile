FROM golang:1.9.7-stretch

RUN apt-get update && apt-get install git -y 
RUN go get github.com/tools/godep && go get github.com/mattn/goveralls
RUN cd /go/src && git clone https://github.com/hawkingrei/redp.git --depth 10 && cd redp && godep restore
