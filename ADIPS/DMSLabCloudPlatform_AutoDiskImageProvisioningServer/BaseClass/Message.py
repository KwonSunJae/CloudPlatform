#!/usr/bin/python
#-*- coding: utf-8 -*-


class Message:
    def __init__(self, code, status, value, msg):
        self.code=code
        self.status=status
        self.value=value
        self.msg=msg
    
    #Code Getter Setter
    @property
    def code(self):
        return self.__code
    
    @code.setter
    def age(self, setVal):
        self.__code=setVal
        
    #Status Getter Setter
    @property
    def status(self):
        return self.__status
    
    @status.setter
    def status(self, setVal):
        self.__status=setVal
    
    #Value Getter Setter
    @property
    def value(self):
        return self.__value
    
    @value.setter
    def value(self, setVal):
        self.__value=setVal
      
    #msg Getter Setter
    @property
    def msg(self):
        return self.__msg
    
    @msg.setter
    def msg(self, setVal):
        self.__msg=setVal
        