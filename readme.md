# Base on [GGH](https://github.com/byawitz/ggh)

### What is GGH?

GGH is a lightweight, fast wrapper around your SSH commands. It helps you to recall your SSH sessions.
This is one of the most useful tools for developers who work with multiple servers.

Thanks to GGH Team.

### What is Worm?

Worm inherits all the features of GGH and expands with several new features (I need), such as:

- Setting up a workspace for each project or company...
- Each workspace contains multiple files, each managing a separate list of servers needing SSH access. Each file can represent a data center or Platform…
- Supporting server access via SSH and TSH
- Saving history for each workspace

### Installation && Configuration

- Clone the repository
- Run the command `go install github.com/1Ginz/jump/cmd@latest`
- Setup workspace

- Setup Configuration file
by default jump look up config at ~/.jump/configs
```shell
mkdir ~/.jump/configs
```
create workspace dir (each workspace are isolate)
```shell
mkdir adtech
```
now each workspace we can create login file. each file manage connect to 1 remote server
```shell
Host ssp-23-115
	HostName 192.168.23.115
	User cuongnv
	Mode TSH
	#other info want to show
	IdentityFile NPS
```

```text
Host {HostNAme}
	HostName {IP}
	User {UserName}
	Mode {SSH|TSH}
	Detail {What ever u want}
````

### Usage

List workspaces

```shell
worm --workspace
```

Switch workspace

```shell
worm --active 
```

Interactive history

```shell
worm 
```

```shell
worm --history
```

Interactive configuration file

```shell
worm -
```

Interactive configuration file with search for groups or hostnames.

```shell
worm - xxxx
```
