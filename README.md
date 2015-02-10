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

    touch ~/.config/email/config.json

Copy the JSON below and update values according.

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

Override recipients in settings file.

    email -t "some.other.email.address@example.com" -b "body"

### Subject

Override subject in settings file

    email -s "subject" -b "body"

### All at once

    email -t "mum@example.com" -s "Some photos from the weekend" dinner.jpg tokyo-skyline.jpg
