from Utils import readYAMLFileToDictType
import redis
import datetime

EXTIME = datetime.datetime(2024,3,1,15,19,10)
EXTIME.strftime('%Y-%m-%d %H:%M:%S %f')

class RedisDatabaseRepositoryData:
    
    def __init__(self):
        database=RedisDatabaseRepositoryData.__getRedisConnectionInfo()
        self.pool = redis.ConnectionPool(host=database["ipaddr"],port=database["port"],db=database["dbindex"], decode_responses=True)
        self.conn=redis.Redis(connection_pool=self.pool)
    
    @staticmethod
    def __getRedisConnectionInfo():
        result=readYAMLFileToDictType("/Documents/Databaseconfiguration.yaml","r")
        database_list=result["database"]
        for database in database_list:
            if (database["type"] == "redis"):
                return database
    
    # Set String      
    def set_str(self, key, value, time=0):
        if time == 0:
            self.conn.set(key, value)
        else:
            self.conn.set(key, time, value)

    # Get String
    def get_str(self, key):        
        value = self.conn.get(key)
        if value:
            value = str(value, encoding='utf8')
        else:
            return None
        return value  
    
    # Delete String
    def del_str(self, key):
        return self.conn.delete(key)
    
    # Collection Insert
    def insert_set(self, key, value):
        self.conn.sadd(key, value)
    
    # Get_set
    def find_set(self, key):
        value = self.conn.smembers(key)
        if value:
            list = []
            for vl in value:
                list.append(str(vl, encoding='utf-8'))
            return list
        else:
            return None
    
    # Insert Hash
    def insert_hash(self, key, params, value):
        self.conn.hset(key, params, value)
    
    # Get single hash
    def get_Hvalue(self, key, params):
        return self.conn.hget(key, params)
    
    # Get all values
    def get_all_Hvalues(self, key):
        return self.conn.hgetall(key)
    
    # Delete hash
    def del_hash(self, key, params):
        self.conn.hdel(key, params)
        
    # Insert list
    def list_push(self, key, value):
        self.conn.lpush(key, value)
        
    # Get list
    def list_pop(self, key):
        self.conn.lrange(key,0,-1)
    
    # Length of list 
    def list_len(self, key):
        if self.conn.exists(key):
            return self.conn.llen(key)
        else:
            return 0
    # Key exists
    def exists_key(self, key):
        return self.conn.exists(key)


