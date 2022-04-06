import { Route, HashRouter as Router, Routes } from 'react-router-dom';
import './App.css';
import AppBar from "./MuAppBar";
import MuMonth from './MuMonth';
import MuYear from './MuYear';

function App() {
  return (
    <div className="App">
      <Router>
        <AppBar />
        <Routes>
          <Route path='/' element={<MuYear />} />
          <Route path='/:year' element={<MuYear />} />
          <Route path='/:year/:month' element={<MuMonth />} />
        </Routes>
      </Router>
    </div>
  );
}

export default App;
