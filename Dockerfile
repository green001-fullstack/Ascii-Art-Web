FROM golang:1.22

WORKDIR /app

COPY . .

RUN go build -o app

EXPOSE 8001

CMD ["./app"]

LABEL maintainer="Your Name <oladimejiemmanuel2015@gmail.com"
LABEL version="1.0"
LABEL description ="Ascii Art Web"