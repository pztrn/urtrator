#!/bin/bash

# URTrator Windows build scripts.
# Requirements: installed and properly configured Go and MSYS2

#####################################################################
# VARIABLES
#####################################################################
# Path of script.
SCRIPT_PATH=$(dirname "`readlink -f "${BASH_SOURCE}"`")
# Go binary path.
GO=$(which go | awk {' print $1 '})

#####################################################################
# Checks if Go installed and properly configured.
check_golang() {
    if [ ${#GO} -eq 0 ]; then
        echo "! Can't find Go binary. Please, install Go and configure your environment."
    fi
}

# Check if we are in valid MSYS2 environment.
check_msys() {
    echo "* Checking for MSYS..."
    if [ -z MSYSTEM ]; then
        echo "! Probably, not MSYS2 environment. Building can continue only with MSYS2!"
        exit 1
    fi

    if [ "${MSYSTEM}" == "MSYS" ]; then
        echo "! Invalid MSYS2 environment. You're launched MSYS console, but building can continue in MINGW64 console only."
        exit 1
    fi
}

# Check if we have neccessary MSYS packages installed.
check_msys_packages() {
    echo "* Installing neccessary MSYS2 packages..."
    pacman -S ${MINGW_PACKAGE_PREFIX}-tools-git ${MINGW_PACKAGE_PREFIX}-gtk2 ${MINGW_PACKAGE_PREFIX}-pkg-config ${MINGW_PACKAGE_PREFIX}-gtk-engines ${MINGW_PACKAGE_PREFIX}-gtk-engine-murrine --noconfirm --needed
}

# Build URTrator
urtrator_build() {
    echo "* Starting URTrator building..."
    # Create temporary Go root.
    if [ -d "${SCRIPT_PATH}/GOROOT" ]; then
        rm -rf "${SCRIPT_PATH}/GOROOT"
    fi
    mkdir -p "${SCRIPT_PATH}/GOROOT"
    # Exporting Go variables.
    export GOPATH="${SCRIPT_PATH}/GOROOT"

    echo "* Obtaining URTrator and dependencies sources..."
    go get -u -d -v github.com/pztrn/urtrator
    echo "* Building URTrator"
    go install -v github.com/pztrn/urtrator
}

# Prepare urtrator distribution.
urtrator_dist() {
    echo "* Preparing URTrator distribution..."

    echo "* Creating distribution directory..."
    if [ -d "${SCRIPT_PATH}/dist" ]; then
        rm -rf "${SCRIPT_PATH}/dist"
    fi
    mkdir "${SCRIPT_PATH}/dist"

    echo "* Copying URTrator binary..."
    cp "${GOPATH}/bin/urtrator.exe" "${SCRIPT_PATH}/dist/"

    echo "* Getting list of library dependencies..."
    DEPS=$(ldd dist/urtrator.exe | grep -v "Windows" | awk {' print $1 '})
    echo "* Copying dependencies..."
    for dep in ${DEPS[@]}; do
        cp /mingw64/bin/${dep} "${SCRIPT_PATH}/dist"
    done

    echo "* Installing GTK+2 theme engines..."
    mkdir -p "${SCRIPT_PATH}/dist/lib/gtk-2.0/2.10.0/engines/"
    cp /mingw64/lib/gtk-2.0/2.10.0/engines/libmurrine.dll "${SCRIPT_PATH}/dist/lib/gtk-2.0/2.10.0/engines/"
    cp /mingw64/lib/gtk-2.0/2.10.0/engines/libredmond95.dll "${SCRIPT_PATH}/dist/lib/gtk-2.0/2.10.0/engines/"
    cp /mingw64/lib/gtk-2.0/2.10.0/engines/libthinice.dll "${SCRIPT_PATH}/dist/lib/gtk-2.0/2.10.0/engines/"
    cp /mingw64/lib/gtk-2.0/2.10.0/engines/libmist.dll "${SCRIPT_PATH}/dist/lib/gtk-2.0/2.10.0/engines/"
    cp /mingw64/lib/gtk-2.0/2.10.0/engines/libindustrial.dll "${SCRIPT_PATH}/dist/lib/gtk-2.0/2.10.0/engines/"
    cp /mingw64/lib/gtk-2.0/2.10.0/engines/libhcengine.dll "${SCRIPT_PATH}/dist/lib/gtk-2.0/2.10.0/engines/"
    cp /mingw64/lib/gtk-2.0/2.10.0/engines/libglide.dll "${SCRIPT_PATH}/dist/lib/gtk-2.0/2.10.0/engines/"
    cp /mingw64/lib/gtk-2.0/2.10.0/engines/libcrux-engine.dll "${SCRIPT_PATH}/dist/lib/gtk-2.0/2.10.0/engines/"
    cp /mingw64/lib/gtk-2.0/2.10.0/engines/libclearlooks.dll "${SCRIPT_PATH}/dist/lib/gtk-2.0/2.10.0/engines/"
    cp /mingw64/lib/gtk-2.0/2.10.0/engines/libwimp.dll "${SCRIPT_PATH}/dist/lib/gtk-2.0/2.10.0/engines/"

    echo "* Checking that we haven't forget any of dependencies..."
    LIBS=$(ls "${SCRIPT_PATH}/dist" | grep ".dll")
    for lib in ${LIBS[@]}; do
        LIBDEPS=$(ldd "${SCRIPT_PATH}/dist/${lib}" | grep -v "Windows" | awk {' print $1 '})
        for libdep in ${LIBDEPS[@]}; do
            if [ ! -f "${SCRIPT_PATH}/dist/${lib}" ]; then
                cp /mingw64/bin/${libdep} "${SCRIPT_PATH}/dist"
            fi
        done
    done

    echo "* Installing default theme..."
    mkdir -p "${SCRIPT_PATH}/dist/themes"
    cp -R /mingw64/share/themes/ "${SCRIPT_PATH}/dist/"
}

echo "Launched in ${SCRIPT_PATH}"
check_msys
check_golang
check_msys_packages
urtrator_build
urtrator_dist

echo "*** URTrator done."
