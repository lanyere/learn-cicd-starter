FROM --platform=linux/arm64 debian:stable-slim

RUN apt-get update && apt-get install -y ca-certificates

GOOS=linux GOARCH=arm64 go build -o notely .

ADD notely /usr/bin/notely

CMD ["notely"]
