package main

import (
	"fmt"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/disk"
	"github.com/shirou/gopsutil/host"
	"github.com/shirou/gopsutil/load"
	"github.com/shirou/gopsutil/mem"
	"math"
	"net/http"
	"strconv"
)

func dealwithErr(err error) {
	if err != nil {
		fmt.Println(err)
		//os.Exit(-1)
	}
}

func GetHardwareData(w http.ResponseWriter, r *http.Request) {

	// host or machine kernel, uptime, platform Info
	hostStat, err := host.Info()
	dealwithErr(err)
	//fmt.Println("hostStat", hostStat)

	html := "<html><br>===================== HOST INFORMATION =============================<br><br>"
	html = html + "OS: " + hostStat.OS + "<br>"
	html = html + "Hostname: " + hostStat.Hostname + "<br>"
	html = html + "Platform: " + hostStat.Platform + "<br>"
	html = html + "Platform Family: " + hostStat.PlatformFamily + "<br>"
	html = html + "Number of processes running: " + strconv.FormatUint(hostStat.Procs, 10) + "<br>"
	html = html + "Uptime: " + strconv.FormatUint(hostStat.Uptime, 10) + "<br>"
	html = html + "Host ID(uuid): " + hostStat.HostID + "<br>"

	// memory
	vmStat, err := mem.VirtualMemory()
	dealwithErr(err)
	//fmt.Println("VirtualMemory", vmStat)

	html = html + "<br><br>===================== Memory RAM INFORMATION =============================<br><br>"

	html = html + "Total memory in bytes: " + fmt.Sprint(vmStat.Total) + " and " + bytesToSize(vmStat.Total) + "<br>"
	html = html + "Used memory in bytes: " + fmt.Sprint(vmStat.Used) + " and " + bytesToSize(vmStat.Used) + "<br>"
	html = html + "Unused memory in bytes: " + fmt.Sprint(vmStat.Available) + " and " + bytesToSize(vmStat.Available) + "<br>"
	html = html + "Free memory  in bytes: " + fmt.Sprint(vmStat.Free) + " and " + bytesToSize(vmStat.Free) + " <br>"
	html = html + "Percentage used memory: " + fmt.Sprint(vmStat.UsedPercent) + " and " + strconv.FormatFloat(vmStat.UsedPercent, 'f', 2, 64) + "%<br>"

	// Disk
	diskStat, err := disk.Usage("/")
	dealwithErr(err)
	//diskPartitions, err := disk.Partitions(true)
	//diskPartitions, err := disk.IOCounters()
	//dealwithErr(err)
	//fmt.Println("diskStat", diskStat)
	//fmt.Println("diskPartitions", diskPartitions)

	html = html + "<br><br>===================== DISK SPACE INFORMATION =============================<br><br>"
	html = html + "Total disk space in bytes: " + fmt.Sprint(diskStat.Total) + " and " + bytesToSize(diskStat.Total) + " <br>"
	html = html + "Used disk space in bytes: " + fmt.Sprint(diskStat.Used) + " and " + bytesToSize(diskStat.Used) + " <br>"
	html = html + "Free disk space in bytes`: " + fmt.Sprint(diskStat.Free) + " and " + bytesToSize(diskStat.Free) + " <br>"
	html = html + "Percentage disk space usage: " + fmt.Sprint(diskStat.UsedPercent) + "and " + strconv.FormatFloat(diskStat.UsedPercent, 'f', 2, 64) + "%<br>"

	// cpu - get CPU number of cores and speed
	cpuStat, err := cpu.Info()
	dealwithErr(err)
	percentage, err := cpu.Percent(0, true)
	dealwithErr(err)

	//fmt.Println("cpuStat", cpuStat)
	//fmt.Println("cpuStat percentage ", percentage)

	html = html + "<br><br>===================== CPU INFORMATION =============================<br><br>"
	html = html + "CPU index number: " + strconv.FormatInt(int64(cpuStat[0].CPU), 10) + "<br>"
	html = html + "VendorID: " + cpuStat[0].VendorID + "<br>"
	html = html + "Family: " + cpuStat[0].Family + "<br>"
	html = html + "Number of cores: " + strconv.FormatInt(int64(cpuStat[0].Cores), 10) + "<br>"
	html = html + "Model Name: " + cpuStat[0].ModelName + "<br>"
	html = html + "Speed: " + strconv.FormatFloat(cpuStat[0].Mhz, 'f', 2, 64) + " MHz <br>"

	for idx, cpupercent := range percentage {
		html = html + "Current CPU utilization: [" + strconv.Itoa(idx) + "] " + strconv.FormatFloat(cpupercent, 'f', 2, 64) + "%<br>"
	}

	// Load
	loadStat, err := load.Avg()
	dealwithErr(err)
	loadMisc, err := load.Misc()
	dealwithErr(err)
	//loadMiscWithContext, err := load.MiscWithContext(loadMisc.ProcsRunning);
	//dealwithErr(err)
	fmt.Println("loadStat", loadStat)
	fmt.Println("loadMisc", loadMisc)
	//fmt.Println("loadMiscWithContext", loadMiscWithContext)

	html = html + "</html>"

	w.Write([]byte(html))

}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", GetHardwareData)

	http.ListenAndServe(":8001", mux)

}

func bytesToSize(bytes uint64) string {
	sizes := []string{"Bytes", "KB", "MB", "GB", "TB"}
	if bytes == 0 {
		return fmt.Sprint(float64(0), "bytes")
	} else {
		var bytes1 = float64(bytes)
		var i = math.Floor(math.Log(bytes1) / math.Log(1024))
		var count = bytes1 / math.Pow(1024, i)
		var j = int(i)
		var val = fmt.Sprintf("%.1f", count)
		//fmt.Println("bytes, count,  val........", bytes, count,  val);
		return fmt.Sprint(val, sizes[j])
		//return fmt.Sprint(float64(count) , sizes[j])
		//return fmt.Sprintf("%.2f", float64(count) )
	}
}

//reference
//https://www.socketloop.com/tutorials/golang-get-hardware-information-such-as-disk-memory-and-cpu-usage
//https://stackoverflow.com/questions/15900485/correct-way-to-conver
