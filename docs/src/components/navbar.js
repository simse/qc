import React from "react"
import { Link } from "gatsby"

import * as styles from "../styles/components/navbar.module.scss"

const Navbar = () => {

    return (
        <nav className={styles.navbar}>
            <div className={styles.inner}>
                <Link to={"/"}>
                    <div className={styles.logo}>
                        <span>quickconvert</span>
                    </div>
                </Link>

                <div className={styles.items}>
                    <div className={styles.item}>
                        <Link to={"/install"}>Install</Link>
                    </div>

                    <div className={styles.item}>
                        <Link to={"/docs"}>Docs</Link>
                    </div>
                </div>
            </div>
        </nav>
    )
}

export default Navbar