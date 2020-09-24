# Running a self-hosted Docker Registry on Azure App Service

This repository contains [the script](https://github.com/groovy-sky/self-registry/blob/master/run.sh) for deploying a self-hosted Docker Registry running on Azure App Service, which uses Storage Account for storing images. Detailed instruction how-to run it you can find [here](https://github.com/groovy-sky/azure/tree/master/docker-private-registry#introduction).

You can use Azure Cloud Shell in Bash mode for the deployment, by executing commands below:

```
export grpName="demo-registry-group"                                                                                 
export grpRegion="westeurope"                                                                                        
export strgConfig="config"                                                                                           
export strgData="data"                                                                                               
export authConfig="htpasswd" 

[ ! -d "self-registry/.git" ] && git clone https://github.com/groovy-sky/self-registry
cd self-registry && git pull
./run.sh

```

Deployment process should take less than 5 minutes:

![](https://github.com/groovy-sky/azure/raw/master/images/docker/registy_build.gif)
