import time

class SnowflakeIdGenerator:
    def __init__(self, datacenter_id, worker_id):
        self.datacenter_id = datacenter_id
        self.worker_id = worker_id
        self.twepoch = 1288834974657
        self.sequence = 0
        self.worker_id_bits = 5
        self.datacenter_id_bits = 5
        self.max_worker_id = -1 ^ (-1 << self.worker_id_bits)
        self.max_datacenter_id = -1 ^ (-1 << self.datacenter_id_bits)
        self.sequence_bits = 12

        self.worker_id_shift = self.sequence_bits
        self.datacenter_id_shift = self.sequence_bits + self.worker_id_bits
        self.timestamp_left_shift = self.sequence_bits + self.worker_id_bits + self.datacenter_id_bits
        self.sequence_mask = -1 ^ (-1 << self.sequence_bits)
        self.last_timestamp = -1
        
    def get_timestamp(self):
        return int(time.time() * 1000)

    def generate_id(self):
        timestamp = self.get_timestamp()

        if timestamp < self.last_timestamp:
            raise Exception("Clock moved backwards. Refusing to generate id for %d milliseconds" % (self.last_timestamp - timestamp))

        if timestamp == self.last_timestamp:
            self.sequence = (self.sequence + 1) & self.sequence_mask
            if self.sequence == 0:
                timestamp = self.til_next_millis(self.last_timestamp)
        else:
            self.sequence = 0

        self.last_timestamp = timestamp

        return ((timestamp - self.twepoch) << self.timestamp_left_shift) | (self.datacenter_id << self.datacenter_id_shift) | (self.worker_id << self.worker_id_shift) | self.sequence

    def til_next_millis(self, last_timestamp):
        timestamp = self.get_timestamp()
        while timestamp <= last_timestamp:
            timestamp = self.get_timestamp()
        return timestamp