FROM golang:alpine

WORKDIR /app

# Go.mod and download
COPY go.mod ./
RUN go mod download

#COPY *.go ./
COPY . /app

RUN go build -o /app/gateway

CMD [ "/app/gateway" ]
