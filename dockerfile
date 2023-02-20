FROM golang:1.20.0-alpine
RUN mkdir /app
ADD . /app
WORKDIR /app
RUN go build -o main .
EXPOSE 1338
CMD ["/app/main"]