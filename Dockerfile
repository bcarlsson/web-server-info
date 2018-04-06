FROM golang:alpine
ADD main.go ./
RUN go build -o main
EXPOSE 8080
ENTRYPOINT [ "/main" ]