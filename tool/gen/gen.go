package gen

import (
	"fmt"
	"math/rand"
	"strings"
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

const (
	loweralpha   = "qwertyuiopasdfghjklzxcvbnm"
	upperalpha   = "QWERTYUIOPASDFGHJKLZXCVBNM"
	numeric      = "0123456789"
	alphanumeric = loweralpha + upperalpha + numeric
)

func AnyChar(s string) byte {
	return s[rand.Intn(len(s))]
}

func GenString(dict string, ln int) string {
	b := strings.Builder{}
	for i := 0; i < ln; i++ {
		b.WriteByte(AnyChar(dict))
	}
	return b.String()
}

func Alnum(ln int) string {
	return GenString(alphanumeric, ln)
}

func Ab(ln int) string {
	return GenString("ab", ln)
}
