#!/usr/bin/python
#-*- coding: utf-8 -*-

class SystemSoftwareInformation:
    def __init__(self, sysSWId, distribution, version,codename, architecture, description):
        self.sysSWId = sysSWId
        self.distribution = distribution
        self.version = version
        self.codename = codename
        self.architecture = architecture
        self.description = description