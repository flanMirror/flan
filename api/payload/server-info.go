package payload

import (
	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/mem"
	"log"
	"os"
	"random.chars.jp/git/misskey/api/structs"
	"random.chars.jp/git/misskey/config"
	"random.chars.jp/git/misskey/data"
	"runtime"
	"time"
)

// Node is not what we're using
const Node = "false"

// NetInterface is insane and completely unnecessary and relies on too much platform specific stuff,
// so I set it to whatever I have on my FreeBSD workstation. If anyone knows of a good reason to properly
// populate this and a cross-platform way to fetch it, do let me know. The parsing route thing in upstream
// is absolutely insane and caused me to explode for hours and end up using a very ugly workaround for it,
// so I wish to not do that in this implementation as it is meant to address flaws as well.
const NetInterface = "igb0"

var (
	// Hostname is set in init
	Hostname string
	// OS is set in init
	OS string
	// PSQL is set in main.go
	PSQL string
	// Redis is set in main.go
	Redis string
	// CPUModel is set in init
	CPUModel string
	// CPUCores is set in init
	CPUCores int
	// MemTotal is set in init (for some insane reason there is no allocated, only total)
	MemTotal int
	// FSTotal is set in init
	FSTotal int
)

func init() {
	if h, err := os.Hostname(); err != nil {
		log.Fatalf("error getting hostname: %s", err)
	} else {
		Hostname = h
	}
	OS = runtime.GOOS
	if info, err := cpu.Info(); err != nil {
		log.Fatalf("error getting CPU info: %s", err)
	} else {
		if len(info) > 0 {
			CPUModel = info[0].ModelName
			CPUCores = int(info[0].Cores)
		} else {
			log.Print("this system seems to have no CPU")
		}
	}
	if vm, err := mem.VirtualMemory(); err != nil {
		log.Fatalf("error getting virtual memory stat: %s", err)
	} else {
		MemTotal = int(vm.Total)
	}
	if u, err := disk.Usage("."); err != nil {
		log.Fatalf("error getting disk total: %s", err)
	} else {
		FSTotal = int(u.Total)
	}
}

var serverInfoDuration = time.Minute * 30

var (
	AdminServerInfo = data.NewExpire(serverInfoDuration, func() interface{} {
		info := serverInfoRefresh().(structs.ServerInfo)
		info.OS = OS
		info.Node = Node
		info.PSQL = PSQL
		info.Redis = Redis
		info.Net = &struct {
			Interface string `json:"interface"`
		}{Interface: NetInterface}
		return info
	})
	ServerInfo = data.NewExpire(serverInfoDuration, serverInfoRefresh)
)

func serverInfoRefresh() interface{} {
	info := structs.ServerInfo{
		Machine: Hostname,
	}
	info.CPU.Model = CPUModel
	info.CPU.Cores = CPUCores
	info.Mem.Total = MemTotal
	info.FS.Total = FSTotal
	if u, err := disk.Usage("."); err != nil {
		if config.Log.Verbose {
			log.Printf("error getting disk usage: %s", err)
		}
	} else {
		info.FS.Used = int(u.Used)
	}

	return info
}