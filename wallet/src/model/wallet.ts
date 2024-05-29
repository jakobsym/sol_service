export interface Wallet {
    walletAddress: string
    tokens?: Map<string, number> // maps cannot be serialized, this only used internally
}
export interface WalletResponse {
    walletAddress: string
    tokens?: {[address:string]: number}
}

export interface HeliusWalletResponse {
    jsonrpc: string
    result: {
        total?: number
        limit?: number
        page?: number
        token_accounts: TokenAccount[]
    }
}

export interface TokenAccount {
    address: string
    mint: string
    owner: string
    amount: number
    delegated_amount: number
    frozen: boolean
}
/*
{
    walletaddress: "B2MhxCCPipp85D6LFdqd94Q5X5e8gy3dsGmHDxFH5GSV",
    tokens: [
        '7GCihgDB8fe6KNjn2MYtkzZcRjQy3t9GHdC8uHYmW2hr',
        '7GCihgDB8fe6KNjn2MYtkzZcRjQy3t9GHdC8uHYmW2hr',
        '7GCihgDB8fe6KNjn2MYtkzZcRjQy3t9GHdC8uHYmW2hr'
    ]
}
*/
