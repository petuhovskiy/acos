package gen

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/bclicn/color"
	"github.com/petuhovskiy/acos/tool/def"
)

func Init() {
	rand.Seed(time.Now().UTC().UnixNano())
}

func LR32(l, r int32) int32 {
	return l + rand.Int31n(r-l+1)
}

func Int32() int32 {
	return int32(rand.Uint32())
}

func Eq(a, b interface{}) error {
	if a != b {
		return fmt.Errorf("not equal, %v != %v", a, b)
	}
	return nil
}

func Test(f func() error, cnt int) {
	ok := 0
	wa := 0
	for i := 0; i < cnt; i++ {
		fmt.Print(color.Bold(fmt.Sprintf("> Test %3d = ", i)))
		err := f()
		if err == nil {
			fmt.Println(def.OK)
			ok++
		} else {
			fmt.Println(def.WA, err)
			wa++
		}
	}
	fmt.Println(def.OK, ok)
	fmt.Println(def.WA, wa)
}
