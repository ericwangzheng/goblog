#goblog
A simple blog based on beego

#Install or Update
Run command line:

    go get -u -v github.com/nsecgo/goblog

New a file at the conf directory named app.conf

    appname = goblog
    httpport = 9000
    runmode = dev
    enablexsrf = true
    xsrfkey = 61oATzKXQAGaYdkL5gAmGeAAAuYh7EQnp2XdTP1o
    xsrfexpire = 3600
    cookiesecret = 123456
    sqlitepath = goblog.sqlite
    
Build app

    go build
    
Run app

    ./goblog