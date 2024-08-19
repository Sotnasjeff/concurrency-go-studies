package main

import (
	"io"
	"os"
	"strings"
	"testing"
)

func Test_updateMessage(t *testing.T) {
	wg.Add(1)

	go updateMessage("Hello Universe", &wg)

	wg.Wait()

	if !strings.Contains(msg, "Hello Universe") {
		t.Error("Expected to find Hello Universe but it's not there")
	}

}

func Test_printMessage(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	msg = "Real deal"
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Real deal") {
		t.Errorf("Expected to find Real deal but it's not there")
	}
}

func Test_main(t *testing.T) {
	stdOut := os.Stdout

	r, w, _ := os.Pipe()
	os.Stdout = w

	wg.Add(1)
	go updateMessage("Hello Jimmy", &wg)
	wg.Wait()
	printMessage()

	_ = w.Close()

	result, _ := io.ReadAll(r)
	output := string(result)

	os.Stdout = stdOut

	if !strings.Contains(output, "Hello Jimmy") {
		t.Errorf("Expected to find Hello Jimmy but it's not there")
	}

}
