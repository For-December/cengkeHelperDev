FROM ubuntu:22.04
RUN apt-get update && apt-get install -y golang
WORKDIR /app
COPY . /app
RUN go build -o out/mainApp src/main.go