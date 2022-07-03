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

const url = "http://localhost:8000/api/app/blog_all"



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


const home: NextPage = () => {



    async function example(url: string): Promise<boolean | null> {
        const response = await fetch(url)
        return response.json();
    }
    const { data, error } = useSWR(url, example);
    

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

    if (data === undefined) {
        return (
            <div id={styles.home}>
                <div className={styles.header}>
                    <div className={styles.navStyle}>
                    <Header />
                    </div>
                </div>
                <SearchBar />
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

    const blogs = data.user

    return (
        <>            
            <div id={styles.home}>
                <div className={styles.header}>
                    <div className={styles.navStyle}>
                    <Header />
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



