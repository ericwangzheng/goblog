FROM ubuntu
WORKDIR /root
RUN apt update && apt install wget -y
RUN wget https://storage.googleapis.com/golang/go1.7.3.linux-amd64.tar.gz -O go.linux-amd64.tar.gz
RUN tar -C /usr/local -xzf go.linux-amd64.tar.gz && rm go.linux-amd64.tar.gz
RUN echo 'export GOPATH=/root/gopath' >> /etc/profile
RUN echo 'export PATH=$PATH:/usr/local/go/bin:$GOPATH/bin' >> /etc/profile
RUN ["/bin/bash","-c","source /etc/profile"]
RUN mkdir gopath
RUN ["/bin/bash","-c","go get -u github.com/nsecgo/goblog"]
RUN cd $GOPATH/src/github.com/nsecgo/goblog && cp goblog.sqlite.example goblog.sqlite \
    && cp conf/app.conf.example conf/app.conf
EXPOSE 80
ENTRYPOINT ["/bin/bash","-c","goblog"]