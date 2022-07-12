import React,{ useEffect ,useState } from 'react'
import type { NextPage } from "next"
import Header from "../components/Header"
import styles from "../styles/blogForm.module.css"
import { parseCookies } from 'nookies'
import FormData from 'form-data'
import useSWR from "swr"
import { useRouter } from 'next/router'
import LoadingMotion from '../components/LoadMotion/LoadMotion'



const url = "http://localhost:8000/api/app/blog_push"



const blogForm: NextPage = () => {
    const router = useRouter()
    const [redirect, setRedirect] = useState<boolean>(false)
    const [loading, setLoading] = useState<boolean>(false)
    
    const [title, setTitle] = useState<string>("")
    const [subtitle, setSubtitle] = useState<string>("")
    const [content, setContent] = useState<string>("")
    const [tag, setTag] = useState<number>(1)
    const [File, setFile] = useState<React.SetStateAction<FileList> | undefined>()
    const ClientValue = parseCookies().ClientKey

    const url = "http://localhost:8000/api/check"
  
    async function check(): Promise<boolean | null> {
        const resp = await fetch(url, {
            method: "GET",
            headers: {
                "ClientKey": ClientValue, 
            },
        })
        return resp.json()
    }

    const { data, error } = useSWR(url, check);
    
    console.log(data)

    if (!data) {
        return (
            <>
                <LoadingMotion />
            </>
        )
    }

    console.log(data)

    if (!data.status) {
        router.replace("/login")
        return
    }
    
    const submit = async (): Promise<void> => {
        try {
            setLoading(true)
            const PushFormData = new FormData();
            PushFormData.append("title", title)
            PushFormData.append("subtitle", subtitle)
            PushFormData.append("content", content)
            PushFormData.append("tag", tag)
            PushFormData.append("file", File ? File: "")
            console.log(PushFormData);
            await fetch("http://localhost:8000/api/app/blog_push", {
                method: "POST",
                headers: {
                    // "Content-Type": "multipart/form-data; boundary=------some-random-characters",
                    "ClientKey": ClientValue,
                },
                body: PushFormData,
            })
            .then(response => response.json())
            .then((response) => {
                console.log(response.status);
                if (response.status) {
                    // setRedirect(true)
                    console.log("実行されてる")
                    router.push("/home")
                }
                
            })
            setLoading(false)
        } catch (e) {
            setLoading(false)
            alert("not create")
            router.replace("/createForm")
        }        
        // if (redirect) {
        //     console.log("ダイレクト")
        // }
    }

    if (loading) {
        return (
            <div id={styles.home}>
                <div className={styles.header}>
                    <div className={styles.navStyle}>
                    <Header />
                    </div>
                </div>
                <div className={styles.content}>
                    <div className={styles.contentColumn}>
                        <form action="{{url}}" method="post" >
                            <p>タイトル</p>
                            <div className="form-floating mb-3">
                                <input type="email" className="form-control" id="floatingInput" placeholder="title" onChange={e => setTitle(e.target.value)} />
                                <label htmlFor="floatingInput">please your blog name....</label>
                            </div>
                            <p>サブタイトル</p>
                            <div className="form-floating">
                                <textarea onChange={e => setSubtitle(e.target.value)} className="form-control" placeholder="subTitle" id="floatingTextarea"></textarea>
                                <label htmlFor="floatingTextarea">subTitle</label>
                            </div>
                            <p style={{margin: "20px 0"  }}>本文</p>
                            <div className="form-floating">
                                <textarea onChange={e => setContent(e.target.value)} className="form-control" placeholder="contents" id="floatingTextarea2" style={{height: "300px"  }}></textarea>
                                <label htmlFor="floatingTextarea2">Contents</label>
                            </div>
                            <div className="mb-3">
                                <label htmlFor="formFile" className="form-label">jpegのみ可能</label>
                                <input type="file" accept="image/jpeg" onChange={(e) => setFile(e.target.files[0])} className="form-control"  id="formFile" />
                            </div>
                            <LoadingMotion />
                        </form>
                    </div>
                </div>
            </div>
        )
    }


    return (
        <div id={styles.home}>
        <div className={styles.header}>
            <div className={styles.navStyle}>
            <Header userName={data.userName} certification={data.status} />
            </div>
        </div>
        <div className={styles.content}>
            <div className={styles.contentColumn}>
                <form action="{{url}}" method="post" >
                    <p>タイトル</p>
                    <div className="form-floating mb-3">
                        <input type="email" className="form-control" id="floatingInput" placeholder="title" onChange={e => setTitle(e.target.value)} />
                        <label htmlFor="floatingInput">please your blog name....</label>
                    </div>
                    <p>サブタイトル</p>
                    <div className="form-floating">
                        <textarea onChange={e => setSubtitle(e.target.value)} className="form-control" placeholder="subTitle" id="floatingTextarea"></textarea>
                        <label htmlFor="floatingTextarea">subTitle</label>
                    </div>
                    <p style={{margin: "20px 0"  }}>本文</p>
                    <div className="form-floating">
                        <textarea onChange={e => setContent(e.target.value)} className="form-control" placeholder="contents" id="floatingTextarea2" style={{height: "300px"  }}></textarea>
                        <label htmlFor="floatingTextarea2">Contents</label>
                    </div>
                    <div className="mb-3">
                        <label htmlFor="formFile" className="form-label">jpegのみ可能</label>
                        <input type="file" accept="image/jpeg" onChange={(e) => setFile(e.target.files[0])} className="form-control"  id="formFile" />
                    </div>
                    <input style={{margin: "50px 0", width: "100%"}} className="btn btn-primary" onClick={submit} type="button" value="送信" />
                </form>
            </div>
        </div>
    </div>
    )
}

export default blogForm