FROM golang:1.20.0-alpine3.14 
COPY . /app
WORKDIR /app
RUN go build -o /build/app
CMD ["sh", "-c", "/build/app"]
