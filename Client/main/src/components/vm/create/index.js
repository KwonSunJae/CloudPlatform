// VMCreateForm.js
import React, { useState, useEffect } from 'react';
import './index.css';
import instance from '../../../apis/instance'
import SpinnerOurs from '../../sppinner';
const VMCreateForm = () => {
    const [vmName, setVmName] = useState('');

    const [externalIP, setExternalIP] = useState(false);
    const [internalIP, setInternalIP] = useState(false);
    const [image, setImage] = useState('');
    const [flavor, setFlavor] = useState('');
    const [unionMountImage, setUnionMountImage] = useState('');
    const [keyPair, setKeyPair] = useState('');
    const [selectedSecurityGroup, setSelectedSecurityGroup] = useState('');
    const [newKeyPairName, setNewKeyPairName] = useState('');
    const [flavorList, setFlavorList] = useState([]);
    const [imageList, setImageList] = useState([]);
    const [keypairList, setKeyPairList] = useState([]);
    const [securityGroups, setSecurityGroups] = useState([]);
    const [loading, setLoading] = useState(true);

    const [step, setStep] = useState(1);
    const [pemKey, setPemKey] = useState('');
    useEffect(() => {
        // Fetch Flavor options from API
        instance.post("/transaction", {
            "dest": "/resource/flavor",
            "method": "GET",
            "data": ""
        }).then((response) => {
            var flavors = JSON.parse(response.data.data).flavors;
            
            setFlavorList(flavors);
            setLoading(false);
            console.log(flavors);
        }).catch((error) => {
            setLoading(false);
            console.error('Error fetching OS options:', error);
        });

        // Fetch Keypair options from API
        instance.post("/transaction", {
            "dest": "/resource/keypair",
            "method": "GET",
            "data": ""
        }).then((response) => {
            var keypairs = JSON.parse(response.data.data).keypairs;
            setKeyPairList(keypairs);
            setLoading(false);
            console.log(keypairs);
        }).catch((error) => {
            setLoading(false);
            console.error('Error fetching OS options:', error);
        });

        // Fetch SecurityGroups options from API
        instance.post("/transaction", {
            "dest": "/resource/securitygroup",
            "method": "GET",
            "data": ""
        }).then((response) => {
            var securityGroups = JSON.parse(response.data.data).securityGroups;
            setSecurityGroups(securityGroups);
            setLoading(false);
            console.log(securityGroups);
        }
        ).catch((error) => {
            setLoading(false);
            console.error('Error fetching OS options:', error);
        }
        );

        // Fetch Image options from API
        instance.post("/transaction", {
            "dest": "/resource/image",
            "method": "GET",
            "data": ""
        }).then((response) => {
            var images = JSON.parse(response.data.data).images;
            setImageList(images);
            setLoading(false);
            console.log(images);
        }
        ).catch((error) => {
            setLoading(false);
            console.error('Error fetching OS options:', error);
        }
        );
    }, []);




    const handleCreateKeyPair = (event) => {
        if (newKeyPairName === '') {
            alert('Please enter a key pair name');
            return;
        }
        setLoading(true);
        const datas = JSON.stringify({ keypairName: newKeyPairName });

        instance.post("/transaction", {
            "dest": "/resource/keypair",
            "method": "POST",
            "data": datas
        }).then((response) => {
            setPemKey(response.data.pemKey);
            setKeyPairList([...keypairList, response.data]);
            setLoading(false);
            console.log(response.data);
        }).catch((error) => {
            setLoading(false);
            console.error('Error creating key pair:', error);
        });
    };

    const handleCopyPemKey = () => {
        navigator.clipboard.writeText(pemKey).then(() => {
            alert('Pem key copied to clipboard!');
        }).catch(err => {
            console.error('Failed to copy: ', err);
        });
    };

    const handleSubmit = (event) => {
        event.preventDefault();
        setLoading(true);
        instance.post("/vm", {
            Name: vmName,
            FlavorID: flavor,
            ExternalIP: externalIP ? "true" : "false",
            internalIP: internalIP ? "e3102f8e-6d1c-41b3-b4ef-3ca5d81b79da" : "false",
            SelectedOS: image,
            unionMountImage: unionMountImage,
            KeyPair: keyPair,
            SelectedSecurityGroup: selectedSecurityGroup,
        })
            .then(response => {
                console.log(response.data);
                window.location.href = "/vm/";
            })
            .catch(error => {
                console.error(error);
            });
    };

    if (loading) {
        return <SpinnerOurs />;
    }

    return (
        <div className="vm-create-form">
            <h2>Create New Virtual Machine</h2>
            {step === 1 && (
                <div>
                    <label>
                        VM Name:
                        <input type="text" value={vmName} onChange={(e) => setVmName(e.target.value)} />
                    </label>
                    <label>
                        External IP:
                        <input type="checkbox" checked={externalIP} onChange={() => setExternalIP(!externalIP)} />
                    </label>
                    <label>
                        Internal IP:
                        <input type="checkbox" checked={internalIP} onChange={() => setInternalIP(!internalIP)} />
                    </label>
                    <label>
                        Flavor:
                        <select value={flavor} onChange={(e) => setFlavor(e.target.value)}>
                            {flavorList.map((f) => (
                                <option key={f.id} value={f.id}>{f.name}</option>
                            ))}
                        </select>
                    </label>
                    <button onClick={() => setStep(2)}>Next</button>
                </div>
            )}
            {step === 2 && (
                <div>
                    <label>
                        Image:
                        <select value={image} onChange={(e) => setImage(e.target.value)}>
                            {imageList.map((img) => (
                                <option key={img.id} value={img.id}>{img.name}</option>
                            ))}
                        </select>
                    </label>
                    <button onClick={() => setStep(1)}>Back</button>
                    <button onClick={() => setStep(3)}>Next</button>
                </div>
            )}
            {step === 3 && (
                <div>
                    <label>
                        Key Pair:
                        <select value={keyPair} onChange={(e) => setKeyPair(e.target.value)}>
                            {keypairList.map((kp) => (
                                <option key={kp.id} value={kp.id}>{kp.name}</option>
                            ))}
                        </select>
                    </label>
                    <label>
                        New Keypair Name :
                        <input type="text" value={newKeyPairName} onChange={(e) => setNewKeyPairName(e.target.value)} />
                    </label>
                    <button onClick={handleCreateKeyPair}>Create Key Pair</button>
                    {pemKey && (
                        <div>
                            <textarea value={pemKey} readOnly />
                            <button onClick={handleCopyPemKey}>Copy Pem Key</button>
                        </div>
                    )}
                    <button onClick={() => setStep(2)}>Back</button>
                    <button onClick={() => setStep(4)}>Next</button>
                </div>
            )}
            {step === 4 && (
                <div>
                    <label>
                        Security Group:
                        <select value={selectedSecurityGroup} onChange={(e) => setSelectedSecurityGroup(e.target.value)}>
                            {securityGroups.map((sg) => (
                                <option key={sg.id} value={sg.id}>{sg.name}</option>
                            ))}
                        </select>
                    </label>
                    <label>
                        Union Mounted Image:
                        <span>서비스 준비중입니다.</span>
                    </label>
                    <button onClick={() => setStep(3)}>Back</button>
                    <button onClick={handleSubmit}>Create VM</button>
                </div>
            )}
        </div>
    );
};

export default VMCreateForm;
