import React from "react"

import * as styles from "../styles/components/navbar.module.scss"

const Navbar = () => {

    return (
        <nav className={styles.navbar}>
            <div className={styles.inner}>
                <div className={styles.logo}>
                    <span>quickconvert</span>
                </div>

                <div className={styles.items}>

                </div>
            </div>
        </nav>
    )
}

export default Navbar