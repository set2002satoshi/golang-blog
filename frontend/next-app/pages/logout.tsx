import type { NextPage } from 'next'
import { useRouter } from 'next/router'
import { destroyCookie } from 'nookies';
import { parseCookies } from 'nookies';
import { useEffect } from 'react';

const Logout: NextPage = () => {
    const router = useRouter()
    const cookie = parseCookies()

    useEffect(() => {
        console.log(cookie.ClientKey)
        if (cookie.ClientKey !== undefined) {
            destroyCookie(null, "ClientKey")
        }
        router.replace('/login')
    }, [])
    
    return (
        <>
        </>
    )
}

export default Logout

