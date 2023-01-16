## Gitlab runner

Option 1: Use local system volume mounts to start the Runner container
This example uses the local system for the configuration volume that is mounted into the gitlab-runner container. This volume is used for configs and other resources.

docker run -d --name gitlab-runner --restart always \
  -v /srv/gitlab-runner/config:/etc/gitlab-runner \
  -v /var/run/docker.sock:/var/run/docker.sock \
  gitlab/gitlab-runner:latest



## Gitlab temp password
https://www.czerniga.it/2021/11/14/how-to-install-gitlab-using-docker-compose/
docker exec -it web-1 grep 'Password:' /etc/gitlab/initial_root_password

var/opt/gitlab/gitlab-kas/

## Register the runner
https://docs.gitlab.com/runner/register/

Vanha:
docker run --rm -it -v /srv/gitlab-runner/config:/etc/gitlab-runner gitlab/gitlab-runner register

Uusi:
docker exec -it gitlab-runner gitlab-runner register --url "git@localhost" --clone-url "git@localhost"

User: sieni
url: 
git@localhost:gitlab-instance-9d36923c/
http://localhost/gitlab-instance-9d36923c/
http://localhost:8080/gitlab-instance-9d36923c/
http://localhost/
http://127.0.0.1/
http://docker/
http://host.docker.internal/
http://host.docker.internal:8080/
http://kuutlab.com

RATKAISU: paina vaan "Enter" kun kysytään urlia. LOL.
token: SasFrsDntzJHXpESQjEM
new token: GR1348941NrpwYax3PvVyqzBpxjdy

sudo nano gitlab/gitlab-runner/config.toml
--------------------------------------

https://stackoverflow.com/questions/41559660/gitlab-ci-runner-not-able-to-expose-ports-of-nested-docker-containers

ERROR: Registering runner... failed                 runner=SasFrsDn status=couldn't execute POST against http://localhost/api/v4/runners: Post "http://localhost/api/v4/runners": dial tcp 127.0.0.1:80: connect: connection refused
PANIC: Failed to register the runner. 

Onko pakko käyttää EE:tä vai voiko käyttää CE?
Kokeile ohjeiden mukainen CE jos toimiskin.
RATKAISU: CE toimi.

Configure GitLab for your system by editing /etc/gitlab/gitlab.rb file
And restart this container to reload settings.
To do it use docker exec:

  docker exec -it gitlab editor /etc/gitlab/gitlab.rb
  docker restart gitlab

For a comprehensive list of configuration options please see the Omnibus GitLab readme
https://gitlab.com/gitlab-org/omnibus-gitlab/blob/master/README.md
/opt/gitlab/embedded/bin/runsvdir-start: line 24: ulimit: pending signals: cannot modify limit: Operation not permitted
/opt/gitlab/embedded/bin/runsvdir-start: line 37: /proc/sys/fs/file-max: Read-only file system
ffi-libarchive could not be loaded, libarchive is probably not installed on system, archive_file will not be available

If this container fails to start due to permission problems try to fix it by executing:

  docker exec -it gitlab update-permissions
  docker restart gitlab
  '

## Onnistunut rekisteröinti
Mitä tein:
1. Käytin gitlab-ce
2. gitlab-runner:alpine
3. komennot:
- docker exec -it gitlab-runner gitlab-runner register --url "http://gitlab-ce" --clone-url "http://gitlab-ce"
- "Press enter for url"
4. Runner pysyy rekisteröitynä ee-versiossakin :o

docker exec -it gitlab-runner gitlab-runner register --url "http://gitlab-ee" --clone-url "http://gitlab-ee"

## Add new remote origin
https://docs.github.com/en/get-started/getting-started-with-git/managing-remote-repositories?platform=linux#removing-a-remote-repository

git remote add local-gitlab http://localhost/gitlab-instance-9d36923c/COMP.SE.140.git
git remote set-url local-gitlab http://localhost/gitlab-instance-9d36923c/COMP.SE.140.git
git remote set-url local-gitlab http://localhost/gitlab-instance-9d36923c/COMP.SE.140.git
git remote -v

git log

## Multiple remote origins at once
Define a git remote which will point to multiple git remotes.
Say, we call it “all”: 
git remote add all REMOTE-URL-1.
git remote add all https://github.com/KaluTheKova/COMP.SE.140.git

