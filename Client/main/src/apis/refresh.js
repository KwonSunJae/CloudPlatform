import instance from "./instance";

const Refresh = (refreshToken) => {
    const promise = new Promise((resolve, reject) => {
        instance.post("/auth/reissue", {
        }, {
            headers: {
                'Authorization': `${refreshToken}`
            },
            withCredentials: true
        })
        .then((response) => {
            if(response.status === 200){
                localStorage.setItem("accessToken", response.data.accessToken);
                resolve(true); 
            } else {
                resolve(false)
            }
        })
        .catch((error) => {
             //exception 처리
            if(error.response.status === 401){
                alert("토큰이 만료되었습니다. 로그인을 재시도해주세요.")
                window.location.href = "/";
            }
            console.error(error);
            reject(error);
        });
    });
    return promise;
};

export default Refresh;
