package main

import (
	"fmt"
	"log"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
)

func main() {
	r := gin.Default()

	// ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(http.StatusOK, "pong")
	})

	// system info
	r.GET("/system", getSystemInfo)

	log.Println("get system info at http://127.0.0.1:8080/system")

	log.Fatal(r.Run())
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
