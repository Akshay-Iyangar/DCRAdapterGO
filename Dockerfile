FROM kiasaki/alpine-golang
WORKDIR /gopath/src/app
ADD . /gopath/src/app/
RUN go get app
ENTRYPOINT ["/gopath/bin/app"]

# Define bash as default command
EXPOSE 12285

COPY hostname.sh /etc/my_init.d/hostname.sh
