FROM alpine:3.4

RUN apk update && apk add curl git mercurial bzr 'go=1.5.3-r0' && rm -rf /var/cache/apk/*

ENV GOROOT /usr/lib/go
ENV GOPATH /gopath
ENV GOBIN /gopath/bin
ENV PATH $PATH:$GOROOT/bin:$GOPATH/bin


WORKDIR /gopath/src/app
ADD . /gopath/src/app/
RUN go get app
ENTRYPOINT ["/gopath/bin/app"]

# Define bash as default command
EXPOSE 12285

COPY hostname.sh /etc/my_init.d/hostname.sh
