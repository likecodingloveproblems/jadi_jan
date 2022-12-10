from locust import HttpUser, task, between
import requests


def on_start():
    print("Create comments")
    for _ in range(20):
        requests.post(
            "http://localhost:1323/comments/",
            json={"full_name": "full name", "description": "some description"},
        )


on_start()


class QuickstartUser(HttpUser):
    wait_time = between(1, 5)

    @task(5)
    def get_counter(self):
        self.client.get("/counter/")

    @task(3)
    def incr_counter(self):
        self.client.post("/counter/")

    @task(2)
    def get_comments(self):
        self.client.get("/comments/")
