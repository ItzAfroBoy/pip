<div align="center">
<pre>
 ____    ______   ____
/\  _`\ /\__  _\ /\  _`\  
\ \ \L\ \/_/\ \/ \ \ \L\ \
 \ \ ,__/  \ \ \  \ \ ,__/
  \ \ \/    \_\ \__\ \ \/
   \ \_\    /\_____\\ \_\
    \/_/    \/_____/ \/_/
<br>
A shitty public IP fetcher written in Go  
</pre>
</div>

## Installation

### Install with go

```shell
go install github.com/ItzAfroBoy/pip
pip ...
```

### Build from source

```shell
git clone https://github.com/ItzAfroBoy/pip
cd pip
go install
pip ...
```

## Usage

`Usage: pip [--dns] [--isp] [--location]`  

- `--dns`: Gets data of DNS server used
- `--isp`: Gets ISP data
- `--location`: Gets your location data

**Note: You are rate limited to 45 requests per min due to API service used**  
Powered by [IP-API](https://ip-api.com)
