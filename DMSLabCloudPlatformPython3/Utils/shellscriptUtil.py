import subprocess
from BaseClass import Message

timeout=20
def execShellScript(cmd_string):
    print("Execution Command is " + cmd_string)
    try:
        out_bytes = subprocess.check_output(cmd_string, stderr=subprocess.STDOUT, timeout=timeout, shell=True)
        res_code = 0
        value = out_bytes.decode('utf-8')
        message = Message(res_code,True, value, "")
    except subprocess.CalledProcessError as e:
        out_bytes = e.output
        msg = "[ERROR]CallError :" + out_bytes.decode('utf-8')
        res_code = e.returncode
        message = Message(res_code,False, None, msg)
    except subprocess.TimeoutExpired as e:
        res_code = 100
        msg = "[ERROR]Timeout : " + str(e)
        message = Message(res_code,False, None, msg)
    except Exception as e:
        res_code = 200
        msg = "[ERROR]Unknown Error : " + str(e)
        message = Message(res_code,False, None, msg)
    return message
