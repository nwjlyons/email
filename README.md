# Email

Email is a command line program that can send attachments, and stdin as the body of an email. It's most useful as a way of emailing yourself files.

## Install

```
    go get github.com/nwjlyons/email
    cd $GOPATH/src/github.com/nwjlyons/email
    go build
    go install
```

## Setup

    mkdir -p ~/.config/email;touch ~/.config/email/config.json

Copy the JSON below and update values accordingly.

    {
        "mailbox": "youremail@gmail.com",
        "from": "youremail@gmail.com",
        "host": "smtp.gmail.com",
        "port": "587",
        "password": "",
        "to": [
            "youremail@gmail.com"
        ],
        "subject": ""
    }

## Usage

### Body

Body from command line flag

    email -b "body"

or from a pipe

    echo "body" | email

or from stdin

    email < file.txt

### Attachments

    email attachment.pdf attachment.jpg

### Recipients

Override recipients from settings file. Accepts a comma separated list.

    email -t "one@example.com" -b "body"

multiple recipients

    email -t "one@example.com,two@example.com" -b "body"

### Subject

Override subject from settings file

    email -s "subject" -b "body"

### All at once

    email -t "mum@example.com" -s "Some photos from the weekend" dinner.jpg tokyo-skyline.jpg

## Tests

To run tests

    go test

These tests will actually send an email using the settings in your config file. To avoid this set `EMAIL_SEND` enviroment variable to `0`.

    EMAIL_SEND=0 go test
