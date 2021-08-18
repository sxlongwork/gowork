package main

import (
	"bufio"
	"fmt"
	"io"
	"net/url"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type LogProcess struct {
	rc    chan string
	wc    chan *Message
	read  Read
	write Write
}

type Read interface {
	read(rc chan<- string)
}
type Write interface {
	write(wc <-chan *Message)
}

type ReadFromFile struct {
	path string
}

type WriteToDb struct {
	dbinfo string
}

type Message struct {
	remoteAddr string
	localtime  time.Time
	reqMethod  string
	reqPath    string
	reqScheme  string
	status     string
	reqBytes   int
}

func (r *ReadFromFile) read(rc chan<- string) {
	// 打开日志文件
	f, err := os.Open(r.path)
	if err != nil {
		panic(fmt.Sprintf("open file failed, %s", err.Error()))
	}
	// 从文件末尾开始读取，不读取历史信息
	// f.Seek(0, 2)
	reader := bufio.NewReader(f)
	for {
		// 逐行读取文件内容
		data, err := reader.ReadBytes('\n')
		if err == io.EOF {
			time.Sleep(500 * time.Millisecond)
			continue
		} else if err != nil {
			panic(fmt.Sprintf("read content err, %s", err.Error()))
		}
		rc <- string(data[:len(data)-1])
	}
}

func (w *WriteToDb) write(wc <-chan *Message) {
	for v := range wc {
		fmt.Println(v)
	}

}

func (l *LogProcess) ParseLog() {

	/*
				'$remote_addr - $remote_user [$time_local] "$request" '
			    '$status $body_bytes_sent "$http_referer" '
			    '"$http_user_agent" "$http_x_forwarded_for"';
		124.64.8.242 - 5e572e694e4d61763b567059273a4d3d [29/Jun/2020:17:19:08 +0800]
		"GET /uploads/common/2019/0909/OSOSSFAP201909091551010719000003.png HTTP/1.1" 200 1898 "-"
		"LaiYueRegenerate/10060 CFNetwork/1126 Darwin/19.5.0" "-"
	*/
	pattern := `(\d)`
	reg := regexp.MustCompile(pattern)

	loc, _ := time.LoadLocation("Asia/Shanghai")
	for v := range l.rc {
		result := reg.FindStringSubmatch(v)
		msg := &Message{}

		msg.remoteAddr = result[0]
		t, err := time.ParseInLocation("02/Jan/2006:15:04:05", result[2], loc)
		if err != nil {
			fmt.Println("parse time field failed, ", result[2], err.Error())
			continue
		}
		msg.localtime = t
		request := strings.Fields(result[3])
		if len(request) != 3 {
			fmt.Println("parse request failed: ", result[3], err.Error())
			continue
		}
		msg.reqMethod = request[0]
		u, _ := url.Parse(request[1])
		msg.reqPath = u.Path
		msg.reqScheme = request[2]

		msg.status = result[4]
		bytes, err := strconv.Atoi(result[5])
		if err != nil {
			fmt.Println("parse failed: ", request[5], err.Error())
			continue
		}
		msg.reqBytes = bytes

		l.wc <- msg
	}

}

func main() {
	r := &ReadFromFile{path: "access.log"}
	w := &WriteToDb{dbinfo: "admin@admin"}
	logp := &LogProcess{
		rc:    make(chan string),
		wc:    make(chan *Message),
		read:  r,
		write: w,
	}
	go logp.read.read(logp.rc)
	go logp.ParseLog()
	go logp.write.write(logp.wc)

	time.Sleep(60 * time.Second)

}
