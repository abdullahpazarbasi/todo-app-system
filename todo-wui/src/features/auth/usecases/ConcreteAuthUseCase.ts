import type AuthUseCase from "@/features/auth/usecases/AuthUseCase";
import type AuthService from "@/features/auth/services/AuthService";

import type {Router} from "vue-router";

export default class ConcreteAuthUseCase implements AuthUseCase {

    protected router: Router;
    protected authService: AuthService;

    constructor(router: Router, authService: AuthService) {
        this.router = router;
        this.authService = authService;
    }

    login = (username: string, password: string): void => {
        this.authService.claimToken(username, password).then((tokenClaim) => {
            localStorage.setItem("token", tokenClaim.token);
            this.router.back();
        });
    }

    logout = (): void => {
        localStorage.removeItem("token");
    }

}
