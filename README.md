# dongfeng
聚合支付，支付系统，golang，go，go语言，前后端齐全（管理后台，商户后台，代理后台，网关，结算，代付，等）
# 目的
接触了很多支付系统，绝大部分都是用java、php写的，本人想着为go献出一份绵薄之力。该开源项目只能用于学习，不准备用于任何的非法商业活动，否则后果自负；可集成到小型的电商平台。
# 安装
详细的安装文档在doc文件夹中。
使用IntelliJ IDEA 导入项目后，若出现了：cannot resolve directory....，删除.idea目录，重新导入即可。

# 关于mysql文件
1：sql文件已经上传。</br>
2：测试超级管理员账号：10086，密码：123456

# 后话
1：因为是模拟了三方的支付系统，所有很多功能都是实现了，但是该系统截图上面的资金都是模拟的。</br>
2：还有一部分功能没有写上去，有需要的可以在上面进行二次开发。


$ cd server
$ go mod tidy
$ go build -o server main.go (windows编译命令为 go build -o server.exe main.go )

# 运行二进制
$ ./server (windows运行命令为 server.exe)

$ cd web
$ npm install
$ npm run dev

activemq
https://zhaobiao666.github.io/middleware/ActiveMQ/

docker run -d --name myactivemq -p 61613:61613 -p 61616:61616 -p 8161:8161 webcenter/activemq
docker run -d --name myactivemq -p 61613:61613 -p 61616:61616 -p 8161:8161 rmohr/activemq
docker run -d --name myactivemq -p 61613:61613 -p 61616:61616 -p 8161:8161 symptoma/activemq

查看WEB管理页面：
浏览器输入http://127.0.0.1:8161/，点击Manage ActiveMQ broker使用默认账号/密码：admin/admin进入查看
