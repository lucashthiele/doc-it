import GitHubLoginButton from "./components/GitHubLoginButton";

export default function LoginPage() {
  const githubClientId = import.meta.env.VITE_GH_CLIENT_ID;
  return (
    <main className="flex items-center justify-center min-h-screen">
      <div className="flex items-center justify-center w-full max-w-md p-8 space-y-4 bg-white rounded shadow-lg">
        <a href={"https://github.com/login/oauth/authorize?client_id=" + githubClientId}>
          <GitHubLoginButton />
        </a>
      </div>
    </main>
  );
}
