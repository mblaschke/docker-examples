.PHONY: hello-world hello-world-multistage http-server php go-func

hello-world:
	bash run.sh hello-world

hello-world-multistage:
	bash run.sh hello-world-multistage

http-server:
	bash run.sh http-server

php:
	bash run.sh php

go-func:
	bash run.sh go-func
