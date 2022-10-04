FROM golang:alpine

WORKDIR /app

# Go.mod and download
COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /service2

CMD [ "/service2" ]
