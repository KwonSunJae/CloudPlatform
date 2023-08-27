from FlaskRestulFulAPI import SystemSoftwareInformationOperator, SystemSoftwareInformationDTOObject


class SystemSoftwareInformationService:
    def __init__(self):
        self.systemsoftwareinfooperator = SystemSoftwareInformationOperator()
    
    def get_all_systeminformation(self):
        systemsoftwareinformationlist =self.systemsoftwareinfooperator.queryAllSystemsoftwareInformation()
        all_list = []
        for value in systemsoftwareinformationlist:
            obj = SystemSoftwareInformationDTOObject(value)
            all_list.append(obj.to_dict())
        return all_list