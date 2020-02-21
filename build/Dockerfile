FROM golang

RUN mkdir -p /go/src/github.com/SeamPay/gokita

COPY ./ /go/src/github.com/SeamPay/gokita

RUN go get  -t -v ./...
RUN go get  github.com/canthefason/go-watcher
RUN go install github.com/canthefason/go-watcher/cmd/watcher

ENTRYPOINT  watcher -run github.com/SeamPay/gokita/cmd  -watch github.com/SeamPay/gokita
