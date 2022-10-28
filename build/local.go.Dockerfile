FROM golang:1.17

RUN apt-get update

RUN GO111MODULE=on go get golang.org/x/tools/cmd/goimports

RUN GO111MODULE=on go get github.com/vektra/mockery/v2/.../@v2.10.0

RUN GO111MODULE=on go install github.com/volatiletech/sqlboiler/v4@v4.8.6 && \
    GO111MODULE=on go install github.com/volatiletech/sqlboiler/v4/drivers/sqlboiler-psql@v4.8.6 \
