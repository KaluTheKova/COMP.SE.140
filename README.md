# COMP.SE.140



1. Instructions for the teaching assistant
Implemented optional features
- Static analysis check using go vet

Instructions for examiner to test the system.
TO DO

- Project pushes to two remotes simulatenously, Github and local Gitlab with CI/CD pipeline.

NOTE: Unit tests are not located in "tests" - folder, because Go requires tests to reside in the same folder as the file that is being tested.

2. Description of the CI/CD pipeline
Briefly document all steps:
• Version management; use of branches etc
• Building tools
• Testing; tools and test cases
• Packing
• Deployment
• Operating; monitoring


3. Example runs of the pipeline
Include some kind of log of both failing test and passing.

### Succeeding pipeline:

### Failing pipeline:
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

4. Reflections
### Main learnings and worst difficulties:
Main learnings include a yearning for Jenkins automation server or any pipeline automation server where I could replay a pipeline and adjust the gitlab-ci.yml without having to always recommit.

I also learned that one can push into multiple repositories simultaneously. Previously, I had never needed this feature, but now it proved to be interesting to implement and use. So here I configured the project so that it pushes simultaneously to my Github and to my locally hosted Gitlab that ran the pipeline.

My worst difficulties were actually with RabbitMQ's Golang implementation. The library had some under-the-hood implementations that caused a lot of lost time and headache, such as somewhat unsual use of channels without documenting them. 

Instead of using HTTP request from Gateway to and resepond on that, I wanted to try out using Docker CLI via Gateway's code. It was interesting to implement INIT, PAUSED, RUNNING and SHUTDOWN via that way. Using code to run container functions is also pretty DevOps thinking. 

Of course, in a real world production application this might cause problems if the CLI library is reliant on Docker, since the production application would have to always run on Docker instead of being portable to another container engine. In a real world application, I probably would have used the HTTP method. In this case however, we can assume that Docker engine is used to run the application.

And naturally, I would have implemented errorhandling and basically everything more robustly.

### Amount effort (hours) used
Around 50 I believe.
