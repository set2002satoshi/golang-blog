import React, {FC, useState } from 'react'; 
import styles from './RegisterForm.module.css';
import { useRouter } from 'next/router'

const url = "http://localhost:8000/api/create-user"

const RegisterForm: FC  = () => {
    const [Name, setName] = useState<string>("");
    const [Email, setEmail] = useState<string>("");
    const [Password, setPassword] = useState<string>("");
    const [Redirect, setRedirect] = useState<boolean>(false);
    const router = useRouter()

    const submit = async (): Promise<void> => {
        // const data = {
        //     "Email": Email,
        //     "Password": Password,
        //     "Name": Name
        // }
        const data = {
            Email,
            Password,
            Name
        }
        
        console.log(data);
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
                    setRedirect(true);
                }}
                ) 
            } catch {
                console.log("aaaf");
            }
            if (Redirect) {
                await router.push('/home');
            }
        
            
            
        }
        

    return (
        
        <div className={styles.box_P}>
            <div className={styles.box_C}>
                <form method="POST">
                    <div className="mb-3">
                        <label htmlFor="exampleInputPassword1" className="form-label">Id Name</label>
                        <input type="text" className="form-control" id="exampleInputPassword1" value={Name} onChange={e => setName(e.target.value)} />
                    </div>
                    <div className="mb-3">
                        <label htmlFor="exampleInputEmail1" className="form-label">Email address</label>
                        <input type="email" className="form-control" id="exampleInputEmail1" aria-describedby="emailHelp" value={Email} onChange={e => setEmail(e.target.value)} />
                    </div>
                    <div className="mb-3">
                        <label htmlFor="exampleInputPassword1" className="form-label">Password</label>
                        <input type="password" className="form-control" id="exampleInputPassword1" value={Password} onChange={e => setPassword(e.target.value)} />
                    </div>
                    <input type="submit" className="btn btn-primary" onClick={submit}/>
                </form>
            </div>
        </div>
    )
}

export default RegisterForm;