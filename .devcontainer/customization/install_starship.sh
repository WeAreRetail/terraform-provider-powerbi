#! /bin/bash
echo "Installing Starship"
curl -sS https://starship.rs/install.sh -o install.sh
chmod +x ./install.sh
./install.sh --yes
rm -f ./install.sh
echo "Installed Starship"
