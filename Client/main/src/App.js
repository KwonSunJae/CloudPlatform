import Routers from './routes';
import React from 'react';
import { QueryClientProvider } from 'react-query';

import { queryClient } from './configs/reactQuery';
function App() {
  return (
    <QueryClientProvider client={queryClient}>
      <Routers />
    </QueryClientProvider>
  );
}

export default App;
