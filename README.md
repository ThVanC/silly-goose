# Silly Goose

Silly Goose is a very simple, vulnerable web application with 2 forms, used to familiarize users with pentest tools such as OWASP Zed Attack Proxy.
It can be run from a single binary file, without requiring any other dependencies installed.

## Install and run Silly Goose

see `/windows`, `/linux`, `/mac` for the binary file for each platform (i386)

Add executable permissions if needed to the file, and run `silly-goose`.


Visit [http://localhost:3016](), after configuring your browser to connect to your local OWASP ZAP Proxy.

## Modules

In module 1 you can log-in with a username / password.


The username/password combination is admin/admin.

In Module 2 you can specify an xml file. If the user is logged in, the username is shown.
This can be used to test a generated CSRF form.


![honk](https://i.imgur.com/1YGFqUu.jpg)