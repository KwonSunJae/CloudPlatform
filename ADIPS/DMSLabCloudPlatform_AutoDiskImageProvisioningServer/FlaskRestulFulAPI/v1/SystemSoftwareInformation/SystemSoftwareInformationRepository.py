from sqlalchemy.orm import sessionmaker
from sqlalchemy import create_engine
from sqlalchemy.orm.exc import NoResultFound
from sqlalchemy.dialects.postgresql import insert

from Utils import readYAMLFileToDictType
from FlaskRestulFulAPI import SystemSoftwareInformationModelObject

docPath="Documents/Databaseconfiguration.yaml"
opts="r"

class SystemSoftwareInformationOperator():

    def __init__(self):
        postgresConn = self.__getDBConnection(docPath, opts)
        self.db_url = postgresConn['type']+"://"+postgresConn["user"]+":"+postgresConn["password"]+"@"+postgresConn["host"]+":"+postgresConn["port"]+"/"+postgresConn["dbname"]
        self.engine = create_engine(self.db_url,echo=True, future=True)
        Session = sessionmaker(bind=self.engine)
        self.dbsession = Session()

    @staticmethod
    def __getDBConnection(documentsPath,opts):
        dbconn=readYAMLFileToDictType(documentsPath,opts)
        postgresConn=None
        for db in dbconn["database"]:
            print(db)
            if (db['type'] == 'postgresql'):
               postgresConn=db
               return postgresConn
    
    def queryAllSystemsoftwareInformation(self):
        all_data = self.dbsession.query(SystemSoftwareInformationModelObject).all()
        return all_data
    
    def querySystemsoftwareinformationdataById(self, id):           
        return self.dbsession.query(SystemSoftwareInformationModelObject).filter(SystemSoftwareInformationModelObject.sysswid==id)
    
    def insertSingleSystemSoftwareInformatioDAO(self, sysSWDAO):
        if (self.dataExisted(sysSWDAO) == True):
            self.dbsession.add(sysSWDAO)
            self.dbsession.commit()
            self.dbsession.close()
            print("Succeed to insert")
        else:
            print("Failed to insert")
    
    def insertMultipleSystemSoftwareInformatioDAO(self, SystemSoftwareInformationList):
        self.dbsession.add_all(SystemSoftwareInformationList)
        self.dbsession.commit()    
        self.dbsession.close()
    
    def dataDuplicateExisted(self, sysSWDAO):
        # Data to insert or merge
        try:
            #insert_systeminformationdao = insert(SystemSoftwareInformation.__tablename__).values(sysSWDAO)
            #on_duplicate_key = insert_systeminformationdao.du
            data2 = self.dbsession.query(SystemSoftwareInformationModelObject).filter(SystemSoftwareInformationModelObject.distribution==sysSWDAO.distribution,
                                                                           SystemSoftwareInformationModelObject.architecture==sysSWDAO.architecture,
                                                                           SystemSoftwareInformationModelObject.version==sysSWDAO.version)
            return False
        except NoResultFound:
            return True

