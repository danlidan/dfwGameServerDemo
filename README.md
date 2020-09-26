大复翁后台服务器demo
基于golang的leaf框架

leaf框架直接已经直接加入项目，即刻修改即刻享受
需要先安装golang, protobuf, golang/protobuf包

mySql库使用方式https://github.com/jmoiron/sqlx，配置于bin/conf/mysql.json

redis库github.com/gomodule/redigo/redis，同样有配置文件
注意及时close连接防止泄漏

启动方法（win10）：
1. 执行 go install server
2. 在bin路径中执行 ./server.exe