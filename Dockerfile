FROM golang:1.19.3-alpine

WORKDIR /app

COPY go.mod ./
COPY go.sum ./

RUN go mod download
COPY . ./
EXPOSE 80
RUN go build -o /bin/main cmd/main/main.go
CMD [ "/bin/main" ]