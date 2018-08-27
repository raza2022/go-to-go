package main

import ()
import (
	"fmt"
	"github.com/c9s/goprocinfo/linux"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func getCPUSample() (idle, total uint64) {
	contents, err := ioutil.ReadFile("/proc/stat")
	if err != nil {
		return
	}
	lines := strings.Split(string(contents), "\n")
	for _, line := range lines {
		fields := strings.Fields(line)
		if fields[0] == "cpu" {
			numFields := len(fields)
			for i := 1; i < numFields; i++ {
				val, err := strconv.ParseUint(fields[i], 10, 64)
				if err != nil {
					fmt.Println("Error: ", i, fields[i], err)
				}
				total += val // tally up all the numbers to get total ticks
				if i == 4 {  // idle is the 5th field in the cpu line
					idle = val
				}
			}
			return
		}
	}
	return
}

func main() {

	stat, err := linuxproc.ReadStat("/proc/stat")
	if err != nil {
		log.Fatal("stat read fail")
	}

	for _, s := range stat.CPUStats {
		// s.User
		// s.Nice
		// s.System
		// s.Idle
		// s.IOWait
	}

	// stat.CPUStatAll
	// stat.CPUStats
	// stat.Processes
	// stat.BootTime
	// ... etc
}
