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

Init sqlite datebase

    --
    -- 由SQLiteStudio v3.1.0 产生的文件 周一 9月 26 18:12:32 2016
    --
    -- 文本编码：UTF-8
    --
    PRAGMA foreign_keys = off;
    BEGIN TRANSACTION;
    
    -- 表：article
    DROP TABLE IF EXISTS article;
    
    CREATE TABLE article (
        id          INTEGER       NOT NULL
                                  PRIMARY KEY AUTOINCREMENT,
        title       VARCHAR (255) NOT NULL
                                  DEFAULT '',
        content     TEXT          NOT NULL,
        author      VARCHAR (255) NOT NULL
                                  DEFAULT '',
        create_time DATETIME      NOT NULL,
        update_time DATETIME      NOT NULL
    );
    
    
    -- 表：tag
    DROP TABLE IF EXISTS tag;
    
    CREATE TABLE tag (
        id         INTEGER       NOT NULL
                                 PRIMARY KEY AUTOINCREMENT,
        name       VARCHAR (255) NOT NULL
                                 DEFAULT '',
        article_id INTEGER       NOT NULL
                                 DEFAULT 0
    );
    
    
    -- 表：user
    DROP TABLE IF EXISTS user;
    
    CREATE TABLE user (
        id    INTEGER       NOT NULL
                            PRIMARY KEY AUTOINCREMENT,
        uname VARCHAR (255) NOT NULL
                            DEFAULT ''
                            UNIQUE,
        upass VARCHAR (255) NOT NULL
                            DEFAULT '',
        email VARCHAR (255) NOT NULL
                            DEFAULT ''
                            UNIQUE
    );
    
    INSERT INTO user (
                         id,
                         uname,
                         upass,
                         email
                     )
                     VALUES (
                         1,
                         'admin',
                         '8c6976e5b5410415bde908bd4dee15dfb167a9c873fc4bb8a81f6f2ab448a918',
                         'nsecgo@gmail.com'
                     );
    
    
    COMMIT TRANSACTION;
    PRAGMA foreign_keys = on;
