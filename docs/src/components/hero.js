import React, { useState } from "react"

import * as styles from "../styles/components/hero.module.scss"

const { detect } = require('detect-browser');

const Hero = () => {
    let guessedBrowser = "windows"

    const browser = detect();
    if (browser.os === "Mac OS") {
        guessedBrowser = "mac"
    } else if (browser.os === "Linux") {
        guessedBrowser = "linux"
    }

    const [scriptOS, setScriptOS] = useState(guessedBrowser)
    

    return (
        <div className={styles.hero}>
            <div className={styles.inner}>
                <div className={styles.innerBox}>
                    <div className={styles.text}>
                        <h1 className={styles.colored}>Convert Any File Format</h1>
                        <h1>Without the Hassle</h1>

                        <p>qc is a tool for converting between file formats. It supports tons of formats, using either native Go libraries or good ol' C libraries.</p>
                    
                        <button>Read the docs</button>
                    </div>
                </div>
                
                <div className={styles.innerBox}>
                    <div className={styles.demo}>
                        <div className={styles.osSelect}>
                            <button onClick={() => {setScriptOS("windows")}} className={scriptOS === "windows" ? styles.active : ""}>Windows</button>

                            <button onClick={() => {setScriptOS("mac")}} className={scriptOS === "mac" ? styles.active : ""}>macOS</button>

                            <button onClick={() => {setScriptOS("linux")}} className={scriptOS === "linux" ? styles.active : ""}>Linux</button>
                        </div>

                        {scriptOS === "windows" && <pre>
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

                        {scriptOS === "mac" && <pre>
                            <div className={styles.codeSection}>
                                <code><span className={styles.comment}># Install qc using the terminal</span></code>
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

                        {scriptOS === "linux" && <pre>
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

export default Hero