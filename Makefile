SHELL=cmd

DSN="root@(localhost:3306)/todolist?parseTime=true&tls=false"

## clean: cleans all binaries and runs go clean
clean:
	@echo Cleaning...
	@echo y | DEL /S dist
	@go clean
	@echo Cleaned and deleted binaries

## Complete
build: build_front build_back

start: build_front build_back start_front start_back

stop: stop_front stop_back

restart: stop build start
## Front End

## build_front: builds the front end
build_front:
	@echo Building front end...
	@go build -o dist/frontend.exe ./src/frontend
	@echo Front end built!

## start_front: starts the front end
start_front: build_front
	@echo Starting the front end...
	@start /B .\dist\frontend.exe -dsn=${DSN}
	@echo Front end running!

## stop_front: stops the front end
stop_front:
	@echo Stopping the front end...
	@taskkill /IM frontend.exe /F
	@echo Stopped front end

restart_front: stop_front start_front

## Back End

## build_front: stops the back end
build_back:
	@echo Building back end...
	@go build -o dist/backend.exe ./src/backend
	@echo Back end built!

## start_front: starts the back end
start_back: build_front
	@echo Starting the back end...
	@start /B .\dist\backend.exe -dsn=${DSN}
	@echo Back end running!

## stop_front: stops the back end
stop_back:
	@echo Stopping the back end...
	@taskkill /IM backend.exe /F
	@echo Stopped back end

restart_front: stop_front start_front