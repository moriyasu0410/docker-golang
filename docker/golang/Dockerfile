FROM golang:1.7.4-alpine

# apk
RUN apk update \
    && apk --no-cache add git curl

# glide
ENV GLIDE_VERSION 0.12.3

ENV GLIDE_DOWNLOAD_URL https://github.com/Masterminds/glide/releases/download/v${GLIDE_VERSION}/glide-v${GLIDE_VERSION}-linux-amd64.tar.gz

RUN curl -fsSL "$GLIDE_DOWNLOAD_URL" -o glide.tar.gz \
    && tar -xzf glide.tar.gz \
    && mv linux-amd64/glide /usr/bin/ \
    && rm -r linux-amd64 \
    && rm glide.tar.gz

# Golang
RUN go get github.com/golang/lint/golint \
    && go get github.com/motemen/gore \
    && go get github.com/nsf/gocode \
    && go get github.com/k0kubun/pp \
    && go get golang.org/x/tools/cmd/godoc

WORKDIR /go/src/github.com/mediba-moritake/docker-golang

CMD ["/sbin/init"]
