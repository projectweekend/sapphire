FROM golang
COPY ./sapphire /go/src/github.com/projectweekend/sapphire
WORKDIR /go/src/github.com/projectweekend/sapphire
RUN go get && go build
ENTRYPOINT ["sapphire"]
