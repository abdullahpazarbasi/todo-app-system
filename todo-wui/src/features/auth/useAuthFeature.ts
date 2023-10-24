import type AuthUseCase from "@/features/auth/usecases/AuthUseCase";
import ConcreteAuthUseCase from "@/features/auth/usecases/ConcreteAuthUseCase";
import type AuthService from "@/features/auth/services/AuthService";
import ConcreteAuthService from "@/features/auth/services/ConcreteAuthService";

import type {AxiosInstance} from "axios";
import type {Router} from "vue-router";

export function useAuthFeature(router: Router, httpClient: AxiosInstance): AuthUseCase {
    httpClient.interceptors.request.use(
        (config) => {
            const token = localStorage.getItem('token');
            if (token != null && token.length > 0) {
                config.headers.Authorization = `Bearer ${token}`;
            }

            return config;
        },
    );
    httpClient.interceptors.response.use(
        response => {
            return response;
        },
        error => {
            if (error.response.status === 401) {
                return router.push({name: "Login"});
            }

            return Promise.reject(error);
        }
    );
    const authService: AuthService = new ConcreteAuthService(
        httpClient,
    );

    return new ConcreteAuthUseCase(
        router,
        authService,
    );
}
