package main

import (
	"runtime"
	"strconv"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"

	log "github.com/zhangxiaofeng05/systeminfo"
)

func main() {
	getSystemInfo()
	println("------------------")
	getGoInfo()
}

func getSystemInfo() {
	log.Print("操作系统", runtime.GOOS)
	log.Print("CPU类型", runtime.GOARCH)
	log.Print("CPU数量", strconv.Itoa(runtime.NumCPU()))
	is, err := host.Info()
	if err != nil {
		panic(err)
	}
	log.Print("Hostname", is.Hostname)
	log.Print("KernelArch", is.KernelArch)

	log.Debugln("host.Info", is.String())

	ci, err := cpu.Info()
	if err != nil {
		panic(err)
	}
	for i := range ci {
		log.Print("CPU核数", strconv.Itoa(int(ci[i].Cores)))
		log.Print("CPU处理器", ci[i].ModelName)

		log.Debugln("cpu.Info", ci[i].String())
	}
	vm, err := mem.VirtualMemory()
	if err != nil {
		panic(err)
	}
	u := vm.Total / 1024 / 1024 / 1024
	log.Print("总内存(GB)", u)
	log.Print("内存使用比(%)", vm.UsedPercent)
	log.Debugln("mem.VirtualMemory", vm.String())

}

func getGoInfo() {
	log.Print("Goroutine数量", strconv.Itoa(runtime.NumGoroutine()))
}
