#!/bin/bash
#go get -u golang.org/x/crypto/bcrypt
#go get -u golang.org/x/crypto/ssh/terminal
#go run Script/main.go

grpName="registry-group"
grpRegion="westeurope"
strgConfig="config"
strgData="data"
authConfig='htpasswd'


#az deployment sub create --location $grpRegion --template-uri https://raw.githubusercontent.com/groovy-sky/self-registry/master/Template/group.json --parameters Location=$grpRegion Name=$grpName

strgName=$(az deployment group create --resource-group $grpName --template-uri https://raw.githubusercontent.com/groovy-sky/self-registry/master/Template/storage.json --parameters containerConfigName=$strgConfig containerDataName=$strgData | jq -r '.properties.outputs | map(.value) | add' )

strgKey=$(az storage account keys list --account-name $strgName | jq -r '. | first.value')

az storage blob upload --account-name $strgName --account-key $strgKey --container-name $strgConfig --name $authConfig --file $authConfig

registryName=$(az deployment group create --resource-group $grpName --template-uri https://raw.githubusercontent.com/groovy-sky/self-registry/master/Template/registry.json --parameters htpasswdPath=$strgConfig/$authConfig storageName=$strgName storageKey=$strgKey storageContainer=$strgData | jq -r '.properties.outputs | map(.value) | add' )

#https://docs.microsoft.com/en-us/cli/azure/webapp/config/storage-account?view=azure-cli-latest
