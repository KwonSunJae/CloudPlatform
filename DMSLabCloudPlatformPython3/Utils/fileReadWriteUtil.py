import yaml
import json
import os


def readYAMLFileToDictType(filePath, opts):
    with open(os.getcwd()+"/"+filePath,opts) as  file:
       data = yaml.load(file, Loader=yaml.FullLoader)
    return data
 
def readJSONFileToDictType(filePath, opts):
    with open(os.getcwd()+"/"+filePath,opts) as  file:
       data = json.load(file)
    return data