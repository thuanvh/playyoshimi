package main
import (
	"fmt"
	"github.com/rakyll/portmidi"
	"log"
	"time"
	"os/exec"
	"os"
)
func main(){
	argsWithProg := os.Args
	deviceConnected := false
	for{
		portmidi.Initialize()
		count := portmidi.CountDevices()
		fmt.Println("Device count = ", count)
		// var deviceID = portmidi.DefaultInputDeviceID()
		//fmt.Println("Device id = ", deviceID)
		var deviceID portmidi.DeviceID
		var devices int = 0
		for deviceID = 0; int(deviceID) < count; deviceID++ {
			var deviceInfo = portmidi.Info(deviceID)
			if deviceInfo != nil && deviceInfo.IsInputAvailable {
				fmt.Println("Device id ", deviceID, " is ", deviceInfo)
				fmt.Println("Open device stream")
				in, err := portmidi.NewInputStream(deviceID, 1024)
				if err != nil {
					log.Fatal(err)
					fmt.Println("Error 1 ", err)
				}
				defer in.Close()
				fmt.Println("Read stream")
				events, err := in.Read(1024)
				if err != nil {
					log.Fatal(err)
					fmt.Println("Error 2 ", err)
					fmt.Println(events)
				}
				devices++
			}
		}
		portmidi.Terminate()

		if devices >= 2{
			if !deviceConnected {
				fmt.Println("Call Yoshimi")
				if len(argsWithProg) > 1 {
					fullcommand := os.Args[1:]
					fmt.Println(argsWithProg)
					fmt.Println(fullcommand)
					if len(fullcommand) > 1 {
						args := fullcommand[1:]
						cmd := exec.Command(fullcommand[0], args...)
						cmd.Start()
					} else {	
						cmd := exec.Command(fullcommand[0])
						cmd.Start()
					}
				}else{
					cmd := exec.Command("yoshimi");
					cmd.Start()
				}
				deviceConnected = true
			}
		}else{
			if deviceConnected {
				fmt.Println("Killall yoshimi")
				cmd := exec.Command("killall", "yoshimi", "-9")
				cmd.Start()
				deviceConnected = false
			}
		}
		time.Sleep(3 * time.Second)
	}
}