Register 1st push URL: 
git remote set-url --add --push all https://github.com/KaluTheKova/COMP.SE.140.git
Register 2nd push URL: 
git remote set-url --add --push all http://localhost/gitlab-instance-9d36923c/COMP.SE.140.git

Push a branch to all the remotes with 
git push all BRANCH – replace BRANCH with a real branch name.


### Git push & pull all
git push all project
git pull all project
git fetch all project

You cannot pull from multiple remotes, but you can fetch updates from multiple remotes with git fetch --all.

## TO DO
1. CI
- unit tests to gateway
- Pipeline prakaa: go: cannot find main module, but found .git/config in
https://stackoverflow.com/questions/57182988/gitlab-ci-and-go-modules
- HOX: Hyödynnä Teemun https://gitlab.com/oliviasau/smart-home-designer/-/blob/master/docker-compose.dev.yml docker-compose in test - ideaa!

- deployment to somewhere
2. Unit tests for Gateway
- mock API responses from API Gateway
- GET /messages response
- PUT /state (payload “INIT”, “PAUSED”, “RUNNING”, “SHUTDOWN”)
- GET /state
- GET /run-log (as text/plain)
3. Dev
- Implement gateway API's
- Joko client tai irtofunccarit. Melkein sama asia. Irtofunkkarit ihan yhtä ok testata.
- Modify the ORIG service to send messages forever until pause paused or stopped.

## CURL COMMANDS
curl localhost:8083/state -X PUT -d "PAUSED"
curl localhost:8083/state -X PUT -d "RUNNING"
curl localhost:8083/state -X PUT -d "SHUTDOWN"
curl localhost:8083/state -X PUT -d "INIT"
curl localhost:8080
curl localhost:8083/messages
curl -v -X GET localhost:8083/messages
curl -v GET localhost:8083/state
curl -v GET localhost:8083/run-log

curl localhost:8083/state
curl localhost:8083/run-log
curl localhost:8080
curl localhost:8083/messages

## GITLAB RUNNER REGISTER
gitlab-runner register -n \
  --url https://Gitlab_Url/ \
  --registration-token TOKEN \
  --executor docker \
  --description "My Docker Runner" \
  --docker-image "docker" \
  --docker-privileged \
  --docker-volumes "/certs/client"

docker exec -it gitlab-runner gitlab-runner register --url "http://gitlab-ce" --clone-url "http://gitlab-ce" --docker-privileged

docker exec -it gitlab-runner gitlab-runner verify

docker exec -it gitlab-runner gitlab-runner register --url "http://gitlab-ce" --clone-url "docker:8081" --docker-privileged

new token: GR1348941NrpwYax3PvVyqzBpxjdy

docker logs gitlab-runner
So, after these steps, you can find the toml file inside the container volume: /etc/gitlab-runner


----------
fatal: unable to access 'http://localhost/gitlab-instance-9d36923c/COMP.SE.140.git/': Failed to connect to localhost port 80 after 0 ms: Connection refused

https://stackoverflow.com/questions/63766919/build-step-in-pipeline-is-failing-with-connection-refused-error-while-running-gi

Gitlab admin -> settings -> network -> Outbound requests
----------
Kai se ny on vaan pakko mockata unit testeille
Paitsi että ongelma on ettei runnerit ollu yhistettynä networkiin ja siksi heittää erroria. ASD.

- Kokeile uudestaan integraatiota nyt kun network_mode toimii

compse140-httpserv-1  | ls: go run /httpserv.go: No such file or directory
compse140-httpserv-1 exited with code 1

sudo apt-get install tree, katso kansiorakenne

