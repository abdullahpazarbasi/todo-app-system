export default interface AuthUseCase {
    login(username: string, password: string): void
    logout(): void
}
