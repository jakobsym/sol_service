import 'dotenv/config'
import HeliusRepository from '../repository/helius/helius'
import { Wallet, WalletResponse } from '../model/wallet';
import { Request, Response } from 'express';
require('dotenv').config()




class WalletController {
    // calls helius api returning a wallet based on address
    // curl curl -X GET localhost:8081/wallet?id=B2MhxCCPipp85D6LFdqd94Q5X5e8gy3dsGmHDxFH5GSV
    private WalletResponse: WalletResponse
    public setWallet(address: string, wr: WalletResponse): WalletResponse  {
        return this.WalletResponse = {
            walletAddress: address,
            tokens: wr.tokens
        }
    }

    public getWalletHelius = async(req: Request, res: Response): Promise<WalletResponse> => {
        try {
            const walletAddress: string | undefined = req.query.id as string
            if(!walletAddress) {
                res.status(400).json({error: "Invalid wallet address"})
                return
            }
            let tmpWallet = await HeliusRepository.getWalletContent(walletAddress)
            this.WalletResponse = this.setWallet(walletAddress, tmpWallet)
            //console.log(`WalletController response: ${JSON.stringify( this.WalletResponse )}`)
            res.json( this.WalletResponse )
        } catch(error){
            res.status(500).json({ error: "Unable to fetch wallet content" });
        }
    }

    // getWalletMemory()
}

export default new WalletController()