# COMP.SE.140

## info
- Pushes to two remotes simulatenously, Github and local Gitlab with CI/CD pipeline.

1. Instructions for the teaching assistant
Implemented optional features
List of optional features implemented.
Instructions for examiner to test the system.
Pay attention to optional features.

NOTE: Unit tests are not located in "tests" - folder, because Go requires tests to reside in the same folder as the file that is being tested.

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
Main learnings include a yearning for Jenkins automation server or any pipeline automation server where I could replay a pipeline and adjust the gitlab-ci.yml without having to always recommit.

I also learned that one can push into multiple repositories simultaneously. Previously, I had never needed this feature, but now it proved to be interesting to implement and use. So here I configured the project so that it pushes simultaneously to my Github and to my locally hosted Gitlab that ran the pipeline.

My worst difficulties were actually with RabbitMQ's Golang implementation. The library had some under-the-hood implementations that caused a lot of lost time and headache, such as somewhat unsual use of channels without documenting them. 

Instead of using HTTP request from Gateway to and resepond on that, I wanted to try out using Docker CLI via Gateway's code. It was interesting to implement INIT, PAUSED, RUNNING and SHUTDOWN via that way. Using code to run container functions is also pretty DevOps thinking. 

Of course, in a real world production application this might cause problems if the CLI library is reliant on Docker, since the production application would have to always run on Docker instead of being portable to another container engine. In a real world application, I probably would have used the HTTP method. In this case however, we can assume that Docker engine is used to run the application.

And naturally, I would have implemented errorhandling and basically everything more robustly.

### Amount effort (hours) used
Around 50 I believe.
