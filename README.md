# sampcmd-go

This tool allows you to start SA-MP via the command line, alternatively it can be used as a go library.

## Commandline Usage

```cmd
sampcmd-go.exe -h 127.0.0.1 -p 7777 -n Name
```

## Golang usage

```GO
returnValue := sampcmd.LaunchSAMP("-h 127.0.0.1 -p 7777 -n Name")
if(returnValue != 0) {
    fmt.Printf("An error occured. The application returned %d", returnValue)
}
```

## Available parameters

* `-d`
  > Enables the debugging mode.
* `-c <RCON password>`
  > Passes an RCON password to directly login as RCON administrator.
* `-h <IP address / hostname>`
  > Specifies the server that you want to join.
* `-p <port>`
  > Specified the servers port.
* `-n <name>`
  > Specifies the ingame username.
* `-z <server password>`
  > Specifies a password necessary to join the server.

## Credits

Big thanks to [BigETI](https://github.com/BigETI/) for helping me on this one!