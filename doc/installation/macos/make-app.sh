#!/bin/bash

#####################################################################
# HELPER FUNCTIONS
#####################################################################
# Libraries work.
change_framework_library_load_path() {
    local bin_to_fix=$1
    local path=$2

    # First iteration - main libraries.
    echo "Copying libraries..."
    dylibbundler -of -b -x ${bin_to_fix} -d ./URTrator.app/Contents/Framework/ -p ${path}

    # Fix shit for dylibbundler. By this moment we should have everything
    # we needed in Framework directory.
    for lib in $(ls ./URTrator.app/Contents/Framework); do
        libname=$(echo ${lib} | awk -F"/" {' print $NF '})
        DEPS=$(otool -L ./URTrator.app/Contents/Framework/${lib} | grep "/usr/local")
        for dep in ${DEPS[@]}; do
            dep_name=$(echo ${dep} | awk -F"/" {' print $NF '})
            install_name_tool -change ${dep} ${path}/${dep_name} ./URTrator.app/Contents/Framework/${libname}
        done
    done
}
# More permissive UMASK.
umask 002

# Do some checks.

# Do we have brew installed?
brew config | grep HOMEBREW_VERSION &>/dev/null
if [ $? -ne 0 ]; then
    echo "Please, install brew (http://brew.sh)"
    exit 1
fi

# Do we have Go installed?
# ToDo: more proper check.
if [ -z ${GOPATH} ]; then
    echo "Please, install Go >= 1.7 for continue and configure your shell. See Go installation docs for more information."
    exit 1
fi

# Do we have GTK+ installed?
GTK_STATUS=$(brew info gtk+ | grep Cellar)
if [ $? -ne 0 ]; then
    echo "Please, install GTK+ (brew install gtk+ gtksourceview)"
    exit 1
fi

# Okay, let's compile.
echo "Getting URTrator (and dependencies) sources"
go get -u -v -d github.com/pztrn/urtrator
if [ $? -ne 0 ]; then
    echo "Failed to get URTrator sources"
    exit 1
fi

echo "Building URTrator..."
go install -v github.com/pztrn/urtrator
if [ $? -ne 0 ]; then
    echo "Failed to build URTrator! Please, create a new bug report at https://github.com/pztrn/urtrator and attach FULL console output!"
    exit 1
fi

echo "Creating app bundle..."
mkdir -p URTrator.app/Contents/{MacOS,Framework,Resources}
# Copying URTrator binary
cp $GOPATH/bin/urtrator URTrator.app/Contents/MacOS/
# Copying main resources.
cp $GOPATH/src/github.com/pztrn/urtrator/artwork/urtrator.icns ./URTrator.app/Contents/Resources/
cp -R ./Resources/themes ./URTrator.app/Contents/Resources/

#####################################################################
# APP BUNDLE INFO.PLIST
#####################################################################
INFOPLIST='<?xml version="1.0" encoding="UTF-8"?>
<!DOCTYPE plist PUBLIC "-//Apple Computer//DTD PLIST 1.0//EN" "http://www.apple.com/DTDs/PropertyList-1.0.dtd">
<plist version="1.0">
<dict>
  <key>CFBundleGetInfoString</key>
  <string>URTrator</string>
  <key>CFBundleExecutable</key>
  <string>urtrator.sh</string>
  <key>CFBundleIdentifier</key>
  <string>name.pztrn.urtrator</string>
  <key>CFBundleName</key>
  <string>URTrator</string>
  <key>CFBundleIconFile</key>
  <string>urtrator.icns</string>
  <key>CFBundleShortVersionString</key>
  <string>0.1.0</string>
  <key>CFBundleInfoDictionaryVersion</key>
  <string>6.0</string>
  <key>CFBundlePackageType</key>
  <string>APPL</string>
  <key>IFMajorVersion</key>
  <integer>0</integer>
  <key>IFMinorVersion</key>
  <integer>1</integer>
</dict>
</plist>'

echo -e ${INFOPLIST} > URTrator.app/Contents/Info.plist

echo -e '#!/bin/bash\ncd "${0%/*}"\nexport GTK_PATH="../lib/gtk-2.0/"\nexport GTK_MODULES="../lib/gtk-2.0/"\nexport GDK_PIXBUF_MODULE_FILE="../lib/gdk-pixbuf-2.0/2.10.0/loaders.cache"\nexport GDK_PIXBUF_MODULEDIR="../lib/gdk-pixbuf-2.0/2.10.0/loaders/"\nexport GTK_EXE_PREFIX="../lib"\n./urtrator' > ./URTrator.app/Contents/MacOS/urtrator.sh
chmod +x ./URTrator.app/Contents/MacOS/urtrator.sh
#####################################################################
# Copying helper binaries.
cp /usr/local/Cellar/gdk-pixbuf/2.36.0_2/bin/gdk-pixbuf-query-loaders ./URTrator.app/Contents/MacOS/

# Copy GTK engines as needed for default theme.
cp /usr/local/lib/gtk-2.0/2.10.0/engines/* ./URTrator.app/Contents/Framework/
cp /usr/local/Cellar/gdk-pixbuf/2.36.0_2/lib/gdk-pixbuf-2.0/2.10.0/loaders/*.so ./URTrator.app/Contents/Framework/
chmod -R 0644 ./URTrator.app/Contents/Framework/*
chmod -R 0755 ./URTrator.app/Contents/MacOS/*

# Libraries works.
change_framework_library_load_path "./URTrator.app/Contents/MacOS/urtrator" "@executable_path/../Framework"
change_framework_library_load_path "./URTrator.app/Contents/MacOS/gdk-pixbuf-query-loaders" "@executable_path/../Framework"

#####################################################################
# Directory structure for GTK things. We will symlink neccessary
# libraries from /Framework here.
echo "Creating libraries structure with symlinks"
# GTK engines
mkdir -p ./URTrator.app/Contents/lib/gtk-2.0/2.10.0/engines/
cd ./URTrator.app/Contents/lib/gtk-2.0/2.10.0/engines/
ln -s ../../../../Framework/libclearlooks.so libclearlooks.so
ln -s ../../../../Framework/libmurrine.so libmurrine.so
ln -s ../../../../Framework/libpixmap.so libpixmap.so
# Pixbuf loaders
cd ../../../../
mkdir -p lib/gdk-pixbuf-2.0/2.10.0/loaders/
cd lib/gdk-pixbuf-2.0/2.10.0/loaders/
for file in $(ls ../../../../Framework | grep libpixbufloader); do
    ln -s ../../../../Framework/${file} ${file}
done
# Fix pixbuf loaders to load things from "../Framework".
for file in $(ls . | grep libpixbufloader); do
    DEPS=$(otool -L ${file} | grep "executable_path")
    for dep in ${DEPS[@]}; do
        dep_name=$(echo ${dep} | awk -F"/" {' print $NF '})
        install_name_tool -change ${dep} ../Framework/${dep_name} ${file}
    done
done
cd ..
ln -s ../../../Framework .
cd ../../../MacOS
GDK_PIXBUF_MODULE_FILE="../lib/gdk-pixbuf-2.0/loaders.cache" GDK_PIXBUF_MODULEDIR="../lib/gdk-pixbuf-2.0/2.10.0/loaders/" GTK_EXE_PREFIX="../lib" GTK_PATH="../Framework" ./gdk-pixbuf-query-loaders > ../lib/gdk-pixbuf-2.0/2.10.0/loaders.cache

echo "Finishing..."

echo "URTrator is ready! Copy URTrator.app bundle to Applications and launch it!"
