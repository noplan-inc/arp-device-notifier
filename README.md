# arp-device-notifier
It is a tool to get a list of devices by Arp protocol and request it to the server.

## dependent
- arp

## install

```bash
$ go get github.com/noplan-inc/arp-device-notifier
```

or

Download [Release Package](https://github.com/noplan-inc/arp-device-notifier/releases)

## Usage

### normal
```bash
$ arp-device-notifier post -e http://example.com
```

### Verbose
```bash
arp-device-notifier pst -e http://example.com -v
```
### Post Interval(default is 10s)
```bash
$ arp-device-notifier post -e http://example.com -i 3
```
### Authorization Header
```bash
$ arp-device-notifier post -e http://example.com -a 'Bearer <token>'
```


## Authors
- [serinuntius](https://twitter.com/_serinuntius)

## License
This project is licensed under the MIT License - see the LICENSE.md file for details
