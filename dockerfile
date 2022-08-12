# this is a two step docker build to reduce container size

FROM golang:1.16-alpine AS builder 

WORKDIR /app  

ADD . /app/

RUN CGO_ENABLED=1 \
 GOOS=linux \ 
 go mod download  

FROM alpine:latest 

COPY --from=builder /app .

 CMD [ "sh entrypoint.sh" ]