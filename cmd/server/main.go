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
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

var (
	Port int // server port

	Pprof bool // pprof gin middleware
)

func init() {
	flag.IntVar(&Port, "port", 8080, "server port")
	flag.BoolVar(&Pprof, "pprof", false, "use pprof gin middleware")

	// customize -h param
	flag.Usage = func() {
		fmt.Println("Usage: systeminfo [-h] [-port int] [-pprof bool]")
		flag.PrintDefaults()
	}

	// parse
	flag.Parse()
}

func main() {
	// Create context that listens for the interrupt signal from the OS.
	ctx, stop := signal.NotifyContext(context.Background(), syscall.SIGINT, syscall.SIGTERM)
	defer stop()

	r := gin.Default()
	if Pprof {
		pprof.Register(r)
		log.Printf("get pprof info at http://127.0.0.1:%d/debug/pprof", Port)
	}

	// ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	// system info
	r.GET("/system", getSystemInfo)

	log.Printf("get system info at http://127.0.0.1:%d/system", Port)

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
		u := vm.Total / 1024 / 1024 / 1024
		memMap["total memory(GB)"] = u
		memMap["memory usedPercent(%)"] = vm.UsedPercent
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

	if all || len(ci) != 1 {
		cpuMap["all"] = ci
	} else {
		cpuMap["cpu cores"] = ci[0].Cores
		cpuMap["cpu modelName"] = ci[0].ModelName
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
