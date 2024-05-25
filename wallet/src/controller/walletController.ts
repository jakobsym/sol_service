import 'dotenv/config'
import HeliusRepository from '../repository/helius/helius'
import { Wallet } from '../model/wallet';
import { Request, Response } from 'express';
require('dotenv').config()




class WalletController {
    // calls helius api returning a wallet based on address
    // curl curl -X GET localhost:8081/wallet?id=B2MhxCCPipp85D6LFdqd94Q5X5e8gy3dsGmHDxFH5GSV
    async getWalletHelius(req: Request, res: Response): Promise<void> {
        try {
            const walletAddress: string | undefined = req.query.id as string
            if(!walletAddress) {
                res.status(400).json({error: "Invalid wallet address"})
                return
            }
            const wallet = HeliusRepository.getWalletContent(walletAddress)
            res.json(wallet)
        } catch(error){
            res.status(500).json({ error: "Unable to fetch wallet content" });
        }
    }

    // getWalletMemory()
}

export default new WalletController()