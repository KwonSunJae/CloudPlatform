from Utils import readYAMLFileToDictType

docPath="Documents/Databaseconfiguration.yaml"
opts="r"

class PostgreSQLDBConfigData:
    def __init__(self):
        postgresConn = self.__getDBConnection(docPath, opts)
        self.db_url = postgresConn['type']+"+psycopg2"+"://"+postgresConn["user"]+":"+postgresConn["password"]+"@"+postgresConn["host"]+":"+postgresConn["port"]+"/"+postgresConn["dbname"]

    @staticmethod
    def __getDBConnection(documentsPath,opts):
        dbconn=readYAMLFileToDictType(documentsPath,opts)
        postgresConn=None
        for db in dbconn["database"]:
            if (db['type'] == 'postgresql'):
               postgresConn=db
               return postgresConn