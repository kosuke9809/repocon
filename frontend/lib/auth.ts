import NextAuth, { Session } from "next-auth";
import { JWT } from "next-auth/jwt";
import GitHub from "next-auth/providers/github";

export interface CustomSession extends Session {
  accessToken?: string;
}

interface ExtendedToken extends JWT {
  accessToken?: string;
}

export const { handlers, auth, signIn, signOut } = NextAuth({
  providers: [GitHub],
  callbacks: {
    authorized({ request, auth }) {
      const { pathname } = request.nextUrl;
      return !!auth;
    },
    async jwt({ token, account, profile, user }) {
      const extendedToken: ExtendedToken = token as ExtendedToken;
      if (account) {
        extendedToken.accessToken = account.access_token;
      }
      return { ...extendedToken, ...user };
    },
    async session({ session, token, user }) {
      const extendedToken: ExtendedToken = token;
      const customSession: CustomSession = {
        ...session,
        accessToken: extendedToken.accessToken,
      };
      return customSession;
    },
  },
  session: {
    strategy: "jwt",
  },
  trustHost: true,
  // pages: {
  //   signIn: "/login",
  // },
});
