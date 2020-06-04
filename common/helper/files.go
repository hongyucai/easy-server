package helper

import(
	"io/ioutil"
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func IoutilRead(path string)string{
	fi,err := os.Open(path)
	if err != nil{panic(err)}
	defer fi.Close()
	fd,err := ioutil.ReadAll(fi)
	// fmt.Println(string(fd))
	return string(fd)
}
