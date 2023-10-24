import type AuthService from "@/features/auth/services/AuthService";

import type {Axios} from "axios";
import type TokenClaim from "@/features/auth/models/TokenClaim";

export default class ConcreteAuthService implements AuthService {

    protected client: Axios;

    constructor(client: Axios) {
        this.client = client;
    }

    claimToken = async (username: string, password: string): Promise<TokenClaim> => {
        const formData = new URLSearchParams();
        formData.append('username', username);
        formData.append('password', password);
        const response = await this.client?.post(
            '/auth/token-claims',
            formData,
            {
                headers: {
                    'Content-Type': 'application/x-www-form-urlencoded',
                },
            }
        );

        return response.data;
    }

}
