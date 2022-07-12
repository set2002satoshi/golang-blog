import Head from 'next/head'
import React,{ useState } from 'react'
import type { NextPage } from 'next'
import { useRouter } from 'next/router'
import { setCookie } from 'nookies'
import LoginForm from '../components/certifications/LoginForm'
import { NextPageContext } from 'next';
import Header from "../components/Header"
import styles from '../styles/Home.module.css';

const url = "http://localhost:8000/api/certification"


const Login: NextPage = () => {
    const [Redirect, setRedirect] = useState<boolean>(false)
    const router = useRouter()

    const [email, setEmail] = useState<string>("");
    const [password, setPassword] = useState<string>("");

    const submit = async (): Promise<void> => {
        const data = {
            email,
            password
            // email: "root@a.com",
            // password: "pass"
        }
        try {
            await fetch(url, {
                method: "POST",
                headers: {
                    'Content-Type': 'application/json',
                    'Access-Control-Allow-Origin': 'http://localhost:8000'
                },
                body: JSON.stringify(data),
            }).then(resp => resp.json())
            .then(resp => {
                console.log(resp.status);
                console.log(resp.ok);
                if (resp.status) {
                    setCookie(null,"ClientKey", resp.ClientKey, {
                        maxAge: 30 * 60 * 24 * 60,
                        path: "/",
                    })
                    // router.push('/home')
                    router.replace('/home')
                    setRedirect(true)
                }
                console.log(resp)
            })
        } catch (e) {
            alert("パスワードがちげぇーよばーか ミスるならloginするなよ")
        }
    }
    return (
        <div>
            <div className={styles.header}>
                <div className={styles.navStyle}>
                <Header userName={""} certification={false} />
                </div>
            </div>
            <LoginForm
                setEmail={setEmail} 
                Email={email} 
                setPassword={setPassword}
                Password={password}
                EventsTrigger={submit}
            />
        </div>
    )

}

export default Login