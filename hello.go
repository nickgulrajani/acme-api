package main
// Change one on develop in GL
import (
	"fmt"

	"github.com/hello-jenkins-larkintuckerllc/hello-jenkins/sum"
)

func main() {
	s := sum.Sum(1, 2)
	fmt.Println(s)
}
