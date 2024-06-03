import instance from "./instance";

const login = (username, password) => {
    instance.post("/auth/login", {
        dest : "/user/login",
        method : "POST",
        data : { "UserID" : username, "Password" : password}
    })
    .then((response) => {
        localStorage.setItem("accessToken", response.data.accessToken);
        localStorage.setItem("refreshToken", response.data.refreshToken);

        return true;
    })
    .catch((error) => {
        console.error(error);
        return false;
    });
};

export default login;