import * as React from "react"

import SEO from "../components/seo"
import Navbar from "../components/navbar"

import Hero from "../components/hero"

const IndexPage = () => (
  <>
    <SEO title="World's fastest conversion tool" /> {/* eslint-disable-line react/jsx-pascal-case*/}

    <Navbar />
    
    <Hero />
    
  </>
)

export default IndexPage
