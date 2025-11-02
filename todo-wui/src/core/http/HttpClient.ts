export interface HttpRequestConfig {
    url: string;
    method?: string;
    headers?: Record<string, string>;
    data?: BodyInit | null;
}

export interface HttpRequestOptions extends Omit<HttpRequestConfig, "url"> {}

export interface HttpClientConfig {
    baseURL?: string;
}

export interface HttpResponse<T = unknown> {
    data: T;
    status: number;
    statusText: string;
    headers: Headers;
    config: HttpRequestConfig;
}

export class HttpError extends Error {
    public readonly response?: HttpResponse<unknown>;
    public readonly config: HttpRequestConfig;
    public readonly cause?: unknown;

    constructor(message: string, config: HttpRequestConfig, response?: HttpResponse<unknown>, cause?: unknown) {
        super(message);
        this.name = "HttpError";
        this.config = config;
        this.response = response;
        this.cause = cause;
    }
}

type MaybePromise<T> = T | Promise<T>;

type RequestInterceptor = (config: HttpRequestConfig) => MaybePromise<HttpRequestConfig>;
type ResponseFulfilled = (response: HttpResponse<unknown>) => MaybePromise<HttpResponse<unknown>>;
type ResponseRejected = (error: HttpError) => MaybePromise<HttpResponse<unknown>>;

interface InterceptorPair {
    onFulfilled: ResponseFulfilled;
    onRejected?: ResponseRejected;
}

class RequestInterceptorManager {
    private readonly handlers: RequestInterceptor[] = [];

    use(onFulfilled: RequestInterceptor): number {
        this.handlers.push(onFulfilled);

        return this.handlers.length - 1;
    }

    clear(): void {
        this.handlers.length = 0;
    }

    async run(config: HttpRequestConfig): Promise<HttpRequestConfig> {
        let currentConfig = config;
        for (const handler of this.handlers) {
            currentConfig = await handler(currentConfig);
        }

        return currentConfig;
    }
}

class ResponseInterceptorManager {
    private readonly handlers: InterceptorPair[] = [];

    use(onFulfilled: ResponseFulfilled, onRejected?: ResponseRejected): number {
        this.handlers.push({onFulfilled, onRejected});

        return this.handlers.length - 1;
    }

    clear(): void {
        this.handlers.length = 0;
    }

    getHandlers(): InterceptorPair[] {
        return [...this.handlers];
    }
}

export default class HttpClient {
    readonly interceptors = {
        request: new RequestInterceptorManager(),
        response: new ResponseInterceptorManager(),
    } as const;

    private readonly baseURL?: string;

    constructor(config: HttpClientConfig = {}) {
        this.baseURL = config.baseURL;
    }

    async request<T = unknown>(config: HttpRequestConfig): Promise<HttpResponse<T>> {
        const preparedConfig: HttpRequestConfig = {
            method: config.method ?? "GET",
            url: this.combineURL(config.url),
            headers: {...config.headers},
            data: config.data ?? null,
        };

        preparedConfig.headers = preparedConfig.headers ?? {};

        const interceptedConfig = await this.interceptors.request.run(preparedConfig);

        const requestInit = this.createRequestInit(interceptedConfig);

        let fetchResponse: Response;
        try {
            fetchResponse = await fetch(interceptedConfig.url, requestInit);
        } catch (error) {
            throw new HttpError("Failed to execute request", interceptedConfig, undefined, error);
        }

        const httpResponse = await this.buildResponse<T>(fetchResponse, interceptedConfig);

        if (!fetchResponse.ok) {
            return this.handleResponseError(httpResponse);
        }

        return this.handleResponseSuccess(httpResponse);
    }

    get<T = unknown>(url: string, config: HttpRequestOptions = {}): Promise<HttpResponse<T>> {
        return this.request<T>({
            ...config,
            url,
            method: "GET",
        });
    }

    delete<T = unknown>(url: string, config: HttpRequestOptions = {}): Promise<HttpResponse<T>> {
        return this.request<T>({
            ...config,
            url,
            method: "DELETE",
        });
    }

    post<T = unknown>(url: string, data?: BodyInit | null, config: HttpRequestOptions = {}): Promise<HttpResponse<T>> {
        return this.request<T>({
            ...config,
            url,
            method: "POST",
            data: data ?? null,
        });
    }

    patch<T = unknown>(url: string, data?: BodyInit | null, config: HttpRequestOptions = {}): Promise<HttpResponse<T>> {
        return this.request<T>({
            ...config,
            url,
            method: "PATCH",
            data: data ?? null,
        });
    }

    private async handleResponseSuccess<T>(response: HttpResponse<T>): Promise<HttpResponse<T>> {
        let currentResponse: HttpResponse<unknown> = response;
        const handlers = this.interceptors.response.getHandlers().reverse();
        for (const {onFulfilled} of handlers) {
            currentResponse = await onFulfilled(currentResponse);
        }

        return currentResponse as HttpResponse<T>;
    }

    private async handleResponseError<T>(response: HttpResponse<T>): Promise<HttpResponse<T>> {
        let error: HttpError = new HttpError("Request failed with status code " + response.status, response.config, response);
        const handlers = this.interceptors.response.getHandlers().reverse();
        for (const {onRejected} of handlers) {
            if (onRejected == null) {
                continue;
            }
            try {
                const maybeResponse = await onRejected(error);
                if (maybeResponse != null) {
                    return maybeResponse as HttpResponse<T>;
                }
            } catch (interceptorError) {
                error = interceptorError instanceof HttpError ? interceptorError : new HttpError(
                    interceptorError instanceof Error ? interceptorError.message : String(interceptorError),
                    response.config,
                    response,
                    interceptorError,
                );
            }
        }

        throw error;
    }

    private combineURL(url: string): string {
        if (this.baseURL == null || /^https?:/i.test(url)) {
            return url;
        }

        return `${this.baseURL.replace(/\/$/, "")}/${url.replace(/^\//, "")}`;
    }

    private createRequestInit(config: HttpRequestConfig): RequestInit {
        const headers = new Headers();
        if (config.headers != null) {
            for (const [key, value] of Object.entries(config.headers)) {
                if (value != null) {
                    headers.set(key, value);
                }
            }
        }

        const method = config.method?.toUpperCase();
        let body: BodyInit | undefined;
        if (method !== "GET" && method !== "HEAD") {
            body = config.data ?? undefined;
        }

        return {
            method,
            headers,
            body,
        };
    }

    private async buildResponse<T>(response: Response, config: HttpRequestConfig): Promise<HttpResponse<T>> {
        let data: unknown = null;

        if (response.status !== 204) {
            const contentType = response.headers.get("content-type");
            if (contentType != null && contentType.includes("application/json")) {
                data = await response.json();
            } else if (contentType != null && (contentType.includes("text/") || contentType.includes("application/xml"))) {
                data = await response.text();
            } else {
                data = await response.arrayBuffer();
            }
        }

        return {
            data: data as T,
            status: response.status,
            statusText: response.statusText,
            headers: response.headers,
            config,
        };
    }
}
