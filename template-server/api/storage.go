package api

import (
	"bytes"
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
	"io"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

func EditFileName(name string) string {
	if strings.HasSuffix(name, ".md") {
		return name
	} else {
		return name + ".md"
	}
}

func CreateTemplateFile(name string, text string) error {
	loc, err := os.Getwd()
	if err != nil {
		return err
	}

	loc += viper.GetString("upload.location") + viper.GetString("upload.dir_name") + `\`
	dst, err := os.Create(loc + name)
	if err != nil {
		return err
	}
	defer dst.Close()
	logrus.Print("file created")

	src := bytes.NewBuffer([]byte(text))

	_, err = io.Copy(dst, src)
	if err != nil {
		return err
	}
	logrus.Print("buf created")

	/*err = os.Chmod(loc+name, syscall.FILE_SHARE_READ|syscall.FILE_SHARE_WRITE)*/
	if err != nil {
		return err
	}

	logrus.Print("context written")
	return nil
}

func DeleteFile(name string) error {
	loc, err := os.Getwd()
	if err != nil {
		return err
	}

	loc += viper.GetString("upload.location") + viper.GetString("upload.dir_name") + `\`

	err = os.Remove(loc + name)
	if err != nil {
		return err
	}
	print("file deleted")
	return nil
}

func UpdateTemplateFile(filename string, text string) error {
	loc, err := os.Getwd()
	if err != nil {
		return err
	}

	loc += viper.GetString("upload.location") + viper.GetString("upload.dir_name") + `\`

	err = os.Chmod(loc+filename, 0666)
	err = ioutil.WriteFile(loc+filename, []byte(text), 0666)
	if err != nil {
		return err
	}
	return nil
}

func ReadFile(name string) ([]byte, error) {
	loc, err := os.Getwd()
	if err != nil {
		return nil, err
	}

	loc += viper.GetString("upload.location") + viper.GetString("upload.dir_name") + `\`
	data, err := ioutil.ReadFile(loc + name)
	if err != nil {
		log.Panicf("failed reading data from file: %s", err)
	}
	fmt.Printf("\nFile Content: %s", data)
	return data, nil
}
