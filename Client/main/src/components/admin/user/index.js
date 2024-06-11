import React, { useEffect, useState } from "react";
import instance from "../../../apis/instance";
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Container from 'react-bootstrap/Container';
import ListGroup from 'react-bootstrap/ListGroup';
import Table from 'react-bootstrap/Table';
import Modal from 'react-bootstrap/Modal';
import './index.css';
import { PriorityLevelList, RoleList, SpotList } from '../../../constants';

export default function Admin() {
    const [pendingUserList, setPendingUserList] = useState([]);
    const [selectedPriority, setSelectedPriority] = useState({});
    const [selectedRole, setSelectedRole] = useState({});
    const [checkedUsers, setCheckedUsers] = useState({});
    const [showDeleteModal, setShowDeleteModal] = useState(false);
    const [userToDelete, setUserToDelete] = useState(null);

    useEffect(() => {
        instance.post("/transaction", {
            "dest": "/user",
            "method": "GET",
            "data": ""
        }).then((response) => {
            const datas = response.data.data;
            setPendingUserList(datas.sort((a, b) => PriorityLevelList.indexOf(a.Priority) - PriorityLevelList.indexOf(b.Priority)));
            const priorityData = {};
            const roleData = {};
            datas.forEach(user => {
                priorityData[user.Id] = user.Priority;
                roleData[user.Id] = user.Role;
            });
            
            setSelectedPriority(priorityData);
            setSelectedRole(roleData);
        }).catch((error) => {
            setPendingUserList([]);
            setSelectedPriority({});
            setSelectedRole({});
            console.error(error);
        });
    }, []);

    function approveSelectedUsers() {
        Object.keys(checkedUsers).forEach(Id => {
            if (checkedUsers[Id]) {
                console.log("Approving user:", Id);
                approveUser(Id);
            }
        });
    }

    function approveUser(Id) {
        console.log("Priority:", selectedPriority[Id], "Role:", selectedRole[Id]);
        const datas = JSON.stringify({
            "Priority": selectedPriority[Id],
            "Role": selectedRole[Id]
        });
        instance.post("/transaction", {
            "dest": "/user/approve/" + Id,
            "method": "POST",
            "data": datas
        }).then(() => {
            alert("User(s) approved successfully!");
            window.location.reload();
        }).catch((error) => {
            console.error("Error approving user:", error);
        });
    }

    const handleChangePriority = (e, Id) => {
        setSelectedPriority(prev => ({ ...prev, [Id]: e.target.value }));
    }

    const handleChangeRole = (e, Id) => {
        setSelectedRole(prev => ({ ...prev, [Id]: e.target.value }));
    }

    const handleCheckUser = (e, Id) => {
        setCheckedUsers(prev => ({ ...prev, [Id]: e.target.checked }));
    }

    const handleDeleteUser = (Id) => {
        instance.post("/transaction", {
            "dest": "/user/" + Id,
            "method": "DELETE",
            "data": ""
        }).then(() => {
            alert("User deleted successfully!");
            window.location.reload();

        }).catch((error) => {
            console.error("Error deleting user:", error);
        });
        setShowDeleteModal(false);
    }

    return (
        <Container className="mt-5">
            <Table striped bordered hover>
                <thead>
                    <tr>
                        <th>선택</th>
                        <th>고유ID</th>
                        <th>이름</th>
                        <th>소속</th>
                        <th>아이디</th>
                        <th>우선순위</th>
                        <th>신분</th>
                        <th>액션</th>
                    </tr>
                </thead>
                <tbody>
                    {pendingUserList.map((user) => (
                        <tr key={user.Id}>
                            <td>
                                <Form.Check 
                                    type="checkbox" 
                                    onChange={(e) => handleCheckUser(e, user.Id)} 
                                    checked={checkedUsers[user.Id] || false}
                                />
                            </td>
                            <td>{user.Id}</td>
                            <td>{user.Name}</td>
                            <td>{user.Spot}</td>
                            <td>{user.UserID}</td>
                            <td>{user.Priority}</td>
                            <td>{user.Role}</td>
                            <td>
                                <Form.Select value={selectedPriority[user.Id]} onChange={(e) => handleChangePriority(e, user.Id)}>
                                    {PriorityLevelList.map((level) => (
                                        <option key={level} value={level}>{level}</option>
                                    ))}
                                </Form.Select>
                            </td>
                            <td>
                                <Form.Select value={selectedRole[user.Id]} onChange={(e) => handleChangeRole(e, user.Id)}>
                                    {RoleList.map((role) => (
                                        <option key={role} value={role}>{role}</option>
                                    ))}
                                </Form.Select>
                            </td>
                            <td>
                            {user.Priority === 'Denied' && (
                                    <Button variant="secondary" onClick={() => approveUser(user.Id)}>승인</Button>
                                )}
                                <Button variant="secondary" onClick={() => { setUserToDelete(user.Id); setShowDeleteModal(true); }}>삭제</Button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </Table>
            <Button variant="primary" className="mb-3" onClick={approveSelectedUsers}>선택한 사용자 승인</Button>

            <Modal show={showDeleteModal} onHide={() => setShowDeleteModal(false)}>
                <Modal.Header closeButton>
                    <Modal.Title>사용자 삭제</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    정말로 이 사용자를 삭제하시겠습니까?
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="secondary" onClick={() => setShowDeleteModal(false)}>
                        취소
                    </Button>
                    <Button variant="danger" onClick={() => handleDeleteUser(userToDelete)}>
                        삭제
                    </Button>
                </Modal.Footer>
            </Modal>
        </Container>
    );
};
