import style from './ErrorPage.module.css';
import cat from '../../../assets/cat.gif';
export default function ErrorPage() {
    return (
        <div className={style.container}>
            <h1 style={{fontSize: '100px'}}>Error 404</h1>
            <p style={{fontSize: '50px'}}>Page not found</p>
            <img style={{height: '300px', marginTop: '20px'}} src={cat} alt="Error GIF" />
        </div>
    );
}