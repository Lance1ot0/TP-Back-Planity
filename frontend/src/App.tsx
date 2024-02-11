import { Route, Routes } from "react-router-dom";
import style from "./App.module.css";

import Homepage from "./components/pages/Homepage/Homepage";

import AdminPage from "./components/pages/adminPage/adminPage";
import LoginAdmin from "./components/login/loginAdmin/LoginAdmin";

import ProfessionalPage from "./components/pages/professionalPage/ProfessionalPage";
import LoginProfessional from "./components/login/loginProfessional/LoginProfessional";
import RegisterProfessional from "./components/register/registerProfessional/RegisterProfessional";

import ErrorPage from "./components/pages/errorPage/ErrorPage";
import Navbar from "./components/Navbar/Navbar";
import LoginSelection from "./components/pages/Login/LoginSelection";
import RegisterSelection from "./components/pages/Register/RegisterSelection";

function App() {
  return (
    <div className={style.appWrapper}>
      <Navbar />
      <Routes>
        <Route path="/" element={<Homepage />} />
        <Route path="/login" element={<LoginSelection />} />
        <Route path="/register" element={<RegisterSelection />} />

        <Route path="/admin" element={<AdminPage />} />
        <Route path="/admin/login" element={<LoginAdmin />} />

        <Route path="/professional" element={<ProfessionalPage />} />
        <Route path="/professional/login" element={<LoginProfessional />} />
        <Route
          path="/professional/register"
          element={<RegisterProfessional />}
        />

        <Route path="*" element={<ErrorPage />} />
      </Routes>
    </div>
  );
}

export default App;
