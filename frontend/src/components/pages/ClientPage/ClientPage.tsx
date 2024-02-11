import style from './ClientPage.module.css'
import { useNavigate } from 'react-router-dom';
import { useEffect, useState } from 'react';

function ClientPage() {
    const navigate = useNavigate();
    const [role, setRole] = useState<string | null>(null);
    const [token, setToken] = useState<string | null>(null);

    const logout = () => {
        localStorage.removeItem('role');
        localStorage.removeItem('token');
        setRole(null);
        setToken(null);
        navigate('/admin/login');
    };

    useEffect(() => {
        const role = localStorage.getItem('role');
        const storedToken = localStorage.getItem('token');
        if (role && storedToken) {
            setRole(role);
            setToken(storedToken);
        } else {
            navigate('/client/login'); 
        }
    }, [navigate]);

    if (!role || role !== 'client' || !token) {
        navigate('/client/login');
        return null;
    }
  return (
    <div className={style.container}>
        <div className={style.menu}>
                <h1 style={{ fontSize: '38px' }}>Client Page</h1>
                <button
                    className={style.logoutButton}
                    onClick={() => { logout(); }}
                >Déconnexion</button>
            </div>
    </div>
  )
}

export default ClientPage
