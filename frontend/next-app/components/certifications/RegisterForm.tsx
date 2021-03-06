import React, { Dispatch, SetStateAction } from 'react';
import styles from './RegisterForm.module.css';


type Props = {
    setName: Dispatch<SetStateAction<string>>
    Name: string
    setEmail: Dispatch<SetStateAction<string>>
    Email: string
    setPassword: Dispatch<SetStateAction<string>>
    Password: string
    EventsTrigger: VoidFunction
}

const RegisterForm = (props: Props) => {
    return (
        <div className={styles.box_P}>
            <div className={styles.box_C}>
                <form method="POST">
                    <div className="mb-3">
                        <label htmlFor="exampleInputPassword1" className="form-label">Id Name</label>
                        <input 
                            type="text" 
                            className="form-control" 
                            id="exampleInputPassword1" 
                            value={props.Name} 
                            onChange={e => props.setName(e.target.value)} 
                        />
                    </div>
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
                    <input type="submit" className="btn btn-primary" onClick={props.EventsTrigger}/>
                    {/* <input type="submit" className="btn btn-primary" /> */}
                    
                </form>
            </div>
        </div>
    )
}

export default RegisterForm;