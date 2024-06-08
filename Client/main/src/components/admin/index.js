import React, { useEffect, useState } from "react";
import instance from "../../apis/instance";
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Container from 'react-bootstrap/Container';
import ListGroup from 'react-bootstrap/ListGroup';
import './index.css';
export default function Admin() {
    const [pendingUserList, setPendingUserList] = useState([]);
    const [selectedPriority, setSelectedPriority] = useState({});
    const [selectedRole, setSelectedRole] = useState({});

    useEffect(() => {
        instance.post("/transaction", {
            "dest": "/user",
            "method": "GET",
            "data": ""
        }).then((response) => {
            const datas = response.data.data;
            setPendingUserList(datas);
        }).catch((error) => {
            console.error(error);
        });
    }, []);

    const PriorityLevelList = ["선택안함", "Denied", "Low", "Medium", "High", "Urgent"];
    const RoleList = ["선택안함", "Admin", "Student", "Master", "Researcher", "Others"];

    function approveAllUsers() {
        pendingUserList.forEach(user => {
            approveUser(user.Id);
        });
    }

    function approveUser(userID) {
        const datas = JSON.stringify({
            "Priority": selectedPriority[userID],
            "Role": selectedRole[userID]
        });
        instance.post("/transaction", {
            "dest": "/user/" + userID,
            "method": "PATCH",
            "data": datas
        }).then(() => {
            alert("User(s) approved successfully!");
        }).catch((error) => {
            console.error("Error approving user:", error);
        });
    }

    const handleChangePriority = (e, userId) => {
        setSelectedPriority(prev => ({ ...prev, [userId]: e.target.value }));
    }

    const handleChangeRole = (e, userId) => {
        setSelectedRole(prev => ({ ...prev, [userId]: e.target.value }));
    }

    return (
        <Container className="mt-5">
            <h1>Admin</h1>
            <Button variant="primary" className="mb-3" onClick={approveAllUsers}>모두 승인</Button>
            <ListGroup>
                {pendingUserList.map((user) => (
                    <ListGroup.Item key={user.UserID} className="d-flex justify-content-between align-items-center">
                        이름 : {user.Name} , 소속 : {user.Spot}
                        <div>
                            <Form.Label className="me-2">우선순위 :</Form.Label>
                            <Form.Select className="me-2" style={{width: "auto", display: "inline-block"}} value={selectedPriority[user.Id] || PriorityLevelList[0]} onChange={(e) => handleChangePriority(e, user.Id)}>
                                {PriorityLevelList.map((level) => (
                                    <option key={level} value={level}>{level}</option>
                                ))}
                            </Form.Select>
                            <Form.Label className="me-2">신분 :</Form.Label>
                            <Form.Select className="me-2" style={{width: "auto", display: "inline-block"}} value={selectedRole[user.Id] || RoleList[0]} onChange={(e) => handleChangeRole(e, user.Id)}>
                                {RoleList.map((role) => (
                                    <option key={role} value={role}>{role}</option>
                                ))}
                            </Form.Select>
                            <Button variant="success" size="sm" onClick={() => approveUser(user.UserID)}>승인</Button>
                        </div>
                    </ListGroup.Item>
                ))}
            </ListGroup>
        </Container>
    );
}
