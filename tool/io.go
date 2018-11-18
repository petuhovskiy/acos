package tool

import "io/ioutil"

func ReadText(src string) (string, error) {
	b, err := ioutil.ReadFile(src)
	if err != nil {
		return "", err
	}
	return string(b), nil
}
