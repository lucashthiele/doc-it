export default function GitHubLoginButton() {
  return (
    <button className="flex items-center justify-center flex-row px-4 py-2 gap-2 bg-black rounded text-white cursor-pointer">
      <img className="w-8" src="/github-mark-white.svg" />
      Login with GitHub
    </button>
  );
}
