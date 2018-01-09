如何使用go发送邮件。
在net/smtp中为我们提供了相关的功能
首先什么是smtp
SMTP(Simple Mail Transfer Protocol)即简单邮件传输协议，
它是一组用于由源地址到目的地址传送邮件的规则，由它来控制信件的
中转方式。SMTP协议属于TCP/IP协议族，它帮助每台计算机在发送或
中转信件时找到下一个目的地。通过SMTP所指定的服务器，就可以把
E-mail寄到收件人的服务器上了。SMTP服务器则是遵循SMTP协议的
发送邮件服务器，用来发送或者中转发出的电子邮件

send mail error: 535 Error
http://service.mail.qq.com/cgi-bin/help?subtype=1&&id=28&&no=1001256
你开始怀疑你的代码有问题，又开始怀疑你的qq邮箱的密码有问题，但是最终都无功而返。
这个时候需要对你的qq邮箱进行设置。

解决方案就是把password改为授权码！！！！！
什么是授权码？
授权码是QQ邮箱推出的，用于登录第三方客户端的专用密码。
适用于登录以下服务：POP3/IMAP/SMTP/Exchange/CardDAV/CalDAV服务。
问题：如何通过同一host发送不同的邮箱