#!/bin/bash
#go get -u golang.org/x/crypto/bcrypt
#go get -u golang.org/x/crypto/ssh/terminal
#go run Script/main.go

grpName="registry-group"
grpRegion="westeurope"
strgConfig="config"
strgData="data"

#az deployment sub create --location $grpRegion --template-uri https://raw.githubusercontent.com/groovy-sky/self-registry/master/Template/group.json --parameters Location=$grpRegion Name=$grpName

az deployment group create --resource-group $grpName --template-uri https://raw.githubusercontent.com/groovy-sky/self-registry/master/Template/storage.json --parameters containerConfigName=$strgConfig containerDataName=$strgData
