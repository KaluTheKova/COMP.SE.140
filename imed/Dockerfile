FROM golang:alpine

WORKDIR /app

# Go.mod and download
COPY go.mod ./
RUN go mod download

#COPY *.go ./
COPY . ./

RUN go build -o /imed

CMD [ "/imed" ]
