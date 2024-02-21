#! /bin/bash
echo "Installing Task"
curl -sS https://taskfile.dev/install.sh -o install.sh
chmod +x ./install.sh
./install.sh -d -b /usr/local/bin
rm -f ./install.sh
echo "Installed Task"
