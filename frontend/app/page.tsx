import AuthButton from "@/components/auth-button";
import { CustomSession } from "@/lib/auth";
import { auth } from "@/lib/auth";


export default async function Home() {
  const session: CustomSession | null = await auth();
  return (
    <main className="flex min-h-screen flex-col items-center justify-between p-24">
      <div>
        <AuthButton/>
        {session ? (
          <div>
            <div>user: {session.user?.name || "Guest"}</div>
            <div>AccessToken: {session.accessToken}</div>
          </div>
          
          ) : <div>user: not signed in</div>}
      </div>
    </main>
  );
}
