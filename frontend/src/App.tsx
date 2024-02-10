import { Route, Routes} from 'react-router-dom';
import style from './App.module.css';

import AdminPage from './components/pages/adminPage/AdminPage';
import LoginAdmin from './components/login/loginAdmin/LoginAdmin';

import ProfessionalPage from './components/pages/professionalPage/ProfessionalPage';
import LoginProfessional from './components/login/loginProfessional/LoginProfessional';
import RegisterProfessional from './components/register/registerProfessional/RegisterProfessional';

import ErrorPage from './components/pages/errorPage/ErrorPage';

function App() {
  return (
    <div className={style.appWrapper}>
      <Routes>
        <Route path="/" element={<ErrorPage />} />

        <Route path="/admin" element={<AdminPage />} />
        <Route path="/admin/login" element={<LoginAdmin />} />

        <Route path="/professional" element={<ProfessionalPage />} />
        <Route path="/professional/login" element={<LoginProfessional />} />
        <Route path="/professional/register" element={<RegisterProfessional />} />

        <Route path="*" element={<ErrorPage />} /> 
      </Routes>
    </div>
  );
}

export default App;