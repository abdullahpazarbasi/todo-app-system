import type TokenClaim from "@/features/auth/models/TokenClaim";

export default interface AuthService {
    claimToken(username: string, password: string): Promise<TokenClaim>;
}
