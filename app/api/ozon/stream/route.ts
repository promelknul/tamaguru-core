import { NextRequest } from 'next/server';
export const runtime = 'edge';
export async function GET(_: NextRequest) {
  const ozonRes = await fetch('http://95.163.222.133:9000/orders/stream', {
    headers: { Accept: 'text/event-stream' }
  });
  return new Response(ozonRes.body, {
    headers: {
      'Content-Type': 'text/event-stream',
      'Cache-Control': 'no-cache',
      Connection: 'keep-alive'
    }
  });
}
