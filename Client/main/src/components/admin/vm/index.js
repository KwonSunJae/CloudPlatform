import React, { useEffect, useState } from "react";
import instance from "../../../apis/instance";
import Button from 'react-bootstrap/Button';
import Form from 'react-bootstrap/Form';
import Container from 'react-bootstrap/Container';
import ListGroup from 'react-bootstrap/ListGroup';
import Table from 'react-bootstrap/Table';
import Modal from 'react-bootstrap/Modal';
import './index.css';
import { VmStatusList } from '../../../constants';

export default function Admin() {
    const [pendingVmList, setPendingVmList] = useState([]);
    const [checkedVms, setCheckedVms] = useState({});
    const [showDeleteModal, setShowDeleteModal] = useState(false);
    const [VmToDelete, setVmToDelete] = useState(null);

    useEffect(() => {
        instance.post("/transaction", {
            "dest": "/vm",
            "method": "GET",
            "data": ""
        }, {
            headers: { "Access-Token": `${localStorage.getItem("accessToken")}` }
        })
            .then((response) => {
                const datas = response.data.data;
                setPendingVmList(datas);
                console.log(datas);

            }).catch((error) => {
                console.error(error);
            });
    }, []);

    function approveSelectedVms() {
        Object.keys(checkedVms).forEach(Id => {
            if (checkedVms[Id]) {
                console.log("Approving Vm:", Id);
                approveVm(Id,pendingVmList.find(vm => vm.Id === Id).UUID);
            }
        });
    }

    function approveVm(Id,approveUserUUID) {
        const datas = JSON.stringify({
            "approveUserUUID": approveUserUUID
        });
        instance.post("/transaction", {
            "dest": "/action/approve/" + Id,
            "method": "POST",
            "data": datas

        }).then(() => {
            alert("Vm(s) approved successfully!");
            window.location.reload();
        }).catch((error) => {
            console.error("Error approving Vm:", error);
        });
    }

    const handleCheckVm = (e, Id) => {
        setCheckedVms(prev => ({ ...prev, [Id]: e.target.checked }));
    }

    const handleDeleteVm = (Id) => {
        instance.post("/transaction", {
            "dest": "/vm/" + Id,
            "method": "DELETE",
            "data": ""
        }).then(() => {
            alert("Vm deleted successfully!");
            window.location.reload();

        }).catch((error) => {
            console.error("Error deleting Vm:", error);
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
                    {pendingVmList.map((vm) => (
                        <tr key={vm.Id}>
                            <td>
                                <Form.Check
                                    type="checkbox"
                                    onChange={(e) => handleCheckVm(e, vm.Id)}
                                    checked={checkedVms[vm.Id] || false}
                                />
                            </td>
                            <td>{vm.Id}</td>
                            <td>{vm.Name}</td>
                            <td>{vm.ExternalIP}</td>
                            <td>{vm.FlavorID}</td>
                            <td>{vm.InternalIP}</td>
                            <td>{vm.Keypair}</td>
                            <td>{vm.SelectedOS}</td>
                            <td>{vm.SelectedSecuritygroup}</td>
                            <td>{vm.Status}</td>
                            <td>{vm.UUID}</td>
                            <td>{vm.UnionmountImage}</td>
                            <td>
                                {vm.Status === 'Pending' && (

                                    <Button variant="secondary" onClick={() => approveVm(vm.Id,vm.UUID)}>승인</Button>

                                )}
                                <Button variant="secondary" onClick={() => { setVmToDelete(vm.Id); setShowDeleteModal(true); }}>삭제</Button>
                            </td>
                        </tr>
                    ))}
                </tbody>
            </Table>
            <Button variant="primary" className="mb-3" onClick={approveSelectedVms}>선택한 VM 승인</Button>

            <Modal show={showDeleteModal} onHide={() => setShowDeleteModal(false)}>
                <Modal.Header closeButton>
                    <Modal.Title>VM 삭제</Modal.Title>
                </Modal.Header>
                <Modal.Body>
                    정말로 이 VM을 삭제하시겠습니까?
                </Modal.Body>
                <Modal.Footer>
                    <Button variant="secondary" onClick={() => setShowDeleteModal(false)}>
                        취소
                    </Button>
                    <Button variant="danger" onClick={() => handleDeleteVm(VmToDelete)}>
                        삭제
                    </Button>
                </Modal.Footer>
            </Modal>
        </Container>
    );
};
