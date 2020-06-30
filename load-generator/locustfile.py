from locust import User, task

# content is this file is dummy
# the task definitions are in slave.go

class Dummy(User):
    @task(20)
    def hello(self):
        pass
