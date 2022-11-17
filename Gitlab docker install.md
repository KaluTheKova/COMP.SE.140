export GITLAB_HOME=/srv/gitlab

$GITLAB_HOME/data	    /var/opt/gitlab	    For storing application data.
$GITLAB_HOME/logs	    /var/log/gitlab	    For storing logs.
$GITLAB_HOME/config	    /etc/gitlab	For     storing the GitLab configuration files.

# Using Docker Engine:
Once you’ve set up the GITLAB_HOME variable, you can run the image:

sudo docker run --detach \
  --hostname gitlab.example.com \
  --publish 443:443 --publish 80:80 --publish 22:22 \
  --name gitlab \
  --restart always \
  --volume $GITLAB_HOME/config:/etc/gitlab \
  --volume $GITLAB_HOME/logs:/var/log/gitlab \
  --volume $GITLAB_HOME/data:/var/opt/gitlab \
  --shm-size 256m \
  gitlab/gitlab-ee:latest

https://docs.gitlab.com/ee/install/docker.html

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