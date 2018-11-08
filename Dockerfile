FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

RUN mkdir /app
WORKDIR /app
ADD microTest /app
ADD conf /app/conf
RUN chmod +x /app/microTest

ENTRYPOINT [ "/app/microTest" ]