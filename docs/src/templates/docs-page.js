import * as React from "react"
import { graphql } from "gatsby"

import SEO from "../components/seo"


export default function DocsPage({data, pageContext}) {
    const page = data.markdownRemark
    // console.log(pageContext)

    return (
        <>
            <SEO title={page.frontmatter.title} /> {/* eslint-disable-line react/jsx-pascal-case*/}

            <h1>{page.frontmatter.title}</h1>
            
        </>
    )
  }

export const query = graphql`
  query($id: String!) {
    markdownRemark(id: {eq: $id}) {
      html
      frontmatter {
        title
      }
    }
  }
`