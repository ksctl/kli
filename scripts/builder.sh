#!/bin/bash

cd .. || echo -e "\033[1;31mUnable to cd into ksctl root\033[0m\n"

go mod tidy
go build -v -o kli .
chmod +x kli

sudo mv -v kli /usr/local/bin/kli

echo -e "\033[1;32mINSTALL COMPLETE\033[0m\n"
