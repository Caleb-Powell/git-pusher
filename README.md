# Git Repository Pusher

This CLI application is meant to be ran by a scheduled runner such as chron jobs in linux. Given a dir path it will git push all of the repos in that directory.

## Installing the application

To build the bianary download this package and run the following command. Note you will need to have go installed on your system

```go build .```

You then have an executable named git-pusher in the root of the repo. Copy this to wherever you want it to run from.

To run the program you will need to run the exectable and provide the path. (The default path is the dir that it is running in)

```/the/path/to/your/executable -path="/the/path/to/your/git/dir"```

For Linux, you can set up a cron job to run this application. A good article on how to set this up can be found [here](https://www.ostechnix.com/a-beginners-guide-to-cron-jobs/ "https://www.ostechnix.com/a-beginners-guide-to-cron-jobs/").
