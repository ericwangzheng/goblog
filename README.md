**1.配置conf/app.conf**

**2.运行**
```
./goblog-linux-amd64
```

**3.添加用户**
```
$echo -n "123456" | sha256sum
8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92  -
$sqlite3 sqlite.db           
SQLite version 3.20.1 2017-08-24 16:21:36
Enter ".help" for usage hints.
sqlite> INSERT INTO user ('uname','upass') VALUES ('admin','8d969eef6ecad3c29a3a629280e686cf0c3f5d5a86aff3ca12020c923adc6c92');
sqlite> .exit
```

**4.登录**
```
http://site-url/login
```