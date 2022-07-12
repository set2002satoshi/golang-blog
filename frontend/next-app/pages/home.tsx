import React, { useEffect, useState } from 'react'
import useSWR from 'swr'  
import type { NextPage } from 'next'
import Head from 'next/head'
import RegisterForm from '../components/certifications/RegisterForm'
import Header from '../components/Header';
import SearchBar from '../components/home/SearchBar';
import PostCard from '../components/home/PostCard';
import LoadingMotion from '../components/LoadMotion/LoadMotion'
import styles from '../styles/Home.module.css';
import LoadMotion from '../components/LoadMotion/LoadMotion'
import { parseCookies } from "nookies"
import { useRouter } from "next/router"

// const url = "http://localhost:8000/api/app/blog_all"
const url = "http://localhost:8000/api/main/home"



type blogstype = {
    CustomerInfoID: number
	BlogImage: string
	Subtitle:  string 
	Content:  string 
	Tags: tags
}

type tags = {
    tags: string
}


// const [objects, setObjects] = useState<null | blogstype >()
const ClientValue = parseCookies().ClientKey

const home: NextPage = () => {
    
    const router = useRouter()
    useEffect(() => {
        router.replace("/home")
    },[])


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
    
    
    
    

    if (error) {
        return (
            <>            
                <div id={styles.home}>
                    <div className={styles.header}>
                        <div className={styles.navStyle}>
                        <Header userName={""} certification={false} />
                        </div>
                    </div>
                    <SearchBar />
                    <div className={styles.content}>
                        <div className={styles.contentColumn}>
                            <h1>エラー</h1>
                            <LoadMotion />
                        </div>
                    </div>
                </div>
            </>
        )
    }

    if (data === undefined) {
        return (
            <>            
                <div id={styles.home}>
                    <div className={styles.header}>
                        <div className={styles.navStyle}>
                        <Header userName={""} certification={false} />
                        </div>
                    </div>
                    <SearchBar />
                    <div className={styles.content}>
                        <div className={styles.contentColumn}>
                            <LoadMotion />
                        </div>
                    </div>
                </div>
            </>
        )
    }

    const user = data.username
    const userCheck = data.certification
    const blogs = data.blogs
    console.log(data)
    return (
        <>            
            <div id={styles.home}>
                <div className={styles.header}>
                    <div className={styles.navStyle}>
                    <Header userName={user} certification={userCheck} />
                    </div>
                </div>
                <SearchBar />
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
        </>
    )
}


export default home



