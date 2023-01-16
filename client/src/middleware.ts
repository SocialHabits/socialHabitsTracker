import { NextRequest, NextResponse } from 'next/server';

export async function middleware(req: NextRequest) {
  const cookie = req.cookies.get('jwt')?.value;
  const url = req.nextUrl.clone();

  if (req.nextUrl.pathname === '/' && !cookie) {
    return;
  }

  if (req.nextUrl.pathname === '/' && cookie) {
    console.log(req.nextUrl.pathname === '/');
    url.pathname = '/dashboard';
    return NextResponse.redirect(url);
  }

  if (req.nextUrl.pathname === '/auth/login' && cookie) {
    url.pathname = '/dashboard';
    return NextResponse.redirect(url);
  }

  if (req.nextUrl.pathname === '/auth/sign-up' && cookie) {
    url.pathname = '/dashboard';
    return NextResponse.redirect(url);
  }

  if (
    !cookie &&
    req.nextUrl.pathname !== '/auth/login' &&
    req.nextUrl.pathname !== '/auth/sign-up'
  ) {
    url.pathname = '/';
    return NextResponse.redirect(url);
  }

  return NextResponse.next();
}

export const config = {
  matcher: [
    '/',
    '/auth/login',
    '/auth/sign-up',
    '/dashboard',
    '/dashboard/:path*',
  ],
};
