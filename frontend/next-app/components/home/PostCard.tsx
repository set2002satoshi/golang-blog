import styles from './PostCard.module.css';

type Props = {
    UserName: string
    UserThumbnail : string
    Title: string
    Thumbnail : string
    LinkID : number
    date: string

}

const PostCard = (props: Props) => {

    return (
        <>
            <a className={styles.link} href={props.LinkID}>
                <div className={styles.a1_0}>
                    <div className={styles.b1_0}>
                        <img className={styles.c1_0} src={props.Thumbnail} />
                    </div>
                    <div className={styles.b1_1}>
                        <div className={styles.c1_1}>
                            <p>{props.Title}</p>
                        </div>
                        <div className={styles.c1_2}>
                            <div className={styles.d1_0}>
                                <img className={styles.e1_0} src={props.UserThumbnail}/>
                            </div>
                            <div className={styles.c1_3}>
                                <p className={styles.d1_2}>{props.UserName}</p>
                            </div>
                        </div>
                        <div className={styles.timearea}>
                            <time>{props.date}</time>
                        </div>
                    </div>
                </div>
            </a>
        </>
    )
}


export default PostCard