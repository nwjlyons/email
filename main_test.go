package main

import (
	"encoding/binary"
	"io/ioutil"
	"log"
	"math"
	"os"
	"testing"
)

// tempFile is a helper function to mock stdin during testing.
func tempFile() *os.File {

	file, err := ioutil.TempFile(os.TempDir(), "stdin")
	if err != nil {
		log.Fatal(err)
	}

	return file
}

func TestSettingsFromFile(t *testing.T) {

	_, err := settingsFromFile()
	if err != nil {
		t.Error(err)
	}
}

func TestSettingsFromFlags(t *testing.T) {

	parseFlags([]string{"-t", "joe.bloggs@example.com,user@example.com", "-s", "subject from flag", "-b", "body from flag"}, nil)

	email, err := settingsFromFlags()

	if err != nil {
		t.Error(err)
	}

	if len(email.To) == 0 || email.To[0] != "joe.bloggs@example.com" || email.To[1] != "user@example.com" {
		t.Error("Email recipients is wrong")
	}

	if email.Subject != "subject from flag" {
		t.Error("Email subject is wrong")
	}

	if email.body != "body from flag" {
		t.Error("Email body is wrong")
	}
}

func TestReadingBodyFromStdin(t *testing.T) {

	// Mock stdin
	file := tempFile()
	defer os.Remove(file.Name())
	content := "content from stdin"
	file.WriteString(content)

	parseFlags([]string{"-t", "", "-s", "", "-b", ""}, file)

	email, err := settingsFromFlags()

	if err != nil {
		t.Error(err)
	}

	if email.body != content {
		t.Error("body from stdin is wrong")
	}
}

func TestBodyIsNotBinary(t *testing.T) {

	// Mock stdin
	file := tempFile()
	defer os.Remove(file.Name())
	var pi float64 = math.Pi
	err := binary.Write(file, binary.LittleEndian, pi)
	if err != nil {
		t.Error("binary.Write failed:", err)
	}

	parseFlags([]string{"-t", "", "-s", "", "-b", ""}, file)

	_, err = settingsFromFlags()

	if err.Error() != "Body is not text. Send as attachment instead." {
		t.Error("Failed to detect binary data from stdin.")
	}
}

func TestBodyIsRequiredWhenThereAreNoAttachments(t *testing.T) {

	parseFlags([]string{"-t", "", "-s", "", "-b", ""}, nil)
	_, err := settingsFromFlags()

	if err.Error() != "Body or attachment is required." {
		t.Error("Failed to detect empty body and no attachments.")
	}
}

func TestSettings(t *testing.T) {

	parseFlags([]string{"-t", "a@b.com", "-s", "subject", "-b", "body"}, nil)

	email, err := settings()
	if err != nil {
		t.Error(err)
	}

	if email.To[0] != "a@b.com" {
		t.Error("Email setting is wrong.", email)
	}

	if email.Subject != "subject" {
		t.Error("Subject setting is wrong.")
	}

	if email.body != "body" {
		t.Error("Body setting is wrong.")
	}
}

func TestSendingEmail(t *testing.T) {
	parseFlags([]string{"-t", "", "-s", "", "-b", "body"}, nil)

	email, err := settings()
	if err != nil {
		t.Error(err)
	}

	err = sendMail(email)
	if err != nil {
		t.Error(err)
	}
}
