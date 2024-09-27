FROM ubuntu:22.04
RUN apt-get update && apt-get install -y golang ca-certificates
WORKDIR /app
COPY . /app
RUN go env -w GOPROXY=https://goproxy.cn,direct
RUN go build -o out/mainApp src/main.go