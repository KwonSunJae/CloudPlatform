import instance from "./instance";
/**
 * 
 * @param {*} name 
 * @param {*} userID 
 * @param {*} pw 
 * @param {*} role 
 * @param {*} spot 
 * @param {*} priority 
 */
const signup = (name, userID, pw, role,spot,priority) => {
    const datas = JSON.stringify({ 
        "Name" : name,
        "UserID" : userID,
        "PW" : pw,
        "Role" : role,
        "Spot" : spot,
        "Priority" : priority
    })
    const promise = new Promise((resolve, reject) => {
        instance.post("/auth/signup", {
            dest : "/user",
            method : "POST",
            data : datas
        })
        .then((response) => {
            console.log(response);
            if(response.status === 200){
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
            if(error.response.status === 409){
                alert("이미 존재하는 아이디입니다.");
            }
            console.error(error);
            resolve(false);
        });
    });
    return promise;
}

export default signup;