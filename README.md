# COMP.SE.140

## info
- Pushes to two remotes simulatenously, Github and local Gitlab with CI/CD pipeline.

1. Instructions for the teaching assistant
Implemented optional features
List of optional features implemented.
Instructions for examiner to test the system.
Pay attention to optional features.
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

4. Reflections
### Main learnings and worst difficulties:
Main learnings include a yearning for Jenkins library or any pipeline automation server where I could replay a pipeline and adjust the gitlab-ci.yml without having to always recommit.

My worst difficulties were actually with RabbitMQ's Golang implementation. The library had some under-the-hood implementations that caused a lot of lost time and headache, such as somewhat unsual use of channels without documenting them. 

Instead of using HTTP request from Gateway to and resepond on that, I wanted to try out using Docker CLI via Gateway's code. It was interesting to implement INIT, PAUSED, RUNNING and SHUTDOWN via that way. Using code to run container functions is also pretty DevOps thinking. 

Of course, in a real world production application this might cause problems is the CLI library is realiant on Docker, since the production application would have to always run on Docker instead of being portable to another container engine. In a real world application, I probably would have used normal HTTP method. In this case however, we can assume that Docker engine is used to run the application

### Amount effort (hours) used
Around 50 I believe, but most of that time was simply used to debug libraries.
