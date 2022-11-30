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
├── logic_bus.go
├── logic_dday.go
├── logic_stock.go
├── logic_weather.go
├── main.go
├── requests_commmon.go
├── requests_handler.go
├── static
│   ├── files
│   │   ├── NationalRegionCodeutf8.csv
│   │   ├── nugu_banner.jpg
│   │   └── nugu_logo_white.png
│   ├── js
│   │   ├── index.js
│   │   └── jquery-3.6.0.min.js
│   └── manifest.json
├── struct_db.go
├── struct_nugu.go
├── struct_openapi_bus.go
├── struct_openapi_stock.go
├── struct_openapi_weather.go
└── templates
    ├── footer.html
    ├── header.html
    └── index.html
```
