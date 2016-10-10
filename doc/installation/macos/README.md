# MacOS installation

For proper application bundle building you have to:

* Install Homebrew (http://brew.sh)
* Install Golang:

```
brew install go
```

* Install GTK+2:

```
brew install gtk+ --with-quartz-relocation
```

*Note: default GTK+2 build from Brew might not work for you!*

* Reinstall gdk-pixbuf with additional option:

```
brew install --with-relocations
```

*Note: default build might not work for you!*

* Install dylibbundler:

```
brew install dylibbundler
```

* Execute ``make-app.sh`` script from current directory. If everything
went fine - you will see URTrator.app right in this directory.

# Some descriptions

## Resources

This directory will be copied inside bundle. It contains resources needed
for proper working or/and launching.

Of course, it *might* be able to work without this directory, but it will
be superugly.
