FROM golang:1.16-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download

COPY . ./

RUN go mod download golang.org/x/sys
RUN go build -o /server

EXPOSE 4040  

CMD ["/server", "redis"]
