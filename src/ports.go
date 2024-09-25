package main

import (
    "bufio"
    "fmt"
    "net"
    "sync"
    "time"
)

type Port struct {
    Port    int
    Service string
    Version string
}

const startPort = 1
const endPort = 65525

var services = map[int]string{
    20:   "FTP Data Transfer",
    21:   "FTP Control",
    22:   "SSH",
    23:   "Telnet",
    25:   "SMTP",
    53:   "DNS",
    67:   "DHCP",
    69:   "TFTP",
    80:   "HTTP",
    110:  "POP3",
    143:  "IMAP",
    443:  "HTTPS",
    3306: "MySQL",
    5432: "PostgreSQL",
    6379: "Redis",
    8080: "HTTP Alt",
}

func scanPort(ip string, port int, results chan<- Port, wg *sync.WaitGroup) {
    defer wg.Done()
    address := fmt.Sprintf("%s:%d", ip, port)

    conn, err := net.DialTimeout("tcp", address, 1*time.Second)
    if err == nil {
        portInfo := Port{
            Port:    port,
            Service: services[port],
        }

        var version string
        switch port {
        case 21: // FTP
            _, _ = conn.Write([]byte("USER anonymous\r\n")) 
            reader := bufio.NewReader(conn)
            version, _ = reader.ReadString('\n')
        case 22: // SSH
            version = "SSH (протокол 2)"
        case 23: // Telnet
            _, _ = conn.Write([]byte("HELP\r\n")) 
            reader := bufio.NewReader(conn)
            version, _ = reader.ReadString('\n')
        case 25: // SMTP
            _, _ = conn.Write([]byte("EHLO example.com\r\n")) 
            reader := bufio.NewReader(conn)
            version, _ = reader.ReadString('\n')
        case 53: // DNS
            version = "DNS (не вимагає запиту для версії)"
        case 67: // DHCP
            version = "DHCP (не вимагає запиту для версії)"
        case 69: // TFTP
            version = "TFTP (не вимагає запиту для версії)"
        case 80: // HTTP
            _, _ = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n")) 
            reader := bufio.NewReader(conn)
            version, _ = reader.ReadString('\n')
        case 110: // POP3
            _, _ = conn.Write([]byte("USER anonymous\r\n")) 
            reader := bufio.NewReader(conn)
            version, _ = reader.ReadString('\n')
        case 143: // IMAP
            _, _ = conn.Write([]byte("A001 LOGIN user pass\r\n")) 
            reader := bufio.NewReader(conn)
            version, _ = reader.ReadString('\n')
        case 443: // HTTPS
            version = "HTTPS (не вимагає запиту для версії, використовуйте TLS)"
        case 3306: // MySQL
            version = "MySQL (можливо, потрібно використовувати специфічні команди)"
        case 5432: // PostgreSQL
            version = "PostgreSQL (можливо, потрібно використовувати специфічні команди)"
        case 6379: // Redis
            _, _ = conn.Write([]byte("PING\r\n"))   
            reader := bufio.NewReader(conn)
            version, _ = reader.ReadString('\n')
        case 8080: // HTTP Alternative
            _, _ = conn.Write([]byte("HEAD / HTTP/1.0\r\n\r\n"))  
            reader := bufio.NewReader(conn)
            version, _ = reader.ReadString('\n')
        }

        portInfo.Version = version
        results <- portInfo 
        conn.Close()
    }
}

func scanAllPorts(ip string, wg *sync.WaitGroup) []Port {
    results := make(chan Port) 
    var portList []Port

    go func() {
        for result := range results {
            portList = append(portList, result)
        }
    }()

    for port := startPort; port <= endPort; port++ {
        wg.Add(1)
        go scanPort(ip, port, results, wg) 
    }

    wg.Wait() 
    close(results) 

    return portList 
}
