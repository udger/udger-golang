package main

import (
	"fmt"
	"os"

	"github.com/yoavfeld/udger"
)

func main() {
	//if len(os.Args) != 3 {
	//	fmt.Println("Usage:\n\tgo run main.go ./udgerdb.dat \"Opera/9.50 (Nintendo DSi; Opera/507; U; en-US)\"")
	//	os.Exit(0)
	//}

	file := "../udgerdb_v3.dat"
	usag := "Mozilla/5.0 (Linux; Android 4.4.4; MI PAD Build/KTU84P) AppleWebKit/537.36 (KHTML, like Gecko) Version/4.0 Chrome/33.0.0.0 Safari/537.36"
	u, err := udger.New(file, &udger.Flags{Device: true})
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(-1)
	}

	fmt.Println(len(u.Browsers), "browsers loaded")
	fmt.Println(len(u.OS), "OS loaded")
	fmt.Println(len(u.Devices), "device types loaded")
	fmt.Println(len(u.Devices), "device types loaded")
	fmt.Printf("%#v", u.Devices)

	ua, err := u.Lookup(usag)
	if err != nil {
		fmt.Println("error: ", err)
		os.Exit(-1)
	}

	fmt.Printf("%+v\n", ua)
}
