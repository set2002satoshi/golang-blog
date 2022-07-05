import styles from './profile.module.css';


const Profile = () => {

    return (
        <>
            <div className={styles.profile}>
                <div className={styles.contents}>
                    <div className={styles.myImg}>
                        <img src="https://via.placeholder.com/1400x500" width="100" height="100"/>
                    </div>
                    <div className={styles.NameAdditionallyMessage}>
                        <h2>sssssaaaaaaaaaaaaaaa</h2>
                        <p>message:</p>
                        <p>sssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssssss</p>
                    </div>
                    <div className={styles.EditButton}>
                        <button>Profile Edit</button>
                    </div>
                </div>
            </div>
        </>
    )


}

export default Profile