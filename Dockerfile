FROM    golang:latest

RUN mkdir /build
WORKDIR /build

RUN export GO111MODULE=on
RUN go get github.com/Imomali1/video-service/cmd
RUN cd /build && git clone