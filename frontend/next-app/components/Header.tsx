import styles from "./header.module.css";
import Link from "next/link";



const Header = () => {
    // return (
    //     <>
    //         <header id={styles.a1_0}>
    //             <div className={styles.b1_0}>
    //                 <h1>logo</h1>
    //             </div>
    //             <div className={styles.b1_2}>
    //                 <div>ID:</div>
    //                 <div>MyProfile</div>
    //                 <div>Logout</div>
    //             </div>
    //         </header>
    //     </>
    // )
    
    return (
        <>
            <header id={styles.a1_0}>
                <div className={styles.b1_0}>
                    <h1>logo</h1>
                </div>
                <div className={styles.b1_2}>
                    <div>ID:</div>
                    <Link href="/myprofile">
                        <div>MyProfile</div>
                    </Link>
                    <Link href="/login">
                        <div>login</div>
                    </Link>
                    <div>singup</div>
                </div>
            </header>
        </>
    )
}

export default Header



