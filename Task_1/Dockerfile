FROM golang:1.20.2-bullseye
COPY . /app
WORKDIR /app
RUN go build -o /build/app
CMD ["sh", "-c", "/build/app"]
