import { useState } from 'react';
import style from './LoginClient.module.css';
import { loginClientApi } from '../../../api/login/loginClientApi/loginClientApi';
import { useNavigate } from 'react-router-dom';

export default function loginClient() {
    const navigate = useNavigate();
    const  [errorMessage, setErrorMessage] = useState(false);

    const onHandleLogin = async (e: React.FormEvent<HTMLFormElement>) => {
        e.preventDefault();
        const email = (e.target as HTMLFormElement).email.value;
        const password = (e.target as HTMLFormElement).password.value;

        const response = await loginClientApi(email, password);
        if (response.success) {
            if (response.role && response.token) {
                localStorage.setItem('role', response.role);
                localStorage.setItem('token', response.token);

                navigate('/client');
            }
        } else {
            setErrorMessage(true);
        }
    };

    return (
        <div className={style.loginWrapper}>
            <form className={style.loginForm} onSubmit={onHandleLogin}>
                <h2 className={style.loginTitle}>Login Client</h2>
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
            <a className={style.link} href="/client/register">Inscription</a>
            {errorMessage && (
                <p className={style.errorMessage}>Invalid email or password</p>
            )}
        </div>
    );
}