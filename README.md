### 代理池想法源头
```
才开始学习web安全的时候用的Sqlmap扫描网站后IP总是被别人家的WAF给拦
 
截了，查百度以后有大佬自己写的代理池供Sqlmap使用，但是源代码并没有开放，那

时也不知道其中的原理，随着学的东西越多，自己也明白其中的原理，所以自己写了

一个代理池，供Sqlmap获取能够使用的IP。这个代理池只是获取代理并存储到

Mysql数据库中。下一步我会写一个HTTP代理服务器，用来转发sqlmap的数据包，

这样再也不用担心IP被封了。
```
<!--more-->
### Golang初遇

```
从最开始接触编程，就感受到编程的魅力。在遇到golang之前，一直迷惑到底
    
哪个编程语言适合自己，当深入了解Golang之后，就下定决心认真的学习这们
   
语言。同时也感叹创造Golang的人，真的很厉害。通过一个月的入门，我写了
    
这个小程序，证明一下自己这一个月的学习结果，以后再接再厉。
```
		
### Go环境配置
	
开始敲代码之前配置好自己的环境：
1.安装Golang的环境
2.配置好GoPath
3.安装Mysql数据库
4.下载第三方开发库
	
    从github中获取第三方库：
    github.com/Unknwon/goconfig
	github.com/PuerkitoBio/goquery
	github.com/go-sql-driver/mysql
    
这是我的GOPATH:
![](/images/one.png)
          
        
### 运行结果

配置conf.ini

```
[mysql]
    user = root					//数据库用户名
    password = zad090924.			//数据库密码
    ip = 127.0.0.1 				//数据库地址
    port = 3306					//数据库端口
    database = test				//所使用的数据库
```


![](/images/two.png)


代码分享在 [Github](https://github.com/hatmagic/GolangProxyPool)