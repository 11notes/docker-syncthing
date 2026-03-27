package main

import (
	"os"
	"os/exec"
	"log"
	"encoding/xml"
	"golang.org/x/crypto/bcrypt"
	"syscall"
	"io"
	"time"
	"io/ioutil"
	"bytes"
	"strings"
)

// defaults
const config = "/syncthing/etc/config.xml"
const bin = "/usr/local/bin/syncthing"
const share = "/syncthing/share"

var (
	args = []string{"syncthing", "serve", "--config=/syncthing/etc", "--data=/syncthing/var", "--no-upgrade", "--no-browser", "--no-restart", "--gui-address=0.0.0.0:8384"}
)

// config structure
type gui struct {
	User     string `xml:"user"`
	Password string `xml:"password"`
	APIKey   string `xml:"apikey"`
	Metrics	 string `xml:"metricsWithoutAuth"`
}

type options struct {
	TelemetryAccepted string `xml:"urAccepted"`
	TelemetrySeen     string `xml:"urSeen"`
}

func main(){
	// check if config file exists already
	if _, err := os.Stat(config); err != nil {
		// syncthing config does not exist, create one
		firstRun()
	}else{
		// syncthing config does exist, run syncthing
		syncthing()
	}
}

func syncthing(){
	// execute syncthing and force to foreground
	if err := syscall.Exec(bin, args, os.Environ()); err != nil {
		os.Exit(1)
	}
}

func firstRun(){
	// run syncthing once to create default config
	if runOnce() {
		// check if config was created
		if _, err := os.Stat(config); err == nil {
			// create bcrypt encrypted password
			password, err := bcrypt.GenerateFromPassword([]byte(os.Getenv("SYNCTHING_PASSWORD")), 12)
			if err != nil {
				log.Fatal(err)
			}

			// update default
			updateConfig(string(password))
		}else{
			// syncthing could not create default config, abort
			log.Fatal("could not find config: ", config)
			os.Exit(1)
		}
	}
}

func runOnce() bool{
	// create process
	cmd := exec.Command(bin, args[1:]...)
	cmd.SysProcAttr = &syscall.SysProcAttr{Setpgid:true}
	err := cmd.Start()
	if err != nil {
		log.Fatal("could not start syncthing: ", err)
		return false
	}
	time.Sleep(5 * time.Second)
	if err := syscall.Kill(-cmd.Process.Pid, syscall.SIGTERM); err != nil {
		log.Fatal("could not terminate syncthing: ", err)
		return false
	}
	return true
}

func updateConfig(password string){
	// update the XML config with pre-set values or environment values
	file, err := os.Open(config)
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()
	var buf bytes.Buffer
	decoder := xml.NewDecoder(file)
	encoder := xml.NewEncoder(&buf)
	for {
		token, err := decoder.Token()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Printf("error getting token: %v\n", err)
			break
		}

		switch v := token.(type) {
			case xml.StartElement:
				switch v.Name.Local {
					case "gui":
						// set admin password, API key and allow metrics with no authentication
						var g gui
						if err = decoder.DecodeElement(&g, &v); err != nil {
							log.Fatal(err)
						}
						g.User = "admin"
						g.Password = password
						g.APIKey = os.Getenv("SYNCTHING_API_KEY")
						g.Metrics = "true"
						if err = encoder.EncodeElement(g, v); err != nil {
							log.Fatal(err)
						}
						continue

					case "options":
						// disable telemetry
						var o options
						if err = decoder.DecodeElement(&o, &v); err != nil {
							log.Fatal(err)
						}
						o.TelemetryAccepted = "-1"
						o.TelemetrySeen = "3"
						if err = encoder.EncodeElement(o, v); err != nil {
							log.Fatal(err)
						}
						continue
				}
		}
		
		if err := encoder.EncodeToken(xml.CopyToken(token)); err != nil {
			log.Fatal(err)
		}
	}
	if err := encoder.Flush(); err != nil {
		log.Fatal(err)
	}
	// replace default share path /Sync with correct one
	xmlConfig := strings.Replace(buf.String(), `<folder id="default" label="Default Folder" path="/Sync"`, `<folder id="default" label="Default Folder" path="` + share + `"`, -1)

	// write config file
	err = ioutil.WriteFile(config, []byte(xmlConfig), os.ModePerm)	
	if err != nil {
		os.Exit(1)
	}
	log.Println("default syncthing configuration with environment variables created, restarting ...")

	// start syncthing
	syncthing()
}