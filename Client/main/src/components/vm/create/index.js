// VMCreateForm.js
import React, { useState, useEffect } from 'react';
import './index.css';
import instance from '../../../apis/instance'
const VMCreateForm = () => {
    const [vmName, setVmName] = useState('');
    const [gpu, setGPU] = useState(false);
    const [cpuCores, setCpuCores] = useState(1);
    const [ram, setRam] = useState(2);
    const [externalIP, setExternalIP] = useState(false);
    const [internalIP, setInternalIP] = useState('');
    const [selectedOS, setSelectedOS] = useState('');
    const [unionMountImage, setUnionMountImage] = useState([]);
    const [keyPair, setKeyPair] = useState('');
    const [selectedSecurityGroup, setSelectedSecurityGroup] = useState('');

    const [osOptions, setOSOptions] = useState([]);
    const [internalIPOptions, setInternalIPOptions] = useState([]);
    const [keyPairOptions, setKeyPairOptions] = useState([]);

    useEffect(() => {
        // Fetch OS options from API
        fetch('osApiUrl')
            .then(response => response.json())
            .then(data => setOSOptions(data))
            .catch(error => console.error('Error fetching OS options:', error));

        // Fetch Internal IP options from API
        fetch('internalIpApiUrl')
            .then(response => response.json())
            .then(data => setInternalIPOptions(data))
            .catch(error => console.error('Error fetching Internal IP options:', error));

        // Fetch Key Pair options from API
        fetch('keyPairApiUrl')
            .then(response => response.json())
            .then(data => setKeyPairOptions(data))
            .catch(error => console.error('Error fetching Key Pair options:', error));
    }, []);
    const handleCreateKeyPair = () => {
        // Implement logic to create a new key pair
    };

    const handleSubmit = (event) => {
        event.preventDefault();
        instance.post("/vm",{
            Name : vmName,
            FlavorID : "df2ca4c3-715a-4749-b955-1c73a9aca1a6",
            ExternalIP : "true",
            internalIP : "false",
            SelectedOS : "45adb171-6051-43db-af19-ec62c00a1bf2",
            unionMountImage: "false",
            KeyPair : keyPair,
            SelectedSecurityGroup : "default",
            UserID : "myuser"
        })
        .then(response => {
            console.log(response.data);
            window.location.href="/vm/";
        })
        .catch(error => {
            console.error(error);
        });


        // Implement logic to create the virtual machine
    };

    return (
        <div className="vm-create-form">
            <h2>Create New Virtual Machine</h2>
            <form onSubmit={handleSubmit}>
                <label>
                    VM Name:
                    <input
                        type="text"
                        value={vmName}
                        onChange={(e) => setVmName(e.target.value)}
                    />
                </label>

                <label>
                    GPU:
                    <input
                        type="checkbox"
                        checked={gpu}
                        onChange={() => setGPU(!gpu)}
                    />
                </label>

                <label>
                    CPU Cores:
                    <input
                        type="number"
                        value={cpuCores}
                        onChange={(e) => setCpuCores(Number(e.target.value))}
                        min={1}
                    />
                </label>

                <label>
                    RAM (GB):
                    <input
                        type="number"
                        value={ram}
                        onChange={(e) => setRam(Number(e.target.value))}
                        min={1}
                    />
                </label>

                <label>
                    External IP:
                    <input
                        type="checkbox"
                        checked={externalIP}
                        onChange={() => setExternalIP(!externalIP)}
                    />
                </label>

                <label>
                    Internal IP:
                    <select
                        value={internalIP}
                        onChange={(e) => setInternalIP(e.target.value)}
                    >
                        <option value="">Select Internal IP</option>
                        {internalIPOptions.map(option => (
                            <option key={option.id} value={option.value}>{option.label}</option>
                        ))}
                    </select>
                </label>

                <label>
                    Selected OS:
                    <select
                        value={selectedOS}
                        onChange={(e) => setSelectedOS(e.target.value)}
                    >
                        <option value="">Select OS</option>
                        {osOptions.map(option => (
                            <option key={option.id} value={option.value}>{option.label}</option>
                        ))}
                    </select>
                </label>

                <label>
                    Union Mount Image:
                    <div>
                        <input
                            type="checkbox"
                            value="PX4"
                            checked={unionMountImage.includes('PX4')}
                            onChange={() => {
                                if (unionMountImage.includes('PX4')) {
                                    setUnionMountImage(unionMountImage.filter(item => item !== 'PX4'));
                                } else {
                                    setUnionMountImage([...unionMountImage, 'PX4']);
                                }
                            }}
                        />
                        PX4
                    </div>
                    {/* Repeat similar checkboxes for other options */}
                </label>

                <label>
                    Key Pair:
                    <select
                        value={keyPair}
                        onChange={(e) => setKeyPair(e.target.value)}
                    >
                        <option value="">Select Key Pair</option>
                        <option value="dmslab">dmslab</option>
                        <option value="batman">batman</option>
                    </select>
                    <button type="button" onClick={handleCreateKeyPair}>Create New Key Pair</button>
                </label>

                <label>
                    Security Group:
                    <select
                        value={selectedSecurityGroup}
                        onChange={(e) => setSelectedSecurityGroup(e.target.value)}
                    >
                        <option value="">Select Security Group</option>
                        {keyPairOptions.map(option => (
                            <option key={option.id} value={option.value}>{option.label}</option>
                        ))}
                    </select>
                </label>

                <button type="submit">Create VM</button>
            </form>
        </div>
    );
};

export default VMCreateForm;
