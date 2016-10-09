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
go install -v github.com/pztrn/urtrator
if [ $? -ne 0 ]; then
    echo "Failed to build URTrator! Please, create a new bug report at https://github.com/pztrn/urtrator and attach FULL console output!"
    exit 1
fi

echo "Creating app bundle..."
mkdir -p URTrator.app/Contents/{MacOS,Framework,Resources}
cp $GOPATH/bin/urtrator URTrator.app/Contents/MacOS/
cp $GOPATH/src/github.com/pztrn/urtrator/artwork/urtrator.icns URTrator.app/Contents/Resources/

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

echo ${INFOPLIST} > URTrator.app/Contents/Info.plist

echo -e '#!/bin/bash\ncd "${0%/*}"\n./urtrator' > ./URTrator.app/Contents/MacOS/urtrator.sh
chmod +x ./URTrator.app/Contents/MacOS/urtrator.sh
#####################################################################

# Libraries works.
# First iteration - main libraries.
echo "Copying libraries..."
dylibbundler -of -b -x ./URTrator.app/Contents/MacOS/urtrator -d ./URTrator.app/Contents/Framework/ -p @executable_path/../Framework/

# Fix shit for dylibbundler. By this moment we should have everything
# we needed in Framework directory.
for lib in $(ls ./URTrator.app/Contents/Framework); do
    libname=$(echo ${lib} | awk -F"/" {' print $NF '})
    DEPS=$(otool -L ./URTrator.app/Contents/Framework/${lib} | grep "/usr/local")
    for dep in ${DEPS[@]}; do
        dep_name=$(echo ${dep} | awk -F"/" {' print $NF '})
        install_name_tool -change ${dep} @executable_path/../Framework/${dep_name} ./URTrator.app/Contents/Framework/${libname}
    done
done

echo "Finishing..."

echo "URTrator is ready! Copy URTrator.app bundle to Applications and launch it!"
