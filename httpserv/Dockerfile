FROM golang:alpine

WORKDIR /app

# Go.mod and download
COPY go.mod ./
RUN go mod download

COPY *.go ./

RUN go build -o /service1

EXPOSE 8001
ENV SERVICE2_ADDRESS = "service2://service2:8002"

CMD [ "/service1" ]
