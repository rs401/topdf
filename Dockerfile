FROM golang:buster
WORKDIR /build/
COPY ./go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM alpine:latest
RUN apk update
RUN apk --no-cache add libreoffice
RUN apk --no-cache add font-misc-misc terminus-font ttf-inconsolata ttf-dejavu font-noto ttf-font-awesome font-noto-extra
ENV PORT=8888
WORKDIR /app/
COPY --from=0 /build/app ./
EXPOSE 8888
ENTRYPOINT ["./app"]