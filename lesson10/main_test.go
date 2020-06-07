package main

import (
	"bytes"
	"io/ioutil"
	"testing"
)

func TestCopy(t *testing.T) {

	b := make([]byte, 100)
	//ioutil.WriteFile("~/1",b,644)
	err := MyCopy("/Users/vitaliy/1", "/Users/vitaliy/2", 0, 0)
	if err != nil {
		t.Errorf("Ошибка! %v", err)
	}
	x, err := ioutil.ReadFile("/Users/vitaliy/2")
	if err != nil {
		t.Errorf("Ошибка!! %v", err)
	}
	if bytes.Equal(x, b) {
		t.Errorf("Ошибка!!! - файлы не равны")
	}
}
