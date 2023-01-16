# COMP.SE.140

## Instructions for the teaching assistant

###  Implemented optional features
- Static analysis check is done using go vet (but it does look horrible since ./... didn't work in container for some reason)

### Instructions for examiner to test the system.
0. Make sure to clean any other compose140 exercises from containers and images first.
1. Start the system with "docker compose up -d" (or "docker-compose up" if using Docker Compose standlone)
2. Run commands in any order you wish:
- curl localhost:8083/state -X PUT -d "INIT"
- curl localhost:8083/state -X PUT -d "PAUSED"
- curl localhost:8083/state -X PUT -d "RUNNING"
- curl localhost:8083/state -X PUT -d "SHUTDOWN"
- curl localhost:8083/state
- curl localhost:8083/run-log
- curl localhost:8083/messages

NOTE: As per exercise instructions, "INIT" is logged into state log as "RUNNING".

NOTE: Unit tests are not located in "tests" - folder, because Go requires tests to reside in the same folder as the file that is being tested.

## Description of the CI/CD pipeline
### Version management; use of branches etc
- Project pushes to two remotes simulatenously, Github and local Gitlab with CI/CD pipeline.
- Gitlab repo is locally run in the Docker, using Gitlab CE.
- Github repo: 
### Building tools
- Built using go build
- Image: golang:alpine
### Testing; tools and test cases
- Testing uses go's testing: go test
- All the external outputs are tested.
### Packing
- Packaged as containers and run in a dind-container on the gitlab-runner.
### Deployment
- Deployed using dind-container on the gitlab-runner.
### Operating; monitoring
Operated using the following curl commands:
- curl localhost:8083/state -X PUT -d "INIT"
- curl localhost:8083/state -X PUT -d "PAUSED"
- curl localhost:8083/state -X PUT -d "RUNNING"
- curl localhost:8083/state -X PUT -d "SHUTDOWN"
- curl localhost:8083/state
- curl localhost:8083/run-log
- curl localhost:8083/messages
No monitoring implemented.

## Example runs of the pipeline

### Succeeding pipeline:
```
Running with gitlab-runner 15.5.1 (7178588d)
  on priviledged-runner k4L6VW4c
Preparing the "docker" executor
Using Docker executor with image docker ...
Starting service docker:dind ...
Pulling docker image docker:dind ...
Using docker image sha256:0e865cd51cb00cea31ca715846714e36966d0a404bbc50490633e3c1416caa6d for docker:dind with digest docker@sha256:2e0135466bcb3398e7f3943b87aef5c036dbaf1683805b8bfe992a477f7269e9 ...
Waiting for services to be up and running (timeout 30 seconds)...
Pulling docker image docker ...
Using docker image sha256:0f8d12a73562adf6588be88e37974abd42168017f375a1e160ba08a7ee3ffaa9 for docker with digest docker@sha256:75026b00c823579421c1850c00def301a6126b3f3f684594e51114c997f76467 ...
Preparing environment
00:01
Running on runner-k4l6vw4c-project-36-concurrent-0 via b597f1396b2f...
Getting source from Git repository
00:01
Fetching changes with git depth set to 20...
Reinitialized existing Git repository in /builds/gitlab-instance-9d36923c/COMP.SE.140/.git/
Checking out 46f993a0 as project...
Skipping Git submodules setup
Executing "step_script" stage of the job script
Using docker image sha256:0f8d12a73562adf6588be88e37974abd42168017f375a1e160ba08a7ee3ffaa9 for docker with digest docker@sha256:75026b00c823579421c1850c00def301a6126b3f3f684594e51114c997f76467 ...
$ docker-compose -H $DOCKER_HOST up --build -d
#1 [compse140-obse internal] load build definition from Dockerfile
#1 transferring dockerfile: 205B done
#1 DONE 0.1s
#2 [compse140-httpserv internal] load build definition from Dockerfile
#2 ...
#3 [compse140-rabbitmq internal] load build definition from Dockerfile
#3 transferring dockerfile: 162B done
#3 DONE 0.1s
#2 [compse140-httpserv internal] load build definition from Dockerfile
#2 transferring dockerfile: 213B done
#2 DONE 0.2s
#4 [compse140-gateway internal] load build definition from Dockerfile
#4 transferring dockerfile: 211B done
#4 DONE 0.2s
#5 [compse140-gateway internal] load .dockerignore
#5 transferring context: 2B done
#5 DONE 0.2s
#6 [compse140-imed internal] load build definition from Dockerfile
#6 transferring dockerfile: 195B done
#6 DONE 0.2s
#7 [compse140-orig internal] load build definition from Dockerfile
#7 transferring dockerfile: 195B done
#7 DONE 0.3s
#8 [compse140-rabbitmq internal] load .dockerignore
#8 transferring context: 2B done
#8 DONE 0.2s
#9 [compse140-orig internal] load .dockerignore
#9 transferring context: 2B done
#9 DONE 0.2s
#10 [compse140-imed internal] load .dockerignore
#10 transferring context: 2B done
#10 DONE 0.2s
#11 [compse140-obse internal] load .dockerignore
#11 transferring context: 2B done
#11 DONE 0.2s
#12 [compse140-gateway internal] load metadata for docker.io/library/golang:alpine
#12 ...
#13 [compse140-httpserv internal] load .dockerignore
#13 transferring context: 2B done
#13 DONE 0.2s
#12 [compse140-imed internal] load metadata for docker.io/library/golang:alpine
#12 ...
#14 [compse140-rabbitmq internal] load metadata for docker.io/library/rabbitmq:3-management
#14 DONE 2.5s
#12 [compse140-obse internal] load metadata for docker.io/library/golang:alpine
#12 DONE 2.6s
#15 [compse140-orig internal] load build context
#15 transferring context: 15.46kB done
#15 DONE 0.0s
#16 [compse140-obse 1/6] FROM docker.io/library/golang:alpine@sha256:2381c1e5f8350a901597d633b2e517775eeac7a6682be39225a93b22cfd0f8bb
#16 resolve docker.io/library/golang:alpine@sha256:2381c1e5f8350a901597d633b2e517775eeac7a6682be39225a93b22cfd0f8bb 0.0s done
#16 ...
#17 [compse140-httpserv internal] load build context
#17 transferring context: 2.36kB done
#17 DONE 0.1s
#18 [compse140-gateway internal] load build context
#18 transferring context: 30.24kB done
#18 DONE 0.1s
#19 [compse140-imed internal] load build context
#19 transferring context: 8.02kB done
#19 DONE 0.2s
#20 [compse140-obse internal] load build context
#20 transferring context: 8.26kB done
#20 DONE 0.2s
#16 [compse140-obse 1/6] FROM docker.io/library/golang:alpine@sha256:2381c1e5f8350a901597d633b2e517775eeac7a6682be39225a93b22cfd0f8bb
#16 sha256:2381c1e5f8350a901597d633b2e517775eeac7a6682be39225a93b22cfd0f8bb 1.65kB / 1.65kB done
#16 sha256:3183b343d01a6e8c9bee7b7c4eb7208518e68dc0cdac8623e1575820342472ed 1.16kB / 1.16kB done
#16 sha256:feb4bbda921c3f51bc84ae6f12414c53189d33e193e37d930b4a44076b7fb348 5.11kB / 5.11kB done
#16 ...
#21 [compse140-rabbitmq 1/2] FROM docker.io/library/rabbitmq:3-management@sha256:65a1412faeae95d260c8fcab7e0aa138fd4842726713b8496a2be87b26facca3
#21 resolve docker.io/library/rabbitmq:3-management@sha256:65a1412faeae95d260c8fcab7e0aa138fd4842726713b8496a2be87b26facca3 0.0s done
#21 sha256:65a1412faeae95d260c8fcab7e0aa138fd4842726713b8496a2be87b26facca3 1.42kB / 1.42kB done
#21 sha256:e7b677df5e85506845468791bf8ff9267523ceda705a32008b4cbddba460db09 2.62kB / 2.62kB done
#21 sha256:854c78c564838cd1e532115a0eda5dbdef260c164ac1849da3e3815ee0db67c9 8.52kB / 8.52kB done
#21 sha256:03dea6b5e70e630bb5bf621528463788e46cf2571f98f3f591e10633f1b650d9 334.49kB / 334.49kB 1.3s done
#21 sha256:cae1c228689d2f294439828746d67d06536e5cf331393d0529e18ce4c8f222f8 8.92kB / 8.92kB 1.3s done
#21 sha256:846c0b181fff0c667d9444f8378e8fcfa13116da8d308bf21673f7e4bea8d580 7.34MB / 28.58MB 5.4s
#21 sha256:3ddf939b446499c0937604d712f7e13340db9687e710d61c8b0b278700ba9e38 6.29MB / 51.10MB 5.4s
#21 sha256:92e035229b114a43bd976252b210632f191d41ae606b4eafaf029f0d06c0f454 5.19kB / 5.19kB 1.7s done
#21 sha256:d5e103ef1475feb2f050483aa4216a95752e3c27c02404fcc0a815938da1c0ad 2.10MB / 21.54MB 5.4s
#21 sha256:d5e103ef1475feb2f050483aa4216a95752e3c27c02404fcc0a815938da1c0ad 4.19MB / 21.54MB 6.6s
#21 sha256:3ddf939b446499c0937604d712f7e13340db9687e710d61c8b0b278700ba9e38 9.44MB / 51.10MB 7.4s
#21 sha256:846c0b181fff0c667d9444f8378e8fcfa13116da8d308bf21673f7e4bea8d580 9.44MB / 28.58MB 8.1s
#21 sha256:d5e103ef1475feb2f050483aa4216a95752e3c27c02404fcc0a815938da1c0ad 6.29MB / 21.54MB 8.5s
#21 sha256:3ddf939b446499c0937604d712f7e13340db9687e710d61c8b0b278700ba9e38 12.58MB / 51.10MB 9.4s
#21 sha256:d5e103ef1475feb2f050483aa4216a95752e3c27c02404fcc0a815938da1c0ad 8.39MB / 21.54MB 10.0s
#21 sha256:846c0b181fff0c667d9444f8378e8fcfa13116da8d308bf21673f7e4bea8d580 11.53MB / 28.58MB 11.3s
#21 sha256:3ddf939b446499c0937604d712f7e13340db9687e710d61c8b0b278700ba9e38 15.73MB / 51.10MB 11.4s
#21 sha256:d5e103ef1475feb2f050483aa4216a95752e3c27c02404fcc0a815938da1c0ad 10.49MB / 21.54MB 11.8s
#21 sha256:3ddf939b446499c0937604d712f7e13340db9687e710d61c8b0b278700ba9e38 18.87MB / 51.10MB 12.7s
#21 sha256:3ddf939b446499c0937604d712f7e13340db9687e710d61c8b0b278700ba9e38 22.02MB / 51.10MB 14.0s
#21 sha256:d5e103ef1475feb2f050483aa4216a95752e3c27c02404fcc0a815938da1c0ad 12.58MB / 21.54MB 14.9s
#21 sha256:3ddf939b446499c0937604d712f7e13340db9687e710d61c8b0b278700ba9e38 25.17MB / 51.10MB 15.4s
#21 sha256:846c0b181fff0c667d9444f8378e8fcfa13116da8d308bf21673f7e4bea8d580 13.63MB / 28.58MB 15.8s
#21 sha256:3ddf939b446499c0937604d712f7e13340db9687e710d61c8b0b278700ba9e38 28.31MB / 51.10MB 16.7s
#21 sha256:d5e103ef1475feb2f050483aa4216a95752e3c27c02404fcc0a815938da1c0ad 14.68MB / 21.54MB 17.7s
#21 sha256:3ddf939b446499c0937604d712f7e13340db9687e710d61c8b0b278700ba9e38 31.46MB / 51.10MB 18.1s
#21 sha256:3ddf939b446499c0937604d712f7e13340db9687e710d61c8b0b278700ba9e38 34.60MB / 51.10MB 19.6s
#21 sha256:d5e103ef1475feb2f050483aa4216a95752e3c27c02404fcc0a815938da1c0ad 16.78MB / 21.54MB 20.3s
#21 sha256:846c0b181fff0c667d9444f8378e8fcfa13116da8d308bf21673f7e4bea8d580 14.68MB / 28.58MB 20.9s
#21 sha256:3ddf939b446499c0937604d712f7e13340db9687e710d61c8b0b278700ba9e38 37.75MB / 51.10MB 20.9s
#21 sha256:3ddf939b446499c0937604d712f7e13340db9687e710d61c8b0b278700ba9e38 40.89MB / 51.10MB 22.2s
#21 sha256:3ddf939b446499c0937604d712f7e13340db9687e710d61c8b0b278700ba9e38 44.04MB / 51.10MB 23.4s
#21 sha256:d5e103ef1475feb2f050483aa4216a95752e3c27c02404fcc0a815938da1c0ad 18.87MB / 21.54MB 23.7s
#21 sha256:846c0b181fff0c667d9444f8378e8fcfa13116da8d308bf21673f7e4bea8d580 16.78MB / 28.58MB 24.7s
#21 sha256:3ddf939b446499c0937604d712f7e13340db9687e710d61c8b0b278700ba9e38 47.19MB / 51.10MB 24.7s
#21 sha256:3ddf939b446499c0937604d712f7e13340db9687e710d61c8b0b278700ba9e38 50.33MB / 51.10MB 26.0s
#21 sha256:3ddf939b446499c0937604d712f7e13340db9687e710d61c8b0b278700ba9e38 51.10MB / 51.10MB 26.1s done
#21 sha256:f0926bc87948c9673fbbee4d4bff60af175dfae14be1b5b8e6a3d0c0037099ef 0B / 274B 26.2s
#21 sha256:f0926bc87948c9673fbbee4d4bff60af175dfae14be1b5b8e6a3d0c0037099ef 274B / 274B 26.6s
#21 sha256:d5e103ef1475feb2f050483aa4216a95752e3c27c02404fcc0a815938da1c0ad 20.97MB / 21.54MB 26.7s
#21 sha256:f0926bc87948c9673fbbee4d4bff60af175dfae14be1b5b8e6a3d0c0037099ef 274B / 274B 26.6s done
#21 sha256:9ae105c01998d72f7a495f74a8674275001eb3ee74e232d458dce9091b2706c3 0B / 107B 26.7s
#21 sha256:d5e103ef1475feb2f050483aa4216a95752e3c27c02404fcc0a815938da1c0ad 21.54MB / 21.54MB 26.9s done
#21 sha256:9ae105c01998d72f7a495f74a8674275001eb3ee74e232d458dce9091b2706c3 107B / 107B 26.9s done
#21 sha256:553e1f2f08f51bf343735e0e580955a4c851bd0ba8b2a2f0ed461538e37a15bb 0B / 501B 27.0s
#21 sha256:46f5869ee28d3f9a1362b44150f944d7f8ea10bb2be4a01f1c9ee635619af4be 0B / 834B 27.0s
#21 sha256:553e1f2f08f51bf343735e0e580955a4c851bd0ba8b2a2f0ed461538e37a15bb 501B / 501B 27.2s done
#21 sha256:2ed5fb86c5042c27df38884a97dc40996ae0891099ef6365b2ac9b1d4ca8b081 0B / 9.90MB 27.3s
#21 sha256:46f5869ee28d3f9a1362b44150f944d7f8ea10bb2be4a01f1c9ee635619af4be 834B / 834B 27.3s done
#21 ...
#16 [compse140-obse 1/6] FROM docker.io/library/golang:alpine@sha256:2381c1e5f8350a901597d633b2e517775eeac7a6682be39225a93b22cfd0f8bb
#16 sha256:8921db27df2831fa6eaa85321205a2470c669b855f3ec95d5a3c2b46de0442c9 0B / 3.37MB 27.3s
#16 sha256:a2f8637abd914a8a62416e027a351293d0472bc4b4f44383c6f425fd0e03861c 0B / 284.81kB 30.8s
#16 sha256:8921db27df2831fa6eaa85321205a2470c669b855f3ec95d5a3c2b46de0442c9 1.05MB / 3.37MB 31.1s
#16 sha256:a2f8637abd914a8a62416e027a351293d0472bc4b4f44383c6f425fd0e03861c 284.81kB / 284.81kB 31.3s done
#16 sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 0B / 122.33MB 31.5s
#16 sha256:8921db27df2831fa6eaa85321205a2470c669b855f3ec95d5a3c2b46de0442c9 2.10MB / 3.37MB 33.4s
#16 sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 6.29MB / 122.33MB 34.8s
#16 sha256:8921db27df2831fa6eaa85321205a2470c669b855f3ec95d5a3c2b46de0442c9 3.15MB / 3.37MB 36.2s
#16 sha256:8921db27df2831fa6eaa85321205a2470c669b855f3ec95d5a3c2b46de0442c9 3.37MB / 3.37MB 36.6s done
#16 extracting sha256:8921db27df2831fa6eaa85321205a2470c669b855f3ec95d5a3c2b46de0442c9
#16 sha256:4f26d270037d8578282b40f86c0a1816fc5d034d6213f7dcab8440056a589309 0B / 155B 36.7s
#16 sha256:8921db27df2831fa6eaa85321205a2470c669b855f3ec95d5a3c2b46de0442c9 3.37MB / 3.37MB 36.6s done
#16 extracting sha256:8921db27df2831fa6eaa85321205a2470c669b855f3ec95d5a3c2b46de0442c9 0.1s done
#16 extracting sha256:a2f8637abd914a8a62416e027a351293d0472bc4b4f44383c6f425fd0e03861c 0.0s done
#16 ...
#21 [compse140-rabbitmq 1/2] FROM docker.io/library/rabbitmq:3-management@sha256:65a1412faeae95d260c8fcab7e0aa138fd4842726713b8496a2be87b26facca3
#21 sha256:846c0b181fff0c667d9444f8378e8fcfa13116da8d308bf21673f7e4bea8d580 25.17MB / 28.58MB 37.5s
#21 sha256:2ed5fb86c5042c27df38884a97dc40996ae0891099ef6365b2ac9b1d4ca8b081 9.90MB / 9.90MB 30.9s done
#21 ...
#16 [compse140-obse 1/6] FROM docker.io/library/golang:alpine@sha256:2381c1e5f8350a901597d633b2e517775eeac7a6682be39225a93b22cfd0f8bb
#16 sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 12.58MB / 122.33MB 38.0s
#16 sha256:4f26d270037d8578282b40f86c0a1816fc5d034d6213f7dcab8440056a589309 155B / 155B 38.7s done
#16 sha256:4f26d270037d8578282b40f86c0a1816fc5d034d6213f7dcab8440056a589309 155B / 155B 38.7s done
#16 sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 18.87MB / 122.33MB 40.2s
#16 ...
#21 [compse140-rabbitmq 1/2] FROM docker.io/library/rabbitmq:3-management@sha256:65a1412faeae95d260c8fcab7e0aa138fd4842726713b8496a2be87b26facca3
#21 sha256:846c0b181fff0c667d9444f8378e8fcfa13116da8d308bf21673f7e4bea8d580 28.58MB / 28.58MB 39.1s done
#21 extracting sha256:846c0b181fff0c667d9444f8378e8fcfa13116da8d308bf21673f7e4bea8d580 0.4s done
#21 extracting sha256:03dea6b5e70e630bb5bf621528463788e46cf2571f98f3f591e10633f1b650d9 0.0s done
#21 extracting sha256:cae1c228689d2f294439828746d67d06536e5cf331393d0529e18ce4c8f222f8 done
#21 extracting sha256:3ddf939b446499c0937604d712f7e13340db9687e710d61c8b0b278700ba9e38 0.5s done
#21 extracting sha256:92e035229b114a43bd976252b210632f191d41ae606b4eafaf029f0d06c0f454 done
#21 extracting sha256:d5e103ef1475feb2f050483aa4216a95752e3c27c02404fcc0a815938da1c0ad 0.3s done
#21 extracting sha256:f0926bc87948c9673fbbee4d4bff60af175dfae14be1b5b8e6a3d0c0037099ef done
#21 extracting sha256:9ae105c01998d72f7a495f74a8674275001eb3ee74e232d458dce9091b2706c3 done
#21 extracting sha256:553e1f2f08f51bf343735e0e580955a4c851bd0ba8b2a2f0ed461538e37a15bb done
#21 extracting sha256:46f5869ee28d3f9a1362b44150f944d7f8ea10bb2be4a01f1c9ee635619af4be done
#21 extracting sha256:2ed5fb86c5042c27df38884a97dc40996ae0891099ef6365b2ac9b1d4ca8b081 0.2s done
#21 DONE 41.3s
#16 [compse140-obse 1/6] FROM docker.io/library/golang:alpine@sha256:2381c1e5f8350a901597d633b2e517775eeac7a6682be39225a93b22cfd0f8bb
#16 sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 25.17MB / 122.33MB 42.0s
#16 ...
#22 [compse140-rabbitmq 2/2] RUN rabbitmq-plugins enable --offline rabbitmq_mqtt rabbitmq_federation_management rabbitmq_stomp
#0 0.640 Enabling plugins on node rabbit@buildkitsandbox:
#0 0.640 rabbitmq_mqtt
#0 0.640 rabbitmq_federation_management
#0 0.640 rabbitmq_stomp
#0 0.959 The following plugins have been configured:
#0 0.959   rabbitmq_federation
#0 0.959   rabbitmq_federation_management
#0 0.959   rabbitmq_management
#0 0.959   rabbitmq_management_agent
#0 0.959   rabbitmq_mqtt
#0 0.959   rabbitmq_prometheus
#0 0.959   rabbitmq_stomp
#0 0.959   rabbitmq_web_dispatch
#0 0.959 Applying plugin configuration to rabbit@buildkitsandbox...
#0 0.960 The following plugins have been enabled:
#0 0.960   rabbitmq_federation
#0 0.960   rabbitmq_federation_management
#0 0.960   rabbitmq_mqtt
#0 0.960   rabbitmq_stomp
#0 0.960 
#0 0.960 set 8 plugins.
#0 0.960 Offline change; changes will take effect at broker restart.
#22 DONE 1.2s
#16 [compse140-obse 1/6] FROM docker.io/library/golang:alpine@sha256:2381c1e5f8350a901597d633b2e517775eeac7a6682be39225a93b22cfd0f8bb
#16 ...
#23 [compse140-rabbitmq] exporting to image
#23 exporting layers 0.0s done
#23 writing image sha256:afd3fa418a1741146c51637819fcfb75ef62a53952c605f88f3d59d243f7ab68 done
#23 naming to docker.io/library/compse140-rabbitmq done
#23 DONE 0.1s
#16 [compse140-obse 1/6] FROM docker.io/library/golang:alpine@sha256:2381c1e5f8350a901597d633b2e517775eeac7a6682be39225a93b22cfd0f8bb
#16 sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 31.46MB / 122.33MB 43.8s
#16 sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 37.75MB / 122.33MB 45.6s
#16 sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 44.04MB / 122.33MB 47.4s
#16 sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 50.33MB / 122.33MB 49.3s
#16 sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 56.62MB / 122.33MB 51.0s
#16 sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 62.91MB / 122.33MB 52.8s
#16 sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 69.21MB / 122.33MB 54.7s
#16 sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 75.50MB / 122.33MB 56.5s
#16 sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 81.79MB / 122.33MB 58.5s
#16 sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 88.08MB / 122.33MB 60.2s
#16 sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 94.37MB / 122.33MB 62.0s
#16 sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 100.66MB / 122.33MB 63.9s
#16 sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 106.95MB / 122.33MB 65.6s
#16 sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 113.25MB / 122.33MB 67.5s
#16 sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 119.54MB / 122.33MB 69.2s
#16 sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 122.33MB / 122.33MB 70.3s done
#16 extracting sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613
#16 extracting sha256:d48e7ca896ec974241d86439d34188186e085705e42d914d8b7757c27eea8613 1.7s done
#16 extracting sha256:4f26d270037d8578282b40f86c0a1816fc5d034d6213f7dcab8440056a589309 done
#16 DONE 72.6s
#24 [compse140-httpserv 2/6] WORKDIR /app
#24 DONE 0.4s
#25 [compse140-obse 3/6] COPY go.mod ./
#25 DONE 0.1s
#26 [compse140-imed 3/6] COPY go.mod ./
#26 ...
#27 [compse140-gateway 3/6] COPY go.mod ./
#27 DONE 0.1s
#28 [compse140-httpserv 3/6] COPY go.mod ./
#28 DONE 0.2s
#26 [compse140-imed 3/6] COPY go.mod ./
#26 DONE 0.2s
#29 [compse140-orig 3/6] COPY go.mod ./
#29 DONE 0.2s
#30 [compse140-orig 4/6] RUN go mod download
#30 ...
#31 [compse140-httpserv 4/6] RUN go mod download
#0 1.069 go: no module dependencies to download
#31 DONE 1.1s
#32 [compse140-httpserv 5/6] COPY . /app
#32 DONE 0.0s
#33 [compse140-httpserv 6/6] RUN go build -o /app/httpserv
#33 ...
#34 [compse140-obse 4/6] RUN go mod download
#34 DONE 1.9s
#35 [compse140-obse 5/6] COPY . /app
#35 DONE 0.1s
#36 [compse140-obse 6/6] RUN go build -o /app/obse
#36 ...
#33 [compse140-httpserv 6/6] RUN go build -o /app/httpserv
#33 DONE 1.3s
#23 [compse140-httpserv] exporting to image
#23 exporting layers 0.1s done
#23 writing image sha256:1f1162696d7e323bfbf44c1538d1ade9d45c4422691d4692a861fdb04c167978 done
#23 naming to docker.io/library/compse140-httpserv done
#23 DONE 0.2s
#37 [compse140-imed 4/6] RUN go mod download
#37 DONE 2.9s
#38 [compse140-imed 5/6] COPY . ./
#38 DONE 0.0s
#39 [compse140-imed 6/6] RUN go build -o /imed
#39 ...
#36 [compse140-obse 6/6] RUN go build -o /app/obse
#36 DONE 1.3s
#23 [compse140-obse] exporting to image
#23 exporting layers 0.1s done
#23 writing image sha256:a8abacc7c52c487bc6a8302071479f6bb0d3cf71ef00357a8a213a504d5b814d done
#23 naming to docker.io/library/compse140-obse done
#23 DONE 0.3s
#40 [compse140-gateway 4/6] RUN go mod download
#40 ...
#39 [compse140-imed 6/6] RUN go build -o /imed
#39 DONE 1.8s
#23 [compse140-imed] exporting to image
#23 exporting layers 0.1s done
#23 writing image sha256:3d73637e92e705f1289a04be3dab4e202ff800712a1a3366daa5f2a019ac5b2d done
#23 naming to docker.io/library/compse140-imed done
#23 DONE 0.4s
#40 [compse140-gateway 4/6] RUN go mod download
#40 ...
#30 [compse140-orig 4/6] RUN go mod download
#30 DONE 12.2s
#41 [compse140-orig 5/6] COPY . ./
#41 DONE 0.1s
#42 [compse140-orig 6/6] RUN go build -o /orig
#42 DONE 1.3s
#40 [compse140-gateway 4/6] RUN go mod download
#40 ...
#23 [compse140-orig] exporting to image
#23 exporting layers
#23 exporting layers 0.7s done
#23 writing image sha256:98c10fffbfcff4b586d4f4f0a385c1e9835633efc1ffeb34a3756d1b184a1fb9 done
#23 naming to docker.io/library/compse140-orig done
#23 DONE 1.1s
#40 [compse140-gateway 4/6] RUN go mod download
#40 DONE 22.4s
#43 [compse140-gateway 5/6] COPY . /app
#43 DONE 0.0s
#44 [compse140-gateway 6/6] RUN go build -o /app/gateway
#44 DONE 3.2s
#23 [compse140-gateway] exporting to image
#23 exporting layers
#23 exporting layers 1.5s done
#23 writing image sha256:fbd342c467feb8244b2f5f69c2fdee0ec5888b63cc5ad44702812ad39105c620 done
#23 naming to docker.io/library/compse140-gateway done
#23 DONE 2.7s
Network compse140_default  Creating
Network compse140_default  Created
Volume "compse140_message-storage"  Creating
Volume "compse140_message-storage"  Created
Container compse140-rabbitmq-1  Creating
Container compse140-rabbitmq-1  Created
Container compse140-imed-1  Creating
Container compse140-obse-1  Creating
Container compse140-httpserv-1  Creating
Container compse140-gateway-1  Creating
Container compse140-imed-1  Created
Container compse140-httpserv-1  Created
Container compse140-gateway-1  Created
Container compse140-obse-1  Created
Container compse140-orig-1  Creating
Container compse140-orig-1  Created
Container compse140-rabbitmq-1  Starting
Container compse140-rabbitmq-1  Started
Container compse140-gateway-1  Starting
Container compse140-rabbitmq-1  Waiting
Container compse140-rabbitmq-1  Waiting
Container compse140-rabbitmq-1  Waiting
Container compse140-gateway-1  Started
Container compse140-rabbitmq-1  Healthy
Container compse140-rabbitmq-1  Healthy
Container compse140-rabbitmq-1  Healthy
Container compse140-httpserv-1  Starting
Container compse140-obse-1  Starting
Container compse140-imed-1  Starting
Container compse140-obse-1  Started
Container compse140-httpserv-1  Started
Container compse140-imed-1  Started
Container compse140-orig-1  Starting
Container compse140-orig-1  Started
$ echo "Application successfully deployed."
Application successfully deployed.
Job succeeded
```
-----------------------------
### Failing pipeline:

```
Running with gitlab-runner 15.5.1 (7178588d)
  on priviledged-runner k4L6VW4c
Preparing the "docker" executor
00:46
Using Docker executor with image golang:alpine ...
Pulling docker image golang:alpine ...
Using docker image sha256:cae57157ceaa07c0fffad419f2c4cedc683c31981f466a0fb9fd8cdd434e05d8 for golang:alpine with digest golang@sha256:a9b24b67dc83b3383d22a14941c2b2b2ca6a103d805cac6820fd1355943beaf1 ...
Preparing environment
00:02
Running on runner-k4l6vw4c-project-36-concurrent-0 via b597f1396b2f...
Getting source from Git repository
00:02
Fetching changes with git depth set to 20...
Reinitialized existing Git repository in /builds/gitlab-instance-9d36923c/COMP.SE.140/.git/
Checking out caa18e8d as project...
Skipping Git submodules setup
Executing "step_script" stage of the job script
00:47
Using docker image sha256:cae57157ceaa07c0fffad419f2c4cedc683c31981f466a0fb9fd8cdd434e05d8 for golang:alpine with digest golang@sha256:a9b24b67dc83b3383d22a14941c2b2b2ca6a103d805cac6820fd1355943beaf1 ...
$ echo "Waiting for 30 seconds..."
Waiting for 30 seconds...
$ sleep 30
$ cd ./gateway
$ CGO_ENABLED=0 go test
go: downloading github.com/docker/docker v20.10.22+incompatible
go: downloading github.com/gin-gonic/gin v1.8.1
go: downloading github.com/stretchr/testify v1.8.1
go: downloading github.com/davecgh/go-spew v1.1.1
go: downloading github.com/pmezard/go-difflib v1.0.0
go: downloading gopkg.in/yaml.v3 v3.0.1
go: downloading github.com/gin-contrib/sse v0.1.0
go: downloading golang.org/x/net v0.2.0
go: downloading github.com/mattn/go-isatty v0.0.16
go: downloading google.golang.org/protobuf v1.28.1
go: downloading github.com/ugorji/go/codec v1.2.7
go: downloading github.com/pelletier/go-toml/v2 v2.0.6
go: downloading gopkg.in/yaml.v2 v2.4.0
go: downloading github.com/go-playground/validator/v10 v10.11.1
go: downloading golang.org/x/sys v0.2.0
go: downloading github.com/leodido/go-urn v1.2.1
go: downloading github.com/go-playground/universal-translator v0.18.0
go: downloading golang.org/x/crypto v0.3.0
go: downloading golang.org/x/text v0.4.0
go: downloading github.com/go-playground/locales v0.14.0
go: downloading github.com/docker/distribution v2.8.1+incompatible
go: downloading github.com/opencontainers/go-digest v1.0.0
go: downloading github.com/opencontainers/image-spec v1.0.2
go: downloading github.com/docker/go-connections v0.4.0
go: downloading github.com/pkg/errors v0.9.1
go: downloading github.com/docker/go-units v0.5.0
go: downloading github.com/sirupsen/logrus v1.9.0
go: downloading github.com/gogo/protobuf v1.3.2
# github.com/KaluTheKova/COMP.SE.140/httpserv
./gateway.go:164:3: log.Panic call has possible formatting directive %s
./gateway.go:173:3: log.Panic call has possible formatting directive %s
FAIL	github.com/KaluTheKova/COMP.SE.140/httpserv [build failed]
ERROR: Job failed: exit code 2
```
----------------------
## Reflections
### Main learnings and worst difficulties
Main learnings include a yearning for Jenkins automation server or any pipeline automation server where I could replay a pipeline and adjust the gitlab-ci.yml without having to always recommit.

I also learned that one can push into multiple repositories simultaneously. Previously, I had never needed this feature, but now it proved to be interesting to implement and use. So here I configured the project so that it pushes simultaneously to my Github and to my locally hosted Gitlab that ran the pipeline.

It was also quite interesting to setup Gitlab on my local docker. In real world jobs, DevOps engineer would be responsible for setting a up a company's Gitlab and managing it, so this was quite an interesting thing to practise.

My worst difficulties were actually with RabbitMQ's Golang implementation. The library had some under-the-hood implementations that caused a lot of lost time and headache, such as somewhat unsual use of channels without documenting them.

Also, implementing local deployment was a somewhat pain. I would have liked to deploy to Heroku, but Heroku is no longer free starting November 28, 2022.

Instead of using HTTP request from Gateway to and resepond on that, I wanted to try out using Docker CLI via Gateway's code. It was interesting to implement INIT, PAUSED, RUNNING and SHUTDOWN via that way. Using code to run container functions is also pretty DevOps thinking. 

Of course, in a real world production application this might cause problems if the CLI library is reliant on Docker, since the production application would have to always run on Docker instead of being portable to another container engine. In a real world application, I probably would have used the HTTP method. In this case however, we can assume that Docker engine is used to run the application.

And naturally, I would have implemented error handling and basically everything more robustly.

### Amount effort (hours) used
Around 50 I believe.
