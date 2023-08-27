from FlaskRestulFulAPI import SystemSoftwareInformationdDBObject




obj = SystemSoftwareInformationdDBObject()

class SystemSoftwareInformationModelObject(obj.systemSoftwareInformationdb.Model):
    __tablename__='systemsoftwareinformationtable'
    
    sysswid = obj.systemSoftwareInformationdb.Column(obj.systemSoftwareInformationdb.Integer, primary_key=True, autoincrement=True)
    distribution = obj.systemSoftwareInformationdb.Column(obj.systemSoftwareInformationdb.String)
    version = obj.systemSoftwareInformationdb.Column(obj.systemSoftwareInformationdb.String)
    codename = obj.systemSoftwareInformationdb.Column(obj.systemSoftwareInformationdb.String)
    architecture = obj.systemSoftwareInformationdb.Column(obj.systemSoftwareInformationdb.String)
    description = obj.systemSoftwareInformationdb.Column(obj.systemSoftwareInformationdb.String)