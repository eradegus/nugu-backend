# NUGU Bus (server)
Backend proxy of nugu bus project


### How to run
```bash
$ go run .

or

$ PORT=8088 go run .
```

### File structure
```bash
nugu-bus
├── README.md
├── go.mod
├── go.sum
├── logic_dummy.go
├── main.go
├── requests_commmon.go
├── requests_handler.go
├── static
│   ├── files
│   │   ├── nugu_banner.png
│   │   └── nugu_logo_white.png
│   ├── js
│   │   ├── index.js
│   │   └── jquery-3.6.0.min.js
│   └── manifest.json
├── struct_db.go
├── struct_nugu.go
├── struct_openapi.go
└── templates
    ├── footer.html
    ├── header.html
    └── index.html
```
