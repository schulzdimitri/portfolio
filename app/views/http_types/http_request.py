class HttpRequest:
    def __init__(
        self,
        headers: dict = None,
        body: dict = None,
        path_params: dict = None
    ) -> None:
        self.headers = headers
        self.body = body
        self.path_params = path_params
