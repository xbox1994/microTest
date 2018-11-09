FROM alpine:latest
RUN apk update && apk add ca-certificates && rm -rf /var/cache/apk/*

WORKDIR /app
ADD microTest .
ADD conf conf
RUN chmod +x microTest

ENTRYPOINT [ "./microTest" ]