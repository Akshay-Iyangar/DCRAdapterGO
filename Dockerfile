FROM kiasaki/alpine-golang
WORKDIR /gopath/src/app
ADD . /gopath/src/app/
RUN go get app
ENTRYPOINT ["/gopath/bin/app"]
# Install base packages
RUN apk update && apk upgrade && \
    apk-install curl wget bash tree && \
    echo -ne "Alpine Linux 3.2 image. (`uname -rsv`)\n" >> /root/.built

# Define bash as default command
EXPOSE 12285

COPY hostname.sh /etc/my_init.d/hostname.sh
