# GoRedy
# Configuration file

Create a file named `redis_conf.json` as:

```json
{
	"IpAddr": "192.168.99.100",
	"Port": 32769
}
```

# JSON output
```bash
# string
$ curl http://127.0.0.1/cmd/get/$myKey
{"key":"myKey","value":"myValue"}

```
