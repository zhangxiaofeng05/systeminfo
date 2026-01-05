package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"net/http"
	"os/signal"
	"runtime"
	"syscall"
	"time"

	"github.com/gin-contrib/pprof"
	"github.com/gin-gonic/gin"
	"github.com/prometheus/client_golang/prometheus"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

var (
	Port int // server port

	Pprof bool // pprof gin middleware
)

func init() {
	flag.IntVar(&Port, "port", 8080, "server port")
	flag.BoolVar(&Pprof, "pprof", false, "use pprof gin middleware")

	// parse
	flag.Parse()
}

func main() {
	// log init
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	version := getVersion()
	log.Printf("version: %s", version)

	localHostPre := "http://"
	localHost := fmt.Sprintf("127.0.0.1:%d", Port)

	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := gin.Default()
	if Pprof {
		pprof.Register(r)
		log.Printf("get pprof info at %s:%d/debug/pprof", localHost, Port)
	}

	// Register the metrics with the Prometheus collector.
	prometheus.MustRegister(requests)
	prometheus.MustRegister(requestDuration)

	r.Use(prometheusMiddleware())

	// ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})
	r.GET("/version", func(c *gin.Context) {
		c.JSON(http.StatusOK, version)
	})

	// system info
	r.GET("/system", getSystemInfo)

	// home
	r.GET("/", func(c *gin.Context) {
		// 获取不到scheme的相关issue: https://github.com/gin-gonic/gin/issues/1233
		if c.Request.TLS != nil {
			localHostPre = "https://"
		}
		if c.Request.Host != "" {
			localHost = c.Request.Host
		}
		res := make(map[string]map[string]string)

		allUrl := make(map[string]string)
		allUrl["/ping"] = fmt.Sprintf("%s%s%s", localHostPre, localHost, "/ping")
		allUrl["/version"] = fmt.Sprintf("%s%s%s", localHostPre, localHost, "/version")
		allUrl["/system"] = fmt.Sprintf("%s%s%s", localHostPre, localHost, "/system")
		allUrl["/"] = fmt.Sprintf("%s%s%s", localHostPre, localHost, "/")

		res["all url"] = allUrl
		c.JSON(http.StatusOK, res)
	})

	log.Printf("server listening at %s%s", localHostPre, localHost)

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", Port),
		Handler: r,
	}

	// Initializing the server in a goroutine so that
	// it won't block the graceful shutdown handling below
	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("listen: %s\n", err)
		}
	}()

	// Listen for the interrupt signal.
	<-ctx.Done()

	// Restore default behavior on the interrupt signal and notify user of shutdown.
	stop()
	log.Println("shutting down gracefully, press Ctrl+C again to force")

	// The context is used to inform the server it has 5 seconds to finish
	// the request it is currently handling
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := srv.Shutdown(ctx); err != nil {
		log.Fatal("Server forced to shutdown: ", err)
	}

	log.Println("Server exiting")

}

func getSystemInfo(c *gin.Context) {
	// prometheus
	// Increment the request counter.
	requests.Inc()

	// Record the request duration.
	start := time.Now()
	defer func() {
		requestDuration.Observe(time.Since(start).Seconds())
	}()

	// Simulate processing time.
	//time.Sleep(time.Duration(rand.Intn(1000)) * time.Millisecond)

	type params struct {
		All bool `form:"all" binding:"-"`
	}

	var p params
	err := c.ShouldBindQuery(&p)
	if err != nil {
		log.Printf("params err:%v", err)
		c.JSON(http.StatusInternalServerError, fmt.Errorf("params err:%v", err).Error())
		return
	}

	mp := make(map[string]any)
	if !p.All {
		mp[c.FullPath()+"?all=true"] = "get all system info"
	}

	mp["golang"] = getGo()
	mp["host"] = getHost(p.All)
	mp["cpu"] = getCpu(p.All)
	mp["memory"] = getMemory(p.All)

	c.JSON(http.StatusOK, mp)
}

func getMemory(all bool) map[string]any {
	memMap := make(map[string]any)
	vm, err := mem.VirtualMemory()
	if err != nil {
		log.Printf("get memory info err:%v", err)
		memMap["error"] = fmt.Errorf("get memory info err:%v", err)
		return memMap
	}
	if all {
		memMap["all"] = vm
	} else {
		memMap["Total Memory(MB)"] = vm.Total / 1024 / 1024
		memMap["Free Memory(MB)"] = vm.Free / 1024 / 1024
		memMap["Available Memory(MB)"] = vm.Available / 1024 / 1024
		memMap["Used Memory(MB)"] = vm.Used / 1024 / 1024
		memMap["Used Percent(%)"] = vm.UsedPercent
	}
	return memMap
}

func getCpu(all bool) map[string]any {
	cpuMap := make(map[string]any)
	ci, err := cpu.Info()
	if err != nil {
		log.Printf("get cpu info err:%v", err)
		cpuMap["error"] = fmt.Errorf("get cpu info err:%v", err)
		return cpuMap
	}

	if all {
		cpuMap["all"] = ci
	} else {
		for i, c := range ci {
			cpuMap[fmt.Sprintf("index:%d cpu cores", i)] = c.Cores
			cpuMap[fmt.Sprintf("index:%d cpu modelName", i)] = c.ModelName
		}
	}
	return cpuMap
}

func getHost(all bool) map[string]any {
	hostMap := make(map[string]any)
	is, err := host.Info()
	if err != nil {
		log.Printf("get host info err:%v", err)
		hostMap["error"] = fmt.Errorf("get host info err:%v", err)
		return hostMap
	}
	if all {
		hostMap["all"] = is
	} else {
		hostMap["hostname"] = is.Hostname
		hostMap["kernelArch"] = is.KernelArch
	}
	return hostMap
}

func getGo() map[string]any {
	goMap := make(map[string]any)
	goMap["os"] = runtime.GOOS
	goMap["arch"] = runtime.GOARCH
	goMap["cpu num"] = runtime.NumCPU()
	goMap["goroutine num"] = runtime.NumGoroutine()
	return goMap
}
