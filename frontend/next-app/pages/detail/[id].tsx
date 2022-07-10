import type { NextPage } from "next"
import Header from '../../components/Header';
import styles from '../../styles/detail.module.css';
import { useRouter } from 'next/router'
import LoadingMotion from '../../components/LoadMotion/LoadMotion'
import useSWR from "swr"
import { parseCookies } from 'nookies'



const Detail: NextPage = () => {
    const router = useRouter()
    const { id } = router.query
    const ClientValue = parseCookies().ClientKey
    const url = `http://localhost:8000/api/app/blog?id=${id}`
    async function example(url: string): Promise<Boolean | null> {
        
        const response = await fetch(url, {
            method: 'GET',
            headers: {
                "ClientKey": ClientValue,
            },
        })
        return response.json()
    }
    console.log(url)
    const { data, error } = useSWR(url, example)
    
    
    
    

    if (data === undefined) {
        return (
            <div id={styles.home}>
                <div className={styles.header}>
                    <div className={styles.navStyle}>
                    <Header />
                    </div>
                </div>
                <div className={styles.content}>
                    <div className={styles.contentColumn}>
                        <div className={styles.loadContent}>
                            <LoadingMotion />
                        </div>
                    </div>
                </div>
            </div>
        )
    }
    
    if (error) {
        return (
            <div id={styles.home}>
                <div className={styles.header}>
                    <div className={styles.navStyle}>
                    <Header />
                    </div>
                </div>
                <div className={styles.content}>
                    <div className={styles.contentColumn}>
                    <div className="d-flex justify-content-center align-items-center" id="main">
                        <h1 className="mr-3 pr-3 align-top border-right inline-block align-content-center">404</h1>
                        <div className="inline-block align-middle">
                            <h2 className="font-weight-normal lead" id="desc">The page you requested was not found.</h2>
                        </div>
                    </div>
                    </div>
                </div>
            </div>
        )
    }
    
    console.log(data.blog)
    const BlogData = data.blog[0]
    
    return (
        <>
            <div className={styles.header}>
                <div className={styles.navStyle}>
                <Header />
                </div>
                </div>  
                <div className={styles.mainImg}>
                <img src={ BlogData.blog_image !== undefined ?`https://mysatoshitest.s3.ap-northeast-1.amazonaws.com/${ BlogData.blog_image}` : "https://via.placeholder.com/1400x500"}  />
                </div>
                <div className={styles.content}>
                <div className={styles.contentColumn}>
                <h1>{ BlogData.title }</h1>
                        <br/>
                        <p>{ BlogData.content }</p>
                    </div>
                </div>
        </>
    )
}

export default Detail