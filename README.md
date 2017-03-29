# basicauth

[![Build Status](https://travis-ci.org/northbright/basicauth.svg?branch=master)](https://travis-ci.org/northbright/basicauth)
[![Go Report Card](https://goreportcard.com/badge/github.com/northbright/basicauth)](https://goreportcard.com/report/github.com/northbright/basicauth)
[![GoDoc](https://godoc.org/github.com/northbright/basicauth?status.svg)](https://godoc.org/github.com/northbright/basicauth)

basicauth is [Golang](http://golang.org) package which provides helper functions of HTTP basic auth.

#### Go Version Requirement
* Go1.4 and later

#### Example

See [./example](./example):  

    func hello(w http.ResponseWriter, r *http.Request) {
        // Creates a BasicAuth with username and password
        ba := basicauth.New("admin", "admin")
        if ba.IsOK(w, r) {  // Basic Authorized
            io.WriteString(w, "/: you are in.")
        } else {  // UnAuthrorized
            io.WriteString(w, "/: 401: unauthorized.")
        }
    }
    
#### Documentation
* [API Reference](http://godoc.org/github.com/northbright/basicauth)

#### License
* [MIT License](./LICENSE)
