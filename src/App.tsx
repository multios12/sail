import { Route, HashRouter as Router, Routes } from 'react-router-dom';
import './App.css';
import AppBar from "./components/MuAppBar";
import MuSalaryYear from './components/MuSalaryYear';
import MuSalaryMonth from './components/MuSalaryMonth';

function App() {
  return (
    <div className="App">
      <Router>
        <AppBar />
        <Routes>
          <Route path='/' element={<MuSalaryYear />} />
          <Route path='/:year' element={<MuSalaryYear />} />
          <Route path='/:year/:month' element={<MuSalaryMonth />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
