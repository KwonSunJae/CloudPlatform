import React, { useEffect, useState } from "react";
import instance from "../../../apis/instance";
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Container from 'react-bootstrap/Container';
import ListGroup from 'react-bootstrap/ListGroup';
import Table from 'react-bootstrap/Table';
import Modal from 'react-bootstrap/Modal';
import './index.css';
import { ConatinerStatusList } from '../../../constants';

export default function Admin() {
    const [pendingContainerList, setPendingContainerList] = useState([]);
    const [checkedContainers, setCheckedContainers] = useState({});
    const [showDeleteModal, setShowDeleteModal] = useState(false);
    const [containerToDelete, setContainerToDelete] = useState(null);

    useEffect(() => {
        instance.post("/transaction", {
            "dest": "/service",
            "method": "GET",
            "data": ""
        }, {
            headers: { "Access-Token": `${localStorage.getItem("accessToken")}` }
        })
            .then((response) => {
                const datas = response.data.data;
                setPendingContainerList(datas);
                console.log(datas);

            }).catch((error) => {
                console.error(error);
            });
    }, []);

    function approveSelectedContainers() {
        Object.keys(checkedContainers).forEach(Id => {
            if (checkedContainers[Id]) {
                console.log("Approving Container:", Id);
                approveContainer(Id, pendingContainerList.find(container => container.Id === Id).UUID);
            }
        });
    }

    function approveContainer(Id, approveUserUUID) {
        instance.post("/transaction", {
            "dest": "/approve/service/" + Id,
            "method": "POST",
            "data": ""
        }).then(() => {
            alert("Container(s) approved successfully!");
            window.location.reload();
        }).catch((error) => {
            console.error("Error approving Container:", error);
        });
    }

    const handleCheckContainer = (e, Id) => {
        setCheckedContainers(prev => ({ ...prev, [Id]: e.target.checked }));
    }

    const handleDeleteContainer = (Id) => {
        instance.post("/transaction", {
            "dest": "/service/" + Id,
            "method": "DELETE",
            "data": ""
        }).then(() => {
            alert("Container deleted successfully!");
            window.location.reload();

        }).catch((error) => {
            console.error("Error deleting Container:", error);
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
                        <th>External IP</th>
                        <th>Flavor ID</th>
                        <th>Internal IP</th>
                        <th>Keypair</th>
                        <th>Selected OS</th>
                        <th>Security Group</th>
                        <th>Status</th>
                        <th>UUID</th>
                        <th>Unionmount Image</th>
                        <th>액션</th>
                    </tr>
                </thead>
                <tbody>
                    {pendingContainerList.map((container) => (
                        <tr key={container.Id}>
                            <td>
                                <Form.Check
                                    type="checkbox"
                                    onChange={(e) => handleCheckContainer(e, container.Id)}
                                    checked={checkedContainers[container.Id] || false}
                                />
                            </td>
                            <td>{container.Id}</td>
                            <td>{container.MetadataName}</td>
                            <td>{container.SpecExternalname}</td>
                            <td>{container.SpecSelectorType}</td>
                            <td>{container.SpecClusterIP}</td>
                            <td>{container.SpecSelectorApp}</td>
                            <td>{container.ApiVersion}</td>
                            <td>{container.SpecType}</td>
                            <td>{container.Status}</td>
                            <td>{container.UUID}</td>
                            <td>{container.Kind}</td>
                            <td>
                                {container.Status === 'Pending' && (
                                    <Button variant="secondary" onClick={() => approveContainer(container.Id, container.UUID)}>승인</Button>
                                )}
                                <Button variant="secondary" onClick={() => { setContainerToDelete(container.Id); setShowDeleteModal(true); }}>삭제</Button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </Table>
            <Button variant="primary" className="mb-3" onClick={approveSelectedContainers}>선택한 Container 승인</Button>

            <Modal show={showDeleteModal} onHide={() => setShowDeleteModal(false)}>
                <Modal.Header closeButton>
                    <Modal.Title>Container 삭제</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    정말로 이 Container를 삭제하시겠습니까?
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="secondary" onClick={() => setShowDeleteModal(false)}>
                        취소
                    </Button>
                    <Button variant="danger" onClick={() => handleDeleteContainer(containerToDelete)}>
                        삭제
                    </Button>
                </Modal.Footer>
            </Modal>
        </Container>
    );
};
