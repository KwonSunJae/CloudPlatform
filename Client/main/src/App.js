import Routers from './routes';
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
