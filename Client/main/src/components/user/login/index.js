import React, { useState } from 'react';
import './index.css';
import login from '../../../apis';


const LoginModal = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');

    const handleSubmit = (event) => {
        event.preventDefault();
        login(username, password)
            .then((response) => {
                if (response) {
                    console.log(response);
                    window.location.href = "/user/";
                }
            });
    };

    function closeLoginModal() {
        document.getElementById("login-modal").style.display = "none";
    }

    return (
        <div id="login-modal" className="modal">
            <div className="login-modal-content">
                <span className="login-modal-close" onClick={closeLoginModal}>&times;</span>
                <form className="login-modal-form" onSubmit={handleSubmit}>
                    <div className="login-form-group">
                        <label htmlFor="username">Username:</label>
                        <input type="text" id="username" value={username} onChange={(e) => setUsername(e.target.value)} />
                    </div>
                    <div className="login-form-group">
                        <label htmlFor="password">Password:</label>
                        <input type="password" id="password" value={password} onChange={(e) => setPassword(e.target.value)} />
                    </div>
                    <button type="submit" className="btn btn-primary">Login</button>
                </form>
            </div>
        </div>
    );
}
export default LoginModal;