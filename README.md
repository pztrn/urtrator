# URTrator

[![Join the chat at https://gitter.im/urtrator/Lobby](https://badges.gitter.im/urtrator/Lobby.svg)](https://gitter.im/urtrator/Lobby?utm_source=badge&utm_medium=badge&utm_campaign=pr-badge&utm_content=badge)

![Main Window](/doc/screenshots/0.1-main_window.png)

URTrator is a desktop application that should (eventually) replace
Urban Terror's and IRC client interfaces for you, because they're
pretty shitty :).

Right now it can:

* Obtaining list of Urban Terror servers from master server and
updating information about them.
* Local caching of whole data (in SQLite3 database).
* Extended Urban Terror launching capabilities (e.g. launching game
in another X session).
* Supporting of multiple profiles with multiple game versions.
When you're launching Urban Terror URTrator will check profile you're
trying to use and, if version incompatability found, will not launch
the game.
* Favorites servers.

Planning:

* Updating single server.
* Showing information about servers (like in UrT Connector).
* Friends searching.
* RCON console.
* Game updating (not from official servers yet, sorry).
* Pickup/matchmaking interfaces.
* All kinds of notifications.
* Extended profile editor, so every profile could have own configuration
files, etc.
* Clipboard monitoring.
* ...maybe more :)

# Installation

## Release

You don't need to install anything, thanks to Go's static linking.
URTrator executable contains everything we need. Just download
approriate binary and launch it! :) The only thing is to make
sure you have GTK2 installed.

## Development version

URTrator written in Go and GTK2, so you should have them installed.
Make sure your ``GOPATH`` variable is defined.

Then execute:

```
go get -d github.com/pztrn/urtrator
go install github.com/pztrn/urtrator
```

First command will get sources of URTrator and dependencies, second
command will build executable for you and place it in ``$GOROOT/bin``.


### Updating

``go get`` will do initial repo clone for you, but flag ``-u`` is
required to get updated URTrator source. So, for updating sources
just issue:

```
go get -d -u github.com/pztrn/urtrator
```

# Important information

## GTK warnings in console

Many GTK warnings in console may appear while using URTrator. Unfortunately,
they are out of control, because they are related to Go GTK+2 bindings
and nothing can be done on URTrator side.

## About box cause app crash on macOS

As previous, it is related to GTK+2, nothing I can do there.
