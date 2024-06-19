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
    const [pendingDeploymentList, setPendingDeploymentList] = useState([]);
    const [checkedDeployments, setCheckedDeployments] = useState({});
    const [showDeleteModal, setShowDeleteModal] = useState(false);
    const [deploymentToDelete, setDeploymentToDelete] = useState(null);

    useEffect(() => {
        instance.post("/transaction", {
            "dest": "/deployment",
            "method": "GET",
            "data": ""
        }, {
            headers: { "Access-Token": `${localStorage.getItem("accessToken")}` }
        })
            .then((response) => {
                const datas = response.data.data;
                setPendingDeploymentList(datas);
                console.log(datas);

            }).catch((error) => {
                console.error(error);
            });
    }, []);

    function approveSelectedDeployments() {
        Object.keys(checkedDeployments).forEach(Id => {
            if (checkedDeployments[Id]) {
                console.log("Approving Deployment:", Id);
                approveDeployment(Id, pendingDeploymentList.find(deployment => deployment.Id === Id).UUID);
            }
        });
    }

    function approveDeployment(Id, approveUserUUID) {
        instance.post("/transaction", {
            "dest": "/approve/deployment/" + Id,
            "method": "POST",
            "data": ""
        }).then(() => {
            alert("Deployment(s) approved successfully!");
            window.location.reload();
        }).catch((error) => {
            console.error("Error approving Deployment:", error);
        });
    }

    const handleCheckDeployment = (e, Id) => {
        setCheckedDeployments(prev => ({ ...prev, [Id]: e.target.checked }));
    }

    const handleDeleteDeployment = (Id) => {
        instance.post("/transaction", {
            "dest": "/deployment/" + Id,
            "method": "DELETE",
            "data": ""
        }).then(() => {
            alert("Deployment deleted successfully!");
            window.location.reload();

        }).catch((error) => {
            console.error("Error deleting Deployment:", error);
        });
        setShowDeleteModal(false);
    }

    return (
        <Container className="mt-5">
            <Table striped bordered hover>
                <thead>
                    <tr>
                        <th>선택</th><th>고유ID</th><th>Name</th><th>Labels</th><th>Replicas</th>
                        <th>Selector</th><th>Template Labels</th><th>Container Name</th>
                        <th>Container Image</th><th>Status</th><th>UUID</th><th>Container Port</th>
                        <th>액션</th>
                    </tr>
                </thead>
                <tbody>
                    {pendingDeploymentList.map((deployment) => (
                        <tr key={deployment.Id}>
                            <td>
                                <Form.Check
                                    type="checkbox"
                                    onChange={(e) => handleCheckDeployment(e, deployment.Id)}
                                    checked={checkedDeployments[deployment.Id] || false}
                                />
                            </td>
                            <td>{deployment.Id}</td>
                            <td>{deployment.MetadataName}</td>
                            <td>{deployment.MetadataLabelsApp}</td>
                            <td>{deployment.SpecReplicas}</td>
                            <td>{deployment.SpecSelectorMatchlabelsApp}</td>
                            <td>{deployment.SpecTemplateMetadataLabelsApp}</td>
                            <td>{deployment.SpecTemplateSpecContainersName}</td>
                            <td>{deployment.SpecTemplateSpecContainersImage}</td>
                            <td>{deployment.Status}</td>
                            <td>{deployment.UUID}</td>
                            <td>{deployment.SpecTemplateSpecContainersPortsContainerport}</td>
                            <td>
                                {deployment.Status === 'Pending' && (
                                    <Button variant="secondary" onClick={() => approveDeployment(deployment.Id, deployment.UUID)}>승인</Button>
                                )}
                                <Button variant="secondary" onClick={() => { setDeploymentToDelete(deployment.Id); setShowDeleteModal(true); }}>삭제</Button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </Table>
            <Button variant="primary" className="mb-3" onClick={approveSelectedDeployments}>선택한 Deployment 승인</Button>

            <Modal show={showDeleteModal} onHide={() => setShowDeleteModal(false)}>
                <Modal.Header closeButton>
                    <Modal.Title>Deployment 삭제</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    정말로 이 Deployment를 삭제하시겠습니까?
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="secondary" onClick={() => setShowDeleteModal(false)}>
                        취소
                    </Button>
                    <Button variant="danger" onClick={() => handleDeleteDeployment(deploymentToDelete)}>
                        삭제
                    </Button>
                </Modal.Footer>
            </Modal>
        </Container>
    );
};
