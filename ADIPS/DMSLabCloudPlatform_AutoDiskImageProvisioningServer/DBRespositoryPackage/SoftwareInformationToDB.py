from Utils import RedisDatabaseRepositoryData, readYAMLFileToDictType
import json

documentsPath="Documents/SoftwareInfo.yaml"
opts="r"
r = RedisDatabaseRepositoryData()

def readSoftwareInformationToDB(docPath):
    result = readYAMLFileToDictType(docPath,opts)
    for key, value in result.items():
        for val in value:
            for valKey in val.keys():               
                if (valKey == 'id'):
                    r.insert_hash(key, val[valKey], json.dumps(val))

def readConfigurationFilesToDB(docPath):
    result = readYAMLFileToDictType(docPath,opts)
    for key, value in result.items():
        r.insert_hash("ConfigurationOfPlatform", key, json.dumps(value))


def readVirtualDiskCreatorDefaultPropertyToDB(docPath):
    result = readYAMLFileToDictType(docPath,opts)
    for key, value in result.items():
        r.insert_hash("VirtualDiskImageCreatorDefaultProperty", key, json.dumps(value))

def initFilesToDB():
    softwareInfoDocPath="Documents/SoftwareInfo.yaml"
    virtualDiskImageCreatorDefaultPropertyDocPath="Documents/VirtualDiskImageCreatorDefaultProperty.yaml"
    configurationPlaformDocPath="Documents/ConfigurationOfPlatform.yaml"
    readSoftwareInformationToDB(softwareInfoDocPath)
    readConfigurationFilesToDB(configurationPlaformDocPath)
    readVirtualDiskCreatorDefaultPropertyToDB(virtualDiskImageCreatorDefaultPropertyDocPath)