FROM nsecgo/ubuntu-go
WORKDIR /root
RUN go get -u github.com/nsecgo/goblog
RUN cd $GOPATH/src/github.com/nsecgo/goblog && cp goblog.sqlite.example goblog.sqlite \
    && cp conf/app.conf.example conf/app.conf
EXPOSE 80
ENTRYPOINT ["goblog"]