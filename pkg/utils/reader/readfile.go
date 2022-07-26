package reader

import "io/ioutil"

func ReadFile(fileName string) []byte {
	data, err := ioutil.ReadFile(fileName)
	if err != nil {
		panic(err)
	}
	return data
}
