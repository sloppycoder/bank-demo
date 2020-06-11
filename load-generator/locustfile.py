import grpc
import time

from locust import User, TaskSet, between, events, task

import demo_bank_pb2
import demo_bank_pb2_grpc


def init_tracer():
    config = Config(
        config={
            'sampler': {
                'type': 'const',
                'param': 1,
            },
            'logging': True,
            'propagation': 'b3',  # zipkin headers
        },
        validate=True,
        service_name='load-generator')
    tracer = config.initialize_tracer()
    interceptor = open_tracing_client_interceptor(tracer, log_payloads=True)
    return tracer, interceptor


class DashboardTask(TaskSet):
    def on_start(self):
        self.channel = grpc.insecure_channel(self.user.host)
        self.stub = demo_bank_pb2_grpc.DashboardServiceStub(self.channel)

    @task
    def get_dashboard(self):
        start_time = time.time()
        try:
            test_id = self.get_test_id()
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

    def get_test_id(self):
        return 'user001'


class MobileAppUser(User):
    tasks = [DashboardTask]
    wait_time = between(1, 2)

    def __init__(self, *args, **kwargs):
        super(MobileAppUser, self).__init__(*args, **kwargs)


if __name__ == '__main__':
    print('''
This script does not work as a standalone program, please use locust utility instead.

    locust --host="dashboard:50051" --headless -u 1
''')
