# MacOS installation

For proper application bundle building you have to:

* Install Homebrew (http://brew.sh)
* Install Golang:

```
brew install go
```

* Install GTK+2:

```
brew install gtk+
```

* Install dylibbundler:

```
brew install dylibbundler
```

* Execute ``make-app.sh`` script from current directory. If everything
went fine - you will see URTrator.app right in this directory.
