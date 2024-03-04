#!/bin/sh

test -f DevMate-Installer.dmg && rm DevMate-Installer.dmg

wails build -u -platform darwin/universal
create-dmg \
    --volname "DevMate Installer" \
    --volicon "build/bin/matedev.app/Contents/Resources/iconfile.icns" \
    --icon-size 75 \
    --window-size 600 400 \
    --icon "DevMate.app" 200 170 \
    --app-drop-link 400 170 \
    --hide-extension "DevMate.app" \
    "DevMate-Installer.dmg" \
    "build/bin/"
