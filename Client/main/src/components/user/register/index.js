
import React, { useState, useEffect } from 'react';
import './index.css';
import signup from '../../../apis/signup';
import { PriorityLevelList, RoleList, SpotList } from '../../../constants';

const RegisterModal = () => {
    const [name, setName] = useState('');
    const [userID, setUserID] = useState('');
    const [pw, setPw] = useState('');
    const [role, setRole] = useState('선택안함');
    const [spot, setSpot] = useState('선택안함');
    const priority = PriorityLevelList[0]; // Denied

    const handleSubmit = async (event) => {
        event.preventDefault();

        if (!name.trim() || !userID.trim() || !pw.trim()) {
            alert("모든 필드를 채워주세요.");
            return;
        }
        
        if (role === '선택안함' || spot === '선택안함') {
            alert("Role 및 Spot은 필수 선택 사항입니다.");
            return;
        }
        try {
            const response = await signup(name, userID, pw, role, spot, priority);
            console.log(response);
            if(response){
                alert("회원가입 성공");
                closeRegisterModal();
                window.location.href = "/";
            }
            else {
                alert("회원가입 실패");
            }
        }
        catch(error) {
            console.error(error);
        }
        
    };

    function closeRegisterModal() {
        document.getElementById("register-modal").style.display = "none";
    }

    return (
        <div id="register-modal" className="modal fade" tabIndex="-1" role="dialog">
            <div className="modal-dialog" role="document">
                <div className="modal-content">
                    <div className="modal-header">
                        <h5 className="modal-title">Register</h5>
                        <button type="button" className="close" onClick={closeRegisterModal}>
                            <span aria-hidden="true">&times;</span>
                        </button>
                    </div>
                    <div className="modal-body">
                        <div>
                            <div className="form-group">
                                <label htmlFor="name">Name:</label>
                                <input type="text" className="form-control" id="name" value={name} onChange={(e) => setName(e.target.value)} />
                            </div>
                            <div className="form-group">
                                <label htmlFor="userID">User ID:</label>
                                <input type="text" className="form-control" id="userID" value={userID} onChange={(e) => setUserID(e.target.value)} />
                            </div>
                            <div className="form-group">
                                <label htmlFor="pw">Password:</label>
                                <input type="password" className="form-control" id="pw" value={pw} onChange={(e) => setPw(e.target.value)} />
                            </div>
                            <div className="form-group">
                                <label htmlFor="role">Role:</label>
                                <select className="form-control" id="role" value={role} onChange={(e) => setRole(e.target.value)}>
                                    <option value="선택안함">선택안함</option>
                                    {RoleList.map((role, index) => (
                                        <option key={index} value={role}>{role}</option>
                                    ))}
                                </select>
                            </div>
                            <div className="form-group">
                                <label htmlFor="spot">Spot:</label>
                                <select className="form-control" id="spot" value={spot} onChange={(e) => setSpot(e.target.value)}>
                                    <option value="선택안함">선택안함</option>
                                    {SpotList.map((spot, index) => (
                                        <option key={index} value={spot}>{spot}</option>
                                    ))}
                                </select>
                            </div>
                        </div>
                    </div>
                    <div className="modal-footer">
                        <button type="button" className="btn btn-secondary" onClick={closeRegisterModal}>Close</button>
                        <button type="submit" className="btn btn-primary" onClick={handleSubmit}>Register</button>
                    </div>
                </div>
            </div>
        </div>
    );
}
export default RegisterModal;