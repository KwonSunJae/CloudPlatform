class ResourceNotFoundError(Exception):
    def __init__(self, obj):
        self.message = f"{obj} not found"
        super().__init__(self.message)
