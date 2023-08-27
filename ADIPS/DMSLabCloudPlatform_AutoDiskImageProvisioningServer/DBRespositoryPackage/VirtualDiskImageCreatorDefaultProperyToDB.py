from Utils import RedisDatabaseRepositoryData, readYAMLFileToDictType

import json

documentsPath="Documents/VirtualDiskImageCreatorDefaultProperty.yaml"
opts="r"
r = RedisDatabaseRepositoryData()

def readSoftwareInformationToDB():
    result = readYAMLFileToDictType(documentsPath,opts)
    for key, value in result.items():
        r.insert_hash("VirtualDiskImageCreatorDefaultProperty", key, json.dumps(value))
    
#readSoftwareInformationToDB()
#print(r.get_Hvalue("VirtualDiskImageCreatorDefaultProperty", "MAXPACKAGES"))