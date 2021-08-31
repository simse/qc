import React from "react"
import { Link } from "gatsby"

import * as styles from "../styles/components/hero.module.scss"

const { detect } = require('detect-browser');

class Hero extends React.Component {

    constructor(props) {
        super(props)

        this.state = {
            os: "windows"
        }
    }

    componentDidMount() {
        const browser = detect();
        if (browser.os === "Mac OS") {
            this.setState({os: "mac"})
        } else if (browser.os === "Linux") {
            this.setState({os: "linux"})
        }
    }

    render () {
        return (
            <div className={styles.hero}>
                <div className={styles.inner}>
                    <div className={styles.innerBox}>
                        <div className={styles.text}>
                            <h1 className={styles.colored}>Convert Any File Format</h1>
                            <h1>Without the Hassle</h1>

                            <p>qc is a tool for converting between file formats. It supports tons of formats, using either native Go libraries or good ol' C libraries.</p>
                        
                            <Link to={"/install"}>Install qc now</Link>
                        </div>
                    </div>
                    
                    <div className={styles.innerBox}>
                        <div className={styles.demo}>
                            <div className={styles.osSelect}>
                                <button onClick={() => {this.setState({os: "windows"})}} className={this.state.os === "windows" ? styles.active : ""}>Windows</button>

                                <button onClick={() => {this.setState({os: "mac"})}} className={this.state.os === "mac" ? styles.active : ""}>macOS</button>

                                <button onClick={() => {this.setState({os: "linux"})}} className={this.state.os === "linux" ? styles.active : ""}>Linux</button>
                            </div>

                            {this.state.os === "windows" && <pre>
                                <div className={styles.codeSection}>
                                    <code><span className={styles.comment}>&gt; Install qc using PowerShell</span></code>
                                    <br />
                                    <code><span className={styles.prompt}>PS C:&gt;</span> install</code>
                                </div>
                                
                                <div className={styles.codeSection}>
                                    <code><span className={styles.comment}>&gt; Navigate to a folder with pictures</span></code>
                                    <br />
                                    <code><span className={styles.prompt}>PS C:&gt;</span> cd <span className={styles.comment}>%USER%</span>\Haiwaii_2019</code>
                                </div>

                                <div className={styles.codeSection}>
                                    <code><span className={styles.comment}>&gt; Convert all files to JPG</span></code>
                                    <br />
                                    <code><span className={styles.prompt}>PS C:&gt;</span> qc <span className={styles.arg}>jpg</span></code>
                                </div>
                            </pre>}

                            {this.state.os === "mac" && <pre>
                                <div className={styles.codeSection}>
                                    <code><span className={styles.comment}># Install qc using Homebrew</span></code>
                                    <br />
                                    <code><span className={styles.prompt}>$</span> brew install simse/tap/qc</code>
                                </div>
                                
                                <div className={styles.codeSection}>
                                    <code><span className={styles.comment}># Navigate to a folder with pictures</span></code>
                                    <br />
                                    <code><span className={styles.prompt}>$</span> cd <span className={styles.comment}>~</span>/Pictures/Haiwaii_2019</code>
                                </div>

                                <div className={styles.codeSection}>
                                    <code><span className={styles.comment}># Convert all files to JPG</span></code>
                                    <br />
                                    <code><span className={styles.prompt}>$</span> qc <span className={styles.arg}>jpg</span></code>
                                </div>
                            </pre>}

                            {this.state.os === "linux" && <pre>
                            <div className={styles.codeSection}>
                                    <code><span className={styles.comment}># Install qc using the command line</span></code>
                                    <br />
                                    <code><span className={styles.prompt}>$</span> brew install simse/tap/qc</code>
                                </div>
                                
                                <div className={styles.codeSection}>
                                    <code><span className={styles.comment}># Navigate to a folder with pictures</span></code>
                                    <br />
                                    <code><span className={styles.prompt}>$</span> cd <span className={styles.comment}>~</span>/Haiwaii_2019</code>
                                </div>

                                <div className={styles.codeSection}>
                                    <code><span className={styles.comment}># Convert all files to JPG</span></code>
                                    <br />
                                    <code><span className={styles.prompt}>$</span> qc <span className={styles.arg}>jpg</span></code>
                                </div>
                            </pre>}
                        </div>
                    </div>
                </div>
            </div>
        )
    }
}

export default Hero