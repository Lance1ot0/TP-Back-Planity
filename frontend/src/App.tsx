import { Route, Routes } from 'react-router-dom';
import style from './App.module.css';

import AdminPage from './components/pages/adminPage/AdminPage';
import LoginAdmin from './components/login/loginAdmin/LoginAdmin';

function App() {
  return (
    <div className={style.appWrapper}>
      <Routes>
        <Route path="/" element={""} />
        <Route path="/admin" element={<AdminPage />} />
        <Route path="/admin/login" element={<LoginAdmin />} />
      </Routes>
    </div>
  );
}

export default App;