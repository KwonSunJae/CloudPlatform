import instance from "./instance";

const signup = (name, userID, pw, role,spot,priority) => {
    instance.post("/auth/signup", {
        dest : "/user",
        method : "POST",
        data : { 
            "Name" : name,
            "UserID" : userID,
            "Password" : pw,
            "Role" : role,
            "Spot" : spot,
            "Priority" : priority
        }
    })
    .then((response) => {
        if(response.status === 200){
            return response.data.data;
        }
        else{
            return response.data.error
        }
    })
    .catch((error) => {
        console.error(error);
        return false;
    });
}

export default signup;