from SystemSoftwarePackage import SystemSoftwareInformation



class SystemSoftwareInformationDTOObject(SystemSoftwareInformation):
    def __init__(self, systemSoftwareInformationModelObject):
        super(SystemSoftwareInformationDTOObject, self).__init__(
            systemSoftwareInformationModelObject.sysswid,
            systemSoftwareInformationModelObject.distribution,
            systemSoftwareInformationModelObject.version,
            systemSoftwareInformationModelObject.codename,
            systemSoftwareInformationModelObject.architecture,
            systemSoftwareInformationModelObject.description
        )
         
    def to_dict(self):
        return {
            "id":self.sysSWId,
            "distribution":self.distribution,
            "version":self.version,
            "codename":self.codename,
            "architecture":self.architecture,
            "description":self.description            
        }

    
    
