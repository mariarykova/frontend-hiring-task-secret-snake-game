
REPO := frontend-hiring-task-secret-snake-game

.PHONY: build clean

build: ${REPO}

${REPO}: $(shell find src-backend)
	@cd src-backend && go build -o ../${REPO}

clean:
	@rm -f ${REPO}
