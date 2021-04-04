const express = require('express')

const app = express()

const fibonnaci = (nth) => {
    let x = 0
    let y = 1

    for (let i = 0; i <= nth; i++) {
        let tmp = y
        y += x
        x = tmp
    }

    return y
}

app.get("/hc", (_, res) => {
    console.log("[GET] on health check")
    return res.json({ status: "ok" })
})

app.get("/fibonacci/:nth", (req, res) => {
    const nth = req.params.nth
    const result = fibonnaci(nth)

    console.log(`[GET] on fibonacci/${nth}`)
    return res.json({ fibonacci: result })
})

app.listen(9090, () => console.log("NodeJS Listen on port 9090"))