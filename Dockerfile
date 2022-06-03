FROM golang:1.18.2-alpine

WORKDIR /app/boilerplate
COPY ./boilerplate ./

RUN apk add --no-cache libc6-compat
RUN chmod +x boilerplate

EXPOSE 8080
CMD [ "./boilerplate" ]
