# timo
A timvt(less) version implement in go

## Why
Make a minimal mvt tile server (function) which can tile a single table (layer) in postgis. 

## Less 
- Not support custom TMS ([TileMatrixSets](https://github.com/vincentsarago/TileMatrixSets))


## Reference
- [timvt](https://github.com/developmentseed/timvt)
- [pg_tileserv](https://github.com/CrunchyData/pg_tileserv)
- [tegola-postgis](https://github.com/go-spatial/tegola-postgis)

## More
- Add row filter, enable select special rows in a table to tile

## Use


Install timo with go get or clone repo to then `go build`
```
$ go get github.com/yuhangch/timo
```
Create a config file named ~/.timo.yaml like .timo.yaml.example
```
$ echo "URL: postgresql://postgres:passwd@127.0.0.1/timo" >> ~/.timo.yaml
```
Run Timo
```
$ timo serve 
$ timo serve --config=./.timo.yaml # select a config file
```

## Endpoint
```
http://host:port/tiles/{table_name}/{z}/{x}/{y}.pbf
```
### Parameters
- columns 
```
http://host:port/tiles/{table_name}/{z}/{x}/{y}.pbf?column=fid,name
```
- filters
```
http://host:port/tiles/{table_name}/{z}/{x}/{y}.pbf?filter=fid=3
```