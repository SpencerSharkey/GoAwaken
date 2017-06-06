package main

import (
	"flag"
	"os"

	"fmt"

	gs "github.com/ReviveNetwork/GoRevive/GameSpy"
	log "github.com/ReviveNetwork/GoRevive/Log"
)

var (
	// BuildTime of the build provided by the build command
	BuildTime = "Not provided"
	// GitHash of build provided by the build command
	GitHash = "Not provided"
	// GitBranch of the build provided by the build command
	GitBranch = "Not provided"
	// compileVersion we are receiving by the build command
	CompileVersion = "0"
	// Version of the Application
	Version = "0.0.0"
)

func EncryptDecrypt(input string, key []rune) (output string) {
	a := []rune(input)
	for i := 0; i < len(input); i++ {
		new := string(a[i] ^ key[i%len(key)])
		fmt.Println([]byte(new))
		output += new
	}

	return output
}

func main() {
	var (
		logLevel = flag.String("loglevel", "error", "LogLevel [error|warning|note|debug]")
	)
	flag.Parse()

	if CompileVersion != "0" {
		Version = Version + "." + CompileVersion
	}

	log.SetLevel(*logLevel)
	log.Notef("Starting up v%s", Version)

	// Startup done

	// Generate session key
	var len = len("MakaHost")
	var nameIndex = 0
	var session rune
	runeName := []rune("MakaHost")

	for {
		len = len - 1
		if len < 0 {
			break
		}
		fmt.Println("Char: ", runeName[nameIndex])
		fmt.Println("Index: ", ((runeName[nameIndex]^session)&0xff)%256)
		fmt.Println("Operator: ", (session >> 8))
		fmt.Println("Crc: ", gs.CrcLookup[((runeName[nameIndex]^session)&0xff)%256])
		tmpSession := session >> 8
		session = gs.CrcLookup[((runeName[nameIndex]^session)&0xff)%256] ^ (tmpSession)
		fmt.Println("Result:", session)

		nameIndex = nameIndex + 1
	}

	fmt.Println(session)

	test := gs.ShortHash("Bla")
	log.Noteln(test)

	test2, err := gs.ProcessCommand("\\pi\\\\profileid\\1234\\nick\\MakaHost\\userid\\4321\\\\final\\")
	if err != nil {
		log.Errorln(err)
	}
	log.Noteln(test2, err)

	test3 := new(gs.Socket)
	eventsChannel, err := test3.New("Testing", "10000")
	if err != nil {
		log.Errorln(err)
	}

	for {
		select {
		case event := <-eventsChannel:
			switch {
			case event.Name == "newClient":
				log.Debugln(event)
			case event.Name == "error":
				log.Debugln(event)
			case event.Name == "close":
				log.Debugln(event)
				os.Exit(0)
			default:
				log.Debugln(event)
			}
		}
	}

}
