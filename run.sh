#!/bin/bash

# Variables definition
grpName="demo-registry-group"
grpRegion="westeurope"
strgConfig="config"
strgData="data"
authConfig='htpasswd'

echo -e "[INFO] Working on htpasswd generation\n"
go get -u golang.org/x/crypto/bcrypt
go get -u golang.org/x/crypto/ssh/terminal
go run Script/main.go

echo -e "[INFO] Resource group deployment\n"
az deployment sub create --location $grpRegion --template-uri https://raw.githubusercontent.com/groovy-sky/self-registry/master/Template/group.json --parameters Location=$grpRegion Name=$grpName

echo -e "[INFO] Storage account deployment\n"
strgName=$(az deployment group create --resource-group $grpName --template-uri https://raw.githubusercontent.com/groovy-sky/self-registry/master/Template/storage.json --parameters containerConfigName=$strgConfig containerDataName=$strgData | jq -r '.properties.outputs | map(.value) | add' )

echo -e "[INFO] Uploading configuration to storage\n"
strgKey=$(az storage account keys list --account-name $strgName | jq -r '. | first.value')

az storage blob upload --account-name $strgName --account-key $strgKey --container-name $strgConfig --name $authConfig --file $authConfig

echo -e "[INFO] Registry deployment\n"
registryName=$(az deployment group create --resource-group $grpName --template-uri https://raw.githubusercontent.com/groovy-sky/self-registry/master/Template/registry.json --parameters htpasswdPath=$strgConfig/$authConfig storageName=$strgName storageKey=$strgKey storageContainer=$strgData | jq -r '.properties.outputs | map(.value) | add' )

echo -e "[INFO] Mounting storage container\n"
#https://docs.microsoft.com/en-us/cli/azure/webapp/config/storage-account?view=azure-cli-latest
az webapp config storage-account add --access-key $strgKey --account-name $strgName --share-name $strgConfig --storage-type AzureBlob --mount-path /$strgConfig --name $registryName --custom-id $strgConfig  --resource-group $grpName

echo -e "[INFO] Registry $registryName.azurewebsites.net is ready for use"
