import * as fs from 'fs';
import 'dotenv/config';
import express from 'express';
import router from './handler/walletHandler';
/*
- running le code
    npm run build
    node dist/index.js 
*/

const app = express()
const port = 8082

app.use('/', router)
app.listen(port, () => {
    console.log(`wallet service running at port 8082`)
})