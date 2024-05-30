import walletController from "../controller/walletController";
import { Router } from 'express'

const router = Router()
router.get('/wallet', (req,res) => {
    walletController.getWalletHelius(req, res)
})


export default router