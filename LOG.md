# Work Log

## 2021-03-14
For the sake of keeping historical information on work I did for this repository along with the tools I used, I will be keeping this work log. Previous work that is not logged on this file will be discussed later. 

Here are the list of things I did prior to logging work:
* Creation of the `address-book` application
* Application file structure
* Unit test using [Testify](https://github.com/stretchr/testify)
* Building Docker image for Go applications
* Basic commands for deploying the app in Kubernetes
* [Newman](https://github.com/postmanlabs/newman) for preparation of API/end-to-end testing
* Building Docker image for MySQL datasource
* Azure DevOps pipeline

Today, I've been continuing work on the build pipeline specifically on the _SonarCloud_ integration. Details on how to integrate _Azure Devops_ can easily be found in the web. Previous instructions left me with the following _task_ definitions added to _azure-pipelines.yml_:
```
- task: SonarCloudPrepare@1
  inputs:
    SonarCloud: 'sonar-source'
    organization: 'esgeronimo-github'
    scannerMode: 'CLI'
    configMode: 'file'
- task: SonarCloudAnalyze@1
- task: SonarCloudPublish@1
  inputs:
    pollingTimeoutSec: '300'
```
I left the pipeline having the following error on `SonarCloudAnalyze` task:
```
You must define the following mandatory properties for 'Unknown': sonar.projectKey
```
Issue can be addressed by defining the value as parameter during Sonar execution or adding `sonar-project.properties`. Given the abstraction of running Sonar-cloud related tasks in Azure Devops, I preferred doing the latter approach to have more control on such kind of parameters. The said file is added under the base directory of the project and contained the following values:
```
sonar.projectKey=esgeronimo_address-book
sonar.organization=esgeronimo-github
```
>IMPORTANT: Values for both `projectKey` and `organization` cannot be arbitrarily declared. If the project is already present in SpringCloud (probaby because of previous POC/integration attempts), refer to the dashboard page of the project and look for _Project Key_ and _Organization_ values. Have a look at [address-book's](https://sonarcloud.io/dashboard?id=esgeronimo_address-book) page for sample.