# zeroconf-beacon

Spawn a simple zeroconf service over the local network

## Install

```
go get -u github.com/sosedoff/zeroconf-beacon
```

Using make:

```
make install
```

Build binaries:

```
make linux
make osx
make all
```

## Usage

```
Usage of zeroconf-beacon:
  -domain string
      Service domain (default "local.")
  -name string
      Service name
  -port int
      Service port (default 80)
  -protocol string
      Service protocol (default "_http._tcp")
  -txt string
      Service text records: foo=bar,hello=world
```

Start a service on your local machine:

```
zeroconf-beacon -name=demo -txt=token=foo
```

### Discover on OSX

```
$ dns-sd -B _http._tcp

Browsing for _http._tcp
DATE: ---Fri 23 Mar 2018---
20:06:48.338  ...STARTING...
Timestamp     A/R    Flags  if Domain               Service Type         Instance Name
20:06:48.339  Add        2   4 local.               _http._tcp.          demo

$ dns-sd -L demo _http._tcp

Lookup demo._http._tcp.local
DATE: ---Fri 23 Mar 2018---
20:07:42.854  ...STARTING...
20:07:42.855  demo._http._tcp.local. can be reached at Dan-S-Macbook.local.local.:80 (interface 4)
 token=foo
```

### Discover on Linux

```
$ avahi-browse -r _http._tcp

+  wlan0 IPv4 demo                                          Web Site             local
=  wlan0 IPv4 demo                                          Web Site             local
   hostname = [Dan-S-Macbook.local.local]
   address = [192.168.0.107]
   port = [80]
   txt = ["token=foo"]
```

## License

MIT