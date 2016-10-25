FROM ubuntu
WORKDIR /root
RUN apt update && apt dist-upgrade -y && apt install wget
RUN wget https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz -o go.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go.linux-amd64.tar.gz
RUN echo 'export GOPATH=/root/gopath' >> /etc/profile
RUN echo 'export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin' >> /etc/profile
RUN source /etc/profile
RUN mkdir gopath
RUN go get -u github.com/nsecgo/goblog
RUN cd $GOPATH/src/github.com/nsecgo/goblog && cp goblog.sqlite.example goblog.sqlite \
    && cp conf/app.conf.example conf/app.conf
EXPOSE 80
ENTRYPOINT ["goblog"]