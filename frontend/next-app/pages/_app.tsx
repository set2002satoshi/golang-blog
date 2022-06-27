import { useEffect } from 'react';

import '../styles/globals.css'
import 'bootstrap/dist/css/bootstrap.min.css';
import type { AppProps } from 'next/app'

function MyApp({ Component, pageProps, router }: AppProps) {

  useEffect(() => {
    const AuthenticationRequired = [];
    for (const page of AuthenticationRequired) {
      
    }

    if (router.pathname === "/logout") {
      
    }

  })

  return <Component {...pageProps} />
}

export default MyApp
