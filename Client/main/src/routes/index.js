import React from 'react';
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import Error from './error';
import Home from './home';
import Profile from './profile';
import Cluster from './createCluster'
import Machine from './createMachine'
import Container from './createContainer';
import VM from './createVM'
const Routers = () => {

    return (
        <Router>
            <Routes>
                <Route  path="/" element={<Home />} />
                <Route  path="/profile" element={<Profile />} />
                <Route  path="/create/cluster" element={<Cluster />} />
                <Route path= "/create/container" element={<Container/>}/>
                <Route  path="/create/machine" element={<Machine />} />
                <Route path='/create/vm' element={<VM/>}/>
                <Route path = "*" element={<Error/>}/>
            </Routes>
        </Router>
    );
};

export default Routers;