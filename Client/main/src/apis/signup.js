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
            const data = response.data
            console.log(data.status === 200)
            if(data.status === 200){
                resolve(true)
            }
            else{
                resolve(false)
            }
        })
        .catch((error) => {
            console.error(error);
            resolve(false);
        });
    });
    return promise;
}

export default signup;