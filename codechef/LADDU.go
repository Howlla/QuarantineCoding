package main

import (
	"bufio"
	"fmt"
	"strings"
	"strconv"
	"os"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	var T int
	fmt.Scan( &T)
	for ; T>0; T-- {
		var activities, sum int
		var origin string
		scanner.Scan() // use `for scanner.Scan()` to keep reading
    	line := scanner.Text()
		// fmt.Println("captured:",line)
		activities,_ = strconv.Atoi( strings.Split(line," ")[0])
		origin = strings.Split(line," ")[1]
		for ;activites>0;activites--{
			scanner.Scan() // use `for scanner.Scan()` to keep reading
    		line := scanner.Text()
		}
	}
}
