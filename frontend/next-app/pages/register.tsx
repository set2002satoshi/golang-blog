import React,{ useState } from 'react'; 
import type { NextPage } from 'next'
import Head from 'next/head'
import { useRouter } from 'next/router'
import RegisterForm from '../components/certifications/RegisterForm'

const url = "http://localhost:8000/api/create-user"


const Register: NextPage = () => {
    const [Redirect, setRedirect] = useState<boolean>(false);
    const router = useRouter()

    const [name, setName] = useState<string>("");
    const [email, setEmail] = useState<string>("");
    const [password, setPassword] = useState<string>("");

    const submit = async (): Promise<void> => {
        const data = {
            email,
            password,
            name
        }

        try {
            await fetch(url, {
                method: "POST",
                headers: {
                    'Content-Type': 'application',
                    'Access-Control-Allow-Origin': 'http://localhost:8000',
                },
                body: JSON.stringify(data),
        }).then(resp => {
            if (resp.ok) {
                console.log(resp.status)
                setRedirect(true)
            }}
            ) 
        } catch {
            console.log("err");
        }    
        if (Redirect) {
            router.replace('/home')
        }
    }
    return (
        <div>
            <RegisterForm
                setName={setName}
                Name={name}
                setEmail={setEmail} 
                Email={email} 
                setPassword={setPassword}
                Password={password}
            />
            <input type="submit" className="btn btn-primary" onClick={submit}/>
        </div>
    )

}

export default Register