package utils

import (
	"fmt"
	"hash/crc32"
	"math/rand"
	"sort"
	"time"
)

type HttpServes []*HttpServer

func (h HttpServes) Len() int {
	return len(h)
}

func (h HttpServes) Less(i, j int) bool {
	return h[i].Cweight > h[j].Cweight
}

func (h HttpServes) Swap(i, j int) {
	h[i], h[j] = h[j], h[i]
}

//HttpServer
type HttpServer struct {
	Host   string
	Weight int
	Cweight int
}

//LoadBalancing
type LoadBalancing struct {
	HttpServers  HttpServes
	CurrentIndex int
}

func (this *LoadBalancing) AddHttpServer(server *HttpServer) {
	this.HttpServers = append(this.HttpServers, server)
}

func (this *LoadBalancing) SelectServerByRand() *HttpServer { //随机算法
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(this.HttpServers))
	return this.HttpServers[index]
}

func (this *LoadBalancing) SelectServerByIpHash(ip string) *HttpServer { //iphash算法
	index := int(crc32.ChecksumIEEE([]byte(ip))) % len(this.HttpServers)
	return this.HttpServers[index]
}

func (this *LoadBalancing) SelectServerByWeightRand() *HttpServer { //加权随机算法
	rand.Seed(time.Now().UnixNano())
	index := rand.Intn(len(ServerIndexs))
	return this.HttpServers[ServerIndexs[index]]
}

func (this *LoadBalancing) SelectServerByWeightRand2() *HttpServer { //加权随机算法 改良版
	rand.Seed(time.Now().UnixNano())
	sumList := make([]int, len(this.HttpServers))
	sum := 0

	for i := 0; i < len(this.HttpServers); i++ {
		sum += this.HttpServers[i].Weight
		sumList[i] = sum
	}

	read := rand.Intn(sum)

	for index, value := range sumList {
		if read < value {
			return this.HttpServers[index]
		}
	}

	return this.HttpServers[0]
}

func (this *LoadBalancing) SelectServerByPolling() *HttpServer { //轮询算法
	server := this.HttpServers[this.CurrentIndex]
	this.CurrentIndex = (this.CurrentIndex + 1) % len(this.HttpServers)
	return server
}

func (this *LoadBalancing) SelectServerByWeightPolling() *HttpServer { //加权轮询算法
	server := this.HttpServers[ServerIndexs[this.CurrentIndex]]
	this.CurrentIndex = (this.CurrentIndex + 1) % len(ServerIndexs)
	return server
}

func (this *LoadBalancing) SelectServerByWeightPolling2() *HttpServer { //加权轮询算法，使用区间
	server := this.HttpServers[0]
	sum := 0
	for i := 0; i < len(this.HttpServers); i++ {
		sum += this.HttpServers[i].Weight //3 4 5
		//fmt.Printf("CurrentIndex: %d, sum: %d\n", this.CurrentIndex, sum)
		if this.CurrentIndex < sum {
			server = this.HttpServers[i]
			if this.CurrentIndex == (sum -1) && i != len(this.HttpServers) - 1 {
				this.CurrentIndex++
			}else {
				this.CurrentIndex = (this.CurrentIndex + 1) % sum
			}
			fmt.Println(this.CurrentIndex)
			break
		}
	}
	return server
}


//平滑加权轮询
func (this *LoadBalancing) SelectServerByWeightPolling3() *HttpServer {
	for _,server := range this.HttpServers{
		server.Cweight = server.Cweight + server.Weight
	}

	//排序
	sort.Sort(this.HttpServers)
	//得出最大的Cweight
	max := this.HttpServers[0]
	max.Cweight = max.Cweight - SumWeight

	test := ""
	for _,server := range this.HttpServers{
		test += fmt.Sprintf("%d,", server.Cweight)
	}
	fmt.Println(test)
	return max
}

//NewLoadBalancing
func NewLoadBalancing() *LoadBalancing {
	return &LoadBalancing{HttpServers: make([]*HttpServer, 0)}
}

func NewHttpServer(host string, weight int) *HttpServer {
	return &HttpServer{Host: host, Weight: weight, Cweight: 0}
}

var LB *LoadBalancing
var ServerIndexs []int
var SumWeight int

func init() {
	LB = NewLoadBalancing()
	LB.AddHttpServer(NewHttpServer("http://localhost:9091", 3))
	LB.AddHttpServer(NewHttpServer("http://localhost:9092", 1))
	LB.AddHttpServer(NewHttpServer("http://localhost:9093", 1))

	for index, server := range LB.HttpServers {
		if server.Weight > 0 {
			for i := 0; i < server.Weight; i++ {
				ServerIndexs = append(ServerIndexs, index)
			}
		}

		SumWeight += server.Weight

	}
}
