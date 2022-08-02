#!/bin/sh

ASSETS="cmd/misskey/assets"
ASSETS_URL="https://cronut.cafe/~rand/misskey/assets-$(./build/misskey target).tar.gz"

rm -rf "$ASSETS/public" "$ASSETS/template"
if type "fetch" >/dev/null 2>/dev/null ; then
  fetch -o "build/assets" "$ASSETS_URL"
elif type "curl" >/dev/null 2>/dev/null ; then
  curl "$ASSETS_URL" > "build/assets"
elif type "wget" >/dev/null 2>/dev/null ; then
  wget -o "build/assets" "$ASSETS_URL";
fi

tar -zx -f "build/assets" -C "cmd/misskey/assets"
rm "build/assets"
