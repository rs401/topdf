FROM golang:buster
WORKDIR /build/
COPY ./go.mod ./
RUN go mod download
COPY . .
RUN CGO_ENABLED=0 GOOS=linux go build -a -o app .

FROM pandoc/core:latest
# This brings it from 194MB to 817 MB
# RUN apk --no-cache add texlive
# Guess I'll have to look into tinytex
RUN apk --no-cache add perl wget
RUN wget -qO- "https://yihui.org/tinytex/install-bin-unix.sh" | sh
ENV PATH=/root/.TinyTeX/bin/x86_64-linuxmusl/:$PATH
WORKDIR /app/
COPY --from=0 /build/app ./
EXPOSE 8888
ENTRYPOINT ["./app"]