import { HeliusWalletResponse, TokenAccount, Wallet } from "../../model/wallet";
import 'dotenv/config'
//import MemoryRepository from "../memory/memory";
require('dotenv').config()


class HeliusRepository {
    private Wallet: Wallet
    private url: string = `https://mainnet.helius-rpc.com/?api-key=${process.env.HELIUS_API_KEY}`;

    // TODO: Assign response values to HeliusWalletResponse:TokenAccount
    // then pass these values to a Wallet, tokens will contain <TokenAccount.mint:TokenAccount.amount>
    getWalletContent = async(address: string): Promise<Wallet> => {
        //var token_accounts: TokenAccount
        var HeliusWalletResponse: HeliusWalletResponse
        const fetch = (await import("node-fetch")).default
        const res = await fetch(this.url, {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
            },
            body: JSON.stringify({
                jsonrpc: "2.0",
                method: "getTokenAccounts",
                id: "helius-test",
                params: {
                    page: 1,
                    limit: 100,
                    "displayOptions": {
                        "showZeroBalance": false,
                    },
                    owner: `${address}`,
                },
            }),
        });

        const data = await res.json()
        console.log(JSON.stringify(data, null, 2))
    
        return this.Wallet
    }
}

export default new HeliusRepository()