#!/bin/bash

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
go build -v github.com/pztrn/urtrator
if [ $? -ne 0 ]; then
    echo "Failed to build URTrator! Please, create a new bug report at https://github.com/pztrn/urtrator and attach FULL console output!"
    exit 1
fi

echo "Creating app bundle..."
mkdir -p /Applications/URTrator.app/Contents/{MacOS,Library,Resources}
cp $GOPATH/bin/urtrator /Applications/URTrator.app/Contents/MacOS/
cp $GOPATH/src/github.com/pztrn/urtrator/artwork/urtrator.icns /Applications/URTrator.app/Contents/Resources/

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
  <string>urtrator</string>
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

echo ${INFOPLIST} > /Applications/URTrator.app/Contents/Info.plist

#####################################################################

# Libraries works.
echo "Copying libraries..."
LIBS_TO_COPY=$(otool -L urtrator | awk {' print $1 '} | grep "/usr/local")

for lib in ${LIBS_TO_COPY[@]}; do
    cp ${lib} /Applications/URTrator.app/Contents/Library
    libname=$(echo ${lib} | awk -F"/" {' print $NF '})
    install_name_tool -change ${lib} @executable_path/../Library/${libname} /Applications/URTrator.app/Contents/MacOS/urtrator
done

echo "Finishing..."

echo "URTrator is ready! Launch from Applications!"
