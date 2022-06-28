import styles from './SearchBar.module.css';


const SearchBar = () => {

    return (
        <>
            <div id={styles.a1_0}>
                <div className={styles.b1_0}>
                    <div className={styles.c1_0}>
                        <input className={styles.d1_0} type="text" value="キーワード... name..." />
                        <input className={styles.d1_1} type="submit" value="button" value="Search" />
                    </div>
                </div>

            </div>
        </>
    )
}


export default SearchBar