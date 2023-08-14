# blowfish

NOTE: This is an old project I am not actively developing, but pushing out purely as a reference to a prototype that may be useful to others.

blowfish is a command line interface for interacting with [DMTF
Redfish](https://www.dmtf.org/standards/redfish) and [SNIA
Swordfish](https://www.snia.org/forums/smi/swordfish) enabled devices based on [gofish](https://github.com/stmcginnis/gofish).

Currently the CLI is based on human readable output via tablewriter, but output via JSON could be added in the future. The intent of the human readable output is for quick access to information relevant to debugging.

## build

This repo uses [taskfile](https://taskfile.dev/#/usage) instead of Makefile; [installation](https://taskfile.dev/#/installation) is easy and available for nearly all environments. Works similarly to Make by simply running `task lint` or `task build`.

## config

For convenience when testing / using you can add, get, remove, or set existing server information; see the details via the `blowfish config` command as well as setting a default if none is specified (you can setup a local mock if you desire)

NOTE: the information is simply written to a yaml file, not encrypted or managed in terms of username and password! (this can be enhanced in the future)

```
 blowfish config add --help
Adds new connection to use with the given name.

Usage:
  blowfish config add NAME [flags]

Flags:
      --default           Set this connection as the default.
  -h, --help              help for add
  -e, --host string       The host name or IP address of the system.
  -p, --password string   The password to connect with.
      --port uint16       Port used to connect (defaults to 443, or port 80 if 'http' protocol is specified.)
      --protocol string   Protocol to use (https (default) or http). (default "https")
      --secure            Enforce certificate validation with https connections (default allows self-signed certs).
  -u, --user string       The user name to connect as.

Global Flags:
      --config string   config file (default is $HOME/.blowfish.yaml)
```

example:

```
blowfish config add dfw2-ynot --host "10.0.0.1" --user "root" --password "mattisthebest"
     NAME       USER  ENDPOINT
     dfw2-ynot  root  https://10.0.0.1:443
```
you can then grab that information back / verify:

```
blowfish config get
     NAME       USER  ENDPOINT
     dfw2-ynot  root  https://10.0.0.1:443
```

This information can be written or fed into the CLI via the .blowfish.yaml file:

```yaml
default: ""
systems:
  dfw2-ynot:
    host: 10.0.0.1
    port: 443
    protocol: https
    username: root
    password: mattisthebest
    secure: false
```

## fetching info

you can provide a connection string or use the basic `get` function with the connection info you may have already added

```
blowfish chassis get
  NAME                     POWER  STATUS  MANUFACTURER  SERIAL          MODEL
  Computer System Chassis  On     OK      Dell Inc.     XXXXXXXXXXXXX  PowerEdge R6515
  PCIe SSD Backplane 1     On     OK                                    PCIe SSD Backplane 1
  BP14G+ 0:1               On     OK                                    BP14G+ 0:1

blowfish drive get
  NAME                    SIZE       STATUS  MANUFACTURER  MODEL          SERIAL NUMBER
  SSD 0                   223.57 GB  OK      MICRON        MTFDDAV240TCB  XXXXXXXXXXXXX
  SSD 1                   223.57 GB  OK      MICRON        MTFDDAV240TCB  XXXXXXXXXXXXX
  Solid State Disk 0:1:1  447.13 GB  OK      MICRON        MTFDDAK480TDC  XXXXXXXXXXXXX
  Solid State Disk 0:1:0  447.13 GB  OK      MICRON        MTFDDAK480TDC  XXXXXXXXXXXXX
```
