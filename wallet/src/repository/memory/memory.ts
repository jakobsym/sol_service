// In-memory mock database
import { Wallet } from "../../model/wallet";

class MemoryRepository {
    // key: string ; val: Wallet
    private db: Map<string, Wallet> = new Map()
    constructor() {}

    // add a wallet
    add(wallet: Wallet): void {
        this.db.set(wallet.walletAddress, wallet)
    }

    // delete by address
    delete(address: string): void {
        this.db.delete(address)
    }

    // get wallet by wallet address
    getByAddress(address: string): Wallet {
        return this.db.get(address)
    }

    // get all wallets
    // return as an object
    getAllWallets(): { [address: string]: Wallet }  {
        const obj: {[address: string]: Wallet} = {}
        this.db.forEach((value, key) => {
            obj[key] = value
        })
        return obj
    }

    // get total # of wallets
    getTotalWallets(): number {
        return this.db.size
    }
}


export default new MemoryRepository()