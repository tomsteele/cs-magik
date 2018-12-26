# cs-magik
Implements an events channel and job queue using Redis for Cobalt Strike. This is implemented by listening on `*` events on the teamserver and executing dynamic Sleep using a standard job queue implementation.

```
.\cs-magik-call.exe -redis-addr=192.168.24.137:6379 -server-id=x "return beacon_commands();"
DEBUG: encoding return beacon_commands();
DEBUG: Encoded job: NmQ4ODlmNGUtM2E3ZS00NDdjLTgyMmMtMTA4NTllM2I1M2Vl|cmV0dXJuIGJlYWNvbl9jb21tYW5kcygpOw==|
DEBUG: polling for result
Result:
{"result":["cancel","runas","psexec_psh","ps","bypassuac","upload","portscan","rportfwd","ls","psinject","ssh","mode dns6","run","download","checkin","dllload","reg","powershell","kerberos_ticket_use","net","elevate","execute-assembly","mkdir","steal_token","socks"
,"powershell-import","mv","winrm","shspawn","spawnto","execute","ppid","exit","make_token","dllinject","getuid","drives","logonpasswords","shell","psexec","covertvpn","rm","mode smb","pwd","shinject","note","link","setenv","powerpick","getsystem","screenshot","getp
rivs","jobkill","kerberos_ticket_purge","runasadmin","spawnu","sleep","wmi","desktop","downloads","kerberos_ccache_use","rev2self","dcsync","wdigest","mode dns","mimikatz","hashdump","cd","ssh-key","pth","jobs","mode http","clear","kill","cp","timestomp","help","sp
awn","keylogger","unlink","spawnas","browserpivot","socks stop","inject","mode dns-txt","runu"],"id":"6d889f4e-3a7e-447c-822c-10859e3b53ee"}
```

## Code Execution
Yep, you can send dynamic sleep code that is executed on the teamserver. Be smart, use authentication, use network filtering, etc on your Redis instance.

## But Why?
First, why not? Second, use your imagination, imagine being able to see new beacons coming in and reacting to them on the server-side using an established proxy using Python with impacket, could be cool? 

## Up and Running
Copy `jar_files` to `/opt` and change permissions accordingly. Then on the servers-side set the following environment variables:
* `REDIS_URL`
* `TEAMSERVER_ID`

Run either the `queue.cna` or `events.cna` using `agscript`. You can now subscribe to `events:$teamserver_id` for events. See `cmd/cs-magik-call/main.go` for an example of how to use the job queue.
