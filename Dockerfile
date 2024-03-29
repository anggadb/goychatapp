FROM golang:latest as builder

LABEL maintainer="Angga Bachtiar <bachtiar.angga@gmail.com>"

WORKDIR /app

COPY go.mod go.sum ./

RUN go mod download

COPY . .

RUN CGO_ENABLED=0 GOOS=linux go build -a -installsuffix cgo -o main .

FROM alpine:latest  

RUN apk --no-cache add ca-certificates

WORKDIR /root/

COPY --from=builder /app/main .

COPY .env /root/

EXPOSE 3000

CMD ["./main"]