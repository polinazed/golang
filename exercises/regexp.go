package main 
import ( 
	"fmt"
	"regexp"
)

func main(){
	logs := `192.168.1.1 - GET
	10.0.0.2 - POST
	192.168.1.1 - GET
	10.0.0.3 - PUT`

	ips := re.FindAllString(logs, -1)
	fmt.Println(ips) // что будет?

	count := make(map[string]int)
	for _, ip := range ips {
	    count[ip]++
	}
	fmt.Println(count)
}