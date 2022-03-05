FROM golang:alpine AS builder
RUN mkdir /app
ADD . /app
WORKDIR /app
COPY . .
RUN go build -o main .


FROM alpine
WORKDIR /app
COPY --from=builder /app/main .

CMD ["/app/main"]
