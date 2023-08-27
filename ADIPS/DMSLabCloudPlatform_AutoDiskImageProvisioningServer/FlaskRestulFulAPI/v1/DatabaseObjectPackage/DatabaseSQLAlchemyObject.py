from flask_sqlalchemy import SQLAlchemy

class SystemSoftwareInformationdDBObject:
    def __init__(self):
        self.systemSoftwareInformationdb = SQLAlchemy()