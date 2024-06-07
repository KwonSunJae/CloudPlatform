import React, { useState,useEffect } from 'react';
import { Form, Button } from 'react-bootstrap';
import instance from '../../../apis/instance';
const Mypage = () => {
    const [name, setName] = useState(''); // [1
    const [userID, setUserID] = useState(''); // [1
    const [password, setPassword] = useState(''); // [1
    const [role, setRole] = useState('');
    const [spot, setSpot] = useState('');
    const [priority, setPriority] = useState('');

    useEffect(() => {
        // Add your logic here to fetch user data
        console.log('Fetching user data...');
        instance.get("/member", { headers: { "Access-Token" : `${localStorage.getItem("accessToken")}` } })
        .then((response)=>{
            console.log(response);
            setName(response.data.data.Name);
            setUserID(response.data.data.UserID);
            setRole(response.data.data.Role);
            setSpot(response.data.data.Spot);
            setPriority(response.data.data.Priority);
        })
        .catch((err)=>{
            console.log(err);
        });
    }
    , []);

    const handleSubmit = (e) => {
        e.preventDefault();
        // Add your logic here to update the user data
        instance.patch("/transaction",{
            "dest" : "/user/" + userID,
            "data" : {
                "Password" : password,
                "Role" : role,
                "Spot" : spot
            }
        })
    };

    return (
        <div className="container">
            <h1>My Page</h1>
            <Form onSubmit={handleSubmit}>
                <Form.Group controlId="name">
                    <Form.Label>Name</Form.Label>
                    <Form.Control
                        type="text"
                        placeholder="Enter new name"
                        value={name}
                        disabled
                    />
                </Form.Group>
                <Form.Group controlId="userID">
                    <Form.Label>User ID</Form.Label>
                    <Form.Control
                        type="text"
                        placeholder="Enter new user ID"
                        value={userID}
                        disabled
                    />
                </Form.Group>
            
                <Form.Group controlId="role">
                    <Form.Label>Role</Form.Label>
                    <Form.Control
                        type="text"
                        placeholder="Enter new role"
                        value={role}
                        onChange={(e) => setRole(e.target.value)}
                    />
                </Form.Group>

                <Form.Group controlId="spot">
                    <Form.Label>Spot</Form.Label>
                    <Form.Control
                        type="text"
                        placeholder="Enter new spot"
                        value={spot}
                        onChange={(e) => setSpot(e.target.value)}
                    />
                </Form.Group>
                <Form.Group controlId="priority">
                    <Form.Label>Priority</Form.Label>
                    <Form.Control
                        type="text"
                        placeholder="Enter new priority"
                        value={priority}
                        disabled
                    />
                </Form.Group>
                <Button variant="primary" type="submit">
                    Update
                </Button>
            </Form>
        </div>
    );
};

export default Mypage;