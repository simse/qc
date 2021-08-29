import * as React from "react"

import SEO from "../components/seo"
import Navbar from "../components/navbar"

// Styles
import * as styles from "../styles/pages/install.module.scss"

const InstallPage = () => (
  <>
    <SEO title="Install" /> {/* eslint-disable-line react/jsx-pascal-case*/}

    <Navbar />
    
    <div className={styles.install}>
      <div className={styles.inner}>
        <h1>Install qc</h1>
        <p>qc is <strong>easy</strong> to install on all platforms, except Windows (because it is not supported).</p>
      </div>
    </div>
    
  </>
)

export default InstallPage
