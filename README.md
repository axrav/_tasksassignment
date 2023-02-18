# TASKS Assignment
This repository contains the tasks given by an org as an assignment


### PATHS
- Task_1 contains the first task 
- Task_2 contains the second task  


## TASK_1

### File Structure
- config (Loading environment variables)
- controllers (The API controllers for routes)
- db (The database package where the database is connected)
- router (The routes to apis)
- main.go (The endpoint to run )


### Instructions to Install TASK_1 on a linux machine
``` console
git clone https://github.com/axrav/_tasksassignment
cd _tasksassignment/Task_1
# either export or env or use a .env file refer to sample.env
go build . app
chmod +x app
./app


```

### Install using docker
```console
# use a .env file or export the env to docker later, make sure to edit values of env before build
sudo docker image build -t tasks .
# not using env
sudo docker run -p machine_port:container_port -e PORT=container_port -e DATABASE_URL="your_database_url" -d -it tasks  
# if using .env 
sudo docker run -p machine_port:container_port -d -it tasks


```

## Install with docker compose
```console
```

### Stack used




## TASK_2

### File Structure
- data.csv (Sample Csv file)
- helpers (functions used in completing the task)
- structs (the struct according to the sample csv file)
- main.go (the main file to run the project)

### Instructions to use TASK_2
The structs are made according to sample csv file, all what is needed to follow this

``` console_
git clone https://github.com/axrav/_tasksassignment
cd _tasksassignment/Task_2
go run .
```



The instructions above are given for a linux machine, for different OS, the instructions might differ

- These Tasks were completed by [axrav](https://github.com/axrav)

