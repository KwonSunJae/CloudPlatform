import instance from "./instance";

const Login = (userID, password) => {
    const datas = JSON.stringify({ 
        
        "UserID" : userID,
        "PW" : password,
    })
    const promise = new Promise((resolve, reject) => {
        instance.post("/auth/login", {
            dest : "/user/login",
            method : "POST",
            data : datas
        }, { withCredentials: true })
        .then((response) => {
            if(response.status === 200){
                localStorage.setItem("accessToken", response.data.accessToken);
                resolve(true); 
            }
            else {
                resolve(false);
            }

        })
        .catch((error) => {
            if(error.response.status === 401){
                alert("아이디 또는 비밀번호가 잘못되었습니다.");
            }
            else{
                alert("서버 오류입니다. 잠시 후 다시 시도해주세요.");
            }
            reject(error);
        });
    });
    return promise;
};

export default Login;