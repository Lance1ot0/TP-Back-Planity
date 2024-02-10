import { useState } from 'react';
import style from './LoginAdmin.module.css';
import { loginAdminApi } from '../../../api/login/loginAdminApi/loginAdminApi';
import { useNavigate } from 'react-router-dom';

export default function loginAdmin() {
    const navigate = useNavigate();
    const  [errorMessage, setErrorMessage] = useState(false);

    const onHandleLogin = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const email = (e.target as HTMLFormElement).email.value;
        const password = (e.target as HTMLFormElement).password.value;

        const response = await loginAdminApi(email, password);
        if (response.success) {
            if (response.token) {
                localStorage.setItem('token', response.token);

                navigate('/admin');
            }
        } else {
            setErrorMessage(true);
        }
    };

    return (
        <div className={style.loginWrapper}>
            <form className={style.loginForm} onSubmit={onHandleLogin}>
                <h2 className={style.loginTitle}>Login Admin</h2>
                <label className={style.loginLabel} htmlFor="email">
                    Email
                </label>
                <input
                    className={style.loginInput}
                    type="text"
                    id="email"
                    name="email"
                    required
                />
                <label className={style.loginLabel} htmlFor="password">
                    Password
                </label>
                <input
                    className={style.loginInput}
                    type="password"
                    id="password"
                    name="password"
                    required
                />
                <button className={style.loginButton} type="submit">
                    Login
                </button>
            </form>
            {errorMessage && (
                <p className={style.errorMessage}>Invalid email or password</p>
            )}
        </div>
    );
}