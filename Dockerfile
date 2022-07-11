FROM golang:alpine

WORKDIR /var/www/html/apps/balancing
COPY . /var/www/html/apps/balancing

RUN go build -o main .

CMD ["/var/www/html/apps/balancing/main"]