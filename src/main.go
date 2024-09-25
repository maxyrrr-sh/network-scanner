package main

import (
	"fmt"
	"strings"
	"os"
	"os/user"
  //"net"
  	"log"
	"sync"
	
)

func main() {

	currentUser, err := user.Current()
    if err != nil {
        log.Fatalf("Не вдалося отримати інформацію про користувача: %v", err)
    }

    if currentUser.Uid != "0" {
        fmt.Println("Ця програма потребує прав адміністратора.")
        os.Exit(1)
    }



	localIP, err := getLocalIP()
	if err != nil {
		fmt.Println(err)
		return
	}

	baseIP := strings.Join(strings.Split(localIP, ".")[:3], ".")
	fmt.Printf("Сканування мережі %s.0/24...\n", baseIP)

	availableHosts := scanNetwork(baseIP)

	if len(availableHosts) == 0 {
		fmt.Println("Пристрої не знайдені.")
	} else {
		fmt.Println("Доступні пристрої:")
		for _, host := range availableHosts {
			fmt.Println(host)
		}
	}

	var wg sync.WaitGroup
	// ress:=make(chan map[string]Port)
	// iptoports
	// for _, host := range availableHosts{
		
	// }
	portList := scanAllPorts(availableHosts[0], &wg)
	for _, port := range portList {
        fmt.Printf("Порт: %d, Сервіс: %s, Версія: %s\n", port.Port, port.Service, port.Version)
    }
	

}
