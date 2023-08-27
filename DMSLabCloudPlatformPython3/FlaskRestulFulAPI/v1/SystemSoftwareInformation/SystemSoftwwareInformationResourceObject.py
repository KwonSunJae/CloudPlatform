from FlaskRestulFulAPI import SystemSoftwareInformationService

from flask_restful import Resource

sys_sw_info_service=SystemSoftwareInformationService()

class SystemSoftwareInformationDTOResourceObject(Resource):
    def get(self):
        return sys_sw_info_service.get_all_systeminformation()