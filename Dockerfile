FROM golang:1.16.3-alpine
RUN mkdir /app
ADD . /app/
WORKDIR /app