#14 [compse140-gateway 4/6] RUN go mod download
#14 DONE 71.7s
#17 [compse140-gateway 5/6] COPY . /app
#17 DONE 0.1s
#18 [compse140-gateway 6/6] RUN go build -o /app/gateway
#18 DONE 3.4s
#16 [compse140-gateway] exporting to image
#16 exporting layers
#16 exporting layers 0.8s done
#16 writing image sha256:6dd1922242479d915b244bd6716ee2713b711cb47d96393daa564d2359ee89ee done
#16 naming to docker.io/library/compse140-gateway done
#16 DONE 0.9s
Network compse140_default  Creating
Network compse140_default  Created
Volume "compse140_message-storage"  Creating
Volume "compse140_message-storage"  Created
Container compse140-httpserv-1  Creating
Container compse140-httpserv-1  Created
Container compse140-gateway-1  Creating
Container compse140-gateway-1  Created
Attaching to compse140-gateway-1, compse140-httpserv-1
compse140-httpserv-1  | Dockerfile
compse140-httpserv-1  | go.mod
compse140-httpserv-1  | go.sum
compse140-httpserv-1  | httpserv
compse140-httpserv-1  | httpserv.go
compse140-httpserv-1  | httpserv_test.go
compse140-httpserv-1  | messages.txt
compse140-httpserv-1 exited with code 0
Aborting on container exit...
Container compse140-gateway-1  Stopping
Container compse140-gateway-1  Stopped
Container compse140-httpserv-1  Stopping
Error response from daemon: failed to create shim task: OCI runtime create failed: runc create failed: unable to start container process: exec: "CGO_ENABLED=0 go test go /app": stat CGO_ENABLED=0 go test go /app: no such file or directory: unknown

Aaah, tiedostot ei kopioidu. Test - filet jää pois.

go test -c -o foo if I remember correctly will compile all your tests into the executable binary named foo.

See: https://golang.org/pkg/cmd/go/internal/test/
Kato saako binaryt testattua ja ajettua

-----------
TO DO:
- RabbitMQ: tee viesteistä kestäviä, eli eivät katoa kun ne consumataan. Koska haistakaa paska. DONE
- Tee getState testi
- Implement getState (lue filestä state)
- implement staten tallentaminen gatewayn tekstitiedostoon

ONGELMA: curl localhost:8083/state -X PUT -d "RUNNING" ei sais lukita vaan pitää palauttaa vastaus. Ratkaise.

ONGELMA: IMED ei ny lue taaskaan. Kysy Sepolta prefetchin asetukset

ONGELMA: Pause ei toimi. Täytyy jotenkin lukita channelit kokonaan. Tai close ja reopen connection + channel

https://stackoverflow.com/questions/32864644/rabbitmq-multiple-consumers-on-a-queue-only-one-get-the-message

https://www.ribice.ba/golang-rabbitmq-client/

Mitä jos pausettais vain container executionin? :>>>>
https://docs.docker.com/engine/reference/commandline/pause/

Jos ei onnistu niin tee erilliset handlefuncit origiin
Toimii.

Sitten get State

JA TEE SHUTDOWN SEKÄ INIT
                                                     
#0 2.414 # github.com/KaluTheKova/COMP.SE.140/httpserv                                                                                                                        
#0 2.414 ./gateway.go:201:23: undefined: container.StopOptions
#0 2.414 ./gateway.go:235:23: undefined: container.StopOptions

Gatewat

Vielä time-to-live
2023-01-11 12:36:09 compse140-rabbitmq-1  | 2023-01-11 10:36:09.442346+00:00 [error] <0.1019.0> closing AMQP connection <0.1019.0> (172.22.0.7:52594 -> 172.22.0.2:5672):
2023-01-11 12:36:09 compse140-rabbitmq-1  | 2023-01-11 10:36:09.442346+00:00 [error] <0.1019.0> missed heartbeats from client, timeout: 10s
pitää laittaa pidemmäksi, että pause sallitaan.

Tai sitte vedät vaan toisen handlefucin kautta. Kokeile eka. 
Ongelma: Sama, channel jää lukkoon eikä vapaudu ikinä.

https://stackoverflow.com/questions/41991926/how-to-detect-dead-rabbitmq-connection
https://golang.hotexamples.com/examples/github.com.streadway.amqp/Config/Heartbeat/golang-config-heartbeat-method-examples.html

curl localhost:8083/state -X PUT -d "SHUTDOWN"
curl: (52) Empty reply from server <- Korjaa

privileged: true in /etc/gitlab-runner/config.toml

Testikomennot:
curl localhost:8083/state -X PUT -d "PAUSED"
curl localhost:8083/state -X PUT -d "RUNNING"
curl localhost:8083/state -X PUT -d "SHUTDOWN"
curl localhost:8083/state -X PUT -d "INIT"

curl localhost:8083/state
curl localhost:8083/run-log
curl localhost:8083/messages

Tarkista toimiiko httpserv ilman portteja