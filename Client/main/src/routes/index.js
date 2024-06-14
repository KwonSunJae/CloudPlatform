import React from 'react';
import { BrowserRouter as Router, Routes, Route } from "react-router-dom";

import Error from './error';
import Home from './home';
import Introduction from './introduction';
import Monitoring from './monitoring';
import System from './system';
import MyPage from './mypage';
import Cluster from './createCluster'
import Machine from './createMachine'
import Container from './createContainer';
import VM from './createVM'
import VMStat from './statusVM';
import ContainerStat from './statusContainer';
import AdminPage from './admin';
import AdminPageVM from './admin/vm';
import AdminPageContainer from './admin/container';
const Routers = () => {

    return (
        <Router>
            <Routes>
                <Route  path="/" element={<Home />} />
                <Route  path="/mypage" element={<MyPage />} />
                <Route  path="/introduction" element={<Introduction />} />
                <Route  path="/system" element={<System />} />
                <Route  path="/monitoring" element={<Monitoring />} />
                <Route  path="/create/cluster" element={<Cluster />} />
                <Route path= "/create/container" element={<Container/>}/>
                <Route  path="/create/machine" element={<Machine />} />
                <Route path='/create/vm' element={<VM/>}/>
                <Route path='/vm' element={<VMStat/>}/>
                <Route path='/container' element={<ContainerStat/>}/>
                <Route path="/admin" element={<AdminPage />} />
                <Route path="/admin/vm" element={<AdminPageVM />} />
                <Route path="/admin/container" element={<AdminPageContainer />} />
                <Route path = "*" element={<Error/>}/>
            </Routes>
        </Router>
    );
};

export default Routers;