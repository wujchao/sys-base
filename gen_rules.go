package sys_base

import (
	"fmt"
	"github.com/gogf/gf/v2/os/gfile"
)

func Gen() {
	fmt.Println(gfile.Pwd())
	fmt.Println(gfile.SelfPath())
}
