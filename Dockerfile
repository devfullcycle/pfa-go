FROM golang:1.19
WORKDIR /app
ENTRYPOINT [ "tail", "-f", "/dev/null" ]