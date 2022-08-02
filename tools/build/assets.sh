#!/bin/sh -e

if test "$@" ; then
  echo "copying from $@"
else 
  echo "path to prepared misskey root required"
  exit 1
fi

TARGET=$(./build/misskey target)
ASSETS="cmd/misskey/assets"
PUBLIC="$ASSETS/public"
MISSKEY_ROOT=$@

rm -rf $PUBLIC
mkdir $PUBLIC

# general files
cp -r "$MISSKEY_ROOT/packages/backend/assets" "$PUBLIC/static-assets"
cp -r "$MISSKEY_ROOT/packages/client/assets" "$PUBLIC/client-assets"
cp -r "$MISSKEY_ROOT/built/_client_dist_" "$PUBLIC/assets"

# special cases
cp "$MISSKEY_ROOT/packages/backend/assets/favicon.ico" "$PUBLIC/favicon.ico"
cp "$MISSKEY_ROOT/packages/backend/assets/apple-touch-icon.png" "$PUBLIC"
cp "$MISSKEY_ROOT/packages/backend/assets/robots.txt" "$PUBLIC"
cp "$MISSKEY_ROOT/built/_client_dist_/sw.$TARGET.js" "$PUBLIC/sw.js"
cp -r "$MISSKEY_ROOT/packages/client/node_modules/@discordapp/twemoji/dist/svg" "$PUBLIC/twemoji"

# this is done this way for the sole purpose of not including misskey code
# there is no good reason for this otherwise
./build/prairie -w -o "$ASSETS/template" -i "$@"
patch --posix -d "cmd/misskey/assets/template/" < cmd/misskey/patch/template.patch

tar -zcvf "build/assets-$TARGET.tar.gz" -C "cmd/misskey/assets/" .
