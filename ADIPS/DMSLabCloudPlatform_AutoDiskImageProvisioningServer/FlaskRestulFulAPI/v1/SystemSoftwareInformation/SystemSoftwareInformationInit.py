from flask import Flask
from FlaskRestulFulAPI import SystemSoftwareInformationdDBObject, PostgreSQLDBConfigData


def create_app():
    psqlData = PostgreSQLDBConfigData()
    objDB = SystemSoftwareInformationdDBObject()
    flask_app = Flask(__name__)
    flask_app.config['SQLALCHEMY_DATABASE_URI'] = psqlData.db_url
    flask_app.config['SQLALCHEMY_TRACK_MODIFICATIONS'] = False
    flask_app.app_context().push()
    objDB.systemSoftwareInformationdb.init_app(flask_app)
    return flask_app