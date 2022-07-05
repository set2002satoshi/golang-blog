import type { NextPage } from 'next'
import useSWR from "swr"
import Header from '../components/Header';
import Profile from '../components/myprofile/profile';
import PostCard from '../components/home/PostCard';
import LoadingMotion from '../components/LoadMotion/LoadMotion'
import styles from '../styles/myprofile.module.css';
import { parseCookies } from 'nookies'
import { style } from '@mui/system';


const url = "http://localhost:8000/api/main/my-profile"

const myprofile: NextPage = () => {
    
    // async function MyProfile(url: string): Promise<Boolean | null> {
    //     const response = await fetch(url)
    //     return response.json()
    // }

    // const { data, error } = useSWR(url, MyProfile)


    const ClientValue = parseCookies().ClientKey
    async function example(url: string): Promise<boolean | null> {
        const response = await fetch(url, {
            method: 'GET',
            headers: {
                "ClientKey": ClientValue, 
            },
        })
        return response.json();
    }
    const { data, error } = useSWR(url, example);
    
    // const Customer = data.Customer




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


    const blogs = data.Blogs;
    const Customer = data.Customer;

    return (
        <div id={styles.home}>
            
            <div className={styles.header}>
                <div className={styles.navStyle}>
                <Header />
                </div>
            </div>
            <div className={styles.navStyle}>
                    <div className={styles.profile}>
                        <div className={styles.contents}>
                            <div className={styles.myImg}>
                                <img src="https://via.placeholder.com/1400x500" width="100" height="100"/>
                            </div>
                            <div className={styles.NameAdditionallyMessage}>
                                <h2>{Customer.name}</h2>
                                <p>message:</p>
                                <p>{Customer.message}</p>
                            </div>
                            <div className={styles.EditButton}>
                                <button>Profile Edit</button>
                            </div>
                        </div>
                    </div>
                </div>
            <br />
            <div className={styles.content}>
                <div className={styles.contentColumn}>
                    {blogs.map((blog) => (
                        <div>
                            <PostCard 
                                UserName={blog.UserName}
                                UserThumbnail={"https://via.placeholder.com/1400x500"}
                                Title={blog.title}
                                Thumbnail={blog.blog_image?`https://mysatoshitest.s3.ap-northeast-1.amazonaws.com/${blog.blog_image}`:"https://via.placeholder.com/1400x500"}
                                LinkID={blog.ID}
                                date={blog.CreatedAt}
                            />
                        </div>
                    ))}
                </div>
            </div>
            
        </div>
    )
}

export default myprofile;