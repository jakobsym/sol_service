import { HeliusWalletResponse, TokenAccount, Wallet } from "../../model/wallet";
require('dotenv').config()

// TODO: The returned token 'amount' needs to be converted into a readable form
// currently 1008892272740 actually looks like this -> 1_008_892.272740
// TODO: returned tokens are sometimes scam tokens, need a way to filter these out

const url = `https://mainnet.helius-rpc.com/?api-key=${process.env.HELIUS_API_KEY}`;

class HeliusRepository {
    private Wallet: Wallet

    getWalletContent = async(address: string): Promise<Wallet> => {
        var HeliusWalletResponse: HeliusWalletResponse
        try {
            HeliusWalletResponse = await HeliusWalletRequest(address)
        } catch (error) {
            console.error("error obtaining wallet contents: ", error)
        }
        if(!HeliusWalletResponse.result) {
            console.error(`empty wallet: ${address}`)
            return this.Wallet
        }

        let tokenAccount: TokenAccount[] = HeliusWalletResponse.result.token_accounts 
        this.Wallet.walletAddress = address
        tokenAccount.forEach(tokenData => {
            this.Wallet.tokens[tokenData.mint][tokenData.amount]
        })
        return this.Wallet
    }
}

export const HeliusWalletRequest = async (address: string): Promise<HeliusWalletResponse>  => {
    const fetch = (await import("node-fetch")).default
    var resData: HeliusWalletResponse

    try {
        const res = await fetch(url, {
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
        resData = await res.json() as HeliusWalletResponse
        
    } catch (error) {
        console.error(`error fetching wallet contents <${address}>: `, error)
        // will this return an empty HeliusWalletResponse?
        let res: HeliusWalletResponse
        return res
    }
    return resData
}
export default new HeliusRepository()