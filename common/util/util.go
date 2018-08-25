package util

import (
	"crypto/rand"
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"mime/multipart"
	"time"
)

func NewOrderID() (string, error) {
	var orderid string
	uuid, err := NewUUID()
	if err != nil {
		return "", err
	}
	orderid = fmt.Sprintf("%v%s", time.Now().UnixNano(), uuid)
	return orderid, nil
}

func NewUUID() (string, error) {
	uuid := make([]byte, 16)
	n, err := io.ReadFull(rand.Reader, uuid)
	if n != len(uuid) || err != nil {
		return "", err
	}
	// variant bits; see section 4.1.1
	uuid[8] = uuid[8]&^0xc0 | 0x80
	// version 4 (pseudo-random); see section 4.1.3
	uuid[6] = uuid[6]&^0xf0 | 0x40
	return fmt.Sprintf("%x%x%x%x%x", uuid[0:4], uuid[4:6], uuid[6:8], uuid[8:10], uuid[10:]), nil
}

func SaveFile(file multipart.File, filepath string) (fileURL string, errR error) {
	data, err := ioutil.ReadAll(file)
	if err != nil {
		err = errors.New("read file error：" + err.Error())
		return "", err
	}

	filename, err := NewUUID()
	if err != nil {
		err = errors.New("create file name error：" + err.Error())
		return "", err
	}
	filepath = filepath + filename
	err = ioutil.WriteFile(filepath, data, 0666)
	return filepath, err
}
