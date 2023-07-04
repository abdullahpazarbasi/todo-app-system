import type AuthUseCase from "@/features/auth/usecases/AuthUseCase";
import ConcreteAuthUseCase from "@/features/auth/usecases/ConcreteAuthUseCase";
import type AuthService from "@/features/auth/services/AuthService";
import ConcreteAuthService from "@/features/auth/services/ConcreteAuthService";

import axios from "axios";
import type {Router} from "vue-router";

export function useAuthFeature(router: Router): AuthUseCase {
    const httpClient = axios.create({
        baseURL: import.meta.env.VITE_TODO_WBFF_BASE_URL + "/auth/token-claims",
    });
    const authService: AuthService = new ConcreteAuthService(
        httpClient,
    );

    return new ConcreteAuthUseCase(
        router,
        authService,
    );
}