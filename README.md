#goblog
A simple blog based on beego

#Install or Update
Run command line:

    go get -u -v github.com/nsecgo/goblog

New a file at the conf directory named app.conf

    appname = goblog
    httpport = 9000
    runmode = dev
    sqlitepath = sqlite.db
    
Build app

    go build
    
Run app

    ./goblog