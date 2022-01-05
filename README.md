

<p align="center">
  <img alt="Github top language" src="https://img.shields.io/github/languages/top/ugurkorkmaz/shorten?color=56BEB8">

  <img alt="Github language count" src="https://img.shields.io/github/languages/count/ugurkorkmaz/shorten?color=56BEB8">

  <img alt="License" src="https://img.shields.io/github/license/ugurkorkmaz/shorten?color=56BEB8">



<br>

## :dart: About ##
Url Shortener is a web application that allows you to shorten your long URLs. Golang, Gofiber, Sqlite and Redis are the main technologies used in this project.

## :checkered_flag: Starting ##

```bash
# Clone this project
$ git clone https://github.com/ugurkorkmaz/shorten.git

# Access
$ cd shorten

# Install dependencies
$ docker-compose up -d

# Install dependencies
$ go mod vendor

# Run the project
$ go run .

# The server will initialize in the <http://localhost:3000>
```
```
#!/bin/bash

curl -X GET http://127.0.0.1/shorten?url=https://github.com/ugurkorkmaz/shorten

// 
{"long":"https://github.com/ugurkorkmaz/shorten","uid":"31ae8b"}
```
```
#!/bin/bash

curl -X GET http://127.0.0.1/31ae8b

// Redirects to long url


```
## :memo: License ##

This project is under license from GPL-3.0. For more details, see the [LICENSE](LICENSE.md) file.


Made with :heart: by <a href="https://github.com/ugurkorkmaz" target="_blank">Uğur Korkmaz</a>

&#xa0;

<a href="#top">Back to top</a>
