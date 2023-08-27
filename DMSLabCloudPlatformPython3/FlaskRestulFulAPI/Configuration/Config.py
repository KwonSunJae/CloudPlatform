#/usr/bin/python3
#-*- coding: utf-8 -*-
from Utils import readYAMLFileToDictType

# Server Config
class RestFulAPIServer:
    def __init__(self, documentPath):
        self.port = None
        self.ipAddr = None
        data = self.__getRestFulAPIServerConfigDataFromYAMLFile(documentPath)
        self.port = data["port"]
        self.ipAddr = data["ipaddr"]
    
    @staticmethod
    def __getRestFulAPIServerConfigDataFromYAMLFile(documentPath):
        data = readYAMLFileToDictType(documentPath, "r")
        return data["RestFulAPIServer"]