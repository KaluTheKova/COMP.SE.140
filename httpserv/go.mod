module github.com/KaluTheKova/COMP.SE.140/httpserv

go 1.19

replace (
	github.com/docker/cli => github.com/docker/cli v20.10.3-0.20221013132413-1d6c6e2367e2+incompatible // 22.06 master branch
	github.com/docker/docker => github.com/docker/docker v20.10.3-0.20221021173910-5aac513617f0+incompatible // 22.06 branch
	github.com/moby/buildkit => github.com/moby/buildkit v0.10.1-0.20220816171719-55ba9d14360a // same as buildx
)
