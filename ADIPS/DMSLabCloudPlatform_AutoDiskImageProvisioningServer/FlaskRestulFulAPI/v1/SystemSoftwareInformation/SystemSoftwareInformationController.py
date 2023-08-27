from flask import Blueprint
from flask_restful import Api
from FlaskRestulFulAPI import create_app,RestFulAPIServer,SystemSoftwareInformationDTOResourceObject,ResourceNotFoundError

systemsoftwareinformationbp = Blueprint('systemsoftwareinformation',__name__)
api = Api(systemsoftwareinformationbp)


@api.errorhandler(ResourceNotFoundError)
def handle_resource_not_found_error(error):
    return {'error': error.message}, 404

api.add_resource(SystemSoftwareInformationDTOResourceObject, '/systemsoftwareinformationlist')
