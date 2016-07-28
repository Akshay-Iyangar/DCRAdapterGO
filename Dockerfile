FROM artifactory.secureserver.net:10000/jdk:8u74
MAINTAINER John Rudolf Lewis <JohnRLewis@godaddy.com>

RUN apt-get update && apt-get install --no-install-recommends -y \
    ca-certificates \
    curl \
    mercurial \
    git-core
RUN curl -s https://storage.googleapis.com/golang/go1.4.2.linux-amd64.tar.gz | tar -v -C /usr/local -xz

ENV GOPATH /go
ENV GOROOT /usr/local/go
ENV PATH /usr/local/go/bin:/go/bin:/usr/local/bin:$PATH

WORKDIR /go/src/app
ADD . /go/src/app/
RUN go get app
ENTRYPOINT ["/go/bin/app"]

CMD ["bash"]
