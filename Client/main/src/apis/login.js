import instance from "./instance";

const Login = (username, password) => {
    const datas = JSON.stringify({ 
        
        "UserID" : username,
        "PW" : password,
    })
    const promise = new Promise((resolve, reject) => {
        instance.post("/auth/login", {
            dest : "/user/login",
            method : "POST",
            data : datas
        })
        .then((response) => {
            if(response.status === 200){
                localStorage.setItem("accessToken", response.data.accessToken);
                localStorage.setItem("refreshToken", response.data.refreshToken);
                resolve(true); 
            }
            else{
                resolve(false)
            }
        })
        .catch((error) => {
            console.error(error);
            reject(error);
        });
    });
    return promise;
};

export default Login;