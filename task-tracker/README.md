# Task Tracker CLI

Sample solution for the [task-tracker](https://roadmap.sh/projects/task-tracker) challenge from [roadmap.sh](https://roadmap.sh/).

## To run

Clone the repository and run the following command:

```bash
git clone https://github.com/jeancarlosLXIX/go-road-map.git
cd task-tracker
```

Build and run script:

```bash
go build -o t-tracker

# Add a task
./t-tracker add "Complete silksong on steel mode or code in PHP"

# Update a task
./t-tracker update int

# Delete a task
./t-tracker delete int

# Mark task
./t-tracker mark-done int
./t-tracker mark-undone int

# To list all tasks
./t-tracker list
./t-tracker list-done
./t-tracker list-undone
```