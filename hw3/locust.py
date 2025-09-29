from locust import FastHttpUser, task, between

class MyUser(FastHttpUser):
    wait_time = between(1, 2)

    @task
    def get_albums(self):
        self.client.get("/albums")
        self.client.get("/albums")
        self.client.get("/albums")

    @task
    def post_album(self):
        self.client.post("/albums", json={
            "id": "99", "title": "Test Album", "artist": "Tester", "price": 9.99
        })
