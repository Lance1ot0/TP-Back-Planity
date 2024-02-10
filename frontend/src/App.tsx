import { Route, Routes } from 'react-router-dom';
import s from './App.module.css';

import AdminPage from './components/pages/adminPage/adminPage';
import LoginAdmin from './components/login/loginAdmin/LoginAdmin';

function App() {
  return (
    <div className={s.appWrapper}>
      <Routes>
        <Route path="/" element={""} />
        <Route path="/admin" element={<AdminPage />} />
        <Route path="/admin/login" element={<LoginAdmin />} />
      </Routes>
    </div>
  );
}

export default App;