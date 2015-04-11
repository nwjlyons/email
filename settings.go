package main

import (
	"encoding/json"
	"errors"
	"flag"
	"io/ioutil"
	"net/http"
	"os"
	"os/user"
	"path"
	"strings"
)

type Email struct {
	Mailbox  string
	From     string
	Host     string
	Port     string
	Password string

	To      []string
	Subject string
	// Lowercased so as not to parse body from JSON config
	body        string
	attachments []string
}

func settingsFromFile(filepath string) (email Email, err error) {
	// If user can't provide own config file, use default
	if len(filepath) == 0 {
		// This is the desired location.
		// "~/.config/email/config.json"
		// Need to use the user package to expand the tilde dynamically.
		usr, err := user.Current()
		if err != nil {
			return Email{}, err
		}
		filepath = path.Join(usr.HomeDir, ".config/email/config.json")
	}
	// Check file exists.
	if _, err := os.Stat(filepath); os.IsNotExist(err) {
		// File does not exist.
		return email, err
	}

	file, err := ioutil.ReadFile(filepath)
	if err != nil {
		// Problem reading file. eg. Permission error.
		return email, err
	}

	err = json.Unmarshal(file, &email)

	// Validate data
	if email.Mailbox == "" {
		return email, errors.New("Mailbox from settings file is blank.")
	}
	if email.From == "" {
		return email, errors.New("From from settings file is blank.")
	}
	if email.Host == "" {
		return email, errors.New("Host from settings file is blank.")
	}
	if email.Port == "" {
		return email, errors.New("Port from settings file is blank.")
	}
	if email.Password == "" {
		return email, errors.New("Password from settings file is blank.")
	}

	return email, err
}

func settingsFromFlags() (email Email, err error) {

	// Email recipients
	if *to != "" {

		email.To = strings.Split(*to, ",")

		// Trim space from substrings.
		for key, value := range email.To {
			email.To[key] = strings.Trim(value, " ")
		}
	}

	// Email subject
	email.Subject = *subject

	// Email body

	// stat stdin to get statistics about the file, like filesize which is what
	// I'm most interested in
	fi, err := stdin.Stat()
	if err != nil {
		return email, err
	}

	// read from stdin if the file size is greater than zero or the file type is
	// pipe
	if fi.Size() > 0 || fi.Mode()&os.ModeNamedPipe != 0 {
		stdinBody, err := ioutil.ReadFile(stdin.Name())
		if err != nil {
			return email, err
		}

		if strings.HasPrefix(http.DetectContentType(stdinBody), "text") == false {
			return email, errors.New("Body is not text. Send as attachment instead.")
		}

		email.body = string(stdinBody)
	} else if *body != "" {
		email.body = *body
	}

	email.attachments = flag.Args()

	if len(email.attachments) <= 0 && email.body == "" {
		return email, errors.New("Email body or attachment is required.")
	}

	return email, err
}

func settings() (Email, error) {
	fromFile, err := settingsFromFile(*config)
	if err != nil {
		return Email{}, err
	}

	fromFlags, err := settingsFromFlags()
	if err != nil {
		return Email{}, err
	}

	if fromFlags.Subject != "" {
		fromFile.Subject = fromFlags.Subject
	}

	if len(fromFlags.To) > 0 {
		fromFile.To = fromFlags.To
	}

	// Body and attachments aren't parsed from settings file. Therefore always
	// merge.
	fromFile.body = fromFlags.body
	fromFile.attachments = fromFlags.attachments

	return fromFile, nil
}
