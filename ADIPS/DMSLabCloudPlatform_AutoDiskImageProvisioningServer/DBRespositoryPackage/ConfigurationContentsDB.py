from Utils import RedisDatabaseRepositoryData, readYAMLFileToDictType
import json

docPath="Documents/ConfigurationOfPlatform.yaml"
opts="r"

def readConfigurationFilesToDB():
    result = readYAMLFileToDictType(docPath,opts)
    r = RedisDatabaseRepositoryData()
    for key, value in result.items():
        r.insert_hash("ConfigurationOfPlatform", key, json.dumps(value))