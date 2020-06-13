from datetime import datetime
import grpc
import random
import time

from locust import User, TaskSet, between, events, task

import demo_bank_pb2
import demo_bank_pb2_grpc


class DataPool:
    def __init__(self, file):
        with open(file, 'r') as f:
            self.id_pool = [line.rstrip() for line in f.readlines()]
        self.pool_size = len(self.id_pool)

        random.seed(int(datetime.today().timestamp() * 100))

    def next_id(self):
        i = random.randrange(0, self.pool_size - 1, 1)
        return self.id_pool[i]


pool = DataPool('ids.txt')


class DashboardTask(TaskSet):
    def on_start(self):
        self.channel = grpc.insecure_channel(self.user.host)
        self.stub = demo_bank_pb2_grpc.DashboardServiceStub(self.channel)

    @task
    def get_dashboard(self):
        start_time = time.time()
        try:
            test_id = pool.next_id()
            req = demo_bank_pb2.GetDashboardRequest(login_name=test_id)
            self.stub.GetDashboard(req)
        except grpc.RpcError as e:
            total_time = int((time.time() - start_time) * 1000)
            events.request_failure.fire(
                request_type='grpc',
                name='get_dashboard',
                response_time=total_time,
                response_length=0,
                exception=e
            )
        else:
            total_time = int((time.time() - start_time) * 1000)
            events.request_success.fire(
                request_type='grpc',
                name='get_dashboard',
                response_time=total_time,
                response_length=0,
            )


class MobileAppUser(User):
    tasks = [DashboardTask]
    wait_time = between(1, 2)

    def __init__(self, *args, **kwargs):
        super(MobileAppUser, self).__init__(*args, **kwargs)
        DataPool('ids.txt')


if __name__ == '__main__':
    for i in range(100):
        print(pool.next_id())

    print('''
This script does not work as a standalone program, please use locust utility instead.

    locust --host="dashboard:50051" --headless -u 1
''')
