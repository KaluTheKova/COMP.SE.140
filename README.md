# COMP.SE.140

# Information about the host:
Linux DESKTOP-GBFKV5A 5.10.16.3-microsoft-standard-WSL2 #1 SMP Fri Apr 2 22:23:49 UTC 2021 x86_64 x86_64 x86_64 GNU/Linux
Docker version 20.10.20, build 9fdeb9c
Docker Compose version v2.12.1

# Perceived (in your mind) benefits of the topic-based communication compared to request-response (HTTP):
Asynchronous communication allows scaling of workers to answer the needs of a increasing message load. Using topics allows us distribute our messages and read them however we want. And the with the consumer acknowledgement, we can make sure that even if a consumer crashes, we can spool up a new consumer and our message can be recovered.

# Main learnings
This was a good exercise. The best thing about this was that I got to refresh my memory on many parts of programming messaging systems and container communication. I had totally forgotten about healthchecking and dependency starting a cluster, since things like that have already been solved at work.
I left debugging messages in, since according to the instructions we'll be using this later. However, those messages shouldn't be visible to you if you check the system according to instructions.