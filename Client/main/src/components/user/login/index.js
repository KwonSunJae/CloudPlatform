import React, { useState } from 'react';
import './index.css';
import Login from '../../../apis/login';
import RegisterModal from '../register';
import SpinnerOurs from '../../sppinner';
const LoginModal = () => {
    const [username, setUsername] = useState('');
    const [password, setPassword] = useState('');
    const [isMemorized, setIsMemorized] = useState(false);
    const [loading, setLoading] = useState(false);
    const handleSubmit = async (event) => {
        event.preventDefault();
        if(await Login(username, password,isMemorized)) {
            alert("로그인 성공");
            setLoading(true);
            closeLoginModal();
            window.location.href = "/";
        }
        else {
            setLoading(false);
            alert("로그인 실패");
        }
    };

    function closeLoginModal() {
        document.getElementById("login-modal").style.display = "none";
    }

    function openRegisterModal() {
        closeLoginModal();
        document.getElementById("register-modal").style.display = "block";
    }

    if (loading) {
        return <SpinnerOurs />
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
                    <div className="login-form-group">
                        <label htmlFor="isMemorized">
                            <input type="checkbox" id="isMemorized" checked={isMemorized} onChange={(e) => setIsMemorized(e.target.checked)} />
                            자동 로그인
                        </label>
                    </div>
                    <div className="button-group">
                        <button type="submit" className="btn btn-primary">로그인</button>
                        <button type="button" className="btn btn-primary" onClick={openRegisterModal}>회원가입</button>
                    </div>
                </form>
            </div>
        </div>
    );
}
export default LoginModal;