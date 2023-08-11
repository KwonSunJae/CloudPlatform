// VMCreateForm.js
import React, { useState } from 'react';
import './index.css';
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

    const handleCreateKeyPair = () => {
        // Implement logic to create a new key pair
    };

    const handleSubmit = (event) => {
        event.preventDefault();
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
                        <option value="192.168.100.0">192.168.100.0</option>
                        <option value="192.168.200.0">192.168.200.0</option>
                    </select>
                </label>

                <label>
                    Selected OS:
                    <select
                        value={selectedOS}
                        onChange={(e) => setSelectedOS(e.target.value)}
                    >
                        <option value="">Select OS</option>
                        <option value="ubuntu 20.04">Ubuntu 20.04</option>
                        <option value="ubuntu 18.04">Ubuntu 18.04</option>
                        <option value="Windows 10">Windows 10</option>
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
                        <option value="Default">Default</option>
                        <option value="release">release</option>
                    </select>
                </label>

                <button type="submit">Create VM</button>
            </form>
        </div>
    );
};

export default VMCreateForm;
