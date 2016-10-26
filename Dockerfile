FROM ubuntu

MAINTAINER nsecgo <nsecgo@gmail.com>

ENV GOLANG_VERSION 1.7.3
ENV GOLANG_SRC_URL https://storage.googleapis.com/golang/go$GOLANG_VERSION.linux-amd64.tar.gz

RUN apt update && apt install -y \
	gcc \
    	wget \
	git \
    	libc-dev \
	&& wget -q "$GOLANG_SRC_URL" -O golang.tar.gz \
	&& tar -C /usr/local -xzf golang.tar.gz \
	&& rm golang.tar.gz
    
ENV GOPATH /go
ENV PATH $GOPATH/bin:/usr/local/go/bin:$PATH

RUN mkdir -p "$GOPATH/src" "$GOPATH/bin" && chmod -R 777 "$GOPATH"
WORKDIR $GOPATH
RUN go get -u github.com/nsecgo/goblog
RUN cd $GOPATH/src/github.com/nsecgo/goblog && cp goblog.sqlite.example goblog.sqlite \
        && cp conf/app.conf.example conf/app.conf \
        && cp $GOPATH/bin/goblog $GOPATH/src/github.com/nsecgo/goblog/goblog
EXPOSE 80
ENTRYPOINT ["src/github.com/nsecgo/goblog/goblog"]
