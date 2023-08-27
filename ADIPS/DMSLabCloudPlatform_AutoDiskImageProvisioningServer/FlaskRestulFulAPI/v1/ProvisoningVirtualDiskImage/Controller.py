#!/usr/bin/python
#-*- coding: utf-8 -*-

from . import provinVDIapp


@provinVDIapp.route('/helloworld', methods=['GET'])
def test():
    return "helloworld v1"