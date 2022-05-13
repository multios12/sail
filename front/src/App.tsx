import { Route, HashRouter as Router, Routes } from 'react-router-dom';
import AppBar from "./components/MuAppBar";
import MuHome from './components/MuHome';
import MuSalaryYear from './components/MuSalaryYear';
import MuSalaryMonth from './components/MuSalaryMonth';
import MuCostYear from './components/MuCostYear';

function App() {
  return (
    <div className="App">
      <Router>
        <AppBar />
        <Routes>
          <Route path='/' element={<MuHome />} />
          <Route path='/salary/' element={<MuSalaryYear />} />
          <Route path='/salary/:year' element={<MuSalaryYear />} />
          <Route path='/salary/:year/:month' element={<MuSalaryMonth />} />
          <Route path='/cost/' element={<MuCostYear />} />
          <Route path='/cost/:year' element={<MuCostYear />} />
          <Route path='/:year' element={<MuHome />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
