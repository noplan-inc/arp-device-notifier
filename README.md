# arp-device-notifier
It is a tool to get a list of devices by Arp protocol and request it to the server.

## Usage

### normal
arp-device-notifier -e http://example.com

### Verbose
arp-device-notifier -e http://example.com -v

### Post Interval(default is 10s)
arp-device-notifier -e http://example.com -i 3

### Authorization Header
arp-device-notifier -e http://example.com -a 'Bearer <token>'
