package status

import (
	"encoding/json"
	"fmt"
	"os"
	"regexp"
	"strings"
)

type Process struct {
	Name                       string
	Umask                      string
	State                      string
	Tgid                       string
	Ngid                       string
	Pid                        string
	PPid                       string
	TracerPid                  string
	Uid                        string
	Gid                        string
	FDSize                     string
	Groups                     string
	NStgid                     string
	NSpid                      string
	NSpgid                     string
	NSsid                      string
	Kthread                    string
	VmPeak                     string
	VmSize                     string
	VmLck                      string
	VmPin                      string
	VmHWM                      string
	VmRSS                      string
	RssAnon                    string
	RssFile                    string
	RssShmem                   string
	VmData                     string
	VmStk                      string
	VmExe                      string
	VmLib                      string
	VmPTE                      string
	VmSwap                     string
	HugetlbPages               string
	CoreDumping                string
	THP_enabled                string
	untag_mask                 string
	Threads                    string
	SigQ                       string
	SigPnd                     string
	ShdPnd                     string
	SigBlk                     string
	SigIgn                     string
	SigCgt                     string
	CapInh                     string
	CapPrm                     string
	CapEff                     string
	CapBnd                     string
	CapAmb                     string
	NoNewPrivs                 string
	Seccomp                    string
	Seccomp_filters            string
	Speculation_Store_Bypass   string
	SpeculationIndirectBranch  string
	Cpus_allowed               string
	Cpus_allowed_list          string
	Mems_allowed               string
	Mems_allowed_list          string
	voluntary_ctxt_switches    string
	nonvoluntary_ctxt_switches string
}

type Status struct{}

const (
	PROC_PATH = "/proc"
)

func getStatusFileToObj(str string) map[string]string {
	infoMap := make(map[string]string)
	infoLineList := strings.Split(str, "\n")
	for _, line := range infoLineList {
		line1 := strings.ReplaceAll(line, " ", "")
		line2 := strings.ReplaceAll(line1, "\t", "")
		info := strings.Split(line2, ":")
		if len(info) == 2 {
			infoMap[info[0]] = info[1]
			// fmt.Println(infoMap)
		}
	}
	return infoMap
}
func GetProcStatusList() []Process {
	validate := regexp.MustCompile("^[0-9]+$")
	processObjList := make([]Process, 0)
	processList, err := os.ReadDir(PROC_PATH)
	if err != nil {
		fmt.Println(err)
	}
	for _, info := range processList {
		var process Process
		if info.IsDir() && validate.MatchString(info.Name()) {
			filebuffer, e := os.ReadFile(PROC_PATH + "/" + info.Name() + "/status")
			if e != nil {
				fmt.Println(e)
			}
			file := string(filebuffer)
			j, _ := json.Marshal(getStatusFileToObj(file))
			json.Unmarshal(j, &process)
			processObjList = append(processObjList, process)
		}
	}
	return processObjList
}
