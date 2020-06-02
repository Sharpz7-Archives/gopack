#/bin/bash

# Uninstall any previous versions.
echo "Checking for any previous version..."
sudo rm -r /usr/local/bin/gopack

# Download and unpack
wget https://github.com/Sharpz7/gopack/releases/download/XXXXX/linux.tar.gz
sudo tar -C /usr/local/bin/ -zxvf linux.tar.gz
rm -r linux.tar.gz

# Permissions
chmod u+x /usr/local/bin/gopack

echo ""
echo "GOPACK IS NOW INSTALLED"
echo "======================="
echo "Do gopack -h for more info!"