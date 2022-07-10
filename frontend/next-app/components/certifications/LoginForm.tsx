import React, { Dispatch, SetStateAction } from 'react';
import styles from './LoginForm.module.css';

type Props = {
    setEmail: Dispatch<SetStateAction<string>>
    Email: string
    setPassword: Dispatch<SetStateAction<string>>
    Password: string
    EventsTrigger: VoidFunction
}


const LoginForm = (props: Props) => {
    return (
        <div className={styles.box_P}>
            <div className={styles.box_C}>
                {/* <form method="POST"> */}
                <form method="POST">
                    <div className="mb-3">
                        <label 
                            htmlFor="exampleInputEmail1" 
                            className="form-label"
                        >
                            Email address
                        </label>
                        <input 
                            type="email" 
                            className="form-control" 
                            id="exampleInputEmail1" 
                            aria-describedby="emailHelp" 
                            value={props.Email} 
                            onChange={e => props.setEmail(e.target.value)} 
                        />
                    </div>
                    <div className="mb-3">
                        <label htmlFor="exampleInputPassword1" className="form-label">Password</label>
                        <input type="password" className="form-control" id="exampleInputPassword1" value={props.Password} onChange={e => props.setPassword(e.target.value)} />
                    </div>
                    {/* <button className="btn btn-primary" onClick={props.EventsTrigger}>Login</button> */}
                    <input type="button" value="Login" className="btn btn-primary" onClick={props.EventsTrigger}/>
                </form>
            </div>
        </div>
    )
}

export default LoginForm





