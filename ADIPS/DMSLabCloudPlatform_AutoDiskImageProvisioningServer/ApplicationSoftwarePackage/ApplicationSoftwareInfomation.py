#!/usr/bin/python
#-*- coding: utf-8 -*-

class ApplicationSoftwareInfomation:
    def __init__(self, appSWName, provider, url, projectName, latestVer, dependenciesList):
        self.appSWId = None
        self.appSWName = appSWName
        self.provider = provider
        self.url = url
        self.projectName = projectName
        self.latestVer = latestVer
        self.dependenciesList = dependenciesList

