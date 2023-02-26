FROM golang:1.18-alpine

WORKDIR /opt/api

COPY go.mod .
COPY go.sum .

RUN go mod download

RUN mkdir bin

COPY . .

RUN go build -o ./bin ./cmd/importer/...

ENTRYPOINT ["./bin/importer"]

CMD [ "-file=./cmd/importer/data/names.csv" ]

