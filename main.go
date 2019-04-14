package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"strings"
	"time"

	"github.com/Unknwon/goconfig"

	"github.com/PuerkitoBio/goquery"
	_ "github.com/go-sql-driver/mysql"
)

type Proxy struct {
	IP      string `json:"ip"`
	Port    string `jsoon:"port"`
	Method  string `json:"method"`
	Addr    string `json:"addr"`
	GetTime interface{}
}

/*
*判断代理是否可以正常使用
 */
func CheckProxy(host, port string) bool {
	proxyUrl, _ := url.Parse(fmt.Sprintf("http://%s:%s", host, port))
	transport := &http.Transport{
		Proxy: http.ProxyURL(proxyUrl),
	}
	c := &http.Client{
		Transport: transport,
		Timeout:   time.Second * 10,
	}
	//http://httpbin.org/get 此网站可以自行修改，用来判断代理是否可以正常使用
	request, err := http.NewRequest("GET", "http://httpbin.org/get", nil)
	if err != nil {
		log.Println(err)
	}
	request.Header.Add("user-agent", "I am Gofer")
	resp, err := c.Do(request)
	if err != nil {
		//log.Println(err)
		return false
	}
	defer resp.Body.Close()
	if resp.StatusCode == 200 {
		//fmt.Printf("[+] http://%s:%s\t%d\n", host, port, resp.StatusCode)
		return true
	} else {
		//fmt.Printf("[-] http://%s:%s\t%d\n", host, port, resp.StatusCode)
		return false
	}
}

/*
*得到代理网站一页的内容，并判断代理是否可以，若可以使用则写入到数据库中
 */

func getResponse(host string, db *sql.DB) {
	c := &http.Client{
		Timeout: time.Second * 10,
	}
	request, err := http.NewRequest("GET", host, nil)
	if err != nil {
		log.Fatal(err)
	}
	request.Header.Add("user-agent", "I am Gofer")
	resp, err := c.Do(request)
	if err != nil {
		log.Println(err)
	}
	defer resp.Body.Close()
	//此处使用goqurey css选择器
	doc, err := goquery.NewDocumentFromReader(resp.Body)
	if err != nil {
		log.Println("create docment error:", err)
		return
	}
	doc.Find("body > div > div.container.mt-4 > div.mt-0.mb-2.table-responsive > table > tbody >tr").Each(func(i int, s *goquery.Selection) {
		IP := s.Find("td").Eq(0).Text()
		Method := s.Find("td").Eq(1).Text()
		Addr := s.Find("td").Eq(3).Text()
		ips := strings.Split(IP, ":")
		proxy := &Proxy{
			IP:      ips[0],
			Port:    ips[1],
			Method:  Method,
			Addr:    Addr,
			GetTime: time.Now(),
		}
		//fmt.Println(proxy.IP, proxy.Port)
		if CheckProxy(proxy.IP, proxy.Port) {
			fmt.Println(IP, Method, Addr, " Write Ok")
			WriteToMysql(proxy, db)
		} else {
			fmt.Println(IP, Method, Addr, " Write False")
		}
	})
}

/*
*	将Proxy结构体里的数据写入到mysql数据库
 */
func WriteToMysql(p *Proxy, db *sql.DB) {
	stmt, _ := db.Prepare("INSERT INTO proxy(ip,port,addr,time) VALUES(?,?,?,?)")
	stmt.Exec(p.IP, p.Port, p.Addr, time.Now())
	stmt.Close()
}

/*
*	创建数据库链接
 */
func LinkToMysql(user, password, ip, port, database string) *sql.DB {
	//fmt.Sprintf("%s:%s@tcp(%s:%s)/%s",user,password,ip,port,database)
	db, err := sql.Open("mysql", fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, ip, port, database))
	if err != nil {
		log.Println("Open Mysql Error : ", err)
		return nil
	}
	return db
}

/*
* 	清除数据库中不能使用的代理
 */
func deleteProxy(db *sql.DB) {
	rows, _ := db.Query("select ip,port from proxy")
	defer rows.Close()
	for rows.Next() {
		//fmt.Println("hello")
		var ip string
		var port string
		if err := rows.Scan(&ip, &port); err != nil {
			log.Fatal(err)
		}
		if CheckProxy(ip, port) {
			fmt.Printf("[+] http://%s:%s\tOk\n", ip, port)
		} else {
			db.Exec("DELETE FROM proxy WHERE ip = ?", ip)
			fmt.Println("DELETE ", ip, port)
		}
	}
}

/*
*执行爬取代理的操作，循环调用getRespone函数来爬取代理地址并保存在mysql中
*代理网站采用了防爬虫机制，所以爬取速度不能太快。为了控制调用的速度，所有操作
*封装到这个函数中 num控制爬取代理多少 例如：num = 2  代理数： 2 * 3 * 50
 */
func Working(num int, done chan int, db *sql.DB) {
	for i := num; i < num+3; i++ {
		getResponse(fmt.Sprintf("http://www.xiladaili.com/gaoni/%d/", i), db)
	}
	done <- num
}

func main() {
	//获取conf.ini 里的基本配置
	conf, err := goconfig.LoadConfigFile("conf.ini")
	if err != nil {
		fmt.Println("Can't Open File , Error: ", err)
	}
	user, _ := conf.GetValue("mysql", "user")
	password, _ := conf.GetValue("mysql", "password")
	ip, _ := conf.GetValue("mysql", "ip")
	port, _ := conf.GetValue("mysql", "port")
	database, _ := conf.GetValue("mysql", "database")

	//创建mysql数据库链接
	db := LinkToMysql(user, password, ip, port, database)
	defer db.Close()

	done := make(chan int)
	for i := 1; i < 10; i++ {
		go Working(i, done, db)
	}
	for i := 1; i < 10; i++ {
		<-done
	}
	//清理不能使用的代理
	deleteProxy(db)
}
