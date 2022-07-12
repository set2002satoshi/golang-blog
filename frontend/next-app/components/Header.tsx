import styles from "./header.module.css";
import Link from "next/link";


type Props = {
    userName: string
    certification: boolean
}


const Header = (props: Props) => {

    if (!props.certification) {
        return (
            <>
                <header id={styles.a1_0}>
                    <div className={styles.b1_0}>
                        <h1>logo</h1>
                    </div>
                    <div className={styles.b1_2}>

                        <Link href="/login">
                            <div>login</div>
                        </Link>
                        <Link href="/register">
                            <div>singup</div>
                        </Link>
                    </div>
                </header>
            </>

        )
    }

    return (
        <>
            <header id={styles.a1_0}>
                <div className={styles.b1_0}>
                    <h1>logo</h1>
                </div>
                <div className={styles.b1_2}>
                    <div>ID:{props.userName}</div>
                    <Link href="/createForm">
                        <div>create</div>
                    </Link>
                    <Link href="/myprofile">
                        <div>MyProfile</div>
                    </Link>
                    <Link href="/logout">
                        <div>logout</div>
                    </Link>
                </div>
            </header>
        </>
    )
}

export default Header



