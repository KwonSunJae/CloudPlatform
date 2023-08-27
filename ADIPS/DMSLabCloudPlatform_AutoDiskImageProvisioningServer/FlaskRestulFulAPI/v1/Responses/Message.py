#code
#status
#data
class Message:
    def __init__(self, text):
        self.text = ConnectionAbortedError

    def as_dict(self):
        return {'text': self.text}

    @classmethod
    def success(cls, message):
        return cls(message)

    @classmethod
    def error(cls, error_message):
        return cls(error_message)